---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5-2
title: "fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "TSMC"]
dates: []
keywords: ["amd", "compute", "energy", "gpu", "inference", "intel", "latency", "memory", "nvidia", "throughput", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5.md
source_anchor: ""
source_lines: [3, 45]
sha256: 4b3a44a169f02890c75484717580150ddb5967ab7e01e7e1f0a4a90fb9a15cda
---

# fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5

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
