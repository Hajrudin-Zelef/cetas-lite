---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-10vo9xm-adoption-failed-for-all-my-devices-f991f08b
title: "r-ubiquiti-comments-10vo9xm-adoption-failed-for-all-my-devices-f991f08b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-10vo9xm-adoption-failed-for-all-my-devices-f991f08b.md
source_anchor: ""
source_lines: [1, 35]
sha256: 84a05af6101cc34667032f4132a46399183426f7f362348a308251740fa8caa8
---

# r-ubiquiti-comments-10vo9xm-adoption-failed-for-all-my-devices-f991f08b

adoption failed for all my devices 
        
        
        
    
    
    After things have been running fairly smoothly for a while, today my internet started acting up. Since I had the time I logged into the network application and saw that all my devices are listed with a status of adoption failed.
I tried upgrading to the latest network application version of Unifi Network Application 7.4.140
I tried the Advanced adoption option
but I still fail to adopt.
the devices haven't been previously managed anywhere else, just from the one Mac computer, which also ran its own update in the last few days. Log files are downloaded as *.supp, not sure how to open those
Is factory reset the last option ? will I lose any of my settings by doing that or are all settings saved on the Network Application on my computer ?
anyone else seeing this issue, where status of all the devices became adoption failed
Edit: I also restored config file to an older version from December which was auto backed up and still same issue.Edit 2: Managed to roll back the NanoHD APs to version 6.0.21 and they are working now. I had to reset them, and push the update through SSH for it to work. Ofcourse now network application asks for me to update them, but I bet that will break them again. Next step is hard reboot the other devices and see if it works again or I have to roll them back too.
Edit 3: ended up resetting all my devices, removing them and then everything worked again. not fun
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Try just power failing each device and see if it connects.
You can also try to ssh into the device and reset the host and see if it connects.
I can't connect through SSH either, I'm getting connection refused error using the default MacOS terminal tool. I tried turning on and off the SSH Authentication. but still same error.
when you say power fail each device do you mean a regular restart from the touch screen on the switch or pulling the plug ?
Pull the plug so that you do a hard reboot.
Hi u/fsolo23 We apologize for the inconvenience. Please share more info and any related support tickets here so we can properly escalate and assist: http://community.ui.com/social-feedback
Issue resolved now see posts above. But I did file through the link provided.
All settings are saved in Network App, so you can reset the AP. That's the beauty of Unifi
What would be beautiful is an actual solution to the problem.
not randomly failing to adopt would be nice :)
I just had pretty much the same thing happen to me, I'm assuming after the 1.12.38 upgrade. I had my network drop out during work today and when I took a look at the devices two of them are "Adoption Failed".
This is on a UDM Pro. The two devices are a
USW-Aggregationswitch and aUSW-24-G1
I'm going to open a case with Ubiquiti using the link in the comments below but wanted to report here too in case other experience the same thing.
