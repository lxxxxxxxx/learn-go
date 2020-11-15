package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello go")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// Handler echoes r.RUL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Paht = %q\n", req.URL.Path)
}

// Handler echoes r.RUL.Handler
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
