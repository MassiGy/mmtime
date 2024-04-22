#!/bin/bash

BINARY_NAME=$(cat ./BINARY_NAME)
VERSION=$(cat ./VERSION)


if [ ! -f  $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/${BINARY_NAME}-applet ]; then
    echo "Creating $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/${BINARY_NAME}-applet file";

    # cat the local information about the binary to the applet file.
    echo "
#!/bin/bash

BINARY_NAME=$BINARY_NAME
VERSION=$VERSION

" > $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/${BINARY_NAME}-applet;

    # cat the remaining code to the applet file.
    cat ./applet >> $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/${BINARY_NAME}-applet;

    # add execute permissions to the file.
    chmod +x $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/${BINARY_NAME}-applet;
fi




	# ln -s $(shell pwd)/applet ${SHARED_DIR}/${BINARY_NAME}-applet;
	# @echo "The ${BINARY_NAME}-applet  file can be found in ${SHARED_DIR}";
	# @echo ""