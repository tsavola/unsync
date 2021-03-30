// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !race

package unsync

import (
	"sync/atomic"
	"unsafe"
)

func atomicStorePointer(addr *unsafe.Pointer, val unsafe.Pointer) {
	atomic.StorePointer(addr, val)
}
