# Mercury Blockchain
[![GoReportCard example](https://goreportcard.com/badge/github.com/driemworks/mercury-blockchain)](https://goreportcard.com/report/github.com/driemworks/mercury-blockchain)

[![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](http://perso.crans.org/besson/LICENSE.html)

Mercury is an unfinished experiment without a clearly defined scope as to what 'done' means.
The current iteration of mercury is a blockchain based p2p-chat platform. It allows nodes to define topics within transactions in the blockchain. Peers are then able to subscribe to topics as referenced by their unique transaction hash using rpc endpoints. Though there is a wallet and a token, there is currently no value assigned nor means to exchange tokens.

With no specific direction in mind, intentions for Mercury is to expand to a decentralized state management platform (i.e. something along the lines of holochain)

## Getting Started
Mercury uses [go-libp2p](https://github.com/libp2p/go-libp2p) to enable p2p communication and exposes rpc endpoints to allow communication with the the node. Security was *not* considered whatsoever in the current implementation. 

### Pre requisites
- install `go`
- install [grpcurl](https://github.com/fullstorydev/grpcurl) (recommended)

### Installation 
- Install mercury from the source by cloning this repo and run `go install ./cli/...` from the root directory `mercury-blockchain/`.
- Install the latest version using:
```
go get github.com/driemworks/mercury-blockchain/cli/...
```
Note: if you're running linux you may need to set `export GO111MODULE=on;go get github.com/driemworks/mercury-blockchain/cli/...`

## Usage
This section explains how to get started using mercury. 

### CLI commands
`mercury [command] [options]`
- Available Commands:
  - `help`: Help about any command
  - `version`: Print the current version
  - `run`:  Run the mercury node
    -  options:
      - `--name`: (optional) The name of your node - Default: `""`
      - `--datadir`: (optional) the directory where local data will be stored - Default: `.mercury`
      - `--host`: (optional) the ip addreses of the mercury node - Default: `0.0.0.0`
      - `--port`: (optional) the port of the mercury node. The RPC server will be run - Default: `8080`
      - `--rpc-host`: (optional) the ip addreses of the rpc server - Default: `0.0.0.0`
      - `--rpc-port`: (optional) the port to run the rpc server on - Default: `9080`
      - `--address`: (required) the address to use (found in keystore generated by wallet new-address command, or provide your own keystore)
      - `--bootstrap`: (required) Multihash of the peer you want to use as a bootstrap. This will be in the form `/ip4/<peer-ip>/tcp/<peer-port>/p2p/<peer node hash>` - Defaut: `""`
  - `wallet`: Access the node's wallet
    - `new-address` Generate a new address
        -  options:
            - `--datadir`: (required) the directory where local data will be stored (ex: blockchain transactions)

 example:
  ```
  mkdir .mercury
  # generate a new address
  mercury wallet new-address --datadir=./.mercury
  >  0x27084384033F90d96c3769e1b4fCE0E5ffff720B
  # start a node using the new address as the miner
  mercury run --datadir=./.mercury --port=8081 --rpc-port=9081 --miner=0x27084384033F90d96c3769e1b4fCE0E5ffff720B --bootstrap="/ip4/172.31.78.60/tcp/8080/p2p/QmWPgXq1ZXAMkdDMSaJok9VQsBVn69bk71y3yWYefd7nSr"
  ```

### Connect to test network
A bootstrap node is available at `/ip4/3.224.116.20/tcp/8080/p2p/QmVZMMmtvYLyUxeJjGP7LRZqEa957Z3DHZvJ1pkhDXTpXj`

example: `mercury run --name=<your-name> --datadir=.mercury/ --miner=<your-address> --port=<your port> --bootstrap=/ip4/172.31.78.60/tcp/8080/p2p/QmWPgXq1ZXAMkdDMSaJok9VQsBVn69bk71y3yWYefd7nSr`

Subscribe to the tx hash `197e33d7b4b7c987c3739689978a4a88745e3ef095b3df7878774d10b09b7e7c` and publish a message to say hello!

### RPC
Mercury uses gRPC to let you communicate directly with a node.
Authentication and Security is pending.

#### GetNodeStatus -> Not working?
Query the node for a status report

`rpc GetNodeStatus(NodeInfoRequest) returns (NodeInfoResponse) {}`

example with grpcurl:
```
grpcurl -plaintext 127.0.0.1:9081 proto.NodeService/GetNodeStatus
> {
>   "address": "0xEA3d0650a05d8F94DFFEd9514594BE2532Bec001",
>   "balance": 8,
>   "channels": [
>     "test",
>     "test"
>   ]
> }

```
#### AddTransaction
The main functionality (to be extended...): Create a new pending transaction that, once mined, will allow us to send generic tx payloads across nodes. Security has not been considered whatsoever with the current implementation.

In the current implementation this is synonymous with defining a new topic.
`rpc AddTransaction(AddPendingPublishCIDTransactionRequest) returns (AddPendingPublishCIDTransactionResponse) {}`

Example 
```
grpcurl -plaintext -d @ 127.0.0.1:9081 proto.NodeService/AddTransaction <<EOM
{
    "label": "hello",
    "password": "test"
}
EOM
```

### ListBlocks
TODOS:
1) this needs to be updated so we can actually stream blocks instead of just list them
List blocks from some given block hash, fromBlock. If empty, it is assumed to be from the genesis block.
2) BlockHeader is empty in response


In the current implementation this is synonymous with listing all defined topics and topic creators.
`rpc ListBlocks(ListBlocksRequest) returns (stream BlockResponse) {}`

Example 
```
./grpcurl -plaintext -d @ 127.0.0.1:9081 proto.NodeService/ListBlocks <<EOM
{
    "fromBlock": ""
}
EOM
{
  "blockHeader": {

  },
  "txs": [
    {
      "author": "0x1176b917AD034B2886c8B4F692121BFc1A8EA659",
      "topic": "welcome",
      "hash": "197e33d7b4b7c987c3739689978a4a88745e3ef095b3df7878774d10b09b7e7c"
    }
  ]
}
...(stream continues)
```

### Subscribe
Allows a node to subscribe to a topic, identified by the transaction hash within which it was defined. Any messages published to the topic are streamed.
`rpc Subscribe(JoinChannelRequest) returns (stream ChannelData) {}`

Example:
```
grpcurl -plaintext -d @ 127.0.0.1:9081 proto.NodeService/Subscribe <<EOM
{
    "txHash": "197e33d7b4b7c987c3739689978a4a88745e3ef095b3df7878774d10b09b7e7c"
}
EOM
{
  "data": "hello world",
  "from": "QmXiDmUAVczWgqG8dHyu3P7pciKs2TNy7mr9Ffd1ej7ysz",
  "topic": "197e33d7b4b7c987c3739689978a4a88745e3ef095b3df7878774d10b09b7e7c"
}
(stream continues)...
```

### Publish
Publish to a pubsub topic which is defined within a transaction.
`rpc Publish(PublishRequest) returns (PublishResponse) {}`

```
grpcurl -plaintext -d @ 127.0.0.1:9081 proto.NodeService/Publish <<EOM
{
  "txHash": "197e33d7b4b7c987c3739689978a4a88745e3ef095b3df7878774d10b09b7e7c",
	"message": "hello world"
}
EOM
```

## Development

#### proto
To update node proto run:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/node.proto
```


### Issues
- test coverage is nil

### Testing
- example: $ go test ./node/ -test.v -test.run ^TestValidBlockHash$ 

## Contributing
For now, there are no contributing guidelines. 
If you'd like to contribute send me an email at tonyrriemer@gmail.com or message me on discord: driemworks#1849 and let's chat, or just open an issue or pull request.
