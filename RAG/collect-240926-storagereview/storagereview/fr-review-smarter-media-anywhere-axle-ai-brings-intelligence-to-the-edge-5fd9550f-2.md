---
id: collect-240926-storagereview/storagereview/fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f-2
title: "fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "Samsung"]
dates: []
keywords: ["intel", "accelerator", "amd", "benchmark", "compute", "gpu", "gpus", "inference", "latency", "memory", "nvidia", "parameters"]
source: docs/RAG/clean_en/storagereview/fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f.md
source_anchor: ""
source_lines: [3, 31]
sha256: f39e3a4e38abc5c58b4a000fccbd5dc97061a31a3ff42352d13ce858cdc41364
---

# fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f

The demand for fast, high-quality video content has never been greater. In today's media landscape, small press teams, independent content creators, and event production teams must capture, edit, and publish professional-grade footage in real time, often without the support of a post-production team or centralized infrastructure.
This is why on-site AI search, powered by edge production systems, is becoming essential, given the explosion in footage volume and the growing need to reuse it. By enabling teams to work directly with footage at the moment of capture, edge workflows eliminate bottlenecks and allow creators to spend less time searching for and managing footage, and more time producing the best possible content.
Axle AI, an AI-based media asset management and automation platform, is purpose-built for decentralized, dynamic workflows. Combined with cutting-edge hardware, such as the HPE ProLiant DL145 Gen11 server, an NVIDIA L4 GPU, and Solidigm PCIe Gen5 SSDs, it delivers a powerful, portable production environment. It brings professional AI-assisted media capabilities within reach of even the smallest teams.
System architecture overview
The HPE ProLiant DL145 Gen11 edge server is an ideal candidate for Axle AI deployments in the field and at events where space, power, and cooling are limited. With its 2U short-depth design (only 16 cm), this server is purpose-built for portability and installation in confined spaces such as mobile racks and flight cases. Unlike traditional professional equipment, it supports an extended operating temperature range of -5 °C to 55 °C, making it suitable for less controlled environments. Despite its compact size, it offers professional features such as redundant power supplies and boot drive options, making it ideal for teams that need robust edge infrastructure without the overhead of standard hardware.
On the processor side, the DL145 supports AMD EPYC 8004 series (Siena) processors, compatible with 8 to 64 cores. Our test system is equipped with an EPYC 8434P, offering 48 cores and 96 threads. With a TDP of 200 W, it delivers an excellent balance between efficiency and compute power, making it ideal for intensive workloads. This processor supports 96 PCIe 5.0 lanes and offers six DDR5 memory channels, enabling speeds up to 4,800 MT/s with excellent memory bandwidth. While the minimum configuration required for Axle AI is 16 cores and 32 GB of RAM, our configuration far exceeds these requirements with 48 cores and 256 GB of DDR5 memory.
Storage performance is critical for any media pipeline, and the DL145 supports up to six NVMe EDSFF E3.S drives to meet these demands. In our configuration, we used six Solidigm D7-PS1010 E3.S SSDs, each with a capacity of 7.68 TB. These PCIe 5.0 x4 drives deliver read speeds of up to 14,500 MB/s and write speeds of up to 10,000 MB/s, for high-throughput storage of nearly 46 TB. Designed for AI and media workflows, these SSDs offer impressive power efficiency and can deliver up to 50% higher throughput than conventional NVMe drives in certain phases of the pipeline. For teams requiring additional capacity, this configuration can scale up to 15.36 TB per drive, providing fast, dense storage of 92 TB in a compact format.
The DL145 board supports up to three single-slot GPUs. In our configuration, we used a single NVIDIA L4 GPU, a low-power accelerator based on the Ada Lovelace architecture. Designed for video processing, AI inference, visual computing, and virtualization workloads, the L4 delivers robust performance while consuming only 72 watts. It features 24 GB of GDDR6 VRAM and connects via a PCIe 4.0 x16 interface for power and data. Axle AI recommends a minimum of 16 GB of VRAM, and the L4 far exceeds this requirement with 8 GB extra, ensuring sufficient headroom for high-resolution media tasks and AI-based operations.
For the application layer, we deployed Axle AI MAM and Axle AI Tags using the Proxmox virtual environment (8.3.5). The system is configured with two dedicated virtual machines: one running Axle AI MAM and the other Axle AI Tags. This structure ensures a clear separation between media management and AI metadata processing, while ensuring smooth communication between components.
Running in parallel with MAM, Axle AI Tags provides on-site AI capabilities, including vector semantic search, object and logo recognition, and trainable face detection. Deployed in a Docker container, it leverages the NVIDIA L4 GPU via PCIe pass-through for efficient, real-time inference. The browser-based training and administration interface allows for flexible tuning, and its modular design supports Intel/NVIDIA or AMD/NVIDIA hardware configurations. Fully integrated with MAM in our configuration, Axle AI Tags also offers REST API compatibility for use with other platforms.
The Solidigm PS1010 drives are integrated into the Axle AI MAM virtual machine, ensuring fast, low-latency access to footage. By virtualizing the entire stack and allocating dedicated compute resources to each component, this configuration provides increased availability, better resource management, and far superior multi-user support compared to a single workstation. It is therefore ideal for a high-volume collaborative production environment.
Together, these components form a tightly integrated hardware solution that enables fast, local, AI-optimized media workflows. It is an efficient, ready-to-use platform for remote teams that need reliable performance in environments where space and power are limited, but where AI processing remains essential.
HPE ProLiant DL145 Gen11 performance
Before moving on to the benchmark analysis, the table below presents the system configuration of the HPE ProLiant DL145 Gen11. The tests with Axle AI were performed on a virtualized installation running Proxmox, and the system was migrated to Ubuntu Server 22.04.5 in order to measure storage performance during our GDSIO and FIO tests.
| HPE ProLiant DL145 generation 11 | Hardware overview |
|---|---|
| Processor | Single AMD EPYC 8434P processor |
| RAM | 256 GB of ECC DDR5 memory |
| Storage | 6 Solidigm D7-PS1010 E3s 7.68 TB SSDs |
| Boot storage | NS204 2x populated port (480 GB Samsung PM9A3 M.2) |
| GPU | Single NVIDIA L4 |
| Operating system | Ubuntu Server 22.04.5 |
Peak synthetic performance
The FIO test is a flexible and powerful benchmarking tool for measuring the performance of storage devices, including SSDs and hard drives. It evaluates parameters such as bandwidth, IOPS (input/output operations per second), and latency under various workloads, such as sequential and random read/write operations. This test is designed to capture peak performance and subject the storage system to multiple workloads, making it particularly useful for comparing different devices or configurations. In this case, a full-surface test was performed, testing the total capacity of the drives in order to obtain a complete view of their sustained performance.
In the sequential read test using 128 K blocks, the system delivered a bandwidth of 56.4 GB/s and sustained 430,000 IOPS with an average latency of 1.78 milliseconds. Sequential write performance at the same block size reached 45.4 GB/s, producing 346,000 IOPS and an average latency of 2.22 milliseconds. For random read operations using 4 K blocks, the system reached 46.6 GB/s with 11.4 million IOPS and a low average latency of 0.269 milliseconds, highlighting the high-throughput potential of the NVMe storage array under intensive access conditions. Random write operations at 4 K measured 29.1 GB/s, reaching 7.1 million IOPS with an average latency of 0.432 milliseconds, confirming robust sustained write capability, even under fragmented access.
| HPE DL145 Gen 11 FIO benchmark summary | Bandwidth – GB/s | IOPS | Average latency |
|---|---|---|---|
| Sequential read (128 KB) | 56.4 GB / s | 430K | 1.78 ms |
| Sequential write (128 KB) | 45.4 GB / s | 346K | 2.22 ms |
