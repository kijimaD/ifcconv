.PHONY: test
test:
	docker-compose exec app go test ./... -v
