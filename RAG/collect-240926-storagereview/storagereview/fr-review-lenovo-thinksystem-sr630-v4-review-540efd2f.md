---
id: collect-240926-storagereview/storagereview/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f
title: "fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Microsoft", "Samsung"]
dates: []
keywords: ["benchmark", "compute", "cost", "decode", "energy", "gpu", "gpus", "intel", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f.md
source_anchor: ""
source_lines: [1, 158]
sha256: 1ca117cb05fa3ec759948865df7dc00f99cf4f70b7da6d111765ae69ab58dd92
---

# fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f

<!-- source: https://www.storagereview.com/fr/review/lenovo-thinksystem-sr630-v4-review -->

The Lenovo ThinkSystem SR630 V4 is a flexible and powerful 1-socket 2U rack server, designed to meet the needs of sectors such as cloud services and telecommunications. Whether optimizing scalable workloads or future-proofing your data center, the SR630 V4 offers significant upgrades over its predecessor, the SR630 V3. In this review, we examined what's new and how Lenovo has refined this enterprise server to meet the challenges of modern IT environments.
Differences Between Lenovo SR630 V4 and SR630 V3
Processors
One of the major improvements in the SR630 V4 lies in its processing capabilities. While the SR630 V3 relied on 4th and 5th generation Intel Xeon Scalable processors, featuring up to 64 cores and Hyper-Threading technology, the SR630 V4 integrates Intel Xeon 6700 series processors, with up to 144 high-efficiency cores (E-cores). This evolution doubles the number of cores while prioritizing efficiency, despite the absence of Hyper-Threading technology. Additionally, the V4 version plans support for Intel Xeon P-cores, which could significantly improve performance for certain workloads. The higher core density of the V4 version allows businesses to consolidate more applications on the same number of servers, thereby reducing operating costs and physical server requirements.
Here is an overview of all 6300 series processors supported by the SR4 V6700:
| CPU Model | Cores / Threads | Core Speed (Base/Max Turbo) | L3 Cache | Memory Chan | Max Memory Speed | UPI 2.0 Links and Speed | PCIe Lanes | TDP | 
| 6710E | 64/64 | 2.4 / 3.2 GHz | 94 MB | 8 | 5600 MHz | 4 / 16 GT/s | 88 | 205W | 
| 6731E | 96/96 | 2.2 / 3.1 GHz | 96 MB | 8 | 5600 MHz | None‡ | 88 | 250W | 
| 6740E | 96/96 | 2.4 / 3.2 GHz | 96 MB | 8 | 6400 MHz | 4 / 20 GT/s | 88 | 250W | 
| 6746E | 112/112 | 2.2 / 2.7 GHz | 96 MB | 8 | 5600 MHz | 4 / 16 GT/s | 88 | 250W | 
| 6756E | 128/128 | 1.8 / 2.6 GHz | 96 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 225W | 
| 6766E | 144/144 | 1.9 / 2.7 GHz | 108 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 250W | 
| 6780E | 144/144 | 2.2 / 3 GHz | 108 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 330W | 
Memory
Regarding memory, the SR630 V4 improves upon the V3's DDR5 memory, operating up to 5600 MHz by supporting DDR5 memory speeds up to 6400 MHz for E-cores. Both systems have 32 DIMMs (16 per processor) and two DIMMs per channel across eight channels per CPU. Nevertheless, the V4 introduces future-proofing with planned support for advanced memory technologies such as Compute Express Link (CXL) and MCRDIMM for P-cores.
While the SR630 V3 could support up to 8 TB of memory, the V4 focuses on a more targeted capacity of 2 TB for E-cores, thereby optimizing costs and performance for specific workloads.
Storage
The storage capabilities of the Lenovo SR630 V4 demonstrate an evolution toward high-performance NVMe drives while reducing reliance on traditional SAS/SATA options. The SR630 V3 offers flexibility with 3.5-inch SAS/SATA drive bays and up to 16 integrated NVMe ports. However, the SR630 V4 supports up to 12 NVMe drives in front and rear configurations, adding future options for E3.S drive formats, which promise superior capacity and density.
The V4 eliminates support for 3.5-inch drives but introduces hot-swap M.2 options for operating system boot, thereby improving performance and serviceability. By focusing on NVMe and eliminating the need for additional adapters, the V4 maximizes I/O bandwidth and system efficiency.
Networking
For networking, the SR630 V4 builds on the V3's single OCP slot by offering two OCP 3.0 slots, both supporting PCIe Gen 5 x16. This upgrade doubles networking flexibility, enabling improved connectivity and throughput with dual-port 200 GbE network adapters or other advanced networking solutions.
The increased PCIe bandwidth (32 GT/s in Gen 5 versus 16 GT/s in Gen 4) means the V4 can better handle the demanding workloads of data centers and the cloud, making it a more versatile option for modern networking needs.
Tuning Engine
Finally, the SR630 V4 offers improved power options, moving from the SR630 V3's Platinum/Titanium AC 750 W to 1800 W options to a broader range of 800 W to 2000 W, including models compliant with the ErP Lot 9 standard for optimal energy efficiency. The V4 also retains support for -48 V DC power compatible with telecommunications operators while introducing new HVDC 1300 W options for specific regional requirements. These improvements make the V4 more adaptable to various power environments, ensuring it meets the energy requirements of more complex and higher-performance configurations.
Overall, the SR630 V4 represents a solid advance on paper to better meet the need for more performance, flexibility, and efficiency in modern IT environments.
Lenovo ThinkSystem SR630 V4 Specifications
| Lenovo ThinkSystem SR630 V4 Specifications |  | 
| Form Factor | 1U rack | 
| Processor | One or two Intel Xeon 6700E series processors (up to 144 cores, 2.4 GHz and 330 W TDP). Support for Intel Xeon 6700P series processors planned for Q1 2025. | 
| Memory | 32 DIMM slots (16 per processor), supports TruDDR5 RDIMMs up to 6400 MHz (1DPC) or 5200 MHz (2DPC). CXL memory is planned for the Intel Xeon 6700P series in Q1 2025. | 
| Maximum Memory | Up to 2 TB using 32 x 64 GB RDIMMs | 
| Drive Bays |  | 
| Maximum Internal Storage | 184.3 TB with 12 x 15.36 TB 2.5-inch NVMe SSDs | 
| Storage Controller | Up to 16 integrated NVMe ports with RAID support (Intel VROC). Planned support for 12 Gb SAS/SATA RAID and non-RAID adapters. | 
| Network Interfaces | Two OCP 3.0 SFF slots with PCIe 5.0 host interface (x8 or x16), supporting up to 100 GbE network adapters. | 
| PCI Expansion Slots |  | 
| GPU Support | Planned support for up to 3 single-width GPUs | 
| Ports | Front:  Rear:  | 
| Cooling | Up to 8 hot-swap fans (N+1 redundancy), with an additional fan integrated into each power supply. | 
| Power Supply | Up to two redundant hot-swap AC power supplies (800 W, 1300 W, 2000 W). 80 PLUS Platinum and Titanium certifications. | 
| Video | Integrated graphics card with two video ports (rear VGA and optional Mini DisplayPort), supporting resolutions up to 1920 × 1200 at 60 Hz. | 
| Hot-Swap Parts | Drives, power supplies, and fans | 
| Systems Management |  | 
| Security Features | Chassis intrusion switch, power-on and administrator passwords, TPM 2.0, and optional lockable front security bezel. | 
| Supported Operating Systems | Microsoft Windows Server, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, Ubuntu Server. | 
| Warranty | Three years or one year (depending on model) with optional service upgrades for faster response times and extended coverage. | 
| Dimensions | Width: 440 mm (17.3 in), Height: 43 mm (1.7 in), Depth: 788 mm (31 in). | 
| Weight | Maximum weight: 20.2 kg (44.5 lb) | 
Lenovo ThinkSystem SR630 V4 Design and Build
The Lenovo ThinkSystem SR630 V4 retains the compact 1U form factor that is standard for many enterprise rack servers. Its design focuses on a combination of functionality, accessibility, and flexibility. We appreciated its simple yet effective layout, which maximizes airflow, modularity, and ease of use for IT administrators.
Let's move on to the details.
Front Panel
The front panel supports up to 10 hot-swap 2.5-inch drive bays and offers flexibility for various storage configurations, including SAS, SATA, NVMe, or AnyBay drives. This allows businesses to customize storage based on their workloads, whether they prioritize speed, capacity, or cost-effectiveness.
The front panel can also be configured with an optional Mini DisplayPort video port, which can be used for quick local monitoring and diagnostics without having to access the rear of the rack. Up to two optional USB 3.0 ports are also available; one is explicitly designated for connection to the Lenovo XClarity Controller (XCC). This connectivity simplifies management tasks, such as downloading firmware updates or running diagnostics directly from a USB device.
Lenovo has integrated an external diagnostic port on our system, which can be very useful for on-site IT staff who need to troubleshoot hardware issues (such as system status and failures). This can speed up problem resolution and minimize downtime. It also has a removable information tag for quick access to essential system details, such as serial numbers, configurations, and network information.
The indicators and controls on the front control panel provide the usual at-a-glance information on system status and activity, including power and reset buttons and LEDs for drive status and health.
Rear Panel
The rear panel includes hot-swap power supplies (from 800 W to 2000 W) located on each side of the system, which provide redundancy and can be replaced without shutting down the system. The Dual OCP 3.0 slots support PCIe Gen 5 x16 for advanced networking, allowing for high-bandwidth adapters such as dual-port 200 GbE cards. The panel also includes a video port, two USB 3.0 ports, and a dedicated XClarity Controller (XCC) management port for local and remote system management. LEDs provide quick visual updates on system status.
The rear panel includes a mix of low-profile and full-height PCIe slots for expansion, allowing users to add GPUs, storage controllers, or other adapters. Storage options include hot-swap 2.5-inch drives and hot-swap M.2 drives, offering flexible configurations for boot devices or additional storage. The rear panel is also available in four air-cooled configurations and two water-cooled configurations.
Internal
When you open the Lenovo ThinkSystem SR630 V4, you will see the two processors surrounded by their respective DIMM slots. This chassis offers 16 DIMMs per processor, for a total of 32, allowing up to 2 TB of RAM to be installed.
From the front, you will notice up to eight hot-swap fans lined up, channeling airflow over critical components and keeping everything cool under pressure. Our system is equipped with directly connected NVMe SSDs, which are cabled directly to the motherboard.
This direct-attach NVMe offers the highest available NVMe performance, although for RAID, users must choose between software or hardware options like Graid.
At the rear, the PCIe slots are ready to accommodate GPUs or other expansion cards, while the OCP slots add an extra layer of versatility for specialized networking needs. The hot-swap power supplies are easy to access and replace, minimizing downtime during maintenance.
XClarity Controller 3
The Lenovo ThinkSystem SR630 V4 is equipped with the XClarity Controller 3 (XCC3) for remote and lifecycle management. It offers out-of-band management capabilities via a dedicated LAN port, allowing users to configure and deploy new hardware, interact with the system if primary networks are down, perform firmware management activities, and carry out countless other tasks.
Users can get a quick overview of all major components and warnings from the main home screen. System status, active events, and power are the key areas here.
You can see how the platform handles updates by exploring several areas, such as firmware update. You import a firmware package into local storage inside the XClarity controller and click Update System to start the process of updating the downloaded firmware payload.
Remote control is another common function that Lenovo XClarity handles well via an HTML5 web interface. This allows a wide range of client systems to manage the platform, including mobile platforms. Although an iPad is not the primary support mechanism, you sometimes have to use what's nearby.
Lenovo ThinkSystem SR630 V4 Performance
This section examines the performance test results of y-cruncher, Cinebench, Blackmagic, 7-Zip, and Geekbench. We compared the dual-processor Lenovo ThinkSystem SR630 V4 with the recently tested single-processor Supermicro Hyper 1U SYS-112H-TN. Both systems are equipped with the Intel Xeon 6780E processor, allowing us to see how the 6780E scales from single- to dual-processor configurations.
In addition to the Supermicro, we added an older Intel Ice Lake server, provided at the launch of the first Ice Lake Xeon 8380 processors. This shows how the E-core models position themselves as a cost-effective upgrade for legacy platforms that prioritize efficiency over raw processing power. This comparison between the dual-processor SR630 V4 platform and Intel Ice Lake and the single-processor Supermicro Hyper 1U illustrates the performance difference.
Here are the configurations for each system.
Lenovo ThinkSystem SR630 V4 Configuration
- CPU: 2 x Intel Xeon 6780E (144 cores)
- RAM: 512GB DDR5
- SSD Samsung MZWL6960HFJA-00AW7
- Operating System: Windows Server 2025
Supermicro Hyper 1U SYS-112H-TN Configuration
- CPU: Intel Xeon 6780E (144 cores)
- RAM: 512GB DDR5
- SSD Micron 7450 NVMe Data Center SSD
- Operating System: Windows Server 2022
Intel Ice Lake Server
- CPU: 2 x Intel Xeon 8380 (80 cores)
- RAM: 512GB DDR5
- SSD
- Operating System: Windows Server 2025
OptiX Blender
First, we will move on to the Blender test, an open-source 3D modeling application. This test was run using the Blender Benchmark utility. The score is expressed in samples per minute, with the higher being the better.
The Blender OptiX tests reveal interesting information about the systems tested. The Lenovo ThinkSystem SR630 V4 excelled with Blender 4.2.0, delivering 1,432 samples per minute in the Monster scene, while the Intel Ice Lake Server reached 569 and the Supermicro Hyper 1U 112H-TN (running Blender 4.0) reached 781. Lenovo reported 914 samples in the Junkshop scene, with 403 for the Intel Ice Lake Server and 514 for Supermicro. The Classroom scene showed Lenovo had 657 samples, Ice Lake with 280, and Supermicro with 371.
| Blender Processor | Supermicro Hyper 1U 112H-TN (1x Xeon 6780E, 512 GB DDR5) Blender 4.0 | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) Blender 4.2 | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) Blender 4.2 | 
| Monster | 781.42 | 1432.09 | 569.10 | 
| Junkshop | 514.658 | 914.75 | 403.96 | 
| Classroom | 370.52 | 656.68 | 280.86 | 
Geekbench 6
Geekbench 6 is a cross-platform performance evaluation tool that measures the overall performance of a system. The Geekbench browser allows you to compare any system using this tool.
Lenovo's single-core test scored 1,173, while the dual-processor Supermicro recorded 1,154, highlighting its potential for tasks that depend on the efficiency of each core. In the multi-core test, Supermicro recorded 15,167 and Lenovo 13,868. Not all applications scaled well with the increase in core count. Geekbench struggled to scale from 144 to 288 cores on the dual-socket Lenovo SR630 V4 platform. The older Ice Lake processors showed higher single-core and multi-core scores, reaching 1,668 and 17,409 respectively.
| Geekbench 6 (Higher is Better) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| Single-Core CPU | 1,154 | 1,173 | 1,668 | 
| Multi-Core CPU | 15,167 | 13,868 | 17,409 | 
Cinebench R23
The Cinebench R23 benchmarking tool evaluates a system's CPU performance by rendering a complex 3D scene using the Cinema 4D engine. It measures single-core and multi-core performance, providing a comprehensive view of the CPU's capabilities in handling 3D rendering tasks.
The table below shows that the Supermicro and Lenovo systems recorded good results. The Lenovo ThinkSystem SR630 V4 outperformed in multi-core and single-core performance, with 99,266 and 894 points respectively. The Supermicro Hyper 1U 112H-TN recorded 92,516 and 888 points. The additional processor had some benefits in this benchmark, but the numbers did not double, going from one to two processors. The Ice Lake processors recorded 74,020, with limited scaling compared to Lenovo's dual Xeon 6780E.
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| Multi-Core CPU | 92,516 | 99,266 pts | 74,020 pts | 
| Single-Core CPU | 888 pts | 894 pts | 1,059 pts | 
| MP Ratio | 104.20 x | 111.00 x | 69.87 x | 
Cinebench 2024
Cinebench 2024 extends the benchmarking capabilities of R23 by adding a GPU performance evaluation. It continues to test CPU performance but also includes tests that measure the GPU's ability to handle rendering tasks.
The results of the 2024 version of Cinebench told a similar story. Here, the Lenovo ThinkSystem SR630 V4 with its two 6780E processors reflected the advantage over the single-socket Supermicro Hyper 1U 112H-TN in multi-core CPU performance with a score of 2,884 points. The Supermicro reported 2,565 points. The Intel Ice Lake 8380 processors demonstrated their power with a multi-core score of 4,131. For single-core tests, the Intel Ice Lake 8380 scored 61 points.
| Cinebench 2024 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| Multi-Core CPU | 2,565 pts | 2,884 pts | 4,131 pts | 
| Single-Core CPU | 53 pts | 53 pts | 61 pts | 
| MP Ratio | 48.38 x | 54.43 x | 68.22 x | 
y-cruncher
y-cruncher is a popular benchmarking and stress testing application launched in 2009. This test is multithreaded and scalable, calculating Pi and other constants to trillions of digits. Faster is better in this test. This software has been fantastic for testing high-core-count platforms and showing the computational advantages between single- or dual-socket platforms.
In the y-cruncher performance tests, the dual-socket ThinkSystem SR630 V4 took 5.997 seconds to calculate Pi to 1 billion digits. The Intel Xeon 6780E processor took 8.757 seconds. To calculate 50 billion digits, a single processor needed 674.299 seconds, compared to 476.826 seconds for dual processors. While not all workloads respond well to the high core count of the new e-core processors, y-cruncher had no trouble exploiting them. The older Intel Ice Lake Server completed the Pi calculation to 1 billion digits in 7.074 seconds in the first 1 billion digit test, while reaching 617.828 seconds for 50 billion digits.
| y-cruncher (0.8.5.9) (Lower is Better) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| 1 billion | 8.757 seconds | 5.997 seconds | 7.074 seconds | 
| 2.5 billion | 24.928 seconds | 17.573 seconds | 19.203 seconds | 
| 5 billion | 53.489 seconds | 37.793 seconds | 42.300 seconds | 
| 10 billion | 113.727 seconds | 81.046 seconds | 93.886 seconds | 
| 25 billion | 308.218 seconds | 220.025 seconds | 272.679 seconds | 
| 50 billion | 674.299 seconds | 476.826 seconds | 617.828 seconds | 
Blackmagic RAW Speed Test
The Blackmagic RAW Speed Test is a benchmarking tool designed to measure a system's ability to handle video playback and editing using the Blackmagic RAW codec. It evaluates a system's ability to decode and play high-resolution video files, providing frame rates for CPU-based and GPU-based processing.
The Lenovo ThinkSystem SR630 V4 achieved slightly higher results than the Supermicro and Ice Lake systems, with a score of 120 FPS in 8K CPU, making it an excellent choice for video playback and editing tasks. The Supermicro Hyper 1U 112H-TN achieved 116 FPS in 8K CPU. Although the dual-socket Lenovo performed better, it wasn't by much, given the dual-socket configuration with 116 FPS (8K CPU) and 0 FPS (8K GPU). The Intel Ice Lake server was on par with the Lenovo and Supermicro with 116 FPS in the 8K CPU benchmark. However, it did not record a GPU score due to the absence of a dedicated GPU.
7-Zip
The built-in memory benchmark of the popular 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, indicating how well the system can handle data-intensive operations.
Regarding compression tasks, the Supermicro system achieved slightly higher results than the SR630 in terms of CPU usage and resulting ratings, with a total compression rating of 245.823 GIPS and Lenovo's at 224.313 GIPS. This suggests a slight advantage when handling heavily threaded compression workloads. Lenovo's decompression tasks showed a higher resulting rating of 288.457 GIPS, while Supermicro's rating indicated 269.373 GIPS. This translates to better performance for workloads that require reading and extracting data. The Intel Ice Lake server demonstrated balanced performance, with a compression rating of 235.437 GIPS and a decompression rating of 253.692 GIPS.
The overall performance of the two systems is nearly identical, with total ratings of 257.598 GIPS for the single-socket Supermicro and 256.385 GIPS for the dual-socket Lenovo. The older Xeon Ice Lake reached 244.565 GIPS. All three systems are very high-performing.
| 7-Zip Compression | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| Compression |  |  |  | 
| Current CPU Usage | 5287 % | 5064 % | 5835 % | 
| Current Rating/Usage | 4.647 GIPS | 4.341 GIPS | 4.030 GIPS | 
| Current | 245.699 GIPS | 219.840 GIPS | 235.143 GIPS | 
| Resulting CPU Usage | 5296 % | 5156 % | 5839 | 
| Resulting Rating/Usage | 4.642 GIPS | 4.350 GIPS | 4.032 GIPS | 
| Resulting Rating | 245.823 GIPS | 224.313 GIPS | 235.437 GIPS | 
| Decompression |  |  |  | 
| Current CPU Usage | 6236 % | 6184 % | 6230 % | 
| Current Rating/Usage | 4.261 GIPS | 4.688 GIPS | 4.050 GIPS | 
| Current | 265.709 GIPS | 289.879 GIPS | 252.326 GIPS | 
| Resulting CPU Usage | 6236 % | 6205 % | 6245 % | 
| Resulting Rating/Usage | 4.341 GIPS | 4.649 GIPS | 4.062 GIPS | 
| Resulting Rating | 269.373 GIPS | 288.457 GIPS | 253.692 GIPS | 
| Total Rating |  |  |  | 
| Total CPU Usage | 5751 % | 5681 % | 6042 % | 
| Total Rating/Usage | 4.491 GIPS | 4.500 GIPS | 4.047 GIPS | 
| Total Rating | 257.598 GIPS | 256.385 GIPS | 244.565 GIPS | 
Conclusion
The Lenovo ThinkSystem SR630 V4 is a versatile 1U rack server that, while not representing a significant upgrade over its predecessor, is nonetheless a reliable advance in the evolution of Lenovo's mainstream enterprise systems. With support for Intel Xeon 6700E series processors, it doubles core density compared to its predecessor, offering better scalability for demanding workloads. Improved DDR5 memory speeds up to 6400 MHz and planned support for emerging technologies such as Compute Express Link (CXL) and MCRDIMM show that Lenovo intends to keep this server line equipped for future demands.
The SR630 V4's evolution toward NVMe storage and flexible drive configurations also means it prioritizes performance and scalability in data-intensive workloads. Additionally, the two OCP 3.0 slots supporting PCIe 5.0 enable advanced networking options, including hot-swap components and improved cooling systems, simplifying maintenance and operational efficiency.
Regarding its performance in our benchmark tests, the SR630 V4 performed well, excelling in multithreaded workloads such as Cinebench R23 and Y-Cruncher. The way the new Intel Sierra Forest E-core Xeon processors fit into this space is intended for businesses looking to upgrade their existing platforms for greater efficiency. Not all workloads require maximum performance, but they could benefit from new improvements such as density and reduced power consumption.
That said, even though the efficiency of E-core processors does not fully meet the requirements of workloads demanding high single-thread performance, Lenovo has planned support for P-core processors. Additional drive configurations are also available, supporting a wide range of enterprise applications.
