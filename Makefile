build:
	@ go build -o bin/comprosaur .

run: build
	@ ./bin/comprosaur
