---
id: collect-240926-storagereview/storagereview/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903-3
title: "fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Samsung"]
dates: []
keywords: ["acquisition", "amd", "benchmark", "compute", "cost", "energy", "latency", "licenses", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903.md
source_anchor: ""
source_lines: [51, 92]
sha256: e8ac5dc636a2daf837bebfb0c12ebf82292d04652fdc4870e5ec9ef9d03a5f65
---

# fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903

The 32-core 9335 model is the most powerful in our range and offers the greatest flexibility for compute-intensive workloads. The 24-core 9255 model comes closest to the best price-performance ratio observed in our tests, particularly for databases, where marginal gains from core count become negligible beyond 24 cores. The 16-core 9135 and 8-core 9015 models offer the best value. Since most software used by SMBs is licensed per core, including Windows Server, many relational databases, and some hypervisor and backup platforms, the number of cores chosen at purchase represents a recurring cost for the entire life of the deployment.
Choosing an 8 or 16-core processor suited to the workload, rather than oversizing idle cores, reduces acquisition costs and licensing expenses. The 16-core 9135 model is particularly relevant because it matches the minimum number of cores required by Windows Server licenses. It is therefore an ideal starting point for Windows environments looking to properly size their processor according to licensing constraints, without compromising performance. The 8-core 9015 model is the most economical and energy-efficient option, ideal for storage-focused workloads or roles where the processor is not a limiting factor. Each processor runs at the same memory speed and supports the same number of PCIe Gen5 lanes; higher-tier component configuration choices are therefore not limited by the model choice.
Performance testing
| Test configurations | Dell PowerEdge R4715 | Dell PowerEdge R5715 | 
|---|---|---|
| Processors tested | AMD EPYC 9335, 9255, 9135, 9015 | AMD EPYC 9015 | 
| Memory | 384GB DDR5 | 384GB DDR5 | 
| Boot storage | BOSS RAID1 | BOSS RAID1 | 
| Front storage configuration | 8x Samsung PM9D3a RI U.2 Gen5 NVMe SSDs (1.92 TB) RAID 10 x 6 | 12x 20 TB hard drives in RAID 6 | 
Database performance: HammerDB MariaDB TPC-C
The primary workload for this evaluation is HammerDB running TPC-C against MariaDB 12.3.1. TPC-C is a recognized OLTP benchmark that produces measurable, comparable results regardless of processor and storage configuration, and represents the type of transactional database workload at the heart of most SMB applications. We tested two distinct profiles: a CPU-intensive profile, which heavily stresses transactional processing, and an I/O-intensive profile, which places a greater load on the storage subsystem. Both profiles were run on all four processor options of the R4715 flash configuration to obtain a clear CPU scaling curve, then on the R5715 hard drive configuration to observe variations based on storage type.
Dell PowerEdge R4715
The HammerDB results show a clear improvement in performance as core count increases on the R4715 platform. With the 8-core EPYC 9015 processor, the system reached 480,818 NOPM in CPU-intensive mode and 296,105 NOPM in I/O-intensive mode before leveling off despite increasing virtual user counts. Moving to the 16-core EPYC 9135 processor delivered a considerable throughput gain, bringing performance to 737,445 NOPM in CPU-intensive mode and 493,093 NOPM in I/O-intensive mode, while allowing the system to support a higher number of virtual users before saturation.
Moving to the 24-core EPYC 9255 processor allowed the platform to exceed one million NOPM in CPU-intensive mode, peaking at 1,017,429 NOPM, while in I/O-intensive mode it peaked at 740,574 NOPM. At this point, additional cores continued to translate directly into usable transactional throughput, while the NVMe storage subsystem kept pace with the growing database load.
At the top of the range, the 32-core EPYC 9335 processor delivered the best results on both profiles, reaching 1,133,714 NOPM for the CPU-intensive workload and 910,321 NOPM for the I/O-intensive one. The performance curve remained relatively stable even with a high number of virtual users, indicating that the R4715 flash configuration effectively leveraged the more powerful CPU configurations without storage bottlenecks prematurely limiting performance.
Dell PowerEdge R5715
We then tested the Dell PowerEdge R5715, configured with 12x 20 TB hard drives in RAID 6, paired with the 8-core AMD EPYC 9015 processor. In the CPU-intensive profile, the platform peaked at 484,715 NOPM with 16 virtual users, with throughput increasing steadily as more users were added before leveling off near saturation.
The I/O-intensive activity profile peaked at 308,012 NOPM with 24 virtual users, demonstrating the excellent transactional performance of the high-capacity hard drive array under moderate concurrency. As the workload increased, the scaling curve flattened, with the hard drive subsystem approaching its practical performance limits under sustained concurrent database activity.
Windows Server shared storage
The second practical use case targets another type of SMB workload: Windows shared storage. For organizations using file shares, departmental applications, or general-purpose Windows servers, these platforms meet current SMB performance requirements. We ran FIO on Windows Server to characterize the sequential and random performance of two storage configurations: a RAID 6 hard drive array (R5715) as a baseline, and a JBOD SSD array with 8 drives (R4715) representing a high-performance storage configuration. The comparison highlights the performance gap between the two storage tiers in this category of platforms.
The gap between the two storage configurations is clearly visible in the FIO results. While the R5715's RAID 6 hard drive array offered respectable sequential throughput for a high-capacity hard drive platform, the SSD-equipped R4715 showed significantly higher performance, particularly for random workloads, where SMB environments are especially sensitive to storage latency.
Sequential performance on the hard drive array reached up to 3.7 GB/s write and 2.2 GB/s read in the 4-core tests, which is more than sufficient for classic file sharing, backup, and mass storage tasks. In contrast, the SSD configuration achieved sequential throughput of several tens of gigabytes per second, exceeding 56 GB/s read and 26 GB/s write, while maintaining considerably reduced latency.
The gap widened further in the 4K random write tests. The hard drive array capped at less than 1,300 IOPS in random write, with latency exceeding 100 ms, while the SSD configuration reached over 4 million IOPS with sub-millisecond latency. In concrete terms, this translates directly into better application responsiveness, increased performance for multi-user file sharing, optimized virtual machine storage behavior, and the ability to handle simultaneous SMB workloads without storage becoming a bottleneck.
| FIO workload | R5715 RAID6 HDD 1T | R5715 RAID6 HDD 4T | R4715 8x SSD 1T | R4715 8x SSD 4T | 
|---|---|---|---|---|
| Sequential read (128 KB) |  |  |  |  | 
| Bandwidth | 1,475.89 MB/s | 2,198.89 MB/s | 56,861.09 MB/s | 56,866.76 MB/s | 
| IOPS | 11,807 | 17,589 | 454,885 | 454,918 | 
| Latency | 2.70ms | 7.28ms | 0.56ms | 2.25ms | 
| Sequential write (128 KB) |  |  |  |  | 
| Bandwidth | 2,665.31 MB/s | 3,726.63 MB/s | 26,739.52 MB/s | 26,753.48 MB/s | 
| IOPS | 21,322 | 29,811 | 213,912 | 214,011 | 
| Latency | 1.49ms | 4.39ms | 1.20ms | 4.78ms | 
| Random read (4K) |  |  |  |  | 
| Bandwidth | 1.00 MB/s | 3.60 MB/s | 7,268.78 MB/s | 16,143.19 MB/s | 
| IOPS | 256 | 919 | 1,860,803 | 4,132,645 | 
| Latency | 125.11ms | 139.00ms | 0.13ms | 0.17ms | 
| Random write (4K) |  |  |  |  | 
| Bandwidth | 4.88 MB/s | 4.67 MB/s | 7,555.53 MB/s | 16,010.24 MB/s | 
| IOPS | 1,248 | 1,195 | 1,934,214 | 4,098,613 | 
| Latency | 25.63ms | 106.98ms | 0.07ms | 0.13ms | 
Proxmox backup server
