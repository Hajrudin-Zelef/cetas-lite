---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-bgp-allowas-in-enable-or-as-override-when-local-as-ta-b693e214-2
title: "t5-fortigate-technical-tip-bgp-allowas-in-enable-or-as-override-when-local-as-ta-b693e214"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-bgp-allowas-in-enable-or-as-override-when-local-as-ta-b693e214.md
source_anchor: ""
source_lines: [148, 163]
sha256: 1bbd4a30a3d69fb6e20f934ad090b19cf93371e8f77ec75a11f12e3931f1eb26
---

# t5-fortigate-technical-tip-bgp-allowas-in-enable-or-as-override-when-local-as-ta-b693e214

       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
       * - candidate default
S*      0.0.0.0/0 [10/0] via 10.109.31.254, port1
C       10.109.16.0/20 is directly connected, port1
B       10.201.0.0/20 [20/0] via 10.109.19.146, port1, 00:02:23
B       10.205.0.0/20 [20/0] via 10.109.19.146, port1, 00:02:23
Location B # get  router  info bgp network 10.201.0.0/20
BGP routing table entry for 10.201.0.0/20
Paths: (1 available, best #1, table Default-IP-Routing-Table)
  Not advertised to any peer
  1111 1111
    10.109.19.146 from 10.109.16.172 (10.201.0.172)
      Origin incomplete metric 0, localpref 100, valid, external, best
      Last update: Sun Oct  6 14:25:57 2019
Note:
After enabling the 'allowas-in-enable' or 'as-override', the BGP neighbor gets down and comes up.
