---
id: collect-250926-servers-hardware/servers-hardware/intel-announces-xeon-600-series-this-is-granite-rapids-for-workstations
title: "intel-announces-xeon-600-series-this-is-granite-rapids-for-workstations"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["intel", "amd", "chiplet", "compute", "ethernet", "latency", "memory", "pricing"]
source: docs/RAG/clean4/intel-announces-xeon-600-series-this-is-granite-rapids-for-workstations.md
source_anchor: ""
source_lines: [1, 84]
sha256: dddb5285b8ef627a8cdb6905ce9ab7b75907d3c885ccc6b4996c11477960d2ac
---

# intel-announces-xeon-600-series-this-is-granite-rapids-for-workstations

Intel this afternoon is taking the wraps off of the Xeon 600 series, the company’s upcoming workstation platform. Based on Intel’s Granite Ridge CPUs, the release of the Xeon 600 series will see Intel’s latest server technology finally cascade down into workstation parts, replacing the current Xeon W-2500/W-3500 (Sapphire Rapids) platforms that have been the basis of Intel’s Xeon workstation lineup for the last few years.

As with the Xeon 6 family of server parts, the big-ticket improvements in the Xeon 600 Processors for Workstations family include the newer and faster Redwood Cove CPU cores, as well as offering broader core counts overall, with top Xeon 600 chips now hitting 86 CPU cores. The Granite Rapids hardware at the heart of the latest processors also brings a larger number of PCIe lanes, CXL 2.0 support, and even MRDIMM support for higher transfer rates. Ultimately, the end product is meant to be a thorough generational update that brings more I/O, more memory bandwidth, and more multi-threaded performance to Intel’s high-end P-core workstation platform.

The launch of the Xeon 600 series will also mark the end of Intel’s Xeon W branding, which for the last several years has been the company’s designation for workstation-class processors. In its place will be the singular Xeon 600 series, which besides being a fresh start in terms of naming that is aligned with the server-focused Xeon 6 series, will also see Intel consolidate their workstation offerings a bit. The separate Xeon W-2×00 and W-3×00 CPUs and associated platforms are being reduced to a single offering under the Xeon 600 umbrella. Though note that does not mean that all of these chips are identical in features, as Intel is still shipping two tiers of chips: lower-end chips will feature 4 memory channels and 80 PCIe lanes, while higher-end chips will get all 8 memory channels and 128 PCIe lanes that the platform supports.

Finally, with the launch of the Xeon 600 family comes a new chipset: W890. Intel’s latest workstation chip is pretty similar to its predecessor on the whole, but this newer chipset includes Wi-Fi 7 support, helping it to keep pace with Intel’s desktop platforms.


## Xeon 600: Redwood Cove CPU Cores & MRDIMMs For Workstations

At the heart of the Xeon 600 series is Intel’s Granite Rapids silicon, which has been shipping in servers as the Xeon 6 series for the better part of a year now. Granite Rapids, in turn, is Intel’s high-end chip offering based around their Redwood Cove CPU P cores, which were first introduced with client Meteor Lake (Core Ultra 100 Series) processors in early 2024.

Compared to the Raptor Cove cores that Redwood Core replaced, Redwood itself is not a massive step up in terms of either IPC or attainable clockspeeds. What it has going for it instead are both area efficiency and power efficiency, which combined with the use of the company’s Intel 3 fab node, has allowed Intel to pack a much larger number of CPU cores into a single Granite Rapids chip. As a result, while the top Xeon W-3500 chip offered 60 CPU cores, the top Xeon 600 chip brings that to 86 cores – or about a 43% increase in core counts, running at roughly the same peak frequency as before.

The end result is that the bulk of the Xeon 600’s performance gains come from multi-threaded workloads that can take advantage of those additional cores. Intel’s own promotional material is pretty straightforward on this, listing the single-threaded performance gains for Intel’s top SKU at just 9%, versus 61% higher multi-threaded performance. Single-threaded gains are hard to get these days period, and with the Xeon 600 series that is even more true.

