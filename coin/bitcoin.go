package coin

import (
	"log"

	"github.com/btcsuite/btcrpcclient"
	"github.com/prometheus/client_golang/prometheus"
)

// Bitcoin Coin instance to be monitoring using btcd client
type Bitcoin struct {
	client  *btcrpcclient.Client
	compare func(int64) (float64, error)
}

// NewBitcoinCoin Create new bitcoin compatible instance
func NewBitcoinCoin(host string, username string, password string, coinType Type) (*Bitcoin, error) {
	connCfg := &btcrpcclient.ConnConfig{
		Host:         host,
		User:         username,
		Pass:         password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	// create new client connection using websocket
	client, err := btcrpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}

	coin := &Bitcoin{
		client: client,
	}

	if coinType == DogecoinType {
		coin.compare = compareDogecoinBlockCount
	}

	return coin, nil
}

// MonitorCount monitor coin current block count
func (coin *Bitcoin) MonitorCount(gauge prometheus.Gauge) {
	count, err := coin.client.GetBlockCount()
	if err != nil {
		log.Print(err)
	}

	// update block count
	gauge.Set(float64(count))
}

// MonitorStatus monitor node state
func (coin *Bitcoin) MonitorStatus(gauge prometheus.Gauge) {
	statusCode := 200

	err := coin.client.Ping()
	if err != nil {
		statusCode = 404
	}

	gauge.Set(float64(statusCode))
}

// Ping check bitcoin compatible connection status
func (coin *Bitcoin) Ping() error {
	return coin.client.Ping()
}

// MonitorDifferences monitor differences in block count
func (coin *Bitcoin) MonitorDifferences(gauge prometheus.Gauge) {
	count, err := coin.client.GetBlockCount()
	if err != nil {
		log.Panic(err)
	}

	diff, err := coin.compare(count)
	if err != nil {
		log.Panic(err)
	}

	gauge.Set(diff)
}
