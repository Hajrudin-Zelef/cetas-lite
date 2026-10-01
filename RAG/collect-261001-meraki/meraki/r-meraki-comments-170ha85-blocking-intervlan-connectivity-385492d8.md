---
id: collect-261001-meraki/meraki/r-meraki-comments-170ha85-blocking-intervlan-connectivity-385492d8
title: "r-meraki-comments-170ha85-blocking-intervlan-connectivity-385492d8"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-170ha85-blocking-intervlan-connectivity-385492d8.md
source_anchor: ""
source_lines: [1, 21]
sha256: a81d50204112b902dfa0545968cd99a6956be1e93111137f5a3060a04059d212
---

# r-meraki-comments-170ha85-blocking-intervlan-connectivity-385492d8

Blocking Inter-VLAN connectivity 
        
    We have VLAN1 10.110.1.0/24 and VLAN2 10.113.1.0/24 - devices on each VLAN are able to ping each other. We want VLAN2 to be fully isolated as it is open to the internet. How would we achieve this.
Are Firewall rules the only way or is there an easy way to only allow 10.113.1.0/24 traffic on the specific port we set?
We also looked further in the Per-Port VLAN settings at the ALLOWED VLANs area - what does this do? We didn't notice this making a difference.
Thank you.
Section des commentaires
You can set layer 3 firewall..
Deny vlan 2 to vlan 1 Then deny vlan 1 to vlan 2 And then allow any for last rule..
This way, in this case, both vlans can't get to each other..
If you plan future nee subnets, you can also deny vlan 2 to 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.. These are the rfc1918 local IP ranges. So, Vlan 2 would never be able to do anything internally and can only go to the internet, public IPs.
You can also look into isolated ports in Meraki KB
Trunks with allowed vlans are on an uplink between switches. With that, you can (dis)allow certain vlans to reach a switch, or not... You can segment the network so specific vlans are available in certain places..
Your client sends untagged traffic and your switch tags that traffic with the native vlan or the data vlan set on the port.. your switch then sends out that same traffic tagged to its neighbor, upstream switch where your L3 interface resides... Si trunks between switches need to be allowed that vlan to traverse traffic between them of that specific vlan.
Few things to add as well as some questions for ya:
How are these VLANs propagated on your network? Multiple switches? One switch?
How are these VLANs communicating with each other? Do you have a layer 3 switch? Where is the L3 device that allows for communication across VLANs?
On that layer 3 device (this can be firewall or switch), you want to define an ACL as others mentioned above. Check this out https://documentation.meraki.com/MS/Layer_3_Switching/Configuring_ACLs
Assuming you are not using a L3 switch, and are attempting to do this on an MX, then yes this is what firewall rules are for.
The default meraki firewall rule allows any traffic to be routed. You can create a rule to deny all local traffic from being permitted, and work backwards from there.
Create a full that denies 10.0.0.0/8 to 10.0.0.0/8 and that will block the communication for that and any future VLANs in the 10 range.
