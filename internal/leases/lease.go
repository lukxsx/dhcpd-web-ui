package leases

import (
	"fmt"
	"net/netip"
	"time"
)

type Lease struct {
	IP        netip.Addr
	Hostname  string
	MAC       string
	StartTime time.Time
	EndTime   time.Time
}

func (l Lease) String() string {
	return fmt.Sprintf("Lease %s:\n  MAC:\t%s\n  Hostname:\t%s\n  Starts:\t%s\n  Ends:\t\t%s\n", l.IP, l.MAC, l.Hostname, l.StartTime, l.EndTime)
}
