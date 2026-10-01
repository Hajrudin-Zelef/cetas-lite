---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-96
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [13371, 13510]
sha256: 534cfe9aa486f6701b1870db9a62097230f8e79b5b3cf687e2ed87630c3f7f1d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         213
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration


         Step 5 Configure SPEs as RRs and specify UPEs and NPEs as RR clients. The following
                example uses SPE1.
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
                    [SPE1-bgp-af-vpnv4] undo policy vpn-target
                    [SPE1-bgp-af-vpnv4] peer 1.1.1.1 enable
                    [SPE1-bgp-af-vpnv4] peer 1.1.1.1 reflect-client
                    [SPE1-bgp-af-vpnv4] peer 1.1.1.1 next-hop-local
                    [SPE1-bgp-af-vpnv4] peer 2.2.2.2 enable
                    [SPE1-bgp-af-vpnv4] peer 2.2.2.2 reflect-client
                    [SPE1-bgp-af-vpnv4] peer 2.2.2.2 next-hop-local
                    [SPE1-bgp-af-vpnv4] peer 5.5.5.5 enable
                    [SPE1-bgp-af-vpnv4] peer 5.5.5.5 reflect-client
                    [SPE1-bgp-af-vpnv4] peer 5.5.5.5 next-hop-local
                    [SPE1-bgp-af-vpnv4] peer 6.6.6.6 enable
                    [SPE1-bgp-af-vpnv4] peer 6.6.6.6 reflect-client
                    [SPE1-bgp-af-vpnv4] peer 6.6.6.6 next-hop-local
                    [SPE1-bgp-af-vpnv4] quit
                    [SPE1-bgp] quit

         Step 6 Configure a route-policy to adjust the local preference of the primary and backup
                routes.

                    # Configure SPE1.
                    [SPE1] route-policy NPE1 permit node 10
                    [SPE1-route-policy] apply local-preference 200
                    [SPE1-route-policy] quit
                    [SPE1] route-policy NPE2 permit node 10
                    [SPE1-route-policy] apply local-preference 190
                    [SPE1-route-policy] quit
                    [SPE1] route-policy pref permit node 10
                    [SPE1-route-policy] apply local-preference 150
                    [SPE1-route-policy] quit
                    [SPE1] bgp 100
                    [SPE1-bgp] ipv4-family vpnv4
                    [SPE1-bgp-af-vpnv4] peer 1.1.1.1 route-policy pref export
                    [SPE1-bgp-af-vpnv4] peer 2.2.2.2 route-policy pref export
                    [SPE1-bgp-af-vpnv4] peer 5.5.5.5 route-policy NPE1 import
                    [SPE1-bgp-af-vpnv4] peer 6.6.6.6 route-policy NPE2 import
                    [SPE1-bgp-af-vpnv4] quit
                    [SPE1-bgp] quit

                    The configuration of SPE2 is similar to that of SPE1. On SPE2, configure the local
                    preference as 180 for routes imported from NPE1, 170 for routes imported from
                    NPE2, and 50 for routes to be advertised to UPEs.

                    # Configure NPE1.
                    [NPE1] route-policy SPE1 permit node 10
                    [NPE1-route-policy] apply local-preference 200
                    [NPE1-route-policy] quit
                    [NPE1] route-policy SPE2 permit node 10
                    [NPE1-route-policy] apply local-preference 190
                    [NPE1-route-policy] quit
                    [NPE1] bgp 100
                    [NPE1-bgp] ipv4-family vpnv4
                    [NPE1-bgp-af-vpnv4] peer 3.3.3.3 route-policy SPE1 import
                    [NPE1-bgp-af-vpnv4] peer 4.4.4.4 route-policy SPE2 import


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          214
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    [NPE1-bgp-af-vpnv4] quit
                    [NPE1-bgp] quit

