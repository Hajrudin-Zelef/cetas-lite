---
id: collect-261001-cisco/cisco/redistribution-between-cisco-eigrp-into-ospf-and-vice-versa-example-2
title: "redistribution-between-cisco-eigrp-into-ospf-and-vice-versa-example"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/redistribution-between-cisco-eigrp-into-ospf-and-vice-versa-example.md
source_anchor: ""
source_lines: [180, 246]
sha256: 3a62fe8d7f0e0424b09274d4f00c100fe227df86901ec4f39f23588c7382ba7f
---

# redistribution-between-cisco-eigrp-into-ospf-and-vice-versa-example

Basic redistribution into OSPF is pretty simple, just remember the word **subnets** because if you issue the redistribution command without the subnets statement you will only redistribute classful routes. Again the redistribution must be done on ENT2 router that has both protocols.

__ENT2 Router__

router ospf 1                      < —-First go into OSPF routing instance

router-id 1.1.1.2

log-adjacency-changes

 **redistribute eigrp 1 subnets**           < —- Redistribute EIGRP into OSPF

network 1.1.1.0 0.0.0.3 area 0

network 10.10.0.0 0.0.255.255 area 0

Now, the routing table on ENT1 shows the EIGRP redistributed routes as E2 with a metric of 20

**ENT1#sh ip route ospf**O – OSPF, IA – OSPF inter area, E2 – OSPF external type 2

2.0.0.0/30 is subnetted, 1 subnets

**O E2 2.2.2.0 [110/20] via 1.1.1.2, 00:00:08, FastEthernet4**

10.0.0.0/32 is subnetted, 2 subnets

O 10.10.0.1 [110/2] via 1.1.1.2, 00:00:08, FastEthernet4

O 10.10.100.1 [110/2] via 1.1.1.2, 00:00:08, FastEthernet4

**O E2 192.168.1.0/24 [110/20] via 1.1.1.2, 00:00:08, FastEthernet4**

**O E2 192.168.100.0/24 [110/20] via 1.1.1.2, 00:00:08, FastEthernet4**

To verify that the remote EIGRP network is reachable you can issue a ping or a traceroute

**ENT1#ping 192.168.1.1**Type escape sequence to abort.

Sending 5, 100-byte ICMP Echos to 192.168.1.1, timeout is 2 seconds:

!!!!!

Success rate is 100 percent (5/5), round-trip min/avg/max = 1/1/4 ms

**ENT1#traceroute 192.168.1.1**Type escape sequence to abort.

Tracing the route to 192.168.1.1

VRF info: (vrf in name/id, vrf out name/id)

1 1.1.1.2 0 msec 0 msec 0 msec

2 2.2.2.2 0 msec 0 msec *

192.168.1.1 0 msec 0 msec

**ENT1#sh ip route 192.168.1.1**Routing entry for 192.168.1.0/24

Known via “ospf 1”, distance 110, metric 20, type extern 2, forward metric 1

Last update from 1.1.1.2 on FastEthernet4, 00:07:23 ago

Routing Descriptor Blocks:

* 1.1.1.2, from 1.1.1.2, 00:07:23 ago, via FastEthernet4

Route metric is 20, traffic share count is 1
