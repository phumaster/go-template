dev:
	cp env/dev.env .env && go run main.go
build:
	cp env/prod.env .env && go build
linux:
	cp env/dev.env .env && GOOS=linux go build
up:
	docker-compose up
down:
	docker-compose down
