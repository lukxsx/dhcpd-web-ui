package main

import (
	"net/netip"
)

type Lease struct {
	IP       netip.Addr
	Hostname string
	MAC      string
}
