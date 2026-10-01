---
id: collect-261001-general-networking/general-networking/t5-switching-subinterfaces-and-vlans-td-p-2057667-9845676b-3
title: "t5-switching-subinterfaces-and-vlans-td-p-2057667-9845676b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-general-networking/t5-switching-subinterfaces-and-vlans-td-p-2057667-9845676b.md
source_anchor: ""
source_lines: [176, 232]
sha256: e12bcf3c249fa725a35dcac90696bf11682df3317feb1a8cf3a3306224a1b00f
---

# t5-switching-subinterfaces-and-vlans-td-p-2057667-9845676b

Access Port: Such ports belong to 1 Vlan only ( 1 data Vlan + 1 Voice Vlan)
Data from and to access ports are always untagged.
TrunkPort: It offers data belonging to multiple VLAN to pass through. Trunk port does that by tagging.
Frame Tagging helps the recieving port to switch to differentiate between data of many VLANs.
Now this VLAN tag is not understood by the machines (Host or Server) which is connected to Switch.
To conclude:
1. Trunk port does Vlan tagging of frames.
2. Host Machine doesn't understand VLAN.
Therefore Host Machine is always connected to the switch's access port.
Interconnectivity of switch is using trunk ports.
Plz correct if I am wrong somewhere or it is improperly explained.
Regards,
Azmun
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-05-2017 12:33 PM
Hey,
I have a Cisco 2650 router with a NM-16ESW module attached, I'm looking for the same result as in the vlans inter-connecting. I have a DHCP server on vlan 10, PC's on vlan 30, etc... I can get IP addresses to each devices on each vlan from DHCP server, but when I go to ping or access resources from vlan 10 no luck. I have this all working without the use of sub-interfaces, do I need the sub-interfaces for this to work properly? Is there a document on this for my type of setup?
Thanks in advance for the help.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-23-2017 10:45 AM
Thanks JohnNathan,
Nicely explained , Iwas also having confussion with sub interfaces , now its cleared.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-29-2018 01:03 AM
I have an Extreme Network environment where I have a L3 core switch and other L2 access switches across the office. I did created 4 VLANS with 24 bit subnet. The topology is something like this.
For IT users, the VLAN Network range is X.X.1.0/24 with X.X.1.254 being the gateway.
For FINANCE users, the VLAN Network range is X.X.2.0/24 with X.X.2.254 being the gateway
For HR users, the VLAN Network range is X.X.3.0/24 with X.X.3.254 being the gateway
Now all of these .254 ip is configured in Core switch. Intervlan routing is configured and working fine. Now the trouble start when users from FINANCE and HR are trying to access our remote office networks. The router (Cisco 4331) is connected to Core switch with its LAN port ip as X.X.1.1/24. Now IT users are facing no problem in accessing remote location servers as the traffic is directly being routed to Router LAN port and from there to remote site through WAN.
But when I give a traceroute from FINANCE or HR users, the traffic doesn't cross over from their local gateway which is X.X.2.254 and X.X.3.254 respectively.
Now my question is, if I do configure the port in Switch connected to router interface as Trunk and create sub interfaces in Router LAN port, would my issue be resolved?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-03-2018 04:06 AM
