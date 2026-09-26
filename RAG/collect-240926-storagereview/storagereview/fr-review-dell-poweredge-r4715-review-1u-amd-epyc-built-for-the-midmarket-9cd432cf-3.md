---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf-3
title: "fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "compute", "ethernet", "gpus", "license", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf.md
source_anchor: ""
source_lines: [66, 110]
sha256: 43bc20c18e5267323a35285264fa48f49c3b3fa7012b7b050d846cc91588bef0
---

# fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf

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
