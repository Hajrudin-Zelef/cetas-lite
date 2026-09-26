---
id: collect-250926-servers-hardware/servers-hardware/amd-ups-ante-with-192gb-ryzen-ai-max-pro-400-chips-for-ai-systems-1
title: "amd-ups-ante-with-192gb-ryzen-ai-max-pro-400-chips-for-ai-systems"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "compute", "fp4", "gpu", "inference", "intel", "lpddr5x", "memory", "nvidia", "throughput"]
source: docs/RAG/clean4/amd-ups-ante-with-192gb-ryzen-ai-max-pro-400-chips-for-ai-systems.md
source_anchor: ""
source_lines: [1, 45]
sha256: b3728757f3bbeab26c0699bdb65a3fcaa5b3f874c7c55032cefcda284a15dd78
---

# amd-ups-ante-with-192gb-ryzen-ai-max-pro-400-chips-for-ai-systems

Ahead of the Computex trade show a bit later this month AMD is getting the ball rolling early with the announcement of a new series of Ryzen AI Max processors. Joining the existing Ryzen AI Max 300 series of chips, AMD Ryzen AI Max PRO 400 family will be continuing the Strix Halo legacy with a minor performance bump and a major memory capacity bump – which will see the chips paired with up to 192GB of RAM. Systems based on the new Ryzen AI Max chips will begin shipping in the third quarter of this year.


Diving right in to the hardware, the AMD Ryzen AI Max PRO 400 family is an additional set of chip SKUs based around AMD’s Strix Halo silicon (or as AMD calls the chips in this family, Gorgon Halo). This means we are looking at a high-performance SoC combining up to 16 Zen 5 CPU cores with AMD’s powerful RDNA 3.5 architecture integrated GPU, and feed by a sizable 256-bit memory bus. The combination of the high-performance iGPU and LPDDR5X-based memory subsystem has made Strix Halo a chip to contend with in the AI space over the last year, and AMD is looking to get a bit more out of the silicon over the coming year with a mid-generation refresh.

Compared to their Ryzen AI Max 300 counterparts, the Ryzen AI Max PRO 400 chips are a rather minor spec bump in terms of chip clockspeeds and core configurations – in fact most of the SKUs are unchanged in this regard. Only the top-end flagship SKU, the Ryzen AI Max+ PRO 495, offers meaningfully different specs in regards to computational throughput, getting a small boost to clockspeeds for its CPU, GPU, and NPU. All together, these bumps bring it up to 5.2GHz on the CPU while the GPU gets the new Radeon 8065S branding.

| **AMD Ryzen AI Max PRO 400 Chip SKUs** |  |  |  |  |  |  |  | 
|  | **CPU Cores** | **Max Boost** | **L3 Cache** | **NPU** | **GPU** | **TDP** | **Max RAM** | 
| **Ryzen AI Max+ PRO 495** | 16 | 5.2GHz | 64MB | 55 TOPS | 8065S 40 CUs | 45 – 120W | 192GB 8533MT | 
| **Ryzen AI Max PRO 490** | 12 | 5.0GHz | 64MB | 50 TOPS | 8050S 32 CUs |  |  | 
| **Ryzen AI Max PRO 485** | 8 | 5.0GHz | 32MB | 50 TOPS | 8050S 32 CUs |  |  | 
| **AMD Ryzen AI Max 300 Chip SKUs** |  |  |  |  |  |  |  | 
| **Ryzen AI Max+ 395** | 16 | 5.1GHz | 64MB | 50 TOPS | 8060S 40 CUs | 45 – 120W | 128GB 8000MT | 
| **Ryzen AI Max+ 392** | 12 | 5.0GHz | 64MB |  | 8060S 40 CUs |  |  | 
| **Ryzen AI Max 390** | 12 | 5.0GHz | 64MB |  | 8050S 32 CUs |  |  | 
| **Ryzen AI Max+ 388** | 8 | 5.0GHz | 32MB |  | 8060S 40 CUs |  |  | 
| **Ryzen AI Max 385** | 8 | 5.0GHz | 32MB |  | 8050S 32 CUs |  |  | 

