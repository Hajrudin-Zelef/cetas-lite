---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-hyper-superserver-sys-221h-tn24r-review-e7dd20c1-3
title: "fr-review-supermicro-hyper-superserver-sys-221h-tn24r-review-e7dd20c1"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "benchmarks", "energy", "gpu", "gpus", "intel", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-hyper-superserver-sys-221h-tn24r-review-e7dd20c1.md
source_anchor: ""
source_lines: [89, 146]
sha256: 090501c8248f9b9c777db9e93047783dcd86490af29b21b068c125ba389a2a5a
---

# fr-review-supermicro-hyper-superserver-sys-221h-tn24r-review-e7dd20c1

The power consumption area provides historical and real-time data on the server's power consumption. As shown in the image below, it displays the minimum, average, and maximum power consumption over the last hour, and below the graph is a summary of power consumption since the server was powered on. This includes peak values and their corresponding times. Additionally, it also provides users with historical trends for the last hour, last day, and last week, detailing average usage, maximum consumption, and minimum power consumption.
This information is extremely useful for those looking to manage their system's energy efficiency as well as ensure the server operates within its power capacity limits.
Supermicro Hyper SuperServer SYS-221H-TN24R Performance
We continue to leverage the performance results obtained during our initial analysis of 5th Gen Intel Xeon Scalable processors. Our study focuses on the SYS-221H-TN24R system, equipped with two 8562Y+ processors. This configuration will be compared to our Supermicro E1.S platform as a reference, the latter being equipped with two 8460H processors.
Supermicro Hyper SuperServer SYS-221H-TN24R Specifications
- 2 x Intel Xeon Platinum 8562Y+ processors
- 512GB DDR5
- Windows Server 2022
Supermicro Storage SuperServer SSG-121E-NES24R Specifications
- 2 x Intel Xeon Platinum 8460H processors
- 512GB DDR5
- Windows Server 2022
The Intel Xeon Platinum 8562Y and 8460H, belonging to the 5th and 4th generations of Intel Xeon Scalable processors respectively, present notable differences. The 8562Y+ operates at a base frequency of 2.80 GHz and offers 32 cores with 64 threads, contrasting with the 8460H's 40 cores and 80 threads at a base frequency of 2.2 GHz.
The 8460H also stands out with a significantly larger cache of 105 MB, compared to the 60 MB of the 8562Y+, which could prove advantageous in applications requiring intensive data processing. Both models support DDR5 memory. However, the 8460H is less energy-efficient, with a TDP of 330W versus the 8562Y+'s 300W. In terms of price, the 8460H is significantly more expensive, at over $10,000, while the 8562Y+ costs a little under $6,000.
To provide a comprehensive performance analysis, our tests encompass a variety of intensive benchmarks, each targeting different aspects of CPU performance. We will use Blender OptiX for rendering performance, Blackmagic to evaluate video processing capabilities, and Geekbench for overall system performance evaluation. Additionally, Cinebench will give us insight into graphics and rendering efficiency, y-cruncher for mathematical calculations, and 7-Zip compression tests to evaluate data processing and compression speeds.
These various benchmarks will help us provide a comprehensive evaluation, covering both general and specialized computing tasks, to see how they perform in real-world scenarios.
Blender OptiX
The first is the Blender test, an open-source 3D modeling application. This benchmark was run using the Blender Benchmark utility. The score is expressed in samples per minute, with the higher being the better.
In Blender version 4.0, here are the scores for the two Intel CPUs:
| Blender 4.0 CPU | 2 x Platinum 8562Y+ 32 cores | 2x Platinum 8460H 40 cores | 
| Monster | 805.254137 | 671.717058 | 
| Junkshop | 513.800608 | 424.825295 | 
| Classroom | 414.999628 | 347.791937 | 
Blackmagic RAW Speed Test
We also started running the Blackmagic RAW speed test, which tests video playback performance. Additionally, this test is more of a hybrid test combining both CPU and GPU in a real-world scenario for RAW decoding.
| Blackmagic RAW Speed Test | 2 x Platinum 8562Y+ 32 cores | 2x Platinum 8460H 40 cores | 
| 8k CPU | 175 FPS | 145 FPS | 
Geekbench 6
Geekbench 6 is a cross-platform evaluation tool measuring the overall performance of a system. However, it would be interesting to analyze single-core and multi-core performance, as well as OpenCL benchmark results. A high score indicates better performance. Let's clarify that we only examined CPU results, as no GPU is installed on the servers.
| Geekbench 6 | 2 x Platinum 8562Y+ 32 cores | 2x Platinum 8460H 40 cores | 
| Single Core | 2,149 | 1,111 | 
| Multi-Core | 22,494 | 8,589 | 
You can find comparisons with any system in the Geekbench browser.
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that utilizes all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better.
| Cinebench R23 | 2 x Platinum 8562Y+ 32 cores | 2x Platinum 8460H 40 cores | 
| Multi-core CPU | 103,848 | 75,720 | 
| Single-core CPU | 1,500 | 1,068 | 
| MP Ratio | 69.25x | 70.92x | 
Cinebench 2024
Here are the results for the 2024 version of Cinebench, looking at the CPU.
| Cinebench R24 | 2 x Platinum 8562Y+ 32 cores | 2x Platinum 8460H 40 cores | 
| Multi-core CPU | 5,289 | 4,243 | 
| Single-core CPU | 82 | 62 | 
| MP Ratio | 62.1 | 68.52 | 
y-cruncher
y-cruncher is a multithreaded and scalable program that can calculate Pi and other mathematical constants to billions of digits. Since its launch in 2009, y-cruncher has become a popular benchmarking and stress testing application among overclockers and hardware enthusiasts. Faster is better in this test.
| y-cruncher | 2 x Platinum 8562Y+ 32 cores | 2x Platinum 8460H 40 cores | 
| 50 billion | 363.758 seconds | N/A | 
| 25 billion | 164.066 seconds | 191.909 seconds | 
| 10 billion | 57.868 seconds | 67.484 seconds | 
| 5 billion | 26.934 seconds | 30.813 seconds | 
| 2.5 billion | 12.036 seconds | 14.059 seconds | 
| 1 billion | 4.344 seconds | 5.134 seconds | 
Conclusion
The Supermicro Hyper SuperServer SYS-221H-TN24R is designed for high-demand computing workloads, offering support for 5th Gen Xeon Scalable processors and up to 8 TB of DDR5 memory. This 2U rack server, capable of handling tasks such as virtualization, AI, and cloud computing, stands out for its flexibility and power. The inclusion of 24 hot-swappable NVMe/SATA/SAS drive bays and two internal M.2 NVMe/SATA slots offers fairly extensive storage options, enhancing its overall utility and adaptability in various settings. Its design also emphasizes easy maintenance and comprehensive monitoring of crucial components, ensuring optimal performance.
The server's networking and expansion capabilities, including AIOM slots and optional PCIe configurations, make it a versatile choice for a range of applications. Additionally, its support for up to four GPUs means it is well-suited for graphics-intensive tasks and high-performance computing, while security features such as TPM 2.0, secure boot, and regular firmware updates ensure a secure and reliable server environment. As for its performance results in our various benchmarks, it certainly demonstrated its efficiency in handling demanding tasks, making it a serious contender for businesses needing a high-performance and scalable server solution.
Ultimately, its NVMe storage configurations, dual-processor configuration, and various expansion options make the SYS-221H-TN24R an ideal solution for businesses looking for a flexible, versatile platform to handle most enterprise applications.
