---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-125
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [17623, 17815]
sha256: 819d5005c102511ba2c8ca0a74e1197ccdf84c4a7e05973b6456edaaf813176f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Verifying the Configuration
                 # Run the reset mpls rsvp-te and display interface tunnel commands in
                 sequence on LSR1. The command output shows that the tunnel interface is up.
                 # Check the RSVP-TE authentication configuration on LSR1.
                 [LSR1] display mpls rsvp-te interface eth-trunk 1
                 Interface: Eth-Trunk1
                  Interface Address: 10.1.1.1
                  Interface state: UP           Interface Index: 0x36
                  Total-BW: 0                  Used-BW: 0
                  Hello configured: NO            Num of Neighbors: 1
                  SRefresh feature: DISABLE         SRefresh Interval: 30 sec
                  Mpls Mtu: 1500                 Retransmit Interval: 5000 msec
                  Increment Value: 1
                  Authentication: ENABLE
                  Challenge: ENABLE               WindowSize: 32
                  Next Seq # to be sent:2767789282 0 Key ID: 0xa4ff1cdc0000
                  Bfd Enabled: DISABLE             Bfd Min-Tx: 1000
                  Bfd Min-Rx: 1000               Bfd Detect-Multi: 3
                  RSVP instance name: RSVP0


Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                        #
                        interface Eth-Trunk1
                         undo portswitch
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te authentication cipher %^%#D-hX<^i%{I*n!l)w-_hP0cjUIu'4h7I2qDYx2gwA%^%#


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     297
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration

                         mpls rsvp-te authentication handshake
                         mpls rsvp-te authentication window-size 32
                        #
                        interface 10GE1/0/1
                         eth-trunk 1
                        #
                        interface 10GE1/0/2
                         eth-trunk 1
                        #
                        interface 10GE1/0/3
                         eth-trunk 1
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.3
                         mpls te tunnel-id 1
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Eth-Trunk1
                         undo portswitch
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te authentication cipher %^%#H/jsR,*!wO%dy25}vKYO<&@;1z=M<X6gIL7M[!N#%^%#
                         mpls rsvp-te authentication handshake
                         mpls rsvp-te authentication window-size 32
                        #
                        interface Vlanif100
                         ip address 10.1.2.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         eth-trunk 1
                        #
                        interface 10GE1/0/2
                         eth-trunk 1
                        #
                        interface 10GE1/0/3
                         eth-trunk 1
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 100


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 298
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                          mpls-te enable
                        #
                        return

                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif100
                         ip address 10.1.2.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                          mpls-te enable
                        #
                        return


4.10.4 Example for Configuring RSVP-TE Authentication
(Manual TE FRR)
Networking Requirements
                 On the network shown in Figure 4-18, an RSVP-TE tunnel is established from LSR1
                 to LSR4 along the path LSR1 -> LSR2 -> LSR3 -> LSR4. Configure TE FRR to protect
                 the link LSR2 -> LSR3. Establish a bypass CR-LSP along the path LSR2 -> LSR5 ->
                 LSR3. LSR2 is a PLR, and LSR3 is an MP. Use explicit paths to establish the primary
                 and bypass MPLS TE tunnels.
                 In addition, configure RSVP-TE authentication between LSR2 and LSR3.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      299
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


                         NOTE

