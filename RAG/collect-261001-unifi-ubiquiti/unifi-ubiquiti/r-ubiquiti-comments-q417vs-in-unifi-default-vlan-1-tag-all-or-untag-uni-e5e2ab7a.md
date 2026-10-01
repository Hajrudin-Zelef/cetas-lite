---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-q417vs-in-unifi-default-vlan-1-tag-all-or-untag-uni-e5e2ab7a
title: "r-ubiquiti-comments-q417vs-in-unifi-default-vlan-1-tag-all-or-untag-uni-e5e2ab7a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-q417vs-in-unifi-default-vlan-1-tag-all-or-untag-uni-e5e2ab7a.md
source_anchor: ""
source_lines: [1, 46]
sha256: 849a1482dad9998d53b4f956c7380df4b86302dd3055d09096968cec36057628
---

# r-ubiquiti-comments-q417vs-in-unifi-default-vlan-1-tag-all-or-untag-uni-e5e2ab7a

In Unifi - Default vlan 1 - Tag all or untag????? **Uni** 
        
        
        
    
    
    Hello ,
I have a UDM and 2x nanoHD (APs) and a non unifi switch which support vlans (avaya 4526t-pwr+ 24-)
So from the UDM its connected to port 24 on switch and i made it a trunk port Tagged vlans 107,168,172 ,1. - Works great when i untag a particular port with a desired vlan ,
vlan 1 - 192.168.100.xx | vlan 107 - 192.168.107.xx. | vlan 168 - 192.168.168.xx | vlan 172 -172.172.172.xx
But my APs wont work - Here is the setup
I connected APs on port 21 & 22 and made it trunk with all vlans tagged , and pvid as 1 it throws DHCP for vlan 1 which is default 192.168.100.xx but the clients wont get IPs fron vlan 107 and 172
Below are the commands for port 21 , 22
vlan ports 21,22 tagging tagAll
vlan mem add 1,107,168,172 21,22
vlan ports 21,22 pvid 107
vlan ports 21,22 pvid 168
vlan ports 21,22 pvid 172
vlan ports 21,22 pvid 1
int eth 21,22
no spanning-tree bpdu-filtering
no shut
exit
wr mem
Note : the APs shows connected in controller only when i make the port 21,22 untag 1 , when i leave them as tag default vlan 1 , they keep disconnecting from controller and keeps looping on adopting
Same setup works fine on Netgear GS110 smart switch (vlans works fine when i leave it on vlan 1 tagged and also tag all other vlans )
hope iam clear on this ....
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Network topology, configuration, etc.
i just updated my post ...
By default, Unifi doesn't tag the management VLAN.
Not a whole lot of info to go on. It would help for you to describe your setup.
Did you set up a DHCP server in each VLAN or, alternatively, DHCP relay?
DHCP is in the VLAN itself (UDM) ,no DHCP relay i. i have update my post above pls have a look at it ...
I don't quite understand the Avaya configuration. It looks like you have multiple PVIDs on a port 21. That doesn't make sense. Same with port 22.
Anyway, I see two problems on the Unifi side of things.
As I mentioned before, the nanoHDs, by default, expect the management VLAN to be untagged. You should set VLAN 1 to be untagged on Avaya ports 21 and 22. Plus set pvid to 1.
Alternatively, if you really want the APs to operate on VLAN 1 tagged, then you have to go to the AP's config, go to Services and set the Management VLAN to a network that has a VLAN tag set. 2. The other problem is that I don't think you have set up the IoT and CCTV SSIDs to use VLAN tags. Under the controller's site settings, you have to enable Advanced Features. This will expose a VLAN setting for each wireless network in the controller. That's where you need to set the VLAN tag (i.e. 107 for IoT and 168 for CCTV).
create a switch profile with the default VLAN of 1 (untagged) then tag all the other vlans you want to send upstream.
assign profile
win
