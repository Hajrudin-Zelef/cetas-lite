---
id: collect-261001-general-networking/general-networking/questions-5888-vpn-client-to-multiple-locations-simultaneously-705ffa11
title: "questions-5888-vpn-client-to-multiple-locations-simultaneously-705ffa11"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-5888-vpn-client-to-multiple-locations-simultaneously-705ffa11.md
source_anchor: ""
source_lines: [1, 3]
sha256: a1021b15f2935c65dd943b2c7bd5e58a69ce70b66fed9da35a2312480e36ce0a
---

# questions-5888-vpn-client-to-multiple-locations-simultaneously-705ffa11

I need to connect to 15+ locations to run network scans weekly. All the locations have Fortigate firewalls over which I have full control. The current solution I have is to connect via IPSec VPN to each location one by one to run my scans. I have considered scripting the connection and scanning process but it seems like network connectivity to all locations simultaneously is a superior solution.
All locations have the same internal subnet (which I can't control) and many locations have dynamic IP addresses. Because of this I was thinking I would need some type of NAT for VPN to give each subnet a unique address on my end allowing me to reach them all. But I am not sure how this would work.
I am wondering what the best solutions is, either hardware or through a VPN client, that would allow me to more easily gain access to these networks (preferably simultaneously). I've looked into StrongSwan as a VPN server and a Fortigate firewall on my side for form the mesh but don't have a solution yet.
