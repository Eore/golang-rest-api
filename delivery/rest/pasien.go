package rest

import (
	"../../pkg"
	c "../../usecase/pasien"
)

var Pasien = []pkg.API{
	pkg.API{
		Method:     "GET",
		URL:        "/pasien",
		Controller: c.TambahPasien,
		Privillage: []string{"admin"},
	},
}
