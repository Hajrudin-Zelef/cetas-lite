---
id: collect-250926-servers-hardware/servers-hardware/hpe-shows-off-amd-epyc-venice-and-sp7-supercomputing-node-at-sc25
title: "hpe-shows-off-amd-epyc-venice-and-sp7-supercomputing-node-at-sc25"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "compute", "intel", "liquid cooling", "memory", "mi400", "nvidia", "rubin", "vera rubin"]
source: docs/RAG/clean4/hpe-shows-off-amd-epyc-venice-and-sp7-supercomputing-node-at-sc25.md
source_anchor: ""
source_lines: [1, 29]
sha256: f05d7d27b9a54ac779980dbbaba96394548e87c4c390b5f0b24b7067b9d755f7
---

# hpe-shows-off-amd-epyc-venice-and-sp7-supercomputing-node-at-sc25

HPE has its new GX5000 blade on display. We recently covered the launch of HPE’s New AMD EPYC Venice Instinct MI400 and NVIDIA Vera Rubin Compute Blades, but now we get to show at least the CPU-focused blade. As a bonus, this is one of the first AMD EPYC Venice platforms and socket SP7 systems that have been shown publicly.

## HPE Shows off AMD EPYC Venice and SP7 Supercomputing Node at SC25

As a PCIe Gen6 CPU, we are going to see more EDSFF SSDs in future systems since PCIe Gen5 appears to be the end-of-the-line for 2.5″ drives. HPE has what appears to be eight E1.S drive bays up front.

Here is the full node. You can see the bare sockets to the left, then there are sockets with the Slingshot 400Gbps NICs in the middle left. On the right we see four sockets that are all liquid-cooled both CPUs and memory. If you have seen our previous HPE Cray EX series coverage this will have a lot of familiar features.

Earlier this week, we discussed how next-generation processors will have 16 memory channels, an important distinction as Intel Cancels its Mainstream 8-channel Next-Gen Xeon Server Processors. Here we have a TE Connectivity AMD SP7 socket flanked by eight DIMMs on each side to fill all of the channels. We expect production units to have the AMD EPYC Venice series installed, but also much faster memory. The DIMMs here are DDR5-5600 RDIMMs that seem to be there just for show.

HPE’s key feature is fitting eight of these CPUs onto a compute blade and liquid cooling them. It should be noted that these blades are much larger than typical enterprise blades.

Here is a look at the liquid-cooled side. You can see the memory and the CPUs are liquid-cooled.

For iLO fans, that is on the motherboards as well.

A small detail above is that those might be the new PCIe Gen6-era MCIO connectors.

## Final Words

This appears to be an early prototype node, but it was still great to see. It is also great to see an AMD EPYC Venice platform that we expect will sell well into government installations. There is a lot of code in those domains built to run on CPUs only, so CPU-only clusters like this are very popular. HPE also has the ability in the new GX5000 generation to mix blades, so one could, in theory, have a Venice plus MI430X blade and a Venice-only blade in the same rack.

The last picture looks to be the new KickStart connector that is starting to show up on more advanced designs, I’ve seen it in some of the high end PCIe 5.0 hardware raid cards coming out.

16 memory channels per CPU… holy crap that is is a lot of memory traces. You basically would need a supercomputer just for layout. I wonder how many layers the PCB has.

SP7 boards will typically have either 22 or 24 PCB layers

Twenty years ago, servers had two single core processors per rack unit (U). Assuming EPYC Zen 6 has up to 192 P cores per socket, 8 sockets is 1536 P cores. That’s equivalent to more than eighteen 42U racks with 2 cores per U.
