---
id: collect-260926-mikrotik/mikrotik/questions-10769-two-or-more-different-networks-without-vlans-using-mikrotik-devi-e085a238
title: "Two or more different networks without VLANs using MikroTik devices"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-10769-two-or-more-different-networks-without-vlans-using-mikrotik-devi-e085a238.md
source_anchor: ""
source_lines: [1, 59]
sha256: d5b8778079f132336630616c3abc26e22de6969747ae1fe3384c238e3b323778
---

# Two or more different networks without VLANs using MikroTik devices

*Source : https://networkengineering.stackexchange.com/questions/10769/two-or-more-different-networks-without-vlans-using-mikrotik-devices | Site : networkengineering | Score : 0*

There is need to add WiFi network to boardroom. I would like to add one more LAN segment to my 192.168.1.0 network. I wish that network have diferent IP address scheme - 192.168.2.0. Edge router is connected to my new WiFi router by a separate switch.

Also I would like to separate 192.168.2.0 traffic from 192.168.1.0 traffic. Hosts from 192.168.2.0 network must not have acces to devices in 192.168.1.0 network.

Is it possible to do that without VLANs ?

How should be configured MikroTik WiFi router ?

In what mode the WiFi router should work ?

---

## Reponse — score 1

Is it possible to do this without VLANs?

It can be done without VLANs. The simplest way might be to connect the Wifi router directly to the Cloud Core router/firewall instead of to the HP switch. This way, the firewall functionality in the CCR can prevent 192.168.2.0 from reaching 192.168.1.0.

A quick and dirty solution that might meet your requirements is configuring your wifi router as any home wifi router, NATting the wireless subnet onto your network. No vlans, and the rest of the network cannot contact the wireless subnet. I would advise against this because NAT is not necessary inside your network, but it is a quick fix.

---

## Reponse — score 0

Is it possible to do that without VLANs ?
yes, but why, you have good switch?! ;-)

How should be configured MikroTik WiFi router ?
In what mode the WiFi router should work ?

If your switch is in default config, then you just need to connect wifi-router ethernet1 port to the switch. This will be "wan" for wifi-router, add address to this interface for. example 192.168.1.254, set gateway 192.168.1.1. It will works. 

If you want to limit connections from 192.168.2.0 to 192.168.1.0, set acl on mikrotik. But, if you want, to permit 192.168.1.0 to 192.168.2.0 and deny 192.168.2.0 to 192.168.1.0, both, you need:

1. disable nat on mikrotik(routing mode), allow forwarding from "wan". 
2. set route on cloud core, a-la: ip route -net 192.168.2.0 -mask 255.255.255.0 -gateway 192.168.2.254
3. set mikrotik firewall(allow from 192.168.1.0 to 2.0, deny 192.168.2.0 to 1.0, except 1.1)

---

## Reponse — score 0

Yes. It possible in different ways.

Gerben and pyatka describe some possible solutions.

I can add more ways.

- 
Create EoIP tunnel between RB2011 and CCR and bridge Wireless and this tunnel at RB2011. It allow make all wireless settings on AP (RB2011) and all routing, dhcp, firewall... on router (CCR).

- 
Use CAPsMAN on CCR and make Wi-Fi AP controlled. And use Central forwarding. Then virtual Wireles interface arrive on controller (CCR). Then almost all settings for AP done on CCR.

PS. EoIP on my experience can limit bandwidth to 10 - 30 Mbit, but it can be enough for one AP.
