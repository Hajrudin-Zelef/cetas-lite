---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-27
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [3741, 3894]
sha256: 78a73120db29e77f11eec565c0d8a176b14cc6f8ba5c9e28889b45ca06c7e7d2
---

# A line starting with the # sign is comments.

l Tunnel binding policy: This policy binds a TE tunnel to a s specified VPN by binding a
specified destination address to the TE tunnel to provide QoS guarantee.
Auto Routes
An Interior Gateway Protocol (IGP) uses an auto route related to a CR-LSP in a TE tunnel that
functions as a logical link to calculate a path. The tunnel interface is used as an outbound interface
in the auto route. The TE tunnel is considered a P2P link with a specified metric value. The
following auto routes are supported:
l IGP shortcut: A route related to a CR-LSP is not advertised to neighbor nodes, preventing
other nodes from using the CR-LSP.
l Forwarding adjacency: A route related to a CR-LSP is advertised to neighbor nodes,
allowing these nodes to use the CR-LSP.
The forwarding adjacency advertises CR-LSP routes with neighbor IP addresses by sending
link-state advertisements (LSAs) or IS-IS link state packets (LSPs). Type 10 Opaque LSAs
carry the neighbor IP addresses in the Remote IP Address sub-type-length-value (sub-
TLV), and LSPs carry the neighbor IP addresses in intermediate system (IS) reachability
TLV's Remote IP Address sub-TLV.
If the forwarding adjacency is used, nodes on both ends of a CR-LSP must be in the same
area.
The following example demonstrates the IGP shortcut and forwarding adjacency.
Figure 3-18 Schematic diagram for IGP shortcut and forwarding adjacency
R7
R8
R1
R2
R3 R4
R5
R6
5
10
10
10
10
10
10
10
Destination Nexthop CostNode Mode
R2 R7 20
Forwarding 
Adjacency
R5
R7R1 30
R2 Tunnel0/0/1 10
R7
Tunnel0/0/1R1 20
R2 R4 25
IGP Shortcut
R5
R4R1 35
R2 Tunnel0/0/1 10
R7
Tunnel0/0/1R1 20
TE Metric=10
MPLS TE Tunnel0/0/1 
 
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
77

A CR-LSP over the path R7 → R6 → R2 is established on the network shown in Figure 3-18,
and the TE metric values are specified. When finding a route from R5 and R7 to R2 and R1
respectively, the result depends on whether the auto route is used.
l The auto route is not used. R5 uses R4 as the next hop in a route to R1 and a route to R2;
R7 uses R6 as the next hop in a route to R1 and a route to R2.
l The auto route is used. Either IGP shortcut or forwarding adjacency can be configured:
– The IGP shortcut is used to advertise the route of Tunnel0/0/1. R5 uses R4 as the next
hop in the route to R1 and the route to R2; R7 uses Tunnel0/0/1 as the next hop in the
route to R1 and the route to R2. R7, unlike R5, uses Tunnel0/0/1 in IGP path calculation.
– The forwarding adjacency is used to advertise the route of Tunnel0/0/1. R5 uses R7 as
the next hop in the route to R1 and the route to R2; R7 uses Tunnel0/0/1 as the next hop
in the route to R1 and the route to R2. Both R5 and R7 use Tunnel0/0/1 in IGP path
calculation.
3.2.7 Tunnel Re-optimization
An MPLS TE tunnel can be automatically reestablished over a new optimal path (if one exists)
if topology information is updated.
Background
MPLS TE tunnels are used to optimize traffic distribution over a network. An MPLS TE tunnel
is configured using static information, such as a bandwidth setting and a calculated path. Without
the optimization function, an MPLS TE tunnel cannot be automatically updated after the service
bandwidth or a tunnel management policy changes. This wastes network resources. MPLS TE
tunnels need to be optimized after being established.
Implementation
A specific event that occurs on the ingress can trigger optimization for a CR-LSP bound to an
MPLS TE tunnel. The optimization enables the CR-LSP to be reestablished over the optimal
path with the smallest metric.
NOTE
l Re-optimization is disabled by default. If enabled, re-optimization is performed every 3600 seconds
by default.
l The fixed filter (FF) reservation style and CR-LSP re-optimization cannot be configured together.
l Re-optimization cannot be performed for a CR-LSP that is established over an explicit path.
Re-optimization is classified into the following modes:
l Automatic re-optimization
When the interval at which a CR-LSP is optimized elapses, Constraint Shortest Path First
(CSPF) attempts to calculate a new path. If the calculated path has a metric smaller than
that of the existing CR-LSP, a new CR-LSP is set up over the new path. After the CR-LSP
is successfully set up, the ingress instructs the forwarding plane to switch traffic to the new
CR-LSP and tear down the original CR-LSP. Re-optimization is then complete. If the CR-
LSP fails to be set up, traffic is still forwarded along the existing CR-LSP.
l Manual re-optimization
A re-optimization command is run in the user view to trigger re-optimization.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
78

The Make-Before-Break mechanism is used to ensure uninterrupted service transmission
during the re-optimization process. Traffic must switch to a new CR-LSP before the original
CR-LSP is torn down.
3.2.8 MPLS TE Security
RSVP authentication verifies digest messages carried in RSVP messages to prevent attacks
initiated by modified or forged messages. Authentication enhancements can also be used to
prevent replay attacks and packet mis-sequence. RSVP authentication and its enhancements
improve MPLS TE security.
Background
RSVP uses raw IP to transmit packets. Raw IP has no security mechanism and is prone to attacks.
RSVP authentication can be used to verify packets based on keys to prevent attacks.
Original RSVP authentication, however, cannot prevent replay attacks or the problem of
neighbor relationship termination resulted from RSVP message mis-sequence. The RSVP
authentication enhancements are used to address this problem. The authentication lifetime,
handshake, and message window are added as enhanced functions. The authentication
enhancements improve security and RSVP neighborhood authentication in a harsh network
environment, such as network congestion.
Related Concepts
l Raw IP: similar to UDP but unreliable. No control is provided for raw IP. Whether raw IP
datagrams reach their destinations is uncertain.
l Spoofing attack: An unauthorized router establishes a neighbor relationship with a local
router or attacks the local router by generating pseudo RSVP messages to establish an RSVP
neighbor relationship. The pseudo RSVP messages can reserve lots of bandwidths.
l Replay attack: A remote router repeatedly sends a large number of packets with a sequence
number less than the maximum sequence number on a local router. After the local router
receives such RSVP packets, the local router terminates the RSVP neighbor relationship
with the remote router and tears down the CR-LSP.
Implementation
l Key authentication
RSVP authentication uses keys carried in packets exchanged between RSVP neighboring
nodes to verify those packets, preventing spoofing attacks. The same key must be
configured on two RSVP neighboring nodes before they perform RSVP authentication. A
local node uses Keyed-Hashing for Message Authentication Message Digest 5 (HMAC-
MD5) to calculate a digest for a key, adds this digest as an integrity object into an RSVP
message, and sends that message to the remote node. After the remote node receives the
message, the node uses the same key and algorithm to calculate a digest and checks whether
the local digest is the same as the received one. If they match, the remote node accepts the
message. If they do not match, the remote node discards the message.
l Authentication lifetime
Authentication lifetime specifies how long the RSVP neighbor relationship can last. It
provides the following functions:
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
79

