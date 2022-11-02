package leases

import (
	"log"
	"os"
	"time"
)

type LeaseStore struct {
	leases      []Lease
	updated     time.Time
	Filename    string
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
}

// Update the lease store
func (s *LeaseStore) Update() error {
	s.InfoLogger.Println("Updating lease data")

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

	s.updated = time.Now()
	return nil
}

// Return all leases in the database
func (s *LeaseStore) GetAllLeases() []Lease {
	currentTime := time.Now()

	// Update lease data if it's over 10 seconds old
	if currentTime.Sub(s.updated).Seconds() > 10 {
		s.InfoLogger.Println("Lease data over 10 seconds old, updating...")
		s.Update()
	}

	return s.leases
}

// Return only active leases (lease ends after current time)
func (s *LeaseStore) GetActiveLeases() []Lease {
	currentTime := time.Now()

	// Update lease data if it's over 10 seconds old
	if currentTime.Sub(s.updated).Seconds() > 10 {
		s.InfoLogger.Println("Lease data over 10 seconds old, updating...")
		s.Update()
	}

	var filtered []Lease
	for _, l := range s.leases {
		if l.EndTime.After(currentTime) {
			filtered = append(filtered, l)
		}
	}

	return filtered
}
