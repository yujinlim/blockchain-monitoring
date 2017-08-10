package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	log.Println("starting...")
	// load env
	err := godotenv.Load()
	if err != nil {
		// ignore err
		log.Println(err)
	}

	port := os.Getenv("PORT")
	namespace := os.Getenv("COIN_NAMESPACE")
	if len(namespace) == 0 {
		namespace = "coin"
	}

	if len(port) == 0 {
		port = "8080"
	}
	addr := strings.Join([]string{":", port}, "")

	coin, err := NewCoin()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Print("creating...")

	blockCounter := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "block_total",
		Help:      "Total block count for current node",
	})

	stateSummary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Name:      "status",
		Help:      "status of coin node",
	}, []string{"status"})

	prometheus.MustRegister(blockCounter)
	prometheus.MustRegister(stateSummary)

	// spin go routine
	// monitor block counter
	go func() {
		for {
			coin.MonitorCount(blockCounter)
		}
	}()

	// monitor node state
	go func() {
		for {
			coin.MonitorStatus(stateSummary)
		}
	}()

	// setup routes with handler
	http.Handle("/metrics", prometheusHandler())
	http.Handle("/probe", probHandler(coin))

	log.Fatal(http.ListenAndServe(addr, nil))
}
