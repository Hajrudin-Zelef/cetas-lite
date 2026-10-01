---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-208
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [30512, 30686]
sha256: 8d259817f854026cc39882fc752cb96a64d3ebdc178ab44882008ce5f67935dc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 2.2.2.9
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         ip address 10.2.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif20
                         ip address 10.1.1.2 255.255.255.0
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
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.2.2.0 0.0.0.255
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 1.1.1.9
                         remote-ip 1.1.1.9
                        #
                        interface Vlanif10
                         ip address 10.2.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif20


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   488
VPN Configuration
VPN Configuration                                                                                   5 VPWS Configuration

                         mpls static-l2vc destination 1.1.1.9 transmit-vpn-label 200 receive-vpn-label 100
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.2.2.0 0.0.0.255
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 20
                        #
                        interface Vlanif20
                         ip address 10.10.1.2 255.255.255.0
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        return



5.9 Configuring Inter-AS VPWS

5.9.1 Understanding Inter-AS VPWS

Definition
                    In real-world situations, VPN users may need to communicate across multiple
                    autonomous systems (ASs). This requires an interworking model that is different
                    from the MPLS VPN model. This interworking model features a VPN that spans
                    multiple ASs, also known as an inter-AS VPN.

                    Inter-AS VPWS depends on the implementation mode. The CCC mode uses a
                    single layer of label, meaning an inter-AS VPWS connection can be established
                    over a static LSP between autonomous system boundary routers (ASBRs). The
                    following describes the implementation of an inter-AS L2VPN in comparison with
                    the three methods of implementing an inter-AS L3VPN.

                    The SVC or LDP mode can implement inter-AS VPWS using Option A (VRF-to-
                    VRF). In inter-AS L2VPN networking, the link type between ASBRs must be the
                    same as the VC link type. The disadvantage of inter-AS VPWS Option A is that
                    each ASBR must reserve a sub-interface for each inter-AS VC. Option A can be
                    used when there are a small number of inter-AS VCs. Compared with L3VPN,
                    inter-AS L2VPN Option A consumes more resources and involves more
                    configurations. Therefore, inter-AS L2VPN Option A is not recommended.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         489
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


Purpose
                    With the increasing popularity of the MPLS VPN solution, the scale and scope of
                    served users are on a constant rise. As an enterprise deploys more sites, sites in
                    different geographical locations often connect to different ISP networks. This
                    causes the inter-AS issue where a VPN spans the metro or backbone networks of
                    different carriers in different ASs. This is where inter-AS L2VPN comes in.

                    Figure 5-28 Origin of inter-AS technology




                    On the L2VPN shown in Figure 5-28, some users belong to AS1, and some to AS2.
                    If MPLS forwarding is not implemented, L2VPN users in different ASs cannot
                    communicate with each other. The figure above shows a scenario where L2VPN
                    users in two ASs need to communicate. In real-world situations, L2VPN users in
                    more ASs may need to communicate.

Inter-AS Option A
                    In this mode, the ASBRs of the two ASs directly connect to each other and
                    function as PEs in their respective ASs. The two ASBRs regard each other as a CE.

                    Figure 5-29 Inter-AS Option A networking




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           490
VPN Configuration
VPN Configuration                                                                5 VPWS Configuration


