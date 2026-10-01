---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-210
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [30517, 30647]
sha256: 8769d96afdb76a3ec4c5554d43ef1f18f2629e0e0d6b7e0d95c09840177464f5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Prerequisites
                 Before configuring synchronization between TE FRR and CR-LSP backup, complete
                 the following tasks:
                 ●      Configure CR-LSP backup: Configure hot standby or ordinary backup.
                 ●      Configure manual MPLS TE FRR or configure MPLS TE Auto FRR.

Context
                 If both TE FRR local protection and end-to-end CR-LSP backup protection are
                 deployed, the system supports synchronization between a bypass tunnel and a
                 backup CR-LSP. After synchronization between TE FRR and CR-LSP backup is
                 enabled:

                 ●      If ordinary backup is enabled, the following situations occur:
                        If the protected link or node fails, TE FRR switches traffic to the bypass CR-LSP
                        and attempts to restore the primary CR-LSP and to set up a backup CR-LSP.
                        If the backup CR-LSP is set up successfully and the primary CR-LSP is not
                        restored, traffic is switched to the backup CR-LSP.
                        After the primary CR-LSP restores successfully, traffic is switched back to the
                        primary CR-LSP, regardless of whether traffic is transmitted along the bypass
                        or backup CR-LSP.
                        If the backup CR-LSP fails to be set up and the primary CR-LSP is not restored,
                        traffic is transmitted along the bypass CR-LSP.
                 ●      If CR-LSP hot standby is enabled, the following situations occur:
                        If the protected link or node fails and the backup CR-LSP is up, traffic is
                        switched to the bypass CR-LSP and then immediately to the backup CR-LSP.
                        At the same time, the system attempts to restore the primary CR-LSP.
                        If the backup CR-LSP is down, the processing is the same as that in ordinary
                        backup.

                 If the primary CR-LSP is up, the system keeps attempting to set up a hot-standby
                 CR-LSP. If the hot-standby CR-LSP is set up successfully, extra bandwidth is
                 occupied. An ordinary backup CR-LSP is set up only when the primary CR-LSP is in
                 the FRR-in-use state. If the primary CR-LSP works properly, no ordinary backup
                 CR-LSP is set up, and no extra bandwidth is used. Synchronization between CR-LSP
                 ordinary backup and TE FRR is recommended. After synchronization between TE
                 FRR and CR-LSP backup is enabled, the entire CR-LSP can be protected.

                 Perform the following configuration on the ingress of a primary MPLS TE tunnel.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            505
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS TE tunnel interface view.
                 interface tunnel tunnel-number

         Step 3 Configure synchronization between TE FRR and CR-LSP backup.
                 mpls te backup frr-in-use

                 If the primary CR-LSP fails, the system starts the bypass CR-LSP (that is, the
                 primary CR-LSP is in the FRR-in-use state) and attempts to restore the primary CR-
                 LSP. At the same time, the system attempts to set up a backup CR-LSP.

                 ----End

Verifying the Configuration
                 ●      Run the display mpls te tunnel-interface [ tunnel interface-number | auto-
                        bypass-tunnel [ autoname ] ] command to check tunnel states.

4.30.2 Example for Configuring Synchronization Between TE
FRR and CR-LSP Backup
Networking Requirements
                 On the network shown in Figure 4-54, establish a primary tunnel along the
                 explicit path LSR1 -> LSR2 -> LSR3 -> LSR4. Establish a TE FRR bypass tunnel on
                 LSR2 along the path LSR2 -> LSR5 -> LSR3. Configure CR-LSP ordinary backup on
                 LSR1 to establish a backup tunnel along the path LSR1 -> LSR6 -> LSR3 -> LSR4. If
                 the link between LSR2 and LSR3 is faulty (the primary CR-LSP is in the FRR-in-use
                 state), the system starts the TE FRR bypass tunnel and attempts to restore the
                 primary CR-LSP. At the same time, the system attempts to establish a backup CR-
                 LSP.

                         NOTE

                        To avoid loops in this scenario, ensure that all connected interfaces have STP disabled and
                        are removed from VLAN1. If STP is enabled and VLANIF interfaces of switches are used to
                        construct a Layer 3 ring network, an interface on the network will be blocked. As a result,
                        Layer 3 services on the network cannot run properly.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       506
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Figure 4-54 Network diagram of synchronization between TE FRR and CR-LSP
                 backup




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign IP addresses to interfaces and configure OSPF to ensure that public
                        network routes between the nodes are reachable.
                 2.     Configure LSR IDs for the nodes, and enable MPLS, MPLS TE, RSVP-TE, CSPF,
                        and OSPF TE on the nodes globally and on the interfaces involved.
                 3.     On the ingress of the primary tunnel, create a tunnel interface and specify the
                        tunnel IP address, tunneling protocol, destination address, tunnel ID, and
                        dynamic signaling protocol (RSVP-TE).
                 4.     Enable TE FRR on the primary tunnel interface of the ingress.
                 5.     Configure a TE FRR bypass tunnel along the path LSR2 -> LSR5 -> LSR3 on
                        LSR2 to protect the link between LSR2 and LSR3.
                 6.     Configure an ordinary backup CR-LSP on the ingress along the path LSR1 ->
                        LSR6 -> LSR3 -> LSR4.
                 7.     Configure synchronization between the bypass tunnel and the backup CR-LSP
                        in the primary tunnel interface view on the ingress.

Procedure
         Step 1 Configure interface IP addresses for the nodes.
                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100 600
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 24


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            507
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration

