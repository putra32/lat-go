package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/config"
	"rest-api/mahasiswa"
	"rest-api/models"
	"rest-api/utils"
	"strconv"
)

// GetMahasiswa
func GetMahasiswa(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		mahasiswas, err := mahasiswa.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(rw, mahasiswas, http.StatusOK)
		return
	}

	http.Error(rw, "Tidak diijinkan", http.StatusNotFound)
}

// PostMahasiswa
func PostMahasiswa(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.Mahasiswa

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.ResponseJSON(rw, err, http.StatusBadRequest)
		}

		if err := mahasiswa.Insert(ctx, mhs); err != nil {
			utils.ResponseJSON(rw, err, http.StatusInternalServerError)
			return
		}
		res := map[string]string{
			"status": "Successfully",
		}

		utils.ResponseJSON(rw, res, http.StatusCreated)
		return
	}

	http.Error(rw, "Tidak di ijinkan", http.StatusMethodNotAllowed)

}

// Update mahasiswa
func UpdateMahasiswa(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.Mahasiswa

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.ResponseJSON(rw, err, http.StatusBadRequest)
			return
		}

		fmt.Println(mhs)

		if err := mahasiswa.Update(ctx, mhs); err != nil {
			utils.ResponseJSON(rw, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(rw, res, http.StatusCreated)
		return
	}

	http.Error(rw, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}

// Delete Mahasiswa
func DeleteMahasiswa(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var mhs models.Mahasiswa

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(rw, "id tidak boleh kosong", http.StatusBadRequest)
		}

		mhs.ID, _ = strconv.Atoi(id)

		if err := mahasiswa.Delete(ctx, mhs); err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(rw, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(rw, res, http.StatusOK)
		return
	}
	utils.ResponseJSON(rw, "Tidak diijinkan", http.StatusMethodNotAllowed)
}

func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()

	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	// Routing
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	http.HandleFunc("/mahasiswa/create", PostMahasiswa)
	http.HandleFunc("/mahasiswa/update", UpdateMahasiswa)
	http.HandleFunc("/mahasiswa/delete", DeleteMahasiswa)

	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
