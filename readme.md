# Blockchain Monitoring
> prometheus exporter to monitor bitcoin compatible blockchain node

## Purpose
This is an API that provides metrics and probe endpoints for prometheus monitoring. It is use to monitor:
- node state (ping/pong)
- current node block count
- get difference of highest blockcount and current blockcount

## Supported Blockchain
- bitcoin (main, testnet)
- dash (main, testnet)
- dogecoin (main, testnet)
- ethereum (main, testnet)
- litecoin (main, testnet)

## Supported Network
- main
- testnet

## Environmental variables
```bash
# default service port
PORT=8080
# type of blockchain, please refer to supported blockchain for variable naming
COIN_TYPE=ethereum
# node ip and port, if ethereum, please include protocol
COIN_HOST=http://<ip>:8545
# type of network, main/testnet
COIN_NETWORK=main
# for bitcoin compatible rpc
COIN_USER=username
COIN_PASSWORD=password
# namespace for prometheus exporter
COIN_NAMESPACE=ethereum
```

## Requirements
- Go 1.8 and above

## Usage
### Via git
```bash
git clone git@github.com:yujinlim/blockchain-monitoring.git

cd blockchain-monitoring

go build

# browse to localhost:8080
./blockchain-monitoring
```

### via docker
```bash
```
