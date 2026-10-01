---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-160
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [23145, 23287]
sha256: c9c477463114a7d5753367a88ebc55b756c71dc8feb42fd2b3ed1013f1202229
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 4 (Optional) Set CR-LSP hot standby parameters as required. For details, see Table
                4-17.

                 After the hot standby mode is configured, the system automatically selects a
                 backup CR-LSP path. To allow traffic to pass through a specified backup CR-LSP,
                 perform one or more steps in Table 4-17.


                 Table 4-17 Configure hot-standby CR-LSP parameters.

                  Operation                   Command                     Purpose

                  Configure path              mpls te backup hot-         By default, a hot-standby
                  overlapping to allow a      standby overlap-path        CR-LSP must be disjoined
                  hot-standby CR-LSP path                                 from a primary CR-LSP. If
                  to overlap with a                                       no CR-LSP meets this
                  primary CR-LSP path.                                    requirement, a hot-
                                                                          standby CR-LSP fails to
                                                                          be established. The path
                                                                          overlapping function
                                                                          allows a hot-standby CR-
                                                                          LSP to share some links
                                                                          with a primary CR-LSP,
                                                                          but it cannot fully
                                                                          overlap the primary CR-
                                                                          LSP.

                  Allow the path of a hot-    mpls te backup hot-         By default, the device
                  standby CR-LSP to           standby overlap-path        disallows the path of a
                  partially or completely     frr-in-use                  hot-standby CR-LSP and
                  overlap with the path of                                the path of a primary
                  a primary CR-LSP in the                                 CR-LSP in the FRR-in-use
                  FRR-in-use state.                                       state to overlap,
                                                                          meaning that their paths
                                                                          must be separated.
                                                                          When network resources
                                                                          are insufficient, you can
                                                                          enable path overlapping
                                                                          to increase the
                                                                          probability of successful
                                                                          hot-standby CR-LSP
                                                                          setup.

                  Configure an explicit       mpls te path explicit-      -
                  path for the backup CR-     path path-name
                  LSP.                        secondary

                  Configure the affinity      mpls te affinity            By default, the affinity
                  attribute for the backup    property properties         attribute of a backup CR-
                  CR-LSP.                     [ mask mask-value ]         LSP is 0x0.
                                              secondary

                  Set a hop limit for the     mpls te hop-limit hop-      By default, the hop limit
                  backup CR-LSP.              limit-value secondary       of a backup CR-LSP is 32.



Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                          386
MPLS Configuration
MPLS Configuration                                                                    4 MPLS TE Configuration


         Step 5 (Optional) Configure forcible traffic switching.
                 hotstandby-switch { force | clear }

                 If the primary CR-LSP goes down, traffic is switched to the hot-standby CR-LSP. If
                 the primary LSP goes up, traffic is switched back to the primary LSP by default.
                 This configuration provides the flexibility to control the traffic switching behavior.

                 If force is specified, traffic is temporarily switched to the hot-standby CR-LSP. If
                 clear is specified, traffic is switched back to the primary CR-LSP.

         Step 6 (Optional) Configure CSPF fast switching.
                 quit
                 mpls
                 mpls te tedb fast-notice

                 If BFD is not configured and the primary LSP fails, TE hot-standby switching relies
                 on protocol convergence. To speed up traffic switching, you can configure CSPF
                 fast switching to ensure that traffic can be rapidly switched to the hot-standby
                 LSP.

                 ----End


Verifying the Configuration
                 ●      Run the display mpls te tunnel-interface command to check information
                        about tunnel interfaces, backup CR-LSP status, switchback policy of hot-
                        standby CR-LSPs, and enabling status of path overlapping.
                 ●      Run the display mpls te hot-standby state all verbose command to check
                        the hot-standby status of a tunnel.
                 ●      Run the display mpls te tunnel path command to check primary tunnel path
                        information on the local node.

4.23.3 Configuring CR-LSP Ordinary Backup

Prerequisites
                 Before configuring CR-LSP ordinary backup, complete the following task:

                 ●      Configure a dynamic MPLS TE tunnel.
                 ●      Enable MPLS, MPLS TE, and RSVP-TE globally and on interfaces on each node
                        of the backup CR-LSP. For details, see 4.7.1 Enabling MPLS TE and RSVP-TE.


Context
                 Configure an ordinary backup CR-LSP to take over traffic if the primary CR-LSP
                 fails, preventing service interruptions.

                         NOTE

                        Ordinary CR-LSP backup cannot be configured together with the best-effort path or hot
                        standby mode.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      387
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS TE tunnel interface view.
                 interface tunnel tunnel-number

         Step 3 Configure the ordinary CR-LSP establishment mode.
                 mpls te backup ordinary

         Step 4 (Optional) Set CR-LSP ordinary backup parameters as required. For details, see
                Table 4-18.

                 After the ordinary backup mode is configured, the system automatically selects a
                 backup CR-LSP path. To allow traffic to pass through a specified backup CR-LSP,
                 perform one or more steps in Table 4-18.

                 Table 4-18 Configure CR-LSP ordinary backup parameters.

                  Operation                       Command                  Purpose

