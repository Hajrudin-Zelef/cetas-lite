---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-108
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [15078, 15222]
sha256: f8375709a9f4b7f80748d8fbf1e78166cb080c097a381affaf5e362ca87fa890
---

                    # Configure Spoke-PE2.
                    <Spoke-PE2> system-view
                    [Spoke-PE2] ip vpn-instance vpna
                    [Spoke-PE2-vpn-instance-vpna] ipv4-family
                    [Spoke-PE2-vpn-instance-vpna-af-ipv4] route-distinguisher 100:3
                    [Spoke-PE2-vpn-instance-vpna-af-ipv4] vpn-target 100:1 export-extcommunity
                    [Spoke-PE2-vpn-instance-vpna-af-ipv4] vpn-target 200:1 import-extcommunity
                    [Spoke-PE2-vpn-instance-vpna-af-ipv4] quit
                    [Spoke-PE2-vpn-instance-vpna] quit
                    [Spoke-PE2] interface Vlanif 100
                    [Spoke-PE2-Vlanif100] ip binding vpn-instance vpna
                    [Spoke-PE2-Vlanif100] ip address 10.4.1.2 24
                    [Spoke-PE2-Vlanif100] quit

                    # Configure the Hub-PE.
                    <Hub-PE> system-view
                    [Hub-PE] ip vpn-instance vpn_in
                    [Hub-PE-vpn-instance-vpn_in] ipv4-family
                    [Hub-PE-vpn-instance-vpn_in-af-ipv4] route-distinguisher 100:21
                    [Hub-PE-vpn-instance-vpn_in-af-ipv4] vpn-target 100:1 import-extcommunity
                    [Hub-PE-vpn-instance-vpn_in-af-ipv4] quit
                    [Hub-PE-vpn-instance-vpn_in] quit
                    [Hub-PE] ip vpn-instance vpn_out
                    [Hub-PE-vpn-instance-vpn_out] ipv4-family
                    [Hub-PE-vpn-instance-vpn_out-af-ipv4] route-distinguisher 100:22
                    [Hub-PE-vpn-instance-vpn_out-af-ipv4] vpn-target 200:1 export-extcommunity
                    [Hub-PE-vpn-instance-vpn_out-af-ipv4] quit
                    [Hub-PE-vpn-instance-vpn_out] quit
                    [Hub-PE] interface Vlanif 300
                    [Hub-PE-Vlanif300] ip binding vpn-instance vpn_in
                    [Hub-PE-Vlanif300] ip address 10.2.1.2 24
                    [Hub-PE-Vlanif300] quit
                    [Hub-PE] interface Vlanif 400
                    [Hub-PE-Vlanif400] ip binding vpn-instance vpn_out
                    [Hub-PE-Vlanif400] ip address 10.3.1.2 24
                    [Hub-PE-Vlanif400] quit

                    # Configure IP addresses for interfaces on CEs, as shown in Figure 3-48. For
                    detailed configurations, see Configuration Scripts.
                    After completing the configuration, run the display ip vpn-instance verbose
                    command on each PE to check the VPN instance configuration. Each PE can ping
                    its connected CEs using the ping -vpn-instance vpn-name ip-address command.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   240
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration


                          NOTE

                        If a PE has multiple interfaces bound to the same VPN instance, use the -a source-ip-
                        address parameter to specify a source IP address when running the ping -vpn-instance
                        vpn-instance-name -a source-ip-address dest-ip-address command to ping the CE
                        connected to the remote PE. If the source IP address is not specified, the ping operation
                        may fail.

         Step 4 Establish EBGP peer relationships between PEs and CEs to import VPN routes.
                          NOTE

                        Configure the Hub-PE to allow AS numbers to be repeated once in the AS_Path attribute, so
                        that it can receive the routes advertised by the Hub-CE.
                        You do not need to configure the Spoke-PEs to allow AS numbers to be repeated once,
                        because the device does not check the AS_Path attributes in routes received from IBGP
                        peers.

                    # Configure Spoke-CE1.
                    [Spoke-CE1] interface loopback 1
                    [Spoke-CE1-Loopback1] ip address 11.11.11.11 32
                    [Spoke-CE1-Loopback1] quit
                    [Spoke-CE1] bgp 65410
                    [Spoke-CE1-bgp] peer 10.1.1.2 as-number 100
                    [Spoke-CE1-bgp] network 11.11.11.11 32
                    [Spoke-CE1-bgp] quit

                    # Configure Spoke-PE1.
                    [Spoke-PE1] bgp 100
                    [Spoke-PE1-bgp] ipv4-family vpn-instance vpna
                    [Spoke-PE1-bgp-vpna] peer 10.1.1.1 as-number 65410
                    [Spoke-PE1-bgp-vpna] quit
                    [Spoke-PE1-bgp] quit

                    # Configure Spoke-CE2.
                    [Spoke-CE2] interface loopback 1
                    [Spoke-CE2-Loopback1] ip address 22.22.22.22 32
                    [Spoke-CE2-Loopback1] quit
                    [Spoke-CE2] bgp 65420
                    [Spoke-CE2-bgp] peer 10.4.1.2 as-number 100
                    [Spoke-CE2-bgp] network 22.22.22.22 32
                    [Spoke-CE2-bgp] quit

                    # Configure Spoke-PE2.
                    [Spoke-PE2] bgp 100
                    [Spoke-PE2-bgp] ipv4-family vpn-instance vpna
                    [Spoke-PE2-bgp-vpna] peer 10.4.1.1 as-number 65420
                    [Spoke-PE2-bgp-vpna] quit
                    [Spoke-PE2-bgp] quit

                    # Configure the Hub-CE.
                    [Hub-CE] interface loopback 1
                    [Hub-CE-Loopback1] ip address 33.33.33.33 32
                    [Hub-CE-Loopback1] quit
                    [Hub-CE] bgp 65430
                    [Hub-CE-bgp] peer 10.2.1.2 as-number 100
                    [Hub-CE-bgp] peer 10.3.1.2 as-number 100
                    [Hub-CE-bgp] network 33.33.33.33 32
                    [Hub-CE-bgp] quit

                    # Configure the Hub-PE.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         241
