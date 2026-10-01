---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-88
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [12168, 12331]
sha256: 2bf2362b3eb698600c51dc679e8a45b4e0fc68f5def9d5001fe2658ce91e484f
---

                    # Configure UPE1.
                    [UPE1] ip vpn-instance vpna
                    [UPE1-vpn-instance-vpna] ipv4-family
                    [UPE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [UPE1-vpn-instance-vpna-af-ipv4] vpn-target 1:1 both
                    [UPE1-vpn-instance-vpna-af-ipv4] quit
                    [UPE1-vpn-instance-vpna] quit
                    [UPE1] interface Vlanif300
                    [UPE1-Vlanif300] ip binding vpn-instance vpna
                    [UPE1-Vlanif300] ip address 10.1.1.2 24
                    [UPE1-Vlanif300] quit
                    [UPE1] bgp 100
                    [UPE1-bgp] ipv4-family vpn-instance vpna
                    [UPE1-bgp-vpna] peer 10.1.1.1 as-number 65410
                    [UPE1-bgp-vpna] import-route direct
                    [UPE1-bgp-vpna] quit
                    [UPE1-bgp] quit

                    # Configure NPE1.
                    [NPE1] ip vpn-instance vpna
                    [NPE1-vpn-instance-vpna] ipv4-family
                    [NPE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [NPE1-vpn-instance-vpna-af-ipv4] vpn-target 1:1 both
                    [NPE1-vpn-instance-vpna-af-ipv4] quit
                    [NPE1-vpn-instance-vpna] quit
                    [NPE1] interface Vlanif300
                    [NPE1-Vlanif300] ip binding vpn-instance vpna
                    [NPE1-Vlanif300] ip address 10.4.1.1 24
                    [NPE1-Vlanif300] quit
                    [NPE1] bgp 100
                    [NPE1-bgp] ipv4-family vpn-instance vpna
                    [NPE1-bgp-vpna] peer 10.4.1.2 as-number 65420
                    [NPE1-bgp-vpna] import-route direct
                    [NPE1-bgp-vpna] quit
                    [NPE1-bgp] quit

                    # Configure the CE.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE
                    [CE] interface Vlanif100
                    [CE-Vlanif100] ip address 10.4.1.2 24
                    [CE-Vlanif100] quit
                    [CE] interface Vlanif200
                    [CE-Vlanif200] ip address 10.2.1.2 24
                    [CE-Vlanif200] quit
                    [CE] interface Vlanif300
                    [CE-Vlanif300] ip address 10.3.1.1 24
                    [CE-Vlanif300] quit
                    [CE] interface loopback 1
                    [CE-Loopback1] ip address 7.7.7.7 32
                    [CE-Loopback1] quit
                    [CE] bgp 65420
                    [CE-bgp] peer 10.4.1.1 as-number 100
                    [CE-bgp] peer 10.2.1.1 as-number 100
                    [CE-bgp] import-route direct
                    [CE-bgp] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           196
VPN Configuration
VPN Configuration                                                              3 IPv4 L3VPN Configuration


                    The configuration of NPE2 is similar to that of NPE1. For detailed configurations,
                    see Configuration Scripts.

                    After completing the configuration, run the display ip vpn-instance verbose
                    command on UPE1 or an NPE. The command output shows VPN instance
                    configuration.

         Step 4 Establish MP-IBGP peer relationships between UPEs and SPEs and between SPEs
                and NPEs.

                    # Configure UPE1.
                    [UPE1] bgp 100
                    [UPE1-bgp] router-id 1.1.1.1
                    [UPE1-bgp] peer 3.3.3.3 as-number 100
                    [UPE1-bgp] peer 3.3.3.3 connect-interface loopback 1
                    [UPE1-bgp] peer 4.4.4.4 as-number 100
                    [UPE1-bgp] peer 4.4.4.4 connect-interface loopback 1
                    [UPE1-bgp] ipv4-family vpnv4
                    [UPE1-bgp-af-vpnv4] peer 3.3.3.3 enable
                    [UPE1-bgp-af-vpnv4] peer 4.4.4.4 enable
                    [UPE1-bgp-af-vpnv4] quit
                    [UPE1-bgp] quit

                    The configuration of UPE2 is similar to that of UPE1. For detailed configurations,
                    see Configuration Scripts.

                    # Configure SPE1.
                    [SPE1] bgp 100
                    [SPE1-bgp] router-id 3.3.3.3
                    [SPE1-bgp] peer 1.1.1.1 as-number 100
                    [SPE1-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [SPE1-bgp] peer 2.2.2.2 as-number 100
                    [SPE1-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [SPE1-bgp] peer 5.5.5.5 as-number 100
                    [SPE1-bgp] peer 5.5.5.5 connect-interface loopback 1
                    [SPE1-bgp] peer 6.6.6.6 as-number 100
                    [SPE1-bgp] peer 6.6.6.6 connect-interface loopback 1
                    [SPE1-bgp] ipv4-family vpnv4
                    [SPE1-bgp-af-vpnv4] peer 1.1.1.1 enable
                    [SPE1-bgp-af-vpnv4] peer 2.2.2.2 enable
                    [SPE1-bgp-af-vpnv4] peer 5.5.5.5 enable
                    [SPE1-bgp-af-vpnv4] peer 6.6.6.6 enable
                    [SPE1-bgp-af-vpnv4] quit
                    [SPE1-bgp] quit

                    The configuration of SPE2 is similar to that of SPE1. For detailed configurations,
                    see Configuration Scripts.

                    # Configure NPE1.
                    [NPE1] bgp 100
                    [NPE1-bgp] peer 3.3.3.3 as-number 100
                    [NPE1-bgp] peer 3.3.3.3 connect-interface loopback 1
                    [NPE1-bgp] peer 4.4.4.4 as-number 100
                    [NPE1-bgp] peer 4.4.4.4 connect-interface loopback 1
                    [NPE1-bgp] ipv4-family vpnv4
                    [NPE1-bgp-af-vpnv4] peer 3.3.3.3 enable
                    [NPE1-bgp-af-vpnv4] peer 4.4.4.4 enable
                    [NPE1-bgp-af-vpnv4] quit
                    [NPE1-bgp] quit

                    The configuration of NPE2 is similar to that of NPE1. For detailed configurations,
                    see Configuration Scripts.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         197
VPN Configuration
VPN Configuration                                                                          3 IPv4 L3VPN Configuration


         Step 5 Configure a VPN instance on SPEs and specify UPEs as the understratum PEs of
                SPEs.
                    # Configure a VPN instance.
                    [SPE1] ip vpn-instance vpna
                    [SPE1-vpn-instance-vpna] ipv4-family
                    [SPE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [SPE1-vpn-instance-vpna-af-ipv4] vpn-target 1:1 both
                    [SPE1-vpn-instance-vpna-af-ipv4] quit
                    [SPE1-vpn-instance-vpna] quit

                    # Specify the corresponding UPE as an understratum PE.
                    [SPE1] bgp 100
                    [SPE1-bgp] ipv4-family vpnv4
                    [SPE1-bgp-af-vpnv4] peer 1.1.1.1 upe
                    [SPE1-bgp-af-vpnv4] peer 2.2.2.2 upe
                    [SPE1-bgp-af-vpnv4] quit
                    [SPE1-bgp] quit

         Step 6 Configure a static default route and use a route-policy to ensure that an SPE
                advertises only this route to UPEs.
                    # Configure SPE1.
                    [SPE1] ip route-static vpn-instance vpna 0.0.0.0 0.0.0.0 55.55.55.55
                    [SPE1] route-policy default permit node 10
                    [SPE1-route-policy] apply local-preference 200
                    [SPE1-route-policy] quit
                    [SPE1] bgp 100
                    [SPE1-bgp] ipv4-family vpn-instance vpna
                    [SPE1-bgp-vpna] network 0.0.0.0 route-policy default
                    [SPE1-bgp-vpna] quit
                    [SPE1-bgp] quit

