LOCAL_BIN=./bin

.PHONY: up
up:
	docker-compose --file ./build/docker-compose.yml up -d --remove-orphans

.PHONY: down
down:
	docker-compose --file ./build/docker-compose.yml down

.PHONY: dev-up
dev-up:
	docker-compose --file ./build/docker-compose.yml up -d --build --remove-orphans

.PHONY: dev-down
dev-down:
	docker-compose --file ./build/docker-compose.yml down --rmi all -v

.PHONY: run
run:
	CGO_ENABLED=0 go build -ldflags='-w -s' -o $(LOCAL_BIN)/app ./cmd/xlsx-builder/main.go && HTTP_ADDR=:8080 $(LOCAL_BIN)/app

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v -shuffle=on -count=2 -short -cover ./...

.PHONY: test-race
test-race:
	go test -race ./...

.PHONY: clean-bin
clean-bin:
	rm -fr $(LOCAL_BIN)
