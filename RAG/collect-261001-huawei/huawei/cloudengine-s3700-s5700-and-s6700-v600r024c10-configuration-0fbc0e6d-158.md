---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-158
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [22874, 23008]
sha256: 9b39acfe3665b5c8f0cf354510a24d4b28b4f831b00c37ebe675493ff8d2eb1b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 1.     When LSR3 enters the IS-IS overload state, IS-IS encapsulates the overload
                        information into IS-IS protocol packets and floods the packets throughout the
                        entire IS-IS domain.
                 2.     LSR1 detects that LSR3 is overloaded and recalculates paths to LSR2 for all
                        LSPs passing through LSR3.
                 3.     LSR1 calculates a new path LSR1 -> LSR4 - >LSR2, which bypasses the
                        overloaded IS-IS node. LSR1 triggers the creation of a new LSP over the new
                        path LSR1 -> LSR4 -> LSR2.
                 4.     After the new LSP is established, LSR1 switches traffic from the original LSP to
                        the new LSP, ensuring the quality of the service carried on Tunnel1.

4.22.2 Configuring Synchronization Between CR-LSP
Establishment and Overload Status

Prerequisites
                 Before configuring synchronization between CR-LSP establishment and overload,
                 complete the following task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 A node becomes overloaded in the following situations:
                 ●      When the node is transmitting a large number of services and its system
                        resources are exhausted, the node marks itself overloaded.
                 ●      An administrator runs the set-overload command to mark the node
                        overloaded if the node is transmitting a large number of services and its CPU
                        is busy.

                 If there are overloaded nodes on an MPLS TE network, associate CR-LSP
                 establishment with overload to ensure that CR-LSPs are established over paths
                 excluding overloaded nodes. This configuration prevents overloaded nodes from
                 being further burdened and improves CR-LSP reliability.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure synchronization between CR-LSP establishment and IS-IS overload
                status.
                 mpls te path-selection overload

                 The association allows CSPF to exclude overloaded IS-IS nodes when calculating
                 paths.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             381
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                         NOTE

                        Before the association is configured, run the mpls te cspf command to enable CSPF, and
                        the mpls te record-route command to enable route and label recording.
                        Traffic travels through an existing CR-LSP before a new CR-LSP is established. After the new
                        CR-LSP is established, traffic is switched to the new CR-LSP, and the original CR-LSP is
                        deleted. This traffic switchover operates on the make-before-break mechanism, ensuring
                        that no traffic is lost during the switchover.

                 After the mpls te path-selection overload command is run:
                 ●      Existing CR-LSPs are re-optimized. CSPF recalculates paths to prevent traffic
                        from passing through overloaded nodes.
                 ●      Traffic is not dropped during the switchover. CSPF calculates paths excluding
                        overloaded nodes for new CR-LSPs.
                         NOTE

                        This function does not take effect on bypass tunnels.
                        If the ingress or egress is marked overloaded, the mpls te path-selection overload
                        command does not take effect. The established CR-LSPs associated with the ingress or
                        egress will not be re-optimized, and new CR-LSPs associated with the ingress or egress will
                        not be established.

                 ----End


4.23 Configuring CR-LSP Backup

4.23.1 Understanding CR-LSP Backup
                 A CR-LSP that is used to protect a primary CR-LSP on the same MPLS TE tunnel is
                 called a backup CR-LSP. A backup CR-LSP protects traffic of an important CR-LSP.
                 If the primary CR-LSP fails, the ingress can detect that the primary CR-LSP is
                 unavailable and switch traffic to the backup CR-LSP. After the primary CR-LSP
                 recovers, traffic is switched back to the primary CR-LSP.
                 CR-LSP backup has two modes: hot standby and ordinary backup. To further
                 improve the reliability of MPLS TE tunnels, the best-effort path technology is
                 provided.
                 ●      Hot standby: A backup CR-LSP is set up immediately after the primary CR-LSP
                        is set up. If the primary CR-LSP fails, traffic is switched to the backup CR-LSP.
                        If the primary CR-LSP recovers, traffic is switched back to the primary CR-LSP
                        by default. Hot standby CR-LSPs support best-effort paths.
                 ●      Ordinary backup: A backup CR-LSP is set up after a primary CR-LSP fails. If
                        the primary CR-LSP fails, traffic is switched to the backup CR-LSP. If the
                        primary CR-LSP recovers, traffic is switched back to the primary CR-LSP by
                        default.
                        For details about the differences between hot standby and ordinary backup,
                        see Table 4-16.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      382
MPLS Configuration
MPLS Configuration                                                                    4 MPLS TE Configuration


                        Table 4-16 Differences between hot standby and ordinary backup
                         Item             Hot Standby                       Ordinary Backup

                         Time when a      Created immediately after         Created only after the primary
                         backup CR-       the primary CR-LSP is             CR-LSP fails.
                         LSP is           established.
                         established

                         Primary and      You can specify whether           The path of the backup CR-LSP
                         backup           the primary and backup            can partially overlap the path
                         explicit paths   paths can overlap. If an          of the primary CR-LSP,
                                          explicit path is allowed for      regardless of whether the
                                          a backup CR-LSP, the              backup CR-LSP is set up over
                                          explicit path is used as the      an explicit path.
                                          constraint to set up the
                                          backup CR-LSP.

                         Whether a        Yes                               No
                         best-effort
                         path is
                         supported


