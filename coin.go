package main

import (
	"log"
	"os"
	"time"

	"github.com/btcsuite/btcrpcclient"
	"github.com/prometheus/client_golang/prometheus"
)

// Coin Coin instance to be monitoring using btcd client
type Coin struct {
	client *btcrpcclient.Client
}

// NewCoin Create new coin instance
func NewCoin() (*Coin, error) {
	connCfg := &btcrpcclient.ConnConfig{
		Host:         os.Getenv("COIN_HOST"),
		User:         os.Getenv("COIN_USER"),
		Pass:         os.Getenv("COIN_PASSWORD"),
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	// create new client connection using websocket
	client, err := btcrpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}

	coin := &Coin{
		client: client,
	}

	return coin, nil
}

// MonitorCount monitor coin current block count
func (coin *Coin) MonitorCount(gauge prometheus.Gauge) {
	count, err := coin.client.GetBlockCount()
	if err != nil {
		log.Print(err)
	}

	// update block count
	gauge.Set(float64(count))
	log.Print("monitoring...")
	time.Sleep(10 * time.Minute)
}

// MonitorStatus monitor node state
func (coin *Coin) MonitorStatus(summary *prometheus.SummaryVec) {
	labels := prometheus.Labels{"status": "success"}

	err := coin.client.Ping()
	if err != nil {
		labels = prometheus.Labels{"status": "fail"}
	}

	summary.With(labels)
	time.Sleep(1 * time.Minute)
}
