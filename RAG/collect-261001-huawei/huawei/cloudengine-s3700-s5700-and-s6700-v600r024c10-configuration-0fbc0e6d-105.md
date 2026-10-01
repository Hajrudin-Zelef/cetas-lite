---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-105
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [14794, 14919]
sha256: 3dc55499f500459ad01d0a8b20cb10f2b6d0eefea79ababe5f0994fd295173e4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Verifying the Configuration
                 ●      Run the display ip routing-table command to check whether an MPLS TE
                        tunnel interface is used as the outbound interface of a route.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              250
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


4.8.5 Example for Configuring IGP Shortcut to Steer Traffic to
a TE Tunnel
Networking Requirements
                 If an MPLS TE tunnel is created, traffic is not automatically steered to the TE
                 tunnel. This requires you to configure a traffic steering mode. IGP shortcut is a
                 commonly used mode. In this mode, a TE tunnel is used as a logical link to
                 participate in local IGP route calculation. Set a metric value for the TE tunnel to
                 make it the preferred route. Traffic can then be steered to the TE tunnel.
                 On the network shown in Figure 4-12, the devices run OSPF to communicate.
                 Establish a TE tunnel from LSR1 to LSR3, with LSR2 being the transit node. The
                 number marked on a link indicates its cost. If both traffic destined for LSR5 and
                 traffic destined for LSR3 are present on LSR1, they are forwarded through the
                 same interface, VLANIF400, according to the OSPF route selection result. Assume
                 that the bandwidth of the link between LSR1 and LSR4 is 100 Mbit/s, the
                 bandwidth required by the traffic destined for LSR3 is 50 Mbit/s, and the
                 bandwidth required by the traffic destined for LSR5 is 60 Mbit/s. A total of 110
                 Mbit/s bandwidth is required. In this case, congestion occurs on the link between
                 LSR1 and LSR4, causing traffic delay or loss.
                 To resolve this issue, configure IGP shortcut on the TE tunnel interface of LSR1 to
                 steer traffic destined for LSR3 to the TE tunnel. Then, this part of traffic is
                 forwarded through VLANIF100, preventing network congestion.

                         NOTE

                        ● After IGP shortcut is configured, LSR1 does not advertise the TE tunnel as a route to its
                          neighbors. Therefore, the TE tunnel can only participate in local route calculation on
                          LSR1.
                        ● To avoid loops in this scenario, ensure that all connected interfaces have STP disabled
                          and are removed from VLAN1. If STP is enabled and VLANIF interfaces of switches are
                          used to construct a Layer 3 ring network, an interface on the network will be blocked.
                          As a result, Layer 3 services on the network cannot run properly.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      251
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 Figure 4-12 Network diagram of IGP shortcut to steer traffic to a TE tunnel




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign IP addresses to interfaces and configure OSPF to generate reachable
                        public network routes for node communication. Configure OSPF route costs.
                 2.     On LSR1, create a TE tunnel destined for LSR3 along the path LSR1 -> LSR2 ->
                        LSR3. In this example, use RSVP-TE to establish dynamic MPLS TE tunnels.
                        Configure LSR IDs and enable MPLS TE, RSVP-TE, CSPF, and OSPF TE globally
                        on each node and involved interfaces. On the tunnel ingress, create a tunnel
                        interface and specify the tunnel IP address, tunneling protocol, destination
                        address, tunnel ID, and dynamic signaling protocol (RSVP-TE).
                 3.     Enable the IGP shortcut function on the TE tunnel interface of LSR1 and set
                        the IGP metric value of the TE tunnel.

Procedure
         Step 1 Assign an IP address to each interface and configure OSPF and OSPF costs for
                links.
                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100 400
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 255.255.255.0
                 [LSR1-Vlanif100] ospf cost 15
                 [LSR1-Vlanif100] quit
                 [LSR1] interface vlanif 400
                 [LSR1-Vlanif400] ip address 10.1.4.1 255.255.255.0
                 [LSR1-Vlanif400] ospf cost 10
                 [LSR1-Vlanif400] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         252
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration

                 [LSR1] interface 10ge 1/0/2
                 [LSR1-10GE1/0/2] port link-type trunk
                 [LSR1-10GE1/0/2] port trunk allow-pass vlan 400
                 [LSR1-10GE1/0/2] quit
                 [LSR1] interface loopback 1
                 [LSR1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                 [LSR1-LoopBack1] quit
                 [LSR1] ospf 1
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.4.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2, LSR3, LSR4, and LSR5 are similar to the configuration
                 of LSR1. For detailed configurations, see Configuration Scripts.
                 # After the configuration is complete, check the IP routing tables on LSR1, LSR2,
                 and LSR3. The following example uses the command output on LSR1.
                 [LSR1] display ip routing-table
                 Proto: Protocol        Pre: Preference
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                        Destinations : 16        Routes : 16

                 Destination/Mask     Proto Pre Cost        Flags NextHop                          Interface

