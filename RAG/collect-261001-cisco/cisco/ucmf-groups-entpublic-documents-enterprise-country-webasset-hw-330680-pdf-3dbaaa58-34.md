---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-34
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [4826, 5013]
sha256: f99e6ae88a8c492069e7fdfe5542df28945a3ac6a6b102c9dbdda47d378f1169
---

# A line starting with the # sign is comments.

l Protection switchover: switches traffic from a faulty working tunnel to a protection tunnel
in a tunnel protection group, which improves network reliability.
Figure 3-27 shows a tunnel protection group.
Figure 3-27 Tunnel protection group
Data flow when  primary
tunnel is normal
Data flow when primary
tunnel is failed
LSRA LSRB
Working tunnel-1
Protection tunnel-3
Working tunnel-1 is failed
 
Primary tunnels tunnel-1, and the bypass tunnel tunnel-3 are established on the ingress LSRA
shown in Figure 3-27. Tunnel-3 is specified as a protection tunnel for primary CR-LSPs tunnel-1
on LSRA. If the configured fault detection mechanism on the ingress detects a fault in tunnel-1,
traffic switches to tunnel-3. LSRA attempts to reestablish tunnel-1. If tunnel-1 is successfully
established, traffic switches back to the primary CR-LSP.
Implementation
A TE tunnel protection group uses a configured protection tunnel to protect traffic on the working
tunnel to improve tunnel reliability. To ensure the improved performance of the protection
tunnel, the protection tunnel must exclude links and nodes through which the working tunnel
passes during network planning.
Table 3-18 describes the implementation procedure of a tunnel protection group.
Table 3-18 Implementation procedure of a tunnel protection group
Process Description
Establish
ment
The working and protection tunnels must have the same ingress and egress. The
protection tunnel is established in the same procedure as a regular tunnel. The
protection tunnel can use attributes that differ from those for the working tunnel.
Ensure that the working and protection tunnels are established over different
paths as much as possible.
NOTE
l A protection tunnel cannot be protected or enabled with TE FRR.
l Attributes for a protection tunnel can be configured independently of those for the
working tunnel, which facilitates the network planning.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
98

Process Description
Binding
between
the
working
and
protection
tunnels
The protection tunnel is bound to the tunnel ID of the working tunnel so that the
two tunnels form a tunnel protection group.
Fault
detection
In addition to MPLS TE's own detection mechanism, MPLS OAM and BFD
for CR-LSP are used to detect faults in a tunnel protection group to speed up
protection switching.
Protection
switching
The tunnel protection group supports either of the following protection
switching modes:
l Manual switching: Traffic is forcibly switched to the protection tunnel.
l Automatic switching: Traffic automatically switches to the protection tunnel
if the working tunnel fails.
A time interval can be set for automatic switching.
Switchbac
k
After a traffic switchover is implemented, the ingress attempts to reestablish the
working tunnel. If the working tunnel is reestablished, the ingress can switch
traffic back to the working tunnel or still forward traffic over the protection
tunnel.
 
Other Usage
A tunnel protection group works in either 1:1 or N:1 mode. The 1:1 mode enables a protection
tunnel to protect only a single working tunnel. The N:1 mode enables a protection tunnel to
protect more than one working tunnel.
Figure 3-28 N:1 protection mode
LSRA LSRB
Working tunnel-1
Working tunnel-2
Protection tunnel-3
Data flow when  primary
tunnel is normal
Data flow when primary
tunnel is failed
 
Differences Between CR-LSP Backup and a Tunnel Protection Group
CR-LSP backup and a tunnel protection group are both E2E protection mechanisms for MPLS
TE. Table 3-19 shows the comparison between these two mechanisms.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
99

Table 3-19 Comparison between CR-LSP backup and a tunnel protection group
Item CR-LSP Backup Tunnel Protection Group
Object to be
protected
Primary and backup CR-LSPs
are established on the same
tunnel interface. A backup CR-
LSP protects traffic on a primary
CR-LSP.
One tunnel protects traffic over
another tunnel in a tunnel protection
group.
TE FRR A primary CR-LSP supports TE
FRR. A backup CR-LSP does not
support TE FRR.
A working tunnel supports TE FRR. A
protection tunnel does not support TE
FRR.
LSP attributes Primary and backup CR-LSPs
have the same attributes, except
for the TE FRR attribute. In
addition, the bandwidth for the
backup CR-LSP can be set
separately.
The attributes of one tunnel in a tunnel
protection group are independent of
the attributes of the other tunnel. For
example, a protection tunnel with no
bandwidth can protect traffic on a
working tunnel that has a bandwidth.
Protection mode The 1:1 protection mode is
supported. Each primary CR-
LSP is protected by a backup CR-
LSP.
Apart from the 1:1 protection mode,
the N:1 protection mode is supported.
Multiple working tunnels share one
protection tunnel. When any one
working tunnel fails, data is switched
to the protection tunnel.
 
3.2.9.8 BFD for MPLS TE
Bidirectional forwarding detection (BFD) can monitor MPLS TE tunnels, CR-LSPs bound to
the MPLS TE tunnels, and RSVP neighbor relationships. If BFD detects a fault, the BFD module
instructs the MPLS module to perform a traffic switchover, improving network reliability.
Background
TE FRR, CR-LSP backup, and tunnel protection groups can be used to improve the reliability
of MPLS TE networks. A fault, however, occurs if no message arrives after the refresh period
of RSVP Hello or RSVP messages, which leads to a slow detection speed. When a Layer 2
device (such as a switch or hub) exists on the faulty link, slow detection delays a traffic
switchover and causes some traffic to be dropped. BFD can send packets to quickly detect faults
in MPLS TE tunnels and trigger a rapid traffic switchover to minimize traffic loss.
Related Concepts
BFD sessions are classified into the following types:
l Static BFD session: Local and remote discriminators are configured manually.
l Dynamic BFD session: Local and remote discriminators are allocated automatically.
NOTE
For details about BFD, see the chapter "BFD" in the Feature Description - Reliability.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
100

Implementation
The following BFD functions are supported for MPLS TE:
l BFD for RSVP
BFD monitors RSVP. BFD can detect faults in links between RSVP neighboring nodes in
milliseconds. BFD for RSVP applies to a TE FRR network, on which Layer 2 devices exist
between the PLR and its RSVP neighboring nodes over the primary CR-LSP.
l BFD for CR-LSP
BFD monitors CR-LSPs. After BFD detects a fault in a CR-LSP, the BFD module
immediately instructs the forwarding plane to trigger a rapid traffic switchover. BFD for
CR-LSP is used together with a hot-standby CR-LSP or a tunnel protection group.
l BFD for TE tunnel
BFD can monitor MPLS TE tunnels that are used as public network tunnels to transmit
VPN traffic. BFD monitors a whole TE tunnel. If BFD detects a fault in a tunnel that
transmits private network traffic, the BFD module instructs the VPN or virtual leased line
(VLL) FRR module to perform a traffic switchover.
BFD for RSVP
When a Layer 2 device exists between RSVP neighboring nodes, the two nodes can detect a link
fault only using the Hello mechanism in seconds. This process results in the loss of lots of data.
BFD monitors RSVP neighbor relationships. BFD for RSVP rapidly detects faults in a link
between RSVP neighboring nodes within milliseconds. BFD for RSVP applies to TE FRR
networks, on which Layer 2 devices exist on a primary CR-LSP between the PLR and its RSVP
neighboring node, as shown in Figure 3-29.
Figure 3-29 BFD for RSVP
BFD Session
BFD Session
BFD Session
BFD Session
 
