.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -o api-gateway.o ./cmd/api-gateway/api-gateway.go

.PHONY: gen
gen:
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