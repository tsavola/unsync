// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unsync

import (
	"testing"
)

func TestMutexUninit(t *testing.T) {
	var mu Mutex

	defer func() {
		if recover() == nil {
			t.Error("LockMaybe did not panic")
		}
	}()

	mu.LockMaybe()
}
