---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-25
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [2626, 2752]
sha256: 98eaa8c25e78157fe7d605e6c29e495685e9d169a6d555988d084ddbddbc1bb2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Context
                    With rapid network development, carriers have increasing demand for IP network
                    reliability. Conventional non-stop forwarding (NSF) and graceful restart (GR)
                    techniques cannot prevent traffic interruptions if a peer does not support GR or
                    multiple peers fail simultaneously during a GR process. Traffic is interrupted
                    temporarily before the GR process is complete because a GR-enabled router
                    cannot obtain routing information from or establish control plane connections to
                    its peers during the GR process. This is where NSR comes into play. NSR is an
                    innovation, compared with NSF. NSR can be used to ensure uninterrupted traffic
                    transmission and retain control plane connections if a software or hardware fault
                    occurs on the control plane of a device. In addition, the fault is transparent to the
                    control planes of its peers.

Related Concepts
                    BGP NSR involves the following concepts:
                    ●   High availability (HA): supports data backup between the active and slave
                        main boards.
                    ●   NSR: a routing technique that prevents neighbors or peers from detecting a
                        control plane protocol fault on a device with a backup control plane. If a fault
                        occurs, NSR ensures that the neighbor or peer relationships set up through
                        specific routing protocols, as well as the sessions of signaling protocols and
                        the protocols that are used to meet service requirements, are not interrupted.
                    ●   NSF: enables a device to use the GR mechanism to ensure uninterrupted
                        service forwarding during an active/slave main board switchover.
                    ●   Active main board (AMB) and slave main board (SMB): boards that are used
                        to carry control plane processes.

Implementation
                    The AMB backs up VPN data to the standby main board (SMB) on a specific node
                    to implement NSR. The following key VPN data is synchronized between the AMB
                    and SMB:
                    ●   VPN routes, including:
                        –    Routes imported by running the import-route or network command in
                             the BGP VPN instance IPv4 address family view or BGP VPN instance IPv6
                             address family view.
                        –    Locally leaked routes
                        –    Remotely leaked routes
                    ●   Attributes carried by routes
                    ●   VPN labels

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               39
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


                    ●   Next-hop information about routes

                         NOTE

                        The active and standby control planes of an NSR-capable device must run on the AMB and
                        SMB, respectively.

                    NSR is implemented as follows:
                    ●   Batch backup: The AMB backs up data in batches to the SMB immediately
                        after the SMB starts.
                    ●   Real-time backup: During service running, the AMB backs up data in real time
                        to the SMB. Moreover, the AMB and SMB both receive packets.
                    ●   Switchover: If the AMB fails, the SMB takes over services. Because data is
                        synchronous between the AMB and SMB, neither the control plane nor the
                        forwarding plane is interrupted.


Related Functions
                    An NSR-enabled device can function as a GR helper to communicate with NSR-
                    incapable devices and respond to peers' GR help requests during an AMB/SMB
                    switchover on the device.


Benefits
                    ●   NSR ensures that neither the control plane nor the forwarding plane is
                        interrupted when a fault occurs on the VPN control plane.
                    ●   A node can implement NSR independently, without the assistance of
                        neighboring nodes. NSR can help multiple MPLS nodes implement AMB/SMB
                        switchovers if the control planes of these nodes fail.

3.2.3 Inter-AS VPN
                    With the wide application of IPv4 L3VPN solutions, different MANs of a carrier or
                    collaborating backbone networks of different carriers frequently span multiple
                    ASs.

                    In most cases, an IPv4 L3VPN architecture runs within an AS in which VPN routing
                    information is flooded on demand. The VPN routing information within the AS
                    cannot be flooded to other ASs. To implement exchange of VPN routes between
                    different ASs, the inter-AS IPv4 L3VPN model is used. The model is an extension to
                    the basic IPv4 L3VPN framework. Through this model, route prefixes and labels
                    can be advertised over links between different carrier networks.

                    The following inter-AS VPN solutions are proposed in related standards:

                    ●   Inter-Provider Backbones Option A (inter-AS VPN Option A): VPN instances
                        spanning multiple ASs manage their own VPN routes through dedicated
                        interfaces between ASBRs. This solution is also called VRF-to-VRF.
                    ●   Inter-Provider Backbones Option B (inter-AS VPN Option B): ASBRs advertise
                        labeled VPN-IPv4 routes to each other through MP-EBGP. This solution is also
                        called EBGP redistribution of labeled VPN-IPv4 routes.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  40
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


3.2.4 Label Allocation Modes of IPv4 L3VPN over MPLS
Context
                    In IPv4 L3VPN networking, a device assigns an MPLS label to each VPN instance
                    (known as one-label-per-instance) by default. On a network with a large number
                    of VPN routes, the one-label-per-instance mode helps conserve label resources
                    and reduce capacity requirements for PEs. Table 3-1 lists three label allocation
                    modes that are currently available.

                    Table 3-1 Comparison of label allocation modes
                     Mode     Definition               Applicable                  Configuration
                                                       Networking                  Location

                     One-     All the VPN routes       It is applicable to all     Devices on which VPN
                     label-   from a VPN instance      types of IPv4 L3VPN         instances are
                     per-     are assigned the same    networking.                 configured
                     instan   label.
                     ce

                     One-     Each VPN route is        It is applicable to all     Devices on which VPN
                     label-   assigned a label.        types of IPv4 L3VPN         instances are
                     per-                              networking.                 configured
                     route

