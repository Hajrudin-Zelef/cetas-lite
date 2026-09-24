---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79
title: "fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "compute", "cost", "intel", "latency"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79.md
source_anchor: ""
source_lines: [1, 62]
sha256: bcb7910b3de7a762b687866e43e0f56595b8085e5cf49145ee4600dda110c979
---

# fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79

<!-- source: https://www.storagereview.com/fr/review/supermicro-superchassis-846be1c-r1k28b-review -->

The SuperMicro SuperChassis 846BE1C-R1K28B is a 24-bay JBOD. The drive trays are hot-swappable and designed for 3.5-inch hard drives (although an adapter can be used to install 2.5-inch hard drives or SSDs). If one were to use 8TB drives, such as HGST's He8 drives, this would bring the total maximum capacity to 192TB. The SuperChassis can also be used as a central unit and supports single and dual Intel and AMD processors, and supports a variety of motherboards.
The concept of a storage shelf, or JBOD, is one of the most basic in storage architecture. The chassis essentially houses the disks, connecting to a host machine via a SAS cable and an HBA in the host. This type of arrangement continues to be popular when enterprise users want to keep storage local on the host, but may have outgrown the available drive bays, or have other unique requirements and don't need a full SAN with its own storage controllers. In fact, in the future, we will show these same hard drive configurations paired with caching solutions, to show how flash and software can benefit large hard drive arrays in an enterprise environment. JBOD use cases continue to expand with new technologies and powerful host-side compute power.
The SuperMicro SuperChassis 846BE1C-R1K28B is designed with budget and cost-effectiveness in mind. This DAS is designed to be easy to use and maintain. The drives are hot-swappable coupled with redundant power supplies; this makes the SuperChassis unlikely to cause serious downtime issues. The JBOD is intended for businesses that need an inexpensive method to add storage to their existing system.
The SuperMicro SuperChassis 846BE1C-R1K28B has a list price of $1,400 (unpopulated) and comes with a warranty that covers 3 years labor, 1 year parts, and 3 months advanced replacement.
SuperMicro SuperChassis 846BE1C-R1K28B Specifications:
- Form factor: 4U
- Processor support: single and dual Intel and AMD processors
- Storage capacity:
  - 24 hot-swappable 3.5-inch SAS or SATA hard drives
  - Optional 2.5" adapter trays can be added
- Expansion slots: 7 x full height and full length
- Expander: LSI SAS2 expanders
- Environment:
  - Operating temperature: 5 ° C ~ 35 ° C (41 ° F ~ 95 ° F)
  - Non-operating temperature: -40 °C ~ 60 °C (-40 °F ~ 140 °F)
  - Operating relative humidity: 8% ~ 90% (non-condensing)
  - Non-operating relative humidity: 5% to 95% (non-condensing)
- Cooling
  - 3 hot-swappable 80mm PVM cooling fans
  - 2 hot-swappable 80mm rear exhaust PWM fans
- Power supply: 1280W high-efficiency redundant digital power supplies with PMBus 1.2
- AC input:
  - 1000W output @ 100-140V, 8-12A, 50-60Hz
  - 1280W output @ 180-240V, 6-8A, 50-60Hz
- DC input:
  - 1000W: +12V/83A; +5Vsb/4A
  - 1280W: +12V/106.7A, +5Vsb/4A
