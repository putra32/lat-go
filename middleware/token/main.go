package main

import (
	"fmt"
	"net/http"
)

// Fungsi CekLogin yang berguna sebagai middleware
func CekLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("token") != "12345" {
			fmt.Fprintf(rw, "Token tidak tersedia atau salah\n")
			return
		}
		next.ServeHTTP(rw, r)
	})
}

// Fungsi GetMahasiswa untuk menampilkan text string di browser
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Anda Berhasil Mengakses Fungsi GetMahasiswa()</h1>"))
}

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":5000",
	}

	// routing
	http.Handle("/", CekLogin(http.HandlerFunc(GetMahasiswa)))

	// jalankan server
	server.ListenAndServe()
}
