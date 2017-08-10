# Blockchain Monitoring
> prometheus exporter to monitor bitcoin compatible blockchain node

## Purpose
This is an API that provides metrics and probe endpoints for prometheus monitoring. It is use to monitor:
- node state (ping/pong)
- current node block count

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
