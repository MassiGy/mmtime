BINARY_NAME:=$(shell cat ./BINARY_NAME)
VERSION:=$(shell cat ./VERSION)
CONFIG_DIR:=${HOME}/.config/${BINARY_NAME}-${VERSION}
SHARED_DIR:=${HOME}/.local/share/${BINARY_NAME}-${VERSION}


clean:
	rm -rf bin/* 

compile:
	go build -o bin/${BINARY_NAME} cmd/main.go

runbin:
	./bin/${BINARY_NAME}

run:
	go run cmd/main.go

test:
	echo "(Makefile) tests are not setup yet";


make_bin_shared: 
	ln -s $(shell pwd)/bin/${BINARY_NAME} ${SHARED_DIR}/${BINARY_NAME};
	echo "The ${BINARY_NAME} binary file can be found in ${SHARED_DIR}";

	bash ./scripts/setup_aliases.sh;
	echo "Added ${BINARY_NAME} aliases to .bashrc|.bash_aliases|.zshrc|.zsh_aliases";

launch: 
	${SHARED_DIR}/${BINARY_NAME}

build: clean compile 

setup: 
	bash ./scripts/setup_config_dir.sh 
	bash ./scripts/setup_targets_file.sh 


	bash ./scripts/setup_shared_dir.sh
	bash ./scripts/setup_db_file.sh
	bash ./scripts/setup_log_file.sh

install: setup make_bin_shared launch 

rm_local_bin: 
	rm -rf ${SHARED_DIR}/${BINARY_NAME}




uninstall: rm_local_bin 
	rm -rf ${CONFIG_DIR}
	rm -rf ${SHARED_DIR}




	


