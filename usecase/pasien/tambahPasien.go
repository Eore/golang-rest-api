package controllers

import (
	"fmt"
	"net/http"

	"../../model"
)

func TambahPasien(w http.ResponseWriter, r *http.Request) {
	m := model.Pasien{
		NomorPengenal: "1234",
	}

	fmt.Println(m)
	fmt.Fprintf(w, "Test")
}
