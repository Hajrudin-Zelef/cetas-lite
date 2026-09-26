---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf-2
title: "fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia"]
dates: []
keywords: ["amd", "energy", "gpu", "gpus", "memory", "nvidia", "parameters", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf.md
source_anchor: ""
source_lines: [3, 80]
sha256: 3d454b9dbc38e968582e13cab21fe3ba85a17dbab5939eaedaf103e9ba598c42
---

# fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf

The Supermicro AS-1115SV-WTNRT server is a versatile and robust 1U rackable server, designed to meet the demands of the most resource-intensive applications such as virtualization, database management, and edge computing. Suited for environments requiring high throughput and high reliability, this server offers a balanced architecture integrating the AMD EPYC 8004 series processor, thus providing a scalable solution to modern computing challenges.
Supermicro AS-1115SV-WTNRT Main Features and Hardware Specifications
At the heart of the AS-1115SV-WTNRT is a single-socket configuration using the AMD EPYC 8004 series processor, capable of handling up to 64 cores and 128 threads, which allows for significant parallel processing power. Designed for dense data centers and service providers, the "Siena" series processors offer an optimal combination of high core count and energy efficiency at a competitive price, operating within a modest power envelope starting at just 70 W and reaching up to 200 W.
The server supports up to 576 GB of DDR5 4800 2.5 MHz memory across six DIMM slots, ensuring high-speed data transfer and efficient management of large datasets. Storage versatility is another important attribute, with ten hot-swappable XNUMX″ drive bays that support a combination of NVMe, SAS, and SATA drives.
The Supermicro AS-1115SV-WTNRT is designed with a compact 1U form factor, measuring 437 mm x 43 mm x 597 mm, optimizing space without compromising internal hardware capacity. Expansion capabilities include three PCIe 5.0 x16 slots, two of which are full-height and full-length slots, providing ample space for larger cards, while the third is a low-profile slot, suitable for smaller expansion cards.
The server features two redundant 860 W power supplies for continuous power and system reliability.
Under the hood, the server's internal layout optimizes airflow and cooling efficiency, ensuring consistent performance under various operating conditions, including the six internal counter-rotating fans (FAN-0163L4).
The smaller H13SVW-NT motherboard is in the foreground, housing a single AMD EPYC 8004 series processor in an SP6 socket, surrounded by its six DDR5 DIMM slots.
Supermicro AS-1115SV-WTNRT System Management
Supermicro has equipped the AS-1115SV-WTNRT with a robust suite of management tools, including SuperCloud Composer and Supermicro Server Manager (SSM), which facilitate streamlined system management and monitoring. The BIOS is equipped with the contemporary UEFI 2.8 specification, offering enhanced security, faster boot times, and support for larger disk partitions.
The AS-1115SV-WTNRT also features a comprehensive baseboard management controller (BMC) system that provides administrators with control and monitoring capabilities. This BMC interface (accessible via a dedicated IP address) provides critical information about system status, allows for comprehensive configuration changes, and enables remote control actions without physical access to the hardware. It is an essential tool for efficiently managing server operations, ensuring availability, and quickly resolving issues.
The dashboard offers a comprehensive overview of server status and essential metrics. It presents system status indicators, IP addresses, and firmware versions, providing a snapshot of critical operational data. Power consumption is displayed graphically over time, detailing minimum, average, and maximum usage. This visualization allows administrators to assess at a glance the server's energy efficiency and operational costs. A remote console overview is also available, facilitating quick access to the server's current operational status, which is invaluable for rapid troubleshooting and continuous monitoring.
The CPU section offers detailed information about the server's processor specifics (in our case, the AMD EPYC 8534P 64-core processor), such as speed, thermal design power (TDP), number of cores, number of threads, and manufacturer. This section is crucial for verifying the processor's operational parameters and ensuring it operates within specified limits.
In the Firmware Update tab, users can manage and update the firmware of various server components, such as BMC, BIOS, and CPLD. The interface allows users to choose the firmware type, preserve existing configurations, and ensure that all critical updates are applied without affecting the server's operational settings. This function is essential for maintaining server security, as firmware updates often contain patches for vulnerabilities and performance improvements.
The Memory tab provides detailed information about all installed memory modules, displaying each module's health status, type, error correction code capacity, operating speed, size, and serial number. The Memory tab allows users to easily track and manage the server's physical memory, which is essential for diagnosing memory-related issues or planning upgrades.
The Power tab displays detailed power consumption statistics over different periods, including the last hour, day, and week. It shows historical trends and peak values, useful for capacity planning and operational analysis. Users and administrators can also use this data to optimize power settings, schedule low-power modes during off-peak hours, and make informed decisions regarding energy usage and efficiency.
Supermicro AS-1115SV-WTNRT Specifications
| Processor |  | 
| Processor | Single Processor(s) – AMD EPYC 8004 Series Processor | 
| Number of Cores | Up to 64C/128T | 
| Note | Supports processors with TDP up to 225 W (air-cooled) | 
| GPU |  | 
| Maximum Number of GPUs | Up to 1 double-width GPU or 2 single-width GPUs | 
| System Memory |  | 
| Memory | Number of Slots: 6 DIMM slots – Maximum Memory (1DPC): up to 576 GB 4800 5 MT/s ECC DDRXNUMX RDIMM | 
| Onboard Devices |  | 
| Chipset | System on Chip | 
| Network Connectivity | 2 RJ45 10GBASE-T with Broadcom BCM57416 | 
| Input / Output |  | 
| LAN | 1 dedicated IPMI LAN port(s) RJ45 1 GbE – 2 RJ45 10 GBASE-T LAN port(s) | 
| USB | 2 ports (front) – 4 ports (rear) | 
| Video | 1 VGA port(s) | 
| Serial Port | 1 COM port(s) (rear) | 
| TPM | 1 onboard TPM/80 port | 
| System BIOS |  | 
| BIOS Type | AMI 32 MB SPI Flash EEPROM | 
| BIOS Features | ACPI 6.4, SMBIOS 3.5 or later, UEFI 2.8, Plug and Play (PnP), USB keyboard support | 
| Management |  | 
| Software |  | 
| Power Configurations | Redundant 1U 800/860 W Titanium AC Power Supply | 
| Security |  | 
| Hardware | Trusted Platform Module (TPM) 2.0, Silicon Root of Trust (RoT) – NIST 800-193 Compliant | 
| Feature |  | 
| PC Health Monitoring |  | 
| FAN |  | 
| Temperature |  | 
| Chassis |  | 
| Form Factor | 1U | 
| Model | CSE-116BTS-R000WNP | 
| Dimensions and Weight |  | 
| Height | 1.7 mm (43 in) | 
| Width | 17.2 mm (437 in) | 
| Depth | 23.5 mm (597 in) | 
| Package | 7.75 "(H) x 23.5" (L) x 31.5 "(D) | 
| Weight |  | 
| Available Color | Black | 
| Front Panel |  | 
| LED |  | 
| Buttons |  | 
| Expansion Slots |  | 
| PCI-Express (PCIe) Configuration |  | 
| M.2 | 2 M.2 PCIe 3.0 x4 NVMe slots (M-key) | 
| Drive/Storage Bays |  | 
| Drive Bay Configuration |  | 
| M.2 | 2 M.2 PCIe 3.0 x4 NVMe slots (M-key 2280/22110) View M.2 options | 
| System Cooling |  | 
| Fans | 6 robust 40 x 40 x 56 mm counter-rotating fans. | 
| Power Supply |  | 
|  |  | 
Supermicro 1115SV-WTNRT Performance
Our Supermicro Storage A+ 1115SV-WTNRT review unit is configured with the following:
- Processor: AMD EPYC 8534P 64 cores
- RAM: 194GB DDR5
- GPU: NVIDIA A2 (15,360 6 MB GDDRXNUMX)
- SSD: Ultrastar SN655 NVMe Data Center SSD (15.36 TB)
- 2x Mellanox ConnectX-6 DX 100G
- 10 64-bit Windows
We recently tested the Graid SupremeRAID SR-1001 RAID solution with the Supermicro AS-1115SV-WTNRT SSD. To discover their respective performance, see our full test.
