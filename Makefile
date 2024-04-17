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

create_directories: 
	# create the needed directories
	# file creation is taken care of by the program itself
	mkdir ~/.config/mmtime &&
	mkdir ~/.local/share/mmtime &&

make_bin_global: 
	mv ./bin/mmtime /usr/local/bin 

install: clean build create_directories make_bin_global

launch: /usr/local/bin/mmtime 

	


