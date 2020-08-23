# bench-go

# Execution

```sh
$ cd malloc
$ go test -bench . -benchmem
```

# Result

```sh
goos: windows
goarch: amd64
pkg: bench-go/malloc
BenchmarkEveryStringAllocation/Try_10_times-4             750192              1371 ns/op             496 B/op          5 allocs/op
BenchmarkEveryStringAllocation/Try_100_times-4            106195             11352 ns/op            4260 B/op         98 allocs/op
BenchmarkEveryStringAllocation/Try_1,000_times-4           10000            122652 ns/op           45342 B/op       1745 allocs/op
BenchmarkEveryStringAllocation/Try_100,000_times-4            51          20960763 ns/op        10845747 B/op     199773 allocs/op
BenchmarkEveryStringAllocation/Try_10,000,000_times-4          1        2267867000 ns/op        987535144 B/op  19999825 allocs/op
BenchmarkOnceStringAllocation/Try_10_times-4             1443093               845 ns/op             160 B/op          1 allocs/op
BenchmarkOnceStringAllocation/Try_100_times-4             123711              9476 ns/op            1972 B/op         91 allocs/op
BenchmarkOnceStringAllocation/Try_1,000_times-4            10000            112544 ns/op           28972 B/op       1735 allocs/op
BenchmarkOnceStringAllocation/Try_100,000_times-4             99          12504959 ns/op         3202803 B/op     199738 allocs/op
BenchmarkOnceStringAllocation/Try_10,000,000_times-4           1        1296228800 ns/op        320004248 B/op  19999744 allocs/op
BenchmarkEveryStructAllocation/Try_10_times-4             237570              4563 ns/op            3464 B/op         75 allocs/op
BenchmarkEveryStructAllocation/Try_100_times-4             27057             43485 ns/op           35512 B/op        714 allocs/op
BenchmarkEveryStructAllocation/Try_1,000_times-4            1920            602038 ns/op          366683 B/op      12417 allocs/op
BenchmarkEveryStructAllocation/Try_100,000_times-4            13          78786515 ns/op        45684740 B/op    1299437 allocs/op
BenchmarkEveryStructAllocation/Try_10,000,000_times-4          1        9854731400 ns/op        4726933712 B/op 129999462 allocs/op
BenchmarkOnceStructAllocation/Try_10_times-4              266668              4428 ns/op            3296 B/op         71 allocs/op
BenchmarkOnceStructAllocation/Try_100_times-4              26666             44407 ns/op           34368 B/op        707 allocs/op
BenchmarkOnceStructAllocation/Try_1,000_times-4             1966            600767 ns/op          358498 B/op      12407 allocs/op
BenchmarkOnceStructAllocation/Try_100,000_times-4             14          73180943 ns/op        41833191 B/op    1299407 allocs/op
BenchmarkOnceStructAllocation/Try_10,000,000_times-4           1        7428980000 ns/op        4383433536 B/op 129999407 allocs/op
PASS
ok      bench-go/malloc 84.684s
```

# Reference
* [go/walk.go at go1.13.5 · golang/go](https://github.com/golang/go/blob/go1.13.5/src/cmd/compile/internal/gc/walk.go#L2585)
* [go/slice.go at release-branch.go1.13 · golang/go](https://github.com/golang/go/blob/release-branch.go1.13/src/runtime/slice.go#L76)