---
id: collect-260926-mikrotik/mikrotik/questions-32962-configurating-vlan-access-and-trunk-ports-on-routeros-mikrotik-b7811aa5
title: "questions-32962-configurating-vlan-access-and-trunk-ports-on-routeros-mikrotik-b7811aa5"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-32962-configurating-vlan-access-and-trunk-ports-on-routeros-mikrotik-b7811aa5.md
source_anchor: ""
source_lines: [1, 31]
sha256: 331588843f52c215988596824ae499d56d0d2b52362cf12cbb446008d28e0f9b
---

# questions-32962-configurating-vlan-access-and-trunk-ports-on-routeros-mikrotik-b7811aa5

I have RouteBOARD RB951G-2HnD, RouterOS 6.34.2 and I am trying to achive a network configuration something like what is depict on the next picture. The picture is from Mikrotik Wiki http://wiki.mikrotik.com/wiki/Vlans_on_Mikrotik_environment.
I'm trying to configure the router same way as SW3 on the picture. Two trunk port and one access port. I tried to configure the router as it is described on the wiki page and the access port is not working correctly. I was trying to find similar problem on the Internet but with no luck.
I'm missing something in the configuration. So, here are the steps that I used for configuration. The router have default configuration and I'm trying to configure next:
- ether4 (on the picture same as SW3 eth3) and ether5 (on the picture same as SW3 eth4) are trunk ports
- two vlans id, 100 and 200
- ether3 (on the picture same as SW3 eth1) is access port for vlan 100
Steps
#Removing parts of default configuration
/interface ethernet set ether3 master-port=none
/interface ethernet set ether4 master-port=none
/interface ethernet set ether5 master-port=none
#Configuration for trunk ports
/interface bridge add name=bridge-trunk disabled=no
/interface bridge port add interface="ether3" bridge="bridge-trunk" disabled=no
/interface bridge port add interface="ether4" bridge="bridge-trunk" disabled=no
#configuration for the access port
/interface vlan add name="bridge-trunk-vlan100" vlan-id=100 interface=bridge-trunk disabled=no
/interface bridge add name=bridge-vlan100 disabled=no
/interface bridge port add interface="bridge-trunk-vlan100" bridge="bridge-vlan100" disabled=no
/interface bridge port add interface="ether3" bridge="bridge-vlan100" disabled=no
When I do these steps and connect my laptop to the ether3, any type of network connection to other network devices is not working. Trunk ports are working properly.
Correct me if I'm wrong, if I introduce the router configured this way into existing network, APR tables must refresh itself on other network devices because bridge have it's own MAC address.
I would be grateful if you can provide with hints, tutorials or books to read more about a bridging and VLANs.
Update 1
I found a way but I do not know if it is right way by the RouterOS methodology.
I added the filter rule for the bridges
/interface bridge filter add chain=forward mac-protocol=vlan vlan-id=100 action=accept
Also, I could run all VLAN traffic through IP Firewall with similar rule.
Update 2
Previous situation was done in a experimental environment on separate RouterBOARD. I wanted to apply this solution on the real device and it did not worked. The device is Cloud Router Switch CRS125-24G-1S. The moment when I add bridge-trunk-vlan100 interface to the bridge-vlan100 bridge, the traffic that is passing through bridge-trunk is dropped and new connections cannot be established.
bridge-trunk disabled=no. I would think it should be:bridge-trunk disabled=yes.
