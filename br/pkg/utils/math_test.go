// Copyright 2020 PingCAP, Inc. Licensed under Apache-2.0.

package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClampInt(t *testing.T) {
	require.Equal(t, 3, ClampInt(100, 1, 3))
	require.Equal(t, 2, ClampInt(2, 1, 3))
	require.Equal(t, 1, ClampInt(0, 1, 3))
	require.Equal(t, 1, ClampInt(0, 1, 1))
	require.Equal(t, 1, ClampInt(100, 1, 1))
}

func TestNextPowerOfTwo(t *testing.T) {
	require.Equal(t, int64(1), NextPowerOfTwo(1))
	require.Equal(t, int64(4), NextPowerOfTwo(3))
	require.Equal(t, int64(256), NextPowerOfTwo(255))
	require.Equal(t, int64(1024), NextPowerOfTwo(1024))
	require.Equal(t, int64(0x100000000), NextPowerOfTwo(0xabcd1234))
}
