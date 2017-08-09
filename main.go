package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("starting...")
	// load env
	err := godotenv.Load()
	if err != nil {
		// ignore err
		log.Println(err)
	}

	coin, err := NewCoin()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Print("creating...")

	namespace := os.Getenv("COIN_NAMESPACE")
	log.Print(namespace)
	if len(namespace) == 0 {
		namespace = "coin"
	}

	blockCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "block_total",
		Help:      "Total block count for current node",
	})

	prometheus.MustRegister(blockCounter)

	// spin go routine
	go func() {
		for {
			coin.Monitor(blockCounter)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
