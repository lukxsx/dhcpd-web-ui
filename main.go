package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type webApp struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
}

func main() {
	// Parse command line flags
	leasefile := flag.String("leasefile", "/var/db/dhcpd.leases", "dhcpd.leases file location")
	flag.Parse()

	// Open lease file
	f, err := os.Open(*leasefile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse leases
	leases := ParseLeases(f)
	for _, l := range leases {
		fmt.Println(l)
	}

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

	err = server.ListenAndServe()
	log.Fatal(err)
}
