---
id: collect-240926-storagereview/storagereview/fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f-3
title: "fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f"
domain: storagereview
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["compute", "cost", "gpu", "gpus", "inference", "latency", "memory", "nvidia", "throughput", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f.md
source_anchor: ""
source_lines: [32, 63]
sha256: 33e204334f8dabedf3feb3ca75385d4f7416ba7a50ddb13e0d2a786c1213b8f8
---

# fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f

| Random read (4 KB) | 46.6 GB / s | 11.4M | 0.269 ms |
| Random write (4K) | 29.1 GB / s | 7.1M | 0.432 ms |
GPU Direct Storage
One of the tests we conducted on this test bench was the Magnum IO GPU Direct Storage (GDS) test. GDS is a feature developed by NVIDIA that allows GPUs to bypass the CPU when accessing data stored on NVMe drives or other high-speed storage devices. Instead of routing data through the CPU and system memory, GDS enables direct communication between the GPU and the storage device, significantly reducing latency and improving throughput.
How GPU Direct Storage works
Traditionally, when a GPU processes data stored on an NVMe drive, the data must first pass through the processor and system memory before reaching the GPU. This process introduces bottlenecks, as the processor becomes an intermediary, adding latency and consuming valuable system resources. GPU Direct Storage eliminates this inefficiency by allowing the GPU to access data directly from the storage device via the PCIe bus. This direct path reduces the overhead associated with moving data, enabling faster and more efficient data transfers.
AI workloads, particularly those involving deep learning, are highly data-intensive. Training large neural networks requires processing terabytes of data, and any delay in data transfer can lead to GPU underutilization and longer training times. GPU Direct Storage addresses this challenge by ensuring that data is delivered to the GPU as quickly as possible, minimizing idle time and maximizing compute efficiency.
Furthermore, GDS is particularly useful for workloads involving the streaming of large datasets, such as video processing, natural language processing, or real-time inference. By reducing reliance on the processor, GDS accelerates data movement and frees up processor resources for other tasks, further improving overall system performance.
GDSIO sequential read
In the GDSIO sequential read test of the 7.68 TB Solidigm PS1010 drive, performance scaled significantly with block size and I/O depth. With the smallest block size of 16 KB, throughput started at only 0.2 GiB/s with a queue depth of 1, then gradually increased to 1.3 GiB/s at a depth of 128 KB, showing limited scalability at this granularity. With 128 KB blocks, performance scaled more dramatically, starting at 1.1 GiB/s and reaching 6.5 GiB/s at the highest depth. The best results were obtained with a block size of 1 MB, where throughput initially reached 2.4 GiB/s and peaked at 8.5 GiB/s with a queue depth of 128 KB, indicating that the drive's optimal performance profile is achieved with large sequential reads and deeper queues.
GDSIO sequential write
The sequential write performance of the Solidigm PS1010 shows solid scalability for larger block sizes, but shows a slight regression at higher queue depths in medium-sized workloads. With the smallest block size of 16 KB, write speeds started at 0.5 GiB/s and peaked modestly at 0.9 GiB/s between queue depths of 8 to 64, before dropping slightly to 0.8 GiB/s at a depth of 128. With 128 KB blocks, performance started at 2.2 GiB/s, peaked at 4.3 GiB/s at a depth of 32, then dropped to only 1.9 GiB/s at the highest queue depth, indicating saturation or a potential write limitation. The best sustained performance was obtained with block sizes of 1 M, where throughput scaled markedly from 4.1 GiB/s at a depth of 1 to 5.6 GiB/s at depths of 32 and 64, remaining stable up to a depth of 128.
GDSIO summary
This table provides a detailed analysis of the GDSIO performance measurements for latency and IOPS collected on the Solidigm D7-PS1010 SSD, measured for block sizes of 16 KB, 128 KB, and 1 MB with an I/O depth of 128. With a queue depth of 128, latency and IOPS scale predictably with block size. Reading a 16 KB block averaged 1.549 ms with 82.3 K IOPS, while write latency was 2.429 ms with 52.6 K IOPS. With 128 KB, read latency reached 2.414 ms (52.9 K IOPS) and write latency 8.050 ms (15.9 K IOPS). At 1 MB, read latency reached 14.643 ms with 8.7 K IOPS, and write latency reached 23.030 ms with 5.6 K IOPS.
| GDSIO chart (average block sizes 16, 128, 1 K and 1 M) | HPE DL145 Gen 11 (6 Solidigm D7-PS1010 E3s 7.68 TB SSDs) |
|---|---|
| (16 KB block size, I/O depth 128) Average read | 1.3 GiB/s (1.549 ms) IOPS: 82.3 K |
| (16 KB block size, I/O depth 128) Average write | 0.8 GiB/s (2.429 ms) IOPS: 52.6 K |
| (128 KB block size, I/O depth 128) Average read | 6.5 GiB/s (2.414 ms) IOPS: 52.9 K |
| (128 KB block size, I/O depth 128) Average write | 1.9 GiB/s (8.050 ms) IOPS: 15.9 K |
| (1 M block size, I/O depth 128) Average read | 8.5 GiB/s (14.643 ms) IOPS: 8.7 K |
| (1 M block size, I/O depth 128) Average write | 5.4 GiB/s (23.030 ms) IOPS: 5.6 K |
Cutting-edge media production
What is Axle AI?
Axle AI is an on-site, AI-optimized media asset management (MAM) platform that simplifies video workflows for small and medium-sized media teams. It automates key tasks, such as ingesting, tagging, and searching media, without the cost or complexity of traditional MAM systems. Designed to be fast and easy to use, Axle AI enables teams to manage and distribute content more efficiently, whether they are in the office or working remotely.
The platform supports the entire production process. It automatically ingests footage, creates proxies, applies AI-based metadata through features such as object recognition, face detection, and semantic search, and integrates with standard industry editing tools, including Adobe Premiere Pro and DaVinci Resolve. Teams can thus quickly locate, edit, and publish content, without bottlenecks or delays.
Real-world use cases
Axle AI is trusted by a wide range of media teams, including broadcasters, documentary filmmakers, corporate video departments, marketing agencies, and live event producers. These teams often work across multiple sites and need reliable access to their media libraries without requiring heavy infrastructure.
With Axle AI, users can collaborate remotely through a web interface that provides instant access to proxy media. Editors can begin editing footage immediately in their preferred non-linear editing (NLE) systems, while producers and stakeholders review, tag, or approve clips in real time (no large file transfers or special technical skills required).
Whether you are a team of five or fifty, Axle AI enables faster, more efficient workflows. Its simple interface, rapid deployment, and compatibility with existing storage systems make it a smart solution for high-volume distributed media production.
The strong demand for instant social media has significantly accelerated the pace of media production. The idea is widespread: "If you're not first, you're last," especially when it comes to live events and the content creation surrounding them. Media companies needed an efficient way to quickly organize and review large amounts of footage. Tools like Axle AI, with their ability to scale to small operations, make this workflow practical in the field. Combined with a suitable workflow, such as cameras capable of simultaneously recording proxies and raw footage, the possibilities for efficiency increase.
Recording high- and low-resolution proxy media at the time of capture and immediately processing smaller proxy files in Axle AI Tags makes it possible to search all footage almost instantly. Editors can get to work immediately thanks to Axle AI's integration with most major non-linear editing software such as Adobe Premiere Pro, DaVinci Resolve, and Avid Media Composer. This also significantly lightens the production team's workload, allowing them to focus on other tasks or avoid burnout.
