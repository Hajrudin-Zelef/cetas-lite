---
id: collect-240926-tomshardware/tomshardware/amd-ryzen-5-8600g-review-the-new-value-igpu-gaming-champion-2
title: "amd-ryzen-5-8600g-review-the-new-value-igpu-gaming-champion"
domain: tomshardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "gpu", "compute", "cost", "gpus", "memory", "pricing"]
source: docs/RAG/clean_en/tomshardware/amd-ryzen-5-8600g-review-the-new-value-igpu-gaming-champion.md
source_anchor: ""
source_lines: [76, 123]
sha256: 64acc54669319bff889fea240dfd7556af8f0744f4aef38c0fb67da109e764a4
---

# amd-ryzen-5-8600g-review-the-new-value-igpu-gaming-champion

The standard Ryzen 7000 desktop processors expose 24 usable PCIe 5.0 lanes, but the Ryzen 8700G and 8600G only expose 16 usable PCIe 4.0 lanes, which is a big step back on available bandwidth due to both fewer lanes and a reduction in PCIe interface speed. However, the x8 PCIe 4.0 connection to the CPU won't be a constraint with current GPUs. Besides, these chips aren't really meant to be used with a discrete GPU. The system has two x4 NVMe SSD connections available, which is sufficient connectivity for a lower-end platform.


The previous-gen Ryzen 5 5600G was the killer value APU because it provided up to 96% of the performance of the higher-end model, but there are differences between this generation. 

The previous-gen Ryzen 7 5700G and Ryzen 5 5600G come with Vega graphics with either seven or eight compute units (CUs). However, only one compute unit (CU) and 100 MHz separated the graphics engines on the prior-gen models. In contrast, the new Ryzen chips have a much larger gap, with the 8700G's Radeon 780M iGPU having 12 CU compared to the 8600G's Radeon 760M with eight CU. This resulted in a larger performance gap between the two models than we saw with the prior gen, but since memory bandwidth is the primary constraint for the iGPUs, overclocking the 5600G does help level the playing field, as you'll see on the following pages.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Current page: The Return of the APU

Next Page AMD Ryzen 5 8600G Hyper-RX, Power Consumption, Overclocking, Test Setup
- 
>The current-gen flagship APU, the Ryzen 7 8700G, is 10% faster than the 8600G, but it costs $100 (44%) more, making the 8600G the clear value winner for this generation of APUs.Reply
 
 This is incorrect, because you do not use the APU by itself, but as a part of a system. The correct value calculation is to take the price delta of the entire system but with different APUs.
 
 Assuming a $1K SFF build, the $100 APU price diff would come out to ~10%, which is the same perf diff between 8700G and 8600G. AMD did its pricing homework. The value proposition is the same for both APUs.
 
 I agree both are niche, as their main appeal would be for small SFF (eg NUC), and in that space, mobile parts, eg MTL & AMD 780M parts, may have more functionality and be better value.
 
 
 >...the 8600G will look great if it hits the low prices seen for the 5600G along with other platform costs getting cheaper.
 
 That'll happen if you're willing to wait ~2 years until the 9600G's release.
 
 5600G pricing was very stable, staying close to its $259 launch price, and only gradually dropping after the 7000 series release on Sep'22.
 
https://camelcamelcamel.com/product/B092L9GF5N
- 
"Higher DDR5 pricing, no 8GB options"Reply
 
I am so tired of hearing about DDR5 pricing being a con. While it is more expensive, you won't have a choice for a new CPU in the Zen 5 or 15th Gen anyways. Not to mention due to the added bandwidth the 8600G averages 72% more iGPU performance at 1080p vs the 5600G. If you are looking at doing only iGPU gaming that added cost is minimal compared to extra performance. The lack of an 8GB option is also not a con. Gaming on 8GB is with a dGPU isn't good, unless it is old games. Doing it on an iGPU can kill your performance completely.
- 
I think in most cases the APU will be preferable over the CPU with the old graphics card.Reply
 
If price is your only reason to go for a system like this, and you badly want to game, the APU isn't the best solution. But case/enclosure size, heat and noise will be the primary reason to choose the APU. The fact that you actually can game on it will what makes the sale.
- 
Reply
A few years ago I played through the entirety of Star Wars Jedi: Fallen Order on a Ryzen 3200G (no dedicated GPU) and with a 4GB x 2 kit of RAM. The Series S also works fine with what is functionally only 8 GB of RAM shared between the CPU and GPU (there's technically +2 GB for a total of 10, but the extra 2 GB is clocked slow and is for the OS). Memory needs tend to be overstated.jeremyj_83 said:The lack of an 8GB option is also not a con. Gaming on 8GB is with a dGPU isn't good, unless it is old games. Doing it on an iGPU can kill your performance completely.
- 
Reply
Have you tried to do basic work on a computer with Win 10, 8GB RAM, and 2GB RAM reserved for the iGPU? It is painfully slow as you are page swapping all the time. Any computer that isn't a Chromebook should come with less than 16GB RAM now.HopefulToad said:Memory needs tend to be overstated.
- 
Reply
 My desktop is a 32GB ram machine. I also have an i7 / 32gb / 2TB nvme gaming laptop.jeremyj_83 said:Have you tried to do basic work on a computer with Win 10, 8GB RAM, and 2GB RAM reserved for the iGPU? It is painfully slow as you are page swapping all the time. Any computer that isn't a Chromebook should come with less than 16GB RAM now.
 
When I want portability on the road, I have a dell notebook that's education-oriented. 4GB ram, no fans, small CPU. It runs 10/11 perfectly well, starts up quickly and performs my work requirements (Office / Firefox / NextCloud) pretty much as quick as my main machines.
