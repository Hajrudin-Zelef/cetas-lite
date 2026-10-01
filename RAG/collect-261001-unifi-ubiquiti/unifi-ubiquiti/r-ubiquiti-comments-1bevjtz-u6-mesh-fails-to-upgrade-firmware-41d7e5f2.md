---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1bevjtz-u6-mesh-fails-to-upgrade-firmware-41d7e5f2
title: "r-ubiquiti-comments-1bevjtz-u6-mesh-fails-to-upgrade-firmware-41d7e5f2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1bevjtz-u6-mesh-fails-to-upgrade-firmware-41d7e5f2.md
source_anchor: ""
source_lines: [1, 46]
sha256: 4fafb7a4d298909d9d58bbfa83ebb0fa309df4d1bc5effa45ae00a7f4ce970f7
---

# r-ubiquiti-comments-1bevjtz-u6-mesh-fails-to-upgrade-firmware-41d7e5f2

U6 Mesh fails to upgrade firmware 
        
        
        
    
    
    Apologies in advance for the long post, I am at a total loss here. I have followed many different posts to upgrade and nothing seems to work. I have 3 U6 Mesh APs, 2 upgrade just fine but a 3rd always fails. I have tried upgrading from the controller and command line without any luck.
Version on the affected AP is 6.6.58.
When upgrading from the Web UI I get the following message:
"u6-ap-west could not be updated due to an incorrect network configuration."
Messages received when trying to upgrade from SSH:
After performing a "set-default"
u6-ap-west-BZ.6.6.58# upgrade https://dl.ui.com/unifi/firmware/UAP6MP/6.6.56.152
00/BZ.ipq50xx_6.6.56+15200.231207.1909.bin
Downloading firmware from 'https://dl.ui.com/unifi/firmware/UAP6MP/6.6.56.15200/BZ.ipq50xx_6.6.56+15200.231207.1909.bin'.
Couldn't open image file: /tmp/fwupdate.bin!
Invalid firmware. Exit status for firmware check = 65024.
I have also scp'd the latest firmware to the AP and ran "syswrapper.sh upgrade2 &" and received the following:
u6-ap-west-BZ.6.6.58# syswrapper.sh upgrade2 &
u6-ap-west-BZ.6.6.58# syswrapper: [set_state] upgrading
Could not open MTD device /dev/mmcblk0p4
syswrapper: do_upgrade cfgmtd: 255
Could not open MTD device /dev/mmcblk0p4
ls: /tmp/run/persistent/: No such file or directory
ls: /tmp/run/read.cfg: No such file or directory
open(/dev/mmcblk0) failed: No such file or directory
open(/dev/mmcblk0p1) failed: No such file or directory
open(/dev/mmcblk0p2) failed: No such file or directory
open(/dev/mmcblk0p3) failed: No such file or directory
open(/dev/mmcblk0p4) failed: No such file or directory
open(/dev/mmcblk0p5) failed: No such file or directory
FW Image partition "kernel0"(6) ends at address, 0x02244400 outside the flash memory map. Valid range is 0x00000000-0x00400000.
syswrapper: [set_state] read
Any guidance is greatly appreciated!
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Just FYI... There's a different U6Mesh binary on the release page at https://community.ui.com/releases/UniFi-Access-Point-6-6-65/3df9b0c0-7e51-4843-93a2-276c2d63e10a
The U6Mesh binary there ends with .2328.bin
Thanks for the info. I tried that one too, no luck.
I submitted a support ticket. Now just waiting on Ubiquiti to review logs.
Hi, u/Senior_Case110.
That's not the expected experience our users should have. We would like to help you and get our team to check the support file of your's so that they can help you to find what is the root cause. We have reached out to you in the DMs with a request for your support ticket number. Thanks!
