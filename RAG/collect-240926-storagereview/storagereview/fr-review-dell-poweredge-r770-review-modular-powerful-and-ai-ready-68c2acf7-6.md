---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-6
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "compute", "gpu", "intel", "nvidia"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [133, 197]
sha256: 2def66ad1c3d30d70a1a3ef8049a203cc537d3159cf9d2b17889cce70b34e455
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

In the 7-Zip benchmark, the Dell system achieved a higher score (266.425 GIPS) than Lenovo (224.313 GIPS) for compression tasks, with the Dell system showing slightly lower CPU utilization. However, Lenovo outperformed Dell in decompression, with a higher score (288.457 GIPS vs. 256.154 GIPS) and slightly higher CPU utilization. Dell achieved a slightly higher overall score (261.290 GIPS), demonstrating better overall efficiency for compression and decompression tasks.
| 7-Zip Compression and Decompression | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| Compression – Current CPU usage | 5267 % | 5064 % | 
| Compression – Current rating/usage | 5.061 GIPS | 4.341 GIPS | 
| Compression – Current rating | 266.591 GIPS | 219.840 GIPS | 
| Compression – Resulting CPU usage | 5270 % | 5156 % | 
| Compression – Resulting rating/usage | 5.056 GIPS | 4.350 GIPS | 
| Compression – Resulting rating | 266.425 GIPS | 224.313 GIPS | 
| Decompression – Current CPU usage | 5623 % | 6184 % | 
| Decompression – Current rating/usage | 4.586 GIPS | 4.688 GIPS | 
| Decompression – Current rating | 257.909 GIPS | 289.879 GIPS | 
| Decompression – Resulting CPU usage | 5627 % | 6205 % | 
| Decompression – Resulting rating/usage | 4.553 GIPS | 4.649 GIPS | 
| Decompression – Resulting rating | 256.154 GIPS | 288.457 GIPS | 
| Total – Total CPU usage | 5448 % | 5681 % | 
| Total – Total rating/usage | 4.804 GIPS | 4.500 GIPS | 
| Total – Total rating | 261.290 GIPS | 256.385 GIPS | 
y-cruncher
y-cruncher is a popular benchmarking and stress-testing application launched in 2009. This test is multithreaded and scalable, computing Pi and other constants to thousands of billions of digits. Faster is better in this test. This software has been fantastic for testing high-core-count platforms and showing the compute advantages between single- and dual-socket platforms.
The Y-Cruncher benchmark results show a significant performance gap between the Dell PowerEdge R770, equipped with P-core processors, and the Lenovo ThinkSystem SR630 V4 equipped with E-core processors, particularly as dataset size increases. This is less about determining which system is better than comparing processor types under this workload.
For smaller-scale computations, the Dell system was already ahead, computing 1 billion digits of Pi in 2.753 seconds, while the Lenovo system took more than double, at 5.997 seconds. As the workload increased, the gap widened. At 10 billion digits, the Dell system finished in 34.873 seconds, less than half of the Lenovo system's 81.046 seconds. Beyond 50 billion digits, Dell maintained its lead, completing the task in 221.255 seconds, versus 476.826 seconds for Lenovo, making Dell 53% faster.
At 100 billion digits, Lenovo could not complete the test due to its current 512 GB RAM configuration. With 2 TB of RAM, Dell handled the workload efficiently, finishing in 491.737 seconds.
| Y-cruncher (shorter time is better) | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| 1 billion | 2.753 seconds | 5.997 seconds | 
| 2.5 billion | 7.365 seconds | 17.573 seconds | 
| 5 billion | 16.223 seconds | 37.793 seconds | 
| 10 billion | 34.873 seconds | 81.046 seconds | 
| 25 billion | 99.324 seconds | 220.025 seconds | 
| 50 billion | 221.255 seconds | 476.826 seconds | 
| 100 billion | 491.737 seconds |  | 
OptiX Blender
An open-source 3D modeling application. This benchmark was performed with the Blender Benchmark utility. The score is expressed in samples per minute, with the highest being the best.
The Blender benchmark results show a clear performance advantage for the Dell PowerEdge R770 over the Lenovo ThinkSystem SR630 V4, particularly in CPU rendering. In the "CPU Monster" test, Dell reached 1,706.002 samples per minute, a 19% lead over Lenovo's 1,432.09 samples per minute. The "CPU Junkshop" test further widened this gap, with Dell reaching 1,169.370 samples per minute, surpassing Lenovo's 914.75 samples per minute by 28%. Similarly, Dell recorded 791.475 samples per minute in the "CPU Classroom" test, while Lenovo lagged at 656.68 samples per minute, a difference of 20%.
The absence of a GPU in the Lenovo system also meant it could not participate in GPU-based rendering, where Dell's NVIDIA L4 scored 1,895.71 samples/min for Monster, 950.42 samples/min, and a Classroom score of 968.43 samples/min.
| Blender CPU Benchmark | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| CPU Monster (Blender 4.3) | 1,706.002 samples/min | 1432.09 samples/min | 
| CPU Junkshop (Blender 4.3) | 1,169.370 samples/min | 914.75 samples/min | 
| CPU Classroom (Blender 4.3) | 791.475 samples/min | 656.68 samples/min | 
| GPU Monster (Blender 4.3) | 1,895.712 samples/min | (no GPU) | 
| GPU Junkshop (Blender 4.3) | 950.424 samples/min | (no GPU) | 
| GPU Classroom (Blender 4.3) | 968.432 samples/min | (no GPU) | 
Cinebench R23
The Cinebench R23 benchmark tool evaluates a system's CPU performance by rendering a complex 3D scene using the Cinema 4D engine. It measures single-core and multi-core performance, providing a comprehensive view of the processor's capabilities in handling 3D rendering tasks.
In Cinebench R23, the benchmark results highlight notable CPU performance differences between the Dell PowerEdge R770 and the Lenovo ThinkSystem SR630 V4, particularly in terms of cores per processor. The Lenovo ThinkSystem SR630 V4, equipped with two Intel Xeon 6780E processors (144 cores per processor), outperformed Dell in the multi-core CPU test with a score of 99,266 points, versus 74,710 points for Dell. This gap reflects Lenovo's advantage in multithreaded workloads, thanks to its higher core count (288 cores total) compared to Dell's two Intel Xeon 6787P processors (86 cores per processor), which limits its multi-core performance.
In the Single-Core CPU test, Dell performed better with a score of 1,272 points, surpassing Lenovo's 894 points, highlighting Dell's superior single-threaded processor efficiency despite its lower core count.
| Cinebench R23 | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| Multi-core CPU | 74,710 pts | 99,266 pts | 
| Single-core CPU | 1,272 pts | 894 pts | 
| MP ratio | 58.74 x | 111.00 x | 
Cinebench 2024
Cinebench 2024 extends R23's benchmarking capabilities by adding GPU performance evaluation. It continues to test CPU performance but also includes tests that measure the GPU's ability to handle rendering tasks.
In this updated benchmark, the Dell PowerEdge R770 scored 12,996 points for GPU performance, highlighting its ability to handle GPU-accelerated rendering tasks. The Lenovo ThinkSystem SR630 V4 has no dedicated GPU and therefore recorded no GPU score.
In the multi-core CPU test, Lenovo scored 2,884 points, slightly ahead of Dell's 2,831 points, demonstrating a slight advantage in multi-core performance. In single-core CPU, Dell outperformed Lenovo, scoring 71 points versus 53 points for Lenovo, demonstrating Dell's superior single-core performance despite a reduced core count.
| Cinebench R24 | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| GPU score | 12,996 pts |  | 
| Multi-core CPU | 2,831 pts | 2,884 pts | 
| Single-core CPU | 71 pts | 53 pts | 
| MP ratio | 39.77 x | 54.43 x | 
Geekbench 6
Geekbench 6 is a cross-platform benchmarking tool that measures a system's overall performance. The Geekbench browser allows comparison of any system with this tool.
