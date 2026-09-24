---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887
title: "fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "accelerator", "apache", "benchmark", "benchmarks", "compute", "distribution", "energy", "gpu", "gpus", "inference", "latency"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887.md
source_anchor: ""
source_lines: [1, 169]
sha256: ed397f1806200ae03a0ebf0c9e96d0d9a10cb9be9239f7fb43d8e09eeb597399
---

# fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887

<!-- source: https://www.storagereview.com/fr/review/supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin -->

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
Fan speed is dynamically controlled via PWM connectors and monitored by integrated tachometers, while thermal management is handled at the system level. This design allows the platform to adapt its cooling in real time based on CPU, memory, and chassis temperatures, ensuring stable operation across a wide range of workloads.
Compute Core: Processors and Memory
At the heart of the system is the Super H14DSH motherboard, compatible with two AMD EPYC 9005 or 9004 series processors. Supporting up to 384 cores and 768 threads, the AS-2126HS-TN is perfectly suited for highly parallel workloads such as virtualization, high-performance computing (HPC), and AI inference.
The platform supports processors with TDP up to 500W, with air cooling possible above 400W under certain thermal configurations. Each processor is paired with a full set of memory channels feeding 24 DDR5 DIMM slots, supporting up to 6 TB of DDR5 ECC RDIMM memory in a 1DPC configuration. Memory speed scales with the processor generation, reaching up to 6,400 MT/s with EPYC 9005 processors and 4,800 MT/s with EPYC 9004 processors.
This symmetrical memory layout minimizes latency and ensures balanced airflow across all DIMMs, even in fully populated configurations.
Boot Storage and Onboard Connectivity
For the operating system and management storage, the system includes two integrated M.2 PCIe 3.0 NVMe slots, isolated from the main NVMe backplane. This separation allows administrators to dedicate the front bays to application or data storage while preserving thermal and logical independence of the boot media.
Network management is handled via a PCIe 5.0 x16 AIOM (OCP 3.0) slot, enabling flexible network configurations without monopolizing standard PCIe expansion slots. This approach preserves available PCIe space for additional accelerators or storage controllers.
Expansion and GPU Support
The rear of the chassis is dedicated to PCIe expansion and accelerator support. The AS-2126HS-TN motherboard supports up to three double-width PCIe GPUs, making it ideal for compute-intensive environments. Compatible accelerators include NVIDIA PCIe GPUs such as the H100 NVL, RTX 6000 Ada Generation, and L4, as well as AMD Instinct MI210 accelerators.
All expansion slots are PCIe Gen 5 and arranged to ensure optimal ventilation and perfect mechanical stability for full-height, full-length cards. GPU-to-GPU communication is via PCIe, ensuring platform flexibility and vendor compatibility while supporting a wide range of accelerator configurations.
Rear I/O and Management
The rear I/O panel centralizes local and out-of-band management access. It includes a dedicated 1 GbE IPMI LAN port, two USB 3.0 ports, and a VGA output for direct console access. Management is handled via IPMI 2.0, with KVM over LAN and virtual media support integrated into Supermicro's management software suite, which includes SuperCloud Composer, Supermicro Server Manager, and IPMIView.
Power Distribution and Electrical Flexibility
Power is provided by redundant high-efficiency (Titanium) power supplies available in 1,200W, 1,300W, 1,600W, 2,000W, and 2,600W configurations. This wide range of power supplies allows the platform to scale from CPU-dense configurations to fully GPU-equipped systems without over- or under-sizing power capacity.
This system supports 120 V AC, 240 V AC, and 48-60 V DC input voltages, making it suitable for traditional enterprise racks, high-density data centers, and telecommunications environments. Integrated ACPI/APM power management enables controlled recovery after a power outage, configurable power-on behavior, and a power button priority function. With its redundant power design, the platform offers both optimal energy efficiency and high operational reliability under sustained load.
Security, Monitoring, and Reliability
The AS-2126HS-TN controller incorporates a comprehensive security and monitoring framework, including TPM 2.0, a hardware root of trust, Secure Boot, cryptographically signed firmware, and automatic firmware recovery. Runtime BMC protections and system lockdown features help ensure platform integrity throughout its lifecycle.
System health monitoring covers the CPU, memory, power rails, fan speed, and chassis temperature, with real-time telemetry integrated into Supermicro's management tools. This enables proactive monitoring and rapid response to thermal or power anomalies in production environments.
Supermicro Management (BMC)
The AS-2126HS-TN router comes with Supermicro's integrated BMC, a platform we have already covered for several A+ systems, and which remains an essential element of daily management. The dashboard provides a consolidated view of system status and configuration, allowing immediate visibility into firmware status, hardware inventory, and sensor telemetry data. From the home page, administrators can quickly check BMC, BIOS, CPLD, and Redfish versions, confirm network configuration, and validate host identity before proceeding with deeper troubleshooting or maintenance operations.
In the "Component Information" section, the BMC accurately displays the installed AMD EPYC processors, including the two 192-core processors present in the system. Each socket is equipped with an AMD EPYC 9965, featuring 192 active cores and 384 threads per processor, with a configured TDP of 500W. This view allows quick verification of processor configuration and other component characteristics directly from the BMC, ensuring optimal operation.
In the Cooling tab of Component Information, the BMC allows control of four fan modes: Standard Speed, Maximum Speed, Optimal Speed, and Heavy I/O Speed. These presets allow quick adjustment of airflow based on workload and thermal requirements, without modifying the BIOS. Below the mode selection, the interface displays the status of each fan and its real-time rotation speed, providing a clear view of the system's cooling status.
For remote access, the AS-2126HS-TN BMC includes a KVM console accessible via an HTML5 interface or a Java plugin. Console settings allow selection of the basic mouse mode to adapt to different operating systems, and an IKVM reset option is available if the session hangs. This provides full out-of-band console access for installation, troubleshooting, and recovery, without requiring local peripherals.
Supermicro AS-2126HS-TN Performance Testing
For testing, the Supermicro AS-2126HS-TN motherboard was not configured with dense storage architecture or GPU-optimized power and cabling. Therefore, our evaluation focused exclusively on CPU performance, using Blender, y-cruncher, and Phoronix benchmarks to characterize raw compute throughput on the dual-processor AMD EPYC platform.
The AS-2126HS-TN server was equipped with two 192-core AMD EPYC 9965 processors, providing a total of 384 cores and 768 threads. For comparison, results are obtained against those of a Dell PowerEdge R7725 server configured with the same two AMD EPYC 9965 processors and 1.5 TB of system memory. Both systems were tested with the same test suites and parameters to ensure consistency of results.
Supermicro A+ Server AS -2126HS-TN Configuration
- CPU: 2x AMD EPYC 9965 (192 cores)
- RAM: 1.5 TB 24 x 64 GB DDR5 6000 MHz
- SSD: Micron 7.68 TB Data Center NVMe SSD
Blender 4.5
Blender is an open-source 3D modeling application. This benchmark was performed with the Blender Benchmark utility. The score is measured in samples per minute, with higher values indicating better performance.
In the Blender CPU SMT benchmark, the Supermicro A+ AS-2126HS-TN server, equipped with two AMD EPYC 9965 processors, delivers excellent results, illustrating the performance gains of a 192-core architecture for multithreaded rendering workloads. With the Blender Benchmark utility, the system recorded 3,070.84 samples per minute for Monster, 2,063.61 for Junkshop, and 1,527.39 for Classroom.
For comparison, the Dell PowerEdge R7725, equipped with the same two EPYC 9965 processors, shows slightly higher throughput in this scenario with SMT enabled, reaching 3,193.11 samples per minute in Monster, 2,174.63 in Junkshop, and 1,608.79 in Classroom. The relatively small gap between the two platforms suggests broadly similar behavior under SMT-intensive rendering loads, with differences likely due to platform-level optimization rather than raw compute power.
| Blender CPU SMT (Samples per minute; higher is better) | Supermicro A+ Server AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (dual AMD EPYC 9965 192C) | 
|---|---|---|
| Monster | 3,070.84 | 3,193.11 | 
| Junkshop | 2,063.61 | 2,174.63 | 
| Classroom | 1,527.39 | 1,608.79 | 
With SMT disabled, the Supermicro system shows a notable increase in raw throughput across all three scenes. Performance reaches 4,018.10 samples per minute in Monster, 2,707.10 in Junkshop, and 1,990.51 in Classroom, demonstrating how Blender's CPU render engine can benefit from reduced thread contention on very high core count systems.
In this configuration, the Dell PowerEdge R7725 still shows higher absolute throughput, with 4,304.21, 2,870.34, and 2,079.79 samples per minute in Monster, Junkshop, and Classroom, respectively. Although the Dell system retains an advantage in peak performance, the Supermicro AS-2126HS-TN shows substantial gains with SMT disabled, highlighting the significant influence of multithreading strategy on the performance of platforms of this scale.
| Blender CPU without SMT (Samples per minute; higher is better) | Supermicro A+ Server AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (dual AMD EPYC 9965 192C) | 
|---|---|---|
| Monster | 4,018.10 | 4,304.21 | 
| Junkshop | 2,707.10 | 2,870.34 | 
| Classroom | 1,990.51 | 2,079.79 | 
y-cruncher
y-cruncher is a multithreaded and scalable program capable of calculating Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application among overclockers and hardware enthusiasts.
In the y-cruncher benchmark, the Supermicro AS-2126HS-TN demonstrates consistent and predictable scaling as problem sizes increase. The system performs the calculation of one billion decimal places in 8.092 seconds, then moves to 15.245 seconds for 2.5 billion decimal places and 24.961 seconds for 5 billion decimal places.
As the workload increases, computation times evolve linearly, reaching 44.350 seconds for 10 billion digits, 114.107 seconds for 25 billion, and 246.688 seconds for 50 billion. In the most demanding calculation, involving 100 billion digits, the system performs the operation in 572.800 seconds, demonstrating its ability to handle a high number of threads over long execution periods.
The Dell PowerEdge R7725 executes the same workloads slightly faster across all test sizes, completing the 100 billion digit run in 481.207 seconds.
| Y-Cruncher (total computation time) | Supermicro A+ Server AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (dual AMD EPYC 9965 192C) | 
|---|---|---|
| 1 billion | 8.092 s | 7.879 seconds | 
| 2.5 billion | 15.245 s | 13.811 seconds | 
| 5 billion | 24.961 s | 22.107 seconds | 
| 10 billion | 44.350 | 40.111 seconds | 
| 25 billion | 114.107 | 98.445 seconds | 
| 50 billion | 246.688 | 211.567 seconds | 
| 100 billion | 572.800 | 481.207 seconds | 
Phoronix Benchmarks
Phoronix Test Suite is an open-source automated benchmarking platform that supports over 450 test profiles and over 100 test suites via OpenBenchmarking.org. It manages the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration.
Stream Memory Bandwidth
In the Stream memory bandwidth test, the Supermicro AS-2126HS-TN achieves a sustained throughput of 807,766 MB/s. The Dell PowerEdge R7725, meanwhile, peaks at 883,312 MB/s, indicating slightly higher maximum memory bandwidth. However, both systems are within the expected performance range for dual-processor EPYC 9005 platforms.
7-Zip Compression
For 7-Zip compression, the Supermicro system achieves 1,262,832 MIPS, demonstrating excellent integer compute performance and efficient multithreaded scaling. The Dell system shows 1,326,967 MIPS, representing a slight advantage for this workload, while remaining in the same overall performance category.
Kernel Compilation
During kernel compilation (allmod), the Supermicro AS-2126HS-TN executes the task in 117.97 seconds, ahead of the Dell PowerEdge R7725 which performs the same operation in 139.36 seconds. This result highlights the efficiency of the Supermicro platform in parallel compilation scenarios common in development and continuous integration environments.
Apache Web Server
The Apache performance test shows the Supermicro system processing 90,623.69 requests per second, while the Dell system reaches 96,782.75 requests per second. Both systems show high throughput for web servers, although the Dell configuration shows a slightly higher request peak.
OpenSSL Verification
In OpenSSL tests, the Supermicro AS-2126HS-TN reaches 3.55 TB/s, demonstrating significant cryptographic throughput, suited to encryption-intensive workloads. The Dell PowerEdge R7725 goes even further with 4.41 TB/s, indicating higher peak cryptographic performance. Both platforms offer performance well beyond typical enterprise needs.
| Phoronix Benchmarks | Supermicro A+ Server AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (dual AMD EPYC 9965 192C) | 
|---|---|---|
| Stream | 807,766.0 MB / s | 883,312.0 MB / s | 
| 7-ZIP | 1,262,832 MIPS | 1,326,967 MIPS | 
| Kernel Compilation (allmod) | 117.97 seconds | 139.356 seconds | 
| Apache (requests per second) | 90,623.69 R/s | 96,782.75 R/s | 
| OpenSSL | 3,545,769,484,910 verifications | 4,409,642,672,307 verifications | 
Conclusion
The Supermicro A+ AS-2126HS-TN server offers a balanced dual-processor Turin platform, prioritizing compute density, power scalability, and flexible PCIe expansion in a 2U form factor. Compatible with two AMD EPYC 9005 series processors, high-TDP configurations, and flexible I/O configurations, it is ideal for enterprise, cloud, and high-performance computing (HPC) environments where sustained parallel performance and configurability are essential.
In our enterprise test suite, the AS-2126HS-TN delivered consistent and predictable performance for CPU-intensive workloads. Although it slightly trailed on some absolute throughput metrics, the Supermicro platform remained competitive, demonstrating excellent scalability and high efficiency for rendering, compute, and mixed workloads. These results confirm that the AS-2126HS-TN is a high-performance and flexible foundation for high core count deployments, especially when platform versatility and power headroom are as important as peak benchmark performance.
