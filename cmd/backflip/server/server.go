package server

import (
	"log"
	"net/http"

	"github.com/DaveVED/backflip/cmd/backflip/proxy"
)

func ServeHTTP(targetUrl string) {
	log.Printf("Server running on :8080")

	http.HandleFunc("/", proxy.ProxyHandler(targetUrl))

	http.ListenAndServe(":8080", nil)
}
