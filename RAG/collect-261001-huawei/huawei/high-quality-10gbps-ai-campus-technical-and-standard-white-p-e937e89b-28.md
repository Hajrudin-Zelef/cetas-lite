---
id: collect-261001-huawei/huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b-28
title: "high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["energy"]
source: docs/RAG/collect-261001-huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b.md
source_anchor: ""
source_lines: [1056, 1110]
sha256: e73e8d4a1e1d61c4f7eadd3b9cce98f540ab5f76cc9adc492c7743ddfbc54ef4
---

# high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b

Network energy consumption visualization refers            Energy saving can also be achieved by simplifying
                                                                                                                  2.6.2.3 Intelligent Building Collaboration
to collecting energy consumption information               the network architecture. For example, making the
reported by devices such as APs to monitor                 access layer lightweight and the aggregation layer     Intelligent building collaboration is a solution in
the energy usage of the campus net work ,                  passive enables lightweight operation and reduces      which the WLAN collaborates with the building
buildings, floors, and specific APs in real time.          energy consumption.                                    management system to save energy on building
The information is analyzed and displayed on the                                                                  facilities such as air conditioners and lights. Over
intelligent analysis system — usually a component          . The passive architecture saves the deployment        75% of building energy consumption comes from
of the NMS or SDN controller. As shown in the                of active devices at the aggregation layer, saving   air conditioning and lighting. As air conditioners
following figure, the system can display the                 energy at the aggregation layer.                     and lights are mainly for people's use, they can be
network status at the region level, including              . The access layer uses the simplified architecture,   turned off at night when no one uses them to save
regional overall energy consumption by day or                replacing the original access devices that           energy. As technology advances, in addition to
hour, regional network usage (channel utilization),          consume dozens of watts with lightweight             communication functions, a WLAN now supports
and regional overall network quality by day or               devices that consume only several watts.             spatial and personnel sensing based on channel
hour. This visualized management helps network                                                                    state information (CSI), providing information
administrators easily grasp the overall energy                                                                    input for the building management system to take
usage and implement time- and region-specific                                                                     decisions on energy saving.
energy-saving measures.
                                                                                                                  As the WLAN collaborates with the building
                                                                                                                  management system to save energy, the running
                                                                                                                                                                              Figure 2-16 Intelligent building collaboration
                                                                                                                  time of the air conditioning and lighting systems
                                                                                                                  is significantly reduced, lowering their energy
                                                                                                                  consumption by 15% to 30%. This lowers power
                                                                                                                  costs for building operators and offers substantial
                                                                                                                  economic and environmental benefits.




                                                                                                                   2.6.3           Key Metrics
    Figure 2-14 Campus energy consumption
   monitoring through the WLAN visualization
                    system                                                                                                                          Table 2-8 Key metrics for energy saving

                                                                                                                     Category                        Metric                           Excellent                  Good

                                                                                                                                                                                 Supported, saving
                                                                                                                   Device energy
                                                                                                                                       Chip-level dynamic energy saving          more than 10% in            Not supported
                                                                                                                      saving
                                                                                                                                                                                energy consumption

                                                                                                                                        Energy saving based on NE tidal
                                                                                                                                                                                 AI-based prediction       Manual prediction
                                                                                                                                                   prediction

                                                                                                                                       Shutdown when unoccupied, sleep
                                                                                                                                                                                      Supported              Not supported
                                                                                                                     Network                 mode under low load
                                                                                                                   energy saving
                                                                                                                                                 Wake-up time                          Seconds                  Minutes

                                                                                                                                             Energy-saving benefit                       30%                      10%

