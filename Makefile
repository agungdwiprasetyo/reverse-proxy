.PHONY : build test cover

TEST_PACKAGES = ./helper

build:
	go build -o bin

test:
	$(foreach pkg, $(TEST_PACKAGES),\
	go test $(pkg);)

cover:
	if [ -f coverage.txt ]; then rm coverage.txt; fi;
	$(foreach pkg, $(TEST_PACKAGES), \
	go test -coverprofile=coverage.out -covermode=atomic $(pkg); \
	tail -n +2 coverage.out >> coverage.txt;)
	rm coverage.out
