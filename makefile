GO = $(shell which go)

export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_NAME=zweb

run:
	@if [ -z ${DB_PASS} ] ; then echo DB_PASS undefined, set DB_PASS first ; else $(GO) run main.go ; fi
	