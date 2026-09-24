---
id: collect-240926-storagereview/storagereview/fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba
title: "fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba"
domain: storagereview
role: reference
task: reference
actors: ["DeepSeek", "Hugging Face", "Nvidia", "SpaceX"]
dates: []
keywords: ["deepseek", "inference", "nvidia", "cost", "distribution", "energy", "ethernet", "gpu", "latency", "memory", "nand", "parameters"]
source: docs/RAG/clean_en/storagereview/fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba.md
source_anchor: ""
source_lines: [1, 66]
sha256: 5ef26d6aadc7245a1af4d17ea06d3d257abcf914c72d23ca24c67d6c56b8c9be
---

# fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba

<!-- source: https://www.storagereview.com/fr/review/nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-the-edge -->

The Jetson Orin Nano Super is a compact computing powerhouse that brings sophisticated AI capabilities to edge devices. It combines performance, affordable price, and solid integration options, making it an ideal candidate for prototyping and commercial product development. Whether used in robotic kits or integrated into larger machines, its flexible design allows engineers to deploy AI in scenarios that demand efficiency and low power consumption - for only $249.
The Jetson platform is specially designed for edge deployments, ensuring that projects in environments with limited space or power can still leverage high-end AI performance. With a scalable format and extensive connectivity options, it offers a gateway to innovative solutions in robotics, intelligent surveillance, and even wildlife conservation.
The Jetson Orin Nano Super board is renowned for delivering projects requiring embedded artificial intelligence, whether in traditional robotic kits using classic programming or in more advanced configurations integrating frameworks such as ROS (Robot Operating System). Available as a complete development kit and as a standalone SoC daughterboard, it easily integrates into a vast range of products and machines. This versatility makes it popular for applications ranging from small educational projects to large-scale industrial deployments.
Jetson Orin Nano Super Development Kit Specifications
The Jetson Orin Nano Super packs impressive features into a compact format. The 78-core Arm Cortex-A6AE processor provides a solid foundation for computing, while the 1024-core NVIDIA Ampere GPU with Tensor Cores accelerates various workloads, including deep learning and computer vision tasks. With 67 TOPS (Tera Operations Per Second) of AI performance and 5 GB of high-bandwidth LPDDR8 memory, this platform is designed to perform complex operations at the edge.
| Specifications | DETAILS | 
|---|---|
| Processor | Arm Cortex-A6AE v78 8.2-bit 64-core processor, 3 MB L2 + 4 MB L3 | 
| GPU | 1024-core NVIDIA Ampere architecture GPU with 32 Tensor Cores | 
| AI Performance | 67 TOPS | 
| Memory | 8 GB 128-bit LPDDR5 102 GB/s | 
| Storage | NVMe SSD support 16 GB eMMC 5.1, microSD, M.2 Key M 1x M.2 Key M slot with x4 PCIe Gen3 1x M.2 Key M slot with x2 PCIe Gen3 | 
| Networking | Gigabit Ethernet 1x | 
| Display | 1x HDMI, 1x eDP 1.4 | 
| Connectivity | 4 USB 3.2 Type-A ports, 1 USB Type-C port | 
| Power | DC barrel jack accepts 7 V to 20 V power supply | 
| Camera | 2x MIPI CSI camera connectors | 
| Expansion | 40-pin GPIO expansion headers | 
| Power Consumption | 7W – 25W configurable | 
| Operating System | Ubuntu-based Linux with NVIDIA JetPack SDK | 
| Dimensions | 103mm x x 90.5mm 34.77mm | 
Connectivity options are numerous, making the Nano Super extremely versatile for many applications. Four USB 3.2 Type-A ports and one USB Type-C port allow you to easily connect a range of peripherals, from external storage devices to input devices or sensors. The built-in Gigabit Ethernet ensures reliable networking, while the two MIPI CSI camera connectors allow for the integration of two cameras. This feature is particularly advantageous for applications requiring depth perception, essential in robotics and autonomous systems where accurate environmental mapping is critical.
Storage capabilities include 16 GB of eMMC 5.1 memory, a microSD card, and dual M.2 NVMe SSDs via dedicated slots with PCIe Gen3 connectivity. This provides sufficient storage space for operating systems, software, and datasets and supports the high-speed data transfers needed for real-time analytics and AI inference tasks. Additionally, the inclusion of HDMI and eDP 1.4 interfaces allows the Nano Super to support displays, making it ideal for kiosk-type applications or digital signage.
Pushing the Nano Super to its limits: LLM Inference at the Edge
Our work with the Nano Super focused on exploring its potential for performing AI development tasks, particularly large language model (LLM) inference. We recognized that onboard memory limitations complicated running models with billions of parameters, so we implemented an innovative approach to work around these constraints. Typically, the Nano Super's 8 GB of graphics memory limits it to smaller models, but we sought to run a model 45 times larger than what would traditionally fit.
We enhanced the Nano Super's storage by integrating the recently launched Solidigm D5-P5336 122.88 TB SSD, a very high-capacity NVMe drive designed for data center environments, to support this ambitious task.
The Solidigm D5-P5336 122 TB SSD is a revolutionary storage solution for data-intensive workloads, particularly in AI and data center domains. Here are its detailed specifications:
- Capacities: 122.88TB
- Technology: Quad-Level Cell (QLC) NAND
- Interface: PCIe x4 Gen4
- Performance: Up to 15% improvement on data-intensive workloads compared to previous models
- Form Factor: U.2 Approximately the size of a deck of cards
- Use Cases: Ideal for AI training, data collection, media capture, and transcoding
Performance Indicators
- Sequential Read/Write Speed: Up to 7.1 GB/s (read) and 3.3 GB/s (write)
- Random Performance: Up to 1,269,000 XNUMX XNUMX IOPS
Endurance Measures
- Endurance: The Solidigm 122 TB SSD is designed for data-intensive workloads and offers high endurance. You can use the Solidigm SSD Endurance Estimator to calculate the expected lifespan based on specific workloads.
Power Measures
- TB per watt=122 TB25 W=4.88 TB/WTB per watt=25 W122 TB=4.88 TB/W. With these power measures, this drive offers approximately 4.88 terabytes of storage per watt of power consumed, highlighting its efficiency for data-intensive applications.
The Nano Super includes two M.2 NVMe bays, which we tested as part of this evaluation. Both slots offer a PCIe Gen3 connection, with a 30 mm slot supporting 2 PCIe lanes and an 80 mm slot supporting 4 full PCIe lanes. We used the 80 mm slot paired with a breakout cable to provide the highest bandwidth to the Solidigm D5-P5336 122 TB QLC SSD. Our USB-C power cable wasn't ready for the demonstration, so we used an ATX power supply that provided 12 V and 3.3 V to the U.2 drive.
The result was an overpowered storage solution that allowed us to handle massive models and highlighted the role of robust storage in cutting-edge AI workflows. This configuration allowed us to store and transport most popular models from Hugging Face while retaining sufficient additional space.
How were we able to run DeepSeek R1 70B Distilled, a model 45 times larger than expected, on such a device? To do this, we used AirLLM, a project that sequentially loads model layers into memory as needed, rather than loading all weights at once. This layer-by-layer approach allowed us to perform inference on a model far exceeding the device's VRAM limits. One caveat, however: computing performance. In terms of storage performance, via the 4-lane PCIe 3 connection, the NVIDIA Orin Nano board could achieve approximately 2.5 GB/s on the Solidigm D5-P5336 122 TB QLC SSD. With our inference workload running on the QLC SSD, read speeds hovered around 1.7 GB/s.
Even though we managed to work around the VRAM limitations, we were still stuck with 67 TOPS of performance. Additionally, as the model size increases, the layer size also increases, meaning the time per token increases. So we went from a few tokens per second with smaller LLMs, such as ChatGLM3-6B, to one every 4.5 minutes with DeepSeek R1 70B Distilled.
Practical Applications of Large-Scale Storage and Cutting-Edge AI
While our LLM experiment was more of a proof of concept, the combination of the Jetson Orin Nano Super and a high-capacity Solidigm drive has practical applications. The Jetson's SODIMM format makes it easy to integrate into custom printed circuit boards, making the connection of professional-grade U.2 drives simpler and more plausible. This configuration is advantageous for long-term, low-power AI deployments in remote or sensitive environments.
Artificial intelligence (AI) is increasingly used in wildlife conservation. In a previous article, we explained how AI contributes to monitoring hedgehog populations. Similarly, Indigenous nations in British Columbia are using AI to monitor fish populations. These installations often need to operate uninterrupted for years, requiring large storage capacities, low power consumption, and minimal environmental impact. A solution based on the Jetson Orin Nano Super, equipped with a high-capacity hard drive, can meet these requirements while consuming only 15 W (or 50 W at full power). With backup batteries and a small solar panel, such an installation can be the size of a standard desk phone, making it discreet and practical for long-term use.
Another interesting use case is using the system as a large local repository for model distribution. By downloading hundreds of models from Hugging Face, we noticed that not all models were identical. The most popular models downloaded faster than older or less popular models. However, all downloads are generally very slow at the edge, even with Starlink. In such cases, a package like the Nano Super, equipped with an additional network card and a high-capacity drive, would serve perfectly as a cache or intermediate store to efficiently redistribute models at the edge.
Many Use Cases
Here are some compelling use cases for leveraging an NVIDIA Jetson device with substantial storage capacity:
- Autonomous Vehicles: Storing and processing large amounts of sensor and camera data in real time for navigation and obstacle detection.
- Intelligent Surveillance: Managing high-resolution video streams from multiple cameras for security and monitoring purposes, with the ability to store and analyze footage locally.
- Health Diagnostics: Real-time processing and storage of medical imaging data for immediate diagnostics and treatment decisions in remote or resource-limited environments.
- Industrial Automation: Enhancing factory automation with AI-driven quality control and predictive maintenance, storing large datasets for analysis and model training.
- Retail Analytics: Analyzing customer behavior and inventory data in real time to optimize stock levels and improve the shopping experience.
- Environmental Monitoring: Using AI to track and analyze ecological data, such as air and water quality, to support conservation efforts and public health initiatives.
- Smart Agriculture: Monitoring crop health and soil conditions using AI-powered sensors and cameras to optimize farming practices and increase yield.
- Telecommunications: Managing and processing data in cell towers to improve network performance and reduce latency.
Conclusion: Finding its place in the Jetson family
The Jetson Orin Nano Super sits at the heart of NVIDIA's Jetson lineup, offering a balance between high performance and energy efficiency for edge AI tasks. The Jetson family ranges from entry-level models like the Jetson Nano, designed for basic AI and robotics applications, to the powerful Jetson AGX Orin, which provides up to 275 TOPS for demanding autonomous machine workloads. In between, the Jetson Orin Nano Super offers flexible performance and power profiles, meeting the needs of developers who need more power without the bulk of a full AGX platform.
Solidigm's QLC SSD lineup offers a range of high-capacity storage solutions designed for read-intensive workloads. The lineup includes models like the D5-P5336, with up to 122.88 TB of storage and smaller drive capacities starting at 7.68 TB. These SSDs are optimized for performance, density, and cost-effectiveness, making them ideal for applications such as content delivery networks, AI, data pipelines, and object storage. Thanks to QLC technology, Solidigm SSDs offer substantial storage capacity while maintaining solid read performance and proven reliability.
The Nano Super's ability to integrate serious AI capabilities into compact, low-power environments sets it apart from others. While the original Jetson Nano was a favorite for hobbyists and light AI tasks, the Nano Super raises the bar by offering 67 TOPS, enough to handle complex LLM inference and other demanding AI applications. This makes it an attractive option for developers looking to deploy sophisticated AI models at the edge without the overhead of larger, more power-hungry systems. Paired with a high-capacity QLC offering, such as the Solidigm D122-P5 5336 TB SSD, it allows edge locations to operate with a wide range of AI models and without capacity constraints requiring storage swapping once provisioned.
The Nano Super costs $249. While it's more expensive than a Raspberry Pi, it offers significantly higher performance and includes all necessary components. The heatsink, equipped with a fan, allows operation at maximum power even in a poorly ventilated 3D-printed case. It also comes with a power adapter, making it ideal for those interested in AI.
StorageReview thanks the Solidigm team for the new D122-P5 5336 TB SSD. The capacity and speed of this drive allowed us to carry out a large part of the testing.