VPN Configuration
VPN Configuration                                                                        3 IPv4 L3VPN Configuration

                    [Hub-PE] bgp 100
                    [Hub-PE-bgp] ipv4-family vpn-instance vpn_in
                    [Hub-PE-bgp-vpn_in] peer 10.2.1.1 as-number 65430
                    [Hub-PE-bgp-vpn_in] quit
                    [Hub-PE-bgp] ipv4-family vpn-instance vpn_out
                    [Hub-PE-bgp-vpn_out] peer 10.3.1.1 as-number 65430
                    [Hub-PE-bgp-vpn_out] peer 10.3.1.1 allow-as-loop 1
                    [Hub-PE-bgp-vpn_out] quit
                    [Hub-PE-bgp] quit

                    After completing the configuration, run the display bgp vpnv4 all peer command
                    on each PE. The command output shows that BGP peer relationships have been
                    established between the PEs and CEs and are in Established state.
         Step 5 Establish MP-IBGP peer relationships between the PEs.
                    # Configure Spoke-PE1.
                    [Spoke-PE1] bgp 100
                    [Spoke-PE1-bgp] peer 2.2.2.9 as-number 100
                    [Spoke-PE1-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [Spoke-PE1-bgp] ipv4-family vpnv4
                    [Spoke-PE1-bgp-af-vpnv4] peer 2.2.2.9 enable
                    [Spoke-PE1-bgp-af-vpnv4] quit

                    # Configure Spoke-PE2.
                    [Spoke-PE2] bgp 100
                    [Spoke-PE2-bgp] peer 2.2.2.9 as-number 100
                    [Spoke-PE2-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [Spoke-PE2-bgp] ipv4-family vpnv4
                    [Spoke-PE2-bgp-af-vpnv4] peer 2.2.2.9 enable
                    [Spoke-PE2-bgp-af-vpnv4] quit

