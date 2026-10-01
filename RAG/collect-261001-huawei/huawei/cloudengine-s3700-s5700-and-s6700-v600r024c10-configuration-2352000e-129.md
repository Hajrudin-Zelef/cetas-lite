---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-129
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [18253, 18399]
sha256: 037b948360471e4d3b4821849d263050f43732845e26c51ce80c5d6320235781
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                              After the import-route bgp command is run in the RIPng view, the PE
                              can import the VPN-IPv6 routes learned from the remote PE into the
                              RIPng routing table and advertise them to the attached CE.
                        d.    Return to the system view.
                              quit

                        e.    Enter the view of the interface connected to the CE.
                              interface interface-type interface-number

                        f.    Enable RIPng on the interface.
                              ripng process-id enable

                                      NOTE

                                     Before running this command, ensure that IPv6 has been enabled in the interface
                                     view.
                        g.    Return to the system view.
                              quit

                        h.    Enter the BGP view.
                              bgp as-number

                        i.    Enter the BGP-VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name

                        j.    Import RIPng routes into the routing table of the BGP-VPN instance IPv6
                              address family.
                              import-route ripng process-id [ med med | route-policy route-policy-name ] *

                              After the import-route ripng command is run in the BGP-IPv6 VPN
                              instance IPv6 address family view, the PE imports the IPv6 routes learned
                              from the attached CE into the BGP routing table and advertises VPN-IPv6
                              routes to the remote PE.

                                      NOTE

                                     If a RIPng multi-instance process is deleted, RIPng will be disabled on all the
                                     interfaces in the process.
                                     Deleting a VPN instance or disabling a VPN instance IPv6 address family will
                                     delete all the RIPng processes bound to the VPN instance or VPN instance IPv6
                                     address family.
                        k.    (Optional) Run either of the following commands to configure the device
                              to advertise specific routes in a BGP-VPN routing table to a BGP-VPNv6
                              routing table:

                              ▪      Configure the device to advertise only valid routes in a BGP-VPN
                                     routing table to a BGP-VPNv6 routing table.
                                     advertise valid-routes

                                     By default, the device advertises all routes in the BGP-VPN routing
                                     table to the BGP-VPNv6 routing table. The advertise valid-routes
                                     command allows the device to advertise only valid routes to the
                                     BGP-VPNv6 routing table.
                    ●   Configure RIPng on the CE. The configuration details are not provided here.

                    ----End


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                         290
VPN Configuration
VPN Configuration                                                                    4 IPv6 L3VPN Configuration


4.6.9 Verifying the Configuration

Procedure
                    ●   Run the display ip vpn-instance vpn-instance-name command to check brief
                        information about a specified VPN instance.
                    ●   Run the display ip vpn-instance verbose vpn-instance-name command to
                        check detailed information about a specified VPN instance.
                    ●   Run the display ip vpn-instance import-vt ivt-value command to check
                        information about VPN instances with the specified import VPN target.

                        This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                        S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-
                        V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.
                    ●   Run the display ip vpn-instance [ vpn-instance-name ] interface command
                        to check information about the interface bound to a specified VPN instance.
                    ●   Run the display ipv6 routing-table vpn-instance vpn-instance-name
                        command on a PE to check route information in the VPN instance IPv6
                        address family.
                    ●   Run the display ipv6 routing-table command on a CE to check route
                        information.

                    ----End

4.6.10 Example for Configuring Basic IPv6 L3VPN over MPLS

Networking Requirements
                    IPv6 L3VPN applies to scenarios where different user sites communicate through
                    the public network without letting the public network detect their internal routing
                    information. IPv6 L3VPN can isolate VPN services from each other by allowing
                    intra-VPN access and prohibiting inter-VPN access.

                    On the network shown in Figure 4-2, CE1 and CE3 belong to VPNA, and CE2 and
                    CE4 belong to VPNB. It is required that IPv6 L3VPN be configured to allow the
                    sites in VPNA and those in VPNB to communicate with each other through an
                    MPLS backbone network instead of directly communicating with each other. It is
                    also required that different methods be used to exchange routes between PEs and
                    CEs:
                    ●   BGP4+ between PE1 and CE1, and between PE2 and CE4
                    ●   IPv6 static route between PE1 and CE2
                    ●   OSPFv3 between PE2 and CE3


                    Figure 4-2 Configuring basic IPv6 L3VPN
                         NOTE

                        In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200,
                        and VLANIF 300, respectively.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      291
VPN Configuration
VPN Configuration                                                           4 IPv6 L3VPN Configuration




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure IGP on the IPv4 backbone network for PEs to communicate.
                    2.   Configure MPLS and MPLS LDP on each PE and the P to establish LDP LSPs
                         between PEs.
                    3.   Configure MP-IBGP on PE1 and PE2 to enable PEs to exchange IPv6 VPN
                         routing information through BGP.
                    4.   Configure an IPv6-address-family-enabled VPN instance on both PE1 and PE2,
                         and bind the interfaces connected to CEs to the corresponding VPN instances.
                    5.   Configure IPv6 routing protocols between PEs and CEs for them to exchange
                         IPv6 routing information.

Data Plan
                    To complete the configuration, you need the following data:

                    ●    AS numbers of PEs and CEs
                    ●    VPN instance names
                    ●    Attributes of the VPN instance IPv6 address family, such as the RD and VPN
                         targets

Procedure
         Step 1 Configure IPv4 or IPv6 addresses for device interfaces.

                    # Configure an IPv6 address for the interface on CE1.

