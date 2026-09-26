---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r5715-review-bc780bb6-2
title: "fr-review-dell-poweredge-r5715-review-bc780bb6"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Microsoft"]
dates: []
keywords: ["amd", "compute", "datacenter", "ethernet", "gpu", "gpus", "liquid cooling", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r5715-review-bc780bb6.md
source_anchor: ""
source_lines: [3, 61]
sha256: 656b1cd518fe18bf62735e198486ba1cd293b5a63e192972bb1f898e85c9fe43
---

# fr-review-dell-poweredge-r5715-review-bc780bb6

The PowerEdge R5715 is the second model in Dell's 17th generation PowerEdge lineup, dedicated to SMBs. It distinguishes itself from its 1U counterpart through different priorities. While the R4715 prioritizes compute density and core count efficiency per rack unit, the R5715 is designed around storage capacity and I/O expandability in a 2U single-socket form factor. Readers of our R4715 review will recognize the platform fundamentals: same 5th generation AMD EPYC processor family, same DDR5 memory architecture with 24 slots, and same iDRAC 10 management system. Only the task assigned to the R5715 differs slightly.
Our test server was equipped with an AMD EPYC 9015 processor, the 8-core model in the Turin lineup, paired with 384 GB of DDR5 memory and a BOSS RAID 1 configuration for boot. Our testing focused on the R5715's 12-bay 3.5-inch storage backplane, an area where the 9015 makes perfect sense. Workloads such as file sharing, data backup, and point-of-sale video surveillance don't need 32 cores; they prioritize storage density, sustained throughput, and reliable management. The 9015 helps limit power consumption and licensing costs while offering up to 288 TB of raw storage capacity in a single 2U node.
The R5715 increases the number of PCIe Gen5 ports to four, compared to three for the R4715, and adds an additional OCP 3.0 network port, providing greater flexibility to meet growing I/O demands. Both platforms support 100 GbE and 400 GbE protocols via PCIe AIC, making them perfectly suited for environments requiring high bandwidth. However, neither officially supports Fibre Channel connectivity. Additionally, neither supports GPUs or DPUs. They run on the same 800 W and 1100 W power supplies, available in Platinum and Titanium versions, with fault-tolerant redundancy and air cooling.
Dell PowerEdge R5715 Specifications
The table below highlights the physical and hardware specifications of the Dell PowerEdge R5715 platform.
| Specifications | Dell PowerEdge R5715 | 
|---|---|
| Processor |  | 
| Processor | One 5th generation AMD EPYC 9005 series processor, up to 32 cores | 
| Form Factor | 2U rack server | 
| Memory |  | 
| DIMM Slots | 24 DDR5 DIMM slots | 
| Maximum Memory | 1.5 TB (up to 64 GB per DIMM) | 
| Memory Speed | Up to 5200 MT/s | 
| Memory Type | Registered DDR5 ECC RDIMM modules only | 
| Storage |  | 
| Internal Controllers (RAID) | PERC H365i, H965i | 
| Internal Boot | BOSS-N1 DC-MHS | 
| External HBA | N/A | 
| Front Drive Bays | 12 x 3.5-inch SAS/SATA ports 16 x 2.5-inch SAS/SATA ports | 
| Engine Tuning |  | 
| Power Supplies | 800 W Platinum, 1100 W Titanium 800 W, 1100 W FTR supported | 
| Cooling and Fans |  | 
| Cooling Options | air cooling | 
| Fans | Up to six hot-swappable fans | 
| Dimensions |  | 
| Height | 86.8 mm (3.41 inches) | 
| Width | 482.0 mm (18.97 inches) | 
| Depth (with bezel) | 802.4 mm (31.59 inches) | 
| Depth (without bezel) | 801.51 mm (31.55 inches) | 
| Bezel | Optional metal bezel | 
| Networking and Expansion |  | 
| OCP Network Options | 2 OCP 3.0 network cards (optional), 1 GbE, 10 GbE, 25 GbE Slot 4: 1×16 Gen5 OCP 3.0 Slot 10: 1×16 Gen5 OCP 3.0 | 
| Integrated Network Card | Dedicated 1 Gb BMC Ethernet port | 
| PCIe AIC Network Card | 100 GbE and 400 GbE; NDR VPI (400 GbE) | 
| PCIe Slots | Up to 4 PCIe Gen5 slots (x16 connectors) Slot 2: 1×16 Gen5 Full Height Slot 3: 1×16 Gen5 Full Height Slot 7: 1×16 Gen5 Full Height Slot 9: 1×16 Gen5 Full Height | 
| GPU Options | N/A | 
| Ports |  | 
| Front Ports | 1 USB 2.0 Type-A port (optional KVM LCP) 1 USB 2.0 Type-C port (HOST/BMC Direct) 1 MiniDisplayPort (optional KVM LCP) | 
| Rear Ports | 2x USB 3.1 Type-A 1x VGA Dedicated 1 Gb BMC Ethernet port | 
| Internal Ports | 1x USB 3.1 Type-A | 
| Management |  | 
| Integrated Management | iDRAC10, iDRAC Direct, iDRAC RESTful API with Redfish, RACADM command-line interface, Quick Sync 2 wireless module | 
| OpenManage Software | OpenManage Enterprise (OME), OME Power Manager, OME Services, OME Update Manager, OME APEX AIOps Observability, OME Integration for VMware vCenter, OME Integration for Microsoft System Center, OpenManage Integration for Windows Admin Center | 
| Tools | IPMI | 
| Integrations | OpenManage Integrations: Red Hat Ansible Collections, Terraform providers | 
| Change Management | Dell Repository Manager, Dell System Update, Enterprise Catalogs, Server Update Utility (SUU) | 
| Security |  | 
| Security Features | Cryptographically signed firmware, data-at-rest encryption (SED with local or external key management), secure boot, secure component verification (hardware integrity check), secure erase, silicon root of trust, system lock (requires iDRAC10 Enterprise or Datacenter), FIPS/CC-TCG certified TPM 2.0, chassis intrusion detection, AMD Secure Encrypted Virtualization (SEV), AMD Secure Memory Encryption (SME) | 
| Operating Systems and Hypervisors |  | 
| Supported Operating Systems/Hypervisors | Canonical Ubuntu Server LTS, Microsoft Windows Server with Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi | 
The Dell PowerEdge R5715 is a single-processor 2U rack server based on the 5th generation AMD EPYC 9005 series platform. Designed as a high-performance storage platform for businesses requiring high capacity and reliable I/O connectivity without the overhead of a dual-processor architecture, the R5715 targets workloads such as databases, file shares, backups, and virtualization, where a single powerful EPYC processor handles the load more efficiently than two previous-generation processors. We also used this chassis for our power consumption comparison between hard drives and Micron 6600 ION flash memory, replacing eight 30 TB hard drives with a single 245 TB SSD. Supporting up to 288 TB of raw storage and featuring four PCIe Gen5 expansion slots, the R5715 delivers performance well above its price category.
Exterior and Front Panel
The R5715 ships with an optional metal bezel featuring Dell's iconic hexagonal pattern. This bezel fits perfectly onto the chassis and reveals the front panel controls on the right side: a power button, a USB 2.0 Type-C port for direct BMC access, an iDRAC Direct port, and a system identification button. Without the bezel, the chassis measures 3.41 cm in height, 18.97 cm in width, and 31.55 cm in depth, allowing it to fit into standard 2U racks. Build quality is impeccable, with tool-less drive bay latches and blue retention clips used consistently on internal components for quick access.
Storage Configuration
The tested unit is equipped with a 12 x 3.5-inch SAS/SATA front bay, with four bays occupied by 20 TB 6 Gb/s 7,200 RPM SATA hard drives and eight bays free for future expansion. An alternative configuration with a 16 x 2.5-inch SAS/SATA backplane is also available, depending on needs. RAID management is handled by the internal PERC H365i controller or the more powerful PERC H965i. Boot is managed separately by a dedicated BOSS-N1 DC-MHS module at the rear, isolating the operating system from the data. This clever design avoids the common mistake of running the operating system and data storage on the same bay.
Processor and Cooling
The R5715 is a single-socket platform based on the AMD EPYC 9005 series processor, supporting up to 32 cores. Its imposing heatsink is a finned tower model incorporating copper heat pipes, secured to the SP5 socket by six captive screws. Cooling is entirely air-based; up to six hot-swappable fans circulate air from front to rear of the chassis. Liquid cooling is not compatible with this platform.
Memory
