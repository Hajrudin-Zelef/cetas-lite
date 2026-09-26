---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r5715-review-bc780bb6-3
title: "fr-review-dell-poweredge-r5715-review-bc780bb6"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "compute", "consumer", "ethernet", "gpu", "license", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r5715-review-bc780bb6.md
source_anchor: ""
source_lines: [62, 104]
sha256: 3e7c3fcd5133391cf7a2d76265647996d2cda460318ad739807882d7ee0711b9
---

# fr-review-dell-poweredge-r5715-review-bc780bb6

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
