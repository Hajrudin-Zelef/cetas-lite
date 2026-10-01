---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-161
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [23288, 23448]
sha256: bae04e7e7998fd01debea870a97cace73840e9691db8ac15019c55b2a93101b6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                  Configure an explicit           mpls te path explicit-   -
                  path for the backup CR-         path path-name
                  LSP.                            secondary

                  Configure the affinity          mpls te affinity         By default, the affinity
                  attribute for the backup        property properties      attribute of a backup CR-
                  CR-LSP.                         [ mask mask-value ]      LSP is 0x0.
                                                  secondary

                  Set a hop limit for the         mpls te hop-limit hop-   By default, the hop limit
                  backup CR-LSP.                  limit-value secondary    of a backup CR-LSP is 32.



                 ----End

Verifying the Configuration
                 ●      Run the display mpls te tunnel-interface command to check information
                        about tunnel interfaces and the status of backup CR-LSPs.
                 ●      Run the display mpls te tunnel path command to check primary tunnel path
                        information on the local node.


4.23.4 (Optional) Configuring a Best-Effort Path for a CR-LSP

Prerequisites
                 Before configuring a best-effort path for a CR-LSP, complete the following task:

                 ●      Configure a dynamic MPLS TE tunnel.
                 ●      Enable MPLS, MPLS TE, and RSVP-TE globally and on interfaces on each node
                        of the backup CR-LSP. For details, see 4.7.1 Enabling MPLS TE and RSVP-TE.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            388
MPLS Configuration
MPLS Configuration                                                                    4 MPLS TE Configuration


Context
                 Configure a best-effort path on the ingress of a primary CR-LSP to take over traffic
                 if both the primary and backup CR-LSPs fail.

                        NOTE

                 A best-effort path does not provide bandwidth guarantee for traffic. Configure the affinity
                 attribute and hop limit as needed.
                 CR-LSP hot standby can work with a best-effort path to further enhance reliability. CR-LSP
                 ordinary backup cannot work with a best-effort path.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS TE tunnel interface view.
                 interface tunnel tunnel-number

         Step 3 Configure a best-effort path for a CR-LSP.
                 mpls te backup ordinary best-effort

         Step 4 (Optional) Set parameters for the best-effort CR-LSP path as required. For details,
                see Table 4-19.

                 To allow traffic to pass through a specified best-effort CR-LSP path, perform one
                 or more steps in Table 4-17.


                 Table 4-19 Parameters for configuring a best-effort CR-LSP path

                  Operation                            Command                     Purpose

                  Configure the affinity               mpls te affinity            The default affinity
                  attribute for a best-effort          property properties         attribute of a best-effort
                  path.                                [ mask mask-value ]         path is 0x0.
                                                       best-effort

                  Set a hop limit for the              mpls te hop-limit hop-      The default hop limit of
                  best-effort path.                    limit-value best-effort     a best-effort path is 32.



                 ----End


Verifying the Configuration
                 ●      Run the display mpls te tunnel-interface command to check tunnel
                        interface information and best-effort path status.
                 ●      Run the display mpls te tunnel path command to check primary tunnel path
                        information on the local node.



Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                      389
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


4.23.5 Example for Configuring CR-LSP Hot Standby
Networking Requirements
                 On the MPLS network shown in Figure 4-35, establish a TE tunnel from LSR1 to
                 LSR3 and configure CR-LSP hot standby and best-effort path creation. Here:
                 ●      The path of the primary CR-LSP is LSR1 -> LSR2 -> LSR3.
                 ●      The path of the backup CR-LSP is LSR1 -> LSR4 -> LSR3.
                 The requirements are as follows: If the primary CR-LSP fails, traffic is switched to
                 the backup CR-LSP. After the primary CR-LSP recovers, traffic is switched back to
                 the primary CR-LSP after a 15-second delay. If both the primary and backup CR-
                 LSPs fail, traffic is switched to the best-effort path. Explicit paths can be
                 configured for the primary and backup CR-LSPs. A best-effort path can be
                 generated automatically. In this example, the best-effort path is LSR1 -> LSR4 ->
                 LSR2 -> LSR3. The calculated best-effort path varies according to the faulty node.

                 Figure 4-35 Network diagram of CR-LSP hot standby




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign IP addresses to interfaces and configure OSPF to ensure that public
                        network routes between the nodes are reachable.
                 2.     Configure LSR IDs for the nodes, and enable MPLS, MPLS TE, RSVP-TE, CSPF,
                        and OSPF TE on the nodes globally and on the interfaces involved.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             390
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration


                 3.     Configure primary and backup explicit paths on the ingress LSR1.
                 4.     Configure a dynamic primary MPLS TE tunnel, bind the tunnel to an explicit
                        path, enable hot standby and best-effort path, and set the switchback delay
                        to 15s.

Procedure
         Step 1 Configure interface IP addresses for the devices.
                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100 500
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 24
                 [LSR1-Vlanif100] quit
                 [LSR1] interface vlanif 500
                 [LSR1-Vlanif600] ip address 10.1.5.1 24
                 [LSR1-Vlanif600] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface 10ge 1/0/2
                 [LSR1-10GE1/0/2] port link-type trunk
                 [LSR1-10GE1/0/2] port trunk allow-pass vlan 500
                 [LSR1-10GE1/0/2] quit
                 [LSR1] interface loopback 1
                 [LSR1-loopback1] ip address 1.1.1.9 32
                 [LSR1-loopback1] quit

