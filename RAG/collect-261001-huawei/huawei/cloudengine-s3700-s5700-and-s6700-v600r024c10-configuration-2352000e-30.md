---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-30
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [3359, 3458]
sha256: 134d09ed8c34e4d4c3cc23ba20702589cea822952a81d61d03db91a3f94233c4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   CE1
                        #
                        sysname CE1
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        ip route-static 10.2.1.0 255.255.255.0 10.1.1.2
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.2.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        ip route-static 10.1.1.0 255.255.255.0 10.2.1.2
                        #
                        return



3.6 Configuring Basic IPv4 L3VPN over MPLS

3.6.1 Understanding IPv4 L3VPN over MPLS

Definition
                    IPv4 L3VPN over MPLS uses BGP to advertise VPN routes and uses MPLS to
                    forward VPN packets on the SP backbone network.
                    Figure 3-9 shows the basic IPv4 L3VPN model.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          52
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                    Figure 3-9 IPv4 L3VPN model




                    The basic IPv4 L3VPN model consists of three parts:

                    ●    Customer edge (CE): edge device on the customer network. A CE provides
                         interfaces to directly connect to the SP network. A CE can be a device, switch,
                         or host. Generally, a CE is unaware of VPNs and does not need to support
                         MPLS.
                    ●    Provider edge (PE): edge device on the SP network. A PE directly connects to
                         CEs. On an MPLS network, PEs process all VPN services and must have high
                         performance.
                    ●    Provider device (P): backbone device on the SP network. A P does not directly
                         connect to CEs. Ps only need to possess basic MPLS forwarding capabilities
                         and do not maintain VPN information.

                    Generally, PEs and Ps are managed by SPs, and CEs are managed by users. Users
                    can also delegate CE management to SPs.

                    A PE can connect to multiple CEs. A CE can connect to multiple PEs of the same
                    SP or of different SPs.

IPv4 L3VPN Route Advertisement
                    On an IPv4 L3VPN, CEs and PEs are responsible for advertising VPN routing
                    information, whereas Ps only need to maintain backbone network routes without
                    knowing VPN routing information. Generally, a PE maintains only the routes of
                    VPNs that the PE accesses, not all VPN routes. VPN route advertisement consists
                    of the following phases: route advertisement from the local CE to the ingress PE,
                    route advertisement from the ingress PE to the egress PE, and route advertisement
                    from the egress PE to the remote CE. After the process of route advertisement is
                    complete, the local and remote CEs can obtain reachable routes to each other, and
                    VPN routing information can be advertised on the backbone network. The
                    following describes the three phases of route advertisement:
                    1.   Route advertisement from the local CE to the ingress PE
                         After a peer relationship is established between a CE and the directly
                         connected PE, the CE advertises the local IPv4 routes to the PE. Static routes
                         or a routing protocol (RIP, OSPF, IS-IS, or BGP) can be used between the CE

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                53
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


