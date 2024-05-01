# How will we integrate mmtime to the userspace ? 

We will use yad to create a simple gtk based systray icon, that the user can use to interact with the 
mmtime deamon that is running in the background.

# Why yad ? 

I was planning to use alltray since it was simpler but it seems to be that alltray is not available in the 
debian repos anymore, but yad is. Thus, and since I am using a debian instance on my machine, I will go with 
yad.

# How will I use yad ?

We will use yad to create a small systray appelt, this applet will be it own script and it will be called
${BINARY_NAME}-applet, so for us it will be something like: mmtime-applet. 

This script will be a bash script that setups the yad command/applet and we will add it to the user shared 
directroy under: ~/.loacl/share/{BINARY_NAME}-{VERSION}/
- Maybe make the applet a symlink too to be consistent with what we've did with the mmtime binary. 
(NOTE): we will also add a small text file named about in the same directroy in which we'll define the author
and contact email adresse and some other useful info.

Basically, in this script we need to setup the applet is such a way so as the user can : 
- See a small tooltip text on hover that says "${BINARY_NAME}-applet".
- See the about file content by clicking directly (single left click) on the applet.
- See a 3 item menu by clicking with a single right click on the applet. This menu will contain an option to:
    * See the current stats/logs/usage data.
    * Reload the config, in case the targets file under .config has changed. 
    * Quit and kill the applet ( maybe the deamon too )


So, about the implementation, the first option will basically read the logs file under the shared directroy. 
The tricky part is to be sure that yad reload the content every time that menu item is selected, and that can
be done by re-reading the same file. Of course the file will not be read entirely, since it is a log file, only 
the last part of it are what matters for the user. ( last log cycle results ). These few lines can be read 
using the tail command with the --lines argument set to the exact lenght of the targets file under the .config
directroy.

The second option will be used to send a SIGUSR1 signal to the deamon process so as this one can reload the
targets file to monitor the newly added tasks. Also, to be sure and to give a visual effect to the user, we 
can kill the applet and relaunch it back on so as the user gets the lastest logs but also, by doing so, he/she 
will notice that the applet had disappeared and re-appeared again, giving a true reload feel to it.


For the third option, it is simply a kill command on top of a pgrep command to get the pid of the deamon 
process and the applet. ( See if we need to kill the deamon process too ) 


# The en-countred problems ? 

- How to reload the text once the yad notification applet is already drawn ? 
    Some solutions: 
    * You can either use a pipe to poll text from it and diplay the yad window content.  
    ``
        One simple thing you can do is use a --text-info dialog with a --tail option and pipe text into it.

        Using a named pipe like Victor does in his example just adds additional flexibility as to where the text comes from.

        As new input comes into yad, it will fill up the dialog box and then add a scrollbar when more is added.
        The --tail flag will cause the dialog to automatically scroll so that the latest info is displayed.
    ``

    * You can pipe the content to it as an input stream. 
    ``
        https://sourceforge.net/p/yad-dialog/wiki/LogViewer/
    ``

    * You can open a terminal instead of a yad window, in this terminal you can use the watch command to update
    the displayed content ( this is rather heavy but that was my intent in the first place with alltray )


- How to kill the parent yad dialog through a menu item action ? 
    Some solutions:
    * Killall instances of yad. This will of course kill all the other applications that use yad.
    * Figure out how to use the YAD_PID; 
    ``
        https://groups.google.com/g/yad-common/c/FPmSpoizb6M
    ``
    * Figure out a way to retreive the pid through a shell command, but this is really annoying since it gets quiet tidious with
    all the scripts embedding inside of quoats.


