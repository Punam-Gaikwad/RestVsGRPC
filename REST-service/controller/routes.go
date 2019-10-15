package controller

import "net/http"

//Register -
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/records", movieRecords())
	return mux
}
