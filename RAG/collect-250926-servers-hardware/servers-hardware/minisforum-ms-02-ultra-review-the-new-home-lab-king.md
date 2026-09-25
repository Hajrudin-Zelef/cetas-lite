---
id: collect-250926-servers-hardware/servers-hardware/minisforum-ms-02-ultra-review-the-new-home-lab-king
title: "minisforum-ms-02-ultra-review-the-new-home-lab-king"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Intel"]
dates: []
keywords: ["agents", "amd", "benchmark", "compute", "cost", "ethernet", "gpu", "gpus", "intel", "license", "memory", "power delivery"]
source: docs/RAG/clean4/minisforum-ms-02-ultra-review-the-new-home-lab-king.md
source_anchor: ""
source_lines: [1, 129]
sha256: 2eb5583ff9ae5a44909389c71a74ff26166e360a2e5ba74e4799a088d9860b63
---

# minisforum-ms-02-ultra-review-the-new-home-lab-king

Over the past 18 months or so, we have seen several mini-PC vendors put together small form factor systems designed for (or at least marketed to) the AI development crowd. With this being a quickly growing and fairly lucrative market, there continues to be a bit of a race to be in the business of selling the pickaxes in this modern gold rush.

And for specialist builders like Minisforum in particular, we are now seeing them offer multiple takes on the same concept as they tune their offerings to address different sub-sets of the AI market. The company already offers a highly integrated system with the Ryzen AI Max-based Minisforum MS-S1 Max. Now for their latest system, they are going the other direction, offering a rather modular/discrete Intel-based SFF system with their latest design, the MS-02 Ultra.

| **Minisforum MS-02 Ultra Key Specs (As Configured)** |  | 
| **Processors** | Intel Cora Ultra 9 285HX (5.5GHz) | 
| **Operating System** | Windows 11 Pro | 
| **Memory** | 64GB SO-DIMM DDR5-4800 (2x32GB) | 
| **Storage** | 1TB SSD (1TB PCIe 4.0 M.2 2280) | 
| **GPU** | Intel Graphics (Xe-LPG, 4 cores) | 
| **PSU** | 350W Internal PSU | 
| **Form Factor** | Mini-PC (4.8L) | 
| **Dimensions** | 222 x 225 x 97 mm (8.74 x 8.86 x 3.82 in) | 
| **Weight** | 3.45kg (7.6 lbs) | 
| **Wireless** | Wi-Fi 7 (2×2) + Bluetooth 5.4 | 
| **Color** | Black | 
| **Ports** | **Front:**  1x USB-A 10Gbps, 2x USB-C 80Gbps, 1x Combo Audio**Rear:**  1x HDMI 2.1 FRL, 3x USB-A 10Gbps, 1x USB-C 40Gbps, 1x 10GbE LAN (RJ45, RTL8127), 1x 2.5GBE LAN (RJ45, I226-LM), 2x 25 GbE LAND (SFP28, E810) | 

Officially, Minisforum bills this as the successor to the popular MS-01, a 2024-era system based around Intel’s Core 13<sup>th</sup> Gen (Raptor Lake) platform. And while that heritage is easy to see – especially with the use of an Intel platform – the MS-02 Ultra is not just an MS-01 with upgraded guts. It is a whole new design that, at nearly three times the volume of the MS-01, changes virtually everything about the design to allow for more expandability and performance. In short, it is a design that earns its “Ultra” moniker.

So what all are we looking at? In short, the MS-02 Ultra is a sizable mini-PC based on Intel’s Core Ultra 2 series (Arrow Lake-HX) platform. Minisforum is counting on several design aspects to set it apart from the crowd, including room for a dual-slot half-height video card, 80Gbps USB support, and in the high-end 285HX model, ECC memory support and a very rare set of dual 25GbE SFP28 ports. As a result, the MS-02 Ultra leans more towards a workstation in terms of features and component selection, with a copious amount of networking bandwidth to integrate it into larger networks.

If you wanted to find the Minisforum MS-02 Ultra online, here is an Amazon Affiliate link.


## Minisforum MS-02 Ultra External Hardware Overview

