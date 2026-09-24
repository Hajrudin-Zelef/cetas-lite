---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662
title: "fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["memory", "accelerator", "amd", "benchmark", "compute", "cost", "energy", "ethernet", "gpu", "gpus", "latency", "liquid cooling"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662.md
source_anchor: ""
source_lines: [1, 178]
sha256: 187d2ddab5c27535492720782fe1b73be008ea7d2e9a21b14f75e866432f5f01
---

# fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662

<!-- source: https://www.storagereview.com/fr/review/dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storage -->

The Dell PowerEdge R7715 is a 2U single-socket server designed for traditional enterprise workloads and software-defined storage. It uses AMD EPYC 9005 series processors, supporting up to 160 cores and 12 memory channels, and has 24 DDR5 DIMM slots for up to 6 TB of high-speed memory. The R7715 also supports PCIe Gen5 on its front storage bays and risers, delivering high-speed connectivity for SSDs, GPUs, and network adapters. In high-density configurations, some Gen5 SSDs can operate with x2 link widths, a design choice that balances I/O scalability and bandwidth without requiring PCIe switches. With configuration options for EDSFF E3.S and U.2 NVMe drives and support for multiple GPUs or accelerators, the R7715 delivers a perfect balance between compute density and flexibility in an energy-efficient form factor.
Where the Dell PowerEdge R7715 fits in the AMD server lineup
Dell's AMD-based PowerEdge rack server lineup includes four models: R6715, R7715, R6725, and R7725, designed to scale from density-optimized single-socket systems to high-performance dual-socket systems. The PowerEdge R7715 sits in the middle of this range, with its 2U single-socket system offering a perfect balance between compute power, memory capacity, and I/O flexibility.
| Rack server | R6715 | R7715 | R6725 | R7725 | 
| Form factor | 1U | 2U | 1U | 2U | 
| Processor | 1x AMD EPYC9005 | 1x AMD EPYC9005 | 2x AMD EPYC9005 | 2x AMD EPYC9005 | 
| Maximum core count | Up to 160 | Up to 160 | Up to 192 | Up to 192 | 
| Max memory | 3 TB (24 DDR5 DIMM) | 6 TB (24 DDR5 DIMM) | 3 TB (24 DDR5 DIMM) | 6 TB (24 DDR5 DIMM) | 
| NVMe drives | Up to 22 | Up to 40 | Up to 22 | Up to 40 | 
| PCIe Gen5 slots | Up to 3 | Up to 8 | Up to 3 | Up to 8 | 
| Accelerator support | 3x SW | 3 x DW, 6 x SW | 3x SW | 2 x DW, 6 x SW | 
| Target workloads | Analytics, virtualization, SDS | Analytics, virtualization, SDS | HPC, VDI, virtualization | HPC, VDI, virtualization | 
The R6715 and R6725 are both 1U systems. However, the R6725 is a dual-socket system, while the R6715 is a single-socket system, similar to the R7715. The main difference lies in expandability and maximum memory capacity. The R7715 supports up to 6 TB of memory, twice the capacity offered by the R6715, and offers more extensive NVMe drive and storage configuration options. It also supports more PCIe Gen5 slots and accelerator cards, making it better suited for I/O- or GPU-intensive environments.
The R6725 and R7725 offer greater raw compute potential in dual-socket, with scalability up to 192 cores. The R7725, in particular, is the most powerful of the four, offering up to 40 NVMe drives, 8 PCIe Gen5 slots, and support for high-density E3.S storage. However, this entails higher power consumption and potentially higher costs. For many use cases, such as data analytics or software-defined storage, where single-socket performance is sufficient, the R7715 offers a good balance between core count, storage density, and thermal envelope.
Regarding use cases, the R7715 is ideal for businesses seeking Zen5 performance, generous DDR5 memory bandwidth, and scalable NVMe storage without having to migrate to dual-socket systems. It is perfectly suited for dense virtualization, modern databases, and analytical workloads that benefit from a high core count, without requiring the complexity of two processors.
Dell PowerEdge R7715 front bay storage configuration options
The Dell PowerEdge R7715 server offers a comprehensive range of storage backplane configurations, addressing traditional enterprise workloads and next-generation NVMe deployments. For compute-only configurations, buyers can opt for configurations without a backplane or with a minimal configuration of 2 U.2 NVMe SSDs, making the system suitable for boot-focused or highly customized deployments. Additionally, the system supports up to 8 universal 2.5-inch bays or 8 EDSFF E3.S Gen5 NVMe drives, giving businesses a choice between compatibility with existing motherboards and cutting-edge PCIe Gen5 performance. Mid-range options include 12 3.5-inch SAS/SATA bays or 16 2.5-inch SAS/SATA SSDs. Dell also offers hybrid configurations, such as 16 SAS/SATA bays and 8 U.2 NVMe SSDs, for environments requiring a combination of cost-effective storage and high-speed cache or tiering layers.
Dell supports high-capacity NVMe configurations for dense, high-performance use cases, with scalability up to 32 or even 40 EDSFF E3.S Gen5 drives (though some options are planned for 2025). These Gen5 NVMe options are ideal for high-throughput, low-latency workloads, such as software-defined storage or real-time analytics. Interestingly, the R7715 does not support tri-mode RAID configurations, meaning the controller must separate SAS, SATA, and NVMe drives. This strategic limitation simplifies validation but reinforces Dell's commitment to dedicated, high-performance platforms. Overall, the R7715's storage configurability makes it a highly flexible server, capable of adapting to both classic deployments and demanding all-flash NVMe arrays.
A key point regarding storage performance is understanding how Dell designed the E3.S storage backplane for high storage density configurations, particularly in terms of PCIe lane allocation per SSD. To avoid resorting to a PCIe switch, configurations like the R7715 we are examining only include Gen2 x5 width SSD bays. Systems equipped with 8 or 16 E3.S bays use x4 width bays, which improves individual SSD performance.
Dell PowerEdge R7715 expansion card and PCIe card
The Dell PowerEdge R7715 offers several expansion card configuration options, allowing businesses to tailor the server's expansion capabilities to their specific workload. These options allow for adding high-speed network or storage accelerators. One of the simplest options includes 2 full-height x16 Gen5 slots and support for a second OCP network card, providing a clear and efficient configuration for basic expansion needs. This is an ideal solution for businesses that need high-speed cards without overloading the server with PCIe devices.
The system can be configured for more demanding environments with 4 x 16 full-height Gen5 slots, with a second OCP network card included. This configuration is ideal for use cases requiring multiple expansion cards, such as dual network cards combined with GPUs or other accelerators. Imagine pushing the limits of PCIe expansion. In this case, the R7715 supports a high-density configuration with 6 x 16 full-height Gen5 slots, optimizing available bandwidth and slot count, making it ideal for I/O-intensive or multi-accelerator workloads.
For users requiring more flexibility in terms of lane width and card type, a hybrid card option is available: 2 x 16 full-height Gen5 slots and 6 x 8 full-height Gen5 slots. This configuration is handy for combining high-bandwidth devices with others that can operate efficiently with fewer PCIe lanes. The R7715 also supports a configuration with 2 x 16 low-profile Gen5 slots and 3 x 16 full-height double-width Gen5 slots, designed to accommodate larger, power-hungry accelerator cards, with optimized airflow spacing.
Finally, Dell offers a GPU-focused version of this configuration for liquid-cooled environments, with 2 x 16 extra-low-profile Gen5 slots and 3 x 16 full-height double-width Gen5 slots, validated for direct liquid cooling configurations. This ensures compatibility with the physical and thermal design requirements of liquid-cooled chassis and components, making the R7715 a flexible platform for high-performance deployments.
Dell PowerEdge R7715 memory and power configurations
The PowerEdge R7715 supports up to 24 DDR5 DIMM slots, delivering an impressive 6 TB of registered ECC DDR5 memory capacity. This memory configuration can handle up to 5200 MT/s and is designed to be scalable, offering high bandwidth and low latency. It is therefore ideal for use cases such as virtualization, database management, or intensive data analytics.
The system also offers various power supply options. It supports hot-pluggable redundant power supplies from 800 W to 3200 W, providing flexibility suited to different needs and configurations. You have Titanium-class high-performance units (3200 W, 2400 W, 1800 W, 1500 W, 1100 W, and 800 W), as well as Platinum-class options (1100 W and 800 W). Special configurations are also available, such as 277 VCA and CCHT (3200 W and 1500 W), as well as a DC power supply (1400 W -48--60 VCC).
Dell PowerEdge R7715 and iDRAC 10
iDRAC, or Integrated Dell Remote Access Controller, is Dell's integrated remote management tool that simplifies monitoring, updating, and troubleshooting of PowerEdge servers, such as the R7715, without requiring physical presence. With the latest iDRAC10 version, Dell has made several improvements to enhance security and usability. It now incorporates a dedicated security processor with an integrated root of trust, improved encryption algorithms, and device-level attestation, making server management more secure than ever.
The user interface has also been redesigned for a more consistent experience across all Dell Technologies consoles, with simplified navigation that makes managing your server even more intuitive. Additionally, iDRAC10 allows for creating custom user roles and introduces a simplified licensing structure, specifically for 17th generation PowerEdge. AC power recovery is now managed directly by iDRAC, instead of being controlled by the BIOS, giving administrators more centralized control, an appreciated feature.
Dell PowerEdge R7715 specifications
| Feature | PowerEdge R7715 | 
| Processor | One 5th generation AMD EPYC 9005 Series processor with up to 160 cores for the Zen5 processor | 
| Chipset | AMD chipset | 
| Accelerators | Up to three 400 W double-width GPUs or six 75 W single-width GPUs | 
| Memory |  | 
| DIMM module speed | Up to 5200 MT/S | 
| Memory type | RDIMM | 
| Memory module slots | 24 DDR5 DIMM slots | 
|  | Supports only registered DDR5 ECC DIMM modules. | 
| Storage |  | 
| Front bays |  | 
| Rear bays | N/A | 
| Storage controllers |  | 
| Internal controllers | PERC H365i, H965i, H975i | 
| External controllers | HBA465e, H965e | 
| Software RAID | N/A | 
| Internal boot |  | 
| Power supply |  | 
| Cooling options |  | 
| Fans | Up to six hot-plug Gold/Very High Performance fans | 
| Ports |  | 
| Network options | Dedicated 1 Gb BMC Ethernet port | 
|  | 2 OCP NIC 3.0 cards (optional) | 
| Front ports | 1 x USB 2.0 Type-A (optional LCP KVM) | 
|  | 1 x USB 2.0 Type-C (HOST/BMC Direct) | 
|  | 1 x Mini-DisplayPort (optional LCP KVM) | 
| Rear ports | Dedicated 1 Gb BMC Ethernet port | 
|  | 2 x USB 3.1 | 
|  | 1 x VGA | 
| Internal ports | 1 USB 3.0 port (optional) | 
| Slots |  | 
| PCIe | Up to eight PCIe Gen5 slots | 
| Form factor | 2U rack server | 
| Height | 86.8 mm (3.41 inches) | 
| Width | 482.0 mm (18.97 inches) | 
| Depth | 802.4 mm (31.59 inches) with power handle | 
| Weight | Maximum 28.68 kg (63.22 lb) | 
| Bezel | Optional metal bezel | 
| System management |  | 
| Embedded management |  | 
| OpenManage Console |  | 
| Mobility | N/A | 
| Tools | IPMI | 
| Change Management |  | 
| OpenManage Integrations |  | 
| Security |  | 
| Operating systems and hypervisors |  | 
Dell PowerEdge R7715 design and construction
The PowerEdge R7715 features a clean design engineered for efficient cooling and consistent performance. Its modular, tool-less installation simplifies maintenance and upgrades, allowing you to maintain optimal operation with minimal downtime. It is also equipped with an optional locking bezel, which provides enhanced security and gives the front a clean, professional look. Easy to remove, it allows quick access for maintenance or upgrades, for simple and convenient use.
Once the front panel is removed, you can access the entire front panel. On the right side, the control panel includes the power button, a USB port (for connecting external devices such as USB keys or devices for maintenance or direct access), a micro iDRAC Direct port, and the iDRAC Direct status indicator.
The front and center consist of 32 1.6 TB EDSFF Gen5 E3.S NVMe drives (at least for our configuration), each operating on a PCIe Gen5 x2 interface. Although x4 connectivity is more common, this configuration leverages the higher bandwidth of Gen5 to maintain high throughput even at x2, thus preserving valuable PCIe lanes for other expansion needs without resorting to PCIe switches. Nevertheless, these drives are hot-swappable, allowing them to be added or replaced without shutting down the system. This design minimizes downtime and offers flexibility to adjust storage capacity as needed. The ventilation panel is located between the drive rows to ensure consistent airflow and maintain optimal operating temperatures.
The bottom of the rear panel houses several ports, including the BOSS-N1 DC-MHS module and a dedicated BMC Ethernet port for remote management via Open Server Manager. This configuration allows administrators to monitor and control the server from a separate network connection, essential for maintaining control in the event of a network outage or software failure.
Additionally, two USB 3.1 ports (9-pin and 3.0 compatible) provide high-speed connectivity for external devices such as USB keys or external hard drives. A VGA port also allows connecting displays, ensuring compatibility with existing systems. The system is equipped with two power supply units (PSU1 and PSU2) on each side of the rear panel, providing built-in redundancy. As always, this configuration ensures uninterrupted operation even in the event of a power supply failure, an essential feature for enterprise environments where availability is critical.
Removing the top panel of the Dell PowerEdge R7715 reveals a well-organized interior. The first thing you notice is likely the large black ventilation shrouds that cover most components and guide air to critical parts to ensure optimal cooling. PCIe expansion cards can be inserted into the metal risers, such as Riser 3 and Riser 5. Many blue plastic clips and tabs are scattered throughout, designed to simplify upgrades and maintenance as much as possible, without tools. Everything seems designed to ensure adequate airflow and preserve component accessibility.
Once the covers are removed, the internal board of the Dell PowerEdge R7715 is visible. It presents a well-organized configuration, including the AMD EPYC processor, DIMM slots, PCIe risers, BOSS-N1 DC-MHS modules, and fans.
At the center of the motherboard is the single-socket AMD EPYC 9665P processor, topped with an imposing heatsink featuring copper heat pipes for efficient cooling. Of the 24 available slots, 12 DIMM slots, equipped with 768 GB of RAM, surround the processor, providing significant room for future upgrades.
The image below shows a close-up of the PowerEdge R7715's BOSS-N1 DC-MHS module, equipped with a Micron 7450 NVMe SSD. This drive uses Micron's advanced 176-layer NAND technology, offering excellent latency and PCIe Gen4 performance. Designed for reliable and secure booting, the BOSS-N1 is generally dedicated to the server's operating system. BOSS SSDs can be configured in JBOD, RAID 0, or RAID 1. They are generally used in RAID 1 for durable boot storage. Their mounting is secured by blue tool-less retention clips, making installation and replacement easy. The Dell configuration also ensures good airflow and easy access for maintenance.
Here is a close-up of the rear of one of the system's power supplies. The black plastic housing surrounding the connector holds it firmly in place and protects it from accidental disconnection. The power cables, carefully organized and routed, are elegant, ensure optimal airflow, and easy maintenance.
Dell PowerEdge R7715 performance
The performance evaluation of the Dell PowerEdge R7715 will focus primarily on its processor: the AMD EPYC 9665P. This 96-core, 192-thread processor is designed for a single-socket (1P) configuration and offers a maximum boost frequency of 4.5 GHz, an all-core boost frequency of 4.1 GHz, and a base frequency of 2.6 GHz. With 384 MB of L3 cache and a configurable TDP from 320 W to 400 W, the EPYC 9665P is designed to excel in multi-core workloads, delivering an optimal balance between speed, efficiency, and scalability.
Other key system specifications to note include:
- Storage: 32 x 1.6 TB E1.S SSDs
- Memory: 768 GB of RAM
We compared its performance to that of other systems equipped with models from the same EPYC 9005 series. Among these is the EPYC 9665, a high-end model with 192 cores and 384 threads, available in single- and dual-processor configurations (1P/2P). The EPYC 9965, however, has lower clock frequencies (3.7 GHz in maximum boost mode and 3.35 GHz in all-core boost mode), but compensates for this weakness with twice the number of cores as the 9665. The EPYC 9755 is another comparable model, offering 128 cores and 256 threads. It has a slightly lower maximum boost frequency (4.1 GHz) and a lower base frequency (2.7 GHz), but a larger L3 cache (512 MB) and a higher default TDP (500 W). It is also available in single- and dual-processor configurations. Finally, the EPYC 9575F is a high-frequency processor designed for demanding workloads. With 64 cores and 128 threads, it offers a maximum turbo frequency of 5 GHz and an all-core turbo frequency of 4.5 GHz. All processors in the EPYC 9005 series were tested in single-processor configuration for this test.
Additionally, the comparative tests include the Genoa (2 processors/96 cores) and Bergamo (2 processors/128 cores) models, which leverage dual-socket configurations. The comparison with these dual-socket multi-core systems will highlight its scalability against larger configurations in terms of processor count, efficiency, clock frequency, and processing power.
Geekbench 6
The cross-platform Geekbench 6 benchmark measures a system's performance and provides a comparison score. Designed to run on multiple platforms, it provides consistent performance measurements across different devices, including smartphones, tablets, desktops, and servers.
With a single-core score of 2,806, it far surpasses other models, including the EPYC 9965, 9755, and even the high-frequency 9575F. This is likely due to its high maximum boost clock frequency of 4.5 GHz and its all-core boost speed of 4.1 GHz, among the highest in the lineup. The multi-core score of 28,645 is also higher, despite having fewer cores than the 9965 or 9755. Compared to dual-processor systems like Genoa (2P/96c) and Bergamo (2P/128c), it offers better single-core and multi-core performance.
| Geekbench 6 | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| Single Core | 2,806 | 1,453 | 1,641 | 1,865 | 2,048 | 1,723 | 
| Multi-Core | 28,645 | 11,199 | 11,800 | 13,219 | 20,217 | 17,916 | 
Blackmagic RAW Speed Test
We performed the Blackmagic RAW Speed Test to evaluate the PowerEdge R7715's ability to handle Blackmagic RAW decoding tasks using only the processor, without a GPU installed. This test measures performance at different resolutions and compression levels.
The R7715 achieved a score of 171 FPS for 8K content, demonstrating fairly solid CPU performance for high-resolution video processing, even without GPU acceleration.
Maxon Cinebench
Cinebench is a widely used benchmarking tool that measures the performance of processors and graphics cards (CPU) using Maxon Cinema 4D for rendering. It provides a score to compare the performance of different systems and components. We tested four popular versions of Cinebench so you can compare results on the most popular online rankings.
Here, the AMD EPYC 9665P demonstrated its high-performance single-socket processor capabilities. With a multi-core score of 121,254 points in Cinebench R23, it stands out from comparable systems, slightly ahead of the 131,846-point score of the EPYC 9755, but surpassing the Genoa (2P/96c) and Bergamo (2P/128c) systems. However, the most remarkable feature here is the single-core performance, where the 9665P scored 1,845 points, significantly higher than all other processors tested.
In Cinebench 2024, the 9665P continues to excel, scoring 7,501 points in multi-core tests, surpassing the 9965, 9755, and 9575F. Its single-core score of 109 points is also impressive, well ahead of the 9965 (77 points) and 9755 (84 points). The 9665P offers a perfect balance between single-core speed and multi-core power.
| Test | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) |  | 
| Cinebench R23 |  |  |  |  |  |  |  | 
| CPU (multi-core) | 121,254 pts | N/A | 131,846 pts | 111,149 pts | 116,744 pts | 102,125 points |  | 
| CPU (single-core) | 1,845 pts | N/A | 1,400 pts | 1,052 pts | 1,294 pts | 1,089 points |  | 
| Cinebench 2024 |  |  |  |  |  |  |  | 
| CPU (multi-core) | 7,501 pts | 4,845 pts | 5,921 pts | 4,324 | N/A | N/A |  | 
| CPU (single-core) | 109 pts | 77 pts | 84 pts | 103 pts | N/A | N/A |  | 
y-cruncher
y-cruncher is a multithreaded and scalable program capable of calculating Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application among overclockers and hardware enthusiasts.
Despite having fewer cores than the other processors tested, the 9665P showed solid performance on most calculations. Its time of 6.836 seconds for the 1 billion digit calculation is respectable, even though it is slower than the AMD EPYC 9575F (64c), which achieved the best time at 4.476 seconds. As the number of digits increases, the 9665P maintains competitive performance, particularly excelling in the 10 billion digit test, which it completes in 51.851 seconds. This result is comparable to other processors with higher core counts, such as the 9965 (41.750 seconds) and 9755 (41.512 seconds).
However, for higher digit counts (25 and 50 billion in this case), the 9665P takes 140.339 and 303.842 seconds respectively. This likely means that while the 9665P handles heavy computational tasks well, it begins to lose ground as workload complexity increases, particularly compared to processors with higher core counts. Nevertheless, the 9665P remains a serious contender for most workloads and offers excellent performance for its core count.
| y-cruncher Total Computation Time (Lower is Better) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, 128t) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| 1 billion | 6.836 seconds | 7.346 seconds | 7.747 seconds | 5.408 seconds | 4.476 seconds | 8.882 seconds | 9.184 seconds | 
| 2.5 billion | 13.720 seconds | 13.661 seconds | 14.113 seconds | 11.376 seconds | 10.067 seconds | N/A | N/A | 
| 5 billion | 25.795 seconds | 23.211 seconds | 22.820 seconds | 20.177 seconds | 20.030 seconds | N/A | N/A | 
| 10 billion | 51.851 seconds | 41.750 seconds | 41.512 seconds | 40.767 seconds | 41.518 seconds | 51.071 seconds | 55.683 seconds | 
| 25 billion | 140.339 seconds | 115.091 seconds | 98.981 seconds | 103.650 seconds | 104.737 seconds | N/A | N/A | 
| 50 billion | 303.842 seconds | N/A | N/A | N/A | N/A | N/A | N/A | 
| 100 billion | 707.391 seconds | N/A | N/A | N/A | N/A | N/A | N/A | 
Blender OptiX
Blender OptiX is an open-source 3D modeling application. This benchmark was performed using the Blender Benchmark CLI utility. The score is measured in samples per minute, with higher values being better.
The Blender OptiX benchmark results indicate that the EPYC 9665P offers good performance given its core count, but falls behind higher-core models for intensive rendering tasks. In the Monster scene, the 9665P achieves 1,026.50 samples per minute, a solid result, but far from the results of the 9965 (2,558.43) and 9755 (2,606.54). The trend continues in the Junkshop and Classroom scenes, where the 9665P achieves 795.44 and 511.45 samples per minute respectively, well below dual-processor systems like Bergamo (2p/128c) and Genoa (2p/96c).
Although the 9665P offers decent performance for most workloads, it is clear that there are more suitable processors with higher core counts. However, the 9665P holds its own against the EPYC 9575F, particularly in the Junkshop section, where its performance is nearly identical (795.44 vs 802.00).
| Blender 4.0 CPU Samples per Minute (higher is better) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| Monster | 1,026.50 | 2,558.43 | 2,606.54 | 1,196.15 | 1,700.65 | 2,038.71 | 
| Junkshop | 795.44 | 1,866.65 | 1,843.48 | 802.00 | 1,101.84 | 1,382.58 | 
| Classroom | 511.45 | 1,270.17 | 1,251.54 | 637.13 | 869.48 | 1,045.96 | 
Hammer DB TPROC-C
We also tested database performance with Hammer DB on the PowerEdge R7715. This system demonstrated excellent OLTP performance across all databases tested under the HammerDB TPROC-C workload (based on TPC-C, 800 warehouses).
| Database Engine | Transaction Performance (TPM) | 
|---|---|
| MariaDB 11.4.4 | 3,600,000 | 
| MySQL 8.4.4 | 3,300,000 | 
| PostgreSQL 17.2 | 3,100,000 | 
| MariaDB 10.11.12 (MDEV-21923) | 2,950,000 | 
| MariaDB 10.6.22 | 2,850,000 | 
| MySQL 5.7.44 | 2,700,000 | 
MariaDB 11.4.4 showed the best transactional performance. It outperformed older versions of MariaDB, such as 10.6.22 and the 10.11.12 version optimized for specific needs (MDEV-21923). MySQL 8.4.4 also showed excellent performance, closely trailing MariaDB 11.4.4. PostgreSQL 17.2 achieved competitive results but remained slightly behind MariaDB and the new MySQL version. MySQL 5.7.44 was the weakest database among those tested.
7-Zip Compression Benchmark
The built-in memory test in the 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, indicating how well the system can handle data-intensive operations. We run this test with a 128 MB dictionary size when possible.
Although it has fewer cores than some other processors in the EPYC 9005 series, it achieved a total score of 378.469 GIPS. While this score is undoubtedly respectable, it is significantly lower than that of the EPYC 9755 (443.029 GIPS) and EPYC 9575F (394.900 GIPS). Interestingly, the 9965 (266.740 GIPS) lags behind in this benchmark, suggesting that a higher core count does not always translate to better compression performance.
In decompression tasks, the 9665P maintains its current level with 395.502 GIPS but remains behind the 9755 (487.263 GIPS) and 9575F (425.580 GIPS). Its single-socket architecture and high clock frequencies allow it to remain competitive, but higher-throughput models have the advantage in raw throughput.
| 7-Zip Compression Benchmark (Higher is Better) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, SMT disabled) |  | 
| Compression |  |  |  |  |  | 
| Current CPU Usage | 5881 % | 4302 % | 5233 % | 4406 % |  | 
| Current Rating/Usage | 6.112 GIPS | 5.830 GIPS | 7.597 GIPS | 7.975 GIPS |  | 
| Current | 359.471 GIPS | 250.827 GIPS | 397.536 GIPS | 351.358 GIPS |  | 
| Resulting CPU Usage | 5875 % | 4041 % | 5306 % | 4555 % |  | 
| Resulting Rating/Usage | 6.140 GIPS | 5.804 GIPS | 7.720 GIPS | 8.070 GIPS |  | 
| Resulting Rating | 360.695 GIPS | 234.317 GIPS | 409.652 GIPS | 367.358 GIPS |  | 
| Decompression |  |  |  |  |  | 
| Current CPU Usage | 6168 % | 4322 % | 6041 % | 5017 % |  | 
| Current Rating/Usage | 6.412 GIPS | 7.078 GIPS | 8.065 GIPS | 8.483 GIPS |  | 
| Current | 395.502 GIPS | 305.909 GIPS | 487.263 GIPS | 425.580 GIPS |  | 
| Resulting CPU Usage | 6159 % | 4556 % | 5921 % | 4940 % |  | 
| Resulting Rating/Usage | 6.434 GIPS | 6.577 GIPS | 8.045 GIPS | 8.569 GIPS |  | 
| Resulting Rating | 396.243 GIPS | 299.163 GIPS | 476.405 GIPS | 422.441 GIPS |  | 
| Total Rating |  |  |  |  |  | 
| Total CPU Usage | 6017 % | 4298 % | 5613 % | 4747 % |  | 
| Total Rating/Usage | 6.287 GIPS | 6.190 GIPS | 7.883 GIPS | 8.319 GIPS |  | 
| Total Rating | 378.469 GIPS | 266.740 GIPS | 443.029 GIPS | 394.900 GIPS |  | 
Conclusion
The Dell PowerEdge R7715 offers an impressive balance of performance, scalability, and efficiency for modern enterprise workloads. Compatible with AMD EPYC 9005 series processors, offering up to 160 cores and 24 DDR5 DIMM slots for up to 6 TB of memory, the R7715 is perfectly equipped to handle data-intensive applications in virtualization, analytics, and software-defined storage environments.
With high clock frequencies and a high-performance architecture, the R7715 excels in single-core performance and offers competitive multi-core performance thanks to its single-socket design. Although it does not match dual-socket systems in raw parallel computing, it comes surprisingly close, offering a more energy-efficient and cost-effective alternative for many real-world workloads.
This configuration is not ideal for all use cases, particularly when maximum core density or GPU acceleration is essential. However, Dell is expected to roll out GPU support for the R7715 later this year. Additionally, businesses needing larger capacity drives for specific workloads might consider more extensive storage configurations.
Ultimately, the R7715 is an ideal platform for IT environments that prioritize high throughput, fast memory, and Gen5 I/O flexibility, without the complexity or cost of dual-socket deployments. The R7715 stands out as a wise option for businesses looking to optimize their efficiency without sacrificing capabilities.
Product configuration page
