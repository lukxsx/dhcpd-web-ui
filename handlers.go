package main

import (
	"html/template"
	"net/http"
)

// Leases view handler
func (a *webApp) leaseViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// HTML template files
	files := []string{
		"./assets/html/layout.html",
		"./assets/html/leases.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.errorLogger.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		a.errorLogger.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
