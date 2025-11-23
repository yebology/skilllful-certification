.PHONY: run test

-include .env

run: 
	go run main.go

test:
	go test -v ./app/controller 

restart-migrate:
	cd cmd/reset && go run main.go && cd ../seeder && go run main.go