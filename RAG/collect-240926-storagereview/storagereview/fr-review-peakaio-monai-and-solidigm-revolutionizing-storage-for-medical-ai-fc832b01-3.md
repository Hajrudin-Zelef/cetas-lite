---
id: collect-240926-storagereview/storagereview/fr-review-peakaio-monai-and-solidigm-revolutionizing-storage-for-medical-ai-fc832b01-3
title: "fr-review-peakaio-monai-and-solidigm-revolutionizing-storage-for-medical-ai-fc832b01"
domain: storagereview
role: reference
task: reference
actors: ["Nvidia", "Oracle"]
dates: []
keywords: ["gpu", "nvidia", "regulation", "research"]
source: docs/RAG/clean_en/storagereview/fr-review-peakaio-monai-and-solidigm-revolutionizing-storage-for-medical-ai-fc832b01.md
source_anchor: ""
source_lines: [25, 51]
sha256: 2b517541c80370680f8988ece9d764b031f1b75afc6fda2415c5f210cee1097e
---

# fr-review-peakaio-monai-and-solidigm-revolutionizing-storage-for-medical-ai-fc832b01

The collaboration between PEAK:AIO and Solidigm continues to evolve, including ongoing innovations such as the use of Single Layer Cell (SLC) cache in PEAK:AIO's Gen 2 system to optimize write performance, reduce write amplification, and ensure long-term reliability of AI workloads. The partnership provides medical AI applications with tailored storage solutions that handle the growing complexity of data and the need for scalability without compromising performance.
Enhancing MONAI deployments with Solidigm and PEAK:AIO
Although MONAI operates primarily as a software layer on GPU servers and is not directly tied to storage solutions, it (like many complex software platforms) relies heavily on high-performance storage to operate optimally. A useful analogy provided by PEAK:AIO from traditional computing environments would be Oracle Database, which requires robust servers and a high-performance SAN for optimal operation. Although the SAN itself is not part of the Oracle DB offering, its design and infrastructure are essential to ensuring the database's superior performance.
Extending this concept to the partnership between PEAK:AIO and Solidigm, they have carefully crafted a solution to enhance MONAI deployments by providing storage that is not only suited to MONAI's I/O-intensive requirements but also addresses larger challenges related to capacity, power consumption, and scalability. As AI projects grow in scale and complexity, these storage solutions become essential to ensuring that MONAI can handle the vast amounts of data in medical imaging and AI-driven healthcare workloads.
PEAK:AIO's collaboration with Solidigm enables MONAI deployments to overcome the storage obstacles faced by AI projects in healthcare. Although MONAI is just one of many AI frameworks in which PEAK:AIO's high-performance storage solutions add significant value, it serves as an excellent case study of integration performance in a real AI vertical. In many ways, MONAI is representative of the types of challenges that PEAK:AIO regularly faces in AI-focused sectors, where efficient storage, low power consumption, and scalability are paramount to success.
PEAK:AIO Storage System: Feature Overview
| Feature | DETAILS | 
|---|---|
| Capacity / Drives / Protection | 30 TB – 1.4 PB RAW (per typical 2U) Hard drives of 7.68 TB/15.36 TB/30.72 TB/61.44 TB PEAK:PROTECT, RAID 0, 10, 5, 6, N+2 | 
| Performance (bandwidth) | 80 GB/s with 2xCX-7 40 GB/s with CX-6 Up to 160 GB/s max per 2U (scalable per 2U) | 
| Protocols | NFS3/4 (RDMA/TCP) NVMe-oF (RDMA/TCP) NVIDIA GPUDirect® | 
| Interface | QSFP6/112 (NVIDIA ConnectX-6/7) 200 Gb/400 Gb (IB/ETH) per CX-6/7 Recommended: 2 x CX-X Supports up to 6 CX-8 compatible (800 Gb per port) | 
The PEAK:AIO storage system offers exceptional scalability, performance, and reliability, making it ideal for demanding AI-focused healthcare environments. System capacities range from 30 TB to 1.4 PB of raw storage in a standard 2U chassis, offering flexibility that meets the varying data needs of medical institutions. Drives are available in capacities of 7.68 TB, 15.36 TB, 30.72 TB, and 61.44 TB, offering a range of configurations to meet different workload requirements. Additionally, PEAK:PROTECT supports multiple RAID levels (0, 5, 6, 10) and N+2 redundancy, ensuring data integrity and resilience.
Performance is a standout feature of the PEAK:AIO system, with bandwidth capabilities reaching 80 GB/s with 2 CX-7 adapters and 40 GB/s with CX-6 adapters. The system can achieve up to 160 GB/s per 2U chassis, increasing performance as units are added. These performance figures are remarkably achievable with a single host, demonstrating the system's efficiency and power.
When it comes to connectivity, PEAK:AIO supports advanced protocols such as NFS3/4 over RDMA or TCP, NVMe-oF over RDMA or TCP, and NVIDIA GPUDirect®, ensuring compatibility with high-performance computing environments. The system interfaces via QSFP6/112 with NVIDIA ConnectX-6 or ConnectX-7 adapters, supporting 200 Gb or 400 Gb connections, making it scalable and ready for the most demanding applications. The recommended configuration includes two CX-X adapters, although the system supports up to six, offering unmatched flexibility and scalability.
Initial use of Solidigm in AI-focused storage solutions
The collaboration between PEAK:AIO and Solidigm began with a targeted effort on optimizing read performance for AI applications. During a presentation sponsored by SK Hynix with EMC3, the potential of using Solidigm's technology for AI-specific storage solutions was discussed. At the time, Solidigm wanted to succeed in the AI sector, and PEAK:AIO was well-positioned to achieve this, a goal it continues to pursue.
Initially, the focus was on developing an AI-specific archiving solution that could meet the evolving needs of AI-focused research and applications. As AI models are increasingly created based on their long-term usefulness, there is a growing need to be able to trace decisions and models through the various stages of development. This need is particularly evident in projects undertaken by AI centers or NHS trials.
Through close collaboration with Professor Sébastien Ourselin, PEAK:AIO developed the PEAK:ARCHIVE solution. This solution, designed to meet AI-specific archiving needs, is evolving into what will be known as PEAK:AUDIT. When these intelligent systems are merged, this new model will help ensure compliance with strict healthcare standards and regulations.
Progress with Gen 2 PEAK:AIO QLC
Building on the success of the initial deployment, PEAK:AIO continued to refine its storage solutions, notably with the development of the PEAK:AIO QLC Gen 2 system. Although this project was not made public, significant advances were made, particularly in write path optimization. This includes the integration of an SLC cache, which is more than just a front-end buffer. For example, large sequential writes can be directed straight to QLC. In contrast, smaller random writes are staged on the SLC, enabling structured data management that reduces write amplification and improves overall write performance.
Although the Gen 2 system is not yet production-ready, the underlying technology has been licensed to a cloud provider for over two years and has proven stable and efficient. The current focus is on transitioning this technology, which simply accelerates front-end access to hard drives (HDDs) in cloud environments, toward creating a more dedicated QLC front-end. This approach would operate at the block layer, intercepting writes while ensuring that read operations remain direct, unless the data resides on the SLC. There is also consideration of offering this technology as a cloud service in the future. However, privacy regulation faces obstacles.
PEAK:AIO: Powering medical AI with scalable storage solutions
PEAK:AIO specializes in high-capacity storage systems optimized for medical AI applications. These solutions, powered by Solidigm QLC SSDs, offer a balanced combination of capacity, performance, and efficiency. They are specifically designed to meet the data requirements of modern healthcare institutions while ensuring secure and immediate access to patient information.
Designed for the future of medical AI, the PEAK:AIO platform is scalable to accommodate the exponential growth of health data. This scalability ensures that, as data loads increase, the system maintains consistent performance and efficiency, preparing hospitals and research institutes for future challenges.
PEAK:Archive: Optimized for medical data
PEAK:AIO's flagship offering is PEAK:ARCHIVE, a solution designed specifically for long-term data storage. While most storage solutions focus on immediate data retrieval and processing, PEAK:ARCHIVE is optimized for read performance, making it superior for storing historical medical data that is not often accessed but is nonetheless essential.
