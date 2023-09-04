.PHONY: dev-up dev-down run test clean

dev-up:
	docker-compose --file ./build/docker-compose.yml up -d --build --remove-orphans

dev-down:
	docker-compose --file ./build/docker-compose.yml down --rmi all -v

run:
	CGO_ENABLED=0 go build -ldflags='-w -s' -o app ./main.go && HTTP_ADDR=:8080 ./app
	#go build -o app cmd/port-service/main.go && HTTP_ADDR=:8080 ./app

test:
	go test -v -shuffle=on -count=2 -short -cover ./...
	#go test -race ./...

clean:
	rm app
