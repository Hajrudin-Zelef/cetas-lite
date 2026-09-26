---
id: collect-250926-servers-hardware/servers-hardware/diving-deeper-on-nvidia-s-vera-cpu-new-architectural-details-and-spec-cpu-2026-benchmarks-2
title: "diving-deeper-on-nvidia-s-vera-cpu-new-architectural-details-and-spec-cpu-2026-benchmarks"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["nvidia", "agentic", "amd", "consumer", "decode", "disclosure", "fp8", "intel", "latency", "memory", "throughput"]
source: docs/RAG/clean4/diving-deeper-on-nvidia-s-vera-cpu-new-architectural-details-and-spec-cpu-2026-benchmarks.md
source_anchor: ""
source_lines: [49, 81]
sha256: 5c8d0046ef18cc384d1285e1561b6eef5a18e53bee8b01e921403634a4f6d17e
---

# diving-deeper-on-nvidia-s-vera-cpu-new-architectural-details-and-spec-cpu-2026-benchmarks

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
