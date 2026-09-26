---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb-3
title: "fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "decode", "gpu", "inference", "intel", "latency", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb.md
source_anchor: ""
source_lines: [46, 123]
sha256: aa9180d52ea5a2a89edfbe5d13c1e224c8a00137db739da8b79d7c914c6545f7
---

# fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb

In our tests, we compare the performance of the Supermicro Hyper 1U SYS-112H-TN with the results from our initial review of the Intel launch server equipped with the same Xeon 6 processor family. This comparison is useful because it examines the performance of these processors across different server architectures, allowing us to provide more valuable benchmark results. It can indicate the impact of server design and configuration on overall performance and efficiency in real-world applications.
Blender OptiX 4.0 / 4.1
Blender is an open-source 3D modeling application. This benchmark was run using the Blender Benchmark utility. The score is expressed in samples per minute, with the higher being the better.
The Supermicro Hyper 1U 112H-TN loaded with a single Xeon 6780E delivers solid performance but is, as expected, below the Intel (launch) server equipped with two Xeon processors. For example, in the "Monster" scene, the Supermicro configuration achieves 781.42 samples per minute, while the dual Xeon 6780E configuration achieves 1,410.463 samples per minute. This shows that while the single-processor configuration is competent for 3D rendering, the dual-processor configuration nearly doubles the performance, making it better suited for more demanding 3D workloads. The Supermicro would likely outperform the Intel launch server if equipped with two processors.
| Blender 4.0 CPU | Supermicro Hyper 1U 112H-TN (1x Xeon 6780E, 512 GB DDR5) | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 
| Monster | 781.42 | 1410.463 | 1297.715 | 
| Junkshop | 514.658 | 862.418 | 777.716 | 
| Classroom | 370.52 | 696.543 | 628.960 | 
| Blender 4.1 CPU |  |  |  | 
| Monster | 764.112 | N/A | N/A | 
| Junkshop | 511.872 | N/A | N/A | 
| Classroom | 363.672 | N/A | N/A | 
Geekbench 6
Geekbench 6 is a cross-platform performance evaluation tool that measures the overall performance of a system. The Geekbench browser allows you to compare any system with this tool.
The single-core CPU score of 1,154 reflects decent single-threaded performance, sufficient for tasks that depend on individual core speed. The multi-core CPU score of 15,167 highlights the Xeon 6780E's ability to efficiently handle multithreaded tasks. This makes the system well-suited for workloads that leverage multiple cores, such as virtualization, database management, and multi-user environments. While these scores are solid, they suggest that the Xeon 6780E in this configuration is more oriented toward parallel processing tasks rather than excelling in latency-sensitive single-core applications.
| Geekbench 6 (Higher is better) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | 
| Single-Core CPU | 1,154 | 
| Multi-Core CPU | 15,167 | 
Cinebench R23
The Cinebench R23 benchmark tool evaluates a system's CPU performance by rendering a complex 3D scene using the Cinema 4D engine. It measures both single-core and multi-core performance, providing a comprehensive view of the CPU's capabilities in handling 3D rendering tasks.
The Supermicro Xeon configuration scored 92,516 points in the multi-core test, which is more than the original Intel launch dual Xeon 6780E platform (67,984 pts).
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 
| Multi-Core CPU | 92,516 | 67,984 pts | 64,326 pts | 
| Single-Core CPU | 888 pts | 873 pts | 793 pts | 
| MP Ratio | 104.20 x | 77.91 x | 81.10 x | 
Cinebench 2024
Cinebench 2024 extends the benchmarking capabilities of R23 by adding a GPU performance evaluation. It continues to test CPU performance but also includes tests that measure the GPU's ability to handle rendering tasks.
The Supermicro Hyper 112H-TN scores 2,941 pts in multi-core, surpassing the dual Xeon 6780E configuration, which scored 2,687 pts. This is another example where the single-processor configuration shows robust optimization in the Supermicro server, delivering superior performance in specific multithreaded applications.
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 
| Multi-Core CPU | 2,565 pts | 2,687 pts | 2,347 pts | 
| Single-Core CPU | 53 pts | 43 pts | 45 pts | 
| MP Ratio | 48.38 x | 62.85 x | 52.65 x | 
Y-Cruncher
Y-cruncher is a popular benchmarking and stress testing application launched in 2009. This test is multithreaded and scalable, calculating Pi and other constants to billions of digits. Faster is better in this test.
In this test, the Supermicro configuration was slower, especially in larger calculations, compared to the Intel launch server equipped with two processors. For example, calculating Pi to 50 billion digits takes 680.090 seconds on the Supermicro configuration, while the dual Xeon 6780E configuration completes it in 565.913 seconds.
| Y-Cruncher (0.8.3.9) (lower is better) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Xeon 6780E (256 GB DDR5) | Xeon 6766E (256 GB DDR5) | 
| 1 billion | 9.754 seconds | 6.927 seconds | 7.254 seconds | 
| 2.5 billion | 25.215 seconds | 17.898 seconds | 19.507 seconds | 
| 5 billion | 55.242 seconds | 38.454 seconds | 41.116 seconds | 
| 10 billion | 118.657 seconds | 81.146 seconds | 87.403 seconds | 
| 25 billion | 315.085 seconds | 217.530 seconds | 238.813 seconds | 
| 50 billion | 680.090 seconds | 565.913 seconds | 502.245 seconds | 
Here are the results for version 0.8.5.9:
| Y-Cruncher (0.8.5.9) (lower is better) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | 
| 1 billion | 8.757 seconds | 
| 2.5 billion | 24.928 seconds | 
| 5 billion | 53.489 seconds | 
| 10 billion | 113.727 seconds | 
| 25 billion | 308.218 seconds | 
| 50 billion | 674.299 seconds | 
Blackmagic Disk Speed Test
In the Blackmagic Disk Speed Test, the Micron 7450 NVMe SSD produced read speeds of 3,627.4 MB/s and write speeds of 2,849.5 MB/s.
Blackmagic RAW Speed Test
The Blackmagic RAW Speed Test is a performance benchmarking tool designed to measure a system's ability to handle video playback and editing using the Blackmagic RAW codec. It evaluates a system's ability to decode and play high-resolution video files, providing frame rates for both CPU-based and GPU-based processing.
The Supermicro Hyper 112H-TN achieved 116 FPS with an 8K CPU. We did not perform the GPU portion of this test.
7-Zip
The built-in memory benchmark of the popular 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, providing an indication of the system's ability to handle data-intensive operations.
Here, the Supermicro Hyper 112h-tn achieves a total score of 205.449 GIPS, significantly lower than the 271.217 GIPS achieved by the dual Xeon 6780E configuration.
| 7-Zip Compression | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Xeon 6780E (256 GB DDR5) | Xeon 6766E (256 GB DDR5) | 
| Compression |  |  |  | 
| Current CPU Usage | 5287% | 5,891% | 4,768% | 
| Current Rating/Usage | 4.647 GIPS | 4.985 GIPS | 4.614 GIPS | 
| Current | 245.699 GIPS | 293.689 GIPS | 220.001 GIPS | 
| Resulting CPU Usage | 5296% | 5,603% | 5,103% | 
| Resulting Rating/Usage | 4.642 GIPS | 4.954 GIPS | 4.638 GIPS | 
| Resulting Rating | 245.823 GIPS | 277.670 GIPS | 236.910 GIPS | 
| Decompression |  |  |  | 
| Current CPU Usage | 6236% | 5,962% | 5,798% | 
| Current Rating/Usage | 4.261 GIPS | 4.550 GIPS | 4.152 GIPS | 
| Current | 265.709 GIPS | 271.266 GIPS | 240.693 GIPS | 
| Resulting CPU Usage | 6236% | 5,832% | 6,029% | 
| Resulting Rating/Usage | 4.341 GIPS | 4.540 GIPS | 4.161 GIPS | 
| Resulting Rating | 269.373 GIPS | 264.764 GIPS | 250.853 GIPS | 
| Total Rating |  |  |  | 
| Total CPU Usage | 5751% | 5,717% | 5,566% | 
| Total Rating/Usage | 4.491 GIPS | 4.747 GIPS | 4.399 GIPS | 
| Total Rating | 257.598 GIPS | 271.217 GIPS | 243.882 GIPS | 
UL Procyon AI Inference
