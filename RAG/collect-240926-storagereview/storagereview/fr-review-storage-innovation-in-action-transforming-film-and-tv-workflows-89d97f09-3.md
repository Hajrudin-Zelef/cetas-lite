---
id: collect-240926-storagereview/storagereview/fr-review-storage-innovation-in-action-transforming-film-and-tv-workflows-89d97f09-3
title: "fr-review-storage-innovation-in-action-transforming-film-and-tv-workflows-89d97f09"
domain: storagereview
role: reference
task: reference
actors: []
dates: []
keywords: ["energy", "latency"]
source: docs/RAG/clean_en/storagereview/fr-review-storage-innovation-in-action-transforming-film-and-tv-workflows-89d97f09.md
source_anchor: ""
source_lines: [22, 53]
sha256: 95f75873fa3e0b61b489957616da17f75f3633d67588e0912299038b0e26d424
---

# fr-review-storage-innovation-in-action-transforming-film-and-tv-workflows-89d97f09

Switching to Tuxera's Fusion software had a considerable impact, with no changes needed on the client side. Looking at the TCP protocol, we measured 11.1 GB/s and 11.1 GB/s for each client, giving us 22.2 GB/s of read bandwidth. This was the limit of the 100 GbE connection to each host. For write bandwidth, we measured 5.7 GB/s and 5.8 GB/s, giving us a total of 11.5 GB/s. Read latency averaged 6 ms while write latency was 11.65 ms.
In addition to TCP, Tuxera's Fusion File Share also supports RDMA. We measured the Fusion RDMA protocol, which gave us read bandwidth measuring 11.6 GB/s and 11.6 GB/s for each host, or 23.2 GB/s. Write bandwidth was 5 GB/s and 5.2 GB/s, for a total of 10.2 GB/s. Read latency in this configuration was 5.8 ms, while write latency was 13.2 ms.
Comparing Samba to Fusion showed huge gains for Windows clients. Read bandwidth was nearly 3x higher, with latency only 33% of that offered by Samba. Write bandwidth was also 2.3x higher, with latency only 44% of that measured with Samba.
| Gateway | Metric | Client1 | Client2 | Total | 
|---|---|---|---|---|
| Samba | Read bandwidth | 4GB / s | 3.8GB / s | 7.8GB / s | 
|  | Read latency | 16.77ms | 17.68ms | 17.23ms | 
|  | Write bandwidth | 2.4GB / s | 2.6GB / s | 5GB / s | 
|  | Write latency | 27.5ms | 25.7ms | 26.6ms | 
| Fusion TCP | Read bandwidth | 11.1GB / s | 11.1GB / s | 22.2GB / s | 
|  | Read latency | 6ms | 6ms | 6ms | 
|  | Write bandwidth | 5.7GB / s | 5.8GB / s | 11.5GB / s | 
|  | Write latency | 11.8ms | 11.5ms | 11.65ms | 
| FusionRDMA | Read bandwidth | 11.6GB / s | 11.6GB / s | 23.2GB / s | 
|  | Read latency | 5.8ms | 5.8ms | 5.8ms | 
|  | Write bandwidth | 5GB / s | 5.2GB / s | 10.2GB / s | 
|  | Write latency | 13.4ms | 13ms | 13.2ms | 
Huge SSDs keep pace with M&E
Massive SSDs like Solidigm's 61.44 TB P5336 bring substantial benefits to the media and entertainment sectors, where speed and reliability are crucial for handling large files such as high-definition video, complex graphics, and extended audio tracks. Unlike traditional Hard Disk Drives (HDDs), SSDs offer faster data access times, higher read/write speeds, and a greater number of I/O operations per second. This performance advantage enables more efficient editing, rendering, and processing workflows, significantly reducing the time needed to load and use large media files. The absence of moving parts in SSDs improves their reliability and durability, making them less susceptible to mechanical failures and data loss (critical issues when it comes to valuable media content).
The adoption of dense SSDs in media production and post-production environments streamlines workflows, enabling real-time editing, color grading, and effects processing without compromising quality or efficiency. These drives can seamlessly handle multiple 4K video streams, eliminating the need for proxy files or low-resolution placeholders. Additionally, the compact format of SSDs, combined with their high storage capacities, simplifies data management by allowing entire projects to be stored on a single drive or minimal array, facilitating access to and management of large volumes of data.
Beyond performance and capacity benefits, massive SSDs contribute to a more conducive production environment through their silent operation, low heat output, and energy efficiency. These features are particularly useful in densely packed or mobile editing suites, where noise reduction and cooling are constant concerns. The energy efficiency of SSDs not only reduces operational costs but also helps create a cooler and quieter workspace, improving overall productivity and comfort for media professionals. Essentially, massive SSDs are transforming the media and entertainment landscape, enabling faster, more reliable, and more efficient production processes capable of keeping pace with the growing demand for high-quality digital content.
Bridge Digital is a big fan
In our own testing with the CheetahRAID platform, we found impressive benefits for M&E workloads. But we wanted another industry opinion, to see if our conclusions were consistent with those who are deeply involved in on-set data management. We contacted our friend Richie Murray, founder and president of Bridge Digital.
Bridge Digital is a company with expertise in digital video workflows and the technologies that make them work better.
They help creators and digital content owners build infrastructure to efficiently create, manage, distribute, and monetize their video assets. Bridge Digital's solutions cover the entire digital media workflow, from ingestion to final delivery.
Richie notes that "performance, compatibility, and support are our clients' absolute priorities. That's exactly what this solution brings together. The data portability offered by the CheetahRAID system is unmatched in the industry, without sacrificing the speed and reliability of less 'ship-friendly' hardware."
Richie continued, saying: "Fusion File Share is more performant and is fully compatible with existing networked creative workstations. Moreover, the massive Solidigm SSDs combined with Graid's RAID implementation both contribute to the idea that 'this solution is greater than the sum of its parts.' The combined solution is extremely compelling for modern M&E workloads."
Conclusion
Tuxera's Fusion File Share sequential speeds are ideal for quickly downloading and verifying footage, especially when shooting with many cameras. Fast storage is not unfamiliar in on-set use, and speed alone won't impress many, but the ability to be networked and offload multiple systems simultaneously without slowing down is a particular workflow in which Tuxera Fusion excels.
Most of the time on set, the limiting factor is not storage but the camera media, and this bottleneck becomes problematic, especially when you start having double-digit camera counts. It is not uncommon for multiple systems to offload multiple cards simultaneously to independent drives, which must then be consolidated either on set to an identical pair of master drives, or at the post relay on their servers.
The ability to network many systems to the same Tuxera Fusion storage pool and offload them simultaneously without slowing down constitutes a huge time savings. More importantly, a standard NAS would struggle to efficiently scale this shared storage with many clients if it used Samba, enabling offloading, transcoding, management, review, and quality control to all happen simultaneously.
While we have discussed the benefits offered by this solution on set, it is also essential to understand that all of this is very low in complexity. The CheetahRAID server, Solidigm QLC SSDs, Graid Technologies RAID management, and Tuxera protocol are all very easy to configure in a basic Linux installation. The entire solution is assembled in minutes, not hours, and brings together the best technologies for M&E professionals.
