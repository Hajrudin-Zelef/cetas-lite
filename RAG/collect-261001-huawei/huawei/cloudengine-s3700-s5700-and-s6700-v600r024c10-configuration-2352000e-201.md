---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-201
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2020-07-13", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [29451, 29603]
sha256: 268a8ccecb88ee798b6d80aa18303a4f839025da17f5d564df84433b591420f2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    After completing the configurations, run the display this interface command in
                    the tunnel interface view. The command output shows that the Line protocol
                    current state field displays UP, indicating that the MPLS TE tunnel has been
                    established.
                    [PE1-Tunnel10] display this interface
                    Tunnel10 current state : UP (ifindex: 37)


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          471
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration

                    Line protocol current state : UP
                    Last line protocol up time : 2020-07-13 01:29:54
                    Description:
                    Route Port,The Maximum Transmit Unit is 1500, Current BW: 2Mbps
                    Internet Address is unnumbered, using address of LoopBack1(1.1.1.9/32)
                    Encapsulation is TUNNEL, loopback not set
                    Tunnel destination 3.3.3.9
                    Tunnel up/down statistics 1
                    Tunnel ct0 bandwidth is 2000 Kbit/sec
                    Tunnel protocol/transport MPLS/MPLS, ILM is available
                    primary tunnel id is 0x8001, secondary tunnel id is 0x0
                    Current system time: 2020-07-13 01:38:44
                       0 seconds output rate 0 bits/sec, 0 packets/sec
                       0 seconds output rate 0 bits/sec, 0 packets/sec
                       0 packets output, 0 bytes
                       0 output error
                       0 output drop
                       Last 300 seconds input utility rate: 0.00%
                       Last 300 seconds output utility rate: 0.00%

         Step 6 Establish LDP sessions between PEs.
                    # Configure PE1.
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] mpls ldp remote-peer 3.3.3.9
                    [PE1-mpls-ldp-remote-3.3.3.9] remote-ip 3.3.3.9
                    [PE1-mpls-ldp-remote-3.3.3.9] quit

                    # Configure PE2.
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] mpls ldp remote-peer 1.1.1.9
                    [PE2-mpls-ldp-remote-1.1.1.9] remote-ip 1.1.1.9
                    [PE2-mpls-ldp-remote-1.1.1.9] quit

                    After completing the configuration, check whether LDP sessions are established
                    between PEs.
                    The following example uses the command output on PE1. The command output
                    shows that the status of the remote LDP session between PE1 and PE2 is
                    Operational, indicating that an LDP session has been established.
                    [PE1] display mpls ldp session
                     LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.

                    --------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge              KASent/Rcv
                    --------------------------------------------------------------------------
                     3.3.3.9:0        Operational DU Passive 000:00:00 4/5
                    --------------------------------------------------------------------------
                    TOTAL: 1 Session(s) Found.

         Step 7 Configure a tunnel policy and establish a VPWS connection.
                    # Configure PE1.
                    [PE1] tunnel-policy policy1
                    [PE1-tunnel-policy-policy1] tunnel select-seq cr-lsp load-balance-number 1
                    [PE1-tunnel-policy-policy1] quit
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] mpls l2vc 3.3.3.9 10 tunnel-policy policy1
                    [PE1-Vlanif20] quit


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                  472
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration


                    # Configure PE2.
                    [PE2] tunnel-policy policy1
                    [PE2-tunnel-policy-policy1] tunnel select-seq cr-lsp load-balance-number 1
                    [PE2-tunnel-policy-policy1] quit
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] interface vlanif 10
                    [PE2-Vlanif10] mpls l2vc 1.1.1.9 10 tunnel-policy policy1
                    [PE2-Vlanif10] quit

                           NOTE

                         ● The VC IDs at the two ends of a VPWS connection must be the same. Otherwise, the VC
                           cannot go up.
                         ● No IP address needs to be configured on PE interfaces connecting to CEs.

                    ----End

Verifying the Configuration
                    Run the display mpls lsp verbose command on PE1. The command output shows
                    that an MPLS RSVP-TE tunnel has been established between the addresses 1.1.1.9
                    and 3.3.3.9. The value of LspIndex is the same as the LSP index in the MPLS
                    forwarding table, indicating that packets sent to 3.3.3.9 are forwarded over the
                    MPLS TE tunnel.
                    <PE1> display mpls lsp verbose
                    ----------------------------------------------------------------------
                                 LSP Information: RSVP LSP
                    ----------------------------------------------------------------------
                      No                : 1
                      SessionID             : 10
                      IngressLsrID           : 1.1.1.9
                      LocalLspID             : 1
                      Tunnel-Interface : Tunnel10
                      Fec              : 3.3.3.9/32
                      Type               : Main
                      Nexthop                : 10.1.1.2
                      In-Label            : NULL
                      Out-Label              : 1024
                      In-Interface          : ----------
                      Out-Interface           : Vlanif10
                      LspIndex             : 1
                      Token               : ----------
                      LsrType             : Ingress
                      Mpls-Mtu                : 1500
                      LspAge               : 1511 sec
                      Entropy Label Flag : None

                     No              : 2
                     SessionID          : 10
                     IngressLsrID        : 3.3.3.9
                     LocalLspID          : 1
                     Tunnel-Interface : Tunnel10
                     Fec            : 1.1.1.9/32
                     Type             : Main
                     Nexthop             : ----------
                     In-Label          : 3
                     Out-Label           : NULL
                     In-Interface       : Vlanif10
                     Out-Interface        : ----------
                     LspIndex           : 2
                     Token             : ----------
                     LsrType           : Egress
                     Mpls-Mtu             : -----


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                  473
VPN Configuration
VPN Configuration                                                                                     5 VPWS Configuration

