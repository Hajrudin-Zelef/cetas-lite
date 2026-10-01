---
id: collect-261001-huawei/huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b-47
title: "high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/collect-261001-huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b.md
source_anchor: ""
source_lines: [2066, 2120]
sha256: 901037f369a6961dac349b0de7ec3a349a8cb5baedb3f79e382c2aff4ff9497c
---

# high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b

                                                                                                                                    3.5.4.1 RiaStone Factory (Visabeira Group)
                                                                                                                                    Key Challenges and Requirements                                making timely alerts ineffective.
                                                                                                                                    The RiaStone Factory, part of the Visabeira Group,            . AGV: The existence of AGVs at the factory
                                                                                                                                    manufactures ceramic tableware. The factory                     provided a testing opportunity for both improved
                                                                                                                                    provided a live industrial environment used to                  communicat ions and seamless handover
                                                                                                                                    assess communication technologies for Industry 4.0              performance.
                                                                                                                                    use cases, validating the usage of Wi-Fi 7 and TSN              Pressing machines impacted by network
                                                                                                                                    technologies to deploy newly envisioned use cases.              traffic: The factory's pressing machines were
                                                                                                                                                                                                  . halted because their network-enabled latency-
                                                                                                                                    The factory manifested different challenges,                    sensitive Manufacturing Executing System (MES,
                                                                                                                                    highlighted in the following:                                   which handles, e.g., monitoring and control)
                                                                                                                                                                                                    features were being disrupted by the underlying
                                                                                                                                    . Quality Control Point (QCP): This entity                    . network traffic.
                                                                                                                                      required real-time monitoring and fault detection
                                 Figure 3-21 AI-based HD quality inspection                                                           by transferring high-resolution images (over 15             Key Technologies
                                                                                                                                      to 120 MB, compressed to 15 to 18.5 MB) to an               The Wi-Fi 7 network was coupled with Multi-
Manufacturing enterprises need to protect their             promotes seamless flow of high-value data and                             AI server (YOLOv5) for defect detection. This               Link Operation (MLO) capabilities, with the tests
fixed assets and information assets from internal           blurs the original security boundaries. In this case,                     demands high throughput and high reliability.               utilizing frequency bands including 2.4 GHz, 5 GHz,
and external attacks. As digital transformation             ensuring network security is particularly important                     . Security video streaming: In the factory, there are         and 6 GHz. Wi-Fi 7 leverages such bands, along
continues to sweep across the manufacturing                 for secure operations of enterprises.                                     multiple high-quality video streams from cameras,           with 320 MHz bandwidth, 4096-QAM modulation,
industry, manufacturing enterprises typically have                                                                                    requiring the network to handle their transport             Multiple Resource Unit (MRU), and MLO to increase
tens of thousands of production devices that are            The following table lists the recommended                                 over a fast and reliable wireless medium.                   throughput to 23 Gbps. TSN was also deployed by
connected to or about to connect to the network             i n d i c a t o r s o f a h i g h l y re l i a b l e c o n v e rg e d   . Smartwatches: Due to the large size and high                implementing traffic prioritization. Internal press
in a single factory. The convergence of IT and OT           production network.                                                       noise levels of the factory, its operators struggled        communication was placed in queues 5, 6, and 7
                                                                                                                                      to constantly monitor all equipment. The factory's          based on criticality, while factory network traffic
          Table 3-13 Recommended indicators of a highly reliable converged production network                                         existing Wi-Fi network provided unreliable                  used queues 0, 1, 2, and 3.
                                                                                                                                      connectivity for small smartwatch antennas,
                               Item                                               Recommended Indicator

                        WLAN AP standard                                                       Wi-Fi 7

                 Maximum single-user Wi-Fi rate                                              ≥ 3.5 Gbps

                 Packet loss rate in Wi-Fi roaming                                             ≤ 0.1%

                 Switchover time upon a link fault                          ≤ 50 ms, without service interruption

      Switchover time upon an industrial ring network fault                                    ≤ 20 ms

                                                                         Dual fed and selective receiving, ensuring
                  Going-wireless of workstations
                                                                            service continuity during roaming

                 1588v2 clock precision @ 64 hops                                                1 µs

   Video definition and latency of key applications such as video
                                                                                    1080p: latency < 100 ms
                  conferencing and cloud desktop

                    Service latency of key users                                               ≤ 50 ms

                 Unified O&M for IT/OT networks                                              Supported
                                                                                                                                                                           Figure 3-22 Networking diagram


