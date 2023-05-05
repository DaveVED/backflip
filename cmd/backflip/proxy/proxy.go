package proxy

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func proxy(targetUrl string, w http.ResponseWriter, r *http.Request) {
	target, err := url.Parse(targetUrl)
	if err != nil {
		log.Printf("Error while parsing target URL: %v", err)
	}

	// Set route criteria.
	r.Host = target.Host
	r.URL.Scheme = target.Scheme
	r.URL.Host = target.Host

	// Transport to target.
	transport := http.DefaultTransport
	response, err := transport.RoundTrip(r)
	if err != nil {
		log.Printf("Error while sending request to target server: %v", err)
		http.Error(w, "Error while sending request to target server", http.StatusBadGateway)
		return
	}

	// Copy response headers
	for key, values := range response.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set response status code
	w.WriteHeader(response.StatusCode)

	// Copy response body
	if response.Body != nil {
		defer response.Body.Close()
		_, err := io.Copy(w, response.Body)
		if err != nil {
			log.Printf("Error while copying response body: %v", err)
		}
	}
}

func ProxyHandler(targetUrl string) http.HandlerFunc {
	// Entry point to proxy.
	return func(w http.ResponseWriter, r *http.Request) {
		proxy(targetUrl, w, r)
	}
}
