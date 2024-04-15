# Monitor My Time (mmtime)

This CLI tool aims at offering a deamon-like process that works in the background and monitors the
time that the user had spent on a collection of applications.

This collection of applications can be defined in the ~/.config/mmtime/targets file, where every
application that the user wants to track the amount of time he/she spend in is listed in a sperated
time. i.e

`` Discord Code Slack Mpv Firefox Vim ``


The mmtime tool will then track all the time spent on these applications and write that to a sqlite
db file.

Besides, while mmtime is running in the background, the user can go to thier web browser and check
thier current stats by visiting the localhost:9999 address.





