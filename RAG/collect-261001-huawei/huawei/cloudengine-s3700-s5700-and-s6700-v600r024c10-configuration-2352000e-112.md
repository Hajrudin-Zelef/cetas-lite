---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-112
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [15671, 15810]
sha256: 2c51b3ccd02ac2a39817b5c4fb96f0515dcc7e64299a7bb02eb9169f8e87d756
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Procedure
         Step 1 Configure IGP on the backbone network for the Hub-PE and Spoke-PEs to
                communicate.
                    OSPF is used as IGP in this example. For detailed configurations, see Configuration
                    Scripts.
                    After the configuration is complete, OSPF neighbor relationships are established
                    between the Hub-PE and Spoke-PEs. Run the display ospf peer command. The
                    command output shows that the neighbor status is Full. Run the display ip
                    routing-table command. The command output shows that the Hub-PE and
                    Spoke-PEs have learned the routes to each other's loopback interface.
         Step 2 Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                backbone network.
                    For detailed configurations, see Configuration Scripts.
                    After the configuration is complete, LDP peer relationships are established
                    between the Hub-PE and Spoke-PEs. Run the display mpls ldp session command
                    on each device. The command output shows that Session State is Operational.
         Step 3 Configure a VPN instance on each PE, enable the IPv4 address family for the
                instance, and bind the interface that connects each PE to a CE to the VPN instance
                on that PE.
                    # Configure Spoke-PE1.
                    <Spoke-PE1> system-view
                    [Spoke-PE1] ip vpn-instance vpna
                    [Spoke-PE1-vpn-instance-vpna] ipv4-family
                    [Spoke-PE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [Spoke-PE1-vpn-instance-vpna-af-ipv4] vpn-target 100:1 export-extcommunity
                    [Spoke-PE1-vpn-instance-vpna-af-ipv4] vpn-target 200:1 import-extcommunity
                    [Spoke-PE1-vpn-instance-vpna-af-ipv4] quit
                    [Spoke-PE1-vpn-instance-vpna] quit
                    [Spoke-PE1] interface Vlanif 100
                    [Spoke-PE1-Vlanif100] ip binding vpn-instance vpna
                    [Spoke-PE1-Vlanif100] ip address 10.1.1.2 24
                    [Spoke-PE1-Vlanif100] quit

                    # Configure Spoke-PE2.
                    <Spoke-PE2> system-view
                    [Spoke-PE2] ip vpn-instance vpna
                    [Spoke-PE2-vpn-instance-vpna] ipv4-family
                    [Spoke-PE2-vpn-instance-vpna-af-ipv4] route-distinguisher 100:3
                    [Spoke-PE2-vpn-instance-vpna-af-ipv4] vpn-target 100:1 export-extcommunity
                    [Spoke-PE2-vpn-instance-vpna-af-ipv4] vpn-target 200:1 import-extcommunity


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  249
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration

                    [Spoke-PE2-vpn-instance-vpna-af-ipv4] quit
                    [Spoke-PE2-vpn-instance-vpna] quit
                    [Spoke-PE2] interface Vlanif 100
                    [Spoke-PE2-Vlanif100] ip binding vpn-instance vpna
                    [Spoke-PE2-Vlanif100] ip address 10.4.1.2 24
                    [Spoke-PE2-Vlanif100] quit

                    # Configure the Hub-PE.
                    <Hub-PE> system-view
                    [Hub-PE] ip ip-prefix defaultip index 10 permit 0.0.0.0 0
                    [Hub-PE] route-policy policy_in permit node 1
                    [Hub-PE-route-policy] if-match ip-prefix defaultip
                    [Hub-PE-route-policy] quit
                    [Hub-PE] route-policy policy_in deny node 2
                    [Hub-PE-route-policy] quit
                    [Hub-PE] ip vpn-instance vpnhub
                    [Hub-PE-vpn-instance-vpnhub] ipv4-family
                    [Hub-PE-vpn-instance-vpnhub-af-ipv4] route-distinguisher 100:21
                    [Hub-PE-vpn-instance-vpnhub-af-ipv4] export route-policy policy_in
                    [Hub-PE-vpn-instance-vpnhub-af-ipv4] vpn-target 200:1 export-extcommunity
                    [Hub-PE-vpn-instance-vpnhub-af-ipv4] vpn-target 100:1 import-extcommunity
                    [Hub-PE-vpn-instance-vpnhub-af-ipv4] apply-label per-route pop-go
                    [Hub-PE-vpn-instance-vpnhub-af-ipv4] quit
                    [Hub-PE-vpn-instance-vpnhub] quit
                    [Hub-PE] interface Vlanif 300
                    [Hub-PE-Vlanif300] ip binding vpn-instance vpn_in
                    [Hub-PE-Vlanif300] ip address 10.2.1.2 24
                    [Hub-PE-Vlanif300] quit

                    # Configure IP addresses for interfaces on CEs, as shown in Figure 3-49. For
                    detailed configurations, see Configuration Scripts.
                    After completing the configuration, run the display ip vpn-instance verbose
                    command on each PE to check the VPN instance configurations. Each PE can ping
                    its connected CEs using the ping -vpn-instance vpn-name ip-address command.

                          NOTE

                        If a PE has multiple interfaces bound to the same VPN instance, use the -a source-ip-
                        address parameter to specify a source IP address when running the ping -vpn-instance
                        vpn-instance-name -a source-ip-address dest-ip-address command to ping the CE
                        connected to the remote PE. If the source IP address is not specified, the ping operation
                        may fail.

         Step 4 Establish EBGP peer relationships between PEs and CEs to import VPN routes.
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


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        250
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration

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
                    [Hub-CE-bgp] peer 10.2.1.2 default-route-advertise
                    [Hub-CE-bgp] network 33.33.33.33 32
                    [Hub-CE-bgp] quit

