package main

import (
	"net/http"
)

func (a *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", a.staticServer))
	mux.HandleFunc("/", a.leaseListHandler)

	return mux
}
