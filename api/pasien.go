package api

var Pasien = []API{
	API{
		Method:     "GET",
		URL:        "/pasien",
		Controller: "",
		Privillage: []string{"admin"},
	},
}
