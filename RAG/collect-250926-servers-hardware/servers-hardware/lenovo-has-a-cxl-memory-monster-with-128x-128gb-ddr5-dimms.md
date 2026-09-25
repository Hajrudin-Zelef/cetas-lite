---
id: collect-250926-servers-hardware/servers-hardware/lenovo-has-a-cxl-memory-monster-with-128x-128gb-ddr5-dimms
title: "lenovo-has-a-cxl-memory-monster-with-128x-128gb-ddr5-dimms"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["memory", "amd", "consumer", "dram", "gpus", "intel"]
source: docs/RAG/clean4/lenovo-has-a-cxl-memory-monster-with-128x-128gb-ddr5-dimms.md
source_anchor: ""
source_lines: [1, 49]
sha256: a475f92220de2d70ebb3edb0f7e13def10c324ab1029da694ef5e5577ea0db39
---

# lenovo-has-a-cxl-memory-monster-with-128x-128gb-ddr5-dimms

At OCP Summit 2024, we saw one more CXL monster, the Lenovo ThinkSystem SR860 V3. This system is listed on Lenovo’s website as having “*Up to 16TB of TruDDR5 memory in 64x slots*.” At the OCP Summit 2024, we saw this beast and how Lenovo is doing it with the help of CXL and Astera Labs Leo.

## Lenovo Has a CXL Memory Monster with 128x 128GB DDR5 DIMMs

The idea of having a four-socket server is not new. A challenge, however, is that with a 8 channel Intel Xeon, you are limited by the number of DDR5 DIMM slots. Each 8-channel Xeon supports 2 DIMMs Per Channel (2DPC) for 16 DIMMs per socket. Four sockets gives us 64 DDR5 SODIMMs. Filling those with 128GB DDR5 RDIMMs only gets one to a paltry 8TB of memory. For some, a large consumer SSD-sized memory capacity is simply not enough.

At OCP Summit 2024, we saw Lenovo’s answer to add another 64 DIMMs to the system.

Using stacked boards with Astera Labs Leo CXL controllers, Lenovo is able to add many memory slots to the system.

Better said, up to 64 DDR5 DIMM slots. Each Astera Labs Leo CXL memory controller can handle up to four DDR5 DIMMs.

With this option installed, the top part of the chassis becomes a giant memory forest. The four CPUs sit on the motherboard at the bottom each with 16 DDR5 DIMM slots. Above that, is this CXL memory forest with another 64 DIMM slots.

The 64 DIMMs are connected directly to the Xeon CPUs, and then 64 additional DIMM slots are connected via CXL memory expansion. That gives us 128 DIMM slots. With 128GB RDIMMs, that is 16TB of memory capacity.

Something neat here is that the configuration with four CPUs and 128 DIMMs can also support four double-width GPUs.

That makes these very big systems.

## Final Words

For some scale-up workloads where massive memory capacity is needed, solutions like these using CXL Type-3 memory create topologies that would not be possible otherwise. The recent AI build-out has delayed the CXL deployment, but these solutions are coming. What is more, with CXL 2.0 (or really CXL 3.1), the vision is that in the future, one can hook up multiple memory expansion shelves, like the Inventec 96 DIMM CXL Expansion Box to CXL switches, building massive memory pools and then allocating and sharing them among different systems.

CXL is getting fun again. In Q1 2025 with Granite Rapids-SP we expect to see eight CPU systems like the Inventec with even more memory that will make this Lenovo 4-way system look small in comparison.

How long would posting take with 16TB?

All I gotta say is WOW. That’s a crap ton of RAM.

@Just Wondering, initializing ECC DRAM during POST and accessing CXL memory to determine its presence and size are two different things; CXL replacing persistent memory may well hold its contents between boots and be available to BIOS immediately, if not its ECC is likely to be different from DRAM and it would seem sensible that a new design would eliminate an old problem.

AMD explains it (though not for x86, but equally applicable) thusly:

“When the ECC mode is enabled, a write operation computes and stores an ECC code along with the data, and a read operation reads and checks the data against the stored ECC code. Consequently, it is possible to receive ECC errors when reading uninitialized memory locations. To avoid this problem, *** all memory locations must be written before being read ***. Writing to the entire DDR DRAM through the CPU can be time intensive. It might be worthwhile to use a DMA device to generate larger bursts to the DDR controller initialization and offload the CPU.”.

I’m not aware of DMA being used to initialize DRAM and 16TBs would take over an hour, don’t know why (with a new memory controller) this problem wouldn’t be resolved with DDR5 but I also don’t know that it is.

In addition server motherboards take a very long time to boot with very little memory, don’t know how they get their 9’s without failover to hide boot times.

One (other) benefit of CXL is that the large single pool of memory can be divided between the processors dynamically, to avoid overprovisioning; don’t know that is happening here (or in the prior system reviewed here), but if it’s not then it’s simply the ability to add additional memory beyond what would fit on the slots on the motherboard.

Something isn’t adding up correctly. Each of those Astera Labs boards have 24 slots arranged in six groups of 4. Three of those boards would equate to 96 slots. Four Xeons with 8 channels at 2 DPC would be 64 slots total from the Xeons. In aggregate, that’d be 160 slots in a single system and using 128 GB DIMMs in each would permit 18 GB of memory. That’s comfortably below the 64 TB traditional x86-64 memory cap.

My presumption is that the Xeon side is only leverage 1 DPC either with different system board or a dual socket 2 DPC configuration to get to 128 DIMMs in total.

Other alternative is that since this is a SAP HANA appliance, it simply has an artificial limit of 16 TB regardless of hardware.

Comments are closed.
