package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"dhcpd-ui/internal/leases"
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

	f, err := os.Open(conf.LeaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	leases := leases.ParseLeases(f)
	for _, lease := range leases {
		fmt.Println(lease)
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
