---
id: collect-240926-tomshardware/tomshardware/intel-s-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery-1
title: "intel-s-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Intel"]
dates: []
keywords: ["intel", "amd", "compute", "gpu", "gpus", "memory", "panther lake", "power delivery"]
source: docs/RAG/clean_en/tomshardware/intel-s-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery.md
source_anchor: ""
source_lines: [1, 56]
sha256: bebf7c475fe46877540c729eb2b590b3494527cba25a25b7770c4ed2e39c1273
---

# intel-s-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery

<!-- source: https://www.tomshardware.com/pc-components/cpus/intels-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery-for-its-igpu-leaker-claims-beefy-integrated-graphics-could-require-two-vccgt-phases-for-12-xe3p-cores -->

Intel's next-gen Nova Lake family of desktop CPUs is shaping up to be the company's most exciting launch in years, with a major performance leap expected across the board. A new leak from *Jaykihn* now claims that one of the SKUs from this lineup will have an iGPU so strong that it will require 65W of power delivery to achieve full performance. It will also apparently need two VCCGT phases to handle said power.

The SKU in question is a 16-core part composed of 8 P-cores, 4 E-cores, and 4 LP-E cores, while the iGPU is said to have 12 Xe3P cores. Just as a reminder, Xe3, also known as "Celestial," is an enhanced, optimized version of the Xe3 (Battlemage) graphics architecture that already debuted on Panther Lake. Similarly, Nova Lake is expected to use Coyote Cove P-cores and Arctic Wolf E-cores (and LP-E cores).

Top-end Panther Lake SKUs come equipped with the Arc B390 iGPU, which also has 12 Xe3 cores, so this Nova Lake SKU with 12 Xe3P cores should perform even better. Especially when you consider the thermal and power headroom at its disposal. AMD's desktop APUs, the Ryzen G-series, are usually 65W parts as well, but that's the TDP for the entire chip, not just the integrated graphics.

Requiring 65W of dedicated power delivery via two VCCGT phases would constitute a kind of top-end iGPU performance we haven't seen before. In fact, this rumored Nova Lake chip is arguably veering into Strix Halo territory where the up to 40 Compute Units on flagship SKUs can sip around 70W-80W of power. However, that's still a mobile part with an integrated SMU that can dynamically allocate power between the CCD and iGPU.

| Nova Lake-S Rumored SKUs |  |  |  | 
|---|---|---|---|
| SKU | Core Config (P+E+LP-E) | bLLC | TDP (Unlocked/Locked) | 
|---|---|---|---|
| 52 Cores (dual-tile) | (8+16)+(8+16)+4 | 288MB | 175W | 
| 44 Cores (dual-tile) | (8+12)+(8+12)+4 | 264MB | 175W | 
| 28 Cores | 8+16+4 | 144MB | 125W | 
| 28 Cores | 8+16+4 | - | 125W / 65W | 
| 24 Cores | 8+12+4 | 132MB | 125W | 
| 24 Cores | 8+12+4 | - | 125W / 65W | 
| 22 Cores | 6+12+4 | 108MB | 125W / 65W | 
| 22 Cores | 6+12+4 | - | 125W / 65W | 
| 16 Cores | 4+8+4 | - | 65W / 35W | 
| 12 Cores | 4+4+4 | - | 65W / 35W | 
| 8 Cores | 4+0+4 | - | 65W / 35W | 
| 6 Cores | 2+0+4 | - | 65W / 35W | 

Nova Lake-S would likely use a traditional desktop rail split where the motherboard's VRMs must deliver up to 65W via two separate dedicated VCCGT phases (since one wouldn't be sufficient) to power the integrated graphics tile independently from the VCCCore CPU phases. That's unprecedented territory, and it serves as just one of many rumors from the Nova Lake launch that have us excited for the competition that's brewing for next year.

Current reports pin both AMD and Intel's next-gen releases to show up at CES 2027 instead of later this year. Therefore, take everything with a grain of salt; a lot could change between now and then, and we don't have confirmation on any SKU from Intel. The last time we saw Intel put out a strong desktop APU was 2015's Core i7-5775C, so a product like this has been a long time coming.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Hassam Nasir is a die-hard hardware enthusiast with years of experience as a tech editor and writer, focusing on detailed CPU comparisons and general hardware news. When he’s not working, you’ll find him bending tubes for his ever-evolving custom water-loop gaming rig or benchmarking the latest CPUs and GPUs just for fun.

-
The prospect of this is genuinely exciting to me if the SKU is priced right and there at least some affordable motherboards with appropriate VRM. Entry level discrete graphics is effectively dead since the only parts that make sense under the B570/5050 are used purchases. By allowing the GPU to have free reign power wise it should maintain boost clocks extremely well. Intel should be sticking with a large cache (for low core count) design which will mitigate slower desktop memory speeds if there's no reasonably priced high speed DDR5 available. This all combined with an improved architecture I really wouldn't be surprised to see somewhere around desktop RTX 3050 performance. This would be a great place to be as either a starter gaming system or small size/low powered system.Reply
 
 The biggest problem with all better integrated graphics on desktop to date has been that they're tied to upper CPU SKUs. This should sit in the lower/middle of Intel's desktop stack.
 
The biggest problem with AMD APUs has been that they're compromised should you upgrade to discrete. With less cache and PCIe connectivity it means not only are you no longer using the integrated graphics you paid for but the CPU you're left with is worse than the regular core count equivalent. This for sure won't have a compromised cache setup and there's no reason it should have compromised PCIe connectivity either.
- 
Reply
 It's an afterthought. AMD puts the mobile die onto the socket (although desktop APUs could be made with chiplets soon), you get whatever the compromises were. Hopefully the user understood the compromises going in, or got it at a budget price.thestryker said:The biggest problem with all better integrated graphics on desktop to date has been that they're tied to upper CPU SKUs. This should sit in the lower/middle of Intel's desktop stack.
 
 The biggest problem with AMD APUs has been that they're compromised should you upgrade to discrete. With less cache and PCIe connectivity it means not only are you no longer using the integrated graphics you paid for but the CPU you're left with is worse than the regular core count equivalent. This for sure won't have a compromised cache setup and there's no reason it should have compromised PCIe connectivity either.
 
 Competition is key. AMD phoned it in with high MSRPs, skipped/late generations, and now an OEM-only Krackan Point lineup (no Strix models with 880M/890M). The desktop APU is often derided for poor price/performance, and that's been 100% tied to AMD's decisions. Now we could see a different way of doing things from Intel.
 
If AMD wants to challenge 12 Xe3P graphics performance, they're going to have to use Medusa Halo Mini (24 CUs RDNA5, 128-bit). CPU performance should technically be in the lower/middle of AMD's desktop stack, if it is packing 4+8+2 cores (no optional compute chiplets added), which should be slower than an Olympic Ridge CPU with a single 12-core CCD. Probably with 10 MiB graphics L2 cache to challenge Intel's ~16 MiB L2 cache, which was a big part of Panther Lake's secret sauce.
- 
Reply
 The i7-5775C was an experiment for Intel, with eDRAM + large iGPU only carried forward to a handful of mobile processors and BGA desktop packages instead of socketed desktops. Kaby Lake G was another weird experiment, obviously not using Intel's own graphics.Nota ReAlperson said:Intel has made some powerful igpus since the i7 5775c. The iris pro p580 and the Kaby lake g with amd vega graphics come to mind.
 
