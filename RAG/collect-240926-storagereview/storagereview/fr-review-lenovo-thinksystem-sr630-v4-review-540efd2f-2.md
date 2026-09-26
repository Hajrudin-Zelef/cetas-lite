---
id: collect-240926-storagereview/storagereview/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f-2
title: "fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["compute", "cost", "energy", "gpu", "gpus", "intel", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f.md
source_anchor: ""
source_lines: [3, 55]
sha256: e2998102d91d559588883b05216f440efbd54914ea9b508341aae24025b10f46
---

# fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f

The Lenovo ThinkSystem SR630 V4 is a flexible and powerful 1-socket 2U rack server, designed to meet the needs of sectors such as cloud services and telecommunications. Whether optimizing scalable workloads or future-proofing your data center, the SR630 V4 offers significant upgrades over its predecessor, the SR630 V3. In this review, we examined what's new and how Lenovo has refined this enterprise server to meet the challenges of modern IT environments.
Differences Between Lenovo SR630 V4 and SR630 V3
Processors
One of the major improvements in the SR630 V4 lies in its processing capabilities. While the SR630 V3 relied on 4th and 5th generation Intel Xeon Scalable processors, featuring up to 64 cores and Hyper-Threading technology, the SR630 V4 integrates Intel Xeon 6700 series processors, with up to 144 high-efficiency cores (E-cores). This evolution doubles the number of cores while prioritizing efficiency, despite the absence of Hyper-Threading technology. Additionally, the V4 version plans support for Intel Xeon P-cores, which could significantly improve performance for certain workloads. The higher core density of the V4 version allows businesses to consolidate more applications on the same number of servers, thereby reducing operating costs and physical server requirements.
Here is an overview of all 6300 series processors supported by the SR4 V6700:
| CPU Model | Cores / Threads | Core Speed (Base/Max Turbo) | L3 Cache | Memory Chan | Max Memory Speed | UPI 2.0 Links and Speed | PCIe Lanes | TDP | 
| 6710E | 64/64 | 2.4 / 3.2 GHz | 94 MB | 8 | 5600 MHz | 4 / 16 GT/s | 88 | 205W | 
| 6731E | 96/96 | 2.2 / 3.1 GHz | 96 MB | 8 | 5600 MHz | None‡ | 88 | 250W | 
| 6740E | 96/96 | 2.4 / 3.2 GHz | 96 MB | 8 | 6400 MHz | 4 / 20 GT/s | 88 | 250W | 
| 6746E | 112/112 | 2.2 / 2.7 GHz | 96 MB | 8 | 5600 MHz | 4 / 16 GT/s | 88 | 250W | 
| 6756E | 128/128 | 1.8 / 2.6 GHz | 96 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 225W | 
| 6766E | 144/144 | 1.9 / 2.7 GHz | 108 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 250W | 
| 6780E | 144/144 | 2.2 / 3 GHz | 108 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 330W | 
Memory
Regarding memory, the SR630 V4 improves upon the V3's DDR5 memory, operating up to 5600 MHz by supporting DDR5 memory speeds up to 6400 MHz for E-cores. Both systems have 32 DIMMs (16 per processor) and two DIMMs per channel across eight channels per CPU. Nevertheless, the V4 introduces future-proofing with planned support for advanced memory technologies such as Compute Express Link (CXL) and MCRDIMM for P-cores.
While the SR630 V3 could support up to 8 TB of memory, the V4 focuses on a more targeted capacity of 2 TB for E-cores, thereby optimizing costs and performance for specific workloads.
Storage
The storage capabilities of the Lenovo SR630 V4 demonstrate an evolution toward high-performance NVMe drives while reducing reliance on traditional SAS/SATA options. The SR630 V3 offers flexibility with 3.5-inch SAS/SATA drive bays and up to 16 integrated NVMe ports. However, the SR630 V4 supports up to 12 NVMe drives in front and rear configurations, adding future options for E3.S drive formats, which promise superior capacity and density.
The V4 eliminates support for 3.5-inch drives but introduces hot-swap M.2 options for operating system boot, thereby improving performance and serviceability. By focusing on NVMe and eliminating the need for additional adapters, the V4 maximizes I/O bandwidth and system efficiency.
Networking
For networking, the SR630 V4 builds on the V3's single OCP slot by offering two OCP 3.0 slots, both supporting PCIe Gen 5 x16. This upgrade doubles networking flexibility, enabling improved connectivity and throughput with dual-port 200 GbE network adapters or other advanced networking solutions.
The increased PCIe bandwidth (32 GT/s in Gen 5 versus 16 GT/s in Gen 4) means the V4 can better handle the demanding workloads of data centers and the cloud, making it a more versatile option for modern networking needs.
Tuning Engine
Finally, the SR630 V4 offers improved power options, moving from the SR630 V3's Platinum/Titanium AC 750 W to 1800 W options to a broader range of 800 W to 2000 W, including models compliant with the ErP Lot 9 standard for optimal energy efficiency. The V4 also retains support for -48 V DC power compatible with telecommunications operators while introducing new HVDC 1300 W options for specific regional requirements. These improvements make the V4 more adaptable to various power environments, ensuring it meets the energy requirements of more complex and higher-performance configurations.
Overall, the SR630 V4 represents a solid advance on paper to better meet the need for more performance, flexibility, and efficiency in modern IT environments.
Lenovo ThinkSystem SR630 V4 Specifications
| Lenovo ThinkSystem SR630 V4 Specifications |  | 
| Form Factor | 1U rack | 
| Processor | One or two Intel Xeon 6700E series processors (up to 144 cores, 2.4 GHz and 330 W TDP). Support for Intel Xeon 6700P series processors planned for Q1 2025. | 
| Memory | 32 DIMM slots (16 per processor), supports TruDDR5 RDIMMs up to 6400 MHz (1DPC) or 5200 MHz (2DPC). CXL memory is planned for the Intel Xeon 6700P series in Q1 2025. | 
| Maximum Memory | Up to 2 TB using 32 x 64 GB RDIMMs | 
| Drive Bays |  | 
| Maximum Internal Storage | 184.3 TB with 12 x 15.36 TB 2.5-inch NVMe SSDs | 
| Storage Controller | Up to 16 integrated NVMe ports with RAID support (Intel VROC). Planned support for 12 Gb SAS/SATA RAID and non-RAID adapters. | 
| Network Interfaces | Two OCP 3.0 SFF slots with PCIe 5.0 host interface (x8 or x16), supporting up to 100 GbE network adapters. | 
| PCI Expansion Slots |  | 
| GPU Support | Planned support for up to 3 single-width GPUs | 
| Ports | Front:  Rear:  | 
| Cooling | Up to 8 hot-swap fans (N+1 redundancy), with an additional fan integrated into each power supply. | 
| Power Supply | Up to two redundant hot-swap AC power supplies (800 W, 1300 W, 2000 W). 80 PLUS Platinum and Titanium certifications. | 
| Video | Integrated graphics card with two video ports (rear VGA and optional Mini DisplayPort), supporting resolutions up to 1920 × 1200 at 60 Hz. | 
| Hot-Swap Parts | Drives, power supplies, and fans | 
| Systems Management |  | 
| Security Features | Chassis intrusion switch, power-on and administrator passwords, TPM 2.0, and optional lockable front security bezel. | 
| Supported Operating Systems | Microsoft Windows Server, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, Ubuntu Server. | 
| Warranty | Three years or one year (depending on model) with optional service upgrades for faster response times and extended coverage. | 
| Dimensions | Width: 440 mm (17.3 in), Height: 43 mm (1.7 in), Depth: 788 mm (31 in). | 
| Weight | Maximum weight: 20.2 kg (44.5 lb) | 
Lenovo ThinkSystem SR630 V4 Design and Build
The Lenovo ThinkSystem SR630 V4 retains the compact 1U form factor that is standard for many enterprise rack servers. Its design focuses on a combination of functionality, accessibility, and flexibility. We appreciated its simple yet effective layout, which maximizes airflow, modularity, and ease of use for IT administrators.
Let's move on to the details.
Front Panel
The front panel supports up to 10 hot-swap 2.5-inch drive bays and offers flexibility for various storage configurations, including SAS, SATA, NVMe, or AnyBay drives. This allows businesses to customize storage based on their workloads, whether they prioritize speed, capacity, or cost-effectiveness.
