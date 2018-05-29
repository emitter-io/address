// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package address

import (
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
