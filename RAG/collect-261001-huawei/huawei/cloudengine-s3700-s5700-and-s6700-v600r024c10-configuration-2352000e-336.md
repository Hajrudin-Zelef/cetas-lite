---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-336
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [50221, 50360]
sha256: 523d271bce438912d4afb99f270d47f5c7a517e151d53d30cc108432c97d1b74
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        b.    Enter the VPN instance view.
                              ip vpn-instance vpn-instance-name

                        c.    Enter the VPN instance IPv4 address family view.
                              ipv4-family

                        d.    Apply a tunnel policy to the VPN instance IPv4 address family.
                              tnl-policy policy-name

                    ●   Apply a tunnel policy to an IPv6 L3VPN.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the VPN instance view.
                              ip vpn-instance vpn-instance-name

                        c.    Enter the VPN instance IPv6 address family view.
                              ipv6-family

                        d.    Apply a tunnel policy to the VPN instance IPv6 address family.
                              tnl-policy policy-name

                    ●   Apply a tunnel policy to a 6PE peer.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    808
VPN Configuration
VPN Configuration                                                                7 Tunnel Management Configuration

                              bgp as-number

                        c.    Enter the BGP-IPv6 unicast address family view.
                              ipv6-family unicast

                        d.    Apply a tunnel policy to a specified IPv4 peer.
                              peer ipv4-address tnl-policy policy-name

                    ●   Apply a tunnel policy to VPWS.
                        –     Apply a tunnel policy to SVC VPWS.
                              For details about how to configure SVC VPWS, see 5.8 Configuring SVC
                              VPWS. Perform the following steps on the PEs with VCs configured:
                              i.     Enter the system view.
                                     system-view

                              ii.    Enter the AC interface view.
                                     interface interface-type interface-number

                              iii.   Apply a tunnel policy to SVC VPWS.
                                     mpls static-l2vc { { destination ip-address | pw-template pw-template-name vc-id } * |
                                     destination ip-address [ vc-id ] } transmit-vpn-label transmit-label-value receive-vpn-
                                     label receive-label-value [ tunnel-policy tnl-policy-name | [ control-word | no-control-
                                     word ] | [ raw | tagged ] ] *

                        –     Apply a tunnel policy to LDP VPWS.
                              For details about how to configure LDP VPWS, see 5.7 Configuring LDP
                              VPWS. Perform the following steps on the PEs with VCs configured:
                              i.     Enter the system view.
                                     system-view

                              ii.    Enter the AC interface view.
                                     interface interface-type interface-number

                              iii.   Apply a tunnel policy to LDP VPWS.
                                     mpls l2vc { ip-address | pw-template pw-template-name } * vc-id [ [ control-word | no-
                                     control-word ] | [ raw | tagged ] | tunnel-policy policy-name | ignore-standby-state ] *

                        –     Apply a tunnel policy to BGP VPWS.
                              For details about the configuration, see 5.6.2 Configuring a Local BGP
                              VPWS Connection or 5.6.3 Configuring a Remote BGP VPWS
                              Connection based on service requirements.
                    ----End

7.4.4 Verifying the Configuration
Procedure
                    ●   Run the display tunnel { tunnel-id | all | statistics } command to check
                        tunnel information in the existing system.
                    ●   Run the display tunnel-policy [ tnl-policy-name ] command to check
                        information about tunnel policies in the existing system.
                    ●   Run the display tunnel-policy tnl-policy-name subscriber statistics
                        command to check the number of times that a specified tunnel policy is used
                        by external services.
                    ●   Run the display ip vpn-instance verbose [ vpn-instance-name ] command to
                        check the tunnel policy applied to a VPN instance.
                    ----End

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                             809
VPN Configuration
VPN Configuration                                                   7 Tunnel Management Configuration




7.5 Configuring a Tunnel Selector

7.5.1 Understanding Tunnel Selectors
                    In VPN networking, after a tunnel policy is applied to a VPN instance, all the
                    routes of the VPN instance recurse to the same tunnel according to the tunnel
                    policy. In inter-AS VPN Option B networking, the ASBR receives all VPN routes
                    from PE peers. Currently, the system recurses VPN routes to LSPs. Sometimes,
                    these VPN routes need to recurse to tunnels to guarantee bandwidth. If you do
                    not want to create VPN instances on ASBRs, you cannot use tunnel policies. This is
                    where tunnel selectors come in.

                    A tunnel selector can apply a tunnel policy to VPN routes so that desired tunnels
                    can be selected based on the tunnel policy.


Implementation
                    A tunnel policy selector consists of one or more nodes, and the relationship
                    between these nodes is OR. The system checks the nodes according to index
                    numbers. If a route matches a node in the tunnel selector, the route stops the
                    matching process. Each node comprises a set of if-match and apply clauses:

                    ●   The if-match clauses define the matching rules that are used to match certain
                        route attributes, such as the next hop and RD. The relationship between the
                        if-match clauses of a node is AND. A route matches a node only when the
                        route meets all the matching rules specified by the if-match clauses of the
                        node.
                    ●   The apply clauses specify actions. When a route matches a node, the apply
                        clauses select a corresponding tunnel policy for the route. This tunnel policy
                        can select other types of tunnels to carry services by means of tunnel type
                        prioritizing or tunnel binding.

                    The node matching modes of a tunnel policy selector are as follows:

                    ●   Permit: If a route matches all the if-match clauses of a node, the route
                        matches the tunnel policy selector and all the actions defined by apply
                        clauses are performed on the route. If a route does not match any if-match
                        clauses of a node, the route continues to be match against the next node.
                    ●   Deny: In this mode, the apply clauses are not implemented. If a route matches
                        all the if-match clauses of the node, the route is denied and no longer
                        matched against other nodes of the tunnel policy selector.

7.5.2 Configuring a Tunnel Selector

