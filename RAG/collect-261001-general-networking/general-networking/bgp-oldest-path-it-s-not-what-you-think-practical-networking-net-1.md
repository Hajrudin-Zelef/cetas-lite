---
id: collect-261001-general-networking/general-networking/bgp-oldest-path-it-s-not-what-you-think-practical-networking-net-1
title: "bgp-oldest-path-it-s-not-what-you-think-practical-networking-net"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/bgp-oldest-path-it-s-not-what-you-think-practical-networking-net.md
source_anchor: ""
source_lines: [1, 234]
sha256: 637d927f6354d9207750e4ab99c5c2800a9ebfd42dc172a4009fb67c3fde748f
---

# bgp-oldest-path-it-s-not-what-you-think-practical-networking-net

BGP is the predominant dynamic routing protocol used to exchange routes between different autonomous systems. BGP’s popularity lies in the Path Selection process which allows extremely granular control of the path for incoming and outgoing traffic. One of the steps in this process states that if a tie still exists between two paths, BGP will prefer the Oldest Path. However, the BGP Oldest Path criteria does not work as most people think.

In this article, we will describe the BGP Oldest Path selection criteria, illustrate how it works, and explain specifically how it doesnât work as most people think.

To provide context, below is the full BGP Path Selection criteria that BGP uses to select a single best path among multiple known paths:

Path Selection:

1. If next-hop is inaccessible, drop the update
2. Prefer path with the largest Weight (Cisco proprietary)
3. Prefer path with the largest Local-Preference
4. Prefer path with locally originated routes vs externally learned
5. Prefer path with shortest AS-Path
6. Prefer path with best Origin (IGP > EGP > Incomplete)
7. Prefer path with lowest MED
8. Prefer path learned from eBGP over iBGP
9. Prefer path with the lowest IGP metric to the next-hop IP
10. **Prefer path with the greatest age (oldest path â eBGP only)**
11. Prefer path learned from neighbor with lowest Router ID
12. Prefer path learned from neighbor with lowest Neighbor IP Address

If two paths to a particular prefix exist, and all of the attributes in Steps 1-9 are identical, Step 10 will attempt to break the tie by preferring the oldest path.


### Topology

We will use the following topology to put the BGP Oldest Path criteria to the test. R5 in AS-55 will announce the 5.5.5.0/24 network, which will be shared via R2, R3, and R4 in AS-22, AS-33, and AS-44 (respectively). Each of these AS’s will then share the path to R1 in AS-11.

This is the initial configuration of each Router. All BGP adjacencies have been configured and R5 is announcing the only prefix in the BGP topology for the 5.5.5.0/24 network). R1 has learned of three paths to the 5.5.5.0/24 network â one via each of R2, R3, and R4.

*Click tabs to view initial configuration for each router.*

!
hostname **R1**
!
!
interface Loopback0
 ip address 1.1.1.1 255.255.255.0
!
interface Ethernet1/0
 description Link to R2 in AS-22
 ip address 9.22.11.1 255.255.255.0
!
interface Ethernet1/1
 description Link to R3 in AS-33
 ip address 9.33.11.1 255.255.255.0
!
interface Ethernet1/2
 description Link to R4 in AS-44
 ip address 9.44.11.1 255.255.255.0
!
!
**router bgp 11**
 bgp log-neighbor-changes
 neighbor 9.22.11.2 remote-as 22
 neighbor 9.33.11.3 remote-as 33
 neighbor 9.44.11.4 remote-as 44
!

!
hostname **R2**
!
!
interface Loopback0
 ip address 2.2.2.2 255.255.255.0
!
interface Ethernet0/0
 description Link to R5 in AS-55
 ip address 9.55.22.2 255.255.255.0
!
interface Ethernet1/0
 description Link to R1 in AS-11
 ip address 9.22.11.2 255.255.255.0
!
!
**router bgp 22**
 bgp log-neighbor-changes
 neighbor 9.22.11.1 remote-as 11
 neighbor 9.55.22.5 remote-as 55
!

!
hostname **R3**
!
!
interface Loopback0
 ip address 3.3.3.3 255.255.255.0
!
interface Ethernet0/1
 description Link to R5 in AS-55
 ip address 9.55.33.3 255.255.255.0
!
interface Ethernet1/1
 description Link to R1 in AS-11
 ip address 9.33.11.3 255.255.255.0
!
!
**router bgp 33**
 bgp log-neighbor-changes
 neighbor 9.33.11.1 remote-as 11
 neighbor 9.55.33.5 remote-as 55
!

!
hostname **R4**
!
!
interface Loopback0
 ip address 4.4.4.4 255.255.255.0
!
interface Ethernet0/2
 description Link to R5 in AS-55
 ip address 9.55.44.4 255.255.255.0
!
interface Ethernet1/2
 description Link to R1 in AS-11
 ip address 9.44.11.4 255.255.255.0
