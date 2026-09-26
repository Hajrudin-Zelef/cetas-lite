---
id: collect-240926-storagereview/storagereview/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6-3
title: "fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "cost", "fine-tuning", "gpu", "gpus", "inference", "latency", "memory", "nvidia", "parameters"]
source: docs/RAG/clean_en/storagereview/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6.md
source_anchor: ""
source_lines: [36, 89]
sha256: aa9ddf021d839c85d663659573c7d36483a5a27e997b763de74ca4781c79b655
---

# fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6

Suppose you are working on a single server and fine-tuning models suited to a handful of locally connected GPUs. In that case, you benefit from local storage, which is simpler to set up and manage than networked storage solutions. Configuring, administering, and maintaining networked storage can be complex and time-consuming, often requiring specialized knowledge and additional infrastructure. In contrast, local storage solutions such as NVMe SSDs are simpler to integrate into existing server configurations.
This simplicity of setup and maintenance allows IT teams to focus more on optimizing AI workloads rather than on the intricacies of network storage management. As a result, deploying and managing storage for AI applications becomes simpler and more efficient with local NVMe SSDs.
Cost and scalability
Even though NAS solutions can scale horizontally by adding more storage devices, they also entail costs related to network infrastructure and potential performance bottlenecks. Conversely, investing in high-capacity local storage can offer immediate performance benefits without major network upgrades.
Local storage solutions are often more cost-effective than network-attached storage (NAS) systems, as they eliminate the need for expensive network hardware and complex configurations. Setting up and maintaining a NAS involves significant investments in network equipment, such as high-speed switches and routers, as well as ongoing network management and maintenance costs.
High-capacity local SSDs integrated directly into the server are used as a staging area, thereby reducing the need for additional infrastructure. This direct integration reduces hardware costs and simplifies the setup process, making it more economical for organizations seeking to optimize their AI workloads without incurring high expenses.
To thoroughly evaluate the performance of Solidigm 61.44TB QLC SSDs in an AI server configuration, we will compare a range of four Solidigm P5336 61.44TB SSDs installed in a Lenovo ThinkSystem SR675 V3. This server configuration also includes a set of four NVIDIA L40S GPUs. The benchmarking tool used for this purpose is GDSIO, a specialized utility designed to measure storage performance in GPU Direct Storage (GDS) environments. We examined two configurations: one GPU for single-drive performance and one GPU for four drives configured for RAID0.
Stay with us. The following sections cover the specifics of the tests and how they mimic the different stages of the AI pipeline.
Test parameters
The benchmarking process involves various test parameters that simulate different stages of the AI pipeline. These parameters include io_sizes, threads, and transfer_type, each chosen to represent specific aspects of AI workloads.
1. I/O sizes:
- 4K, 128K, 256K, 512K, 1M, 4M, 16M, 64M, 128M: These different I/O sizes allow for simulating different data transfer patterns. Smaller I/O sizes (128K, 256K, 512K) mimic scenarios in which small chunks of data are frequently accessed, which is typical during data preparation stages. Larger I/O sizes (1M, 4M, 16M, 64M, 128M) represent bulk data transfers, often observed during training and inference stages, during which entire batches of data are moved.
2. Threads:
- 1, 4, 16, 32: The number of threads represents the level of concurrency in data access. A single thread tests baseline performance, while a higher number of threads (4, 16, 32) simulates more intensive parallel data processing activities, similar to what occurs during large-scale training sessions where multiple data streams are handled simultaneously.
3. Transfer types:
- Storage->GPU (GDS): This transfer type leverages GPU Direct Storage (GDS), enabling direct data transfers between SSDs and GPUs, bypassing the CPU. This configuration is ideal for testing the efficiency of direct data paths and minimizing latency, reflecting real-time inference scenarios.
- Storage->CPU->GPU: This traditional data transfer path involves moving data from storage to the CPU before transferring it to the GPU. This method simulates scenarios in which intermediate processing or caching may occur at the processor level, which is expected during the data preparation phase. We could argue that this data path would represent performance regardless of the GPU vendor.
- Storage->PAGE_CACHE->CPU->GPU: This path uses the page cache for data transfers, where data is first cached in memory before being processed by the CPU and then transferred to the GPU. This configuration is useful for testing the impact of caching mechanisms and memory bandwidth on overall performance, which is relevant during training when data can be preprocessed and cached for efficiency. Again, we could argue that this data path would represent performance regardless of the GPU vendor.
Mimicking the stages of the AI pipeline
The benchmarks are designed to reflect the different stages of the AI pipeline, ensuring that the performance measurements obtained are relevant and comprehensive.
Data preparation:
- I/O sizes: Smaller (128K, 256K, 512K)
- Threads: 1, 4
- Transfer types: "Storage->CPU->GPU", "Storage->PAGE_CACHE->CPU->GPU"
- Purpose: Evaluate how SSDs handle frequent small data transfers and CPU involvement, essential during data ingestion, cleaning, and augmentation phases.
Training and fine-tuning:
- I/O sizes: Medium to large (1M, 4M, 16M)
- Threads: 4, 16, 32
- Transfer types: "Storage->GPU (GDS)", "Storage->CPU->GPU"
- Purpose: Evaluate performance under high data throughput conditions with multiple simultaneous data streams, representing the intensive data handling required during model training and fine-tuning.
Inference:
- I/O sizes: Large to very large (16M, 64M, 128M) and 4K
- Threads: 1, 4, 16
- Transfer types: Storage->GPU (GDS)
- Purpose: Measure the efficiency of direct large-scale data transfers to the GPU, essential for real-time inference applications where rapid data access and minimal latency are paramount. 4K is designed to examine lookups performed in the RAG database.
By varying these parameters and testing different configurations, we can obtain a detailed performance profile of Solidigm 61.44TB QLC SSDs in a high-performance AI server environment, thereby providing insight into their suitability and optimization for various AI workloads. We examined the data by performing over 1200 tests over a few weeks.
Server configuration
- Lenovo ThinkSystem SR675 V3
- AMD EPYC 9254 24-core processor
- 6 X 64GB DDR5, total capacity of 384GB
- 4X NVIDIA L40S GPUs
- 4 Solidigm P5336 61.44TB QLC NVMe SSDs
- Ubuntu Server 22.04
- NVIDIA driver version: 535.171.04
- CUDA version: 12.2
Benchmark results
First, let's look at training and inference-type workloads. The 1024K GPU Direct I/O size represents model loading, training data loaded onto the GPU, and other large batch inference tasks such as in image or video work.
| 4Drive | I/O Type | Transfer Type | Threads | Dataset Size (KiB) | I/O Size (KB) | Throughput (GiB/sec) | Average Latency (usecs) | 
|---|---|---|---|---|---|---|---|
|  | WRITE | GPUD | 8 | 777,375,744 | 1024 | 12.31 | 634.55 | 
|  | READ | GPUD | 8 | 579,439,616 | 1024 | 9.30 | 840.37 | 
|  | RANDWRITE | GPUD | 8 | 751,927,296 | 1024 | 12.04 | 648.67 | 
|  | RANDREAD | GPUD | 8 | 653,832,192 | 1024 | 10.50 | 743.89 | 
Next, examine the smaller I/O sizes, for a RAG-type workload, for example, where rapid random access to 4K data to a RAG database stored on disk is needed. Efficient random I/O is necessary for scenarios in which inference workloads must access data non-sequentially, such as with recommendation systems or search applications. The RAID0 configuration shows good performance for both sequential and random operations, which is crucial for AI applications that involve a mix of access patterns like RAG. Read latency values are particularly low, especially in GPUD.
