---
id: collect-240926-storagereview/storagereview/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi-d918d5c2
title: "fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["apache", "benchmark", "benchmarks", "blackwell", "compute", "distribution", "fine-tuning", "fp4", "governance", "gpu", "gpus", "inference"]
source: docs/RAG/clean_en/storagereview/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2.md
source_anchor: ""
source_lines: [1, 101]
sha256: 8617c4f6784f1e25ce6bacb6b9bf43324167f256b1be396fb0af136b6931f890
---

# fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2

<!-- source: https://www.storagereview.com/fr/review/hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi-gpu-ai -->

The HPE ProLiant Compute DL380a Gen12 server is aimed at enterprise AI teams looking for high compute density without changing their rack configuration. This 4U air-cooled chassis integrates easily, supports up to eight double-width GPUs, and offers full PCIe Gen5 connectivity. It can be configured with two Intel Xeon 6 processors, each with 144 cores, 4 TB of DDR5 memory across 32 DIMM modules, and sixteen E3.S NVMe bays for high throughput and capacity. The goal is simple: achieve production-grade inference capacity and precise fine-tuning at scale without resorting to liquid cooling.
For accelerators, HPE offers a full range including NVIDIA H200 NVL, H100 NVL, L40S, L20, L4 cards and the RTX PRO 6000 Blackwell Server Edition, with power options compatible with high-consumption components. In this test, we focus on the RTX PRO 6000 Server, which offers an excellent compromise for enterprise AI. Each card carries 96 GB of ECC GDDR7 memory, a PCIe Gen5 x16 interface, FP4-compatible Tensor cores, and a 600 W thermal envelope, suited to air-cooled racks. Our configuration was equipped with four cards, a sensible starting point for high-throughput inference and targeted optimization, with room for growth.
HPE completes the platform with essential operational elements. iLO 7 handles out-of-band configuration, system status, and power management, thanks to the Silicon Root of Trust (a secure enclave ensuring firmware integrity), RSA 4096-bit encryption support, and a detachable iLO DC-MHS module that strengthens supply chain verification. The server also integrates with the HPE Private Cloud AI platform for multi-team governance and reproducible deployments at scale.
HPE ProLiant Compute DL380a Gen12 – Technical Specifications
| Categories | Specifications | 
|---|---|
| Processor type | HPE ProLiant Compute DL380a Gen12 | 
| Processor family | 6th Generation Intel® Xeon® Scalable Processors | 
| Available processor cores | 64 to 144 cores, depending on processor | 
| Number of processors | 2 | 
| Processor speed | Up to 2.4 GHz, depending on processor | 
| Maximum memory | RDIMM 4 TB (2 TB per processor) | 
| Memory slots | 32 DIMM slots | 
| Memory type | HPE DDR5 Smart Memory | 
| Memory protection | RAS: Advanced ECC, online spare memory, mirroring, combined channel functionality (locking), HPE Fast Fault-Tolerant Memory (ADDDC) | 
| Drive support | SFF NVMe and EDSFF | 
| Security | Optional locking bezel, intrusion detection, and integrated HPE TPM 2.0 module | 
| Infrastructure management | HPE iLO Standard with intelligent provisioning (embedded), HPE OneView Standard (requires download) • Optional: HPE iLO Advanced and HPE OneView Advanced (licenses required) | 
| Power supply | Up to 8 M-CRPS. Single 1+1 redundancy for the motherboard. Dual 2+1 redundancy for GPUs. | 
| Expansion connectors | 6 | 
| System fans | 4 dual-rotor fans and 8 hot-swappable single-rotor fans included | 
| Form factor | 4U rack | 
| Warranty | 3/3/3: Server warranty | 
Design and assembly of the HPE ProLiant DL380a Gen12 server
The HPE ProLiant Compute DL380a Gen12 rack server is a 4U dual-socket server designed for scalable, high-performance deployments. Measuring 6.88 x 17.63 x 31.60 cm, it combines dense CPU and GPU compute power with efficient air cooling for reliable operation even under heavy workloads.
Weighing between 82.7 and 137.8 kg depending on configuration, the chassis supports high-capacity components, redundant power, and offers easy front access for maintenance. Its design prioritizes performance, scalability, and efficient thermal management, making it ideal for enterprise and data center environments.
On the storage side, the HPE ProLiant DL380a Gen12 offers 4- or 8-bay configurations in SFF or EDSFF formats. Our test model was equipped with the HPE DL380a Gen12 NS204i-u front bay kit, supporting two hot-swappable NVMe M.2 boot devices. The chassis also included eight 2.5-inch bays, occupied by two HPE U.3 SSDs with a total capacity of 15.36 TB. HPE offers several front bay options, providing great flexibility to adapt to different deployment needs.
The unit can be transported using its two side handles, which means at least two people are required for safe racking and installation. It uses a 2U rail kit with telescopic rails, allowing for easy installation and simplified maintenance without fully dismantling the rack.
At the rear of the HPE ProLiant DL380a Gen12 server, the optimized layout promotes airflow, expandability, and ease of maintenance. The system supports up to eight MCRPS power supplies (1 to 8) thanks to an integrated ventilation panel, ensuring optimal cooling even at full load. Expandability is significant, with multiple PCIe Gen5 x16 slots (slots 1 to 6) compatible with integrated and optional expansion cards, as well as OCP A and B slots for flexible network adapter configuration.
Connectivity includes a dedicated iLO network port, multiple USB 3.2 Gen 1 ports, and a VGA port for local management. It is important to note that slot 1 is available only when the HPE DL380a Gen12 4EDSFF Direct Cable for NVD (P74716-B21) is installed and cannot be used with SFF NVMe drives. As for slot 4, it is not supported in configurations with 4 or 8 DW GPUs.
Power for the HPE ProLiant DL380a Gen12 server is provided by hot-swappable M-CRPS Titanium modular power supply kits. Compatible models include the 1,500 W (P67244-B21), 2,400 W (P67252-B21), and 3,200 W (P67248-B21) versions. The system supports up to eight power supplies, offering N+1 redundancy to ensure continuous operation even if a power module fails. Power requirements and distribution may vary depending on GPU configuration. Our test model was equipped with five 2,400 W M-CRPS power supplies, providing sufficient capacity to power the system's four GPUs (600 W TDP) while ensuring reliable redundancy.
Observing the interior of the HPE ProLiant DL380a Gen12 from the top, it is clear that HPE designed this chassis with a focus on cooling the GPUs, positioned at the front of the system to benefit from direct, unobstructed airflow. The cooling system includes four hot-swappable fan groups, each consisting of a 92 x 56 mm dual-rotor fan and two 40 x 28 mm single-rotor fans. The smaller fans concentrate airflow on the processor and memory modules, ensuring efficient thermal management of lower components. By comparison, the larger dual-rotor fans are specifically designed to generate significant airflow directly onto the GPU array. This balanced design ensures optimal cooling of compute and acceleration components, even under high and sustained workloads.
When examining the GPU configuration, our unit was pre-wired for four PCIe 5.0 GPUs, each cleanly installed in the front GPU cage. The system was configured with NVIDIA RTX PRO 6000 GPUs (Blackwell Server Edition, 96 GB), belonging to NVIDIA's new professional range optimized for AI, rendering, and compute workloads. Depending on configuration, the DL380a Gen12 can support 4 or 8 double-width GPUs or up to 16 single-width accelerators, offering great flexibility for a wide range of enterprise and AI deployments.
The GPUs compatible with this platform are as follows:
- NVIDIA RTX PRO 6000 Server Edition (96 GB)
- NVIDIA H200 NVL (141 GB)
- NVIDIA H100 NVL (94 GB)
- NVIDIA L40S (48 GB)
- NVIDIA L20 (48 GB)
- NVIDIA L4 (24 GB)
This flexible GPU configuration, combined with high-bandwidth PCIe Gen5 lanes, ensures that the DL380a Gen12 is ready for dense inference tasks and large-scale AI training environments. Inside the chassis, once the cooling shroud is installed, HPE's ingenuity in airflow management becomes apparent. This shroud features precisely molded deflectors that efficiently direct air toward the processors, memory modules, and VRMs, ensuring uniform system cooling.
Regarding the CPU cooling system, HPE designed the DL380a Gen12 with a focus on thermal balance. Each Xeon 6 processor is equipped with a tall, high-density heatsink designed to handle heat generated by the processor and that generated by the GPUs located at the front of the chassis. This design ensures consistent cooling performance, even under heavy mixed workloads, where heat from the GPUs can raise the ambient temperature around the processor. When high-consumption GPUs are installed, the combination of taller heatsinks and HPE's front-to-back airflow design provides the surface area and cooling efficiency needed to effectively manage the additional thermal load.
iLO 7 Overview
As noted, the system includes a dedicated iLO port offering out-of-band management capabilities for comprehensive server control and monitoring. This enclosure ships with the new HPE iLO 7 interface, providing administrators with a modernized interface and enhanced features, integrated with HPE Compute Ops Management for simplified configuration, monitoring, and lifecycle management. Below is the new HPE iLO 7 login page of our demonstration system.
From the dashboard, it is clear that HPE iLO 7 offers a modernized interface that immediately highlights system status and key health indicators. The main panel provides an overview of power status, health status, and host connection with HPE Compute Ops Management. On the right, general system information such as iLO IP address, hostname, and license type are displayed for quick reference.
The dashboard groups key indicators (fan redundancy, power status, and temperatures) into a clear, colorful interface, allowing quick assessment of server status. Access to virtual media and the remote console is also direct from the main page, simplifying common remote administration tasks without additional navigation.
In the Firmware tab, HPE iLO 7 offers a clear and organized view of all components and update management tasks. The interface presents firmware inventory, active installation queues, and verification results as cards, making navigation easy. Administrators can quickly launch updates, upload packages to the iLO repository, or create installation sets for batch deployment.
Firmware verification and repository management are integrated into this interface, allowing users to ensure integrity and version consistency across components. The Quick Actions menu, located on the right, simplifies essential tasks such as updating firmware or uploading new files. The Firmware Settings section allows control over downgrade policies and acceptance of third-party packages.
In the Host section, HPE iLO 7 provides quick access to essential server management functions, including power control, virtual media, hardware status, and system performance. Administrators can view hardware redundancy status in real time, access the integrated management log, or launch the remote console directly from this interface. The interface also offers quick actions such as graceful shutdown, restart, and reset, allowing full remote system control without physical intervention.
The right panel displays host settings, including TPM status, platform policy configuration, and hardware module information. This section highlights iLO 7's role as a centralized control platform, allowing administrators to manage power, monitor events, and oversee operations securely from a single interface.
In the Host > Hardware view, we can see the GPUs installed in the system. As indicated, the unit is equipped with four NVIDIA RTX PRO 6000 Blackwell GPUs, all enabled and operating normally. iLO 7 provides detailed hardware information, including model, part, and serial numbers, allowing administrators to verify component status at a glance.
In the Security tab, HPE iLO 7 centralizes all essential controls related to system protection and access management. The overview panel provides a security status overview, highlighting risk levels, configuration locks, and certificate status. Administrators can thus manage encryption settings, authentication methods, and TLS certificates, as well as configure secure erase and remote key management.
The interface clearly flags critical points, such as self-signed certificates or unconfigured key management, while ensuring secure operation where applicable. Security logs, user management policies, and access controls are easily accessible, giving administrators a comprehensive overview of system security directly from the iLO environment.
Under the HPE Applications tab, iLO 7 provides access to integrated tools that optimize server deployment and lifecycle management. From this interface, administrators can launch Intelligent Provisioning, an integrated utility designed to simplify operating system installation, firmware updates, and system configuration without requiring external support.
The iLO Settings tab groups all configuration and administration options for the iLO interface. Administrators can control user access, network port configuration, authentication methods, and activity logs. The menu also offers options for troubleshooting, applying security policies, license management, and time synchronization.
Quick actions, such as backing up or restoring the iLO configuration and resetting, are easily accessible on the right, simplifying maintenance tasks. The interface, like the rest of iLO 7's modern card-based interface, offers clear and organized management of security, connectivity, and operational settings from a centralized location.
Performance Testing
To evaluate the real-world capabilities of the DL380a Gen12, we conducted a comprehensive series of performance tests covering both AI inference and general compute workloads. These tests include vLLM online serving benchmarks for large language models (LLMs) and Phoronix Test Suite benchmarks to measure CPU throughput, memory bandwidth, web service efficiency, and cryptographic performance.
System Configuration
- CPU: 2 Intel Xeon 6527P processors
- Memory: HPE 64 GB 2Rx4 PC5-3400B-R Smart Kit (16 disks)
- GPU: 4 x NVIDIA RTX PRO 6000 (96 GB)
- Storage: 2 x 15.63 TB PM1733a U.3
vLLM Online Serving – LLM Inference Performance
vLLM is the most popular high-throughput inference and serving engine for LLMs. The vLLM online serving benchmark is a performance evaluation tool that measures the actual serving capabilities of this inference engine under concurrent requests. It simulates production workloads by sending requests to a running vLLM server with configurable parameters, such as request rate, input/output length, and number of concurrent clients. The benchmark measures key metrics, including throughput (tok/s), time to first token, and time per output token, allowing users to understand vLLM performance under different load conditions.
We tested inference performance on three representative models covering different scales and quantization approaches, evaluating how the four NVIDIA RTX PRO 6000 GPUs in the HPE ProLiant DL380a Gen12 handle production inference workloads.
Dense Model Performance
Dense models represent the conventional LLM architecture, where all parameters and activations are used during inference. We evaluated two dense model configurations: Llama-2-70b-chat-hf and Llama-3.2-90B-Vision-Instruct.
Llama-2-70B-Chat Performance
In single-user configuration (BS=1) with TP=4, the model achieves 32.89 tok/s per user and a TPOT of 30.18 ms. With BS=8, performance reaches 15.68 tok/s per user, for a total throughput of 433.62 tok/s and a TPOT of 35.98 ms. Moving to BS=32, total throughput reaches 741.62 tok/s while maintaining 8.00 tok/s per user and a TPOT of 43.44 ms.
Llama 3.2-90B Vision Instruction Performance
With a base station (BS) of 1 and a total throughput (TP) of 4, the model achieves a throughput of 20.59 tok/s per user and a TPOT of 38.27 ms. With a BS of 16, performance increases to 7.20 tok/s per user, for a total throughput of 806.14 tok/s and a TPOT of 54.98 ms. The maximum total throughput of 1,372.21 tok/s is reached with a BS of 128, or 2.59 tok/s per user and a TPOT of 122.75 ms.
Microscaling Data Type Performance
Microscaling represents an advanced quantization approach that applies precise scaling factors to small blocks of weights, rather than uniform quantization across large groups of parameters. NVIDIA's NVFP4 format implements this technique through a block floating-point representation, where each microscaling block of 8 to 32 values shares a common exponent serving as a scaling factor. This granular approach preserves numerical precision while ensuring a 4-bit representation, thus maintaining the dynamic range essential to transformer architectures. This format integrates with NVIDIA's Tensor Core architecture on the RTX PRO 6000, enabling efficient mixed-precision computation with on-the-fly decompression during matrix operations.
GPT-OSS-120B Performance
We evaluated OpenAI's GPT-OSS-120B model with NVFP4 quantization. In single-user mode (TP=2), the model achieves 176.09 tok/s per user with a TPOT of 5.46 ms, the lowest latency in our test suite. With BS=4 and TP=4, performance reaches 105.79 tok/s per user, for a total throughput of 1155.94 tok/s and a TPOT of 7.79 ms. With BS=32 and TP=4, throughput increases to 47.54 tok/s per user and 3956.44 tok/s total, with a TPOT of 13.86 ms. The maximum total throughput of 4015.77 tok/s is reached with BS=64, or 25.38 tok/s per user and a TPOT of 14.78 ms.
Phoronix Benchmarks
Phoronix Test Suite is an open-source automated benchmarking platform supporting over 450 test profiles and more than 100 test suites via OpenBenchmarking.org. It handles the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration. We will focus on the following tests: Stream, 7-Zip, Linux kernel compilation, Apache, and OpenSSL.
Stream Memory Bandwidth
In the Stream benchmark, which measures raw memory throughput, the HPE DL380a Gen12 achieved an impressive score of 542 GB/s, demonstrating the platform's ability to maintain high data transfer rates under continuous load. This level of bandwidth makes the system particularly performant for workloads such as data modeling, simulation, and AI inference, where large datasets must be transferred quickly between memory and compute resources.
7-Zip Compression
The 7-Zip compression test measured 305,000 MIP, highlighting the system's excellent multithreaded efficiency for compute-intensive compression and decompression operations. These results make the DL380a Gen12 an ideal choice for environments involving frequent data packaging, archiving, or backup operations, which require consistent and reproducible CPU performance.
Kernel Compilation
When compiling a full Linux kernel (allmodconfig), the DL380a Gen12 completed the operation in 316 seconds. This result demonstrates the system's ability to easily handle complex, parallelized workloads. Increased compilation performance translates directly into shorter build times and improved iteration speed for developers working on large-scale software projects or in CI/CD environments.
Apache Web Server
In terms of web server performance, the DL380a Gen12 sustained 94,348 requests per second in the Apache performance test. This result demonstrates balanced I/O handling and high cache efficiency, providing the throughput and responsiveness needed for enterprise web applications, virtualization interfaces, or internal service hosting.
OpenSSL Verification
Cryptographic performance was equally remarkable, with the DL380a Gen12 verifying 803 billion operations per second under OpenSSL. This demonstrates the system's ability to handle encryption, authentication, and secure communications workloads at scale.
| Phoronix Benchmarks | HPE ProLiant DL380a Gen 12 (2x Intel Xeon 6527P) | 
| Discussions | 542,720.7 MB / s | 
| 7-ZIP | 304,907 MIP/s | 
| Kernel Compilation (allmod) | 316.166 seconds | 
| Apache (requests per second) | 94,347.52 R/s | 
| OpenSSL | 803,597,895,087 Verifications | 
Conclusion
The HPE ProLiant DL380a Gen12 server stands out as one of the most practical and balanced AI servers for the enterprise AI market. Its 4U air-cooled form factor offers exceptional compute density thanks to its two Xeon 6 processors, support for up to 8 double-width GPUs or 16 single-width GPUs, and 16 E3.S NVMe bays, while ensuring reliability and ease of maintenance. HPE's engineering approach to airflow and thermal balance ensures consistent performance even under heavy workloads, demonstrating that advanced AI acceleration can work perfectly in traditional air-cooled environments.
The integration of iLO 7 significantly improves management, a major asset for leading server vendors like HPE. The modernized interface, integration with HPE Compute Ops Management, and detailed hardware telemetry make remote administration intuitive and efficient. Each section (Dashboard, Firmware, Host, Security, Applications, and Settings) demonstrates HPE's commitment to delivering a smoother, more cloud-integrated experience, without sacrificing the on-premises control essential to enterprise teams.
In performance testing, the server achieved excellent results. The four RTX PRO 6000 GPUs delivered impressive throughput on dense and microscaling LLM models, with vLLM serving performance rivaling that of liquid-cooled systems. Phoronix CPU benchmarks also highlight its balance, with memory bandwidth over 540 GB/s, 94,000 requests per second under Apache, and over 800 billion OpenSSL verifications per second, demonstrating its robustness for both AI and general compute.
HPE's design goal is clear: to deliver high-density, production-ready AI performance through air cooling compatible with existing rack and power infrastructure. For data center teams looking for a reliable, secure, and easy-to-manage air-cooled compute solution, the DL380a Gen12 is a performant and innovative solution for the rapidly growing mainstream AI market.
