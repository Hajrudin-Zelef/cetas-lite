---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf-2
title: "fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "datacenter", "ethernet", "gpu", "gpus", "intel", "license", "liquid cooling", "memory", "nvidia", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf.md
source_anchor: ""
source_lines: [3, 65]
sha256: c1fe5e88a2a3d63384aa9b2743991066f055a1dce4d08bb95421667503cd264d
---

# fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf

Dell's 17th generation PowerEdge lineup is already well established, and with the R4715 and R5715, it now targets SMBs and mid-sized businesses more specifically. These two single-processor servers are based on the same 5th generation AMD EPYC architecture as the entire 17th generation PowerEdge lineup. They are optimized for organizations where the right number of cores, license management, and ease of use take priority over maximum throughput. The R4715 is a more compact 1U model, designed for virtualization, large-scale databases, and network edge deployments. The R5715 adopts a 2U form factor, offering more drive bays and PCIe expansion, and is suited to configurations where storage capacity and I/O performance are essential.
The R4715 server is designed for organizations running virtualization workloads, large-scale databases, and edge computing, for which licensing efficiency and ease of use are essential. Our test model was equipped with an AMD EPYC 9335 processor, the top-tier 32-core processor available on this platform, paired with 384 GB of DDR5 memory and a BOSS RAID 1 boot configuration. In addition to its 32 cores, the 9335 offers 128 MB of L3 cache and a TDP of 210 W. Users who don't need as many cores can opt for the EPYC 9255 (24 cores), EPYC 9135 (16 cores), or EPYC 9015 (8 cores) models.
A few important points to clarify upfront about the platform: the R4715 does not support GPUs or DPUs. This is not an oversight. This platform is designed specifically for CPU-focused workloads, and Dell made deliberate compromises to limit component costs and footprint. For workloads requiring accelerators, the R6715 and R7715 models are better suited.
The R4715 offers a compact air-cooled chassis, with up to three PCIe Gen5 slots, 24 DDR5 RDIMM slots, and flexible 3.5-inch and 2.5-inch storage options, including U.2 NVMe. Additionally, it integrates iDRAC10 with OpenManage Enterprise and hardware security via a silicon-level root of trust. Network options include 25 GbE via OCP 3.0, 100 GbE, and 400 GbE via PCIe AIC. Broadcom, Intel, and NVIDIA complete the network card ecosystem; note that no Fibre Channel connectivity is officially supported. It runs on 800 W or 4715 1100 W power supplies, available in Platinum or Titanium versions, and supports fault-tolerant redundancy. For an entry-level server, the enterprise management and security features are generally satisfactory.
Dell PowerEdge R4715 Specifications
The table below highlights the physical and hardware specifications of the Dell PowerEdge R4715 platform.
| Specifications | Dell PowerEdge R4715 | 
|---|---|
| Processor |  | 
| Processor | One 5th generation AMD EPYC 9005 series processor, up to 32 cores | 
| Form factor | 1U rack server | 
| Memory |  | 
| DIMM slots | 24 DDR5 DIMM slots | 
| Maximum memory | 1.5 TB (up to 64 GB per DIMM) | 
| Memory speed | Up to 5200 MT/s | 
| Memory type | Registered DDR5 ECC RDIMM modules only | 
| Storage |  | 
| Internal controllers (RAID) | PERC H365i, H965i | 
| Internal boot | BOSS-N1 DC-MHS | 
| External HBAs | N/A | 
| Front drive bays | 4x 3.5-inch SAS 8-port SAS/SATA 2.5-inch 8-port U.2 NVMe G4 | 
| Tuning Engine |  | 
| Power supplies | Platinum 800 W, 1100 W Titanium 800 W, 1100 W FTR supported | 
| Cooling and fans |  | 
| Cooling options | air cooling | 
| Fans | Up to four assemblies (dual-fan module) of hot-swappable fans | 
| Dimensions |  | 
| Height | 42.8 mm (1.68 inches) | 
| Width | 482.0 mm (18.97 inches) | 
| Depth (with bezel) | 816.921 mm (32.16 inches) | 
| Depth (without bezel) | 815.141 mm (32.09 inches) | 
| Bezel | Optional metal bezel | 
| Networking and expansion |  | 
| OCP network options | 2 OCP 3.0 network cards (optional), 1 GbE, 10 GbE, 25 GbE Slot 2: 1×16 Gen5 OCP 3.0 Slot 5: 1×16 Gen5 OCP 3.0 | 
| Integrated network card | Dedicated 1 Gb BMC Ethernet port | 
| PCIe AIC network card | 100 GbE and 400 GbE; NDR VPI (400 GbE) | 
| PCIe slots | Up to 3 PCIe Gen5 slots (x16 connectors) Slot 1: 1×16 Gen5 full height or low profile Slot 2: 1×16 Gen5 Low Profile or 1×16 OCP3.0 Slot 4: 1×16 Gen5 full height or low profile | 
| GPU options | N/A | 
| Ports |  | 
| Front ports | 1 USB 2.0 Type-A port (optional KVM LCP) 1 USB 2.0 Type-C port (HOST/BMC Direct) 1 MiniDisplayPort (optional KVM LCP) | 
| Rear ports | 2x USB 3.1 Type-A 1x VGA Dedicated 1 Gb BMC Ethernet port | 
| Internal ports | 1x USB 3.1 Type-A | 
| Management |  | 
| Integrated management | iDRAC10, iDRAC Direct, iDRAC RESTful API with Redfish, RACADM command-line interface, Quick Sync 2 wireless module | 
| OpenManage software | OpenManage Enterprise (OME), OME Power Manager, OME Services, OME Update Manager, OME APEX AIOps Observability, OME Integration for VMware vCenter, OME Integration for Microsoft System Center, OpenManage Integration for Windows Admin Center | 
| Tools | IPMI | 
| Integrations | OpenManage Integrations: Red Hat Ansible Collections, Terraform providers | 
| Change management | Dell Repository Manager, Dell System Update, Enterprise Catalogs, Server Update Utility (SUU) | 
| Security |  | 
| Security features | Cryptographically signed firmware, data-at-rest encryption (SED with local or external key management), secure boot, secure component verification (hardware integrity check), secure erase, silicon root of trust, system lockdown (requires iDRAC10 Enterprise or Datacenter), FIPS/CC-TCG certified TPM 2.0, chassis intrusion detection, AMD Secure Encrypted Virtualization (SEV), AMD Secure Memory Encryption (SME) | 
| Operating systems and hypervisors |  | 
| Supported operating systems/hypervisors | Canonical Ubuntu Server LTS, Microsoft Windows Server with Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi | 
Dell PowerEdge R4715 Design and Installation
The R4715 is a 1U rack server measuring 1.68 mm in height, 18.97 mm in width, and 32.09 mm in depth without the optional metal bezel (32.16 mm with the bezel). The front panel includes a power button, a system identification button, a USB 2.0 Type-A port (used with the optional KVM LCP module), a USB 2.0 Type-C port for direct iDRAC access, and an optional MiniDisplayPort for the same KVM configuration. The drive bays open tool-free across the entire chassis.
Storage Configuration
The R4715 is available with three front bay configurations: 4x 3.5-inch SAS bays, 8x 2.5-inch SAS/SATA bays, or 8x U.2 NVMe Gen4 bays. The latter two configurations share the same 8-bay 2.5-inch form factor, with the choice of backplane determining drive protocol compatibility. Internal RAID is handled by the PERC H365i controller or the higher-performance PERC H965i controller. OS booting uses a dedicated BOSS-N1 DC-MHS module, which completely isolates the boot volume from the data storage pool and avoids having to allocate OS space within an active bay.
The test unit was shipped with a single 480 GB SATA SSD and two 1.92 TB U.2 NVMe drives.
Processor and Memory
The R4715 motherboard supports one 5th generation AMD EPYC 9005 series processor, up to 32 cores. Memory is handled by 24 DDR5 DIMM slots compatible only with RDIMM modules (neither UDIMM nor LRDIMM). Maximum capacity is 1.5 TB with 64 GB DIMM modules per slot, for speeds up to 5,200 MT/s.
Cooling
The R4715 is exclusively air-cooled. The processor is equipped with a five-heatpipe heatsink and a sizable fin stack that extends into the space usually empty next to the socket, thereby increasing the heat exchange surface area. Ventilation is provided by eight hot-swappable high-performance fan modules, arranged in groups, which move air from the front to the rear of the chassis. No liquid cooling system is available on this platform.
Tuning Engine
The R4715 supports hot-swappable redundant power supplies, available in two power ratings: 800 W and 1,100 W, each offered with 80 PLUS Platinum or Titanium certifications. FTR (Flex Titanium Rating) certification is also supported across the entire power supply range.
