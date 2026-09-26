---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887-2
title: "fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "accelerator", "compute", "distribution", "gpu", "gpus", "inference", "memory", "nvidia", "packaging"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887.md
source_anchor: ""
source_lines: [3, 86]
sha256: 7d42aae93a062adfca3c630b5c86319a1af55ebf8585d36c42fffb7051e22d71
---

# fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887

We recently received the Supermicro AS-2126HS-TN server, a dual-processor A+ server based on the Turin architecture, for testing. This 2U platform is designed to support extremely resource-intensive workloads and offers flexible PCIe expansion. Based on two AMD EPYC 9005 (Turin) and 9004 (Genoa) series processors, this system targets a wide range of enterprise and data center use cases, including virtualization, software-defined storage, AI inference and machine learning, cloud computing, enterprise server consolidation, and high-performance computing (HPC).
The AS-2126HS-TN board supports multiple PCIe slot configurations, allowing prioritization of accelerator density or I/O flexibility. Depending on the configuration, the platform can accommodate up to four PCIe 5.0 x16 slots, up to eight PCIe 5.0 x8 slots, or mixed configurations to integrate GPUs, high-speed network controllers, and storage controllers. This modular approach allows the system to adapt to compute-intensive, I/O-intensive, or accelerator-focused workloads without requiring chassis modifications.
To evaluate the platform, we subjected the AS-2126HS-TN to our enterprise test environment, focusing on CPU-intensive workloads reflecting real-world deployment scenarios. Our tests emphasize sustained multithreaded performance, memory behavior, and overall platform efficiency, providing insight into the performance of Supermicro's Turin architecture under demanding enterprise and high-performance computing (HPC) conditions.
Supermicro AS-2126HS-TN Specifications
The table below presents the hardware specifications of the Supermicro AS-2126HS-TN, providing an overview of its platform design, compute capabilities, expansion options, and power and cooling characteristics.
| Specifications | Supermicro A+ Server AS -2126HS-TN | 
|---|---|
| System Overview |  | 
| SKU | A+ Server AS -2126HS-TN | 
| Motherboard | Super H14DSH | 
| Form Factor | 2U Rackmount | 
| Chassis Model | CSE-HS201-R000NFP | 
| Processors |  | 
| CPU Support | Dual processor; AMD EPYC™ 9005/9004 series processors | 
| Maximum Core Count | Up to 384C / 768C | 
| Maximum CPU TDP | Supports processors up to 500W TDP* | 
| CPU Cooling Note | Air-cooled processors with TDP above 400W are only supported under specific conditions. | 
| GPU and Acceleration |  | 
| Maximum GPU Count | Up to 3 double-width GPUs | 
| Supported GPUs (PCIe) | NVIDIA: H100 NVL, RTX 6000 Ada Generation, L4 AMD: Instinct™ MI210 | 
| GPU-GPU Interconnect | PCIe | 
| Memory |  | 
| DIMM Slots | 24 DIMM slots | 
| Maximum Memory (1DPC, EPYC 9005) | Up to 6 TB of DDR5 ECC RDIMM memory at 4,800 MT/s | 
| Maximum Memory (1DPC, EPYC 9004) | Up to 6 TB of DDR5 ECC RDIMM memory at 4,800 MT/s | 
| Memory Voltage | 1.1V | 
| Onboard Devices and Networking |  | 
| Chipset | System-on-chip | 
| Network Connectivity | Via AIOM (AOC options available) | 
| BMC / IPMI | IPMI 2.0 with Virtual Media over LAN and KVM-over-LAN support | 
| Input / Output |  | 
| LAN (BMC) | 1 dedicated RJ45 1GbE BMC LAN port | 
| USB | 2 USB 3.0 ports (rear) | 
| Video | 1×VGA | 
| TPM | 1 onboard TPM / port 80 | 
| BIOS |  | 
| BIOS Type | AMI 64 MB SPI Flash EEPROM | 
| BIOS Features | Plug and Play (PnP); UEFI 2.8; USB keyboard support; ACPI 6.5; SMBIOS 3.7 or later | 
| Management |  | 
| Software / Tools | SuperCloud Composer®; Supermicro Server Manager (SSM); Super Diagnostics Offline (SDO); KVM with dedicated LAN; IPMI 2.0; Supermicro Update Manager (SUM); SuperServer Automation Assistant (SAA); IPMIView | 
| Power Configurations | ACPI/APM power management; power-on mode control for AC power loss recovery; power button bypass mechanism | 
| Security |  | 
| Hardware Security | TPM 2.0; Silicon Root of Trust (RoT) – NIST 800-193 compliant | 
| Security Features | Cryptographically signed firmware; Secure Boot; Secure firmware updates; Automatic firmware recovery; Supply chain security (remote attestation); real-time BMC protections; system lockdown | 
| PC Health Monitoring |  | 
| Fan Monitoring | Tachometer monitoring; speed control status; PWM fan connectors | 
| Temperature Monitoring | CPU and chassis environment monitoring; thermal control of fan connectors | 
| Voltage / Sensors | System temperature; Memory temperature; CPU temperature; 3.3V standby; +5V standby; +5V; +3.3V; +12V; CPU overheat protection | 
| Front Panel |  | 
| LEDs | Hard drive activity; LAN activity; power status; system information | 
| Buttons | Power On/Off; UID | 
| Expansion and Interconnect |  | 
| PCIe Slot Configuration | Option A*: 4 PCIe 5.0 x16 FHFL double-width slots; 1 PCIe 5.0 x16 AIOM slot (OCP 3.0 compatible) Option B*: 8 PCIe 5.0 x8 slots (in x16) FHFL; 1 PCIe 5.0 x16 AIOM module (OCP 3.0 compatible) *Requires additional parts; see optional parts list. Refer to system diagrams for more details. | 
| CXL Support | Up to 4 CXL 2.0 x16 devices | 
| Storage |  | 
| Drive Bays (default) | 8 bays total; 8 front hot-swap bays for 2.5-inch NVMe*/SATA* drives | 
| Drive Bays (Option A) | 24 bays total; 24 front hot-swap bays for 2.5-inch NVMe*/SATA* drives | 
| M.2 | 2 × M.2 PCIe 3.0 x4 NVMe (M-key 2280/22110) | 
| Storage Note | NVMe/SATA support may require an additional storage controller and/or cables. | 
| Cooling |  | 
| Fans | Up to 6 counter-rotating 60×60×56 mm fans | 
| Air Shroud | 2 air shrouds | 
| Power Engine |  | 
| Power Supply Options | 2× 1200W / 1300W / 1600W / 2000W / 2600W (varies by configuration) | 
| Included/Listed Power Supply | 2 × 2000W redundant (1+1) Titanium Level (96%) | 
| Power Supply Dimensions (L×H×W) | 73.5 × 40 x 265 mm | 
| Input (varies by power supply) | 1000W: 100–127 Vac / 50–60 Hz 1800W: 200–220 Vac / 50–60 Hz 2700W: 200–240 Vac 1980W: 220–230 Vac / 50–60 Hz 2000W: 220–240 Vac / 50–60 Hz (UL certified only) 2000W: 230–240 Vac / 50–60 Hz 2000W: 230–240 Vdc / 50–60 Hz (CQC only) | 
| +12V Rail (varies by input) | Max 83 A (100–127 Vac) / Max 150 A (200–220 Vac) / Max 225 A (200–240 Vac) Max 165 A (220–230 Vac) / Max 166 A (230–240 Vac) | 
| SB 12V | Max. 3.5 A / Min. 0 A | 
| Output Type | Backplane (gold finger) | 
| Dimensions and Weight |  | 
| Height | 3.5 mm (88.9 in) | 
| Width | 17.2 mm (437 in) | 
| Depth | 31.74 mm (806.2 in) | 
| Packaging (H×W×D) | 9.96 "× 26.46" × 43.31 " | 
| Weight | Gross weight: 34 kg (75 lb); Net weight: 20.5 kg (45 lb) | 
| Available Color | Silver | 
Supermicro AS-2126HS-TN Design and Mounting
The Supermicro A+ AS-2126HS-TN server is a compact dual-socket 2U platform designed for very high core count configurations, high TDP processors, and multiple accelerators. Its internal architecture optimizes front-to-back airflow, ensuring a balance between thermal efficiency, power scalability, and ease of maintenance for compute, memory, storage, and expansion subsystems.
Front: Storage and Air Intake
The front of the chassis is designed for high-density storage and optimal airflow. Depending on the configuration, the system supports up to 8 or 24 bays for NVMe drives, providing direct PCIe connectivity for high-performance storage workloads. These bays are positioned upstream of the main airflow, ensuring even distribution of incoming cool air before it reaches the compute and expansion zones.
This configuration allows storage to operate independently of CPU and accelerator temperatures, maintaining consistent performance under sustained I/O load while minimizing airflow disruptions within the chassis.
Central Chassis Cooling Architecture
Just behind the front storage compartment is a set of six high-performance hot-swap internal fans. These fans create a direct cooling channel, generating sufficient static pressure to feed two AMD EPYC processors with thermal envelopes up to 500W and multiple double-width PCIe graphics cards.
