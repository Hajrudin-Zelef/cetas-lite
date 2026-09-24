---
id: collect-240926-tomshardware/tomshardware/amd-ryzen-9-9950x-review-zen-5-at-full-power
title: "amd-ryzen-9-9950x-review-zen-5-at-full-power"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: ["2023-04"]
keywords: ["amd", "benchmark", "benchmarks", "chiplet", "compute", "intel", "latency", "memory", "pricing"]
source: docs/RAG/clean_en/tomshardware/amd-ryzen-9-9950x-review-zen-5-at-full-power.md
source_anchor: ""
source_lines: [1, 155]
sha256: cbe44c1c0a8e9dc3075e199753cd28ad5f0e5213a1b46fd59a3e5d3068a00359
---

# amd-ryzen-9-9950x-review-zen-5-at-full-power

<!-- source: https://www.tomshardware.com/pc-components/cpus/amd-ryzen-9-9950x-cpu-review -->

### Tom's Hardware Verdict

The Ryzen 9 9950X offers performance improvements across the board and the highest performance available on a mainstream PC platform in multi-threaded workloads. However, it lags behind competing chips in gaming, and the generational gains are small enough in some productivity workloads that the previous-gen Ryzen 9 7950X is an attractive alternative.

#### Pros

- +Leading multi-thread performance
- +Strong single-threaded performance
- +Native AVX-512 support
- +Leverages AM5 motherboards, backward compatible
- +Overclockable

#### Cons

- -Core parking and thread targeting issues
- -Pricing
- -Gaming performance isn't as competitive as other chips
- -Slim generational gains in some workloads

AMD’s $650 Ryzen 9 9950X comes with 16 cores and 32 threads, slotting in as the flagship model for the company’s new ‘Granite Ridge’ family of processors sporting the Zen 5 architecture. AMD has infused its innovative thread-targeting tech, previously reserved for the gaming-optimized 3D V-Cache models, into the upper-tier Ryzen 9000 models. However, the improvements in gaming performance aren’t enough to take the crown from Intel’s competing Core i9-14900K flagship, let alone AMD’s own X3D processors that lead our list of the best CPUs for gaming. That's not to mention that AMD's core parking feature has multiple issues, which we cover below.


The Ryzen 9 9950X has made significant gains in single- and multi-threaded productivity applications, carving out convincing leads in several workloads. That’s at least partly due to AMD’s different approach with the 9950X — all other Zen 5 desktop PC processors have lower gen-on-gen TDP ratings, but the 9950X stays at the same 170/230W rating as the prior-gen models while benefiting from Zen 5’s other improvements, like a claimed 16% improvement in IPC, fueling a strong lead in multi-threaded workloads.


Unlike Intel’s competing Raptor Lake processors, Ryzen 9000 also has full native AVX-512 support, a boon for productivity work. The 9950X also delivers a notable gain in single-threaded performance, significantly reducing the gap with the 14900K. 

| AMD Zen 5 Ryzen 9000 'Granite Ridge' Specifications and Pricing |  |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|---|
|  | Street/MSRP | Arch | Cores / Threads | Base / Boost Clock (GHz) | Cache (L2/L3) | TDP / PPT | Memory | 
|---|---|---|---|---|---|---|---|
| **Ryzen 9 9950X** | **$649** | **Zen 5** | **16 / 32** | **4.3 / 5.7** | **80MB (16+64)** | **170W / 230W** | **DDR5-5600** | 
| **Ryzen 9 9900X** | **$499** | **Zen 5** | **12 / 24** | **4.4 / 5.6** | **76MB (12+64)** | **120W / 162W**  | **DDR5-5600** | 
| **Ryzen 7 9700X** | **$359** | **Zen 5** | **8 /16** | **3.8 / 5.5** | **40MB (8+32)** | **65W / 88W** | **DDR5-5600** | 
| **Ryzen 5 9600X** | **$279** | **Zen 5** | **6 / 12** | **3.9 / 5.4** | **38MB (6+32)** | **65W / 88W** | **DDR5-5600** | 

Due to a scheduling delay, AMD’s Ryzen 9 9950X arrives alongside the $500 Ryzen 9 9900X as the second salvo of the company’s Zen 5 lineup for desktop PC processors. AMD released the Ryzen 5 9600X and Ryzen 7 9700X last week, and that launch has been exceptionally rocky, particularly on the gaming front — inconsistent benchmark results and poor baseline performance metrics in the company’s materials have led to conflicting findings about the state of Zen 5 gaming.


Intel has had its own difficulties, too. AMD’s Zen 5 launch arrives as Intel has finally begun to address long-running instability issues with its 13th- and 14th-Gen processors. Intel’s instability issue can potentially impact all its 65W and higher CPUs, and an unknown percentage of processors have experienced the problems. Intel has now delivered a new microcode path to address the issue, and it has also extended its warranty for the impacted product lines by two additional years. We’ve retested the Intel processors with the new microcode for this review.


AMD’s Ryzen 9 9950X takes the lead in our suite of multi-threaded application benchmarks. While it delivers some uplift in gaming, it relies on the strength of performance in productivity workloads to justify the $650 price tag, especially because stepping up to higher-end workstation-class fare like AMD’s Threadripper Pro and Intel’s Xeon W has become exceedingly expensive. Intel has its competing Arrow Lake processors in the hopper for release, ostensibly by the end of the year, leaving the Ryzen 9 9950X time to reign as the leader in heavily threaded work. However, the 9950X's $650 price tag leaves room for lower-priced alternatives to offer a better value proposition for productivity-focused users.

