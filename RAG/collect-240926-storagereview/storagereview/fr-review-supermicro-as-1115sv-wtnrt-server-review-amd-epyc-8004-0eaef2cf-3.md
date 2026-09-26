---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf-3
title: "fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Samsung"]
dates: []
keywords: ["amd", "benchmark", "compute", "gpu", "inference"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf.md
source_anchor: ""
source_lines: [81, 145]
sha256: 4d91f656da04a4d05f1493068da35f71acb64afe05a874647b7f7fdd8fd47725
---

# fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf

In this performance analysis, however, we will compare it to the Supermicro Storage A+ ASG-1115S-NE316R. It is important to note that this will not be a direct evaluation. Rather, we aim to provide a perspective on the scale of the two AMD processors.
The A+ ASG-1115S-NE316R review unit is configured with the following:
- Processor: EPYC 84 at 9634 cores
- RAM: 384 GB DDR5 (12 x 32 GB DIMMs)
- SSD: 1 TB Samsung PM9A3
- 2x Mellanox ConnectX-6 DX 100G
- Windows Server 2022
Blender OptiX 4.0
The first is Blender OptiX, an open-source 3D modeling application. This benchmark was run using the Blender Benchmark CLI utility. The score is expressed in samples per minute, with the higher being the better.
The Supermicro system showed variable performance in 3D rendering tasks across different scenes when tested with and without GPU acceleration. In all scenarios (Monster, Junkshop, and Classroom), the CPU-only configuration provided higher samples per minute than the GPU-assisted configuration.
Specifically, the Monster scene reached 492.88 samples/min without GPU, surpassing the score of 427.25 samples/min with GPU. Similarly, in the Junkshop scene, the CPU-only test reached 349.35 samples/min, far exceeding the GPU's 266.65 samples/min. The Classroom scene also recorded better performance with the CPU alone, with a score of 255.98 samples/min versus 238.33 samples/min for the GPU.
| Blender 4.0 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5, GPU) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5, CPU) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| Monster | 427.25 | 492.88 | 673.21 | 
| Junkshop | 266.65 | 349.35 | 475.17 | 
| Classroom | 238.33 | 255.98 | 342.06 | 
Blender OptiX 4.1
Blender OptiX 4.1 brings new features, such as GPU-accelerated denoising, streamlining the rendering process and reducing the time needed for denoising tasks. Despite these advances, the overall performance improvements in benchmark scores compared to version 4.0 are minimal, indicating only slight improvements in efficiency (as you will notice in the results below).
The results show that the CPU-only configuration consistently outperformed the GPU-enhanced configuration in all tests, as in version 4.0. Specifically, CPU-only mode reached 493.167 samples/min in Monster, 356.36 samples/min in Junkshop, and 250.87 samples/min in Classroom, compared to GPU-assisted scores of 428.90, 269.39, and 238.58 samples/min respectively.
| Blender 4.0 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5, GPU) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5, CPU) | 
| Monster | 428.90 | 493.167 | 
| Junkshop | 269.39 | 356.36 | 
| Classroom | 238.58 | 250.87 | 
Blackmagic RAW Speed Test
We performed the Blackmagic RAW speed test to extend video playback. This is more of a hybrid test including CPU and GPU performance for actual RAW decoding. Here, we only tested the CPU, which was able to reach 117 FPS.
| Blackmagic RAW Speed Test (Higher is Better) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| CPU 8K | 117 FPS | 131 FPS | 
Blackmagic Disk Speed Test
The Blackmagic Disk Speed Test is another test for which we only have results for the Supermicro. This test runs a 5 GB sample file for read and write speeds. This test showed read speeds of 3.57 GB/s and nearly 2.54 GB/s read on the Ultrastar SN655 NVMe Data Center SSDs installed by Supermicro.
| Blackmagic Disk Speed Test (Higher is Better) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| Write | 2,536 MB/s | 1,415.1 MB/s | 
| Read | 3,568 MB/s | 3,031.0 MB/s | 
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better. Here are the results for all EPYC chips.
Here, the Supermicro 1115SV-WTNRT scored 63,332 points in the multi-core test, demonstrating the processor's robust ability to handle multiple threads simultaneously. On the other hand, the single-core test score was 1,093 points, reflecting its efficiency in tasks requiring single-threaded performance. The MP ratio, which compares multi-core performance to single-core performance, stands at 57.92, indicating a well-balanced architecture for parallel processing and per-core efficiency.
| Cinebench R23 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| CPU (multi-core) (points) | 63,332 | 81,148 | 
| CPU (single-core) (points) | 1,093 | 1,309 | 
| MP Ratio | 57.92 | 61.99x | 
Cinebench 2024
Maxon's Cinebench 2024 is a CPU and GPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. We do not have these figures because the ASG-1115S-NE316R configuration does not have a GPU. Higher scores are better.
The 1115SV-WTNRT scored 3,928 points in the multi-core test, highlighting its strong performance in terms of processing power for multitasking and demanding computational tasks. Single-core performance was measured at 69 points, while the GPU benchmark reached a score of 3,634 points, demonstrating its mastery of graphics processing. The MP ratio was calculated at 57.27, demonstrating a consistent balance between multi-threaded and single-threaded processing capabilities.
| Cinebench 2024 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| CPU (multi-core) (Points) | 3,928 | 4,913 | 
| CPU (single-core) (Points) | 69 | 81 | 
| GPU | 3,634 | N/A | 
| MP Ratio | 57.27 | 60.47x | 
Geekbench CPU Benchmark
Geekbench 6 is a cross-platform performance evaluation tool that measures the overall performance of a system. However, it would be interesting to compare single-core and multi-core performance, as well as OpenCL performance. A high score indicates better performance.
Here, the system demonstrated solid overall performance in terms of CPU and GPU metrics. The single-core processor benchmark recorded a score of 1,718, while the multi-core processor test showed 19,455. Additionally, GPU performance was evaluated using the OpenCL framework, with a score of 35,764, indicating solid graphics processing power suitable for various compute-intensive applications.
You can find comparisons with any system in the Geekbench browser.
| Geekbench 6 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| CPU Benchmark - Single-Core | 1,718 | 2,055 | 
| CPU Benchmark - Multi-Core | 19,455 | 22,868 | 
| GPU Benchmark – OpenCL | 35,764 | N/A | 
y-cruncher
y-cruncher is a multi-threaded and scalable program that can calculate Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress testing application for overclockers and hardware enthusiasts.
In terms of results, we have results ranging from 1 billion to 25 billion for the Supermicro 1115SV-WTNRT.
| y-cruncher (Total Computation Time) (lower is better) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| 1 billion digits (seconds) | 9.989 seconds | 7.274 seconds | 
| 2.5 billion digits (seconds) | 24.974 seconds | 17.055 seconds | 
| 5 billion digits (seconds) | 52.117 seconds | 34.336 seconds | 
| 10 billion digits (seconds) | 110.483 seconds | 71.336 seconds | 
| 25 billion digits (seconds) | 303.372 seconds | 196.695 seconds | 
| 50 billion digits (seconds) | N/A | 439.435 seconds | 
UL Procyon AI Inference
