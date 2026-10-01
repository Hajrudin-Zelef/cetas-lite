---
id: collect-261001-huawei/huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b-19
title: "high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "throughput"]
source: docs/RAG/collect-261001-huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b.md
source_anchor: ""
source_lines: [680, 744]
sha256: cdaf6f797cfa05992f638e24a396215bce9272fa899243078d86e19d01244707
---

# high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b

                                                                                                                                                                             Figure 2-8 Smart roaming handover
 2.3.2            Key Technologies
                                                                                                              2.3.2.3 10 Gigabit Ethernet
2.3.2.1 Continuous Networking                                                                                 A cost-effective network acceleration solution is
Continuous networking allows a WLAN with                 Continuous networking implements AP grouping         to reuse the existing Cat5E/Cat6 cables to increase
multiple APs to provide seamless signal coverage         and priority division by synchronizing data          the rate to 2.5 Gbps (as shown in Figure 2-9),
and ensure optimal service experience for users.         between AP radio chips at the microsecond (μs)       which meets the basic requirements of Wi-Fi 7.
Key technologies of continuous networking include        level. After an AP preempts a channel, it becomes
smart antennas and multi-AP coordination.                the primary AP through negotiation, and instructs
                                                         secondary APs to adjust the coverage angles
                                                         of their smart antennas and send data packets
                                                         synchronously. In this way, multiple APs can work
                                                         on the same channel, significantly improving
                                                         spectrum efficiency and system capacity.                                                                        Figure 2-9 2.5GE Wi-Fi 7 traditional network


                                                    24                                                                                                              25
When the bandwidth requirement exceeds 2.5               · Passive Ethernet Network (PEN) Solution:                                                     Table 2-4 Key metrics
Gbps, another option is optical cables, which can        uses passive all-optical technology to build a
support future bandwidth requirements of over            simplified 10 Gbps all-optical network with high       Category                       Metric                           Excellent               Good
10 Gbps. The following solutions are available           performance. The solution extends high-speed
based on the evolution of all-optical Ethernet           10GE optical fibers to rooms to meet the future                       (Wireless) Single-user maximum rate            ≥ 3.5 Gbps              ≤ 3.5 Gbps
technologies:                                            service evolution needs, and saves the deployment
                                                         of active ELV rooms, providing security and                              (Wireless) Total throughput of
· Active Ethernet Network Solution: uses hybrid          reliabilit y. Figure 2-11 shows the solution                            three-AP 80 MHz intra-frequency               960 Mbps               700 Mbps
cables to provide long-distance PoE power supply         networking.                                                                 networking with 18 STAs
                                                                                                                Bandwidth
and high-speed access capabilities, as shown in                                                                                (Wired) High bandwidth at the access
                                                                                                                                                                            Exclusive 10GE           Exclusive GE
Figure 2-10.                                                                                                                        side (access device uplink)

                                                                                                                                 (Wired) High bandwidth based on
                                                                                                                                                                           Single module: >
                                                                                                                                 the passive, simplified architecture                           Single module: 10 Gbps
                                                                                                                                                                               100 Gbps
                                                                                                                                (aggregation/core device downlink)

                                                                                                                                 Number of 4K HD video channels                    60                     40

                                                                                                               Concurrency

                                                                                                                                     Passive high-density access                 96/PCS                80/PCS

                                                                                                                                                                           Continuous large-    Single-AP working, no
                                                                                                                                 (Wireless) Distributed architecture       scale networking     continuous networking
                                                                                                                                                                           and SSID sharing        or zero roaming
                                                                                                               Architecture
                                                                                                                                 (Wired) Integrated networking of
                                                                                                                                                                               Unified
                                                                                                                                aggregation and access devices and                                  Not supported
                                                                                                                                                                             management
                                                                                                                                       unified management

                                                                                                              Long-distance
                                                                                                                                  Long-distance PoE++ capability                ≥ 250 m                ≥ 200 m
                                                                                                                 access




  Figure 2-10 Active Ethernet Network Solution                       Figure 2-11 PEN Solution
                                                                                                             2.4 Connect Everything


