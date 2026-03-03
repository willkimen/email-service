up:
	docker compose up

start:
	docker compose start

stop:
	docker compose stop

list:
	docker compose ps -a

build:
	docker compose up --build 

logs-email:
	mkdir -p logs
	docker compose logs --timestamps emailservice > logs/logs-email.txt

logs-redis:
	mkdir -p logs
	docker compose logs --timestamps redis-broker > logs/logs-redis.txt

logs-recent:
	mkdir -p logs
	docker compose logs --since=2m --timestamps > logs/logs-recent.txt

logs-email-recent:
	mkdir -p logs
	docker compose logs --since=2m --timestamps emailservice > logs/logs-email-recent.txt

logs-redis-recent:
	mkdir -p logs
	docker compose logs --since=2m --timestamps redis-broker > logs/logs-redis-recent.txt

logs-live:
	docker compose logs -f --timestamps

logs-email-live:
	docker compose logs -f --timestamps emailservice

logs-redis-live:
	docker compose logs -f --timestamps redis-broker

# Run default tests (unit + local integration, no containers, no external calls)
test:
	go test ./...

# Run integration tests using testcontainers (slower, requires Docker)
test-lazy:
	go test -tags=lazy ./...

# Run email integration tests (calls Resend API, requires valid credentials)
test-email:
	go test -tags=email ./...

test-all:
	go test -tags=email,lazy ./...

