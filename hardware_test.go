// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, uint64(0x313233343536), encode([]byte("123456")))
}

func TestHardware(t *testing.T) {
	assert.NotEqual(t, Fingerprint(0), GetHardware())
	assert.NotEqual(t, "", GetHardware().String())
	assert.NotEqual(t, "", GetHardware().Hex())
}
