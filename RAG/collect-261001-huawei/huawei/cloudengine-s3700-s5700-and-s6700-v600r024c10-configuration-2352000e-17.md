---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-17
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [1406, 1583]
sha256: 13a4359bf67a26394c71d46ac05cbe1c4a5e91e98ab11bbd47ccfb1c2a8325c7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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
                          ip address 10.1.1.2 255.255.255.0
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
                          ip address 172.22.1.1 255.255.255.0
                          tunnel-protocol gre
                          source 172.20.1.1
                          destination 172.21.1.2
                         #
                         ospf 1
                          area 0.0.0.0
                           network 172.20.1.0 0.0.0.255
                         #
                         ip route-static 10.2.1.0 255.255.255.0 Tunnel1
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
                         interface Vlanif2
                          ip address 172.21.1.1 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type access
                          port default vlan 1
                         #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                              18
VPN Configuration
VPN Configuration                                                                             2 GRE Configuration

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

                    ●    DeviceC
                         #
                         sysname DeviceC
                         #
                         vlan batch 1 2
                         #
                         interface Vlanif1
                          ip address 172.21.1.2 255.255.255.0
                         #
                         interface Vlanif2
                          ip address 10.2.1.2 255.255.255.0
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
                          ip address 172.22.1.2 255.255.255.0
                          tunnel-protocol gre
                          source 172.21.1.2
                          destination 172.20.1.1
                         #
                         ospf 1
                          area 0.0.0.0
                          network 172.21.1.0 0.0.0.255
                         #
                         ip route-static 10.1.1.0 255.255.255.0 Tunnel1
                         #
                         return


2.4.6 Example for Configuring an IPv6 over IPv4 GRE Tunnel

Networking Requirements
                    In Figure 2-8, DeviceA, DeviceB, and DeviceC belong to the IPv4 backbone
                    network and run OSPF. An IPv6 direct link needs to be established between
                    DeviceA and DeviceC to ensure that PC1 and PC2 can communicate with each
                    other. To meet such a requirement, an IPv6 over IPv4 GRE tunnel needs to be
                    established between DeviceA and DeviceC and IPv6 static routes need to be
                    configured so that packets between PC1 and PC2 can be forwarded through
                    tunnel interfaces on both ends of the tunnel. DeviceA and DeviceC are the default
                    gateways of PC1 and PC2, respectively.

                    Figure 2-8 Configuring an IPv6 over IPv4 GRE tunnel
                         NOTE

                    In this example, interface 1 and interface 2 represent VLANIF1 and VLANIF2, respectively.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    19
VPN Configuration
VPN Configuration                                                                   2 GRE Configuration




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure an IPv4 dynamic routing protocol for communication between the
                         devices.
                    2.   Create tunnel interfaces on DeviceA and DeviceC and specify the tunnel
                         source and destination addresses. The tunnel source address is the IPv4
                         address of the interface that sends packets, and the tunnel destination
                         address is the IPv4 address of the interface that receives packets.
                    3.   Configure IPv6 addresses for the tunnel interfaces to generate GRE tunnel
                         routes.
                    4.   Configure an IPv6 static route between DeviceA and PC1 and between
                         DeviceC and PC2 and specify the local tunnel interface as the outbound
                         interface of the static route, so that IPv6 traffic between PC1 and PC2 can be
                         transmitted through the GRE tunnel.

Procedure
         Step 1 Configure IPv4 or IPv6 addresses for interfaces.
                    Configure an IP address for each involved interface according to Figure 2-8. For
                    detailed configurations, see the configuration scripts.
         Step 2 Configure IGP on the IPv4 backbone network.
                    # Configure DeviceA.
                    [DeviceA] ospf 1
                    [DeviceA-ospf-1] area 0
                    [DeviceA-ospf-1-area-0.0.0.0] network 172.20.1.0 0.0.0.255
                    [DeviceA-ospf-1-area-0.0.0.0] quit
                    [DeviceA-ospf-1] quit

                    # Configure DeviceB.
                    [DeviceB] ospf 1
                    [DeviceB-ospf-1] area 0
                    [DeviceB-ospf-1-area-0.0.0.0] network 172.20.1.0 0.0.0.255
                    [DeviceB-ospf-1-area-0.0.0.0] network 172.21.1.0 0.0.0.255


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           20
VPN Configuration
VPN Configuration                                                                                            2 GRE Configuration

                    [DeviceB-ospf-1-area-0.0.0.0] quit
                    [DeviceB-ospf-1] quit

                    # Configure DeviceC.
                    [DeviceC] ospf 1
                    [DeviceC-ospf-1] area 0
                    [DeviceC-ospf-1-area-0.0.0.0] network 172.21.1.0 0.0.0.255
                    [DeviceC-ospf-1-area-0.0.0.0] quit
                    [DeviceC-ospf-1] quit

