package node

import (
	"context"
	com "ftp2p/common"
	"ftp2p/state"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func TestNode_Run(t *testing.T) {
	datadir := getTestDataDirPath()
	err := state.RemoveDir(datadir)
	if err != nil {
		t.Fatal(err)
	}

	n := NewNode("testAlias", datadir, "127.0.0.1", 8085, state.NewAddress("test"), "", com.NewPeerNode(
		"", "127.0.0.1", 8080, false, common.Address{}, "", true,
	))

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	err = n.Run(ctx)
	if err != nil {
		t.Fatal("")
	}
}

func TestNode_Mining(t *testing.T) {
	// Remove the test directory if it already exists
	datadir := getTestDataDirPath()
	err := state.RemoveDir(datadir)
	if err != nil {
		t.Fatal(err)
	}

	// Required for AddPendingTX() to describe
	// from what node the TX came from (local node in this case)
	nInfo := com.NewPeerNode(
		"test",
		"127.0.0.1",
		8085,
		false,
		state.NewAddress("0x9F0d31dFE801cc74ED9e50F06aDC7B168FF2F35b"),
		"",
		true,
	)

	// Construct a new Node instance and configure
	// Andrej as a miner
	n := NewNode("testAlias", datadir, nInfo.IP, nInfo.Port, state.NewAddress("test"), "", nInfo)

	// Allow the mining to run for 30 mins, in the worst case
	ctx, closeNode := context.WithTimeout(
		context.Background(),
		time.Minute*30,
	)

	// Schedule a new TX in 3 seconds from now, in a separate thread
	// because the n.Run() few lines below is a blocking call
	go func() {
		time.Sleep(time.Second * miningIntervalSeconds / 3)
		tx := state.SignedTx{state.NewTx(state.NewAddress("tony"), state.NewAddress("tonay"),
			state.TransactionPayload{state.NewCID("QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH", "ipfs.io")},
			10, 0, state.TX_TYPE_001), []byte{}}

		_ = n.AddPendingTX(tx)
	}()

	// Schedule a new TX in 12 seconds from now simulating
	// that it came in - while the first TX is being mined
	go func() {
		time.Sleep(time.Second*miningIntervalSeconds + 2)
		tx := state.SignedTx{state.NewTx(state.NewAddress("tony"), state.NewAddress("theo"),
			state.TransactionPayload{state.NewCID("QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH", "ipfs.io")}, 10, 0, state.TX_TYPE_001), []byte{}}

		_ = n.AddPendingTX(tx)
	}()

	go func() {
		// Periodically check if we mined the 2 blocks
		ticker := time.NewTicker(10 * time.Second)

		for {
			select {
			case <-ticker.C:
				if n.state.LatestBlock().Header.Number == 1 {
					closeNode()
					return
				}
			}
		}
	}()

	// Run the node, mining and everything in a blocking call (hence the go-routines before)
	_ = n.Run(ctx)

	if n.state.LatestBlock().Header.Number != 1 {
		t.Fatal("2 pending TX not mined into 2 under 30m")
	}
}

// func TestNode_MiningStopsOnNewSyncedBlock(t *testing.T) {
// 	// Remove the test directory if it already exists
// 	datadir := getTestDataDirPath()
// 	err := fs.RemoveDir(datadir)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Required for AddPendingTX() to describe
// 	// from what node the TX came from (local node in this case)
// 	nInfo := NewPeerNode(
// 		"127.0.0.1",
// 		8085,
// 		false,
// 		database.NewAccount(""),
// 		true,
// 	)

// 	andrejAcc := database.NewAccount("andrej")
// 	babayagaAcc := database.NewAccount("babayaga")

// 	n := NewNode("test", datadir, nInfo.IP, nInfo.Port, babayagaAcc, nInfo)

// 	// Allow the test to run for 30 mins, in the worst case
// 	ctx, closeNode := context.WithTimeout(context.Background(), time.Minute*30)

// 	tx1 := database.NewTx("andrej", "babayaga", 1, "")
// 	tx2 := database.NewTx("andrej", "babayaga", 2, "")
// 	tx2Hash, _ := tx2.Hash()

// 	// Pre-mine a valid block without running the `n.Run()`
// 	// with Andrej as a miner who will receive the block reward,
// 	// to simulate the block came on the fly from another peer
// 	validPreMinedPb := NewPendingBlock(database.Hash{}, 0, andrejAcc, []database.Tx{tx1})
// 	validSyncedBlock, err := Mine(ctx, validPreMinedPb)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Add 2 new TXs into the BabaYaga's node
// 	go func() {
// 		time.Sleep(time.Second * (miningIntervalSeconds - 2))

// 		err := n.AddPendingTX(tx1, nInfo)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		err = n.AddPendingTX(tx2, nInfo)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 	}()

// 	// Once the BabaYaga is mining the block, simulate that
// 	// Andrej mined the block with TX1 in it faster
// 	go func() {
// 		time.Sleep(time.Second * (miningIntervalSeconds + 2))
// 		if !n.isMining {
// 			t.Fatal("should be mining")
// 		}

// 		_, err := n.state.AddBlock(validSyncedBlock)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		// Mock the Andrej's block came from a network
// 		n.newSyncedBlocks <- validSyncedBlock

// 		time.Sleep(time.Second * 2)
// 		if n.isMining {
// 			t.Fatal("synced block should have canceled mining")
// 		}

// 		// Mined TX1 by Andrej should be removed from the Mempool
// 		_, onlyTX2IsPending := n.pendingTXs[tx2Hash.Hex()]

// 		if len(n.pendingTXs) != 1 && !onlyTX2IsPending {
// 			t.Fatal("synced block should have canceled mining of already mined TX")
// 		}

// 		time.Sleep(time.Second * (miningIntervalSeconds + 2))
// 		if !n.isMining {
// 			t.Fatal("should be mining again the 1 TX not included in synced block")
// 		}
// 	}()

// 	go func() {
// 		// Regularly check whenever both TXs are now mined
// 		ticker := time.NewTicker(time.Second * 10)

// 		for {
// 			select {
// 			case <-ticker.C:
// 				if n.state.LatestBlock().Header.Number == 1 {
// 					closeNode()
// 					return
// 				}
// 			}
// 		}
// 	}()

// 	go func() {
// 		time.Sleep(time.Second * 2)

// 		// Take a snapshot of the DB balances
// 		// before the mining is finished and the 2 blocks
// 		// are created.
// 		startingAndrejBalance := n.state.Manifest[andrejAcc]
// 		startingBabaYagaBalance := n.state.Manifest[babayagaAcc]

// 		// Wait until the 30 mins timeout is reached or
// 		// the 2 blocks got already mined and the closeNode() was triggered
// 		<-ctx.Done()

// 		endAndrejBalance := n.state.Manifest[andrejAcc]
// 		endBabaYagaBalance := n.state.Manifest[babayagaAcc]

// 		// In TX1 Andrej transferred 1 TBB token to BabaYaga
// 		// In TX2 Andrej transferred 2 TBB tokens to BabaYaga
// 		expectedEndAndrejBalance := startingAndrejBalance - tx1.Value - tx2.Value + database.BlockReward
// 		expectedEndBabaYagaBalance := startingBabaYagaBalance + tx1.Value + tx2.Value + database.BlockReward

// 		if endAndrejBalance != expectedEndAndrejBalance {
// 			t.Fatalf("Andrej expected end balance is %d not %d", expectedEndAndrejBalance, endAndrejBalance)
// 		}

// 		if endBabaYagaBalance != expectedEndBabaYagaBalance {
// 			t.Fatalf("BabaYaga expected end balance is %d not %d", expectedEndBabaYagaBalance, endBabaYagaBalance)
// 		}

// 		t.Logf("Starting Andrej balance: %d", startingAndrejBalance)
// 		t.Logf("Starting BabaYaga balance: %d", startingBabaYagaBalance)
// 		t.Logf("Ending Andrej balance: %d", endAndrejBalance)
// 		t.Logf("Ending BabaYaga balance: %d", endBabaYagaBalance)
// 	}()

// 	_ = n.Run(ctx)

// 	if n.state.LatestBlock().Header.Number != 1 {
// 		t.Fatal("was suppose to mine 2 pending TX into 2 valid blocks under 30m")
// 	}

// 	if len(n.pendingTXs) != 0 {
// 		t.Fatal("no pending TXs should be left to mine")
// 	}
// }

func getTestDataDirPath() string {
	return filepath.Join(os.TempDir(), ".ftp2p2_test")
}
