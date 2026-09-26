---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770ap-review-84b71047-3
title: "fr-review-dell-poweredge-r770ap-review-84b71047"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "compute", "ethernet", "intel", "license", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770ap-review-84b71047.md
source_anchor: ""
source_lines: [56, 92]
sha256: 52552f70c6b5a937c48f6d4907ae42dbf64900e0680dcc7370a34796b9f6b389
---

# fr-review-dell-poweredge-r770ap-review-84b71047

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