!
!
**router bgp 44**
 bgp log-neighbor-changes
 neighbor 9.44.11.1 remote-as 11
 neighbor 9.55.44.5 remote-as 55
!

!
hostname **R5**
!
!
**interface Loopback0
 ip address 5.5.5.5 255.255.255.0**
!
interface Ethernet0/0
 description Link to R2 in AS-22
 ip address 9.55.22.5 255.255.255.0
!
interface Ethernet0/1
 description Link to R3 in AS-33
 ip address 9.55.33.5 255.255.255.0
!
interface Ethernet0/2
 description Link to R4 in AS-44
 ip address 9.55.44.5 255.255.255.0
!
!
**router bgp 55**
 bgp log-neighbor-changes
 **network 5.5.5.0 mask 255.255.255.0**
 neighbor 9.55.22.2 remote-as 22
 neighbor 9.55.33.3 remote-as 33
 neighbor 9.55.44.4 remote-as 44
!

None of the path selection attributes have been modified, ensuring Steps 1-9 are all identical and do not explicitly prefer a path. Leaving only Steps 10-12 to break any ties that may exist.


### Initial Setup

R1 has three neighbor adjacencies, each of which announce a path to the 5.5.5.0/24 prefix. In order to deterministically set the age of each path for our experiment we will shut down all neighbor adjacencies:

R1(config)# **router bgp 11**
R1(config-router)# **neighbor 9.22.11.2 shutdown**
R1(config-router)# **neighbor 9.33.11.3 shutdown**
R1(config-router)# **neighbor 9.44.11.4 shutdown**
R1(config-router)#
*May  9 19:00:10.057: %BGP-5-ADJCHANGE: neighbor 9.22.11.2 Down Admin. shutdown
*May  9 19:00:10.058: %BGP-5-ADJCHANGE: neighbor 9.33.11.3 Down Admin. shutdown
*May  9 19:00:10.892: %BGP-5-ADJCHANGE: neighbor 9.44.11.4 Down Admin. shutdown

Then re-enable them one by one:

R1(config-router)# **no neighbor 9.22.11.2 shutdown**
***May  9 19:01:48.258: %BGP-5-ADJCHANGE: neighbor 9.22.11.2 Up**
R1(config-router)#
R1(config-router)# **no neighbor 9.33.11.3 shutdown**
***May  9 19:03:07.507: %BGP-5-ADJCHANGE: neighbor 9.33.11.3 Up**
R1(config-router)#
R1(config-router)# **no neighbor 9.44.11.4 shutdown**
***May  9 19:04:22.675: %BGP-5-ADJCHANGE: neighbor 9.44.11.4 Up**
R1(config-router)#

Between each âno shutdownâ command, we waited about a minute. This can be verified against the time stamp in the debug messages stating the neighbor adjacency came back up.

This gives us a topology where R1 has three neighbor adjacencies, with R2 being the oldest, followed by R3, followed by R4:

R1# **show ip bgp summary**
BGP router identifier 1.1.1.1, local AS number 11
...
Neighbor        V           AS MsgRcvd MsgSent   TblVer  InQ OutQ Up/Down  State/PfxRcd
**9.22.11.2**       4           22      11       9        7    0    0 **00:04:58**        1
**9.33.11.3**       4           33       9       9        7    0    0 **00:03:39**        1
**9.44.11.4**       4           44       8       8        7    0    0 **00:02:24**        1

R1 has three paths to the 5.5.5.0/24 prefix, with the path through R2 currently selected as the best path (9.22.11.2):

R1# **show ip bgp**
...
     Network          Next Hop            Metric LocPrf Weight Path
 *   5.5.5.0/24       9.44.11.4                              0 44 55 i
 *                    9.33.11.3                              0 33 55 i
 ***>                   9.22.11.2**                              0 22 55 i

R1# **show ip bgp 5.5.5.0/24**
BGP routing table entry for 5.5.5.0/24, version 7
Paths: (3 available, best #3, table default)
  Advertised to update-groups:
     3
  Refresh Epoch 2
  44 55
    9.44.11.4 from 9.44.11.4 (4.4.4.4)
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 2
  33 55
    9.33.11.3 from 9.33.11.3 (3.3.3.3)
      Origin IGP, localpref 100, valid, external
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 2
  22 55
    **9.22.11.2 from 9.22.11.2 (2.2.2.2)**
      Origin IGP, localpref 100, valid, external, **best**
      rx pathid: 0, tx pathid: 0x0

The best path is indicated by the “**>****show ip bgp**`best`” in theÂ **show ip bgp 5.5.5.5.0/24**


### BGP Oldest Path – Trial 1

Weâll begin by shutting down the present oldest path to the 5.5.5.0/24 prefix â the path through R2:

R1(config)# **router bgp 11**
R1(config-router)# **neighbor 9.22.11.2 shutdown**
*May  9 19:23:25.912: %BGP-5-ADJCHANGE: neighbor 9.22.11.2 Down Admin. shutdown

