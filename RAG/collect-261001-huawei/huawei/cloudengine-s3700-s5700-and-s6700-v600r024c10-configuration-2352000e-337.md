---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-337
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [50361, 50514]
sha256: 2d519c031bb2561ffd3b56472d1d46f98bfc5436f9b3c539543e785bb84d0748
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Prerequisites
                    A tunnel policy has been configured.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            810
VPN Configuration
VPN Configuration                                                                7 Tunnel Management Configuration


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a tunnel selector and enter its view.
                    tunnel-selector name { permit | deny } node node

         Step 3 Apply a tunnel policy to routes.
                    apply tunnel-policy tunnel-policy-name

                    By default, no tunnel policy is applied.

         Step 4 (Optional) Configure an if-match clause.

                    By default, all routes match the filtering rules.

                    Run the following commands as needed to configure one or more filtering rules
                    for the current tunnel selector node.

                    ●    Configure a filtering rule based on the RD attribute of routes.
                         if-match rd-filter rd-filter-number

                    ●    Configure a filtering rule based on the IPv4 next hop attribute of routes.
                         if-match ip next-hop { acl { acl-number | acl-name } | ip-prefix ip-prefix-name }

                    ●    Configure a filtering rule based on the IPv6 next hop attribute of routes.
                         if-match ipv6 next-hop prefix-list ipv6-prefix-name

                    ●    Configure a filtering rule based on the IP address prefix of routes.
                         if-match ip-prefix ip-prefix-name

                    ●    Configure a matching rule based on a community attribute of routes.
                         if-match community-filter { basIndex [ whole-match ] | AdvIndex [sort-match ] } &<1-16>
                         if-match community-filter cfName [ whole-match | sort-match ]

                    ----End

7.5.3 Applying a Tunnel Selector

Context
                    The system can use a tunnel selector to have routes recurse to a proper tunnel
                    only after the tunnel selector is applied to the routes on the PE or ASBR.

                    A tunnel selector applies to the following types of routes:
                    ●    VPNv4 routes: A tunnel selector can be applied to the BGP-VPNv4 address
                         family so that the ASBR in inter-AS VPN Option B networking can apply
                         tunnel policies to VPNv4 routes and recurse these routes to appropriate
                         tunnels.
                    ●    VPNv6 routes: A tunnel selector can be applied to the BGP-VPNv6 address
                         family so that the ASBR in inter-AS VPN Option B networking can apply
                         tunnel policies to VPNv6 routes and recurse these routes to appropriate
                         tunnels.

                    Perform the following steps on the PE or ASBR where a tunnel selector needs to
                    be applied:

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       811
VPN Configuration
VPN Configuration                                                    7 Tunnel Management Configuration


Procedure
                    ●   Apply a tunnel selector to VPNv4 routes.
                        a.    Enter the system view.
                              system-view
                        b.    Enter the BGP view.
                              bgp as-number
                        c.    Enter the BGP-VPNv4 address family view.
                              ipv4-family vpnv4
                        d.    Apply a tunnel selector to VPNv4 routes.
                              tunnel-selector tunnel-selector-name

                              After the tunnel selector is applied to VPNv4 routes, the VPNv4 routes
                              that match the if-match clause recurse to tunnels according to the
                              tunnel policy specified in the apply clause. The VPNv4 routes that do not
                              match the if-match clause recurse to LSPs by default.
                    ●   Apply a tunnel selector to VPNv6 routes.
                        a.    Enter the system view.
                              system-view
                        b.    Enter the BGP view.
                              bgp as-number
                        c.    Enter the BGP-VPNv6 address family view.
                              ipv6-family vpnv6
                        d.    Apply a tunnel selector to VPNv6 routes.
                              tunnel-selector tunnel-selector-name

                              After the tunnel selector is applied to VPNv6 routes, the VPNv6 routes
                              that match the if-match clause recurse to tunnels according to the
                              tunnel policy specified in the apply clause. The VPNv6 routes that do not
                              match the if-match clause recurse to LSPs by default.
                    ----End

7.5.4 Verifying the Configuration
Procedure
                    ●   Run the display tunnel-selector tunnel-selector-name command to check
                        detailed information about the tunnel selector.
                    ●   Run the display tunnel { tunnel-id | all | statistics } command to check
                        tunnel information in the existing system.
                    ●   Run the display tunnel-policy tnl-policy-name subscriber statistics
                        command to check the number of times that a specified tunnel policy is used
                        by external services.
                    ●   Run the display bgp vpnv4 all routing-table ipv4-address [ mask [ longer-
                        prefixes ] | mask-length [ longer-prefixes ] ] command on an ASBR to check
                        information about the tunnels to which VPNv4 routes recurse.
                    ●   Run the display bgp vpnv6 all routing-table ipv6-address [ prefix-length ]
                        command on an ASBR to check information about the tunnels to which
                        VPNv6 routes recurse.
                    ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         812
VPN Configuration
VPN Configuration                                                  7 Tunnel Management Configuration




7.6 Maintaining Tunnel Management

7.6.1 Monitoring the Running Status of Tunnels
                    After configuring and applying a tunnel selector, run the following commands to
                    check tunnel selector and tunnel policy information in the system.

Procedure
                    ●   Run the display interface tunnel interface-number command to check tunnel
                        interface information.
                    ●   Run the display tunnel all command to check tunnel information.
                    ●   Run the display tunnel tunnel-id command to check detailed tunnel
                        information.
                    ●   Run the display tunnel-policy policy-name command to check the
                        configuration of a specified tunnel policy.
                    ●   Run the display ip vpn-instance verbose [ vpn-instance-name ] command to
                        check the tunnel policy applied to a VPN instance.
                    ●   Run the display ip routing-table vpn-instance [ ip-address ] verbose or
                        display ipv6 routing-table vpn-instance [ ipv6-address ] verbose command
                        to check tunnels used by VPN routes.
                    ----End




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                        813

