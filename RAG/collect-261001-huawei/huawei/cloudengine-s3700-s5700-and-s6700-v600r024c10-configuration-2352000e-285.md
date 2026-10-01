---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-285
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [41888, 42055]
sha256: a5c27c58aa99da56caa6c646cf2a07d3f15cc0582e6ec44e99be3b2219ab6773
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ***VSI Name             : v123
                       Administrator VSI       : no
                       Isolate Spoken        : disable
                       VSI Index          :0
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
                       Create Time          : 0 days, 0 hours, 1 minutes, 3 seconds
                       VSI State         : up

                       VSI ID            : 123
                      *Peer Router ID         : 3.3.3.9
                       Negotiation-vc-id       : 123
                       primary or secondary : primary
                       ignore-standby-state : no
                       VC Label            : 4096
                       Peer Type            : dynamic
                       Session            : up
                       Tunnel ID            : 0x1c5
                       Broadcast Tunnel ID : 0x1c5
                       Broad BackupTunnel ID : 0x0
                       CKey              :9
                       NKey               :3
                       Stp Enable            :0
                       PwIndex              :0
                       Control Word           : disable


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                         673
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration

                     *Peer Router ID        : 1.1.1.9
                      Negotiation-vc-id      : 123
                      primary or secondary : primary
                      ignore-standby-state : no
                      VC Label           : 4097
                      Peer Type           : dynamic
                      Session           : up
                      Tunnel ID           : 0x1c3
                      Broadcast Tunnel ID : 0x1c3
                      Broad BackupTunnel ID : 0x0
                      CKey              :5
                      NKey              : 12
                      Stp Enable           :0
                      PwIndex             :0
                      Control Word          : disable

                     **PW Information:

                     *Peer Ip Address        : 1.1.1.9
                      PW State             : up
                      Local VC Label         : 4097
                      Remote VC Label           : 4096
                      Remote Control Word : disable
                      PW Type              : MEHVPLS
                      Local VCCV            : alert lsp-ping bfd
                      Remote VCCV              : alert lsp-ping bfd
                      Tunnel ID           : 0x1c3
                      Broadcast Tunnel ID : 0x1c3
                      Broad BackupTunnel ID : 0x0
                      Ckey             : 0x5
                      Nkey              : 0xc
                      Main PW Token             : 0x1c3
                      Slave PW Token           : 0x0
                      Tnl Type           : LSP
                      OutInterface          : Vlanif30
                      Backup OutInterface :
                      Stp Enable           :0
                      PW Last Up Time           : 2014/11/12 11:34:08
                      PW Total Up Time           : 0 days, 0 hours, 0 minutes, 22 seconds
                     *Peer Ip Address        : 3.3.3.9
                      PW State             : up
                      Local VC Label         : 4096
                      Remote VC Label           : 4096
                      Remote Control Word : disable
                      PW Type              : label
                      Local VCCV            : alert lsp-ping bfd
                      Remote VCCV              : alert lsp-ping bfd
                      Tunnel ID           : 0x1c5
                      Broadcast Tunnel ID : 0x1c5
                      Broad BackupTunnel ID : 0x0
                      Ckey             : 0x9
                      Nkey              : 0x3
                      Main PW Token             : 0x1c5
                      Slave PW Token           : 0x0
                      Tnl Type           : LSP
                      OutInterface          : Vlanif40
                      Backup OutInterface :
                      Stp Enable           :0
                      PW Last Up Time           : 2014/11/12 11:34:18
                      PW Total Up Time           : 0 days, 0 hours, 0 minutes, 12 seconds

                    The command output shows that the VSI named v123 is up and the corresponding
                    PW is also up.
                    CE1, CE2, and CE3 can ping each other successfully. After the shutdown command
                    is run on the UPE's or PE1's interface bound to the VSI, CE2 and CE3 cannot ping
                    each other. This indicates that data is transmitted through the PW of the VSI.



Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                               674
VPN Configuration
VPN Configuration                                                              6 VPLS Configuration


Configuration Scripts
                    ●   UPE
                        #
                        sysname UPE
                        #
                        vlan batch 10 20 30
                        #
                        mpls lsr-id 1.1.1.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi v123 static
                         pwsignal ldp
                          vsi-id 123
                          peer 2.2.2.9
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         l2 binding vsi v123
                        #
                        interface Vlanif20
                         l2 binding vsi v123
                        #
                        interface Vlanif30
                         ip address 192.0.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 192.0.2.0 0.0.0.255
                        #
                        return

