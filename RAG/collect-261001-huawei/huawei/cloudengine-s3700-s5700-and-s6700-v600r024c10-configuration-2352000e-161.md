---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-161
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [23127, 23296]
sha256: 5433681f1868762bc9c522469f6c05f823702486420816c85e9de01ec5b2f64b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Procedure
         Step 1 Configure IGP on the MPLS backbone networks in AS100 and AS200 for IP
                connectivity between the ASBR and PE within each MPLS backbone network.

                    OSPF is used as IGP in this example. For detailed configurations, see Configuration
                    Scripts.

                          NOTE

                        The 32-bit IP address of the loopback interface that functions as the LSR ID needs to be
                        advertised using OSPF.

                    After completing the configuration, run the display ospf peer command on an
                    ASBR and PE. The command output shows that the OSPF neighbor relationship is
                    in the Full state, indicating that the OSPF neighbor relationship has been
                    established between the ASBR and PE in the same AS.

                    The ASBR and PE in the same AS can learn and successfully ping the address of
                    each other's loopback interface.

         Step 2 Configure basic MPLS capabilities and MPLS LDP on the MPLS backbone networks
                in AS100 and AS200 and establish MPLS LDP LSPs.

                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface Vlanif100
                    [PE1-Vlanif100] mpls
                    [PE1-Vlanif100] mpls ldp
                    [PE1-Vlanif100] quit

                    # Configure ASBR1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR1
                    [ASBR1] mpls lsr-id 2.2.2.9
                    [ASBR1] mpls
                    [ASBR1-mpls] quit
                    [ASBR1] mpls ldp
                    [ASBR1-mpls-ldp] quit
                    [ASBR1] interface Vlanif100
                    [ASBR1-Vlanif100] mpls
                    [ASBR1-Vlanif100] mpls ldp
                    [ASBR1-Vlanif100] quit

                    # Configure ASBR2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR2
                    [ASBR2] mpls lsr-id 3.3.3.9
                    [ASBR2] mpls
                    [ASBR2-mpls] quit
                    [ASBR2] mpls ldp
                    [ASBR2-mpls-ldp] quit
                    [ASBR2] interface Vlanif100
                    [ASBR2-Vlanif100] mpls
                    [ASBR2-Vlanif100] mpls ldp
                    [ASBR2-Vlanif100] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       366
VPN Configuration
VPN Configuration                                                                                4 IPv6 L3VPN Configuration


                    # Configure PE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] mpls lsr-id 4.4.4.9
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface Vlanif100
                    [PE2-Vlanif100] mpls
                    [PE2-Vlanif100] mpls ldp
                    [PE2-Vlanif100] quit

                    After completing the configuration, run the display mpls ldp session command
                    on a PE or ASBR. The command output shows that the session status is
                    Operational, indicating that an LDP peer relationship has been established
                    between the PE and ASBR in the AS.

                    The following example uses the command output on PE1.
                    <PE1> display mpls ldp session
                     LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.
                     -------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                    --------------------------------------------------------------------------
                     2.2.2.9:0        Operational DU Passive 0000:02:30 604/604
                    --------------------------------------------------------------------------
                    TOTAL: 1 Session(s) Found.


         Step 3 Configure basic IPv6 L3VPN functions in AS100 and AS200.
                          NOTE

                         The VPN targets of the IPv6-address-family-enabled VPN instance configured on the ASBR
                         and PE in the same AS must match.

                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] interface Vlanif100
                    [CE1-Vlanif100] ipv6 enable
                    [CE1-Vlanif100] ipv6 address 2001:db8:1::1 64
                    [CE1-Vlanif100] quit
                    [CE1] interface loopback 1
                    [CE1-Loopback1] ip address 10.10.10.10 32
                    [CE1-Loopback1] quit
                    [CE1] bgp 65001
                    [CE1-bgp] router-id 10.10.10.10
                    [CE1-bgp] peer 2001:db8:1::2 as-number 100
                    [CE1-bgp] ipv6-family unicast
                    [CE1-bgp-af-ipv6] peer 2001:db8:1::2 enable
                    [CE1-bgp-af-ipv6] import-route direct
                    [CE1-bgp-af-ipv6] quit
                    [CE1-bgp] quit

                    # Configure PE1 to establish an EBGP peer relationship with CE1.
                    [PE1] ip vpn-instance vpn1
                    [PE1-vpn-instance-vpn1] ipv6-family
                    [PE1-vpn-instance-vpn1-af-ipv6] route-distinguisher 100:1
                    [PE1-vpn-instance-vpn1-af-ipv6] vpn-target 1:1 both
                    [PE1-vpn-instance-vpn1-af-ipv6] quit
                    [PE1-vpn-instance-vpn1] quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          367
VPN Configuration
VPN Configuration                                                                             4 IPv6 L3VPN Configuration

                    [PE1] interface Vlanif200
                    [PE1-Vlanif200] ip binding vpn-instance vpn1
                    [PE1-Vlanif200] ipv6 enable
                    [PE1-Vlanif200] ipv6 address 2001:db8:1::2 64
                    [PE1-Vlanif200] quit
                    [PE1] bgp 100
                    [PE1-bgp] ipv6-family vpn-instance vpn1
                    [PE1-bgp6-vpn1] peer 2001:db8:1::1 as-number 65001
                    [PE1-bgp6-vpn1] import-route direct
                    [PE1-bgp6-vpn1] quit
                    [PE1-bgp] quit

                    # Configure PE1 to set up an MP-IBGP peer relationship with ASBR1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.9 as-number 100
                    [PE1-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [PE1-bgp] ipv6-family vpnv6
                    [PE1-bgp-af-vpnv6] peer 2.2.2.9 enable

                    # Configure ASBR1 to set up an MP-IBGP peer relationship with PE1.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] peer 1.1.1.9 as-number 100
                    [ASBR1-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [ASBR1-bgp] ipv6-family vpnv6
                    [ASBR1-bgp-af-vpnv6] peer 1.1.1.9 enable

                            NOTE

                         The configurations of CE2, PE2, and ASBR2 are similar to the configurations of CE1, PE1,
                         and ASBR1, respectively.

