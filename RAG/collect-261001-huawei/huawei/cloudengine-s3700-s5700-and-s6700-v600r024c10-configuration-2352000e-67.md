---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-67
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [8880, 9009]
sha256: d803a1f85747ec8f9034d70a9f33b982ad002a7b08693cebe87f257b5cd3e590
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    After completing the configuration, run the display mpls ldp session command
                    on a PE or ASBR. The command output shows that the session status is
                    Operational, indicating that an LDP peer relationship has been established
                    between the PE and ASBR in the same AS.
                    The following example uses the command output on PE1.
                    <PE1> display mpls ldp session
                     LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.
                     -------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                    --------------------------------------------------------------------------
                     2.2.2.9:0        Operational DU Passive 0000:02:30 604/604


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          141
VPN Configuration
VPN Configuration                                                                                3 IPv4 L3VPN Configuration

                    --------------------------------------------------------------------------
                    TOTAL: 1 Session(s) Found.


         Step 3 Configure basic IPv4 L3VPN functions for AS 100 and AS 200.
                           NOTE

                         The VPN targets of the VPN instances on the ASBR and PE in the same AS must match. In
                         different ASs, the VPN targets of the PEs do not need to match.

                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 100
                    [CE1] interface 10GE 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 100
                    [CE1-10GE1/0/1] quit
                    [CE1] interface Vlanif 100
                    [CE1-Vlanif100] ip address 10.1.1.1 24
                    [CE1-Vlanif100] quit
                    [CE1] interface loopback 1
                    [CE1-Loopback1] ip address 11.11.11.11 32
                    [CE1-Loopback1] quit
                    [CE1] bgp 65001
                    [CE1-bgp] peer 10.1.1.2 as-number 100
                    [CE1-bgp] network 11.11.11.11 32
                    [CE1-bgp] quit

                    # Configure PE1 to establish an EBGP peer relationship with CE1.
                    [PE1] ip vpn-instance vpn1
                    [PE1-vpn-instance-vpn1] ipv4-family
                    [PE1-vpn-instance-vpn1-af-ipv4] route-distinguisher 100:1
                    [PE1-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 both
                    [PE1-vpn-instance-vpn1-af-ipv4] quit
                    [PE1-vpn-instance-vpn1] quit
                    [PE1] vlan batch 100 200
                    [PE1] interface 10GE1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 200
                    [PE1-10GE1/0/2] quit
                    [PE1] interface Vlanif 200
                    [PE1-Vlanif200] ip binding vpn-instance vpn1
                    [PE1-Vlanif200] ip address 10.1.1.2 24
                    [PE1-Vlanif200] quit
                    [PE1] bgp 100
                    [PE1-bgp] ipv4-family vpn-instance vpn1
                    [PE1-bgp-vpn1] peer 10.1.1.1 as-number 65001
                    [PE1-bgp-vpn1] quit
                    [PE1-bgp] quit

                    # Configure PE1 to establish an MP-IBGP peer relationship with ASBR1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.9 as-number 100
                    [PE1-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [PE1-bgp] ipv4-family vpnv4
                    [PE1-bgp-af-vpnv4] peer 2.2.2.9 enable

                    # Configure ASBR1 to establish an MP-IBGP peer relationship with PE1.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] peer 1.1.1.9 as-number 100
                    [ASBR1-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [ASBR1-bgp] ipv4-family vpnv4
                    [ASBR1-bgp-af-vpnv4] peer 1.1.1.9 enable


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                        142
VPN Configuration
VPN Configuration                                                                        3 IPv4 L3VPN Configuration


                          NOTE

                        The configurations of CE2, PE2, and ASBR2 are similar to the configurations of CE1, PE1,
                        and ASBR1, respectively.

                    After completing the configuration, run the display bgp vpnv4 vpn-instance vpn-
                    instance-name peer command on a PE. The command output shows that a BGP
                    peer relationship has been established between the PE and CE and is in the
                    Established state. Run the display bgp vpnv4 all peer command on a PE. The
                    command output shows that the PE has established a BGP peer relationship with
                    the CE and ASBR in the same AS, and the BGP peer relationships are in the
                    Established state.
                    The following example uses the command output on PE1.
                    <PE1> display bgp vpnv4 vpn-instance vpn1 peer
                     BGP local router ID : 10.10.1.2
                     Local AS number : 100

                    VPN-Instance vpn1, Router ID 10.10.1.2:
                    Total number of peers : 1          Peers in established state : 1

                      Peer         V       AS MsgRcvd MsgSent OutQ Up/Down          State PrefRcv
                      10.1.1.1     4     65001        79  80 0 01:05:48 Established       1
                    <PE1> display bgp vpnv4 all peer
                     Status codes: * - Dynamic
                     BGP local router ID : 10.10.1.2
                     Local AS number : 100
                     Total number of peers : 2
                     Peers in established state : 2
                     Total number of dynamic peers : 0
                      Peer         V       AS MsgRcvd MsgSent OutQ Up/Down          State PrefRcv
                      2.2.2.9      4      100       180  180 0 02:33:25 Established     1

                     Peer of IPv4-family for vpn instance :

                     VPN-Instance vpn1, Router ID 10.10.1.2:
                     Peer       V       AS MsgRcvd MsgSent OutQ Up/Down             State PrefRcv
                     10.1.1.1    4    65001      80      80  0 01:06:34 Established      1

