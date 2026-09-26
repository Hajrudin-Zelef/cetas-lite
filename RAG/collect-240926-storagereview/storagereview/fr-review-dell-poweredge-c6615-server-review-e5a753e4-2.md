---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-c6615-server-review-e5a753e4-2
title: "fr-review-dell-poweredge-c6615-server-review-e5a753e4"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "China", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "benchmark", "datacenter", "distribution", "ethernet", "intel", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-c6615-server-review-e5a753e4.md
source_anchor: ""
source_lines: [3, 57]
sha256: 45f2a70be9b729b2da21f627c1deb5da6aad0938d2085852cdb1e945341951d9
---

# fr-review-dell-poweredge-c6615-server-review-e5a753e4

The Dell PowerEdge C-Series platform has a 2U chassis supporting four servers in the Dell Modular Infrastructure category. Depending on the workload, the C-series system can be configured with two different node types: a single-socket AMD C6615 node or a dual-socket Intel C6620 node.
Our review will focus on the C-series chassis, which features four single-socket AMD EPYC nodes connected to an E8.S PCIe Gen3 5-bay disk backplane.
From a storage perspective, the platform can be configured with a 2.5-inch SFF disk backplane, which supports up to 24 NVMe SSDs or Gen5 support by leveraging a 3-bay E8.S backplane. Internally, these drives are connected directly to each node, with equal distribution across the four servers. For example, in the 24-bay configuration, each node sees six drives; in the 8-bay configuration, each node sees two drives.
The C6600 chassis provides shared redundant power supplies and cooling for the four installed nodes, though beyond that, each node is managed independently. So, unlike a blade chassis managed with a chassis management portal, this is more like four small PowerEdge servers under one metal roof. Each C6615 node has dedicated network connections, an iDRAC interface, and PCIe slots for expansion.
Dell PowerEdge C6615 Node Specifications
| C6615 Specifications |  | 
|---|---|
| Processor | One AMD EPYC processor with up to 64 cores | 
| Memory | 6 DDR5 DIMM slots, supports up to 576 GB RDIMM (6 x 96 GB), speeds up to 4800 MT/s | 
| Storage Controllers | Internal controllers (RAID): PERC H755N, PERC H355 Internal Boot: Boot Optimized Storage Subsystem (NVMe BOSS-N1): HWRAID 1, 2 x M.2 SSD Internal SAS 12 Gbit/s HBA (non-RAID): HBA355i Software RAID: S160 | 
| Availability | Hot-plug redundant hard drives and power supplies | 
| Drive Bays | Front bays: Up to 16 x 2.5-inch SAS/SATA (HDD/SSD) drives, 61 TB maximum Up to 16 x 2.5-inch SATA/NVMe drives, 15.36 TB maximum on a universal backplane configuration Up to 16 x 2.5-inch on NVMe backplane Up to 8 x E3.s on the NVMe SSD hard drive backplane | 
| Hot-swap, Redundant Power Supplies | 3200 W 277 VAC or 336 VDC 2800 W Titanium 200-240 VAC or 240 VDC Platinum 2400 W 100-240 VAC or 240 VDC 1800 W Titanium 200-240 VAC or 240 VDC | 
| Dimensions | Height – 40.0 mm (1.57 inches) Width – 174.4 mm (6.86 inches) Depth – 549.7 mm (21.64 inches), 561.3 mm (22.10 inches) – SAS/SATA or NVMe or E3.S or universal configuration | 
| Weight | 3.7 kg (8.15 pounds) | 
| Integrated Management | iDRAC9 iDRAC Direct RESTful API iDRAC with Redfish iDRAC Service Module | 
| OpenManage Software | CloudIQ plugin for PowerEdge OpenManage Enterprise OpenManage Enterprise Integration for VMware Vcenter OpenManage Integration for Microsoft System Center OpenManage Integration with Windows Admin Center OpenManage Power Manager plugin OpenManage Service plugin OpenManage Update Manager plugin | 
| Integrations | BMC TrueSight Microsoft System Center OpenManage Integration with ServiceNow OpenManage Integration with Windows Admin Center OpenManage Power Manager plugin OpenManage Service plugin OpenManage Update Manager plugin | 
| Security | AMD Secure Encrypted Virtualization (SEV) AMD Secure Memory Encryption (SME) Cryptographically signed firmware Data encryption at rest (SED with local or external key management) Secure BootSecured Component Verification (hardware integrity verification) Secure Erase Silicon Root of Trust System Lockdown (requires iDRAC9 Enterprise or Datacenter) TPM 2.0 FIPS, CC-TCG certified, TPM 2.0 China NationZ | 
| Integrated Network Card | 1 x 1 Gb | 
| Rear Ports | 1 x USB 3.0 1 iDRAC Ethernet port 1 iDRAC Direct port (Micro-AB USB) 1 x Mini DisplayPort | 
| PCIE Slots | Up to 2 Low-Profile PCIe x16 Gen5 slots 1 x OCP 3.0 x16 Gen5 | 
| Operating Systems and Hypervisors | Canonical Ubuntu LTS Server Microsoft Windows Server with Hyper-V Red Hat Enterprise Linux SUSE Linux Enterprise Server VMware ESXi/vSAN | 
Build and Design
The Dell PowerEdge C6600 chassis and C6615 nodes offer an exceptionally dense computing option for deployment scenarios that need to minimize physical space used in a rack-mount environment. This is suitable for hyperconverged solutions running in a clustered environment, requiring multiple nodes or heavy workloads that do not require consuming 4U or 8U through traditional 1U or 2U server designs. The chassis has a 2U footprint with a depth of 30 inches. The weight of the chassis can increase depending on the final configuration. Dell indicates a maximum weight of a C16 6600-bay configuration with all drives installed at 93.69 pounds.
The front of the system is fairly basic compared to other PowerEdge platforms, without much Dell branding. This type of server does not offer the standard PowerEdge bezel but places the drives and fan intakes at the forefront. The front of the E3.S C6600 version features eight Gen5 NVMe SSDs in the middle, flanked by cooling fan intakes.
The side ears of the chassis contain dedicated power buttons for each node and information buttons indicating the status or issues of that node.
Each C6615 node has a condensed port layout at the rear of the chassis compared to a traditional 1U or 2U server. Ports include USB, iDRAC, a display connector, and a USB service port.
For networking, an OCP slot is available for various interface options (ours has a four-port 25GbE network card), and two PCIe slots are also available. The OCP and dual PCIe slots provide a Gen5 interface.
Opening the PowerEdge C6600 chassis gives you visibility into the layout of how cooling, power distribution, and disk I/O paths are managed. The PCIe/SAS cabling from the disk backplane is routed directly to each node via quick-connect fittings that also transmit data and power.
Depending on the internal configuration of each node, the drive connections connect directly to the motherboard or to a PERC card for hardware RAID options.
Apart from cooling and power, the nodes do not share any other resources.
Dell PowerEdge C6615 Performance
Tested Node Specifications
Our four C6615 nodes have identical configurations. We will compare them and show the average performance across the nodes.
- 1 AMD EPYC 8534P processor 64 cores
- 6 x 96 GB DDR5 4800 576 MB/s (XNUMX GB)
- Windows Server Standard 2022
- Dell RAID1 BOSS boot SSD
- 2 PCIe Gen5 E3.S SSDs
During our performance tests, the nodes ran in parallel to give an overall score taking into account shared power and cooling resources.
Storage Performance
Each of the four Dell Power Edge C6615 nodes includes a RAID1 BOSS SSD for booting and two E3.S bays for Gen5 enterprise SSDs. Although the BOSS card is no slouch, it offers a very different performance profile than the E3.S SSDs.
Although much of this review focuses on overall system-level performance, we briefly covered both types of storage on this system with corner-case workloads. Our first test focused on the BOSS RAID1 boot SSD group.
| Dell BOSS RAID1 | Read Performance | Write Performance | 
|---|---|---|
| Sequential 1 MB Q32/4T | 2,963 MB/s | 1,067 MB/s | 
| Random 4K Q32/8T | 600,786 IOPS (0.426 ms) | 249,819 IOPS (1.024 ms) | 
Next, we looked at a single Gen5 E3.S SSD, which included the KIOXIA CM7 7.68 TB read-intensive SSD in our test system.
| KIOXIA 7.68 TB CM7-R | Read Performance | Write Performance | 
|---|---|---|
| Sequential 1 MB Q32/4T | 13,736 MB/s | 7,089 MB/s | 
| Random 4K Q32/8T | 931,671 IOPS (0.266 ms) | 768,739 IOPS (0.329 ms) | 
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better. Here are the results for all EPYC chips.
