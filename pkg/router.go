package pkg

import (
	"net/http"
)

type API struct {
	Method     string
	URL        string
	Controller func(w http.ResponseWriter, r *http.Request)
	Privillage []string
}

func Router(apiRoute []API) {
	for _, val := range apiRoute {
		http.HandleFunc("/api/v1"+val.URL, val.Controller)
	}
}
