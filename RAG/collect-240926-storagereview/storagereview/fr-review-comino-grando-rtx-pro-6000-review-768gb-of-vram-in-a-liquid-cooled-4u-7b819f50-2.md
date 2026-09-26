---
id: collect-240926-storagereview/storagereview/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u-7b819f50-2
title: "fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Apple", "Google", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "blackwell", "distribution", "gpu", "gpus", "inference", "intel", "liquid cooling", "memory", "nvidia"]
source: docs/RAG/clean_en/storagereview/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50.md
source_anchor: ""
source_lines: [3, 51]
sha256: 50a3b4a8338331c50dd021da70910673768751afb201c24c470a8429149f7cdc
---

# fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50

Comino recently sent us the latest version of the Comino Grando for testing. This configuration includes eight NVIDIA RTX PRO 6000 Blackwell cards, each with 96 GB of VRAM, for a total of 768 GB of GPU memory. We had tested the Comino in 2024, with a configuration of six RTX 4090s offering 144 GB of GPU memory, as well as a version equipped with NVIDIA H100. This new version represents a significant generational leap, both in terms of raw memory capacity and the range of supported workloads.
The Grando is a 4U platform designed specifically to resolve the crucial conflict between high-density GPU computing power and thermal management. While classic air-cooled chassis collapse under the continuous consumption of more than 600 W from modern professional graphics cards, the Grando adopts a fundamentally different approach. Entirely designed around a liquid-cooling architecture, it is capable of dissipating a massive 6.5 kW of heat continuously. This is not an adaptation or a last-minute addition: the entire chassis, from the inverted motherboard layout to its color-coded quick-connect system, was designed around the cooling loop.
The result is a platform capable of accommodating eight professional GPUs at full consumption (TDP) in a single 4U chassis, operating 24/7 in ambient environments of 3 to 38 °C, without thermal throttling, without the noise nuisance of high-RPM air cooling, and without compromising maintenance. For businesses deploying AI inference, machine learning, or high-performance simulation workloads at scale, the Grando offers a truly rare solution: a server that does not force you to choose between density, heat dissipation, and reliability.
Comino Grando Specifications
The table below presents the physical specifications and supported hardware configurations for the Comino Grando platform.
| Specifications / Features | Comino Grando | 
|---|---|
| Comino Grando Server and Rackable Workstation |  | 
| Cooling capacity | 6.5 kW (Maximum 6,500 W at 20 °C intake air temperature) | 
| Motherboards | Up to EATX and EBB | 
| GPU (Server) | Up to 8; NVIDIA: RTX A6000, RTX 6000 ADA, RTX PRO 6000, A40, L40, L40S, A100, H100, H200 | 
| GPU (Rackable Workstation) | Up to 6; NVIDIA: 3090, 4090, 5080, 5090, RTX A6000, RTX 6000 ADA, RTX PRO 6000, A40, L40, L40S, A100, H100, H200; AMD: W7800, W7900 | 
| CPU | Up to 2; Single-socket processor: Intel Xeon W-2400/2500 and 3400/3500, Intel Xeon Scalable 4th Gen, 5th Gen, Xeon 6, AMD Threadripper PRO 5000WX, 7000WX, 9000WX, AMD EPYC 9004/9005 Dual-socket: Intel Xeon Scalable 4th and 5th Gen, Xeon 6, AMD EPYC 9004/9005 | 
| RAM | Up to 2TB | 
| M2 Drives | Up to 8 NVMe slots | 
| Storage | Hot-swap cages on the rear panel: up to 4 hot-swap SSDs (4 x 7 mm or 2 x 15 mm) and up to 4 more (4 x 7 mm or 2 x 15 mm) in place of a 4th power supply; Internal 3.5″ cage up to 4 x 3.5″ or 4 x 2.5″ 15 mm or 12 x 2.5″ 7 mm; Internal 2.5″ slots: up to 4 x 2.5″ 7 mm SSDs | 
| Power supply and operating voltage | Up to 4 hot-swap CRPS power supplies of 1,000 W at 180-264 V Up to 4 hot-swap CRPS power supplies of 1,000 W at 90-140 V Redundancy modes: 4+0, 3+1, 2+2 | 
| Noise level | 39dB-70dB | 
| LAN | Up to 2 x 10 Gbit/s on the motherboard and up to 400 Gbit/s via PCIe | 
| OS | Ubuntu / Windows 11 (Pro/Home) / Windows Server | 
| Physical and cooling specifications |  | 
| Liquid cooling | CPU with VRM and GPU with GDDR and VRM | 
| Reservoir | Custom Comino 450ml with integrated pumps | 
| Fans | 3x Ultra High Flow 6200 RPM (high noise level) or 3x High Flow 3000 RPM (low noise level) | 
| Installation | 19-inch rack mount or standalone use as a workstation | 
| Rack space required | 4U | 
| Size | 439 x 681 x 177 mm (without handles or protruding parts) | 
| Weight | 4 GPU: 49 kg (net), 67 kg (gross) 6 GPU: 52 kg (net), 70 kg (gross) 8 GPU: 55 kg (net), 72 kg (gross) | 
| Operating and storage temperature range | Storage: -5 to 50 °C / 23 to 122 °F Operating temperature: 3 to 38 °C / 38 to 100 °F | 
| Comino Monitoring System (CMS) |  | 
| Market | Control board with sensors and software for real-time monitoring | 
| Key advantages | Cooling system and CPU/GPU monitoring, web interface, cooling system log, centralized monitoring for workgroups | 
| Connected sensors and devices | Temperature (air and coolant), relative humidity, voltage, coolant flow rate, coolant level in the reservoir, fans, pumps, motherboard, display, and buttons | 
| Integration possibilities | Set up monitoring via a REST API and send sensor data to monitoring software (e.g., Zabbix, Grafana) or databases (e.g., InfluxDB). | 
| CMS technical requirements |  | 
| OS | Windows 11 / 10 Ubuntu 22.04/20.4 (Dependency for Ubuntu: the target system must have nvidia-smi and sensors utilities installed) | 
| Web Browsers | Mozilla Firefox, Google Chrome, Chromium, Apple Safari, Microsoft Edge (Note: Internet Explorer 11 is not supported) | 
| Hard drive | 300MB | 
| Controller firmware version | 1.0.6 or newer | 
| Controller PCB version | 2.xx.xx | 
GPU Design, Construction, and Density
Chassis Layout and Deployment
The Grando server is a model of space optimization, with dimensions of 17.3 x 26.8 x 6.97 cm (4U). Unlike traditional servers, it places the rear of the motherboard at the front of the chassis, reversing the classic internal layout. This ensures that air-cooled components, such as RAM modules and VRMs, benefit from optimal fresh airflow before it reaches the liquid cooling radiator located at the rear.
The chassis itself is manufactured to the same rigorous standards, with solid steel construction and a matte black epoxy paint finish applied both inside and out. This deliberate choice extends to the tubes, cables, radiator, and PCB solder mask, demonstrating a clear commitment to a polished and professional aesthetic. Additionally, the system offers great installation flexibility and can be used both in a 19-inch rack and as a standalone desktop unit. Depending on the configuration, its weight varies between 148 and 159 kg.
GPU Cold Plates and Cooling Blocks
The proprietary copper waterblocks form the heart of the Grando's density, cooling not only the graphics chip but also other components such as memory and voltage regulators. Each GPU is delivered as a standard card, onto which Comino installs a custom cooling system. Specifically, this ultra-thin design reduces each card's footprint to a single slot, allowing six or even eight professional GPUs to be installed side by side in a 4U chassis. Our test model was equipped with eight NVIDIA RTX PRO 6000 Blackwell cards, each with a TDP of 600 W, representing a total cooling requirement of 4,800 W at full load.
Achieving the density of 8 GPUs per Comino slot would be nearly impossible with air cooling, as standard NVIDIA RTX PRO 6000 cards each occupy two slots and require significant airflow. Conversely, these custom-cooled cards occupy only a single slot each. The cooling plates are robust, which adds noticeable weight to each card, but this weight reflects the quality and cooling performance required at this level.
Each pair of GPUs is connected to a dedicated secondary manifold that combines the two cards into a single inlet and outlet connection to the main cooling manifold. This paired design simplifies the overall loop architecture, reduces the number of connections at the main manifold, and allows a technician to disconnect a single pair of quick-connects to remove two cards simultaneously, further simplifying maintenance.
Water Distribution and Manifold
