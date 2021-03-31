### Benchmark results

The benchmark uses a varying number of goroutines, each of which simulate
access to a shared resource in a loop.

- BenchmarkSyncMutex uses a `sync.Mutex`.
- BenchmarkUnsyncMutexLocking uses an `unsync.Mutex` which always synchronizes.
- BenchmarkUnsyncMutexNopLock uses an `unsync.Mutex` which never synchronizes.
- BenchmarkAtomicLoad does an atomic load from a uint32 variable.

Run on a t4g.2xlarge VM with 8 vCPUs, with increasing number of goroutines:

	goos: linux
	goarch: arm64
	pkg: github.com/tsavola/unsync
	BenchmarkSyncMutex-4              	 4302242	       358.2 ns/op
	BenchmarkSyncMutex-8              	 1686327	       855.0 ns/op
	BenchmarkSyncMutex-16             	  685428	      1639 ns/op
	BenchmarkUnsyncMutexLocking-4     	 3636868	       382.6 ns/op
	BenchmarkUnsyncMutexLocking-8     	 1266651	      1001 ns/op
	BenchmarkUnsyncMutexLocking-16    	  704870	      2124 ns/op
	BenchmarkUnsyncMutexNopLock-4     	206438330	         5.824 ns/op
	BenchmarkUnsyncMutexNopLock-8     	205063986	         5.873 ns/op
	BenchmarkUnsyncMutexNopLock-16    	100000000	        12.57 ns/op
	BenchmarkAtomicLoad-4             	998210292	         1.202 ns/op
	BenchmarkAtomicLoad-8             	997223380	         1.204 ns/op
	BenchmarkAtomicLoad-16            	476909344	         2.555 ns/op

Summary:

- A synchronizing `unsync.Mutex` is slower than a `sync.Mutex`.
- An unsynchronizing `unsync.Mutex` is fast and scalable.
- An atomic load is faster than the overhead caused by this abstraction.

