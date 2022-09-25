package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type webApp struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
}

type Config struct {
	LeaseFile string
}

var conf Config

func main() {
	// Parse command line flags
	leasefile := flag.String("leasefile", "/var/db/dhcpd.leases", "dhcpd.leases file location")
	flag.Parse()
	conf.LeaseFile = *leasefile

	// Initialize loggers
	infoLogger := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime)

	// Initialize the web application struct
	app := &webApp{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}

	server := &http.Server{
		Addr:     ":3000",
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
