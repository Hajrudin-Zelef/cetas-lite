---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-144
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [20406, 20605]
sha256: efc4ba45fed4e544923e521aeae964c791f9001a420aa43517b7321ce601da7d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpna
                         ipv6-family
                          route-distinguisher 200:1
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                        #
                        ip vpn-instance vpnb
                         ipv6-family
                          route-distinguisher 200:2
                          vpn-target 222:2 export-extcommunity
                          vpn-target 222:2 import-extcommunity
                        #
                        ospfv3 100 vpn-instance vpna
                         router-id 10.5.5.5
                         area 0.0.0.1
                        #
                        ospfv3 200 vpn-instance vpnb
                         router-id 10.6.6.6
                         area 0.0.0.2
                        #
                        interface Vlanif100
                         ip binding vpn-instance vpna
                         ipv6 enable
                         ipv6 address 2001:DB8:8::1/64
                         ospfv3 100 area 0.0.0.1 instance 1
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpnb
                         ipv6 enable
                         ipv6 address 2001:DB8:9::1/64
                         ospfv3 200 area 0.0.0.2 instance 2
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        return

                    ●   MCE
                        #
                        sysname MCE
                        #
                        vlan batch 100 200 300 400
                        #
                        ip vpn-instance vpna
                         ipv6-family
                          route-distinguisher 100:1
                        #
                        ip vpn-instance vpnb
                         ipv6-family
                          route-distinguisher 100:2
                        #
                        ospfv3 100 vpn-instance vpna
                         router-id 10.7.7.7
                         vpn-instance-capability simple
                         import-route ripng 100
                         area 0.0.0.1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         324
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                        #
                        ospfv3 200 vpn-instance vpnb
                         router-id 10.8.8.8
                         vpn-instance-capability simple
                         import-route ripng 200
                         area 0.0.0.2
                        #
                        interface Vlanif100
                         ip binding vpn-instance vpna
                         ipv6 enable
                         ipv6 address 2001:DB8:8::2/64
                         ospfv3 100 area 0.0.0.1 instance 1
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpnb
                         ipv6 enable
                         ipv6 address 2001:DB8:9::2/64
                         ospfv3 200 area 0.0.0.2 instance 2
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpna
                         ipv6 enable
                         ipv6 address 2001:DB8:3::2/64
                         ripng 100 enable
                        #
                        interface Vlanif400
                         ip binding vpn-instance vpnb
                         ipv6 enable
                         ipv6 address 2001:DB8:4::2/64
                         ripng 200 enable
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
                        ripng 100 vpn-instance vpna
                         import-route ospfv3 100
                        #
                        ripng 200 vpn-instance vpnb
                         import-route ospfv3 200
                        #
                        return
                    ●   DeviceA
                        #
                        sysname DeviceA
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:DB8:3::1/64
                         ripng 100 enable
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         325
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                        interface LoopBack1
                         ipv6 enable
                         ipv6 address 2001:DB8:13::3/128
                         ripng 100 enable
                        #
                        ripng 100
                        #
                        return

                    ●   DeviceB
                        #
                        sysname DeviceB
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:DB8:4::1/64
                         ripng 200 enable
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ipv6 enable
                         ipv6 address 2001:DB8:14::4/128
                         ripng 200 enable
                        #
                        ripng 200
                        #
                        return



4.10 Configuring IPv6 L3VPN FRR

4.10.1 Configuring IPv6 VPN FRR

Context
                    IPv6 VPN FRR applies to scenarios where IPv6 services are sensitive to packet loss
                    and delay and the requirements for VPNs transmitting IPv6 services are high. If a
                    fault occurs on the backbone network after IPv6 VPN FRR is enabled, IPv6 VPN
                    services can be quickly switched to another link, ensuring service continuity.


Prerequisites
                    Before configuring IPv6 VPN FRR, you have completed the following tasks:

