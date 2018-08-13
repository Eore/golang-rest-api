package libs

import (
	"net/http"

	"../api"
)

func Router() {
	for _, val := range api.Pasien {
		http.HandleFunc(val.URL)
	}
}
