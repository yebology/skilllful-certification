.PHONY: run

-include .env

run: 
	go run main.go

restart-migrate:
	cd cmd/reset && go run main.go && cd ../seeder && go run main.go