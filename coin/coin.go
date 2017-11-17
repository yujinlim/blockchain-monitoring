package coin

import "github.com/prometheus/client_golang/prometheus"

// Type coin type
type Type string

//type of coins
const (
	BitcoinType  Type = "bitcoin"
	DogecoinType Type = "dogecoin"
	LitecoinType Type = "litecoin"
	EthereumType Type = "ethereum"
)

// Coin an interface for all coin client instance
type Coin interface {
	MonitorCount(gauge prometheus.Gauge)
	MonitorStatus(gauge prometheus.Gauge)
	MonitorDifferences(gauge prometheus.Gauge)
	Ping() error
}
