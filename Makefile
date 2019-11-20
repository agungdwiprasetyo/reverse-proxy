.PHONY : build test docker

PACKAGES = $(shell go list ./... | grep -v -e . | tr '\n' ',')
GCP_PROJECT_ID = mantul-tenan

build:
	go build -o bin

docker: build
	docker build -t reverse-proxy:latest .

kubernetes:
	docker tag reverse-proxy:latest gcr.io/$(GCP_PROJECT_ID)/reverse-proxy:latest
	docker push gcr.io/$(GCP_PROJECT_ID)/reverse-proxy:latest
	kubectl set image deployment/reverse-proxy reverse-proxy-sha256=gcr.io/$(GCP_PROJECT_ID)/reverse-proxy:latest

test: build
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=set -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.txt

