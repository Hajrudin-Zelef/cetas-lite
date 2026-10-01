---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1avx8wl-port-configuration-native-vlan-network-and-tagged-2185f7ae
title: "r-ubiquiti-comments-1avx8wl-port-configuration-native-vlan-network-and-tagged-2185f7ae"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1avx8wl-port-configuration-native-vlan-network-and-tagged-2185f7ae.md
source_anchor: ""
source_lines: [1, 30]
sha256: e66764c8114699f8f95cd828d5331a7a7c65fc3461f06d3df463fec0a60335c5
---

# r-ubiquiti-comments-1avx8wl-port-configuration-native-vlan-network-and-tagged-2185f7ae

Port Configuration: Native VLAN / Network and Tagged VLAN Management 
        
        
        
    
    
    Hello All,
I’m trying to figure out what the fields Native VLAN / Network and Tagged VLAN Management mean in a Unifi controller (USG-3) under port configuration.
I’m not clear on what untagged traffic is. When is traffic tagged and what device tags it? Is this for outgoing traffic from this port?
Regarding the Tagged VLAN Management, if I know that all data going to that port should on be on VLAN1, then can I only Allow VLAN1 traffic? Is that the same as blocking all other VLANS but VLAN, or is it different because of untagged traffic?
To make matters more confusing. I have one port that has an unmanaged switch attached to it. It has my computer on it (VLAN1) and my Printer (VLAN2). How should I set the Native VLAN / Network and Tagged VLAN Management on this port?
Thank you in advance.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
A basic "flat" lan (no VLANs) passes only a single network thru each port. with VLAN aware network equipment, multiple networks can pass thru ports in the form of VLANs.
A Port configured to carry VLANs has a single network defined as "Native" or untagged. Ant device directly connected to that port will be considered to be on that "untagged" network by default. Traffic for any additional VLANs passing thru that has to be "Tagged".
In your example, having a port with only a single vlan, that VLAN would be Untagged (native) and would be the only traffic passing thru that port. Unifi sometimes refers to this as the "default" network for that port. The LAN port on your USG should have your main network as untagged ans all other VLANs as tagged.
Unmanaged switches do not support tagging of VLAN traffic and as a result resides on the untagged network - as do all devices directly connected to it. If your PC and printer are both connected to the unmanaged switch, then they will be on the same VLAN. You need a managed switch that supports VLANs to do what you describe.
Thank you OtherTechnician.
Just so I understand:
If I have one device on a port it can be an access port.
If I have an unmanaged switch on a port and have two devices on that switch on the same VLAN it can be an access port
If I have a switch on a port with two devices on that switch that I want to have those devices on different VLANS than I must use a managed switch and I must make it a trunk port
Yes
Here a video that may help
https://youtu.be/JszGeQPTo4w?feature=shared
