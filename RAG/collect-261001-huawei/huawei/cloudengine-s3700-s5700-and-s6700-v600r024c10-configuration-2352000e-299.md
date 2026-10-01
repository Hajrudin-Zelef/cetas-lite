---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-299
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [44223, 44385]
sha256: f1ee2ec7b1a61c6024c64fc435df24750ef83c92220ba7d1b0df8e6ae335fde0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The configurations of PE2, PE3, PE4, and PE5 are similar to the configuration of
                    PE1. For detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display ip routing-table command
                    on PEs. The command outputs show that the PEs have learned each other's
                    loopback interface address.
         Step 3 Configure MPLS and public network tunnels.
                    In this example, LDP LSPs are configured.
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
                    [PE1] interface vlanif 40
                    [PE1-Vlanif40] mpls
                    [PE1-Vlanif40] mpls ldp
                    [PE1-Vlanif40] quit

                    The configurations of PE2, PE3, PE4, and PE5 are similar to the configuration of
                    PE1. For detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display mpls ldp session command
                    on PEs. The command outputs show that the status of the peer relationship
                    between PEs is Operational, indicating that the peer relationship has been
                    established. Run the display mpls lsp command. The command output shows LSP
                    establishment information.
         Step 4 Configure PE1, PE2, and PE3 to form an LDP VPLS network.
                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] vsi vsi1 static
                    [PE1-vsi-vsi1] pwsignal ldp
                    [PE1-vsi-vsi1-ldp] vsi-id 1
                    [PE1-vsi-vsi1-ldp] peer 2.2.2.9


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          711
VPN Configuration
VPN Configuration                                                                6 VPLS Configuration

                    [PE1-vsi-vsi1-ldp] peer 3.3.3.9
                    [PE1-vsi-vsi1-ldp] quit
                    [PE1-vsi-vsi1] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] vsi vsi1 static
                    [PE2-vsi-vsi1] pwsignal ldp
                    [PE2-vsi-vsi1-ldp] vsi-id 1
                    [PE2-vsi-vsi1-ldp] peer 1.1.1.9
                    [PE2-vsi-vsi1-ldp] peer 3.3.3.9
                    [PE2-vsi-vsi1-ldp] quit
                    [PE2-vsi-vsi1] quit

                    # Configure PE3.
                    [PE3] mpls l2vpn
                    [PE3-l2vpn] quit
                    [PE3] vsi vsi1
                    [PE3-vsi-vsi1] pwsignal ldp
                    [PE3-vsi-vsi1-ldp] vsi-id 1
                    [PE3-vsi-vsi1-ldp] peer 1.1.1.9 upe
                    [PE3-vsi-vsi1-ldp] peer 2.2.2.9 upe
                    [PE3-vsi-vsi1-ldp] quit
                    [PE3-vsi-vsi1] quit

                    # On PE1, bind an AC interface to a VSI.
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] l2 binding vsi vsi1
                    [PE1-Vlanif10] quit

         Step 5 Configure PE3, PE4, and PE5 to form a BGP AD VPLS network.
                    1.   Enable BGP peers to exchange VPLS member information.
                         # Configure PE3.
                         [PE3] bgp 100
                         [PE3-bgp] peer 4.4.4.9 as-number 100
                         [PE3-bgp] peer 4.4.4.9 connect-interface loopback1
                         [PE3-bgp] peer 5.5.5.9 as-number 100
                         [PE3-bgp] peer 5.5.5.9 connect-interface loopback1
                         [PE3-bgp] l2vpn-ad-family
                         [PE3-bgp-af-l2vpn-ad] peer 4.4.4.9 enable
                         [PE3-bgp-af-l2vpn-ad] peer 5.5.5.9 enable
                         [PE3-bgp-af-l2vpn-ad] quit
                         [PE3-bgp] quit

                         # Configure PE4.
                         <PE4> system-view
                         [PE4] bgp 100
                         [PE4-bgp] peer 3.3.3.9 as-number 100
                         [PE4-bgp] peer 3.3.3.9 connect-interface loopback1
                         [PE4-bgp] peer 5.5.5.9 as-number 100
                         [PE4-bgp] peer 5.5.5.9 connect-interface loopback1
                         [PE4-bgp] l2vpn-ad-family
                         [PE4-bgp-af-l2vpn-ad] peer 3.3.3.9 enable
                         [PE4-bgp-af-l2vpn-ad] peer 5.5.5.9 enable
                         [PE4-bgp-af-l2vpn-ad] quit
                         [PE4-bgp] quit

                         # Configure PE5.
                         <PE5> system-view
                         [PE5] bgp 100
                         [PE5-bgp] peer 3.3.3.9 as-number 100


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                    712
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration

                         [PE5-bgp] peer 3.3.3.9 connect-interface loopback1
                         [PE5-bgp] peer 4.4.4.9 as-number 100
                         [PE5-bgp] peer 4.4.4.9 connect-interface loopback1
                         [PE5-bgp] l2vpn-ad-family
                         [PE5-bgp-af-l2vpn-ad] peer 3.3.3.9 enable
                         [PE5-bgp-af-l2vpn-ad] peer 4.4.4.9 enable
                         [PE5-bgp-af-l2vpn-ad] quit
                         [PE5-bgp] quit
                    2.   Create a VSI and configure BGP AD signaling.
                         # Configure PE3.
                         [PE3] vsi vsi1
                         [PE3-vsi-vsi1] bgp-ad
                         [PE3-vsi-vsi1-bgpad] vpls-id 192.168.0.0:1
                         [PE3-vsi-vsi1-bgpad] vpn-target 100:1 import-extcommunity
                         [PE3-vsi-vsi1-bgpad] vpn-target 100:1 export-extcommunity
                         [PE3-vsi-vsi1-bgpad] quit
                         [PE3-vsi-vsi1] quit

                                NOTE

                         On PE3, LDP and BGP AD PWs must be configured in the same VSI.

                         # Configure PE4.
                         [PE4] mpls l2vpn
                         [PE4-l2vpn] quit
                         [PE4] vsi vsi1
                         [PE4-vsi-vsi1] bgp-ad
                         [PE4-vsi-vsi1-bgpad] vpls-id 192.168.0.0:1
                         [PE4-vsi-vsi1-bgpad] vpn-target 100:1 import-extcommunity
                         [PE4-vsi-vsi1-bgpad] vpn-target 100:1 export-extcommunity
                         [PE4-vsi-vsi1-bgpad] quit
                         [PE4-vsi-vsi1] quit

                         # Configure PE5.
                         [PE5] mpls l2vpn
                         [PE5-l2vpn] quit
                         [PE5] vsi vsi1
                         [PE5-vsi-vsi1] bgp-ad
                         [PE5-vsi-vsi1-bgpad] vpls-id 192.168.0.0:1
                         [PE5-vsi-vsi1-bgpad] vpn-target 100:1 import-extcommunity
                         [PE5-vsi-vsi1-bgpad] vpn-target 100:1 export-extcommunity
                         [PE5-vsi-vsi1-bgpad] quit
                         [PE5-vsi-vsi1] quit
                    3.   On PE5, bind an AC interface to the VSI.
                         [PE5] interface vlanif 80
                         [PE5-Vlanif80] l2 binding vsi vsi1
                         [PE5-Vlanif80] quit

                    ----End

