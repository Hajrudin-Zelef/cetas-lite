---
id: collect-240926-tomshardware/tomshardware/intel-core-ultra-9-285k-and-core-5-245k-review-intel-throws-a-lateral-with-arrow-2
title: "intel-core-ultra-9-285k-and-core-5-245k-review-intel-throws-a-lateral-with-arrow"
domain: tomshardware
role: reference
task: reference
actors: ["Intel", "Microsoft", "Nvidia", "TSMC"]
dates: []
keywords: ["intel", "3nm", "attention", "benchmarks", "chiplet", "compute", "consumer", "copilot", "cost", "disaggregated", "dram", "energy"]
source: docs/RAG/clean_en/tomshardware/intel-core-ultra-9-285k-and-core-5-245k-review-intel-throws-a-lateral-with-arrow.md
source_anchor: ""
source_lines: [68, 114]
sha256: 1473aee57d0ea0cb7be0a05c2151a521f3165a5be6006a16ed2e618f25e460bb
---

# intel-core-ultra-9-285k-and-core-5-245k-review-intel-throws-a-lateral-with-arrow

Intel made numerous adjustments to the stack, including reductions in peak clock speeds — a somewhat expected byproduct of moving to TSMC's N3B process node. The Ultra 9 peaks at 5.7 GHz, 300 MHz less than the prior-gen 14900K's blistering 6 GHz boost, while the Ultra 7 and 5 boost clocks see 100 MHz reductions. However, Intel has adjusted P-core base clocks upward by 500 to 700 MHz. The E-cores also have boost clock improvements spanning from 200 to 600 MHz, and a 600 MHz to 1 GHz improvement in base clock speeds, all of which vary by model.

Despite Arrow Lake's claimed lower operating power consumption, the chips still have similar maximum TDP (MTP) ratings of 250W for Ultra 9 and 7 (3W less than Raptor Lake) and 159W (22W less) for the Ultra 5 model. Intel says the lower power consumption occurs during normal workloads with an up to 40% reduction in package power consumption. Intel also increased the maximum CPU temperature (TJMax) to 105C for Arrow Lake, which is 5C higher than its traditional limit with mainstream PC processors.

The new chips come with 24 lanes of PCIe 5.0 support, with an additional 20 PCIe 4.0 lanes provided by the chipset. The Ultra 9, 7, and 5 all have the same Xe-LPG graphics engine with four Xe cores — the same GPU as the Meteor Lake chips, not the newer Battlemage Xe2 engine found in the Lunar Lake mobile chips. The Ultra 9 and 7 iGPU have a 2.0 GHz graphics boost clock, while the Ultra 5 drops to 1.9 GHz. Intel claims the iGPU offers twice the performance of the graphics on the 14th-gen processors, but we haven't yet had time to put those claims to the test. Also, integrated graphics performance on a desktop chip largely won't matter for many users, as if you actually care about GPU performance you just add a discrete graphics card.

The chips also feature the same in-built NPU engine for AI acceleration as Meteor Lake — not Lunar Lake. This engine provides up to 13 TOPS of INT8 throughput, but that doesn't meet Microsoft's requirement of 40+ TOPS to unlock Copilot+ features. Intel already has a larger NPU design in the market — the (up to) 48 TOPS engine found in Lunar Lake — but used the smaller engine to optimize the die area for the desktop PC market. A larger engine would chew into the space available for other additives, like cores and cache. Intel says it hasn't seen enough interest in the desktop PC market yet to sacrifice other areas of performance for the NPU. As with the iGPU, the decision makes sense: If you care about AI compute, even Nvidia's lowest tier RTX 4060 offers 242 TOPS of INT8 performance.

Arrow Lake drops into the LGA 1851 socket, so the chips are incompatible with existing motherboards. Existing coolers should be compatible with the requisite mounting hardware, but the need for a kit could vary by vendor. Intel hasn't committed to using the LGA 1851 socket for future processor generations. We've seen signs of a Core Ultra 2000S Refresh generation in the works, but that isn't yet confirmed. That could mean that LGA 1851 will end up as a single-generation socket.

## CUDIMMs and Memory Support Matrix

Arrow Lake supports up to 192GB of DDR5 memory, but now in two flavors with two different base speeds. The chips support DDR5-6400 with DDR5 CUDIMMs, a new type of DIMM with an integrated clock driver (ckd) that boosts easily attainable stable clock frequencies by amplifying the signal, thus stabilizing the data eye. Unlike the clock redrivers present on fully-buffered registered DIMMs, the CUDIMM redrivers are said not to impose an additional clock cycle of latency (they use a less complicated and cheaper design).

