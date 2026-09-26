---
id: collect-240926-storagereview/storagereview/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f-4
title: "fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "compute", "decode", "gpu", "intel", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f.md
source_anchor: ""
source_lines: [106, 156]
sha256: 4faf15a4cbbd63e40111f84532d85bb4ffb4995af14a2ab5bfcab1d7a7032037
---

# fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f

The table below shows that the Supermicro and Lenovo systems recorded good results. The Lenovo ThinkSystem SR630 V4 outperformed in multi-core and single-core performance, with 99,266 and 894 points respectively. The Supermicro Hyper 1U 112H-TN recorded 92,516 and 888 points. The additional processor had some benefits in this benchmark, but the numbers did not double, going from one to two processors. The Ice Lake processors recorded 74,020, with limited scaling compared to Lenovo's dual Xeon 6780E.
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| Multi-Core CPU | 92,516 | 99,266 pts | 74,020 pts | 
| Single-Core CPU | 888 pts | 894 pts | 1,059 pts | 
| MP Ratio | 104.20 x | 111.00 x | 69.87 x | 
Cinebench 2024
Cinebench 2024 extends the benchmarking capabilities of R23 by adding a GPU performance evaluation. It continues to test CPU performance but also includes tests that measure the GPU's ability to handle rendering tasks.
The results of the 2024 version of Cinebench told a similar story. Here, the Lenovo ThinkSystem SR630 V4 with its two 6780E processors reflected the advantage over the single-socket Supermicro Hyper 1U 112H-TN in multi-core CPU performance with a score of 2,884 points. The Supermicro reported 2,565 points. The Intel Ice Lake 8380 processors demonstrated their power with a multi-core score of 4,131. For single-core tests, the Intel Ice Lake 8380 scored 61 points.
| Cinebench 2024 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| Multi-Core CPU | 2,565 pts | 2,884 pts | 4,131 pts | 
| Single-Core CPU | 53 pts | 53 pts | 61 pts | 
| MP Ratio | 48.38 x | 54.43 x | 68.22 x | 
y-cruncher
y-cruncher is a popular benchmarking and stress testing application launched in 2009. This test is multithreaded and scalable, calculating Pi and other constants to trillions of digits. Faster is better in this test. This software has been fantastic for testing high-core-count platforms and showing the computational advantages between single- or dual-socket platforms.
In the y-cruncher performance tests, the dual-socket ThinkSystem SR630 V4 took 5.997 seconds to calculate Pi to 1 billion digits. The Intel Xeon 6780E processor took 8.757 seconds. To calculate 50 billion digits, a single processor needed 674.299 seconds, compared to 476.826 seconds for dual processors. While not all workloads respond well to the high core count of the new e-core processors, y-cruncher had no trouble exploiting them. The older Intel Ice Lake Server completed the Pi calculation to 1 billion digits in 7.074 seconds in the first 1 billion digit test, while reaching 617.828 seconds for 50 billion digits.
| y-cruncher (0.8.5.9) (Lower is Better) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| 1 billion | 8.757 seconds | 5.997 seconds | 7.074 seconds | 
| 2.5 billion | 24.928 seconds | 17.573 seconds | 19.203 seconds | 
| 5 billion | 53.489 seconds | 37.793 seconds | 42.300 seconds | 
| 10 billion | 113.727 seconds | 81.046 seconds | 93.886 seconds | 
| 25 billion | 308.218 seconds | 220.025 seconds | 272.679 seconds | 
| 50 billion | 674.299 seconds | 476.826 seconds | 617.828 seconds | 
Blackmagic RAW Speed Test
The Blackmagic RAW Speed Test is a benchmarking tool designed to measure a system's ability to handle video playback and editing using the Blackmagic RAW codec. It evaluates a system's ability to decode and play high-resolution video files, providing frame rates for CPU-based and GPU-based processing.
The Lenovo ThinkSystem SR630 V4 achieved slightly higher results than the Supermicro and Ice Lake systems, with a score of 120 FPS in 8K CPU, making it an excellent choice for video playback and editing tasks. The Supermicro Hyper 1U 112H-TN achieved 116 FPS in 8K CPU. Although the dual-socket Lenovo performed better, it wasn't by much, given the dual-socket configuration with 116 FPS (8K CPU) and 0 FPS (8K GPU). The Intel Ice Lake server was on par with the Lenovo and Supermicro with 116 FPS in the 8K CPU benchmark. However, it did not record a GPU score due to the absence of a dedicated GPU.
7-Zip
The built-in memory benchmark of the popular 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, indicating how well the system can handle data-intensive operations.
Regarding compression tasks, the Supermicro system achieved slightly higher results than the SR630 in terms of CPU usage and resulting ratings, with a total compression rating of 245.823 GIPS and Lenovo's at 224.313 GIPS. This suggests a slight advantage when handling heavily threaded compression workloads. Lenovo's decompression tasks showed a higher resulting rating of 288.457 GIPS, while Supermicro's rating indicated 269.373 GIPS. This translates to better performance for workloads that require reading and extracting data. The Intel Ice Lake server demonstrated balanced performance, with a compression rating of 235.437 GIPS and a decompression rating of 253.692 GIPS.
The overall performance of the two systems is nearly identical, with total ratings of 257.598 GIPS for the single-socket Supermicro and 256.385 GIPS for the dual-socket Lenovo. The older Xeon Ice Lake reached 244.565 GIPS. All three systems are very high-performing.
| 7-Zip Compression | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| Compression |  |  |  | 
| Current CPU Usage | 5287 % | 5064 % | 5835 % | 
| Current Rating/Usage | 4.647 GIPS | 4.341 GIPS | 4.030 GIPS | 
| Current | 245.699 GIPS | 219.840 GIPS | 235.143 GIPS | 
| Resulting CPU Usage | 5296 % | 5156 % | 5839 | 
| Resulting Rating/Usage | 4.642 GIPS | 4.350 GIPS | 4.032 GIPS | 
| Resulting Rating | 245.823 GIPS | 224.313 GIPS | 235.437 GIPS | 
| Decompression |  |  |  | 
| Current CPU Usage | 6236 % | 6184 % | 6230 % | 
| Current Rating/Usage | 4.261 GIPS | 4.688 GIPS | 4.050 GIPS | 
| Current | 265.709 GIPS | 289.879 GIPS | 252.326 GIPS | 
| Resulting CPU Usage | 6236 % | 6205 % | 6245 % | 
| Resulting Rating/Usage | 4.341 GIPS | 4.649 GIPS | 4.062 GIPS | 
| Resulting Rating | 269.373 GIPS | 288.457 GIPS | 253.692 GIPS | 
| Total Rating |  |  |  | 
| Total CPU Usage | 5751 % | 5681 % | 6042 % | 
| Total Rating/Usage | 4.491 GIPS | 4.500 GIPS | 4.047 GIPS | 
| Total Rating | 257.598 GIPS | 256.385 GIPS | 244.565 GIPS | 
Conclusion
The Lenovo ThinkSystem SR630 V4 is a versatile 1U rack server that, while not representing a significant upgrade over its predecessor, is nonetheless a reliable advance in the evolution of Lenovo's mainstream enterprise systems. With support for Intel Xeon 6700E series processors, it doubles core density compared to its predecessor, offering better scalability for demanding workloads. Improved DDR5 memory speeds up to 6400 MHz and planned support for emerging technologies such as Compute Express Link (CXL) and MCRDIMM show that Lenovo intends to keep this server line equipped for future demands.
The SR630 V4's evolution toward NVMe storage and flexible drive configurations also means it prioritizes performance and scalability in data-intensive workloads. Additionally, the two OCP 3.0 slots supporting PCIe 5.0 enable advanced networking options, including hot-swap components and improved cooling systems, simplifying maintenance and operational efficiency.
