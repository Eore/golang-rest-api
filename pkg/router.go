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
		http.HandleFunc(val.URL, val.Controller)
	}
}
