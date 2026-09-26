---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r7615-review-536a6ae7-3
title: "fr-review-dell-poweredge-r7615-review-536a6ae7"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "compute", "consumer", "dram", "ethernet", "gpu", "gpus", "intel", "liquid cooling", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r7615-review-536a6ae7.md
source_anchor: ""
source_lines: [40, 89]
sha256: 516776b74efaf5abc08077df0b1beead5745d9244abbcaf1fcad90005005f7bd
---

# fr-review-dell-poweredge-r7615-review-536a6ae7

The storage area has a scope section providing useful information. The summary page offers an overview of the disk environment, displaying the status of each physical disk and summarizing the RAID configuration. Below is the event log, which records all storage-related activities. This is extremely valuable for ongoing maintenance and quick troubleshooting when issues arise.
Overall, with its intuitive design and detailed reporting features, the iDRAC9 interface is a server management engine. It enables businesses to administer their Dell PowerEdge R7615 efficiently and seamlessly and provides system administrators with everything they need to maintain optimal server performance and reliability.
Dell PowerEdge R7615 Review Version
The Dell PowerEdge R7615 we received for review is essentially the most basic version of what this server range can offer. It's like the entry-level model of a high-end vehicle, stripped of its premium features while retaining the potential of what it could be.
As such, our review unit is powered by an AMD EPYC 9354P 32-core processor coupled with 16 GB of RAM, the latter certainly lower than that of most consumer desktop computers. The server's storage and connectivity follow this theme of simplicity. Equipped with a PERC 11 RAID card and a SATA/SAS/NVMe backplane, it offers basic but versatile storage options. Dual 1GbE Ethernet provides standard network connectivity, suitable for everyday tasks.
Storage is managed by a single 960 GB read-intensive NVMe SSD, which should provide adequate data access speeds for general purposes. Although it reflects the general theme of our server, which is sufficient for typical operations, it is far from reaching the full potential of the R7615. The PCIe backplane, Riser Config 3, and its various slots hint at the system's scalability. Although it has available slots, we do not have a GPU installed in this review version. This will undoubtedly affect how we test the R7615 in the performance section below.
Essentially, this PowerEdge R7615 configuration is a glimpse of the possibilities of the Dell server. It's a starting point, a basic, no-frills server that gets the job done but with much more room for growth and expansion.
Dell PowerEdge R7615 Specifications
| Feature | Specifications |
| Processor | One 4th Generation AMD EPYC 9004 Series processor with up to 128 cores |
| Memory | 12 DDR5 DIMM slots, supports up to 3 TB RDIMM, speeds up to 4800 MT/s |
| Storage Controllers | PERC H965i, PERC H755, PERC H755N, PERC H355, HBA355i |
| Drive Bays | Front bays: up to 32 bays with different configurations, max 368.64 TB; Rear bays: up to 4 bays, max 61.44 TB |
| Power Supplies | 2400 1800 W Platinum, 1400 1100 W Titanium, 1100 W Platinum, 1100 W Titanium, 1100 W LVDC options |
| Cooling Options | Air cooling, optional Direct Liquid Cooling (DLC) |
| Fans | High-performance Silver/Gold fans, up to 6 hot-plug fans |
| Dimensions | Height: 86.8 mm, Width: 482 mm, Depth: 772.13 mm (with bezel) |
| Form Factor | 2U rack server |
| Integrated Management | iDRAC9, iDRAC Direct, iDRAC RESTful API with Redfish, iDRAC Service Module, Quick Sync 2 wireless module |
| Bezel | Optional LCD bezel or security bezel |
| OpenManage Software | CloudIQ plugin for PowerEdge, OpenManage Enterprise, various OpenManage integrations |
| Security | AMD Secure Memory Encryption, Secure Boot, TPM 2.0 and other features |
| Integrated Network Card | 2 x 1 GbE LOM card (optional), 1 x OCP 3.0 card (optional) |
| Network Options | Various configurations, including LOM and OCP cards |
| GPU Options | Up to 3 x 300 W DW or 6 x 75 W SW |
| Ports | iDRAC Direct Micro-AB USB, USB 2.0, VGA, serial (optional), USB 3.0 (optional) |
| PCIe Slots | Up to eight PCIe slots, including various configurations |
| Operating Systems and Hypervisors | Support for Canonical Ubuntu Server LTS, Microsoft Windows Server with Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi |
Dell PowerEdge R7615 Design and Build
The overall design of the PowerEdge R7615 emphasizes accessibility and ease of maintenance. Whether upgrading RAM, replacing a processor, or adding expansion cards, the server's layout and modular design simplify these tasks. Users have a wide choice of components to upgrade as their business needs grow.
From the front panel, the chassis can accommodate various drive configurations. In this configuration, we offer options for drives up to 16 × 2.5″, providing flexibility for storage needs. Whether you need high-capacity hard drives for mass storage or fast SSDs for quicker access, this chassis can support it.
Under the hood, we see an intelligent and spacious design. At the front are the drive bays and the backplane next to a bank of cooling fans. These are crucial for maintaining airflow over heat-generating components, and this configuration certainly seems to dissipate heat and keep the server running within safe temperature thresholds.
In the middle of the motherboard is the processor, which is the AMD EPYC 9354 processor. On each side of the processor/heatsink are 12 DDR5 DIMM slots (which support up to 3 TB RDIMM and speeds up to 4800 MT/s). In our case, we only have 16 GB of RAM.
Moving to the expansion slots, you will see 2 full-height x8 slots, 2 low-profile x16 slots, and 2 double-width x16 slots, all supporting PCIe Gen5. This variety allows for a wide range of expansion cards, giving you the flexibility to add anything from additional network cards to high-end storage controllers.
Regarding power supplies, the PowerEdge R7615 offers a range of choices. Depending on the power needs of your specific version, you can opt for single or dual, hot-plug, fully redundant power supplies, with capacities up to 2400 W. This versatility ensures that the server can be customized to meet specific power requirements, which is crucial for high-performance configurations.
Dell PowerEdge R7615 Performance
As we indicated above, the configuration of our Dell PowerEdge R7615 server is as basic as possible. This means we don't expect much performance in our benchmarking. The R7615 system in our review includes an AMD EPYC 9354 processor (3.25 GHz, 32 cores/64 threads, 256 MB cache) and 16 GB of DDR5 RAM.
Although the server is designed to handle two GPUs, our system has none. It is also limited on DRAM, so we adapted our benchmarking approach to focus on tests that heavily leverage CPU capabilities, meaning we will not perform benchmarks such as ESRI, Luxmark, SPECworkstation 3, 7zip, and SPECviewperf 2020. Instead, our evaluation will focus on benchmarks that primarily emphasize the CPU to clearly show its performance in compute-intensive tasks.
Blender OptiX
The first is the Blender test, an open-source 3D modeling application. This benchmark was run using the Blender Benchmark utility. The score is expressed in samples per minute, with the higher being the better.
|  | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cores, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor (32 cores, 3.7 GHz max)) |
| Blender OptiX version 3.6 (CPU) (Samples per minute; higher is better) | Score | Score |
| Monster | 397 | 279 |
| Junkshop | 245 | 177 |
| Classroom | 200 | 135 |
In our benchmarking tests using Blender OptiX version 3.6, the basic configuration of the Dell PowerEdge R7615, equipped with an AMD EPYC 9354P processor and 16 GB of DDR5 RAM, produced modest results. The server scored 397 samples per minute on the "Monster" test, 245 samples per minute on "Junkshop," and 200 samples per minute on "Classroom."
| Blender OptiX version 4.0 (CPU) (Samples per minute; higher is better) | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cores, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Gen Xeon-G 6430 Processor (32 cores, 3.7 GHz max)) |
| Monster | 375 | 952 |
| Junkshop | 248 | 613 |
| Classroom | 198 | 437 |
