package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/locustbaby/httpProxy/config"
)

func startServer(c *config.Config) {
	endpoint := fmt.Sprintf("http://%s:%d", c.Backend.Endpoints[0].Host, c.Backend.Endpoints[0].Port)
	url, _ := url.Parse(endpoint)
	proxy := httputil.NewSingleHostReverseProxy(url)
	err := http.ListenAndServe(":8888", proxy)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}

}

func main() {
	var c config.Config
	config.Init(&c, "../testdata/yml_base.yaml")
	startServer(&c)
}
