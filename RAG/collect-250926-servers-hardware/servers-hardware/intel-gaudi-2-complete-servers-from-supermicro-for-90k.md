---
id: collect-250926-servers-hardware/servers-hardware/intel-gaudi-2-complete-servers-from-supermicro-for-90k
title: "intel-gaudi-2-complete-servers-from-supermicro-for-90k"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["intel", "accelerator", "cost", "datacenter", "inference", "memory", "mlperf", "nvidia", "pricing"]
source: docs/RAG/clean4/intel-gaudi-2-complete-servers-from-supermicro-for-90k.md
source_anchor: ""
source_lines: [1, 27]
sha256: bd7bee3d8f6dbf34853eeacf7d67d78992bb3b6041b6aa92e96146ba00eb5a56
---

# intel-gaudi-2-complete-servers-from-supermicro-for-90k

Intel has not been shy in saying that it is heavily discounting Gaudi 2 versus the NVIDIA H100 competition. Now, we have some sense of that. Supermicro is currently running a promotion for complete Gaudi 2 OAM AI systems for $90K.

## Intel Gaudi 2 Complete Servers from Supermicro for $90K

We covered this in the NVIDIA MLPerf Inference v4.0 release, but something that was highlighted is the performance per dollar of Gaudi 2 versus the NVIDIA H100. It seems as though the value proposition may have changed.

Normally, the pricing for AI servers is held as a relatively tight secret. That is because, with the NVIDIA HGX A100 and HGX H100 designs, server OEMs are paying a similar amount for everything, so margins are generally under a lot of pressure. For the first time in a long time, we saw a price on an AI server. Here is what $90K gets you in a configuration for the Supermicro SYS-820GH-TNR2-01.

The interesting part about this is that it is still using an Ice Lake Xeon platform, so it is on PCIe Gen4 and lower-cost DDR4 memory. The list price for the Intel Xeon Platinum 8368 processors is $7214 each. Those 64GB DIMMs usually price out in new systems at around $3000-$3200 for 16 of them. There is probably another $1500-2000 in SSD content and a bit more for the NVIDIA ConenctX-6 VPI NIC. We only have one of those NICs because Gaudi 2 uses its onboard 100GbE networking directly from the AI accelerators. Here is the Gaudi 3 version that bumps the speed, but it is similar to how the Gaudi 2 networking works.

That is a lot of server for $90,000, especially with eight Gaudi 2 AI accelerators and all of the 100GbE networking for the inter-accelerator traffic.

## Final Words

A few things here. First, we have already shown Supermicro’s Gaudi 3 systems, and Gaudi 3 will start ramping up sales and production very soon. Second, we reached out to Supermicro, and that the $90K configuration can be customized, and we were told the pricing should change based on the specified options. If you want to add another NIC for example, the pricing should change by the cost of adding the NIC.

Still, this entire exercise shows that the Gaudi 2 is now being sold for well under $10K/ AI accelerator.

If you want to see these in action, we actually saw some running when Touring the Intel AI Playground – Inside the Intel Developer Cloud.

Clearly they aren’t getting Nvidia margins on these; but do we know(or have reasonably plausible estimates) of where such pricing lands relative to the cost of production and to the margins of other things they could be producing on their N7 process; and their margins more generally?

Is this “practically giving them away, for some combination of avoiding being stuck with dead stock when Gaudi 3 comes out and expanding the install base”? “Not losing money on them; but it’s more for the install base than for the money”? “In line with Xeon margins”? “Not Nvidia money; but better than average for the Datacenter division as a whole”?

Based on price quotes for the Gaudi2 (HL-225H) they’re either giving away the server for free, or they’re eliminating margins on the accelerators. So I expect the sale of these systems at $90k to be at negative margin, either to Intel or SuperMicro.

Comments are closed.
