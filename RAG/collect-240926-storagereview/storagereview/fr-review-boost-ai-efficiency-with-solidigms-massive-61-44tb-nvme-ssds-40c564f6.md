---
id: collect-240926-storagereview/storagereview/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6
title: "fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "compute", "cost", "energy", "fine-tuning", "gpu", "gpus", "inference", "latency", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6.md
source_anchor: ""
source_lines: [1, 148]
sha256: 66aa2139b6a3227b4d7934adce478fa4f665075babaca605d4533abc8027677d
---

# fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6

<!-- source: https://www.storagereview.com/fr/review/boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds -->

It is no secret that we love the massive density of Solidigm 61.44TB U.2 NVMe SSDs. We have conducted numerous endurance and performance tests, made scientific discoveries, and pushed world record calculations to extraordinary new heights. So, with the AI boom developing at a frenetic pace all around us, the next logical step was to see how Solidigm NVMe drives compare in the dynamic world of AI 2024.
Understanding the benefits of extreme storage density
Solidigm's 61.44TB QLC SSDs stand out for their remarkable storage capacity, allowing data centers to pack more storage onto fewer drives. This extreme density is particularly advantageous in AI servers, where datasets are experiencing exponential growth and efficient storage solutions are paramount. Thanks to these high-capacity SSDs, data centers can reduce the number of physical drives, decrease footprint, reduce power consumption, and simplify maintenance.
Limited PCIe lanes in GPU servers
One of the main challenges of modern GPU servers is the limited number of PCIe lanes available once the GPUs have gotten their share. Essential for AI workloads, GPUs require significant PCIe bandwidth, often leaving limited lanes for other components, including storage peripherals and networking. This constraint makes optimizing the use of available PCIe lanes indispensable. Solidigm's 61.44TB QLC SSDs offer a solution by providing massive storage capacity on a single drive, reducing the need for multiple drives and preserving PCIe lanes for GPUs and other essential components.
AI workloads and storage requirements
AI workloads can be broadly classified into three phases: data preparation, training and fine-tuning, and inference. Each phase has unique storage requirements, and Solidigm's high-capacity SSDs can significantly improve performance and efficiency during these phases. Deploying high-capacity QLC drives, such as the Solidigm D5-P5336, benefits all AI workloads. Most of the benefits range from data preparation to training and from fine-tuning to inference.
Data preparation
Data preparation is the foundation of any AI project and involves collecting, cleaning, transforming, and augmenting data. This phase requires extensive storage because raw datasets can be enormous. Solidigm's 61.44TB QLC SSDs can store large amounts of raw data without compromising performance. Additionally, the high sequential read and write speeds of these SSDs ensure rapid data access, thereby accelerating the preparation process. For data preparation, Solidigm 61.44TB QLC SSDs meet all the requirements described above with benefits such as:
- Massive storage capacity: Efficient management of large datasets.
- High sequential speeds: Rapid data access and processing.
- Reduced latency: Minimized delays in data retrieval, thereby improving workflow efficiency.
Training and fine-tuning
Training AI models is an intensive process that involves feeding numerous datasets into neural networks to adjust weights and biases. This phase is compute-demanding and requires high IOPS (input/output operations per second) and low-latency storage to keep up with the rapid data exchanges between storage and GPUs. Solidigm SSDs excel in this regard, offering high performance and endurance. The extreme density of these SSDs allows for the use of larger datasets in training, potentially leading to more accurate models. To meet the demands of training and fine-tuning, Solidigm SSDs offer the following:
- High IOPS: Supports the rapid data exchanges essential for training.
- Endurance: QLC technology optimized for read/write-heavy workloads, ideal for repeated training cycles.
- Scalability: Expand storage without adding physical drives, while maintaining efficient use of PCIe lanes.
Inference
Once trained, AI models are deployed to make predictions or decisions based on new data, which is called inference. This phase often requires rapid access to preprocessed data and efficient handling of increased read requests. Solidigm's 61.44TB QLC SSDs provide the necessary read performance and low latency to ensure that inference operations are performed smoothly and quickly. Solidigm SSDs exceed performance and low latency expectations by offering the following benefits:
- Fast read performance: Ensures rapid data access for real-time inference.
- Low latency: Critical for applications requiring immediate responses.
- High capacity: Efficiently store numerous inference data and historical results.
QLC technology offers significant benefits for inference applications, including high storage capacity, cost-effectiveness, fast read speeds, efficient PCIe utilization, endurance, and improved workflow efficiency. These benefits collectively enhance the performance, scalability, and cost-effectiveness of inference tasks, making QLC drives an ideal choice for modern AI and machine learning deployments.
Why is it important to have large storage as close as possible to the GPU?
For AI and machine learning, the proximity of storage to the GPU can have a significant impact on performance. Designing an AI data center requires careful consideration of several factors to ensure optimal functionality and efficiency. This is why it is crucial to have extensive storage as close as possible to the GPU. As we recently explored, access to a large-scale networked storage solution is beginning to transform into a single tool, but relying solely on it may not always be the optimal choice.
Latency and bandwidth
One of the main reasons for placing significant storage close to the GPU is to minimize latency and optimize bandwidth. AI workloads, particularly during training, involve frequent and massive data transfers between storage and the GPU. High latency can block the entire process, slowing down training times and reducing efficiency.
In AI workloads, where rapid data availability is essential, low latency ensures that GPUs receive data quickly, thereby reducing idle times and improving overall computational efficiency. During the training phase, enormous volumes of data must be continuously fed into the GPU for processing. By minimizing latency, DAS ensures that the high-speed demands of AI applications are met, resulting in faster training times and more efficient workflows.
Data throughput and I/O performance
Local NVMe SSDs excel at handling a large number of input/output operations per second (IOPS), which is crucial for the read/write-intensive nature of AI workloads. During the training phase, AI models require rapid access to vast data repositories, necessitating storage solutions capable of meeting the high demand for data transactions.
The Solidigm D5-P5336, designed for high-capacity and high-performance scenarios, offers exceptional IOPS, enabling faster data retrieval and write processes. This capability ensures that GPUs stay busy with computation rather than waiting for data, thereby maximizing efficiency and reducing training times. The high IOPS performance of local NVMe SSDs makes them ideal for the demanding environments of AI applications, where rapid data access and processing are essential for optimal performance.
Data management
While in some scenarios, having sufficient storage directly connected to the GPU simplifies data management, it adds a layer of data management necessary to store data on the GPU server. In a perfect world, your GPU is busy computing and your CPU connects to the network to save checkpoints or fetch new data. Solidigm 61.44TB drives help reduce the number of data transactions required. You can also take this into account by using a simplified network configuration and distributed file systems. This simple approach can streamline workflows and reduce the risk of data-related errors or delays.
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
8 worker threads were selected here, which do not fully saturate the SSD but provide a more representative snapshot of what you might find in a RAG-type workload. This provides an out-of-the-box application context around the GPU's perspective with a limited number of worker tasks and higher queue depth, which is worth noting as it shows that there is more performance left on the table that can be obtained through further software optimizations.
| 4Drive | I/O Type | Transfer Type | Threads | Dataset Size (KiB) | I/O Size (KB) | Throughput (GiB/sec) | Average Latency (usecs) | 
|---|---|---|---|---|---|---|---|
|  | WRITE | GPUD | 8 | 69,929,336 | 4 | 1.12 | 27.32 | 
|  | READ | GPUD | 8 | 37,096,856 | 4 | 0.59 | 51.52 | 
|  | RANDWRITE | GPUD | 8 | 57,083,336 | 4 | 0.91 | 33.42 | 
|  | RANDREAD | GPUD | 8 | 27,226,364 | 4 | 0.44 | 70.07 | 
If you are not using GPU Direct due to unsupported libraries or GPUs, here are those two types if you use the CPU in the data transfer. In this specific server, the Lenovo ThinkSystem SR675 V3, given that all PCIe devices go through the CPU root complex, we see comparable bandwidth but take a hit on our latency. An improvement can be expected from a system with PCIe switches.
| 4Drive | I/O Type | Transfer Type | Threads | Dataset Size (KiB) | I/O Size (KB) | Throughput (GiB/sec) | Average Latency (usecs) | 
|---|---|---|---|---|---|---|---|
|  | WRITE | CPU_GPU | 8 | 767,126,528 | 1024 | 12.24 | 638.05 | 
|  | READ | CPU_GPU | 8 | 660,889,600 | 1024 | 10.58 | 738.75 | 
|  | RANDWRITE | CPU_GPU | 8 | 752,763,904 | 1024 | 12.02 | 649.76 | 
|  | RANDREAD | CPU_GPU | 8 | 656,329,728 | 1024 | 10.47 | 746.26 | 
|  | WRITE | CPU_GPU | 8 | 69,498,220 | 4 | 1.11 | 27.47 | 
|  | READ | CPU_GPU | 8 | 36,634,680 | 4 | 0.58 | 52.31 | 
The table shows high throughput for read operations, particularly with the GPUD transfer type. For example, read operations in GPUD mode reach over 10.5 GiB/sec. This benefits AI workloads, which often require rapid data access to train large models.
The balanced performance between random and sequential operations makes this configuration suitable for inference tasks, which often require a mix of these access patterns. Even though latency values are not extremely low, they remain within acceptable limits for many inference applications.
Additionally, we see impressive throughput, with write operations reaching up to 12.31 GiB/s and read operations up to 9.30 GiB/s. This high throughput benefits AI workloads that require rapid data access for model training and inference.
Sequential reads and optimization
Moving to a 128M I/O size and cycling through worker threads, we can see the result of optimizing a workload for a storage solution.
| Transfer Type | Threads | Throughput (GiB/s) | Latency (usec) | 
|---|---|---|---|
| Storage->CPU->GPU | 16 | 25.134916 | 79528.88255 | 
| Storage->CPU->GPU | 4 | 25.134903 | 19887.66948 | 
| Storage->CPU->GPU | 32 | 25.12613 | 159296.2804 | 
| Storage->GPU (GDS) | 4 | 25.057484 | 19946.07198 | 
| Storage->GPU (GDS) | 16 | 25.044871 | 79770.6007 | 
| Storage->GPU (GDS) | 32 | 25.031055 | 159478.8246 | 
| Storage->PAGE_CACHE->CPU->GPU | 16 | 24.493948 | 109958.4447 | 
| Storage->PAGE_CACHE->CPU->GPU | 32 | 24.126103 | 291792.8345 | 
| Storage->GPU (GDS) | 1 | 23.305366 | 5362.611458 | 
| Storage->PAGE_CACHE->CPU->GPU | 4 | 21.906704 | 22815.52797 | 
| Storage->CPU->GPU | 1 | 15.27233 | 8182.667969 | 
| Storage->PAGE_CACHE->CPU->GPU | 1 | 6.016992 | 20760.22778 | 
Properly writing any application to interact with storage is paramount and must be taken into account as companies seek to maximize their GPU investment.
GPU direct
By isolating GPU Direct performance only across all tests, we can get a general idea of how NVIDIA technology shines.
| I/O Type | Transfer Type | Threads | Dataset Size (KiB) | I/O Size (KB) | Throughput (GiB/sec) | Average Latency (usecs) | 
|---|---|---|---|---|---|---|
| WRITE | GPUD | 8 | 777,375,744 | 1024 | 12.31 | 634.55 | 
| READ | GPUD | 8 | 579,439,616 | 1024 | 9.30 | 840.37 | 
| RANDWRITE | GPUD | 8 | 751,927,296 | 1024 | 12.04 | 648.67 | 
| RANDREAD | GPUD | 8 | 653,832,192 | 1024 | 10.50 | 743.89 | 
| WRITE | GPUD | 8 | 69,929,336 | 4 | 1.12 | 27.32 | 
| READ | GPUD | 8 | 37,096,856 | 4 | 0.59 | 51.52 | 
| RANDWRITE | GPUD | 8 | 8,522,752 | 4 | 0.14 | 224.05 | 
| RANDREAD | GPUD | 8 | 21,161,116 | 4 | 0.34 | 89.99 | 
| RANDWRITE | GPUD | 8 | 57,083,336 | 4 | 0.91 | 33.42 | 
| RANDREAD | GPUD | 8 | 27,226,364 | 4 | 0.44 | 70.07 | 
Closing thoughts
Since this article focuses on the Solidigm 61.44TB P5336, let's take a step back and address the TLC vs QLC debate around performance versus capacity. When looking at other products in the Solidigm portfolio, such as the D7 series, which uses TLC 3D NAND, capacity is limited in exchange for performance. In our testing, particularly with the Solidigm 61.44TB drives, we found overall throughput performance that can properly keep GPUs fed with data at low latencies. We hear feedback from ODMs and OEMs regarding the demand for storage increasingly closer to the GPU, and the Solidigm D5-P5336 drive seems to fit the bill. As there are generally a limited number of NVMe bays available on GPU servers, Solidigm dense drives are at the top of the list for local GPU server storage.
Ultimately, the massive storage capacity offered by these drives, combined with GPUs, is only part of the solution; their performance remains indispensable. By adding up the performance of a single drive across multiple drives, we see that sufficient throughput is available, even for the most demanding tasks. In the case of a four-drive RAID 0 configuration using GDSIO, total write throughput can reach 12.31 GiB/s, and read throughput 25.13 GiB/s.
This level of throughput is more than sufficient for the most demanding AI tasks, such as training large deep learning models on massive datasets or performing real-time inference on high-resolution video streams. The ability to increase performance by adding more drives to the RAID0 array makes it a must-have choice for AI applications where rapid and efficient data access is crucial.
However, it is important to note that RAID0 configurations, while offering high performance, do not provide any data redundancy. It is therefore essential to implement appropriate backup and data protection strategies to avoid any data loss in the event of a drive failure.
Another unique consideration in data centers today is power. As AI servers consume more energy than ever and show no signs of slowing down, total available power is one of the biggest bottlenecks for those looking to integrate GPUs into their data centers. This means that the focus is even more on economy for every possible watt. If you can get more TB per watt, we touch on some interesting thought processes around total cost of ownership and infrastructure costs. Even removing these drives from the GPU server and placing them in a rack-scale storage server can provide massive throughput with extreme capacities.
Integrating Solidigm D5-P5336 61.44TB QLC SSDs with AI servers with limited NVMe slots represents a significant advancement in addressing the storage challenges of modern AI workloads. Their extreme density, performance characteristics, and TB/watt make them ideal for the data preparation, training and fine-tuning, and inference phases. By optimizing PCIe lane utilization and providing high-capacity storage solutions, these SSDs enable the modern AI Factory to focus on developing and deploying more sophisticated and accurate models, thereby driving innovation in the field of AI.
Lenovo ThinkSystem SR675 V3 Page
This report is sponsored by Solidigm. All points of view and opinions expressed in this report are based on our impartial view of the product(s) under study.
