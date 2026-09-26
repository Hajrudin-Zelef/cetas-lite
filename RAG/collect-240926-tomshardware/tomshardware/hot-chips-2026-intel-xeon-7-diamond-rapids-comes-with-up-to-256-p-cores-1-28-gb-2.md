---
id: collect-240926-tomshardware/tomshardware/hot-chips-2026-intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb-2
title: "hot-chips-2026-intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb-"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["intel", "18a", "amd", "benchmarks", "dram", "latency", "memory", "nvlink", "packaging", "panther lake"]
source: docs/RAG/clean_en/tomshardware/hot-chips-2026-intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb-.md
source_anchor: ""
source_lines: [59, 103]
sha256: 55c4f87777b31bd9f21cfca5d563ec916d60bb796ecb6f801f5557afcfae10e0
---

# hot-chips-2026-intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb-

 I think it's likely to be an offshoot of the Cougar Cove cores in Panther Lake. Those are discussed a bit, here:The article said:Intel didn't detail the core architecture (known as Panther Cove) in Diamond Rapids during its Hot Chips 2026 presentation, so we'll likely have at least one more technical deep dive on Diamond Rapids before it arrives, and possibly more.
 https://www.tomshardware.com/pc-components/cpus/intel-takes-the-wraps-off-panther-lake-first-18a-client-processor-brings-the-best-of-lunar-lake-and-arrow-lake-together-in-one-package
 Some additional details, here: https://www.techpowerup.com/review/intel-panther-lake-technical-deep-dive/5.html
 The main takeaway for me is that it's more of a tuned version of Arrow Lake's Lion Cove than a full, new generation. For server CPUs, even that much is welcome, since Intel's current P-core flagship, Granite Rapids, is still based on the Redwood Cove that's basically a tuned version of Raptor Cove.
 
 
 They certainly didn't move it from memory. DRAM doesn't know anything about caches. It wouldn't make sense, from either a bandwidth or latency perspective, to put any part of the directory cache in DRAM.The article said:In the memory fabric, you can see the standard flow through the DDR PHY into the memory controller, but Intel includes some special sauce at the end of the chain, notably an on-die snoop filter. A snoop filter is a directory to maintain cache coherency, and moving it onto the CPU removes directory storage and cache coherency tasks from the memory.
 
 Perhaps what they meant is that they centralized snoop filtering, compared to their older designs that might've had it more distributed?
 
 
 I wonder if this move could have anything to do with their upcoming NVLink integration.The article said:Intel isn't leveraging its advanced EMIB packaging to connect the fabric hubs to the CBBs. Instead, Intel is using a standard UCIe-S connection through copper in the substrate.
 
 
 Oops. This is a tricky area, so the precise wording is very important.The article said:AVX 10.1 served as a transition step off of AVX-512 and only supported 512-bit vector instructions. AVX 10.2 supports converged 256-bit vectors, enabling execution on both P-cores and E-cores.
 AVX-512 and AVX10.1 always supported multiple vector lengths of*up to* 512-bit. The vector length is determined at compile-time and can be: 128-bit, 256-bit, or 512-bit. All implementations support all vector lengths.
 AVX10.2 was originally announced to come in two flavors: a version for server CPUs that supported up to 512-bit vectors and a version for hybrid/client CPUs that supported only up to 256-bit vectors.
 The 256-bit limits on AVX10.2 were formally withdrawn by Intel. So, now all implementations of AVX10 will support up to 512-bit vectors (**source:** https://www.phoronix.com/news/Intel-AVX10-Drops-256-Bit ).
 
 To be clear: it's backward-compatible, but you can't reap the benefits of APXThe article said:Diamond Rapids also supports Intel's APX.
 ...
 Intel says software will see a performance improvement when recompiled with APX, and without source code changes.
 ...
 It doesn't require a code change, either, with full compatibility with previous code bases.*without* recompiling. In this regard, it's just like any of their AVX-family extensions, where you need to generate code which utilizes the new instructions to get any benefit from them.
 
 However, with a lot of web technologies being just-in-time compiled, users could see*some* benefits from APX on day 1 (or whenever web browsers start shipping JIT engines that support it).
 
 How soon other stuff would support it is quite an open question, since I doubt most game developers will want to maintain separate builds of their code for what's probably a < 10% performance improvement. That's a much smaller gain than various vector extensions had provided for vector-intensive code.
 
 It'll be good for certain benchmarks, like the SPEC CPU benchmarks that are more popular among server customers and CPU architects. Those get natively compiled for the machine they're running on.
 
 
 A key difference I see is that AMD has stuck with their approach of integrating L3 cache in the core dies, while Intel is making it more centralized. This is a compromise vs. Intel's previous approach of having L3 cache being fully distributed (i.e. in their Ring and Mesh architectures), so it could still be a win for them.The article said:borrows a lot of key design points from AMD's evolution with EPYC.
 
 
This is the closest I have yet heard of a confirmation that Diamond Rapids will lack SMT. That's going to hurt more in a P-core server than E-core servers or hybrid clients.The article said:Intel has continued to double down on Coral Rapids; however, the generation that will follow Diamond Rapids will reintroduce SMT to Xeon.
- 
Reply
 Distance is the reason. The part I mostly wonder about is why the Fabric Hubs aren't closer together and connected with EMIB since all cores can communicate cross hubs.bit_user said:I wonder if this move could have anything to do with their upcoming NVLink integration.
 
 DMR does not have SMT.bit_user said:This is the closest I have yet heard of a confirmation that Diamond Rapids will lack SMT. That's going to hurt more in a P-core server than E-core servers or hybrid clients.
 Diamond Rapids is Intel’s answer to the high-core-count, high-memory-bandwidth server lane as the platform heads toward 2027 against AMD’s EPYC push. It seems like Intel updated its roadmap recently to match AMD’s 256 cores since AMD’s 256-core design is an all-P-core design (albeit with 512 threads). https://www.servethehome.com/intel-diamond-rapids-the-2027-intel-xeon-at-hot-chips-2026/
 The lack of SMT was also mentioned by Ian Cutress.
 
Definitely seems likely that it's Cougar Cove with a little Coyote Cove on top.bit_user said:I think it's likely to be an offshoot of the Cougar Cove cores in Panther Lake.
