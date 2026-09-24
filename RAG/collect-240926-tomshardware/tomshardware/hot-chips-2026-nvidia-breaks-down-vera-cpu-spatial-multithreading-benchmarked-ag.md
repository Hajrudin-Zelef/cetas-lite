---
id: collect-240926-tomshardware/tomshardware/hot-chips-2026-nvidia-breaks-down-vera-cpu-spatial-multithreading-benchmarked-ag
title: "hot-chips-2026-nvidia-breaks-down-vera-cpu-spatial-multithreading-benchmarked-ag"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Cohere", "Intel", "Nvidia", "SpaceX"]
dates: []
keywords: ["benchmark", "nvidia", "agent", "agentic", "agents", "amd", "compute", "consumer", "decode", "intel", "lpddr5x", "memory"]
source: docs/RAG/clean_en/tomshardware/hot-chips-2026-nvidia-breaks-down-vera-cpu-spatial-multithreading-benchmarked-ag.md
source_anchor: ""
source_lines: [1, 45]
sha256: 9887dc45254fd6e47ab7e6bf6de40b47b587e9721cdb944438b0c0acb4bf686e
---

# hot-chips-2026-nvidia-breaks-down-vera-cpu-spatial-multithreading-benchmarked-ag

<!-- source: https://www.tomshardware.com/pc-components/cpus/hot-chips-2026-nvidia-breaks-down-88-core-vera-cpu-spatial-multithreading-benchmarked-1-2-tb-s-socamm2-memory-agentic-workloads-detailed-and-more -->

Nvidia has spent the last several months providing key disclosures about its next-gen Vera CPU for agentic data centers, which it continued at Hot Chips 2026. Although we've already learned a lot about Vera, how it compares to AMD's next-gen Venice CPUs, and the inner workings of the Olympus core, Nvidia provided a bit more color at Hot Chips on spatial multithreading, the memory subsystem, and what types of workloads it's targeting with Vera.

As a quick refresher, Vera is the first CPU with a custom Nvidia core, following up on Grace, which used a stock Arm design. It's shipping as a single, 88-core SKU, and it has some key design differences compared to Nvidia's x86 competition, most notably a multi-threading implementation that Nvidia calls spatial multi-threading, an LPDDR5X memory subsystem, and a monolithic compute die rather than using compute chiplets.

Nvidia says it's designed Vera specifically for agentic AI workloads, a category that's still being defined in terms of performance benchmarking. Many CPU-intensive tasks serve as proxies for agentic workloads (i.e., code compilation), though measuring performance across a full agentic chain is complex and inconsistent. Nvidia, in its own slides (see the end of this article), calls agentic AI the "most complex computing workload in history," after all.

Nvidia provided an example of a headless browser to show the benefits of Vera, using optimized code to mimic how an agent would use a browser. Compared to the 96-core EPYC 9655P, Nvidia says Vera runs 24% faster as browser instances scale.

This slide is a good demonstration of the complexities in measuring traditional workloads and applying that performance to agentic workflows. Agents will often fetch websites for information, but there are several layers where agents can trim back compared to humans; in this case, agents can run through a browsing workflow 4.5x faster by cutting things like GUI rendering, fonts, media decoding, and more.

Another touchstone for agentic performance is code compilation, as agents seek out software to compile on the system. This might be the most direct benchmark of agentic AI performance with current workflows right now. Though, as previously mentioned, agentic chains are long, complex, and involve several different workloads.

Once again, compared to the 96-core EPYC 9655P, Nvidia claims Vera can compile the Linux kernel 22% faster with a native AArch64 target, and 14% faster when cross-compiling for x86.

## Nvidia's big cores for agentic AI — another look at Olympus and how it fits into Vera

Nvidia reiterated the importance of the large cores inside Vera, including the large BPU, neural branch predictor, and 10-wide decode. Nvidia has previously disclosed the Olympus core architecture, which you can read about in our Vera deep dive. Broadly speaking, however, it's a wide core optimized for high single-core throughput.

