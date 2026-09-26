---
id: collect-250926-servers-hardware/servers-hardware/minisforum-ms-02-ultra-review-the-new-home-lab-king-2
title: "minisforum-ms-02-ultra-review-the-new-home-lab-king"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Alibaba", "Intel"]
dates: []
keywords: ["agents", "amd", "benchmark", "compute", "gpu", "gpus", "intel", "license", "memory", "power delivery"]
source: docs/RAG/clean4/minisforum-ms-02-ultra-review-the-new-home-lab-king.md
source_anchor: ""
source_lines: [54, 129]
sha256: dc0482db40f2085b88dad996e8daf3a1760cb1d035c8a098db323babb45c7a37
---

# minisforum-ms-02-ultra-review-the-new-home-lab-king

Fittingly, this configuration also allows for the blower on any actively cooled card to be facing an external vent, giving it access to the cool air it needs to keep itself cool.

Finally, since we happen to still have Minisforum’s similar MS-S1 Max system on hand, here are a couple of shots of the AI-focused systems side-by-side, showing how they differ, and how the MS-02 Ultra ends up being an even larger system.

To us, the MS-02 Ultra is not a replacement for the MS-S1 Max. Instead, the MS-S1 Max is the better AI machine, whereas the MS-02 Ultra is the better server node. Granted, if you saw our Minisforum MS-S1 Max review, you might also think about it as an AI server. We just checked, and ours is serving OpenClaw Agents a Qwen3-Coder-Next model so you could rightly argue it is a server as well. Still, the MS-02 Ultra feels more like it is geared towards traditional desktop or home lab server use.

And with that said and done, it is time to take a look inside the MS-02 Ultra.

I wish they’d have managed to make this fit into 2U so you could fit 2 of them side-by-side in a rack (via a presumably-optional bracket). As it is, it’s about 8mm too tall.

The MS-01 has a similar problem, where it’s ~3.5mm too tall for 1U.

These are slick. Wish I had a use case for one.

I would love to see a comparison between native Intel ECC memory support and the software based ECC capabilities offered on some of the mini platforms. I have soft ecc enabled via bios on 4x alderlake mini boxes and they have been stable for 1 yr+.

I wonder when MCIO pinout and cable finally landed in SFF or normal desktop mainboard? So many benefits come from it.

The article says that the 2 M.2 slots on the PCIe board are limited to PCIe3. The Minisforum site says that it is limited to PCIe3 for 8TB, and PCIe4 for <4TB. Was a less that 4TB SSD tested to see if it runs at PCIe4?

The lack of a TB5 name on the front most likely means some part of the data path is make by someone other than Intel. While they released TB to be an “open standard”, they still own the name and are the only ones that can grant its use. It’s been a good thing, allowing USB4 and now USB4v2 to be compatible. It’s dumb of them, really; in the couple of years nobody is going to bother paying Intel for the name at all, and as a result it will become worthless.

15W power delivery is nowhere near the TB5. Please turn you bs knob down a little bit.

That slot for the 25 Gbit NIC I see as a long term curse: you can’t use it for a 16x device even though it’d physically fit. There are several reasons for it. The first is that there are two sources of PCIe lanes going to it which only a few devices actively support in this fashion. The second is that it is 12 lanes in total. Leveraging 12 lanes in a 16x physical spot is permitted under the PCIe spec but few 16x devices support. Furthermore, those 12 lanes I believe are arranged in a nonstandard fashion for those devices that do support 12x. Thus the best fit for that slot would be a device that inherently needs bifurcation but users have to realize that one of the middle devices is not going to work as it doesn’t have PCIe lane going to it. For example, a quad M.2 carrier card would not have the second M.2 slot functional while the first, third and forth work fine.

I’d almost advocate for having all the lanes from that PCIe slot come from the PCH but it’d be bandwidth constrained due to the DMI link. I’d have implemented the slot as a 4x physical slot and then leverage an 8 lane MCIO connector to a proprietary card to get the additional two M.2 slots off of it. The 8 lane MCIO could find other uses, though the system would likely be physically space constrained for other cards.

Dual 10 Gbit NIC on the motherboard would have been nice to see.

I do hope that Nova Lake adds more PCIe lanes to the desktop platform. The oddities here are all examples of lane shortages forcing compromises into the design.

Why would anyone prefer this crap over Minisforum’s won BG-795SE MoBo with extra 25GbE NIC ?

Or mini-ITX MoBo with 9950X and NIC as needed ?

Why are you hyping those included NICs so hard ?

Hi Tinkering Ted – Just as a FYI as I have a system with the X3D version of that motherboard next to me, and we have another one with the non-X3D motherboard as well. The MS-02 has better I/O like the USB4 ports, it can also take an internal RTX 5060 LP while it has the 25GbE and 10GbE/2.5GbE ports, and more SSD. That AMD board is great, but this has way more expansion potential.

@Andrew

I based that section off of their user guide, which has the most comprehensive details on the matter.

To quote said guide: “For compatibility, the E810 expansion card’s NVME slots are set to PCIe 3.0 x4 by default. You can change them to PCIe 4.0 in BIOS, but after switching it is recommended to run a storage benchmark. If you see anomalies, revert the slot to PCIe 3.0.”

At least there, Minisforum isn’t guaranteeing anything. It

canrun at PCIe Gen4 speeds, but they clearly aren’t 100% confident about it. Which is why it defaults to Gen3 speeds (and why you’d need to go out of your way to enable Gen4).
Still, I’ve tweaked the language in the article a bit to make it clearer that Gen3 is the out-of-box setting rather than a hard technical limitation.

@Patrick Kennedy:

Who cares bout USB4 on such box.

Also BD-795SE has way faster PCIe5 lanes, that can be split with a simple splitter and use port bifurcation.

PCIE5x4 is enough for 100GbE. One could use 8 lanes with PCIe4 and be still left with 8 PCIe5 lanes for GPU.

Not to mention there are mini-ITX boards with AM5 that have 2x M.2 PCIe5 directly on CPU + PCIe5x16 for GPU etc.

@Illrigger . Read the block diagram. The USB4v2 80Gbps ports are supported with an Intel JHL9580 which is TB5 certified. Minis Forum just chose not to call it TB5.

@Tinkering Ted i care.

@spuwho: Without Intel (paid) certification for the whole device they can’t call it TB5.

When you say adding the Sparkle Arc A310 was a mistake, do you mean compared to the 5060 or were there problems more broadly?

Have one currently running transcoding/encoding duties on a plex server and was potentially planning to build an MS-02 with the A310 as a replacement. Bad idea?

Billy Baroo – I think the biggest reason was that the A310 is not a huge upgrade over integrated graphics if you just want basic GPU/ transcoding capabilities. It also does not offer the bigger gaming and small local AI model jump that you get with the RTX 5060. Better said, it is not a bad card. It just left us feeling like it was not a big enough jump to warrant adding that GPU to this system. Here is an example https://browser.geekbench.com/v6/compute/compare/5697728?baseline=5349589

I go KVM for more flexibility and have been running Debian KVM and Proxmox, so VM license is no longer my challenge.

The 350W power supply makes this system a non starter for anything AI-optimized: there is no way it can accommodate one of the nice 2-slot blower style GPUs (~800w), not counting the CPU, RAM, 4xNVME and NICs. Once again, Minisforum snatches defeat from the jaws of victory.
