package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requeste
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q \n", req.URL.Path)
	case "/hello":
		for v, k := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
		}
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
