## To execute de application

### Run docker compose  
docker compose up  
or   
docker-compose up

## To run tests

### It is necessary to have Golang installed  

In the root directory application, run:
#### Tests coverage   
go test -cover  ./...  
#### Benchmark
go test -bench=. -cpu=1 -benchmem -benchtime=5s -count=5 ./benchmark/
