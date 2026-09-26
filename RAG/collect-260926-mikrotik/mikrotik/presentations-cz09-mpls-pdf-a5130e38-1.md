---
id: collect-260926-mikrotik/mikrotik/presentations-cz09-mpls-pdf-a5130e38-1
title: "on R1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet", "voice"]
source: docs/RAG/lot-mikrotik/RouterOS/presentations-cz09-mpls-pdf-a5130e38.md
source_anchor: ""
source_lines: [1, 396]
sha256: 34d6753e2b21034a4f79d08f9a4dd33c0df12d7c92f1200915644c59c73736e2
---

# on R1

© MikroTik 2009  
MikroTik RouterOS
Introduction to MPLS
Prague
MUM Czech Republic 2009

© MikroTik 2009  
Q : W h y  h a v e n 't  y o u  h e a r d  
a b o u t  M P L S  b e f o r e ?
A: Probably because of the 
availability and/or price range

© MikroTik 2009  
Q : W h y  s h o u ld  y o u  c a r e  
a b o u t  M P L S  n o w ?
A: Probably because of the 
availability and/or price range...
A: ...and the reasons mentioned further in this presentation!

© MikroTik  2009
Networking
There are 3 networking methods available to 
manage computer networks:
Routing
Protocols: RIP, OSPF, BGP
Bridging
Protocols: STP, RSTP, MESH
Switching
Protocols: MPLS, ATM, Frame Relay

© MikroTik  2009
Concept of Switching

© MikroTik  2009
Switching
Switching is a network communications method 
that groups all transmitted data (no matter of 
content, type, or structure) into suitably-sized 
blocks
Each block is then transmitted over the network 
independently of each other
Network is capable of allocating transmission 
resources as needed, in this way optimizing 
utilization of link capacity and robustness of 
communication

© MikroTik  2009
MPLS
MPLS stands for Multi Protocol Label Switching
MPLS is a packet forwarding method based on 
labels attached to the packet and a label 
forwarding table with minimal lookup overhead
With MPLS the packet forwarding decision is no 
longer based on IP header and routing table
Efficiency of forwarding process is the main 
benefit of MPLS

© MikroTik  2009 8
MPLS Header
Also called Layer2.5 (because it is placed 
between OSI Layer2 and Layer3)
Header can consist of one or several 32bit shims:
Label (20 bits)
EXP (3 bits) – Class of Service
End of stack flag(1 bit) – is it last label? 
TTL (8 bits)
L3MPLSL2
TTLSEXPLabel

© MikroTik  2009
MPLS LDP
MPLS labels are assigned and distributed by 
the Label Distribution Protocol (LDP)
LDP requirements:
IP connectivity – properly configured IP routing 
(static,OSPF,RIP) between all hosts
“loopback” IP address that isn't attached to any real 
network interface (recommended)
Homogeneous MPLS cloud – all devices inside the 
MPLS cloud must have MPLS support

© MikroTik  2009 10
MPLS Basics
LER – Label Edge Router
LSR – Label Switch Router
LER
LSR
 LSR
LER
MPLS Backbone
Packets are classified and 
labeled at ingress LER LSRs forward packets 
using label swapping
Label is removed at 
egress LER
Packets are classified and 
labelled at ingress LER
IP packet

© MikroTik  2009 11
MPLS Benefits 
Increased scalability of the network 
Increased forwarding performance of the network 
Increased amount of possible VPN solutions that 
an ISP can offer to clients
Traffic engineering
Quality of Service
Redundancy and failover

© MikroTik  2009
BGP Scalability with MPLS
Traditionally you have 
to run BGP on all core 
routers
With MPLS, you only 
need to run BGP on 
network edges
Note: it is easy to migrate 
from routed backbone to 
MPLS enabled backbone
BGP
MPLS backbone
Routed backbone
BGP
C
C
C
E
E
C
C
C
E
E

© MikroTik  2009
MPLS enabled L2 VPNs
Layer2 service without the 
drawbacks of Layer2 network
Uses split-horizon method to 
prevent loops ( RSTP is not 
required)
New service is configured at 
the edge routers (no need to 
make changes to the network 
core)
Simpler to configure, easier 
to manage
Complete separation 
between providers network 
and customers network
MPLS backbone
PE1
PE2
CE2
CE1
Site 1
Site 2
Pseudo wire
Label stack
Customer L2 frame
L2 header

© MikroTik  2009 14
Split Horizon
Forward Ethernet frame coming from PE to connected CEs
Packets are not forwarded to interfaces with the same 
horizon value
Horizon value is set in bridge port configuration
   /interface bridge port 
      add bridge=vpn interface=vpls1 horizon=1
PE1
PE2
PE3
CE1
CE2
CE3
CE4

