// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	thechar        = '5'
	_BigEndian     = 0
	_CacheLineSize = 32
	_PhysPageSize  = 65536*goos_nacl + 4096*(1-goos_nacl)
	_PCQuantum     = 4
	_Int64Align    = 4
	hugePageSize   = 0
	minFrameSize   = 4
)

type uintreg uint32
type intptr int32 // TODO(rsc): remove
