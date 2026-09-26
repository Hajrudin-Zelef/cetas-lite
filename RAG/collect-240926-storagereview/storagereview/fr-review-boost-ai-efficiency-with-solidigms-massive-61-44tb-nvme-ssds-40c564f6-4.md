---
id: collect-240926-storagereview/storagereview/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6-4
title: "fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6"
domain: storagereview
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["cost", "energy", "fine-tuning", "gpu", "gpus", "inference", "latency", "nand", "nvidia", "rack-scale", "throughput", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6.md
source_anchor: ""
source_lines: [90, 148]
sha256: 243ec6d7d31ef6c893b855102bfe37b50517eaa731cd9c9b7d0ddbca19006594
---

# fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6

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
