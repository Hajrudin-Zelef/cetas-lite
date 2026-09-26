---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887-3
title: "fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "accelerator", "benchmark", "benchmarks", "compute", "distribution", "energy", "gpu", "gpus", "inference", "latency", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887.md
source_anchor: ""
source_lines: [87, 120]
sha256: 2bd12816617bcf2bc6b2123c4fcbc5c9ec2be60e10217a09e34680049de27589
---

# fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887

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
