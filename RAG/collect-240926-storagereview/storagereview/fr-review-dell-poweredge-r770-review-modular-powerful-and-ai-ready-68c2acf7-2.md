---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-2
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["accelerator", "attention", "compute", "cost", "datacenter", "distribution", "energy", "ethernet", "gpu", "gpus", "inference", "intel"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [3, 36]
sha256: 4eef32e5f8a7c36a2f42b028d13b60cb3e3026d0ee360e9432a336bf6ea0aa38
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

Dell PowerEdge R7x0 series servers have long been a cornerstone of data centers, renowned for their exceptional build quality, thoughtful design, performance, density, and reliability, all in a versatile 2U form factor. These servers have constantly evolved to meet changing requirements. Today, with the introduction of the Dell PowerEdge R770, the series takes a decisive step forward.
The R770 debuts the new Intel Xeon 6 processor family, featuring Xeon 6500 and 6700 P-core and E-core processors. It marks Dell's first full adoption of the OCP Data Center Modular Hardware System (DC MHS) standard in its mainstream server lineup. Together, these two developments promise a significant leap forward in both performance and design.
Meeting the demands of modern data centers
The launch of the R770 comes as data centers face increasing pressure. Workloads are becoming increasingly diverse and demanding. The relentless growth of data reinforces the need for robust analytics and databases. From training complex models to deploying real-time inference, artificial intelligence is no longer a niche application but a core business engine requiring substantial compute power and specialized acceleration.
At the same time, energy efficiency and total cost of ownership optimization are receiving heightened attention. Additionally, the industry is increasingly turning to open standards to foster innovation, improve interoperability, and potentially reduce vendor lock-in. The R770, with its new processor options and adoption of OCP DC MHS, is designed to meet these challenges.
Intel Xeon 6 P-Core processors
The R770 utilizes Intel Xeon 6 series processors, specifically the 6700 and 6500 series, featuring Performance and Efficiency cores based on the Socket E2 platform (LGA4710-2). In this analysis, we focus specifically on the P-series SKUs.
Intel builds these processors on a tile-based design, combining I/O tiles with one or two compute tiles. This enables scalability within the series, with configurations up to 86 P-cores (XCC) with two compute tiles, and up to 48 P-cores (HCC) or 16 P-cores (LCC) with a single compute tile.
Compared to previous-generation Sapphire and Emerald Rapids processors, these processors stand out for the universal availability of built-in accelerators across all Xeon 6 processors. These include Intel QuickAssist Technology for encryption and compression, the Intel Data Streaming Accelerator for data movement, the Intel In-Memory Analytics Accelerator for database and analytics acceleration, and the Intel Dynamic Load Balancer for network processing efficiency.
Memory and I/O bandwidth also see substantial improvements. The Xeon 6700/6500 P-core processors support 8-channel DDR5 memory. They also pave the way for MRDIMM (Multiplexed Rank DIMM) modules, which offer speeds up to 8,800 MT/s. On the I/O side, these processors support PCIe 5.0 and CXL 2.0 standards. In a dual-socket configuration, the platform can deliver up to 88 PCIe lanes per socket (176 lanes total).
Despite the differentiation between P-core and E-core processors, the Xeon 6 family maintains consistency in instruction sets, BIOS, drivers, operating system and application support, as well as RAS features, simplifying integration and management across different deployment types. P-core processors are intended for workloads where per-core performance, AI acceleration, high memory bandwidth, and substantial I/O are paramount; think demanding databases, HPC simulations, advanced analytics, and a wide range of AI applications.
Dell PowerEdge R770 specifications
| Specifications | Dell PowerEdge R770 | 
| Processor | Two Intel Xeon 6 processors with up to 144 E-cores or 86 P-cores per processor | 
| Memory | 32 DDR5 DIMM slots, supports RDIMM 8 TB max, speeds up to 6400 MT/s, supports only registered ECC DDR5 DIMMs | 
| Storage controllers | Internal boot: Boot Optimized Storage Subsystem (BOSS-N1 DC-MHS): HWRAID 1, 2 x M.2 NVMe SSDs or M.2 interposer card (DC-MHS): 2 x M.2 NVMe SSDs or USB, Internal controllers: PERC H965i front, PERC H975i front, PERC H365i front | 
| Front and rear bays |  | 
| Hot-plug power supplies |  | 
| Cooling options | Air cooling and direct liquid cooling (DLC is a rack solution and requires rack manifolds and a cooling distribution unit (CDU) to operate) | 
| Fans | High-performance Silver fans (HPR SLVR)/High-performance Gold fans (HPR GOLD), up to 6 hot-pluggable fans | 
| Dimensions and weight | Height – 86.8 mm (3.42 inches), width – 482 mm (18.97 inches), weight – 28.53 kg (62.89 lbs), depth (for rear I/O configuration) – 802.40 mm (31.59 inches) with bezel, 801.51 mm (31.56 inches) without bezel, depth (for front I/O configuration) – 814.52 mm (32.07 inches) without bezel | 
| Form factor | 2U rack server | 
| Integrated management | iDRAC, iDRAC Direct, iDRAC RESTful API with Redfish, RACADM CLI, iDRAC Service Module (iSM), NativeEdge endpoint, NativeEdge orchestrator | 
| Bezel | Optional security bezel | 
| Security | Cryptographically signed firmware, data-at-rest encryption (SED with local or external key management), Secure Boot, Secure Component Verification (hardware integrity check), silicon root of trust, system lockdown, system lock-down (requires iDRAC10 Enterprise or Datacenter), chassis intrusion detection, TPM 2.0 FIPS, CC-TCG certified | 
| Network options |  | 
| GPU options | Up to 6 x 75 W FHHL or up to 2 x 350 W DWFL | 
| Ports | Front ports: 1 USB 2.0 Type-C port, 1 USB 2.0 Type-A port (optional), 1 Mini-DisplayPort (optional), 1 DB9 serial port (with front I/O configuration), 1 iDRAC dedicated management Ethernet port; Rear ports: 1 iDRAC dedicated management Ethernet port, 1 VGA port, 2 USB 3.1 Type-A ports; Internal ports: 1 USB 3.1 Type-A port | 
| PCIe |  | 
| Operating systems and hypervisors | Canonical Ubuntu LTS Server, Microsoft Windows Server with Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware with vSphere | 
Dell PowerEdge R770: Modularity with OCP DC MHS
The Dell PowerEdge R770 features notable advances and flexibility in its physical design and component architecture, adopting the Open Compute Project's Data Center Modular Hardware System (OCP DC MHS) standard.
Continuing the R7x0 series lineage, the R770 offers numerous configuration options to meet various deployment needs. A significant first for this range is the choice between a traditional rear I/O configuration and a front I/O configuration accessible from the cold aisle, providing greater flexibility to adapt to different data center layouts and maintenance requirements. Storage options are equally versatile, ranging from compute nodes with minimal or no local storage to high-density configurations supporting up to 40 E3.S drives for storage-centric workloads.
To meet the growing need for accelerated computing, particularly for AI and HPC, the R770 offers robust expansion capabilities. Depending on the chassis and riser configuration, the server can accommodate up to six full-height, full-length (FHFL) PCIe Gen 5 x16 cards. Additionally, it supports the installation of two double-width GPUs, making it a high-performance platform for a wide range of tasks. Network flexibility is provided by OCP 3.0 mezzanine slots, supporting x8 or x16 cards depending on the configuration.
