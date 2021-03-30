// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package unsync provides a lock type which can be initially used for mutual
// exclusion, but avoids synchronization overhead in the long run.
//
// The trade-off is that an extra heap-allocation may be needed, and the total
// memory usage is higher than with sync.Mutex.  But after the synchronization
// period is over, the allocation is released and remaining memory usage is
// comparable to sync.Mutex (on 32-bit systems it is even lower).
//
// See https://github.com/tsavola/unsync for benchmark results.
package unsync

// The implementation relies on lazy propagation of the memory write in the
// Unsync call to memory reads in LockMaybe calls.  LockMaybe is on the fast
// path: its memory access non-atomic.  It is assumed that the write will
// eventually be seen by the readers.  Until that, LockMaybe calls are locking
// the "dangling" sync.Mutex unnecessarily.

import (
	"sync"
	"unsafe"
)

// Mutex is a special kind of mutual exclusion lock.  The zero value represents
// an invalid state: Init must be called before LockMaybe.
//
// The value must not be copied after first use.
type Mutex struct {
	locker unsafe.Pointer // *sync.Mutex
}

// Init this mutex before it is shared between goroutines.  If locker is
// non-nil, it is used for locking (until the Unsync call).  Otherwise a
// sync.Mutex is allocated.
func (m *Mutex) Init(locker *sync.Mutex) {
	if locker == nil {
		locker = new(sync.Mutex)
	}
	m.locker = unsafe.Pointer(locker)
}

// LockMaybe might lock the mutex.  Programs must be structured in such a way
// that there is no concurrent access to the resources protected by the mutex
// after Unsync is called.  If Unsync is never called, LockMaybe will always
// lock the mutex.
func (m *Mutex) LockMaybe() Unlocker {
	l := (*sync.Mutex)(m.locker) // Non-atomic load on purpose.
	if l != &sentinel {
		l.Lock()
	}
	return Unlocker{l}
}

// Unsync releases the mutex from the burden of synchronization.  Future
// LockMaybe calls may or may not lock the mutex.
func (m *Mutex) Unsync() {
	atomicStorePointer(&m.locker, unsafe.Pointer(&sentinel))
}

// Unlocker of a mutex.
type Unlocker struct {
	locker *sync.Mutex
}

// Unlock the associated mutex if it was locked by the matching LockMaybe call.
func (u Unlocker) Unlock() {
	if u.locker != &sentinel {
		u.locker.Unlock()
	}
}

var sentinel sync.Mutex // Represents the unsynchronizing state.
