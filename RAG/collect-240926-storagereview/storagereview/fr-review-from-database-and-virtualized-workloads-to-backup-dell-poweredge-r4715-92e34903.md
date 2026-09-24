---
id: collect-240926-storagereview/storagereview/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903
title: "fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Microsoft", "Samsung"]
dates: ["2026-03"]
keywords: ["accelerator", "acquisition", "amd", "benchmark", "compute", "cost", "energy", "gpus", "latency", "licenses", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903.md
source_anchor: ""
source_lines: [1, 99]
sha256: 86877047eb1ef57d83eb5caf70610a970eafa3cdfd4e267434626737e22e75aa
---

# fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903

<!-- source: https://www.storagereview.com/fr/review/from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-and-r5715-for-smb-realities -->

Although the Dell PowerEdge R4715 and R5715 servers are two distinct products, they can be considered a configurable solution. This solution includes two chassis, four AMD EPYC 9005 series processor options, a wide choice of storage configurations, and the complete Dell management and support ecosystem. It is specially designed for SMBs that need to adapt their infrastructure investments to their actual workload needs, as well as for the partners who support them in this process.
Both platforms were launched in March 2026. We analyzed them individually in our tests of the R4715 and R5715 servers. This article is different. Instead of evaluating each server separately, we examine the performance of both platforms and all four processor options for the workloads actually used by SMBs, and we identify the areas where configuration choices have the most impact.
Hypervisor flexibility is a major asset that justifies the current interest in these platforms. The virtualization market is evolving and businesses of all sizes are reevaluating the foundations of their infrastructure. Some are keeping their existing architecture, others are migrating, and many are running two or more hypervisors in parallel for the near future. The R4715 and R5715 support all common options, including VMware ESXi, Microsoft Hyper-V, Proxmox VE, and the major KVM Linux distributions, with consistent management and provisioning regardless of the hypervisor chosen. For SMBs that cannot afford to standardize their infrastructure on a single platform, this flexibility is a definite advantage and explains why we conducted tests on multiple hypervisors for this article.
The advantage of the Dell ecosystem for SMBs and distributors
Discussions around server platforms often focus on silicon, which makes sense at the technical specifications level. But for SMBs and the resellers and system integrators who support them, the operational experience with the silicon is often decisive. Dell's PowerEdge ecosystem is mature and well understood, and its benefits are particularly important for organizations with small IT teams.
iDRAC10 and OpenManage Enterprise are the most visible components. The same management platform applies to the entire 17th generation PowerEdge range. Thus, an SMB that acquires an R4715 today can later add an R7725 or any other PowerEdge model without having to learn new tools. For resellers and system integrators managing many clients, this consistency is even more valuable. A technician who knows iDRAC understands every PowerEdge client's infrastructure. The platform supports remote console access, firmware management, hardware health monitoring, and full Redfish API access for automation. For clients without dedicated infrastructure staff, this capability often makes the difference between a simple call to a partner and an on-site intervention.
Beneath the management layer, Dell offers unmatched security and supply chain. Silicon-level root of trust, cryptographically signed firmware, secure component verification, and FIPS-certified TPM 2.0 are standard. ProSupport and ProDeploy services are available globally, a major asset for distributed SMBs and partners operating in multiple regions. Dell's supply chain is one of the few in the industry to guarantee predictable delivery times at scale. For value-added resellers (VARs) looking to close sales despite stock shortage uncertainty, this is an undeniable competitive advantage.
For an SMB whose IT team has two or three people, or for a business partner managing dozens of clients with limited resources, the Dell ecosystem considerably reduces the operational surface area. The R4715 and R5715 servers benefit from all these advantages.
Overview of the R4715 and R5715 models
Here is a brief recap for readers who have not consulted our individual tests. The R4715 is a 1U single-processor server optimized for high compute density. It supports up to 24 DDR5 RDIMM modules, three PCIe Gen5 slots, and various storage options, including 2.5-inch and 3.5-inch SAS/SATA configurations as well as an 8-bay 2.5-inch NVMe U.2 configuration. It is the ideal choice when rack density and compute power per rack unit take priority over disk count.
The R5715 is a 2U single-processor server optimized for storage capacity and I/O expandability. It supports the same 24 DDR5 RDIMM modules and the four processor options. However, the R5715 adds a fourth PCIe Gen5 slot and offers up to 12 bays for 3.5-inch SAS/SATA drives or 16 bays for 2.5-inch SAS/SATA drives. The 3.5-inch configuration can reach 288 TB of raw capacity on a single node; this is the configuration we used to design our R5715 in this article.
Both platforms are air-cooled, shipped with an iDRAC10, and compatible with 800 W and 1,100 W power supplies, at Platinum or Titanium efficiency levels. These PowerEdge servers do not support GPUs, DPUs, or Fibre Channel, which aligns with Dell's positioning: platforms with dimensions suited to a specific use rather than platforms offering maximum flexibility.
It is important to understand the place of these two servers within Dell's AMD-based PowerEdge range. The R4715 and R5715 are the entry-level models optimized for price-performance ratio, designed specifically for the SMB workloads presented in this article. Customers needing accelerators, more cores, or greater expansion capacity can easily move up to the R6715 and R7715, which include support for GPUs and DPUs, processors offering well over 32 cores, and additional PCIe capacity for performance-demanding, accelerator-based workloads. This tiering is an asset for resellers: a VAR can install an R4715 or R5715 at a client site and upgrade their configuration to an R6715 or R7715 as their needs evolve, all within the same management platform, the same deployment process, and the same support model.
Platform specifications
| Specifications | Dell PowerEdge R4715 | Dell PowerEdge R5715 | 
|---|---|---|
| Processor |  |  | 
| Processor | One 5th Gen AMD EPYC 9005 series processor, up to 32 cores |  | 
| Form factor | 1U rack server | 2U rack server | 
| Memory |  |  | 
| DIMM slots | 24 DDR5 DIMM slots |  | 
| Maximum memory | 1.5 TB (up to 64 GB per DIMM) |  | 
| Memory speed | Up to 5200 MT/s |  | 
| Memory type | Registered DDR5 ECC RDIMM modules only |  | 
| Storage |  |  | 
| Internal controllers (RAID) | PERC H365i, H965i |  | 
| Internal boot | BOSS-N1 DC-MHS |  | 
| External HBAs | N/A |  | 
| Front drive bays | 4x 3.5-inch SAS 8-port SAS/SATA 2.5-inch 8-port U.2 NVMe Gen4 | 12-port SAS/SATA 3.5-inch 16-port SAS/SATA 2.5-inch | 
| Tuning Engine |  |  | 
| Power supplies | 800W Platinum, 1100W Titanium 800W, 1100W FTR supported |  | 
| Cooling and fans |  |  | 
| Cooling options | air cooling |  | 
| Fans | Up to four sets (dual-fan module) of hot-swappable fans | Up to six hot-swappable fans | 
| Dimensions |  |  | 
| Height | 42.8 mm (1.68 inches) | 86.8 mm (3.41 inches) | 
| Width | 482.0 mm (18.97 inches) |  | 
| Depth (with bezel) | 816.921 mm (32.16 inches) | 802.4 mm (31.59 inches) | 
| Depth (without bezel) | 815.141 mm (32.09 inches) | 801.51 mm (31.55 inches) | 
| Bezel | Optional metal bezel |  | 
Four AMD EPYC 9005 series processor options
Dell offers four specific processor SKUs for its two servers. This choice is deliberate and covers all SMB needs without overlap or unnecessary complexity. Each processor is based on AMD's Zen 5 microarchitecture and shares the same memory and PCIe characteristics at the platform level.
| Processor | Cores | Default TDP | cTDP range | Base clock | Max Boost | L3 Cache | 
|---|---|---|---|---|---|---|
| EPYC 9335 | 32 | 210W | 200-240W | 3.0 GHz | 4.4 GHz | 128 MB | 
| EPYC 9255 | 24 | 200W | 200-240W | 3.2 GHz | 4.3 GHz | 128 MB | 
| EPYC 9135 | 16 | 200W | 200-240W | 3.65 GHz | 4.3 GHz | 64 MB | 
| EPYC 9015 | 8 | 125W | 120-155W | 3.6 GHz | 4.1 GHz | 64 MB | 
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
Beyond raw performance, the R5715 with HDD storage is the ideal platform for a virtualized backup workload. To confirm this, we deployed Proxmox Backup Server on the R5715 configured with the 8-core EPYC 9015 processor and the same 12x 3.5-inch hard drive array. Proxmox is a representative example of the open-source hypervisor and infrastructure ecosystem that has seen considerable growth among SMBs, and Proxmox Backup Server, in particular, is perfectly suited to this server's storage profile.
We deployed Proxmox Backup Server 4.2.0 and used it to back up the virtual machines that power our Proxmox community Discord server environment.
In our configuration, backup and restore operations were somewhat limited by the system's 1 GbE network connection, which was the main bottleneck during large transfers. However, the platform supports simple network upgrades via OCP expansion cards, facilitating migration to 10 GbE or even 25 GbE connectivity. With a faster network, the R5715 would be capable of handling significantly higher backup throughput and restore performance, particularly in environments with larger virtual machine datasets or more demanding backup windows.
Conclusion
The Dell PowerEdge R4715 and R5715 servers owe their success to a perfectly adapted configuration. Two chassis with clearly differentiated form factors, four processor options covering all SMB needs without overlap, and a range of storage solutions broad enough to meet every need, from economical mass storage to 100% flash performance. This configuration flexibility is not purely theoretical. In our tests, the optimal configuration proved different depending on the workload. The 24-core 9255 processor offered the best compromise for transactional database performance on flash memory. The 8-core 9015 processor provided the compute power needed to deploy a Proxmox backup server with high-capacity hard drives. The R4715 with flash storage was the ideal choice for Windows shared storage, while the R5715 with hard drives was the optimal choice for capacity-focused workloads, where I/O spikes are not a limiting factor.
For SMBs and their partners, the value of these platforms lies in the ability to adapt infrastructure to the workload, rather than investing in oversized capacity. The Dell ecosystem, including iDRAC10 management, ProSupport, the security suite, and supply chain predictability, reinforces this value at the operational level. Agile IT teams and partners both benefit from a thoroughly mastered platform.
Dell's positioning regarding these servers, presented as a solution to consolidate existing infrastructures and reduce per-socket and per-core licensing costs, is confirmed by the data collected. With four processor options ranging from 8 to 32 cores on a single platform, clients and partners benefit from great flexibility to adapt acquisition, licensing, and operating costs to their actual needs. That is the entire added value, and the R4715 and R5715 servers fully deliver on it.
