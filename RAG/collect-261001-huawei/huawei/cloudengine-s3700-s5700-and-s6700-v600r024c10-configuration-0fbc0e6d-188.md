---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-188
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [27310, 27429]
sha256: 9fbc2a5f29ee19898f8834d0ec5f07e0d76fb5d9144b7e60923180e08a9f535d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

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

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     456
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration


                        If the primary CR-LSP is up, the system keeps attempting to set up a hot-
                        standby CR-LSP. If the hot-standby CR-LSP is set up successfully, extra
                        bandwidth is occupied. An ordinary backup CR-LSP is set up only when the
                        primary CR-LSP is in the FRR-in-use state. If the primary CR-LSP works
                        properly, no extra bandwidth is used. Synchronization between CR-LSP
                        ordinary backup and TE FRR is recommended.

4.27.2 Enabling MPLS TE Auto FRR
Prerequisites
                 Before enabling MPLS TE Auto FRR, complete the following task:
                 ●      Configure a dynamic MPLS TE tunnel.
                 ●      Enable MPLS, MPLS TE, and RSVP-TE globally and on interfaces on each node
                        of a bypass tunnel. For details, see 4.7.1 Enabling MPLS TE and RSVP-TE.
                 ●      Enable CSPF on a PLR.

Context
                 On a network that requires high service reliability, FRR is usually configured to
                 improve network reliability. If the network topology is complex and a great
                 number of links must be configured, the configuration procedure is complex. Auto
                 FRR automatically establishes an eligible bypass tunnel, reducing the configuration
                 workload.
                 Upgrade binding is supported, allowing a device to automatically bind to a higher-
                 priority bypass tunnel when it becomes available. A bypass tunnel is selected
                 based on the following rules prioritized in descending order:
                 ●      SRLG
                        If MPLS TE Auto FRR and an SRLG attribute are configured, the primary and
                        bypass tunnels must be in different SRLGs. Otherwise, the bypass tunnel may
                        fail to be established.
                 ●      Bandwidth protection takes precedence over non-bandwidth protection.
                 ●      Node protection takes precedence over link protection.
                 ●      Manual protection takes precedence over automatic protection.
                 When configuring MPLS TE Auto FRR, you need to enable MPLS TE Auto FRR
                 globally on the PLR of a bypass tunnel. If link protection is required, enable link
                 protection on the corresponding interfaces.

                         NOTE

                        ● RSVP-TE tunnels using bandwidth reserved in SE style support FRR.
                        ● To implement millisecond-level fast switchover, you also need to configure dynamic BFD
                          for RSVP.
                        ● Only a primary CR-LSP supports MPLS TE Auto FRR.


Procedure
         Step 1 Enter the system view.
                 system-view


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  457
MPLS Configuration
MPLS Configuration                                                                    4 MPLS TE Configuration


         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enable MPLS TE Auto FRR globally.
                 mpls te auto-frr [ self-adapting ]

                 After TE Auto FRR is enabled globally, node protection is enabled on all MPLS TE-
                 enabled interfaces by default.
                 To enable an automatic bypass tunnel to dynamically select node protection or
                 link protection based on network conditions, specify self-adapting. If the self-
                 adapting parameter is not specified, the current automatic bypass tunnel selects
                 node protection by default.
         Step 4 (Optional) Set the interval at which the TE FRR binding relationship is refreshed.
                 mpls te timer fast-reroute [ weight ]

                 After TE FRR is configured, the PLR periodically refreshes the TE FRR binding
                 relationship. By default, the system searches for an optimal bypass tunnel for each
                 primary tunnel every 1 second and binds the bypass tunnel to the primary tunnel.
         Step 5 (Optional) Configure MPLS TE Auto FRR on an interface. You only need to run the
                command in the view of the outbound interface of the primary tunnel.
                 quit
                 interface interface-type interface-number
                 mpls te auto-frr { block | default | link | node | self-adapting }

                        NOTE

