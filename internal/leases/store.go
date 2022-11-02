package leases

import (
	"log"
	"os"
	"sync"
	"time"
)

type LeaseStore struct {
	leases      []Lease
	updateMtx   sync.RWMutex
	updated     time.Time
	Filename    string
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
}

// Update the lease store
func (s *LeaseStore) Update() error {
	s.InfoLogger.Println("Updating lease data")
	s.updateMtx.Lock()
	defer s.updateMtx.Unlock()

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
	// Update lease data if it's over 10 seconds old
	currentTime := time.Now()
	if currentTime.Sub(s.updated).Seconds() > 10 {
		s.InfoLogger.Println("Lease data over 10 seconds old, updating...")
		s.Update()
	}

	s.updateMtx.RLock()
	defer s.updateMtx.RUnlock()

	return s.leases
}

// Return only active leases (lease ends after current time)
func (s *LeaseStore) GetActiveLeases() []Lease {
	// Update lease data if it's over 10 seconds old
	currentTime := time.Now()
	if currentTime.Sub(s.updated).Seconds() > 10 {
		s.InfoLogger.Println("Lease data over 10 seconds old, updating...")
		s.Update()
	}

	s.updateMtx.RLock()
	defer s.updateMtx.RUnlock()

	var filtered []Lease
	for _, l := range s.leases {
		if l.EndTime.After(currentTime) {
			filtered = append(filtered, l)
		}
	}

	return filtered
}
