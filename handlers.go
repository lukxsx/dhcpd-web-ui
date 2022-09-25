package main

import (
	"html/template"
	"net/http"
	"net/netip"
)

type templateData struct {
	Leases []Lease
}

// Lease list handler
func (a *webApp) leaseListHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// HTML template files
	files := []string{
		"./assets/html/layout.html",
		"./assets/html/leaselist.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.errorLogger.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := &templateData{
		Leases: []Lease{
			{IP: netip.MustParseAddr("10.242.11.1"), Hostname: "asdfff"},
		},
	}

	err = ts.ExecuteTemplate(w, "layout", data)
	if err != nil {
		a.errorLogger.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
