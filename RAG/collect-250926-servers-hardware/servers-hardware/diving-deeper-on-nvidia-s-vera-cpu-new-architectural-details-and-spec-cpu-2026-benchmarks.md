---
id: collect-250926-servers-hardware/servers-hardware/diving-deeper-on-nvidia-s-vera-cpu-new-architectural-details-and-spec-cpu-2026-benchmarks
title: "diving-deeper-on-nvidia-s-vera-cpu-new-architectural-details-and-spec-cpu-2026-benchmarks"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "benchmarks", "nvidia", "agentic", "amd", "consumer", "decode", "disclosure", "fp8", "gpu", "gpus", "intel"]
source: docs/RAG/clean4/diving-deeper-on-nvidia-s-vera-cpu-new-architectural-details-and-spec-cpu-2026-benchmarks.md
source_anchor: ""
source_lines: [1, 81]
sha256: 7c2c530b8ce5c698d97b970b0c0bcb93e5c0a4b04429fca5b4fea62d494f0b24
---

# diving-deeper-on-nvidia-s-vera-cpu-new-architectural-details-and-spec-cpu-2026-benchmarks

This week is quickly shaping up to be a busy week in the world of server hardware. Later this week earnings season for the industry will kick off, starting, as always with Intel. Meanwhile AMD’s annual enterprise an AI event, Advancing AI, is taking place in San Francisco mid-week.

But before either of those chip vendors get to present, the first salvo for this week comes from NVIDIA. This morning the company is releasing a significant amount of new technical information on its upcoming Vera CPUs and Rubin GPUs, offering the deepest look yet into the products since they were “launched” at the start of this year.

Among NVIDIA’s publications this morning includes a long-awaited tech blog on the GPU architecture and additional details on the NVLink interface that will be connecting those GPUs (and CPUs) together. But the most interesting disclosure is arguably around Vera, NVIDIA’s new CPU, which is getting a full architectural whitepaper release. Though it does not answer all of our questions about the chip or the Olympus CPU cores inside, it is our best look yet at NVIDIA’s first custom CPU core design in years.

At the same time, today’s revelations come with some fresh performance figures for Vera, most notably performance benchmarks on the industry-standard SPEC CPU 2026 benchmark suite. The performance figures certainly paint a vivid picture that outlines just what aspects of CPU design that NVIDIA has focused on. Though at the same time the benchmarks give us our first real look at the CPU’s overall performance – and it is not all rosy for NVIDIA. But before we get too far ahead of ourselves here, let’s take a quick look at the new architectural details around Vera.


## NVIDIA Vera Exposed: Olympus CPU Architecture

Back around the time of GTC 2026 we had a chance to dig a bit deeper into NVIDIA’s Vera CPU, which in many respects was the star of the show for NVIDIA. While the company makes the bulk of its fortune off of GPUs in various forms, NVIDIA is making a hard press to expand its reach into the server CPU market with Vera. While the chip will not be going toe-to-toe with the links of Intel’s Xeon or AMD’s EPYC lineups across their full footprints, NVIDIA wants to elevate Vera towards being a chip that is used for many kinds of servers, and not just as a CPU that is used to drive their GPUs, as was the case with Grace.

Of the full Vera Rubin product stack, this makes Vera the most interesting development at NVIDIA. Not only is it their first custom CPU architecture in many years, but it means that NVIDIA is diving headfirst into a far more competitive market than GPUs or even networking gear (both markets where NVIDIA is now the world’s largest vendor). Server CPUs, by comparison, is a very large and very lucrative market where NVIDIA is still trying to climb the ladder. The company is hardly a newcomer to producing CPUs, but it is going to be the hardest market to grow; between Intel, AMD, and a slew of Arm-based CPUs, there is no shortage of capable CPUs.

Which is why for NVIDIA there is a strong desire to put their best foot forward on Vera. NVIDIA wants to win customers for Vera – as well as dissuading customers from pairing up Rubin GPUs with someone else’s CPU – and as a result it is the product that is getting the most support from NVIDIA, be it in terms of promotion or architectural disclosures.

