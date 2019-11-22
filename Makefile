.PHONY : build test docker deploy

PACKAGES = $(shell go list ./... | grep -v -e . | tr '\n' ',')
GCP_PROJECT_ID = mantul-tenan
APP_NAME = reverse-proxy
ts = $(shell date +%Y%m%d%H%M%S)
IMAGE_TAG = $(ts)

build:
	go build -o bin

docker: build
	docker build -t $(APP_NAME):latest .

deploy:
	docker build -t $(APP_NAME):$(IMAGE_TAG) .
	docker tag $(APP_NAME):$(IMAGE_TAG) gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME):$(IMAGE_TAG)
	docker push gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME):$(IMAGE_TAG)
	kubectl set image deployment/$(APP_NAME) $(APP_NAME)-sha256=gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME):$(IMAGE_TAG)

test: build
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=set -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.txt

