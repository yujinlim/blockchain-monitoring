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

// Monitor monitor coin instance
func (coin *Coin) Monitor(gauge prometheus.Gauge) {
	count, err := coin.client.GetBlockCount()
	if err != nil {
		log.Print(err)
	}

	// update block count
	gauge.Set(float64(count))
	log.Print("monitoring...")
	time.Sleep(10 * time.Minute)
}
