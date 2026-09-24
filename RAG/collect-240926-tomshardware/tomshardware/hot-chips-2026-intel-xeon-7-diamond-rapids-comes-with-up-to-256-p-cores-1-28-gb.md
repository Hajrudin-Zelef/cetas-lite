---
id: collect-240926-tomshardware/tomshardware/hot-chips-2026-intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb
title: "hot-chips-2026-intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb-"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["intel", "18a", "accelerator", "agentic", "amd", "benchmarks", "chiplet", "clearwater forest", "compute", "dram", "latency", "memory"]
source: docs/RAG/clean_en/tomshardware/hot-chips-2026-intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb-.md
source_anchor: ""
source_lines: [1, 103]
sha256: f0108991d86f698495642105b8af1af7ee24fb8bd516afdaa8b283d25406c2e9
---

# hot-chips-2026-intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb-

<!-- source: https://www.tomshardware.com/pc-components/cpus/intel-xeon-7-diamond-rapids-comes-with-up-to-256-p-cores-1-28-gb-of-last-level-cache-next-gen-18a-p-cpu-also-brings-avx-10-2-and-uses-ucie-s-instead-of-emib -->

After teasing the chips earlier this year, Intel has provided some details on its next-gen Xeon 7, codenamed Diamond Rapids, CPUs. Featuring up to 256 P-cores and 1.28 GB of last-level cache, the new range of CPUs is set to release in the data center in 2027. The range brings forth several advancements we've expected on Intel's roadmap, including the enhanced 18A-P process, UCIe interconnects, AVX 10.2, and Intel's new "fan-out" fabric.

Intel didn't detail the core architecture (known as Panther Cove) in Diamond Rapids during its Hot Chips 2026 presentation, so we'll likely have at least one more technical deep dive on Diamond Rapids before it arrives, and possibly more. Although there are still questions about Panther Cove, Intel shared a technical breakdown of how Diamond Rapids chips are built more broadly, including a look at the compute tiles and how they come together across the chip.

Intel calls the compute tiles Compute Building Blocks, or CBBs, and they hold the core chiplet stacked on top of the base tile that holds the LLC. Each core chiplet can hold up to 16 cores, and based on the scaled-up Diamond Rapids SoC, up to four of those chiplets can live in a CBB. Each chiplet connects to the base tile with a 3D Xbar. A full Diamond Rapids SoC includes four base tiles built on Intel 3-T, two fabric hub tiles built on Intel 3, and 16 core chiplets built on Intel 18A-P.

Bringing everything together are two advanced packaging techniques. Intel is once again using its own Foveros Direct 3D to bond the compute tiles to the base tiles, as seen with Xeon 6+ 'Clearwater Forest' CPUs. Intel is using UCIe-S to connect the fabric hub tiles to the cores via a copper connection. Notably, Intel isn't using its own Embedded Multi-die Interconnect Bridge (EMIB) that it's broadly deployed in past products.

## Intel Xeon 7 'Diamond Rapids' compute chiplet

Diamond Rapids is built with four Compute Building Blocks, each of which includes four core chiplets that house 16 P-cores each. The cores have access to private L2 within each chiplet, and they share an L3 cache located on the base tile. The chiplets are connected to the base tile with a 3D crossbar, packaged with Foveros Direct 3D.

Within each CBB, there's 3D packaging, but Intel leverages 2D communication via a UCIe-S interconnect to connect the CBBs to two centralized fabric hubs, allowing the cores (and caches) to communicate with each other. Although there are two fabric hubs, each of the CBBs is connected to both fabric hubs, so communication routes are clear across the chip.

Compared to Granite Rapids, Intel has quite literally flipped the layout, centralizing memory and I/O while pushing the cores out to the edges of the chip. It's much closer to a layout we'd expect to see from AMD.

Thermal improvements will likely follow. With the highest-clocked and hottest components pushed out to the edges, there's much less concern for hot spots in the middle of the chip, as is the case with Granite Rapids-AP, where the cores are at the center.

