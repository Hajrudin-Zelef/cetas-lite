---
id: collect-240926-storagereview/storagereview/fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655-4
title: "fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "gpus", "inference", "latency", "memory", "nvidia", "throughput", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655.md
source_anchor: ""
source_lines: [57, 83]
sha256: ccf964fd1c2b969a8e8efde378e722b31e24ee3a8f5c221d17b23a31833c7e7b
---

# fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655

The result is a platform perfectly suited to PEAK:AIO workloads. The R7725xd offers high-density NVMe capacity, PCIe Gen5 throughput, two AMD EPYC 9005 processors for parallelism, and network capacity enabling multi-client data ingestion at several hundred gigabits per client. All of these characteristics are essential for achieving the performance expected from PEAK:AIO.
PEAK:AIO – NVMe-oF RDMA – Bandwidth
Analysis of the NVMe-oF RDMA bandwidth results on the PowerEdge R7725xd with PEAK:AIO confirms expectations for a system with such PCIe and network bandwidth. Throughput increases rapidly with block size until it stabilizes near the platform's practical limit.
With small block sizes, performance starts around 20 GB/s for both reads and writes, which is normal since 4 and 8 KB transfers place more demand on the IOPS path than the throughput path. Once 16 and 32 KB blocks are reached, throughput improves considerably. Reads climb to about 154 GB/s at 32 KB and continue to progress to about 160 GB/s, which corresponds to the expected performance of a two-client configuration on four 200 Gb/s links.
Random read performance is nearly identical to sequential read performance. PEAK:AIO efficiently manages command queue feeding, so random read bandwidth closely follows sequential read bandwidth, stabilizing between approximately 159 and 161 GB/s from 32,000 to 1 million accesses. This indicates that the storage stack presents no bottleneck under mixed access, and that the R7725xd's PCIe topology distributes the load evenly across the 24 Gen5 NVMe drives.
Write performance follows a similar curve, although it peaks slightly below reads. Sequential writes remain between 140 and 148 GB/s for medium-sized blocks, dropping to about 117 GB/s for 128 KB blocks, before rising again as block size increases. Random writes behave differently and stabilize around 110-117 GB/s, which is normal for mixed-queue workloads that introduce additional overhead.
The main takeaway from this section is that the R7725xd effortlessly maintains extremely high bandwidth via NVMe-oF, even with multiple clients heavily loading the system. Once block size reaches 32 KB or more, the server consistently saturates its available network and storage bandwidth. This is precisely the type of performance PEAK:AIO is designed to achieve, confirming the platform's ability to scale under real-world conditions.
PEAK AIO – NVMe-oF IOPS RDMA
On the IOPS side, the PowerEdge R7725xd displays excellent performance for small blocks, despite results initially below expectations; this issue is expected to be resolved with better network driver support. Despite this, the overall scaling trend corresponds exactly to the typical behavior of NVMe-oF RDMA as block size increases.
With the smallest block size, the system can deliver over 6 million IOPS for sequential and random workloads. Read, write, random read, and random write performance remain substantially the same at 4K and 8K, indicating that the front-end clients, PCIe infrastructure, and NVMe drives themselves can easily keep pace with request rates.
As block size increases, the expected decline in IOPS begins. At 32 KB, reads reach about 4.7 million IOPS, while writes follow slightly behind at about 4.4 million. Random writes are the most affected, dropping to about 3.3 million IOPS, which corresponds to the additional queue and processor overhead induced by mixed access.
With large block sizes, IOPS continue to decrease linearly and predictably. From 256,000 and 512,000 transfers, throughput becomes the predominant criterion and IOPS naturally drop to a few hundred thousand. For a 1 MB block size, all workloads converge to 140,000 to 153,000 IOPS, which corresponds to the bandwidth values observed in the previous section.
GPUDirect Storage Performance
One of the tests we performed on the R7725xd was the Magnum IO GPUDirect Storage (GDS) test. GDS is a feature developed by NVIDIA that allows GPUs to bypass the CPU when accessing data stored on NVMe drives or other high-speed storage devices. Instead of routing data through the CPU and system memory, GDS enables direct communication between the GPU and the storage device, significantly reducing latency and improving data throughput.
How GPUDirect Storage works
Traditionally, when a GPU processes data stored on an NVMe drive, that data must pass through the CPU and system memory before reaching the GPU. This process creates bottlenecks, with the CPU acting as an intermediary, which increases latency and consumes valuable system resources. GPUDirect Storage eliminates this inefficiency by allowing the GPU to access data directly from the storage device via the PCIe bus. This direct path reduces the overhead associated with data transfers, enabling faster and more efficient transfers.
AI workloads, particularly those involving deep learning, are extremely data-hungry. Training large neural networks requires processing terabytes of data, and any delay in data transfer can lead to GPU underutilization and longer training times. GPUDirect Storage addresses this challenge by ensuring that data is delivered to the GPU as quickly as possible, minimizing idle time and maximizing computational efficiency.
Furthermore, GDS is particularly useful for workloads involving the streaming of large datasets, such as video processing, natural language processing, or real-time inference. By reducing dependence on the processor, GDS accelerates data movement and frees processor resources for other tasks, further improving overall system performance.
Beyond raw bandwidth, GPUDirect with NVMe-oF (TCP/RDMA) also provides very low-latency I/O. Thus, GPUs never starve for data, making it the ideal system for real-time AI inference, analytics pipelines, and video replay.
GDSIO Sequential Read
When analyzing PEAK:AIO with a client using GDSIO, read throughput shows a clear correlation with increasing block size and thread count. This single client was connected via two 400G links, limiting its total potential to 90 GB/s.
With small block sizes and a low thread count, performance remains modest: 4K reads reach about 189 MiB/s with a single thread. As soon as the thread count increases, the system responds instantly, reaching 691 MiB/s with four threads and exceeding GiB/s for larger blocks.
Intermediate block sizes show the greatest sensitivity to thread count. At 32,000 threads, throughput goes from 1.3 GiB/s with a single thread to nearly 20 GiB/s with 64 threads, with a slight decrease beyond that. Similar behavior is observed at 64,000 and 128,000 threads: the system goes from a few GiB/s at low parallelism to over 30 GiB/s as the workload increases.
Once larger block sizes are reached, throughput stabilizes, with the system approaching its maximum performance for a single client. At 1 MiB, performance goes from 11 GiB/s with a single thread to about 88 GiB/s with a high thread count. The 5 MiB and 10 MiB transfers show the same plateau, peaking around 89-90 GiB/s, whether the test is run with 64, 128, or 256 threads.
GDSIO Sequential Write
In writes, scaling behavior follows a trend similar to reads, but with slightly lower performance for most block sizes, which is normal for sequential write workloads. For the smallest block sizes, initial throughput is 165 MiB/s for a single thread at 4 KB and increases steadily with parallelism. At four threads, it reaches just over 619 MiB/s before exceeding 1 GiB/s at eight threads.
Intermediate block sizes show more pronounced gains as thread count increases. At 32 K, initial throughput is slightly below 1 GiB/s and exceeds 21 GiB/s at higher thread levels. The 64 K and 128 K ranges confirm this trend, going from a few GiB/s to about 30 GiB/s, or even 50 GiB/s, as the workload becomes more parallelized.
