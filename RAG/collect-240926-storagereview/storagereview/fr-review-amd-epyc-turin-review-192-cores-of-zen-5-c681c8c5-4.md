---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5-4
title: "fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "compute", "inference", "latency", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5.md
source_anchor: ""
source_lines: [76, 126]
sha256: 2da6cc3f8866b0aed50c4c7e853a161873997cbf8ed72cf4446977fc4500a5a7
---

# fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5

Here, the EPYC 9575F (SMT disabled) achieves the fastest computation times for 1 and 2.5 billion digit calculations, at 4.476 and 10.067 seconds respectively. This speed even surpasses the high-core 9965 and 9755 models, demonstrating its efficiency for high-precision tasks that do not require an extreme number of threads.
The 9965 and 9755 processors also offer good times, notably surpassing those of Genoa and Bergamo at 10 billion digits, in just over 41 seconds versus 51 seconds for Genoa. This shows that Turin-based EPYC processors are well optimized for compute-intensive applications, demonstrating scalable performance that allows for more precise control of thread and resource management. The 64-core 9575F processor really shone here with its incredibly high clock speed, showing a strong lead over the cores of the 9755 or 9965 samples in the y-cruncher ranges we examined.
| y-cruncher Total Computation Time (Lower is Better) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, 128t) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| 1 billion | 7.346 seconds | 7.747 seconds | 5.408 seconds | 4.476 seconds | 8.882 seconds | 9.184 seconds | 
| 2.5 billion | 13.661 seconds | 14.113 seconds | 11.376 seconds | 10.067 seconds | N/A | N/A | 
| 5 billion | 23.211 seconds | 22.820 seconds | 20.177 seconds | 20.030 seconds | N/A | N/A | 
| 10 billion | 41.750 seconds | 41.512 seconds | 40.767 seconds | 41.518 seconds | 51.071 seconds | 55.683 seconds | 
| 25 billion | 115.091 seconds | 98.981 seconds | 103.650 seconds | 104.737 seconds | N/A | N/A | 
This Y-Cruncher performance test uses the Bailey-Borwein-Plouffe (BBP) formulas to calculate massive hexadecimal digits of Pi, measuring the CPU's total computation time, utilization, and multi-core efficiency. The results provide a detailed analysis of the performance of AMD Turin processors in these intensive calculations.
The EPYC 9575F with SMT disabled excels in single-thread efficiency, completing 1 BBP in just 0.179 seconds with a multi-core efficiency of 44.73%. This demonstrates its efficiency in scenarios where single-thread optimization is a priority. By comparison, the 9965, with its 192 cores, shows a lower multi-core efficiency of 4.61% in this single-BBP test, as its high core count introduces latency in low-thread scenarios.
Multi-core efficiency becomes more evident with higher workloads, such as 10 and 100 BBP. The EPYC 9575F with SMT disabled maintains a strong lead with 86.59% efficiency at 10 BBP, completing it in 0.896 seconds. For 100 BBP, the 9575F continues to excel, reaching nearly 99% efficiency and displaying balanced performance in highly parallel workloads, completing the run in 8.065 seconds. In this test, the 9965 achieves 85.10% efficiency, while the 9755 and 9575 maintain competitive multi-core efficiencies of 94% and 96.83%, respectively, demonstrating that Turin-based EPYC processors can efficiently handle both low-thread and high-thread workloads.
| Benchmark | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | AMD EPYC9575F (64c, SMT disabled) | 
| 1 BBP |  |  |  |  | 
| 10 BBP |  |  |  |  | 
| 100 BBP |  |  |  |  | 
The memory test built into the 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, indicating how well the system can handle data-hungry operations. We run this test with a dictionary size of 128 MB when possible.
The 7-Zip compression test demonstrates the efficiency of the 9755, achieving 443.029 GIPS at 5613% CPU utilization, the highest in the EPYC lineup. This level of performance makes the 9755 ideal for intensive data compression tasks, such as those in backup or archiving solutions, where fast data handling is essential. The 9575F (SMT disabled) also performs well, achieving 394.9 GIPS. These results indicate that the 9755 and 9575F have solid compression and decompression capabilities, with high CPU throughput and efficient resource management for applications that leverage high compression ratios or process large volumes of data.
| 7-Zip Compression Benchmark (Higher is Better) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c, SMT disabled) |  | 
| Compression |  |  |  |  | 
| Current CPU Usage | 4302% | 5233% | 4406% |  | 
| Current Rating/Usage | 5.830 GIPS | 7.597 GIPS | 7.975 GIPS |  | 
| Current | 250.827 GIPS | 397.536 GIPS | 351.358 GIPS |  | 
| Resulting CPU Usage | 4041% | 5306% | 4555% |  | 
| Resulting Rating/Usage | 5.804 GIPS | 7.720 GIPS | 8.070 GIPS |  | 
| Resulting Rating | 234.317 GIPS | 409.652 GIPS | 367.358 GIPS |  | 
| Decompression |  |  |  |  | 
| Current CPU Usage | 4322% | 6041% | 5017% |  | 
| Current Rating/Usage | 7.078 GIPS | 8.065 GIPS | 8.483 GIPS |  | 
| Current | 305.909 GIPS | 487.263 GIPS | 425.580 GIPS |  | 
| Resulting CPU Usage | 4556% | 5921% | 4940% |  | 
| Resulting Rating/Usage | 6.577 GIPS | 8.045 GIPS | 8.569 GIPS |  | 
| Resulting Rating | 299.163 GIPS | 476.405 GIPS | 422.441 GIPS |  | 
| Total Rating |  |  |  |  | 
| Total CPU Usage | 4298% | 5613% | 4747% |  | 
| Total Rating/Usage | 6.190 GIPS | 7.883 GIPS | 8.319 GIPS |  | 
| Total Rating | 266.740 GIPS | 443.029 GIPS | 394.900 GIPS |  | 
Blender OptiX is an open-source 3D modeling application. This benchmark was run using the Blender Benchmark CLI utility. The score is expressed in samples per minute, with the highest being the best.
Here, the 9755 slightly outperforms with a score of 2,606.54 samples per minute in the "Monster" scene, closely followed by the 9965 with 2,558.43. Both surpass Genoa and Bergamo, with respective scores of 1,700.65 and 2,038.71, highlighting the advantage that Turin-based processors bring to 3D rendering. These scores highlight the suitability of the 9755 and 9965 for high-resolution 3D modeling tasks and content creation workflows involving complex scene rendering.
| Blender 4.0 CPU Samples per Minute (Higher is Better) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| Monster | 2,558.43 | 2,606.54 | 1,196.15 | 1,700.65 | 2,038.71 | 
| Junkshop | 1,866.65 | 1,843.48 | 802.00 | 1,101.84 | 1,382.58 | 
| Classroom | 1,270.17 | 1,251.54 | 637.13 | 869.48 | 1,045.96 | 
UL's Procyon AI inference test suite evaluates the performance of different AI inference engines using state-of-the-art neural networks. These tests were run on the CPU only. Each value represents an average inference time; the lower the value, the better the performance. The last line indicates an overall score; the higher the value, the better the performance.
Here, the EPYC 9575F with SMT disabled delivers solid inference times across various neural networks, with exceptional performance in MobileNet V3 (7.02 ms) and ResNet 50 (10.45 ms). These scores surpass both Genoa (3.63 ms for MobileNet and 6.34 ms for ResNet) and Bergamo (4.16 ms and 8.22 ms, respectively) for AI workloads, suggesting that the SMT-off configuration of the 9575F prioritizes latency control, making it highly effective for real-time AI tasks where low inference times are essential.
The EPYC 9965 and 9755, while slower in single inference tasks than Genoa, maintain solid performance across the board, with the 9755's 20.07 ms on MobileNet and 20.11 ms on ResNet 50 indicating stable and reliable performance suited to moderate workloads.
While Genoa and Bergamo offer consistently lower latency, the 9575F, especially with SMT disabled, offers excellent latency control.
| UL Procyon Average | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) |  | 
| Inference Time (Lower is Better) |  |  |  |  |  |  | 
| Mobile Net V3 | 32.86 ms | 20.07 ms | 7.02 ms | 3.63 ms | 4.16 ms |  | 
| ResNet 50 | 38.63 ms | 20.11 ms | 10.45 ms | 6.34 ms | 8.22 ms |  | 
| Creation V4 | 116.38 ms | 66.22 ms | 33.69 ms | 25.99 ms | 30.68 ms |  | 
