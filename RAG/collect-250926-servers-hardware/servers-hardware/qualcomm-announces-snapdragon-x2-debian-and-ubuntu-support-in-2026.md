---
id: collect-250926-servers-hardware/servers-hardware/qualcomm-announces-snapdragon-x2-debian-and-ubuntu-support-in-2026
title: "qualcomm-announces-snapdragon-x2-debian-and-ubuntu-support-in-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia", "Qualcomm"]
dates: []
keywords: ["agentic", "agents", "amd", "benchmark", "compute", "disclosure", "intel", "llama", "llama.cpp", "nvidia", "vllm"]
source: docs/RAG/clean4/qualcomm-announces-snapdragon-x2-debian-and-ubuntu-support-in-2026.md
source_anchor: ""
source_lines: [1, 28]
sha256: 21c5d0d2708baaa543b8c43f2339f66c78c918a2b6d2b2db3100875b304eeaf6
---

# qualcomm-announces-snapdragon-x2-debian-and-ubuntu-support-in-2026

Today, Qualcomm announced that it will (finally) support Linux on Snapdragon, with a few caveats. At the same time, the path forward is close, with both Debian and Ubuntu support planned for 2026. This is a highly anticipated feature, so it is one that we are excited about since we run a lot of Ubuntu at STH.

As a quick disclosure, Patrick, our Chief Analyst, is at the Snapdragon Summit, and Qualcomm paid for his travel there, but we are doing this piece on our own.


## Qualcomm Announces Snapdragon X2 Debian and Ubuntu Support in 2026

Qualcomm has been doing a lot of work with Microsoft to get its processors working with Windows. Still, there are many customer segments and geographies that prefer Ubuntu Linux. These days, a lot of AI development and embedded development is being done on Linux, so the lack of Snapdragon Linux support is notable.

Qualcomm says it will focus on systemd-boot for UEFI. It will also focus on its I/O for devices like USB and PCIe, as well as lower-speed UART, I2C, and SPI. It will also support FastRPC driver upstreaming for its Hexagon NPU, which is a big part of Qualcomm’s go-forward AI strategy. The company is banking on AI agents running on compute ranging from earbuds and glasses to phones to PCs, with varying levels of compute capabilities.

Beyond that support, Qualcomm also needs to work on power management. Some other platforms we have tried “support” Linux, but their power management is not implemented well, or at all. A big part of Qualcomm’s value proposition is tied to its power efficiency, so saying that it will work on power management as part of its initial Linux support is important.

Another interesting one is the graphics and display support. Qualcomm actually has a very large gaming ecosystem since it powers so many mobile devices. Bringing graphics support to Linux has become very important for companies like NVIDIA and AMD, as their devices have become popular in the AI boom. It is good to see Qualcomm is working on this, but that brings us to our next point: OS support.

Qualcomm said it is starting with Debian and will quickly support Ubuntu as well, with Canonical expressing support at the Snapdragon Summit. Now the big one, especially for those thinking about the graphics possibilities. There is no SteamOS support that has been announced yet. Also, the Red Hat ecosystem sounds like it is not part of the 2026 initial push.

Another caveat is that the support is for the current-generation Snapdragon X2, so it is not focused on going back to the Snapdragon X1 platform.

## Final Words

Now the big question: will AgentSTH V7, our agentic AI CPU benchmark suite, run on Snapdragon X2 and Ubuntu-like chips from NVIDIA, AMD, Intel, and Arm? If so, how would a Snapdragon X2 perform compared to others? Hopefully we will get some answers later this year. Still, it is exciting to get another platform on Linux.

Finally! Without good Linux support, these machines are totally irrelevant to me and possibly lots of other people who might want to use them for development. For people who dislike MacOS, there really aren’t many good ARM based laptops available if you don’t want to be stuck with Windows.

Unless they offer full llama.cpp/vllm support and more than that poor 32gb unified RAM (sweetspot seem 64gb on the lower side and 96GB at the mid), the Strix Halosl and RTX/DGX Spark laptops/mini-PCs plus the Mac Mini/Studios will keep eating that developer slice Qualcomm is targeting…

Robust Linux support is crucial for the success of this product line, and I agree with others like Cristian that 64 and 96 GB RAM options must be available. 32GB is usable for AI workloads but just barely, we need to be able to run larger models with full context to get the most out of our equipment.
