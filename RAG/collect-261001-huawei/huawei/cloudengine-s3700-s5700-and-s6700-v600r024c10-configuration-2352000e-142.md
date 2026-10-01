---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-142
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [20132, 20278]
sha256: 6d5668d1276b1762543e27a47fce6ceb37070395942b9b0c39925718ef0591a8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Procedure
         Step 1 Configure VPN instances on the MCE and PE1 and bind interfaces to the VPN
                instances.
                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv6-family
                    [PE1-vpn-instance-vpna-af-ipv6] route-distinguisher 200:1
                    [PE1-vpn-instance-vpna-af-ipv6] vpn-target 111:1 both
                    [PE1-vpn-instance-vpna-af-ipv6] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE1] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv6-family
                    [PE1-vpn-instance-vpnb-af-ipv6] route-distinguisher 200:2
                    [PE1-vpn-instance-vpnb-af-ipv6] vpn-target 222:2 both
                    [PE1-vpn-instance-vpnb-af-ipv6] quit
                    [PE1-vpn-instance-vpnb] quit
                    [PE1] vlan batch 100 200
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] port trunk allow-pass vlan 100
                    [PE1-10GE1/0/1] quit
                    [PE1] interface Vlanif 100
                    [PE1-Vlanif100] ip binding vpn-instance vpna
                    [PE1-Vlanif100] ipv6 enable
                    [PE1-Vlanif100] ipv6 address 2001:DB8:8::1 64
                    [PE1-Vlanif100] quit
                    [PE1] interface 10GE1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 200
                    [PE1-10GE1/0/2] quit
                    [PE1] interface Vlanif 200
                    [PE1-Vlanif200] ip binding vpn-instance vpnb
                    [PE1-Vlanif200] ipv6 enable
                    [PE1-Vlanif200] ipv6 address 2001:DB8:9::1 64
                    [PE1-Vlanif200] quit

                    # Configure the MCE.
                    <HUAWEI> system-view
                    [HUAWEI] sysname MCE
                    [MCE] ip vpn-instance vpna
                    [MCE-vpn-instance-vpna] ipv6-family
                    [MCE-vpn-instance-vpna-af-ipv6] route-distinguisher 100:1
                    [MCE-vpn-instance-vpna-af-ipv6] quit
                    [MCE-vpn-instance-vpna] quit
                    [MCE] ip vpn-instance vpnb


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          320
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration

                    [MCE-vpn-instance-vpnb] ipv6-family
                    [MCE-vpn-instance-vpnb-af-ipv6] route-distinguisher 100:2
                    [MCE-vpn-instance-vpnb-af-ipv6] quit
                    [MCE-vpn-instance-vpnb] quit
                    [MCE] vlan batch 100 200 300 400
                    [MCE] interface 10ge 1/0/1
                    [MCE-10GE1/0/1] port link-type trunk
                    [MCE-10GE1/0/1] port trunk allow-pass vlan 100
                    [MCE-10GE1/0/1] quit
                    [MCE] interface Vlanif 100
                    [MCE-Vlanif100] ipv6 enable
                    [MCE-Vlanif100] ip binding vpn-instance vpna
                    [MCE-Vlanif100] ipv6 address 2001:DB8:8::2 64
                    [MCE-Vlanif100] quit
                    [MCE] interface 10GE1/0/2
                    [MCE-10GE1/0/2] port link-type trunk
                    [MCE-10GE1/0/2] port trunk allow-pass vlan 200
                    [MCE-10GE1/0/2] quit
                    [MCE] interface Vlanif 200
                    [MCE-Vlanif200] ipv6 enable
                    [MCE-Vlanif200] ip binding vpn-instance vpnb
                    [MCE-Vlanif200] ipv6 address 2001:DB8:9::2 64
                    [MCE-Vlanif200] quit
                    [MCE] interface 10GE1/0/3
                    [MCE-10GE1/0/3] port link-type trunk
                    [MCE-10GE1/0/3] port trunk allow-pass vlan 300
                    [MCE-10GE1/0/3] quit
                    [MCE] interface Vlanif 300
                    [MCE-Vlanif300] ipv6 enable
                    [MCE-Vlanif300] ip binding vpn-instance vpna
                    [MCE-Vlanif300] ipv6 address 2001:DB8:3::2 64
                    [MCE-Vlanif300] quit
                    [MCE] interface 10GE1/0/4
                    [MCE-10GE1/0/4] port link-type trunk
                    [MCE-10GE1/0/4] port trunk allow-pass vlan 400
                    [MCE-10GE1/0/4] quit
                    [MCE] interface Vlanif 400
                    [MCE-Vlanif400] ipv6 enable
                    [MCE-Vlanif400] ip binding vpn-instance vpnb
                    [MCE-Vlanif400] ipv6 address 2001:DB8:4::2 64
                    [MCE-Vlanif400] quit

         Step 2 Configure OSPFv3 multi-instance on PE1 and the MCE.
                    # Configure PE1.
                    [PE1] ospfv3 100 vpn-instance vpna
                    [PE1-ospfv3-100] router-id 10.5.5.5
                    [PE1-ospfv3-100] quit
                    [PE1] interface Vlanif100
                    [PE1-Vlanif100] ospfv3 100 area 1 instance 1
                    [PE1-Vlanif100] quit
                    [PE1] ospfv3 200 vpn-instance vpnb
                    [PE1-ospfv3-200] router-id 10.6.6.6
                    [PE1-ospfv3-200] quit
                    [PE1] interface Vlanif200
                    [PE1-Vlanif200] ospfv3 200 area 2 instance 2
                    [PE1-Vlanif200] quit

                    # Configure the MCE.
                    [MCE] ospfv3 100 vpn-instance vpna
                    [MCE-ospfv3-100] router-id 10.7.7.7
                    [MCE-ospfv3-100] quit
                    [MCE] interface Vlanif100
                    [MCE-Vlanif100] ospfv3 100 area 1 instance 1
                    [MCE-Vlanif100] quit
                    [MCE] ospfv3 200 vpn-instance vpnb
                    [MCE-ospfv3-200] router-id 10.8.8.8
                    [MCE-ospfv3-200] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          321
VPN Configuration
VPN Configuration                                                              4 IPv6 L3VPN Configuration

                    [MCE] interface Vlanif200
                    [MCE-Vlanif200] ospfv3 200 area 2 instance 2
                    [MCE-Vlanif200] quit

         Step 3 RIPng must be configured on the MCE to import VPN routes from Site 1 and Site
                2.
                    # Configure the MCE.
                    [MCE] ripng 100 vpn-instance vpna
                    [MCE-ripng-100] import-route ospfv3 100
                    [MCE-ripng-100] quit
                    [MCE] interface Vlanif300
                    [MCE-Vlanif300] ripng 100 enable
                    [MCE-Vlanif300] quit
                    [MCE] ripng 200 vpn-instance vpnb
                    [MCE-ripng-200] import-route ospfv3 200
                    [MCE-ripng-200] quit
                    [MCE] interface Vlanif400
                    [MCE-Vlanif400] ripng 200 enable

