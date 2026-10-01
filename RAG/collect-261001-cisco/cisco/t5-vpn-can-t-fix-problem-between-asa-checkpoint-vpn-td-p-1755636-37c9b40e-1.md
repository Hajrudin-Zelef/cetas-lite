---
id: collect-261001-cisco/cisco/t5-vpn-can-t-fix-problem-between-asa-checkpoint-vpn-td-p-1755636-37c9b40e-1
title: "t5-vpn-can-t-fix-problem-between-asa-checkpoint-vpn-td-p-1755636-37c9b40e"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-vpn-can-t-fix-problem-between-asa-checkpoint-vpn-td-p-1755636-37c9b40e.md
source_anchor: ""
source_lines: [1, 49]
sha256: 8c6260412eba9bc6a97d2c7df5feac15ee87d449beb24753214bc04378aa5d2d
---

# t5-vpn-can-t-fix-problem-between-asa-checkpoint-vpn-td-p-1755636-37c9b40e

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-27-2011 02:29 PM
Hi Team,
I have a strange problem with a VPN L2L between an ASA on my side and a CheckPoint as the peer.
The IPsec tunnel works fine, but from time to time, traffic stops passing through the tunnel.
Scenario:
172.31.250.0/28 -- ASA --- Internet --- CheckPoint --- 200.122.x.y/32
I've done plenty tunnels between ASAs and CheckPoints but this time we've found this:
access-list outside_1_cryptomap extended permit ip 172.31.250.0 255.255.255.240 host 200.122.164.165
local ident (addr/mask/prot/port): (172.31.250.0/255.255.255.240/0/0)
remote ident (addr/mask/prot/port): (200.122.164.165/255.255.255.255/0/0)
#pkts encaps: 0, #pkts encrypt: 0, #pkts digest: 0
#pkts decaps: 1148, #pkts decrypt: 1148, #pkts verify: 1148
local ident (addr/mask/prot/port): (172.31.250.8/255.255.255.248/0/0)
remote ident (addr/mask/prot/port): (200.122.164.0/255.255.255.0/0/0)
#pkts encaps: 27682, #pkts encrypt: 27683, #pkts digest: 27683
#pkts decaps: 27683, #pkts decrypt: 27683, #pkts verify: 27683
local ident (addr/mask/prot/port): (172.31.250.8/255.255.255.248/0/0)
remote ident (addr/mask/prot/port): (200.122.164.165/255.255.255.255/0/0)
#pkts encaps: 3579, #pkts encrypt: 3579, #pkts digest: 3579
#pkts decaps: 10443, #pkts decrypt: 10443, #pkts verify: 10443
Traffic is defined between 172.31.250.0/28 and a single host, but I see three SAs:
1. 172.31.250.0/28 - 200.122.164.165/32
2. 172.31.250.8/32 - 200.122.164.0/24
3. 172.31.250.8/32 - 200.122.164.165/32
What is the reason for this??
The reason I paste the above is because the CheckPoint defines the ''interesting traffic'' as two rules (one in each direction).
On CheckPoint:
Rule 1: Traffic from 200.122.164.165/32 to 172.31.250.0/28
Rule 2: Traffic from 172.31.250.0/28 to 200.122.164.165/32
So, I believe the problem happens because we define the phase 2 SAs as bidirectional rules (crypto ACLs), and CheckPoint defines the phase 2 SAs as unidirectional rules. Even though the traffic matches, I see the above output.
I think it means that the ASA receives some traffic in one SA and send it via another, and I'm not sure if that causes the problem, and if so, how to fix it?
The problem is totally random. We lowered the rekey time to 2 minutes on phase 2 and 5 minutes on phase 1 and there's no problem during rekey.
We hadn't been able to capture the log at the exact moment of the problem. Then the tunnel suddenly comes up again and start working.
ASA 5510 version 8.2(5)
Any help is appreciated!
Federico.
Solved! Go to Solution.
- Labels:
- 
						
							
		
