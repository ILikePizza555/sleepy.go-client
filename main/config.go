package main

import (
	"net"
)

// Config holds the configuration data for the client.
// HubAddr is the IP address of the local hub. Only local ip addresses are allowed.
// Key is a 4-byte key shared between the hub and client.
type Config struct {
	HubAddr net.Addr
	Key     [4]byte
}
