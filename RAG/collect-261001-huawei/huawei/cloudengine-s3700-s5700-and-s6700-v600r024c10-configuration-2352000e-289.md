---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-289
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2009-08-15", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [42583, 42767]
sha256: 74361d59adf57e742846d99dfc035663aa9289c5cfded482117833f20ced947f
---

                    # Configure ASBR_PE2.
                    [ASBR_PE2] vsi a1 static
                    [ASBR_PE2-vsi-a1] pwsignal ldp
                    [ASBR_PE2-vsi-a1-ldp] vsi-id 3
                    [ASBR_PE2-vsi-a1-ldp] peer 4.4.4.4
                    [ASBR_PE2-vsi-a1-ldp] quit
                    [ASBR_PE2-vsi-a1] quit
                    [ASBR_PE2] interface vlanif 30
                    [ASBR_PE2-Vlanif30] l2 binding vsi a1
                    [ASBR_PE2-Vlanif30] quit

                    # Configure PE2.
                    [PE2] vsi a1 static
                    [PE2-vsi-a1] pwsignal ldp
                    [PE2-vsi-a1-ldp] vsi-id 3
                    [PE2-vsi-a1-ldp] peer 3.3.3.3
                    [PE2-vsi-a1-ldp] quit
                    [PE2-vsi-a1] quit
                    [PE2] interface vlanif 50
                    [PE2-Vlanif50] l2 binding vsi a1
                    [PE2-Vlanif50] quit

                    ----End

Verifying the Configuration
                    # Check detailed information about the VSI named a1 on PE1.
                    [PE1] display vsi name a1 verbose



Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                      684
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration

                    ***VSI Name             : a1
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
                       Create Time          : 0 days, 3 hours, 30 minutes, 31 seconds
                       VSI State         : up

                      VSI ID            :2
                     *Peer Router ID        : 2.2.2.2
                      Negotiation-vc-id      :2
                      primary or secondary : primary
                      ignore-standby-state : no
                      VC Label           : 23552
                      Peer Type           : dynamic
                      Session           : up
                      Tunnel ID           : 0x20020
                      Broadcast Tunnel ID : 0x20020
                      Broad BackupTunnel ID : 0x0
                      CKey              :6
                      NKey              :5
                      Stp Enable           :0
                      PwIndex             :0
                      Control Word          : disable

                      Interface Name        : Vlanif10
                      State          : up
                      Access Port       : false
                      Last Up Time        : 2009-08-15 15:41:59
                      Total Up Time       : 0 days, 0 hours, 1 minutes, 2 seconds

                     **PW Information:

                     *Peer Ip Address        : 2.2.2.2
                      PW State             : up
                      Local VC Label         : 23552
                      Remote VC Label           : 23552
                      Remote Control Word : disable
                      PW Type              : label
                      Local VCCV            : alert lsp-ping bfd
                      Remote VCCV              : alert lsp-ping bfd
                      Tunnel ID           : 0x20020
                      Broadcast Tunnel ID : 0x20020
                      Broad BackupTunnel ID : 0x0
                      Ckey             : 0x6
                      Nkey              : 0x5
                      Main PW Token             : 0x20020
                      Slave PW Token           : 0x0
                      Tnl Type           : LSP
                      OutInterface          : Vlanif20
                      Backup OutInterface :
                      Stp Enable           :0
                      PW Last Up Time           : 2009-08-15 15:41:59
                      PW Total Up Time           : 0 days, 0 hours, 1 minutes, 3 seconds

                    The command output shows that a PW to PE2 has been established for the VSI
                    named a1 and the VSI status is Up.

                    # Perform a ping test to check the connectivity.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                              685
VPN Configuration
VPN Configuration                                                                    6 VPLS Configuration

                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=172 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=156 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=156 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=156 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=156 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 156/159/172 ms


Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         vlan 10
                         #
                         interface Vlanif10
                          ip address 10.1.1.1 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 10
                         #
                         return
                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 10 20
                         #
                         mpls lsr-id 1.1.1.1
                         mpls
                         #
                         mpls l2vpn
                         #
                         vsi a1 static
                          pwsignal ldp
                           vsi-id 2
                           peer 2.2.2.2
                         #
                         mpls ldp
                         #
                         isis 1
                          network-entity 10.0000.0000.0001.00
                         #
                         interface Vlanif10
                          l2 binding vsi a1
                         #
                         interface Vlanif20
                          ip address 1.1.1.1 255.255.255.0
                          isis enable 1
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


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         686
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

