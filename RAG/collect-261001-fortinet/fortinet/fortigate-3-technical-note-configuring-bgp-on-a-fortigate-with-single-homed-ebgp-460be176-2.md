---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-note-configuring-bgp-on-a-fortigate-with-single-homed-ebgp-460be176-2
title: "fortigate-3-technical-note-configuring-bgp-on-a-fortigate-with-single-homed-ebgp-460be176"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-note-configuring-bgp-on-a-fortigate-with-single-homed-ebgp-460be176.md
source_anchor: ""
source_lines: [359, 598]
sha256: 8ed14e6b95da8bf9ff6eaef179da51404658770f882c0b6c99f448030255f86b
---

# fortigate-3-technical-note-configuring-bgp-on-a-fortigate-with-single-homed-ebgp-460be176

**FGT-1 # get router info ospf neighbor**


OSPF process 0:

Neighbor ID Pri State Dead Time Address Interface

10.0.0.3 1 Full/DR 00:00:30 10.160.0.75 dmz



**FGT-1 # get router info routing-table all**

Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP

O - OSPF, IA - OSPF inter area

N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2

E1 - OSPF external type 1, E2 - OSPF external type 2

i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area

* - candidate default


B* 0.0.0.0/0 [20/0] via 10.142.0.110, wan1, 2d17h58m

B 1.0.0.0/8 [20/0] via 10.142.0.110, wan1, 2d17h58m

B 2.0.0.0/8 [20/0] via 10.142.0.110, wan1, 2d17h58m

S 10.0.0.1/32 [10/0] via 10.142.0.110, wan1

C 10.0.0.2/32 is directly connected, loopback

O 10.0.0.3/32 [110/110] via 10.160.0.75, dmz, 2d18h18m

C 10.142.0.0/23 is directly connected, wan1

O 10.143.0.0/23 [110/20] via 10.160.0.75, dmz, 2d18h18m

C 10.160.0.0/23 is directly connected, dmz




**FGT-2 # get router info bgp summary**

BGP router identifier 10.0.0.3, local AS number 1000

BGP table version is 1

2 BGP AS-PATH entries

0 BGP community entries


Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd

10.0.0.2 4 1000 4555 4563 0 0 0 2d18h28m 6


Total number of neighbors 1



**FGT-2 # get router info bgp neighbors**

BGP neighbor is 10.0.0.2, remote AS 1000, local AS 1000, internal link

BGP version 4, remote router ID 10.0.0.2

BGP state = Established, up for 2d18h28m

Last read 00:00:06, hold time is 180, keepalive interval is 60 seconds

Configured hold time is 180, keepalive interval is 60 seconds

Neighbor capabilities:

Route refresh: advertised and received (old and new)

Address family IPv4 Unicast: advertised and received

Received 4556 messages, 0 notifications, 0 in queue

Sent 4563 messages, 0 notifications, 0 in queue

Route refresh request: received 0, sent 0

Minimum time between advertisement runs is 30 seconds

Update source is loopback


For address family: IPv4 Unicast

BGP table version 1, neighbor version 0

Index 1, Offset 0, Mask 0x2

Community attribute sent to this neighbor (both)

6 accepted prefixes

0 announced prefixes


Connections established 1; dropped 0

Local host: 10.0.0.3, Local port: 179

Foreign host: 10.0.0.2, Foreign port: 1101

Nexthop: 10.0.0.3



**FGT-2 # get router info bgp network**

BGP table version is 1, local router ID is 10.0.0.3

Status codes: s suppressed, d damped, h history, * valid, > best, i - internal,

S Stale

Origin codes: i - IGP, e - EGP, ? - incomplete


Network Next Hop Metric LocPrf Weight Path

*>i0.0.0.0/0 10.142.0.110 0 100 0 1 ?

*>i1.0.0.0 10.142.0.110 0 100 0 1 ?

*>i2.0.0.0 10.142.0.110 0 100 0 1 ?

*>i10.0.0.2/32 10.0.0.2 0 100 0 ?

*>i10.142.0.0/23 10.0.0.2 0 100 0 ?

*>i10.160.0.0/23 10.0.0.2 0 100 0 ?


Total number of prefixes 6



**FGT-2 # get router info ospf neighbor**


OSPF process 0:

Neighbor ID Pri State Dead Time Address Interface

10.0.0.2 1 Full/Backup 00:00:38 10.160.0.74 dmz




**FGT-2 # get router info routing-table all**

Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP

O - OSPF, IA - OSPF inter area

N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2

E1 - OSPF external type 1, E2 - OSPF external type 2

i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area

* - candidate default


B* 0.0.0.0/0 [200/0] via 10.142.0.110 (recursive via 10.160.0.74), 2d18h00m

B 1.0.0.0/8 [200/0] via 10.142.0.110 (recursive via 10.160.0.74), 2d18h00m

B 2.0.0.0/8 [200/0] via 10.142.0.110 (recursive via 10.160.0.74), 2d18h00m

O 10.0.0.2/32 [110/110] via 10.160.0.74, dmz, 2d18h20m

C 10.0.0.3/32 is directly connected, loopback

O 10.142.0.0/23 [110/20] via 10.160.0.74, dmz, 2d18h20m

C 10.143.0.0/23 is directly connected, internal

C 10.160.0.0/23 is directly connected, dmz

**Troubleshooting**

 “get router info bgp <subcommand>”, where subcommand can be :


cidr-only                  display routes with non-natural netmasks

community display routes matching the communities

community-info list all bgp community information

community-list display routes matching the community-list

dampening display router dampening infomation

filter-list display routes conforming to the filter-list

inconsistent-as display routes with inconsistent AS Paths

neighbors show BGP neighbors

network show BGP info for network

network-longer-prefixes show BGP info for route and more specific routes

paths path information

prefix-list display routes conforming to the prefix-list

regexp display routes matching the AS path regular expression

quote-regexp display routes matching the AS path regular expression

route-map display routes conforming to the route-map

scan display BGP scan status

summary summary of BGP neighbor status

memory                     BGP memory table



 FGT# diag ip router bgp all enable (or “disable” to stop the trace)

And the CLI sniffer (see related article for using the sniffer)

**Related Articles**
