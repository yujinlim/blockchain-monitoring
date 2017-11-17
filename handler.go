package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yujinlim/prometheus-coin-monitoring/coin"
)

// prometheusHandler wrapper around prometheus http
func prometheusHandler() http.Handler {
	return promhttp.Handler()
}

// probHandler capture current coin instance and does ping and pong with coin node
func probHandler(coin coin.Coin) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := coin.Ping()
		if err != nil {
			w.WriteHeader(500)
		}

		w.WriteHeader(200)
	})
}
