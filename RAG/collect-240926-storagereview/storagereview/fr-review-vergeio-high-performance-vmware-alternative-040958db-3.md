---
id: collect-240926-storagereview/storagereview/fr-review-vergeio-high-performance-vmware-alternative-040958db-3
title: "fr-review-vergeio-high-performance-vmware-alternative-040958db"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "cost", "gpus", "liquid cooling", "memory", "nand", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-vergeio-high-performance-vmware-alternative-040958db.md
source_anchor: ""
source_lines: [24, 55]
sha256: c9334962d2c07c13f066a09fec6124f5876e5f9707018ae3eb215a489d95191e
---

# fr-review-vergeio-high-performance-vmware-alternative-040958db

The platform we used for this VergeIO HCI configuration offered 24 2.5″ NVMe bays, with six available per node. Two nodes required metadata disks, leaving us five remaining slots for corresponding SSDs to create the storage tiers. VergeIO is very flexible in managing different types of storage devices, so we created two storage pools with Solidigm SSDs. Although not used for storage capacity, VergeIO uses a metadata tier that stores logs for tracking the data reduction index. It is recommended that the tier be built around multi-drive daily-write SSDs.
For tier 1, we used 16 Solidigm 15.36 TB P5520 SSDs, which are 1DWPD SSDs, giving us a total storage capacity of 111.7 TB. For tier 2, we used the Solidigm 61.44 TB SSD, with each node receiving a single SSD. This also corresponds to the same storage capacity of 111.7 TB, as the drives were 4 times larger than the P5520 SSDs. It is important to note that the usable storage capacity of the VergeIO platform will be even greater as it offers global inline data deduplication. For the metadata tier 0, only two SSDs were needed for mirror redundancy. We used Solidigm P5620 6.4 TB SSDs installed in nodes 1 and 2.
Solidigm's portfolio offers a wide range of SSDs to cover many different use cases. For many systems, the type of SSD, whether Gen4 or Gen5, can be a deciding factor, while others may focus on the drive form factor: U.2, E1.S, or E3.S. Drive endurance also plays a major role in terms of the type of NAND used in SSD model ranges. For customers looking to deploy a VergeIO cluster, Solidigm offers products for every tier, making it easy to become a one-stop shop for flash storage.
Data density plays a major role in storage solutions. If you want to leverage the largest capacity SSDs on the market, such as the Solidigm 61.44 TB P5336, the storage system must be compatible with its native 4K block size. Customers can use QLC media for large data footprints, like our build, where we used only four drives to achieve 111 TB. VMware currently does not support QLC, which limits the platform density for customers and the ability to select the drives they want.
Test Platform – GIGABYTE H273-Z80-LAW1
To evaluate VergeIO's performance, we enlisted several leading partners to set up the cluster. We assembled a cutting-edge hardware configuration focused on density and efficiency, two key trends in modern data centers. At the heart of our test cluster is the GIGABYTE H273-Z80-LAW1, a 4N2U high-density liquid-cooled server that perfectly meets these goals.
The H273-Z80-LAW1 is an impressive server, offering up to 1024 cores and 48 TB of RAM. Our specific configuration includes eight AMD EPYC 9554 64-core processors and 2 TB of DDR5 memory, providing a solid foundation for our VergeIO cluster.
This impressive density is made possible by Giga Computing's partnership with CoolIT, which enables liquid cooling to be integrated into the system. The system fits perfectly with our existing CoolIT CDU and manifold configuration, a setup that is becoming increasingly crucial in modern data centers. This direct liquid cooling solution is not just about improving performance; it is necessary to meet the thermal requirements of today's high-power processors, GPUs, and HPC server configurations.
Our previous experiences with similar configurations have shown significant power consumption, CPU temperature reductions, and modest performance improvements from reduced thermal throttling. The efficiency and ease of use of the CoolIT system make it an ideal choice for high-density computing environments.
Each node in the H273-Z80-LAW1 supports up to 6 U.2 NVMe/SAS/SATA drives, for a total of 24 drives on the server. As noted above, we incorporated three tiers of Solidigm flash storage, allowing us to test VergeIO's ability to optimize data placement based on performance requirements. The server also features high-bandwidth PCIe Gen 5 slots for networking and expansion, ensuring first-rate connectivity for our HCI environment.
In addition to its high density and impressive expansion capacity for its size, one of the most notable features of the H273-Z80-LAW1 is its centralized management controller (CMC). This integrated solution simplifies management of the four nodes, providing a single control point for the entire server. The CMC supports IPMI 2.0 and Redfish APIs, offering flexible remote management and monitoring options.
VergeIO Performance
To measure VergeIO platform performance, we deployed 16 VMs (4 per node) to measure the overall performance of this HCI cluster. These VMs were used to orchestrate a Vdbench workload run uniformly across the cluster, all reporting to a single VM. These Vdbench sessions were also configured to test incompressible data to see how the cluster performed in the worst case, as they support data reduction. Regarding the data footprint, each VM had a 500 GB data disk, totaling 8 TB across the cluster.
We focused on four-corner performance as well as synthetic database performance using the following workloads:
- 2 MB sequential read and write
- 4K random read and write
- SQL workload
Regarding maximum sequential bandwidth, we measured 4.7 GB/s read on our TLC tier and 4.2 GB/s on the QLC tier. Moving to write bandwidth, the TLC tier measured 6.9 GB/s while the QLC tier measured 5 GB/s.
| Vdbench Workload | VergeIO Tier 1 Solidigm TLC | VergeIO Tier 2 Solidigm QLC | 
|---|---|---|
| 2 MB sequential read | 4.7 GB/s (27 ms) | 4.2 GB/s (30 ms) | 
| 2 MB sequential write | 6.9 GB/s (17.6 ms) | 5.0 GB/s (21.5 ms) | 
| 4K random read | 215 MB/s (2.6 ms) | 243 MB/s (8.2 ms) | 
| 4K random write | 263 MB/s (0.96 ms) | 200 MB (0.85 ms) | 
| SQL | 533 MB/s (0.89 ms) | 525 MB/s (0.97 ms) | 
VergeIO VSAN shows impressive performance on Solidigm TLC and QLC storage tier workloads. Sequential operations show excellent throughput, with the TLC tier reaching 6.9 GB/s for writes and 4.7 GB/s for reads. Random I/O performance is respectable, with both tiers reaching over 200 MB/s for 4K operations. The platform notably excels in SQL and VDI boot workloads, maintaining sub-millisecond latencies and high throughput.
These results indicate that the Solidigm SSDs had no trouble keeping up with the storage tiers integrated into our platform, with network and platform constraints being the main limiting factors rather than the drives themselves. VergeIO's ultraconverged infrastructure can efficiently support a wide range of enterprise applications, from large file transfers to database operations and virtual desktop environments, with the TLC tier generally offering superior performance for write-intensive scenarios.
VDI Performance
With VDI being a common workload deployed on VergeIO platforms, we wanted to test an extreme bootstorm to see how the cluster performs with 1000 VMs running simultaneously. Each VM had 2 processors, 10 GB of RAM, and a 22.04 GB disk. The VM had a standard Ubuntu XNUMX installation (not minimal) to represent a real image. Once fully booted, a script is called via systemd that uses curl to send its MAC and timestamp via HTTP to a remote collector.
Our tests revealed that the TLC and QLC drives offered very similar final performance. Examining the back-end storage data, the TLC SSDs had an advantage in total IOPS, although for this platform, the processors became the bottleneck before storage. The 1000 71 VMs were able to boot in approximately XNUMX seconds. These results highlight the benefits of combining different SSDs in the VergeIO platform. Customers can easily use QLC storage for VDI tasks, which in this case allows for fantastic density and cost-effectiveness.
We also tested high availability. We measured 138 seconds for the VM to become available after a total loss of one node, which roughly matches what VergeIO claims.
Conclusion
