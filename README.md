# Monitor My Time (mmtime)

This CLI tool aims at offering a deamon-like process that works in the background and monitors the
time that the user had spent on a collection of applications.

This collection of applications can be defined in the `~/.config/mmtime-vX.X/targets` file, where every
application, that the user wants to track the amount of time he/she spend in, is listed in a sperated
time. i.e

```
    # My target apps.

    Discord 
    Code 
    Slack 
    Mpv 
    Firefox 
    Vim 

```


The `mmtime` tool will then track all the time spent on these applications and write that to a 
db file (`~/.local/share/mmtime-vX.X/targets.stats.db`).

Besides, there will be a system tray applet that will allow the user to see the current monitoring related stats, this will give the user the ability to see the stats live and also through a menu item reload the configuration file (`targets` file).

## Setup & Installation ( Using Go tools )

```sh

    # make sure that you have xterm installed (for some use these commands)
    [[ "$HOSTNAME" == "debian" ]] && sudo apt install xterm
    [[ "$HOSTNAME" == "ubuntu" ]] && sudo apt install xterm
    [[ "$HOSTNAME" == "fedora" ]] && sudo dnf install xterm


    # build the binary (this requires Go to be installed )
    make build

    # setup the environment and mv the binary to ~/.local/share
    make install

    # (Optional) you can create a symlink to the binary
    # to /usr/local/bin, or directly copy the binary to it.

```

## Setup & Installation ( Without Go tools )

First, download from Github the latest release. 

```sh

    # clone the repo
    git clone <repo_url>

    # download from github the latest release

    # move the binary of the release to the project /bin directory

    # run make install 
    make install

```


## Uninstall 

```sh

    # just run the following command to remove all the used
    # directroies and the ~/.local/share binary.
    make uninstall

    # (NOTE) extra setup from your part should be undone by
    # yourself.

```




