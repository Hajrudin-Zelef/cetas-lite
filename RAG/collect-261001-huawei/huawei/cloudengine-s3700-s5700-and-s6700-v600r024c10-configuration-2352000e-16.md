---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-16
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [1265, 1405]
sha256: e5c63db0357b75324668bfb136e31138882ff10301f505b3f1ae788530ea34ae
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Configure an IP address for each involved interface according to Figure 2-7. For
                    detailed configurations, see the configuration scripts.

         Step 2 Configure IGP on the VPN backbone network.

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
                    [DeviceB-ospf-1-area-0.0.0.0] quit
                    [DeviceB-ospf-1] quit

                    # Configure DeviceC.
                    [DeviceC] ospf 1
                    [DeviceC-ospf-1] area 0
                    [DeviceC-ospf-1-area-0.0.0.0] network 172.21.1.0 0.0.0.255
                    [DeviceC-ospf-1-area-0.0.0.0] quit
                    [DeviceC-ospf-1] quit

                    After the configuration is complete, run the display ip routing-table command on
                    DeviceA and DeviceC. The command output shows that DeviceA and DeviceC can
                    learn OSPF routes to the network segments of the remote interfaces.

                    # The following example uses the command output on DeviceA.
                    [DeviceA] display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : _public_
                           Destinations : 11        Routes : 11

                    Destination/Mask     Proto Pre Cost        Flags NextHop         Interface

                         10.1.1.0/24 Direct 0 0            D 10.1.1.2     Vlanif2
                         10.1.1.2/32 Direct 0 0            D 127.0.0.1    Vlanif2
                       10.1.1.255/32 Direct 0 0             D 127.0.0.1     Vlanif2
                       172.20.1.0/24 Direct 0 0             D 172.20.1.1    Vlanif1
                       172.20.1.1/32 Direct 0 0             D 127.0.0.1     Vlanif1
                      172.20.1.255/32 Direct 0 0             D 127.0.0.1     Vlanif1
                       172.21.1.0/24 OSPF 10 2               D 172.20.1.2     Vlanif1
                        127.0.0.0/8 Direct 0 0             D 127.0.0.1    InLoopBack0
                        127.0.0.1/32 Direct 0 0            D 127.0.0.1     InLoopBack0
                    127.255.255.255/32 Direct 0 0             D 127.0.0.1      InLoopBack0
                    255.255.255.255/32 Direct 0 0             D 127.0.0.1      InLoopBack0


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                16
VPN Configuration
VPN Configuration                                                                                            2 GRE Configuration


         Step 3 Configure tunnel interfaces.
                    # Configure DeviceA.
                    [DeviceA] interface tunnel 1
                    [DeviceA-Tunnel1] tunnel-protocol gre
                    [DeviceA-Tunnel1] ip address 172.22.1.1 255.255.255.0
                    [DeviceA-Tunnel1] source 172.20.1.1
                    [DeviceA-Tunnel1] destination 172.21.1.2
                    [DeviceA-Tunnel1] quit

                    # Configure DeviceC.
                    [DeviceC] interface tunnel 1
                    [DeviceC-Tunnel1] tunnel-protocol gre
                    [DeviceC-Tunnel1] ip address 172.22.1.2 255.255.255.0
                    [DeviceC-Tunnel1] source 172.21.1.2
                    [DeviceC-Tunnel1] destination 172.20.1.1
                    [DeviceC-Tunnel1] quit

                    After the configuration is complete, the tunnel interfaces become up and can ping
                    each other successfully.
                    # The following example uses the command output on DeviceA.
                    [DeviceA] ping -a 172.22.1.1 172.22.1.2
                     PING 172.22.1.2: 56 data bytes, press CTRL_C to break
                       Reply from 172.22.1.2: bytes=56 Sequence=1 ttl=255 time=24 ms
                       Reply from 172.22.1.2: bytes=56 Sequence=2 ttl=255 time=33 ms
                       Reply from 172.22.1.2: bytes=56 Sequence=3 ttl=255 time=48 ms
                       Reply from 172.22.1.2: bytes=56 Sequence=4 ttl=255 time=33 ms
                       Reply from 172.22.1.2: bytes=56 Sequence=5 ttl=255 time=36 ms
                     --- 172.22.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 24/34/48 ms

         Step 4 Configure static routes.
                    # Configure DeviceA.
                    [DeviceA] ip route-static 10.2.1.0 255.255.255.0 tunnel1

                    # Configure DeviceC.
                    [DeviceC] ip route-static 10.1.1.0 255.255.255.0 tunnel1

                    ----End

Verifying the Configuration
                    After the configuration is complete, run the display ip routing-table command on
                    DeviceA and DeviceC. The command output shows the static route with the local
                    tunnel interface as the outbound interface and destined for the user-side network
                    segment of the remote end.
                    # The following example uses the command output on DeviceA.
                    [DeviceA] display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : _public_
                           Destinations : 15        Routes : 15

                    Destination/Mask     Proto Pre Cost        Flags NextHop         Interface


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                17
VPN Configuration
VPN Configuration                                                                         2 GRE Configuration


                          10.1.1.0/24 Direct 0 0        D 10.1.1.2     Vlanif2
                          10.1.1.2/32 Direct 0 0        D 127.0.0.1    Vlanif2
                        10.1.1.255/32 Direct 0 0         D 127.0.0.1     Vlanif2
                       10.2.1.0/24 Static 60 0         D 172.22.1.1    Tunnel1
                        172.20.1.0/24 Direct 0 0         D 172.20.1.1    Vlanif1
                        172.20.1.1/32 Direct 0 0         D 127.0.0.1     Vlanif1
                      172.20.1.255/32 Direct 0 0          D 127.0.0.1     Vlanif1
                        172.21.1.0/24 OSPF 10 2           D 172.20.1.2     Vlanif1
                        172.22.1.0/24 Direct 0 0         D 172.22.1.1    Tunnel1
                        172.22.1.1/32 Direct 0 0         D 127.0.0.1     Tunnel1
                      172.22.1.255/32 Direct 0 0          D 127.0.0.1     Tunnel1
                         127.0.0.0/8 Direct 0 0         D 127.0.0.1    InLoopBack0
                         127.0.0.1/32 Direct 0 0         D 127.0.0.1    InLoopBack0
                    127.255.255.255/32 Direct 0 0          D 127.0.0.1      InLoopBack0
                    255.255.255.255/32 Direct 0 0          D 127.0.0.1      InLoopBack0


