---
id: collect-240926-storagereview/storagereview/fr-review-dell-powerstore-gen-3-ac26ebce-2
title: "fr-review-dell-powerstore-gen-3-ac26ebce"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["compute", "dram", "intel", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-powerstore-gen-3-ac26ebce.md
source_anchor: ""
source_lines: [3, 33]
sha256: c15500189d9c790af4f6293c388854cc593a7fd0f56142f784effba3ae44a64b
---

# fr-review-dell-powerstore-gen-3-ac26ebce

Storage updates typically come in two forms. There is the discreet upgrade, where a vendor drops in a new processor, announces a performance gain of a few percentage points, and ships the same chassis with a different label. And then there is the generational redesign, where the chassis, drives, interconnect, cache architecture, and management plane are all replaced at once. Dell PowerStore Gen 3, initially marketed as PowerStore Elite, belongs firmly in the second category, and by a wide margin. Every major subsystem of the platform has been changed, and most of them in ways that fundamentally redefine what a unified array should look like in 2026.
We have been following PowerStore since its inception. In our view, this is the most significant update since its initial launch in 2020, and arguably the most ambitious storage platform redesign from a major vendor in years. Dell did not simply improve Gen 2. The company completely rebuilt the platform, from the chassis to the enclosure, betting on form factors and architectural choices that most of the industry has not yet adopted, and designed the whole thing for a ten-year lifespan with multiple controller upgrades. To see these new models, Dell invited us to Hopkinton, Massachusetts.
Storage density is a major selling point of the new Dell PowerStore, and Dell does not disappoint on this front. The platform supports up to 40 E3.S NVMe drives in a 3U chassis, with planned compatibility for E3.L drives, while retaining user-addressable slots for data instead of reserving slots for cache SSDs. Dell has also thoroughly modernized the underlying hardware platform, adopting next-generation Intel processors, DDR5 memory, end-to-end PCIe Gen5 connectivity, and OCP 3.0 modules that replace the previous Dell-specific SLIC card design.
Controller-to-controller connectivity now reaches 200 GbE RDMA, with the potential to reach even higher throughput through a future I/O card upgrade. On the software side, PowerStoreOS 5.0 introduces new autonomous intelligence for data path management and log-structured metadata to optimize performance and endurance of high-capacity QLC flash. I/O-level telemetry lays the groundwork for built-in ransomware detection, and dynamic resource sharing between block and file services is also built in. All PowerStore appliances are now unified from the start, offering improved scale-out support for file and block, as well as non-disruptive data mobility between clusters.
Dell has also added unaligned deduplication and improved compression offloads, a key factor in the company's decision to increase its data reduction guarantee from 5:1 in previous generations to 6:1 with the new platform.
Hardware changes and software updates alone do not explain the full scope of the changes. Dell's underlying architectural choices are essential, as they determine whether the platform will remain competitive over the next decade or whether it will seem obsolete within a few years. The move to E3.S/L drives, the larger chassis, the adoption of software-defined persistent memory, the wireless backplane interconnect, and the decoupling of the inter-node architecture from the processor generation are all forward-looking decisions. Together, they make the 3rd-generation chassis a solid foundation for Dell's proposed multi-generational upgrade strategy with lifecycle extension.
The PowerStore lineup gains three new models. The PowerStore 1500 is a single-processor platform with 24 drive bays at launch and a 100 GbE RDMA inter-node interface. The 5500 and 9500 models, in a 3U form factor, are dual-processor and offer 40 drive bays as well as 200 GbE RDMA inter-node connectivity. The 9500 offers twice the memory and a higher core count than the 5500. A future upgrade via controller replacement will allow the 1500 to scale up to 40 drives and 200 GbE RDMA connectivity. All three models run PowerStoreOS 5.0, share the same OCP 3.0 I/O architecture, and support both TLC and QLC media, with no performance loss when moving to QLC technology. There is therefore no longer any need to choose between different models to meet various needs or performance requirements.
Second-generation platforms remain available, and existing PowerStore customers have a clear future-proofing path through intelligent clustering. First-, second-, and third-generation appliances can coexist within the same cluster, ensuring non-disruptive workload mobility. This solution is ideal for an installed base that does not want to change platforms every two or three years.
Dell PowerStore Gen3 Specifications
All three third-generation models share a 3U dual-node chassis, the same OCP 3.0 I/O architecture, and a unified PowerStoreOS operating system with native support for block or file storage. They differ in processor socket count, DRAM memory capacity, drive count, and inter-node bandwidth.
| Specifications | Power Store 1500 | Power Store 5500 | Power Store 9500 | 
|---|---|---|---|
| Market |  |  |  | 
| placement | Midrange | midrange/high-end | High-end flagship model | 
| Chassis | 3U, dual node |  |  | 
| Compute and memory |  |  |  | 
| CPU platform | Single-socket Intel | Dual-socket Intel | Dual-socket Intel | 
| Processor per appliance | 2 x 24 GHz cores | 4 x 24 GHz cores | 4 x 32 GHz cores | 
| Memory per appliance | 512 GB (16 × 32 GB) | 1,024 GB (32 × 32 GB) | 2,048 GB (64 × 32 GB) | 
| Storage |  |  |  | 
| Drives per base appliance | Up to 24 EDSFF | Up to 40 EDSFF | Up to 40 EDSFF | 
| Drives per expansion (after RTS) | 44EDSFF |  |  | 
| Roadside assistance | TLC: 3.84 / 7.68 / 15.36 TB · QLC: 30.72 TB |  |  | 
| Minimum hard drive configuration | TLC 6 × 3.84 TB · QLC 7 × 30.72 TB | TLC 6 × 3.84 TB · QLC 11 × 30.72 TB | TLC 6 × 3.84 TB · QLC 11 × 30.72 TB | 
| Maximum raw capacity (base) | ~737 TB (24 × 30.72 TB) | ~1.2 PB (40 × 30.72 TB) | ~1.2 PB (40 × 30.72 TB) | 
| I/O and network |  |  |  | 
| OCP 3.0 line cards per node | 3 (+1 reserved) | 5 (+1 reserved) | 5 (+1 reserved) | 
| Inter-node interconnect | 100 GbE RDMA | 200 GbE RDMA | 200 GbE RDMA | 
Dell PowerStore Gen 2 vs. Gen 3
Every major hardware subsystem of the Gen 3 series PowerStore appliances has received at least a one-generation improvement. The most significant changes for capacity planning are the 3U chassis, the E3.S drive bays, and the move from 2 × 10 GbE to up to 200 GbE on the node interconnect.
The base 3U chassis of the 5500 and 9500 models has 40 drive slots. The 1500 model ships with 24 populated bays and 16 additional bays accessible through a future Data-In-Place upgrade. That works out to 13.3 drives per rack unit (RU), compared to 10.5 for generation 2, or roughly 40% more drives per RU in the base appliance and an 83% increase in density with the expansion to 44 drives after the RTS release. The drives used are standard 1 TB E3.S NVMe SSDs from various vendors, with no proprietary carrier, which reduces risks related to supply constraints and dependence on a single vendor. Capacities range from 3.84 TB, 7.68 TB, and 15.36 TB in TLC, plus 30.72 TB in QLC, and the same set of options is supported across all three models (1500, 5500, and 9500) without compromising performance between media types.
