package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, str string) {
	fmt.Fprintf(writer, "Hello "+str)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "http")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}