Traditional CPU performance improvements aside, Granite Rapids also brings with it one other notable feature that Intel is hoping will give the company a wider edge in the AI market: FP16 support in its AMX units. The latest generation of Intel’s dedicated matrix multiplication hardware for mid-to-low precision math, the company previously only supported INT8 and BFloat16 in these cores – so the addition of FP16 support is a modest expansion in its capabilities. For Intel, this also marks one of the few major feature differentiators between it and AMD, as the company’s rival does not have anything equivalent to AMX on its CPU cores.

As with Xeon 6 servers, the other half of the Granite Rapids story is the I/O and memory bandwidth available to the processor. Intel bumped up both of those things with their new processor architecture, and all of those improvements are cascading down to the Xeon 600 series – at least for higher-end SKUs.

On the I/O front, Intel now offers a full 128 lanes of PCIe Gen5 coming from the host CPU, which is up from 112 lanes on the previous-generation Xeons – in practice allowing for one more x16 slot. The PCie controller on Granite Rapids also gains improved CXL capabilities with CXL 2.0 support, which allows for memory pooling and single-level switching.

As for memory, Granite Rapids comes with a thorough increase in memory bandwidth. All Xeon 600 SKUs will support DDR5-6400 speeds for 1 DPC, or slower speeds at 2 DPC. Conversely, for users who need even more bandwidth, the top half of the Xeon 600 series (28 cores and higher) also brings support for Multiplexed Rank DIMMs (MRDIMM), which allows for effective memory speeds of up to DDR5-8000.

But as noted before, not all Xeon 600 SKUs will get these features. Lower-end chips (16 cores and below) will be limited to just 80 PCIe lanes and 4 memory channels, which is very similar to the outgoing Xeon W-2400/2500 parts. So while Intel has gotten rid of the Xeon W-2×00 family as an official lineup, they have not stopped offering more limited chips on the low-end – a minimal buy-in of 18 cores is needed to get access to the rest of those DDR5 memory channels and PCIe lanes.

Finally, paired with all of the Xeon 600 processors will be Intel’s new W890 chipset. This is connected to the host CPU via a DMI Gen4 x8 link, offering just shy of 16GB/second of bandwidth between the two. Compared to the new Xeon CPUs, the W890 is admittedly unremarkable here: Intel’s USB speeds still top out at 20Gbps USB 3.2 (Gen 2×2), with BMCs, SATA ports, additional PCIe Gen4 lanes, and networking all hanging off of the chip. Wired Ethernet users can look forward to 1Gbps or 2.5Gbps ports, while for wireless users one of the handful of new features here is Wi-Fi 7 support (though while unconfirmed by Intel, I suspect this is using a fully discrete controller rather than CNVIo3).

Asus, Gigabyte, and Supermicro are among the initial vendors who are slated to offer W890 motherboards.

## SKUs: 11 Chips, From 12 to 86 Cores

For their latest workstation offering, Intel will be supplying a product stack of 11 SKUs, with core counts ranging from 12 cores all the way up to 86 cores. With the entire platform being a workstation configuration of Granite Rapids-SP, server customers should be quick to clue in on what is going on here, with Intel offering chips based on their low (LCC), high (HCC), and extreme (XCC) Granite Rapids configurations, which respectively use either 1 or 2 Granite Rapids compute tiles on a single socket LGA 4710 chip.

It is notable here that the Xeon 600 stack is somewhat slow to ramp up in regards to core counts. 5 of the SKUs have 20 or fewer CPU cores, and 8 of them have 32 or fewer cores. Only the top 3 SKUs go beyond this at 48, 64, and 86 CPU cores respectively. So at present time, for users with workloads can that leverage more than a couple dozen CPU cores, the steps up in performance and pricing get very steep very quickly.

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
