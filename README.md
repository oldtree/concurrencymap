# concurrencymap
concurrencymap implement and test
## test result

## flow operation
>       go test -cpu 4 -bench .

        Benchmark_ConcurrencyMapWithLock_Add-4           1000000              1067 ns/op
        Benchmark_ConcurrencyMapWithLock_Get-4           5000000               275 ns/op
        Benchmark_ConcurrencyMapWithLock_Remove-4       10000000               205 ns/op
        Benchmark_ConcurrencyMapWithChan_Add-4           1000000              1810 ns/op
        Benchmark_ConcurrencyMapWithChan_Get-4           1000000              1021 ns/op
        Benchmark_ConcurrencyMapWithChan_Remove-4        2000000               957 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Add-4         1000000              1045 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Get-4         5000000               276 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Remove-4     10000000               210 ns/op
        PASS
        ok      github.com/oldtree/ConcurrencyMap       22.319s    

## use go concurrency
> go test -cpu 4 -bench .

        Benchmark_ConcurrencyMapWithLock_Add-4           1000000              9338 ns/op
        Benchmark_ConcurrencyMapWithLock_Get-4           1000000              1882 ns/op
        Benchmark_ConcurrencyMapWithLock_Remove-4        1000000              2460 ns/op
        Benchmark_ConcurrencyMapWithLock_Random-4        1000000              1959 ns/op
        Benchmark_ConcurrencyMapWithChan_Add-4           1000000              5019 ns/op
        Benchmark_ConcurrencyMapWithChan_Get-4           1000000              1672 ns/op
        Benchmark_ConcurrencyMapWithChan_Remove-4        1000000              1297 ns/op
        Benchmark_ConcurrencyMapWithChan_Random-4        1000000              1336 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Add-4         1000000              1566 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Get-4         2000000               625 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Remove-4      2000000               760 ns/op
        Benchmark_ConcurrencyMapWithRWLock_Random-4      1000000              1069 ns/op
        PASS
        ok      github.com/oldtree/ConcurrencyMap       147.969s
