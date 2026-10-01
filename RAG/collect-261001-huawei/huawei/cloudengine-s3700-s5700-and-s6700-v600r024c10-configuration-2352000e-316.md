---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-316
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [46962, 47122]
sha256: a75cdb72e36ca8850c9fb8c1fe18eeee79455835d776e933308e1f2005f1867f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 755
VPN Configuration
VPN Configuration                                                               6 VPLS Configuration


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
                    [PE2] interface vlanif 100
                    [PE2-Vlanif100] l2 binding vsi a2
                    [PE2-Vlanif100] quit
                    [PE2] interface vlanif 200
                    [PE2-Vlanif200] l2 binding vsi a2
                    [PE2-Vlanif200] quit
                    [PE2] interface vlanif 300
                    [PE2-Vlanif300] l2 binding vsi a2
                    [PE2-Vlanif300] quit

         Step 8 Configure VPLS service isolation on PE2 and set the VSI attribute of VLANIF 100 to
                hub so that CE2 can communicate with CE3 and CE4 and that CE3 is isolated from
                CE4.
                    # Configure VPLS service isolation in the VSI named a2.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                   756
VPN Configuration
VPN Configuration                                                                        6 VPLS Configuration

                    [PE2] vsi a2 static
                    [PE2-vsi-a2] isolate spoken
                    [PE2-vsi-a2] quit
                    [PE2] interface vlanif 100
                    [PE2-Vlanif100] hub-mode enable
                    [PE2-Vlanif100] quit

                    ----End

Verifying the Configuration
                    # Check detailed information about the VSI named a2 on PE1.
                    [PE1]display vsi name a2 verbose

                    ***VSI Name                : a2
                       Work Mode                 : normal
                       Administrator VSI          : no
                       Isolate Spoken          : disable
                       VSI Index            :3
                       PW Signaling             : ldp
                       Member Discovery Style : static
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
                       Create Time            : 0 days, 1 hours, 5 minutes, 50 seconds
                       VSI State           : up
                       Resource Status           : --

                      VSI ID            :2
                     *Peer Router ID          : 3.3.3.9
                      Negotiation-vc-id        :2
                      Encapsulation Type        : vlan
                      primary or secondary : primary
                      ignore-standby-state : no
                      VC Label            : 1029
                      Peer Type            : dynamic
                      Session            : up
                      Tunnel ID            : 0x0000000001004c8b43
                      Broadcast Tunnel ID : --
                      Broad BackupTunnel ID : --
                      CKey              : 129
                      NKey               : 16777408
                      Stp Enable            :0
                      PwIndex              : 129
                      Control Word            : disable
                      BFD for PW             : unavailable

                      Interface Name         : Vlanif10
                      State          : up
                      Ac Block State      : unblocked
                      Access Port       : false
                      Last Up Time         : 2024/08/06 18:01:14
                      Total Up Time        : 0 days, 1 hours, 5 minutes, 44 seconds

                     **PW Information:

                     *Peer Ip Address       : 3.3.3.9
                      PW State            : up
                      Local VC Label        : 1029
                      Remote VC Label          : 1030


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                            757
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration

                      Remote Control Word : disable
                      PW Type              : label
                      Local VCCV            : alert lsp-ping bfd
                      Remote VCCV              : alert lsp-ping bfd
                      Tunnel ID           : 0x0000000001004c8b43
                      Broadcast Tunnel ID : --
                      Broad BackupTunnel ID : --
                      Ckey             : 129
                      Nkey              : 16777408
                      Main PW Token             : 0x0
                      Slave PW Token           : 0x0
                      Tnl Type           : ldp
                      OutInterface          : --
                      Backup OutInterface : --
                      Stp Enable           :0
                      Mac Flapping            :0
                      PW Last Up Time           : 2024/08/06 18:02:43
                      PW Total Up Time           : 0 days, 1 hours, 4 minutes, 15 seconds

