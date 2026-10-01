---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-19
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [1718, 1899]
sha256: 87ffab651fd3bf4974603fe335c50109326dc9d8929346aa2990bc249f1bf9cd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Destination : 2001:DB8:1::2                 PrefixLength : 128
                    NextHop      : ::1                      Preference : 0
                    Cost      :0                          Protocol    : Direct
                    RelayNextHop : ::                        TunnelID      : 0x0
                    Interface : Vlanif2                    Flags       :D

                    Destination : 2001:DB8:2::                  PrefixLength : 64
                    NextHop      : ::1                     Preference : 60
                    Cost      :0                          Protocol    : Static
                    RelayNextHop : ::                        TunnelID      : 0x0
                    Interface : Tunnel1                    Flags       :D




Configuration Scripts
                    ●    DeviceA
                         #
                         sysname DeviceA
                         #
                         vlan batch 1 2
                         #
                         interface Vlanif1
                          ip address 172.20.1.1 255.255.255.0
                         #
                         interface Vlanif2
                          ipv6 enable
                          ipv6 address 2001:DB8:1::2/64
                         #
                         interface 10GE1/0/1
                          port link-type access
                          port default vlan 1
                         #
                         interface 10GE1/0/2
                          port link-type access
                          port default vlan 2
                         #
                         interface Tunnel1
                          ipv6 enable
                          ipv6 address 2001:DB8:3::1/64
                          tunnel-protocol gre
                          source 172.20.1.1
                          destination 172.21.1.2
                         #
                         ospf 1
                          area 0.0.0.0
                           network 172.20.1.0 0.0.0.255
                         #
                         ipv6 route-static 2001:DB8:2:: 64 Tunnel1
                         #
                         return
                    ●    DeviceB
                         #
                         sysname DeviceB
                         #
                         vlan batch 1 2
                         #
                         interface Vlanif1
                          ip address 172.20.1.2 255.255.255.0
                         #


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                        23
VPN Configuration
VPN Configuration                                                                   2 GRE Configuration

                        interface Vlanif2
                         ip address 172.21.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 1
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 2
                        #
                        ospf 1
                         area 0.0.0.0
                          network 172.20.1.0 0.0.0.255
                          network 172.21.1.0 0.0.0.255
                        #
                        return

                    ●   DeviceC
                        #
                        sysname DeviceC
                        #
                        #
                        vlan batch 1 2
                        #
                        interface Vlanif1
                         ip address 172.21.1.2 255.255.255.0
                        #
                        interface Vlanif2
                         ipv6 enable
                         ipv6 address 2001:DB8:2::2/64
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 1
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 2
                        #
                        interface Tunnel1
                         ipv6 enable
                         ipv6 address 2001:DB8:3::2/64
                         tunnel-protocol gre
                         source 172.21.1.2
                         destination 172.20.1.1
                        #
                        ospf 1
                         area 0.0.0.0
                         network 172.21.1.0 0.0.0.255
                        #
                        ipv6 route-static 2001:DB8:1:: 64 Tunnel1
                        #
                        return


2.4.7 Example for Enlarging the Operation Scope of a
Network with a Hop Limit
Networking Requirements
                    In Figure 2-9, DeviceA, DeviceB, DeviceC, and DeviceD run RIP to ensure
                    reachability. Data sent from DeviceA to DeviceD should pass through only one
                    hop. That is, the route cost is 1. Without changing the RIP network topology, there
                    are two hops between DeviceA and DeviceD. To eliminate one hop, you need to
                    set up a GRE tunnel between DeviceA and DeviceC. Although the logical hop

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          24
VPN Configuration
VPN Configuration                                                                             2 GRE Configuration


                    count is 1, there are two devices on the path from DeviceA to DeviceD. In this way,
                    the hop count allowed on a RIP network is increased.

                    Figure 2-9 Enlarging the operation scope of a network with a hop limit
                          NOTE

                    In this example, interface1, interface2, and interface3 represent VLANIF 10, VLANIF 20, and
                    VLANIF 30, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Run RIP process 1 on DeviceA, DeviceB, and DeviceC to ensure reachability
                         between them.
                    2.   Set up a GRE tunnel between DeviceA and DeviceC to hide DeviceB.
                    3.   Run RIP process 2 on DeviceA, DeviceC, and DeviceD to forward packets over
                         the GRE tunnel. The actual hop counts allowed on a RIP network are
                         increased.

Procedure
         Step 1 Configure an IP address for each physical interface.

                    # Configure DeviceA.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan batch 10
                    [DeviceA] interface 10GE1/0/1
                    [DeviceA-10GE1/0/1] port link-type trunk
                    [DeviceA-10GE1/0/1] port trunk allow-pass vlan 10
                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] interface Vlanif 10
                    [DeviceA-Vlanif10] ip address 10.1.1.1 255.255.255.0
                    [DeviceA-Vlanif10] quit

                    # Configure DeviceB.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      25
VPN Configuration
VPN Configuration                                                              2 GRE Configuration

