---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-286
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [42056, 42266]
sha256: 0b309db09a698ef138df9686147c9755dd62c5a569097856b6bde0ae545e0e1f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   SPE
                        #
                        sysname SPE
                        #
                        vlan batch 30 40
                        #
                        mpls lsr-id 2.2.2.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi v123 static
                         pwsignal ldp
                          vsi-id 123
                          peer 3.3.3.9
                          peer 1.1.1.9 upe


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   675
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        #
                        mpls ldp
                        #
                        interface Vlanif30
                         ip address 192.0.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif40
                         ip address 198.51.100.1 255.255.255.0
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
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 192.0.2.0 0.0.0.255
                          network 198.51.100.0 0.0.0.255
                        #
                        return
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 40 50
                        #
                        mpls lsr-id 3.3.3.9
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
                        interface Vlanif40
                         ip address 198.51.100.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif50
                         l2 binding vsi v123
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
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   676
VPN Configuration
VPN Configuration                                                                6 VPLS Configuration

                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 198.51.100.0 0.0.0.255
                        #
                        return

                    ●   CE1
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

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan 20
                        #
                        interface Vlanif20
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        return

                    ●   CE3
                        #
                        sysname CE3
                        #
                        vlan 50
                        #
                        interface Vlanif50
                         ip address 10.1.1.3 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        return



6.10 Configuring Inter-AS VPLS

6.10.1 Understanding Inter-AS VPLS
Definition
                    Inter-AS VPLS refers to VPLS applications across ASs. The implementation mode is
                    Option A.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      677
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration


                         NOTE

                        The fundamentals and implementation of PW establishment on PEs in inter-AS VPLS are
                        similar to those in inter-AS L2VPN, differing only in the learning and forwarding functions
                        of VSIs.


Purpose
                    Inter-AS VPLS allows a VPLS network to span multiple ASs on an MPLS backbone
                    network so that VPLS users in different ASs can communicate with each other.


Implementation of Inter-AS VPLS Option A
                    Inter-AS VPLS Option A can be used when there are a small number of inter-AS
                    VCs. ASBRs are required to support VSIs and be able to manage VPLS label blocks.
                    In addition, ASBRs need to reserve dedicated interfaces for each inter-AS VPLS.
                    This solution poses high requirements on ASBRs, but does not need ASBRs to have
                    inter-AS configurations.

                    On the network shown in Figure 6-23, the implementation of inter-AS VPLS
                    Option A is as follows:

                    ●   An IGP is configured on the backbone network to allow communication
                        between devices in the same AS.
                    ●   MPLS is enabled on the backbone network and a dynamic LSP is established
                        between the PE and ASBR.
                    ●   An IBGP peer relationship and a VPN LSP are established between the PE and
                        ASBR in the same AS.
                    ●   The PE and ASBR in the same AS have VSIs configured, and their AC
                        interfaces are bound to these VSIs. Each ASBR regards its peer ASBR as a local
                        CE.

                    Figure 6-23 Network diagram of inter-AS VPLS Option A




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       678
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


