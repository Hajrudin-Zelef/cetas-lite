---
id: collect-240926-storagereview/storagereview/fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626
title: "fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Meta", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "cost", "deepseek", "energy", "gpu", "inference", "latency", "llama", "memory", "nvidia"]
source: docs/RAG/clean_en/storagereview/fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626.md
source_anchor: ""
source_lines: [1, 79]
sha256: 435907d0da26a05a84ba950d83c55f61fec01dbf847a2a995e8f5674153edc0d
---

# fr-review-western-digital-ultrastar-dc-hc590-review-dfc98626

<!-- source: https://www.storagereview.com/fr/review/western-digital-ultrastar-dc-hc590-review -->

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
In the 3DMark Storage Benchmark, the Seagate x24 24TB achieved the best score among the hard drives tested, with an overall score of 234. The WD HC590 26TB comes in second place with a score of 168, placing it about 12% ahead of the WD Red Pro and 12% ahead of the WD Gold, whose score is below 150. Compared to the Seagate x24, the HC590 is about 28% behind, but remains well positioned in the upper rankings.
| 3DMark Storage Benchmark (higher is better) | Overall Score | 
| Seagate x24 24TB | 234 | 
| WD Ultrastar DC HC590 26TB | 168 | 
| WD Red Pro 22TB | 156 | 
| WD Gold 24TB | 150 | 
BlackMagic Disk Speed Test
The Blackmagic Disk Speed Test is designed to evaluate a drive's ability to handle high-resolution video workflows, often used in multimedia production. The Ultrastar HC590 displayed read and write speeds of 267.0 MB/s and 264.5 MB/s respectively. These results indicate that the drive can support playback and editing of compressed high-resolution formats, but does not meet the requirements of uncompressed RAW workflows. Again, this performance is in line with expectations for a 7200 RPM enterprise hard drive and is sufficient for archiving, backup, and offline editing.
In the Blackmagic Disk Speed Test, the WD HC590 26TB displayed balanced performance with read speeds of 267.0 MB/s and write speeds of 264.5 MB/s. While its read speed was slightly lower by about 2% compared to the WD Gold and the Seagate x24, its write speed stood out as the highest in the group. It surpassed the WD Gold by 24% and the Seagate by 61% in write.
| BlackMagic Disk Speed (MB/s, higher is better) | Read MB/s | Write MB/s | 
| WD Gold 24TB | 272.8 | 213.0 | 
| Seagate x24 24TB | 271.0 | 164.4 | 
| WD Ultrastar DC HC590 26TB | 267.0 | 264.5 | 
| WD Red Pro 22TB | 260.9 | 258.3 | 
PCMark 10 Storage
PCMark 10 storage benchmarks evaluate real-world storage performance using application traces. They test system and data drives, measuring bandwidth, access times, and consistency under load. These benchmarks provide practical insights beyond synthetic tests, allowing users to effectively compare modern storage solutions.
In the PCMark 10 Data Drive Benchmark, the WD Ultrastar DC HC590 26TB achieved the best performance in the group with an overall score of 853. This is about 27% more than the Seagate x24 and more than double the scores of the WD Gold and WD Red Pro, which reached 397 and 380 respectively.
| PCMark 10 Data Drive (higher is better) | Overall Score | 
| WD Ultrastar DC HC590 26TB | 853 | 
| Seagate x24 24TB | 671 | 
| WD Gold 24TB | 397 | 
| WD Red Pro 22TB | 380 | 
Conclusion
The Western Digital Ultrastar DC HC590 26TB stands out as a serious contender in the high-capacity enterprise hard drive market, offering a judicious balance of density, reliability, and real-world performance. Its 11-platter CMR design delivers a maximum capacity of 26TB in a standard 3.5-inch form factor without requiring changes to existing infrastructure, a key advantage for enterprise and hyperscale environments seeking to scale efficiently.
In our tests, the HC590 consistently displayed excellent performance. In synthetic benchmarks like FIO, it kept pace with its competitors in sequential write and demonstrated excellent random write capability. It also held up well in application tests, displaying competitive LLM load times, write speeds superior to the Blackmagic Disk Speed Test, and first place in the PCMark 10 Data Drive benchmark. The 3DMark Storage results place it firmly in the leading pack of the hard drives tested, confirming its ability to handle storage-hungry workloads.
