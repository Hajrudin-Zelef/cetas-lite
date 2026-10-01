---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-72
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [9707, 9861]
sha256: ce18a3c837d8177116d4c4442233eed19c82263390a3d941c4432419a983c619
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 2 Configure different filters to filter routes.
                    ●      Configure a VPN target filter using either of the following methods:
                           ip extcommunity-filter { basic-extcomm-filter-num | basic basic-extcomm-filter-name } { deny |
                           permit } { rt { as-number:nn | 4as-number:nn | ipv4-address:nn } } &<1-16>
                           ip extcommunity-filter { advanced-extcomm-filter-num | advanced advanced-extcomm-filter-
                           name } { deny | permit } regular-expression
                    ●      Configure an SoO filter using either of the following methods:
                           ip extcommunity-list soo basic basic-extcomm-filter-name [ index index-number ] { permit | deny }
                           { site-of-origin } &<1-16>
                           ip extcommunity-list soo advanced advanced-extcomm-filter-name [ index index-number ]
                           { permit | deny } regular-expression

                    ●      Configure an RD filter.
                           ip rd-filter rdfIndex [ index index-number ] matchMode rdStr &<1-10>

         Step 3 Configure a route-policy.
                    route-policy route-policy-name permit node node

         Step 4 Configure filtering conditions in the roue-policy.
                    ●      Configure a matching rule based on the VPN target filter for the route-policy.
                           if-match extcommunity-filter { { basic-extcomm-filter-num | adv-extcomm-filter-num } &<1-16> |
                           basic-extcomm-filter-name | advanced-extcomm-filter-name }
                    ●      Configure a matching rule based on the SoO filter for the route-policy.
                           if-match extcommunity-list soo extcomm-filter-name

                    ●      Configure a matching rule based on the RD filter for the route-policy.
                           if-match rd-filter rd-filter-number

         Step 5 Return to the system view.
                    quit

         Step 6 Enter the BGP view.
                    bgp as-number

         Step 7 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4 [ unicast ]


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                              154
VPN Configuration
VPN Configuration                                                                          3 IPv4 L3VPN Configuration


         Step 8 Apply the route-policy to control the import and export of VPNv4 routes.
                    peer ipv4-address route-policy route-policy-name { export | import }

                    ----End

3.12.6 Configuring Route Exchange Between the CE and PE

Context
                    BGP, IGP, or static routes (including the default routes) can be used between a CE
                    and a PE. Determine which one to use as required.

Procedure
         Step 1 Determine whether to use BGP, IGP, or static routes between a PE and a CE as
                required. For details, see 3.6 Configuring Basic IPv4 L3VPN over MPLS.

                    ----End

3.12.7 Verifying the Configuration

Prerequisites
                    All the configurations about inter-AS VPN Option B are complete.

Procedure
                    ●    Run the display bgp vpnv4 all peer command on PEs or ASBRs to check the
                         establishment of all BGP peer relationships.
                    ●    Run the display bgp vpnv4 all routing-table command on PEs or ASBRs to
                         check VPNv4 routes.
                    ●    Run the display ip routing-table vpn-instance vpn-instance-name command
                         on PEs to check VPN routing table information.
                    ----End

3.12.8 Example for Configuring IPv4 L3VPN over MPLS Inter-
AS Option B (Basic Networking)
Networking Requirements
                    On the network shown in Figure 3-30, CE1 and CE2 belong to the same VPN. CE1
                    connects to PE1 in AS 100, and CE2 connects to PE2 in AS 200. It is required that
                    an MP-EBGP peer relationship be established between the ASBRs to transmit
                    VPNv4 routes, implementing inter-AS VPN Option B.

                    Figure 3-30 Inter-AS VPN Option B (basic networking)
                          NOTE

                    In this example, interface 1 and interface 2 represent VLANIF 100 and VLANIF 200, respectively.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     155
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration




Precautions
                    Note the following during the configuration:

                    ●    Configure an MP-EBGP peer relationship between ASBR1 and ASBR2, and
                         disable the ASBRs from filtering received VPNv4 routes based on VPN targets.


Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure IGP on the backbone network for IP connectivity between the ASBR
                         and PE in the same AS, and establish an MPLS LDP LSP between the ASBR
                         and PE in the same AS.
                    2.   Set up EBGP peer relationships between PEs and CEs and MP-IBGP peer
                         relationships between PEs and ASBRs.
                    3.   Configure VPN instances on the PEs rather than ASBRs.
                    4.   Enable MPLS on the interfaces that connect the ASBRs, establish an MP-EBGP
                         peer relationship between the ASBRs, and configure the ASBRs not to filter
                         received VPNv4 routes based on VPN targets.


Procedure
         Step 1 Configure IGP on the MPLS backbone networks in AS 100 and AS 200 for
                communication between PEs on each backbone network. OSPF is used as IGP in
                this example. For detailed configurations, see Configuration Scripts.
                          NOTE

                         The 32-bit address of the loopback interface that functions as the LSR ID needs to be
                         advertised using OSPF.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      156
VPN Configuration
VPN Configuration                                                              3 IPv4 L3VPN Configuration


                    After completing the configuration, run the display ospf peer command on an
                    ASBR or PE. The command output shows that the OSPF neighbor relationship is in
                    the Full state, indicating that the OSPF neighbor relationship has been established
                    between the ASBR and PE in the same AS.
                    The ASBR and PE in the same AS can learn and ping the address of each other's
                    loopback interface.
         Step 2 Configure basic MPLS functions and MPLS LDP on the MPLS backbone networks in
                AS 100 and AS 200 respectively to establish MPLS LDP LSPs.
                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] mpls lsr-id 1.1.1.1
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface Vlanif 100
                    [PE1-Vlanif100] mpls
                    [PE1-Vlanif100] mpls ldp
                    [PE1-Vlanif100] quit

