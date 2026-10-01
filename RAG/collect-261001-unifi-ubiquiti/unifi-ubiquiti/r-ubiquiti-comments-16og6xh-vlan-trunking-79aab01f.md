---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-16og6xh-vlan-trunking-79aab01f
title: "r-ubiquiti-comments-16og6xh-vlan-trunking-79aab01f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-16og6xh-vlan-trunking-79aab01f.md
source_anchor: ""
source_lines: [1, 28]
sha256: 4e056ae3142136b2bcb85ae3deff4f25360b10de4a8299f2f665d22840a11db0
---

# r-ubiquiti-comments-16og6xh-vlan-trunking-79aab01f

VLAN Trunking 
        
        
        
    
    
    Hey all,
I'm relatively new to the Unifi eco system. I integrated a UAP into my current network and the dashboard has started pulling me more and more into wanting more Unifi. So I've now got myself a US8-150w that I would like to provision and replace my current switch. I am use a third party gateway (pfSense)
Currently, clients are connected to the native LAN (default). However, I do have a Secure VLAN that is tagged to it's own SSID on the UAP, and this works as expected. As with tagging/trunking any VLANs on my own switch.
I'd like to replicate this on the US8-150. However, while I've worked out potentially doing access ports I've yet to fully figure out how to trunk multiple.
I've attached a screenshot to which I think it should be done, but not entirely sure.
Thanks in advance!
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Trunk in UniFi is “none”
I am also relatively new to Ubiquity and running pfSense, and here is the issue I ran into the minute I connected my main switch to pfSense (which happens to be my DHCP server):
My UniFi network manager lost contact with the switch, saw the switch again as not managed and asked if I wanted to adopt it.
The issue is that the switch now is looking for “default” (VLAN 1) on the trunk line and pfSense doesn’t use VLAN 1 ever.
The solution:
Move the switch back to the configuration you had when it could talk to your UniFi manager
In the settings for the switch find the “Management VLAN” setting (I think it was Management Network on the browser based interface but am currently on mobile and it is called VLAN there)
Set that explicitly to the VLAN you want it on.
Now move the switch cable back to the trucked ports on both ends.
You got it. In your example 2 VLANs will pass (VLAN 1/default VLAN, and VLAN 1337) the port, VLAN1 passed untagged, and VLAN1337 passed tagged.
