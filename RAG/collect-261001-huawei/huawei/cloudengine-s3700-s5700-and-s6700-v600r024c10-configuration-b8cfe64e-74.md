---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-74
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [10523, 10674]
sha256: 05d379c6cfbbba7277b258d2b0e6a2257e622ba809a4d421c7d2ec3bf4090b4d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    After the configuration is complete, OSPF neighbor relationships can be
                    established between PE1, P, and PE2. Run the display ip routing-table command.
                    The command output shows that the PEs have learned the routes to each other's
                    Loopback1.
         Step 2 Configure basic MPLS functions and MPLS LDP on the MPLS backbone network to
                establish LDP LSPs.
                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface 10ge 1/0/3
                    [PE1-10GE1/0/3] mpls
                    [PE1-10GE1/0/3] mpls ldp
                    [PE1-10GE1/0/3] quit

                    # Configure P.
                    [P] mpls lsr-id 2.2.2.9
                    [P] mpls
                    [P-mpls] quit
                    [P] mpls ldp
                    [P-mpls-ldp] quit
                    [P] interface 10ge 1/0/1
                    [P-10GE1/0/1] mpls
                    [P-10GE1/0/1] mpls ldp


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                        192
QoS Configuration
QoS Configuration                                                                     12 MPLS QoS Configuration

                    [P-10GE1/0/1] quit
                    [P] interface 10ge 1/0/2
                    [P-10GE1/0/2] mpls
                    [P-10GE1/0/2] mpls ldp
                    [P-10GE1/0/2] quit

                    # Configure PE2.
                    [PE2] mpls lsr-id 3.3.3.9
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface 10ge 1/0/3
                    [PE2-10GE1/0/3] mpls
                    [PE2-10GE1/0/3] mpls ldp
                    [PE2-10GE1/0/3] quit

                    After the configuration is complete, LDP sessions can be established between PE1
                    and P and between P and PE2. Run the display mpls ldp session command. The
                    command output shows that the Status field displays Operational.
                    The following example uses the command output on PE1.
                    [PE1] display mpls ldp session
                      LDP Session(s) in Public Network Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                    A '*' before a session means the session is being deleted.
                     ------------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                     ------------------------------------------------------------------------------
                     2.2.2.9:0        Operational DU Active 0000:00:01 6/6
                     ------------------------------------------------------------------------------
                     TOTAL: 1 session(s) Found.

         Step 3 Configure VPN instances on PEs and connect CEs to PEs.
                    # Configure PE1.
                    [PE1] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv4-family
                    [PE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [PE1-vpn-instance-vpna-af-ipv4] vpn-target 111:1 both
                    [PE1-vpn-instance-vpna-af-ipv4] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE1] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv4-family
                    [PE1-vpn-instance-vpnb-af-ipv4] route-distinguisher 100:2
                    [PE1-vpn-instance-vpnb-af-ipv4] vpn-target 222:2 both
                    [PE1-vpn-instance-vpnb-af-ipv4] quit
                    [PE1-vpn-instance-vpnb] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] undo portswitch
                    [PE1-10GE1/0/1] ip binding vpn-instance vpna
                    [PE1-10GE1/0/1] ip address 10.1.1.2 24
                    [PE1-10GE1/0/1] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] undo portswitch
                    [PE1-10GE1/0/2] ip binding vpn-instance vpnb
                    [PE1-10GE1/0/2] ip address 10.2.1.2 24
                    [PE1-10GE1/0/2] quit

                    # Configure PE2.
                    [PE2] ip vpn-instance vpna
                    [PE2-vpn-instance-vpna] ipv4-family
                    [PE2-vpn-instance-vpna-af-ipv4] route-distinguisher 200:1
                    [PE2-vpn-instance-vpna-af-ipv4] vpn-target 111:1 both
                    [PE2-vpn-instance-vpna-af-ipv4] quit
                    [PE2-vpn-instance-vpna] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      193
QoS Configuration
QoS Configuration                                                                    12 MPLS QoS Configuration

                    [PE2] ip vpn-instance vpnb
                    [PE2-vpn-instance-vpnb] ipv4-family
                    [PE2-vpn-instance-vpnb-af-ipv4] route-distinguisher 200:2
                    [PE2-vpn-instance-vpnb-af-ipv4] vpn-target 222:2 both
                    [PE2-vpn-instance-vpnb-af-ipv4] quit
                    [PE2-vpn-instance-vpnb] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] undo portswitch
                    [PE2-10GE1/0/1] ip binding vpn-instance vpna
                    [PE2-10GE1/0/1] ip address 10.3.1.2 24
                    [PE2-10GE1/0/1] quit
                    [PE2] interface 10ge 1/0/2
                    [PE2-10GE1/0/2] undo portswitch
                    [PE2-10GE1/0/2] ip binding vpn-instance vpnb
                    [PE2-10GE1/0/2] ip address 10.4.1.2 24
                    [PE2-10GE1/0/2] quit

                    # Configure IP addresses for interfaces on CEs according to Figure 12-8. The
                    configuration procedure is not mentioned here.

                    After the configuration is complete, each PE can ping the connected CE
                    successfully.

                          NOTE

                        If multiple interfaces on a PE are bound to the same VPN, you need to specify the source IP
                        address when running the ping -vpn-instance command to ping the CE connected to the
                        remote PE. That is, you need to specify the -a source-ip-address parameter when running
                        the ping -vpn-instance vpn-instance-name -a source-ip-address command. Otherwise, the
                        ping may fail.

                    The following example uses the command output on PE1 to show that PE1 can
                    ping CE1.
                    [PE1] ping -vpn-instance vpna 10.1.1.1
                     PING 10.1.1.1: 56 data bytes, press CTRL_C to break
                       Reply from 10.1.1.1: bytes=56 Sequence=1 ttl=255 time=5 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=2 ttl=255 time=3 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=3 ttl=255 time=3 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=4 ttl=255 time=3 ms
                       Reply from 10.1.1.1: bytes=56 Sequence=5 ttl=255 time=16 ms

                     --- 10.1.1.1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 3/6/16 ms

         Step 4 Establish an MP-IBGP peer relationship between PEs.