One of the more interesting design points of Vera is spatial multi-threading, which Nvidia described in more detail during its Hot Chips 2026. In short, Nvidia separates core resources on two pipelines, though data and cache can move between threads as needed. To demonstrate the benefit, Nvidia shared the results from SPEC CPU 2017 intrate that you can see below.

This shows the "noisy neighbor" effect. Nvidia measured single-core performance and then measured the same workload with another thread active. Nvidia's data shows that Vera is less concerned with the neighboring thread, whereas a "traditional CPU" sees a larger slowdown. Nvidia didn't clarify which CPU it's comparing Vera to here, however.

Nvidia's slide does a good job illustrating, but it's worth noting the difference compared to traditional SMT nonetheless. With traditional SMT, resources are time-sliced between threads, leading to gaps between BP and decode, as illustrated in the slide. With spatial multithreading in Vera, threads are still fighting for resources within the core. However, spatial multithreading allows Nvidia to deal with the demand of neighboring threads in a deterministic way, leading to a more consistent downturn in per-core performance when the second thread is working.

Nvidia's second-gen Scalable Coherency Fabric (SCF) moves data across the die. Nvidia didn't provide any new disclosures around SCF at Hot Chips, but you can see how the fabric is laid out in the slide below. Centralized Coherency Switch Nodes (CSNs) connect the cores to pools of L3 cache totaling 164 MB and the broader memory subsystem.

At a system level, one of the more interesting choices Nvidia made was to use LPDDR5X as opposed to traditional RDIMMs, a choice that it was only able to make due to the serviceable SOCAMM2 design. Nvidia includes eight SOCAMM2 slots per Vera CPU on a board, offering up to 1.5 TB of capacity with 1.2 TB/s of bandwidth.

LPDDR5X can deliver transfer rates higher than DDR5 RDIMMs, at least compared to single-rank DIMMs. However, it seems the driving force behind LPDDR5X wasn't performance but rather power consumption. One of the pillars of Vera, according to Nvidia's Hot Chips presentation, was to deliver a CPU for power-limited data centers. Micron says its LPDDR5X consumes about a third of the power compared to a traditional RDIMM.

Nvidia demonstrated that point a little differently, using bandwidth per watt as a point of comparison between the LPDDR5X system in Vera and traditional RDIMMs. This illustration does the job, though it could be a bit misleading, measuring power draw against peak bandwidth.

Nvidia tells us that a fully loaded memory system with Vera consumes between 30W and 40W, with 1.5 TB at 9600 MT/s. Power demands for RDIMMs vary wildly depending on capacity, channels, and transfer rate, though power consumption can easily climb over 100W depending on the configuration.

Although Nvidia has deployed Grace in the data center — to the tune of "hundreds of thousands" of standalone servers, apparently — Vera represents Nvidia's first big push to gobble up market share in the expanding agentic CPU market. It's highly targeted, as evidenced by the fact that Nvidia is only delivering a single 88-core SKU, and it's already being put to use in large-scale deployments, with Nvidia announcing yesterday a deployment of Vera at SpaceXAI.

Nvidia has shared the slide above before, which it once again showed at Hot Chips 2026. It's normalizing per-core performance in SPEC CPU 2026 against the AMD EPYC 9755. The core differences explain the big disparity in numbers; in reality, Vera led in overall score by 3%. Regardless, this is the slide Nvidia is using to pitch Vera, claiming it offers a big improvement in the workloads that are most relevant for agentic AI.

The most formidable opponent for Vera isn't Turin, however. It's Venice, which AMD launched in June, and Diamond Rapids, which Intel detailed just moments after Nvidia left the stage. Vera has a lot of interesting talking points already, but it'll be interesting to watch how Nvidia scales (or doesn't scale) its data center CPU business over the next few generations. Perhaps we'll see the firm double down on these agentic workflows, or maybe concessions and product segmentation to appeal to hyperscalers. Time will tell.

## Full Nvidia Vera Hot Chips 2026 presentation

Jake Roach is the Senior CPU Analyst at Tom’s Hardware, writing reviews, news, and features about the latest consumer and workstation processors.
