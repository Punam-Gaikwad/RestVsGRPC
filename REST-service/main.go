package main

import (
	"net/http"

	"github.com/Punam-Gaikwad/RestVsGRPC/REST-service/controller"
)

const (
	port = "8000"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":"+port, mux)
}
