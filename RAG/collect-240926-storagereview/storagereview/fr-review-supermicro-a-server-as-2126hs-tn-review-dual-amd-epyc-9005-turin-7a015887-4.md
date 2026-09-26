---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887-4
title: "fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "compute", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887.md
source_anchor: ""
source_lines: [121, 169]
sha256: bced9210f937d1ce573e08ffa66423f30b5783c085ed42baa758735d0e775285
---

# fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887

For comparison, the Dell PowerEdge R7725, equipped with the same two EPYC 9965 processors, shows slightly higher throughput in this scenario with SMT enabled, reaching 3,193.11 samples per minute in Monster, 2,174.63 in Junkshop, and 1,608.79 in Classroom. The relatively small gap between the two platforms suggests broadly similar behavior under SMT-intensive rendering loads, with differences likely due to platform-level optimization rather than raw compute power.
| Blender CPU SMT (Samples per minute; higher is better) | Supermicro A+ Server AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (dual AMD EPYC 9965 192C) | 
|---|---|---|
| Monster | 3,070.84 | 3,193.11 | 
| Junkshop | 2,063.61 | 2,174.63 | 
| Classroom | 1,527.39 | 1,608.79 | 
With SMT disabled, the Supermicro system shows a notable increase in raw throughput across all three scenes. Performance reaches 4,018.10 samples per minute in Monster, 2,707.10 in Junkshop, and 1,990.51 in Classroom, demonstrating how Blender's CPU render engine can benefit from reduced thread contention on very high core count systems.
In this configuration, the Dell PowerEdge R7725 still shows higher absolute throughput, with 4,304.21, 2,870.34, and 2,079.79 samples per minute in Monster, Junkshop, and Classroom, respectively. Although the Dell system retains an advantage in peak performance, the Supermicro AS-2126HS-TN shows substantial gains with SMT disabled, highlighting the significant influence of multithreading strategy on the performance of platforms of this scale.
| Blender CPU without SMT (Samples per minute; higher is better) | Supermicro A+ Server AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (dual AMD EPYC 9965 192C) | 
|---|---|---|
| Monster | 4,018.10 | 4,304.21 | 
| Junkshop | 2,707.10 | 2,870.34 | 
| Classroom | 1,990.51 | 2,079.79 | 
y-cruncher
y-cruncher is a multithreaded and scalable program capable of calculating Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application among overclockers and hardware enthusiasts.
In the y-cruncher benchmark, the Supermicro AS-2126HS-TN demonstrates consistent and predictable scaling as problem sizes increase. The system performs the calculation of one billion decimal places in 8.092 seconds, then moves to 15.245 seconds for 2.5 billion decimal places and 24.961 seconds for 5 billion decimal places.
As the workload increases, computation times evolve linearly, reaching 44.350 seconds for 10 billion digits, 114.107 seconds for 25 billion, and 246.688 seconds for 50 billion. In the most demanding calculation, involving 100 billion digits, the system performs the operation in 572.800 seconds, demonstrating its ability to handle a high number of threads over long execution periods.
The Dell PowerEdge R7725 executes the same workloads slightly faster across all test sizes, completing the 100 billion digit run in 481.207 seconds.
| Y-Cruncher (total computation time) | Supermicro A+ Server AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (dual AMD EPYC 9965 192C) | 
|---|---|---|
| 1 billion | 8.092 s | 7.879 seconds | 
| 2.5 billion | 15.245 s | 13.811 seconds | 
| 5 billion | 24.961 s | 22.107 seconds | 
| 10 billion | 44.350 | 40.111 seconds | 
| 25 billion | 114.107 | 98.445 seconds | 
| 50 billion | 246.688 | 211.567 seconds | 
| 100 billion | 572.800 | 481.207 seconds | 
Phoronix Benchmarks
Phoronix Test Suite is an open-source automated benchmarking platform that supports over 450 test profiles and over 100 test suites via OpenBenchmarking.org. It manages the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration.
Stream Memory Bandwidth
In the Stream memory bandwidth test, the Supermicro AS-2126HS-TN achieves a sustained throughput of 807,766 MB/s. The Dell PowerEdge R7725, meanwhile, peaks at 883,312 MB/s, indicating slightly higher maximum memory bandwidth. However, both systems are within the expected performance range for dual-processor EPYC 9005 platforms.
7-Zip Compression
For 7-Zip compression, the Supermicro system achieves 1,262,832 MIPS, demonstrating excellent integer compute performance and efficient multithreaded scaling. The Dell system shows 1,326,967 MIPS, representing a slight advantage for this workload, while remaining in the same overall performance category.
Kernel Compilation
During kernel compilation (allmod), the Supermicro AS-2126HS-TN executes the task in 117.97 seconds, ahead of the Dell PowerEdge R7725 which performs the same operation in 139.36 seconds. This result highlights the efficiency of the Supermicro platform in parallel compilation scenarios common in development and continuous integration environments.
Apache Web Server
The Apache performance test shows the Supermicro system processing 90,623.69 requests per second, while the Dell system reaches 96,782.75 requests per second. Both systems show high throughput for web servers, although the Dell configuration shows a slightly higher request peak.
OpenSSL Verification
In OpenSSL tests, the Supermicro AS-2126HS-TN reaches 3.55 TB/s, demonstrating significant cryptographic throughput, suited to encryption-intensive workloads. The Dell PowerEdge R7725 goes even further with 4.41 TB/s, indicating higher peak cryptographic performance. Both platforms offer performance well beyond typical enterprise needs.
| Phoronix Benchmarks | Supermicro A+ Server AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (dual AMD EPYC 9965 192C) | 
|---|---|---|
| Stream | 807,766.0 MB / s | 883,312.0 MB / s | 
| 7-ZIP | 1,262,832 MIPS | 1,326,967 MIPS | 
| Kernel Compilation (allmod) | 117.97 seconds | 139.356 seconds | 
| Apache (requests per second) | 90,623.69 R/s | 96,782.75 R/s | 
| OpenSSL | 3,545,769,484,910 verifications | 4,409,642,672,307 verifications | 
Conclusion
The Supermicro A+ AS-2126HS-TN server offers a balanced dual-processor Turin platform, prioritizing compute density, power scalability, and flexible PCIe expansion in a 2U form factor. Compatible with two AMD EPYC 9005 series processors, high-TDP configurations, and flexible I/O configurations, it is ideal for enterprise, cloud, and high-performance computing (HPC) environments where sustained parallel performance and configurability are essential.
In our enterprise test suite, the AS-2126HS-TN delivered consistent and predictable performance for CPU-intensive workloads. Although it slightly trailed on some absolute throughput metrics, the Supermicro platform remained competitive, demonstrating excellent scalability and high efficiency for rendering, compute, and mixed workloads. These results confirm that the AS-2126HS-TN is a high-performance and flexible foundation for high core count deployments, especially when platform versatility and power headroom are as important as peak benchmark performance.
