# goroutine-sum

Simple project to learn more about [goroutines](https://go.dev/tour/concurrency/1). 

In `main.go` you will find 2 implementations of `sum` function:
- `ClassicSum` which iterates through all the integer array sequentially
- `GoroutineSum` which splits the integer array into chunks and computes the sum of each of them through goroutines and [channels](https://go.dev/tour/concurrency/2)

## Benchmark
Benchmark tests with 2 different input size are in `main_test.go`. You can find benchmark results in `results` folder.

To run the benchmarks use `go test -bench=.`.