Minisforum’s MS-02 Ultra employs a new chassis design from the company that, by and large, is dictated by the system’s support for half-height PCIe cards. Whereas the MS-01 could fit a single card in a horizontal orientation, the MS-02 Ultra can fit 3 cards vertically. And that means the small(ish) form factor PC needs enough vertical clearance to accommodate those cards as well as its Intel desktop-class hardware. By SFF PC standards, this is a pretty sizable machine.

Consequently, the bulk of the front of the MS-02 Ultra chassis is reserved for airflow – and we will see why in better detail once we crack the system open. That leaves a relatively small portion along the front edge for expansion ports and buttons. Here we will find a single USB-A 10Gbps port, and a pair of deceptively powerful USB-C ports.

I say “deceptively” because, atypically for an SFF PC, these are 80Gbps USB-C (USB4 v2) ports being driven by Intel’s JHL9580 (Barlow Ridge) Thunderbolt 5 controller. As a result, they are not only far faster than USB-C ports in many other systems (which frequently top out at just 10Gbps), but they also support PCIe tunneling, 15W power delivery, and Thunderbolt networking. In practice, they are Thunderbolt 5 ports in all but name: Minisforum does not advertise the system as being Thunderbolt 5 capable, implying that they have not received official certification for the standard. Still, all of Thunderbolt 5’s major features are supported thanks to the Intel controller.

USB ports aside, rounding out the front panel is the obligatory 3.5mm combo jack and a power button.

Zooming back out and taking a look at the rest of the system, the MS-02 Ultra is designed to work in both horizontal and vertical orientations.

Along with rubber feet on the bottom for the horizontal layout, the system also has a set of feet on its right side, allowing it to safely and securely sit vertically.

Flipping over to the back, we get to see the other piece of the picture with regards to the rest of the various I/O ports and expansion brackets for the system, along with all of the exhaust vents that are helping to keep the guts of the MS-02 cool.

For all of the MS-02 models, Minisforum is including a trio of 10Gbps USB-A ports, as well as a 40Gbps USB-C port that is being driven by the SoC’s integrated controller. This USB-C port pulls double-duty as the system’s rear DisplayPort output as well, thanks to DP alt mode. Otherwise, if you need a dedicated and traditional display output, the system offers a single HDMI 2.1 port that is driven by the SoC’s integrated graphics.

All MS-02 systems also come with a pair of RJ45 Ethernet ports. On the left is a 2.5GbE port driven by Intel’s i226-LM controller. And to the right of that is an even faster 10GbE port that is driven by Realtek’s RTL8127 controller, the first real budget 10GbE controller on the market that we are quickly seeing peripheral and system manufacturers adopt.

If 10Gb Ethernet is still too slow for you, then you are in luck: on the high-end 285HX model like we are reviewing today, Minisforum also offers a pair of SFP28 25GbE ports that are situated on a PCIe expansion card. These connectors are driven by an Intel E810 controller, and suffice it to say, they give the MS-02 Ultra an ample amount of networking bandwidth, as well as access to more enterprise-focused features like RDMA support. Which, in turn, is a big part of the reason that Minisforum is pitching the MS-02 Ultra at the AI developer market (though admittedly, almost everyone is pitching products at the AI developer market these days).

The inclusion of SFP+ ports on the original MS-01 was one of that system’s more novel features. Here, the added SFP28 ports add a lot of cost and considerable power consumption, which is why it is only standard on the higher-end model.

I/O ports aside, the final connector you will find on the back of the MS-02 Ultra is a standard C14 power connector. Minisforum has integrated a 350 Watt power supply into the system – another major upgrade over the MS-01 – so gone is the need for an external power adapter and DC barrel connector. Though this does come at the cost of some space.

Meanwhile, our look at the MS-02’s chassis gives us a hint at the PCIe expandability offered by the system. Along with the PCIe slot that Minisforum’s 25GbE networking card sits in, there are two more unoccupied PCIe slots sitting behind blank PCIe brackets. This is another major upgrade for the MS-02 Ultra as compared to the MS-01, which only had a single expansion slot. The dual slots – or rather the double-width bay – affords the room to install a modern video card, as development of single-slot half-height cards has seemingly come to a halt. As a result, in order to be able to accommodate newer dual-slot video cards, Minisforum has expanded its PCIe bay size accordingly. The cost of that is significantly upscaling the size of the system to match.

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
