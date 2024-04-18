#!/bin/bash

if [ ! -f  $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/targets.stats.db ]; then
    echo "Creating $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/targets.stats.db file";

    echo "
# For every application/process just sum up the usage
# times for any given date to get the total usage during that day.

Process     Date        Usage
" > $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/targets.stats.db;
fi