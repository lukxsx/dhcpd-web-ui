package main

import (
	"dhcpd-ui/internal/leases"
	"html/template"
	"net/http"
)

type templateData struct {
	ActiveLeases []leases.Lease
	AllLeases    []leases.Lease
}

// Lease list handler
func (a *application) leaseListHandler(w http.ResponseWriter, r *http.Request) {
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
		ActiveLeases: a.leaseStore.GetActiveLeases(),
		AllLeases:    a.leaseStore.GetAllLeases(),
	}

	err = ts.ExecuteTemplate(w, "layout", data)
	if err != nil {
		a.errorLogger.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