## AMD Ryzen 9 9950X Specifications and Pricing

In our Ryzen 5 9600X review, we covered the Zen 5 microarchitecture, new motherboards, and new overclocking features. Head there for more detailed information on the architecture. Below, we’ll cover the basics, along with information about AMD’s game-boosting core parking tech for the Ryzen 9 9950X and the problems we’ve encountered with it.

| AMD Zen 5 Ryzen 9000 'Granite Ridge' Specifications and Pricing |  |  |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|---|---|
|  | Street/MSRP | Arch | Cores / Threads (P+E) | P-Core Base / Boost Clock (GHz) | E-Core Base / Boost Clock (GHz) | Cache (L2/L3) | TDP / PBP / MTP | Memory | 
|---|---|---|---|---|---|---|---|---|
| Ryzen 9 7950X3D | $525 ($699) | Zen 4 X3D | 16 / 32 | 4.2 / 5.7 |  | 144MB (16+128) | 120W / 162W | DDR5-5200 | 
| Core i9-14900K / KF | $550 (K) - $530 (KF) | Raptor Lake Refresh | 24 / 32 (8+16) | 3.2 / 6.0 | <2.4 / 4.4 | 68MB (32+36) | 125W / 253W | DDR4-3200 / DDR5-5600 | 
| **Ryzen 9 9950X** | **$649** | **Zen 5** | **16 / 32** | **4.3 / 5.7** |  | **80MB (16+64)** | **170W / 230W** | **DDR5-5600** | 
| Ryzen 9 7950X | $525 ($699) | Zen 4 | 16 / 32 | 4.5 / 5.7 |  | 80MB (16+64) | 170W / 230W | DDR5-5200 | 
| Ryzen 9 7900X3D | $420 ($599) | Zen 4 X3D | 12 / 24 | 4.4 / 5.6 |  | 140MB (12+128) | 120W / 162W | DDR5-5200 | 
| **Ryzen 9 9900X** | **$499** | **Zen 5** | **12 / 24** | **4.4 / 5.6** |  | **76MB (12+64)** | **120W / 162W**  | **DDR5-5600** | 
| Ryzen 9 7900X | $360 ($549) | Zen 4 | 12 / 24 | 4.7 / 5.6 |  | 76MB (12+64) | 170W / 230W | DDR5-5200 | 
| Core i7-14700K / KF | $400 (K) - $375 (KF) | Raptor Lake Refresh | 20 / 28 (8+12) | 3.4 / 5.6 | 2.5 / 4.3 | 61MB (28+33) | 125W / 253W | DDR4-3200 / DDR5-5600 | 
| Ryzen 7 7800X3D | $375 ($449) | Zen 4 X3D | 8 /16 | 4.2 / 5.0 |  | 104MB (8+96) | 120W / 162W | DDR5-5200 | 
| **Ryzen 7 9700X** | **$359** | **Zen 5** | **8 /16** | **3.8 / 5.5** |  | **40MB (8+32)** | **65W / 88W** | **DDR5-5600** | 
| Ryzen 7 7700X | $295 ($399) | Zen 4 | 8 /16 | 4.5 / 5.4 |  | 40MB (8+32) | 105W / 142W | DDR5-5200 | 
| Core i5-14600K / KF | $300 (K) - $294 (KF) | Raptor Lake Refresh | 14 / 20 (6+8) | 3.5 / 5.3 | **2.6 / 4.0** | 44MB (20+24) | 125W / 181W | DDR4-3200 / DDR5-5600 | 
| **Ryzen 5 9600X** | **$279** | **Zen 5** | **6 / 12** | **3.9 / 5.4** |  | **38MB (6+32)** | **65W / 88W** | **DDR5-5600** | 
| Ryzen 5 7600X | $210 ($229) | Zen 4 | 6 / 12 | 4.7 / 5.3 |  | 38MB (6+32) | 105W / 142W | DDR5-5200 | 

The Ryzen 9 9950X replaces the previous-gen Ryzen 9 7950X, but aside from the newer CPU microarchitecture it has the same basic accommodations, like 16 cores, 32 threads, a 5.7 GHz peak frequency, 80MB of combined L2+L3 cache, and a 170W/230W TDP/max rating. AMD also dialed back the base clock by 200 MHz, though that generally doesn't matter much in our testing. The Ryzen 9 9950X will initially compete with Intel’s $548 Core i9-14900K, but Intel’s Arrow Lake processors will be the real competitors for Ryzen 9000.


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
 
 Agree that the driver problem with core parking feature is something unacceptable. But the need to reinstall the OS seems to me something exaggerated, must exist a better solution.
 
 Also is historical the surpass in Adobe benchmarks.
 
Thanks for the review.
- 
There is obviously something fishy with Windows when you look at the results on Linux.Reply
 
 "The Ryzen 9 9950X was 33% faster than the Intel Core i9 14900K performance overall and even the Ryzen 9 9900X was 18% faster than the Core i9 14900K. For those still on AM4, the Ryzen 9 9950X was delivering 1.87x the performance of the Ryzen 9 5950X processor. These are some great gains found with the Ryzen 9 9900 series."
 
 -Phoronix
 
- 
Whilst it maybe disappointing for some, it simply appears that AMD are continuing to separate out their gaming CPU range from their productivity range.Reply
 
 This appears to be a very powerful and efficient chip for multi core workloads that will likely get better when some of the usual teething problems are ironed out.
 
Nothing groundbreaking, just a bit better.
