package main

import (
	"fmt"
	"net/netip"
)

type Lease struct {
	IP       netip.Addr
	Hostname string
	MAC      string
}

func (l Lease) String() string {
	return fmt.Sprintf("IP: %s\tMAC: %s\tHostname: %s", l.IP, l.MAC, l.Hostname)
}