### Benchmark results

The benchmark uses a varying number of goroutines, each of which lock and
unlock a shared mutex in a loop.

- BenchmarkSyncMutex uses a `sync.Mutex`.
- BenchmarkUnsyncMutexLocking uses an `unsync.Mutex` which always synchronizes.
- BenchmarkUnsyncMutexNopLock uses an `unsync.Mutex` which never synchronizes.

Run on a CPU with 4 cores and 8 hyperthreads, with increasing number of
goroutines:

	goos: linux
	goarch: amd64
	pkg: github.com/tsavola/unsync
	cpu: Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz

	BenchmarkSyncMutex                	97114665	        12.32 ns/op
	BenchmarkUnsyncMutexLocking       	73442139	        16.14 ns/op
	BenchmarkUnsyncMutexNopLock       	226991307	         5.286 ns/op

	BenchmarkSyncMutex-2              	21699235	        85.10 ns/op
	BenchmarkUnsyncMutexLocking-2     	19938133	        97.72 ns/op
	BenchmarkUnsyncMutexNopLock-2     	214588460	         5.314 ns/op

	BenchmarkSyncMutex-4              	 4835926	       274.4 ns/op
	BenchmarkUnsyncMutexLocking-4     	 4273599	       300.0 ns/op
	BenchmarkUnsyncMutexNopLock-4     	184960027	         6.639 ns/op

	BenchmarkSyncMutex-6              	 2133309	       521.1 ns/op
	BenchmarkUnsyncMutexLocking-6     	 2037912	       555.3 ns/op
	BenchmarkUnsyncMutexNopLock-6     	153748228	         7.980 ns/op

	BenchmarkSyncMutex-8              	 1544576	       848.7 ns/op
	BenchmarkUnsyncMutexLocking-8     	 1454521	       882.7 ns/op
	BenchmarkUnsyncMutexNopLock-8     	142711644	         8.401 ns/op

	BenchmarkSyncMutex-10             	 1000000	      1165 ns/op
	BenchmarkUnsyncMutexLocking-10    	 1000000	      1199 ns/op
	BenchmarkUnsyncMutexNopLock-10    	86365126	        11.77 ns/op

	BenchmarkSyncMutex-12             	 1000000	      1339 ns/op
	BenchmarkUnsyncMutexLocking-12    	  921028	      1501 ns/op
	BenchmarkUnsyncMutexNopLock-12    	94899232	        12.64 ns/op

Summary:

- A synchronizing `unsync.Mutex` is slower than a `sync.Mutex`.
- An unsynchronizing `unsync.Mutex` is fast and scalable.

