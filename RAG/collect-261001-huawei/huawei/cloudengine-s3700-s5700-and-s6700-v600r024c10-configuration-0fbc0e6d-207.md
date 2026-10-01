---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-207
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [29998, 30150]
sha256: 646a88ee04a52b00f31eb787a0b2ae86177a957b9d54a99b02bd2d73e31e799d
---

                 # Configure PE1.
                 [PE1] interface tunnel 2
                 [PE1-Tunnel2] ip address unnumbered interface loopback 1
                 [PE1-Tunnel2] tunnel-protocol mpls te
                 [PE1-Tunnel2] destination 3.3.3.3
                 [PE1-Tunnel2] mpls te tunnel-id 2
                 [PE1-Tunnel2] mpls te path explicit-path tope2
                 [PE1-Tunnel2] mpls te reserved-for-binding
                 [PE1-Tunnel2] quit
                 [PE1] interface tunnel 1
                 [PE1-Tunnel1] ip address unnumbered interface loopback 1
                 [PE1-Tunnel1] tunnel-protocol mpls te
                 [PE1-Tunnel1] destination 2.2.2.2
                 [PE1-Tunnel1] mpls te tunnel-id 1
                 [PE1-Tunnel1] mpls te path explicit-path tope3
                 [PE1-Tunnel1] mpls te reserved-for-binding
                 [PE1-Tunnel1] quit

                 # Configure PE2.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         497
