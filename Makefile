# Include variables from the .env file
include .env

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go tool air