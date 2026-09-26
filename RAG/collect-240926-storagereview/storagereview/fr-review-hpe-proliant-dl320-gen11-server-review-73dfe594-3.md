---
id: collect-240926-storagereview/storagereview/fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594-3
title: "fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "benchmarks", "compute", "cost", "dram", "gpu", "gpus", "inference", "intel", "memory", "nvidia", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594.md
source_anchor: ""
source_lines: [47, 93]
sha256: 8139103791fc84f4cd29b76578315fb855c9e7b547c60a20b471dd88fcf014e8
---

# fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594

HPE also offers GPU server configurations with the DL320 Gen11. HPE supports up to four single-width or two double-width GPUs in the front cage with the GPU CTO server. This means customers can configure the server with four NVIDIA L4 GPUs or two NVIDIA L40 GPUs, for example, to have a highly cost-effective inference machine for Edge use cases.
Additionally, a configuration option for the GPU variant of the DL320 Gen11 allows for the installation of eight E3.S SSDs, replacing the 4x U.3 storage chassis.
HPE ProLiant DL320 Gen11 Management – iLO 6
The DL320 Gen11 features an advanced management feature suite through the integration of iLO 6, HPE's latest version of its Lights-Out technology. iLO 6 brings a range of benefits (such as remote server configuration, health monitoring, and power and temperature control) that enhance the manageability, security, and efficiency of the DL320 and simplify complex IT environments. While iLO is not new and is commonly found on most HPE servers, we note it here because it represents considerable added value for those who may not be used to such in-depth server management in this particular price range.
iLO's advanced security features are designed to protect against threats and ensure data integrity while providing a streamlined interface and enhanced automation capabilities to facilitate a more intuitive user experience. Ultimately, this will reduce the time and effort required for various management tasks and make your life easier as an administrator.
To enhance the security of your DL320, iLO 6 incorporates HPE's exclusive silicon root of trust, ensuring that only firmware signed by HPE boots to establish a secure foundation from the start. This security is further reinforced by HPE Secure Start, which authenticates firmware via this silicon root of trust to ensure that the booted firmware is safe and secure. HPE iLO 6 also enforces the use of cryptography and algorithms compliant with CNSA (Commercial National Security Algorithm) standards and introduces support for 2-factor authentication via PIV/CAC cards.
iLO 6's runtime firmware verification also regularly checks critical firmware for intrusions after boot. If it detects a breach, it allows for quick restoration of firmware to factory or last known secure settings to help minimize potential damage caused by the intruder.
HPE ProLiant DL320 Gen11 Full Specifications
| Specifications | DETAILS | 
| Processor Type | Intel | 
| Processor Family | 4th Generation Intel Xeon Scalable Processors | 
| Processor Models |  | 
| Number of Processors | 1 | 
| Available Processor Cores | 8 to 32 cores, depending on processor | 
| Processor Cache | 26.25 – 60 MB L3, depending on processor | 
| Processor Speed | Up to 3.7 GHz, depending on processor | 
| Power Supply Type | HPE Flex Slot Power Supplies (options vary: 500 W, 800 W, 1000 1600 W, 1800 2000 W, XNUMX XNUMX-XNUMX XNUMX W) | 
| Expansion Slots | Maximum, 2 PCIe Gen5 and 1 OCP 3.0 PCIe Gen5 | 
| Maximum Memory | 2.0 TB per socket, single socket, when equipped with 128 GB DDR5 memory. | 
| Memory Slots | 16 DIMM slots per socket | 
| Memory Type | HPE DDR5 Smart Memory | 
| Memory Protection Features | HPE Fast Fault Tolerant Advanced ECC Memory, online spare, mirrored memory | 
| Optical Drive Type | Optional HPE 9.5 mm SATA DVD-RW Optical Drive, HPE Mobile USB DVD-RW Drive | 
| System Fan Features | Standard or high-performance fan kit, depending on model | 
| Network Controller | Wide range of options (PCIe riser and OCP3.0 adapter) | 
| Storage Controller | Integrated SATA Controller (AHCI or Intel SATA Software RAID), optional – HPE Smart Array Gen11 Controller | 
| DIMM Capacity | 16 GB to 256 GB | 
| Infrastructure Management | HPE GreenLake for Compute Ops Management, HPE iLO Standard, HPE OneView Standard, optional – HPE iLO Advanced, HPE OneView Advanced | 
| Warranty | 3/3/3: Three years of coverage for parts, labor, and on-site support. Additional coverage is available. | 
| Supported Drives | Up to 8+2 SFF SAS/SATA HDDs or SATA/SAS/NVMe U.2 or U.3 SSDs, depending on model | 
HPE ProLiant DL320 Gen11 vs DL360 Gen11
With several 1U Intel servers in the portfolio, comparing the available configurations of the HPE ProLiant DL320 Gen11 and DL360 Gen11 servers reveals distinct differences suited to different use cases and performance needs. The most obvious difference lies in the processor choice. The DL320 Gen11 offers 32-core processors, and the DL360 Gen11 offers two processors and accommodates Intel's best processor.
The DL320 Gen11 offers lower base memory configurations: 2 TB per socket (single socket) when each of the 16 DIMMs is equipped with 128 GB DDR5 memory keys. This highlights its adaptability to both economical and more demanding scenarios. The DL360 Gen11 doubles this memory total to 4 TB per socket when equipped with 256 GB DDR5 RAM keys.
The PCIe configurations of the DL320 Gen11, typically one Gen5 slot and one OCP 3.0 Gen5 slot, address businesses needing moderate extensibility, perhaps for additional network cards or reasonable external storage. The DL360 Gen11 goes further by offering more expansion slots. This indicates a design suited to environments where high-throughput I/O and extensive external connectivity are crucial, similar to cluster computing or high-speed storage networks.
The storage capabilities of the DL320 Gen11, which support a combination of SFF and LFF drives, suggest versatility, catering to a range of storage requirements from high-speed SSDs to higher-capacity hard drives. The DL360 Gen11's storage controllers, such as the HPE MR408i-o Gen11, focus on high-performance storage solutions to support more data-intensive applications.
HPE ProLiant DL320 Gen11 Performance
Our HPE ProLiant DL320 Gen11 server version features a base configuration suited to SMBs. As such, we don't expect much performance in our benchmarking. That said, many tasks in which these servers are deployed are not particularly intensive, so the performance data is instructive in terms of scalability.
At its core, our DL320 Gen11 review includes the Intel Xeon-G 32 at 6430 cores and 32 GB of DDR5 RAM (via two 16 GB keys) and includes two 1.92 TB MV SFF BC RI SATA SSDs. It also has two 800 W FS Platinum Hot-Plug LH power supplies for reliable and efficient power.
To complete the build, an HPE MR408i-o Gen11 SPDM storage controller, which offers RAID support, enhanced data protection, and a 1-port 4 GB BASE-T adapter to handle network activity.
The server has no GPU and its configuration is light on DRAM. So we adapted our benchmarking approach to focus on tests that heavily leverage the CPU's capabilities, meaning we will not run benchmarks such as ESRI, Luxmark, SPECworkstation 3, 7zip, and SPECviewperf 2020. Instead, our evaluation will focus on benchmarks that primarily emphasize the CPU to clearly show its performance in compute-intensive tasks.
Blender OptiX
The first is the Blender test, an open-source 3D modeling application. This benchmark was run using the Blender Benchmark utility. The score is expressed in samples per minute, with the higher being the better.
| HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor (32 cores, 3.7 GHz max)) |  | 
| Blender OptiX version 3.6 (CPU) (Samples per minute; higher is better) | Score | 
| Monster | 279 | 
| Junkshop | 177 | 
| Classroom | 135 | 
