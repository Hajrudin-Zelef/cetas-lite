---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf
title: "fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia", "Samsung"]
dates: []
keywords: ["amd", "benchmark", "compute", "cost", "energy", "exploit", "gpu", "gpus", "inference", "memory", "nvidia", "parameters"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf.md
source_anchor: ""
source_lines: [1, 181]
sha256: a019e7f461c96b2649cbf2dfb7f23146ea2983bad0a66731e14d2d418b165301
---

# fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf

<!-- source: https://www.storagereview.com/fr/review/supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004 -->

The Supermicro AS-1115SV-WTNRT server is a versatile and robust 1U rackable server, designed to meet the demands of the most resource-intensive applications such as virtualization, database management, and edge computing. Suited for environments requiring high throughput and high reliability, this server offers a balanced architecture integrating the AMD EPYC 8004 series processor, thus providing a scalable solution to modern computing challenges.
Supermicro AS-1115SV-WTNRT Main Features and Hardware Specifications
At the heart of the AS-1115SV-WTNRT is a single-socket configuration using the AMD EPYC 8004 series processor, capable of handling up to 64 cores and 128 threads, which allows for significant parallel processing power. Designed for dense data centers and service providers, the "Siena" series processors offer an optimal combination of high core count and energy efficiency at a competitive price, operating within a modest power envelope starting at just 70 W and reaching up to 200 W.
The server supports up to 576 GB of DDR5 4800 2.5 MHz memory across six DIMM slots, ensuring high-speed data transfer and efficient management of large datasets. Storage versatility is another important attribute, with ten hot-swappable XNUMX″ drive bays that support a combination of NVMe, SAS, and SATA drives.
The Supermicro AS-1115SV-WTNRT is designed with a compact 1U form factor, measuring 437 mm x 43 mm x 597 mm, optimizing space without compromising internal hardware capacity. Expansion capabilities include three PCIe 5.0 x16 slots, two of which are full-height and full-length slots, providing ample space for larger cards, while the third is a low-profile slot, suitable for smaller expansion cards.
The server features two redundant 860 W power supplies for continuous power and system reliability.
Under the hood, the server's internal layout optimizes airflow and cooling efficiency, ensuring consistent performance under various operating conditions, including the six internal counter-rotating fans (FAN-0163L4).
The smaller H13SVW-NT motherboard is in the foreground, housing a single AMD EPYC 8004 series processor in an SP6 socket, surrounded by its six DDR5 DIMM slots.
Supermicro AS-1115SV-WTNRT System Management
Supermicro has equipped the AS-1115SV-WTNRT with a robust suite of management tools, including SuperCloud Composer and Supermicro Server Manager (SSM), which facilitate streamlined system management and monitoring. The BIOS is equipped with the contemporary UEFI 2.8 specification, offering enhanced security, faster boot times, and support for larger disk partitions.
The AS-1115SV-WTNRT also features a comprehensive baseboard management controller (BMC) system that provides administrators with control and monitoring capabilities. This BMC interface (accessible via a dedicated IP address) provides critical information about system status, allows for comprehensive configuration changes, and enables remote control actions without physical access to the hardware. It is an essential tool for efficiently managing server operations, ensuring availability, and quickly resolving issues.
The dashboard offers a comprehensive overview of server status and essential metrics. It presents system status indicators, IP addresses, and firmware versions, providing a snapshot of critical operational data. Power consumption is displayed graphically over time, detailing minimum, average, and maximum usage. This visualization allows administrators to assess at a glance the server's energy efficiency and operational costs. A remote console overview is also available, facilitating quick access to the server's current operational status, which is invaluable for rapid troubleshooting and continuous monitoring.
The CPU section offers detailed information about the server's processor specifics (in our case, the AMD EPYC 8534P 64-core processor), such as speed, thermal design power (TDP), number of cores, number of threads, and manufacturer. This section is crucial for verifying the processor's operational parameters and ensuring it operates within specified limits.
In the Firmware Update tab, users can manage and update the firmware of various server components, such as BMC, BIOS, and CPLD. The interface allows users to choose the firmware type, preserve existing configurations, and ensure that all critical updates are applied without affecting the server's operational settings. This function is essential for maintaining server security, as firmware updates often contain patches for vulnerabilities and performance improvements.
The Memory tab provides detailed information about all installed memory modules, displaying each module's health status, type, error correction code capacity, operating speed, size, and serial number. The Memory tab allows users to easily track and manage the server's physical memory, which is essential for diagnosing memory-related issues or planning upgrades.
The Power tab displays detailed power consumption statistics over different periods, including the last hour, day, and week. It shows historical trends and peak values, useful for capacity planning and operational analysis. Users and administrators can also use this data to optimize power settings, schedule low-power modes during off-peak hours, and make informed decisions regarding energy usage and efficiency.
Supermicro AS-1115SV-WTNRT Specifications
| Processor |  | 
| Processor | Single Processor(s) – AMD EPYC 8004 Series Processor | 
| Number of Cores | Up to 64C/128T | 
| Note | Supports processors with TDP up to 225 W (air-cooled) | 
| GPU |  | 
| Maximum Number of GPUs | Up to 1 double-width GPU or 2 single-width GPUs | 
| System Memory |  | 
| Memory | Number of Slots: 6 DIMM slots – Maximum Memory (1DPC): up to 576 GB 4800 5 MT/s ECC DDRXNUMX RDIMM | 
| Onboard Devices |  | 
| Chipset | System on Chip | 
| Network Connectivity | 2 RJ45 10GBASE-T with Broadcom BCM57416 | 
| Input / Output |  | 
| LAN | 1 dedicated IPMI LAN port(s) RJ45 1 GbE – 2 RJ45 10 GBASE-T LAN port(s) | 
| USB | 2 ports (front) – 4 ports (rear) | 
| Video | 1 VGA port(s) | 
| Serial Port | 1 COM port(s) (rear) | 
| TPM | 1 onboard TPM/80 port | 
| System BIOS |  | 
| BIOS Type | AMI 32 MB SPI Flash EEPROM | 
| BIOS Features | ACPI 6.4, SMBIOS 3.5 or later, UEFI 2.8, Plug and Play (PnP), USB keyboard support | 
| Management |  | 
| Software |  | 
| Power Configurations | Redundant 1U 800/860 W Titanium AC Power Supply | 
| Security |  | 
| Hardware | Trusted Platform Module (TPM) 2.0, Silicon Root of Trust (RoT) – NIST 800-193 Compliant | 
| Feature |  | 
| PC Health Monitoring |  | 
| FAN |  | 
| Temperature |  | 
| Chassis |  | 
| Form Factor | 1U | 
| Model | CSE-116BTS-R000WNP | 
| Dimensions and Weight |  | 
| Height | 1.7 mm (43 in) | 
| Width | 17.2 mm (437 in) | 
| Depth | 23.5 mm (597 in) | 
| Package | 7.75 "(H) x 23.5" (L) x 31.5 "(D) | 
| Weight |  | 
| Available Color | Black | 
| Front Panel |  | 
| LED |  | 
| Buttons |  | 
| Expansion Slots |  | 
| PCI-Express (PCIe) Configuration |  | 
| M.2 | 2 M.2 PCIe 3.0 x4 NVMe slots (M-key) | 
| Drive/Storage Bays |  | 
| Drive Bay Configuration |  | 
| M.2 | 2 M.2 PCIe 3.0 x4 NVMe slots (M-key 2280/22110) View M.2 options | 
| System Cooling |  | 
| Fans | 6 robust 40 x 40 x 56 mm counter-rotating fans. | 
| Power Supply |  | 
|  |  | 
Supermicro 1115SV-WTNRT Performance
Our Supermicro Storage A+ 1115SV-WTNRT review unit is configured with the following:
- Processor: AMD EPYC 8534P 64 cores
- RAM: 194GB DDR5
- GPU: NVIDIA A2 (15,360 6 MB GDDRXNUMX)
- SSD: Ultrastar SN655 NVMe Data Center SSD (15.36 TB)
- 2x Mellanox ConnectX-6 DX 100G
- 10 64-bit Windows
We recently tested the Graid SupremeRAID SR-1001 RAID solution with the Supermicro AS-1115SV-WTNRT SSD. To discover their respective performance, see our full test.
In this performance analysis, however, we will compare it to the Supermicro Storage A+ ASG-1115S-NE316R. It is important to note that this will not be a direct evaluation. Rather, we aim to provide a perspective on the scale of the two AMD processors.
The A+ ASG-1115S-NE316R review unit is configured with the following:
- Processor: EPYC 84 at 9634 cores
- RAM: 384 GB DDR5 (12 x 32 GB DIMMs)
- SSD: 1 TB Samsung PM9A3
- 2x Mellanox ConnectX-6 DX 100G
- Windows Server 2022
Blender OptiX 4.0
The first is Blender OptiX, an open-source 3D modeling application. This benchmark was run using the Blender Benchmark CLI utility. The score is expressed in samples per minute, with the higher being the better.
The Supermicro system showed variable performance in 3D rendering tasks across different scenes when tested with and without GPU acceleration. In all scenarios (Monster, Junkshop, and Classroom), the CPU-only configuration provided higher samples per minute than the GPU-assisted configuration.
Specifically, the Monster scene reached 492.88 samples/min without GPU, surpassing the score of 427.25 samples/min with GPU. Similarly, in the Junkshop scene, the CPU-only test reached 349.35 samples/min, far exceeding the GPU's 266.65 samples/min. The Classroom scene also recorded better performance with the CPU alone, with a score of 255.98 samples/min versus 238.33 samples/min for the GPU.
| Blender 4.0 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5, GPU) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5, CPU) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| Monster | 427.25 | 492.88 | 673.21 | 
| Junkshop | 266.65 | 349.35 | 475.17 | 
| Classroom | 238.33 | 255.98 | 342.06 | 
Blender OptiX 4.1
Blender OptiX 4.1 brings new features, such as GPU-accelerated denoising, streamlining the rendering process and reducing the time needed for denoising tasks. Despite these advances, the overall performance improvements in benchmark scores compared to version 4.0 are minimal, indicating only slight improvements in efficiency (as you will notice in the results below).
The results show that the CPU-only configuration consistently outperformed the GPU-enhanced configuration in all tests, as in version 4.0. Specifically, CPU-only mode reached 493.167 samples/min in Monster, 356.36 samples/min in Junkshop, and 250.87 samples/min in Classroom, compared to GPU-assisted scores of 428.90, 269.39, and 238.58 samples/min respectively.
| Blender 4.0 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5, GPU) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5, CPU) | 
| Monster | 428.90 | 493.167 | 
| Junkshop | 269.39 | 356.36 | 
| Classroom | 238.58 | 250.87 | 
Blackmagic RAW Speed Test
We performed the Blackmagic RAW speed test to extend video playback. This is more of a hybrid test including CPU and GPU performance for actual RAW decoding. Here, we only tested the CPU, which was able to reach 117 FPS.
| Blackmagic RAW Speed Test (Higher is Better) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| CPU 8K | 117 FPS | 131 FPS | 
Blackmagic Disk Speed Test
The Blackmagic Disk Speed Test is another test for which we only have results for the Supermicro. This test runs a 5 GB sample file for read and write speeds. This test showed read speeds of 3.57 GB/s and nearly 2.54 GB/s read on the Ultrastar SN655 NVMe Data Center SSDs installed by Supermicro.
| Blackmagic Disk Speed Test (Higher is Better) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| Write | 2,536 MB/s | 1,415.1 MB/s | 
| Read | 3,568 MB/s | 3,031.0 MB/s | 
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better. Here are the results for all EPYC chips.
Here, the Supermicro 1115SV-WTNRT scored 63,332 points in the multi-core test, demonstrating the processor's robust ability to handle multiple threads simultaneously. On the other hand, the single-core test score was 1,093 points, reflecting its efficiency in tasks requiring single-threaded performance. The MP ratio, which compares multi-core performance to single-core performance, stands at 57.92, indicating a well-balanced architecture for parallel processing and per-core efficiency.
| Cinebench R23 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| CPU (multi-core) (points) | 63,332 | 81,148 | 
| CPU (single-core) (points) | 1,093 | 1,309 | 
| MP Ratio | 57.92 | 61.99x | 
Cinebench 2024
Maxon's Cinebench 2024 is a CPU and GPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. We do not have these figures because the ASG-1115S-NE316R configuration does not have a GPU. Higher scores are better.
The 1115SV-WTNRT scored 3,928 points in the multi-core test, highlighting its strong performance in terms of processing power for multitasking and demanding computational tasks. Single-core performance was measured at 69 points, while the GPU benchmark reached a score of 3,634 points, demonstrating its mastery of graphics processing. The MP ratio was calculated at 57.27, demonstrating a consistent balance between multi-threaded and single-threaded processing capabilities.
| Cinebench 2024 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| CPU (multi-core) (Points) | 3,928 | 4,913 | 
| CPU (single-core) (Points) | 69 | 81 | 
| GPU | 3,634 | N/A | 
| MP Ratio | 57.27 | 60.47x | 
Geekbench CPU Benchmark
Geekbench 6 is a cross-platform performance evaluation tool that measures the overall performance of a system. However, it would be interesting to compare single-core and multi-core performance, as well as OpenCL performance. A high score indicates better performance.
Here, the system demonstrated solid overall performance in terms of CPU and GPU metrics. The single-core processor benchmark recorded a score of 1,718, while the multi-core processor test showed 19,455. Additionally, GPU performance was evaluated using the OpenCL framework, with a score of 35,764, indicating solid graphics processing power suitable for various compute-intensive applications.
You can find comparisons with any system in the Geekbench browser.
| Geekbench 6 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| CPU Benchmark - Single-Core | 1,718 | 2,055 | 
| CPU Benchmark - Multi-Core | 19,455 | 22,868 | 
| GPU Benchmark – OpenCL | 35,764 | N/A | 
y-cruncher
y-cruncher is a multi-threaded and scalable program that can calculate Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress testing application for overclockers and hardware enthusiasts.
In terms of results, we have results ranging from 1 billion to 25 billion for the Supermicro 1115SV-WTNRT.
| y-cruncher (Total Computation Time) (lower is better) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| 1 billion digits (seconds) | 9.989 seconds | 7.274 seconds | 
| 2.5 billion digits (seconds) | 24.974 seconds | 17.055 seconds | 
| 5 billion digits (seconds) | 52.117 seconds | 34.336 seconds | 
| 10 billion digits (seconds) | 110.483 seconds | 71.336 seconds | 
| 25 billion digits (seconds) | 303.372 seconds | 196.695 seconds | 
| 50 billion digits (seconds) | N/A | 439.435 seconds | 
UL Procyon AI Inference
UL Procyon AI Inference is designed to evaluate a workstation's performance in professional applications. It is important to note that this test does not exploit the capabilities of multiple processors. Specifically, this tool evaluates the workstation's ability to handle AI-based tasks and workflows, providing a detailed analysis of its efficiency and speed in processing complex AI algorithms and applications.
The results cover a range of AI models, indicating the server's versatility in handling different types of AI workloads.
Inference times range from very fast (4.11 ms for MobileNet V3) to much slower (2,131.81 130 ms for Real-ESRGAN), indicating a wide variation in how the system handles different AI tasks. Given the complexity and variety of functions, a score of XNUMX suggests that while the system performs adequately across a wide range of AI-based operations, there may be limitations when handling more resource-intensive tasks like Real-ESRGAN.
| UL Procyon Average Inference Times (lower is better) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | 
| Mobile Net V3 | 4.11 ms | 
| ResNet 50 | 8.40 ms | 
| Inception V4 | 30.27 ms | 
| Deep Lab V3 | 30.87 ms | 
| YOLO V3 | 45.66 ms | 
| Real-ESRGAN | 2,131.81 ms | 
| Overall Score | 130 | 
7-Zip Compression
The popular 7-Zip utility has a built-in memory benchmark that demonstrates CPU performance. In this test, we run it with a 128 MB dictionary size when possible. Again, we only have the results for the ASG-1115S-NE316R for this test.
|  | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| Compression |  |  | 
| Current CPU Usage | 5,596 % | 3,574 % | 
| Current Rating/Usage | 4.387 | 5.755 GIPS | 
| Current | 245.508 | 205.670 GIPS | 
| Resulting CPU Usage | 5,615 % | 3,572 % | 
| Resulting Rating/Usage | 4.383 | 5.733 GIPS | 
| Resulting Rating | 246.126 | 204.758 GIPS | 
| Decompression |  |  | 
| Current CPU Usage | 6,198 % | 3,810 % | 
| Current Rating/Usage | 4.224 GIPS | 5.753 GIPS | 
| Current | 261.832 GIPS | 219.191 GIPS | 
| Resulting CPU Usage | 6,057 % | 3,803 % | 
| Resulting Rating/Usage | 4.382 GIPS | 5.779 GIPS | 
| Resulting Rating | 265.381 GIPS | 219.768 GIPS | 
| Total Rating |  |  | 
| Total CPU Usage | 5,836 % | 3,687 % | 
| Total Rating/Usage | 4.383 GIPS | 5.756 GIPS | 
| Total Rating | 255.753 GIPS | 219.768 GIPS | 
Conclusion
The Supermicro Server AS-1115SV-WTNRT stands out as a robust 1U rack-mounted server, capable of handling high-demand applications, including virtualization, database management, and edge computing. Equipped with an AMD EPYC 8004 "Siena" processor, the AS-1115SV-WTNRT appears as an attractive option for service providers looking to improve data center efficiency and performance while controlling costs.
The AMD EPYC 8004 series excels in single-socket platforms, offering many cores at a lower cost than higher-core models, as well as low power consumption (starting at just 70 watts). This makes them particularly suited for dense data centers where space and power are limited. Additionally, these processors support six channels of DDR5 memory in a compact footprint, enabling large configurations that maintain both speed and efficiency.
The integration of these processors into platforms such as the Supermicro Server AS-1115SV-WTNRT aligns with current trends in cost-effective data center operations while efficiently scaling to meet future demands. Service providers who deploy these processors will certainly experience notable improvements in operational efficiency and a reduction in total cost of ownership (TCO). That said, the AS-1115SV-WTNRT offers an efficient and scalable solution that perfectly addresses modern computing challenges, providing flexibility and redundancy in storage options.
