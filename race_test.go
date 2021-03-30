// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build race

package unsync

import (
	"sync"
	"testing"
)

func TestMutexRace(t *testing.T) {
	const maxParallel = 64
	const iterations = 100

	for parallel := 2; parallel < maxParallel; parallel++ {
		var mu Mutex
		mu.Init(nil)

		var wg sync.WaitGroup
		start := make(chan struct{})

		for i := 0; i < parallel; i++ {
			wg.Add(1)

			go func(i int) {
				defer wg.Done()
				<-start

				for j := 0; j < iterations; j++ {
					mu.LockMaybe().Unlock()
					if i == parallel/2 && j == iterations/2 {
						mu.Unsync()
					}
				}
			}(i)
		}

		close(start)
		wg.Wait()
	}
}
