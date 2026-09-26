---
id: collect-240926-tomshardware/tomshardware/amd-ryzen-9-9950x-review-zen-5-at-full-power-2
title: "amd-ryzen-9-9950x-review-zen-5-at-full-power"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: ["2023-04"]
keywords: ["amd", "benchmark", "benchmarks", "chiplet", "compute", "intel", "latency", "memory"]
source: docs/RAG/clean_en/tomshardware/amd-ryzen-9-9950x-review-zen-5-at-full-power.md
source_anchor: ""
source_lines: [73, 137]
sha256: 1eeb56c2198b87ca90d1c7d6940066f76426e3c058c9fddc9e7e3544e32ea19c
---

# amd-ryzen-9-9950x-review-zen-5-at-full-power

AMD improved the 9950X’s memory support from DDR5-5200 to DDR5-5600 and expanded the L1 data cache (L1D) to 48KB. AMD says you can expect generally higher memory overclocking ceilings that could reach up to DDR5-8000, with a particular advantage if you opt for an 800-series motherboard. However, DDR5-6000 remains the price/performance sweet spot for most users. Once again, AMD supports ECC memory with its silicon, but the implementation, validation, and support are up to the motherboard OEMs.


The Ryzen 9 9950X drops into the existing AM5 LGA1718 socket and is backward compatible with all 600-series chipsets. AMD’s OEM partners will soon release a fleet of new 800-series motherboards, ranging from the X870 and X870E motherboards, which are natural homes for this class of processor, to lower-end B850 and B840 motherboards.


The 800-series is built around the same Promontory 21 chipset silicon from ASMedia as the 600-series, so the basic platform capabilities remain the same. However, the X870 and X870E chipsets come with an increase in mandatory feature requirements to add more functionality. For instance, the PCIe 5.0 interface is now standard on the X870 lineup for both storage and graphics, whereas it was previously limited to the E-series boards. All X870 boards will now also have USB4 40 Gbps interfaces courtesy of third-party controllers, like the ASMedia ASM4242 (the controller will consume some of the PCIe lanes from the CPU).


Like Intel’s competing K-series chips, the Ryzen 9 9950X doesn’t come with a bundled cooler. The 9950X is compatible with existing coolers, and AMD recommends a 240-280mm liquid cooler (or equivalent) for the Ryzen 9 9000-series processors.


As we’ve seen in the past, AMD’s chips can operate near the top of the maximum safe temperature range during normal operation to extract the utmost performance, so it isn’t uncommon to see temperatures exceed 90C during heavily threaded workloads (particularly AVX). AMD says temperatures should be in the 70-90C range during heavy work but lists 95C as the maximum safe temp (TjMax). Given the nature of AMD’s Precision Boost 2 algorithms, bulkier coolers can unlock more performance in some workloads, so it’s best to pair this chip with a powerful liquid cooler.

## AMD Core Parking Problems — PPM Provisioning File Driver and Thread Targeting

AMD has roped in its latest software and driver advancements to boost gaming performance with its Ryzen 9 9000-series processors, but it's abundantly clear that problems remain with some facets of this tech.


AMD’s dual-CCD (compute chiplet) Ryzen 7000X3D models introduced an innovation — a new core parking technique that automatically engages during gaming to boost performance. AMD has now implemented that feature with its dual-CCD Ryzen 9 9000-series models as well. AMD says it chose to enable the feature on the Ryzen 9000-series due to notable performance improvements. In contrast, the Ryzen 7000-series didn’t benefit as much, and the feature can result in ‘very large regressions’ in some applications, so AMD will not enable the feature retroactively with older processors.

The core parking tech effectively shuts down one CCD during gaming to boost cache hit rates, reduce cross-CCD traffic, and keep the workload pinned to the fastest CCD, all of which boosts gaming performance. The dual-CCD provisioning approach works exceedingly well for X3D processors because only one CCD has the vast L3 cache enabled by a vertically stacked chiplet. Still, it should also benefit the dual-CCD Ryzen 9000 chips without X3D cache because it keeps latency-sensitive game data close to the execution cores.


The feature requires four components: a new chipset driver, updated BIOS, Windows Game Mode, and the Xbox Game Bar (make sure to update it through the Microsoft store). The Xbox Game Bar contains a KGL (known good list) of games that it detects when active, thus triggering Game Mode (you can also instruct the game bar to recognize unknown games and/or other applications as games). The driver communicates with Windows Game Mode to trigger the AMD PPM Provisioning File Driver (installed with the chipset drivers) to park the cores on a single CCD, thus constraining latency-sensitive workloads (like games) to the higher-performance chiplet.


AMD didn’t tell reviewers this feature was active in the new chipset driver until late in the review process, which was problematic. As we’ve covered in the past, the core parking feature has a major problem: It can’t be uninstalled from the operating system. As such, if you later install another processor but use the same operating system, the feature will persist and can continue to park cores (potentially unbeknownst to the user), thus hamstringing performance with processors that aren’t designed to use the feature. We remarked back in April 2023 that it was ‘almost unbelievable’ that this known issue exists, and it is even **more** unbelievable that it still exists 16 months later, in 2024.


If you swap from a dual-CCD chip to a regular processor, you must completely reinstall Windows. Additionally, we've heard reports that upgrading from a standard single-CCD model to a dual-CCD model could also require a complete reinstall, an unnecessary and quite irritating situation for end users who might not even be aware of this requirement.


Regardless, this also creates problems for reviewers who test multiple processors on the same motherboard. (It’s even conceivable that this issue contributed to the inconsistent results we have seen with the first wave of Ryzen 9000 reviews, as 7000X3D models can also trigger the issue and cause all other chips tested on the platform to not operate as intended.)


We retested to ensure that the issue did not impact our test results, but we also noticed that the core parking feature isn’t bulletproof — we observed cores slipping in and out of a parked state during gaming on several occasions. This would obviously have a negative performance impact, but the feature also worked perfectly fine at other times. This problem could be specific to our test platform, but we weren’t notified about the feature until late in the review process, so we haven’t had time to dive in for a deeper look. As such, take that into account when viewing our gaming benchmarks.

- **MORE:** **Best CPU for gaming**
- **MORE:** **CPU Benchmark Hierarchy**
- **MORE:** **Intel vs AMD**
- **MORE:** **How to Overclock a CPU**

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

- 
Your story mentioned the USB 4.0 60GB/s with 3rd party chips, is that a new feature? its 40GB/s and then the newest standard of Thunderbolt is 80/80 (or 120 if you change one of the channels to make it asymetrical). There is no such thing as 60Reply
 
And pretty much like every other enthusiest on here, I think th is is the most disapointing AMD launch...ever
- 
Reply
 Good eye. Typo, fixed!HideOut said:Your story mentioned the USB 4.0 60GB/s with 3rd party chips, is that a new feature? its 40GB/s and then the newest standard of Thunderbolt is 80/80 (or 120 if you change one of the channels to make it asymetrical). There is no such thing as 60
 
And pretty much like every other enthusiest on here, I think th is is the most disapointing AMD launch...ever
- 
Very disappointing the single thread and gaming results. Anyway I find suspicious results on some benchmarks.Reply
 
 There are regressions respect 7950X in games like "Far Cry 6" and "Hitman 3".
 
 Regression also in Outlook score vs 7950X.
 
 I found abnormal also the NAMD score that for the 9950X is 0.6328 days/ns while Phoronix achieved a 3.14 ns/day that is about twice the performance, for comparision the 14900k is in line with Phoronix.
 
 I suspect there is something to analyse and refine on these benchmarks.
 
