## To execute de application

### Run docker compose  
docker compose up  
or   
docker-compose up

### CURLs to test 

### Create a campaign  
curl --location 'http://localhost:3000/campaigns' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "Campaign Name",
"content": "Content 01",
"emails":["joao@email.com","maria@email.com"]
}'  
  
### Find campaigns  
curl --location 'http://localhost:3000/campaigns'  
  
### Get a campaign by Id  
curl --location 'http://localhost:3000/campaigns/1'

## To execute code tests

### It is necessary to have Golang installed  

In the root directory application, run:
#### Tests coverage   
go test -cover  ./...  
#### Benchmark
go test -bench=. -cpu=1 -benchmem -benchtime=5s -count=5 ./benchmark/
