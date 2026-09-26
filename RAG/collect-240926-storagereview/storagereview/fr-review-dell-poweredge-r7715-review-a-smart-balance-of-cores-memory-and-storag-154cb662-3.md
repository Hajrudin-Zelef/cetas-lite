---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662-3
title: "fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["memory", "amd", "ethernet", "gpus", "latency", "nand", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662.md
source_anchor: ""
source_lines: [29, 92]
sha256: f977b62075372c5557b918616e0e4a1d8fd775591ac07c89c77aa15615fcb72d
---

# fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662

The system also offers various power supply options. It supports hot-pluggable redundant power supplies from 800 W to 3200 W, providing flexibility suited to different needs and configurations. You have Titanium-class high-performance units (3200 W, 2400 W, 1800 W, 1500 W, 1100 W, and 800 W), as well as Platinum-class options (1100 W and 800 W). Special configurations are also available, such as 277 VCA and CCHT (3200 W and 1500 W), as well as a DC power supply (1400 W -48--60 VCC).
Dell PowerEdge R7715 and iDRAC 10
iDRAC, or Integrated Dell Remote Access Controller, is Dell's integrated remote management tool that simplifies monitoring, updating, and troubleshooting of PowerEdge servers, such as the R7715, without requiring physical presence. With the latest iDRAC10 version, Dell has made several improvements to enhance security and usability. It now incorporates a dedicated security processor with an integrated root of trust, improved encryption algorithms, and device-level attestation, making server management more secure than ever.
The user interface has also been redesigned for a more consistent experience across all Dell Technologies consoles, with simplified navigation that makes managing your server even more intuitive. Additionally, iDRAC10 allows for creating custom user roles and introduces a simplified licensing structure, specifically for 17th generation PowerEdge. AC power recovery is now managed directly by iDRAC, instead of being controlled by the BIOS, giving administrators more centralized control, an appreciated feature.
Dell PowerEdge R7715 specifications
| Feature | PowerEdge R7715 | 
| Processor | One 5th generation AMD EPYC 9005 Series processor with up to 160 cores for the Zen5 processor | 
| Chipset | AMD chipset | 
| Accelerators | Up to three 400 W double-width GPUs or six 75 W single-width GPUs | 
| Memory |  | 
| DIMM module speed | Up to 5200 MT/S | 
| Memory type | RDIMM | 
| Memory module slots | 24 DDR5 DIMM slots | 
|  | Supports only registered DDR5 ECC DIMM modules. | 
| Storage |  | 
| Front bays |  | 
| Rear bays | N/A | 
| Storage controllers |  | 
| Internal controllers | PERC H365i, H965i, H975i | 
| External controllers | HBA465e, H965e | 
| Software RAID | N/A | 
| Internal boot |  | 
| Power supply |  | 
| Cooling options |  | 
| Fans | Up to six hot-plug Gold/Very High Performance fans | 
| Ports |  | 
| Network options | Dedicated 1 Gb BMC Ethernet port | 
|  | 2 OCP NIC 3.0 cards (optional) | 
| Front ports | 1 x USB 2.0 Type-A (optional LCP KVM) | 
|  | 1 x USB 2.0 Type-C (HOST/BMC Direct) | 
|  | 1 x Mini-DisplayPort (optional LCP KVM) | 
| Rear ports | Dedicated 1 Gb BMC Ethernet port | 
|  | 2 x USB 3.1 | 
|  | 1 x VGA | 
| Internal ports | 1 USB 3.0 port (optional) | 
| Slots |  | 
| PCIe | Up to eight PCIe Gen5 slots | 
| Form factor | 2U rack server | 
| Height | 86.8 mm (3.41 inches) | 
| Width | 482.0 mm (18.97 inches) | 
| Depth | 802.4 mm (31.59 inches) with power handle | 
| Weight | Maximum 28.68 kg (63.22 lb) | 
| Bezel | Optional metal bezel | 
| System management |  | 
| Embedded management |  | 
| OpenManage Console |  | 
| Mobility | N/A | 
| Tools | IPMI | 
| Change Management |  | 
| OpenManage Integrations |  | 
| Security |  | 
| Operating systems and hypervisors |  | 
Dell PowerEdge R7715 design and construction
The PowerEdge R7715 features a clean design engineered for efficient cooling and consistent performance. Its modular, tool-less installation simplifies maintenance and upgrades, allowing you to maintain optimal operation with minimal downtime. It is also equipped with an optional locking bezel, which provides enhanced security and gives the front a clean, professional look. Easy to remove, it allows quick access for maintenance or upgrades, for simple and convenient use.
Once the front panel is removed, you can access the entire front panel. On the right side, the control panel includes the power button, a USB port (for connecting external devices such as USB keys or devices for maintenance or direct access), a micro iDRAC Direct port, and the iDRAC Direct status indicator.
The front and center consist of 32 1.6 TB EDSFF Gen5 E3.S NVMe drives (at least for our configuration), each operating on a PCIe Gen5 x2 interface. Although x4 connectivity is more common, this configuration leverages the higher bandwidth of Gen5 to maintain high throughput even at x2, thus preserving valuable PCIe lanes for other expansion needs without resorting to PCIe switches. Nevertheless, these drives are hot-swappable, allowing them to be added or replaced without shutting down the system. This design minimizes downtime and offers flexibility to adjust storage capacity as needed. The ventilation panel is located between the drive rows to ensure consistent airflow and maintain optimal operating temperatures.
The bottom of the rear panel houses several ports, including the BOSS-N1 DC-MHS module and a dedicated BMC Ethernet port for remote management via Open Server Manager. This configuration allows administrators to monitor and control the server from a separate network connection, essential for maintaining control in the event of a network outage or software failure.
Additionally, two USB 3.1 ports (9-pin and 3.0 compatible) provide high-speed connectivity for external devices such as USB keys or external hard drives. A VGA port also allows connecting displays, ensuring compatibility with existing systems. The system is equipped with two power supply units (PSU1 and PSU2) on each side of the rear panel, providing built-in redundancy. As always, this configuration ensures uninterrupted operation even in the event of a power supply failure, an essential feature for enterprise environments where availability is critical.
Removing the top panel of the Dell PowerEdge R7715 reveals a well-organized interior. The first thing you notice is likely the large black ventilation shrouds that cover most components and guide air to critical parts to ensure optimal cooling. PCIe expansion cards can be inserted into the metal risers, such as Riser 3 and Riser 5. Many blue plastic clips and tabs are scattered throughout, designed to simplify upgrades and maintenance as much as possible, without tools. Everything seems designed to ensure adequate airflow and preserve component accessibility.
Once the covers are removed, the internal board of the Dell PowerEdge R7715 is visible. It presents a well-organized configuration, including the AMD EPYC processor, DIMM slots, PCIe risers, BOSS-N1 DC-MHS modules, and fans.
At the center of the motherboard is the single-socket AMD EPYC 9665P processor, topped with an imposing heatsink featuring copper heat pipes for efficient cooling. Of the 24 available slots, 12 DIMM slots, equipped with 768 GB of RAM, surround the processor, providing significant room for future upgrades.
The image below shows a close-up of the PowerEdge R7715's BOSS-N1 DC-MHS module, equipped with a Micron 7450 NVMe SSD. This drive uses Micron's advanced 176-layer NAND technology, offering excellent latency and PCIe Gen4 performance. Designed for reliable and secure booting, the BOSS-N1 is generally dedicated to the server's operating system. BOSS SSDs can be configured in JBOD, RAID 0, or RAID 1. They are generally used in RAID 1 for durable boot storage. Their mounting is secured by blue tool-less retention clips, making installation and replacement easy. The Dell configuration also ensures good airflow and easy access for maintenance.
Here is a close-up of the rear of one of the system's power supplies. The black plastic housing surrounding the connector holds it firmly in place and protects it from accidental disconnection. The power cables, carefully organized and routed, are elegant, ensure optimal airflow, and easy maintenance.
Dell PowerEdge R7715 performance
