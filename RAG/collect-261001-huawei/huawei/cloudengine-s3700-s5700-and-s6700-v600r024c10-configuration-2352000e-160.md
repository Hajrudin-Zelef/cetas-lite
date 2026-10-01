---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-160
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [22980, 23126]
sha256: 3f2605404eee308e9c29946f93aa8db8b4750668cd8a8898524a13067a7b50cd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          362
VPN Configuration
VPN Configuration                                                                       4 IPv6 L3VPN Configuration


                        k.    Apply an export routing policy.
                              peer ipv4-address route-policy route-policy-name export

                    ----End

4.11.6 Verifying the Configuration

Procedure
                    ●   Run the display ip vpn-instance vpn-instance-name command to check brief
                        information about a specified VPN instance.
                    ●   Run the display ip vpn-instance verbose vpn-instance-name command to
                        check detailed information about a specified VPN instance, including
                        information about the VPN instance IPv4 address family and IPv6 address
                        family.
                    ●   Run the display ip vpn-instance import-vt ivt-value command to check
                        information about VPN instances with the specified import VPN target.

                        This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                        S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-
                        V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.
                    ●   Run the display ip vpn-instance [ vpn-instance-name ] interface command
                        to check information about the interface bound to a specified VPN instance.
                    ●   Run the display bgp vpnv6 vpn-instance vpn-instance-name peer ipv4-
                        address verbose command on a PE to check the BGP IPv4 peer relationship
                        established with a CE in the VPN instance IPv6 address family view.

                    ----End


4.12 Configuring IPv6 L3VPN over MPLS Inter-AS
Option A

4.12.1 Understand IPv6 L3VPN over MPLS Inter-AS Option A
                    The fundamentals of inter-AS IPv6 VPN Option A are similar to those of inter-AS
                    VPN Option A. For details, see 3.11.1 Understanding IPv4 L3VPN over MPLS
                    Inter-AS Option A.

4.12.2 Configuring IPv6 L3VPN over MPLS Inter-AS Option A

Prerequisites
                    Before configuring inter-AS IPv6 VPN Option A, you have completed the following
                    tasks:

                    ●   Configure IGP for the MPLS backbone network in each AS to ensure IP
                        connectivity for the backbone network in each AS.
                    ●   Enable MPLS on the PEs and ASBRs.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   363
VPN Configuration
VPN Configuration                                                                   4 IPv6 L3VPN Configuration


                    ●   Establish a tunnel (LSP) between the PE and ASBR in the same AS.
                    ●   Enable IPv6 on interfaces to be configured with IPv6 addresses and configure
                        IPv6 addresses for them.

Context
                    If the MPLS backbone network transmitting IPv6 VPN routes spans multiple ASs,
                    the inter-AS IPv6 VPN solution is required.
                    If PEs access only a few VPNs and only a small number of VPN-IPv6 routes exist,
                    inter-AS IPv6 VPN Option A is recommended. In inter-AS IPv6 VPN Option A,
                    ASBRs are required to support VPN instances so that they can manage VPN-IPv6
                    routes. In addition, ASBRs must provide dedicated interfaces (either sub-interfaces
                    or physical interfaces) for each inter-AS IPv6 VPN. This solution poses high
                    requirements on ASBRs, but does not need ASBRs to have inter-AS configurations.
                    In inter-AS IPv6 VPN Option A, an ASBR views the peer ASBR as a CE and uses
                    EBGP+ to advertise IPv6 routes to the peer ASBR.

Procedure
         Step 1 Configure an IPv6 VPN instance for each AS.
                         NOTE

                        In inter-AS IPv6 VPN Option A, for the same IPv6 VPN, the VPN targets of the IPv6-address-
                        family-enabled VPN instances of the ASBR and PE in the same AS must match. This is not
                        required for the PEs in different ASs.

         Step 2 Configure each ASBR to regard its peer ASBR as a local CE.
         Step 3 Configure an IPv6-address-family-enabled VPN instance on each PE and ASBR. For
                details, see Configuring an IPv6 VPN Instance. The VPN instance on the PE is
                used for access by CEs, and the VPN instance on the ASBR is used for access by
                the peer ASBR.

                    ----End

Verifying the Configuration
                    ●   Run the display bgp vpnv6 all peer command on a PE or ASBR to check the
                        status of the BGP VPNv6 peer relationship between the PE and ASBR in the
                        same AS.
                    ●   Run the display bgp vpnv6 all routing-table command on a PE or ASBR to
                        check VPNv6 routes.
                    ●   Run the display ipv6 routing-table vpn-instance [ vpn-instance-name ]
                        command on a PE or ASBR to check the routing table of the VPN instance
                        IPv6 address family.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    364
VPN Configuration
VPN Configuration                                                                   4 IPv6 L3VPN Configuration


4.12.3 Example for Configuring IPv6 L3VPN over MPLS Inter-
AS Option A
Networking Requirements
                    Inter-AS IPv6 VPN Option A applies to scenarios where the carrier backbone
                    network needs to provide inter-AS IPv6 VPN services for customers.
                    Inter-AS IPv6 VPN Option A is easy to configure. You only need to configure an
                    IPv6-address-family-enabled VPN instance on each ASBR and configure each ASBR
                    to regard the peer ASBR as its local CE. If the services of many VPNs need to be
                    transmitted across ASs, the requirements for ASBR performance are high.
                    On the network shown in Figure 4-7, CE1 and CE2 belong to the same VPN. CE1
                    accesses AS100 through PE1, and CE2 accesses AS200 through PE2.
                    It is required that Option A be configured to implement inter-AS IPv6 VPN, so that
                    CE1 and CE2 can communicate.

                    Figure 4-7 Configuring inter-AS IPv6 VPN Option A
                          NOTE

                         In this example, interface 1 and interface 2 represent VLANIF 100 and VLANIF 200,
                         respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Set up EBGP peer relationships between PEs and CEs and MP-IBGP peer
                         relationships between PEs and ASBRs.
                    2.   On each ASBR, configure an IPv6-address-family-enabled VPN instance and
                         bind the interface connecting to the peer ASBR to the VPN instance. Then
                         establish an EBGP peer relationship between the ASBRs.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  365
VPN Configuration
VPN Configuration                                                                    4 IPv6 L3VPN Configuration


