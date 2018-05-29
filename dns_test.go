// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolve(t *testing.T) {
	tests := []struct {
		addr   string
		err    bool
		expect string
	}{
		{addr: "", err: true},
		{addr: "https://google.com"},
		{addr: "google.com"},
		{addr: "google.com/1234/sds"},
		{addr: "https://127.0.0.1", expect: "127.0.0.1:80"},
		{addr: "https://4987e9gs99sxdg8e8e7tgwe5.boomshouldnotwork", err: true},
		{addr: "127.0.0.1", expect: "127.0.0.1:80"},
		{addr: "2001:db8::68", expect: "[2001:db8::68]:80"},
		{addr: "192.0.2.1", expect: "192.0.2.1:80"},
		{addr: "127.0.0.1:22", expect: "127.0.0.1:22"},
		{addr: "127.0.0.1:8xx0", err: true},
		{addr: "tcp://127.0.0.1:8000", expect: "127.0.0.1:8000"},
		{addr: "https://gooo:oogle.com", err: true},
		{addr: "https://google.com:xxx", err: true},
		{addr: "https-----:// ://--", err: true},
	}

	for _, tc := range tests {
		t.Run(tc.addr, func(*testing.T) {
			println("resolving " + tc.addr)

			addr, err := Resolve(tc.addr, 80)
			assert.Equal(t, tc.err, err != nil)
			if err != nil {
				println(err.Error())
			}

			if len(tc.expect) > 0 {
				assert.Equal(t, tc.expect, addr[0].String())
			}
		})
	}
}
