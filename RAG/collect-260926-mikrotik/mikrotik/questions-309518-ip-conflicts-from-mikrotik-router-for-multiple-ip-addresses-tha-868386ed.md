---
id: collect-260926-mikrotik/mikrotik/questions-309518-ip-conflicts-from-mikrotik-router-for-multiple-ip-addresses-tha-868386ed
title: "questions-309518-ip-conflicts-from-mikrotik-router-for-multiple-ip-addresses-tha-868386ed"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-309518-ip-conflicts-from-mikrotik-router-for-multiple-ip-addresses-tha-868386ed.md
source_anchor: ""
source_lines: [1, 19]
sha256: 3ad978e8fcbc08ca6ff7145afbec3086cd84b73676977688ce4294d034607403
---

# questions-309518-ip-conflicts-from-mikrotik-router-for-multiple-ip-addresses-tha-868386ed

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a point to point wireless connection using two MikroTiks. When I plug the MikroTik into a switch with just my laptop I get an IP address conflict on my machine no matter what IP I am assigned.
Using Wireshark, I see the conflicts are from the mac address of the MikroTik on the other end of the wireless connection. Why is it conflicting with multiple IP addresses when the router itself is assigned a single IP address with no NAT entries or anything like that? I included a little diagram to help visualize my issue.
[me]
[MikroTik] --------------[problem MikroTik]----(other equipment on diff subnet)
The problem MikroTik has a WAN on the same subnet as my machine. The LAN is a different subnet. Any ideas? When I plug the equipment into my network I get IP conflicts on a lot of different servers. Took me forever to isolate it to this MikroTik!
All this equipment has been working previously with no known changes made to the configs. It just started acting up recently.
Have you tried clearing the ARP cache from all devices? If multiple devices shared the same IP address, that is recorded in the ARP cache of neighboring network devices. Til that is cleared, the devices will keep thinking there are duplicate addresses on the network
Most likely the "IP Conflict" is caused by several configuration options of problem mikrotik interface that (over the wireless link) is bridged to me:
IP address (with proper subnet mask) belonging to me's subnet is missing (this is the misconfiguration that should be fixed);
proxy-arp is enabled (this might be OK depending on the needs).
It is possible that the above mentioned interface is a bridge itself, then IP address, most likely, should be assigned to the bridge (and not any interfaces belonging to the bridge).
Can you confirm the problematic Mikrotik does not have Hotspot functionality enabled? If it is enabled, probably your side of the point-to-point link is being part of the "client" hotspot feature.
Mikrotik does a lot of L2 and L3 "magic" to capture any wireless client that may connect and "capture" their traffic to enforce any hotspot authentication or restriction. It does that no matter the client IP configuration.
