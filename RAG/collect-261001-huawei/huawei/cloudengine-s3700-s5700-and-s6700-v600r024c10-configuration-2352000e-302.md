---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-302
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [44706, 44876]
sha256: 59c1c31063ed5e469af83385203b19875079e9f508c60d5f3296e892dfd04376
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                          vpn-target 100:1 import-extcommunity
                          vpn-target 100:1 export-extcommunity
                        #
                        mpls ldp
                        #
                        interface Vlanif50
                         ip address 192.168.4.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif70
                         ip address 192.168.6.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif80
                         l2 binding vsi vsi1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 70
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 80
                        #
                        interface LoopBack1
                         ip address 5.5.5.9 255.255.255.255
                        #
                        bgp 100
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         peer 4.4.4.9 as-number 100
                         peer 4.4.4.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 3.3.3.9 enable
                          peer 4.4.4.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 3.3.3.9 enable
                          peer 4.4.4.9 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 5.5.5.9 0.0.0.0
                          network 192.168.4.0 0.0.0.255
                          network 192.168.6.0 0.0.0.255
                        #
                        return

                    ●   CE1
                        #
                        sysname CE1
                        #
                        vlan 10
                        #
                        interface Vlanif10
                         ip address 192.168.10.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   718
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration

                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan 80
                        #
                        interface Vlanif80
                         ip address 192.168.10.2 255.255.255.0
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 80
                        #
                        return



6.13 Configuring VPWS Accessing VPLS

6.13.1 Understanding VPWS Accessing VPLS
                    L2VPN PWs have different encapsulation types, such as Ethernet and VLAN. These
                    encapsulation types are essentially unrelated to services. Services can be
                    transmitted over VPWS PWs and VPLS PWs in succession as long as these PWs use
                    the same encapsulation type.


VPWS Accessing VPLS in Basic Mode
                    In Figure 6-32, CE1, CE2, and CE3 form a broadcast network. Both CE1 and CE2
                    need to communicate with CE3. PE1 and PE2 have limited MAC address capacity
                    and are not expected to store a large number of MAC addresses. In this case, you
                    can deploy a VPWS PW from PE1 to PE3 and from PE2 to PE3 to transmit traffic
                    from CE1 and CE2, respectively, and deploy a VPLS PW from PE3 to PE1 and from
                    PE3 to PE2 to transmit traffic from CE3.

                         NOTE

                        Due to the deployment of VPLS PWs on PE3 and the default use of split horizon in VPLS,
                        after PE3 receives traffic from CE1 and CE2, PE3 does not forward the traffic over VPLS PWs
                        to other devices. As a result, only communication between CE1 and CE3 and between CE2
                        and CE3 can be realized. To enable CE1, CE2, and CE3 to communicate with each other,
                        configure the VPLS PWs on PE3 as spoke PWs.


                    Figure 6-32 VPWS accessing VPLS in basic mode




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    719
VPN Configuration
VPN Configuration                                                                    6 VPLS Configuration


VPWS Accessing VPLS in Dual-Homed Mode
                    On the network shown in Figure 6-33, PE1 and PE2 each are dual-homed to PE3
                    and PE4 to enhance network reliability. Additionally, CE1 and CE2 need to
                    communicate with each other. Given that PE1 and PE2 have limited MAC address
                    capacity and cannot store a large number of MAC addresses, you can dual-home
                    both PE1 and PE2 to PE3 and PE4 using VPWS PWs to transmit traffic from CE1
                    and CE2. To achieve this, configure VPWS PWs from PE1 to PE3 and PE4 and from
                    PE2 to PE3 and PE4. Then, configure spoke PWs from PE3 to PE1 and PE2 and
                    from PE4 to PE1 and PE2. Finally, configure a hub PW between PE3 and PE4.
                    If PE1 and PE2 use LDP VPWS, you can configure PW redundancy to protect traffic.
                    That is, you can configure the primary and secondary PWs on PE1 and PE2 and
                    configure PW redundancy in master/slave mode. Configure PWs from PE1 and PE2
                    to PE3 as primary PWs (PW1 and PW3) and PWs from PE1 and PE2 to PE4 as
                    secondary PWs (PW2 and PW4). When PE1 receives traffic from CE1, PE1 sends
                    the traffic to PE3 over PW1. After receiving the traffic, PE3 broadcasts the traffic to
                    PE2 and PE4 (because a VPLS network is a broadcast domain). As the PW from
                    PE2 to PE3 is a primary PW, PE2 can successfully receive the broadcast traffic and
                    forward the traffic to CE2. In this way, CE1 and CE2 can communicate with each
                    other.

                    Figure 6-33 VPWS accessing VPLS in dual-homed mode




VPWS accessing HVPLS
                    On the network shown in Figure 6-34, when UPEs do not support VPLS, you can
                    configure VPWS on UPEs and VPLS on SPEs for E2E service transmission. VPWS
                    PWs and VPLS PWs are established in static or LDP mode. For an SPE, the PW
                    established with a UPE is a spoke PW that does not follow the split horizon rule.
                    For a UPE, the PW established with an SPE is a VPWS LDP PW. Similarly, in a
                    common VPLS scenario, if a PE does not support VPLS, you can configure VPWS
                    accessing VPLS on the PE.




