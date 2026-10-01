---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-38
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2018-06-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [4407, 4543]
sha256: e49bb42a258b73ea7b4f192cd0dc6c6822818a1fdd84ff7e1baabaf787935558
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Area 0.0.0.0 interface 11.11.11.1(Vlanif300)'s neighbors
                    Router ID: 2.2.2.9      Address: 11.11.11.2
                     State: Full    Mode:Nbr is Slave Priority: 1
                     DR: 1.1.1.9    BDR: 2.2.2.9      MTU: 1500
                     Dead timer due (in seconds) : 38
                     Retrans timer interval     :0
                     Neighbor up time            : 00h00m29s
                     Neighbor up time stamp         : 2018-06-08 01:41:57
                     Authentication Sequence        :0

         Step 2 Configure basic MPLS capabilities and MPLS LDP on the MPLS backbone network
                to establish LDP LSPs.
                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface Vlanif 300
                    [PE1-Vlanif300] mpls
                    [PE1-Vlanif300] mpls ldp
                    [PE1-Vlanif300] quit

                    # Configure the P.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          70
VPN Configuration
VPN Configuration                                                                                      3 IPv4 L3VPN Configuration

                    [P] mpls lsr-id 2.2.2.9
                    [P] mpls
                    [P-mpls] quit
                    [P] mpls ldp
                    [P-mpls-ldp] quit
                    [P] interface Vlanif 300
                    [P-Vlanif300] mpls
                    [P-Vlanif300] mpls ldp
                    [P-Vlanif300] quit
                    [P] interface Vlanif 200
                    [P-Vlanif200] mpls
                    [P-Vlanif200] mpls ldp
                    [P-Vlanif200] quit

                    # Configure PE2.
                    [PE2] mpls lsr-id 3.3.3.9
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface Vlanif 200
                    [PE2-Vlanif200] mpls
                    [PE2-Vlanif200] mpls ldp
                    [PE2-Vlanif200] quit

                    After the configuration is complete, LDP sessions are established between PE1 and
                    the P and between PE2 and the P. Run the display mpls ldp session command.
                    The command output shows that the session status is Operational. Then, run the
                    display mpls ldp lsp command. The command output shows that LDP LSPs have
                    been established.

                    The following example uses the command output on PE1.
                    [PE1] display mpls ldp session

                    LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.
                     -------------------------------------------------------------------------
                     Peer-ID           Status      LAM SsnRole SsnAge             KA-Sent/Rcv
                     -------------------------------------------------------------------------
                     2.2.2.9:0        Operational DU Passive 0006:20:55 39551/39552
                     -------------------------------------------------------------------------
                     TOTAL: 1 session(s) Found.

                    [PE1] display mpls ldp lsp
                     LDP LSP Information
                     -------------------------------------------------------------------------------
                     Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                     -------------------------------------------------------------------------------
                     DestAddress/Mask In/OutLabel UpstreamPeer NextHop                             OutInterface
                     -------------------------------------------------------------------------------
                     1.1.1.9/32        3/NULL          2.2.2.9        127.0.0.1       InLoop0
                    *1.1.1.9/32         Liberal/1024                 DS/2.2.2.9
                     2.2.2.9/32        NULL/3          -            11.11.11.2        Vlanif300
                     2.2.2.9/32        1024/3          2.2.2.9       11.11.11.2        Vlanif300
                     3.3.3.9/32        NULL/1025         -            11.11.11.2        Vlanif300
                     3.3.3.9/32        1025/1025         2.2.2.9       11.11.11.2        Vlanif300
                     -------------------------------------------------------------------------------
                     TOTAL: 5 Normal LSP(s) Found.
                     TOTAL: 1 Liberal LSP(s) Found.
                     TOTAL: 0 Frr LSP(s) Found.
                     An asterisk (*) before an LSP means the LSP is not established
                     An asterisk (*) before a Label means the USCB or DSCB is stale
                     An asterisk (*) before an UpstreamPeer means the session is stale
                     An asterisk (*) before a DS means the session is stale
                     An asterisk (*) before a NextHop means the LSP is FRR LSP


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                               71
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration


         Step 3 On PEs, create VPN instances, enable the IPv4 address family for these instances,
                and bind the interfaces connected to CEs to the VPN instances.

                    # Configure PE1.
                    [PE1] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv4-family
                    [PE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [PE1-vpn-instance-vpna-af-ipv4] vpn-target 111:1 both
                    [PE1-vpn-instance-vpna-af-ipv4] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE1] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv4-family
                    [PE1-vpn-instance-vpnb-af-ipv4] route-distinguisher 100:2
                    [PE1-vpn-instance-vpnb-af-ipv4] vpn-target 222:2 both
                    [PE1-vpn-instance-vpnb-af-ipv4] quit
                    [PE1-vpn-instance-vpnb] quit
                    [PE1] interface 10GE1/0/1
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] port trunk allow-pass vlan 100
                    [PE1-10GE1/0/1] quit
                    [PE1] interface Vlanif 100
                    [PE1-Vlanif100] ip binding vpn-instance vpna
                    [PE1-Vlanif100] ip address 10.1.1.2 24
                    [PE1-Vlanif100] quit
                    [PE1] interface 10GE1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 200
                    [PE1-10GE1/0/2] quit
                    [PE1] interface Vlanif 200
                    [PE1-Vlanif200] ip binding vpn-instance vpnb
                    [PE1-Vlanif200] ip address 10.2.1.2 24
                    [PE1-Vlanif200] quit

