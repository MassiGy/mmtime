#!/bin/bash

if [ ! -f  $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/logs ]; then
    echo "Creating $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/logs file";

    echo "
# This file contains the logs of your usage.
" > $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/logs;
fi