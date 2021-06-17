package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(rw http.ResponseWriter, p interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)

	rw.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(rw, "error tan", http.StatusBadRequest)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write([]byte(ubahkeByte))
}
