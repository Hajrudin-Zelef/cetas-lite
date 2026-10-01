---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-135
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [19273, 19425]
sha256: 28b60460f48ac92385c7dbac026e00345cacfa6f17db972cad668a27cc14ef4b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

4.13.1 Understanding the MPLS TE Bandwidth Flooding
Threshold
                 IGP TE requires link information to be flooded in order to form a unified TEDB on
                 an MPLS TE network. Information flooding is triggered by the establishment of an
                 MPLS TE tunnel or by one of the following conditions:
                 ●      A specific IGP TE flooding interval elapses. This interval is configurable.
                 ●      A link is activated or deactivated.
                 ●      An LSP fails to be established when no adequate bandwidth can be reserved.
                 ●      Link attributes, such as the administrative group attribute or affinity attribute,
                        change.
                 ●      The link bandwidth changes.
                        When the available bandwidth of an MPLS interface changes, the system
                        automatically updates information in the TEDB and floods it. When a lot of
                        tunnels are to be established on a node, the node reserves bandwidth and
                        frequently updates information in the TEDB and floods it. For example, the

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            323
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                        bandwidth of a link is 100 Mbit/s. If 100 TE tunnels, each with bandwidth of 1
                        Mbit/s, are established, the system floods link information 100 times.
                 To suppress the frequency of updating and flooding information in the TEDB
                 caused by link bandwidth changes, a bandwidth flooding mechanism is defined
                 based on the following conditions:
                 ●      The ratio of the bandwidth reserved for an MPLS TE tunnel to the available
                        link bandwidth in the TEDB is greater than or equal to the specified threshold.
                 ●      The ratio of the bandwidth released by an MPLS TE tunnel to the available
                        bandwidth in the TEDB is greater than or equal to the specified threshold.
                 If either of the preceding conditions is met, IGP TE floods link information, and
                 CSPF updates the TEDB accordingly.
                 Assume that the available bandwidth of a link is 100 Mbit/s and 100 TE tunnels,
                 each with bandwidth of 1 Mbit/s, are established over the link. The flooding
                 threshold is 10%. Figure 4-23 shows the ratio of the bandwidth reserved for each
                 MPLS TE tunnel to the available bandwidth in the TEDB.
                 Bandwidth flooding is not performed when tunnels 1 to 9 are created. After tunnel
                 10 is created, the bandwidth information (10 Mbit/s in total) on tunnels 1 to 10 is
                 flooded. The available bandwidth is 90 Mbit/s. Similarly, no bandwidth information
                 is flooded after tunnels 11 to 18 are created. After tunnel 19 is created, bandwidth
                 information of tunnels 11 to 19 is flooded. The process repeats until tunnel 100 is
                 established.

                 Figure 4-23 Ratio of the bandwidth reserved for each MPLS TE tunnel to the
                 available bandwidth in the TEDB




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             324
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration


4.13.2 Configuring the MPLS TE Bandwidth Flooding
Threshold
Prerequisites
                 Before configuring the MPLS TE bandwidth flooding threshold, you have
                 completed the following task:
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 When the link bandwidth changes slightly, you can configure the bandwidth
                 flooding threshold on the ingress or transit node of an MPLS TE tunnel to control
                 the flooding time on the local node. This prevents frequent flooding and conserves
                 network resources.
                 Perform the following configuration on the ingress or transit node of an MPLS TE
                 tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the view of an MPLS TE-enabled link interface.
                 interface interface-type interface-number

         Step 3 Switch the interface mode from Layer 2 to Layer 3.
                 undo portswitch

                 Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                 S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                 Layer 2 mode to Layer 3 mode using the undo portswitch command.
                 Determine whether to perform this step based on the current interface mode.
         Step 4 Configure the bandwidth flooding threshold.
                 mpls te bandwidth change thresholds { down percent-down | up percent-up }

                 ----End

4.13.3 Example for Configuring the MPLS TE Bandwidth
Flooding Threshold
Networking Requirements
                 On the network shown in Figure 4-24, RSVP-TE is used to establish a TE tunnel
                 with the bandwidth of 50 Mbit/s from LSR1 to LSR4. The maximum reservable
                 bandwidth and BC0 bandwidth of each link are both 100 Mbit/s.
                 To reduce the number of flooding times and conserve network resources, the
                 bandwidth flooding threshold is set to 20%. If the ratio of the bandwidth used or
                 released by the MPLS TE tunnel to the available bandwidth in the TEDB is greater

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 325
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration


                 than or equal to 20%, an IGP floods the bandwidth information, and CSPF updates
                 the TEDB.

                 Figure 4-24 Network diagram for configuring the MPLS TE bandwidth flooding
                 threshold




Procedure
         Step 1 Configure interface IP addresses for the nodes.

                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 24
                 [LSR1-Vlanif100] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface loopback 1
                 [LSR1-loopback1] ip address 1.1.1.9 32
                 [LSR1-loopback1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

         Step 2 Configure IS-IS to advertise routes.

                 # Configure LSR1.
                 [LSR1] isis 1
                 [LSR1-isis-1] is-level level-2
                 [LSR1-isis-1] network-entity 00.0005.0000.0000.0001.00
                 [LSR1-isis-1] quit
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] isis enable 1
                 [LSR1-Vlanif100] quit
                 [LSR1] interface loopback 1
                 [LSR1-LoopBack1] isis enable 1
                 [LSR1-LoopBack1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

