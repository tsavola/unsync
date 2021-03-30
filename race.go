// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build race

package unsync

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

func atomicStorePointer(addr *unsafe.Pointer, val unsafe.Pointer) {
	runtime.RaceDisable()
	defer runtime.RaceEnable()

	atomic.StorePointer(addr, val)
}
