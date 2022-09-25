package main

import (
	"net/http"
)

func (a *webApp) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.leaseViewHandler)

	return mux
}
