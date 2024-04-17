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
	rm -rf /home/massigy/.config/mmtime
	rm -rf /home/massigy/.local/share/mmtime 

create_directories: 
	mkdir /home/massigy/.config/mmtime 2>/dev/null
	mkdir /home/massigy/.local/share/mmtime 2>/dev/null

setup_files_ownership: 
	chown -R massigy:massigy /home/massigy/.config/mmtime
	chown -R massigy:massigy /home/massigy/.local/share/mmtime


make_bin_global: 
	cp bin/mmtime /usr/local/bin/ 
	#ln -s /home/massigy/.local/share/mmtime/mmtime /usr/local/bin/mmtime 

launch: 
	mmtime &	

build: clean compile clean_directories create_directories 

install: make_bin_global launch setup_files_ownership

rm_global_bin:
	rm -rf /usr/local/bin/mmtime

rm_local_bin:
	rm -rf /home/massigy/.local/share/mmtime/mmtime


uninstall: rm_global_bin rm_local_bin clean_directories




	


