---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662-4
title: "fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["memory", "amd", "benchmark", "gpu"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662.md
source_anchor: ""
source_lines: [93, 132]
sha256: 709dcbcf705d1cf28a35603956751e8f58b76e46664babc94fcf4bfb32253439
---

# fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662

The performance evaluation of the Dell PowerEdge R7715 will focus primarily on its processor: the AMD EPYC 9665P. This 96-core, 192-thread processor is designed for a single-socket (1P) configuration and offers a maximum boost frequency of 4.5 GHz, an all-core boost frequency of 4.1 GHz, and a base frequency of 2.6 GHz. With 384 MB of L3 cache and a configurable TDP from 320 W to 400 W, the EPYC 9665P is designed to excel in multi-core workloads, delivering an optimal balance between speed, efficiency, and scalability.
Other key system specifications to note include:
- Storage: 32 x 1.6 TB E1.S SSDs
- Memory: 768 GB of RAM
We compared its performance to that of other systems equipped with models from the same EPYC 9005 series. Among these is the EPYC 9665, a high-end model with 192 cores and 384 threads, available in single- and dual-processor configurations (1P/2P). The EPYC 9965, however, has lower clock frequencies (3.7 GHz in maximum boost mode and 3.35 GHz in all-core boost mode), but compensates for this weakness with twice the number of cores as the 9665. The EPYC 9755 is another comparable model, offering 128 cores and 256 threads. It has a slightly lower maximum boost frequency (4.1 GHz) and a lower base frequency (2.7 GHz), but a larger L3 cache (512 MB) and a higher default TDP (500 W). It is also available in single- and dual-processor configurations. Finally, the EPYC 9575F is a high-frequency processor designed for demanding workloads. With 64 cores and 128 threads, it offers a maximum turbo frequency of 5 GHz and an all-core turbo frequency of 4.5 GHz. All processors in the EPYC 9005 series were tested in single-processor configuration for this test.
Additionally, the comparative tests include the Genoa (2 processors/96 cores) and Bergamo (2 processors/128 cores) models, which leverage dual-socket configurations. The comparison with these dual-socket multi-core systems will highlight its scalability against larger configurations in terms of processor count, efficiency, clock frequency, and processing power.
Geekbench 6
The cross-platform Geekbench 6 benchmark measures a system's performance and provides a comparison score. Designed to run on multiple platforms, it provides consistent performance measurements across different devices, including smartphones, tablets, desktops, and servers.
With a single-core score of 2,806, it far surpasses other models, including the EPYC 9965, 9755, and even the high-frequency 9575F. This is likely due to its high maximum boost clock frequency of 4.5 GHz and its all-core boost speed of 4.1 GHz, among the highest in the lineup. The multi-core score of 28,645 is also higher, despite having fewer cores than the 9965 or 9755. Compared to dual-processor systems like Genoa (2P/96c) and Bergamo (2P/128c), it offers better single-core and multi-core performance.
| Geekbench 6 | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| Single Core | 2,806 | 1,453 | 1,641 | 1,865 | 2,048 | 1,723 | 
| Multi-Core | 28,645 | 11,199 | 11,800 | 13,219 | 20,217 | 17,916 | 
Blackmagic RAW Speed Test
We performed the Blackmagic RAW Speed Test to evaluate the PowerEdge R7715's ability to handle Blackmagic RAW decoding tasks using only the processor, without a GPU installed. This test measures performance at different resolutions and compression levels.
The R7715 achieved a score of 171 FPS for 8K content, demonstrating fairly solid CPU performance for high-resolution video processing, even without GPU acceleration.
Maxon Cinebench
Cinebench is a widely used benchmarking tool that measures the performance of processors and graphics cards (CPU) using Maxon Cinema 4D for rendering. It provides a score to compare the performance of different systems and components. We tested four popular versions of Cinebench so you can compare results on the most popular online rankings.
Here, the AMD EPYC 9665P demonstrated its high-performance single-socket processor capabilities. With a multi-core score of 121,254 points in Cinebench R23, it stands out from comparable systems, slightly ahead of the 131,846-point score of the EPYC 9755, but surpassing the Genoa (2P/96c) and Bergamo (2P/128c) systems. However, the most remarkable feature here is the single-core performance, where the 9665P scored 1,845 points, significantly higher than all other processors tested.
In Cinebench 2024, the 9665P continues to excel, scoring 7,501 points in multi-core tests, surpassing the 9965, 9755, and 9575F. Its single-core score of 109 points is also impressive, well ahead of the 9965 (77 points) and 9755 (84 points). The 9665P offers a perfect balance between single-core speed and multi-core power.
| Test | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) |  | 
| Cinebench R23 |  |  |  |  |  |  |  | 
| CPU (multi-core) | 121,254 pts | N/A | 131,846 pts | 111,149 pts | 116,744 pts | 102,125 points |  | 
| CPU (single-core) | 1,845 pts | N/A | 1,400 pts | 1,052 pts | 1,294 pts | 1,089 points |  | 
| Cinebench 2024 |  |  |  |  |  |  |  | 
| CPU (multi-core) | 7,501 pts | 4,845 pts | 5,921 pts | 4,324 | N/A | N/A |  | 
| CPU (single-core) | 109 pts | 77 pts | 84 pts | 103 pts | N/A | N/A |  | 
y-cruncher
y-cruncher is a multithreaded and scalable program capable of calculating Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application among overclockers and hardware enthusiasts.
Despite having fewer cores than the other processors tested, the 9665P showed solid performance on most calculations. Its time of 6.836 seconds for the 1 billion digit calculation is respectable, even though it is slower than the AMD EPYC 9575F (64c), which achieved the best time at 4.476 seconds. As the number of digits increases, the 9665P maintains competitive performance, particularly excelling in the 10 billion digit test, which it completes in 51.851 seconds. This result is comparable to other processors with higher core counts, such as the 9965 (41.750 seconds) and 9755 (41.512 seconds).
However, for higher digit counts (25 and 50 billion in this case), the 9665P takes 140.339 and 303.842 seconds respectively. This likely means that while the 9665P handles heavy computational tasks well, it begins to lose ground as workload complexity increases, particularly compared to processors with higher core counts. Nevertheless, the 9665P remains a serious contender for most workloads and offers excellent performance for its core count.
| y-cruncher Total Computation Time (Lower is Better) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, 128t) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| 1 billion | 6.836 seconds | 7.346 seconds | 7.747 seconds | 5.408 seconds | 4.476 seconds | 8.882 seconds | 9.184 seconds | 
| 2.5 billion | 13.720 seconds | 13.661 seconds | 14.113 seconds | 11.376 seconds | 10.067 seconds | N/A | N/A | 
| 5 billion | 25.795 seconds | 23.211 seconds | 22.820 seconds | 20.177 seconds | 20.030 seconds | N/A | N/A | 
| 10 billion | 51.851 seconds | 41.750 seconds | 41.512 seconds | 40.767 seconds | 41.518 seconds | 51.071 seconds | 55.683 seconds | 
| 25 billion | 140.339 seconds | 115.091 seconds | 98.981 seconds | 103.650 seconds | 104.737 seconds | N/A | N/A | 
| 50 billion | 303.842 seconds | N/A | N/A | N/A | N/A | N/A | N/A | 
| 100 billion | 707.391 seconds | N/A | N/A | N/A | N/A | N/A | N/A | 
Blender OptiX
Blender OptiX is an open-source 3D modeling application. This benchmark was performed using the Blender Benchmark CLI utility. The score is measured in samples per minute, with higher values being better.
