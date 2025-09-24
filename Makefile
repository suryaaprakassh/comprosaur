build:
	 @ go build -o bin/comprosaur .

run: build
	 @ ./bin/comprosaur

test:
	@ gotestsum --format testname 

clean:
	@ rm -rf *.zip comp.log

.PHONY: test run build
