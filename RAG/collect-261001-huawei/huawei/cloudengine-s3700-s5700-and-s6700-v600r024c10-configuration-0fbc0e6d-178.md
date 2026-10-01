---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-178
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [25866, 25983]
sha256: 7bb3b60a718e3eeacb36195651ea1153034e423ad178e042b9bcbc366be1c783
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        FRR does not support multi-point failures. If FRR switching occurs, data is switched from
                        the primary CR-LSP to the bypass CR-LSP. During data forwarding through the bypass CR-
                        LSP, the bypass CR-LSP must remain up. If the bypass CR-LSP fails during this period, the
                        protected data cannot be forwarded through MPLS. As a result, traffic is interrupted and
                        FRR fails. Even if the bypass CR-LSP is reestablished, it cannot forward data. Data
                        forwarding will be restored only after the primary CR-LSP recovers or is reestablished.


Other Functions
                 When TE FRR is in the FRR-in-use state, the RSVP messages sent by the transmit
                 interface do not carry the interface authentication TLV, and the receive interface
                 does not perform interface authentication on the RSVP messages that do not carry
                 the authentication TLV and are in the FRR-in-use state. In this case, you can
                 configure neighbor authentication.

Coexistence of CR-LSP Backup and TE FRR
                 1.     Co-deployment of CR-LSP backup and TE FRR
                        –    Combination of CR-LSP ordinary backup and TE FRR: TE FRR can respond
                             to link faults in a timely manner and switch traffic to a bypass CR-LSP as
                             soon as possible. If both the primary and bypass CR-LSPs fail, a backup
                             CR-LSP is set up to take over traffic.
                        –    Combination of CR-LSP hot standby and TE FRR: TE FRR can respond to
                             link faults in a timely manner and switch traffic to a bypass CR-LSP as
                             soon as possible. Link fault information is sent to the tunnel ingress
                             through RSVP-TE signaling, and then traffic is switched to the backup CR-
                             LSP.
                 2.     Synchronization between CR-LSP backup and TE FRR
                        If both TE FRR local protection and end-to-end CR-LSP backup protection are
                        deployed, the system supports synchronization between a bypass tunnel and
                        a backup CR-LSP. After synchronization between CR-LSP backup and TE FRR is
                        enabled:
                        –    If ordinary backup is enabled, the following situations occur:
                             If the protected link or node fails, TE FRR switches traffic to the bypass
                             CR-LSP and attempts to restore the primary CR-LSP and to set up a
                             backup CR-LSP.
                             If the backup CR-LSP is set up successfully and the primary CR-LSP is not
                             restored, traffic is switched to the backup CR-LSP.
                             After the primary CR-LSP restores successfully, traffic is switched back to
                             the primary CR-LSP, regardless of whether traffic is transmitted along the
                             bypass or backup CR-LSP.
                             If the backup CR-LSP fails to be set up and the primary CR-LSP is not
                             restored, traffic is transmitted along the bypass CR-LSP.
                        –    If hot standby is enabled, the following situations occur:
                             If the protected link or node fails and the backup CR-LSP is up, traffic is
                             switched to the bypass CR-LSP and then immediately to the backup CR-
                             LSP. At the same time, the system attempts to restore the primary CR-LSP.
                             If the backup CR-LSP is down, the processing is the same as that in
                             ordinary backup.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     432
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration


                        If the primary CR-LSP is up, the system keeps attempting to set up a hot-
                        standby CR-LSP. If the hot-standby CR-LSP is set up successfully, extra
                        bandwidth is occupied. An ordinary backup CR-LSP is set up only when the
                        primary CR-LSP is in the FRR-in-use state. If the primary CR-LSP works
                        properly, no extra bandwidth is used. Synchronization between CR-LSP
                        ordinary backup and TE FRR is recommended.

4.26.2 Enabling MPLS TE FRR

Prerequisites
                 Before enabling MPLS TE FRR, you have completed the following tasks:

                 ●      Configure a dynamic MPLS TE tunnel.
                 ●      Enable MPLS, MPLS TE, and RSVP-TE globally and on interfaces on each node
                        of a bypass tunnel. For details, see 4.7.1 Enabling MPLS TE and RSVP-TE.
                 ●      Enable CSPF on the PLR of the bypass tunnel.


Context
                 A bypass tunnel can be created only after TE FRR is enabled for the primary
                 tunnel.

                         NOTE

                        ● The bypass tunnel used by FRR needs to be established in advance, which occupies extra
                          bandwidth. Therefore, if the remaining network bandwidth is insufficient, TE FRR
                          protection should be configured only for key interfaces or links.
                        ● RSVP-TE tunnels using bandwidth reserved in SE style support FRR.
                        ● To implement millisecond-level fast switchover, you also need to configure dynamic BFD
                          for RSVP.

                 Perform the following configuration on the ingress of a primary MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the tunnel interface view of a primary MPLS TE tunnel.
                 interface tunnel tunnel-number

         Step 3 Enable MPLS TE FRR.
                 mpls te fast-reroute [ bandwidth ]

         Step 4 (Optional) Choose one of the following methods to set the Node protection flag
                in messages sent by the tunnel.
                 ●      Set the Node protection flag in messages sent by the current tunnel.
                        mpls te frr { node-protection | no-node-protection }

                 ●      Set the Node protection flag in messages sent by all P2P RSVP-TE tunnels.
                        quit
                        mpls
                        mpls te frr { node-protection | no-node-protection }


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  433
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                         NOTE

