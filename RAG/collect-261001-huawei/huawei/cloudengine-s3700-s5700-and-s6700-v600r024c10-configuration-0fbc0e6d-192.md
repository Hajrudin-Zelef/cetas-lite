---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-192
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [27857, 27953]
sha256: 3ea9ee0a1894924254d3832099a4feb83e5c94478e27b143d5e1ee0b66cb25db
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 6 Create the primary MPLS TE tunnel on the ingress LSR1.
                 # Configure an explicit path for the primary tunnel.
                 [LSR1] explicit-path pri-path
                 [LSR1-explicit-path-pri-path] next hop 10.1.1.2
                 [LSR1-explicit-path-pri-path] next hop 10.1.2.2
                 [LSR1-explicit-path-pri-path] next hop 10.1.3.2
                 [LSR1-explicit-path-pri-path] next hop 4.4.4.9
                 [LSR1-explicit-path-pri-path] quit

                 # Configure the primary MPLS TE tunnel and bind it to the explicit path.
                 [LSR1] interface Tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 4.4.4.9
                 [LSR1-Tunnel1] mpls te tunnel-id 1
                 [LSR1-Tunnel1] mpls te priority 4 3
                 [LSR1-Tunnel1] mpls te path explicit-path pri-path

                 # Enable TE FRR.
                 [LSR1-Tunnel1] mpls te fast-reroute
                 [LSR1-Tunnel1] quit

                 ----End

Verifying the Configuration
                 # Run the display interface tunnel command on LSR1. The command output
                 shows that the tunnel interface state is up.
                 [LSR1] display interface tunnel
                 Tunnel1 current state : UP (ifindex: 28)
                 Line protocol current state : UP
                 Last line protocol up time : 2024-04-08 14:31:35
                 Description:
                 ...

                 # Run the display mpls te tunnel command on each node. The command output
                 shows that two tunnels pass through LSR1 and LSR2, and three tunnels (one
                 primary tunnel and two bypass tunnels) pass through LSR3.
                 [LSR1]display mpls te tunnel
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name
                 -------------------------------------------------------------------------------
                 1.1.1.9        4.4.4.9       1099 -/18              I Tunnel1
                 1.1.1.9        3.3.3.9       1100 -/16              I AutoBypassTunnel_1.1.1.9_3.3.3.9_32769
                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress
                 [LSR2]display mpls te tunnel
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name
                 -------------------------------------------------------------------------------
                 1.1.1.9        4.4.4.9       1099 18/18              T Tunnel1
                 2.2.2.9        3.3.3.9       13 -/17              I AutoBypassTunnel_2.2.2.9_3.3.3.9_32769
                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress
                 [LSR3]display mpls te tunnel
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name
                 -------------------------------------------------------------------------------


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                            466
MPLS Configuration
MPLS Configuration                                                                                4 MPLS TE Configuration

                 1.1.1.9        4.4.4.9       1099 18/3             T Tunnel1
                 1.1.1.9        3.3.3.9       1100 3/-             E AutoBypassTunnel_1.1.1.9_3.3.3.9_32769
                 2.2.2.9        3.3.3.9       13 3/-              E AutoBypassTunnel_2.2.2.9_3.3.3.9_32769
                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress
                 [LSR4]display mpls te tunnel
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name
                 -------------------------------------------------------------------------------
                 1.1.1.9        4.4.4.9       1099 3/-             E Tunnel1
                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress
                 [LSR5]display mpls te tunnel
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name
                 -------------------------------------------------------------------------------
                 2.2.2.9        3.3.3.9       13 17/3              T AutoBypassTunnel_2.2.2.9_3.3.3.9_32769
                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress
                 [LSR6]display mpls te tunnel
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name
                 -------------------------------------------------------------------------------
                 1.1.1.9        3.3.3.9       1100 16/3             T AutoBypassTunnel_1.1.1.9_3.3.3.9_32769
                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress

