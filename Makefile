build:
	 @ go build -o bin/comprosaur .

run: build
	 @ ./bin/comprosaur

test:
	@ gotestsum --format testname 

.PHONY: test run build
