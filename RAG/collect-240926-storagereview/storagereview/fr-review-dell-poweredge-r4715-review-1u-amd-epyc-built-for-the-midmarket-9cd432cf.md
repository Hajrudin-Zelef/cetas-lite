---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf
title: "fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "compute", "datacenter", "energy", "ethernet", "gpu", "gpus", "intel", "license"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf.md
source_anchor: ""
source_lines: [1, 127]
sha256: 6c15b07385fabf74da072a2e5cb853a186272c46f2620df3727ffbf152c97ac2
---

# fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf

<!-- source: https://www.storagereview.com/fr/review/dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket -->

Dell's 17th generation PowerEdge lineup is already well established, and with the R4715 and R5715, it now targets SMBs and mid-sized businesses more specifically. These two single-processor servers are based on the same 5th generation AMD EPYC architecture as the entire 17th generation PowerEdge lineup. They are optimized for organizations where the right number of cores, license management, and ease of use take priority over maximum throughput. The R4715 is a more compact 1U model, designed for virtualization, large-scale databases, and network edge deployments. The R5715 adopts a 2U form factor, offering more drive bays and PCIe expansion, and is suited to configurations where storage capacity and I/O performance are essential.
The R4715 server is designed for organizations running virtualization workloads, large-scale databases, and edge computing, for which licensing efficiency and ease of use are essential. Our test model was equipped with an AMD EPYC 9335 processor, the top-tier 32-core processor available on this platform, paired with 384 GB of DDR5 memory and a BOSS RAID 1 boot configuration. In addition to its 32 cores, the 9335 offers 128 MB of L3 cache and a TDP of 210 W. Users who don't need as many cores can opt for the EPYC 9255 (24 cores), EPYC 9135 (16 cores), or EPYC 9015 (8 cores) models.
A few important points to clarify upfront about the platform: the R4715 does not support GPUs or DPUs. This is not an oversight. This platform is designed specifically for CPU-focused workloads, and Dell made deliberate compromises to limit component costs and footprint. For workloads requiring accelerators, the R6715 and R7715 models are better suited.
The R4715 offers a compact air-cooled chassis, with up to three PCIe Gen5 slots, 24 DDR5 RDIMM slots, and flexible 3.5-inch and 2.5-inch storage options, including U.2 NVMe. Additionally, it integrates iDRAC10 with OpenManage Enterprise and hardware security via a silicon-level root of trust. Network options include 25 GbE via OCP 3.0, 100 GbE, and 400 GbE via PCIe AIC. Broadcom, Intel, and NVIDIA complete the network card ecosystem; note that no Fibre Channel connectivity is officially supported. It runs on 800 W or 4715 1100 W power supplies, available in Platinum or Titanium versions, and supports fault-tolerant redundancy. For an entry-level server, the enterprise management and security features are generally satisfactory.
Dell PowerEdge R4715 Specifications
The table below highlights the physical and hardware specifications of the Dell PowerEdge R4715 platform.
| Specifications | Dell PowerEdge R4715 | 
|---|---|
| Processor |  | 
| Processor | One 5th generation AMD EPYC 9005 series processor, up to 32 cores | 
| Form factor | 1U rack server | 
| Memory |  | 
| DIMM slots | 24 DDR5 DIMM slots | 
| Maximum memory | 1.5 TB (up to 64 GB per DIMM) | 
| Memory speed | Up to 5200 MT/s | 
| Memory type | Registered DDR5 ECC RDIMM modules only | 
| Storage |  | 
| Internal controllers (RAID) | PERC H365i, H965i | 
| Internal boot | BOSS-N1 DC-MHS | 
| External HBAs | N/A | 
| Front drive bays | 4x 3.5-inch SAS 8-port SAS/SATA 2.5-inch 8-port U.2 NVMe G4 | 
| Tuning Engine |  | 
| Power supplies | Platinum 800 W, 1100 W Titanium 800 W, 1100 W FTR supported | 
| Cooling and fans |  | 
| Cooling options | air cooling | 
| Fans | Up to four assemblies (dual-fan module) of hot-swappable fans | 
| Dimensions |  | 
| Height | 42.8 mm (1.68 inches) | 
| Width | 482.0 mm (18.97 inches) | 
| Depth (with bezel) | 816.921 mm (32.16 inches) | 
| Depth (without bezel) | 815.141 mm (32.09 inches) | 
| Bezel | Optional metal bezel | 
| Networking and expansion |  | 
| OCP network options | 2 OCP 3.0 network cards (optional), 1 GbE, 10 GbE, 25 GbE Slot 2: 1×16 Gen5 OCP 3.0 Slot 5: 1×16 Gen5 OCP 3.0 | 
| Integrated network card | Dedicated 1 Gb BMC Ethernet port | 
| PCIe AIC network card | 100 GbE and 400 GbE; NDR VPI (400 GbE) | 
| PCIe slots | Up to 3 PCIe Gen5 slots (x16 connectors) Slot 1: 1×16 Gen5 full height or low profile Slot 2: 1×16 Gen5 Low Profile or 1×16 OCP3.0 Slot 4: 1×16 Gen5 full height or low profile | 
| GPU options | N/A | 
| Ports |  | 
| Front ports | 1 USB 2.0 Type-A port (optional KVM LCP) 1 USB 2.0 Type-C port (HOST/BMC Direct) 1 MiniDisplayPort (optional KVM LCP) | 
| Rear ports | 2x USB 3.1 Type-A 1x VGA Dedicated 1 Gb BMC Ethernet port | 
| Internal ports | 1x USB 3.1 Type-A | 
| Management |  | 
| Integrated management | iDRAC10, iDRAC Direct, iDRAC RESTful API with Redfish, RACADM command-line interface, Quick Sync 2 wireless module | 
| OpenManage software | OpenManage Enterprise (OME), OME Power Manager, OME Services, OME Update Manager, OME APEX AIOps Observability, OME Integration for VMware vCenter, OME Integration for Microsoft System Center, OpenManage Integration for Windows Admin Center | 
| Tools | IPMI | 
| Integrations | OpenManage Integrations: Red Hat Ansible Collections, Terraform providers | 
| Change management | Dell Repository Manager, Dell System Update, Enterprise Catalogs, Server Update Utility (SUU) | 
| Security |  | 
| Security features | Cryptographically signed firmware, data-at-rest encryption (SED with local or external key management), secure boot, secure component verification (hardware integrity check), secure erase, silicon root of trust, system lockdown (requires iDRAC10 Enterprise or Datacenter), FIPS/CC-TCG certified TPM 2.0, chassis intrusion detection, AMD Secure Encrypted Virtualization (SEV), AMD Secure Memory Encryption (SME) | 
| Operating systems and hypervisors |  | 
| Supported operating systems/hypervisors | Canonical Ubuntu Server LTS, Microsoft Windows Server with Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi | 
Dell PowerEdge R4715 Design and Installation
The R4715 is a 1U rack server measuring 1.68 mm in height, 18.97 mm in width, and 32.09 mm in depth without the optional metal bezel (32.16 mm with the bezel). The front panel includes a power button, a system identification button, a USB 2.0 Type-A port (used with the optional KVM LCP module), a USB 2.0 Type-C port for direct iDRAC access, and an optional MiniDisplayPort for the same KVM configuration. The drive bays open tool-free across the entire chassis.
Storage Configuration
The R4715 is available with three front bay configurations: 4x 3.5-inch SAS bays, 8x 2.5-inch SAS/SATA bays, or 8x U.2 NVMe Gen4 bays. The latter two configurations share the same 8-bay 2.5-inch form factor, with the choice of backplane determining drive protocol compatibility. Internal RAID is handled by the PERC H365i controller or the higher-performance PERC H965i controller. OS booting uses a dedicated BOSS-N1 DC-MHS module, which completely isolates the boot volume from the data storage pool and avoids having to allocate OS space within an active bay.
The test unit was shipped with a single 480 GB SATA SSD and two 1.92 TB U.2 NVMe drives.
Processor and Memory
The R4715 motherboard supports one 5th generation AMD EPYC 9005 series processor, up to 32 cores. Memory is handled by 24 DDR5 DIMM slots compatible only with RDIMM modules (neither UDIMM nor LRDIMM). Maximum capacity is 1.5 TB with 64 GB DIMM modules per slot, for speeds up to 5,200 MT/s.
Cooling
The R4715 is exclusively air-cooled. The processor is equipped with a five-heatpipe heatsink and a sizable fin stack that extends into the space usually empty next to the socket, thereby increasing the heat exchange surface area. Ventilation is provided by eight hot-swappable high-performance fan modules, arranged in groups, which move air from the front to the rear of the chassis. No liquid cooling system is available on this platform.
Tuning Engine
The R4715 supports hot-swappable redundant power supplies, available in two power ratings: 800 W and 1,100 W, each offered with 80 PLUS Platinum or Titanium certifications. FTR (Flex Titanium Rating) certification is also supported across the entire power supply range.
PCIe Expansion and Networking
The R4715 motherboard offers up to three PCIe Gen5 slots with x16 connectors. Slot 1 supports full-height or low-profile cards, slot 2 supports low-profile or OCP 3.0 cards, and slot 4 supports full-height or low-profile cards. Two slots for OCP 3.0 network cards (slots 2 and 5, Gen5 x16) cover 1 GbE, 10 GbE, and 25 GbE adapter options. For higher bandwidth needs, PCIe AIC network cards support up to 100 GbE and 400 GbE, with NDR VPI (400 GbE). Out-of-band management is handled via a dedicated 1 Gb BMC Ethernet port, isolating management traffic from the data plane. This platform does not support GPUs.
Rear Panel
Rear connectivity includes two USB 3.1 Type-A ports, one VGA port, and one dedicated 1 Gb Ethernet port (BMC) for iDRAC management. An additional USB 3.1 Type-A port is available internally.
iDRAC10 Management
Remote management of the R4715 uses iDRAC10, the same platform Dell offers as standard across its entire 17th generation PowerEdge lineup, including the previously reviewed PowerEdge R770 and R7725. Since the interface is identical across the lineup, administrators already familiar with iDRAC on other PowerEdge servers will immediately feel at home.
The iDRAC10 dashboard provides a comprehensive overview of the health status of all key subsystems: system status, processor, memory, cooling, storage, voltages, power supplies, batteries, and intrusion detection. The test unit shows all subsystems as operational at the time of testing. System information and firmware version details are displayed directly on the dashboard, along with license status, which is confirmed as Enterprise type on the test unit. The "Task Summary" panel tracks pending, in-progress, and completed tasks. The test unit shows completed tasks from an initial provisioning cycle, including a small number of errors and one failure, which is typical of a new deployment.
Exploring the "System Environments" section gives you access to cooling details, including the status of each fan, PWM speeds, thermal profile settings, and inlet temperature readings, all in real time. This is particularly useful for verifying airflow in dense rack configurations or for diagnosing thermal issues without physical access to the server.
Power consumption visibility follows the same principle. The "Power Information" section details power supply status, current draw, and utilization rate, along with a rolling historical graph. Administrators can thus quickly visualize average and peak consumption over time, which is useful for capacity planning and detecting workload-related power spikes, without needing an additional monitoring tool.
Together, these views make iDRAC10 a powerful out-of-band management solution that covers the entire operational lifecycle of the R4715, from initial deployment to daily monitoring, all accessible remotely via a browser or the Redfish RESTful API.
Dell PowerEdge R4715 Performance
To evaluate the performance of the Dell PowerEdge R4715, we compared it to its 2U equivalent, the Dell PowerEdge R5715. Both platforms share the same memory configuration and PowerEdge architecture, making their comparison relevant. The main difference between the two test units lies in the processor: the R4715 is equipped with a 32-core AMD EPYC 9335 processor, while the R5715 integrates an 8-core AMD EPYC 9015 processor.
It is important to note that both servers support the same range of EPYC 9005 series processors and can be configured with either chip depending on workload requirements. The difference in core count between these two units will be reflected in the results, but these results indicate the actual performance of each platform as delivered, not a comparison of maximum performance between the two.
To evaluate processor capability across systems, we used a targeted set of compute tests. y-cruncher measured raw arithmetic throughput and multithreaded floating-point performance. Blender provided a realistic rendering workload, scalable based on the number of available cores and memory bandwidth. The Phoronix test suite complemented this set of tests by adding a wider variety of CPU-intensive workloads, providing a more comprehensive view of sustained compute performance on both platforms.
Test System Specifications
- Platform: Dell PowerEdge R4715
- CPU: Single AMD EPYC 9335
- Memory: 384GB DDR5
- Storage: Boss RAID1
y-cruncher
y-cruncher is a multithreaded and scalable program capable of computing Pi and other mathematical constants to thousands of billions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application among overclockers and hardware enthusiasts.
The R4715 consistently outperformed the R5715 regardless of workload size, completing calculations approximately 2.8 to 2.9 times faster across the range of 1 to 50 billion digits. At 1 billion digits, the R4715 completed the calculation in 5.305 seconds, versus 14.537 seconds for the R5715, and this gap remained consistent despite increasing workload. At 50 billion digits, the R4715 reached 445.440 seconds, while the R5715 required 1,273.734 seconds. This result directly reflects the difference in core count: the EPYC 9335 has 32 cores versus 8 cores for the R5715's EPYC 9015.
| y-cruncher (shorter duration is better) | Dell PowerEdge R4715 (AMD EPYC 9335 32 cores \| 384 GiB RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cores \| 384 GiB RAM) | 
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
Blender 4.5 is an open-source 3D modeling software. This performance test was conducted using the Blender Benchmark command-line utility. The score is based on the number of samples per minute; the higher the value, the better the performance.
The R4715 displayed rendering throughput approximately 3.8 to 4 times higher than the R5715 across all three scenes, a slightly larger gap than in the y-cruncher test. This illustrates the strong increase in Blender's CPU rendering engine performance based on core count when parallelizing ray-tracing calculations. In the "Monster" scene, the R4715 reached 523.29 samples per minute, versus 135.21 for the R5715. In "Junkshop," the score was 355.43 versus 88.61, and in "Classroom," 264.70 versus 68.48.
| CPU performance test with Blender 4.5 (higher samples per minute is better) | Dell PowerEdge R4715 (AMD EPYC 9335 32 cores \| 384 GiB RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cores \| 384 GiB RAM) | 
|---|---|---|
| Monster | 1,076.122 samples/min | 1,076.122 samples/min | 
| Junkshop | 1,076.122 samples/min | 1,076.122 samples/min | 
| Classroom | 1,076.122 samples/min | 1,076.122 samples/min | 
Phoronix Benchmarks
The Phoronix test suite is an automated, open-source performance evaluation platform that supports over 450 test profiles and over 100 test suites via OpenBenchmarking.org. It handles the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration. We will compare the processor performance of the R4715 and R5715 using the Stream, 7-Zip, Linux kernel compilation, Apache, and OpenSSL tests.
Apache web server throughput was among the closest results in the suite, with the R4715 reaching 177,839.86 requests per second versus 123,710.75 for the R5715. Apache can maintain reasonable throughput with fewer cores when memory bandwidth is sufficient, which explains a smaller gap here than with more heavily parallelized workloads.
OpenSSL transfer throughput showed a larger gap, with the R4715 reaching 533,318,299,283 bytes/s versus 148,168,050,733 bytes/s for the R5715. Cryptographic throughput increases strongly with the number of threads, and this gap directly reflects that.
The Linux kernel compilation test revealed one of the most pronounced gaps in the suite. The R4715 completed in 379.53 seconds versus 1,244.86 seconds for the R5715; kernel compilation being one of the most direct measures of how many threads a system can execute simultaneously.
7-Zip compression reached 260,124 MIPS on the R4715 versus 98,555 MIPS on the R5715, tracking consistently with the rest of the suite.
Stream memory throughput was measured at 370,228.9 MB/s on the R4715, versus 230,123.6 MB/s on the R5715.
| Phoronix Benchmarks | Dell PowerEdge R4715 (AMD EPYC 9335 32 cores \| 384 GiB RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cores \| 384 GiB RAM) | 
|---|---|---|
| Apache requests per second | 177,839.86 | 123,710.75 | 
| OpenSSL transfer throughput (bytes/s) | 533,318,299,283 | 148,168,050,733 | 
| Kernel compilation time (seconds) (shorter time is better) | 379.531 | 1,244.86 | 
| 7-ZIP MIPS | 260,124 | 98,555 | 
| Stream throughput (MB/s) | 370,228.9 | 230,123.6 | 
Conclusion
The Dell PowerEdge R4715 is a high-performance 1U platform that fully justifies a single-processor architecture for typical SMB workloads. Businesses using virtualization, large-scale databases, and network edge deployments, for which licensing efficiency and ease of use are paramount, will find the R4715 perfectly suited. Its 1U form factor, three PCIe Gen5 slots, 24 DDR5 RDIMM slots, and flexible 2.5-inch and 3.5-inch storage options give this platform great versatility without the additional costs of a dual-processor chassis.
The performance results demonstrate the platform's capabilities. Tested with the EPYC 9335 processor, the R4715 consistently outperformed the R5715 across all benchmarks, with the core count advantage particularly visible in heavily parallelized workloads such as kernel compilation, OpenSSL, and Blender. Users who don't need 32 cores can opt for the EPYC 9255 (24 cores), EPYC 9135 (16 cores), or EPYC 9015 (8 cores), choosing a processor suited to their needs and budget.
It is important to clarify what the R4715 is not. It does not support GPUs or DPUs, a deliberate choice to limit costs and footprint. For workloads requiring accelerators, the AMD-based R6715 and R7715 models are best suited. Faster network connectivity is available via PCIe AIC, with 100 GbE and 400 GbE options.
The R4715 consistently excels in areas essential to its target use cases. It offers high compute density in a 1U chassis, optimal energy efficiency thanks to its 800 W and 1,100 W Platinum and Titanium power supplies, flexible NVMe and SAS/SATA storage configurations, and a proven iDRAC10 Enterprise management experience, perfectly integrated into the 17th generation PowerEdge lineup. For SMBs and mid-sized businesses seeking a single-socket compute platform suited to their needs, without overspending on storage or expansion capacity, the R4715 is an excellent choice and a natural complement to the R5715 in Dell's current AMD-based server lineup.
