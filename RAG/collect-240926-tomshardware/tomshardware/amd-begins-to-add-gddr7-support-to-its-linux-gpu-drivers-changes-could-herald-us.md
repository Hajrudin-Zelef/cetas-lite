---
id: collect-240926-tomshardware/tomshardware/amd-begins-to-add-gddr7-support-to-its-linux-gpu-drivers-changes-could-herald-us
title: "amd-begins-to-add-gddr7-support-to-its-linux-gpu-drivers-changes-could-herald-us"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Nvidia"]
dates: ["2025-08"]
keywords: ["amd", "gpu", "chiplet", "compute", "consumer", "cost", "datacenter", "disaggregated", "gpus", "latency", "memory", "nvidia"]
source: docs/RAG/clean_en/tomshardware/amd-begins-to-add-gddr7-support-to-its-linux-gpu-drivers-changes-could-herald-us.md
source_anchor: ""
source_lines: [1, 36]
sha256: add8cc00a8f6fdc92c1aaa56d55ff4cc04e423b0574ad1493c21dc7d13c3c75c
---

# amd-begins-to-add-gddr7-support-to-its-linux-gpu-drivers-changes-could-herald-us

<!-- source: https://www.tomshardware.com/pc-components/gpus/amd-begins-to-add-gddr7-support-to-its-linux-gpu-drivers-changes-could-herald-use-of-advanced-memory-standard-with-next-gen-radeon-gpus -->

AMD has started to add support for GDDR7 memory and several new graphics IP blocks to its open-source Linux kernel driver, which indicates that software enablement for the company's next-generation standalone GPUs is underway, reports Phoronix. While AMD does not identify the upcoming architecture, the changes are likely tied to its RDNA 5 project — but they don't necessarily herald an imminent launch.

AMD's existing Radeon RX 9000-series graphics processors based on the RDNA 4 architecture use GDDR6 memory, so the addition of GDDR7 identification to drivers is arguably the most explicit confirmation that AMD is setting the stage for its next-generation discrete Radeon graphics processors based on the RDNA 5 architecture.

AMD also submitted patches that enable IH 8.0, a new version of its Interrupt Handler IP block, as well as NBIF 7.10, the latest revision of the company's New Bus Interface. Several smaller patches make additional preparations for the new hardware. These changes follow earlier Linux driver work involving Display Core Next 6 (DCN6) as well as GFX 13.0.x, suggesting that AMD is gradually upstreaming support for multiple components of an upcoming GPU architecture. Yet, while enablement of IH8, NBIF 7.10, DCN6, and GFX 13.0.x clearly point to new graphics hardware, it does not necessarily point to new discrete GPUs, unlike the mention of GDDR7.

As revealed in August, 2025, Laks Pappu, Senior Fellow at AMD, is the lead architect for AMD's next generation datacenter GPU and discrete graphics platforms. Pappu was building next-generation "competitive 2.5D/3.5D chiplet-based and monolithic graphics SoCs on various packaging technologies," according to his LinkedIn profile before it was edited to remove those disclosures.

Essentially, Pappu's profile indicated that AMD's next-generation GPU architecture could support multiple physical implementations — including multi-chiplet and monolithic — depending on performance, cost, market requirements, and AMD's willingness to compete in certain market segments

AMD has already used a multi-chiplet design with its Navi 31 GPU, an implementation that kept graphics processing hardware on a large Graphics Compute Die (GCD), but disaggregated memory interfaces and caches into smaller Memory Cache Dies (MCDs). A more ambitious implementation could potentially distribute graphics processing resources between multiple dies, although Pappu's profile did not disclose how AMD intended to partition its future GPUs.

Such an approach would be considerably more complicated than separating memory and cache functionality. Multiple compute dies would require high-bandwidth, low-latency connections as well as mechanisms for synchronization and coherency, while software would ideally continue to see the set of chiplets as a single graphics processor.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Based on conventional GPU development cycles of roughly 2.5 to 3.5 years, RDNA 5 could already have been approaching tape-out or early post-tape-out stages when the information surfaced in August 2025, which means that the company is already likely to be testing the new GPUs internally.

Officially, AMD hasn't disclosed any details about its RDNA 5 architecture nor launch schedule, so take these developments with a grain of salt. For now, all we know for sure is that AMD is working on standalone RDNA 5-based GPUs, and they're likely to use GDDR7 in at least some configurations, assuming the AI-driven memory crunch doesn't choke supply of those chips even further for consumer applications.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Anton Shilov is a contributing writer at Tom’s Hardware. Over the past couple of decades, he has covered everything from CPUs and GPUs to supercomputers and from modern process technologies and latest fab tools to high-tech industry trends.

- 
I can see GDDR7 being needed for stronger Radeon AI Pro GPU's (and the customer base there that has the coin to afford it), but will it make it into gaming-class cards? I'm hoping so for at least a new halo card and halo -1, but unfortunately only the rich & famous will be able to afford those bad boys. AMD doesn't have a lot of incentive to make lower margins on essentially a reworked beefy AI GPU with less VRAM, higher clocks, etc. at lower prices, but we'll see what hand they play when the time comes.Reply
- 
Reply
 https://www.digitaltrends.com/computing/amds-rdna-5-might-deliver-what-gamers-want-most-more-vram/DS426 said:I can see GDDR7 being needed for stronger Radeon AI Pro GPU's (and the customer base there that has the coin to afford it), but will it make it into gaming-class cards? I'm hoping so for at least a new halo card and halo -1, but unfortunately only the rich & famous will be able to afford those bad boys. AMD doesn't have a lot of incentive to make lower margins on essentially a reworked beefy AI GPU with less VRAM, higher clocks, etc. at lower prices, but we'll see what hand they play when the time comes.
 We've seen a leaked document showing what cards they were considering 1-2 years ago, all using 3 GB GDDR7, at 36 Gbps instead of the 28-30 Gbps Nvidia uses now. With that bandwidth on the table, it probably doesn't make sense to use GDDR6 anymore, and 2 GB GDDR7 production could be ending. The end of 2 GB GDDR7 production should help 3 GB GDDR7 module prices to fall to reasonable levels.
 
 3 GB GDDR7 can enable some interesting moves for the "low-end". An obvious one would be a 128-bit 12 GB card. It ends the 8 GB dilemma and provides enough VRAM for 1080p.
 
The document also shows 160-bit 15 GB, which could be a good compromise below 192-bit 18 GB. A 15 GB RDNA5 card with around 2.25x the bandwidth ((36/20.1)*(160/128)) will crush the 9060 XT 16 GB, and nobody will miss that 1 GB.
