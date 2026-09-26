---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-194033p-mikrotik-vlan-routing-interface-with-a-dhcp-a60c7048
title: "r-mikrotik-comments-194033p-mikrotik-vlan-routing-interface-with-a-dhcp-a60c7048"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/r-mikrotik-comments-194033p-mikrotik-vlan-routing-interface-with-a-dhcp-a60c7048.md
source_anchor: ""
source_lines: [1, 11]
sha256: 161aa49a9bc28bc16d77ec74d8ac3a930ae5743766d5075a561eba101018b273
---

# r-mikrotik-comments-194033p-mikrotik-vlan-routing-interface-with-a-dhcp-a60c7048

Mikrotik VLAN routing interface with a DHCP client assigned address in a bridge VLAN filtered VLAN configuration 
        
    I'm pretty new to Mikrotik and I'm still trying to wrap my head around its handling of VLANs.
I have an initial use case which will eventually involve two Mikrotik routers, each with a VLAN trunk port and a couple of access ports.
I followed the Bridge VLAN Filtering docs and setup a basic trunk and a single access port. A client PC plugged into the access port got an IP from the intended access VLAN which was passed as a tagged port from the trunk interface.
What I couldn't get working, though, was adding a VLAN routing interface whose IP address is assigned via DHCP to the access port VLAN. The DHCP source is external and already exists on the access port VLAN. Is this possible? I'm aware its kind of an unusual way of assigning a router interface IP, but its more for accessing the router management interface than actual routing.
I kind of assumed that the VLAN routing interface would have to be associated with the bridge associated with the VLANs, but I could never get an IP address assigned by the DHCP client.
It's kind of irrelevant to this question, but my long term use case for Mikrotek is using EoIP to bridge an ethernet segment to a remote location. I have a Hex S, but I'll probably end up using CHR in a VM on the remote peer site.
Section des commentaires
Remember that EoIP by itself is not encrypted. That isn't a problem in itself, as long as your use case is good with that.
vlans are completely separate interfaces, just like physical interfaces are, as far as dhcp and ip addresses are concerned. so just configure it accordingly, but putting the settings on the vlan interface.
