---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-120
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [16928, 17105]
sha256: a59539f2d7b7e294ebcf6af0db11390c02aa132dbef328c4f7261f7cd43f43e5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         –    Run the following command to display brief information:
                              ping [ ip ] { [ -c count | -i { interface-name | interface-type interface-number } | -nexthop
                              nexthop-address | { -range [ min min-value | max max-value | step step-value ] * | -s
                              packetsize } | -t timeout | -m time | -a source-ip-address | -h ttl-value | -p pattern | { -tos tos-
                              value | -dscp dscp-value } | { -f | ignore-mtu } | -vpn-instance vpn-instance-name | -name |
                              -8021p 8021p-value | -brief ] * host [ ip-forwarding ] }

                    ●    Run the tracert [ -a source-ip-address | -f first-TTL | -m max-TTL | -p port | -q
                         nqueries | -vpn-instance vpn-instance-name | -w timeout ] * host command
                         to check the gateways through which IPv4 packets pass from the source to
                         the destination.

                    ----End

3.16.4 Checking Integrated Route Statistics of IPv4 VPN
Instances



Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                     267
VPN Configuration
VPN Configuration                                                          3 IPv4 L3VPN Configuration


Procedure
                    ●   Run the display ip routing-table { vpn-instance vpn-instance-name | all-
                        vpn-instance } statistics command to check integrated route statistics of a
                        specified IPv4 VPN instance or all IPv4 VPN instances.
                    ----End




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                          268
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration




                                    4          IPv6 L3VPN Configuration


                    4.1 Overview of IPv6 L3VPN
                    4.2 Understanding IPv6 L3VPN
                    4.3 Configuration Precautions for IPv6 L3VPN
                    4.4 Default Settings for IPv6 L3VPN
                    4.5 Configuring Mutual Access Between Local IPv6 L3VPNs
                    4.6 Configuring Basic IPv6 L3VPN over MPLS Functions
                    4.7 Configuring Route Import Between Instances
                    4.8 Setting a Router ID for a BGP VPN Instance IPv6 Address Family
                    4.9 Configuring IPv6 MCE
                    4.10 Configuring IPv6 L3VPN FRR
                    4.11 Configuring 6VPE
                    4.12 Configuring IPv6 L3VPN over MPLS Inter-AS Option A
                    4.13 Configuring IPv6 L3VPN over MPLS Inter-AS Option B
                    4.14 Configuring IPv6 L3VPN over MPLS Hub-Spoke
                    4.15 Maintaining IPv6 L3VPN


4.1 Overview of IPv6 L3VPN
                         NOTE

                        Only the S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2, S5755E-H,
                        S5755-H, and S5732-H-V2 series can use MPLS to forward VPN packets on backbone
                        networks of SPs.


Definition
                    VPN technology is used to establish private networks over public networks.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 269
VPN Configuration
VPN Configuration                                                            4 IPv6 L3VPN Configuration


                    Depending on whether CEs reside on an IPv4 or IPv6 network, L3VPN can be
                    classified as IPv4 L3VPN or IPv6 L3VPN. For more information about IPv6 L3VPN,
                    see 3.1 Overview of IPv4 L3VPN.


4.2 Understanding IPv6 L3VPN
                    IPv6 L3VPN is fundamentally similar to IPv4 L3VPN. For details, see 3.2
                    Understanding IPv4 L3VPN.


4.3 Configuration Precautions for IPv6 L3VPN

4.4 Default Settings for IPv6 L3VPN
                    Table 4-1 lists the default settings for IPv6 L3VPN.


                    Table 4-1 Default settings for IPv6 L3VPN

                     Parameter                      Default Setting

                     Maximum number of              No limit
                     prefixes supported by the
                     VPN instance IPv6 address
                     family




4.5 Configuring Mutual Access Between Local IPv6
L3VPNs

4.5.1 Configuring an IPv6 VPN Instance on a PE

Prerequisites
                    Before configuring an IPv6 VPN instance, you have completed the following task:

                    ●   Configure link layer protocol parameters for interfaces to ensure that these
                        interfaces work properly.


Context
                    A VPN instance is also called a VRF or a per-site forwarding table.

                    VPN instances are used to isolate VPN routes from public network routes. Routes
                    of different VPN instances are isolated from one another.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         270
VPN Configuration
VPN Configuration                                                                              4 IPv6 L3VPN Configuration


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VPN instance and enter the VPN instance view.
                    ip vpn-instance vpn-instance-name

                          NOTE

                         The name of a VPN instance is case-sensitive. For example, vpn1 and VPN1 are considered
                         different VPN instances.

         Step 3 (Optional) Configure a description for the VPN instance.
                    description description-information

         Step 4 Enable the VPN instance IPv6 address family, and enter the VPN instance IPv6
                address family view.
                    ipv6-family

                    Configurations in a VPN instance can be performed only after an address family is
                    enabled for the VPN instance based on the advertised route and type of forwarded
                    data.

         Step 5 Configure an RD for the VPN instance IPv6 address family.
                    route-distinguisher route-distinguisher

                    A VPN instance IPv6 address family takes effect only after having an RD
                    configured. The RDs of different VPN instance IPv6 address families on a PE must
                    be different.

                          NOTE

                         If you perform this step in the view of a newly created VPN instance, the VPN instance IPv6
                         address family is automatically enabled and the VPN instance IPv6 address family view is
                         automatically displayed.
                         The S3710-H series products do not support dynamic routing protocols. Therefore, this
                         command does not take effect after being delivered.

         Step 6 Configure VPN targets for the VPN instance IPv6 address family.
                    vpn-target vpn-target &<1-8> [ both | export-extcommunity | import-extcommunity ]

                    VPN targets are a type of BGP extended community attribute used to control the
                    import and export of VPN routes. You can configure a maximum of eight import
                    VPN targets and eight export VPN targets each time the vpn-target command is
                    run. Run this command multiple times if you want to configure more VPN targets
                    in the VPN instance IPv6 address family.

