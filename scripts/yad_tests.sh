#! /bin/bash

# SOURCE : https://funprojects.blog/2022/06/12/apps-and-scripts-in-the-system-tray/

RECENT_LOG_LINES_COUNT=$(wc --lines ~/.config/mmtime-v0.1/targets | awk '{print $1}')
ABOUT_CONTENT=$(cat ~/.local/share/mmtime-v0.1/about)


yad --notification --image="appointment-new-symbolic" \
 --command="yad --text='$ABOUT_CONTENT' --title='mmtime' --geometry=600x100+600+200 --button=OK" \
 --menu="See Stats! xterm -bg white -fg black -hold -geometry 100x30+600+200 -fa Monospace -fs 10 -T 'mmtime logs' -e 'watch tail --lines $RECENT_LOG_LINES_COUNT ~/.local/share/mmtime-v0.1/logs' \
       |Reload Config ! bash -c 'kill -SIGUSR1 $(pgrep mmtime) && killall yad && ./yad_tests.sh' \
       |Quit ! killall yad " \
 --text="mmtime applet"