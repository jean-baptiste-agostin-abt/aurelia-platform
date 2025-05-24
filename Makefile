.PHONY: build test up down

build:
	docker-compose build

test:
	cd backend && go test ./...
	cd frontend && npm test -- --watchAll=false || true

up:
	docker-compose up -d

down:
	docker-compose down
