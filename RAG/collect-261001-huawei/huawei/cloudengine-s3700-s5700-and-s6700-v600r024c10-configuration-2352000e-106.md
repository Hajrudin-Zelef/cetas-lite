---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-106
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [14792, 14943]
sha256: b329161fd434f241e2c95eea0f783773c0a1fd5b8da1e297324e2cdc7879e5fd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


Context
                    After an interface is bound to a VPN instance, the interface becomes a part of the
                    VPN. Packets entering the interface will be forwarded based on the VRF table of
                    the VPN.
                    In a Hub-CE dual-link access scenario, two interfaces or sub-interfaces are
                    required on the Hub-PE:
                    ●    One is bound to vpn_in for importing routes advertised by Spoke-PEs.
                    ●    The other is bound to vpn_out for exporting the routes of all hub and spoke
                         sites.
                    If a Hub-PE and a Hub-CE are connected through a single link, only one interface
                    or sub-interface on the Hub-PE needs to be bound to the VPN instance vpnhub.
                    Perform the following steps on the Hub-PE and all Spoke-PEs.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the interface to be bound to the VPN instance.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode from Layer 2 to Layer 3.
                    undo portswitch

                    Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                    S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                    Layer 2 mode to Layer 3 mode using the undo portswitch command. Determine
                    whether to perform this step based on the current interface working mode.
         Step 4 Bind the interface to a VPN instance.
                    ip binding vpn-instance vpn-instance-name

                          NOTE

                         Running the ip binding vpn-instance command deletes Layer 3 (including IPv4 and IPv6)
                         configurations, such as IP address and routing protocol configurations, on the involved
                         interface. If needed, reconfigure them after running the command.

         Step 5 Configure an IP address for the interface.
                    ip address ip-address { mask | mask-length }

                    ----End

3.15.4 Configuring Route Exchange Between the Hub-PE and
Spoke-PE

Context
                    An MP-IBGP peer relationship needs to be established between the Hub-PE and
                    each Spoke-PE. Spoke-PEs do not need to exchange routes directly and therefore
                    do not need to establish MP-IBGP peer relationships.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                  235
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration


                    Perform the following steps on the Hub-PE and all Spoke-PEs.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Configure the remote PE as a peer.
                    peer ipv4-address as-number as-number

         Step 4 Specify an interface for setting up a TCP connection with the BGP peer.
                    peer ipv4-address connect-interface loopback interface-number

                          NOTE

                        A PE must use a loopback interface address with a 32-bit mask to set up an MP-IBGP peer
                        relationship with the peer PE so that VPN routes can recurse to tunnels. The route to the
                        local loopback interface is advertised to the peer PE using IGP on the MPLS backbone
                        network.

         Step 5 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4 [ unicast ]

         Step 6 Enable the ability to exchange VPN-IPv4 routes with a specified peer.
                    peer ipv4-address enable

                    ----End

3.15.5 Configuring Route Exchange Between the PE and CE

Context
                    The routing protocol run between Spoke-PEs and Spoke-CEs is related to the
                    routing protocol run between the Hub-PE and Hub-CE. EBGP, IGP, and static routes
                    (including default routes) can be used between the Hub-PE and Hub-CE.
                    Determine which one to use as required.

Procedure
                    ●    Configure EBGP between the Hub-PE and Hub-CE.
                         For configuration details, see Configuring an IPv4 VPN Instance.
                         In this mode, EBGP, IGP, or static routes (including default routes) can be used
                         between Spoke-PEs and Spoke-CEs.

                               NOTE

                              If EBGP runs both between Spoke-PEs and Spoke-CEs and between the Hub-PE and
                              Hub-CE, you must run the peer ip-address allow-as-loop [ number ] command in the
                              BGP-VPN instance IPv4 address family view of the Hub-PE to allow routing loops. If
                              number is set to 1, routes with the local AS number repeated once in the AS_Path list
                              are allowed.
                    ●    Configure IGP between the Hub-PE and Hub-CE.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                  236
VPN Configuration
VPN Configuration                                                                          3 IPv4 L3VPN Configuration


                        For configuration details, see Configuring an IPv4 VPN Instance.

                        In this mode, only IGP or static routes (including default routes) can be used
                        between Spoke-PEs and Spoke-CEs.
                    ●   Configure static routes (including default routes) between the Hub-PE and
                        Hub-CE.

                        For configuration details, see Configuring an IPv4 VPN Instance.

                        In this mode, EBGP, IGP, or static routes (including default routes) can be used
                        between Spoke-PEs and Spoke-CEs.

                        –     In a CE dual-link access scenario, if the Hub-CE adopts the default route
                              to access the Hub-PE, you need to run the following commands on the
                              Hub-PE to enable it to advertise the default route to all the Spoke-PEs.
                              i.     Enter the system view.
                                     system-view

                              ii.    Configure a default route.
                                     ip route-static vpn-instance vpn-instance-name 0.0.0.0 0.0.0.0 nexthop-address [ tag
                                     tag ] [ description text ]

                                     In this example, vpn-instance-name specifies vpn_out and nexthop-
                                     address specifies the IP address of the Hub-CE interface that
                                     connects to the Hub-PE interface bound to vpn_out.
                              iii.   Enter the BGP view.
                                     bgp as-number

                              iv.    Enter the BGP-VPN instance IPv4 address family view.
                                     ipv4-family vpn-instance vpn-instance-name

                              v.     Import the default route of the local site.
                                     network 0.0.0.0 0

