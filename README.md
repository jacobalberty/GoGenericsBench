
## About
This is a collection of various sorts implemented both as []int only and as `constraints.Ordered`[^1] generics.
I split the generics into their own source file, so they aren't ordered in the output, that way you can still run
the normal []int sorting functions under go 1.17.

[^1]:The generic counting sort uses `constraints.Integer` due to limitations in what a counting sort can do.

## Usage
The benchmarks are executed by first installing [gotip](https://pkg.go.dev/golang.org/dl/gotip) with the commands `go install golang.org/dl/gotip@latest` then `gotip download`. After that just launch the benchmarks with `gotip test -bench . -benchmem` you can also pipe the results into sort and it will arrange the benchmarks in a way that makes it easier to compare performance visually.
You can also test just the []int sorts with `go test -bench . -benchmem` without having to install gotip.
### Array size
This defaults to sorting an array of 500 pseudorandom (generated with math/rand) integers. You can change the array size by setting the `ARRAY_SIZE` environment variable.

## Results
All results are from running `gotip test -bench . -benchmem` which by default will use an array that contains 500 integers.
### AMD Ryzen 7 5800x
`go version devel go1.18-c5fee935bb Fri Dec 17 16:05:31 2021 +0000 linux/amd64`
```
goos: linux
goarch: amd64
pkg: github.com/jacobalberty/GoGenericsBench
cpu: AMD Ryzen 7 5800X 8-Core Processor
BenchmarkBubbleSortG-16                     6674            198354 ns/op            4096 B/op          1 allocs/op
BenchmarkCountingSortG-16                 529876              3062 ns/op           12288 B/op          3 allocs/op
BenchmarkInsertSortG-16                    15571             76483 ns/op            4096 B/op          1 allocs/op
BenchmarkMergeSortG-16                     37220             29168 ns/op           40672 B/op        500 allocs/op
BenchmarkParallelMergeSortG-16             29306             41303 ns/op           40856 B/op        503 allocs/op
BenchmarkSelectionSortG-16                 10155            112004 ns/op            4096 B/op          1 allocs/op
BenchmarkBubbleSort-16                      7029            172419 ns/op            4096 B/op          1 allocs/op
BenchmarkCountingSort-16                  496998              3202 ns/op           12288 B/op          3 allocs/op
BenchmarkInsertSort-16                     13842             80079 ns/op            4096 B/op          1 allocs/op
BenchmarkMergeSort-16                      36690             32712 ns/op           40672 B/op        500 allocs/op
BenchmarkParallelMergeSort-16              29331             40627 ns/op           40856 B/op        503 allocs/op
BenchmarkSelectionSort-16                  11428            116093 ns/op            4096 B/op          1 allocs/op
PASS
ok      github.com/jacobalberty/GoGenericsBench 20.662s
```
### Raspbery Pi 4 (32 bit)
`go version devel go1.18-c5fee93 Fri Dec 17 16:05:31 2021 +0000 linux/arm`
```
goos: linux
goarch: arm
pkg: github.com/jacobalberty/GoGenericsBench
BenchmarkBubbleSortG-4                      1273            944464 ns/op            2048 B/op          1 allocs/op
BenchmarkCountingSortG-4                   68049             16914 ns/op            6144 B/op          3 allocs/op
BenchmarkInsertSortG-4                      3372            354606 ns/op            2048 B/op          1 allocs/op
BenchmarkMergeSortG-4                       6297            165212 ns/op           20384 B/op        500 allocs/op
BenchmarkParallelMergeSortG-4               6486            187331 ns/op           20500 B/op        503 allocs/op
BenchmarkSelectionSortG-4                   2169            554638 ns/op            2048 B/op          1 allocs/op
BenchmarkBubbleSort-4                       1264            958238 ns/op            2048 B/op          1 allocs/op
BenchmarkCountingSort-4                    68227             17410 ns/op            6144 B/op          3 allocs/op
BenchmarkInsertSort-4                       3226            374070 ns/op            2048 B/op          1 allocs/op
BenchmarkMergeSort-4                        6884            168408 ns/op           20384 B/op        500 allocs/op
BenchmarkParallelMergeSort-4                6103            188842 ns/op           20497 B/op        503 allocs/op
BenchmarkSelectionSort-4                    2185            554137 ns/op            2048 B/op          1 allocs/op
PASS
ok      github.com/jacobalberty/GoGenericsBench 15.055s
```
### AMD Phenom II X4 925
`go version devel go1.18-c5fee93 Fri Dec 17 16:05:31 2021 +0000 linux/amd64`
```
goos: linux
goarch: amd64
pkg: github.com/jacobalberty/GoGenericsBench
cpu: AMD Phenom(tm) II X4 925 Processor
BenchmarkBubbleSortG-4                      1908            630381 ns/op            4096 B/op          1 allocs/op
BenchmarkCountingSortG-4                   48958             23373 ns/op           12288 B/op          3 allocs/op
BenchmarkInsertSortG-4                      5595            227135 ns/op            4096 B/op          1 allocs/op
BenchmarkMergeSortG-4                      10000            214319 ns/op           40672 B/op        500 allocs/op
BenchmarkParallelMergeSortG-4              10060            108768 ns/op           40857 B/op        503 allocs/op
BenchmarkSelectionSortG-4                   3584            352909 ns/op            4096 B/op          1 allocs/op
BenchmarkBubbleSort-4                       1868            627303 ns/op            4096 B/op          1 allocs/op
BenchmarkCountingSort-4                    57622             21937 ns/op           12288 B/op          3 allocs/op
BenchmarkInsertSort-4                       5494            231653 ns/op            4096 B/op          1 allocs/op
BenchmarkMergeSort-4                        5904            223928 ns/op           40672 B/op        500 allocs/op
BenchmarkParallelMergeSort-4                9300            119276 ns/op           40856 B/op        503 allocs/op
BenchmarkSelectionSort-4                    3284            373200 ns/op            4096 B/op          1 allocs/op
PASS
ok      github.com/jacobalberty/GoGenericsBench 21.828s
```
