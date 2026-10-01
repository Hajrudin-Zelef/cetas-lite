---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-170
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [24590, 24785]
sha256: 46b99909483984997e189334fa2431bf7a6232ba294c52f96447fb4b6acb52ab
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100 500
                        #
                        bfd
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                        #
                        explicit-path backup-path
                         next hop 10.1.5.2
                         next hop 10.1.3.1
                         next hop 3.3.3.9
                        #
                        explicit-path pri-path
                         next hop 10.1.1.2
                         next hop 10.1.2.2
                         next hop 3.3.3.9
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.1 255.255.255.0
                         mpls


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           409
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration

                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te tunnel-id 1
                         mpls te record-route
                         mpls te path explicit-path pri-path
                         mpls te path explicit-path backup-path secondary
                         mpls te backup hot-standby mode revertive wtr 15
                         mpls te backup ordinary best-effort
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                          mpls-te enable
                        #
                        bfd backuplsp2lsrc bind mpls-te interface Tunnel1 te-lsp backup
                         discriminator local 339
                         discriminator remote 439
                         min-tx-interval 500
                         min-rx-interval 500
                         process-
                        pst
                        #
                        bfd prilsp2lsrc bind mpls-te interface Tunnel1 te-lsp
                         discriminator local 139
                         discriminator remote 239
                         min-tx-interval 500
                         min-rx-interval 500
                         process-pst
                        #
                        return

                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 100 200 400
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  410
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.1.4.1 255.255.255.0
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
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                         isis enable 1
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          mpls-te enable
                        #
                        return

                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 200 300
                        #
                        bfd
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 300


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      411
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

