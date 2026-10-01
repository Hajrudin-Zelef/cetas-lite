---
id: collect-261001-cisco/cisco/questions-734988-cisco-bgp-unequal-cost-load-balancing-e870a89b-2
title: "Here are bgp neighbours from global routing table. Not relevant to the question. IP addresses are hidden"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-734988-cisco-bgp-unequal-cost-load-balancing-e870a89b.md
source_anchor: ""
source_lines: [141, 151]
sha256: 1f2e73e3bf66b43bd6e8ff552d34871c8cb53fd95631fac2da39b36ba3240e4c
---

# Here are bgp neighbours from global routing table. Not relevant to the question. IP addresses are hidden

C       10.30.228.0 is directly connected, Vlan3228
C       10.30.227.0 is directly connected, Vlan3227
C       10.30.225.0 is directly connected, Vlan3225
B*   0.0.0.0/0 [20/0] via 10.30.228.228, 3d18h
               [20/0] via 10.30.227.227, 3d18h
               [20/0] via 10.30.225.225, 3d18h
R1# sh ip bgp vpnv4 vrf nat neighbors
R1 sh ip bgp neighbours output
R1# sh run
R1 running config sensitive information is masked
bandwidth 50000for 'interface Vlan3228'? Could you please also attach output of 'sh ip bgp <IP of R5>'?R1# sh ip bgp vpnv4 vrf nat 0.0.0.0shows 125kbs-250kbs-350kbs...sh ip bgp vpnv4 vrf nat 10.30.228.228displays the same assh ip bgp vpnv4 vrf nat 0.0.0.0. Have you missedneighborskeyword? If so, output ofsh ip bgp vpnv4 vrf nat neighbors 10.30.228.228you can see at the end of R1 sh ip bgp neighbours outputneighbor dmzlink-bw, which only enables advertisement of bandwidth to neighbours (presented in your config) andbgp dmzlink-bw, which enables proportional load balancing (and it seems to be MISSING in your config). Could you try to putbgp dmzlink-bwinto your running config?maximum pathsshould be underaddress-family. I need multipath feature in vrf instance but not in global routing table. If I putmaximum pathsunderrouter bgp 100I get in result only one route via R5 and no routes via R3 and R2. Same result foribgpunder address-family, because all neighbours are external. Bandwidth inherit on Port-channel is consistent in config. I removed this line with no effect. commit
