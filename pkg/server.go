package pkg

import (
	"fmt"
	"net/http"
)

const (
	port = ":8000"
)

func StartServer() {
	err := http.ListenAndServe(port, nil)
	fmt.Println(err)
}
