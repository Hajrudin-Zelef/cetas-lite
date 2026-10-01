---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-58
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [8052, 8192]
sha256: 24b69afdeb60be10398f79d42d1f1044c66eeb7189569f901e3eb72d4f7bc1da
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        134
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration


                 timer expires before a new Hello message is received, the LSR considers that the
                 Hello adjacency is interrupted. The Hello mechanism has a drawback — it cannot
                 rapidly detect link faults, especially when a Layer 2 device is deployed between an
                 LSR and its peer.

                 Figure 3-26 Primary and backup LDP LSPs




                 Against the preceding backdrop, the BFD mechanism, which boasts rapid and
                 lightweight, is used in LDP LSP scenarios for quickly detecting LDP LSP faults and
                 triggering a path switchover in case of faults, so as to minimize data loss and
                 improve service reliability.

BFD for LDP LSP
                 BFD for LDP LSP is implemented by establishing a BFD session between the two
                 ends of an LSP and binding the session to the LSP. After BFD for LDP LSP is
                 deployed, BFD can rapidly detect LSP faults and trigger traffic switchover in case
                 of faults. When an LSP is unidirectional, BFD allows its reverse path to be an IP
                 link, an LSP, or a TE tunnel.

                 A BFD for LDP LSP session can be negotiated in either of the following modes:
                 ●      Static configuration: The negotiation of a BFD session is performed using the
                        local and remote discriminators that are manually configured for the BFD
                        session. In this mode, you need to specify the next hop IP address for the LSP
                        and peer IP address for the BFD session, and bind the BFD session to the LSP.
                 ●      Dynamic mode: The negotiation of a BFD session is performed based on the
                        BFD discriminator TLV carried in an LSP ping packet. In this mode, you need
                        to configure a policy for triggering BFD session establishment. The policy can
                        then be used by BFD to establish sessions and bind them to LSPs.
                        –   Host address-based triggering policy: The establishment of BFD sessions
                            is triggered by host addresses. The LSPs for which BFD sessions need to
                            be established can be constrained based on the next hop and outbound
                            interface information.
                        –   FEC list-based triggering policy: The establishment of BFD sessions is
                            triggered by host addresses listed in a configured FEC list.

                 After a BFD session is established, the ingress and egress periodically send BFD
                 packets to check LSP continuity. If either end does not receive BFD packets from
                 the peer end within the detection period, BFD considers the LSP faulty and reports
                 an LSP fault event to the LDP management module.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             135
MPLS Configuration
MPLS Configuration                                                                    3 MPLS LDP Configuration


                         NOTE

                        On the proxy egress, BFD does not establish a return session for the proxy egress LSP even
                        if BFD for LDP is enabled.


BFD for LDP Tunnel
                 BFD for LDP LSP only detects primary LSP faults and switches traffic to an FRR
                 bypass LSP or load-balancing LSPs. If both a primary LSP and its FRR bypass LSP
                 or load-balancing LSPs fail simultaneously, the BFD mechanism does not take
                 effect. LDP can instruct its upper-layer application to perform a protection
                 switchover (such as VPN FRR or VPN equal-cost load balancing) only after LDP
                 itself detects the FRR bypass LSP failure or load-balancing LSP failure. To address
                 this issue, BFD for LDP tunnel can be used. Here, LDP tunnel is a general term for
                 both primary and FRR bypass LSPs. The BFD for LDP tunnel mechanism
                 establishes a BFD session that can simultaneously monitor the primary LSP and its
                 FRR bypass LSP or load-balancing LSPs. If both the primary LSP and its FRR bypass
                 LSP or load-balancing LSPs fail, BFD can rapidly detect the failure and trigger an
                 LDP upper-layer application to perform a protection switchover, reducing traffic
                 loss.
                 BFD for LDP tunnel uses the same mechanism as BFD for LDP LSP to monitor the
                 connectivity of each LSP in an LDP tunnel. Unlike BFD for LDP LSP, BFD for LDP
                 tunnel has the following characteristics:
                 ●      It supports dynamic BFD sessions but does not support static ones.
                 ●      It can trigger a BFD for LDP tunnel session using a host IP address, a FEC list,
                        or an IP prefix list.
                 ●      It does not allow you to specify next hop addresses or outbound interface
                        names in BFD session triggering policies.

Application Scenarios
                 BFD for LDP applies to the following scenarios:
                 ●      BFD for LDP LSP applies to LDP FRR scenarios.
                 ●      BFD for LDP tunnel applies to VPN FRR scenarios.

Benefits
                 BFD for LDP provides a rapid, lightweight fault detection mechanism for LDP LSPs,
                 improving network reliability.

3.16.2 Enabling BFD Globally
                 BFD can rapidly detect faults on LDP LSPs and trigger a primary/backup LSP
                 switchover, improving network reliability.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enable BFD globally and enter the global BFD view.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     136
MPLS Configuration
MPLS Configuration                                                                      3 MPLS LDP Configuration

                 bfd

                 ----End

3.16.3 Enabling the Function to Dynamically Create BFD
Sessions in MPLS Scenarios
                 Enable BFD on the ingress and egress in an MPLS domain, after which BFD
                 sessions can be dynamically created.

Procedure
                 ●      Perform the following steps on the ingress.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the MPLS view.
                             mpls
                        c.   Enable the function to dynamically create BFD sessions for LDP LSPs.
                             mpls bfd enable

                                     NOTE

                                    After this command is run, the ingress does not create a BFD session
                                    immediately. Instead, the ingress waits for the egress to be configured and sends
                                    an LSP ping request carrying the BFD TLV before creating a BFD session.
                 ●      Perform the following steps on the egress.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the BFD view.
                             bfd
                        c.   Enable the function to passively create BFD sessions.
                             mpls-passive

                                     NOTE

