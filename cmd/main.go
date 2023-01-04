package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func startServer() {
	url, _ := url.Parse("http://127.0.0.1:8000")
	proxy := httputil.NewSingleHostReverseProxy(url)
	err := http.ListenAndServe(":8888", proxy)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}

}

func main() {
	startServer()
}
