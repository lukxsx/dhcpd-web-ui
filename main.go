package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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
	leases := ParseLeases(f)
	for _, l := range leases {
		fmt.Println(l)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	err = http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
