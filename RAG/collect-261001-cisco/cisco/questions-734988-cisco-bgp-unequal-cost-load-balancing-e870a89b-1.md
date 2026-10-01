---
id: collect-261001-cisco/cisco/questions-734988-cisco-bgp-unequal-cost-load-balancing-e870a89b-1
title: "Here are bgp neighbours from global routing table. Not relevant to the question. IP addresses are hidden"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "memory"]
source: docs/RAG/collect-261001-cisco/questions-734988-cisco-bgp-unequal-cost-load-balancing-e870a89b.md
source_anchor: ""
source_lines: [1, 140]
sha256: d4b3d10f5407ce2c8fea12247681976aa3c6b1bc250032b9175767bdbfc5dcbf
---

# Here are bgp neighbours from global routing table. Not relevant to the question. IP addresses are hidden

I'm trying to implement BGP Unequal Cost Load Balancing feature in my network. According to cisco manuals (long: http://www.cisco.com/c/en/us/td/docs/ios/12_2s/feature/guide/fsbgplb.html, short: https://ccieblog.co.uk/bgp/bgp-unequal-load-cost-sharing) I have built such net topology:
R1 - router where I'm trying to implement load balancing for outgoing traffic. VRF table with name nat is used.
R2-R4 - NAT servers running quagga, with default route to R5 shared with R1 over eBGP.
R1 IOS version: 12.2(33)SXJ4 (s72033-adventerprisek9_wan-mz.122-33.SXJ4.bin)
R2 configuration (R3 R4 only router-id and vlan differs)
In result I have 3 different default routes on R1 with same share count - 1/1 (1:1:1). But proportion 1:2:3 expexted:
R1# sh ip bgp vpnv4 vrf nat 0.0.0.0
Paths: (6 available, best #5, table nat)
Multipath: eiBGP
  Advertised to update-groups:
     2         
  65000
    10.30.227.227 from 10.30.227.227 (10.30.227.227)
      Origin IGP, localpref 100, valid, external, multipath
      Extended Community: RT:192.168.33.4:13
      DMZ-Link Bw 250 kbytes
  65000, (received-only)
    10.30.227.227 from 10.30.227.227 (10.30.227.227)
      Origin IGP, localpref 100, valid, external
      DMZ-Link Bw 250 kbytes
  65000
    10.30.228.228 from 10.30.228.228 (10.30.228.228)
      Origin IGP, localpref 100, valid, external, multipath
      Extended Community: RT:192.168.33.4:13
      DMZ-Link Bw 375 kbytes
  65000, (received-only)
    10.30.228.228 from 10.30.228.228 (10.30.228.228)
      Origin IGP, localpref 100, valid, external
      DMZ-Link Bw 375 kbytes
  65000
    10.30.225.225 from 10.30.225.225 (10.30.225.225)
      Origin IGP, localpref 100, valid, external, multipath, best
      Extended Community: RT:192.168.33.4:13
      DMZ-Link Bw 125 kbytes
  65000, (received-only)
    10.30.225.225 from 10.30.225.225 (10.30.225.225)
      Origin IGP, localpref 100, valid, external
      DMZ-Link Bw 125 kbytes
R1# sh ip cef vrf nat 0.0.0.0/0 internal
0.0.0.0/0, epoch 3, flags rib only nolabel, rib defined all labels, RIB[B], refcount 7, per-destination sharing
  sources: RIB, D/N, DRH
  feature space:
   NetFlow: Origin AS 0, Peer AS 0, Mask Bits 0
   Broker: linked
   IPRM: 0x00018000
  subblocks:
   DefNet source: 0.0.0.0/0
  ifnums:
   Vlan3225(231): 10.30.225.225
   Vlan3227(232): 10.30.227.227
   Vlan3228(233): 10.30.228.228
  path 541B7858, path list 53E3E0D8, share 1/1, type recursive nexthop, for IPv4, flags resolved
  recursive via 10.30.225.225[IPv4:nat], fib 5496C804, 1 terminal fib
    path 541B7BF8, path list 53E3E170, share 1/1, type adjacency prefix, for IPv4
    attached to Vlan3225, adjacency IP adj out of Vlan3225, addr 10.30.225.225 513F6B60
  path 541B78CC, path list 53E3E0D8, share 1/1, type recursive nexthop, for IPv4, flags resolved
  recursive via 10.30.227.227[IPv4:nat], fib 54969B7C, 1 terminal fib
    path 541B7B10, path list 53E3E08C, share 1/1, type adjacency prefix, for IPv4
    attached to Vlan3227, adjacency IP adj out of Vlan3227, addr 10.30.227.227 513F66E0
  path 541B7DC8, path list 53E3E0D8, share 1/1, type recursive nexthop, for IPv4, flags resolved
  recursive via 10.30.228.228[IPv4:nat], fib 54970EAC, 1 terminal fib
    path 541B79B4, path list 53E3E040, share 1/1, type adjacency prefix, for IPv4
    attached to Vlan3228, adjacency IP adj out of Vlan3228, addr 10.30.228.228 513F6560
  output chain:
    loadinfo 51283B80, per-session, 3 choices, flags 0003, 5 locks
    flags: Per-session, for-rx-IPv4
    15 hash buckets
      < 0 > IP adj out of Vlan3225, addr 10.30.225.225 513F6B60
      < 1 > IP adj out of Vlan3227, addr 10.30.227.227 513F66E0
      < 2 > IP adj out of Vlan3228, addr 10.30.228.228 513F6560
      < 3 > IP adj out of Vlan3225, addr 10.30.225.225 513F6B60
      < 4 > IP adj out of Vlan3227, addr 10.30.227.227 513F66E0
      < 5 > IP adj out of Vlan3228, addr 10.30.228.228 513F6560
      < 6 > IP adj out of Vlan3225, addr 10.30.225.225 513F6B60
      < 7 > IP adj out of Vlan3227, addr 10.30.227.227 513F66E0
      < 8 > IP adj out of Vlan3228, addr 10.30.228.228 513F6560
      < 9 > IP adj out of Vlan3225, addr 10.30.225.225 513F6B60
      <10 > IP adj out of Vlan3227, addr 10.30.227.227 513F66E0
      <11 > IP adj out of Vlan3228, addr 10.30.228.228 513F6560
      <12 > IP adj out of Vlan3225, addr 10.30.225.225 513F6B60
      <13 > IP adj out of Vlan3227, addr 10.30.227.227 513F66E0
      <14 > IP adj out of Vlan3228, addr 10.30.228.228 513F6560
    Subblocks:
     None
What am I doing wrong? According to manuals, different dmzlink bw values should cause different load sharing proportion, but in fact - it does not!
UPDATE 1 -- requested by user bangal
R1# show ip bgp all summary
For address family: IPv4 Unicast
BGP router identifier X.X.X.129, local AS number 41096
BGP table version is 22283352, main routing table version 22283352
34749 network entries using 4065633 bytes of memory
61661 path entries using 3206372 bytes of memory
8119/5337 BGP path/bestpath attribute entries using 1299040 bytes of memory
3752 BGP AS-PATH entries using 155474 bytes of memory
2990 BGP community entries using 138266 bytes of memory
146 BGP extended community entries using 5168 bytes of memory
53 BGP route-map cache entries using 1696 bytes of memory
0 BGP filter-list cache entries using 0 bytes of memory
BGP using 8871649 total bytes of memory
BGP activity 4716897/4682147 prefixes, 11331539/11269872 paths, scan interval 60 secs
# Here are bgp neighbours from global routing table. Not relevant to the question. IP addresses are hidden 
Neighbor     V       AS    MsgRcvd   MsgSent   TblVer  InQ OutQ Up/Down  State/PfxRcd
X.X.X.1      4       XX219    791704  760380 22283352    0    0 6d17h           1
X.X.X.33     4       XX219 112902498 1315655 22283352    0    0 6d17h           0
X.X.X.238    4       XX772    801422  762830 22283352    0    0 2w5d            0
X.X.X.206    4       XX540   2886112 1313917 22283352    0    0 4w4d         9641
X.X.X.70     4       XX772 188343075 1313853 22283352    0    0 6d14h       25881
X.X.X.78     4       XX772 148265282  941127 22283352    0    0 2w6d        26098
# Here are neighbours for vrf nat.
For address family: VPNv4 Unicast
BGP router identifier X.X.X.129, local AS number 41096
BGP table version is 824, main routing table version 824
1 network entries using 137 bytes of memory
6 path entries using 408 bytes of memory
1 multipath network entries and 3 multipath paths
8119/1 BGP path/bestpath attribute entries using 1299040 bytes of memory
3752 BGP AS-PATH entries using 155474 bytes of memory
2990 BGP community entries using 138266 bytes of memory
146 BGP extended community entries using 5168 bytes of memory
53 BGP route-map cache entries using 1696 bytes of memory
0 BGP filter-list cache entries using 0 bytes of memory
BGP using 1600189 total bytes of memory
3 received paths for inbound soft reconfiguration
BGP activity 4716897/4682147 prefixes, 11331539/11269872 paths, scan interval 15 secs
Neighbor        V          AS MsgRcvd MsgSent   TblVer  InQ OutQ Up/Down  State/PfxRcd
10.30.225.225   4       65000   11003   11443      824    0    0 3d18h           1
10.30.227.227   4       65000    9853   10293      824    0    0 3d18h           1
10.30.228.228   4       65000   10992   11432      824    0    0 3d18h           1
R1# sh ip route vrf nat
Routing Table: nat
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route
Gateway of last resort is 10.30.228.228 to network 0.0.0.0
     10.0.0.0/24 is subnetted, 4 subnets
C       10.30.0.0 is directly connected, Vlan30
