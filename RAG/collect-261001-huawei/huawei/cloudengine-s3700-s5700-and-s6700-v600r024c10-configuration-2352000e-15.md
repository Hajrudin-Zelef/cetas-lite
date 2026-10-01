---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-15
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [1127, 1264]
sha256: 8c0aaa5c0c534278827b297358d9214d5fde8fa3b5d3ec1474e3c990731d0725
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

2.4.3 Configuring Tunnel Routes
Context
                    Routes with GRE tunnel interfaces as the outbound interfaces must exist on both
                    the local and remote devices of a GRE tunnel, so that GRE-encapsulated packets
                    can be properly forwarded. These routes can be either static or dynamic routes.
                    When configuring GRE tunnel routes, note the following:
                    ●    If you configure a static route, it must be configured on both the local and
                         remote devices. Additionally, the destination address of the route must be
                         that of the original packet which is not encapsulated using GRE and the
                         outbound interface of the route must be the local tunnel interface. The
                         destination address cannot be the tunnel destination address or the remote
                         tunnel interface address.
                    ●    If you configure a dynamic routing protocol to generate a dynamic route, the
                         protocol must be enabled on both tunnel interfaces and the interfaces
                         connected to private networks. If you configure a route to the IP address of a
                         physical interface on the tunnel destination, ensure that the next-hop address
                         of the route is a physical interface address instead of a tunnel interface
                         address. Otherwise, packets cannot be properly forwarded.
                         For example, in Figure 2-6, the source physical interface of Tunnel 1 is
                         interface 1 on DeviceA, and the destination physical interface of Tunnel 1 is
                         interface 2 on DeviceB. When configuring a dynamic routing protocol to
                         generate a dynamic route, you need to enable the protocol on both the
                         tunnel interfaces and the interfaces connected to the PCs. In the IP routing
                         table of DeviceA, the outbound interface for the route to the subnet where
                         interface 2 on DeviceB resides cannot be Tunnel 1.
                         In real-world configurations, different routing protocols or different processes
                         of the same routing protocol need to be used on tunnel interfaces and public
                         network physical interfaces. This prevents a tunnel interface from being
                         selected as the outbound interface of routes to the tunnel destination, and

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                            13
VPN Configuration
VPN Configuration                                                                                   2 GRE Configuration


                        ensures that packets are forwarded through the GRE tunnel instead of the
                        physical interfaces.


                        Figure 2-6 Configuring a dynamic routing protocol for GRE




                    Perform the following steps on the devices at both ends of a GRE tunnel.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Choose either of the following methods to configure a GRE tunnel route:
                    ●   Configure an IPv4 static route.
                        ip route-static dest-ip-address { mask | mask-length } tunnel interface-number [ description text ]

                    ●   Configure an IPv6 static route.
                        ipv6 route-static dest-ipv6-address prefix-length tunnel interface-number [ description text ]

                    ●   Configure an IPv4 or IPv6 dynamic routing protocol. The dynamic routing
                        protocol can be IGP or BGP. For details about how to configure a dynamic
                        routing protocol, see the Configuration Guide-IP Routing Configuration.

                    ----End

2.4.4 Verifying the Configuration

Procedure
                    ●   Run the display interface tunnel [ interface-number ] command to check the
                        operating status of the tunnel interface.
                    ●   Run the display ip routing-table command to check the IP routing table. The
                        command output shows that the IP routing table contains a route with the
                        tunnel interface as the outbound interface.
                    ●   Run the display ipv6 routing-table command to check the IPv6 routing table.
                        The command output shows that the IPv6 routing table contains an IPv6
                        route with the tunnel interface as the outbound interface.
                    ●   Run the ping -a source-ip-address host command to check whether the two
                        ends of the tunnel can communicate with each other using IPv4 addresses.



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               14
VPN Configuration
VPN Configuration                                                                             2 GRE Configuration


                    ●    Run the ping ipv6 -a source-ipv6-address destination-ipv6-address command
                         to check whether the two ends of the tunnel can communicate with each
                         other using IPv6 addresses.

                    ----End

2.4.5 Example for Configuring an IPv4 over IPv4 GRE Tunnel

Networking Requirements
                    In Figure 2-7, DeviceA, DeviceB, and DeviceC belong to the IPv4 backbone
                    network and run OSPF. A direct link needs to be established between DeviceA and
                    DeviceC to ensure that PC1 and PC2 can communicate with each other. To meet
                    such a requirement, a GRE tunnel needs to be established between DeviceA and
                    DeviceC and static routes need to be configured so that packets between PC1 and
                    PC2 can be forwarded through tunnel interfaces on both ends of the tunnel.
                    DeviceA and DeviceC are the default gateways of PC1 and PC2, respectively.

                    Figure 2-7 Configuring an IPv4 over IPv4 GRE tunnel
                         NOTE

                    In this example, interface 1 and interface 2 represent VLANIF1 and VLANIF2, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure an IPv4 dynamic routing protocol for communication between the
                         devices.
                    2.   Create tunnel interfaces on DeviceA and DeviceC and specify the tunnel
                         source and destination addresses. The tunnel source address is the IP address
                         of the interface that sends packets, and the tunnel destination address is the
                         IP address of the interface that receives packets.
                    3.   Configure IP addresses for the tunnel interfaces to generate GRE tunnel
                         routes.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     15
VPN Configuration
VPN Configuration                                                                                            2 GRE Configuration


                    4.    Configure a static route between DeviceA and PC1 and between DeviceC and
                          PC2 and specify the local tunnel interface as the outbound interface of the
                          static route, so that traffic between PC1 and PC2 can be transmitted through
                          the GRE tunnel.

Procedure
         Step 1 Configure IP addresses for interfaces.

