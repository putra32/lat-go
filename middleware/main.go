package main

import (
	"fmt"
	"net/http"
)

// Fungsi Log yang berguna sebagai middleware
func log(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Ini dari middleware Log...\n")
		fmt.Println(r.URL.Path)
		next.ServeHTTP(rw, r)
	})
}

// Fungsi GetMahasiswa untuk menampilkan text string di browser
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini dari function GetMahasiswa()"))
}

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":5000",
	}
	// Routing
	http.Handle("/", log(http.HandlerFunc(GetMahasiswa)))

	// Jalankan server
	server.ListenAndServe()
}
