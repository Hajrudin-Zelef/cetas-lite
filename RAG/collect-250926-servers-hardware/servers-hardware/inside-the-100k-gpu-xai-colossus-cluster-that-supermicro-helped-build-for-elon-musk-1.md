---
id: collect-250926-servers-hardware/servers-hardware/inside-the-100k-gpu-xai-colossus-cluster-that-supermicro-helped-build-for-elon-musk-1
title: "inside-the-100k-gpu-xai-colossus-cluster-that-supermicro-helped-build-for-elon-musk"
domain: servers-hardware
role: reference
task: reference
actors: ["AWS", "Broadcom", "Google", "Meta", "Microsoft", "Nvidia", "xAI"]
dates: []
keywords: ["gpu", "blackwell", "compute", "data centre", "datacenter", "distribution", "gpus", "liquid cooling", "nvidia", "research", "training"]
source: docs/RAG/clean4/inside-the-100k-gpu-xai-colossus-cluster-that-supermicro-helped-build-for-elon-musk.md
source_anchor: ""
source_lines: [1, 74]
sha256: 51bb38d111b3336ac5e3df3fde0740e2713f75477902cbcb421d88a131637b1b
---

# inside-the-100k-gpu-xai-colossus-cluster-that-supermicro-helped-build-for-elon-musk

Today, we are releasing our tour of the xAI Colossus Supercomputer. For those who have heard stories of Elon Musk’s xAI building a giant AI supercomputer in Memphis, this is that cluster. With 100,000 NVIDIA H100 GPUs, this multi-billion-dollar AI cluster is notable not just for its size but also for the speed at which it was built. In only 122 days, the teams built this giant cluster. Today, we get to show you inside the building.

Of course, we have a video for this one that you can find on X or on YouTube:

Normally, on STH, we do everything entirely independently. This was different. Supermicro is sponsoring this because it is easily the most costly piece for us to do this year. Also, some things will be blurred out, or I will be intentionally vague due to the sensitivity behind building the largest AI cluster in the world. We received special approval by Elon Musk and his team in order to show this.

## Supermicro Liquid Cooled Racks at xAI

The basic building block for Colossus is the Supermicro liquid-cooled rack. This comprises eight 4U servers each with eight NVIDIA H100’s for a total of 64 GPUs per rack. Eight of these GPU servers plus a Supermicro Coolant Distribution Unit (CDU) and associated hardware make up one of the GPU compute racks.

These racks are arranged in groups of eight for 512 GPUs, plus networking to provide mini clusters within the much larger system.

Here, xAI is using the Supermicro 4U Universal GPU system. These are the most advanced AI servers on the market right now, for a few reasons. One is the degree of liquid cooling. The other is how serviceable they are.

We first saw the prototype for these systems at Supercomputing 2023 (SC23) in Denver about a year ago. We were not able to open one of these systems in Memphis because they were busy running training jobs while we were there. One example of this is how the system is on trays that are serviceable without removing systems from the rack. The 1U rack manifold helps usher in cool liquid and out warmed liquid for each system. Quick disconnects make it fast to get the liquid cooling out of the way, and we showed last year how these can be removed and installed one-handed. Once these are removed, the trays can be pulled out for service.

Luckily, we have images of the prototype for this server so we can show you what is inside these systems. Aside from the 8 GPU NVIDIA HGX tray that uses custom Supermicro liquid cooling blocks, the CPU tray shows why these are a next-level design that is unmatched in the industry.

The two x86 CPU liquid cooling blocks in the SC23 prototype above are fairly common. What is unique is on the right-hand side. Supermicro’s motherboard integrates the four Broadcom PCIe switches used in almost every HGX AI server today instead of putting them on a separate board. Supermicro then has a custom liquid cooling block to cool these four PCIe switches. Other AI servers in the industry are built, and then liquid cooling is added to an air-cooled design. Supermicro’s design is from the ground up to be liquid-cooled, and all from one vendor.

