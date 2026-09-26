---
id: collect-240926-tomshardware/tomshardware/amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c-2
title: "amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "AWS", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "nvidia", "clearwater forest", "compute", "cost", "gpu", "gpus", "graviton", "intel", "latency"]
source: docs/RAG/clean_en/tomshardware/amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c.md
source_anchor: ""
source_lines: [47, 109]
sha256: c62c458f24a794db2ce176ef773b75a0d3b8c6d81de4d6e25b40fc7ff8b8247d
---

# amd-targets-nvidia-with-first-official-benchmarks-for-epyc-venice-cpus-company-c

With NVIDIA, you buy an 8-GPU motherboard and then populate it with 1, 2, 4, or all 8 GPUs (16 were possible before, too) — the 900 GB/s NVLink is already there. With CPUs all we ever had was a dual-socket design with a barely-there ~40 GB/s, and even that will most probably go extinct.
- 
Reply
 To the extent that "consumers" are using Epyc... single socket is cheaper and more accessible. AMD's "Siena" Epyc is single socket only, for example.Stomx said:But exist also one bad news for consumers: its massive socket size and larger RAM area most probably will kill dual socket designs within the previous extended ATX motherboard size. Dual socket motherboard if it will be build will look monstrous (take a look at the first single socket example like the MD25-CG0 from Gigabyte) and do not fit into any existing boxes. It is OK with me, the design could be box-free, but I afraid the manufacturers of motherboards will be scared of smaller demand and will stop making them.
 
 As a single socket arrangement instead previous dual socket Turin, Venice has little to no sense (and it is is still C-core not a P-core CPU). To continue with single socket, CPUs have to rise core counts to 500 and beyond same as GPUs doing.
 
 Venice SP8 doesn't support the full 256 cores, but it also cuts the memory channels down to 8 (16 DIMMs possible). A board could offer only 8 DIMMs for 8 channels, making it much smaller.
 
