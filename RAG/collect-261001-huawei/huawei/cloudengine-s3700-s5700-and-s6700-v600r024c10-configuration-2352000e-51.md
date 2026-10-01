---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-51
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [6377, 6568]
sha256: 45ce1af09b1265a4237435ed0e053d0be294d5e3726d4e2f204c5c8bf64ec5aa
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        10.3.1.0/24 Direct 0 0             D 10.3.1.2     Vlanif300
                        10.3.1.2/32 Direct 0 0             D 127.0.0.1    Vlanif300
                       10.3.1.255/32 Direct 0 0             D 127.0.0.1     Vlanif300
                         3.3.3.3/32 RIP 100 1              D 10.3.1.1     Vlanif300
                        10.5.1.0/24 Direct 0 0             D 10.5.1.2     Vlanif100
                        10.5.1.1/32 Direct 0 0             D 10.5.1.1     Vlanif100
                        10.5.1.2/32 Direct 0 0             D 127.0.0.1    Vlanif100
                       10.5.1.255/32 Direct 0 0             D 127.0.0.1     Vlanif100
                    255.255.255.255/32 Direct 0 0             D 127.0.0.1      InLoopBack0


Configuration Scripts
                    ●     PE1
                          #
                          sysname PE1
                          #
                          vlan batch 100 200
                          #
                          ip vpn-instance vpna
                           ipv4-family
                            route-distinguisher 200:1
                            vpn-target 111:1 export-extcommunity
                            vpn-target 111:1 import-extcommunity
                          #
                          ip vpn-instance vpnb
                           ipv4-family
                            route-distinguisher 200:2
                            vpn-target 222:2 export-extcommunity
                            vpn-target 222:2 import-extcommunity
                          #
                          interface Vlanif100
                           ip binding vpn-instance vpna
                           ip address 10.5.1.1 255.255.255.0
                          #
                          interface Vlanif200
                           ip binding vpn-instance vpnb
                           ip address 10.5.2.1 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 100
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           port trunk allow-pass vlan 200
                          #
                          interface LoopBack1
                           ip address 2.2.2.2 255.255.255.255
                          #


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         101
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        ospf 100 vpn-instance vpna
                         area 0.0.0.0
                          network 10.5.1.0 0.0.0.255
                        #
                        ospf 200 vpn-instance vpnb
                         area 0.0.0.0
                          network 10.5.2.0 0.0.0.255
                        #
                        return
                    ●   MCE
                        #
                        sysname MCE
                        #
                        vlan batch 100 200 300 400
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:1
                        #
                        ip vpn-instance vpnb
                         ipv4-family
                          route-distinguisher 100:2
                        #
                        interface Vlanif100
                         ip binding vpn-instance vpna
                         ip address 10.5.1.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpnb
                         ip address 10.5.2.2 255.255.255.0
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpna
                         ip address 10.3.1.2 255.255.255.0
                        #
                        interface Vlanif400
                         ip binding vpn-instance vpnb
                         ip address 10.4.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        ospf 100 vpn-instance vpna
                         import-route rip 100
                         vpn-instance-capability simple
                         area 0.0.0.0
                          network 10.3.1.0 0.0.0.255
                          network 10.5.1.0 0.0.0.255
                        #
                        ospf 200 vpn-instance vpnb
                         import-route rip 200
                         vpn-instance-capability simple
                         area 0.0.0.0
                          network 10.4.1.0 0.0.0.255
                          network 10.5.2.0 0.0.0.255
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         102
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        rip 100 vpn-instance vpna
                         version 2
                         network 10.0.0.0
                         import-route ospf 100
                        #
                        rip 200 vpn-instance vpnb
                         version 2
                         network 10.0.0.0
                         import-route ospf 200
                        #
                        return

                    ●   DeviceA
                        #
                        sysname DeviceA
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.3.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        rip 100
                         version 2
                         network 10.0.0.0
                         network 3.0.0.0
                        #
                        return

                    ●   DeviceB
                        #
                        sysname DeviceB
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.4.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        rip 200
                         version 2
                         network 10.0.0.0
                         network 4.0.0.0
                        #
                        return



