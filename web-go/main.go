package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
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

	// Static router
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	// Tempalte router
	http.HandleFunc("/view", func(rw http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Batman",
		}

		err = tmpl.Execute(rw, data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	var address = "localhost:9000"

	fmt.Printf("Server started at %s\n", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println(err.Error())
	}
}
