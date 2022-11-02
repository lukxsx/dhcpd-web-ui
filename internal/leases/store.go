package leases

import (
	"os"
	"time"
)

type LeaseStore struct {
	Filename string
	leases   []Lease
}

// Update the lease store
func (s *LeaseStore) Update() error {
	// Open lease file
	f, err := os.Open(s.Filename)
	if err != nil {
		return err
	}
	defer f.Close()

	leases := parseLeases(f)
	if len(leases) != 0 {
		s.leases = leases
	}

	return nil
}

// Return all leases in the database
func (s *LeaseStore) GetAllLeases() []Lease {

	return s.leases
}

// Return only active leases (lease ends after current time)
func (s *LeaseStore) GetActiveLeases() []Lease {
	currentTime := time.Now()
	var filtered []Lease

	for _, l := range s.leases {
		if l.EndTime.After(currentTime) {
			filtered = append(filtered, l)
		}
	}

	return filtered
}
