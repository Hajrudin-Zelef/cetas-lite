---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-190
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [27639, 27811]
sha256: e88386867325c91af516bc156d52d6fe981d50d656a612306e2f8ee9e29cdbb6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

5.6.5 Example for Configuring a Local BGP VPWS Connection
Networking Requirements
                    In Figure 5-17, CE1 and CE2 are connected to the same PE. A local BGP VPWS
                    connection needs to be established between CE1 and CE2 for them to
                    communicate. In this case, the PE functions like a Layer 2 switch and can complete
                    label switching without having BGP configured.

                    Figure 5-17 Network diagram of configuring a local BGP VPWS connection
                          NOTE

                         In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                               441
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Enable MPLS functions and MPLS L2VPN on the PE.
                    2.   Configure a local BGP VPWS connection between CE1 and CE2 on the PE.
                                NOTE

                         By default, the Link-type Negotiation Protocol (LNP) is enabled globally on the device. If a
                         VLANIF interface is used as an AC interface for L2VPN, the configuration conflicts with LNP.
                         In this case, run the lnp disable command in the system view to disable LNP.


Procedure
         Step 1 Configure the VLANs that interfaces belong to and assign IP addresses to VLANIF
                interfaces.

                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 10
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/1] quit
                    [CE1] interface vlanif 10
                    [CE1-Vlanif10] ip address 10.1.1.1 255.255.255.0
                    [CE1-Vlanif10] quit

                    # Configure CE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] vlan batch 20
                    [CE2] interface 10ge 1/0/2
                    [CE2-10GE1/0/2] port link-type trunk
                    [CE2-10GE1/0/2] port trunk allow-pass vlan 20
                    [CE2-10GE1/0/2] quit
                    [CE2] interface vlanif 20
                    [CE2-Vlanif20] ip address 10.1.1.2 255.255.255.0
                    [CE2-Vlanif20] quit

                    The configuration of the PE is similar to the configuration of a CE. For detailed
                    configurations, see Configuration Scripts.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     442
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration


         Step 2 Enable basic MPLS functions on the PE.
                    [PE] interface loopback 1
                    [PE-LoopBack1] ip address 1.1.1.9 32
                    [PE-LoopBack1] quit
                    [PE] mpls lsr-id 1.1.1.9
                    [PE] mpls
                    [PE-mpls] quit

         Step 3 Configure a local BGP VPWS connection.
                    [PE] mpls l2vpn
                    [PE-l2vpn] quit
                    [PE] mpls l2vpn vpn1 encapsulation vlan
                    [PE-mpls-l2vpn-vpn1] route-distinguisher 100:1
                    [PE-mpls-l2vpn-vpn1] ce ce1 id 1 range 10
                    [PE-mpls-l2vpn-ce-vpn1-ce1] connection ce-offset 2 interface vlanif 10
                    [PE-mpls-l2vpn-ce-vpn1-ce1] quit
                    [PE-mpls-l2vpn-vpn1] ce ce2 id 2 range 10
                    [PE-mpls-l2vpn-ce-vpn1-ce2] connection ce-offset 1 interface vlanif 20
                    [PE-mpls-l2vpn-ce-vpn1-ce2] quit
                    [PE-mpls-l2vpn-vpn1] quit

                    ----End


Verifying the Configuration
                    # Check VPWS connection information on the PE. The command output shows
                    that two local VPWS connections have been established and are in the up state.
                    [PE] display mpls l2vpn connection
                    2 total connections,
                    connections: 2 up, 0 down, 2 local, 0 remote, 0 unknown

                    VPN name: vpn1,
                    2 total connections,
                    connections: 2 up, 0 down , 2 local, 0 remote, 0 unknown

                      CE name: ce1, id: 1,
                      Rid type status peer-id          route-distinguisher interface
                      primary or not
                    ----------------------------------------------------------------------------
                      2 loc up        ---          ---             Vlanif10
                      primary
                      CE name: ce2, id: 2,
                      Rid type status peer-id          route-distinguisher interface
                      primary or not
                    ----------------------------------------------------------------------------
                      1 loc up        ---          ---             Vlanif20
                      primary

                    # CE1 and CE2 can ping each other successfully.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                       Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=7 ms
                       Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=3 ms
                       Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=4 ms
                       Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=2 ms
                       Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=3 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 2/3/7 ms




Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                    443
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration


Configuration Scripts
                    ●   CE1
                        #
                        sysname CE1
                        #
                        vlan batch 10
                        #
                        interface Vlanif10
                         ip address 10.1.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 20
                        #
                        interface Vlanif20
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        return

