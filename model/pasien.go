package model

import (
	"time"
)

type Pasien struct {
	NomorPengenal        string    `json:"nomorPengenal" sql:"varchar(20) not null primary key"`
	JenisNomorPengenal   string    `json:"jenisNomorPengenal" sql:"varchar(10) not null"`
	Nama                 string    `json:"nama" sql:"varchar(25) not null"`
	TempatLahir          string    `json:"tempatLahir" sql:"varchar(20) not null"`
	TanggalLahir         time.Time `json:"tanggalLahir" sql:"date not null"`
	NamaPenangungJawab   string    `json:"namaPenanggungJawab" sql:"varchar(25) not null"`
	Agama                string    `json:"agama" sql:"varchar(15) not null"`
	Alamat               string    `json:"alamat" sql:"text not null"`
	AlamatDomisili       string    `json:"alamatDomisili" sql:"text not null"`
	NomorTelepon         string    `json:"nomorTelepon" sql:"varchar(12) not null"`
	NomorHandphone       string    `json:"nomorHandphone" sql:"varchar(12) not null"`
	Pekerjaan            string    `json:"pekerjaan" sql:"varchar(20) not null"`
	NamaPerusahaan       string    `json:"namaPerusahaan" sql:"varchar(30)"`
	AlamatPerusahaan     string    `json:"alamatPerusahaan" sql:"text"`
	NamaKeluarga         string    `json:"namaKeluarga" sql:"varchar(25) not null"`
	AlamatKeluarga       string    `json:"alamatKeluarga" sql:"text not null"`
	NomorTeleponKeluarga string    `json:"nomorTeleponKeluarga" sql:"varchar(12) not null"`
	PenjaminBiaya        string    `json:"penjaminBiaya" sql:"varchar(15) not null"`
}
