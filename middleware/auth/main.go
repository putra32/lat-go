package main

import "net/http"

// Fungsi log yang bergunan sebagai middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			rw.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}
		if uname == "kodingin" && pwd == "kodingin" {
			next.ServeHTTP(rw, r)
			return
		}

		rw.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

// Fungsi GetMahasiswa unutk menampilkan text string di browser
func GetMahasiswa(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		rw.Write([]byte("<h1>Anda Berhasil Mengakses Fungsi GetMahasiswa()</h1>"))
	}
}

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":5000",
	}

	// routing
	http.Handle("/", Auth(http.HandlerFunc(GetMahasiswa)))

	// jalankan server
	server.ListenAndServe()
}
