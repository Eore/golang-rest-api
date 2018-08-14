package rest

import (
	c "../../controller/pasien"
	"../../pkg"
)

var Pasien = []pkg.API{
	pkg.API{
		Method:     "POST",
		URL:        "/pasien",
		Controller: c.TambahPasien,
		Privillage: []string{"admin"},
	},
}
