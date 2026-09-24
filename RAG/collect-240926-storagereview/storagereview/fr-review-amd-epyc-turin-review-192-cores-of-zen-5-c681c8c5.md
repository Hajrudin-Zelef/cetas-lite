---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5
title: "fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "TSMC"]
dates: []
keywords: ["amd", "benchmark", "compute", "decode", "distribution", "energy", "gpu", "gpus", "inference", "intel", "latency", "liquid cooling"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5.md
source_anchor: ""
source_lines: [1, 143]
sha256: cd60394f8f28a01e6df5a95234891059f251edf1034f09694092ff798d49a4ed
---

# fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5

<!-- source: https://www.storagereview.com/fr/review/amd-epyc-turin-review-192-cores-of-zen-5?amp -->

AMD's EPYC lineup has long been a staple in the data center sector, delivering unparalleled performance, scalability, and energy efficiency. With the launch of the EPYC Turin 9005 series, powered by the latest Zen 5 architecture, AMD is once again raising the bar for enterprise computing. This new generation is more than just a step forward: it's a giant leap in processing power and efficiency, tailor-made for the ever-increasing demands of modern data centers.
This section summarizes much of the information contained in our previous articles; you can learn more in our announcement articles.
At the heart of the EPYC Turin series is AMD's Zen 5 architecture, unveiled at Tech Day 2024. Designed from the ground up to meet the challenges of the data-driven world, Zen 5 introduces several critical improvements aimed at optimizing workloads in high-performance computing, AI, cloud, and edge environments. EPYC 9005 series processors offer significantly increased core counts and improved threading capabilities, making them a top-tier choice for businesses looking to consolidate their infrastructure while increasing performance.
AMD also emphasizes energy efficiency, a critical factor in today's environmentally conscious data center landscape. Zen 5 chips feature architectural improvements that enhance thermal and power performance, achieved through AMD's ongoing collaboration with TSMC and their cutting-edge manufacturing processes. This allows data centers to maximize performance per watt, helping operators reduce operational costs and achieve sustainability goals without sacrificing throughput.
The Zen 5 EPYC lineup highlights advanced AI capabilities as a flagship feature. The updated math acceleration unit improves performance for machine learning and cryptography tasks, delivering up to 35% improvement in single-core AES-XTS encryption and a 32% increase in single-core machine learning tasks compared to Zen 4. This positions the EPYC Turin series as an essential solution for businesses focused on AI-driven workloads, where speed and efficiency are paramount.
The architecture's improved data bandwidth also plays a crucial role in performance gains. Featuring a redesigned 48KB L1 data cache with twice the bandwidth for the cache and floating-point unit, the EPYC 9005 series ensures that data-hungry applications, such as large-scale simulations and real-time analytics, run more easily and faster than ever. Additionally, the AVX-512 implementation with a full 512-bit data path ensures that floating-point and vector math operations, essential for AI and HPC workloads, are processed with significantly improved efficiency.
For data center operators, the EPYC Turin 9005 series doesn't just represent a performance improvement: it offers a radical shift in how data centers can operate at scale. From its exceptional core density to its advanced AI capabilities and energy efficiency, the EPYC 9005 series is designed to handle the most demanding workloads of today and tomorrow.
AMD's Turin processor is synonymous with performance and adaptability. With configurations up to 128 cores for scalable configurations and an impressive 192 cores for scalable environments, these processors are designed to handle demanding tasks like AI training, simulations, and large databases across various platforms.
With core counts ranging from 8 to 192 and TDPs ranging from 155W to 500W, Turin processors offer versatile power options. Additionally, advanced features such as 12-channel DDR5 memory support up to 6400 MT/s and up to 128 PCIe 5.0 and CXL 2.0 lanes ensure fast and efficient data management, ideal for bandwidth-hungry applications. Processor compatibility with existing SP5 sockets simplifies upgrades, while comprehensive security features (Confidential Compute and Trusted I/O) address the growing need for secure processing. The integration of 12-channel DDR5 support with speeds up to 6400 MT/s is particularly beneficial for bandwidth-intensive workloads.
AMD EPYC processors also perform well in AI-intensive environments thanks to a strategic partnership with NVIDIA. Custom configurations like those on NVIDIA's HGX and MGX platforms demonstrate how AMD EPYC processors work well in configurations requiring powerful GPU support. High-frequency options like the EPYC 9575F, which reaches up to 5 GHz, are specifically designed for applications that demand low latency and fast processing times. This positions AMD well in real-time AI processing and allows users to adapt systems from single-threaded tasks to large-scale parallel workloads while carefully balancing power and efficiency.
AMD claims its 5th Gen EPYC processors outperform Intel's offerings in AI-specific tests, particularly in machine learning and large language modeling tasks. According to AMD's tests, these processors enjoy an advantage of up to 3.8x, making them particularly suited for real-time AI applications where responsiveness and low latency are crucial. AMD's focus on optimizing AI inference enables faster decision-making in data-hungry applications like recommendation systems and similarity searches.
The AMD EPYC Turin processor lineup offers remarkable diversity suited to the many requirements of enterprises. From high-core processors designed for data-hungry environments to low-core, high-frequency options ideal for specialized tasks, this series offers scalable performance and efficiency options tailored to all workloads.
| Colorful | Model/CCD | Base/Boost | TDP | L3 Cache (MB) | Price (1 KU, USD) | 
|---|---|---|---|---|---|
| 192 Cores | 9965 "Zen5c" | 2.25/3.7 | 500W | 384 | $14,813 | 
| 160 Cores | 9845 "Zen5c" | 2.1/3.7 | 390W | 320 | $13,564 | 
| 144 Cores | 9825 "Zen5c" | 2.2/3.7 | 390W | 384 | $13,006 | 
| 128 Cores | 9755 "Zen5" | 2.7/4.1 | 500W | 512 | $12,984 | 
|  | 9745 "Zen5c" | 2.4/3.7 | 400W | 256 | $12,141 | 
| 96 Cores | 9655 "Zen5" | 2.6/4.5 | 400W | 384 | $11,852 | 
|  | 9655P "Zen5" | 2.6/4.5 | 400W | 384 | $10,811 | 
|  | 9645 "Zen5c" | 2.3/3.7 | 320W | 256 | $11,048 | 
| 72 Cores | 9565 "Zen5" | 3.15/4.3 | 400W | 384 | $10,486 | 
| 64 Cores | 9575F "Zen5" | 3.3/5.0 | 400W | 256 | $11,791 | 
|  | 9555 "Zen5" | 3.2/4.4 | 360W | 256 | $9,826 | 
|  | 9555P "Zen5" | 3.2/4.4 | 360W | 256 | $7,983 | 
|  | 9535 "Zen5" | 2.4/4.3 | 300W | 256 | $8,992 | 
| 48 Cores | 9475F "Zen5" | 3.65/4.8 | 400W | 256 | $7,592 | 
|  | 9455 "Zen5" | 3.15/4.4 | 300W | 192 | $5,412 | 
|  | 9455P "Zen5" | 3.15/4.4 | 300W | 192 | $4,819 | 
| 36 Cores | 9365 "Zen5" | 3.4/4.3 | 300W | 192 | $4,341 | 
| 32 Cores | 9375F "Zen5" | 3.8/4.8 | 320W | 256 | $5,306 | 
|  | 9355 "Zen5" | 3.55/4.4 | 280W | 256 | $3,694 | 
|  | 9355P "Zen5" | 3.55/4.4 | 280W | 256 | $2,998 | 
|  | 9335 "Zen5" | 3.0/4.4 | 210W | 128 | $3,178 | 
| 24 Cores | 9275F "Zen5" | 4.1/4.8 | 320W | 256 | $3,439 | 
|  | 9255 "Zen5" | 3.25/4.3 | 200W | 128 | $2,495 | 
| 16 Cores | 9175F "Zen5" | 4.2/5.0 | 320W | 512 | $4,256 | 
|  | 9135 "Zen5" | 3.65/4.3 | 200W | 64 | $1,214 | 
|  | 9115 "Zen5" | 2.6/4.1 | 125W | 64 | $726 | 
| 8 Cores | 9015 "Zen5" | 3.6/4.1 | 125W | 64 | $527 | 
At the top of AMD's lineup sits the EPYC 9965, the flagship of the Turin family, featuring 192 cores and 384 threads. It operates at a base frequency of 2.25 GHz and can reach 3.7 GHz in Turbo mode, while displaying a substantial TDP of 500W. With 384 MB of L3 cache, this processor is designed for large-scale datacenters and enterprise environments where massively parallel processing is essential. Priced at $14,813, this model targets organizations requiring high throughput and scalable power for compute-intensive tasks.
The EPYC 9845 is another example of a multi-core processor offering 160 cores and 320 threads, clocked at 2.1 GHz in Turbo mode and up to 3.7 GHz in Turbo mode. With a TDP of 390W and 320 MB of L3 cache, this model is optimized for intensive multithreaded workloads, while consuming slightly less power than the 9965.
Lower in the lineup, the EPYC 9755 is a versatile 128-core, 256-thread processor, clocked at 2.7 GHz in Turbo mode and capable of reaching 4.1 GHz. Designed to deliver an optimal balance between performance and energy efficiency, it features a TDP of 500W. With 512 MB of L3 cache, it is ideal for mixed use, from general-purpose applications to data-hungry applications, making it a preferred choice for virtualization, cloud environments, and complex data analytics.
The EPYC 9655 processor (96 cores, 192 threads) is another notable model in this category. It operates at a base frequency of 2.6 GHz with a turbo frequency of 4.5 GHz, and its power consumption is limited to 400W. With 384 MB of L3 cache, it is perfectly suited to environments requiring a large number of cores and frequent memory accesses, such as large-scale virtualization and containerized applications.
In the low-core processor range, models like the EPYC 9575F and EPYC 9555 offer high single-core performance, with clock speeds reaching up to 5.0 GHz for the 9575F. The latter, featuring 64 cores and 128 threads, operates at a base frequency of 3.3 GHz.
The EPYC 9475F is another interesting model. This processor has 48 cores and 96 threads, a base frequency of 3.65 GHz, and a turbo frequency reaching up to 4.8 GHz. Priced at $7,592, it offers a good compromise between frequency and core count.
AMD offers options like the EPYC 9375F, featuring 32 cores and 64 threads, for less demanding workloads or entry-level deployments. This model is priced at $5,306 and operates at a base frequency of 3.8 GHz, reaching up to 4.8 GHz, with a TDP of 320W. With 256 MB of L3 cache, this processor offers more than sufficient performance in scenarios where high frequency and a moderate core count can improve daily operational efficiency.
Finally, the EPYC 9175F processor, featuring 16 cores and 32 threads and a high turbo frequency of 5 GHz, is ideal for entry-level enterprise environments. With a TDP of 320W and a price of $4,256, this model is perfectly suited for organizations with simpler and less demanding workloads.
For this launch, AMD provided us with its Volcano reference platform and a few CPU sets to test. The reference system, which appears to be based on a Lenovo platform, was equipped with 1.5 TB of DDR5 running at 6000 MT/s and was cooled using an AIO liquid loop. However, while this system is well suited for general testing, it only had 3 xGMI links between the CPU and memory controllers.
To expand our testing, we used pre-production versions of the new Dell PowerEdge R7725 and R6725 platforms, Dell's latest AMD-based 2U and 1U servers, respectively. These platforms proved particularly useful for our testing, as they support 4 xGMI links, offering superior bandwidth and efficiency for high-intensity workloads compared to the reference design.
Both Dell platforms were equipped with 1.5 TB of DDR5 at 6000 MT/s, cooled by platinum-grade fans capable of handling 500W processors on air. During lab testing, the air cooling of these systems worked exceptionally well, keeping processors at safe operating temperatures. However, to account for ongoing cooling system renovations in our lab, we used a CoolIT SP5 cold plate kit connected to a CDU to ensure stable temperatures and lower noise levels during prolonged workloads. This configuration allowed us to push the Dell servers to their limits. At the same time, the liquid loop solution primarily addressed our unique environmental constraints rather than the performance limitations of the platforms themselves. All tests were performed with Noctua NT-H2 thermal paste and the CollIT SP5 direct-on-chip liquid cooling solution.
We will now examine the performance of AMD's latest EPYC Turin processors, comparing their capabilities across a series of highly demanding tests with the Genoa and Bergamo models. These tests cover a range of scenarios, from AI inference and 3D rendering to compute workloads and data compression, aiming to evaluate multithread efficiency, core scalability, and single-thread power.
With configurations ranging from the massive 192-core EPYC 9965 to the streamlined 64-core EPYC 9575F, we performed all tests with SMT disabled, the best performance profiles in iDRAC, and power determinism to allow the processors to run at their maximum. This will allow us to provide an overview of how each model balances core count, frequency, and efficiency to meet various workloads.
The cross-platform Geekbench 6 benchmark measures a system's performance and provides a comparison score. It is designed to run on multiple platforms and provides a consistent measure of performance across many devices, from smartphones and tablets to desktops and servers.
Here, the 192 cores of the EPYC 9965 display a multi-core score of 11,199, which allows for considerable processing power, but remains behind dual-socket Genoa (20,217) and Bergamo (17,916) configurations. This suggests that while the 9965 is a high-performance model, Genoa and Bergamo remain optimized for the most extreme parallel processing workloads. The 128 cores of the EPYC 9755 demonstrate efficiency gains, slightly surpassing the 9965 in single-core (1,641) and multi-core (11,800) scores, indicating that it is perfectly suited for scalable and multithreaded tasks without the overhead of a larger core count.
The 9575F (with SMT disabled) achieves a solid single-core score of 1,865, showing its potential for applications that benefit from high single-thread performance. In contrast, its multi-core score of 13,219 for 64 cores, while modest, makes it more suited to specialized workloads rather than the heaviest parallel processing tasks.
| Geekbench 6 | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| Single Core | 1,453 | 1,641 | 1,865 | 2,048 | 1,723 | 
| Multi-Core | 11,199 | 11,800 | 13,219 | 20,217 | 17,916 | 
Cinebench is a widely used benchmarking tool that measures the performance of CPUs and GPUs using Maxon Cinema 4D for rendering. It provides a score to compare the performance of different systems and components. We ran four popular versions of Cinebench so you can compare results with popular online rankings.
Cinebench R23 results show that the EPYC 9755 achieves a high multi-core score of 131,846, surpassing Genoa's 116,744 and Bergamo's 102,125, positioning it as an ideal choice for intensive rendering tasks in highly parallel environments. Genoa leads with 1,294 points in single-core performance, giving it a slight advantage in applications that prioritize faster single-threaded operations.
In the updated Cinebench 2024 tests, the 9965 and 9755 continue to perform well with 4,845 and 5,921 points respectively, offering solid multithreaded scores suited to tasks such as complex 3D rendering and content creation workflows.
| Test | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) |  | 
| Cinebench R20 |  |  |  |  |  |  | 
| CPU | N/A | 43,765 pts | N/A | N/A | N/A |  | 
| Cinebench R23 |  |  |  |  |  |  | 
| CPU (multi-core) | N/A | 131,846 pts | 111,149 pts | 116,744 pts | 102,125 pts |  | 
| CPU (single-core) | N/A | 1,400 pts | 1,052 pts | 1,294 pts | 1,089 pts |  | 
| Cinebench 2024 |  |  |  |  |  |  | 
| CPU (multi-core) | 4,845 pts | 5,921 pts | 4,324 | N/A | N/A |  | 
| CPU (single-core) | 77 pts | 84 pts | 103 pts | N/A | N/A |  | 
y-cruncher 0.8.3.9522 is a multithreaded and scalable program that can calculate Pi and other mathematical constants to billions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application among overclockers and hardware enthusiasts.
Here, the EPYC 9575F (SMT disabled) achieves the fastest computation times for 1 and 2.5 billion digit calculations, at 4.476 and 10.067 seconds respectively. This speed even surpasses the high-core 9965 and 9755 models, demonstrating its efficiency for high-precision tasks that do not require an extreme number of threads.
The 9965 and 9755 processors also offer good times, notably surpassing those of Genoa and Bergamo at 10 billion digits, in just over 41 seconds versus 51 seconds for Genoa. This shows that Turin-based EPYC processors are well optimized for compute-intensive applications, demonstrating scalable performance that allows for more precise control of thread and resource management. The 64-core 9575F processor really shone here with its incredibly high clock speed, showing a strong lead over the cores of the 9755 or 9965 samples in the y-cruncher ranges we examined.
| y-cruncher Total Computation Time (Lower is Better) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, 128t) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| 1 billion | 7.346 seconds | 7.747 seconds | 5.408 seconds | 4.476 seconds | 8.882 seconds | 9.184 seconds | 
| 2.5 billion | 13.661 seconds | 14.113 seconds | 11.376 seconds | 10.067 seconds | N/A | N/A | 
| 5 billion | 23.211 seconds | 22.820 seconds | 20.177 seconds | 20.030 seconds | N/A | N/A | 
| 10 billion | 41.750 seconds | 41.512 seconds | 40.767 seconds | 41.518 seconds | 51.071 seconds | 55.683 seconds | 
| 25 billion | 115.091 seconds | 98.981 seconds | 103.650 seconds | 104.737 seconds | N/A | N/A | 
This Y-Cruncher performance test uses the Bailey-Borwein-Plouffe (BBP) formulas to calculate massive hexadecimal digits of Pi, measuring the CPU's total computation time, utilization, and multi-core efficiency. The results provide a detailed analysis of the performance of AMD Turin processors in these intensive calculations.
The EPYC 9575F with SMT disabled excels in single-thread efficiency, completing 1 BBP in just 0.179 seconds with a multi-core efficiency of 44.73%. This demonstrates its efficiency in scenarios where single-thread optimization is a priority. By comparison, the 9965, with its 192 cores, shows a lower multi-core efficiency of 4.61% in this single-BBP test, as its high core count introduces latency in low-thread scenarios.
Multi-core efficiency becomes more evident with higher workloads, such as 10 and 100 BBP. The EPYC 9575F with SMT disabled maintains a strong lead with 86.59% efficiency at 10 BBP, completing it in 0.896 seconds. For 100 BBP, the 9575F continues to excel, reaching nearly 99% efficiency and displaying balanced performance in highly parallel workloads, completing the run in 8.065 seconds. In this test, the 9965 achieves 85.10% efficiency, while the 9755 and 9575 maintain competitive multi-core efficiencies of 94% and 96.83%, respectively, demonstrating that Turin-based EPYC processors can efficiently handle both low-thread and high-thread workloads.
| Benchmark | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | AMD EPYC9575F (64c, SMT disabled) | 
| 1 BBP |  |  |  |  | 
| 10 BBP |  |  |  |  | 
| 100 BBP |  |  |  |  | 
The memory test built into the 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, indicating how well the system can handle data-hungry operations. We run this test with a dictionary size of 128 MB when possible.
The 7-Zip compression test demonstrates the efficiency of the 9755, achieving 443.029 GIPS at 5613% CPU utilization, the highest in the EPYC lineup. This level of performance makes the 9755 ideal for intensive data compression tasks, such as those in backup or archiving solutions, where fast data handling is essential. The 9575F (SMT disabled) also performs well, achieving 394.9 GIPS. These results indicate that the 9755 and 9575F have solid compression and decompression capabilities, with high CPU throughput and efficient resource management for applications that leverage high compression ratios or process large volumes of data.
| 7-Zip Compression Benchmark (Higher is Better) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c, SMT disabled) |  | 
| Compression |  |  |  |  | 
| Current CPU Usage | 4302% | 5233% | 4406% |  | 
| Current Rating/Usage | 5.830 GIPS | 7.597 GIPS | 7.975 GIPS |  | 
| Current | 250.827 GIPS | 397.536 GIPS | 351.358 GIPS |  | 
| Resulting CPU Usage | 4041% | 5306% | 4555% |  | 
| Resulting Rating/Usage | 5.804 GIPS | 7.720 GIPS | 8.070 GIPS |  | 
| Resulting Rating | 234.317 GIPS | 409.652 GIPS | 367.358 GIPS |  | 
| Decompression |  |  |  |  | 
| Current CPU Usage | 4322% | 6041% | 5017% |  | 
| Current Rating/Usage | 7.078 GIPS | 8.065 GIPS | 8.483 GIPS |  | 
| Current | 305.909 GIPS | 487.263 GIPS | 425.580 GIPS |  | 
| Resulting CPU Usage | 4556% | 5921% | 4940% |  | 
| Resulting Rating/Usage | 6.577 GIPS | 8.045 GIPS | 8.569 GIPS |  | 
| Resulting Rating | 299.163 GIPS | 476.405 GIPS | 422.441 GIPS |  | 
| Total Rating |  |  |  |  | 
| Total CPU Usage | 4298% | 5613% | 4747% |  | 
| Total Rating/Usage | 6.190 GIPS | 7.883 GIPS | 8.319 GIPS |  | 
| Total Rating | 266.740 GIPS | 443.029 GIPS | 394.900 GIPS |  | 
Blender OptiX is an open-source 3D modeling application. This benchmark was run using the Blender Benchmark CLI utility. The score is expressed in samples per minute, with the highest being the best.
Here, the 9755 slightly outperforms with a score of 2,606.54 samples per minute in the "Monster" scene, closely followed by the 9965 with 2,558.43. Both surpass Genoa and Bergamo, with respective scores of 1,700.65 and 2,038.71, highlighting the advantage that Turin-based processors bring to 3D rendering. These scores highlight the suitability of the 9755 and 9965 for high-resolution 3D modeling tasks and content creation workflows involving complex scene rendering.
| Blender 4.0 CPU Samples per Minute (Higher is Better) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| Monster | 2,558.43 | 2,606.54 | 1,196.15 | 1,700.65 | 2,038.71 | 
| Junkshop | 1,866.65 | 1,843.48 | 802.00 | 1,101.84 | 1,382.58 | 
| Classroom | 1,270.17 | 1,251.54 | 637.13 | 869.48 | 1,045.96 | 
UL's Procyon AI inference test suite evaluates the performance of different AI inference engines using state-of-the-art neural networks. These tests were run on the CPU only. Each value represents an average inference time; the lower the value, the better the performance. The last line indicates an overall score; the higher the value, the better the performance.
Here, the EPYC 9575F with SMT disabled delivers solid inference times across various neural networks, with exceptional performance in MobileNet V3 (7.02 ms) and ResNet 50 (10.45 ms). These scores surpass both Genoa (3.63 ms for MobileNet and 6.34 ms for ResNet) and Bergamo (4.16 ms and 8.22 ms, respectively) for AI workloads, suggesting that the SMT-off configuration of the 9575F prioritizes latency control, making it highly effective for real-time AI tasks where low inference times are essential.
The EPYC 9965 and 9755, while slower in single inference tasks than Genoa, maintain solid performance across the board, with the 9755's 20.07 ms on MobileNet and 20.11 ms on ResNet 50 indicating stable and reliable performance suited to moderate workloads.
While Genoa and Bergamo offer consistently lower latency, the 9575F, especially with SMT disabled, offers excellent latency control.
| UL Procyon Average | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) |  | 
| Inference Time (Lower is Better) |  |  |  |  |  |  | 
| Mobile Net V3 | 32.86 ms | 20.07 ms | 7.02 ms | 3.63 ms | 4.16 ms |  | 
| ResNet 50 | 38.63 ms | 20.11 ms | 10.45 ms | 6.34 ms | 8.22 ms |  | 
| Creation V4 | 116.38 ms | 66.22 ms | 33.69 ms | 25.99 ms | 30.68 ms |  | 
| Deep Lab V3 | 58.16 ms | 33.02 ms | 19.38 ms | 25.33 ms | 30.57 ms |  | 
| YOLO V3 | 84.20 ms | 38.47 ms | 24.54 ms | 34.13 ms | 41.38 ms |  | 
| REAL-ESRGAN | 2923.75 ms | 1600.73 ms | 1219.26 ms | 2524.03 ms | 2301.35 ms |  | 
| Overall Score (Higher is Better) | 44 | 81 | 148 | N/A | N/A |  | 
The Blackmagic RAW Speed Test performance benchmarking tool measures a system's ability to handle video playback and editing using the Blackmagic RAW codec. It evaluates a system's ability to decode and play high-resolution video files, providing frame rates for both CPU-based and GPU-based processing.
In this test, the EPYC 9755 leads with an impressive 174 fps in 8K CPU decoding, ideal for high-resolution video playback in media production environments. The 9575F (SMT disabled) follows closely with 154 fps, slightly surpassing the 9965 at 134 fps. These results suggest that the 9755 and 9575F offer the best balance between frame rate consistency and decoding speed, crucial for video editing and production workflows that handle large, high-quality video files.
| Blackmagic RAW (Higher is Better) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | 
| CPU 8K | 134 fps | 174 fps | 154 fps | 
The AMD EPYC Turin 9005 series marks a significant advancement in enterprise computing, offering exceptional performance, efficiency, and adaptability across various workloads. Powered by the Zen 5 architecture, these processors are specifically designed to meet the growing demands of AI, cloud, and HPC environments while delivering unmatched scalability for data centers looking to optimize performance and power consumption.
The EPYC 9005 series stands out for its ability to meet the diverse needs of modern enterprises with configurations ranging from entry-level high-frequency models to very high cache density chips to multi-core powerhouses. Whether it's real-time AI inference, computational fluid dynamics, large-scale data analytics, or high-resolution 3D rendering, the EPYC lineup offers single-thread responsiveness and multithread efficiency. Advanced features such as 12-channel DDR5 memory support, PCIe 5.0 lanes, and AMD's secure and confidential computing make this series both a performance upgrade and a comprehensive solution for forward-looking data centers.
Our testing with the Dell PowerEdge R7725 and R6725 platforms highlighted the true potential of these processors, particularly thanks to their board design and cooling solution, which offer a clear advantage in terms of bandwidth and overall performance. The combination of AMD's exceptional Zen 5 architecture and Dell's robust server design creates a compelling platform for businesses looking to push the boundaries of computing power.
The ASUS ExpertCenter Pro ET900N G3 is the second GB300 DGX station we have tested and the first we have been able to handle…
In the field of enterprise computing, most security upgrades generally occur in response to data breaches or other major incidents. The recently developed INCITS system…
Mini-PCs have become one of the most important segments of the client computing market. They form the basis of streamlined IT fleets…
The HP ZBook Ultra G1a 14 takes the top spot in our ranking of the best laptops for local AI, among the large models…
The most frequently asked question we received regarding the MSI XpertStation WS300 after our test revolves around a key theme. The…
Eaton is about to bring to market a type of power distribution unit (PDU) that shouldn't have existed five years ago. The HDXL…
