---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-66
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [8728, 8879]
sha256: 46c31c8f102ce233bda43b729bc397f377ab28ff84c3fedda95aa0f8e1c0b5da
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    If PEs access only a few VPNs and only a small number of VPN routes exist, inter-
                    AS VPN Option A is recommended. In Inter-AS VPN Option A, ASBRs are required
                    to support VPN instances so that they can manage VPN routes. In addition, ASBRs
                    must reserve dedicated interfaces (sub-interfaces, physical interfaces, or bundled
                    logical interfaces) for each inter-AS VPN. This solution poses high requirements on
                    ASBRs, but does not need ASBRs to have special inter-AS configurations.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          138
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration


Procedure
         Step 1 Configure an IPv4 VPN instance for each AS on each ASBR (an ASBR is viewed as
                a PE in this case).
                         NOTE

                        In inter-AS VPN Option A, ensure that the VPN targets of the VPN instance on the ASBR
                        match those of the VPN instance on the PE in the same AS. The VPN targets of the VPN
                        instances on the PEs in different ASs do not need to match each other.

         Step 2 On each ASBR, bind the interface connected to the remote ASBR to a VPN
                instance. For details, see Binding an Interface to an IPv4 VPN Instance.
         Step 3 Configure a routing protocol between ASBRs. This configuration is similar to
                configuring a routing protocol between a PE and a CE in an IPv4 VPN instance. For
                configuration details, see 3.6 Configuring Basic IPv4 L3VPN over MPLS.

                    ----End

Verifying the Configuration
                    ●    Run the display bgp vpnv4 all peer command on a PE or ASBR to check the
                         status of the BGP-VPNv4 peer relationship between the PE and ASBR in the
                         same AS.
                    ●    Run the display bgp vpnv4 all routing-table command on PEs or ASBRs to
                         check VPNv4 routes.
                    ●    Run the display ip routing-table vpn-instance vpn-instance-name [ ip-
                         address ] verbose command on PEs or ASBRs to check whether their VPN
                         routing tables contain all related VPN routes.

3.11.3 Example for Configuring IPv4 L3VPN over MPLS Inter-
AS Option A
Networking Requirements
                    On the network shown in Figure 3-26, CE1 and CE2 belong to the same VPN. CE1
                    connects to PE1 in AS 100, and CE2 connects to PE2 in AS 200.
                    Inter-AS IPv4 L3VPN is implemented in Option A mode, in which VPN routes are
                    managed using the VRF-to-VRF method.

                    Figure 3-26 Inter-AS VPN networking
                         NOTE

                    In this example, interface 1 and interface 2 represent VLANIF 100 and VLANIF 200, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     139
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Set up EBGP peer relationships between PEs and CEs and MP-IBGP peer
                         relationships between PEs and ASBRs.
                    2.   Create a VPN instance on each ASBR, bind the interface connecting each
                         ASBR to the peer ASBR to the VPN instance, and establish an EBGP peer
                         relationship between ASBRs.

Procedure
         Step 1 Configure IGP on the MPLS backbone networks in AS 100 and AS 200 for IP
                connectivity between the ASBR and PE on each MPLS backbone network.
                    OSPF is used as IGP in this example. For detailed configurations, see Configuration
                    Scripts.

                          NOTE

                         The 32-bit address of the loopback interface that functions as the LSR ID needs to be
                         advertised using OSPF.

                    After completing the configuration, run the display ospf peer command on an
                    ASBR or PE. The command output shows that the OSPF neighbor relationship is in
                    the Full state, indicating that the OSPF neighbor relationship has been established
                    between the ASBR and PE in the same AS.
                    The ASBR and PE in the same AS can learn and ping the address of each other's
                    loopback interface.
         Step 2 Configure basic MPLS functions and MPLS LDP on the MPLS backbone networks in
                AS 100 and AS 200 respectively to establish MPLS LDP LSPs.
                    # Configure PE1.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      140
VPN Configuration
VPN Configuration                                                                                3 IPv4 L3VPN Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface Vlanif 100
                    [PE1-Vlanif100] mpls
                    [PE1-Vlanif100] mpls ldp
                    [PE1-Vlanif100] quit

                    # Configure ASBR1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR1
                    [ASBR1] mpls lsr-id 2.2.2.9
                    [ASBR1] mpls
                    [ASBR1-mpls] quit
                    [ASBR1] mpls ldp
                    [ASBR1-mpls-ldp] quit
                    [ASBR1] interface Vlanif 100
                    [ASBR1-Vlanif100] mpls
                    [ASBR1-Vlanif100] mpls ldp
                    [ASBR1-Vlanif100] quit

                    # Configure ASBR2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR2
                    [ASBR2] mpls lsr-id 3.3.3.9
                    [ASBR2] mpls
                    [ASBR2-mpls] quit
                    [ASBR2] mpls ldp
                    [ASBR2-mpls-ldp] quit
                    [ASBR2] interface Vlanif 100
                    [ASBR2-Vlanif100] mpls
                    [ASBR2-Vlanif100] mpls ldp
                    [ASBR2-Vlanif100] quit

                    # Configure PE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] mpls lsr-id 4.4.4.9
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface Vlanif 100
                    [PE2-Vlanif100] mpls
                    [PE2-Vlanif100] mpls ldp
                    [PE2-Vlanif100] quit

