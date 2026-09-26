---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770ap-review-84b71047-2
title: "fr-review-dell-poweredge-r770ap-review-84b71047"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "compute", "datacenter", "ethernet", "gpu", "inference", "intel", "latency", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770ap-review-84b71047.md
source_anchor: ""
source_lines: [3, 55]
sha256: 06352f092ef13e8f384d996839ac5850c44b733053a08a59eb90dee8f2e88804
---

# fr-review-dell-poweredge-r770ap-review-84b71047

The Dell PowerEdge R770AP is not a general-purpose server, and that is precisely its purpose. While most dual-processor 2U platforms prioritize flexibility, the R770AP skips all of it, sacrificing GPU support, mixed storage options, and raw memory capacity in favor of maximum core density, memory bandwidth, and execution stability available in Dell's current Intel lineup. It is a server built around a specific processor architecture for a particular class of workloads, and it fully embraces its limitations.
To understand its reason for being, one must start with the platform it is built on. Dell's PowerEdge R7x0 lineup has always been the brand's most versatile Intel 2U server, with the PowerEdge R7725, equipped with an AMD processor, playing an equivalent role on the EPYC side. The PowerEdge R770 continues this Intel tradition by supporting Xeon 6 processors with P-cores and E-cores, GPU accelerators, mixed SAS/SATA/NVMe storage, up to 8 TB of memory across 32 DIMM slots, and enough PCIe Gen5 expansion capacity to cover all needs, from virtualization to AI inference.
The PowerEdge R770AP is not that server.
The "AP" designation stands for "Advanced Performance," but it does not fully convey the differences between these two machines. While the R770 uses the Intel Granite Rapids-SP chip on the LGA 4710 socket with 8 memory channels and up to 86 performance cores (P-cores), the R770AP adopts the Granite Rapids-AP platform on the LGA 7529 socket, offering up to 128 performance cores per socket (120 cores in our test configuration) and 12 DDR5 memory channels. This distinction is at the heart of Intel's overall strategy for its Xeon 6 6900 series: the 6900P processors on the AP platform represent the pinnacle of Intel server chips, designed specifically for workloads where per-core performance, memory bandwidth, and execution stability take precedence over overall server configuration flexibility.
Intel's extended Xeon 6 architecture divides the data center into two categories. E-core processors prioritize density and power efficiency for cloud-native and scalable workloads, such as microservices and content delivery. P-core processors target compute-intensive tasks where consistent per-thread performance is essential: HPC simulations, real-time analytics, large in-memory databases, and latency-sensitive financial calculations. The 6900P series sits at the top of this P-core processor range, combining the highest number of available cores with 12-channel memory bandwidth, up to 96 PCIe Gen5 lanes per socket, up to 6 UPI 2.0 links, and L3 cache pools reaching 504 MB on high-end models like the Intel Xeon 6978P. The architectural goal is not just raw throughput, but predictable throughput, minimizing scheduling fluctuations and memory access variability that degrade performance in latency-critical environments.
The R770AP is the very embodiment of this philosophy at Dell. It strips away everything superfluous on the Granite Rapids-AP platform: GPU support is entirely removed, SAS and SATA storage options are replaced by exclusively NVMe configurations (up to 16 2.5-inch Gen5 NVMe SSDs or up to 32 E3.S Gen5 NVMe SSDs, depending on configuration), memory capacity is capped at 3 TB across 24 DIMM slots (12 per socket, 1 DPC for maximum per-channel speed), and PCIe expansion is reduced to five Gen5 x16 slots and two OCP 3.0 network cards. The result is a dual-socket 2U platform optimized for compute density, memory bandwidth, and the deterministic behavior demanded by workloads such as high-frequency trading, real-time risk analytics, and massively parallel simulation.
Our test model pairs two Intel Xeon 6978P processors, each with 120 performance cores clocked at 2.1 GHz (base frequency) and 3.2 GHz (all-core turbo mode), and carries 3 TB of DDR5-6400 memory across all 24 DIMM slots. Compared to the R770, equipped with two Xeon 6787P processors (86 cores each, 8 memory channels, and 2 TB of DDR5), the R770AP offers 39.5% more cores and 50% more memory channels. The question remains whether these architectural advantages translate into concrete performance gains and whether the platform's compromises are justified for the workloads targeted by Dell and Intel.
Dell PowerEdge R770AP Specifications
The table below presents the physical configuration and support specifications for the Dell PowerEdge R770AP platform.
| Specifications | Dell PowerEdge R770AP | 
|---|---|
| Processor |  | 
| Processor | Two Intel® Xeon® 6 6900 series processors with P-cores, up to 128 cores each. | 
| Memory |  | 
| DIMM Slots | 24 DDR5 DIMM slots | 
| Maximum Memory | 3 TB | 
| Memory Speed | Up to 6400 MT/s | 
| Memory Type | Registered ECC DDR5 RDIMM modules only | 
| Storage |  | 
| Storage Controllers (RAID) | Front (internal) PERC H975i DC-MHS | 
| Internal Boot | BOSS-N1 DC-MHS: HWRAID 1, 2 M.2 NVMe SSDs or USB | 
| Front Drive Bays | Up to 16 2.5-inch Gen5 x4 NVMe SSDs (maximum capacity of 245.76 TB) Up to 16 2.5-inch Gen5 x2 NVMe SSDs (maximum capacity of 245.76 TB) Up to 32 EDSFF E3.S Gen5 NVMe SSDs (maximum capacity of 491.52 TB) | 
| Rear Drive Bays | N/A | 
| Engine Tuning |  | 
| Power Supplies | 1500 W Titanium, 100-120 LLAC or 200-240 HLAC, 240 V DC, hot-swap redundancy 1800 W Titanium, 200-240 HLAC, 240 VDC, hot-swap redundancy 2400 W Titanium, 100-120 LLAC or 200-240 HLAC, 240 V DC, hot-swap redundancy 3200 W Titanium, 200-220 HLAC or 220.1-240 HLAC, 240 VDC, hot-swap redundancy 3200 W Titanium, 277 VAC and HVDC, hot-swap redundancy* | 
| Cooling and Fans |  | 
| Cooling Options | air cooling | 
| Fans | Up to 6 hot-swap fans | 
| Form Factor and Dimensions |  | 
| Form Factor | 2U rack server | 
| Height | 86.8 mm (3.42 inches) | 
| Width | 482 mm (19.0 inches) | 
| Depth (with bezel) | 802.40 mm (31.59 inches) | 
| Depth (without bezel) | 801.51 mm (31.56 inches) | 
| Bezel | Optional metal bezel | 
| Networking and Expansion |  | 
| OCP Network Options | Up to two OCP NIC 3.0 network cards Slot 4: 1×8 or 1×16 Gen5 OCP 3.0 Slot 10: 1×16 Gen5 OCP 3.0 | 
| Integrated Network Card | Dedicated 1 Gb BMC Ethernet port | 
| PCIe Slots | Up to 5 PCIe Gen5 slots (x16 connectors) Slot 2: 1×16 Gen5, full height, half length Slot 3: 1×16 Gen5, full height/low profile, half length Slot 5: 1×16 Gen5, full height, half length Slot 7: 1×16 Gen5, full height, half length Slot 9: 1×16 Gen5, full height/low profile, half length | 
| GPU Options | N/A | 
| Ports |  | 
| Front Ports | 1x USB 2.0 Type-C | 
| Rear Ports | 1 dedicated BMC Ethernet port 2x USB 3.1 Type-A 1x VGA | 
| Internal Ports | 1x USB 3.1 Type-A | 
| Management |  | 
| Integrated Management | iDRAC10, iDRAC Direct, iDRAC RESTful API with Redfish, RACADM CLI, iDRAC Service Module | 
| Security |  | 
| Security Features | Cryptographically signed firmware, data-at-rest encryption (SED with local or external key management), secure boot, secure component verification (hardware integrity check), secure erase, hardware root of trust, system lock (requires iDRAC10 Enterprise or Datacenter), FIPS/CC-TCG certified TPM 2.0, chassis intrusion detection | 
| Operating Systems and Hypervisors |  | 
| Supported Operating Systems/Hypervisors | Canonical Ubuntu Server LTS, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware vSAN / VMware ESXi*, Microsoft Windows, Microsoft Windows Server, Microsoft Windows Server Datacenter | 
Design and Build
The Dell PowerEdge 770AP is a 2U rack server from Dell's 17th generation of PowerEdge servers, sharing the same design as the R770 we tested. It measures 3.42 cm in height, 19.0 cm in width, and 31.59 cm in depth. The front bezel is optional. The front panel includes iDRAC direct access, a USB 2.0 Type-C port, a power button, and a system identification button.
Storage
