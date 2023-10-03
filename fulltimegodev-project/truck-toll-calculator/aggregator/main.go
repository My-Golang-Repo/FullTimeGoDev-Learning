package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"net/http"
)

func main() {

	listenAddr := flag.String("listenaddr", ":3000", "the listen address of HTTP Server")
	flag.Parse()
	store := NewMemoryStore()

	var (
		svc = NewInvoiceAggregator(store)
	)
	makeHTTPTransport(*listenAddr, svc)
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("HTTP Transport running on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
