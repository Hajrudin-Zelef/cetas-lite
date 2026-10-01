---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-16dvepg-brand-new-udm-pro-from-2020-stuck-on-firmware-4770c59a
title: "r-ubiquiti-comments-16dvepg-brand-new-udm-pro-from-2020-stuck-on-firmware-4770c59a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-16dvepg-brand-new-udm-pro-from-2020-stuck-on-firmware-4770c59a.md
source_anchor: ""
source_lines: [1, 31]
sha256: 1b255ce5700dc29edf903f52e0188a396ef90dcf18b2d9c0909a327274565432
---

# r-ubiquiti-comments-16dvepg-brand-new-udm-pro-from-2020-stuck-on-firmware-4770c59a

Brand New UDM Pro from 2020 Stuck on Firmware Update 
        
        
        
    
    
    Hey all, I'm trying to do a firmware update on a brand new UDM Pro that came with 1.6.8 installed.
I used the UI through the PC I connected to click update available, but it's been about 2-3 hours and it's stuck on "Device Updating". Is it supposed to take this long? I did come across a post on the UI community forums, but no resolution was given.
I do have the UDM Pro connected to my Verizon Fios router just for setup without interrupting my family's useage. Could that have something to do with it? I realized the Verizon router and the UDM Pro use 192.168.1.1, but once I reverted to using DHCP in my PC's network adapter settings, the UDM Pro splash page came right up. No problems until I tried to update firmware.
Thanks for any help!
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
UPDATE: I sent Ubiquiti support a request at about 3AM and I had a good answer by 6:47AM that same day.
I did have to first upgrade firmware to 1.12.38. After doing that manually, I was able to use the web UI to update without issue up to the latest, step by step.
Below is what they sent me. Hope it helps.
I can see that UDM-Pro is running on a very older version (1.6.8 version), so it won't be possible to directly upgrade to the latest version. Please perform the update via recovery mode, here (hyperlinked) is an official document to follow.
https://fw-download.ui.com/data/udm/1adc-udmpro-1.12.38-ca8a490ac2b04247abb3f7d3e3eae01a.bin
Before entering into the recovery mode, use the above link for the 1.12.38 version and download the file and later on use that file to upload version.
First, upgrade it to the 1.12.38 version; after that you will be able to update it via web UI (1.12.38>>2.4.27>>2.5.17>>3.0.20>> 3.1.15); here are the steps, access the UDM via web application, go to UniFi OS >>application/update tab and apply the update.
Before upgrading, take a backup to be safer.
Try the above steps and let me know your feedback.
Dealt with same problem and solved with that response from support. My unit did go thru several automatic / critical updates that I did not initiate / could not stop.
I'm assuming you've tried to reboot first? Upgrades usually take no longer than 20 minutes... and even that is high.
With that being said, 1.6.8 is EXTREMELY old. You'll likely need to go to 1.12.38 first before going up the 2.x path to 3.x... I just haven't seen that many people with an OS so old.
If you don't get an answer, I'd post a support topic on community.ui.com and they can probably help with the upgrade path.
Often times the Network app gets hung up and doesn't refresh properly. First step I try when things look like they hang on update is to get out of the app and log back in. More often than not I will find that the update went fine.
You need to read the release notes and check which version you need before going to the next major version. Had the same issue on a friends one.. if you can get into the GUI, enable ssh as soon as possible if not already done and then do the update via ssh.. (manual update). Personally I also would not recommend leaving “auto update” enabled once you have it up and running again.. Check YouTube and google for more advise.