| **Feature** | **Grace** | **Vera** | 
| **CPU Architecture** | Arm Neoverse V2 | NVIDIA Olympus | 
| **Cores** | 72 | 88 | 
| **Threads** | 72 | 176 (Spatial Multithreading) | 
| **L2 Cache per core** | 1MB | 2MB | 
| **Unified L3 Cache** | 114MB | 164MB | 
| **Memory bandwidth (BW)** | Up to 512GB/s | Up to 1.2TB/s | 
| **Memory capacity** | Up to 480GB LPDDR5X | Up to 1.5TB LPDDR5X | 
| **SIMD** | 4x 128b SVE2 | 6x 128b SVE2 FP8 | 
| **NVLINK-C2C** | 900GB/s | 1.8TB/s | 
| **PCIe/CXL** | Gen5 | Gen6/CXL 3.1 | 
| **TDP** | 250W | 250-450W | 

Diving right in, when we first looked at Vera earlier this year, NVIDA had already disclosed a great deal about its high-level specifications, including memory configurations, core counts, the front-end decoder width, and the L2 and L3 cache amounts for each core and the complete chip, respectively. With NVIDIA’s newest architectural disclosure, we have some further information to fill in the blanks.

At a high level, Olympus looks a lot like other contemporary high-end CPU cores. These logical diagrams are designed to be readable more than they are detailed, so they are far from the final word in analyzing a CPU architecture. But it is fair to say the Olympus architecture is not wildly different from other high-end CPUs; this is not Denver, and rather than thinking outside of the box NVIDIA’s overall direction is to build a better CPU that follows contemporary CPU design paradigms.

Promoted by NVIDIA early-on, one of the key design aspects for the chip is its high front-end performance. One of the better documented aspects of the chip at the time, NVIDIA disclosed that it used a neural branch predictor that could spit out 2 branches per cycle with high accuracy, with the goal to further reduce branch mispredictions and the performance stalls that incurs.

NVIDIA’s latest disclosure unfortunately does not go into any further detail on the branch predictor, but it does give us our first look at the instruction fetch unit it feeds, as well as the path into the 10-wide decoder. In short, Olympus’s instruction fetch unit can feed as many as 16 instructions – 64 bits each – into the decode queue. The queue can hold 48 instructions altogether, and can spit out up to 10 fused instructions to be consumed by the actual decoder.

We also have our first confirmation on the L1 cache sizes on Olympus. Each CPU core has a 64KB, 4-way L1 instruction cache working in the core’s front-end. Notably, this is two-times the size of Zen 5’s L1 cache (and matching Cougar Cove’s), underscoring NVIDIA’s intent to better handle large instruction footprints. Meanwhile that is joined by a 96KB, 6-way L1 data cache working in the back-end.

Moving on to the mid-core, as NVIDIA calls it, we have a traditional out-of-order execution engine. NVIDIA is not disclosing the size of the all-important reorder buffer here. Nor are they disclosing the size of the register files used for register renaming. But the company is confirming that the CPU core can rename up to 10 micro-ops per cycle, keeping up with the decoder.

Absent specific specifications, NVIDIA is offering some additional details on some of the techniques they are using to identify instructions to reorder in order to increase the amount of ILP they can extract from an instruction stream. Of note, NVIDIA is employing memory renaming here, which is a technique we also see on recent AMD designs. Effectively a superset of register renaming, memory renaming takes things one step further to rename whole memory addresses to break up what would normally be dependent instructions. As NVIDIA outlines in their whitepaper, this allows Olympus to execute dependent instructions before a load even completes, so long as the data relationship can be predicted or determined.

Olympus also employs value prediction, which allows for a further form of speculative execution by identifying repetitive patterns and guessing the values before they have even been generates (this is of course verified). Finally, the Olympus mid-core implements move elimination to catch and work-around unnecessary memory value movements.

Finally, we have the backend of Olympus. At the time of our March preview, this was one of the least well documented aspects of the CPU core, with NVIDIA’s software optimization guide indicating that it was even more massive than their front-end. This morning’s disclosure finally confirms just how wide it is, outlining that Olympus is indeed a rather sizeable CPU.

As illustrated in NVIDIA’s Vera whitepaper, each Olympus CPU core features 18 execution pipes in total. 8 of these pipes are for integer arithmetic operations and branching, 6 more are for vector/FP operations, and then finally there are 4 pipes for the load units and another 2 for the store units. Notably here, all of these pipes are in pairwise configurations – there is a multiple of 2 of everything – which is an important element of how NVIDIA’s partition-based spatial multithreading works.

