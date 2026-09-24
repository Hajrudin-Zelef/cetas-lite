---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r5715-review-bc780bb6
title: "fr-review-dell-poweredge-r5715-review-bc780bb6"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Microsoft"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "compute", "consumer", "datacenter", "energy", "ethernet", "gpu", "gpus", "license"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r5715-review-bc780bb6.md
source_anchor: ""
source_lines: [1, 122]
sha256: 0e12883d6d467120fc405be2cc9d49c70b65c70f2f4090b9e4113a5f3d5ea118
---

# fr-review-dell-poweredge-r5715-review-bc780bb6

<!-- source: https://www.storagereview.com/fr/review/dell-poweredge-r5715-review -->

The PowerEdge R5715 is the second model in Dell's 17th generation PowerEdge lineup, dedicated to SMBs. It distinguishes itself from its 1U counterpart through different priorities. While the R4715 prioritizes compute density and core count efficiency per rack unit, the R5715 is designed around storage capacity and I/O expandability in a 2U single-socket form factor. Readers of our R4715 review will recognize the platform fundamentals: same 5th generation AMD EPYC processor family, same DDR5 memory architecture with 24 slots, and same iDRAC 10 management system. Only the task assigned to the R5715 differs slightly.
Our test server was equipped with an AMD EPYC 9015 processor, the 8-core model in the Turin lineup, paired with 384 GB of DDR5 memory and a BOSS RAID 1 configuration for boot. Our testing focused on the R5715's 12-bay 3.5-inch storage backplane, an area where the 9015 makes perfect sense. Workloads such as file sharing, data backup, and point-of-sale video surveillance don't need 32 cores; they prioritize storage density, sustained throughput, and reliable management. The 9015 helps limit power consumption and licensing costs while offering up to 288 TB of raw storage capacity in a single 2U node.
The R5715 increases the number of PCIe Gen5 ports to four, compared to three for the R4715, and adds an additional OCP 3.0 network port, providing greater flexibility to meet growing I/O demands. Both platforms support 100 GbE and 400 GbE protocols via PCIe AIC, making them perfectly suited for environments requiring high bandwidth. However, neither officially supports Fibre Channel connectivity. Additionally, neither supports GPUs or DPUs. They run on the same 800 W and 1100 W power supplies, available in Platinum and Titanium versions, with fault-tolerant redundancy and air cooling.
Dell PowerEdge R5715 Specifications
The table below highlights the physical and hardware specifications of the Dell PowerEdge R5715 platform.
| Specifications | Dell PowerEdge R5715 | 
|---|---|
| Processor |  | 
| Processor | One 5th generation AMD EPYC 9005 series processor, up to 32 cores | 
| Form Factor | 2U rack server | 
| Memory |  | 
| DIMM Slots | 24 DDR5 DIMM slots | 
| Maximum Memory | 1.5 TB (up to 64 GB per DIMM) | 
| Memory Speed | Up to 5200 MT/s | 
| Memory Type | Registered DDR5 ECC RDIMM modules only | 
| Storage |  | 
| Internal Controllers (RAID) | PERC H365i, H965i | 
| Internal Boot | BOSS-N1 DC-MHS | 
| External HBA | N/A | 
| Front Drive Bays | 12 x 3.5-inch SAS/SATA ports 16 x 2.5-inch SAS/SATA ports | 
| Engine Tuning |  | 
| Power Supplies | 800 W Platinum, 1100 W Titanium 800 W, 1100 W FTR supported | 
| Cooling and Fans |  | 
| Cooling Options | air cooling | 
| Fans | Up to six hot-swappable fans | 
| Dimensions |  | 
| Height | 86.8 mm (3.41 inches) | 
| Width | 482.0 mm (18.97 inches) | 
| Depth (with bezel) | 802.4 mm (31.59 inches) | 
| Depth (without bezel) | 801.51 mm (31.55 inches) | 
| Bezel | Optional metal bezel | 
| Networking and Expansion |  | 
| OCP Network Options | 2 OCP 3.0 network cards (optional), 1 GbE, 10 GbE, 25 GbE Slot 4: 1×16 Gen5 OCP 3.0 Slot 10: 1×16 Gen5 OCP 3.0 | 
| Integrated Network Card | Dedicated 1 Gb BMC Ethernet port | 
| PCIe AIC Network Card | 100 GbE and 400 GbE; NDR VPI (400 GbE) | 
| PCIe Slots | Up to 4 PCIe Gen5 slots (x16 connectors) Slot 2: 1×16 Gen5 Full Height Slot 3: 1×16 Gen5 Full Height Slot 7: 1×16 Gen5 Full Height Slot 9: 1×16 Gen5 Full Height | 
| GPU Options | N/A | 
| Ports |  | 
| Front Ports | 1 USB 2.0 Type-A port (optional KVM LCP) 1 USB 2.0 Type-C port (HOST/BMC Direct) 1 MiniDisplayPort (optional KVM LCP) | 
| Rear Ports | 2x USB 3.1 Type-A 1x VGA Dedicated 1 Gb BMC Ethernet port | 
| Internal Ports | 1x USB 3.1 Type-A | 
| Management |  | 
| Integrated Management | iDRAC10, iDRAC Direct, iDRAC RESTful API with Redfish, RACADM command-line interface, Quick Sync 2 wireless module | 
| OpenManage Software | OpenManage Enterprise (OME), OME Power Manager, OME Services, OME Update Manager, OME APEX AIOps Observability, OME Integration for VMware vCenter, OME Integration for Microsoft System Center, OpenManage Integration for Windows Admin Center | 
| Tools | IPMI | 
| Integrations | OpenManage Integrations: Red Hat Ansible Collections, Terraform providers | 
| Change Management | Dell Repository Manager, Dell System Update, Enterprise Catalogs, Server Update Utility (SUU) | 
| Security |  | 
| Security Features | Cryptographically signed firmware, data-at-rest encryption (SED with local or external key management), secure boot, secure component verification (hardware integrity check), secure erase, silicon root of trust, system lock (requires iDRAC10 Enterprise or Datacenter), FIPS/CC-TCG certified TPM 2.0, chassis intrusion detection, AMD Secure Encrypted Virtualization (SEV), AMD Secure Memory Encryption (SME) | 
| Operating Systems and Hypervisors |  | 
| Supported Operating Systems/Hypervisors | Canonical Ubuntu Server LTS, Microsoft Windows Server with Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi | 
The Dell PowerEdge R5715 is a single-processor 2U rack server based on the 5th generation AMD EPYC 9005 series platform. Designed as a high-performance storage platform for businesses requiring high capacity and reliable I/O connectivity without the overhead of a dual-processor architecture, the R5715 targets workloads such as databases, file shares, backups, and virtualization, where a single powerful EPYC processor handles the load more efficiently than two previous-generation processors. We also used this chassis for our power consumption comparison between hard drives and Micron 6600 ION flash memory, replacing eight 30 TB hard drives with a single 245 TB SSD. Supporting up to 288 TB of raw storage and featuring four PCIe Gen5 expansion slots, the R5715 delivers performance well above its price category.
Exterior and Front Panel
The R5715 ships with an optional metal bezel featuring Dell's iconic hexagonal pattern. This bezel fits perfectly onto the chassis and reveals the front panel controls on the right side: a power button, a USB 2.0 Type-C port for direct BMC access, an iDRAC Direct port, and a system identification button. Without the bezel, the chassis measures 3.41 cm in height, 18.97 cm in width, and 31.55 cm in depth, allowing it to fit into standard 2U racks. Build quality is impeccable, with tool-less drive bay latches and blue retention clips used consistently on internal components for quick access.
Storage Configuration
The tested unit is equipped with a 12 x 3.5-inch SAS/SATA front bay, with four bays occupied by 20 TB 6 Gb/s 7,200 RPM SATA hard drives and eight bays free for future expansion. An alternative configuration with a 16 x 2.5-inch SAS/SATA backplane is also available, depending on needs. RAID management is handled by the internal PERC H365i controller or the more powerful PERC H965i. Boot is managed separately by a dedicated BOSS-N1 DC-MHS module at the rear, isolating the operating system from the data. This clever design avoids the common mistake of running the operating system and data storage on the same bay.
Processor and Cooling
The R5715 is a single-socket platform based on the AMD EPYC 9005 series processor, supporting up to 32 cores. Its imposing heatsink is a finned tower model incorporating copper heat pipes, secured to the SP5 socket by six captive screws. Cooling is entirely air-based; up to six hot-swappable fans circulate air from front to rear of the chassis. Liquid cooling is not compatible with this platform.
Memory
The R5715 motherboard has 24 DDR5 DIMM slots divided into two groups on either side of the processor socket. This platform is exclusively compatible with RDIMM modules; UDIMM and LRDIMM modules are not supported. Maximum capacity reaches 1.5 TB with 64 GB RDIMMs per slot, operating at a maximum frequency of 5,200 MT/s. The tested model ships with several slots occupied, leveraging EPYC's multi-channel memory architecture to deliver high aggregate bandwidth across the entire memory subsystem.
PCIe Expansion and Networking
The R5715 offers up to four full-height PCIe Gen5 x16 slots (slots 2, 3, 7, and 9), distributed across five riser slots (Risers 1 through 5) visible inside the chassis. Two additional slots for OCP NIC 3.0 network cards (slots 4 and 10, Gen5 x16) support OCP 1 GbE, 10 GbE, or 25 GbE network adapters. For high-speed connectivity, PCIe AIC network cards support up to 100 GbE and 400 GbE, with NDR VPI (400 GbE). A dedicated 1 Gb BMC Ethernet port is integrated into the rear panel for out-of-band iDRAC management. The R5715 does not offer GPU options; it is a storage and compute platform, not an acceleration chassis.
iDRAC10 Management
Remote management of the R5715 is handled by iDRAC10, the same platform Dell offers as standard across its entire 17th generation PowerEdge lineup, including the PowerEdge R770 and R7725 we previously tested. Since the interface is identical across the lineup, administrators familiar with iDRAC on other PowerEdge platforms will immediately feel at home.
The iDRAC10 dashboard provides a comprehensive, instant overview of the health status of each major subsystem: system status, processor, memory, cooling, storage, voltages, power supplies, batteries, and intrusion detection. The test unit indicates that all subsystems were operational at the time of testing. System information and firmware version details are displayed directly on the dashboard, along with license status, which on the test unit is confirmed as Enterprise type. The "Task Summary" panel tracks pending, in-progress, and completed tasks. The test unit displays completed tasks from an initial provisioning cycle, a few with errors and one that failed, which is typical of a new deployment.
Exploring the "System Environments" section gives you access to cooling details, including the status of each fan, PWM speeds, thermal profile settings, and inlet temperature readings, all in real time. This feature is particularly useful for verifying airflow in dense rack configurations or for diagnosing thermal issues without having to physically access the server.
Power consumption visibility follows the same principle. The "Power Information" section details power supply status, current draw, and utilization rate, along with a rolling historical graph. Administrators can quickly visualize average and peak consumption over time, which is valuable for capacity planning and identifying workload-related power spikes without needing an additional monitoring tool.
Together, these views make iDRAC10 a powerful out-of-band management solution that covers the entire operational lifecycle of the R5715, from initial deployment to daily monitoring, all accessible remotely via a browser or the Redfish RESTful API.
Dell PowerEdge R5715 Performance
To evaluate the performance of the Dell PowerEdge R5715, we compared it to its 1U counterpart, the Dell PowerEdge R4715. Both platforms share the same memory configuration and PowerEdge architecture, making them an obvious point of comparison. The main difference between the two test units lies in the processor. The R4715 was equipped with a 32-core AMD EPYC 9335 processor, while the R5715 had an 8-core AMD EPYC 9015 processor.
It's important to note that both platforms support the same range of EPYC 9005 series processors and can be configured with either chip as needed. The core count difference between these two units will be reflected in the results, but these results correspond to the actual performance of each platform, not a maximum performance comparison.
In order to stress the processors of both systems, we used a targeted set of compute benchmarks. y-cruncher evaluated raw arithmetic throughput and multithreaded floating-point performance. Blender provided a realistic rendering workload, scalable based on available core count and memory bandwidth. The Phoronix test suite completed this set with a broader collection of CPU-intensive workloads, providing a more comprehensive view of sustained compute performance on both platforms.
Test System Specifications
- Platform: Dell PowerEdge R5715
- CPU: Single AMD EPYC 9015
- Memory: 384GB DDR5
- Storage: Boss RAID1
y-cruncher
y-cruncher is a multithreaded and scalable program capable of computing Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application among overclockers and PC hardware enthusiasts.
The R5715 showed predictable performance relative to the R4715, regardless of workload size. At 1 billion digits, the R5715 finished in 14.537 seconds versus 5.305 seconds for the R4715, and this gap held steady. At 50 billion digits, the R5715 reached 1,273.734 seconds while the R4715 finished in 445.440 seconds, approximately 2.8 to 2.9 times faster across the entire range from 4715 to 50 billion digits. Despite having only 8 cores, the EPYC 9015 is a dedicated server processor with significantly higher memory bandwidth and cache than a typical desktop processor. It thus remains far more powerful than most consumer processors on the same workloads.
| Y-cruncher (shorter duration is better) | Dell PowerEdge R4715 (AMD EPYC 9335 32 cores \| 384 GB RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cores \| 384 GB RAM) | 
|---|---|---|
| 25 million | 0.11 seconds | 0.25 seconds | 
| 50 million | 0.23 seconds | 0.51 seconds | 
| 100 million | 0.46 seconds | 1.08 seconds | 
| 250 million | 1.22 seconds | 3.00 seconds | 
| 500 million | 2.49 seconds | 6.60 seconds | 
| 1 billion | 5.30 seconds | 14.53 seconds | 
| 2.5 billion | 14.58 seconds | 41.32 seconds | 
| 5 billion | 32.38 seconds | 92.99 seconds | 
| 10 billion | 71.54 seconds | 202.87 seconds | 
| 25 billion | 203.40 seconds | 576.87 seconds | 
| 50 billion | 445.44 seconds | 1,273.73 seconds | 
Blender 4.5
Blender 4.5 is an open-source 3D modeling application. This benchmark was performed using the Blender Benchmark CLI utility. The score is measured in samples per minute, with higher values indicating better performance.
Blender results follow a similar trend to Y-Cruncher: the R4715's higher core count translates directly into higher rendering throughput. On the Monster scene, the R4715 reached 523.29 samples per minute versus 135.21 for the R5715. The Junkshop scene scored 355.43 versus 88.61, and the Classroom scene 264.70 versus 68.48 for the R5715. Across all three scenes, the R4715 showed approximately 3.8 to 4 times higher rendering throughput than the R5715, a slightly larger gap than on Y-Cruncher. This illustrates the importance of core count for Blender CPU rendering when parallelizing ray tracing calculations within a single scene.
| CPU Performance Test with Blender 4.5 (higher samples per minute is better) | Dell PowerEdge R4715 (AMD EPYC 9335 32 cores \| 384 GB RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cores \| 384 GB RAM) | 
|---|---|---|
| Monster | 1,076.122 samples/min | 1,076.122 samples/min | 
| Junkshop | 1,076.122 samples/min | 1,076.122 samples/min | 
| Classroom | 1,076.122 samples/min | 1,076.122 samples/min | 
Phoronix Benchmarks
Phoronix Test Suite is an automated, open-source benchmarking platform supporting over 450 test profiles and more than 100 test suites via OpenBenchmarking.org. It manages the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration. Here we will compare the R5715 and R4715 processor performance using the Stream, 7-Zip, Linux kernel compilation, Apache, and OpenSSL tests.
In terms of Apache web server throughput, the R4715 reached 177,839.86 requests per second, versus 123,710.75 for the R5715, one of the closest results in the entire series. Apache's ability to deliver acceptable performance even with fewer cores, provided memory bandwidth is sufficient, explains why the gap here is smaller than for more heavily parallelized workloads.
OpenSSL transfer throughput showed a larger gap, with the R4715 reaching 533,318,299,283 bytes per second versus 148,168,050,733 bytes per second for the R5715. Cryptographic throughput is one of the workloads that scales most rapidly with thread count, and this gap clearly reflects that.
The Linux kernel compilation test revealed one of the most pronounced gaps in the suite, with the R4715 finishing in 379.53 seconds versus 1,244.86 seconds for the R5715. Kernel compilation is one of the most direct measures of how many threads a system can run simultaneously.
7-Zip compression reached 260,124 MIPS on the R4715 versus 98,555 MIPS on the R5715, consistent with the results obtained across the rest of the suite.
Stream memory throughput was 370,228.9 MB/s on the R4715, versus 230,123.6 MB/s on the R5715.
| Phoronix Benchmarks | Dell PowerEdge R4715 (AMD EPYC 9335 32 cores \| 384 GB RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cores \| 384 GB RAM) | 
|---|---|---|
| Apache requests per second | 177,839.86 | 123,710.75 | 
| OpenSSL transfer throughput (bytes/s) | 533,318,299,283 | 148,168,050,733 | 
| Kernel compilation time (seconds) (shorter time is better) | 379.531 | 1,244.86 | 
| 7-ZIP MIPS | 260,124 | 98,555 | 
| Stream throughput (MB/s) | 370,228.9 | 230,123.6 | 
Conclusion
The Dell PowerEdge R5715 is a storage-dedicated 2U platform, perfectly designed and fully justifying the use of a single-processor chassis in certain workload contexts. Businesses running file services, backup targets, video surveillance systems, or databases, and prioritizing disk density and I/O expandability over raw compute power, will find the R5715 an ideal solution. Its 12-bay 3.5-inch backplane, supporting up to 288 TB of raw capacity, combined with four PCIe Gen5 slots and support for two OCP 3.0 network cards, gives the platform significant room for growth without requiring the move to a more expensive dual-processor chassis.
The performance results are unequivocal. Tested with the EPYC 9015 processor, the R5715 shows significantly lower performance than the R4715 (32 cores) across all benchmarks, as expected. However, this comparison is somewhat beside the point. The R5715 is not designed for compute-intensive tasks, and the EPYC 9015 is not the processor Dell envisions for most of its customers. Configuring the R5715 with a higher-core-count EPYC 9005 processor can considerably narrow this gap, and the platform architecture is fully compatible.
The R5715 excels in the areas essential to its target use cases: storage density, expansion flexibility, energy efficiency, and management. iDRAC10 Enterprise offers a proven and consistent out-of-band management experience, directly from the 17th generation PowerEdge lineup, reducing operational costs for teams that have already invested in Dell's management suite.
For SMB and midmarket buyers looking to consolidate their storage workloads onto a suitable single-processor platform without oversizing their compute capabilities, the R5715 is an excellent choice and a natural complement to the R4715 in Dell's current AMD-based lineup.
