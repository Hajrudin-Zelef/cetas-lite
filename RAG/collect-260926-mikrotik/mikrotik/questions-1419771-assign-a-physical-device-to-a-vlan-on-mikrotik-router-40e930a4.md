---
id: collect-260926-mikrotik/mikrotik/questions-1419771-assign-a-physical-device-to-a-vlan-on-mikrotik-router-40e930a4
title: "questions-1419771-assign-a-physical-device-to-a-vlan-on-mikrotik-router-40e930a4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-1419771-assign-a-physical-device-to-a-vlan-on-mikrotik-router-40e930a4.md
source_anchor: ""
source_lines: [1, 5]
sha256: 9d396ad7a2ba36ec4d1a43a08277b638373423a8a54f3652ee587028df79f424
---

# questions-1419771-assign-a-physical-device-to-a-vlan-on-mikrotik-router-40e930a4

The router cannot control what happens to packets after they're sent from the ethernet port – if they go to a switch, then that switch makes the next decision. No matter what sort of VLAN tags the router attaches, they're meaningless if the packets go to a switch that simply doesn't understand them.
So in short, if you use unmanaged switches which don't support VLANs, then by definition, you cannot separate devices connected to that switch into different VLANs.
Similarly for traffic between devices on that switch – you cannot enforce anything via your router when the traffic doesn't go through the router. If the switch doesn't have a feature to prevent two ports from communicating, then you can't prevent that.
Your only remaining option is to have multiple subnets on the same VLAN (e.g. 192.168.88.0/24 and 192.168.30.0/24 on the same interface), and use static DHCP leases to define which device is assigned which address from which subnet.
This won't provide good isolation, but it'll provide some isolation – for IPv4 – as your devices won't know that both subnets happen to be on the same link, so all traffic between them will still go through the router (default gateway), as long as the router is configured to not send ICMP "Redirect" packets. This method won't work with IPv6 due to its stateless configuration – the router cannot send an autoconfiguration broadcast to "all devices except that one in a corner", and it cannot prevent devices from having link-local addresses either.
