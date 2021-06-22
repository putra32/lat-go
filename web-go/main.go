package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(rw http.ResponseWriter, r *http.Request) {
	var message = "wellcome"

	rw.Write([]byte(message))
}

func handlerHello(rw http.ResponseWriter, r *http.Request) {
	var message = "Hello World"

	rw.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"

	fmt.Printf("Server started at %s\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println(err.Error())
	}
}
