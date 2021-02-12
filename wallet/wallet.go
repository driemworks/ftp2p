package wallet

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"ftp2p/main/manifest"
	"io/ioutil"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const keystoreDirName = "keystore"
const AndrejAccount = "0x22ba1F80452E6220c7cc6ea2D1e3EEDDaC5F694A"

func GetKeystoreDirPath(dataDir string) string {
	return filepath.Join(dataDir, keystoreDirName)
}

func NewKeystoreAccount(dataDir, password string) (common.Address, error) {
	ks := keystore.NewKeyStore(GetKeystoreDirPath(dataDir), keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.NewAccount(password)
	if err != nil {
		return common.Address{}, err
	}

	return acc.Address, nil
}

func SignTxWithKeystoreAccount(tx manifest.Tx, acc common.Address, pwd, keystoreDir string) (manifest.SignedTx, error) {
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	ksAccount, err := ks.Find(accounts.Account{Address: acc})
	if err != nil {
		return manifest.SignedTx{}, err
	}
	ksAccountJson, err := ioutil.ReadFile(ksAccount.URL.Path)
	if err != nil {
		return manifest.SignedTx{}, err
	}

	key, err := keystore.DecryptKey(ksAccountJson, pwd)
	if err != nil {
		return manifest.SignedTx{}, err
	}

	signedTx, err := SignTx(tx, key.PrivateKey)
	if err != nil {
		return manifest.SignedTx{}, err
	}

	return signedTx, nil
}

func SignTx(tx manifest.Tx, privKey *ecdsa.PrivateKey) (manifest.SignedTx, error) {
	rawTx, err := tx.Encode()
	if err != nil {
		return manifest.SignedTx{}, err
	}

	sig, err := Sign(rawTx, privKey)
	if err != nil {
		return manifest.SignedTx{}, err
	}

	return manifest.NewSignedTx(tx, sig), nil
}

func Sign(msg []byte, privKey *ecdsa.PrivateKey) (sig []byte, err error) {
	msgHash := sha256.Sum256(msg)

	return crypto.Sign(msgHash[:], privKey)
}

func Verify(msg, sig []byte) (*ecdsa.PublicKey, error) {
	msgHash := sha256.Sum256(msg)

	recoveredPubKey, err := crypto.SigToPub(msgHash[:], sig)
	if err != nil {
		return nil, fmt.Errorf("unable to verify message signature. %s", err.Error())
	}

	return recoveredPubKey, nil
}
