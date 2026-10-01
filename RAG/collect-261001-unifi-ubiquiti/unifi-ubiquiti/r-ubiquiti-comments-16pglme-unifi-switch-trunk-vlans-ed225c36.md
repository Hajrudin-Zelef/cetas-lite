---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-16pglme-unifi-switch-trunk-vlans-ed225c36
title: "r-ubiquiti-comments-16pglme-unifi-switch-trunk-vlans-ed225c36"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-16pglme-unifi-switch-trunk-vlans-ed225c36.md
source_anchor: ""
source_lines: [1, 24]
sha256: 2529f4d7f1ec7fa488ed86d750e0517d61e720a18ddd18549bdba9701cf8634e
---

# r-ubiquiti-comments-16pglme-unifi-switch-trunk-vlans-ed225c36

Unifi Switch Trunk VLANS 
        
        
        
    
    
    Hello all, hope you all are having a lovely Friday.
I'm trying to setup trunk ports on the 2 Unifi switches that I have (USW Flex and USW-Lite 8 Port POE). I mostly have a Unifi network, however the Routing/DHCP/VLANs, etc. are done by the Fortigate Firewall that I have.
On the USW-Lite, I have my Cloud Key, U6 Mesh, and 2 of my Unifi Protect Cameras connected to this switch.
I am currently on one VLAN to ensure network connectivity in the house,
I have a few questions:
When setting a network port to be a trunk, which is the proper way to do it for the Unifi Switches? I know you can do it through Ethernet Profiles and setting the Traffic Restrictions, but what would be the preferred way to do it?
If I set one of the ports to be the trunk, I'm assuming that for the ports that I have connected the Cloud Key, U6 Mesh, etc, I'll have to set their native VLANs/Primary Network, on it as well? Whenever I try it, it seems to lose connection from the Cloud Key to the switch, so is there a way to set the trunk port and set the Native VLAN at the same time?
Apologies if I'm missing something, I'm not a networking expert.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
I thought all switch ports were trunks by default and you only restrict them to a VLAN by associating a port profile with that restriction to the port.
Are they? When I tried to assign a VLAN through restriction it wouldn't get an address, and I've already confirmed that DHCP is enabeld on the Firewall.
I have a number of switches with ports dedicated to specific VLANs. The uplinks use default settings and the traffic from the VLANs has no problem getting back to the console.
Now I wonder if it is doing what I thought it was doing. The port says “primary network “ = my IoT network. I thought that used to say port profile rather than primary network. Hmm. Anyway, I know my device is getting a DHCP address from the correct subnet for the IoT VLAN and the traffic is getting to the Internet through a Lite-16 uplink to a Pro-24-PoE and then to a UDM Pro.
