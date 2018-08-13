package models

import (
	"time"

	"../libs"
)

type Pasien struct {
	NomorPengenal        string
	JenisNomorPengenal   string
	Nama                 string
	TempatLahir          string
	TanggalLahir         time.Time
	NamaPenangungJawab   string
	Agama                string
	Alamat               string
	AlamatDomisili       string
	NomorTelepon         string
	NomorHandphone       string
	Pekerjaan            string
	NamaPerusahaan       string
	AlamatPerusahaan     string
	NamaKeluarga         string
	AlamatKeluarga       string
	NomorTeleponKeluarga string
	PenjaminBiaya        string
}

func init() {
	libs.Connection().Query(`
		create table if not exists pasien (
			nomor_pengenal varchar(20) not null primary key,
			jenis_nomor_pengenal varchar(10) not null,
			nama varchar(25) not null,
			tempat_lahir varchar(20) not null,
			tanggal_lahir date not null,
			nama_penanggung_jawab varchar(25) not null,
			agama varchar(15) not null,
			alamat text not null,
			alamat_domisili text not null,
			nomor_telepon varchar(12) not null,
			nomor_handphone varchar(12) not null,
			pekerjaan varchar(20) not null,
			nama_perusahaan varchar(30),
			alamat_perusahaan text,
			nama_keluarga varchar(25) not null,
			alamat_keluarga text not null,
			nomor_telepon_keluarga varchar(12) not null,
			penjamin_biaya varchar(15) not null
		)
	`)
}
