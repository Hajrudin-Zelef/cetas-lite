---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-17
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["advisory", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1733, 1870]
sha256: 280e307b2810cbc9a8698019f8a876da569a2b27b7d3f0ccbaa22068fa9a5da7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 To establish an LDP session between two devices, a link for setting up a TCP
                 connection must be established — this link is called an adjacency. After an
                 adjacency is established, the two devices exchange LDP session messages to
                 establish an LDP session. An LDP peer relationship is then established between the
                 two devices. Finally, the LDP peers exchange label information. The relationships
                 can be described as follows:

                 ●      LDP maintains peer relationships over adjacencies. The type of peer depends
                        on the type of adjacency.
                 ●      A peer can be maintained by more than one adjacency. If a peer is maintained
                        by both local and remote adjacencies, the peer is a coexistent local and
                        remote peer.
                 ●      An LDP session can be established only between LDP peers.

Types of LDP Messages
                 LDP uses the following types of messages:

                 ●      Discovery message: used to notify or maintain the presence of an LSR.
                 ●      Session message: used to establish, maintain, or terminate an LDP session
                        between LDP peers.
                 ●      Advertisement message: used to create, modify, or delete a mapping between
                        a specific FEC and label.
                 ●      Notification message: used to provide advisory information or error
                        information.

                 LDP transmits Discovery messages over UDP and transmits Session,
                 Advertisement, and Notification messages over TCP.

Label Spaces and LDP Identifiers
                 ●      Label space
                        A label space defines a range of labels distributed between LDP peers. Only
                        the global label space is supported. All interfaces on an LSR share a label
                        space.
                 ●      LDP identifier
                        An LDP identifier identifies a label space used by a specified LSR. An LDP
                        identifier consists of an LSR ID and a label space, and is 6 bytes long. An LDP
                        identifier is in the format of <LSR ID>:<Label space ID>.
                        –   LSR ID: indicates the LSR identifier and is 4 bytes long.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             29
MPLS Configuration
MPLS Configuration                                                                 3 MPLS LDP Configuration


                        –    Label space ID: indicates the identifier of the label space and is 2 bytes
                             long.

3.2.2 LDP NSR
                 LDP non-stop routing (NSR) ensures that connections in the control plane and
                 traffic in the forwarding plane are not interrupted if a fault (caused by a main
                 control board failure) occurs in the LDP control plane of the local node, even
                 without the help of neighboring nodes.

Context
                 As their networks rapidly develop, carriers have growing demand for IP network
                 reliability. Conventional non-stop forwarding (NSF) and graceful restart (GR)
                 techniques cannot prevent traffic interruptions if a peer does not support GR or
                 multiple peers fail simultaneously during a GR process. In addition, during the GR
                 switchover, the control flow is interrupted for a short time, during which the
                 topology change cannot be responded to. To solve this issue, NSR is introduced.
                 NSR is an innovation of, but essentially different from, NSF. NSR can be used to
                 ensure uninterrupted traffic transmission and retain connections in the control
                 plane if a software or hardware fault occurs in the control plane of a router. In
                 addition, the fault is transparent to the control planes of its peers.

Related Concepts
                 ●      High availability (HA): supports a backup channel between the active main
                        board (AMB) and slave main board (SMB).
                 ●      NSR: allows a standby control plane to take over traffic from an active control
                        plane if the active control plane fails, while preventing the control planes of
                        neighbor nodes from detecting the fault.
                 ●      NSF: enables a device to use the GR mechanism to ensure uninterrupted
                        service forwarding during an AMB/SMB switchover.
                 ●      AMB and SMB: implement control plane processes.

Implementation
                 LDP NSR implements LDP data synchronization between the AMB and SMB. The
                 AMB backs up the following LDP data to the SMB:
                 ●      LSP forwarding entries
                 ●      Key resources, such as labels and cross connection (XC) entries
                 ●      LDP control blocks
                         NOTE

                        The active and standby control planes of an NSR-capable device must run on the AMB and
                        SMB, respectively.

                 Synchronous backup and switchover processing are as follows:
                 1.     Batch backup
                        The AMB backs up LDP data in a batch to the SMB immediately after the
                        SMB starts.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  30
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


                 2.     Real-time backup
                        The SMB backs up data and LDP packets sent by the AMB in real time,
                        implementing real-time data synchronization.
                 3.     Switchover
                        If the AMB fails, the SMB takes over services. Because data is synchronized
                        between the AMB and SMB, services in the control plane and forwarding
                        plane are not interrupted.


Other Functions
                 An NSR device can function as a GR Helper. The GR Helper communicates with
                 NSR-disabled devices and responds to neighbor nodes' GR requests during an
                 AMB/SMB switchover.


Application Scenario
                 ●      When a single node (such as PE3 in Figure 3-2) is connected to multiple links,
                        load balancing is implemented among the links to protect services. However,
                        if the main control board on the node fails, services are interrupted. To
                        address this issue, NSR needs to be configured on the node to enhance
                        reliability. Figure 3-2 shows a typical LDP NSR application.

                        Figure 3-2 Typical LDP NSR application




                 ●      NSR can be deployed on a high-capacity and heavily loaded network to
                        minimize the impact of a device's control-plane protocol fault on the entire
                        network, thereby preventing route flapping.


Benefits
                 ●      If a fault occurs in the LDP control plane, services in the control plane and
                        forwarding plane are not interrupted.
                 ●      A node can implement NSR independently, without the assistance of
                        neighboring nodes. When a multi-point fault occurs in the control plane of an
                        MPLS network, the active/standby switchover can still be performed on each
                        node.

