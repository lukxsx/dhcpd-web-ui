package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"dhcpd-ui/internal/leases"
)

type application struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
	leaseStore  *leases.LeaseStore
}

func main() {
	// Parse command line flags
	leasefile := flag.String("leasefile", "/var/db/dhcpd.leases", "dhcpd.leases file location")
	flag.Parse()

	// Initialize loggers
	infoLogger := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime)

	store := &leases.LeaseStore{
		Filename: *leasefile,
	}

	err := store.Update()
	if err != nil {
		errorLogger.Fatal(err)
	}

	// Initialize the web application struct
	app := &application{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
		leaseStore:  store,
	}

	server := &http.Server{
		Addr:     ":3000",
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	for _, l := range store.GetActiveLeases() {
		fmt.Println(l)
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
