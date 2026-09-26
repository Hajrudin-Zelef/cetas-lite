---
id: collect-240926-tomshardware/tomshardware/amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-1
title: "amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Intel", "TSMC"]
dates: []
keywords: ["amd", "2nm", "3nm", "compute", "cost", "energy", "gpus", "intel", "reasoning"]
source: docs/RAG/clean_en/tomshardware/amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-.md
source_anchor: ""
source_lines: [1, 61]
sha256: a018fab4574049e5823f6005894668cd9a55bd90a6c9811597cc399d1e4462e5
---

# amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-

<!-- source: https://www.tomshardware.com/pc-components/cpus/amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-in-intels-footsteps-with-new-core-type-for-background-tasks -->

AMD has submitted Linux kernel patches including support for its new low-power CPU cores that will likely emerge in its future heterogeneous processors. The new patch clearly distinguishes between high-performance cores, efficiency cores, and low-power cores, so it is safe to say that AMD's upcoming CPU platforms will use three types of cores, with the low-power one serving light workloads, reports *Phoronix*.

AMD's heterogeneous processors identify CPU types using CPUID Function 0x80000026 (Extended CPU Topology), as EBX bits [31:28] carry the core classification. Up until recently, AMD only classified its cores as Performance and Efficiency, while the latest patch adds Low-Power cores. The patch enables Linux to distinguish between Performance, Efficiency, and Low-Power cores efficiently, and the latter are also correctly supported by AMD's performance management.

According to AMD engineer Vishal Badole, these cores are designed specifically for background and idle tasks where reducing energy consumption is more important than offering high performance.

In recent years, both AMD and Intel released heterogeneous processors featuring both high-performance and energy-efficient types of cores in a bid to wed performance and low power consumption. With its latest CPU platforms, Intel introduced its low-power cores located in the SoC tile to offload light tasks and prolong the battery life of laptops. As it turns out, AMD is going the same route. Although AMD uses two different core types, the underlying architecture is the same. It offers a "dense" core offering that's optimized for space, while Intel uses entirely different microarchitectures.

Beyond the description of the Linux patch, AMD disclosed little about the low-power cores. The company only described them as being optimized for the lowest possible power consumption during background processing and idle operation, but did not reveal how they differ architecturally from today's dense Zen 5c cores. In addition, the kernel patches introduce no new scheduling policies or optimization logic beyond identifying the additional CPU category.

AMD also did not reveal whether these cores are based on the Zen 5, Zen 6, or other future microarchitecture. It should indeed be noted that AMD has traditionally preferred to use the same microarchitecture within one CPU, albeit with different optimizations when it comes to die size (or rather floorplan) and clocks. Such an approach greatly simplifies software development and performance management, but at the cost of higher power consumption compared to what a simplified microarchitecture would have offered.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Anton Shilov is a contributing writer at Tom’s Hardware. Over the past couple of decades, he has covered everything from CPUs and GPUs to supercomputers and from modern process technologies and latest fab tools to high-tech industry trends.

- 
I've argued about the Zen 6 LP cores here back when they were a leak. Desktops should get them, and they can lower idle/low-intensity power consumption and save everybody (small amounts of) money. Governments and businesses will appreciate lowering power across millions of x86 machines.Reply
 
 The E-cores/C-cores are there to maximize multi-threading performance per die area for the most part. LPE-cores and LP cores will be truly efficiency focused, especially if they can turn off the compute die(s) completely when they aren't needed.
 
 Having more of them could prevent tasks from spilling over to compute die(s). Every Nova Lake-S SKU will have 4c/4t of these. I think they should move up to 8c/8t in a future generation.
 
 Your operating system could have a super power saving mode to manually disable CCDs and run everything on LPs. For example, if you're trying to run a mini PC off a battery pack, are working in a hot environment, etc.
 
Zen 6 mobile APUs will definitely get LP cores. Zen 6 Olympic Ridge desktop CPUs are still uncertain, and we've recently heard they may not include an iGPU, but will include a big NPU. If they do include Zen 6 LP cores, it will be 2c/4t. If they make the final cut, I suspect they may include a tiny amount of dedicated L3 cache, e.g. 4 MiB.
- 
Reply
 Strix Point and Kraken Point basically qualify as that. Strix Point is very heterogeneous with 4x Zen 5 + 16 MiB L3 in one CCX, and 8x Zen 5c + 8 MiB in the other.ezst036 said:In other words, is AMD prepping a big.little CPU?
 
 Zen 6 APUs could look very complicated, with all 3 types of cores sitting together in monolithic dies, or I/O dies with these cores and an optional 12-core desktop CCD added.
 
Zen 6 desktop CPUs will have up to 2x 12-core Zen 6 chiplets, and possibly 2-core Zen 6 LP on the I/O die. Playing the same role as LPE cores did on Meteor Lake APUs or upcoming Nova Lake-S desktop CPUs.
- 
Reply
 Turning off compute tiles only makes sense as long as the LP cluster doesn't use more power than they do...usertests said:Having more of them could prevent tasks from spilling over to compute die(s). Every Nova Lake-S SKU will have 4c/4t of these. I think they should move up to 8c/8t in a future generation.
 8 cores and then you would juice them to keep up with the workload...at that point why bother the bigger cores are going to finish the work a lot more efficiently.
 
 Isn't the I/O die still on a pretty ancient node?!? Will the cores be on that old a node or are they going to an I/O die that is tiled to add the LP cores as a tile?usertests said:Zen 6 desktop CPUs will have up to 2x 12-core Zen 6 chiplets, and possibly 2-core Zen 6 LP on the I/O die. Playing the same role as LPE cores did on Meteor Lake APUs or upcoming Nova Lake-S desktop CPUs.
Both are possible but also both have pretty severe downsides, the first option won't give them much of a power saving and the second option is going to add a lot of cost.
- 
Reply
 It remains to be seen. But they can definitely use a cell library optimized for power efficiency.TerryLaze said:Turning off compute tiles only makes sense as long as the LP cluster doesn't use more power than they do...
 8 cores and then you would juice them to keep up with the workload...at that point why bother the bigger cores are going to finish the work a lot more efficiently.
 
 I/O/SoC tiles would never be deactivated. Max LPE clock speeds on Meteor Lake and Arrow Lake-H are 2.5 GHz (yeah, that's your turbo). They won't use much power.
 
 Whether or not this is a good idea, it's coming to every single Nova Lake-S desktop CPU as far as we know. And 4c/4t instead of the 2c/2t of earlier iterations. So there will be plenty of testing done with them.
 
 
 If Zen 6 includes them they should use whatever node the I/O uses. And I don't think it would be possible for AMD to make Zen 6 LP cores on TSMC N6.TerryLaze said:Isn't the I/O die still on a pretty ancient node?!? Will the cores be on that old a node or are they going to an I/O die that is tiled to add the LP cores as a tile?
 Both are possible but also both have pretty severe downsides, the first option won't give them much of a power saving and the second option is going to add a lot of cost.
 
 I think the node would have to be TSMC N3C or N3P, which is hardly ancient since Zen 5 CCDs use N4X. My reasoning is that basically all Zen 6 cores are going to be made on TSMC 3nm (e.g. budget APU) or 2nm (desktop, server) nodes.
 
 So I would rule out TSMC N4C unless no LP cores are included.
 
