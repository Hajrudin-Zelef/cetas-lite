---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-4
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: storagereview
role: reference
task: reference
actors: ["Hugging Face", "Intel", "Nvidia", "Samsung"]
dates: []
keywords: ["benchmark", "benchmarks", "compute", "gpu", "gpus", "inference", "intel", "latency", "memory", "nvidia", "robotics", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [41, 73]
sha256: bb042c95fb6533f321f43e9b01718f6e187a8968f841f43cb8f01e9f2b1c4565
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

This deep integration standardizes interfaces and form factors across subsystems. While Dell emphasizes the importance of using validated components to ensure compatibility and support, the underlying standardization makes many components inherently easier to maintain and potentially interchangeable between compatible systems in the future.
Management and iDRAC
The Dell PowerEdge R770 server builds on the already feature-rich and highly regarded iDRAC 9 with the new-generation iDRAC 10. The latter enhances system administration through seamless integration with the Data Center Secure Control Module (DC-SCM). This integration simplifies firmware updates and configuration management, ensuring consistent and scalable operations across data centers. iDRAC 10 also supports advanced automation and monitoring features, enabling IT administrators to efficiently manage large-scale deployments without compromising performance or reliability.
Security is a cornerstone of the R770's management features, with Dell implementing robust pre-boot and post-boot verification mechanisms. Leveraging silicon-level Root of Trust technology, iDRAC 10 ensures cryptographic verification of all firmware, including BIOS and iDRAC, before execution. This immutable hardware security measure protects against malware tampering and supply chain attacks, providing a secure foundation for system operation. Additionally, the R770 incorporates quantum-resistant boot protocols to mitigate emerging cryptographic threats, reinforcing its role in protecting critical infrastructure.
Dell's commitment to supply chain security is reflected in the R770's design, which relies on a comprehensive chain-of-trust authentication process. Every hardware component undergoes rigorous verification through cryptographic signatures embedded during manufacturing. This process ensures that only authorized firmware and components are used, reducing the risks associated with unauthorized modifications or counterfeits.
Building blocks of AI factories
The R770 is available with numerous GPU and chassis configurations, making it a versatile platform for a wide range of AI workloads. Its flexibility, combined with its robust storage and networking capabilities, makes it an attractive choice for organizations deploying AI solutions in AI Factories. AI Factories refer to the infrastructure and tools needed to create, train, and deploy AI models at scale. These factories are essential for developing advanced systems such as autonomous vehicles and robotics, as they provide the compute power and data pipelines necessary to efficiently process enormous volumes of data.
The development of autonomous vehicles and robotic systems requires comprehensive training data reflecting real-world scenarios. NVIDIA's Cosmos NIM represents a major advancement in this area, offering developers a powerful toolkit to accelerate the creation and deployment of physical AI systems such as world foundation models.
Understanding world foundation models
World Foundation Models (WFMs) are sophisticated neural networks that simulate real-world environments and predict accurate outcomes from various inputs. Unlike traditional AI models focused on specific tasks, WFMs understand the dynamics of the physical world, including the laws of physics and spatial properties. They can generate video from text instructions, images, or other inputs, while faithfully representing motion, forces, and spatial relationships.
NVIDIA Cosmos NIM: a springboard to world foundation models
NVIDIA's Cosmos NIM modules represent a crucial step toward achieving world foundation models. They enable organizations and AI labs to generate synthetic training data, effectively adapting the amount of data needed to train these AI models. We deployed the Cosmos Predict model, a generalist model that generates world states from text or video instructions and synthesizes continuous motion by predicting key frames.
Here are some interesting results we obtained with Cosmos from a single image of our lab. While not perfect, the results obtained from a single image are very impressive.
The R770's ability to support high-performance GPUs, such as the NVIDIA H100, along with its robust storage and networking capabilities, makes it an attractive choice for organizations looking to deploy AI solutions.
By leveraging the R770's capabilities, enterprises can efficiently train and deploy AI models like Cosmos NIM, accelerating the development of autonomous vehicles and robotic systems. Its performance and scalability make it an ideal platform for handling the large amounts of data required to train AI models, and its versatility allows it to support a wide range of AI workloads.
GPU Direct Storage
GPU Direct Storage is a technology that enables direct data transfer between storage devices and GPUs, bypassing the CPU and system memory. In a traditional data transfer, data is read from storage into CPU memory, then copied to GPU memory. This process involves multiple data copies, resulting in increased latency and reduced performance. The CPU acts as a bottleneck, having to manage data transfer between storage and the GPU. GDS eliminates this bottleneck by allowing storage devices to transfer data directly to and from GPU memory.
We performed a GDSIO workload analysis on a storage system composed of 16 drives, progressively increasing the number of drives used to understand storage performance and its ability to saturate a PCIe Gen 5 GPU.
The GDSIO read graph illustrates the impact of increasing the number of KIOXIA CD8P SSDs on the R770's overall and average read throughput. Initially, as the number of drives goes from one to four, overall read throughput increases rapidly, reaching approximately 50.2 GiB/s. This suggests that the system can saturate PCIe Gen 5 x16 with only three or four drives for data loading. Beyond five drives, overall throughput plateaus, indicating that adding more drives does not significantly improve it. Meanwhile, average read throughput per drive remains stable up to four drives, then decreases as more drives are added. This decline in per-drive performance is explained by more drives sharing the available PCIe bus bandwidth, reducing individual reads.
In contrast, the write performance of these drives is well below their read performance. It took all 16 drives to reach a write bandwidth of 46.7 GiB/s, with average write speed remaining nearly constant. Given the lower write capabilities of the KIOXIA CD8 range, higher-capacity versions or other PCIe Gen5 SSDs will perform better.
Comparative analysis of the Dell PowerEdge R770
Regarding benchmarks, the R770 is Dell's flagship system and, as such, will be deployed in highly varied environments. We therefore conducted a comprehensive series of benchmarks for this platform to evaluate its performance in different environments. The Lenovo ThinkSystem SR630 V4 was compared in some tests to highlight the difference between high-end multi-core and many-core processors.
System configuration
- CPU: 2x Intel Xeon 6787P (86 cores each)
- RAM: 32x Micron 64 GB DDR5 6400 MT/s dual-rank memory Total memory: 2 TB
- Power supplies: 2x Delta 1500W
- GPU: 1x NVIDIA H100 for the TGI benchmark, 1x NVIDIA L4 for the remaining tests
- NIC: DELL BRCM 4P 25G SFP 57504S OCP network card
- BOSS card: BOSS-N1 DC-MHS 0 and 1 SK hynix 480 GB Dell NVMe ISE PE9010 RI M.2 480 GB drives
- Drives: 0-5 in backplane 1: Samsung 6.4 TB, Dell NVMe PM1745 MU E3.S 6.4 TB
AI workload performance
Text generation inference benchmark
Text Generation Inference (TGI) is a high-performance LLM inference server developed by Hugging Face. Designed to optimize the deployment and utilization of LLMs, it is an ideal choice for production environments. TGI supports various open-source LLMs and offers features such as tensor parallelism, token streaming, and continuous batching, which enhance its performance and efficiency.
