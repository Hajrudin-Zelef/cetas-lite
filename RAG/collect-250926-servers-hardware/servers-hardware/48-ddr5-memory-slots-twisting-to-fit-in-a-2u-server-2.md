---
id: collect-250926-servers-hardware/servers-hardware/48-ddr5-memory-slots-twisting-to-fit-in-a-2u-server-2
title: "48-ddr5-memory-slots-twisting-to-fit-in-a-2u-server"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "dram", "research"]
source: docs/RAG/clean4/48-ddr5-memory-slots-twisting-to-fit-in-a-2u-server.md
source_anchor: ""
source_lines: [61, 89]
sha256: dbc05af45e1a11cff1e3727bc32d19cbcccaf32c55698f7ed8d90ade2b7b6194
---

# 48-ddr5-memory-slots-twisting-to-fit-in-a-2u-server

Theoretically, a 48-DIMM system could be condensed down into just 12 physical modules, six per processor, using this model and this in turn would alleviate some of the space constraint. Yes, it would mean an extra chip (the LMC) between the RAM and CPUs, but I believe this could work well enough. But then again, I’m not a system design engineer and somebody else might have a better idea than mine.

@Stephen Beets – to see what you’re suggesting in action, look at IBM Power 8, which used the Centaur chip as your LCM, a high-speed serial Axion bus to the CPU (same bus used for interconnect as well), and standard DRAM channels from Centaur to the DIMMs.

Find a detailed description here: https://research.ibm.com/publications/the-cache-and-memory-subsystems-of-the-ibm-power8-processor

@Scott, for 2U they have come up with double height MCRDIMMs, see articles at Tom’s and Anand.

“Micron demonstrated one ‘tall’ 256 GB DDR5-8800 MCRDIMM at GTC (pictured) but also plans to offer MCRDIMMs of standard height for applications like 1U servers. Both 256 GB MCRDIMMs are based on monolithic 32 Gb DDR5 ICs, but the tall one places 80 DRAM chips on both sides of the module, whereas the standard one uses 2Hi stacked packages, which means that they run slightly hotter due to less space for thermal dissipation. In any case, the tall module consumes around 20W, which isn’t bad as Micron’s 128GB DDR5-8000 RDIMM consumes 10W at DDR5-4800.”

That solves capacity, but doesn’t make the coming 32 DIMM slots (per CPU) take up less space.

I suspect that placing DIMMs on both sides of the motherboard would lead to crosstalk problems, not to mention the need for finer traces to run to the DIMM areas. If it were practical I suspect we would have seen it become popular already, if not only to allow more memory for tiny motherboards.

CAMM and DDR6 may also solve the capacity problem, though not the real estate issue, but those won’t be coming before MCRDIMM. Some genius might try rotating the CPUs 90° only to discover that doing so affords the tiny bit of extra space needed.

SODIMMS!

@David Freeman

Thanks. I kind of figured somebody already did something along the lines of what I outlined.

Unfortunately, the link you provided goes straight to an “Error 404: Page not found” on IEEE’s site. And since they don’t let you read full research papers without an account and paying them stupid money, I couldn’t have read it even if it was available. But, yeah.

I’m surprised that line of thinking HASN’T been more widely adopted. Why has only IBM thought of this and why only with POWER8? Surely ONE other company has thought of the same idea that I had, or that IBM had.

One of the newest layouts seems to be Gigabyte’s H253-Z10-AAP1, where each node has the memory and processor turned to a 45° angle; resulting in 24 channels per node, or 48 in a 2 CPU 1U.

Comments are closed.
