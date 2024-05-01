BINARY_NAME:=$(shell cat ./BINARY_NAME)
VERSION:=$(shell cat ./VERSION)
CONFIG_DIR:=${HOME}/.config/${BINARY_NAME}-${VERSION}
SHARED_DIR:=${HOME}/.local/share/${BINARY_NAME}-${VERSION}


clean:
	rm -rf bin/* 

compile:
	go build -o bin/${BINARY_NAME} -ldflags "-X main.version=${VERSION} -X main.binary_name=${BINARY_NAME}" cmd/main.go

runbin:
	./bin/${BINARY_NAME}

run:
	go run cmd/main.go

test:
	@echo "(Makefile) tests are not setup yet";


make_bin_shared: 
	cp $(shell pwd)/bin/${BINARY_NAME} ${SHARED_DIR}/${BINARY_NAME};
	@echo "The ${BINARY_NAME} binary file can be found in ${SHARED_DIR}";
	@echo ""
	bash ./scripts/setup_aliases.sh;
	@echo "Added ${BINARY_NAME} and ${BINARY_NAME}-applet aliases to .bashrc|.bash_aliases|.zshrc|.zsh_aliases";
	@echo "you can run the program by using this command: ${BINARY_NAME} and ${BINARY_NAME}-applet "

launch: 
	${SHARED_DIR}/${BINARY_NAME}

build: clean compile

buildprod: 
	rm -rf bin/*
	@echo "Building a static executable..."
	CGO_ENABLED=0 go build -a -tags netgo,osusergo -ldflags "-X main.version=${VERSION} -X main.binary_name=${BINARY_NAME} -extldflags '-static -s -w'" -o bin/${BINARY_NAME} cmd/main.go 


setup: 
	@echo ""
	@echo "Setting up the config and local shared directories, and the appropriate files."
	bash ./scripts/setup_config_dir.sh 
	bash ./scripts/setup_targets_file.sh 

	bash ./scripts/setup_shared_dir.sh
	bash ./scripts/setup_db_file.sh
	bash ./scripts/setup_about_file.sh
	bash ./scripts/setup_log_file.sh
	bash ./scripts/setup_applet_file.sh

install: setup make_bin_shared  

rm_local_bin: 
	rm -rf ${SHARED_DIR}/${BINARY_NAME}

uninstall: rm_local_bin 
	rm -rf ${CONFIG_DIR}
	rm -rf ${SHARED_DIR}




	


