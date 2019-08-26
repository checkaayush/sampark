# Prefer running recipes over files with same names
.PHONY: start stop test start-db

help:
	@echo
	@echo "Sampark"
	@echo
	@echo "  Commands: "
	@echo
	@echo "    help - Show this message."
	@echo "    start - Start all services."
	@echo "    stop - Stop all services."
	@echo "    start-db - Start database service."
	@echo "    test - Run tests with coverage report."

start:
	docker-compose up --build

stop:
	docker-compose down --build
	@echo 'Services stopped successfully.'

start-db:
	docker-compose up -d mongodb

test: start-db
	bash scripts/test.sh

clean:
	rm coverage.txt app
