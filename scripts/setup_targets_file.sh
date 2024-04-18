#!/bin/bash

if [ ! -f  $HOME/.config/$(cat ./BINARY_NAME)-$(cat ./VERSION)/targets ]; then
    echo "Creating $HOME/.config/$(cat ./BINARY_NAME)-$(cat ./VERSION)/targets file";

    echo "
# Add the applications that you want to track your usage 
# time in, each application in a seperated line.

$(cat ./BINARY_NAME)
" > $HOME/.config/$(cat ./BINARY_NAME)-$(cat ./VERSION)/targets;
fi