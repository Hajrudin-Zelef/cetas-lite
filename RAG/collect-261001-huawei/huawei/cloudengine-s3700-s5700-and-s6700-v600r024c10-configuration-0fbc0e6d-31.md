---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-31
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [3842, 3998]
sha256: 068d2e098cbd1844a6a1204ccd05943c0b606baab5add5489305eae1ae3bb58f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        b.   Enter the MPLS view.
                             mpls

                        c.   Disable the device from establishing proxy egress LSPs.
                             proxy-egress disable

                             If a policy allows LDP to establish LSPs for static and IGP routes or for
                             routes within a specified IP prefix list, the policy also allows LDP to
                             establish proxy egress LSPs. However, these proxy egress LSPs may be
                             useless and unnecessarily consume system resources. To prevent such an
                             issue, run this command to disable the device from establishing proxy
                             egress LSPs.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    65
MPLS Configuration
MPLS Configuration                                                                          3 MPLS LDP Configuration


                        d.   Return to the system view.
                             quit

                        e.   Enter the MPLS-LDP view.
                             mpls ldp

                        f.   Enter the MPLS-LDP-IPv4 view.
                             ipv4-family

                        g.   Configure a policy for triggering LDP LSP establishment.
                             lsp-trigger { all | host | ip-prefix ip-prefix-name | none }

                             ▪      By default, LDP uses IP routes with 32-bit addresses to trigger LDP
                                    LSP establishment.

                             ▪      The default configuration is recommended. Running the lsp-trigger
                                    all command is not recommended, as this command enables LDP
                                    LSPs to be established for all static routes and IGP routes. As a result,
                                    a large number of LSPs are established, consuming excessive label
                                    resources and slowing down LSP convergence on the entire network.
                                    If this command must be run, configure a policy for filtering out
                                    unnecessary routes first, thereby reducing the number of LDP LSPs to
                                    be established and system resource consumption.

                             ▪      If the lsp-trigger all command is run and an IGP route that does not
                                    overlap with that of an MPLS LDP session exists, the system
                                    establishes a proxy egress LSP for the IGP route.

                             ▪      The command allows LDP to establish ingress and egress LSPs for
                                    public network routes and for private network IGP routes. To
                                    configure a policy for triggering transit LSP establishment, run the
                                    propagate mapping command.

                             ▪      If the lsp-trigger command is run in both the MPLS and MPLS-LDP-
                                    IPv4 views, the configuration in the MPLS-LDP-IPv4 view takes effect.
                 ●      Perform the following configurations on a transit node.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   (Optional) Enter the MPLS-LDP-IPv4 view.
                             ipv4-family

                        d.   Configure a policy for triggering LDP LSP establishment.
                             propagate mapping for ip-prefix ip-prefix-name

                             By default, LDP establishes transit LSPs for all routes, without filtering
                             them.

                             The command takes effect in both the MPLS-LDP and MPLS-LDP-IPv4
                             views. If the command is configured in both views, only the latter
                             configuration takes effect.

                 ----End



Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                    66
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


Result
                 ●      Run the display mpls ldp lsp [ destination-address mask-length | all ]
                        command to check LDP LSP information.
                 ●      Run the display mpls lsp [ verbose ] command to check LSP information.
                 ●      Run the display mpls ldp lsp fault-analysis ip-address mask command to
                        check the cause for an LDP LSP establishment failure.

3.9.3 Disabling LDP LSP Flapping Suppression

Context
                 If an LDP LSP changes from up to down on a device due to a protocol or interface
                 fault, the device immediately attempts to reestablish the LDP LSP so as to
                 maximize the hard convergence speed of the LDP LSP. However, if the LDP LSP
                 continuously alternates between up and down when the downstream node
                 frequently sends a label to the upstream node or a label withdraw message, the
                 CPU usage may surge. To address this issue, configure LDP LSP flapping
                 suppression to prevent label flapping.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enter the MPLS-LDP view.
                 mpls ldp

         Step 4 Disable LDP LSP flapping suppression.
                 protocol-packets suppress disable

                 By default, LDP LSP flapping suppression is enabled.

                 ----End

3.9.4 Example for Configuring a Policy for Triggering LDP LSP
Establishment (Ingress and Egress)

Networking Requirements
                 As shown in Figure 3-11, LSRA, LSRB, and LSRC function as core or edge devices
                 on the backbone network. To implement MPLS service deployment, each pair of
                 LSRs are required to advertise labels to each other and establish LDP LSPs after
                 local LDP sessions are set up.

                 Figure 3-11 Configuring a policy for triggering LDP LSP establishment
                         NOTE

                        Interfaces 1 and 2 in this example represent VLANIF100 and VLANIF200, respectively.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    67
MPLS Configuration
MPLS Configuration                                                                                     3 MPLS LDP Configuration




Precautions
                 During the configuration, note the following:
                 ●      Each LSR must have route entries that exactly match FECs for the LSPs to be
                        established.
                 ●      By default, the triggering policy is host, allowing a device to use host IP
                        routes with 32-bit addresses to trigger LDP LSP establishment.
                 ●      If the triggering policy is all, all static routes and IGP routes are used to
                        trigger LDP LSP establishment. The device does not use public network BGP
                        routes to trigger LDP LSP establishment.

Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Configure local LDP sessions.
                 2.     Change the policy for triggering LSP establishment on each LSR.

