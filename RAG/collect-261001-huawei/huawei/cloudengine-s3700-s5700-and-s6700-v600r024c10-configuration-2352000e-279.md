---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-279
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [41001, 41144]
sha256: 94b89e52e59d54e1e314c78988f8d246184a168795f5363f2f9e783083a2554d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration

                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit

                    # Configure PE3.
                    [PE3] mpls l2vpn
                    [PE3-l2vpn] quit

         Step 6 Configure a VSI.
                    # Configure PE1.
                    [PE1] vsi vplsad1
                    [PE1-vsi-vplsad1] bgp-ad
                    [PE1-vsi-vplsad1-bgpad] vpls-id 172.16.1.1:1
                    [PE1-vsi-vplsad1-bgpad] vpn-target 100:1 import-extcommunity
                    [PE1-vsi-vplsad1-bgpad] vpn-target 100:1 export-extcommunity
                    [PE1-vsi-vplsad1-bgpad] quit
                    [PE1-vsi-vplsad1] quit

                    # Configure PE2.
                    [PE2] vsi vplsad1
                    [PE2-vsi-vplsad1] bgp-ad
                    [PE2-vsi-vplsad1-bgpad] vpls-id 172.16.1.1:1
                    [PE2-vsi-vplsad1-bgpad] vpn-target 100:1 import-extcommunity
                    [PE2-vsi-vplsad1-bgpad] vpn-target 100:1 export-extcommunity
                    [PE2-vsi-vplsad1-bgpad] quit
                    [PE2-vsi-vplsad1] quit

                    # Configure PE3.
                    [PE3] vsi vplsad1
                    [PE3-vsi-vplsad1] bgp-ad
                    [PE3-vsi-vplsad1-bgpad] vpls-id 172.16.1.1:1
                    [PE3-vsi-vplsad1-bgpad] vpn-target 100:1 import-extcommunity
                    [PE3-vsi-vplsad1-bgpad] vpn-target 100:1 export-extcommunity
                    [PE3-vsi-vplsad1-bgpad] quit
                    [PE3-vsi-vplsad1] quit

         Step 7 Bind AC interfaces to the VSI.
                    # Bind VLANIF 10 on PE1 to the VSI.
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] l2 binding vsi vplsad1
                    [PE1-Vlanif10] quit

                    # Bind VLANIF 50 on PE2 to the VSI.
                    [PE2] interface vlanif 50
                    [PE2-Vlanif50] l2 binding vsi vplsad1
                    [PE2-Vlanif50] quit

                    # Bind VLANIF 60 on PE3 to the VSI.
                    [PE3] interface vlanif 60
                    [PE3-Vlanif60] l2 binding vsi vplsad1
                    [PE3-Vlanif60] quit

                    ----End

Verifying the Configuration
                    # Check detailed information about the VSI named vplsad1 on PE1.
                    [PE1] display vsi name vplsad1 verbose



Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                       659
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration

                    ***VSI Name                : vplsad1
                       Work Mode                 : normal
                       Administrator VSI          : no
                       Isolate Spoken          : disable
                       VSI Index            :0
                       PW Signaling             : bgpad
                       Member Discovery Style : --
                       PW MAC Learn Style             : unqualify
                       Encapsulation Type           : vlan
                       MTU                 : 1500
                       Diffserv Mode            : uniform
                       Service Class         : cs-
                       Color             : --
                       DomainId               : --
                       Domain Name                  : --
                       Ignore AcState           : disable
                       P2P VSI             : disable
                       Multicast Fast Switch : disable
                       Create Time            : 0 days, 18 hours, 5 minutes, 30 seconds
                       VSI State           : up
                       Resource Status           : --

                      VPLS ID            : 172.16.1.1:1
                      RD               : 172.16.1.1:1
                      Import vpn target      : 100:1
                      Export vpn target     : 100:1
                      BGPAD VSI ID          : 1.1.1.9

                     *Peer Router ID        : 2.2.2.9
                      VPLS ID            : 172.16.1.1:1
                      SAII           : 1.1.1.9
                      TAII           : 2.2.2.9
                      VC Label           : 1024
                      Peer Type           : dynamic
                      Session           : up
                      Tunnel ID           : 0x80003f
                      Broadcast Tunnel ID : 0x80003f
                      CKey             :2
                      NKey              :1

                     *Peer Router ID        : 3.3.3.9
                      VPLS ID            : 172.16.1.1:1
                      SAII           : 1.1.1.9
                      TAII           : 3.3.3.9
                      VC Label           : 1025
                      Peer Type           : dynamic
                      Session           : up
                      Tunnel ID           : 0x800033
                      Broadcast Tunnel ID : 0x800033
                      CKey             :4
                      NKey              :3

                      Interface Name         : Vlanif10
                      State          : up
                      Ac Block State      : unblocked
                      Access Port       : false
                      Last Up Time         : 2012/07/06 15:54:46
                      Total Up Time        : 0 days, 0 hours, 58 minutes, 24 seconds

                     **PW Information:

                     *Peer Ip Address    : 2.2.2.9
                      PW State         : up
                      Local VC Label     : 1024
                      Remote VC Label       : 1024
                      PW Type          : label
                      Local VCCV        : alert lsp-ping bfd
                      Remote VCCV          : alert lsp-ping bfd
                      Tunnel ID       : 0x80003f



Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                             660
VPN Configuration
VPN Configuration                                                                               6 VPLS Configuration

