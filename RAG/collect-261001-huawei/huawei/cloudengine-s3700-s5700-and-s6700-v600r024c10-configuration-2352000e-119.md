---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-119
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [16762, 16927]
sha256: c18bc975b651754b3ea6ad2cc77a31646e0edafa2ffb8a746453fd8b46c8599b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Hub-PE
                        #
                        sysname Hub-PE
                        #
                        vlan batch 100 200 300 400
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:3
                          vpn-target 100:1 export-extcommunity
                          vpn-target 200:1 import-extcommunity
                        #
                        ip vpn-instance vpnhub
                         ipv4-family
                          route-distinguisher 100:21
                          export route-policy policy_in
                          apply-label per-route pop-go
                          vpn-target 200:1 export-extcommunity
                          vpn-target 100:1 import-extcommunity
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 20.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 11.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpnhub
                         ip address 10.2.1.2 255.255.255.0
                        #
                        interface Vlanif400
                         ip binding vpn-instance vpna
                         ip address 10.3.1.2 255.255.255.0
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         265
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration

                        #
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
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                        #
                         ipv4-family vpn-instance vpna
                          peer 10.3.1.1 as-number 65440
                         ipv4-family vpn-instance vpnhub
                          peer 10.2.1.1 as-number 65430
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 20.1.1.0 0.0.0.255
                          network 11.1.1.0 0.0.0.255
                        #
                        ip ip-prefix defaultip index 10 permit 0.0.0.0 0
                        #
                        route-policy policy_in permit node 1
                         if-match ip-prefix defaultip
                        #
                        route-policy policy_in deny node 2
                        #
                        return



3.16 Maintaining IPv4 L3VPN

3.16.1 Setting an Alarm Threshold for the Number of Routes
in an IPv4 VPN Instance

Context
                    As the number of access hosts increases, the number of routes stored on the
                    control plane also increases, consuming a significant number of memory
                    resources. To better monitor the memory usage in this case and ensure that the
                    device does not restart due to insufficient memory, configure an alarm threshold
                    for the number of routes in a VPN instance. When the number of routes exceeds
                    this threshold, a user log is generated. Conversely, when the number of routes falls
                    below the recovery percentage, a recovery log is reported.

                         NOTE

                        This function is supported only by the S6780-H series, S6750-H series, S6730-H-V2 series,
                        S5732-H-V2 series, S6750E-S, S6750-S, S5755-S, S5755E-H, S5755-H series, S5735I-S-V2
                        series, S5735I-H-V2 series, S5735R-S-V2, and S5735-S-V2 series.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    266
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Set an alarm threshold and a recovery percentage for the number of routes in an
                IPv4 VPN instance.
                    alarm-threshold route route-number [ recovery-percentage percentage ] ipv4 vpn-instance vpn-instance-
                    name

                    ----End

3.16.2 Monitoring the Running Status of IPv4 L3VPN Services

Procedure
                    ●    Run the display ip vpn-instance [ verbose ] [ vpn-instance-name ] command
                         to check VPN instance information.

                    ----End

3.16.3 Checking IPv4 L3VPN Connectivity and Reachability

Procedure
                    ●    Run the ping command to check the connectivity of the IPv4 network from
                         the source to the destination. Perform either of the following operations
                         according to the displayed detailed or brief information.
                         –    Run the following command to display detailed information:
                              ping [ ip ] { [ -c count | -i { interface-name | interface-type interface-number } | -nexthop
                              nexthop-address | { -range [ min min-value | max max-value | step step-value ] * | -s
                              packetsize } | -t timeout | -m time | -a source-ip-address | -h ttl-value | -p pattern | { -tos tos-
                              value | -dscp dscp-value } | { -f | ignore-mtu } | -q | -r | -vpn-instance vpn-instance-name | -v | -
                              name | -system-time | -ri | -8021p 8021p-value | -detail ] * host [ ip-forwarding ] }

