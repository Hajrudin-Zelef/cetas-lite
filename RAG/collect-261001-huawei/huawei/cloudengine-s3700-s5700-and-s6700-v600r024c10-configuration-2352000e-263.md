---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-263
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [38547, 38714]
sha256: f7a75443fa6903951d936d4c743b74ed2c5ef2df6baf1ce4b6ecc42c8376f509
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The configurations of PE1, the P, PE2, and CE2 are similar to the configuration of
                    CE1. For detailed configurations, see Configuration Scripts.

                          NOTE

                         Do not add AC-side and PW-side physical interfaces on a PE to the same VLAN. Otherwise,
                         a loop may occur.

         Step 2 Configure a routing protocol for communication between devices.
                    In this example, OSPF is configured. When configuring OSPF, configure PE1, the P,
                    and PE2 to advertise their loopback interfaces' 32-bit IP addresses (used as LSR
                    IDs).
                    # Configure PE1.
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [PE1-LoopBack1] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 192.168.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display ip routing-table command
                    on PE1, PE2, and the P. The command outputs show that PE1, PE2, and the P have
                    learned routes from each other.
         Step 3 Configure basic MPLS functions and LDP.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 617
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] mpls
                    [PE1-Vlanif20] mpls ldp
                    [PE1-Vlanif20] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display mpls ldp session command
                    on PE1, the P, and PE2. The command outputs show that the status of the peer
                    relationship between PE1 and the P or between PE2 and the P is Operational,
                    indicating that the peer relationship has been established. Run the display mpls
                    lsp command. The command output shows LSP establishment information.
         Step 4 Establish a remote LDP session.
                    # Configure PE1.
                    [PE1] mpls ldp remote-peer 3.3.3.9
                    [PE1-mpls-ldp-remote-3.3.3.9] remote-ip 3.3.3.9
                    [PE1-mpls-ldp-remote-3.3.3.9] quit

                    # Configure PE2.
                    [PE2] mpls ldp remote-peer 1.1.1.9
                    [PE2-mpls-ldp-remote-1.1.1.9] remote-ip 1.1.1.9
                    [PE2-mpls-ldp-remote-1.1.1.9] quit

                    After the configurations are complete, run the display mpls ldp session command
                    on PE1 or PE2. The command output shows that the status of the peer
                    relationship between PE1 and PE2 is Operational, indicating that the remote peer
                    relationship has been established.
         Step 5 Enable MPLS L2VPN.
                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit

         Step 6 Configure LDP VPLS.
                    # Configure PE1.
                    [PE1] vsi a2 static
                    [PE1-vsi-a2] pwsignal ldp
                    [PE1-vsi-a2-ldp] vsi-id 2
                    [PE1-vsi-a2-ldp] peer 3.3.3.9
                    [PE1-vsi-a2-ldp] quit
                    [PE1-vsi-a2] quit

                    # Configure PE2.
                    [PE2] vsi a2 static
                    [PE2-vsi-a2] pwsignal ldp


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          618
VPN Configuration
VPN Configuration                                                                       6 VPLS Configuration

                    [PE2-vsi-a2-ldp] vsi-id 2
                    [PE2-vsi-a2-ldp] peer 1.1.1.9
                    [PE2-vsi-a2-ldp] quit
                    [PE2-vsi-a2] quit

         Step 7 Bind a VSI to an interface.

                    # Configure PE1.
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] l2 binding vsi a2
                    [PE1-Vlanif10] quit

                    # Configure PE2.
                    [PE2] interface vlanif 40
                    [PE2-Vlanif40] l2 binding vsi a2
                    [PE2-Vlanif40] quit

                    ----End

Verifying the Configuration
                    # Check detailed information about the VSI named a2 on PE1.
                    [PE1] display vsi name a2 verbose

                    ***VSI Name                 : a2
                       Work Mode                 : normal
                       Administrator VSI          : no
                       Isolate Spoken          : disable
                       VSI Index            :0
                       PW Signaling            : ldp
                       Member Discovery Style : static
                       PW MAC Learn Style             : unqualify
                       Encapsulation Type           : vlan
                       MTU                 : 1500
                       Diffserv Mode            : uniform
                       Service Class         : cs-
                       Color             : --
                       DomainId               :-
                       Domain Name                  :-
                       Ignore AcState          : disable
                       P2P VSI             : disable
                       Multicast Fast Switch : disable
                       Create Time            : 0 days, 0 hours, 1 minutes, 3 seconds
                       VSI State           : up
                       Resource Status           : --

                       VSI ID            :2
                      *Peer Router ID        : 3.3.3.9
                       Negotiation-vc-id      :2
                       Encapsulation Type      : vlan
                       primary or secondary : primary
                       ignore-standby-state : no
                       VC Label           : 4096
                       Peer Type           : dynamic
                       Session           : up
                       Tunnel ID           : 0x0000000001004cab43
                       Broadcast Tunnel ID : --
                       Broad BackupTunnel ID : --
                       CKey              :6
                       NKey              :5
                       Stp Enable           :0
                       PwIndex             :0
                       Control Word          : disable

                      Interface Name         : Vlanif10


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                           619
VPN Configuration
VPN Configuration                                                                             6 VPLS Configuration

