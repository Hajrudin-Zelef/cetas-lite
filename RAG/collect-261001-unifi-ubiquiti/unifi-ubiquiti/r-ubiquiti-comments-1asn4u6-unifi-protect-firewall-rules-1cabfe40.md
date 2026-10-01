---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1asn4u6-unifi-protect-firewall-rules-1cabfe40
title: "r-ubiquiti-comments-1asn4u6-unifi-protect-firewall-rules-1cabfe40"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1asn4u6-unifi-protect-firewall-rules-1cabfe40.md
source_anchor: ""
source_lines: [1, 21]
sha256: cfd2e24c7a1fe142b72260171d8131450c37369cb698611b02514d546cb6086a
---

# r-ubiquiti-comments-1asn4u6-unifi-protect-firewall-rules-1cabfe40

UniFi Protect Firewall Rules 
        
        
        
    
    
    Hello, I am very new to this, so please take it easy on me.
I recently set up a UniFi Protect install for a funeral home after upgrading their cameras, when my youngest brother 20yr passed and we were there making arrangements I started talking to the Funeral Director about his current set up and offered to install UniFi Protect. assigned the port to the camera and isolated the ports.
Would somebody be willing to post a list of firewall rules that are recommended to secure this install
I haven’t been able to find a clear list that I am able to follow on how I need to create the firewall rules. I do have the cameras on their own vlan, and on the 24 port switch I assigned the port to the camera B land and isolated the ports.
Would somebody be willing to post a list of firewall rules that are recommended to secure this install?
UDM Pro and 24 Port POE UniFi switch with G4 bullet cameras and 1 AI 360
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Do you mean firewall rules to view from outside the network? It shouldn’t need any special port forwarding if you have remote access enabled.
Yea to be able to view it from the protect app or a web browser but make sure the cameras cannot talk to the internet or my other vlans
As long as the cameras can reach the UDM and the UDM can access the internet, the app will work fine.
