---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-21
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [2064, 2254]
sha256: d4b487b2ab8e90fbedb85f88aec56e0e6c8825b180a56c2b049e4ebf51068991
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Verifying the Configuration
                    After the configuration is complete, run the display ip routing-table command on
                    DeviceA and DeviceD. The command outputs show that the cost of the route to
                    the destination address of the peer device is 1.
                    # The following example uses the command output on DeviceA.
                    [DeviceA] display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : _public_
                           Destinations : 12        Routes : 12

                    Destination/Mask     Proto Pre Cost       Flags NextHop        Interface

                         10.1.1.0/24 Direct 0 0           D 10.1.1.1     Vlanif10
                         10.1.1.1/32 Direct 0 0           D 127.0.0.1    Vlanif10
                       10.1.1.255/32 Direct 0 0            D 127.0.0.1     Vlanif10
                         10.1.2.0/24 RIP    100 1         D 10.1.1.2     Vlanif10
                         10.1.3.0/24 RIP    100 1         D 192.168.1.2    Tunnel1
                      192.168.1.0/24 Direct 0 0            D 192.168.1.1    Tunnel1
                      192.168.1.1/32 Direct 0 0            D 127.0.0.1     Tunnel1
                     192.168.1.255/32 Direct 0 0            D 127.0.0.1     Tunnel1
                        127.0.0.0/8 Direct 0 0            D 127.0.0.1    InLoopBack0
                        127.0.0.1/32 Direct 0 0           D 127.0.0.1     InLoopBack0
                    127.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0
                    255.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0


Configuration Scripts
                    ●     DeviceA
                          #
                           sysname DeviceA
                          #
                          #
                          vlan batch 10
                          #
                          interface Vlanif10


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                28
VPN Configuration
VPN Configuration                                                             2 GRE Configuration

                         ip address 10.1.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface Tunnel1
                         ip address 192.168.1.1 255.255.255.0
                         tunnel-protocol gre
                         source 10.1.1.1
                         destination 10.1.2.2
                        #
                        rip 1
                         version 2
                         network 10.1.1.0
                        #
                        rip 2
                         version 2
                         network 192.168.1.0
                        #
                        return
                    ●   DeviceB
                        #
                         sysname DeviceB
                        #
                        vlan batch 10 20
                        #
                        interface Vlanif10
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface Vlanif20
                         ip address 10.1.2.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        rip 1
                         version 2
                         network 10.1.1.0
                         network 10.1.2.0
                        #
                        return
                    ●   DeviceC
                        #
                         sysname DeviceC
                        #
                        vlan batch 20 30
                        #
                        interface Vlanif20
                         ip address 10.1.2.2 255.255.255.0
                        #
                        interface Vlanif30
                         ip address 10.1.3.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface Tunnel1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   29
VPN Configuration
VPN Configuration                                                                                2 GRE Configuration

                         ip address 192.168.1.2 255.255.255.0
                         tunnel-protocol gre
                         source 10.1.2.2
                         destination 10.1.1.1
                        #
                        rip 1
                         version 2
                         network 10.1.2.0
                        #
                        rip 2
                         version 2
                         network 10.1.3.0
                         network 192.168.1.0
                        #
                        return

                    ●   DeviceD
                        #
                        sysname DeviceD
                        #
                        vlan batch 30
                        #
                        interface Vlanif30
                         ip address 10.1.3.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        rip 2
                         version 2
                         network 10.1.3.0
                        #
                        return



2.5 Maintaining GRE

2.5.1 Monitoring the Operating Status of GRE
Context
                    In routine maintenance, you can run GRE-related display commands to check the
                    GRE operating status.

Procedure
                    ●   Check the operating status of tunnel interfaces.
                        display interface tunnel [ interface-number ]

                    ●   Check whether the two ends of the tunnel can communicate with each other.
                        ping [ -a source-ip-address | -vpn-instance vpn-instance-name ] * host

                    ●   Check the numbers of keepalive messages and keepalive response messages
                        sent and received by GRE tunnel interfaces.
                        display keepalive packets count

                    ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      30
VPN Configuration
VPN Configuration                                                                       2 GRE Configuration


2.5.2 Enabling GRE Traffic Statistics Collection and Checking
GRE Traffic Statistics
Context
                    To check the network status or locate network faults, you can enable the GRE
                    traffic statistics collection function so that such statistics can be checked.

