package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yujinlim/blockchain-monitoring/coin"
)

// parse cli
func main() {
	var client coin.Coin
	var err error

	log.Println("starting...")
	// load env
	err = godotenv.Load()
	if err != nil {
		// ignore err
		log.Println(err)
	}

	coinType := os.Getenv("COIN_TYPE")
	port := os.Getenv("PORT")
	namespace := os.Getenv("COIN_NAMESPACE")
	host := os.Getenv("COIN_HOST")
	username := os.Getenv("COIN_USER")
	password := os.Getenv("COIN_PASSWORD")
	network := os.Getenv("COIN_NETWORK")

	if len(namespace) == 0 {
		namespace = "coin"
	}

	if len(port) == 0 {
		port = "8080"
	}
	addr := strings.Join([]string{":", port}, "")

	if coinType == "ethereum" {
		client, err = coin.NewEthCoin(host, coin.NetworkType(network))
	} else {
		client, err = coin.NewBitcoinCoin(host, username, password, coin.Type(coinType), coin.NetworkType(network))
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Print("creating...")
	start(client, namespace, addr)
}

// start setup prometheus and http
func start(coin coin.Coin, namespace string, addr string) {
	blockCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "block_total",
		Help:      "Total block count for current node",
	})

	statusCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "status",
		Help:      "Status of coin node",
	})

	diffCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "block_differences",
		Help:      "Block differences between current node and other api service",
	})

	prometheus.MustRegister(blockCounter, statusCounter, diffCounter)

	// spin go routine
	// monitor block counter
	go func() {
		for {
			coin.MonitorCount(blockCounter)
			time.Sleep(30 * time.Second)
		}
	}()

	// monitor node state
	go func() {
		for {
			coin.MonitorStatus(statusCounter)
			time.Sleep(30 * time.Second)
		}
	}()

	// monitor block differences
	go func() {
		for {
			coin.MonitorDifferences(diffCounter)
			time.Sleep(30 * time.Second)
		}
	}()

	// setup routes with handler
	http.Handle("/metrics", prometheusHandler())
	http.Handle("/probe", probHandler(coin))

	log.Fatal(http.ListenAndServe(addr, nil))
}