For the ALU block, all 8 pipes are able to handle simple operations (add/subtract/logical). Meanwhile 2 of these pipes are complex pipes that can also handle complex arithmetic (multiply/divide) as well as more specialized operations such as CRCs and bit shifting. NVIDIA’s whitepaper does not go into instruction latency, but we know from their existing software operation guide that basic ALU operations complete in a single cycle, while multiplies can be done in as little are 2 cycles, and divides are variable at anywhere between 5 and 20 cycles depending on the number of iterations required.

It is notable here that Olympus has more pipes for integer operations than it does floating-point. Which is not to say that it is a slouch in FP operations, but that is going to be contingent on being able to fill those SVE2 units with whole vector’s worth of operations.

Meanwhile the 4 branch units are evenly distributed between the simple and complex ALU pipes.

As for floating point and SIMD operations, the 6 vector pipes largely reiterates NVIDIA’s earlier claims about the processor. Each Olympus core can process up to 6 SVE2 vector instructions at once (up from 4 on Grace), with Olympus also adding support for FP8 precision for low-precision math (as is all the rage these days for AI workloads). Alternatively, 2 of the pipes can consume cryptographic vector operations, which would be used for processing and computing AES, SHA, and other encrypted data/cryptographic hashes.

Moving on, we have the 4 pipes for the load/store units. 2 of these pipes are shared, supporting both loads and stores, while the last 2 pipes are dedicated solely to loads. This is a similar load-heavy design as other architectures. Though the fact that Olympus does not explicitly have dedicated pipes for address generation does set it apart.

Finally, feeding the whole chip is NVIDIA’s cache subsystem that sits between the CPU core and the larger memory subsystem. The large TLBs, including the 3K entry second-level TLB, are a fairly common sight for high-end CPUs. On the other hand, NVIDIA is taking a great deal of pride in their prefetcher, which earned its own section in the Vera whitepaper.

In particular, NVIDIA has implemented a “novel” graph prefetcher, which is designed to accelerate graph-like data structures. The full details go beyond the scope of this short article, but the graph prefetcher is meant to improve on handling of pointer-heavy (and indirection-heavy) memory patterns, which are common to graphs. In those cases, the Olympus core attempts to identify the relationship between a pointer and its resulting destination (a producer-consumer relationship, as NVIDIA calls it), and then prefetch the data from the final target memory address before the pointer is actually resolved. To the best of my knowledge, neither AMD nor Intel have implemented a similar prefetcher, though the idea has been floating around academia for at least several years.

It looks to me like someone set the colour scale on the core to core latency table so the entire table was green. Even so, the chip seems well designed for tightly coupled parallel workloads. Since AI is different than the cloud-style micro-service throughput targeted by existing processors, I think it makes sense for Nvidia to fabricate their own.

The first diagram shows only 16 lanes of PCIe 6. Surely that’s not right. How could a server CPU in 2026 ship with fewer lanes than a desktop CPU, even if they are a newer generation?

It should be noted that if Vera had been compared to the 9575F for the single-core comparisons, which you suggested would have been more appropriate, then Vera would have held a sizable advantage in the estimated SPECrate2026_int_base score. More importantly, Nvidia doesn’t seem to intend this CPU to be a general purpose CPU to go head-to-head with established processors in most data center workloads. They are targeting it for AI servers, and in particular agentic AI servers. Given how much Vera is ahead in various single core performance metrics, and considering the normal generation-on-generation IPC uplift, and that other CPU manufacturers are unlikely to have engineered a special single core monster for this upcoming generation, it seems very likely that Vera will keep an advantage for these single core performance workloads even against the coming generation of server chips. The whole reason the future CPU server TAM has recently been revised sharply upward is because of the agentic AI workload. Nvidia’s Vera CPU will live and die by how it performs against its coming-generation CPU competition in agentic AI workloads.

The microarchitecture block diagram contains multiple errors (LLM-esque). If it comes directly from nvidia (judging by the color scheme), then one has to question the accuracy of other information provided…

NVIDIA’s latest disclosure unfortunately does not go into any further detail on the branch predictor, but it does give us our first look at the instruction fetch unit it feeds, as well as the path into the 10-wide decoder. In short, Olympus’s instruction fetch unit can feed as many as 16 instructions – 64 bits each – into the decode queue. The queue can hold 48 instructions altogether, and can spit out up to 10 fused instructions to be consumed by the actual decoder.

I thought ARM ISA was 4-bytes (32-bits instruction, fixed length).

add x14,x15,x16 (4-bytes)

mul w3,w2,w4 (4-bytes)

Is that a custom ARM ISA that Nvidia is using?
