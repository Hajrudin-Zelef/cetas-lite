---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-108
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [15174, 15323]
sha256: b1197ec2ac4dc370b53b4074a308dcaf56b3dd8dbadb5380555a8c89e3e3cdf2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                          network 10.1.3.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 300 400 500
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         ospf cost 10
                        #
                        interface Vlanif400
                         ip address 10.1.4.2 255.255.255.0
                         ospf cost 10
                        #
                        interface Vlanif500
                         ip address 10.1.5.1 255.255.255.0
                         ospf cost 10
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.1.3.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                        #
                        return
                 ●      LSR5
                        #
                        sysname LSR5
                        #
                        vlan batch 500
                        #
                        interface Vlanif500
                         ip address 10.1.5.2 255.255.255.0
                         ospf cost 10
                        #
                        interface 10GE1/0/1
                         port link-type trunk


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      257
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration

                         port trunk allow-pass vlan 500
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.1.5.0 0.0.0.255
                        #
                        return


4.8.6 Example for Configuring Forwarding Adjacency to Steer
Traffic to a TE Tunnel
Networking Requirements
                 If an MPLS TE tunnel is created, traffic is not automatically steered to the TE
                 tunnel. This requires you to configure a traffic steering mode. Forwarding
                 adjacency is a commonly used mode. With forwarding adjacency configured, the
                 system not only involves a TE tunnel in local IGP route calculation as a logical link,
                 but also advertises it as a common IGP route to neighbors. This process is different
                 from that when IGP shortcut is use. Set a metric value for the TE tunnel to make it
                 preferred locally or by other devices. Traffic can then be steered to the TE tunnel.
                 On the network shown in Figure 4-13, the devices run OSPF to communicate.
                 Establish a TE tunnel from LSR1 to LSR3, with LSR2 being the transit node. The
                 number marked on a link indicates its cost. If traffic destined for LSR3 exists on
                 both LSR1 and LSR5, the traffic is forwarded through VLANIF300 on LSR4
                 according to the OSPF route selection result. Assume that the bandwidth of the
                 link between LSR3 and LSR4 is 100 Mbit/s, the bandwidth required by the traffic
                 sent from LSR1 to LSR3 is 10 Mbit/s, and the bandwidth required by the traffic
                 sent from LSR5 to LSR3 is 100 Mbit/s. A total of 110 Mbit/s bandwidth is required.
                 In this case, congestion occurs on the link between LSR3 and LSR4, causing traffic
                 delay or loss.
                 To resolve this issue, configure forwarding adjacency on the TE tunnel interface of
                 LSR1. In this manner, all traffic from LSR1 to LSR3 is forwarded through the TE
                 tunnel. Some traffic from LSR5 to LSR3 is forwarded through LSR4, and the other
                 traffic is sent to LSR1 and forwarded through the TE tunnel. This prevents
                 congestion on the link between LSR3 and LSR4.

                         NOTE

                        ● After forwarding adjacency is configured, LSR1 advertises the TE tunnel as an OSPF
                          route to its neighbors. OSPF needs to perform bidirectional link check. Therefore, a TE
                          tunnel from LSR3 to LSR1 needs to be established and forwarding adjacency needs to be
                          enabled on the tunnel interface.
                        ● To avoid loops in this scenario, ensure that all connected interfaces have STP disabled
                          and are removed from VLAN1. If STP is enabled and VLANIF interfaces of switches are
                          used to construct a Layer 3 ring network, an interface on the network will be blocked.
                          As a result, Layer 3 services on the network cannot run properly.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    258
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 Figure 4-13 Network diagram of forwarding adjacency to steer traffic to a TE
                 tunnel




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign IP addresses to interfaces and configure OSPF to generate reachable
                        public network routes for node communication. Configure OSPF route costs.
                 2.     On LSR1, create a TE tunnel destined for LSR3 along the path LSR1 -> LSR2 ->
                        LSR3. On LSR3, create a TE tunnel destined for LSR1 along the path LSR3 ->
                        LSR2 -> LSR1. In this example, use RSVP-TE to establish dynamic MPLS TE
                        tunnels. Configure LSR IDs and enable MPLS TE, RSVP-TE, CSPF, and OSPF TE
                        globally on each node and involved interfaces. On the tunnel ingress, create a
                        tunnel interface and specify the tunnel IP address, tunneling protocol,
                        destination address, tunnel ID, and dynamic signaling protocol (RSVP-TE).
                        Enable the forwarding adjacency function on the TE tunnel interfaces of LSR1
                        and LSR3 and set the IGP metric value of the TE tunnel.

