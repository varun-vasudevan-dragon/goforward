package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Get the port to listen on
func getListenAddress() string {
	port := getEnv("PORT", "1338")
	return ":" + port
}

func addCORSHeaders(req *http.Request, res http.ResponseWriter) http.ResponseWriter {
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	res.Header().Add("Access-Control-Allow-Headers", "*")
	res.Header().Add("Access-Control-Max-Age", "1728000")
	return res
}

// Get the url for a given proxy condition
func getProxyURL(req *http.Request) string {
	pathComponents := strings.Split(req.URL.Path, "/")
	port := pathComponents[1]
	path := strings.Join(pathComponents[2:], "/")
	query := req.URL.RawQuery
	destination := "http://localhost:" + port + "/" + path
	if query != "" {
		destination = destination + "?" + query
	}
	log.Printf("Redirecting to %s", destination)

	return destination
}

// Log the env variables required for a reverse proxy
func logSetup() {
	log.Printf("Server will run on: %s\n", getListenAddress())
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	responseWithCORSHeaders := addCORSHeaders(req, res)
	if req.Method == "OPTIONS" {
		responseWithCORSHeaders.WriteHeader(204)
	} else {
		req.Header.Add("X-Forwarded-Host", req.Host)
		returnedURL := getProxyURL(req)
		usableURL, _ := url.Parse(returnedURL)
		req.URL.Scheme = usableURL.Scheme
		req.URL.Path = ""
		req.Header.Add("X-Origin-Host", req.Host)
		proxy := httputil.NewSingleHostReverseProxy(usableURL)
		proxy.ServeHTTP(responseWithCORSHeaders, req)
	}

}

/*
	Entry
*/

func main() {
	// Log setup values
	logSetup()

	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
