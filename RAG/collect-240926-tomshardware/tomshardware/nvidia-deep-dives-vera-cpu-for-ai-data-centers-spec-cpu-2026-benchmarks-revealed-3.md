---
id: collect-240926-tomshardware/tomshardware/nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed-3
title: "nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "AWS", "Cohere", "Intel", "Nvidia"]
dates: []
keywords: ["benchmarks", "nvidia", "amd", "aws", "chiplet", "decode", "gpu", "gpus", "graviton", "intel", "latency", "memory"]
source: docs/RAG/clean_en/tomshardware/nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed.md
source_anchor: ""
source_lines: [101, 152]
sha256: 40e9869442bc68208e42036728acba37a5b5660ed558cfc7d72891919d1b40f9
---

# nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed

 BTW, there are two other obvious points of comparison:
 The Neoverse V2 cores, used in Nvidia's prior Grace CPUs, have a 6-wide decoder* (**source:** https://chipsandcheese.com/p/hot-chips-2023-arms-neoverse-v2 )
 The Cortex-X925 cores, used in Nvidia's RTX Spark, have a 10-wide decoder (**source:** https://chipsandcheese.com/p/arms-cortex-x925-reaching-desktop )
 * Note that Neoverse V2 still has a mOP cache with 8-wide dispatch. So, as with x86 P-cores, the decoder width is a little bit deceptive.
 
 
 You skipped a pretty big detail:The article said:Past the front end, the mid-core rename / allocation engine is built to keep instructions moving while waiting on dependencies.**no mOP cache!**
 Modern x86 P-cores and some 64-bit ARM cores had micro-op caches to avoid having to re-decode the same instructions. Once ARM dropped 32-bit compatibility, they started getting rid of those. Also, Intel either doesn't have them in their E-cores, or perhaps what they did was to put some of that into the I-cache.
 
 So, it's interesting (but not all that surprising) that Nvidia followed ARM's approach of just skipping the mOP cache, entirely.
 
 
 It's darkly ironic that SVE's main selling point was to allow CPUs to scale all the way up to 2048 bits per vector, but all mainstream implementations (except for AWS Graviton 3) are just 128-bit.The article said:For SIMD instructions, the execution engine includes a vector cluster for Arm’s Scalable Vector Extension (SVE), including six vector units that support 128-bit SVE instructions
 
 
 Not really. The way AMD describes SMT in Zen 5 is that certain competitively-shared resources have watermarks that limit how much a single thread is able to use, so that it doesn't starve out the other thread. However, when there's only one thread running on a core, more of those constraints go away and only the statically-partitioned resources (e.g. the decoder) remain exclusive.The article said:Traditional SMT time-slices execution, giving both threads access to all of the core resources and sharing them as instructions execute in parallel.
 
 My read on "Spatial Multithreading" is that it's Nvidia's marking machine trying to spin a weakness to make it sound more like a strength. This being Nvidia's first SMT implementation (AFAIK), it won't have the same sophistication as where Intel and AMD have gotten, over a couple decades of experience implementing and refining theirs.
 
 
 First, Intel has traditionally done the same thing. However, I need to catch up on what they've said about Sierra Forest.The article said:Nvidia’s second-generation Scalable Coherency Fabric (SCF). It underpins Nvidia’s approach of using a monolithic die as opposed to a chiplet-based design, distributing last level cache in a mesh across the die and avoiding the cross-CCD latency penalty with localized L3.
 
 Second, people tend to overestimate the impact of the cross-CCD thing. Unlike Intel CPUs, one AMD CCD won't write data to another's L3 slice. Each is private to that CCD, except for coherency. That said, when you've got multiple threads that are either exchanging or both modifying the same data, and they happen to be scheduled on different CCDs,*that* is when you feel the impact of the die-to-die communication.
 
 Anyway, AMD's approach has clearly scaled better. I think that's why Sierra Forest looks like it moved in the direction of doing the same thing with segmenting its L3 cache.
 
 
 I don't know what happened here, but the slide with the annotated die photo shows it supporting PCIe 6 with only 16 lanes.The article said:For I/O, Vera supports PCIe 6.4 with 88 lanes per CPU and bifurcation support down to x2.
 
 **P.S.** As for the benchmarks, the initial Phoronix review was interesting, but now I'm just waiting for some independent analysis. From a CPU microarchitecture standpoint, what I really want to know is the single-threaded performance across the whole SPEC suite, and how the sub-scores compare with other CPUs. MT scaling is mildly interesting, but it really stacked that deck in its favor by having relatively few cores and huge amounts of memory bandwidth. So, I expect it scales well, but not well enough to outright beat Intel or AMD on anything that's not fundamentally bottlenecked by memory bandwidth.
- 
Reply
 Announced? Venice isusertests said:But their comparisons (obviously) are to Turin rather than newly announced Venice.*launching* in about a week! That's probably why Nvidia is making a bunch of noise about this*now!*
- 
I like NVIDIAs marketing department scaling results per core. This is the height of aerobatics. Interesting how all would compare per core with the AMD 64core Turin 9575F, their king of single core benchmarks ? And how per core comparison will look for just the 64 core Venice? Huang will leave full of tears on leather jacketReply
- 
Reply
 Good question. Where Vera has a big lead over AMD's Turin is that it has basically 2x the memory bandwidth. Also, ARM uses memory bandwidth a little more efficiently, due to its relaxed memory consistency model.Stomx said:I like NVIDIAs marketing department scaling results per core. The height of aerobatics. Interesting how all would compare per core with the AMD 64core Turin 9575F, their king of single core benchmarks ?
 
 Where AMD would win biggest is going to be on AVX-512 stuff that's not memory-bound.
 
 Otherwise, it looks to me like Vera might have a slight edge over Zen 5. I mean it's a 10-way dispatch core vs. an 8-way one. Even Zen 6 will remain 8-way, if the rumors are correct. Not to say dispatch width is everything, since Intel's cores are also wider, but it's at least a good starting point. I think Zen 5 does have 3 integer multiply ports vs. Vera's 2.
 
 
Yeah. With Venice, AMD is going to leap-frog Nvidia's memory bandwidth by approximately 33%, according to what they've said. That's to be accomplished using 16 DIMMs (matching the same data width as Vera of 1024 bits) using MRDDR5-12800. So, on low core-count models, they'd stomp Nvidia on per-core memory bandwidth.Stomx said:And how per core comparison will look for just the 64 core Venice? Huang will leave full of tears on leather jacket
- 
Reply
 I think that explains why it doesn't have a lot of cores, but it doesn't explain why Nvidia went to the trouble of designing their own cores instead of just licensing the next Neoverse V-series core, like they did in their Grace CPU and like Amazon did since Graviton 3.palladin9479 said:Performance isn't really important, it's only real job is to manage the workloads for the GPU like devices connected to it. It only needs to be fast enough to not make the GPUs wait around.
 
 Arm's Neoverse V3 seems to be derived from their Cortex-X4, which is also a 10-way core (source: https://www.androidauthority.com/arm-cortex-x4-explained-3328008/ ), though it has only 4x SVE2 pipes. However, I think the X4 and V3 are probably more efficiency-focused and don't clock very high. Graviton 5 uses Neoverse-V3 at 3.3 GHz, for instance.
 
So, it seems like the main things it seems Nvidia got by making its own core was the ability to achieve higher single-thread perf by hitting higher clock speeds and the ability to add SMT for better utilization.
