---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-208
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [30151, 30344]
sha256: 43558e62410872eea5cfdb7dd165588de59f2b4eb6f5880aab2d14f62efb1c39
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         vpn-target 100:1 export-extcommunity
                         vpn-target 100:1 import-extcommunity
                        #
                        bfd
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        explicit-path tope2
                         next hop 10.2.1.2
                         next hop 3.3.3.3
                        #
                        explicit-path tope3
                         next hop 10.1.1.2
                         next hop 2.2.2.2
                        #
                        interface vlanif 200
                         ip address 10.2.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface vlanif 100
                         ip address 10.1.1.1 255.255.255.252
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 2.2.2.2
                         mpls te tunnel-id 1
                         mpls te path explicit-path tope3
                         mpls te reserved-for-binding
                        #
                        interface Tunnel2
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.3
                         mpls te tunnel-id 2
                         mpls te path explicit-path tope2
                         mpls te reserved-for-binding
                        #
                        bgp 100
                         peer 2.2.2.2 as-number 100
                         peer 2.2.2.2 connect-interface LoopBack1
                         peer 3.3.3.3 as-number 100
                         peer 3.3.3.3 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 2.2.2.2 enable



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      500
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                          peer 3.3.3.3 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 2.2.2.2 enable
                          peer 3.3.3.3 enable
                         #
                         ipv4-family vpn-instance vpn1
                          import-route direct
                          auto-frr
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 10.1.1.0 0.0.0.3
                          network 10.2.1.0 0.0.0.255
                          network 1.1.1.1 0.0.0.0
                          mpls-te enable
                        #
                        tunnel-policy policy1
                         tunnel binding destination 3.3.3.3 te tunnel 2
                         tunnel binding destination 2.2.2.2 te tunnel 1
                        #
                        bfd pe1tope2 bind mpls-te interface Tunnel2
                         discriminator local 12
                         discriminator remote 21
                         min-tx-interval 100
                         min-rx-interval 100
                         process-pst
                        #
                        return

                 ●      PE2
                        #
                        sysname PE2
                        #
                        ip vpn-instance vpn1
                         route-distinguisher 100:2
                         tnl-policy policy1
                         vpn-target 100:1 export-extcommunity
                         vpn-target 100:1 import-extcommunity
                        #
                        bfd
                        #
                        vlan batch 200
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        explicit-path tope1
                         next hop 10.2.1.1
                         next hop 1.1.1.1
                        #
                        interface vlanif 200
                         ip address 10.2.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      501
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        interface Tunnel2
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 1.1.1.1
                         mpls te tunnel-id 3
                         mpls te path explicit-path tope1
                         mpls te reserved-for-binding
                        #
                        bgp 100
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpn-instance vpn1
                          import-route direct
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 10.2.1.0 0.0.0.255
                          network 3.3.3.3 0.0.0.0
                          mpls-te enable
                        #
                        tunnel-policy policy1
                         tunnel binding destination 1.1.1.1 te tunnel 2
                        #
                        bfd pe2tope1 bind mpls-te interface Tunnel2
                         discriminator local 21
                         discriminator remote 12
                         min-tx-interval 100
                         min-rx-interval 100
                        #
                        return

