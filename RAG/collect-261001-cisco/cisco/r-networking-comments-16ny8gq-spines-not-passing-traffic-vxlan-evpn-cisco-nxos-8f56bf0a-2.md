---
id: collect-261001-cisco/cisco/r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a-2
title: "r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2023-09-20"]
keywords: []
source: docs/RAG/collect-261001-cisco/r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a.md
source_anchor: ""
source_lines: [18, 147]
sha256: 6e6ccbfb7fa2a4de1d2ffd2f59c76ddb8f1e8aac387573edef46c05fe162552b
---

# r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a

coresw501(config)# ethanalyzer local interface inband mirror display-filter "icmp" limit-captured-frames 100
Capturing on 'ps-inb'
1     3 2023-09-20 21:45:26.901327766    10.55.4.9 → 10.55.2.9    ICMP 148 Echo (ping) request  id=0x0029, seq=1/256, ttl=64
2     4 2023-09-20 21:45:27.906902294    10.55.4.9 → 10.55.2.9    ICMP 148 Echo (ping) request  id=0x0029, seq=2/512, ttl=64
3     5 2023-09-20 21:45:28.930895980    10.55.4.9 → 10.55.2.9    ICMP 148 Echo (ping) request  id=0x0029, seq=3/768, ttl=64
4     7 2023-09-20 21:45:29.954889685    10.55.4.9 → 10.55.2.9    ICMP 148 Echo (ping) request  id=0x0029, seq=4/1024, ttl=64
5     8 2023-09-20 21:45:30.978902610    10.55.4.9 → 10.55.2.9    ICMP 148 Echo (ping) request  id=0x0029, seq=5/1280, ttl=64
6    10 2023-09-20 21:45:32.002921998    10.55.4.9 → 10.55.2.9    ICMP 148 Echo (ping) request  id=0x0029, seq=6/1536, ttl=64
I don't get any replies on eth1/2 where 10.55.2.9 is. If I start pinging from 10.55.2.9 -> 10.55.4.9 I will see the requests here, but never a reply. I've run tcpdump on my hosts and I can verify I get no ICMP echo request and, as such, there are no ICMP echo replies.
I have the l2routes on the two leaves:
555         6cfe.545d.e870 10.55.2.9                               BGP    --            0         10.254.50.201 (Label: 10555)
555         6cfe.545d.e4f4 10.55.4.9                               HMM    L,            0         Local
switch508#
555         6cfe.545d.e870 10.55.2.9                               HMM    L,            0         Local
555         6cfe.545d.e4f4 10.55.4.9                               BGP    --            0         10.254.50.203 (Label: 10555)
switch503#
The spines can see the routes
coresw501(config)# sho bgp l2vpn evpn 10.55.4.9
BGP routing table information for VRF default, address family L2VPN EVPN
Route Distinguisher: 10.255.50.107:33322
BGP routing table entry for [2]:[0]:[0]:[48]:[6cfe.545d.e4f4]:[32]:[10.55.4.9]/272, version 952
Paths: (1 available, best #1)
Flags: (0x000202) (high32 00000000) on xmit-list, is not in l2rib/evpn, is not in HW
  Advertised path-id 1
  Path type: internal, path is valid, is best path, no labeled nexthop
  AS-Path: NONE, path sourced internal to AS
    10.254.50.203 (metric 2) from 10.255.50.107 (10.255.50.107)
      Origin IGP, MED not set, localpref 100, weight 0
      Received label 10555 50999
      Extcommunity: RT:64550:10555 RT:64550:50999 SOO:10.254.50.203:0 ENCAP:8
          Router MAC:9088.555c.f31b
  Path-id 1 advertised to peers:
    10.255.50.101      10.255.50.102      10.255.50.103      10.255.50.104
    10.255.50.105      10.255.50.106      10.255.50.108      10.255.50.109
    10.255.50.110
Route Distinguisher: 10.255.50.108:33322
BGP routing table entry for [2]:[0]:[0]:[48]:[6cfe.545d.e4f4]:[32]:[10.55.4.9]/272, version 955
Paths: (1 available, best #1)
Flags: (0x000202) (high32 00000000) on xmit-list, is not in l2rib/evpn, is not in HW
  Advertised path-id 1
  Path type: internal, path is valid, is best path, no labeled nexthop
  AS-Path: NONE, path sourced internal to AS
    10.254.50.203 (metric 2) from 10.255.50.108 (10.255.50.108)
      Origin IGP, MED not set, localpref 100, weight 0
      Received label 10555 50999
      Extcommunity: RT:64550:10555 RT:64550:50999 SOO:10.254.50.203:0 ENCAP:8
          Router MAC:9088.555d.2a93
  Path-id 1 advertised to peers:
    10.255.50.101      10.255.50.102      10.255.50.103      10.255.50.104
    10.255.50.105      10.255.50.106      10.255.50.107      10.255.50.109
    10.255.50.110
coresw501(config)# sho bgp l2vpn evpn 10.55.2.9
BGP routing table information for VRF default, address family L2VPN EVPN
Route Distinguisher: 10.255.50.103:33322
BGP routing table entry for [2]:[0]:[0]:[48]:[6cfe.545d.e870]:[32]:[10.55.2.9]/272, version 940
Paths: (1 available, best #1)
Flags: (0x000202) (high32 00000000) on xmit-list, is not in l2rib/evpn, is not in HW
  Advertised path-id 1
  Path type: internal, path is valid, is best path, no labeled nexthop
  AS-Path: NONE, path sourced internal to AS
    10.254.50.201 (metric 2) from 10.255.50.103 (10.255.50.103)
      Origin IGP, MED not set, localpref 100, weight 0
      Received label 10555 50999
      Extcommunity: RT:64550:10555 RT:64550:50999 SOO:10.254.50.201:0 ENCAP:8
          Router MAC:9088.555d.2453
  Path-id 1 advertised to peers:
    10.255.50.101      10.255.50.102      10.255.50.104      10.255.50.105
    10.255.50.106      10.255.50.107      10.255.50.108      10.255.50.109
    10.255.50.110
Route Distinguisher: 10.255.50.104:33322
BGP routing table entry for [2]:[0]:[0]:[48]:[6cfe.545d.e870]:[32]:[10.55.2.9]/272, version 943
Paths: (1 available, best #1)
Flags: (0x000202) (high32 00000000) on xmit-list, is not in l2rib/evpn, is not in HW
  Advertised path-id 1
  Path type: internal, path is valid, is best path, no labeled nexthop
  AS-Path: NONE, path sourced internal to AS
    10.254.50.201 (metric 2) from 10.255.50.104 (10.255.50.104)
      Origin IGP, MED not set, localpref 100, weight 0
      Received label 10555 50999
      Extcommunity: RT:64550:10555 RT:64550:50999 SOO:10.254.50.201:0 ENCAP:8
          Router MAC:9088.555d.257f
  Path-id 1 advertised to peers:
    10.255.50.101      10.255.50.102      10.255.50.103      10.255.50.105
    10.255.50.106      10.255.50.107      10.255.50.108      10.255.50.109
    10.255.50.110
coresw501(config)#
The switch pairs each see l2 routes:
switch503# show bgp l2vpn evpn 10.55.4.9 vrf vxlan_prod
BGP routing table information for VRF default, address family L2VPN EVPN
Route Distinguisher: 10.255.50.103:6    (L3VNI 50999)
BGP routing table entry for [2]:[0]:[0]:[48]:[6cfe.545d.e4f4]:[32]:[10.55.4.9]/272, version 2030
Paths: (2 available, best #2)
Flags: (0x000202) (high32 00000000) on xmit-list, is not in l2rib/evpn, is not in HW
  Path type: internal, path is valid, not best reason: Router Id, no labeled nexthop
             Imported from 10.255.50.108:33322:[2]:[0]:[0]:[48]:[6cfe.545d.e4f4]:[32]:[10.55.4.9]/272
  AS-Path: NONE, path sourced internal to AS
    10.254.50.203 (metric 3) from 10.255.50.1 (1.1.1.1)
      Origin IGP, MED not set, localpref 100, weight 0
      Received label 10555 50999
      Extcommunity: RT:64550:10555 RT:64550:50999 SOO:10.254.50.203:0 ENCAP:8
          Router MAC:9088.555d.2a93
      Originator: 10.255.50.108 Cluster list: 1.1.1.1
  Advertised path-id 1
  Path type: internal, path is valid, is best path, no labeled nexthop
             Imported from 10.255.50.107:33322:[2]:[0]:[0]:[48]:[6cfe.545d.e4f4]:[32]:[10.55.4.9]/272
  AS-Path: NONE, path sourced internal to AS
    10.254.50.203 (metric 3) from 10.255.50.1 (1.1.1.1)
      Origin IGP, MED not set, localpref 100, weight 0
      Received label 10555 50999
      Extcommunity: RT:64550:10555 RT:64550:50999 SOO:10.254.50.203:0 ENCAP:8
          Router MAC:9088.555c.f31b
      Originator: 10.255.50.107 Cluster list: 1.1.1.1
  Path-id 1 not advertised to any peer
switch503#
switch507# show bgp l2vpn evpn 10.55.2.9 vrf vxlan_prod
BGP routing table information for VRF default, address family L2VPN EVPN
Route Distinguisher: 10.255.50.107:6    (L3VNI 50999)
BGP routing table entry for [2]:[0]:[0]:[48]:[6cfe.545d.e870]:[32]:[10.55.2.9]/272, version 2056
Paths: (2 available, best #2)
Flags: (0x000202) (high32 00000000) on xmit-list, is not in l2rib/evpn, is not in HW
  Path type: internal, path is valid, not best reason: Router Id, no labeled nexthop
             Imported from 10.255.50.104:33322:[2]:[0]:[0]:[48]:[6cfe.545d.e870]:[32]:[10.55.2.9]/272
  AS-Path: NONE, path sourced internal to AS
    10.254.50.201 (metric 3) from 10.255.50.1 (1.1.1.1)
      Origin IGP, MED not set, localpref 100, weight 0
      Received label 10555 50999
      Extcommunity: RT:64550:10555 RT:64550:50999 SOO:10.254.50.201:0 ENCAP:8
          Router MAC:9088.555d.257f
      Originator: 10.255.50.104 Cluster list: 1.1.1.1
  Advertised path-id 1
