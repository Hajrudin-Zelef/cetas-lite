---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-63
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [8770, 8914]
sha256: e4fc5dfd843107c280878c9a45ba77d6e95042327cb5aa2052fd402a5196290c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Context
                 If a direct link for a local LDP session fails, the LDP adjacency is torn down, and
                 the session and labels are deleted. After the direct link recovers, the LDP session is
                 reestablished and labels are re-distributed, after which the LDP LSP converges
                 again. It takes time to establish an LSP. Before the LSP is established successfully,
                 traffic is lost. To accelerate LDP LSP convergence and reduce traffic loss, the device
                 supports LDP session protection.
                 LDP session protection helps maintain an LDP session, eliminating the need to
                 reestablish an LDP session or re-distribute labels in the preceding situation.

Fundamentals
                 On the network shown in Figure 3-28, LDP session protection is configured on
                 both nodes of a link. The two nodes exchange Link Hello messages to establish a
                 local LDP session and exchange Targeted Hello messages to establish a remote

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        146
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration


                 LDP session. The remote adjacency established using Targeted Hello messages
                 serves as a backup for the local adjacency established using Link Hello messages,
                 protecting the local LDP session.

                 Figure 3-28 Basic application scenario of LDP session protection




                 On the network shown in Figure 3-28, if the direct link between nodes A and B
                 fails, the adjacency established using Link Hello messages is deleted. Because the
                 indirectly connected link is working properly, the remote adjacency established
                 using Targeted Hello messages is not deleted. In this case, the LDP session is
                 maintained by the remote adjacency, and the FEC-label mapping information of
                 this session is not deleted. After the direct link recovers, the LDP session does not
                 need to be reestablished or the FEC-label mapping information does not need to
                 be re-learned. This speeds up LDP convergence.

Session Hold Time
                 When LDP session protection is configured, you can configure a session hold time.
                 After a local adjacency established using Link Hello messages is torn down, a
                 remote adjacency established using Targeted Hello messages continues to
                 maintain an LDP session within the configured session hold time. If the local
                 adjacency does not recover after the session hold time elapses, the remote
                 adjacency is torn down, and the LDP session maintained using the remote
                 adjacency is also torn down.

3.17.2 Configuring LDP Session Protection
Prerequisites
                 Before configuring LDP session protection, you have completed the following task:
                 ●      Configure a local LDP session.

Context
                 If a direct link for a local LDP session fails, the LDP adjacency is torn down, and
                 the session and labels are deleted. After the direct link recovers, the LDP session is
                 reestablished and labels are re-distributed, after which the LDP LSP converges
                 again. Before the LSP is reestablished, LDP LSP traffic is lost.
                 To accelerate LDP LSP convergence and reduce traffic loss, configure LDP session
                 protection. After LDP session protection is configured, LDP establishes a remote
                 adjacency when establishing a local one. Both of them will be used to maintain

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            147
MPLS Configuration
MPLS Configuration                                                                         3 MPLS LDP Configuration


                 the LDP session. If the direct link of an LDP session fails but another path is
                 available to ensure route reachability, the remote adjacency can maintain the LDP
                 session. After the direct link recovers, the local outgoing label can still be used,
                 and does not need to be re-distributed by the downstream node. In addition, the
                 LDP session does not need to be reestablished. This speeds up LDP LSP
                 convergence and reduces traffic loss.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Enable LDP session protection.
                 session protection [ peer-group peer-group-name ] [ duration { infinite | time-value } ]

                 When the hold time of a local LDP session is set to different values on the two
                 ends of the session, the smaller value between them takes effect if the link of the
                 local LDP session fails. To configure LDP session protection, you are advised to
                 configure the same LDP session protection parameters on the two ends of the
                 session.

                 ----End

Verifying the Configuration
                 Run the display mpls ldp remote-peer command to check the configuration and
                 validity of LDP session protection.

3.17.3 Example for Configuring LDP Session Protection
Networking Requirements
                 On the network shown in Figure 3-29, an LDP session is established between PE1
                 and PE2. If the direct link between PE1 and PE2 fails, the LDP session and the peer
                 relationship are expected to remain connected. To achieve this, if a redundancy
                 link between PE1 and PE2 is available, configure LDP session protection for the
                 LDP session between PE1 and PE2. If the direct link between PE1 and PE2 fails,
                 LDP session protection prevents the LDP session from being disconnected and the
                 LDP peer relationship from being torn down.

                 Figure 3-29 Configuring LDP session protection
                         NOTE

                        Interfaces 1 and 2 in this example represent VLANIF100 and VLANIF200, respectively.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     148
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign an IP address to each interface and configure IGP routes.
                 2.     Configure a local LDP session.
                 3.     Configure LDP session protection.

Procedure
         Step 1 Assign an IP address to each interface and configure IGP routes. For configuration
                details, see Configuration Scripts in this section.
         Step 2 Configure a local LDP session.
                 # Configure PE1.
                 <PE1> system-view
                 [PE1] mpls lsr-id 1.1.1.1
                 [PE1] mpls
                 [PE1-mpls] quit
                 [PE1] mpls ldp
                 [PE1-mpls-ldp] quit
                 [PE1] interface vlanif 200
                 [PE1-Vlanif200] mpls
                 [PE1-Vlanif200] mpls ldp
                 [PE1-Vlanif200] quit

