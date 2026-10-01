---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-20
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [1900, 2063]
sha256: 667a54f021eb4d60d974f66331bebdeb7fbbb20da2a3f9a94e1a2f97d255577a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 10 20
                    [DeviceB] interface 10GE1/0/1
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 10
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface 10GE1/0/2
                    [DeviceB-10GE1/0/2] port link-type trunk
                    [DeviceB-10GE1/0/2] port trunk allow-pass vlan 20
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface Vlanif 10
                    [DeviceB-Vlanif10] ip address 10.1.1.2 255.255.255.0
                    [DeviceB-Vlanif10] quit
                    [DeviceB] interface Vlanif 20
                    [DeviceB-Vlanif20] ip address 10.1.2.1 255.255.255.0
                    [DeviceB-Vlanif20] quit

                    # Configure DeviceC.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceC
                    [DeviceC] vlan batch 20 30
                    [DeviceC] interface 10GE1/0/1
                    [DeviceC-10GE1/0/1] port link-type trunk
                    [DeviceC-10GE1/0/1] port trunk allow-pass vlan 20
                    [DeviceC-10GE1/0/1] quit
                    [DeviceC] interface 10GE1/0/2
                    [DeviceC-10GE1/0/2] port link-type trunk
                    [DeviceC-10GE1/0/2] port trunk allow-pass vlan 30
                    [DeviceC-10GE1/0/2] quit
                    [DeviceC] interface Vlanif 20
                    [DeviceC-Vlanif20] ip address 10.1.2.2 255.255.255.0
                    [DeviceC-Vlanif20] quit
                    [DeviceC] interface Vlanif 30
                    [DeviceC-Vlanif30] ip address 10.1.3.1 255.255.255.0
                    [DeviceC-Vlanif30] quit

                    # Configure DeviceD.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceD
                    [DeviceD] vlan batch 30
                    [DeviceD] interface 10GE1/0/1
                    [DeviceD-10GE1/0/1] port link-type trunk
                    [DeviceD-10GE1/0/1] port trunk allow-pass vlan 30
                    [DeviceD-10GE1/0/1] quit
                    [DeviceD] interface Vlanif 30
                    [DeviceD-Vlanif30] ip address 10.1.3.2 255.255.255.0
                    [DeviceD-Vlanif30] quit

         Step 2 Run RIP process 1 on devices.

                    # Configure DeviceA.
                    [DeviceA] rip 1
                    [DeviceA-rip-1] version 2
                    [DeviceA-rip-1] network 10.1.1.0
                    [DeviceA-rip-1] quit

                    # Configure DeviceB.
                    [DeviceB] rip 1
                    [DeviceB-rip-1] version 2
                    [DeviceB-rip-1] network 10.1.1.0
                    [DeviceB-rip-1] network 10.1.2.0
                    [DeviceB-rip-1] quit

                    # Configure DeviceC.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   26
VPN Configuration
VPN Configuration                                                                                            2 GRE Configuration

                    [DeviceC] rip 1
                    [DeviceC-rip-1] version 2
                    [DeviceC-rip-1] network 10.1.2.0
                    [DeviceC-rip-1] quit

                    # After the configuration is complete, run the display ip routing-table command
                    on DeviceA and DeviceC. The command outputs show that DeviceA and DeviceC
                    can learn RIP routes to the network segments of the remote interfaces.

                    # The following example uses the command output on DeviceA.
                    [DeviceA] display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Tables: _public_
                           Destinations : 8        Routes : 8

                    Destination/Mask     Proto Pre Cost       Flags NextHop        Interface

                         10.1.1.0/24 Direct 0 0           D 10.1.1.1     Vlanif10
                         10.1.1.1/32 Direct 0 0           D 127.0.0.1    Vlanif10
                       10.1.1.255/32 Direct 0 0            D 127.0.0.1     Vlanif10
                         10.1.2.0/24 RIP    100 1         D 10.1.1.2      Vlanif10
                        127.0.0.0/8 Direct 0 0            D 127.0.0.1    InLoopBack0
                        127.0.0.1/32 Direct 0 0           D 127.0.0.1     InLoopBack0
                    127.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0
                    255.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0

         Step 3 Configure tunnel interfaces.

                    # Configure DeviceA.
                    [DeviceA] interface tunnel 1
                    [DeviceA-Tunnel1] tunnel-protocol gre
                    [DeviceA-Tunnel1] ip address 192.168.1.1 255.255.255.0
                    [DeviceA-Tunnel1] source 10.1.1.1
                    [DeviceA-Tunnel1] destination 10.1.2.2
                    [DeviceA-Tunnel1] quit

                    # Configure DeviceC.
                    [DeviceC] interface tunnel 1
                    [DeviceC-Tunnel1] tunnel-protocol gre
                    [DeviceC-Tunnel1] ip address 192.168.1.2 255.255.255.0
                    [DeviceC-Tunnel1] source 10.1.2.2
                    [DeviceC-Tunnel1] destination 10.1.1.1
                    [DeviceC-Tunnel1] quit

                    # After the configuration is complete, the tunnel interfaces go up and can ping
                    each other successfully.

                    # The following example uses the command output on DeviceA.
                    [DeviceA] ping -a 192.168.1.1 192.168.1.2
                     PING 192.168.1.2: 56 data bytes, press CTRL_C to break
                       Reply from 192.168.1.2: bytes=56 Sequence=1 ttl=255 time=1 ms
                       Reply from 192.168.1.2: bytes=56 Sequence=2 ttl=255 time=1 ms
                       Reply from 192.168.1.2: bytes=56 Sequence=3 ttl=255 time=1 ms
                       Reply from 192.168.1.2: bytes=56 Sequence=4 ttl=255 time=1 ms
                       Reply from 192.168.1.2: bytes=56 Sequence=5 ttl=255 time=1 ms

                     --- 192.168.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 1/1/1 ms


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                27
VPN Configuration
VPN Configuration                                                                                            2 GRE Configuration


         Step 4 Run RIP process 2 on tunnel interfaces.
                    # Configure DeviceA.
                    [DeviceA] rip 2
                    [DeviceA-rip-2] version 2
                    [DeviceA-rip-2] network 192.168.1.0
                    [DeviceA-rip-2] quit

                    # Configure DeviceC.
                    [DeviceC] rip 2
                    [DeviceC-rip-2] version 2
                    [DeviceC-rip-2] network 192.168.1.0
                    [DeviceC-rip-2] network 10.1.3.0
                    [DeviceC-rip-2] quit

                    # Configure DeviceD.
                    [DeviceD] rip 2
                    [DeviceD-rip-2] version 2
                    [DeviceD-rip-2] network 10.1.3.0
                    [DeviceD-rip-2] quit

                    ----End

