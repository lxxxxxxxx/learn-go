package main

import (
	"fmt"
)

//Hello :return Hello World.
func Hello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(Hello("World"))
}
