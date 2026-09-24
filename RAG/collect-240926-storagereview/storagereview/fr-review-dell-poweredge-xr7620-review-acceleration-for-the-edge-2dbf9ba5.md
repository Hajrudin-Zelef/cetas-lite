---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5
title: "fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5"
domain: storagereview
role: reference
task: reference
actors: ["Broadcom", "Intel", "Nvidia"]
dates: []
keywords: ["accelerator", "benchmark", "compute", "energy", "exploit", "gpu", "gpus", "inference", "intel", "latency", "memory", "mlperf"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5.md
source_anchor: ""
source_lines: [1, 109]
sha256: ef707eb2b7df30cd7a11b4d0e02d481168c8b79f7aca241b8ed7a4508742695f
---

# fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5

<!-- source: https://www.storagereview.com/fr/review/dell-poweredge-xr7620-review-acceleration-for-the-edge -->

StorageReview has tested several ruggedized servers, focusing on their performance, capacity, and ease of use. We recently tested the Cheetah RAID solution equipped with Solidigm SSDs. Unfortunately, for obvious reasons, we couldn't test the resilience of these servers. However, Dell has raised the bar even higher by providing an Edge PowerEdge XR7620 server and a rugged transport container, capable of withstanding extreme conditions.
Life is hard at the edge
Rugged edge infrastructure is not a new model. Edge servers have existed for quite some time; they collect data and send it back to headquarters to exploit its full potential. Artificial intelligence has emphasized faster data processing for near-immediate use, which involves processing at the source. Edge AI has had a considerable impact, bringing more innovative, faster, and more secure processing to many concrete applications. A personal assistant always within reach, making real-time decisions, saving you time, money, and sometimes even lives!
Dell PowerEdge XR7620 Configuration
Dell introduced these specially designed PowerEdge servers last year. The model we are examining is the Dell PowerEdge XR7620, a rugged dual-socket server, optimized for the edge, short-depth, specially designed and compact, offering acceleration-focused solutions for the edge. The difference between the XR7620 and classic Edge servers is that this server addresses the rapid maturation of AI/ML with support for the most demanding workloads, including industrial automation, video analytics, point-of-sale analytics, AI inference, and Edge device aggregation.
Our XR7620 is configured with 2x Intel Xeon Gold 6426Y, 128 GB DDR5 4800, 480 GB Dell BOSS RAID1, and 4x Solidigm SSDs. Our tests will take place in Jordan's environmental research laboratory in Canada. So we purchased a 60 TB drive to send the data back to the Storagereview laboratory in Cincinnati.
Dell PowerEdge XR7620 Specifications
| Feature | Technical Specifications | 
|---|---|
| Processor | Two 4th Generation Intel® Xeon® Scalable processors with up to 32 cores per processor | 
| Memory | 16 DDR5 DIMM slots, supports RDIMM 1 TB maximum, speeds up to 4800 5 MT/s. Supports only registered DDRXNUMX ECC DIMMs | 
| Storage Controllers |  | 
| Drive Bays | Front bays: up to 4 x 2.5-inch SAS/SATA/NVMe SSDs, 61.44 TB maximum, up to 8 x E3.S direct NVMe drives, 51.2 TB maximum | 
| Power Supplies |  | 
| Cooling Options | air cooling | 
| Fans | Six hot-swappable cooling fans | 
| Dimensions |  | 
| Weight | Maximum 21.16 kg (46.64 lbs) | 
| Form Factor | 2U rack server | 
| Integrated Management | iDRAC9, iDRAC Direct, iDRAC RESTful API with Redfish, iDRAC Service Module | 
| Bezel | Optional security bezel with dust filter (dust sensor available only for front-access configuration systems) | 
| OpenManage Software |  | 
| Mobility | OpenManage Mobile | 
| OpenManage Integrations |  | 
| Security |  | 
| GPU Options | Up to 5 x 75 W (single-width, full-height/half-length, low-profile) GPUs or up to 2 x 300 W (double-width, full-height/full-length) | 
| Integrated Network Card | 2 x 1 GbE LOM | 
| Network Options | 1 x OCP 3.0 card (optional) | 
| Ports |  | 
| PCIe | 2-processor configuration: up to 5 PCIe slots (4 x16 Gen4/5, 1 x16 LP Gen4) | 
| Operating Systems and Hypervisors |  | 
Dell PowerEdge XR7620 – Critical Edge Features
Security
Like all other PowerEdge servers, the XR7620 is designed with security in mind at every stage of the server lifecycle. The security process begins with a secure supply chain before the server is built. It extends to delivered servers, with Secure Lifecycle Management and Silicon Root of Trust, then secures what is created/stored by the server in Data Protection.
This zero-trust security approach assumes privileged access authorizations requiring validation at every access and implementation point with features such as Identity and Access Management (IAM) and multi-factor authentication (MFA). This is essential for edge deployments where servers are typically installed in unsecured environments. This strict security policy provides the ability to detect and react to tampering or intrusion. Dell's silicon-based Root of Trust platform creates a secure environment that ensures firmware comes from a trusted source. Any detected change in firmware versions or configurations can force PowerEdge to lock the configuration and initiate a restore to the last known good environment.
Management
All Dell PowerEdge servers follow a standard three-tier systems management approach that includes the integrated Dell Remote Access Controller (iDRAC), OpenManage Enterprise (OME), and CloudIQ. This approach provides a unified, simple, automated, and secure systems management solution. It allows you to manage a single server using the iDRAC Baseboard Management Controller (BMC) console, manage thousands of servers simultaneously with OME, and use intelligent infrastructure insights and predictive analytics to maximize server productivity with CloudIQ.
Cooling
We discussed the importance of cooling on overall server performance and efficiency, especially regarding AI. It may not be sexy, but delivering optimal thermal performance is essential when designing resilient and rugged Edge servers. PowerEdge XR servers were designed with balanced and cooling-efficient airflow and, when necessary, comprehensive thermal management that optimizes airflow, regulates fan speeds, and reduces power consumption.
This level of detail allows PowerEdge XR servers to operate between -5 °C and 55 °C. Dell is working on solutions that could extend this operating range even further.
All PowerEdge XR servers are designed with multiple dual counter-rotating fans. This is equivalent to placing two fans in the same housing. It supports N+1 fan redundancy, allowing the server to continue operating with a single fan failure at a maximum temperature of 40 °C for at least four hours.
Scalability
The server's ability to host up to eight NVMe drives and support dense acceleration capabilities positions it well for future advancements in edge data applications and AI inference. Its rugged design and comprehensive cooling system also mean it will continue to operate efficiently as workloads and environmental conditions evolve.
Targeted Workloads
Targeted workloads include digital manufacturing workloads for machine aggregation, VDI, AI inference, OT/IT translation, industrial automation, ROBO, and military applications where rugged design is required. The XR7620 benefits the retail market and is designed for applications such as warehouse operations, point-of-sale aggregation, inventory management, robotics, and AI inference.
Servers like these are deployed by the military, energy exploration and extraction companies, telecommunications, and others needing edge nodes in the most demanding environments. The PowerEdge XR server line is a short-depth 2U, 2-socket server with data center-level compute that delivers high performance, high capacity, and reduced latency.
Design and Build
The Dell PowerEdge XR7620 stands out for its rugged and robust design. It is designed to withstand harsh environments and is often used in environments where reliability is crucial. Typical of all PowerEdge servers, the XR7620 chassis is well-built, with high-quality materials to meet various challenges.
In terms of design, it is a compact and efficient server, making the most of available space while maintaining excellent airflow for cooling.
A key feature of the XR7620 is the smart bezel. It plays an essential role in enhancing the functionality and security of the Dell PowerEdge XR7620 server. The smart bezel is not just an aesthetic addition but an essential component that contributes to the overall efficiency and reliability of the server.
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
The Dell PowerEdge XR7620 is a well-designed and rugged server known for its durability and reliability. Its thoughtful design and build quality make it a reliable choice for businesses and organizations that need high-performance computing in harsh environments. Its rugged and resilient design and cutting-edge technological advancements position it as a formidable player in harsh environments. Additionally, the server's versatile connectivity and expansion options, combined with Dell's reliable support and warranty services, ensure that it not only meets but exceeds the requirements of modern edge computing scenarios.
While price is a factor to consider, the Dell PowerEdge XR7620 offers undeniable value for money, especially for industries requiring flawless reliability and performance in extreme conditions. Overall, the PowerEdge XR7620 is a comprehensive solution designed to address the complexities and challenges of modern edge computing, proving to be a wise investment for forward-looking businesses.
Next, we push this system to the extreme for a major research project, with more to come!