- Dimensions (DxWxH): 26.5"x17.2"x7"
- Weight (without drives): 75 lbs.
Design and construction
The SuperMicro SuperChassis 846BE1C-R1K28B 24-bay JBOD is a 4U direct-attached storage. At the front of the platform are the hot-swappable drive bays, four rows of six bays. To power off a drive, simply press the brown button, a small handle appears, and the drive simply needs to be removed and another replaced. At the bottom left are the LEDs. The bottom right side bears the SuperMicro branding. At the top of each side are metal ear handles for racking and unracking the unit.
Going around to the back of the unit, there are two hot-swappable redundant power supplies on the left side. The center is dominated by two cooling fans. On the right side are 4 mini-SAS HD ports and one RJ45 port.
Testing background and comparables
To test the SuperMicro SuperChassis 846BE1C-R1K28B 24-bay JBOD, we used to test 24 HGST Ultrastar Helium He8 8TB drives in a mirror setup. We also performed the same tests with 20 HGST Ultrastar Helium He8 8TB and 4 HGST Ultrastar SSD800MR SAS3 500GB for SSD tiering.
Drives used in the JBOD for testing:
Management
Management isn't the first thing that comes to mind when thinking about JBODs. However, for our SQL server test, we configured a large mirrored storage pool of HGST he8 8TB helium drives as well as a mirrored pool with SSD cache using HGST Ultrastar SSD800MM SAS3 400GB.
On the main Storage Spaces screen, volumes and storage pools need to be configured.
Users need to select the physical disks to use for the storage pool, in this case the He8 drives.
Once done, users need to name the virtual disk.
And select the size of the virtual disk beyond the existing capacity.
Once the location, name, and size have been decided, users need to confirm the settings.
And confirm the tiering settings, the Mirror layout (RAID10) in this case.
Application performance analysis
StorageReview's Microsoft SQL Server OLTP testing protocol uses the current preliminary version of the TPC-C benchmark (Transaction Processing Performance Council's Benchmark C), an online transaction processing benchmark that simulates the activities encountered in complex application environments. The TPC-C benchmark is more representative than synthetic performance benchmarks of real-world performance, allowing accurate assessment of the strengths and bottlenecks of the storage infrastructure in database environments. Our SQL Server protocol for this analysis uses a 685GB SQL Server database (scale 3,000) and measures transactional performance and latency under a load of 30,000 virtual users.
Using Storage Spaces configured Mirror RAID, we saw the JBOD drop 3,874.65 TPS and looking at SSD tiering, the number nearly doubled with 6,307.92 TPS
The average latency shows us an even more fantastic difference. When it was populated only with He8 helium drives, the latency was 2,990 ms. However, when we add the SSD cache, the latency dropped to 6 ms, almost five hundred times faster.
Conclusion
The SuperMicro SuperChassis 846BE1C-R1K28B 24-bay JBOD is a 4U direct-attached storage unit offering cost-effectiveness and flexibility. With 24 drive bays, able to accommodate 3.5" drives or with a 2.5" adapter and supporting SAS and SATA, the JBOD can have a maximum capacity of 192TB with 8TB hard drives. Like the vast majority of what SuperMicro offers, the 846BE1C JBOD offers great flexibility regarding the processors and motherboards it supports if users choose to use the 846BE1C as a central unit.
To test the 846BE1C JBOD instead of comparing it to another JBOD, we chose to compare running all 24 bays with the highest-capacity drive currently available, the HGST Ultrastar Helium He8 8TB, against the same configuration with 4 of the drives replaced by SSDs, HGST Ultrastar SSD800MR SAS3 500GB, for tiering. We also used Storage Spaces to perform the initial configurations. While SSD tiering provided better results, which is not a surprise, what was surprising was the difference in the results. In the SQL server test, we saw almost double the TPS of the non-tiered test, 3,874.65 TPS, compared to the tiered, 6,307.92 TPS. The biggest difference was in latency, with the non-tiered running a very high 2,990 ms and the tiered running almost five hundred times faster at 6 ms.
Pros
- Affordable and flexible design
- Up to 192TB of capacity with 8TB drives
- Excellent performance with SAS3 interconnects
Cons
- Additional 2.5-inch bays at the rear would improve space efficiency
Conclusion
The SuperMicro SuperChassis 846BE1C-R1K28B 24-bay is a 4U SAS3-compatible JBOD offered with a lot of flexibility and capacity at an affordable price.
SuperMicro SuperChassis 846BE1C-R1K28B product page
Discuss this review
