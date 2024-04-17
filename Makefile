clean:
	rm -rf bin/* 

compile:
	go build -o bin/mmtime cmd/main.go

runbin:
	./bin/mmtime

run:
	go run cmd/main.go

test:
	echo "(Makefile) tests are not setup yet"




clean_directories: 
	rm -rf ${HOME}/.config/mmtime
	rm -rf ${HOME}/.local/share/mmtime 

create_directories: 
	mkdir ${HOME}/.config/mmtime 2>/dev/null
	mkdir ${HOME}/.local/share/mmtime 2>/dev/null

make_bin_global: 
	mv ./bin/mmtime /usr/local/bin 

launch: 
	/usr/local/bin/mmtime 

build: clean compile clean_directories create_directories 

install: make_bin_global launch

	


