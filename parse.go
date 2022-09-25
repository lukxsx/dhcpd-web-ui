package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/netip"
	"regexp"
)

// Map of regex matchers and functions to fill the variables in the struct
var matchers = map[*regexp.Regexp]func(*Lease, string){
	// Parse IP address
	regexp.MustCompile("(.*) {"): func(l *Lease, value string) {
		var err error
		l.IP, err = netip.ParseAddr(value)
		if err != nil {
			log.Println("Error parsing IP: "+value, err)
		}
	},

	// Parse hostname
	regexp.MustCompile(`client-hostname "(.*)";`): func(l *Lease, value string) {
		l.Hostname = value
	},
}

// Parse fields from a lease entry into a lease struct
func (l *Lease) parseLeaseEntry(data []byte) {
	buf := bytes.NewBuffer(data)
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanLines)

	// Read every line in the entry
	for s.Scan() {
		line := s.Text()
		// Check for matching fields and add them to the lease struct
		for r, f := range matchers {
			if r.MatchString(line) {
				f(l, r.FindStringSubmatch(line)[1])
			}
		}
	}

}

// Read lease entries from a file
func ParseLeases(r io.Reader) []Lease {
	s := bufio.NewScanner(r)

	// Splice file to lease entries using a custom split function
	s.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF {
			return 0, nil, nil
		}

		// Find line starting with "lease"
		if i := bytes.Index(data, []byte("lease")); i != -1 {
			i += 6

			// Find the closing bracket
			if j := bytes.Index(data[i:], []byte("}")); j != -1 {
				return i + j, data[i : i+j], nil
			}
		}

		return 0, nil, nil
	})

	var leases []Lease
	for s.Scan() {
		l := Lease{}
		l.parseLeaseEntry(s.Bytes())
		leases = append(leases, l)
	}

	return leases
}
