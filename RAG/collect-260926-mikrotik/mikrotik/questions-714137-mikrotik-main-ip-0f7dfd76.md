---
id: collect-260926-mikrotik/mikrotik/questions-714137-mikrotik-main-ip-0f7dfd76
title: "questions-714137-mikrotik-main-ip-0f7dfd76"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-714137-mikrotik-main-ip-0f7dfd76.md
source_anchor: ""
source_lines: [1, 4]
sha256: af70ca8b595fbd0986d02cfd8af8ef31808b3f77207c83816da64e060e0637d1
---

# questions-714137-mikrotik-main-ip-0f7dfd76

On mikrotik (CCR-1009, v6.0), I use 2 interfaces : eth1 for WAN and eth8 for LAN. On LAN interface 2 IPs: 192.168.22.1/24 as main and 10.10.10.1/24 for routing pptp from WAN. I also have some pptp/l2tp and IPSec connections from WAN port. For expamlpe I have IPsec connect, on my side for incoming tunnel IP:192.168.22.1, and on other side 192.168.8.1/24. When do traceroute from 192.168.22.xxx to 192.168.8.1 I see next:
  1     1 ms     1 ms    <1 мс  10.10.10.1
  2    30 ms    30 ms    30 ms  192.168.8.1
Question: How I understand mikrotik using 10.10.10.1 as IP because it first in sort IPs on LAN interface? Its true? If is it can i change it on 192.168.22.1, for example?
