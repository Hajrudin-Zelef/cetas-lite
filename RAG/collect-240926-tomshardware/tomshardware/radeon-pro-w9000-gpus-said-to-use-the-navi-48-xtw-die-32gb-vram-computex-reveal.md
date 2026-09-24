---
id: collect-240926-tomshardware/tomshardware/radeon-pro-w9000-gpus-said-to-use-the-navi-48-xtw-die-32gb-vram-computex-reveal
title: "radeon-pro-w9000-gpus-said-to-use-the-navi-48-xtw-die-32gb-vram-computex-reveal-"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Nvidia"]
dates: []
keywords: ["compute", "gpu", "gpus", "amd", "blackwell", "consumer", "memory", "nvidia"]
source: docs/RAG/clean_en/tomshardware/radeon-pro-w9000-gpus-said-to-use-the-navi-48-xtw-die-32gb-vram-computex-reveal-.md
source_anchor: ""
source_lines: [1, 59]
sha256: 64a854dbabd435f04bce957c054f5635921847a17820ce6421e0889f0d60db13
---

# radeon-pro-w9000-gpus-said-to-use-the-navi-48-xtw-die-32gb-vram-computex-reveal-

<!-- source: https://www.tomshardware.com/pc-components/gpus/radeon-pro-w9000-gpus-said-to-use-the-navi-48-xtw-die-32gb-vram-computex-reveal-suggested -->

AMD would reportedly be preparing to launch its RDNA 4 workstation GPU offerings for desktops, presumably under the Radeon Pro W9000 family. As put forward by Hoang Anh Phu, who frequently obtains inside scoops, AMD is considering using the Navi 48 XTW die for its top-end SKUs, paired with 32GB of video memory, likely GDDR6. As always, this leak shouldn't be taken as definitive, but there's likely some truth to it given the proximity of Computex next month, followed by AMD's Advancing AI event in June.

Radeon PRO GPUs are aimed at workstation setups, rivaling Nvidia's (former) Quadro or (now incumbent) RTX PRO offerings for prosumers. These graphics cards bridge the gap between consumers and server domains, for applications like AI, HPC, DCC, CGI, CAD, VR/AR, and the list goes on.

It seems that AMD is sticking to more conservative figures for its flagship workstation offerings this generation. That's somewhat expected since Navi 48 (356mm<sup>2</sup>) is in the same ballpark as GB203 (378mm²), found in the RTX PRO 4500 Blackwell. Nvidia's top-end GB202 at 750mm<sup>2</sup>, is home to the RTX PRO 6000 Blackwell featuring a massive 96GB frame buffer.

Navi 48, with its 256-bit interface, enables either 16GB of memory (via eight 32-bit channels) or a theoretical maximum of 32GB in clamshell mode, which is the exact configuration being reported here.

Internally, AMD segments each die into XL, XT, and XTX counterparts; each reflecting the degree to which the die's hardware resources are enabled. Based on the available data, Navi 48 wields a total of 64 Compute Units (CUs), a configuration already present in the Radeon RX 9070 XT (Navi 48 XTX).

However, we cannot infer full-enablement just from shader counts, as AMD could still have other IP blocks disabled on the consumer RX 9070 XT, reserving the full-fat die (likely Navi 48 XTW) for its professional Radeon PRO counterpart(s). It stands to reason that there must be some discernible features between Navi 48 XTX and Navi 48 XTW to justify the different designations.

Despite the per-CU improvements with RDNA 4, this supposed Radeon PRO W9070 will likely be outpaced by the W7900 in memory-intensive tasks. Adding salt on this wound, RDNA 4 still remains unsupported by AMD's ROCm platform. Nonetheless, the leaker is suggesting a launch at an upcoming event in Taiwan, likely at Computex. Here's hoping we'll learn more about broader ROCm support either next month or at AMD's Advancing AI event in June.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

*Follow* *Tom's Hardware on Google News* *to get our up-to-date news, analysis, and reviews in your feeds. Make sure to click the Follow button.*

- 
Reply
 Yup. This is probably the limit, for RDNA4. I think no higher-capacity GDDR6 dies will be made. Also, expect the memory clock speed to drop, slightly.The article said:Navi 48, with its 256-bit interface, enables either 16GB of memory (via eight 32-bit channels) or a theoretical maximum of 32GB in clamshell mode, which is the exact configuration being reported here.
 
 At least their PCIe 5.0 will enable multiple GPUs to communicate faster.
 
BTW, I fully expect they'll remedy the ROCm situation for RDNA4, and within a matter of months. They're obviously working on it - just didn't get it across the line on time.
- 
Reply
 24 GB won't happen, since I think there are no 24 Gb GDDR6 dies.YSCCC said:I wonder if they come up with a 9070XT 24gb or 32gb varient, will it kill off the 5080 also
 
As for 32 GB, we'll have to see what this new card costs. But, like I said, you often see reduced memory speeds on cards using GDDR in "clam shell" config.
- 
Reply
 While im annoyed to some degree on my NON 1000usd flagship AMD gpu this year so ive stuck with my 7900xtx ..YSCCC said:I wonder if they come up with a 9070XT 24gb or 32gb varient, will it kill off the 5080 also
 
 I dont see much of a point at this point to going a 32gb gaming 9070xt 32gb is kinda WAY WAY overkill..
 
 Now if AMD dropped a 24 or 32 GB 9080xtx or something sure shut up and take my money ..
 
 But to truly compete ( as ive seen in vids ) they would need to change the whole die etc etc which i dont think is viable ..
 
 Ive noticed over the many comparison tests ive seen with the 7900xtx VS 9070xt that RT is much better with the 9070xt in most cases , raster is still in the 7900xtx favor and in the odd case the 16gb Vram does hold back the 9070xt in a very small amount of games VS 24gb 7900xtx.
 
 With all the miss steps Nvidia has done this gen and the 9070xt skyrocketing sales for the most part ( despite my NO decent flagship this gen bitterness) they have really done well so far ..
 
Fingers crossed they dont offer ANYTHING 8GB !!
- 
Reply
 I was just thinking by the bus width, it could be compatible with GDDR6X, and with RTX 5000 series are all switching to GDDR7 there could be potential for the 4090 style config? just saying it could become a 7900xtx with better RT and slightly worse rasterbit_user said:24 GB won't happen, since I think there are no 24 Gb GDDR6 dies.
 
As for 32 GB, we'll have to see what this new card costs. But, like I said, you often see reduced memory speeds on cards using GDDR in "clam shell" config.
- 
Reply
Well, I've never seen any info or reason to believe existing RDNA4 chips are capable of more than 256-bit. They cancelled the flagship model, like a year ago. I doubt they support GDDR6X, but I suppose it's possible they added support and then decided to ship the RX 9070 XT with regular GDDR6.YSCCC said:I was just thinking by the bus width, it could be compatible with GDDR6X,
- 
Reply
 I’m quite far from poor but the $700 starting point for 70 tier performance has totally run me out of this generation. I’ve always bought $500 GPUs and that’s what I’ll continue to do.hannibal said:Well poor people need gaming setup too...
Not all can affor $1000 middle range gpus...
