clean:
	rm -rf bin/* 

build:
	go build -o bin/mmtime cmd/main.go

runbin:
	./bin/mmtime

run:
	go run cmd/main.go

test:
	echo "(Makefile) tests are not setup yet"
