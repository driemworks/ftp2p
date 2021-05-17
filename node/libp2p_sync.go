package node

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"strings"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	noise "github.com/libp2p/go-libp2p-noise"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
)

/*
	Build a libp2p host
*/
func makeHost(port int, insecure bool) (host.Host, error) {
	r := rand.Reader
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
		libp2p.Security(noise.ID, noise.New),
		libp2p.EnableNATService(),
	}
	if insecure {
		opts = append(opts, libp2p.NoSecurity)
	}
	host, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}
	hostAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ipfs/%s", host.ID().Pretty()))
	addr := host.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	fmt.Printf("I am %s\n", fullAddr)
	fmt.Println(host.ID().Pretty())
	return host, nil
}

/*
	Manually add a peer to the DHT
	If doRelay = true then opena connection with the peer
*/
func addPeers(ctx context.Context, n Node, kad *dht.IpfsDHT, peersArg string, doRelay bool) {
	if len(peersArg) == 0 {
		return
	}
	peerStrs := strings.Split(peersArg, ",")
	for i := 0; i < len(peerStrs); i++ {
		peerID, peerAddr := MakePeer(peerStrs[i])
		n.host.Peerstore().AddAddr(peerID, peerAddr, peerstore.PermanentAddrTTL)
		_, err := kad.RoutingTable().TryAddPeer(peerID, false, false)
		if err != nil {
			log.Fatalln(err)
		}
		// if the peer is already in the DHT, do not call it again
		if doRelay {
			peerinfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
			if err != nil {
				log.Fatalln(err)
			}
			err = n.host.Connect(ctx, *peerinfo)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func (n *Node) runLibp2pNode(ctx context.Context, port int, bootstrapPeer string, name string) error {
	host, err := makeHost(port, false)
	n.host = host
	if err != nil {
		return err
	}
	// 1) Start a DHT
	kademliaDHT, err := dht.New(ctx, host)
	if err != nil {
		log.Fatal(err)
	}
	defer kademliaDHT.Close()
	// add bootstrap nodes if provided
	addPeers(ctx, *n, kademliaDHT, bootstrapPeer, true)
	log.Printf("Listening on %v (Protocols: %v)", host.Addrs(), host.Mux().Protocols())
	var ps *pubsub.PubSub
	if bootstrapPeer == "" {
		host.SetStreamHandler(DiscoveryServiceTag, func(s network.Stream) {
			// TODO this is pretty bad...
			addPeers(ctx, *n, kademliaDHT, s.Conn().RemoteMultiaddr().String()+"/p2p/"+s.Conn().RemotePeer().String(), false)
		})
		// TODO leaving as localhost for now. Should this be configurable?
		bootstrapPeer = fmt.Sprintf("/ip4/%s/tcp/%d/p2p/%s", "127.0.0.1", port, host.ID().Pretty())
	}
	peerinfo, err := peer.AddrInfoFromP2pAddr(multiaddr.StringCast(bootstrapPeer))
	tracer, err := pubsub.NewRemoteTracer(ctx, host, *peerinfo)
	if err != nil {
		panic(err)
	}
	// create a pubsub service using the GossipSub router
	ps, err = pubsub.NewGossipSub(ctx, host, pubsub.WithEventTracer(tracer))
	if err != nil {
		log.Fatalln(err)
	}
	// pending_tx_channel, err := InitChannel(ctx, PENDING_TX_TOPIC, 128, ps, host.ID())
	pending_tx_cr, err := JoinPendingTxExchange(ctx, ps, host.ID())
	new_block_cr, err := JoinNewBlockExchange(ctx, ps, host.ID())
	for {
		select {
		case m := <-pending_tx_cr.PendingTransactions:
			// read new pending txs and apply them to state
			n.AddPendingTX(*m)
		case tx := <-n.newPendingTXs:
			// publish new pending txs
			pending_tx_cr.Publish(&tx)
		// case data := <-pending_tx_channel.data:
		// 	// read new pending txs and apply them to state
		// 	// need to convert map[string]interface{} to signed tx
		// 	signed_tx := state.NewSignedTx(
		// 		,
		// 		data["signature"].([]byte),
		// 	)

		// 	n.AddPendingTX(*tx)
		// case tx := <-n.newPendingTXs:
		// 	// publish new pending txs
		// 	pending_tx_channel.Publish(&tx)
		case b := <-new_block_cr.NewBlocks:
			// when there's a new block in the stream (not yours) => Add the block
			s, _, err := n.state.AddBlock(*b)
			if err != nil {
				if s != nil {
					n.state = s
				}
				return err
			}
			n.newSyncedBlocks <- *b
		case block := <-n.newMinedBlocks:
			// when you've added a new block, publish it to the stream
			new_block_cr.Publish(&block)
		}
	}
}
