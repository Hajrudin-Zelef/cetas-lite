---
id: collect-240926-storagereview/storagereview/fr-review-dell-powerstore-gen-3-ac26ebce-4
title: "fr-review-dell-powerstore-gen-3-ac26ebce"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["dram", "intel", "latency", "memory", "nand", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-powerstore-gen-3-ac26ebce.md
source_anchor: ""
source_lines: [66, 80]
sha256: 765ef8b9ef17c718208441f25f43a598f094f7adf4d47c5def08f1fe07b33613
---

# fr-review-dell-powerstore-gen-3-ac26ebce

PowerStore Gen 2 uses front-end U.2 NVRAM drives for write cache persistence, reducing storage capacity by four slots. Gen 3 introduces software-defined persistent memory (SDPM): standard DDR5 DRAM is presented to the operating system as an ACPI-compatible NVDIMM-N. In the event of a power loss, a BIOS SMI handler transfers the volatile state to the M.2 SSD, powered by a backup battery. Each PowerStore Gen 3 bay provides substantial power through its lithium-ion batteries. On the 5500 and 9500 models, each controller is equipped with two 54 Wh batteries (216 Wh total per chassis), while the 1500 model has one per controller (108 Wh total per chassis).
The image below shows the PowerStore 9500 controller with its M.2 boot drive and two battery blocks that keep the system powered long enough to save state to disk in the event of a power loss.
On the PowerStore 1500, there is a single-connector, single-battery configuration. The SDPM architecture is the same, simply adapted to the smaller controller size.
Power and Management
Power and management have been migrated to standard Dell server components. The BMC controller, previously based on the legacy EMC GEM controller, is replaced by a storage-optimized iDRAC, aligning PowerStore with the rest of the Dell server lineup in terms of ease of maintenance. Power supplies and backup batteries are now standard PowerEdge components rather than EMC-specific hardware. This should translate into improved technical support, reduced downtime, and increased availability. The backup power subsystem supports the processors, DIMMs, drives, fans, and the iDRAC itself, providing enough power to allow volatile data to be transferred to disk in the event of a power loss.
PowerStore Gen 3 Performance
Dell has released preliminary figures comparing 3rd generation to 2nd generation. These figures, while indicative, align with expected performance given the adoption of the latest Intel x86 architectures, DDR5 memory, PCIe Gen 5, and 200 GbE RDMA interconnect between controllers. Compared to the previous generation, Dell is targeting up to three times the IOPS on mixed 8K workloads, up to three times the throughput on 1 MB sequential reads, and up to three times the throughput on 1 MB sequential writes, with significant latency gains in both reads and writes. For an independent comparison, Principled Technologies tested the 9500 against a 3% NVMe competitor (not named in the report) using Vdbench. On an enterprise OLTP workload with data analytics, the 9500 reached 3 3 IOPS, versus 834,558 357,427 for the competitor, a 2.33x advantage.
Latency shows the same profile. With a target of 310,000 IOPS for a mixed small-block database workload, the 9500 showed 0.44 ms latency, while its competitor reached 1.22 ms, a reduction of roughly 64%.
Data processing efficiency completes the comparison and proves essential in the face of rising NAND memory prices. On a dataset optimized for 2:1 compression and 2.5:1 deduplication with an 8 KB deduplication unit, the 9500 achieved an overall reduction of 6.6:1. Its competitor, on the same data, achieved a rate of 2.76:1.
Conclusion
PowerStore Gen 3 is the most significant update since the platform launched in 2020. This generational redesign will force the entire enterprise storage array market to adapt. Dell has completely rethought the chassis and modernized the drive form factor, cache architecture, node interconnect, I/O plane, and management stack in a single major step. Each of these choices is forward-looking, giving the ten-year lifecycle extension genuine credibility rather than wishful thinking.
The benefits of this design are easy to list. Forty E3.S bays in a 3U form factor with no dedicated cache slots. A 200 GbE RDMA backplane compatible with all processors and ready for future chip generations. SDPM memory replaces NVRAM drives, freeing front bays for data and eliminating dependence on a single persistent memory technology. iDRAC, PowerEdge power supplies, and Dell backup batteries replace traditional EMC components, allowing PowerStore to benefit from the same ease of maintenance and supply chain as the rest of Dell's enterprise lineup. Finally, the unified-by-default approach, combined with support for both TLC and QLC media in a single model without performance loss, greatly simplifies the purchasing decision for platforms that must handle diverse workloads within a single architecture.
We look forward to spending more time with the platform in our lab and observing how the 3rd-generation hardware integrates seamlessly with all the software features offered by PowerStoreOS 5.0. Autonomous data flow management, structured metadata for high-capacity QLC flash, I/O-level telemetry, and dynamic sharing of block and file resources are all improvements that, on a chassis offering such headroom, compound over time. It is this synergy between hardware and software that justifies the "Elite" designation for this release.
PowerStore Product Page
This report is sponsored by Dell Technologies. All views and opinions expressed in this report are based on our impartial assessment of the product(s) under review.
