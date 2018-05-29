// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package address

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		addr string
		err  bool
	}{
		{addr: "private"},
		{addr: "private:8080"},
		{addr: "public"},
		{addr: "10.0.0.1:4000"},
		{addr: "", err: true},
		{addr: ":8000"},
		{addr: "external"},
		{addr: "external:1234"},
	}

	for _, tc := range tests {
		t.Run("Parsing"+tc.addr, func(*testing.T) {
			_, err := Parse(tc.addr, 80)
			assert.Equal(t, tc.err, err != nil)
			if err != nil {
				println(err.Error())
			}
		})
	}
}

func TestIsPrivateIP(t *testing.T) {
	tests := []struct {
		ip      string
		private bool
	}{
		// IPv4 private addresses
		{"10.0.0.1", true},    // private network address
		{"100.64.0.1", true},  // shared address space
		{"172.16.0.1", true},  // private network address
		{"192.168.0.1", true}, // private network address
		{"192.0.0.1", true},   // IANA address
		{"192.0.2.1", true},   // documentation address
		{"127.0.0.1", true},   // loopback address
		{"169.254.0.1", true}, // link local address

		// IPv4 public addresses
		{"1.2.3.4", false},

		// IPv6 private addresses
		{"::1", true},         // loopback address
		{"fe80::1", true},     // link local address
		{"fc00::1", true},     // unique local address
		{"fec0::1", true},     // site local address
		{"2001:db8::1", true}, // documentation address

		// IPv6 public addresses
		{"2004:db6::1", false},
	}

	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			ip := net.ParseIP(tt.ip)
			if ip == nil {
				t.Fatalf("%s is not a valid ip address", tt.ip)
			}
			if got, want := isPrivate(ip), tt.private; got != want {
				t.Fatalf("got %v for %v want %v", got, ip, want)
			}
		})
	}
}
