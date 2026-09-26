---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-3
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-3.md
source_anchor: ""
source_lines: [1, 46]
sha256: 6486de3c55264f2a2ce9f6b5e2b5eda36e281ea03c757371e7d8125128248033
---

# Summary

Bidirectional Forwarding Detection (BFD) is a low-overhead and short-duration protocol intended to detect faults in the bidirectional path between two forwarding engines, including physical interfaces, sub-interfaces, data link(s), and to the extent possible the forwarding engines themselves, with potentially very low latency. It operates independently of media, data protocols, and routing protocols.

BFD is basically a hello protocol for checking bidirectional neighbor reachability. It provides sub-second link failure detection support. BFD is not routing protocol specific, unlike protocol hello timers or such.

BFD Control packets are transmitted in UDP packets with destination port 3784, BFD also uses port 4784 for multihop paths. The source port is in the range 49152 through 65535. And BFD Echo packets are encapsulated in UDP packets with destination port 3785.

Standards and Technologies:

- RFC 5880 Bidirectional Forwarding Detection (BFD)
- RFC 5881 BFD for IPv4 and IPv6
- RFC 5882 Generic Application of BFD
- RFC 5883 Bidirectional Forwarding Detection (BFD) for Multihop Paths

# Features not yet supported

- echo mode
- enabling BFD for ip route gateways
- authentication

# Configuration

Allowing or forbidding BFD sessions can be done from the `/routing bfd configuration` menu. For example:

Configuration entries are order sensitive, which means that in the example above we are forbidding BFD sessions explicitly on the "sfp12" interface and allowing on the rest of the interfaces belonging to the "static" interface list.

To be able to filter multi-hop sessions, `addresses` or `address-list` properties can be used to match the destination, as well as the appropriate VRF, if a session is not running in the "main" VRF.

Everything else that is not explicitly listed in the configuration by default is forbidden.

# BFD with BGP

To enable the use of BFD for BGP sessions, enable `use-bfd` for required entries in `/routing bgp connection` menu.

A useful feature is that the BGP session will show that the BFD session for that particular BGP session is down:

# BFD with OSPF

To enable the use of BFD for OSPF neighbors, enable `use-bfd` for required entries in `/routing ospf interface-template` menu.

# Session Status

The status of the currently available sessions can be observed from `/routing bfd session` menu:

BFD is picking the highest value between the local tx interval and remote minimum rx interval as desired transmit interval. If the session is not established then desired minimum tx interval is set to 1 second.
