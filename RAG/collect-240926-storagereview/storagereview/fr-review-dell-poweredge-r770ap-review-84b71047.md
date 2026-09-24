---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770ap-review-84b71047
title: "fr-review-dell-poweredge-r770ap-review-84b71047"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "compute", "datacenter", "ethernet", "exploit", "gpu", "inference", "intel", "latency"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770ap-review-84b71047.md
source_anchor: ""
source_lines: [1, 137]
sha256: f8366cf7b1493136f58ef8b8384e3b68b33280e6e6e03a32e81aa8fc3dfd6f2d
---

# fr-review-dell-poweredge-r770ap-review-84b71047

<!-- source: https://www.storagereview.com/fr/review/dell-poweredge-r770ap-review -->

The Dell PowerEdge R770AP is not a general-purpose server, and that is precisely its purpose. While most dual-processor 2U platforms prioritize flexibility, the R770AP skips all of it, sacrificing GPU support, mixed storage options, and raw memory capacity in favor of maximum core density, memory bandwidth, and execution stability available in Dell's current Intel lineup. It is a server built around a specific processor architecture for a particular class of workloads, and it fully embraces its limitations.
To understand its reason for being, one must start with the platform it is built on. Dell's PowerEdge R7x0 lineup has always been the brand's most versatile Intel 2U server, with the PowerEdge R7725, equipped with an AMD processor, playing an equivalent role on the EPYC side. The PowerEdge R770 continues this Intel tradition by supporting Xeon 6 processors with P-cores and E-cores, GPU accelerators, mixed SAS/SATA/NVMe storage, up to 8 TB of memory across 32 DIMM slots, and enough PCIe Gen5 expansion capacity to cover all needs, from virtualization to AI inference.
The PowerEdge R770AP is not that server.
The "AP" designation stands for "Advanced Performance," but it does not fully convey the differences between these two machines. While the R770 uses the Intel Granite Rapids-SP chip on the LGA 4710 socket with 8 memory channels and up to 86 performance cores (P-cores), the R770AP adopts the Granite Rapids-AP platform on the LGA 7529 socket, offering up to 128 performance cores per socket (120 cores in our test configuration) and 12 DDR5 memory channels. This distinction is at the heart of Intel's overall strategy for its Xeon 6 6900 series: the 6900P processors on the AP platform represent the pinnacle of Intel server chips, designed specifically for workloads where per-core performance, memory bandwidth, and execution stability take precedence over overall server configuration flexibility.
Intel's extended Xeon 6 architecture divides the data center into two categories. E-core processors prioritize density and power efficiency for cloud-native and scalable workloads, such as microservices and content delivery. P-core processors target compute-intensive tasks where consistent per-thread performance is essential: HPC simulations, real-time analytics, large in-memory databases, and latency-sensitive financial calculations. The 6900P series sits at the top of this P-core processor range, combining the highest number of available cores with 12-channel memory bandwidth, up to 96 PCIe Gen5 lanes per socket, up to 6 UPI 2.0 links, and L3 cache pools reaching 504 MB on high-end models like the Intel Xeon 6978P. The architectural goal is not just raw throughput, but predictable throughput, minimizing scheduling fluctuations and memory access variability that degrade performance in latency-critical environments.
The R770AP is the very embodiment of this philosophy at Dell. It strips away everything superfluous on the Granite Rapids-AP platform: GPU support is entirely removed, SAS and SATA storage options are replaced by exclusively NVMe configurations (up to 16 2.5-inch Gen5 NVMe SSDs or up to 32 E3.S Gen5 NVMe SSDs, depending on configuration), memory capacity is capped at 3 TB across 24 DIMM slots (12 per socket, 1 DPC for maximum per-channel speed), and PCIe expansion is reduced to five Gen5 x16 slots and two OCP 3.0 network cards. The result is a dual-socket 2U platform optimized for compute density, memory bandwidth, and the deterministic behavior demanded by workloads such as high-frequency trading, real-time risk analytics, and massively parallel simulation.
Our test model pairs two Intel Xeon 6978P processors, each with 120 performance cores clocked at 2.1 GHz (base frequency) and 3.2 GHz (all-core turbo mode), and carries 3 TB of DDR5-6400 memory across all 24 DIMM slots. Compared to the R770, equipped with two Xeon 6787P processors (86 cores each, 8 memory channels, and 2 TB of DDR5), the R770AP offers 39.5% more cores and 50% more memory channels. The question remains whether these architectural advantages translate into concrete performance gains and whether the platform's compromises are justified for the workloads targeted by Dell and Intel.
Dell PowerEdge R770AP Specifications
The table below presents the physical configuration and support specifications for the Dell PowerEdge R770AP platform.
| Specifications | Dell PowerEdge R770AP | 
|---|---|
| Processor |  | 
| Processor | Two Intel® Xeon® 6 6900 series processors with P-cores, up to 128 cores each. | 
| Memory |  | 
| DIMM Slots | 24 DDR5 DIMM slots | 
| Maximum Memory | 3 TB | 
| Memory Speed | Up to 6400 MT/s | 
| Memory Type | Registered ECC DDR5 RDIMM modules only | 
| Storage |  | 
| Storage Controllers (RAID) | Front (internal) PERC H975i DC-MHS | 
| Internal Boot | BOSS-N1 DC-MHS: HWRAID 1, 2 M.2 NVMe SSDs or USB | 
| Front Drive Bays | Up to 16 2.5-inch Gen5 x4 NVMe SSDs (maximum capacity of 245.76 TB) Up to 16 2.5-inch Gen5 x2 NVMe SSDs (maximum capacity of 245.76 TB) Up to 32 EDSFF E3.S Gen5 NVMe SSDs (maximum capacity of 491.52 TB) | 
| Rear Drive Bays | N/A | 
| Engine Tuning |  | 
| Power Supplies | 1500 W Titanium, 100-120 LLAC or 200-240 HLAC, 240 V DC, hot-swap redundancy 1800 W Titanium, 200-240 HLAC, 240 VDC, hot-swap redundancy 2400 W Titanium, 100-120 LLAC or 200-240 HLAC, 240 V DC, hot-swap redundancy 3200 W Titanium, 200-220 HLAC or 220.1-240 HLAC, 240 VDC, hot-swap redundancy 3200 W Titanium, 277 VAC and HVDC, hot-swap redundancy* | 
| Cooling and Fans |  | 
| Cooling Options | air cooling | 
| Fans | Up to 6 hot-swap fans | 
| Form Factor and Dimensions |  | 
| Form Factor | 2U rack server | 
| Height | 86.8 mm (3.42 inches) | 
| Width | 482 mm (19.0 inches) | 
| Depth (with bezel) | 802.40 mm (31.59 inches) | 
| Depth (without bezel) | 801.51 mm (31.56 inches) | 
| Bezel | Optional metal bezel | 
| Networking and Expansion |  | 
| OCP Network Options | Up to two OCP NIC 3.0 network cards Slot 4: 1×8 or 1×16 Gen5 OCP 3.0 Slot 10: 1×16 Gen5 OCP 3.0 | 
| Integrated Network Card | Dedicated 1 Gb BMC Ethernet port | 
| PCIe Slots | Up to 5 PCIe Gen5 slots (x16 connectors) Slot 2: 1×16 Gen5, full height, half length Slot 3: 1×16 Gen5, full height/low profile, half length Slot 5: 1×16 Gen5, full height, half length Slot 7: 1×16 Gen5, full height, half length Slot 9: 1×16 Gen5, full height/low profile, half length | 
| GPU Options | N/A | 
| Ports |  | 
| Front Ports | 1x USB 2.0 Type-C | 
| Rear Ports | 1 dedicated BMC Ethernet port 2x USB 3.1 Type-A 1x VGA | 
| Internal Ports | 1x USB 3.1 Type-A | 
| Management |  | 
| Integrated Management | iDRAC10, iDRAC Direct, iDRAC RESTful API with Redfish, RACADM CLI, iDRAC Service Module | 
| Security |  | 
| Security Features | Cryptographically signed firmware, data-at-rest encryption (SED with local or external key management), secure boot, secure component verification (hardware integrity check), secure erase, hardware root of trust, system lock (requires iDRAC10 Enterprise or Datacenter), FIPS/CC-TCG certified TPM 2.0, chassis intrusion detection | 
| Operating Systems and Hypervisors |  | 
| Supported Operating Systems/Hypervisors | Canonical Ubuntu Server LTS, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware vSAN / VMware ESXi*, Microsoft Windows, Microsoft Windows Server, Microsoft Windows Server Datacenter | 
Design and Build
The Dell PowerEdge 770AP is a 2U rack server from Dell's 17th generation of PowerEdge servers, sharing the same design as the R770 we tested. It measures 3.42 cm in height, 19.0 cm in width, and 31.59 cm in depth. The front bezel is optional. The front panel includes iDRAC direct access, a USB 2.0 Type-C port, a power button, and a system identification button.
Storage
The 770AP server supports three storage configurations. It ships with up to 16 2.5-inch Gen 5 x4 NVMe SSDs, for a maximum capacity of 245.76 TB. It is also possible to opt for up to 16 2.5-inch Gen 5 x2 NVMe SSDs, capped at 245.76 TB, or up to 32 EDSFF E3.S Gen 5 NVMe SSDs, expandable to 491.52 TB. In 16-bay configurations, Dell splits the drives into two groups of eight, located on the left and right sides of the server, with the central section serving as an air intake.
Looking more closely inside the chassis, the 770AP features clean, direct NVMe cabling. The cables connect directly from the storage backplane to the front edge of the motherboard, shortening the signal path and optimizing internal organization.
Rear I/O and Networking
Two redundant 2400 W power supplies are mounted at the rear of the 770AP, at each end. The BOSS-N1 module handles booting and includes two 480 GB drives for the operating system.
For expansion, the server offers up to five PCIe Gen 5 slots across slots 2, 3, 5, 7, and 9, all equipped with x16 connectors in full-height configuration. OCP 3.0 network connectivity is provided by a maximum of two cards: slot 4 supports Gen 5 x8 or x16 interfaces, and slot 10 offers a dedicated x16 Gen 5 connection. Our machine shipped with one 200 GbE OCP card and several 100 GbE cards, ensuring more than sufficient network bandwidth.
Standard rear I/O includes one dedicated BMC Ethernet port, two USB 3.1 Type-A ports, and one VGA port.
A closer examination of the BOSS-N1 module reveals two 480 GB boot drives side by side, both hot-swappable and easy to access and replace when needed.
Once the top cover and air shrouds are removed, the interior of the R770AP reveals itself to be clean and well organized. Six hot-swap fans push air through the large heat sinks, cooling the Xeon 6900 series processors, whose dual-processor and memory configuration is arranged symmetrically. Also noticeable are the blue tabs on the chassis, which serve as guides for disassembly, cable removal, and component access.
Processor
Once the processor is removed, the imposing size of the Intel Xeon 6900 series chip is immediately apparent. The R770AP motherboard uses the LGA 7529 socket, and our test model was equipped with two Intel Xeon 6978P processors. Each chip has a TDP of 500 W and 120 cores, bringing the total core count to 240 across both sockets.
Cooling and Memory
To handle the processor's 1000 W heat dissipation with air cooling alone, Dell designed a specific cooling system. The front and rear heat sinks use horizontal fins and heat pipes for efficient thermal dissipation. The central section, meanwhile, features a stack of vertical fins that increases contact time and heat exchange surface area, allowing the fans to more efficiently expel heat before it exits the chassis. In total, 24 DIMM slots are integrated into the cooling system: each processor is flanked by 12 slots, six on each side.
Engine Tuning
The R770AP supports four power supply options, all 80 Plus Titanium certified and hot-swappable: 1500 W, 1800 W, 2400 W, and 3200 W. With consumption reaching up to 1000 W for the processors alone, the base 1500 W configuration offers very little headroom once drives and expansion cards are taken into account. Our model was equipped with a 2400 W power supply, delivering 96% efficiency, which represents the practical minimum for a fully populated storage configuration.
iDRAC 10 Management
Remote management of the R770AP is handled by iDRAC10, the same platform Dell offers as standard across its entire 17th generation PowerEdge lineup, including the PowerEdge R770 and R7725 we previously tested. Since the interface is identical across the lineup, administrators already familiar with iDRAC on other PowerEdge platforms will find their way around easily.
The iDRAC10 dashboard provides a complete, instant overview of the health status of each major subsystem: system status, processor, memory, cooling, storage, voltages, power supplies, batteries, and intrusion detection. The test unit indicates that all subsystems were operational at the time of testing. System information and firmware version details are displayed directly on the dashboard, along with license status, which on the test unit is confirmed as Enterprise type. The "Task Summary" panel tracks pending, in-progress, and completed tasks. The test unit shows completed tasks from an initial provisioning cycle, a few with errors and one that failed, which is typical of a new deployment.
Exploring the "System Environments" section gives you access to cooling details, including the status of each fan, PWM speeds, thermal profile settings, and inlet temperature readings, all in real time. This feature is particularly useful for verifying airflow in dense rack configurations or for diagnosing thermal issues without having to physically access the server.
Power consumption visibility follows the same principle. The "Power Information" section details power supply status, current draw, and utilization rate, along with a rolling historical graph. Administrators can thus quickly visualize average and peak consumption over time, which is valuable for capacity planning and identifying workload-related power spikes, without needing an additional monitoring tool.
Together, these views make iDRAC10 a capable out-of-band management solution that covers the entire operational lifecycle of the R770AP, from initial deployment to daily monitoring, all accessible remotely via a browser or the Redfish RESTful API.
Dell PowerEdge R770AP Performance
To evaluate the R770AP, we compared it directly to the R770. The R770AP is equipped with two Intel Xeon 6978P processors, each with 120 cores, for a total of 240 cores and 3 TB of DDR5 memory. The R770, meanwhile, includes two Intel Xeon 6787P processors, for a total of 172 cores and 2 TB of DDR5 memory.
To stress the processors of both systems, we used a targeted set of compute tests. y-cruncher evaluated raw arithmetic throughput and multithreaded floating-point performance. Blender provided a realistic rendering workload, scaling with the number of available cores and memory bandwidth. The Phoronix test suite rounded out the set with a broader collection of CPU-intensive workloads, providing a more complete picture of sustained compute performance on both platforms.
Test System Specifications
- Platform: Dell PowerEdge R770AP
- CPU: Dual Intel Xeon 6978P, 120 cores
- Memory: 3 TB DDR5
- Storage: Boss RAID1
y-cruncher
y-cruncher is a popular application for performance and stress testing systems, launched in 2009. This multithreaded and scalable test calculates Pi and other constants to trillions of decimal places. The faster the test, the better. This software has proven excellent for testing high-core-count platforms and demonstrating the compute advantages between single-processor and dual-processor platforms.
In the y-cruncher benchmark, the R770AP consistently outperformed the R770, regardless of the data size tested. In the 1 billion decimal test, the R770AP finished in 2.692 seconds, versus 2.753 seconds for the R770. At 10 billion decimals, the R770AP achieved a time of 30.399 seconds, versus 34.873 seconds for the R770. At 50 billion decimals, the R770AP posted a time of 192.128 seconds, versus 221.255 seconds for the R770. The gap widened for the largest workload: the 100 billion decimal test was completed in 430.208 seconds by the R770AP, versus 491.737 seconds by the R770, a difference of approximately 61 seconds and a performance gain of approximately 12.5% for the R770AP.
| Y-cruncher (shorter time is better) | Dell PowerEdge R770 (2x Intel Xeon 6787P \| 2 TB RAM) | Dell PowerEdge R770AP (2x Intel Xeon 6978P \| 3 TB RAM) | 
|---|---|---|
| 1 billion | 2.753 seconds | 2.692 seconds | 
| 2.5 billion | 7.365 seconds | 6.747 seconds | 
| 5 billion | 16.223 seconds | 14.235 seconds | 
| 10 billion | 34.873 seconds | 30.399 seconds | 
| 25 billion | 99.324 seconds | 86.298 seconds | 
| 50 billion | 221.255 seconds | 192.128 seconds | 
| 100 billion | 491.737 seconds | 430.208 seconds | 
Blender
An open-source 3D modeling application. This benchmark was run with the Blender Benchmark utility. The score is expressed in samples per minute, with the higher being the better.
In the Blender 4.3 benchmark, the R770AP outperformed the R770 in all three scenes. In the "Monster" scene, the R770AP reached 2,200.116 samples per minute, versus 1,706.002 for the R770. In the "Junkshop" scene, the R770AP achieved 1,565.643 samples per minute, versus 1,169.370 for the R770. Finally, in the "Classroom" scene, the R770AP obtained 1,076.122 samples per minute, versus 791.475 for the R770, a performance gain of approximately 36% on this workload.
| CPU performance test with Blender 4.3 (higher samples per minute is better) | Dell PowerEdge R770 (2x Intel Xeon 6787P \| 2 TB RAM) | Dell PowerEdge R770AP (2x Intel Xeon 6978P \| 3 TB RAM) | 
|---|---|---|
| Monster | 1,076.122 samples/min | 1,076.122 samples/min | 
| Junkshop | 1,076.122 samples/min | 1,076.122 samples/min | 
| Classroom | 1,076.122 samples/min | 1,076.122 samples/min | 
Phoronix Benchmarks
Phoronix Test Suite is an open-source automated benchmarking platform supporting over 450 test profiles and over 100 test suites via OpenBenchmarking.org. It handles the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration. Here we will compare the performance of the R770AP and R770 using the Stream, 7-Zip, Linux kernel compilation, Apache, and OpenSSL tests.
Stream
In the Stream memory bandwidth test, the R770AP achieved a clear improvement over the R770, reaching 869,965.3 MB/s versus 472,135.6 MB/s. This represents nearly double the memory bandwidth of the reference system, reflecting the R770AP's larger and faster memory configuration.
7-Zip
In the 7-Zip compression test, the R770AP scored 806,375 MIPS, versus 628,206 MIPS for the R770, a clear improvement due to the higher core count of the 6978P processors.
Kernel Compilation
In the Linux kernel compilation test, where a shorter time is preferable, the R770AP completed the allmod compilation in 176.391 seconds versus 188.793 seconds on the R770, reducing compilation time by approximately 12 seconds.
Apache
The Apache test was the only area where the R770 slightly outperformed the R770AP, with a score of 60,258.5 requests per second versus 48,729.63 for the R770AP. This result is important because web server workloads do not always scale linearly with core count and can be influenced by memory latency and I/O characteristics.
OpenSSL
In the OpenSSL verification test, the R770AP scored 2,515,270,390,853 verifications/s versus 2,216,883,554,350 verifications/s on the R770, a significant gain in cryptographic throughput that highlights the computational efficiency of the 6978P at scale.
| Phoronix Benchmarks | Dell PowerEdge R770 (2x Intel Xeon 6787P 86C) | Dell PowerEdge R770AP (2x Intel Xeon 6978P \| 3 TB RAM) | 
|---|---|---|
| Stream | 472,135.6 MB/s | 869,965.3 MB/s | 
| 7-ZIP | 628,206 MIPS | 806,375 MIPS | 
| Kernel Compilation (allmod) (lower is better) | 188.793 seconds | 176.391 seconds | 
| Apache (requests per second) | 60,258.5 R/s | 48,729.63 R/s | 
| OpenSSL | 2,216,883,554,350 Verifications | 2,515,270,390,853 Verifications | 
Dell PowerEdge R770AP: Deterministic Performance and High-Frequency Trading
While our standard test suite focuses on compute throughput, memory bandwidth, and general scalability, the R770AP's design priorities extend into an area we typically do not test: microsecond-level execution determinism. To illustrate this platform's capabilities for its most demanding target audience, Dell published a technical note in partnership with Metrum AI, evaluating the R770AP specifically for high-frequency trading workloads. We did not perform these tests, nor did we independently audit the results. Nevertheless, we include a summary here, as it most directly demonstrates how this server is a distinct product from the R770.
Metrum AI's methodology relies on a custom tool called jitter-c, which measures per-core wake-up latency jitter—that is, the regularity with which a thread scheduled to execute at a precise moment actually starts. This metric isolates processor scheduling variability from network, memory, and application factors, providing a reliable point of comparison between processor generations. By comparing an R770AP equipped with two Xeon 6980P processors (256 cores total) to a previous-generation R760 with two Xeon Platinum 8592+ processors (128 cores total), the study found that the Granite Rapids-AP architecture reduced p99 wake-up jitter to approximately 1 microsecond—roughly half that of the older platform—while doubling core density. These jitter profiles were then fed into a backtesting simulation engine to model the financial impact. The results are summarized below.
| Metrum AI HFT Backtest Results | Dell PowerEdge R760 (2x Xeon 8592+, 128 cores) | Dell PowerEdge R770AP (2x Xeon 6980P, 256 cores) | 
|---|---|---|
| p99 Wake-up Jitter | ~2 µs | ~1 µs | 
| Mean Reversion: Total Trades | 5,175 | 6,229 (+20.4%) | 
| Mean Reversion: Trades/sec | 819 | 991 (+21.1%) | 
| Market Making: Total Trades | 21,765 | 32,491 (+49.3%) | 
| Market Making: Trades/sec | 2,067 | 3,072 (+48.6%) | 
As Seamus Jones of Dell noted in his commentary on the study, the added value lies not in the speed, but in the predictability of that speed. Indeed, in trading, a fast but inconsistent system is a source of risk. Conversely, a deterministic system is a strategic asset.
Conclusion
The Dell PowerEdge R770AP occupies a well-defined place within the 17th generation PowerEdge lineup. It does not replace the R770, and Dell does not present it as such. The R770 remains the versatile and highly configurable Intel 2U platform it has always been, with GPU support, mixed SAS/SATA/NVMe storage, E-core and P-core processor options, and up to 8 TB of memory across 32 DIMM slots. For organizations running general virtualization solutions, mixed enterprise applications, or workloads that leverage that configuration flexibility, the R770 remains the ideal choice.
The R770AP is designed for workloads for which the R770 was never optimized. By adopting the Granite Rapids-AP platform, with its 12-channel memory architecture, up to 128 processing cores per socket, and 504 MB of L3 cache, Dell has created a 2U system that prioritizes compute density, memory bandwidth, and execution determinism over versatility. Our performance tests reflect this priority: STREAM bandwidth nearly doubled, Blender rendering improved by 29 to 36%, and CPU scaling extended consistently as working sets exceeded cache capacity. The Apache regression is an important point to note, as it demonstrates that the R770AP's NUMA topology requires workload consideration to fully exploit its performance, and that not all applications will benefit from this platform change without optimization.
The Metrum AI tests published by Dell alongside this platform highlight the underlying determinism. Halving p99 scheduling jitter while doubling core density represents a significant architectural improvement for organizations running high-frequency trading operations, real-time risk engines, large-scale in-memory analytics, and massively parallel simulations. For these workloads, the R770AP is a high-performance platform that is perfectly suited. For all other applications, the R770 and R7725 remain the most relevant options within the PowerEdge lineup.
