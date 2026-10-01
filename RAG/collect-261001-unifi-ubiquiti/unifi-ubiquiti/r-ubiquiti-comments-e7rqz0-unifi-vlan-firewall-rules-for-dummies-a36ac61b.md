---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-e7rqz0-unifi-vlan-firewall-rules-for-dummies-a36ac61b
title: "r-ubiquiti-comments-e7rqz0-unifi-vlan-firewall-rules-for-dummies-a36ac61b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-e7rqz0-unifi-vlan-firewall-rules-for-dummies-a36ac61b.md
source_anchor: ""
source_lines: [1, 25]
sha256: 5a27fa32364ed597368392a4bec0737bbaf02ab7c1e60eadd42fe75f5e56cf0b
---

# r-ubiquiti-comments-e7rqz0-unifi-vlan-firewall-rules-for-dummies-a36ac61b

Unifi VLAN firewall rules for dummies 
        
    I've had a Unifi system set up for a while with one network - its been working fine.
I've just added a Unifi video camera mounted outside and I'd like to put it on its own VLAN in case someone comes along and connects into my network. I really should also move the chinese vacuum cleaner onto its own VLAN as well.
I've looked around a bit and think I understand the concept of VLANs. I can create the VLANs and assign them to switch ports.
But I don't know how to create firewall rules! Is there a guide for dummies somewhere?
Assume
- 
      VLAN10 = main network
- 
      VLAN20 = video cameras
- 
      VLAN30 = IOT devices
I want the cameras to be able to access the Cloud Key 2+ (only). Ideally only the ports required.
I want the VLAN30 to be able to access the internet but not VLAN10/20.
I want VLAN10 to be able to access VLAN20/VLAN30.
I feel it should be pretty simple I'm just not sure how to make it happen...
Section des commentaires
Thanks everyone - I'll go and look at the links/info provided. I'm fighting a few different issues at the moment (see my other post) so it might take me a little while.
I just used the premade GUEST and WAN/LAN firewall rules you can see here: https://help.ubnt.com/hc/en-us/articles/115003173168-UniFi-USG-Firewall-Introduction-to-Firewall-Rules
The site explains what every rule does, so read through it and shoot if you have any more questions.
Crosstalk solutions has some videos that might help. https://youtu.be/6ElI8QeYbZQ
I have a similar setup on an Edge Router. The interface may be a little different, but the concept is the same. I have a device on LAN1 that needs to accept connections from a device on LAN2. LAN2_IN is set to accept by default.
I then have 2 rules under LAN2_IN. The first accepts connections that match the destination IP and port of the device they're supposed to connect to on LAN1. The second rejects all connections to LAN1.
The firewall considers the rules in the order you assign, so if rule 1 isn't matched (connect to a specific device), it considers rule 2 (any connections to LAN1). If no rules are matched (it's trying to connect to it's own or another LAN), then the default rule applies (accept).
