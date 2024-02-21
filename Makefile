.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -o api-gateway.o ./cmd/api-gateway/api-gateway.go

.PHONY: gen/proto
gen/proto:
	protoc --go_out=. --go_opt=paths=import \
              --go-grpc_out=. --go-grpc_opt=paths=import \
              ./api/*.proto


.PHONY: gen
gen: gen/proto
	go generate ./...

.PHONY: test
test:
	go test ./... -v -coverpkg=./internal/... -coverprofile=coverage.out

.PHONY: testshort
testshort:
	go test ./... -v -test.short=true -coverpkg=./internal/... -coverprofile=coverage.out

.PHONY: c
c:
	go tool cover -func coverage.out

.PHONY: tc
tc: testshort c

.PHONY: lines
lines:
	git ls-files | xargs wc -l

.PHONY: dock
dock:
	docker build . --tag="vladmarlo/hell_kitchen_gateway:latest"

.PHONY: dock/push
dock/push:
	docker push vladmarlo/hell_kitchen_gateway:latest