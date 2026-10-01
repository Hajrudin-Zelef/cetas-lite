---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-175
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [25337, 25479]
sha256: acfd56b9efe86b40682ec8a32ce182e0a5cee9c34cd3b6bdfa9d48cc7024b010
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 3 Configure an IPv6-address-family-enabled VPN instance on each PE and bind the
                interface connecting a PE to a CE to the VPN instance on that PE.
                          NOTE

                        The import VPN target list of a VPN instance on the Hub-PE must contain the export VPN
                        targets of all Spoke-PEs.
                        The export VPN target list of the other VPN instance on the Hub-PE must contain the
                        import VPN targets of all Spoke-PEs.

                    # Configure Spoke-PE1.
                    <Spoke-PE1> system-view
                    [Spoke-PE1] ip vpn-instance vpna
                    [Spoke-PE1-vpn-instance-vpna] ipv6-family
                    [Spoke-PE1-vpn-instance-vpna-af-ipv6] route-distinguisher 100:1
                    [Spoke-PE1-vpn-instance-vpna-af-ipv6] vpn-target 100:1 export-extcommunity
                    [Spoke-PE1-vpn-instance-vpna-af-ipv6] vpn-target 200:1 import-extcommunity
                    [Spoke-PE1-vpn-instance-vpna-af-ipv6] quit
                    [Spoke-PE1] interface Vlanif100
                    [Spoke-PE1-Vlanif100] ip binding vpn-instance vpna
                    [Spoke-PE1-Vlanif100] ipv6 enable
                    [Spoke-PE1-Vlanif100] ipv6 address 2001:db8:1::2 64
                    [Spoke-PE1-Vlanif100] quit

                    # Configure Spoke-PE2.
                    <Spoke-PE2> system-view
                    [Spoke-PE2] ip vpn-instance vpna
                    [Spoke-PE2-vpn-instance-vpna] ipv6-family
                    [Spoke-PE2-vpn-instance-vpna-af-ipv6] route-distinguisher 100:3
                    [Spoke-PE2-vpn-instance-vpna-af-ipv6] vpn-target 100:1 export-extcommunity
                    [Spoke-PE2-vpn-instance-vpna-af-ipv6] vpn-target 200:1 import-extcommunity
                    [Spoke-PE2-vpn-instance-vpna-af-ipv6] quit
                    [Spoke-PE2] interface Vlanif100
                    [Spoke-PE2-Vlanif100] ip binding vpn-instance vpna
                    [Spoke-PE2-Vlanif100] ipv6 enable
                    [Spoke-PE2-Vlanif100] ipv6 address 2001:db8:2::2 64
                    [Spoke-PE2-Vlanif100] quit


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   401
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration


                    # Configure the Hub-PE.
                    <Hub-PE> system-view
                    [Hub-PE] ip vpn-instance vpn_in
                    [Hub-PE-vpn-instance-vpn_in] ipv6-family
                    [Hub-PE-vpn-instance-vpn_in-af-ipv6] route-distinguisher 100:21
                    [Hub-PE-vpn-instance-vpn_in-af-ipv6] vpn-target 100:1 import-extcommunity
                    [Hub-PE-vpn-instance-vpn_in-af-ipv6] quit
                    [Hub-PE-vpn-instance-vpn_in] quit
                    [Hub-PE] ip vpn-instance vpn_out
                    [Hub-PE-vpn-instance-vpn_out] ipv6-family
                    [Hub-PE-vpn-instance-vpn_out-af-ipv6] route-distinguisher 100:22
                    [Hub-PE-vpn-instance-vpn_out-af-ipv6] vpn-target 200:1 export-extcommunity
                    [Hub-PE-vpn-instance-vpn_out-af-ipv6] quit
                    [Hub-PE-vpn-instance-vpn_out] quit
                    [Hub-PE] interface Vlanif300
                    [Hub-PE-Vlanif300] ip binding vpn-instance vpn_in
                    [Hub-PE-Vlanif300] ipv6 enable
                    [Hub-PE-Vlanif300] ipv6 address 2001:db8:3::2 64
                    [Hub-PE-Vlanif300] quit
                    [Hub-PE] interface Vlanif400
                    [Hub-PE-Vlanif400] ip binding vpn-instance vpn_out
                    [Hub-PE-Vlanif400] ipv6 enable
                    [Hub-PE-Vlanif400] ipv6 address 2001:db8:4::2 64
                    [Hub-PE-Vlanif400] quit

                    # Configure IP addresses for interfaces on CEs, as shown in Figure 4-11. For
                    detailed configurations, see Configuration Scripts.

                    After completing the configuration, run the display ip vpn-instance verbose
                    command on PEs to check VPN instance configuration. Each PE can ping its
                    connected CEs through the ping ipv6 vpn-instance vpn-name ipv6-address
                    command.

         Step 4 Establish EBGP peer relationships between PEs and CEs to import VPN routes.
                          NOTE

                        Configure the Hub-PE to allow AS numbers to be repeated once in the AS_Path attribute, so
                        that it can receive the routes advertised by the Hub-CE.
                        You do not need to configure the Spoke-PEs to allow AS numbers to be repeated once,
                        because the device does not check the AS_Path attributes in routes received from IBGP
                        peers.

                    # Configure Spoke-CE1.
                    [Spoke-CE1] interface loopback 1
                    [Spoke-CE1-Loopback1] ipv6 enable
                    [Spoke-CE1-Loopback1] ipv6 address 2001:db8:11::1 128
                    [Spoke-CE1-Loopback1] quit
                    [Spoke-CE1] bgp 65410
                    [Spoke-CE1-bgp] ipv6-family unicast
                    [Spoke-CE1-bgp-af-ipv6] peer 2001:db8:1::2 as-number 100
                    [Spoke-CE1-bgp-af-ipv6] network 2001:db8:11::1 128
                    [Spoke-CE1-bgp-af-ipv6] quit
                    [Spoke-CE1-bgp] quit

                    # Configure Spoke-PE1.
                    [Spoke-PE1] bgp 100
                    [Spoke-PE1-bgp] ipv6-family vpn-instance vpna
                    [Spoke-PE1-bgp-6-vpna] peer 2001:db8:1::1 as-number 65410
                    [Spoke-PE1-bgp-6-vpna] quit
                    [Spoke-PE1-bgp] quit

                    # Configure Spoke-CE2.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     402
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration

                    [Spoke-CE2] interface loopback 1
                    [Spoke-CE2-Loopback1] ipv6 enable
                    [Spoke-CE2-Loopback1] ipv6 address 2001:db8:12::2 128
                    [Spoke-CE2-Loopback1] quit
                    [Spoke-CE2] bgp 65420
                    [Spoke-CE2-bgp] ipv6-family unicast
                    [Spoke-CE2-bgp-af-ipv6] peer 2001:db8:2::2 as-number 100
                    [Spoke-CE2-bgp-af-ipv6] network 2001:db8:12::2 128
                    [Spoke-CE2-bgp-af-ipv6] quit
                    [Spoke-CE2-bgp] quit

                    # Configure Spoke-PE2.
                    [Spoke-PE2] bgp 100
                    [Spoke-PE2-bgp] ipv6-family vpn-instance vpna
                    [Spoke-PE2-bgp-6-vpna] peer 2001:db8:2::1 as-number 65420
                    [Spoke-PE2-bgp-6-vpna] quit
                    [Spoke-PE2-bgp] quit

                    # Configure the Hub-CE.
                    [Hub-CE] interface loopback 1
                    [Hub-CE-Loopback1] ipv6 enable
                    [Hub-CE-Loopback1] ipv6 address 2001:db8:13::3 128
                    [Hub-CE-Loopback1] quit
                    [Hub-CE] bgp 65430
                    [Hub-CE-bgp] ipv6-family unicast
                    [Hub-CE-bgp-af-ipv6] peer 2001:db8:3::2 as-number 100
                    [Hub-CE-bgp-af-ipv6] peer 2001:db8:4::2 as-number 100
                    [Hub-CE-bgp-af-ipv6] network 2001:db8:13::3 128
                    [Hub-CE-bgp-af-ipv6] quit
                    [Hub-CE-bgp] quit

