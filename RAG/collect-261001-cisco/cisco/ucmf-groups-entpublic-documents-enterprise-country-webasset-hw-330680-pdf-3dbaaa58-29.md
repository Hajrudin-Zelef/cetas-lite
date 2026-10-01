---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-29
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [4034, 4143]
sha256: 74301b408ef58b45951cf2b88308854d9b367afb6e19051bed12bf99a60ddf39
---

# A line starting with the # sign is comments.

3.2.9.2 Make-Before-Break
The Make-Before-Break mechanism prevents traffic loss during a traffic switchover between
two CR-LSPs. This mechanism improves MPLS TE tunnel reliability.
Background
If an MPLS TE tunnel is no longer the optimal path due to link attribute or tunnel attribute
changes, a new CR-LSP will be established according to new attributes. After the new CR-LSP
is established, traffic is switched to it. If traffic is switched away from the original MPLS TE
tunnel before the new CR-LSP is successfully established, traffic loss occurs. MPLS TE provides
the Make-Before-Break mechanism to prevent this problem.
Principles
Make-Before-Break is a mechanism that allows a CR-LSP to be established using changed
bandwidth and path attributes over a new path before the original CR-LSP is torn down. It helps
minimize data loss and additional bandwidth consumption. The new CR-LSP is called a modified
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
82

CR-LSP. Make-Before-Break is implemented using the shared explicit (SE) resource
reservation style.
The new CR-LSP competes with the original CR-LSP on some shared links for bandwidth. The
new CR-LSP cannot be established if it fails the competition. The Make-Before-Break
mechanism allows the system to reserve bandwidth used by the original CR-LSP for the new
CR-LSP, without calculating the bandwidth to be reserved. Additional bandwidth is used if links
on the new path do not overlap the links on the original path.
Figure 3-19 Schematic diagram for Make-Before-Break
R1 R2 R3 R4
R5
60M 60M 60M
60M60M
Path1
Path2
In this example, the maximum reservable bandwidth on each link is 60 Mbit/s on the network
shown in Figure 3-19. A CR-LSP along the path Path1 is established, with the bandwidth of 40
Mbit/s.
The path is expected to change to Path2 to forward data because R5 has a light load. The
reservable bandwidth of the link between R3 and R4 is just 20 Mbit/s. The total available
bandwidth for the new path is less than 40 Mbit/s. The Make-Before-Break mechanism can be
used in this situation. The Make-Before-Break mechanism allows the newly established CR-
LSP over the path Path2 to use the bandwidth of the original CR-LSP's link between R3 and R4.
After the new CR-LSP is established over the path, traffic switches to the new CR-LSP, and the
original CR-LSP is torn down.
In addition to the preceding method, another method of increasing the tunnel bandwidth can be
used. If the reservable bandwidth of a shared link increases to a certain extent, a new CR-LSP
can be established.
In the example shown in Figure 3-19, the maximum reservable bandwidth on each link is 60
Mbit/s. A CR-LSP along the path Path1 is established, with the bandwidth of 30 Mbit/s.
The path is expected to change to Path2 to forward data because R5 has a light load, and the
bandwidth is expected to increase to 40 Mbit/s. The reservable bandwidth of the link between
R3 and R4 is just 30 Mbit/s. The total available bandwidth for the new path is less than 40 Mbit/
s. The Make-Before-Break mechanism can be used in this situation. The Make-Before-Break
mechanism allows the newly established CR-LSP over the path Path2 to use the bandwidth of
the original CR-LSP's link between R3 and R4. The bandwidth of the new CR-LSP is 40 Mbit/
s, out of which 30 Mbit/s is released by the link between R3 and R4. After the new CR-LSP is
established, traffic switches to the new CR-LSP and the original CR-LSP is torn down.
Delayed Switchover and Deletion
If an upstream node on an MPLS network is busy but its downstream node is idle or an upstream
node is idle but its downstream node is busy, a CR-LSP may be torn down before the new CR-
LSP is established, causing a temporary traffic interruption.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
83

To prevent this temporary traffic interruption, the switching and deletion delays are used together
with the Make-Before-Break mechanism. In this case, traffic switches to a new CR-LSP a
specified delay time later after a new CR-LSP is established. The original CR-LSP is torn down
a specified delay later after a new CR-LSP is established. The switching delay and deletion delay
can be manually configured.
3.2.9.3 RSVP Hello
The RSVP Hello extension can rapidly monitor the reachability of RSVP nodes. If an RSVP
node becomes unreachable, TE FRR protection is triggered. The RSVP Hello extension can also
monitor whether an RSVP GR neighboring node is in the restart process.
Background
RSVP Refresh messages are used to synchronize path state block (PSB) and reservation state
block (RSB) information between nodes. They can also be used to monitor the reachability
between RSVP neighbors and maintain RSVP neighbor relationships.
Using Path and Resv messages to monitor neighbor reachability delays a traffic switchover if a
link fault occurs and therefore is slow. The RSVP Hello extension can address this problem.
Implementation
The principles of the RSVP Hello extension are as follows:
1. Hello handshake mechanism
Figure 3-20 Hello handshake mechanism
LSRA LSRB
Hello Repuest
Hello ACK
LSRA and LSRB are directly connected on the network shown in Figure 3-20.
l If RSVP Hello is enabled on LSRA, LSRA sends a Hello Request message to LSRB.
l After LSRB receives the Hello Request message and is also enabled with RSVP Hello,
LSRB sends a Hello ACK message to LSRA.
l After receiving the Hello ACK message, LSRA considers LSRB reachable.
2. Detecting neighbor loss
After a successful Hello handshake is implemented, LSRA and LSRB exchange Hello
messages. If LSRB does not respond to three consecutive Hello Request messages sent by
LSRA, LSRA considers router B lost and re-initializes the RSVP Hello process.
3. Detecting neighbor restart
If LSRA and LSRB are enabled with RSVP GR, and the Hello extension detects that LSRB
is lost, LSRA waits for LSRB to send a Hello Request message carrying a GR extension.
After receiving such a message, LSRA helps LSRB to restore the RSVP state and sends a
Hello ACK message to LSRB. After receiving the Hello ACK message, LSRB performs
the GR process and restores the RSVP soft state. LSRA and LSRB exchange Hello
messages to maintain the restored RSVP soft state.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
84