MPLS Configuration
MPLS Configuration                                                                            4 MPLS TE Configuration

                 [PE2] interface tunnel 2
                 [PE2-Tunnel2] ip address unnumbered interface loopback 1
                 [PE2-Tunnel2] tunnel-protocol mpls te
                 [PE2-Tunnel2] destination 1.1.1.1
                 [PE2-Tunnel2] mpls te tunnel-id 3
                 [PE2-Tunnel2] mpls te path explicit-path tope1
                 [PE2-Tunnel2] mpls te reserved-for-binding
                 [PE2-Tunnel2] quit

                 # Configure PE3.
                 [PE3] interface tunnel 1
                 [PE3-Tunnel1] ip address unnumbered interface loopback 1
                 [PE3-Tunnel1] tunnel-protocol mpls te
                 [PE3-Tunnel1] destination 1.1.1.1
                 [PE3-Tunnel1] mpls te tunnel-id 4
                 [PE3-Tunnel1] mpls te path explicit-path tope1
                 [PE3-Tunnel1] mpls te reserved-for-binding
                 [PE3-Tunnel1] quit

                 After the configuration is complete, run the display mpls te tunnel-interface
                 tunnel command on each node. The command output shows that the states of
                 Tunnel1 and Tunnel2 on PE1, Tunnel2 on PE2, and Tunnel1 on PE3 are all CR-LSP
                 is Up.
         Step 7 Configure VPN FRR.
                 # Configure a VPN instance on PE1, PE2, and PE3. Set the VPN instance name to
                 vpn1, RDs to 100:1, 100:2, and 100:3 respectively, and all RTs to 100:1. Configure
                 the PEs to access the CEs. For detailed configurations, see Configuration Scripts.
                 # Establish MP IBGP peer relationship between PE1 and PE2, and between PE1
                 and PE3. The BGP AS number of PE1, PE2, and PE3 is 100. Loopback1 is used as
                 the interface to establish BGP sessions. For detailed configurations, see
                 Configuration Scripts.
                 # Configure tunnel policies for PE1, PE2, and PE3 and apply the policies to the
                 VPN instances.
                 # Configure PE1.
                 [PE1] tunnel-policy policy1
                 [PE1-tunnel-policy-policy1] tunnel binding destination 3.3.3.3 te tunnel 2
                 [PE1-tunnel-policy-policy1] tunnel binding destination 2.2.2.2 te tunnel 1
                 [PE1-tunnel-policy-policy1] quit
                 [PE1] ip vpn-instance vpn1
                 [PE1-vpn-instance-vpn1] tnl-policy policy1
                 [PE1-vpn-instance-vpn1] quit

                 # Configure PE2.
                 [PE2] tunnel-policy policy1
                 [PE2-tunnel-policy-policy1] tunnel binding destination 1.1.1.1 te tunnel 2
                 [PE2-tunnel-policy-policy1] quit
                 [PE2] ip vpn-instance vpn1
                 [PE2-vpn-instance-vpn1] tnl-policy policy1
                 [PE2-vpn-instance-vpn1] quit

                 # Configure PE3.
                 [PE3] tunnel-policy policy1
                 [PE3-tunnel-policy-policy1] tunnel binding destination 1.1.1.1 te tunnel 1
                 [PE3-tunnel-policy-policy1] quit
                 [PE3] ip vpn-instance vpn1
                 [PE3-vpn-instance-vpn1] tnl-policy policy1
                 [PE3-vpn-instance-vpn1] quit


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                        498
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


                 # Configure VPN FRR on PE1.
                 [PE1] bgp 100
                 [PE1-bgp] ipv4-family vpn-instance vpn1
                 [PE1-bgp-vpn1] auto-frr
                 [PE1-bgp-vpn1] quit
                 [PE1-bgp] quit

                 After the configuration is complete, CEs can communicate, and traffic flows
                 through PE1, Switch, and PE2. If the cable of any interface between PE1 and PE2 is
                 removed, the switch fails, or PE2 fails, VPN traffic is switched to the backup path
                 PE1- >PE3. Time taken in fault recovery is close to the IGP convergence time.
         Step 8 Configure BFD for TE tunnel.
                 # Configure a BFD session on PE1 to monitor the TE tunnel of the primary path.
                 Set the minimum intervals at which BFD packets are sent and received.
                 [PE1] bfd
                 [PE1-bfd] quit
                 [PE1] bfd pe1tope2 bind mpls-te interface tunnel2
                 [PE1-bfd-lsp-session-pe1tope2] discriminator local 12
                 [PE1-bfd-lsp-session-pe1tope2] discriminator remote 21
                 [PE1-bfd-lsp-session-pe1tope2] min-tx-interval 100
                 [PE1-bfd-lsp-session-pe1tope2] min-rx-interval 100
                 [PE1-bfd-lsp-session-pe1tope2] process-pst

                 # Establish a BFD session on PE2 and specify a TE tunnel as the reverse tunnel. Set
                 the minimum intervals at which BFD packets are sent and received.
                 [PE2] bfd
                 [PE2-bfd] quit
                 [PE2] bfd pe2tope1 bind mpls-te interface tunnel2
                 [PE2-bfd-lsp-session-pe2tope1] discriminator local 21
                 [PE2-bfd-lsp-session-pe2tope1] discriminator remote 12
                 [PE2-bfd-lsp-session-pe2tope1] min-tx-interval 100
                 [PE2-bfd-lsp-session-pe2tope1] min-rx-interval 100

                 After the configuration is complete, run the display bfd session { all |
                 discriminator discr-value | mpls-te interface interface-type interface-number }
                 [ verbose ] command on PE1 and PE2 to check whether the BFD session is up.

                 ----End

Verifying the Configuration
                 Connect tester's Port 1 and Port 2 to CE1 and CE2, respectively. Inject traffic
                 destined for Port 2 into Port 1. When the cable of any interface between PE1 and
                 PE2 is removed, the fault can be rectified in milliseconds.

Configuration Scripts
                         NOTE

                        The configuration scripts for CE1, CE2, and the switch, and the configuration details for
                        accessing the PEs to the CEs are not provided.
                 ●      PE1
                        #
                        sysname PE1
                        #
                        ip vpn-instance vpn1
                         route-distinguisher 100:1
                         tnl-policy policy1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         499
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

