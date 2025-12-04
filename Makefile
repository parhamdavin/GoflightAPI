.PHONY: run build clean migrate up down logs test docker-build docker-run

APP_NAME=airline-reservation
DOCKER_IMAGE=airline-reservation-app
DB_URL="host=localhost user=postgres password=secret dbname=airline sslmode=disable"
PORT=8080

run:
	go run main.go

build:
	go build -o $(APP_NAME) main.go

clean:
	rm -f $(APP_NAME)

migrate:
	migrate -database "$(DB_URL)" -path migrations up

up:
	docker-compose up --build -d

down:
	docker-compose down

logs:
	docker-compose logs -f

test:
	go test ./...

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run -p $(PORT):$(PORT) --env DB_URL=$(DB_URL) $(DOCKER_IMAGE)
