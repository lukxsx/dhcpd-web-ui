package leases

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/netip"
	"regexp"
	"time"
)

// Map of regex matchers and functions to fill the variables in the struct
var matchers = map[*regexp.Regexp]func(*Lease, string){
	// Parse IP address
	regexp.MustCompile("(.*) {"): func(l *Lease, value string) {
		l.IP = netip.MustParseAddr(value)
	},

	// Parse hostname
	regexp.MustCompile(`client-hostname "(.*)";`): func(l *Lease, value string) {
		l.Hostname = value
	},

	// Parse MAC address
	regexp.MustCompile(`hardware ethernet (.*);`): func(l *Lease, value string) {
		l.MAC = value
	},

	// Parse start time
	regexp.MustCompile(`starts (.*);`): func(l *Lease, value string) {
		var err error
		l.StartTime, err = time.Parse("2006/01/02 15:04:05 UTC", value[2:])
		if err != nil {
			log.Println(err)
		}
	},
	// Parse end time
	regexp.MustCompile(`ends (.*);`): func(l *Lease, value string) {
		var err error
		l.EndTime, err = time.Parse("2006/01/02 15:04:05 UTC", value[2:])
		if err != nil {
			log.Println(err)
		}
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
func parseLeases(r io.Reader) []Lease {
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
