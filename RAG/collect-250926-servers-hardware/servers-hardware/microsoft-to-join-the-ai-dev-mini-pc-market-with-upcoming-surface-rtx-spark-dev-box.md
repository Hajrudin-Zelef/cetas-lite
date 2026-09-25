---
id: collect-250926-servers-hardware/servers-hardware/microsoft-to-join-the-ai-dev-mini-pc-market-with-upcoming-surface-rtx-spark-dev-box
title: "microsoft-to-join-the-ai-dev-mini-pc-market-with-upcoming-surface-rtx-spark-dev-box"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Microsoft", "Nvidia", "Qualcomm"]
dates: []
keywords: ["amd", "copilot", "ethernet", "lpddr5x", "memory", "nvidia", "pricing"]
source: docs/RAG/clean4/microsoft-to-join-the-ai-dev-mini-pc-market-with-upcoming-surface-rtx-spark-dev-box.md
source_anchor: ""
source_lines: [1, 27]
sha256: 652b9230458ccaf158804d52abd7d48e65139f712dfda89062e9e8226a8ea4be
---

# microsoft-to-join-the-ai-dev-mini-pc-market-with-upcoming-surface-rtx-spark-dev-box

Between the Computex and Microsoft Build conferences this week, Microsoft and NVIDIA have been spending a fair bit of time promoting NVIDIA’s RTX Spark SoC. With NVIDIA becoming the second Arm chip vendor to join the Windows 11 ecosystem, the two companies have been eager to promote the expansion of Windows-on-Arm as well as the devices that will be built from the RTX Spark. For Microsoft’s part this includes not only a new Surface laptop – the Surface Laptop Ultra, which was introduced during NVIDIA’s Computex keynote – but as revealed by the company during Build, Microsoft is also putting together a Surface-branded RTX Spark-powered Windows development box, which they will be calling the Surface RTX Spark Dev Box.

With the new small form factor PC, this is a case where the device does exactly what it says on the tin: this is a Windows dev box based around NVIDIA’s RTX Spark SoC. Specifically, this is Microsoft’s answer to the recent spate of Linux-based AI development boxes that have been released by both NVIDIA and AMD partners over the last year, where both NVIDIA’s GB10 and AMD’s Ryzen AI Max (Strix Halo) have been showing up in all sorts of developers kits. With the unique value proposition offered by these high-end SoCs – mainly the size benefits of a single-chip solution combined with the memory capacity afforded by a large LPDDR5X memory bus – these small form factor PCs have been a hot topic and a fast seller for PC vendors who are feeling the pinch of overall slowing sales from high memory and storage prices.

Notably, this is not Microsoft’s first foray into SFF dev kits. The company released the Windows Dev Kit 2023 a few years back, however that was a budget offering based around a Qualcomm 8cx Gen3 SoC. Whereas RTX Spark systems, including the Surface RTX Spark Dev Box, are clearly positioned as anything but budget-priced.

For now, Microsoft is not publishing much in the way of hardware details on the system itself, which is consistent with the other RTX Spark devices revealed so far. But besides confirming that the systems will come up to 128GB of RAM, Microsoft has also posted that Surface RTX Spark Dev Box is designed to run within a 100 Watt thermal envelope, which gives us our first official guidance towards what the power requirements are for the RTX Spark chip inside.

Internal specifications aside, the Surface RTX Spark Dev Box is housed in an all-aluminum chassis with a rather distinctive layered design, with a larger layer stacked over a smaller layer – something akin to an inverted ziggurat. The entire top of the system is all ventilation, with Microsoft stating that there are 1000 holes altogether.

The company has also published photos of the rear of the system, confirming what I/O ports will be available. These will include a pair of USB-C ports, a USB-A port, an HDMI port, an Ethernet jack, and a combo audio jack, though the speeds of the USB and Ethernet ports have not been disclosed. Overall, this setup is quite similar to NVIDIA’s own DGX Spark box, minus a couple of USB-C ports and the Spark’s defining ConnectX-7 QSFP Ethernet ports.

Meanwhile from a software perspective, Microsoft says that they will be pre-loading the system with Windows 11 Pro and a full suite of development tools. These will include Visual Studio Code, GitHub Copilot, WSL, and PowerShell 7. By and large this is comparable to what we have seen with Linux AI dev kits released thus far, though with a distinct Windows twist.

As for the intended audience, while Microsoft is releasing the dev kit for a broad audience of any and all Windows developers, there is certainly a heavy emphasis on AI thanks in large part to the capabilities of the underlying SoC. Especially as Microsoft is also in the process of releasing a number of AI-focused software development tools and technologies, which were similarly announced at Build this week.

At present the most powerful options for local, large memory set AI development in the Windows ecosystem are systems based on AMD’s Ryzen AI Max chips. So a large memory capacity NVIDIA offering is a notable addition to the Windows ecosystem – though not one that is going to be exclusive to Microsoft hardware.

Wrapping things up, Microsoft is not announcing pricing or availability for the Surface RTX Spark Dev Box, with the company promising more information to come at a later date.

I’m curious if this will mean broader unified memory support. Probably some gotchas and ideocracies if you want it on already released parts; but when iGPUs are always feeding off system RAM anyway the main reason for them to pretend to have a hard distinction between system and VRAM has mostly been what the OS expects.

I also hope someone had a very good reason for orienting the HDMI port the way they did. I don’t think it’s actually a spec requirement but that is definitely upside down in a way that will confuse the muscle memory of approximately everyone plugging one in.

I would expect 1,000 holes on any Microsoft product. They never fail to deliver.

I really like the design of this. I wonder if the motherboard is reworked from the DGX Spark so a 2282 SSD can fit and not just the same board, different RAM and no ConnectX-7 fitted.
