# concurrencymap
concurrencymap implement and test
## test result

>       grapestree$ go test -cpu 4 -bench .
        Benchmark_ConcurrencyMapWithLock_Add-4           1000000              1266 ns/op
        Benchmark_ConcurrencyMapWithLock_Get-4           5000000               292 ns/op
        Benchmark_ConcurrencyMapWithLock_Remove-4       10000000               242 ns/op
        Benchmark_ConcurrencyMapWithChan_Add-4           1000000              3500 ns/op
        Benchmark_ConcurrencyMapWithChan_Get-4           1000000              1136 ns/op
        Benchmark_ConcurrencyMapWithChan_Remove-4        1000000              1193 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Add-4         1000000              1924 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Get-4         5000000               398 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Remove-4     10000000               222 ns/op
        PASS
        ok      github.com/oldtree/ConcurrencyMap       25.501s     
