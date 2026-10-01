---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-324
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [48281, 48412]
sha256: 6c867a650101867362efa9a824baef4cf024500dccff70733e3a2f3046868113
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    After the configuration is complete, run the display vsi name s1 verbose
                    command on PE3. The command output shows that PE3 has established PWs with
                    PE1 (1.1.1.1) and PE2 (2.2.2.2).
                    [PE3] display vsi name s1 verbose
                     ***VSI Name             : s1
                        Administrator VSI       : no
                        Isolate Spoken        : disable
                        VSI Index          :2
                        PW Signaling          : ldp
                        Member Discovery Style : static
                        PW MAC Learn Style          : unqualify
                        Encapsulation Type        : vlan
                        MTU               : 1500
                        Diffserv Mode          : uniform
                        Mpls Exp           : --
                        DomainId            : 255
                        Domain Name               :
                        Ignore AcState        : disable
                        P2P VSI           : disable
                        Create Time          : 0 days, 1 hours, 19 minutes, 38 seconds
                        VSI State         : up

                       VSI ID            : 10
                      *Peer Router ID          : 1.1.1.1
                       Negotiation-vc-id        : 10
                       primary or secondary : primary
                       ignore-standby-state : no
                       VC Label            : 32891
                       Peer Type            : dynamic
                       Session            : up
                       Tunnel ID            : 0x0000000001004c4b41
                       Broadcast Tunnel ID : --
                       Broad BackupTunnel ID : --
                       CKey              :2
                       NKey               : 1862271177
                       Stp Enable            :0
                       PwIndex              :1
                       Control Word            : disable
                       BFD for PW             : unavailable
                      *Peer Router ID          : 2.2.2.2
                       Negotiation-vc-id        : 10


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                            777
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration

                      primary or secondary : primary
                      ignore-standby-state : no
                      VC Label           : 32892
                      Peer Type           : dynamic
                      Session           : up
                      Tunnel ID           : 0x0000000001004c4b42
                      Broadcast Tunnel ID : --
                      Broad BackupTunnel ID : --
                      CKey             :2
                      NKey              : 1862271178
                      Stp Enable           :0
                      PwIndex             :2
                      Control Word           : disable
                      BFD for PW            : unavailable

                     **PW Information:

                      *Peer Ip Address       : 1.1.1.1
                       PW State            : up
                       Local VC Label        : 32891
                       Remote VC Label          : 32890
                       Remote Control Word : disable
                       PW Type             : label
                       Local VCCV           : alert lsp-ping bfd
                       Remote VCCV             : alert lsp-ping bfd
                       Tunnel ID          : 0x0000000001004c4b41
                       Broadcast Tunnel ID : --
                       Broad BackupTunnel ID : --
                       Ckey             :2
                       Nkey             : 1862271177
                       Main PW Token            : 0x0
                       Slave PW Token          : 0x0
                       Tnl Type          : ldp
                       OutInterface         : LDP LSP
                       Backup OutInterface : --
                       Stp Enable          :0
                       PW Last Up Time          : 2016/06/14 17:35:12
                       PW Total Up Time          : 0 days, 1 hours, 19 minutes, 38 seconds
                      *Peer Ip Address       : 2.2.2.2
                       PW State            : up
                       Local VC Label        : 32892
                       Remote VC Label          : 32893
                       Remote Control Word : disable
                       PW Type             : label
                       Local VCCV           : alert lsp-ping bfd
                       Remote VCCV             : alert lsp-ping bfd
                       Tunnel ID          : 0x0000000001004c4b42
                       Broadcast Tunnel ID : --
                       Broad BackupTunnel ID : --
                       Ckey             :2
                       Nkey             : 1862271178
                       Main PW Token            : 0x0
                       Slave PW Token          : 0x0
                       Tnl Type          : ldp
                       OutInterface         : LDP LSP
                       Backup OutInterface : --
                       Stp Enable          :0
                       PW Last Up Time          : 2016/06/14 10:35:45
                       PW Total Up Time          : 0 days, 1 hours, 19 minutes, 45 seconds

                    The command output on CE2 shows that the link between CE2 and CE1 is blocked.
                    [CE2] display erps
                    D : Discarding
                    F : Forwarding
                    R : RPL Owner
                    N : RPL Neighbour
                    FS : Forced Device
                    MS : Manual Device
                    Total number of rings configured = 1


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                778
VPN Configuration
VPN Configuration                                                                                      6 VPLS Configuration

                    Ring Control WTR Timer Guard Timer Port 1                         Port 2
                    ID VLAN         (min)       (csec)
                    --------------------------------------------------------------------------------
                      1      100         5        200 (F)10GE1/0/1            (D,R)10GE1/0/2
                    --------------------------------------------------------------------------------

                    ----End

