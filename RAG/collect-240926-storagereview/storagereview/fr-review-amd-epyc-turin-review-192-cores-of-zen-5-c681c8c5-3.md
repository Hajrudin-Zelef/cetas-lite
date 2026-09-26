---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5-3
title: "fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "compute", "energy", "gpus", "inference", "liquid cooling", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5.md
source_anchor: ""
source_lines: [46, 75]
sha256: fc68a996d0f23722e48807ac3da2382c56f9715cd2286b1b01686c2919a011d3
---

# fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5

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