So I guess you got it right in your first sentence, the 256-core is not for plebeians at all. Not a really big surprise. We're long past the days of the $200 Threadripper special too.
- 
Reply
Dang, I need to set up a fake company (fake former cryptominer backstory) so I can get some AI swag. Hook a brother up with a DGX Spark.LordVile said:Doesn’t really matter how much better they are than Nvidia’s when Nvidia gives their products out for free, sorry for pinky promised payments and stock.
- 
Thanks for the article, @JakeRoach . Here's my analysis, FWIW.Reply
 
 On the one hand, it's pretty cool that Vera's 88-core/176-thread CPU can beat a 192-core Graviton 5. On the other hand, I'll bet Graviton 5 is clocked much more for efficiency, while Vera is clocked more for performance. That's at least how their previous generations were clocked, and those used the exact same cores (ARM Neoverse V2) as each other. It's also impressive Vera beat Intel's Granite Rapids.
 
 Intel's current-gen Clearwater Forest is missing from the comparison, of course.
 
 
 They more than doubled the memory bandwidth. I'm sure that's a big part of it.The article said:According to these results, the 9996 is around 78% faster than last-gen's 192-core EPYC 9965.
 
 
 No, the GCC spec subscore is telling you how long it takes for the test system to run GCC 11.2.0 on an input file, with the output targeted for an x86-64 machine (i.e. even if you run the benchmark on an ARM CPU, the output of the compiler is still for the same target as if you run the benchmark on an x86-64 CPU). You can read more about it, here:The article said:GCC 16 takes longer to compile due to better optimizations, hence the lower scores on the GCC and LLVM compilations above.
 https://www.spec.org/cpu2026/docs/benchmarks/721.gcc_r/721.gcc_r.html
 Likewise, the LLVM subscore is running the optimizer of LLVM 14.0.0 on an intermediate representation (IR), irrespective of what AMD used to compile the rest of the benchmarks. You can read more about it, here:
 https://www.spec.org/cpu2026/docs/benchmarks/723.llvm_r/723.llvm_r.html
 
 Well, the way SPEC CPU traditionally works is that each vendor supplies their own hardware and software. It's common practice for vendors to make official submissions with a compiler that has custom tuning for their specific CPU.The article said:Regardless, it's not best practice to compare benchmarks using two different compiler versions. It makes sense that AMD used GCC 16.1 — it includes support for Zen 6 — but ideally Vera would also be on GCC 16.1.
 
 This practice dates back to the old days, when some computers were*only* supported by the vendor-provided compiler. It's still common for supercomputers to compile their software with a vendor-provided compiler. This is why Intel and AMD have their own, even though AMD's (called AOCC) is based on LLVM/Clang and just has some custom tuning.
 
 Intel's compiler is called ICC and used to be based on a backend by EDG (Edison Design Group). About 5 years ago, they also adopted LLVM.
 
 This one looks good for AMD,*unless* you know what the raw bandwidth actually is for both systems: 1.23 TB/s and 1.64 TB/s. So, Nvidia's bandwidth*efficiency* is actually way better: 89.4% vs 79.2%. That has to do with the difference in memory coherency model between ARM and x86. I'm actually surprised AMD did that well. I'd have guessed it'd be no higher than 75%.
 
 Anyway, one thing we can say for sure is that Nvidia's*power* efficiency is way better, since they're using LPDDR5X-based SOCAMM2 modules, while AMD is using power-hungry MRDIMMs.
 
 
 The difference in raw bandwidth comes directly from Nvidia's use of LPDDR5X-9600 vs AMD's use of DDR5-12800. Both have the same data width of 1024 bits per CPU. In AMD's case, that works out to 16 DIMMs, while Nvidia uses 8 SOCAMM2 modules. SOCAMM2 have twice the data width, in part, because LPDDR5X multiplexes address and data over the same pins, whereas regular DDR5 has separate pins for each. This makes the practical bandwidth and latency of LPDDR5X a bit lower and longer, but it saves cost and power.The article said:Vera absolutely clobbered the competition in the publication’s original Stream results.
 
 
 Customers will also be looking at perf/W and TCO (Total Cost of Ownership, where operating costs are combined with system purchase price).The article said:Although looking at benchmark results is always interesting, it doesn't say much in the context of a server deployment, at least at the scale that AMD is targeting. Peak performance is only one of the major factors that go into server deployments, after all, and even then, performance can vary wildly depending on what software you're running and how it's built.
 
 
I think the comparison with Nvidia's Vera is a little bit weird. It seems clear to me that Vera is primarily aimed at orchestration, while AMD's Venice is aimed much broader set of tasks, mostly consisting of general compute workloads. I think the comparison against Vera is more about bragging rights and trying to take wins over Nvidia anywhere that AMD can find them.
- 
Reply
 The cores in this CPU will find their way into some AMD-powered desktops and laptops. Others will have coresStomx said:This fat toy is for the future world "elite" and their personal AI datacenters.*derived* from these. So, the way I look at it is as a taste of what's to come.
 
 
 I think the insane memory bandwidth is just to feed the insane core count + PCIe 6.0 I/O. Their current platform is bandwidth-starved in each domain and PCIe 6.0 doubles I/O bandwidth, while more cores at higher clockspeed also increases bandwidth demands from within the CPU.Stomx said:Maybe it will run local AI on the CPU: with 2.5x the RDRAM bandwidth of AMD Turin,
 
 
 Oh, no question about that. Each CPU needs 16 DIMMs to achieve maximum memory bandwidth. So, there's no realistic way you're fitting 2 of those CPUs and 32 DIMM slots on an EATX board.Stomx said:the massive socket size and the larger RAM footprint will most probably kill dual-socket designs within the old extended-ATX form factor.
 
 
 I thought Venice was supposed to go up to 700W. Intel's 256-core Diamond Rapids CPUs are supposedly going to use up to 650W, though I don't know if that's been officially announced.Stomx said:the CPU crowd is still in kindergarten, scared to touch 600 W.
 
 That said, remember that customers don't*want* these CPUs to draw lots of power. Nvidia likes to point out that Vera and its memory run on a budget of only 500W.
 
 
