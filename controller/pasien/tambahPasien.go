package controllers

import (
	"net/http"

	"../../model"
	"../../repository"
)

func TambahPasien(w http.ResponseWriter, r *http.Request) {
	// tanggalLahir, _ := time.Parse("2006-01-02", r.FormValue("tanggalLahir"))
	formValue := model.Pasien{
		NomorPengenal:        r.FormValue("nomorPengenal"),
		JenisNomorPengenal:   r.FormValue("jenisNomorPengenal"),
		Nama:                 r.FormValue("nama"),
		TempatLahir:          r.FormValue("tempatLahir"),
		TanggalLahir:         r.FormValue("tanggalLahir"),
		NamaPenangungJawab:   r.FormValue("namaPenanggungJawab"),
		Agama:                r.FormValue("agama"),
		Alamat:               r.FormValue("alamat"),
		AlamatDomisili:       r.FormValue("alamatDomisili"),
		NomorTelepon:         r.FormValue("nomorTelepon"),
		NomorHandphone:       r.FormValue("nomorHandphone"),
		Pekerjaan:            r.FormValue("pekerjaan"),
		NamaPerusahaan:       r.FormValue("namaPerusahaan"),
		AlamatPerusahaan:     r.FormValue("alamatPerusahaan"),
		NamaKeluarga:         r.FormValue("namaKeluarga"),
		AlamatKeluarga:       r.FormValue("alamatKeluarga"),
		NomorTeleponKeluarga: r.FormValue("nomorTeleponKeluarga"),
		PenjaminBiaya:        r.FormValue("penjaminBiaya"),
	}
	repository.InsertPasien(formValue)
}
