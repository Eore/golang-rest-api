package libs

import (
	"net/http"
)

const (
	PORT = ":8000"
)

func StartServer() {
	http.ListenAndServe(PORT, nil)
}
