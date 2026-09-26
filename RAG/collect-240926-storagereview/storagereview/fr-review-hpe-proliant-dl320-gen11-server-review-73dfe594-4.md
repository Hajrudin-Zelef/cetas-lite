---
id: collect-240926-storagereview/storagereview/fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594-4
title: "fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "compute", "dram", "gpu", "gpus", "inference", "intel", "memory", "nvidia"]
source: docs/RAG/clean_en/storagereview/fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594.md
source_anchor: ""
source_lines: [94, 137]
sha256: fe690cacb3210c1706c0f89c9cc5f88b1ba977d25fa4f9675e48acb9ce4d35f5
---

# fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594

In benchmark tests for Blender version 3.5, the HPE ProLiant DL320 achieved varying results across different rendering scenarios with its 4th Gen Intel Xeon-G 6430 processor. For the "Monster" scene, it recorded a score of 279 samples per minute. The "Junkshop" scene (which involves a moderate level of detail and complexity) yielded a score of 177 samples per minute. Finally, in the "Classroom" scene, which could be the most demanding of the three due to complex details and more lighting calculations, the server recorded 135 samples per minute. These results demonstrate that it is not intended for complex rendering tasks as configured.
| HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor (32 cores, 3.7 GHz max)) |  | 
| Blender OptiX version 4.0 (CPU) (Samples per minute; higher is better) |  | 
| Monster | 952 | 
| Junkshop | 613 | 
| Classroom | 437 | 
In Blender version 4.0, the HPE ProLiant DL320 delivered improved results, displaying 952 samples per minute for the "Monster" scene, 613 for "Junkshop," and 437 for "Classroom."
Blackmagic RAW Speed Test
We also started running the Blackmagic RAW Speed Test, which tests video playback performance. Additionally, this test is more of a hybrid test combining both CPU and GPU in a real-world scenario for RAW decoding.
| Blackmagic RAW Speed Test (Higher is better) | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) | 
| CPU 8K | FPS 99 | 
| CUDA 8K | N/A | 
The HPE ProLiant DL320 demonstrated decent performance in the Blackmagic RAW Speed Test. In video playback and RAW decoding scenarios, it reached 99 FPS in the 8K CPU test, indicating it can efficiently handle certain high-resolution video processing. Since we do not have a GPU in this build, we did not evaluate the 8K CUDA portion of this test.
Geekbench 6
Geekbench 6 is a cross-platform evaluation tool measuring a system's overall performance. However, it would be interesting to analyze single-core and multi-core performance, as well as OpenCL benchmark results. A high score indicates better performance. Note that we only examined CPU results, as this server does not have a graphics card.
You can find comparisons with any system in the Geekbench browser.
| Geekbench 6 | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) | 
| CPU Benchmark - Single-Core | 1,732 | 
| CPU Benchmark - Multi-Core | 12,792 | 
| GPU Benchmark – OpenCL | N/A | 
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better.
| Cinebench R23 | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) | 
| CPU (multi-core) (points) | 38,707 | 
| CPU (single-core) (points) | 1,245 | 
| MP Ratio | 31.10x | 
Cinebench 2024
Here are the results from the 2024 version of Cinebench, in terms of CPU and GPU performance.
| Cinebench R23 | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor, 32 cores, 3.7 GHz max) | 
| CPU (multi-core) (points) | 2,014 | 
| CPU (single-core) (points) | 71 | 
| MP Ratio | 28.29x | 
y-cruncher
y-cruncher is a multithreaded and scalable program that can calculate Pi and other mathematical constants to billions of digits. Since its launch in 2009, y-cruncher has become a popular benchmarking and stress testing application among overclockers and hardware enthusiasts. Faster is better in this test.
| y-cruncher (total computation time) |  | 
| 1 billion digits (seconds) | 21.452 | 
| 2.5 billion digits (seconds) | 50.418 | 
| 10 billion digits (seconds) | 131.135 | 
The y-cruncher benchmark results, while indicating the performance limits of the entry-level build, can still be relevant for small and medium-sized businesses (SMBs) in specific contexts. This type of system could be quite adequate for SMBs whose computing needs are not centered on high-intensity computing tasks.
Conclusion
The HPE ProLiant DL320 Gen11 is a highly adaptable and efficient 1U 1P server, given the appropriate configuration, designed to meet the evolving requirements of Edge AI workloads, ROBO (Remote Office/Branch Office) environments, and SMB applications.
At its core is a single 4th Generation Intel Xeon Scalable processor, supporting up to 32 cores and 270 W, along with memory capacity of up to 2 TB at 4800 MT/s and PCIe Gen5 slots. HPE also offers some interesting configurations, including one supporting a dozen hard drives and another with support for 4x NVIDIA L4 GPUs, which is ideal for inference.
Based on our lighter review configuration, the DL320 Gen11 showed modest performance in benchmark tests. Although the absence of a GPU and a larger DRAM footprint limited our testing scope, the server demonstrated adequate capability to handle compute-intensive tasks, which aligns well with the expectations of a budget-conscious build suited to SMBs with moderate computing needs.
The DL320 Gen11 has a lot of potential thanks to its flexibility in terms of processor options and expansion capabilities. SMBs will be able to customize a solution that meets their specific performance requirements and budget constraints, making the DL320 Gen11 a versatile and scalable solution for a wide range of scenarios.
