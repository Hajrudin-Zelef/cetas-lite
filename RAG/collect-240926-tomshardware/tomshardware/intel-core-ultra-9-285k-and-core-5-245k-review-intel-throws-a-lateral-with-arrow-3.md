---
id: collect-240926-tomshardware/tomshardware/intel-core-ultra-9-285k-and-core-5-245k-review-intel-throws-a-lateral-with-arrow-3
title: "intel-core-ultra-9-285k-and-core-5-245k-review-intel-throws-a-lateral-with-arrow"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "TSMC"]
dates: []
keywords: ["intel", "amd", "benchmarks", "chiplet", "latency", "memory", "pricing"]
source: docs/RAG/clean_en/tomshardware/intel-core-ultra-9-285k-and-core-5-245k-review-intel-throws-a-lateral-with-arrow.md
source_anchor: ""
source_lines: [115, 164]
sha256: 0b3e0e9d6c4a94b3be6a57068230e989e8b6e8e92df4e8d7b9943c7caf0468c8
---

# intel-core-ultra-9-285k-and-core-5-245k-review-intel-throws-a-lateral-with-arrow

| AIDA L3 Cache Latency Measurements |  |  |  | 
|---|---|---|---|
| Memory Latency - Tom's Hardware | DDR5-5600 | CUDIMM DDR5-6400 | L3 Latency | 
| Core Ultra 9 285K | 94.1 ns | 91.9 ns | 16.6 / 15.8 ns | 
| Core i9-14900K | 79.1 ns | N/A | 21.8 ns | 

As you can see in the AIDA tests, on a like-for-like basis with DDR5-5600, the remote memory controller and PHY add 15 ns of memory latency in our test. Intel says we can expect a 15 to 20 ns increase in memory latency for Arrow Lake. That's a pretty significant change, and not in a good way. It will definitely impact certain workloads, gaming in particular.

Intel redesigned its CPU core layout for Arrow Lake, with quad-core E-core clusters interspersed among the P-cores. Previously, Intel placed all the E-cores in their own dedicated block. Intel spread the cores out to reduce hotspots for this design.

Intel also connected the E-cores to the 36MB L3 cache, so they now share L3 with the P-cores for the first time. The P-cores and E-cores still have dedicated L2 caches, with 3MB for the P-cores (a .5MB increase over the prior gen) and 4MB of L2 shared among each E-core cluster. Intel says the new design yielded a 33% reduction in package size and allowed it to quickly port innovations from Lunar Lake to Arrow, thus allowing it to launch the two chips a mere month apart.

The album below contains a latency heatmap for both the new design and the previous-gen Core 9-14900K, showing that the removal of Hyper-Threading has greatly simplified the core-to-core traffic due to the reduced number of threads.

We’ve also included the slides detailing the advances of the P-core and E-core architectures, but we’ve already covered these microarchitectures in-depth in our Lunar Lake deep dive. Overall, Intel claims a 9% increase in IPC over Raptor Lake (lower than the 14% cited with Lunar Lake because Intel made that comparison to Meteor Lake), a 32% increase in integer IPC, and a 72% increase in floating point IPC for the E-cores.

But claims and architectural details are only the foundation. Let's see how everything looks like in actual real-world performance testing and benchmarks on the following pages.

Current page: Chiplet-based design with TSMC nodes come to Intel

Next Page Intel Core Ultra 9 285K Gaming Benchmarks
- 
Honestly if they brought the pricing to be closer to current 14th gen pricing, these would be a good option if power draw is a concern. I dont think the gaming performance drops are as bad as i thought they would be, and honestly the gaming performance is fine for %99 of poeple. Hopefuly we see gaming improvments in the next gen without power draw increasing. I like the direction their going in, but i dont think this is really a compelling upgrade for anyone 12th gen and up.Reply
- 
Talk about "hold my beer" moments... Holy cow. I thought it was going to be bad, but notReply*THIS* bad.
 
 I hope Intel irons out all the reported and shown issues in multiple reviews and get it to a better place, but as an initial showing, makes Zen5 a friggen home run.
 
 And thanks a lot Paul. Great data as always and I'll definitely check later when the missing bits and bobs are added :D
 
Regards.
- 
Got one 14700t for 283usd what I see will have a little less performance than this core ultra with less power wasted. The T family's aways stuck at 65w power max when set the pl1 to max allowed power.Reply
Will wait till ddr6 and pcie 6 droops on desktop yo upgrade.
- 
Thanks for the write-up, Paul! IMO, Intel did what AMD did with their 9xxx series: they laid the foundations for future gens, with workstation and server workloads sped up first.Reply
 
 "Sorry gamers, tough luck this round."
 
 I like this: "We also can't help but wonder how a future Intel CPU that incorporates a cache chiplet — similar to AMD's X3D line — might change the picture."
I think we've seen examples how a large cache, when wisely implemented, can benefit workloads from both AMD and Intel before.
- 
Reply
 computerbase.de clearly made a mistake. those numbers don't match anyone else.TheHerald said:You sure about that?
 
 From computerbase de
 
 
 
They probably are only measuring the power draw at the 8 pin cpu power plug. apparently this chip draws a lot of it's power from the 24pin, something you wouldn't notice unless you measured system power draw everywhere.
