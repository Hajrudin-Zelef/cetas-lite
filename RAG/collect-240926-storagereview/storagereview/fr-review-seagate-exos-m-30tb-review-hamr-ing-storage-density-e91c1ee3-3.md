---
id: collect-240926-storagereview/storagereview/fr-review-seagate-exos-m-30tb-review-hamr-ing-storage-density-e91c1ee3-3
title: "fr-review-seagate-exos-m-30tb-review-hamr-ing-storage-density-e91c1ee3"
domain: storagereview
role: reference
task: reference
actors: ["DeepSeek", "Meta"]
dates: []
keywords: ["benchmark", "benchmarks", "compute", "deepseek", "inference", "latency", "llama", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-seagate-exos-m-30tb-review-hamr-ing-storage-density-e91c1ee3.md
source_anchor: ""
source_lines: [65, 120]
sha256: 66da3a963e07af3f268aa5acca1eb3f6a1018f328480bf02c652824d96ee5c9d
---

# fr-review-seagate-exos-m-30tb-review-hamr-ing-storage-density-e91c1ee3

4K Random Workloads (16 threads, 32 queues)
In small-block random I/O, performance often reveals a drive's ability to handle metadata-intensive workloads or high-concurrency applications. The Exos M 30TB reached 205 IOPS in random read (155.58 ms) and 341 IOPS in random write (93.79 ms).
Read and write performance in IOPS was competitive across all drives, with the WD Gold and WD Red Pro ranking slightly ahead at 214 IOPS. On the other hand, the Exos M's write performance was more significant, surpassing that of the IronWolf Pro (301 IOPS) and clearly that of the Ultrastar HC590 (663 IOPS) in terms of latency, without however matching that of the x24 24TB, which reached 749 IOPS with a much lower latency of 42.70 ms.
| FIO Test (higher MB/s/IOPS throughput is better) | Sequential Read 128K (1T/64Q) | Sequential Write 128K (1T/64Q) | Random 4K Read (16T/32Q) | Random 4K Write (16T/32Q) | 
| Seagate Exos 30TB | 292 MB/s (28.72 ms) | 289 MB/s (29.04 ms) | 205 IOPS (155.58 ms) | 341 IOPS (93.79 ms) | 
| Seagate Iron Wolf Pro 30TB | 287 MB/s (29.23 ms) | 267 MB/s (31.39 ms) | 205 IOPS (155.74 ms) | 301 IOPS (105.95 ms) | 
| Seagate x24 24TB | 285 MB/s (29.42 ms) | 285 MB/s (29.42 ms) | 210 IOPS (152.03 ms) | 749 IOPS (42.70 ms) | 
| WD Gold 24TB | 283 MB/s (29.66 ms) | 286 MB/s (29.36 ms) | 214 IOPS (148.98 ms) | 651 IOPS (49.11 ms) | 
| WD Red Pro 22TB | 271 MB/s (31.00 ms) | 276 MB/s (30.37 ms) | 214 IOPS (149.17 ms) | 421 IOPS (75.92 ms) | 
| WD Ultrastar DC HC590 26TB | 268 MB/s (31.28 ms) | 280 MB/s (30.00 ms) | 198 IOPS (161.18 ms) | 663 IOPS (48.23 ms) | 
Average LLM Load Times
The average LLM load time test evaluated the load times of three different LLMs: DeepSeek R1 7B, Meta Llama 3.2 11B, and DeepSeek R1 32B. Each model was tested 10 times and the average load time was calculated. This test measures the drive's ability to quickly load large language models (LLMs) into memory. LLM load times are essential for AI-related tasks, including real-time inference and processing of large datasets. Faster loading allows the model to quickly process data, improving AI responsiveness and reducing wait times.
The Seagate Exos M 30TB recorded the fastest load times on two of the three AI models tested, outperforming all other drives on DeepSeek R1 7B and Meta Llama 3.2 11B. It edged out the WD Gold 24TB by a small but consistent margin, demonstrating its ability to maintain high throughput on multi-gigabyte workloads.
Although the Exos M is not at the top of the ranking for the 32B models, it remains close to the best drives like the WD Gold and Ultrastar HC590. This consistency across different model sizes demonstrates that the Exos M maintains low latency and excellent read performance, even under demanding, high-volume conditions.
| Average LLM Load Time (lower is better) | DeepSeek R1 7 billion | Meta Llama 3.2 11B Vision | DeepSeek R1 32 billion | 
| Seagate Exos 30TB | 46.4424s | 68.7064s | 72.7249s | 
| WD Gold 24TB | 46.7133s | 68.8183s | 68.9720s | 
| WD Ultrastar DC HC590 26TB | 47.9877s | 71.0063s | 69.7892s | 
| Seagate Iron Wolf Pro 30TB | 48.4175s | 69.9071s | 72.3803s | 
| Seagate x24 24TB | 48.6615s | 71.4855s | 73.8097s | 
| WD Red Pro 22TB | 49.0575s | 71.4783s | 71.1382s | 
3DMark Storage
The 3DMark Storage benchmark tests your SSD's gaming performance by measuring tasks such as loading, saving progress, installing files, and recording gameplay. It evaluates your storage's ability to handle real-world gaming activities and supports the latest storage technologies for accurate performance analysis.
In the 3DMark Storage benchmark, the Seagate Exos M 30TB scored 223, ranking second overall, just one point behind the Seagate x24 24TB with a score of 234 and ahead of the IronWolf Pro 30TB with a score of 231. Additionally, it placed ahead of the WD Ultrastar DC HC590 (168) and WD Red Pro 22TB (156).
| 3DMark Storage Benchmark (higher is better) | Overall Score | 
| Seagate x24 24TB | 234 | 
| Seagate Exos 30TB | 223 | 
| Seagate Iron Wolf Pro 30TB | 231 | 
| WD Ultrastar DC HC590 26TB | 168 | 
| WD Red Pro 22TB | 156 | 
| WD Gold 24TB | 150 | 
BlackMagic Disk Speed Test
The BlackMagic Disk Speed Test evaluates a disk's read and write speeds and assesses its performance, particularly for video editing. It allows users to ensure their storage is fast enough for high-resolution content, such as 4K or 8K video.
In the disk speed test, the Seagate Exos M 30TB recorded 274.6 MB/s read and 275.2 MB/s write, ranking first in both categories. It outperformed all other drives, including the WD Gold 24TB and IronWolf Pro 30TB, which lagged behind in write performance.
| BlackMagic Disk Speed (MB/s, higher is better) | Read MB/s | Write MB/s | 
| Seagate Exos 30TB | 274.6 | 275.2 | 
| WD Gold 24TB | 272.8 | 213.0 | 
| Seagate x24 24TB | 271.0 | 164.4 | 
| Seagate Iron Wolf Pro 30TB | 267.6 | 272.7 | 
| WD Ultrastar DC HC590 26TB | 267.0 | 264.5 | 
| WD Red Pro 22TB | 260.9 | 258.3 | 
PCMark 10 Storage
PCMark 10 Storage benchmarks evaluate real-world storage performance using application traces. They test the system and data drives, measuring bandwidth, access times, and consistency under load. These benchmarks provide practical insights beyond synthetic tests, allowing users to effectively compare modern storage solutions.
In the PCMark 10 Storage ranking, the Seagate Exos M 30TB scored 769 points, ranking third overall. It closely follows the IronWolf Pro 30TB (771 points) and the WD Ultrastar DC HC590 (853 points), while surpassing the Seagate x24 24TB (671 points), the WD Gold 24TB (397 points), and the WD Red Pro 22TB (380 points).
| PCMark 10 Data Drive (higher is better) | Overall Score | 
| WD Ultrastar DC HC590 26TB | 853 | 
| Seagate Iron Wolf Pro 30TB | 771 | 
| Seagate Exos 30TB | 769 | 
| Seagate x24 24TB | 671 | 
| WD Gold 24TB | 397 | 
| WD Red Pro 22TB | 380 | 
Conclusion
The Seagate Exos M 30TB represents a significant advancement in high-capacity enterprise storage. It launches alongside the IronWolf Pro 30TB, within Seagate's new-generation platform. Featuring 3TB per platter capacity, sustained transfer speeds of 275 MB/s, an MTBF of 2.5 million hours, and a workload of 550TB/year, the Exos M is designed to deliver performance, durability, and density at scale.
In our comparative tests, the Exos M 30TB consistently displayed first-rate sequential performance, stood out in most read and write tests, and showed competitive latency, both for AI inference and general workloads. It performed particularly well in LLM model loading, FIO, and Blackmagic disk speed tests, where bandwidth and responsiveness are essential.
The Exos M 30TB offers performance comparable to Seagate's previous-generation Exos X24 24TB in most cases, while providing 25% more usable capacity for an identical footprint. It also complies with Open Compute Project (OCP) specifications, in accordance with standards adopted by major cloud service providers to validate its performance at scale.
Ideal for hyperscale deployments, high-density RAID arrays, distributed file systems like Hadoop or Ceph, and enterprise backup environments, the Exos M 30TB offers a perfect balance between efficiency, capacity, and lifespan. Combined with Seagate Secure protection and a five-year warranty, it constitutes a major asset for businesses building future-proof infrastructure.
