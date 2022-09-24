package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	leasefile := flag.String("leasefile", "/var/db/dhcpd.leases", "dhcpd.leases file location")
	flag.Parse()

	f, err := os.Open(*leasefile)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}
