---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-s9gi99-multiple-devices-changed-to-adoption-failed-e62e7613
title: "r-ubiquiti-comments-s9gi99-multiple-devices-changed-to-adoption-failed-e62e7613"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-s9gi99-multiple-devices-changed-to-adoption-failed-e62e7613.md
source_anchor: ""
source_lines: [1, 40]
sha256: a95751a4c9a8b12cc3a127090b05085c899c49dddd9a93c064f3ffcd6976a3fc
---

# r-ubiquiti-comments-s9gi99-multiple-devices-changed-to-adoption-failed-e62e7613

Multiple devices changed to “Adoption Failed” randomly 
        
        
        
    
    
    I’ve been using Unifi for months with no issues. This morning, I logged into the CloudKey Gen 2 to view WiFi Status, and noticed that my AP and 2 of the 4 switches were listed as “Adoption Failed”, but still seem to be working. The other 2 switches are up and show connected
All devices are reachable from the CloudKey via ping, and it’s asking for an inform password I don’t have.
How do I get the controller to talk to the APs and switches again?
Edit: Devices/Firmware
1x UniFi CloudKey Gen 2 - 6.5.55
1x UAP-AC-PRO - 5.43.43
2x US-8-60w - 5.64
2x USW-Flex-Mini - 1.84
The two USW-Flex-Minis show up as connected, the rest Adoption Failed
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
I was able to somewhat resolve the issue, but there are some lingering questions/issues.
Sometime last night, the DNS resolver on my firewall died and did not restart.
As soon as I restarted the DNS resolver service on my firewall, all the devices came back up as connected. Every device does have an A record, but I don’t understand why the UniFi controller is using DNS for resolving devices. The controller knows the IP address of each device, and I don’t think I used the host names when setting up devices (just clicked adopt IIRC).
Is there a config file I can check to see if the CloudKey is relying on DNS? If it’s relying on DNS names for each device, how do I change it to use the IP address?
Glad to hear you fixed it.
The devices set-inform would be what is using DNS not the cloud key. If your set-inform was using hostname and your DNS was down as you say the devices were never checking in to the controller.
You can issue the set-inform command on the devices to change (there's also a default inform inside of your site which can override the inform parameter). Your better solution would be to spin up a 2nd DNS controller that can be a backup for your primary so if the resolver service on your firewall goes down the secondary will handle it until you get it back up.
Devices and Firmware
Controller version ...
Added to the original post
UAP Firmware 5.43.51
https://community.ui.com/releases/UAP-Firmware-5-43-51/996b2db2-ec78-4774-a2de-2c9b34ef041a
[HW] Avoid intermittent failed readoption loops.
Still digging through the FW notes for the US-8 but you're behind a couple of versions on those. Generally on the inform password bug a factory default and readopt resolves, but since there are FW notes I'd go that route.
Keep in mind if you set-default from SSH traffic will drop as the device reboots, and drop again after readopt.
Edit: FW for US-8-60W
Most people are saying that 5.43.36 or 5.64.8 are the releases that are most stable right now. 5.64.8 MAY be what you're running but it got truncated off on your post edit. If so, you may want to revert down to 5.43.
Not saying that this is the case, but maybe something to look at.
5.76.7 people are saying has spontaneous reboots in some instances so I wouldn't go to that one.
