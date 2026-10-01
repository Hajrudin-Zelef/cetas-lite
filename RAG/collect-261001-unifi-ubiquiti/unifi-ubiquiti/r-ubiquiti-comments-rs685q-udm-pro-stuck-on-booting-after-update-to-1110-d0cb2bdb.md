---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-rs685q-udm-pro-stuck-on-booting-after-update-to-1110-d0cb2bdb
title: "r-ubiquiti-comments-rs685q-udm-pro-stuck-on-booting-after-update-to-1110-d0cb2bdb"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-rs685q-udm-pro-stuck-on-booting-after-update-to-1110-d0cb2bdb.md
source_anchor: ""
source_lines: [1, 53]
sha256: 2ac0a9023033a9595bc0db343cfa9ecca88cd282647a1065f466cea6bbcb16bd
---

# r-ubiquiti-comments-rs685q-udm-pro-stuck-on-booting-after-update-to-1110-d0cb2bdb

UDM Pro Stuck on booting after Update to 1.11.0 
        
        
        
    
    
    Hey, I recently upgraded my home network to a UDM Pro. I also added a USW-Pro-24-POE, three USW-Flex, two U6 LR and a Lite. Until the week before last, everything worked without problems.
Now the update 1.11.0 for the UDM Pro came out and it was unfortunately still on auto update and has installed it.
My problem now is that I can only get to the dashboard page of the UDM Pro and in the network area it says "Getting Ready...". Rebooting via the web interface or pulling the power cable didn't help. The UDM Pro display says "Unifi OS is starting...". Once there was also an error message that something went wrong and I should go to udm.ui/recovery, but that resolved itself. Protect is running
Until the first reboot, I could still access the UDM with the app, but unfortunately that is no longer possible. It gets no connection, but still shows me my devices.
Curiously, I still have Internet connection and new devices also get an IP address assigned via DHCP.
The other Unifi devices work so far without problems.
Has anyone ever had such a problem or can tell me how to solve it?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Hey mate, i had the same problem until 10 minutes ago.
What worked for me was :
- ssh into the udm
- unifi-os shell
- apt update
- apt upgrade
I didnt even had to restart the UDM. It worked instantly.
Hope it helps!
WOW Dude...that was suprisingly....unspectacular...BUT IT WORKED!
Now i have Network version 6.5.55 instead of 6.5.54 installed :D
Thanks a lot!
Thank you so much. 3 months later and this worked for me as well.
I updated & got this same issue, i updated but it is still showing the same offline, should i now reboot the network? what we are trying this evening.
Thanks for any insight
i just got a problem on a Cloud key showing offline, i did the same thing as for the UDM exept the unifi-os shell part because well when you ssh a key it's directly into the shell.
So it did not work at all untill i reboot the key, once rebooted it instantly showed online
if it helps, even a month later.. :)
Now after some more testing and searching for the problem (I didn't find the exact problem, but similar ones) I tried the following:
unifi-os restart via SSH
I can still access the UDM via SSH and entered the command "unifi-os restart", it did that but without success
2. update the unifi firmware again
Using the command "ubnt-upgrade https://fw-download.ubnt.com/data/udm/d530-udmpro-1.11.0-69dd1244f5e94a8ea72b246822fe8fb2.bin" to install the UDM firmware again, did not bring any improvement.
3. update of the network controller
Here I used the "unifi-os shell" to go to the Linux interface and loaded the update file for 6.5.54 in the folder /tmp with this command "curl -o "unifi_sysvinit_all.deb" https://dl.ui.com/unifi/6.5.54/unifi_sysvinit_all.deb" and installed it with "dpkg -i unifi_sysvinit_all.deb".
It then tried for a while to install the update and then aborted with an undefined error
What I can also say is that the CPU usage is between 25-50%, temperature is 44.3° C and memory is 30% (1.25/4.04 GB).
so far to the update, I'll research further
Did you back it up before the update then try a factory reset?
Unfortunately I don't have a backup from before the auto update. I thought I had the cloud backup set up, but unfortunately nothing is saved there.
If I have to reset the UDM to factory settings, I see it as a lesson to check the auto backup as well....
i assume you have pulled the power and let it sit for a minute then turned it back on?
Hi u/Platzhirsch90 Thanks for your patience. Please start a LiveChat or a support ticket at help.ui.com so our team can collect more information to properly review and assist.
Y’all really need to get a phone number for tech support.
I had a very similar issue on a previous firmware and UI-Glenn did a remote SSH session and was able to fix it somehow. I had to step away so I’m not entirely sure but they were able to do it without a full restore.
