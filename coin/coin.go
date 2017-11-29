package coin

import "github.com/prometheus/client_golang/prometheus"

// Type coin type
type Type string

// NetworkType of network
type NetworkType string

//type of coins
const (
	BitcoinType  Type = "bitcoin"
	DashType     Type = "dash"
	DogecoinType Type = "dogecoin"
	LitecoinType Type = "litecoin"
	EthereumType Type = "ethereum"
)

//type of network, ether main or testnet
const (
	Mainnet NetworkType = "main"
	Testnet NetworkType = "testnet"
)

// Coin an interface for all coin client instance
type Coin interface {
	MonitorCount(gauge prometheus.Gauge)
	MonitorStatus(gauge prometheus.Gauge)
	MonitorDifferences(gauge prometheus.Gauge)
	Ping() error
}
