package main

import (
	"./libs"
	// "./models"
)

func main() {
	libs.Router()
	// newPasien := models.Pasien{
	// 	NomorPengenal:      "111",
	// 	JenisNomorPengenal: "KTP",
	// 	Nama:               "miaw",
	// }

	// con := libs.Connection()

	// q, _ := con.Prepare("insert into pasien values ($1, $2, $3)")
	// _, err1 := q.Exec(newPasien.NomorPengenal, newPasien.JenisNomorPengenal, newPasien.Nama)

	// // panic(err0)
	// panic(err1)

	// fmt.Println(newPasien.Nama)
}
