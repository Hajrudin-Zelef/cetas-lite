---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-178
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [25782, 25974]
sha256: 30dbdccd5670f12dd8b6363b9c819779116e85add7c28cf55bffed10143d988b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Hub-CE
                        #
                        sysname Hub-CE
                        #
                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:db8:3::1/64
                        #
                        interface Vlanif200
                         ipv6 enable
                         ipv6 address 2001:db8:4::1/64
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
                         ipv6 enable
                         ipv6 address 2001:db8:13::3/128
                        #



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         407
VPN Configuration
VPN Configuration                                                            4 IPv6 L3VPN Configuration

                        bgp 65430
                         router-id 2.2.2.2
                         peer 2001:db8:3::2 as-number 100
                         peer 2001:db8:4::2 as-number 100
                         #
                         ipv6-family unicast
                          network 2001:db8:13::3 128
                          peer 2001:db8:3::2 enable
                          peer 2001:db8:4::2 enable
                        #
                        return

                    ●   Hub-PE
                        #
                        sysname Hub-PE
                        #
                        vlan batch 100 200 300 400
                        #
                        ip vpn-instance vpn_in
                         ipv6-family
                          route-distinguisher 100:21
                          vpn-target 100:1 import-extcommunity
                        #
                        ip vpn-instance vpn_out
                         ipv6-family
                          route-distinguisher 100:22
                          vpn-target 200:1 export-extcommunity
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.2.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpn_in
                         ipv6 enable
                         ipv6 address 2001:db8:3::2/64
                        #
                        interface Vlanif400
                         ip binding vpn-instance vpn_out
                         ipv6 enable
                         ipv6 address 2001:db8:4::2/64
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


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         408
VPN Configuration
VPN Configuration                                                                4 IPv6 L3VPN Configuration

                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                         #
                         ipv6-family vpn-instance vpn_in
                          peer 2001:db8:3::1 as-number 65430
                         #
                         ipv6-family vpn-instance vpn_out
                          peer 2001:db8:4::1 as-number 65430
                          peer 2001:db8:4::1 allow-as-loop
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.2.1.0 0.0.0.255
                          network 10.1.1.0 0.0.0.255
                        #
                        return



4.15 Maintaining IPv6 L3VPN

4.15.1 Configuring an Alarm Threshold for the Number of
Routes in an IPv6 VPN Instance

Context
                    As the number of access hosts increases, the number of routes stored on the
                    control plane also increases, consuming a significant amount of memory
                    resources. To better monitor the memory usage in this case and ensure the device
                    does not restart due to insufficient memory, configure an alarm threshold for the
                    number of routes in a VPN instance. When the number of routes exceeds this
                    threshold, a user log is generated. Conversely, when the number of routes falls
                    below the recovery percentage, a recovery log is reported.

                         NOTE

                        Only the S6780-H, S6750-H, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S, S5755-S,
                        S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2 series
                        support this function.


Procedure
         Step 1 Enter the system view.
                    system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                               409
VPN Configuration
VPN Configuration                                                                                  4 IPv6 L3VPN Configuration


         Step 2 Set an alarm threshold and a recovery percentage for the number of routes in an
                IPv6 VPN instance.
                    alarm-threshold route route-number [ recovery-percentage percentage ] ipv6 vpn-instance vpn-instance-
                    name

                    ----End

4.15.2 Monitoring the Running Status of IPv6 L3VPN Services
Procedure
                    ●      Run the display ip vpn-instance [ verbose ] [ vpn-instance-name ] command
                           to check VPN instance information.
                    ----End

4.15.3 Checking IPv6 L3VPN Connectivity and Reachability

