package coin

import (
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

// EthCoin eth coin client instance
type EthCoin struct {
	client *rpc.Client
	// compare func(int64) (float64, error)
}

// NewEthCoin get new eth coin client
func NewEthCoin(url string, network NetworkType) (*EthCoin, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		log.Panic(err)
	}

	return &EthCoin{
		client: client,
		// compare: NewCompare(EthereumType, network),
	}, nil
}

func (coin *EthCoin) getBlockCount() (int64, error) {
	var result string

	err := coin.client.Call(&result, "eth_blockNumber")
	if err != nil {
		return 0, err
	}

	blockCount, err := strconv.ParseInt(result, 0, 64)
	if err != nil {
		return 0, err
	}

	return blockCount, nil
}

// MonitorCount monitor ethereum client current block count
func (coin *EthCoin) MonitorCount(gauge prometheus.Gauge) {
	blockCount, err := coin.getBlockCount()
	if err != nil {
		log.Panic(err)
	}

	gauge.Set(float64(blockCount))
}

// MonitorStatus monitor node status
func (coin *EthCoin) MonitorStatus(gauge prometheus.Gauge) {
	var result bool
	statusCode := 200

	if err := coin.client.Call(&result, "net_listening"); err != nil {
		log.Panic(err)
	}

	if !result {
		statusCode = 404
	}

	gauge.Set(float64(statusCode))
}

// Ping check monitoring status of eth node
func (coin *EthCoin) Ping() error {
	return nil
}

// MonitorDifferences monitor differences between current node and other api service
func (coin *EthCoin) MonitorDifferences(gauge prometheus.Gauge) {
	blockCount, err := coin.getBlockCount()
	if err != nil {
		log.Panic(err)
	}

	diff, err := compareEthBlockCount(blockCount)
	if err != nil {
		log.Panic(err)
	}

	gauge.Set(diff)
}
