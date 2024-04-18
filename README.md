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

Besides, while `mmtime` is running in the background, the user can go to thier web browser and check
thier current stats by visiting the `localhost:9999` address.

## Setup & Installation

```sh

    # build the binary 
    make build

    # setup the environment and mv the binary to ~/.local/share
    make install

    # (Optional) you can create a symlink to the binary
    # to /usr/local/bin, or directly copy the binary to it.

```

## Uninstall 

```sh

    # just run the following command to remove all the used
    # directroies and the ~/.local/share binary.
    make uninstall

    # (NOTE) extra setup from your part should be undone by
    # yourself.

```




