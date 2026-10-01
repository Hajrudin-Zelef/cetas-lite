---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-326
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [48625, 48811]
sha256: 906db56eb0259bb916c5256418755f0f842fcecfbe7e9d379bbbb03a29bac75d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   781
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration

                        #
                        erps ring 1
                         control-vlan 100
                         protected-instance 1
                         version v2
                         sub-ring
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         undo port trunk allow-pass vlan 1
                         port trunk allow-pass vlan 10 100
                         stp disable
                         erps ring 1
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         undo port trunk allow-pass vlan 1
                         port trunk allow-pass vlan 10 100
                         stp disable
                         erps ring 1 rpl owner
                        #
                        return


6.18.5 Example for Configuring ERPS over VPLS for CE Access
Through VLANIF Interfaces
Networking Requirements
                    Figure 6-46 shows a VPLS network where CEs are connected to PEs. However, the
                    problem with this networking is that PE3 receives two copies of traffic from the
                    remote CEs. To solve this problem, enable ERPS on PE1, CE1, CE2, and PE2, and
                    configure interface 2 of CE2 as an RPL owner port to block traffic from CE1. In this
                    way, traffic from CE1 is directly transmitted to PE3 through PE1 without passing
                    through CE2, thereby preventing duplicate traffic or loops.

                    Figure 6-46 Network diagram of configuring ERPS over VPLS for CE access
                    through VLANIF interfaces
                         NOTE

                        In this example, interface 1 and interface 2 represent 10GE 1/0/1 and 10GE 1/0/2,
                        respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                 782
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    Table 6-12 lists the IP addresses of interfaces on PE1, PE2, and PE3.

                    Table 6-12 Data plan

                     Device                      Interface                   IP Address

                     PE1                         10GE1/0/1                   -

                                                 10GE1/0/2                   10.1.1.1/24

                                                 Loopback1                   1.1.1.1/32

                     PE2                         10GE1/0/1                   -

                                                 10GE1/0/2                   10.2.1.1/24

                                                 Loopback1                   2.2.2.2/32

                     PE3                         10GE1/0/1                   10.1.1.2/24

                                                 10GE1/0/2                   10.2.1.2/24

                                                 10GE1/0/3                   -

                                                 Loopback1                   3.3.3.3/32

                     CE1                         10GE1/0/1                   -

                                                 10GE1/0/2                   -

                     CE2                         10GE1/0/1                   -

                                                 10GE1/0/2                   -




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure PEs on the VPLS backbone network to run an IGP so that they can
                         communicate with each other.
                    2.   Configure basic MPLS functions and establish LDP LSPs on the VPLS backbone
                         network.
                    3.   Establish VPLS connections between PEs and bind VLANIF interfaces to a VSI.
                    4.   Configure ERPS, including:
                         –    Enable ERPS on PE1, CE1, CE2, and PE2.
                         –    Configure 10GE 1/0/2 of CE2 as an RPL owner port.

Data Plan
                    To complete the configuration, you need the following data:
                    ●    Data required for configuring OSPF: IP address of each interface, OSPF
                         process ID, and OSPF area ID
                    ●    MPLS LSR IDs (used as MPLS peer addresses)

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          783
VPN Configuration
VPN Configuration                                                              6 VPLS Configuration


                    ●    VSI name and VSI ID
                    ●    VLANIF interfaces bound to a VSI
                    ●    ERPS ring ID, control VLAN ID, and RPL owner port

Procedure
         Step 1 Configure interface IP addresses and an IGP on the VPLS backbone network so
                that PEs can communicate with each other. In this example, OSPF is used as the
                IGP. When configuring OSPF, configure PE1, PE2, and PE3 to advertise their 32-bit
                IP addresses (used as LSR IDs) of loopback interfaces.

                    For detailed configurations, see Configuration Scripts.

         Step 2 Configure basic MPLS functions on the MPLS backbone network and establish
                dynamic LDP LSPs between PEs.

                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.1
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] mpls
                    [PE1-10GE1/0/2] mpls ldp
                    [PE1-10GE1/0/2] quit

                    # Configure PE2.
                    [PE2] mpls lsr-id 2.2.2.2
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface 10ge 1/0/2
                    [PE2-10GE1/0/2] mpls
                    [PE2-10GE1/0/2] mpls ldp
                    [PE2-10GE1/0/2] quit

                    # Configure PE3.
                    [PE3] mpls lsr-id 3.3.3.3
                    [PE3] mpls
                    [PE3-mpls] quit
                    [PE3] mpls ldp
                    [PE3-mpls-ldp] quit
                    [PE3] interface 10ge 1/0/1
                    [PE3-10GE1/0/1] mpls
                    [PE3-10GE1/0/1] mpls ldp
                    [PE3-10GE1/0/1] quit
                    [PE3] interface 10ge 1/0/2
                    [PE3-10GE1/0/2] mpls
                    [PE3-10GE1/0/2] mpls ldp
                    [PE3-10GE1/0/2] quit

         Step 3 Enable MPLS L2VPN on PEs.

                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit

                    # Configure PE2.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   784
VPN Configuration
VPN Configuration                                                               6 VPLS Configuration

                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit

                    # Configure PE3.
                    [PE3] mpls l2vpn
                    [PE3-l2vpn] quit

