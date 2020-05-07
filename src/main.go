package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: corsaro <target> [port]")
		fmt.Println("\ttarget: The target host. (Ex: http://www.example.com:8080)")
		fmt.Println("\tport:   The local port to listen (Default: 1337)")
		return
	}
	localPort := "1337"
	if len(os.Args) == 3 {
		localPort = os.Args[2]
	}
	target, err := url.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(fmt.Sprintf(":%s", localPort), nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.Header().Set("forwarded-by", "corasaro")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			return
		}
		p.ServeHTTP(w, r)
	}
}
