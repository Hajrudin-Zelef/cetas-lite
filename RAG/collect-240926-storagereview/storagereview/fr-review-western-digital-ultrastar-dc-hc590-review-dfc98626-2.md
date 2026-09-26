---
id: collect-240926-storagereview/storagereview/fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626-2
title: "fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Meta", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "cost", "deepseek", "energy", "gpu", "inference", "latency", "llama", "memory", "nvidia"]
source: docs/RAG/clean_en/storagereview/fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626.md
source_anchor: ""
source_lines: [3, 54]
sha256: 510e386f5de9fbd387dc67d22d50dba19a110ca2aa877d655ddf9451c41a8ed2
---

# fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626

The Western Digital Ultrastar DC HC590 is WD's latest high-capacity enterprise hard drive. It offers up to 26TB of storage through conventional magnetic recording (CMR) in a standard 3.5-inch form factor. It is the first CMR hard drive to use an 11-platter design, delivering higher raw capacity without requiring changes to existing infrastructure. This allows enterprise and hyperscale environments to scale storage density within the same physical and thermal footprint, while preserving compatibility with current rack designs and enclosures. Ultimately, the HC590 is designed for storage-hungry, read-focused workloads where reliability and cost-effectiveness take priority over peak performance.
WD Ultrastar DC HC590 Features
The HC590 features reliability characteristics familiar to this class of drives, designed to maintain consistent performance in dense, vibration-prone environments. Notably, it combines energy-assisted PMR (ePMR) technology with its OptiNAND architecture, which integrates embedded flash memory (iNAND) to offload internal tasks such as metadata management. Together, these technologies help increase areal density and reduce strain on the magnetic media, improving efficiency and throughput during large-scale sequential or metadata-heavy operations.
To ensure stable performance in dense server racks, the drive is also equipped with RVS (Rotational Vibration Safeguard) technology, which uses two sensors to detect and compensate for environmental vibrations. It also incorporates DFH (Dynamic Fly Height) technology, which adjusts the read/write head on the fly for better contact precision. The helium-sealed enclosure reduces internal resistance and power consumption compared to air-filled solutions.
ArmorCache is also available, allowing users to select write cache disabled (WCD) mode to improve random write performance or write cache enabled (WCE) mode to protect cached data in the event of a power failure. This is particularly useful in systems where caching policy is governed by data integrity or throughput requirements.
On the reliability front, the drive boasts a MTBF of 2.5 million hours and comes with a 5-year limited warranty, both typical characteristics of enterprise hard drives. The HC590 is available in SAS and SATA versions, with capacities of 24TB and 26TB. For this test, we tested the 26TB SATA model.
WD Ultrastar DC HC590 26TB Hard Drive Specifications
| Specifications | DETAILS | 
| Model | WD Ultrastar DC HC590 | 
| Capacities | 26TB | 
| Form Factor | 3.5 inch | 
| Interface | SAS 12 Gb/s or SATA 6 Gb/s | 
| Recording Technology | Conventional Magnetic Recording (CMR) | 
| Number of Platters | 11 | 
| Helium-Sealed | Yes | 
| Advanced Technologies | ePMR, OptiNAND, Dynamic Fly Height, Rotational Vibration Safeguard (RVS) | 
| Cache Options | ArmorCache: WCE (data protection) / WCD (write performance) | 
| Mean Time Between Failures (MTBF) | 2.5 million hours | 
| Security Features | Secure Erase (SE) | 
| Dimensions (L x W x H) | 5.776 "x 4.000" x 1.028 " | 
| Weight | 1.47 lb (approximately 667 g) | 
| Warranty | 5-Year Limited Warranty | 
| Model Number | WUH722626AL5204 | 
WD Ultrastar DC HC590 Performance
Before diving into the benchmarks, here is a list of comparable-capacity hard drives used to compare the performance of the WD Ultrastar DC HC590 26TB drive.
Here is the high-performance test bench we used for storage benchmarking:
- CPU: AMD Ryzen 7 9800X3D
- Motherboard: Asus ROG Crosshair X870E Hero
- RAM: G.SKILL Trident Z5 Royal Series DDR5-6000 (2 x 16GB)
- GPU: NVIDIA GeForce RTX 4090
- Operating System: Windows 11 Pro, Ubuntu 24.10 Desktop
Peak Synthetic Performance
The FIO test is a flexible and powerful benchmarking tool for measuring the performance of storage devices, including SSDs and hard drives. It evaluates metrics such as bandwidth, IOPS (input/output operations per second), and latency under various workloads, such as sequential and random read/write operations. This test allows for assessing the peak performance of storage systems, making it useful for comparing different devices or configurations. We measured peak burst performance for this test, limiting the workload to 10GB across all hard drives.
In our FIO tests, the WD DC HC590 26TB recorded sequential performance slightly lower than its competitors, but remained within a close range. In sequential read at 128K, it reached 268 MB/s, about 6% less than the Seagate Exos X24 and about 5% less than the WD Gold. In sequential write, it reached 280 MB/s, only 2% less than the WD Gold, and close to that of the Seagate.
In 4K random read, the HC590 reached 198 IOPS, about 7% less than the Seagate and the WD Gold. Its 4K random write performance was superior, with 663 IOPS, only 11% less than the Seagate, while surpassing the WD Red Pro by more than 50% and the WD Gold by nearly 2%.
Overall, the HC590 maintains solid sequential throughput and demonstrates strength in random write performance with only modest differences compared to the other high-capacity drives tested.
| FIO Test (higher MB/s/IOPS is better) | Sequential Read 128K (1T/64Q) | Sequential Write 128K (1T/64Q) | Random 4K Read (16T/32Q) | Random 4K Write (16T/32Q) | 
| Seagate x24 24TB | 285 MB/s (29.42 ms) | 285 MB/s (29.42 ms) | 210 IOPS (152.03 ms) | 749 IOPS (42.70 ms) | 
| WD Gold 24TB | 283 MB/s (29.66 ms) | 286 MB/s (29.36 ms) | 214 IOPS (148.98 ms) | 651 IOPS (49.11 ms) | 
| WD Red Pro 22TB | 271 MB/s (31.00 ms) | 276 MB/s (30.37 ms) | 214 IOPS (149.17 ms) | 421 IOPS (75.92 ms) | 
| WD Ultrastar DC HC590 26TB | 268 MB/s (31.28 ms) | 280 MB/s (30.00 ms) | 198 IOPS (161.18 ms) | 663 IOPS (48.23 ms) | 
Average LLM Load Time
The average LLM load time test evaluated the load times of three different LLMs: DeepSeek R1 7B, Meta Llama 3.2 11B, and DeepSeek R1 32B. Each model was tested 10 times and the average load time was calculated. This test measures the drive's ability to quickly load large language models (LLMs) into memory. LLM load times are critical for AI-related tasks, including real-time inference and processing large datasets. Faster loading allows the model to process data quickly, improving AI responsiveness and reducing wait times.
In the average LLM load time test, the WD Gold 24TB achieved the best results among the three models. DeepSeek R1 7B achieved a load time of 46.71 seconds, about 2.7% faster than the WD Ultrastar DC HC590 (47.99 seconds) and about 4% faster than the Seagate x24. On Meta Llama 3.2 11B Vision, the Gold achieved a load time of 68.82 seconds, versus 71.01 seconds for the HC590, a difference of about 3.2%.
The gap remained consistent with DeepSeek R1 32B, where the WD Gold again dominated with 68.97 seconds, followed by the HC590 with 69.79 seconds and the Seagate x24 further back with 73.81 seconds. Overall, the HC590 was only 1 to 3% slower than the WD Gold, while surpassing the Seagate x24 and WD Red Pro in all tests.
| Average LLM Load Time (lower is better) | DeepSeek R1 7B | Meta Llama 3.2 11B Vision | DeepSeek R1 32B | 
| WD Gold 24TB | 46.7133s | 68.8183s | 68.9720s | 
| WD Ultrastar DC HC590 26TB | 47.9877s | 71.0063s | 69.7892s | 
| Seagate x24 24TB | 48.6615s | 71.4855s | 73.8097s | 
| WD Red Pro 22TB | 49.0575s | 71.4783s | 71.1382s | 
3DMark Storage Benchmark
The 3DMark Storage Benchmark tests the gaming performance of your SSD or hard drive by measuring tasks such as loading, saving progress, installing files, and recording gameplay. It evaluates your storage's ability to handle real-world gaming activities and supports the latest storage technologies for accurate performance analysis.
