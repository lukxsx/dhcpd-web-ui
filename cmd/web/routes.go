package main

import (
	"net/http"
)

func (a *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.leaseListHandler)

	return mux
}