It is analogous to cars, where some are designed to be gas-powered first, and then an EV powertrain is fitted to the chassis, versus EVs that are designed from the ground up to be EVs. This Supermicro system is the latter, while other HGX H100 systems are the former. We have had hands-on time with most of the public HGX H100/H200 platforms since they launched, and some of the hyper-scale designs. Make no mistake, there is a big gap in this Supermicro system and others, including some of Supermicro’s other designs that can be liquid or air cooled that we have reviewed previously.

In the latter part (2017-2021) of my almost decade working in the primary data center for an HFT firm, we moved from air cooled servers to immersion cooling.

From the server side that basically meant finding a vendor willing to warranty servers cooled this way, removing the fans, replacing thermal paste with a special type of foil and (eventually) using power cords made of a more expensive outing coating (so they didn’t turn rock hard from the mineral oil cooling fluid.)

But from the switch side (25 GbE) no way the network team was going to let me put their Arista switches in the vats…Which made for some awkwardly long cabling and eventually a problem with oil wicking out the vats via the twinax cabling (yuck!).

I would look at immersion cooling as a crude (but effective) “bridge technology” between the worlds of the past with 100% air cooling for mass market scale out servers, and a future heavy on plumbing connections and water blocks.

This is extremely impressive.

However, coming online in 122 days is not without issue. For example, this facility uses at least 18 unlicensed/unpermitted portable methane gas generators that are of significant concern to the local population – one that already struggles with asthma rates and air quality alerts. There is also some question as to how well the local utility can support the water requirements of liquid cooling at this scale. One of the big concerns about liquid cooling with datacenters is the impact to the water cycle. When water is typically consumed it ends up as wastewater feeding back to treatment facilities where it ends up back in circulation relatively quickly.

Water-based cooling systems used in datacenters use evaporation – which has a much longer cycle: atmosphere -> clouds -> rainwater -> water table.

Other clusters and datacenters used by the likes of Meta, Amazon, Google, Microsoft, etc take the time and caution to minimize these kinds of environmental impact.

Again, very impressive from a technical standpoint but throwing this together to have it online in record time should not have to come at the expense of the local population for the arbitrary bragging rights of a billionaire.

Patrick – great video. Before you were born, there was a terrific movie (based on a book) titled: Colossus: The Forbin Project. https://www.imdb.com/title/tt0064177/ Highly recommended

There is no 9 links per server but only 8 . 1 is for management …

Musk is a shitty person and should not run companies that the USA depends on strategically, but yeah its a cool datacenter.

what what an amazing cluster and datacenter. Super cool that you got to tour it and show all of this to us!

I hope for Supermicro that they prepaid the bill, not like X/Twitter/Musk did with Wiwynn

Interesting data centre. Too bad Musk is behind it. What a jerk.

100% agreed on the Musk comments. So much god worship out there overlooking the accomplishments from Shotwell, Straubel, Eberhard, Tarpenning and countless others. Interesting article though ;-)

What would be even cooler than owning 100k GPUs would be putting out any AI products, models, or research that was interesting and impactful. xAI is still-born as a company because no researcher with a reputation to protect is willing to join it, the same reason Tesla’s self driving models make no significant progress.

Why so many guys hate Mask. Anyway, his gaint XAI DC is amazing project.

> There is no 9 links per server but only 8 . 1 is for management …

On each GPU node: one 400GbE link for each of 8 GPUs, plus another 400GbE for the CPU, plus gigabit IPMI.

Good article, wondering why did the choose x86 based CPUs instead of grace CPU’s?

To Skywalker: I guess it’s most likey caused by schedule(H100 while not blackwell SKU) and X software environment.

With Elon Musk the news are fun and interesting to watch ! :D

What he is doing is amazing !

The most impressive part is that they will soon double that capacity with the new Nvidia H200 batch deployment.

What the Video didn’t explain is whether DX cooling or ground water to water, air to water heat exchangers, etc are used for liquid cooling. I’m not surprised very limited information was given.

