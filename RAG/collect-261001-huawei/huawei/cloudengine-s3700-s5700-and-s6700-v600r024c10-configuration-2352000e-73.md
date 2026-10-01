---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-73
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [9862, 10021]
sha256: 2d612c9a95d841c4740a670a862e9917a92a928b57c6f7118958c340a4b62a27
---

                    # Configure ASBR1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR1
                    [ASBR1] mpls lsr-id 2.2.2.2
                    [ASBR1] mpls
                    [ASBR1-mpls] quit
                    [ASBR1] mpls ldp
                    [ASBR1-mpls-ldp] quit
                    [ASBR1] interface Vlanif 100
                    [ASBR1-Vlanif100] mpls
                    [ASBR1-Vlanif100] mpls ldp
                    [ASBR1-Vlanif100] quit

                    # Configure ASBR2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR2
                    [ASBR2] mpls lsr-id 3.3.3.3
                    [ASBR2] mpls
                    [ASBR2-mpls] quit
                    [ASBR2] mpls ldp
                    [ASBR2-mpls-ldp] quit
                    [ASBR2] interface Vlanif 100
                    [ASBR2-Vlanif100] mpls
                    [ASBR2-Vlanif100] mpls ldp
                    [ASBR2-Vlanif100] quit

                    # Configure PE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] mpls lsr-id 4.4.4.4
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface Vlanif 100
                    [PE2-Vlanif100] mpls
                    [PE2-Vlanif100] mpls ldp
                    [PE2-Vlanif100] quit

                    After completing the configuration, run the display mpls ldp session command
                    on a PE or ASBR. The command output shows that the session status is

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         157
VPN Configuration
VPN Configuration                                                                                3 IPv4 L3VPN Configuration


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
                     2.2.2.2:0        Operational DU Passive 0000:02:30 604/604
                    --------------------------------------------------------------------------
                    TOTAL: 1 Session(s) Found.


         Step 3 Configure basic IPv4 L3VPN functions on PE1 and PE2.
                          NOTE

                         PE1's import and export VPN targets must match PE2's export and import VPN targets,
                         respectively.

                    # Configure CE1 to establish an EBGP peer relationship with PE1.
                    [CE1] bgp 65001
                    [CE1-bgp] peer 10.1.1.2 as-number 100
                    [CE1-bgp] network 5.5.5.5 255.255.255.255
                    [CE1-bgp] ipv4-family unicast
                    [CE1-bgp-af-ipv4] peer 10.1.1.2 enable
                    [CE1-bgp-af-ipv4] quit
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
                    [PE1-bgp] peer 2.2.2.2 as-number 100
                    [PE1-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [PE1-bgp] ipv4-family vpnv4
                    [PE1-bgp-af-vpnv4] peer 2.2.2.2 enable


                    # Configure ASBR1 to establish an MP-IBGP peer relationship with PE1.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] peer 1.1.1.1 as-number 100


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          158
VPN Configuration
VPN Configuration                                                                           3 IPv4 L3VPN Configuration

                    [ASBR1-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [ASBR1-bgp] ipv4-family vpnv4
                    [ASBR1-bgp-af-vpnv4] peer 1.1.1.1 enable


                    The configurations of CE2, PE2, and ASBR2 are similar to the configurations of
                    CE1, PE1, and ASBR1, respectively.

                    After completing the configuration, run the display bgp vpnv4 vpn-instance vpn-
                    instance-name peer command on a PE. The command output shows that a BGP
                    peer relationship has been established between the PE and CE and is in the
                    Established state. Run the display bgp vpnv4 all peer command on a PE. The
                    command output shows that the PE has established a BGP peer relationship with
                    the CE and ASBR in the same AS, and the BGP peer relationships are in the
                    Established state.

                    The following example uses the command output on PE1.
                    <PE1> display bgp vpnv4 vpn-instance vpn1 peer
                     BGP local router ID : 10.16.1.2
                     Local AS number : 100

                    VPN-Instance vpn1, Router ID 10.16.1.2:
                    Total number of peers : 1          Peers in established state : 1

                      Peer        V        AS MsgRcvd MsgSent OutQ Up/Down             State PrefRcv
                      10.1.1.1    4      65001       79  80    0 01:05:48 Established       1
                    <PE1> display bgp vpnv4 all peer
                     BGP local router ID : 10.16.1.2
                     Local AS number : 100
                     Total number of peers : 2          Peers in established state : 2

                     Peer        V        AS MsgRcvd MsgSent OutQ Up/Down         State PrefRcv
                     2.2.2.2     4       100   180   180   0 02:33:25 Established     1

                     Peer of IPv4-family for vpn instance :

                     VPN-Instance vpn1, Router ID 10.16.1.2:
                     Peer       V       AS MsgRcvd MsgSent OutQ Up/Down             State PrefRcv
                     10.1.1.1    4    65001      80      80  0 01:06:34 Established      1

         Step 4 Configure inter-AS VPN Option B.

