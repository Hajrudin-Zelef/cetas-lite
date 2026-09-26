---
id: collect-250926-servers-hardware/servers-hardware/intel-announces-xeon-600-series-this-is-granite-rapids-for-workstations-2
title: "intel-announces-xeon-600-series-this-is-granite-rapids-for-workstations"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["intel", "amd", "chiplet", "compute", "latency", "memory", "pricing"]
source: docs/RAG/clean4/intel-announces-xeon-600-series-this-is-granite-rapids-for-workstations.md
source_anchor: ""
source_lines: [38, 84]
sha256: 7558fc032047abbde7435c155d19d6a5b864f85ac12bf1bdd4214dbe99075b3c
---

# intel-announces-xeon-600-series-this-is-granite-rapids-for-workstations

Overall, TDPs are very similar to the previous generation, with Intel essentially investing all of their efficiency gains into providing more CPU cores at the same tier. The top two chips, the 698X and 696X, will have base TDPs of 350 Watts and peak TDPs of 420 Watts. Meanwhile the rest of the stack drops steps down with the core count, with the lightest 12 core chip, 634, coming in at 150 Watts for its base TDP.

As with the prior generations, a selection of chips will be made available in boxed form as well. These are the 654 (18C), 658X (24C), 676X (32C), 678X (48C), and 696X (64C).

And the X-tier chips are unlocked for overclocking. Though Intel is quick to note that they expect this functionality to primarily be used for system integrators rather than direct customers. But regardless of who ends up using it, Granite Rapids does bring a couple of enhancements here, most notably with the amount of reporting data the processor provides for diagnosing why overclocks are not working/holding.

## Performance Expectations

In discussing the performance of the new chips, Intel is opting to limit performance comparisons to its own Core and Xeon chips. They are not offering any formal performance comparisons to AMD’s Threadripper chips, so make of that what you will. Though with that said, the company says that they are expecting the Xeon 600 series to be pretty competitive here, particularly with its support for a larger amount of memory (thanks to supporting 2 DIMMs per channel) as well as an edge in performance-per-dollar.

As for the Intel-on-Intel comparisons, the long and short of things is that it is a potentially wide range, depending in part on how well a workload can take advantage of the additional CPU cores and/or memory bandwidth. SPEC Workstation 4.0 has everything between a minor regression to a 61% performance gain depending on the sub-test. In other cases Intel is seeing 20-30% performance gains.

Meanwhile, the company is also being surprisingly candid on the Xeon 600 series’ strengths and weaknesses. Rather than pitching the Xeon 600 as the end-all-be-all of workstation computing performance, they are encouraging customers to make sure the platform makes sense for their workload. Namely, is the workload sensitive to latency and peak single-threaded performance, or is it more constrained by core counts and memory bandwidth? With Intel’s Arrow Lake platform (Core Ultra 200) offering both lower latencies and newer CPU cores with higher clockspeeds, it can run right past the Xeon 600 in lightly-threaded workloads.

## Xeon 600 Series: Available In Late March

Wrapping things up, today’s Xeon 600 announcement from Intel is coming roughly two months ahead of product availability. The company expects Xeon 600 components and systems to become available in late March.

Vendors will include all of the usual suspects. ASUS, Supermicro, and Gigabyte are all slated to offer W890 motherboards. Meanwhile Dell, HP, Supermicro, Boxx, Puget Systems, and others will be offering complete systems. Pricing for those systems will be revealed in due time, but with list prices of Xeon Processors ranging between $500 and $7700, expect a rather wide range in terms of system pricing as well.

Finally, while the release of the Xeon 600 series will set Intel’s desktop workstation offerings for the next couple of years or longer (as Intel previously cancelled its next-gen mainstream Xeons), the mobile side of the market is coming due for its own updates. So expect to see mobile workstations based around the recently-launched Core Ultra Series 3 processors a bit farther down the line.

“Pudget Systems”?

My pudgy fingers may have hit one too many keys. Thanks!

I expect threadripper trounces these chips, just like the desktop Zen 5 chips trounce the desktop Core Ultra chips. Top TR chip has more cores as well, and the lack of slides comparing them is very telling. The pricing seems to not be aggressive either. Intel’s 32c part costs the same as AMDs, the 9970x is also $2500. Very unimpressive, honestly.

For Threadripper comparisons, 9975WX would probably be the more apt one. Intel’s sole 32 core SKU has 8 memory channels and 128 PCIe lanes, which only the Pros match.

What dies are used for which SKU? Intel has finally moved to a chiplet style layout but unlike AMD, there are two memory controllers per chiplet. Thus more chiplets, the more memory channels are possible. Going by the slides (which admittedly may not be representative of the real die layout), each chiplet has 44 cores and 4 DDR5 memory channels. That 48 core SKU is mostly dead silicon. Even worse looking at the 18 core part with 8 memory channels. Is there a smaller die used for the low end parts that has 8 DDR5 memory channels but fewer CPU cores? Do these smaller chiplets have the same connections to the IO dies?

In terms of performance, the tighter integration with the memory controllers and higher bandwidth per memory channel will show some gains compared to the AMD competition. Much will come down to the test being performed I would predict. With memory prices being stupid, a performance advantage there may not translate into any $/perf gains at the platform level. Given the time frame between this release and the Threadripper 9000 series, Intel needed a clear win, not playing catch up. Zen 6 Threadrippers are likely a year away to leap frog things yet again. AMD still has the option to release 3D Vcache enabled Threadrippers on a whim to really shake things up. A clock core count, high clock, high cache capacity part from AMD would be interesting to see as a comparison point.

Kevin G, I think it does make sense if there is a lot of dead silicon on those boards- it allows them significantly increase their relative die yields.

Aim for the best, sell the rest (as a lower SKU.)

@Kevin

According to Intel it’s the XCC configuration for the 86C and 64C parts. HCC configuration for 20-48 cores, and LCC for 12 to 16 cores. It’s the same as the server chips, in other words.

HCC and LCC use a single compute tile each. XCC uses two compute tiles.

Honestly, Intel should have stuck to multiples of 8 on the core counts to keep things simplified at launch. That still gives a stack of 9 options (compared to the 11 above), but with cleaner delineation between products.

16, 24, 32, 40, 48, 56, 64, 72, 80
