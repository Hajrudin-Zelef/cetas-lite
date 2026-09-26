---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r7615-review-536a6ae7-4
title: "fr-review-dell-poweredge-r7615-review-536a6ae7"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "compute", "cost", "gpu", "gpus", "inference", "intel", "liquid cooling"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r7615-review-536a6ae7.md
source_anchor: ""
source_lines: [90, 138]
sha256: 3fcdc72f988b205f8c60c2b0dff85fc8f130e87cf73439d799a3ae73af148e25
---

# fr-review-dell-poweredge-r7615-review-536a6ae7

When benchmarking version 4.0, the Dell PowerEdge R7615 server scored 375 samples per minute in the "Monster" test, 248 samples per minute in "Junkshop," and 198 samples per minute in "Classroom." These results reflect stable rendering capability, closely aligned with the performance observed in the previous version of Blender.
Blackmagic RAW Speed Test
We also started running the Blackmagic RAW Speed Test, which tests video playback performance. Additionally, this test is more of a hybrid test combining both CPU and GPU in a real-world scenario for RAW decoding.
| Blackmagic RAW Speed Test (Higher is better) | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cores, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) |
| CPU 8K | 59 | FPS 99 |
| CUDA 8K | N/A | N/A |
The HPE ProLiant DL320 demonstrated modest performance during the Blackmagic RAW Speed Test. In video playback and RAW decoding scenarios, it only reached 59 FPS in the CPU 8K test. Since this server does not have a GPU, the CUDA 8K portion of this test was not evaluated.
Geekbench 6
Geekbench 6 is a cross-platform evaluation tool measuring a system's overall performance. However, it would be interesting to analyze single-core and multi-core performance, as well as OpenCL benchmark results. A high score indicates better performance. Let's clarify that we only examined CPU results, as this server does not have a graphics card.
You can find comparisons with any system in the Geekbench browser.
| Geekbench 6 | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cores, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) |
| CPU Benchmark - Single-Core | 2,070 | 1,732 |
| CPU Benchmark - Multi-Core | 12,049 | 12,792 |
| GPU Benchmark – OpenCL | N/A | N/A |
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better.
| Cinebench R23 | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cores, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) |
| CPU (multi-core) (points) | 52,401 | 38,707 |
| CPU (single-core) (points) | 1,343 | 1,245 |
| MP Ratio | 39.03x | 31.10x |
Cinebench 2024
Here are the results from the 2024 version of Cinebench, in terms of CPU and GPU performance.
| Cinebench R23 | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cores, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) |
| CPU (multi-core) (points) | 2,127 | 2,014 |
| CPU (single-core) (points) | 83 | 71 |
| MP Ratio | 25.61x | 28.29x |
y-cruncher
y-cruncher is a multithreaded and scalable program that can calculate Pi and other mathematical constants to billions of digits. Since its launch in 2009, y-cruncher has become a popular benchmarking and stress-testing application among overclockers and hardware enthusiasts. Faster is better in this test.
| y-cruncher (Total Computation Time) | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cores, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) |
| 1 billion digits (seconds) | 36.686 | 21.452 |
| 2.5 billion digits (seconds) | 98.68 | 50.418 |
| 10 billion digits (seconds) |  | 131.135 |
In the y-cruncher benchmark, designed to test computational performance, our Dell PowerEdge R7615 version showcased its number-crunching capabilities. It completed the calculation of one billion digits in 1 seconds and took 36.686 seconds to process 98.68 billion digits. These times reflect the server's modest mastery in handling compute-intensive tasks, demonstrating solid performance for its basic configuration.
UL Procyon AI Inference (CPU)
UL's Procyon AI Inference benchmark suite evaluates the performance of different AI inference engines using state-of-the-art neural networks. These tests are run on the CPU only. The figures below correspond to average inference times; the overall score is on the last line.
| Test | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cores, 3.25 GHz) |
| Mobile Net V3 | 4.27 |
| ResNet 50 | 18.39 |
| Inception V4 | 70.70 |
| Deep Lab V3 | 59.17 |
| YOLO V3 | 84.89 |
| REAL-ESRGAN | 5233.51 |
| Overall Score | 68 |
The Dell PowerEdge R7615 demonstrated varied performance across different neural network tests. For example, it completed the MobileNet V3 test in 4.27 seconds, while the more complex Inception V4 test took 70.70 seconds. The server showed a longer processing time of 5233.51 7615 seconds in the REAL-ESRGAN test, highlighting the different computational resource requirements of different AI models. Overall, the R68 scored XNUMX, indicating its ability to handle certain AI inference tasks, albeit with varying efficiency across different AI models.
Conclusion
The Dell PowerEdge R7615 server presents enormous potential through its customization options and its role as a major player in the server technology market, combining cost-effectiveness and robust performance capabilities. As a 2U single-socket server, it supports the new 4th generation AMD EPYC 9004 processors, with options ranging from 16 to 128 cores. It can also be equipped with up to 3 TB of DDR5 RAM, significantly increasing its data processing and multitasking capabilities, making it an ideal choice for businesses requiring high processing power.
Storage solutions are equally comprehensive, with the server supporting up to 368.64 TB via multiple drive bays that can accommodate HDDs, SSDs, and NVMe drives. Options such as air cooling, direct liquid cooling, and various power supply choices further enhance this adaptability. The inclusion of Dell's iDRAC9 for server management adds to its appeal, offering advanced control and monitoring capabilities. This variety ensures it can meet diverse performance needs.
Powered by an entry-level AMD EPYC 9354P 32-core processor and 16 GB of RAM, the server performed as adequately as possible on various benchmarks. However, it is important to note that these are only baseline performances and that with high-end configurations, the server should deliver significantly better results. Due to this basic entry-level configuration, having an overall view of performance is therefore a bit tricky. We expect enormous potential from the server when handling more intensive computing and graphics tasks when equipped with more robust configurations featuring high-end processors and multiple GPUs.
Overall, the Dell PowerEdge R7615 presents itself as a server with great potential for scalability and customization. It offers a solid foundation to build upon to meet a wide range of computing needs, from basic tasks to highly demanding ones. Although entry-level performance may be modest, the server's design and capabilities certainly hint at significant improvements with high-end configurations. This makes the R7615 a versatile and promising choice for various enterprise environments, especially for those needing a server capable of scaling alongside their growing computing needs.