© MikroTik  2009
Current Layer2 VPNs
Customer based VPN
ISP
Site 1Site 2 CE
CE
GW
GW
Site 3
CE
GW
EoIP tunnels
Additional 
administration 
expenses
Big Overhead 
(Ethernet+GRE+IP)
Not very scalable
Each new site requires configuration of 
EoIP tunnels to every existing site
ISP is not involved

© MikroTik  2009
MPLS VPLS
Bandwidth 
improvements
Smaller Overhead 
(Ethernet+2 labels)
Can ask provider for 
guaranteed VPLS 
bandwidth
Each new site only requires correct PE 
configuration 
All the work is done by the ISP
ISP is ready to sell new type of service
ISP
Site 1
Site 2
CECE
GW GW
Site 3
CE
GW
PE
PE
PE
Provider based VPN service

© MikroTik  2009
Layer3 VPNs
VPN scalability
Each VPN has unique 
routing table (VRF table)
Customer IP address 
freedom (overlapping 
private IPs)
Can be set over existing 
BGP network 
MPLS cloud
VPN A
Site1
VPN A
Site2
VPN A
Site3
VPN B
Site1
VPN B
Site2
BGP peering

© MikroTik  2009
VRF Table
Means 'Virtual Routing and Forwarding Table'
VRF tables are similar to policy routing, except:
Each VRF table is independent - main routing table 
will not be used if VRF table fails to resolve route
BGP can be used to distribute routes between 
different VRF tables in the router

© MikroTik  2009
IP Routing Limitation
After two IP traffic flows for the same 
destination are merged, it is impossible to split 
them and reroute over different paths
Overloaded link from Router C to Router E
A
B
C
D
E
F
40Mbps traffic from A to F
40Mbps traffic from B to F

© MikroTik  2009
Traffic Engineering
TE tunnels can be used to shift the traffic load 
onto less utilized links
A
B
C
D
E
F
TE Tunnel1 50Mbps
TE Tunnel2 50Mbps

© MikroTik  2009
Customers Bandwidth Protection
Customers do not care how it is offered by the 
provider
With TE it is easy to deliver guaranteed 
bandwidth from point A to point B
Main bandwidth
Backup bandwidthProvider's network

© MikroTik  2009
Bandwidth Optimization
Separate tunnels for voice, video, or data
Backup tunnels over the third link
Backup Voice Tunnel
Backup Data Tunnel
Voice Tunnel
Data Tunnel

© MikroTik  2009
MPLS on RouterOS
Supported features
Static label binding for Ipv4
LDP for Ipv4
Virtual Private Lan Service
LDP based VPLS
MP-BGP based autodiscovery and signaling
RSVP TE Tunnels
OSPF extension for TE tunnels
Explicit path and CSPF path selection
Forwarding VPN traffic on TE tunnels
OSPF as CE-PE

© MikroTik  2009
MPLS on RotuerOS
Not yet supported
Ipv6
LDP features
Downstream on demand
Ordered label distribution protocol
RIP and iBGP as CE-PE protocols
TE features
Fast reroute
link/node protection
Full feature list at
http://wiki.mikrotik.com/wiki/MPLS

© MikroTik  2009
From EoIP to VPLS
Example: We have a routed network between 
R1, R2 and R3
EoIP tunnel is established between R1 and R3 
to guarantee Layer2 connectivity between Site 1 
and Site 2
Site 1 Site 2Routed network
EoIP
R1 R2 R3
lo:10.1.1.1 lo:10.1.1.2 lo:10.1.1.3
e1 e2 e1 e2 e1 e2

© MikroTik  2009
From EoIP to VPLS
Enable LDP
/mpls ldp 
 set enabled=yes lsr-id=10.1.1.x \
 transport-address=10.1.1.x
# on R1
/mpls ldp interface
 add interface=ether2
# on R2
/mpls ldp interface
  add interface=ether2
 add interface=ether2
# on R3
/mpls ldp interface
 add interface=ether1

© MikroTik  2009
From EoIP to VPLS
Configure VPLS
# on R1
/interface vpls add name=R1toR2 remote-peer=10.1.1.3 \
 vpls-id=10:10
/interface bridge port add bridge=vpn interface=R1toR2
# on R3
/interface vpls add name=R2toR1 remote-peer=10.1.1.1 \
 vpls-id=10:10
/interface bridge port add bridge=vpn interface=R2toR1

© MikroTik  2009
Speed tests
64 byte pps 512 byte pps
Bridging 414 000 359 000
MPLS 410 000 358 000
Routing 236 000 229 700
64 byte pps 512 byte pps
EoIP 190 000 183 900
VPLS 332 500 301 000
Label switching on RB1000
Site 1 Site 2
MPLS network
VPLS
Almost 2x faster 
than IP forwarding
The same speed 
as bridging
60% faster than 
EoIP tunnel over 
routed network

