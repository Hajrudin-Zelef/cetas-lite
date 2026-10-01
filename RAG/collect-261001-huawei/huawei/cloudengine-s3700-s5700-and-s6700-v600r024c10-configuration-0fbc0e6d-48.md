---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-48
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [6509, 6656]
sha256: d9408968ffa68e350ff50e198890c73ee80b3d0953a01545cac6f88d8d526693
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

MPLS Configuration
MPLS Configuration                                                         3 MPLS LDP Configuration


                 Figure 3-21 Typical LDP Auto FRR application scenario - ring topology (1)




                 To address this issue, LDP Remote LFA FRR can be used. Remote LFA FRR is
                 implemented based on LDP Auto FRR of IGP remote LFA FRR. Figure 3-22
                 illustrates the typical LDP Auto FRR application scenario. The primary LDP LSP is
                 established over the path PE1 -> PE2. Remote LFA FRR establishes a remote LFA
                 FRR LSP over the path PE1 -> P2 -> PE2 to protect the primary LDP LSP. The
                 detailed implementation process is as follows:

                 Figure 3-22 Typical LDP Auto FRR application scenario - ring topology (2)




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                           110
MPLS Configuration
MPLS Configuration                                                                       3 MPLS LDP Configuration


                 1.     An IGP uses the remote LFA algorithm to calculate a remote LFA route with
                        the PQ node (P2) IP address and the recursed outbound interface's next hop
                        and then notifies the route management module of the information.
                 2.     LDP obtains the remote LFA route from the route management module. PE1
                        automatically establishes a remote LDP peer relationship with the PQ node
                        and a remote LDP session for the relationship. PE1 then establishes an LDP
                        LSP to the PQ node and a remote LFA FRR LSP over the path PE1 -> P2 ->
                        PE2.
                 3.     LDP-enabled PE1 establishes an LDP LSP over the path PE1 -> P1 -> P2 based
                        on the recursed outbound interface's next hop. This LSP is called a remote LFA
                        FRR recursed LSP.
                 If PE1 detects a fault, PE1 rapidly switches traffic to the remote LFA FRR LSP.

3.14.2 Configuring LDP Auto FRR
Prerequisites
                 Before configuring LDP Auto FRR, you have completed the following task:
                 ●      3.6.2 Configuring MPLS LDP Globally

Context
                 LDP Auto FRR depends on IGP Auto FRR or static route FRR. After IGP Auto FRR or
                 static route FRR is enabled, LDP Auto FRR is automatically enabled. If LDP Auto
                 FRR is needed, configure it on the ingress or transit node.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 (Optional) Enter the MPLS-LDP-IPv4 view.
                 ipv4-family

         Step 4 (Optional) Configure a policy for triggering LDP to establish backup LSPs.
                 auto-frr lsp-trigger { all | host | ip-prefix ip-prefix-name | none }

                 By default, LDP uses backup routes with 32-bit addresses to trigger backup LSP
                 establishment.

                         NOTE

                        If both the auto-frr lsp-trigger and the lsp-trigger commands are run, the established
                        backup LSPs satisfy both the policy for triggering LDP LSP establishment and the policy for
                        triggering backup LDP LSP establishment.
                        The auto-frr lsp-trigger command takes effect in both the MPLS-LDP and MPLS-LDP-IPv4
                        views. If the command is configured in both views, only the latter configuration takes
                        effect.

                 ----End

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      111
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


Result
                 ●      Run the display mpls lsp command to check LSP information after LDP Auto
                        FRR is enabled.
                 ●      Run the display mpls ldp event session-down verbose command to check
                        the reason why a session goes down. In the command output, IGP delete the
                        RLFA IID indicates that the session goes down because the RLFA route
                        changes.
                 ●      Run the display mpls ldp event adjacency-down verbose command to
                        check the reason why an adjacency goes down. In the command output, IGP
                        delete the RLFA IID indicates that the adjacency goes down because the
                        RLFA route changes.

3.14.3 (Optional) Configuring LDP Graceful Deletion

Context
                 LDP graceful deletion can be configured in the LDP-IGP synchronization or LDP
                 FRR scenario to speed up traffic switching. It helps implement uninterrupted traffic
                 transmission during traffic switching, which improves reliability of the entire
                 network.

                 If both the primary link and the LDP session on that link also go down, LDP
                 immediately instructs the upstream device to withdraw labels and triggers LDP
                 Auto FRR. LSP convergence on the backup link requires LDP to distribute labels to
                 the upstream device again, which prolongs convergence and FRR traffic switching.
                 As a result, packet loss occurs. If LDP graceful deletion is configured and the LDP
                 session goes down, LDP delays deleting the LDP session and keeps the relevant
                 labels and LSP. The LSP on the backup link does not require LDP to distribute
                 labels to the upstream device again, which shortens FRR traffic switching and
                 reduces packet loss.

                 Perform the following configuration on the LDP FRR-enabled LSR.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Configure LDP graceful deletion.
                 graceful-delete

         Step 4 Configure a timer for LDP graceful deletion.
                 graceful-delete timer timer

                 After the LDP session goes down, forwarding entries on the LSR remain before the
                 graceful deletion timer expires.

                 By default, the value of the graceful deletion timer is 5 seconds.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                       112
MPLS Configuration
MPLS Configuration                                                                      3 MPLS LDP Configuration


                         NOTE

                 If the value of the graceful delete timer is too large, the invalid LSP will be kept for a long time,
                 consuming system resources.

                 ----End