Intel also points to much higher overclocking headroom with CUDIMMs and says DDR5-8000 appears to be the sweet spot (Gear 2). CUDIMMs should enable the use of poorer-quality DRAM ICs in higher-speed kits while simplifying the pricier DIMM PCB designs often required for higher-end memory. But motherboards with CUDIMM support may cost extra, and the CUDIMMs themselves are likely to carry a price premium, so you'll need to pay close attention to the final cost before deciding whether CUDIMMs make sense.

Intel also supports standard DDR5, but at lower base speeds than it supports with CUDIMMs (the same DDR5-5600 as with its 14th Gen CPUs). Naturally, both types of memory are overclockable. The Arrow Lake DDR5 support matrix is in the table below. Arrow Lake does support ECC memory, but it won’t be supported on consumer platforms — instead, that feature is reserved for enterprise-focused W-series motherboards.

| DRAM Config | Official Speeds Supported | 
| Dual Channel, 2 board slots, 2 UDIMMs | DDR5-5600 | 
| Dual Channel, 2 board slots, 2 CUDIMMs | DDR5-6400 | 
| Dual Channel, 4 board slots, 2 UDIMMs | DDR5-5600 | 
| Dual Channel, 4 board slots, 2 CUDIMMs | DDR5-5600 | 
| Dual Channel, 4 board slots, 4 UDIMMs | DDR5-4800 (single rank DIMMs) DDR5-4400 (dual rank DIMMs) | 
| Dual Channel, 4 board slots, 4 CUDIMMS | DDR5-4800 (single rank DIMMs) DDR5-4400 (dual rank DIMMs) | 

## Intel Arrow Lake Core Ultra 200S architecture

Arrow Lake marks Intel's first foray into a disaggregated architecture, meaning that different compute and I/O functionalities are split out into their own dies. Intel refers to its die disaggregation technique as a ‘tiled’ architecture, but the rest of the industry refers to this as a chiplet architecture. Intel says that a ‘tiled’ processor refers to a chip using advanced packaging, which enables parallel communication between the chip units, while standard packaging employs a serial interface that isn’t as performant or energy efficient. However, other competing processors with advanced packaging are still referred to as chiplet-based, so the terms are largely interchangeable.

Instead of the newer Lunar Lake design, Arrow Lake uses a package design similar to its five-tile previous-gen Meteor Lake laptop processors. However, Intel integrated the newer Lion Cove P-core and Skymont E-core microarchitectures for the compute tile instead of the Redwood Cove and Crestmont cores it used in Meteor Lake.

The Arrow Lake design employs Intel chip designs etched on a compute tile (chiplet) fabbed on the TSMC N3B process node, a GPU tile with the TSMC N5P node, while the SoC and I/O tiles use TSMC’s N6 process. Intel uses Foveros 3D packaging to mount those tiles to an underlying base tile fabbed on the Intel 1227.1 process node (22nm FinFET). There are also two ‘dummy’ filler tiles that provide mechanical rigidity.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

| Row 0 - Cell 0 | Arrow Lake - Manufacturer / Node | Meteor Lake - Manufacturer / Node | 
| CPU Tile | TSMC / N3B (3nm) | Intel / 'Intel 4' | 
| 3D Foveros Base Die | Intel / 22FFL (Intel 16) | Intel / 22FFL (Intel 16) | 
| GPU Tile (tGPU) | TSMC / N5P (5nm) | TSMC / N5 (5nm) | 
| SoC Tile | TSMC / N6 (6nm) | TSMC / N6 (6nm) | 
| IOE Tile | TSMC / N6 (6nm) | TSMC / N6 (6nm) | 

Intel says the chip has 17.8 billion transistors spread out over a total die area of 243mm^2. Unfortunately, given the different nodes employed and the fact that Intel included the filler tiles in the total die area, that doesn't tell us much about transistor density.

Intel’s decision to split the memory controller and PHY into their own tile (I/O tile) was to improve yields and preserve the precious 3nm transistors for compute, but this creates memory latency issues that contribute to the lower gaming performance we see on the next page. Intel does offer the option to overclock the tile-to-tile interface, but we haven't had a chance to test the impact yet. Here's the memory latency we measured with AIDA, and our cache and memory latency benchmarks using the Memory Latency tool from the Chips and Cheese team.

