---
id: collect-240926-tomshardware/tomshardware/amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c-1
title: "amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "AWS", "Google", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "nvidia", "agent", "agentic", "aws", "consumer", "gpu", "gpus", "intel", "memory"]
source: docs/RAG/clean_en/tomshardware/amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c.md
source_anchor: ""
source_lines: [1, 46]
sha256: ec7e57709ebae94710c0c873b3f5ceeb5905aff0d4939bb8bb22478decc8e068
---

# amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c

<!-- source: https://www.tomshardware.com/pc-components/cpus/amd-shares-first-official-benchmarks-for-epyc-venice-cpus-targets-nvidia-company-claims-256-core-chip-is-more-than-twice-as-fast-as-nvidia-vera-96-core-model-20-percent-faster-per-core -->

Following the launch of AMD's EPYC 'Venice' CPUs in July, AMD extended the performance claims for its upcoming generation of server chips on Friday. The high-level claim hasn't changed. AMD still says a 96-core, high-frequency Venice chip is around 20% faster than Nvidia's 88-core Vera in SPEC CPU 2026's Integer Rate test. However, the company went into far greater detail about the benchmarks in a new white paper.

There are several configuration differences depending on the benchmark throughout AMD's white paper, and although we'll call out those differences here to the best of our ability, we don't have all of the details. For the Vera comparison, in particular, AMD is mixing data from different sources, and in some cases, using different major releases of the GNU Compiler Collection (GCC). That can have a substantial impact on performance, so keep your salt shaker handy.

First up are results in SPEC CPU 2026 with the intrate test, looking at total throughput. These are older numbers, gathered in July with GCC 15.2. The intrate test runs multiple copies of an application on the same CPU, and the SOP is to run one copy per thread. Presumably, that's what AMD did here, but the white paper doesn't clarify, even in the footnotes.

The 256-core 9996 is 2.37x faster than the Intel Xeon 6980P and 2.24x faster than Vera according to the slide. The white paper clarifies the mystery 9006 CPU is the 256-core flagship. Perhaps most impressive is AMD's gen-on-gen comparison. According to these results, the 9996 is around 78% faster than last-gen's 192-core EPYC 9965.

Although the high-level results bring in data from Intel and AWS, much of the white paper focused squarely on the comparison between Venice and Vera. AMD broke down the individual subtests of SPEC CPU 2026 intrate in the white paper, which you can see below.

The comparison looks good for AMD, naturally, though there are a few wrinkles in the configuration. AMD is testing a down-cored EPYC 9996, dropping from 256 cores to 96 cores. It made no mention of power budget, but when AMD originally shared SPEC numbers, the 96-core model had access to the same 600W as the 256-core model — AMD's 96-core, high-frequency Venice SKU tops out at 500W. More consequential is the compiler, however. AMD is using GCC 16.1 and comparing the results to the ones Nvidia shared in its Vera white paper. Nvidia used GCC 15.2.

Michael Larabel over at *Phoronix* has a nice write-up about the difference between GCC 15 and 16, but the short story is that there are performance differences, not always for the better. GCC 16 takes longer to compile due to better optimizations, hence the lower scores on the GCC and LLVM compilations above. However, that leads to faster binaries. By how much depends on the flags, software, and a whole host of other factors. Regardless, it's not best practice to compare benchmarks using two different compiler versions. It makes sense that AMD used GCC 16.1 — it includes support for Zen 6 — but ideally Vera would also be on GCC 16.1.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Speaking of *Phoronix,* AMD pulled some data for the publication's initial, controlled testing of Vera. Above, you can see the Stream, an industry-standard benchmark for measuring memory bandwidth. Again, AMD is using a down-cored 9996 from 256 cores to 96, and offering it a 600W power budget. Still, this is an impressive showing, as Vera absolutely clobbered the competition in the publication’s original Stream results. Here, AMD is ahead by about 18%, with per-core performance about 8% ahead.

Breaking out of Vera, AMD also showed performance in cloud workloads, including database, Java, and cryptography. Once again, the gen-on-gen comparison stands out, as AMD was already leading in these workloads with its last-gen chips. AMD ran these tests itself, rather than relying on third-party data, though the Graviton5 results came from an AWS cloud instance.

Similarly, in HPC workloads, AMD furthers its lead over Intel's flagship Granite Rapids-AP offering. Intel's next-gen data center CPUs, codenamed Diamond Rapids, are set to be released next year.

Finally, we have "agentic AI workload performance," which uses actual benchmarks for comparison, despite what the names in the chart above suggest. From left to right, AMD used NGINX, TPCx-AI kit, FAISS, TPC-H and TPC-C, and a replay of a multi-persona agent. For TPC-H and TPC-C, AMD says it derived workloads from those benchmarks, so the results here aren't comparable to published results.

Although looking at benchmark results is always interesting, it doesn't say much in the context of a server deployment, at least at the scale that AMD is targeting. Peak performance is only one of the major factors that go into server deployments, after all, and even then, performance can vary wildly depending on what software you're running and how it's built.

Still, Venice looks impressive, perhaps more so in the gen-on-gen comparison than any competitive comparison. Hopefully that bodes well for AMD's future Zen 6 rollout on consumer desktops, but we'll have to wait until Team Red has more to share before drawing any conclusions on that front.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

- 
This fat toy is for the future world "elite" and their personal AI datacenters.Reply
 
 As for consumers — I honestly don't know who this processor is even for. Maybe it will end up in servers. Maybe it will run local AI on the CPU: with 2.5x the RDRAM bandwidth of AMD Turin, it could reach ~20 tokens per second on decent 4-bit quantized trillion-parameter models — though even 20 cores would be enough for that...
 
 But there is one piece of bad news for consumers: the massive socket size and the larger RAM footprint will most probably kill dual-socket designs within the old extended-ATX form factor. A dual-socket motherboard, if one ever gets built, will look monstrous (just take a look at the first single-socket example, Gigabyte's MD25-CG0) and won't fit into any existing cases. That's fine by me — the design could go case-free — but I'm afraid motherboard manufacturers will be scared off by the smaller demand and simply stop even considering making them.
 
 As a single-socket replacement for the previous dual-socket Turin, Venice
 makes little to no sense (and it's still a C-core CPU, not a P-core one). To keep going single-socket, CPUs will have to push core counts (to 500 and beyond), the same way GPUs are doing.
 
 Unfortunately, while GPU designers are bulking up at 2–2.5 kW per GPU, the CPU crowd is still in kindergarten, scared to touch 600 W.
 
 If CEOs of AMD and Intel ever try to make their morning tea/coffee with 600W boiler instead of standard 1500W they'd probably finally understand that :)
 
