
### Tests
go test -cover  ./...  
### Benchmark
go test -bench=. -cpu=1 -benchmem -benchtime=5s -count=5 ./benchmark/
