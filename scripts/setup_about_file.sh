#!/bin/bash

if [ ! -f  $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/about ]; then
    echo "Creating $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/about file";

    echo "
<big>

<i><u>@Application</u></i>: $(cat ./BINARY_NAME).
<i><u>@Version</u></i>: $(cat ./VERSION)
<i><u>@Author</u></i>: Massiles Ghernaout (github.com/MassiGy).

<i><u>@Description</u></i>: $(cat ./BINARY_NAME) is a simple tool that allow the user to specify a list of applications/tasks to monitor, thus the user can see how much time he/she spends on them.

<i><u>@Help</u></i>: [Read the README.md file | Mail : ghernaoutmassi@gmail.com ]
</big>

" > $HOME/.local/share/$(cat ./BINARY_NAME)-$(cat ./VERSION)/about;
fi