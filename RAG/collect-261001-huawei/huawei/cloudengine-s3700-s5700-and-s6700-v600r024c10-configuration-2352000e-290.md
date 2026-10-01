---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-290
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [42768, 42958]
sha256: 72fc9b2cbf025aea033c35d6530f48e66a922fb555416371d05754b8c7017029
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                         isis enable 1
                        #
                        return
                    ●   ASBR_PE1
                        #
                        sysname ASBR_PE1
                        #
                        vlan batch 20 30
                        #
                        mpls lsr-id 2.2.2.2
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi a1 static
                         pwsignal ldp
                          vsi-id 2
                          peer 1.1.1.1
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 10.0000.0000.0002.00
                        #
                        interface Vlanif20
                         ip address 1.1.1.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif30
                         l2 binding vsi a1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                         isis enable 1
                        #
                        return
                    ●   ASBR_PE2
                        #
                        sysname ASBR_PE2
                        #
                        vlan batch 30 40
                        #
                        mpls lsr-id 3.3.3.3
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi a1 static
                         pwsignal ldp
                          vsi-id 3
                          peer 4.4.4.4
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 10.0000.0000.0003.00


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   687
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        #
                        interface Vlanif30
                         l2 binding vsi a1
                        #
                        interface Vlanif40
                         ip address 3.1.1.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                         isis enable 1
                        #
                        return

                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 40 50
                        #
                        mpls lsr-id 4.4.4.4
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi a1 static
                         pwsignal ldp
                          vsi-id 3
                          peer 3.3.3.3
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 10.0000.0000.0004.00
                        #
                        interface Vlanif40
                         ip address 3.1.1.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif50
                         l2 binding vsi a1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                         isis enable 1
                        #
                        return

                    ●   CE2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   688
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration

                        #
                        sysname CE2
                        #
                        vlan 50
                        #
                        interface Vlanif50
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        return


6.10.6 Example for Configuring Inter-AS BGP VPLS Option A
Networking Requirements
                    On the enterprise network shown in Figure 6-25, Site1 connects to the VPLS
                    domain of AS100 by connecting CE1 to PE1, and Site2 connects to the VPLS
                    domain of AS200 by connecting CE2 to PE2. The network environments of the
                    branch sites are unstable. AS100 and AS200 communicate with each other
                    through ASBR_PE1 and ASBR_PE2. IS-IS is used as the IGP on the MPLS backbone
                    network in an AS. Users at Site1 and Site2 need to communicate at Layer 2, and
                    user information needs to be retained in Layer 2 packets when the packets are
                    transmitted over the backbone network.

                    Figure 6-25 Network diagram of configuring inter-AS BGP VPLS Option A
                         NOTE

                        In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    689
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


