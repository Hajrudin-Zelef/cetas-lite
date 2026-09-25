---
id: collect-250926-servers-hardware/servers-hardware/the-amd-instinct-mi350p-is-a-hbm-pcie-ai-accelerator-that-has-been-all-over
title: "the-amd-instinct-mi350p-is-a-hbm-pcie-ai-accelerator-that-has-been-all-over"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["accelerator", "amd", "blackwell", "compute", "fp4", "fp8", "gpu", "gpus", "inference", "memory", "nvidia"]
source: docs/RAG/clean4/the-amd-instinct-mi350p-is-a-hbm-pcie-ai-accelerator-that-has-been-all-over.md
source_anchor: ""
source_lines: [1, 45]
sha256: ea0ca43b64b6be8b813666cd66aadb87c2413df92941d94e822bd152cdb8d15f
---

# the-amd-instinct-mi350p-is-a-hbm-pcie-ai-accelerator-that-has-been-all-over

The AMD Instinct MI350P has been all over trade shows over the past two months. If you saw our recent AMD Intros Instinct MI350P Accelerator: CDNA 4 Comes to PCIe Cards piece, that was extremely popular. Now we are seeing them at Dell Tech World, HPE Discover, Computex 2026, and more. Over the past few weeks, we have learned a bit more about what is driving this. Here is a quick explainer video:

Also, we have to say this is sponsored. Dell and HPE sponsored our travel to the show. Also, AMD gave us access to see some systems. Let us get to it.

## AMD Instinct MI350P Key Specs and NVIDIA Comparison

Specs between GPUs these days can be rough to compare. One of the reasons is just that manufacturers often quote different specs, and in different ways. Here is a best-effort view of what this hopefully looks like for some of the high-level features among the three main PCIe GPUs you might buy today.

The first reason folks are excited about the MI350P is its memory capacity. Realistically, there is a slight edge over the H200 in memory capacity with 141GB over 144GB, but that is probably not the big one. Perhaps the big difference is that this is a more modern GPU while the H200 NVL is a Hopper generation. For the more modern Blackwell generation, the current go-to GPU is the NVIDIA RTX Pro 6000 Blackwell Server Edition GPU. That only has 96GB of capacity and uses GDDR7. NVIDIA has been pushing this since it is Blackwell, which sports a newer architecture and does not need heavily supply-constrained HBM3E memory. One plus side of the Blackwell card is that it supports RT cores for mixed workloads. Still, in the world of inference where folks try to generally push towards more memory capacity and bandwidth while working with the most compact numeric format they can to maximize memory usage, the RTX Pro 6000 Blackwell Server Edition is a trade-off.

On numeric formats, we see the impact. The Hopper generation was really before the FP6/ FP4 push. There are no published FP6 numbers on the spec page for the Blackwell card. Hopper does not have the lower precision hardware support. The MI350P has an edge both at FP4 and FP6 published figures. Another challenge is that some casually thrown-around numbers are listed as sparse when most are more constrained by dense numbers these days. Then there are peak versus delivered numbers. Here is putting all of those in one table since they tend to exist in different places and in different forms.

AMD delivered performance is very solid. The big win, however, is taking advantage of MXFP6, since you get a lot of performance and a format between FP8 and FP4. Even in the FP4 world, you get more performance. Something else to remember is that if you are using FP4 or FP6 for inference instead of FP8, you can fit more into a card’s memory footprint.

Next, video decoding capabilities are important because folks often want to work not just on text and tokens, but also on images and video feeds. Here, the specs pages are quite different, so here is another best effort chart:

If you are wondering about the name, the MI350P is related to the MI350X. It is roughly half of the OAM form factor MI350X in terms of compute, memory, and power. My best sense is that AMD heard the need for a PCIe-based accelerator, then looked at the MI350X and immediately saw the issue of trying to fit that on a PCIe card. It simply uses too much power for a standard PCIe CEM card form factor. The answer was to effectively use half of the OAM MI350X version for the PCIe card MI350P version.

In all three cases, the modern PCIe GPU cards are 600W passive-cooled cards. Here is a shot of the AMD Instinct MI350P from Computex 2026.

Here is an image Ryan used in his piece a few weeks ago that gives you a sense of what lies beneath the airflow shroud.

Like NVIDIA cards, the power connector is on the opposite side of the card (front) from the I/O plate.

Like the H200 NVL, the MI350P has no video outputs.

Something that is different is that we have started to see these in systems all over. Let us check out some examples from the past few weeks next.

While the popular focus these days is low-precision math for AI, these AMD accelerators do not skimp on the double precision floating point used in science and engineering. As a result, the MI350Ps are just a well suited for numerical computing as they are for making chatbots.

Such versatility–not shared with Nvidia’s Blackwell–should allow the AMD accelerators to hold value over a longer time and be repurposed if the AI bubble bursts.

It looks like I may have a long wait before used MI350P boards appear on eBay at my price.

The tradeoff with trying to be general is that by including a wide variety of hardware, every use case has a substantial portion of the chip that isn’t as useful to them.

At $18k-$25k per AIC, I would expect demand for these to be low. At that price, just buy a chip tailored to your particular use case.

> Like NVIDIA cards, the power connector is on the opposite side of the card (front) from the I/O plate.

Having the power connector not on top of the card just means the cards follow the official PCI specification with respect to max height, width, and length.

It’s nice to see some companies following standards. Standards enable interoperability and enhance competition on markets.

The PCIe form factor could make this easier to deploy than tightly integrated accelerator platforms. Power, cooling and software support will probably determine whether that flexibility matters in practice.
