APP_NAME?=leetCodeProblems

clean:
	rm -f &{APP_NAME}

build:	clean
	go build -o &{APP_NAME}

test:
	go test -v -count=1 ./...

test10:
	go test -v -count=20 ./...

race:
	go test -v -race -count=1 ./...

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm .\coverage.out
