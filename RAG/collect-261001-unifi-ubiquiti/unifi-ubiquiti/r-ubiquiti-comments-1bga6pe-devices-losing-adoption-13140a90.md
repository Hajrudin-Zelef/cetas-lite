---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1bga6pe-devices-losing-adoption-13140a90
title: "r-ubiquiti-comments-1bga6pe-devices-losing-adoption-13140a90"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1bga6pe-devices-losing-adoption-13140a90.md
source_anchor: ""
source_lines: [1, 32]
sha256: 92efa8f767006768fbe36309218d3b1284b9b6257f6606dc0370bb624964fa25
---

# r-ubiquiti-comments-1bga6pe-devices-losing-adoption-13140a90

Devices losing adoption 
        
        
        
    
    
    Hi, I have a small domestic set up of a cloud key, USG, a couple of switches, a couple of cameras and a wifi point. Every now and again, my UniFi app tells me that it looks like some of my devices have been previously adopted by another console. I re-adopt them, sometimes it's fine, sometimes it fails for some devices and they need resetting. Any ideas why this may be happening? Also, why do my cameras only occasionally appear in this view?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
I’ve had some devices do odd things… Maybe they wouldn’t adopt or maybe they’d lose adoption after they rebooted. Solution that’s always worked 💯% of the time for me is to factory/full reset the device and re-adopt.
Login with ssh via putty and run command "set-default" , then readopt
If they're DHCP make sure they still have / can still get IPs.
In my setup that is running for ages, I had version 8.0.28 with 1x AC-Lite and 2x AC-Pro since beginning of this year when out of nowhere this week the AC-Lite started readopting /rebooting.
-> My mentioned AC-Lite is re-adopting in a loop over and over again. Uptime max. 2 minutes.
-> Once dismissed from UniFi Controller, I can SSH into it and watch uptime going up for hours.
-> Once adopting to UniFi controller again, everything is fucked up again.
Is there any known issue with Unifi controller 8.0.28 and 8.1.113?
All three APs were running firmware 6.6.55.15189. AC-Pros rock stabil as always.
Replaced cables and uplink port at switch.
Downgraded AC lite to 2 different firmware versions 6.5.62 and 6.5.54.
Upgraded Unifi from 8.0.28 to 8.1.113
Did factory resets.
Defined IP of controller for adoption instead of FQDN.
AC-Lite is running perfectly fine w/o reboots when reset. So, no PoE injector issue, too.
(Background: Used to be Network Admin operating 25+ locations with Firewalls/LAN/WAN/VoiP until 10y ago. Running UniFi since GEN1 APs)
Update: The AC-lite is working just fine if I do not assign any SSID to it. Effectively adopted and just idling w/o clients connected.
Ghosts 👻
They wanted emancipation
