---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-270
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [39650, 39817]
sha256: fa380cb01850dcd0b85166f9704b45f15b5e9d4e7e5b184e8b1dcbb837dbb9b4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


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
         Step 4 Establish BGP peer relationships and enable BGP peers to exchange VPLS
                information.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 3.3.3.9 as-number 100
                    [PE1-bgp] peer 3.3.3.9 connect-interface loopback 1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 636
VPN Configuration
VPN Configuration                                                               6 VPLS Configuration

                    [PE1-bgp] l2vpn-ad-family
                    [PE1-bgp-af-l2vpn-ad] peer 3.3.3.9 enable
                    [PE1-bgp-af-l2vpn-ad] peer 3.3.3.9 signaling vpls
                    [PE1-bgp-af-l2vpn-ad] quit
                    [PE1-bgp] quit

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.9 as-number 100
                    [PE2-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [PE2-bgp] l2vpn-ad-family
                    [PE2-bgp-af-l2vpn-ad] peer 1.1.1.9 enable
                    [PE2-bgp-af-l2vpn-ad] peer 1.1.1.9 signaling vpls
                    [PE2-bgp-af-l2vpn-ad] quit
                    [PE2-bgp] quit

         Step 5 Enable MPLS L2VPN.

                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit

         Step 6 Configure a VSI.
                          NOTE

                        The site IDs at both ends of a VSI must be different.

                    # Configure PE1.
                    [PE1] vsi bgp1 auto
                    [PE1-vsi-bgp1] pwsignal bgp
                    [PE1-vsi-bgp1-bgp] route-distinguisher 8.8.8.1:1
                    [PE1-vsi-bgp1-bgp] vpn-target 100:1 import-extcommunity
                    [PE1-vsi-bgp1-bgp] vpn-target 100:1 export-extcommunity
                    [PE1-vsi-bgp1-bgp] site 1 range 5 default-offset 0
                    [PE1-vsi-bgp1-bgp] quit
                    [PE1-vsi-bgp1] quit

                    # Configure PE2.
                    [PE2] vsi bgp1 auto
                    [PE2-vsi-bgp1] pwsignal bgp
                    [PE2-vsi-bgp1-bgp] route-distinguisher 9.9.9.2:1
                    [PE2-vsi-bgp1-bgp] vpn-target 100:1 import-extcommunity
                    [PE2-vsi-bgp1-bgp] vpn-target 100:1 export-extcommunity
                    [PE2-vsi-bgp1-bgp] site 2 range 5 default-offset 0
                    [PE2-vsi-bgp1-bgp] quit
                    [PE2-vsi-bgp1] quit

         Step 7 Bind AC interfaces to the VSI.

                    # Bind VLANIF 10 on PE1 to the VSI.
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] l2 binding vsi bgp1
                    [PE1-Vlanif10] quit

                    # Bind VLANIF 40 on PE2 to the VSI.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                    637
VPN Configuration
VPN Configuration                                                                        6 VPLS Configuration

                    [PE2] interface vlanif 40
                    [PE2-Vlanif40] l2 binding vsi bgp1
                    [PE2-Vlanif40] quit

                    ----End

Verifying the Configuration
                    # Check detailed information about the VSI named bgp1 on PE1.
                    [PE1] display vsi name bgp1 verbose

                    ***VSI Name                : bgp1
                       Work Mode                 : normal
                       Administrator VSI          : no
                       Isolate Spoken          : disable
                       VSI Index            :9
                       PW Signaling             : bgp
                       Member Discovery Style : auto
                       PW MAC Learn Style             : unqualify
                       Encapsulation Type           : vlan
                       MTU                 : 1500
                       Diffserv Mode            : uniform
                       Service Class         : cs-
                       Color             : --
                       DomainId               :-
                       Domain Name                  :-
                       Ignore AcState           : disable
                       P2P VSI             : disable
                       Multicast Fast Switch : disable
                       Create Time            : 0 days, 0 hours, 2 minutes, 34 seconds
                       VSI State           : up
                       Resource Status           : --

                      BGP RD            : 8.8.8.1:1
                      SiteID/Range/Offset : 1/5/0
                      Import vpn target     : 100:1
                      Export vpn target    : 100:1
                      Remote Label Block      : 1016576/5/0
                      Local Label Block    : 0/1016576/5/0

