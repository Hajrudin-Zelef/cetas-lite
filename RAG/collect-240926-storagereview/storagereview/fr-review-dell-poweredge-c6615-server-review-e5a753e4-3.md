---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-c6615-server-review-e5a753e4-3
title: "fr-review-dell-poweredge-c6615-server-review-e5a753e4"
domain: storagereview
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "gpu", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-c6615-server-review-e5a753e4.md
source_anchor: ""
source_lines: [58, 142]
sha256: 3d514eadd5294cbd9c9b77a9efb046bc9b584ee9acef7d96317369650b955a95
---

# fr-review-dell-poweredge-c6615-server-review-e5a753e4

In Cinebench R23, the four nodes were around 74,000 on the multi-core portion, with node 3 slipping into 75,000. All four nodes remained much closer for single-core scores, with nodes 1 and 4 at 1,088. Node 2 was only 5 points behind and node 3 was 8 points ahead. Overall, all nodes showed only minor performance variations, typical of different processors, even though they all belong to the same model.
| Cinebench R23 | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Multi-core CPU | 74,877 | 74,961 | 75,011 | 74,745 | 74,898.5 | 
| Single-core CPU | 1,088 | 1,093 | 1,084 | 1,088 | 1,088.25 | 
| MP Ratio | 64.84 | 68.60 | 69.17 | 68.70 | 67.83 | 
Cinebench 2024
Maxon's Cinebench 2024 is a CPU and GPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Since these nodes do not have a GPU, we only have the multi-core and single-core numbers.
In Cinebench 2024, all nodes remained close to each other, with minimal variance on the multi-core and single-core portions. The average performance was 4,509 points for multi-core and 67.25 points for single-core, with an MP ratio of 66.98.
| Cinebench 2024 | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Multi-core CPU | 4,544 | 4,577 | 4,436 | 4,481 | 4,509.5 | 
| Single-core CPU | 68 | 68 | 65 | 68 | 67.25 | 
| MP Ratio | 66.79 | 67.23 | 68.21 | 65.69 | 66.98 | 
Geekbench 6 CPU
Geekbench 6 is a cross-platform benchmark that measures the overall performance of a system. This test includes a dedicated CPU portion and a GPU portion, but since these nodes do not have a GPU, we only have the CPU results. The higher the score, the better the performance.
In Geekbench, we saw tight numbers until we got to node 3, which dipped slightly on single-core and multi-core. The average across all nodes was 1,687 single-core and 19,319.5 multi-core.
| Geekbench 6 CPU | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Single-Core | 1,707 | 1,708 | 1,625 | 1,708 | 1,687 | 
| Multi-Core | 19,544 | 19,234 | 18,999 | 19,501 | 19,319.5 | 
Blender 4.0 CPU
The next step is Blender OptiX, an open-source 3D modeling application. This benchmark was run using the Blender Benchmark CLI utility. The score is expressed in samples per minute, with the highest being the best.
The C6615 nodes recorded fairly consistent numbers. The average scores were 591.79 on Monster, 415.88 on Junkshop, and 311.74 on Classroom.
| Blender 4.0 CPU | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Monster | 595.23 | 593.51 | 584.35 | 594.07 | 591.79 | 
| Junkshop | 415.26 | 415.11 | 418.05 | 415.08 | 415.88 | 
| Classroom | 308.57 | 312.91 | 312.69 | 312.78 | 311.74 | 
Blender 4.1 CPU
Blender OptiX 4.1 brings new features, such as GPU-accelerated denoising, streamlining the rendering process and reducing the time required for denoising tasks. Despite these advances, the overall performance improvements in benchmark scores compared to version 4.0 are minimal, indicating only slight improvements in efficiency.
Again, we see consistent numbers across the board, with averages of 587.22 on Monster, 420.20 on Junkshop, and 306.60 on Classroom.
| Blender 4.1 CPU | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Monster | 590.46 | 590.58 | 584.76 | 583.08 | 587.22 | 
| Junkshop | 418.38 | 416.71 | 426.73 | 419.03 | 420.20 | 
| Classroom | 306.86 | 304.81 | 308.95 | 305.79 | 306.60 | 
7-Zip Compression
The popular 7-Zip utility has a built-in memory test that demonstrates CPU performance. In this test, we run it at a dictionary size of 128 MB when possible.
Fair scores were observed across all nodes. In the total scores, we found a total CPU usage of 5,778.75%, a total rating/usage of 4.355 GIPS, and a total rating of 252 GIPS.
| Blender 4.1 CPU | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Compression |  |  |  |  |  | 
| Current CPU Usage | 5,548% | 5,549% | 5,633% | 5,585% | 5,578.75% | 
| Current Rating/Usage | 4.256 GIPS | 4.210 GIPS | 4.156 GIPS | 4.177 GIPS | 4.20 GIPS | 
| Current | 236.158 GIPS | 233.626 GIPS | 234.092 GIPS | 233.285 GIPS | 234.290 GIPS | 
| Resulting CPU Usage | 5,536% | 5,537% | 5,601% | 5,553% | 5,556.75% | 
| Resulting Rating/Usage | 4.193 GIPS | 4.202 GIPS | 4.172 GIPS | 4.168 GIPS | 4.184 GIPS | 
| Resulting Rating | 232.118 GIPS | 232.631 GIPS | 233.691 GIPS | 231.443 GIPS | 232.470 GIPS | 
| Decompression |  |  |  |  |  | 
| Current CPU Usage | 5,973% | 6,027% | 5,992% | 6,014% | 6,001.5% | 
| Current Rating/Usage | 4.543 GIPS | 4.501 GIPS | 4.565 GIPS | 4.509 GIPS | 4.530 GIPS | 
| Current | 271.343 GIPS | 271.287 GIPS | 273.507 GIPS | 271.196 GIPS | 271.833 GIPS | 
| Resulting CPU Usage | 5,997% | 6,015% | 5,999% | 5,990% | 6,000.25% | 
| Resulting Rating/Usage | 4.537 GIPS | 4.519 GIPS | 4.550 GIPS | 4.499 GIPS | 4.526 GIPS | 
| Resulting Rating | 272.066 GIPS | 271.775 GIPS | 272.946 GIPS | 269.509 GIPS | 271.574 GIPS | 
| Total Rating |  |  |  |  |  | 
| Total CPU Usage | 5,767% | 5,776% | 5,800% | 5,772% | 5,778.75% | 
| Total Rating/Usage | 4.365 GIPS | 4.360 GIPS | 4.361 GIPS | 4.333 GIPS | 4.355 GIPS | 
| Total Rating | 252.092 GIPS | 252.203 GIPS | 253.318 GIPS | 250.476 GIPS | 252.022 GIPS | 
Blackmagic Raw Speed Test
We use the Blackmagic Raw Speed Test to evaluate how machines perform actual RAW decoding. This test can integrate both CPU and GPU usage, but we will only test CPU usage.
All four nodes showed extremely close performance, with an average of 119.75 FPS.
| Blackmagic Raw Speed Test | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| CPU 8K | 121 FPS | 121 FPS | 118 FPS | 119 FPS | 119.75 FPS | 
Blackmagic Disk Speed Test
Next is the Blackmagic Disk Speed Test. This test runs a 5 GB file sample for read and write speeds. Since it is single-threaded, it will not show the highest disk speeds, but it still gives a good perspective.
The C6615s have a BOSS card inside, using two M.2 drives in RAID1, so performance is slightly degraded for reliability. For write speeds, we found an average of 991.6 MB/s and for read speeds, an average of 2,801 MB/s.
| Blackmagic Disk Speed Test | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Write | 999.8 MB/s | 977.4 MB/s | 991.4 MB/s | 997.7 MB/s | 991.6 MB/s | 
| Read | 2,807.4 MB/s | 2,790.1 MB/s | 2,828.0 MB/s | 2,780.4 MB/s | 2,801.5 MB/s | 
Y-Cruncher
y-cruncher is a multi-threaded and scalable program that can calculate Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress testing application for overclockers and hardware enthusiasts.
For our average speeds, we saw 9.5 seconds for 1 billion, 24.20 seconds for 2.5 billion, and 50.73 seconds for 5 billion. On the most significant digit calculations, we saw 105.73 seconds for 10 billion, 288.85 seconds for 25 billion, and 633.5 seconds for 50 billion.
| Y Cruncher (total computation time, in seconds) | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| 1 billion | 9.587 | 9.459 | 9.350 | 9.633 | 9.507 | 
| 2.5 billion | 24.490 | 24.225 | 23.334 | 24.740 | 24.197 | 
| 5 billion | 51.427 | 50.990 | 49.303 | 51.214 | 50.734 | 
| 10 billion | 107.084 | 107.646 | 103.772 | 107.443 | 105.736 | 
| 25 billion | 291.918 | 290.944 | 280.632 | 291.902 | 288.849 | 
| 50 billion | 641.709 | 640.289 | 619.100 | 640.917 | 635.504 | 
UL Procyon AI Computer Vision Benchmark
