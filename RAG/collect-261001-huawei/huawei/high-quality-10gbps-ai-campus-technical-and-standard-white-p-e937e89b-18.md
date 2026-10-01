---
id: collect-261001-huawei/huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b-18
title: "high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "latency", "throughput"]
source: docs/RAG/collect-261001-huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b.md
source_anchor: ""
source_lines: [610, 679]
sha256: e176eb0035aef0a8256386052d0c0059111ff160d607054020d4bd82a4e4203e
---

# high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b

 technologies. All-optical Ethernet includes active        The Wi-Fi 7 standard introduces technologies                                                                Classic All-Optical          All-Optical             All-Optical
                                                                                                                                                 Classic Ethernet
 Ethernet optical and passive Ethernet optical             such as the 6 GHz frequency band, 320 MHz                                                                        Ethernet              Ethernet: Active         Ethernet: PEN
 technologies.                                             frequency bandwidth, 4096-QAM modulation,
                                                                                                                                                  2.5 Gbps to 10      2.5 Gbps to 10 Gbps,
                                                           Multiple Resource Unit (MRU), and Multi-Link                          Bandwidth                                                      2.5 Gbps to 10 Gbps     2.5 Gbps to 10 Gbps
                                                                                                                                                       Gbps           25 Gbps to 400 Gbps
                                                           Operation (MLO) to increase the throughput
                                                           of Wi-Fi networks to 23 Gbps and provide low-                             PoE                 Y                      N                         Y                       N
                                                           latency access assurance for users, helping high-
                                                           quality 10 Gbps AI campus networks improve user
                                                                                                                                   Passive
                                                           experience.                                                                                   N                      N                         N                       Y
                                                                                                                                 architecture




                                                                                                                              1. The key to selecting a wired solution is the cable        solution increases the cable deployment cost by
                                                                                                                                 type. If electrical cables are used, as shown in          about 30%. A more economical solution is to
                                                                                                                                 Figure 2-5, cables of Cat6A or higher must be             reuse the existing Cat5E/Cat6 cables: Through
                                                                                                                                 deployed if 5 Gbps or 10 Gbps bandwidth is                access switch reconstruction, the transmission rate
                                                                                                                                 required, so as to meet the high-performance              can be increased to 2.5 Gbps to meet the basic
                                                                                                                                 access requirements of Wi-Fi 7. However, this             performance requirements of Wi-Fi 7.




                             Figure 2-4 Comparison between Wi-Fi standards                                                                                Figure 2-5 How cables evolve with Wi-Fi generations


                                                      22                                                                                                                              23
2. In all-optical scenarios, hybrid copper-fiber         power supply and high-speed access capabilities.
   cables (hybrid cables for short, as shown in          This integration significantly simplifies network
   Figure 2-6) can provide both long-distance            deployment.




                                                                                                                                        Figure 2-7 Smart antennas and multi-AP coordination



                                                                                                              2.3.2.2 Smart Roaming
                        Figure 2-6 Hybrid cable and optical module composition
                                                                                                              Smart roaming uses the network AI capability to
                               Table 2-3 Key capabilities of hybrid cables                                    profile the roaming behaviors of STAs and steer
                                                                                                              STA roaming in a differentiated manner. This
    Cable                            PoE Power Supply         PoE+ Power Supply      PoE++ Power Supply       solves problems in traditional roaming such as
                   Cable Diameter
 Specifications                      Distance (15.4 W)         Distance (30 W)         Distance (60 W)
                                                                                                              sticky STAs, long channel scanning time, and poor
  Hybrid cable-                                                                                               roaming compatibility, improving the roaming
                      6.2 mm               1280 m                    500 m                  250 m
    17AWG                                                                                                     experience of STAs. Key technologies of smart
                                                                                                              roaming include STA identification, STA profiling,
  Hybrid cable-
                      5.7 mm                500 m                    200 m                  100 m             STA quality data learning, and roaming steering.
    21AWG
                                                                                                              Smart roaming replaces STA-triggered roaming
  Hybrid cable-                                                                                               with network-steered roaming, optimizes roaming
                      4.4 mm                200 m                    100 m                  50 m
    24AWG                                                                                                     handover opportunities, and shortens the roaming
                                                                                                              time. It supports the personalized roaming
* The American Wire Gauge (AWG) is a                     value indicates a thicker wire diameter and higher   parameter settings for each type of STA , and
  standard unit for measuring the diameter of            current carrying capacity.                           minimizes the possible adverse impacts caused by
  an electrically conducting wire. A smaller AWG                                                              STAs' protocol compatibility and implementation
                                                                                                              differences.

