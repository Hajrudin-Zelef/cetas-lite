---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-145
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [20773, 20901]
sha256: 70c42456af9869b672013b31df8bf7f1c9f526df72574e779ef5c3994d7573fb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The BW Unreserved field indicates the remaining bandwidth of each priority in
                 the reservable bandwidth. The display shows that the unreserved bandwidth
                 changes for class type 7 on the outgoing interfaces on each node along the
                 tunnel. This indicates that certain tunnels succeed in reserving 40 Mbit/s
                 bandwidth with the priority of 7. The bandwidth information also matches the
                 path of a tunnel. This proves that the affinity and mask match the administrative
                 group of each link.
                 You can also run the display mpls te tunnel diagnostic command on LSR2 to
                 check the outbound interface of the tunnel.
                 [LSR2] display mpls te tunnel diagnostic
                 * means the LSP is detour LSP
                 --------------------------------------------------------------------------------
                 LSP-Id                     Destination       In/Out-If
                 --------------------------------------------------------------------------------
                 1.1.1.1:1:3                3.3.3.3         Vlanif100/Vlanif200
                 --------------------------------------------------------------------------------

                 # Establish Tunnel2 on LSR2.
                 [LSR1] interface tunnel2
                 [LSR1-Tunnel2] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel2] tunnel-protocol mpls te
                 [LSR1-Tunnel2] destination 3.3.3.3
                 [LSR1-Tunnel2] mpls te tunnel-id 101
                 [LSR1-Tunnel2] mpls te bandwidth ct0 40000
                 [LSR1-Tunnel2] mpls te affinity property 10011 mask 11101
                 [LSR1-Tunnel2] mpls te priority 6
                 [LSR1-Tunnel2] quit

                 The mask of Tunnel2's affinity attribute is 0x11101. As such, the first three bits of
                 the affinity attribute value need to be compared, so do the last bit. In contrast, the
                 fourth bit is ignored. Because the affinity value of Tunnel2 is 0x10011, this tunnel
                 selects the link with the second and third bits of the administrative group attribute
                 being 0 and at least one of the first and fifth bits being 1. According to the
                 preceding rules, if the value of the administrative group attribute is 0x10001,
                 0x10000, 0x00001, 0x10011, 0x10010, or 0x00011, the value meets requirements.
                 Finally, Tunnel2 selects VLANIF100 of LSR1 (the administrative group attribute
                 value is 0x10001) and VLANIF300 of LSR2 (the administrative group attribute
                 value is 0x10011).
                 ----End

Verifying the Configuration
                 After the configuration is complete, run the display interface tunnel or display
                 mpls te tunnel-interface command on LSR1. The state of Tunnel1 is down.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           347
MPLS Configuration
MPLS Configuration                                                                                  4 MPLS TE Configuration


                 Because the maximum reservable bandwidth is insufficient, Tunnel2 is of a higher
                 priority and has preempted the bandwidth reserved for Tunnel1.
                 Run the display mpls te cspf tedb node command to check the TEDB again and
                 trace the bandwidth change of each link. The command output shows that
                 Tunnel2 passes through VLANIF300 of LSR2.
                 Alternatively, run the display mpls te tunnel diagnostic command on LSR2 to
                 check the outbound interface of the tunnel.
                 [LSR2] display mpls te tunnel diagnostic
                 * means the LSP is detour LSP
                 --------------------------------------------------------------------------------
                 LSP-Id                     Destination       In/Out-If
                 --------------------------------------------------------------------------------
                 1.1.1.1:1:4                3.3.3.3         Vlanif100/Vlanif300
                 --------------------------------------------------------------------------------


Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls te link administrative group 10001
                         mpls te bandwidth max-reservable-bandwidth 50000
                         mpls te bandwidth bc0 50000
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          mpls-te enable
                          network 1.1.1.1 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.3
                         mpls te tunnel-id 1
                         mpls te affinity property 10101 mask 11011
                         mpls te bandwidth ct0 40000
                        #
                        interface Tunnel2
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           348
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         destination 3.3.3.3
                         mpls te tunnel-id 101
                         mpls te priority 6
                         mpls te affinity property 10011 mask 11101
                         mpls te bandwidth ct0 40000
                        #
                        return

