---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-65
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [9100, 9228]
sha256: aa2e466d49502bc9f85e3127cb8174d5002b2e317d1b01d7da8d26011f0d3aee
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.18.1 Understanding LDP-IGP Synchronization
Context
                 LDP-IGP synchronization is used to synchronize the states of LDP and an IGP so as
                 to minimize the traffic loss time if a fault occurs on the network.
                 On a network with primary and backup links, if the primary link fails, both IGP
                 and LSP traffic is switched to the backup link. After the primary link recovers, IGP
                 routes are switched back to the primary link before LDP convergence is complete.
                 Before being established, the LSP along the primary link takes time to make
                 preparations, such as adjacency restoration, which leads to traffic drop. If an LDP
                 session or adjacency on the primary link fails, the LSP along the primary link is
                 deleted. However, the IGP still uses the primary link. As a result, LSP traffic cannot
                 be switched to the backup link, and is continuously dropped.

                         NOTE

                        LDP-IGP synchronization supports only OSPFv2 and IS-IS in IPv4.

                 LDP-IGP synchronization allows you to set an IGP cost to delay a route switchback
                 until LDP convergence is complete. Specifically, before the LSP on the primary link
                 is established, the LSP on the backup link is retained to continue forwarding
                 traffic. The LSP on the backup link is torn down only after the LSP on the primary
                 link is established successfully.
                 LDP-IGP synchronization timers are as follows:
                 ●      Hold-down timer
                 ●      Hold-max-cost timer
                 ●      Delay timer




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                              152
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


Implementation

                 Figure 3-30 Switchback with LDP-IGP synchronization enabled




                 ●      The network shown in Figure 3-30 has primary and backup links. When the
                        primary link recovers from a fault, traffic is switched from the backup link
                        back to the primary link. During the traffic switchback, the backup LSP cannot
                        be used once IGP route convergence is complete. As a result, before a new
                        LSP is set up over the primary link, traffic is dropped. To prevent this issue,
                        LDP-IGP synchronization can be configured to delay IGP route switchback
                        until LDP convergence is complete. Before convergence of the primary LSP
                        completes, the backup LSP is retained to continue to forward traffic. The
                        backup LSP is torn down after the primary LSP is successfully established. The
                        detailed process is as follows:
                        a.   The link fault is rectified.
                        b.   The IGP advertises the maximum cost of the primary link, delaying the
                             IGP route switchback.
                        c.   Traffic is still forwarded over the backup LSP.
                        d.   After the LDP session and adjacency are established, Label Mapping
                             messages are exchanged to instruct the IGP to start synchronization.
                        e.   The IGP advertises the normal cost of the primary link, and its routes
                             converge to the original forwarding path. The LSP is reestablished and
                             entries are delivered to the forwarding table (within milliseconds).
                 ●      If an LDP session or adjacency on the primary link fails, the LSP along the
                        primary link is deleted. However, the IGP still uses the primary link. As a
                        result, LSP traffic cannot be switched to the backup link, and is continuously
                        dropped. To prevent this issue, LDP-IGP synchronization can be configured. If
                        an LDP session or adjacency fails, LDP informs the IGP that the LDP session or
                        adjacency is faulty. In this case, IGP advertises the maximum cost of the faulty
                        link. In this way, both the route and LSP are switched to the backup link. The
                        process is as follows:
                        a.   The LDP session or adjacency between nodes on the primary link is faulty.
                        b.   LDP informs the IGP that the LDP session or adjacency along the primary
                             link is faulty. IGP advertises the maximum cost of the primary link.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             153
MPLS Configuration
MPLS Configuration                                                                     3 MPLS LDP Configuration


                        c.   The IGP route switches to the backup link.
                        d.   An LSP is set up over the backup link, and then forwarding entries are
                             delivered.
                        To prevent a continuous failure to reestablish the LDP session or adjacency,
                        you can configure the Hold-max-cost timer to permanently advertise the
                        maximum cost so that traffic is always transmitted over the backup link until
                        the LDP session and LDP adjacency are reestablished on the primary link.
                 ●      LDP-IGP synchronization state transition mechanism
                        After LDP-IGP synchronization is enabled on an interface, the IGP queries the
                        state of the interface, LDP session, and LDP adjacency according to the
                        process shown in Figure 3-31, enters the corresponding state according to the
                        query result, and then transitions the state according to Figure 3-31.

                        Figure 3-31 LDP-IGP synchronization query process and state transition




                                 NOTE

                             Note the differences when different IGP protocols are used:
                             ●     When OSPF is used, the state transits based on the flowchart shown in Figure
                                   3-31.
                             ●     When IS-IS is used, the Hold-normal-cost state does not exist. After the Hold-
                                   max-cost timer expires, IS-IS advertises the normal link cost, but the Hold-max-
                                   cost state is displayed even though this state does not exist.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     154
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


Application Scenario
                 LDP-IGP synchronization applies to the following scenarios:
                 On the network shown in Figure 3-32, primary and backup links are connected.

                 Figure 3-32 LDP-IGP synchronization deployment




Benefits
                 This feature reduces the packet loss rate during a primary/backup link switchover
                 and improves the reliability of the entire network.

