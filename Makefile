build:
	 @ go build -o bin/comprosaur .

run: build
	 @ ./bin/comprosaur

test:
	@ go test -v ./... | grep -v "no test files"

.PHONY: test run build