Instead, the big change with the new chips is with memory support: AMD has dialed up memory clockspeeds a bit and greatly increased the the total amount of supported memory, which was one of the defining features of the Ryzen AI Max family to begin with. While the 300 series chips supported a maximum of 128GB of LPDDR5X-8000 memory, the new 400 series chips are boosting that by 50%, bringing the total to 192GB of LPDDR5X-8533 memory. With the recent advent of 24GB(192Gbit) LPDDR5X memory chips, AMD is quickly taking advantage of the tech to boost the Ryzen AI Max family’s already spacious memory capacity to a new level, while boosting memory bandwidth by about 7% to 273GB/sec.

With the chips’ popularity in systems for local AI inference, AMD’s value proposition for a higher memory capacity is rather straightforward: a larger memory pool allows for these systems to be loaded up with even larger models. The chips can be configured to allocate as much as 160GB of their memory pool to the GPU (leaving the last 32GB for the CPU), which is enough space to load up a 300B FP4 parameter model – a first for any single SoC system that is not a Mac Studio (and likely not the last).

Users who are truly limited by memory capacity should find the extra 48GB of GPU memory a significant boon. Though based on what we have seen with the Strix Halo hardware up to this point, this change is not likely to move the needle on performance. As AMD is not meaningfully boosting the compute throughput or memory bandwidth of the chips, existing bottlenecks such as pre-fill performance are going to remain where they are.

Meanwhile, it is notable that AMD is only launching PRO SKUs at this time. Non-PRO chips are not part of the upcoming chip launch, and AMD is not even confirming whether they will eventually release non-PRO chips. The good news for corporate customers at least is that this means all the Ryzen AI Max 400 systems will offer AMD’s full suite of enterprise management features. However, it also means that these chips will carry a further price premium for that functionality, driving up system costs while boosting AMD’s margins on their best AI SoCs.

It also means that, at least as far as these PRO SKUs are concerned, AMD’s focus/promotional push is solely on AI. Whereas the 300 series chips were launched as high-end laptop and small form factor chips with a multitude of uses (particularly gaming), AMD is all-in on AI for the 400 series chips.

Finally, speaking of systems, AMD tells us that new systems based on the 400 series chips will be launching in Q3 of this year. Confirmed partners include ASUS, Lenovo, and HP. Officially, the chips are rated for use in everything from laptops and mobile workstations to full-fledged desktops thanks to a rather wide cTDP window, however if OEM adoption of the Ryzen AI Max PRO 400 series is anything like the 300 series, then we would expect to see these end up almost exclusively in desktop systems. And particularly small form factor desktops such as AMD’s own Ryzen AI Halo, which is confirmed to be getting the 400 series chips later this year.

Is there anything preventing manufacturers from using these new 24GB memory chips with the existing 300-series CPUs? I’m quite certain that they use the exact same silicon, so I expect it should work, unless AMD is artificially limiting the compatibility to the new SKUs.

@Robert, I believe it would be an artificial limitation. But I’m withholding outrage because we don’t yet know if 300-series will actually be prohibited from reaching 192 GB. It could be a similar situation to Alder Lake-N officially supporting 16 GB on Intel’s website, but working with 32/48/64 GB. Or any case where the spec sheets for older chips were not updated to reflect new RAM capacities becoming available.

Maybe AMD would argue a new stepping enables higher capacity and explicit 8533 MT/s support instead of 8000 MT/s, but I’m not sure I’d buy it.

To others who may ask why it matters: Because Strix Halo BGA packages are just out there, spotted on the open market long ago, not tied to a specific amount of memory (like Lunar Lake was). Maybe manufacturers could save you $200 off a $5k 192 GB box by using the old chips that are virtually identical. It’s also possible that someone could buy any Strix Halo product, perhaps one with a faulty memory chip, desolder the memory chips, and add their own 24 GB chips, saving some money.

Those are definitely valid questions. I’m afraid I don’t have a definitive answer for you. Officially, AMD says that the 300 series chips top out at 128GB. Whether AMD is taking steps to actively enforce that via their firmware, I do not know.

I am more curious if they increased the PCI-E lanes from those stupid 16x… Having at least 24x would make adding an Intel E810 and clustering them a feature not a pain. Also, it would bring real PCI-E x16 GPU support. Heck, with 32x lanes, one could build a (Bluefield) DPU-like box for half the money, similarly to how Patrick sees the DGX Spark. But, as always, we are talking about AMD here…more preoccupied with catching up with Nvidia by copying them than hear the community…

Regardless of memory capacity it’s 50% the processing speed of Mac for the same price. Absolutely no reason to buy this unless they run out of Mac’s.

