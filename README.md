# FTP2P
FTP2P is a blockchain for sharing [CID's](https://docs.ipfs.io/concepts/content-addressing/) with peers across ipfs gateways.

## Getting Started

### Summary


If you're familiar with pinning services (such as [this one](https://pinata.cloud/)) then you can best think of FTP2P as an inverted pinning service. FTP2P does not have any direct integration or dependency on IPFS.

### Pre requisites
- install `go`
- install `ipfs` (recommended)

### Installation 
Install f2p2p using:
```
go get github.com/driemworks/ftp2p/cli/...
```
(Recommnded) Install ftp2p from the source by cloning this repo and run `go install ./cli/...` from the root directory `ftp2p/`.

## Usage
### CLI commands
`ftp2p [command] [options]`
- Available Commands:
  - `help`: Help about any command
  - `run`:  Run the ftp2p node
    -  options:
      - `--name`: (optional) The name of your node - Default: ?
      - `--datadir`: (optional) the directory where local data will be stored - Default: `.ftp2p`
      - `--ip`: (optional) the ip addreses of the ftp2p node - Default: `127.0.0.1`
      - `--port`: (optional) the port of the ftp2p node - Default: `8080`
      - `--miner`: (required) the public key to use (see: output of wallet command)
      - `--bootstrap-ip`: (optional) the ip address of the bootstrap node - Default: `127.0.0.1`
      - `--bootstrap-port`: (optional) the port of the bootstrap node - Default: `8080`
  - `wallet`: Access the node's wallet
    - `new-address` Generate a new address
        -  options:
            - `--datadir`: (required) the directory where local data will be stored (ex: blockchain transactions)

 example:
  ```
  # generate a new address
  ftp2p wallet new-address --datadir=./.ftp2p
  >  0x27084384033F90d96c3769e1b4fCE0E5ffff720B
  # start a node using the new address as the miner
  ftp2p run --datadir=./.ftp2p --name=Theo --miner=0x27084384033F90d96c3769e1b4fCE0E5ffff720B --port=8080 --bootstrap-ip=127.0.0.1 --bootstrap-port=8081
  ```

## API
Note: In order to use the API a node must be running. 

See the [API documentation](https://github.com/driemworks/ftp2p/blob/master/docs/api/api.md)

### Node/Sync API
#### Pending Documentation
- `POST /node/sync`
- `POST /node/status`
- `POST /node/peer`


### Development
If you'd like to contribute send me an email at tonyrriemer@gmail.com or message me on discord: driemworks#1849

#### Future Features
[A comprehensive list of planned features will be added here]
- integrate with gojsonq
- expose API via a CLI 
- /node/* endpoints to rpc
- complete encrypt/decrypt functionality (+ expose api/cli)
- enhance 'transactions' to represent generic state mutation request

### Testing
- example: $ go test ./node/ -test.v -test.run ^TestValidBlockHash$ 

## Acknowledgements
- This repository's basis is heavily influenced by this repo and the associated ebook https://github.com/web3coach/the-blockchain-bar