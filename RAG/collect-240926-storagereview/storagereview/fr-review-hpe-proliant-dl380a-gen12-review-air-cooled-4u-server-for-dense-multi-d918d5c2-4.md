---
id: collect-240926-storagereview/storagereview/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi-d918d5c2-4
title: "fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["apache", "benchmark", "benchmarks", "compute", "gpus", "inference", "intel", "latency", "memory", "nvfp4", "nvidia", "packaging"]
source: docs/RAG/clean_en/storagereview/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2.md
source_anchor: ""
source_lines: [74, 101]
sha256: 8427886d0026edf693eeb14e79ac82657e6f0fc89c7e1a6d8f66baa4b349def1
---

# fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2

With a base station (BS) of 1 and a total throughput (TP) of 4, the model achieves a throughput of 20.59 tok/s per user and a TPOT of 38.27 ms. With a BS of 16, performance increases to 7.20 tok/s per user, for a total throughput of 806.14 tok/s and a TPOT of 54.98 ms. The maximum total throughput of 1,372.21 tok/s is reached with a BS of 128, or 2.59 tok/s per user and a TPOT of 122.75 ms.
Microscaling Data Type Performance
Microscaling represents an advanced quantization approach that applies precise scaling factors to small blocks of weights, rather than uniform quantization across large groups of parameters. NVIDIA's NVFP4 format implements this technique through a block floating-point representation, where each microscaling block of 8 to 32 values shares a common exponent serving as a scaling factor. This granular approach preserves numerical precision while ensuring a 4-bit representation, thus maintaining the dynamic range essential to transformer architectures. This format integrates with NVIDIA's Tensor Core architecture on the RTX PRO 6000, enabling efficient mixed-precision computation with on-the-fly decompression during matrix operations.
GPT-OSS-120B Performance
We evaluated OpenAI's GPT-OSS-120B model with NVFP4 quantization. In single-user mode (TP=2), the model achieves 176.09 tok/s per user with a TPOT of 5.46 ms, the lowest latency in our test suite. With BS=4 and TP=4, performance reaches 105.79 tok/s per user, for a total throughput of 1155.94 tok/s and a TPOT of 7.79 ms. With BS=32 and TP=4, throughput increases to 47.54 tok/s per user and 3956.44 tok/s total, with a TPOT of 13.86 ms. The maximum total throughput of 4015.77 tok/s is reached with BS=64, or 25.38 tok/s per user and a TPOT of 14.78 ms.
Phoronix Benchmarks
Phoronix Test Suite is an open-source automated benchmarking platform supporting over 450 test profiles and more than 100 test suites via OpenBenchmarking.org. It handles the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration. We will focus on the following tests: Stream, 7-Zip, Linux kernel compilation, Apache, and OpenSSL.
Stream Memory Bandwidth
In the Stream benchmark, which measures raw memory throughput, the HPE DL380a Gen12 achieved an impressive score of 542 GB/s, demonstrating the platform's ability to maintain high data transfer rates under continuous load. This level of bandwidth makes the system particularly performant for workloads such as data modeling, simulation, and AI inference, where large datasets must be transferred quickly between memory and compute resources.
7-Zip Compression
The 7-Zip compression test measured 305,000 MIP, highlighting the system's excellent multithreaded efficiency for compute-intensive compression and decompression operations. These results make the DL380a Gen12 an ideal choice for environments involving frequent data packaging, archiving, or backup operations, which require consistent and reproducible CPU performance.
Kernel Compilation
When compiling a full Linux kernel (allmodconfig), the DL380a Gen12 completed the operation in 316 seconds. This result demonstrates the system's ability to easily handle complex, parallelized workloads. Increased compilation performance translates directly into shorter build times and improved iteration speed for developers working on large-scale software projects or in CI/CD environments.
Apache Web Server
In terms of web server performance, the DL380a Gen12 sustained 94,348 requests per second in the Apache performance test. This result demonstrates balanced I/O handling and high cache efficiency, providing the throughput and responsiveness needed for enterprise web applications, virtualization interfaces, or internal service hosting.
OpenSSL Verification
Cryptographic performance was equally remarkable, with the DL380a Gen12 verifying 803 billion operations per second under OpenSSL. This demonstrates the system's ability to handle encryption, authentication, and secure communications workloads at scale.
| Phoronix Benchmarks | HPE ProLiant DL380a Gen 12 (2x Intel Xeon 6527P) | 
| Discussions | 542,720.7 MB / s | 
| 7-ZIP | 304,907 MIP/s | 
| Kernel Compilation (allmod) | 316.166 seconds | 
| Apache (requests per second) | 94,347.52 R/s | 
| OpenSSL | 803,597,895,087 Verifications | 
Conclusion
The HPE ProLiant DL380a Gen12 server stands out as one of the most practical and balanced AI servers for the enterprise AI market. Its 4U air-cooled form factor offers exceptional compute density thanks to its two Xeon 6 processors, support for up to 8 double-width GPUs or 16 single-width GPUs, and 16 E3.S NVMe bays, while ensuring reliability and ease of maintenance. HPE's engineering approach to airflow and thermal balance ensures consistent performance even under heavy workloads, demonstrating that advanced AI acceleration can work perfectly in traditional air-cooled environments.
The integration of iLO 7 significantly improves management, a major asset for leading server vendors like HPE. The modernized interface, integration with HPE Compute Ops Management, and detailed hardware telemetry make remote administration intuitive and efficient. Each section (Dashboard, Firmware, Host, Security, Applications, and Settings) demonstrates HPE's commitment to delivering a smoother, more cloud-integrated experience, without sacrificing the on-premises control essential to enterprise teams.
In performance testing, the server achieved excellent results. The four RTX PRO 6000 GPUs delivered impressive throughput on dense and microscaling LLM models, with vLLM serving performance rivaling that of liquid-cooled systems. Phoronix CPU benchmarks also highlight its balance, with memory bandwidth over 540 GB/s, 94,000 requests per second under Apache, and over 800 billion OpenSSL verifications per second, demonstrating its robustness for both AI and general compute.
HPE's design goal is clear: to deliver high-density, production-ready AI performance through air cooling compatible with existing rack and power infrastructure. For data center teams looking for a reliable, secure, and easy-to-manage air-cooled compute solution, the DL380a Gen12 is a performant and innovative solution for the rapidly growing mainstream AI market.