Intel is using its latest enhanced 18A-P node for the compute die, which is said to increase performance by 9% compared to 18A at peak performance, or operate at 18% lower power with iso-performance. Intel announced in June that 18A-P had entered risk production.

## Intel Xeon 7 'Diamond Rapids' fan-out fabric and memory, I/O subsystem

Diamond Rapids comes with 16-channel memory, supporting up to 8,000 MT/s with DDR5 and up to 12,800 MT/s with MRDIMMs. Although Intel bumped memory speeds with Xeon 6+ 'Clearwater Forest,' we're now seeing fast DDR5 support on a P-core Xeon, and with an expansion to 16 channels (Granite Rapids topped out at 12 channels).

Intel centralizes all of the hardware for memory and I/O communication in the middle of the chip across two tiles (the fabric hubs), and each CBB can communicate with both fabric hubs.

Double-clicking into the diagram at the top of this section, you can see the layout of the I/O system above. Across the chip, Intel supports 128 lanes of PCIe 6.0, CXL 3.0, UPI 3, or some combination thereof, courtesy of the flexible I/O subsystem. Intel also includes four PCIe 4.0 lanes (a total of eight per CPU) for platform use.

The I/O fabric also includes complexes for the various accelerators on-chip in Diamond Rapids, including Intel QuickAssist Technology (QAT) and In-Memory Analytics Accelerator (IAA).

In the memory fabric, you can see the standard flow through the DDR PHY into the memory controller, but Intel includes some special sauce at the end of the chain, notably an on-die snoop filter. A snoop filter is a directory to maintain cache coherency, and moving it onto the CPU removes directory storage and cache coherency tasks from the memory.

Interestingly, Intel isn't leveraging its advanced EMIB packaging to connect the fabric hubs to the CBBs. Instead, Intel is using a standard UCIe-S connection through copper in the substrate. Intel says that UCIe-S offered a "low-latency uniform connection to all of the memory hubs" that "made the most sense for Diamond Rapids."

There were a handful of questions around UCIe-S versus an advanced packaging technique, UCIe-A. Intel says the choice mainly came down to distance, with UCIe-A requiring multiple "hops" depending on the distance. UCIe-S provides uniform access across longer distances, enabling lower latencies across the entire chip.

## Intel Advanced Performance Extensions and AVX 10.2 support in 'Diamond Rapids'

Although it's more of a footnote in the headline reveals about Diamond Rapids, the next-gen Xeon CPUs mark an important milestone in Intel's journey with AVX-512 and Intel's Advanced Performance Extensions, or APX, which has been described as a modernization of the x86 ISA. Both were described in 2023, and now they're showing up in Diamond Rapids.

First, AVX. Expectedly, Diamond Rapids marks the move to AVX 10.2, which is supported on both P-cores and E-cores (AVX 10.1 only worked on P-cores). AVX 10.1 served as a transition step off of AVX-512 and only supported 512-bit vector instructions. AVX 10.2 supports converged 256-bit vectors, enabling execution on both P-cores and E-cores.

Diamond Rapids also supports Intel's APX. APX doubles the number of general-purpose registers from 16 to 32 with new encoding for registers 16 through 31. Intel says software will see a performance improvement when recompiled with APX, and without source code changes. We heard about APX support in Panther Cove nearly two years ago for the first time.

APX requires 10% fewer loads and 20% fewer stores in memory, according to Intel, and includes some key instruction updates like condition load and store. It doesn't require a code change, either, with full compatibility with previous code bases.

Between AVX 10.2, AMX, centralized I/O and memory communication, and 18A-P, Diamond Rapids brings forth a lot of innovation that Intel has been talking about for a long time. Whether it's too little, too late remains to be seen with the missteps around Granite Rapids.

Given the explosion of CPU demand for agentic workloads, Intel has a competitive part here that, at least, supports the latest updates to the x86 ISA and borrows a lot of key design points from AMD's evolution with EPYC. Intel has continued to double down on Coral Rapids; however, the generation that will follow Diamond Rapids will reintroduce SMT to Xeon.

## Full Intel Xeon Diamond Rapids Hot Chips 2026 presentation

- 
Thanks for the writeup!Reply
 
 
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
