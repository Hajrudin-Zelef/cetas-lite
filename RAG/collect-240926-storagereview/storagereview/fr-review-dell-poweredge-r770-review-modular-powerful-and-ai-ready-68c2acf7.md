---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: storagereview
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Intel", "Microsoft", "Nvidia", "Samsung", "TensorRT-LLM"]
dates: []
keywords: ["accelerator", "attention", "benchmark", "benchmarks", "compute", "cost", "datacenter", "decode", "deepseek", "distillation", "distribution", "energy"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [1, 219]
sha256: a7ca5535b923238924e272e6e6b70204296999b7a65910b0438b77ec8c391d4e
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

<!-- source: https://www.storagereview.com/fr/review/dell-poweredge-r770-review-modular-powerful-and-ai-ready -->

Dell PowerEdge R7x0 series servers have long been a cornerstone of data centers, renowned for their exceptional build quality, thoughtful design, performance, density, and reliability, all in a versatile 2U form factor. These servers have constantly evolved to meet changing requirements. Today, with the introduction of the Dell PowerEdge R770, the series takes a decisive step forward.
The R770 debuts the new Intel Xeon 6 processor family, featuring Xeon 6500 and 6700 P-core and E-core processors. It marks Dell's first full adoption of the OCP Data Center Modular Hardware System (DC MHS) standard in its mainstream server lineup. Together, these two developments promise a significant leap forward in both performance and design.
Meeting the demands of modern data centers
The launch of the R770 comes as data centers face increasing pressure. Workloads are becoming increasingly diverse and demanding. The relentless growth of data reinforces the need for robust analytics and databases. From training complex models to deploying real-time inference, artificial intelligence is no longer a niche application but a core business engine requiring substantial compute power and specialized acceleration.
At the same time, energy efficiency and total cost of ownership optimization are receiving heightened attention. Additionally, the industry is increasingly turning to open standards to foster innovation, improve interoperability, and potentially reduce vendor lock-in. The R770, with its new processor options and adoption of OCP DC MHS, is designed to meet these challenges.
Intel Xeon 6 P-Core processors
The R770 utilizes Intel Xeon 6 series processors, specifically the 6700 and 6500 series, featuring Performance and Efficiency cores based on the Socket E2 platform (LGA4710-2). In this analysis, we focus specifically on the P-series SKUs.
Intel builds these processors on a tile-based design, combining I/O tiles with one or two compute tiles. This enables scalability within the series, with configurations up to 86 P-cores (XCC) with two compute tiles, and up to 48 P-cores (HCC) or 16 P-cores (LCC) with a single compute tile.
Compared to previous-generation Sapphire and Emerald Rapids processors, these processors stand out for the universal availability of built-in accelerators across all Xeon 6 processors. These include Intel QuickAssist Technology for encryption and compression, the Intel Data Streaming Accelerator for data movement, the Intel In-Memory Analytics Accelerator for database and analytics acceleration, and the Intel Dynamic Load Balancer for network processing efficiency.
Memory and I/O bandwidth also see substantial improvements. The Xeon 6700/6500 P-core processors support 8-channel DDR5 memory. They also pave the way for MRDIMM (Multiplexed Rank DIMM) modules, which offer speeds up to 8,800 MT/s. On the I/O side, these processors support PCIe 5.0 and CXL 2.0 standards. In a dual-socket configuration, the platform can deliver up to 88 PCIe lanes per socket (176 lanes total).
Despite the differentiation between P-core and E-core processors, the Xeon 6 family maintains consistency in instruction sets, BIOS, drivers, operating system and application support, as well as RAS features, simplifying integration and management across different deployment types. P-core processors are intended for workloads where per-core performance, AI acceleration, high memory bandwidth, and substantial I/O are paramount; think demanding databases, HPC simulations, advanced analytics, and a wide range of AI applications.
Dell PowerEdge R770 specifications
| Specifications | Dell PowerEdge R770 | 
| Processor | Two Intel Xeon 6 processors with up to 144 E-cores or 86 P-cores per processor | 
| Memory | 32 DDR5 DIMM slots, supports RDIMM 8 TB max, speeds up to 6400 MT/s, supports only registered ECC DDR5 DIMMs | 
| Storage controllers | Internal boot: Boot Optimized Storage Subsystem (BOSS-N1 DC-MHS): HWRAID 1, 2 x M.2 NVMe SSDs or M.2 interposer card (DC-MHS): 2 x M.2 NVMe SSDs or USB, Internal controllers: PERC H965i front, PERC H975i front, PERC H365i front | 
| Front and rear bays |  | 
| Hot-plug power supplies |  | 
| Cooling options | Air cooling and direct liquid cooling (DLC is a rack solution and requires rack manifolds and a cooling distribution unit (CDU) to operate) | 
| Fans | High-performance Silver fans (HPR SLVR)/High-performance Gold fans (HPR GOLD), up to 6 hot-pluggable fans | 
| Dimensions and weight | Height – 86.8 mm (3.42 inches), width – 482 mm (18.97 inches), weight – 28.53 kg (62.89 lbs), depth (for rear I/O configuration) – 802.40 mm (31.59 inches) with bezel, 801.51 mm (31.56 inches) without bezel, depth (for front I/O configuration) – 814.52 mm (32.07 inches) without bezel | 
| Form factor | 2U rack server | 
| Integrated management | iDRAC, iDRAC Direct, iDRAC RESTful API with Redfish, RACADM CLI, iDRAC Service Module (iSM), NativeEdge endpoint, NativeEdge orchestrator | 
| Bezel | Optional security bezel | 
| Security | Cryptographically signed firmware, data-at-rest encryption (SED with local or external key management), Secure Boot, Secure Component Verification (hardware integrity check), silicon root of trust, system lockdown, system lock-down (requires iDRAC10 Enterprise or Datacenter), chassis intrusion detection, TPM 2.0 FIPS, CC-TCG certified | 
| Network options |  | 
| GPU options | Up to 6 x 75 W FHHL or up to 2 x 350 W DWFL | 
| Ports | Front ports: 1 USB 2.0 Type-C port, 1 USB 2.0 Type-A port (optional), 1 Mini-DisplayPort (optional), 1 DB9 serial port (with front I/O configuration), 1 iDRAC dedicated management Ethernet port; Rear ports: 1 iDRAC dedicated management Ethernet port, 1 VGA port, 2 USB 3.1 Type-A ports; Internal ports: 1 USB 3.1 Type-A port | 
| PCIe |  | 
| Operating systems and hypervisors | Canonical Ubuntu LTS Server, Microsoft Windows Server with Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware with vSphere | 
Dell PowerEdge R770: Modularity with OCP DC MHS
The Dell PowerEdge R770 features notable advances and flexibility in its physical design and component architecture, adopting the Open Compute Project's Data Center Modular Hardware System (OCP DC MHS) standard.
Continuing the R7x0 series lineage, the R770 offers numerous configuration options to meet various deployment needs. A significant first for this range is the choice between a traditional rear I/O configuration and a front I/O configuration accessible from the cold aisle, providing greater flexibility to adapt to different data center layouts and maintenance requirements. Storage options are equally versatile, ranging from compute nodes with minimal or no local storage to high-density configurations supporting up to 40 E3.S drives for storage-centric workloads.
To meet the growing need for accelerated computing, particularly for AI and HPC, the R770 offers robust expansion capabilities. Depending on the chassis and riser configuration, the server can accommodate up to six full-height, full-length (FHFL) PCIe Gen 5 x16 cards. Additionally, it supports the installation of two double-width GPUs, making it a high-performance platform for a wide range of tasks. Network flexibility is provided by OCP 3.0 mezzanine slots, supporting x8 or x16 cards depending on the configuration.
Dell has also made several design improvements aimed at enhancing serviceability and reliability. The evolution of the Boot Optimized Storage Solution (BOSS) card is a perfect example. Previously cabled and integrated into a PCIe card, the R770's BOSS controller is now a standardized OCP card that interfaces directly with the motherboard, eliminating cabling complexity. This new BOSS controller also incorporates faster M.2 NVMe drives and heatsinks to ensure optimal operating temperatures and performance for boot devices. Another subtle but practical improvement for technicians: the replacement of traditional jumpers with more user-friendly DIP switches for functions such as NVRAM clearing.
The most profound architectural change is the full adoption of the OCP DC MHS standard. Dell had already integrated OCP elements in previous generations, notably by adopting OCP 3.0 network card slots. The R770 goes further. Key components are now compliant with OCP specifications, including Host Processor Modules (HPM), commonly known as motherboards, which include components such as GPU slots, now M-XIO connectors. The M-XIO connector provides a standardized interface for riser cards, improving flexibility and scalability. iDRAC is also implemented as an OCP DC-SCM (Server Control Module).
Additionally, the R770 incorporates the new PICPWR power connector for device connections such as GPUs and backplanes. This connector represents a significant advancement, simplifying power delivery and integrating inline power monitoring.
		
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
TGI's benchmarking feature allows evaluation of its performance under different configurations and workloads. It provides a more accurate representation of real-world performance, as it accounts for the complexity of managing LLMs in a production environment.
Text generation using LLMs involves two main stages: prefill and decode. Prefill is the initial stage, where the LLM processes the input prompt to generate the necessary intermediate representations. This stage is compute-intensive, as it involves processing the entire input prompt in a single pass through the model.
During the prefill stage, the input prompt is tokenized and converted into a format usable by the LLM. The LLM then computes the KV cache, which stores information related to the input tokens. This KV cache is an essential data structure that facilitates the generation of output tokens.
In contrast, the decode stage is an autoregressive process where the LLM generates output tokens one by one, relying on the intermediate representations generated during the prefill stage. The decode stage relies heavily on the KV cache generated during the prefill stage, which provides the necessary context for generating coherent and contextually relevant output tokens.
Prefill stage
As batch size increases from 1 to 32, latency for all three models increases; DeepSeek-R1-Distill-Qwen-32B latency increases from 29.97 ms for a batch size of 1 to 76.95 ms for a batch size of 32. Similarly, GEMMA-3-27B-IT and Qwen/QwQ-32B latency increases from 51.84 ms and 29.90 ms to 79.58 ms and 76.30 ms, respectively.
In contrast, token throughput improves significantly with increasing batch size. For a batch of 1, throughput for the three models ranges from 192.95 to 334.46 tokens per second. For a batch of 32, they reach 4158.67, 4021.40, and 4194.13 tokens per second for DeepSeek-R1-Distill-Qwen-32B, GEMMA-3-27B-IT, and Qwen/QwQ-32B, respectively.
| LLM prefill stage performance: latency (ms) and token throughput (tokens/s) |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Batch size | DeepSeek-R1-Distillation-Qwen-32B |  | GEMMA-3-27B-IT |  | Qwen/QwQ-32B |  | 
|---|---|---|---|---|---|---|
|  | Latency (ms) | Token rate | Latency (ms) | Token rate | Latency (ms) | Token rate | 
| 1 | 29.97 | 333.64 | 51.84 | 192.95 | 29.90 | 334.46 | 
| 2 | 30.21 | 662.09 | 52.55 | 380.61 | 29.95 | 667.80 | 
| 4 | 32.40 | 1234.72 | 52.62 | 760.12 | 32.12 | 1245.47 | 
| 8 | 36.98 | 2163.46 | 52.66 | 1519.19 | 36.69 | 2180.66 | 
| 16 | 51.63 | 3125.50 | 60.96 | 2624.64 | 51.29 | 3147.61 | 
| 32 | 76.95 | 4158.67 | 79.58 | 4021.40 | 76.30 | 4194.13 | 
Decode stage
Unlike the prefill stage, latency during the decode stage remains relatively stable regardless of batch size. For example, DeepSeek-R1-Distill-Qwen-32B latency varies from 27.14 ms to 29.52 ms as batch size increases from 2 to 32.
Token throughput during the decode phase improves with batch size, but not as dramatically as during the prefill phase. For a batch of 1, throughput is approximately 36-37 tokens per second for DeepSeek-R1-Distill-Qwen-32B and Qwen/QwQ-32B, and 33.96 tokens per second for GEMMA-3-27B-IT. For a batch of 32, throughput increases to 1083.83, 873.39, and 1084.89 tokens per second, respectively.
| LLM decode performance (token): latency (ms) and token throughput (tokens/s) |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Batch size | DeepSeek-R1-Distillation-Qwen-32B |  | GEMMA-3-27B-IT |  | Qwen/QwQ-32B |  | 
|---|---|---|---|---|---|---|
|  | Latency (ms) | Token rate | Latency (ms) | Token rate | Latency (ms) | Token rate | 
| 1 | 27.24 | 36.71 | 29.45 | 33.96 | 27.24 | 36.71 | 
| 2 | 27.14 | 73.70 | 30.80 | 64.93 | 27.14 | 73.69 | 
| 4 | 27.50 | 145.46 | 31.33 | 127.65 | 27.47 | 145.62 | 
| 8 | 27.91 | 286.61 | 32.54 | 245.83 | 27.90 | 286.78 | 
| 16 | 28.31 | 565.07 | 34.71 | 460.92 | 28.44 | 562.56 | 
| 32 | 29.52 | 1083.83 | 36.64 | 873.39 | 29.50 | 1084.89 | 
This is normal, as the prefill stage computes the initial hidden states and key-value caches for the entire input prompt, which can saturate the GPU, as large batch operations can be executed simultaneously. After processing the prompt, the model generates new tokens, typically one at a time. At each step, the model uses the previous token and cached hidden states to produce the next token. Since this stage proceeds token by token, the batch size is often reduced, resulting in frequent GPU underutilization.
Procyon AI computer vision benchmark
Using concrete artificial vision tasks, the Procyon AI Computer Vision benchmark evaluates AI inference performance on CPUs, GPUs, and AI accelerators. It supports multiple inference engines such as TensorRT, OpenVINO, SNPE, Windows ML, and Core ML, providing insights into efficiency, compatibility, and optimization.
The Procyon AI Computer Vision benchmark results also demonstrate excellent AI inference performance. The system achieved low inference times, with MobileNet V3 at 20.64 ms and ResNet 50 at 22.42 ms. Inception V4 and DeepLab ran at 65.23 ms and 41.37 ms, respectively, efficiently handling more complex vision workloads. YOLO V3, a key object detection model, processed in 37.80 ms, making it particularly suitable for real-time AI applications. REAL-ESRGAN, a computationally intensive super-resolution model, recorded 1,159.22 ms, earning us an overall score of 81 in AI Computer Vision.
| AI Computer Vision (shorter time is better) (higher score is better) | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | 
|---|---|
| MobileNet V3 average inference time | 20.64 ms | 
| ResNet 50 average inference time | 22.42 ms | 
| Inception V4 average inference time | 65.23 ms | 
| DeepLab average inference time | 41.37 ms | 
| YOLO V3 average inference time | 37.80 ms | 
| REAL-ESRGAN average inference time | 1,159.22 ms | 
| AI Computer Vision overall score | 81 | 
Hammer DB TPROC-C
We also evaluated the performance of four popular open-source databases (MariaDB 11.4.4, MySQL 8.4.4, MySQL 5.7.44, and PostgreSQL 17.2) using the HammerDB TPROC-C benchmark to simulate OLTP workloads on 500 warehouses.
MariaDB emerged as the highest-performing solution, particularly in dual-socket configurations, where it scaled efficiently and achieved the highest transaction throughput. MySQL 8.4.4 showed notable improvements over the older 5.7.44 version, highlighting enhancements in recent releases. PostgreSQL 17.2 delivered consistent performance but lagged slightly behind MariaDB and MySQL 8.4.4. MariaDB delivered 3.15 million TPM on a single socket and 5.8 million TPM on two sockets, outperforming the others in both scenarios.
Performance comparison table (Transactions per minute, TPM)
| Database engine | Single-socket TPM | Dual-socket TPM | 
|---|---|---|
| MariaDB 11.4.4 | 3,150,000 | 5,800,000 | 
| MySQL 8.4.4 | 2,850,000 | 5,150,000 | 
| PostgreSQL 17.2 | 2,700,000 | 4,900,000 | 
| MySQL 5.7.44 | 2,300,000 | 4,250,000 | 
Despite the R770's hardware power, with its 86 cores per processor (a mix of high and low priority cores), no database recorded significant performance gains when spread across both sockets. This reflects the general preference of open-source databases for running on a single socket, due to better core locality and reduced memory latency.
Given these results, the R770 is better suited for running multiple database instances in a virtualized environment than for scaling a single instance. The system architecture is ideal for supporting a high-density mixed database workload, leveraging both performance and efficiency cores to ensure consistent throughput across multiple instances.
7-Zip
The built-in memory benchmark tool of the popular 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, indicating how well the system can handle data-intensive operations.
In the 7-Zip benchmark, the Dell system achieved a higher score (266.425 GIPS) than Lenovo (224.313 GIPS) for compression tasks, with the Dell system showing slightly lower CPU utilization. However, Lenovo outperformed Dell in decompression, with a higher score (288.457 GIPS vs. 256.154 GIPS) and slightly higher CPU utilization. Dell achieved a slightly higher overall score (261.290 GIPS), demonstrating better overall efficiency for compression and decompression tasks.
| 7-Zip Compression and Decompression | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| Compression – Current CPU usage | 5267 % | 5064 % | 
| Compression – Current rating/usage | 5.061 GIPS | 4.341 GIPS | 
| Compression – Current rating | 266.591 GIPS | 219.840 GIPS | 
| Compression – Resulting CPU usage | 5270 % | 5156 % | 
| Compression – Resulting rating/usage | 5.056 GIPS | 4.350 GIPS | 
| Compression – Resulting rating | 266.425 GIPS | 224.313 GIPS | 
| Decompression – Current CPU usage | 5623 % | 6184 % | 
| Decompression – Current rating/usage | 4.586 GIPS | 4.688 GIPS | 
| Decompression – Current rating | 257.909 GIPS | 289.879 GIPS | 
| Decompression – Resulting CPU usage | 5627 % | 6205 % | 
| Decompression – Resulting rating/usage | 4.553 GIPS | 4.649 GIPS | 
| Decompression – Resulting rating | 256.154 GIPS | 288.457 GIPS | 
| Total – Total CPU usage | 5448 % | 5681 % | 
| Total – Total rating/usage | 4.804 GIPS | 4.500 GIPS | 
| Total – Total rating | 261.290 GIPS | 256.385 GIPS | 
y-cruncher
y-cruncher is a popular benchmarking and stress-testing application launched in 2009. This test is multithreaded and scalable, computing Pi and other constants to thousands of billions of digits. Faster is better in this test. This software has been fantastic for testing high-core-count platforms and showing the compute advantages between single- and dual-socket platforms.
The Y-Cruncher benchmark results show a significant performance gap between the Dell PowerEdge R770, equipped with P-core processors, and the Lenovo ThinkSystem SR630 V4 equipped with E-core processors, particularly as dataset size increases. This is less about determining which system is better than comparing processor types under this workload.
For smaller-scale computations, the Dell system was already ahead, computing 1 billion digits of Pi in 2.753 seconds, while the Lenovo system took more than double, at 5.997 seconds. As the workload increased, the gap widened. At 10 billion digits, the Dell system finished in 34.873 seconds, less than half of the Lenovo system's 81.046 seconds. Beyond 50 billion digits, Dell maintained its lead, completing the task in 221.255 seconds, versus 476.826 seconds for Lenovo, making Dell 53% faster.
At 100 billion digits, Lenovo could not complete the test due to its current 512 GB RAM configuration. With 2 TB of RAM, Dell handled the workload efficiently, finishing in 491.737 seconds.
| Y-cruncher (shorter time is better) | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| 1 billion | 2.753 seconds | 5.997 seconds | 
| 2.5 billion | 7.365 seconds | 17.573 seconds | 
| 5 billion | 16.223 seconds | 37.793 seconds | 
| 10 billion | 34.873 seconds | 81.046 seconds | 
| 25 billion | 99.324 seconds | 220.025 seconds | 
| 50 billion | 221.255 seconds | 476.826 seconds | 
| 100 billion | 491.737 seconds |  | 
OptiX Blender
An open-source 3D modeling application. This benchmark was performed with the Blender Benchmark utility. The score is expressed in samples per minute, with the highest being the best.
The Blender benchmark results show a clear performance advantage for the Dell PowerEdge R770 over the Lenovo ThinkSystem SR630 V4, particularly in CPU rendering. In the "CPU Monster" test, Dell reached 1,706.002 samples per minute, a 19% lead over Lenovo's 1,432.09 samples per minute. The "CPU Junkshop" test further widened this gap, with Dell reaching 1,169.370 samples per minute, surpassing Lenovo's 914.75 samples per minute by 28%. Similarly, Dell recorded 791.475 samples per minute in the "CPU Classroom" test, while Lenovo lagged at 656.68 samples per minute, a difference of 20%.
The absence of a GPU in the Lenovo system also meant it could not participate in GPU-based rendering, where Dell's NVIDIA L4 scored 1,895.71 samples/min for Monster, 950.42 samples/min, and a Classroom score of 968.43 samples/min.
| Blender CPU Benchmark | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| CPU Monster (Blender 4.3) | 1,706.002 samples/min | 1432.09 samples/min | 
| CPU Junkshop (Blender 4.3) | 1,169.370 samples/min | 914.75 samples/min | 
| CPU Classroom (Blender 4.3) | 791.475 samples/min | 656.68 samples/min | 
| GPU Monster (Blender 4.3) | 1,895.712 samples/min | (no GPU) | 
| GPU Junkshop (Blender 4.3) | 950.424 samples/min | (no GPU) | 
| GPU Classroom (Blender 4.3) | 968.432 samples/min | (no GPU) | 
Cinebench R23
The Cinebench R23 benchmark tool evaluates a system's CPU performance by rendering a complex 3D scene using the Cinema 4D engine. It measures single-core and multi-core performance, providing a comprehensive view of the processor's capabilities in handling 3D rendering tasks.
In Cinebench R23, the benchmark results highlight notable CPU performance differences between the Dell PowerEdge R770 and the Lenovo ThinkSystem SR630 V4, particularly in terms of cores per processor. The Lenovo ThinkSystem SR630 V4, equipped with two Intel Xeon 6780E processors (144 cores per processor), outperformed Dell in the multi-core CPU test with a score of 99,266 points, versus 74,710 points for Dell. This gap reflects Lenovo's advantage in multithreaded workloads, thanks to its higher core count (288 cores total) compared to Dell's two Intel Xeon 6787P processors (86 cores per processor), which limits its multi-core performance.
In the Single-Core CPU test, Dell performed better with a score of 1,272 points, surpassing Lenovo's 894 points, highlighting Dell's superior single-threaded processor efficiency despite its lower core count.
| Cinebench R23 | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| Multi-core CPU | 74,710 pts | 99,266 pts | 
| Single-core CPU | 1,272 pts | 894 pts | 
| MP ratio | 58.74 x | 111.00 x | 
Cinebench 2024
Cinebench 2024 extends R23's benchmarking capabilities by adding GPU performance evaluation. It continues to test CPU performance but also includes tests that measure the GPU's ability to handle rendering tasks.
In this updated benchmark, the Dell PowerEdge R770 scored 12,996 points for GPU performance, highlighting its ability to handle GPU-accelerated rendering tasks. The Lenovo ThinkSystem SR630 V4 has no dedicated GPU and therefore recorded no GPU score.
In the multi-core CPU test, Lenovo scored 2,884 points, slightly ahead of Dell's 2,831 points, demonstrating a slight advantage in multi-core performance. In single-core CPU, Dell outperformed Lenovo, scoring 71 points versus 53 points for Lenovo, demonstrating Dell's superior single-core performance despite a reduced core count.
| Cinebench R24 | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| GPU score | 12,996 pts |  | 
| Multi-core CPU | 2,831 pts | 2,884 pts | 
| Single-core CPU | 71 pts | 53 pts | 
| MP ratio | 39.77 x | 54.43 x | 
Geekbench 6
Geekbench 6 is a cross-platform benchmarking tool that measures a system's overall performance. The Geekbench browser allows comparison of any system with this tool.
The Geekbench 6 benchmark results show clear performance differences between the Dell PowerEdge R770 and the Lenovo ThinkSystem SR630 V4. In the single-core CPU test, Dell outperformed Lenovo with a score of 1,797, while Lenovo scored 1,173, a 53% improvement in single-core performance for Dell.
In the multi-core CPU test, Dell again dominated with 15,880 points, while Lenovo scored 13,868 points, giving Dell a 14% advantage in multi-core performance. This suggests that Dell's Intel Xeon 6787P processors offer superior overall compute power, particularly for tasks requiring multiple cores.
The OpenCL GPU test further highlighted Dell's advantage, with a score of 148,730 thanks to the NVIDIA L4 GPU.
| Geekbench 6 (Higher is better) | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| Single-core CPU | 1,797 | 1,173 | 
| Multi-core CPU | 15,880 | 13,868 | 
| GPU OpenCL score | 148,730 | (no GPU) | 
Blackmagic RAW Speed Test
The Blackmagic RAW Speed Test is a performance benchmarking tool designed to measure a system's ability to handle video playback and editing using the Blackmagic RAW codec. It evaluates a system's ability to decode and play high-resolution video files, providing frame rates for both CPU-based and GPU-based processing.
In the CPU-based test, the Dell PowerEdge R770 reached 141 fps, surpassing the Lenovo ThinkSystem SR630 V4, which scored 120 fps. This indicates that the Dell system handles CPU-based video processing more efficiently than the Lenovo. In the GPU-based test, the Dell PowerEdge R770 achieved 157 fps, benefiting from the presence of an NVIDIA GPU.
| Blackmagic RAW Speed Test (higher is better) | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| CPU FPS | 141 FPS | 120 FPS | 
| CUDA FPS | 157 FPS | 0 FPS (no GPU) | 
Blackmagic Disk Speed Test
The Blackmagic Disk Speed Test evaluates a disk's read and write speeds, assessing its performance, particularly for video editing tasks. It helps users ensure their storage is fast enough for high-resolution content, such as 4K or 8K video.
In the Blackmagic speed test, the Dell PowerEdge R770 BOSS card with mirrored SK hynix 480 GB Dell NVMe achieved a read speed of 3,010.3 MB/s and a write speed of 976.3 MB/s.
Conclusion
The Dell PowerEdge R770 particularly excites us, thanks to its adoption of the Open Compute Project's Data Center Modular Hardware System standard and its cutting-edge hardware. The integration of OCP DC MHS offers numerous benefits, including increased modularity, improved serviceability, and potential cost reductions through greater standardization. This design philosophy is reflected in every aspect of the system, from the implementation of iDRAC as OCP DC-SCM to the ports.
The R770 also offers impressive storage capabilities, supporting up to 40 E3.S drives in a single 2U chassis, making it the ideal solution for storage-intensive workloads. Additionally, the server's flexibility is enhanced by support for various configurations, including a cold-aisle-accessible front I/O configuration, providing greater flexibility to adapt to different data center layouts and maintenance requirements.
Compatible with a wide range of GPUs and Intel Xeon 6 Performance processors, the R770 is a powerful and versatile server platform, perfectly suited to the demands of modern data centers. Its cutting-edge hardware, modular design, and robust security features make the R770 an attractive option for organizations looking to deploy AI applications, high-performance computing (HPC), and traditional enterprise workloads.
