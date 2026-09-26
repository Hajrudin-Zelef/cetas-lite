---
id: collect-250926-servers-hardware/servers-hardware/amds-epyc-venice-instinct-mi455x-helios-hardware-on-display-for-first-time-at-ces-2026-2
title: "amds-epyc-venice-instinct-mi455x-helios-hardware-on-display-for-first-time-at-ces-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Cerebras", "Intel", "Microsoft", "TSMC"]
dates: []
keywords: ["amd", "mi455x", "chiplet", "compute", "cpo", "dram", "gpu", "hbm", "hbm4", "intel", "memory", "optics"]
source: docs/RAG/clean4/amds-epyc-venice-instinct-mi455x-helios-hardware-on-display-for-first-time-at-ces-2026.md
source_anchor: ""
source_lines: [45, 77]
sha256: 3ef8fa8bda0d55dae38c9aa2336bbf8d778e0652529f3f49888d91597a26c115
---

# amds-epyc-venice-instinct-mi455x-helios-hardware-on-display-for-first-time-at-ces-2026

MI455X is an impressive chip in terms of packaging as well but also feels a bit iterative given that AMD did something similar last generation. I do wonder if we’ll see custom implementations of this design mixing and matching a number of Zen 6 chiplets. Last generation AMD produced a custom part for Microsoft’s Azure cloud that only had Zen 4 CPU chiplets. Four sockets of 256 cores of Zen 6c with 432 GB of HBM4 memory each would be impressive for CPU based workloads.

AMD’s strategy feels like they’re rapidly moving toward a schema where CPU and GPU compute are interchangeable. I wonder how long it’ll be before AMD attempts a checkerboard pattern of CPU, GPU and HBM on a wafer scale interposer (think Cerebras large but chiplets). Just use a full wafer as the interposer that all the chiplets are mounted to. Put dedicated IO dies around the edge for things like networking and PCIe. It seems that the packaging tech is there to do it. Don’t ask how much power and cooling such a setup would be.

Imagining co packaged optics on a server cpu…sounds great, but also what sort of socket installation procedure is that going to be. Not that many of us still replace CPU’s on servers…but aren’t most of those switches with CPO are a BGP switch-chip. Neat to think about.

Anyway, agreed on that package size, should have room for expanding the CPU tiles. Extending the x86 capability is the way to stave off RISC/ARM or whatever comes next. This is the way. In the way of that though might be the socket compatibility or some other limitation that keeps it from getting bigger.

Do you have any idea what those four mini chiplets to each side of the IODs are?

AMD has not officially said. I suspect they’re either blanks (structural silicon) or some kind of electronics, but no way to know for sure right now.

Could it be possible that those are PCIe-PHYs? My idea is that maybe AMD connects the IODs to each other directly using Fan-Out, and to go off-package can add the PHYs to the same interface. Would allow to use the same IOD twice instead of a mirrored design and also to have smaller variants (1 IOD+4 CCD) with the same amount of PCIe Lanes.

@MDF While not copackaged optics, Intel did have on package Omnipath interfaces during the Sky Lake/Cascade Lake generation. So there was some cabling running to the ‘wings’ of that socket for the interface. Systems that supported on-package Omnipath had a slightly different bracket on the motherboard to account for the different CPU socket shape. Any sort of copoackaged optics would need to do the same. This does increase system setup time due to the additional steps to connect the optical cables and ensure that they’re routed properly in the chassis.

I don’t think that AMD is close to the maximum socket size for what can be assembled. I know off hand that there are even larger packages, though most of those are ceramic server chips from IBM. What might be the more limiting factor is not the socket size itself but motherboard limitations for housing that socket. Server motherboards are becoming rather cramped with lots of DIMM slots. Given a certain number of memory channels, there likely is a maximum socket size system designers would tolerate. With what AMD does have, there is still a little bit of room to leverage two larger chiplets in the corner of the current arrangement. Conceptually a FPGA or a small GPU chiplet could fit into these posts for a different style of compute in the server socket. Perhaps AMD will factor in more grow when they jump to DDR6 in a few years , SOCAMM and/or go full CXL memory expansion.

In terms of topology, AMD is still leveraging the IO dies as the means to handle routing coherent traffic between the chiplets. This seems to be an aggregation step which also includes the memory controller. Venice does change things up by leveraging two IO dies this time around. I would expect in the future AMD would include a coherency router on the main chiplet and go with a simple cardinal mesh link between chiplets directly (north, south, east, west) while moving the IO dies to the edge of the socket. This simple grid style design would only be limited in scale by room in the socket, power and cooling. This is similar but not exactly how Intel is leveraging chiplets in server. For Intel, they place a memory controller on each CPU chiplet which may or may not be leveraged based on socket IO (ie Intel can put in more chiplets than there are external memory channels). I do think it is wiser to keep the memory controller and PCIe as its own chiplet that sits on the same on package fabric between chiplets. Depending on the socket needs (SP5 vs. SP6 vs. AM5), only the number of IO dies needed would change to scale up.

@Stefan The IODs are flip-chips with I/Os all across their surface, not just at the edges. Genoa supported 128 + 8 = 136 PCIe lanes (differential pairs) from a single processor socket. If Turin has the same number of PCIe lanes, that would be 136 signals on each side of the processor package going in and out of 4 mini chiplets. Chips and Cheese suggested the mini chiplets could be decoupling capacitor chips to keep the voltage stable for high-speed single-ended DRAM signals.

It’s interesting that AMD reduced the number of CPU chiplets from 16 to 8 and went with larger size CPU chiplets. That suggests confidence in TSMC’s N2P yield or confidence that their design can isolate a defect so that a defect won’t ruin a whole chiplet. I agree with Kevin G that the 256 core Venice processor must be using low clock frequency Zen6c dense cores even if AMD calls them “Zen6” cores. The Venice processor with real Zen6 (high clock frequency) cores will probably contain about 192 cores to have an achievable improvement in performance/Watt from Zen5. A 192 core Venice processor could be organized as 8 chiplets with 2 CCX per chiplet and 12 cores per CCX. Desktop processors will probably have a single CCX with 12 Zen6 (high clock frequency) cores.

At CES, AMD also mentioned Venice-X, which has 3D V-Cache like Milan-X and Genoa-X. The 96 core version of Genoa-X has 12MB of L3 per core. Some media outlets have reported that Venice-X will have the option for a stack of two 3D V-Cache die under each CPU chiplet. Two 3D V-Cache die would provide 4 + 8 + 8 = 20MB of L3 per core. A 192 core Venice-X Zen6 processor could have:

192 x 12MB = 2304MB = 2.25GB of total L3 (one 3D V-Cache die per CPU chiplet) or

192 x 20MB = 3840MB = 3.75GB of total L3 (two 3D V-Cache die per CPU chiplet).

For Venice-X, each high clock frequency Zen6 CCX with 12 cores per CCX would share 12 x 12MB = 144MB or 12 x 20MB = 240MB of L3 cache. For comparison, the largest Granite Rapids processor in SNC3 mode has 43 cores sharing 168MB of L3 slices, with the whole processor containing about 3x that.

The pic of her holding it… Never thought I’d say this but, a banana for reference would’ve been nice!
