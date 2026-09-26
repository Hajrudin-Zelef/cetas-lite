---
id: collect-240926-storagereview/storagereview/fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba-3
title: "fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba"
domain: storagereview
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "SpaceX"]
dates: []
keywords: ["inference", "nvidia", "cost", "distribution", "energy", "latency", "robotics", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba.md
source_anchor: ""
source_lines: [48, 66]
sha256: 663540c68683299db818674c9cd069bc61d5684070dd1b194088122323abd60d
---

# fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba

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
