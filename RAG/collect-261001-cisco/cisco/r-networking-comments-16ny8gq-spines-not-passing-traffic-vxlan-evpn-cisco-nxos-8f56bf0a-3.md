---
id: collect-261001-cisco/cisco/r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a-3
title: "r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a.md
source_anchor: ""
source_lines: [148, 217]
sha256: cc9a0bb0d155ce84b62a78a2ff3cab68cd447acd12466a2d12a30f2001c01170
---

# r-networking-comments-16ny8gq-spines-not-passing-traffic-vxlan-evpn-cisco-nxos-8f56bf0a

  Path type: internal, path is valid, is best path, no labeled nexthop
             Imported from 10.255.50.103:33322:[2]:[0]:[0]:[48]:[6cfe.545d.e870]:[32]:[10.55.2.9]/272
  AS-Path: NONE, path sourced internal to AS
    10.254.50.201 (metric 3) from 10.255.50.1 (1.1.1.1)
      Origin IGP, MED not set, localpref 100, weight 0
      Received label 10555 50999
      Extcommunity: RT:64550:10555 RT:64550:50999 SOO:10.254.50.201:0 ENCAP:8
          Router MAC:9088.555d.2453
      Originator: 10.255.50.103 Cluster list: 1.1.1.1
  Path-id 1 not advertised to any peer
switch507#
The NVE interface is clearly transmitting but NOT receiving anything:
switch507# sho int nve1
nve1 is up
admin state is up,  Hardware: NVE
  MTU 9216 bytes
  Encapsulation VXLAN
  Auto-mdix is turned off
  RX
    ucast: 0 pkts, 0 bytes - mcast: 103 pkts, 14098 bytes
  TX
    ucast: 29825 pkts, 4528536 bytes - mcast: 1851 pkts, 220684 bytes
switch507#
switch503# sh int nve1
nve1 is up
admin state is up,  Hardware: NVE
  MTU 9216 bytes
  Encapsulation VXLAN
  Auto-mdix is turned off
  RX
    ucast: 0 pkts, 0 bytes - mcast: 98 pkts, 13436 bytes
  TX
    ucast: 17998 pkts, 2735696 bytes - mcast: 1684 pkts, 199586 bytes
switch503#
I feel like I'm missing really simple. I believe I have the appropriate features enabled on the leaves and spines and the BGP configs are dead simple.
leaf:
switch507# sho run | grep overlay
nv overlay evpn
feature nv overlay
switch507# sho run | sec feature
feature vrrp
feature ospf
feature bgp
feature fabric forwarding
feature interface-vlan
feature vn-segment-vlan-based
feature lacp
feature dhcp
feature vpc
feature nv overlay
switch507#
spine:
coresw501(config)# sho run | sec feature
feature ospf
feature bgp
feature fabric forwarding
feature nv overlay
coresw501(config)# sho run | grep overlay
nv overlay evpn
feature nv overlay
coresw501(config)#
I have far less hair than when I started this (especially since this same basic config is working in CML 2.0 with NX9000v.... *sigh*
Thanks for any and all help that you can provide!
EDIT: Formatting. Why is it so bad? It look right on the original post I was making.. GAH!
Section des commentaires
I Dumped the ICMP packet at the spine and I definitely see the VNI and expected src/dst.
Ethernet II, Src: 90:88:55:5d:24:53, Dst: 24:2a:04:c0:32:f3 Destination: 24:2a:04:c0:32:f3 Address: 24:2a:04:c0:32:f3 .... ..0. .... .... .... .... = LG bit: Globally unique address (factory default) .... ...0 .... .... .... .... = IG bit: Individual address (unicast) Source: 90:88:55:5d:24:53 Address: 90:88:55:5d:24:53 .... ..0. .... .... .... .... = LG bit: Globally unique address (factory default) .... ...0 .... .... .... .... = IG bit: Individual address (unicast) Type: IPv4 (0x0800) Internet Protocol Version 4, Src: 10.254.50.201, Dst: 10.254.50.203 0100 .... = Version: 4 .... 0101 = Header Length: 20 bytes (5) Differentiated Services Field: 0x00 (DSCP: CS0, ECN: Not-ECT) 0000 00.. = Differentiated Services Codepoint: Default (0) .... ..00 = Explicit Congestion Notification: Not ECN-Capable Transport (0) Total Length: 134 Identification: 0x55aa (21930) Flags: 0x0000 0... .... .... .... = Reserved bit: Not set .0.. .... .... .... = Don't fragment: Not set ..0. .... .... .... = More fragments: Not set Fragment offset: 0 Time to live: 255 Protocol: UDP (17) Header checksum: 0xea2c [validation disabled] [Header checksum status: Unverified] Source: 10.254.50.201 Destination: 10.254.50.203 User Datagram Protocol, Src Port: 9834, Dst Port: 4789 Source Port: 9834 Destination Port: 4789 Length: 114 [Checksum: [missing]] [Checksum Status: Not present] [Stream index: 0] [Timestamps] [Time since first frame: 5.141114815 seconds] [Time since previous frame: 1.024004944 seconds] Virtual eXtensible Local Area Network Flags: 0x0800, VXLAN Network ID (VNI) 0... .... .... .... = GBP Extension: Not defined .... .... .0.. .... = Don't Learn: False .... 1... .... .... = VXLAN Network ID (VNI): True .... .... .... 0... = Policy Applied: False .000 .000 0.00 .000 = Reserved(R): 0x0000 Group Policy ID: 0 VXLAN Network Identifier (VNI): 10555 Reserved: 0 Ethernet II, Src: 6c:fe:54:5d:e8:70, Dst: 6c:fe:54:5d:e4:f4 Destination: 6c:fe:54:5d:e4:f4 Address: 6c:fe:54:5d:e4:f4 .... ..0. .... .... .... .... = LG bit: Globally unique address (factory default) .... ...0 .... .... .... .... = IG bit: Individual address (unicast) Source: 6c:fe:54:5d:e8:70 Address: 6c:fe:54:5d:e8:70 .... ..0. .... .... .... .... = LG bit: Globally unique address (factory default) .... ...0 .... .... .... .... = IG bit: Individual address (unicast) Type: IPv4 (0x0800)
Spine can definitely see routes to the src/dest on the outer IP header:
IP Route Table for VRF "default" '*' denotes best ucast next-hop '**' denotes best mcast next-hop '[x/y]' denotes [preference/metric] '%<string>' in via output denotes VRF <string> 10.254.50.203/32, ubest/mbest: 1/0 *via 10.255.50.107, Eth1/4.1530, [110/2], 1d07h, ospf-UNDERLAY, intra coresw501# show ip route 10.254.50.201 IP Route Table for VRF "default" '*' denotes best ucast next-hop '**' denotes best mcast next-hop '[x/y]' denotes [preference/metric] '%<string>' in via output denotes VRF <string> 10.254.50.201/32, ubest/mbest: 1/0 *via 10.255.50.103, Eth1/2.1530, [110/2], 1d07h, ospf-UNDERLAY, intra coresw501#
Src and Dest MACs in the Ethernet header under the VxLAN header are in BGP:
