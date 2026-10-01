---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-292
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [43102, 43266]
sha256: 8516b87bdf5ce7eec2ff7dda2f04446ae79def43f66c86364e3961832ac0ce2d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Establish an MP IBGP connection and enable BGP VPLS.

                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.2 as-number 100
                    [PE1-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [PE1-bgp] l2vpn-ad-family
                    [PE1-bgp-af-l2vpn-ad] peer 2.2.2.2 enable
                    [PE1-bgp-af-l2vpn-ad] peer 2.2.2.2 signaling vpls
                    [PE1-bgp-af-l2vpn-ad] quit
                    [PE1-bgp] quit

                    # Configure ASBR_PE1.
                    [ASBR_PE1] bgp 100
                    [ASBR_PE1-bgp] peer 1.1.1.1 as-number 100
                    [ASBR_PE1-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [ASBR_PE1-bgp] l2vpn-ad-family
                    [ASBR_PE1-bgp-af-l2vpn-ad] peer 1.1.1.1 enable
                    [ASBR_PE1-bgp-af-l2vpn-ad] peer 1.1.1.1 signaling vpls
                    [ASBR_PE1-bgp-af-l2vpn-ad] quit
                    [ASBR_PE1-bgp] quit

                    # Configure PE2.
                    [PE2] bgp 200
                    [PE2-bgp] peer 3.3.3.3 as-number 200
                    [PE2-bgp] peer 3.3.3.3 connect-interface loopBack1
                    [PE2-bgp] l2vpn-ad-family
                    [PE2-bgp-af-l2vpn-ad] peer 3.3.3.3 enable
                    [PE2-bgp-af-l2vpn-ad] peer 3.3.3.3 signaling vpls
                    [PE2-bgp-af-l2vpn-ad] quit
                    [PE2-bgp] quit

                    # Configure ASBR_PE2.
                    [ASBR_PE2] bgp 200
                    [ASBR_PE2-bgp] peer 4.4.4.4 as-number 200
                    [ASBR_PE2-bgp] peer 4.4.4.4 connect-interface loopback 1
                    [ASBR_PE2-bgp] l2vpn-ad-family
                    [ASBR_PE2-bgp-af-l2vpn-ad] peer 4.4.4.4 enable
                    [ASBR_PE2-bgp-af-l2vpn-ad] peer 4.4.4.4 signaling vpls
                    [ASBR_PE2-bgp-af-l2vpn-ad] quit
                    [ASBR_PE2-bgp] quit


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       692
VPN Configuration
VPN Configuration                                                              6 VPLS Configuration


         Step 5 Enable MPLS L2VPN.
                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit

                    # Configure ASBR_PE1.
                    [ASBR_PE1] mpls l2vpn
                    [ASBR_PE1-l2vpn] quit

                    # Configure ASBR_PE2.
                    [ASBR_PE2] mpls l2vpn
                    [ASBR_PE2-l2vpn] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit

         Step 6 Configure a VSI and bind it to an AC interface.
                    # Configure PE1.
                    [PE1] vsi v1 auto
                    [PE1-vsi-v1] pwsignal bgp
                    [PE1-vsi-v1-bgp] route-distinguisher 100:1
                    [PE1-vsi-v1-bgp] vpn-target 1:1 import-extcommunity
                    [PE1-vsi-v1-bgp] vpn-target 1:1 export-extcommunity
                    [PE1-vsi-v1-bgp] site 1 range 5 default-offset 0
                    [PE1-vsi-v1-bgp] quit
                    [PE1-vsi-v1] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] l2 binding vsi v1
                    [PE1-Vlanif10] quit

                    # Configure ASBR_PE1.
                    [ASBR_PE1] vsi v1 auto
                    [ASBR_PE1-vsi-v1] pwsignal bgp
                    [ASBR_PE1-vsi-v1-bgp] route-distinguisher 100:2
                    [ASBR_PE1-vsi-v1-bgp] vpn-target 1:1 import-extcommunity
                    [ASBR_PE1-vsi-v1-bgp] vpn-target 1:1 export-extcommunity
                    [ASBR_PE1-vsi-v1-bgp] site 2 range 5 default-offset 0
                    [ASBR_PE1-vsi-v1-bgp] quit
                    [ASBR_PE1-vsi-v1] quit
                    [ASBR_PE1] interface vlanif 30
                    [ASBR_PE1-Vlanif30] l2 binding vsi v1
                    [ASBR_PE1-Vlanif30] quit

                    # Configure ASBR_PE2.
                    [ASBR_PE2] vsi v1 auto
                    [ASBR_PE2-vsi-v1] pwsignal bgp
                    [ASBR_PE2-vsi-v1-bgp] route-distinguisher 200:1
                    [ASBR_PE2-vsi-v1-bgp] vpn-target 1:1 import-extcommunity
                    [ASBR_PE2-vsi-v1-bgp] vpn-target 1:1 export-extcommunity
                    [ASBR_PE2-vsi-v1-bgp] site 1 range 5 default-offset 0
                    [ASBR_PE2-vsi-v1-bgp] quit
                    [ASBR_PE2-vsi-v1] quit
                    [ASBR_PE2] interface vlanif 30
                    [ASBR_PE2-Vlanif30] l2 binding vsi v1
                    [ASBR_PE2-Vlanif30] quit

                    # Configure PE2.
                    [PE2] vsi v1 auto
                    [PE2-vsi-v1] pwsignal bgp


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   693
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration

                    [PE2-vsi-v1-bgp] route-distinguisher 200:2
                    [PE2-vsi-v1-bgp] vpn-target 1:1 import-extcommunity
                    [PE2-vsi-v1-bgp] vpn-target 1:1 export-extcommunity
                    [PE2-vsi-v1-bgp] site 2 range 5 default-offset 0
                    [PE2-vsi-v1-bgp] quit
                    [PE2-vsi-v1] quit
                    [PE2] interface vlanif 50
                    [PE2-Vlanif50] l2 binding vsi v1
                    [PE2-Vlanif50] quit

                    ----End

Verifying the Configuration
                    # Check information about BGP VPLS connections on PE1.
                    [PE1] display vpls connection bgp verbose
                    VSI Name: v1                     Signaling: bgp
                     **Remote Site ID       :2
                       VC State        : up
                       RD            : 100:2
                       Encapsulation       : vlan
                       MTU             : 1500
                       Peer Ip Address : 2.2.2.2
                       PW Type           : label
                       Local VC Label      : 35842
                       Remote VC Label : 31745
                       Tunnel Policy      : --
                       Tunnel ID        : 0x20020
                       Remote Label Block : 31744/5/0
                       Export vpn target : 1:1

                    The command output shows that the VSI named v1 is up.

                    # Perform a ping test to check the connectivity.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=90 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=77 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=34 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=46 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=94 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 34/68/94 ms

                    The command output shows that CE1 and CE2 can ping each other successfully.

