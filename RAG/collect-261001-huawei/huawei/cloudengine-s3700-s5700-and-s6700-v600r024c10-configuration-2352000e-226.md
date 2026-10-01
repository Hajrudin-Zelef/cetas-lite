---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-226
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [33185, 33349]
sha256: 7c355bfd6f4acf9c9f23a52a7b7752be9bd9db9285b69537be01694af8c5a826
---

                    # Check the IP routing table on CE2. The command output shows that the
                    outbound interface of the default route has changed to VLANIF 20. That is, L2VPN
                    traffic is switched to the secondary PW.
                    [CE2] display ip routing-table 0.0.0.0
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance
                    ------------------------------------------------------------------------------
                    Routing Table : Public
                    Summary Count : 1
                    Destination/Mask Proto Pre Cost                Flags NextHop           Interface


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                        529
VPN Configuration
VPN Configuration                                                                                      5 VPWS Configuration


                          0.0.0.0/0 Static 100 0             D 10.1.2.1         Vlanif20

                    # Resolve the failure that is manually simulated on 10GE1/0/2 of PE3.
                    [PE3] interface 10ge 1/0/1
                    [PE3-10GE1/0/2] undo shutdown
                    [PE3-10GE1/0/2] quit

                    # After the network becomes stable, check the IP routing table on CE2 again. The
                    command output shows that the outbound interface of the default route has
                    changed to VLANIF 10. That is, L2VPN traffic is switched back to the primary PW.
                    [CE2] display ip routing-table 0.0.0.0
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance
                    ------------------------------------------------------------------------------
                    Routing Table : Public
                    Summary Count : 1
                    Destination/Mask Proto Pre Cost                Flags NextHop           Interface

                          0.0.0.0/0 Static 60 0              D 10.1.1.1         Vlanif10


Configuration Scripts
                    ●     CE1
                          #
                          sysname CE1
                          #
                          vlan batch 20 30
                          #
                          interface Vlanif30
                           ip address 10.1.1.1 255.255.255.252
                           ip address 10.1.2.1 255.255.255.252 sub
                          #
                          interface Vlanif20
                           ip address 10.1.3.1 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk pvid vlan 30
                           port trunk allow-pass vlan 30
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           port trunk allow-pass vlan 20
                          #
                          return

                    ●     CE2
                          #
                          sysname CE2
                          #
                          vlan batch 10 20
                          #
                          interface Vlanif10
                           ip address 10.1.1.2 255.255.255.252
                          #
                          interface Vlanif20
                           ip address 10.1.2.2 255.255.255.252
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 10
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           port trunk allow-pass vlan 20
                          #


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                        530
VPN Configuration
VPN Configuration                                                                              5 VPWS Configuration

                        ip route-static 0.0.0.0 0.0.0.0 Vlanif10 10.1.1.1
                        ip route-static 0.0.0.0 0.0.0.0 Vlanif20 10.1.2.1 preference 100
                        #
                        return
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 20 30
                        #
                        bfd
                        #
                        mpls lsr-id 1.1.1.1
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        mpls l2vpn
                        #
                        pw-template 1to2
                         peer-address 2.2.2.2
                         control-word
                        #
                        pw-template 1to3
                         peer-address 3.3.3.3
                         control-word
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 3.3.3.3
                         remote-ip 3.3.3.3
                        #
                        interface Vlanif20
                         ip address 10.0.2.1 255.255.255.252
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif30
                         ip address 10.2.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        tunnel-policy p1
                         tunnel select-seq cr-lsp load-balance-number 1
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         mpls l2vc pw-template 1to3 100 tunnel-policy p1
                         mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100
                         mpls l2vc pw-template 1to2 200 secondary
                         mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 secondary
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        interface Tunnel2
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.3


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    531
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration

