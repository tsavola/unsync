// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unsync

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkSyncMutex(b *testing.B) {
	var mu sync.Mutex

	parallelize(func() {
		for i := 0; i < b.N; i++ {
			mu.Lock()
			runtime.KeepAlive(nil)
			mu.Unlock()
		}
	})
}

func BenchmarkUnsyncMutexLocking(b *testing.B) {
	var mu Mutex
	mu.Init(nil)

	parallelize(func() {
		for i := 0; i < b.N; i++ {
			l := mu.LockMaybe()
			runtime.KeepAlive(nil)
			l.Unlock()
		}
	})

	mu.Unsync()
}

func BenchmarkUnsyncMutexNopLock(b *testing.B) {
	var mu Mutex
	mu.Init(nil)
	mu.Unsync()

	parallelize(func() {
		for i := 0; i < b.N; i++ {
			l := mu.LockMaybe()
			runtime.KeepAlive(nil)
			l.Unlock()
		}
	})
}

func parallelize(f func()) {
	var (
		group sync.WaitGroup
		start = make(chan struct{})
		count = runtime.GOMAXPROCS(0)
	)

	for i := 0; i < count; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			<-start
			f()
		}()
	}

	close(start)
	group.Wait()
}
