---
id: collect-240926-storagereview/storagereview/fr-review-wd-gold-24tb-hdd-review-6de6b88d-2
title: "fr-review-wd-gold-24tb-hdd-review-6de6b88d"
domain: storagereview
role: reference
task: reference
actors: []
dates: ["2023-11"]
keywords: ["benchmarks", "latency", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-wd-gold-24tb-hdd-review-6de6b88d.md
source_anchor: ""
source_lines: [3, 79]
sha256: d9313c77d6f60bfb9ebfec628ede1184a9e8968e0bbd4d8ce150f82e94721d21
---

# fr-review-wd-gold-24tb-hdd-review-6de6b88d

In November 2023, WD launched the 24TB Gold Enterprise hard drive, offering 2TB more than its previous 22TB model. While this update may seem modest, it represents a significant improvement in storage capacity, particularly in terms of storage density in NAS and large-scale server configurations. Following our first HL15 test, where we put fifteen WD Gold 24TB drives through their paces, we are now ready to re-evaluate these drives within a standard 8-drive NAS configuration.
The WD Gold series is designed to meet the diverse needs of businesses, offering a comprehensive range of enterprise-class hard drives from 1TB to 24TB. These drives are designed to withstand demanding storage environments, offering MTBF of up to 2.5 million hours, vibration protection technology, and reduced power consumption thanks to HelioSeal technology for models above 12TB.
WD Gold 24TB Specifications
WD has chosen to market these high-capacity drives exclusively with a SATA interface, continuing the tradition of their helium-filled CMR design at 7200 RPM, a staple of their lineup for several years. Interestingly, the new 24TB Gold model slightly reduces power consumption, using 0.3 watts less in operation and 0.2 watts less in standby mode compared to its 22TB predecessor.
While the Gold HDD series covers capacities from 1TB to 24TB, it is worth noting that the OptiNAND feature, including the innovative OptiNAND-compatible Armor Cache, is exclusive to models of 20TB and above. This technology synergizes the benefits of enabled and disabled write cache modes, offering users a blend of performance and data protection without having to choose between the two.
Additionally, the WD Gold series offers flexibility and scalability to businesses, allowing for a customized storage configuration capable of handling intensive application workloads and supporting up to 550TB of data per year. This makes the WD Gold 24TB hard drive an attractive option for businesses looking to scale their data storage infrastructure while maintaining high levels of reliability and performance.
QNAP Storage & Snapshots Manager
Storage & Snapshots Manager is a versatile tool within the QNAP NAS interface that allows users to oversee and manage various aspects of storage, including RAID configurations, drive health, and snapshot creation for data protection.
The Storage/Snapshots pane displays an overview of storage volumes within the storage pool. In this example, it highlights the system volume and several iSCSI targets. The "Percentage Used" indicator bars provide a quick overview of space usage across different volumes.
The Storage Pool 1 Management window shows us that all drives are configured in a RAID 6 array, have a total capacity of 130.91TB, and are in "good" condition. This indicates that everything is ready for our tests.
The image below shows the Disks/VJBOD section of the Storage & Snapshots Manager. Here, the "Good" and "Ready" statuses confirm the deployment of eight high-capacity WD Gold hard drives. This pane also includes instant information such as model number, drive capacity, current speed, bus type, etc., which changes based on the drive selected in the list. This information is crucial for maintaining optimal performance and anticipating any potential issues.
WD Gold 24TB Hard Drive Specifications
| Model Number | WD241KRYZ |
| Form Factor | 3.5 inch |
| Interface | SATA 6 Gb/s |
| 512n / 512e user sectors per drive4 | 512e |
| Formatted Capacity | 24TB |
| OptiNAND Technology | Yes |
| RoHS Compliant | Yes |
| Performance | |
| Data Transfer Rate (max sustained) | 298MB/s |
| RPM | 7200 |
| Cache | 512MB |
| Power Management | |
| Average Power Requirements (W) | |
| Efficiency | 6.8W |
| Idle | 5.5W |
| Power Efficiency Index (W/TB, idle) | 0.2 |
| Reliability | |
| MTBF (hours, projected) | 2,500,000 |
| Annualized Failure Rate2 (AFR, %) | 0.35 |
| Limited Warranty | 5 years |
| Environmental | |
| Operating Temperature | 5°C to 60°C |
| Non-operating Temperature | -40°C to 70°C |
| Shock (read/write) Operating (half-sine, 2ms) | 40G / 40G |
| Non-operating (half-sine wave, 2ms) | 200G |
| Acoustics (average) | |
| Idle Mode | 20 dBA |
| Seek Mode | 32 dBA |
| Physical Dimensions | |
| Height (max) | 26.1mm |
| Length (max) | 147.0mm |
| Width (± 0.01 in) | 101.6mm |
| Weight | 1.47 lb (0.67 kg) ± 10% |
Performance
Synthetic Enterprise Workload Analysis
Our enterprise hard drive evaluation process conditions each set of drives to a steady state using the same workload with which the device will be tested. This involves a heavy load of 16 threads and an outstanding queue of 16 per thread. The device is then tested at defined intervals of multiple thread/queue depth profiles to show its performance under light and intensive usage conditions. Since hard drives quickly reach their rated performance level, only the main sections of each test are represented graphically.
Primary Preconditioning and Steady-State Tests:
- Throughput (aggregate read + write IOPS)
- Average Latency (read + write latency averaged together)
- Maximum Latency (maximum read or write latency)
- Latency Standard Deviation (read + write standard deviation averaged together)
Our synthetic enterprise workload analysis includes four profiles based on real-world tasks. These profiles were developed to facilitate comparison with our past benchmarks and widely published values, such as the maximum 4K and 8K 70/30 read and write speeds commonly used for enterprise drives.
4K
- 100 percent read or 100 percent write
- 100 percent 4K
8K70/30
- 70 percent read, 30 percent write
- 100 percent 8K
8K (Sequential)
- 100 percent read or 100 percent write
- 100 percent 8K
128K (Sequential)
- 100 percent read or 100 percent write
- 100 percent 128K
4K 100% Read/Write
Our first test measures 4K random performance. In this test, the WD Gold achieved 4,187 IOPS read and 1,417 IOPS write in SMB while displaying 4,373 IOPS read and 1,398 IOPS write in iSCSI.
For average latency, the WD Gold achieved 61.13 ms read and 180.51 ms write, and 58.52 ms read/182.92 ms write in SMB and iSCSI, respectively.
At maximum latency, the WD Gold 24TB drives recorded 759.98 ms read, 4,322.5 ms write in SMB, 1,228.1 ms read and 31,059 ms write in iSCSI.
Regarding standard deviation, in SMB, the WD Gold achieved 37.06 ms read and 373.254 ms write, while iSCSI recorded 104.53 ms read and 1,178 ms write.
8K 100% Read/Write
Both drives showed significant performance improvements in the 8K read/write test compared to the 4K read/write test. In the iSCSI configuration, the WD Gold 24TB displayed 157 IOPS read, 95 IOPS write for SMB, 215 IOPS read and 228 IOPS write in iSCSI.
8K 70 percent Read 30 percent Write
The other type of workload we run transitions from a pure sequential read/write scenario to a mixed workload, demonstrating how performance evolves from 2T/2Q to 16T/16Q. Here, SMB performance showed a minimum of 589 IOPS (coincidentally, this was the same as the Seagate Exos 24TB) and a maximum of 1,648 IOPS. With iSCSI, we see a minimum of 913 IOPS and a maximum of 1,923 IOPS, with generally consistent performance.
For average latency, we see notable peaks (in SMB) reaching up to 152.53 ms while dropping to as low as 6 ms at 2 Threads 4 Queue. We found similar scores with iSCSI, with latency as low as 6.78 ms and a maximum of 132.96 ms at 16 Threads 2 Queue.
Regarding maximum latency, WD Golds showed significantly higher latency with SMB compared to iSCSI (like the Seagate drives). Specifically in SMB, latency started at 5,395.85 ms and peaked at 5736.11 ms at 16 Threads 16 Queue, while iSCSI showed lower and more consistent latency, starting at 1167.74 ms and peaking at 3,618.97 ms.
