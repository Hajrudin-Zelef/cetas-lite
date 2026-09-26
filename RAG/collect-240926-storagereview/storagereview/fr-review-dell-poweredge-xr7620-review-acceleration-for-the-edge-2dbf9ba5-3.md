---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5-3
title: "fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5"
domain: storagereview
role: reference
task: reference
actors: ["Broadcom", "Intel", "Nvidia"]
dates: []
keywords: ["accelerator", "benchmark", "gpu", "gpus", "inference", "intel", "latency", "memory", "mlperf", "nvidia", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5.md
source_anchor: ""
source_lines: [53, 106]
sha256: 5dc90e42eece6182f90c391de29a0348345fd3285b4068beaee053bb2ba0a55f
---

# fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5

One of its main functions is to provide a visual interface for monitoring and managing server status and performance. With built-in LED indicators and a clear, user-friendly display, it provides real-time binary information on overall system status, temperature, power consumption, etc. This feature is invaluable for IT professionals, as it allows them to quickly identify servers and troubleshoot issues, ensuring uninterrupted server operation.
Additionally, the XR7620 smart bezel is a security asset. It helps protect sensitive data and the server by providing physical protection against tampering and unauthorized access. Locking mechanisms and intrusion detection add an extra layer of defense to the server infrastructure.
Another notable feature of the XR7620 smart bezel is its filtering capability. This filtering mechanism goes beyond aesthetics; it actively contributes to the server's performance and reliability. The bezel includes strategically placed air filters that help maintain a clean, dust-free environment within the server chassis. These filters help prevent dust and other large particles from entering the server, which is particularly crucial in harsh environments where the server may be exposed to airborne particles.
By keeping internal components free of dust and contaminants, the filtering mechanism extends the lifespan of critical server components such as fans, heat sinks, and circuits. It also helps maintain optimal airflow for cooling, ensuring the server operates efficiently with minimal risk of overheating.
The XR7620 smart bezel's filtering mechanism is a key element that enhances server reliability by keeping internal components clean and well-ventilated. This feature plays an important role in ensuring the longevity and proper operation of the server in various operating environments.
The PowerEdge XR7620 is designed to operate reliably in harsh conditions, from extreme temperatures and humidity to dusty environments and shocks, and is certified to NEBS Level 3, GR-3108 Class 1, and MIL-STD-810G.
The XR7620 improves upon the previous PowerEdge XR2 and XE2420 servers, with similar core features and the latest components, including:
- A processor upgrade to the recently announced 4th Generation Intel Xeon Scalable processor, with up to 32 cores.
- 2x the memory bandwidth with the upgrade from DDR4 to DDR5
- Higher-performing I/O capabilities with the upgrade from PCIe Gen 4 to PCIe Gen 5, with 5 PCIe slots.
- Improved storage capabilities with up to 8 NVMe drives, BOSS support, and hardware-based NVMe RAID.
- Dense acceleration capabilities at the edge where the XR7620 excels, with support for up to 2 double-width (DW) accelerators up to 300 W each or four single-width (SW) accelerators up to 150 W each. Smart filtered bezel for working in dusty environments
Dell PowerEdge XR7620 Performance
Our evaluation system is configured with the following key components for testing:
- 2 Xeon Gold 6426Y – 16 cores 2.5 GHz
- Nvidia L4
- 8 x 16GB DDR5
- BOSS RAID480 1 GB
- Dell PERC H755 PERC11
- 1400 XNUMX W power supplies
- Broadcom 2 x 10 Gb network card
- Intel 4 x 10 Gb network card
Geekbench 6 is a cross-platform benchmark that measures system performance and provides a score for comparison purposes. It is designed to run on multiple platforms and provides a consistent performance measurement across many devices, from smartphones and tablets to desktops and servers.
| Geekbench6 Multi | 16,507 | 
| Geekbench6 Single | 1,864 | 
Cinebench is a widely used benchmarking tool that measures CPU and GPU performance using Maxon Cinema 4D for rendering. It provides a score that can be used to compare the performance of different systems and components. We have included both R23 and 2024 results.
| Cinebench R23 Multi | 46,544 | 
| Cinebench R23 Single | 1,292 | 
| Cinebench 2024 Multi | 2,383 | 
| Cinebench 2024 Single | 73 | 
y-cruncher 0.8.3.9522 is a multithreaded and scalable program that can calculate Pi and other mathematical constants to billions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application among overclockers and hardware enthusiasts.
| y-cruncher 1b (1 billion digits) | 9.805s | 
| y-cruncher 2.5b (2.5 billion digits) | 27.344s | 
| y-cruncher 5b (5 billion digits) | 60.564s | 
| y-cruncher 10b (10 billion digits) | 133.100s | 
| y-cruncher 25b (25 billion digits) | 383.261s | 
The Blender benchmark measures the 3D rendering performance of a CPU or GPU by rendering a 3D scene in the Blender software. It provides a score that can be used to compare the performance of different systems and components.
|  | Blender 3.6 CPU CLI only | Blender 4.0 CPU CLI only | 
| Monster | 333.737716 | 312.455545 | 
| Junkshop | 216.598456 | 213.333376 | 
| Classroom | 161.74015 | 161.812143 | 
| Total | 657.218016 | 687.601064 | 
GPU Acceleration
The XR7620 supports up to 2 x 300 W GPU accelerator cards to handle demanding workloads, delivering image classification 45% faster than the Dell XR 12 server with a single 300 W GPU accelerator. The combination of low latency and high processing power enables faster and more efficient data analysis, allowing organizations to make real-time decisions for increased monetization.
To test inference performance, we ran the MLPerf 3.1 inference, both offline and on the server. BERT (Bidirectional Encoder Representations from Transformers) is a transformer-based model primarily used for natural language processing tasks such as question answering, language understanding, and sentence classification. ResNet50 is a convolutional neural network (CNN) model widely used for image classification tasks. It is a variant of the 50-layer ResNet model, known for its deep architecture but efficient performance.
In order to test MLPerf 3.1, we equipped our XR7620 with a single NVIDIA L4 GPU. The L4 is a very high-performance card, designed specifically for AI and deep learning workloads. With a huge 24 GB of VRAM on a power-efficient 70-watt half-height, half-length card, the L4 is a perfect accelerator for this rugged Edge server. The L4, which is part of NVIDIA's latest generation of Ada Lovelace data center GPUs, is designed to deliver exceptional performance in machine learning tasks, making it an ideal choice for our benchmarking objectives.
| Resnet50 – Offline: | 13,010.2 | 
| Resnet50 – Server: | 12,204.4 | 
| Bert K99 – Offline: | 973.465 | 
| Bert K99 – Server: | 898.945 | 
- Offline mode: This mode measures system performance when all data is available for simultaneous processing. This is similar to batch processing, where the system processes a large dataset in a single batch. This mode is crucial for scenarios where latency is not a major concern, but throughput and efficiency are.
- Server mode: In contrast, server mode evaluates system performance in a scenario that mimics a real server environment, where requests arrive one by one. This mode is latency-sensitive and measures how quickly the system can respond to each request. This is crucial for real-time applications where an immediate response is needed, such as in web servers or interactive applications.
Our XR7620, equipped with two 4th Generation Xeon Scalable Gold processors, handled the L4 cache perfectly, with results consistent with those validated by NVIDIA during its MLPerf 3.1 inference test. As our MLPerf 3.1 validations are still ongoing, this test will serve as a reference for our future tests.
Conclusion
