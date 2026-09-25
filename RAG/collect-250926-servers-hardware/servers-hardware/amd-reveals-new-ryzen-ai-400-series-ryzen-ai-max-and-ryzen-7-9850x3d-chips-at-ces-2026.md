---
id: collect-250926-servers-hardware/servers-hardware/amd-reveals-new-ryzen-ai-400-series-ryzen-ai-max-and-ryzen-7-9850x3d-chips-at-ces-2026
title: "amd-reveals-new-ryzen-ai-400-series-ryzen-ai-max-and-ryzen-7-9850x3d-chips-at-ces-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "compute", "consumer", "gpu", "lpddr5x", "memory", "pricing"]
source: docs/RAG/clean4/amd-reveals-new-ryzen-ai-400-series-ryzen-ai-max-and-ryzen-7-9850x3d-chips-at-ces-2026.md
source_anchor: ""
source_lines: [1, 58]
sha256: 6dcbb94e8f07a160e8715b6df4a381451e2bac6b294eefe38306aa572f4e5c7e
---

# amd-reveals-new-ryzen-ai-400-series-ryzen-ai-max-and-ryzen-7-9850x3d-chips-at-ces-2026

Kicking off AMD’s slate of consumer announcements for this year’s CES trade show, the company brought to the shows several new chip SKUs for the consumer market. Altogether the company is launching two new Ryzen AI Max+ processors for AI developers, a new desktop Ryzen 9000X3D chip for gamers, and for the mobile market a while new family of SKUs: the Ryzen AI 400 series.

Looking at AMD’s hardware development calendar, the company finds itself in the middle of its hardware release cycle – a pretty typical occurrence given its roughly 2-year development cadence for new chip architectures. As a result, AMD does not have much up their sleeve for new consumer hardware right now: they are technically still ramping the Zen 5 launch as it is. Consequently, CES 2026 for AMD is an opportunity to launch a more curated selection of additional SKUs based on its existing hardware.

Though with that said, even AMD is not immune to the needs of its OEM customers, who ideally want to have new hardware to sell every year. As a result, we are also getting a whole new family of chips based on AMD’s existing Strix Point and Kracken Point silicon, which will be sold as the Ryzen AI 400 series.


## Ryzen AI 400 Series

We will start with the biggest announcement first, which is the launch of the Ryzen AI 400 series. For regular hardware followers, this is essentially the usual off-year hardware refresh, where AMD releases a suite of new chip SKUs with slightly improved clockspeeds while incrementing the product generation digit by one. In most years this is unremarkable (if not a bit annoying), but it keeps the OEMs who are buying these chips happy.

For 2026, AMD is giving the Ryzen AI 300 series chips various specification bumps, giving us the Ryzen AI 400 series. But with an added twist, while the Ryzen AI 300 series chips were exclusively limited to laptops and other systems using soldered-down CPUs, AMD is formally releasing the Ryzen AI 400 series for desktops as well, offering a socketed desktop option for their mono-die Zen 5 APUs for the first time.

With regards to chip specs, the big change here is that peak clockspeeds on the CPU cores, GPU cores, and NPU cores have all been boosted. So every Ryzen AI 400 chip is more powerful than its predecessor in some respect – and with some SKUs, in *every* respect. As well, AMD has validated the memory controller on some of these SKUs for higher clockspeeds, allowing them to hit LPDDR5X-8533, versus the AI 300 series’ peak speed of LPDRR5X-8000.

Unofficially, I hear that the platform name for this generation is Gorgon Point – a name that has been getting thrown around for a year now. But the hardware itself is (at best) a newer revision of the familiar Strix Point and Kracken Point hardware that we have seen AMD shipping for mobile customers all through 2025, including its mix of Zen 5 and Zen 5c CPU cores, RDNA 3.5 architecture, and XDNA 2 NPU architecture.

Altogether, AMD is launching a 7 chip stack for the Ryzen AI 400 family, replacing the 6 chip stack of the current Ryzen AI 300 family. Most of these are one-for-one replacements of 300 series parts, but there are a few xx5 parts that shuffle things around for the lower-end chips in the stack.

The flagship chip is the Ryzen AI 9 HX 475, which is getting a modest clockspeed boost across the board, with AMD slightly increasing the CPU, GPU, NPU, and memory clocks. Notably, this is the sole chip SKU to get the faster 60 TOPS NPU configuration, with AMD seemingly set on ensuring that they are not overtaken in NPU performance by rivals this year.

Meanwhile bringing up the bottom of the stack is the Ryzen AI 5 430, a quad core part that, compared to its predecessor, just about see its GPU performance double as AMD has increased the part to 4 RDNA 3.5 graphics cores (CUs). This eliminates the 2 CU option from AMD’s lineup entirely; no Ryzen AI 400 chip has fewer than 4 CUs.

Based on AMD’s published specifications, it looks like the 430, 435, and 445 chips are all using the Krackan Point silicon. Meanwhile the 475, 470, and 465 are all using Strix Point silicon. The one potential outlier here is the 450: the chip configuration matches Krackan Point, but it is the sole Krackan-level chip validated for faster LPDDR5X-8553 memory speeds. So AMD may be using this SKU to sell salvaged Strix Point chips, rather than validating Krackan Point for higher memory speeds.

Finally, AMD has not clarified their plans for socketed desktop chips at this time. So while we know they are coming, when and with what SKUs/configurations remains to be seen. Right now the initial focus is on the mobile market, where this silicon has already been residing for the past year.

On which note, Ryzen AI 400 laptops will be available almost immediately. AMD says that laptops with the new chips should hit store shelves this quarter, making for a relatively quick rollout for the latest generation of mobile chips.

## Ryzen AI PRO 400 Series To Launch as Well

With less fanfare, AMD is also announcing the Ryzen AI PRO 400 series this evening as well. These chips will be coming later in the quarter, and will serve as AMD’s enterprise-equivalent SKUs of the consumer Ryzen AI 400 chips.

As with past generations, we are looking at the same silicon with its PRO enterprise features enabled; so at the same product number there is no difference in terms of configuration or clockspeeds. AMD is forgoing a low-end 430 chip, however: the Ryzen AI Pro 400 series starts with the Ryzen AI PRO 435.

## AMD Adds Ryzen AI Max+ 392 & 388 For AI Developers

Meanwhile for AI developers and other customers who need a compact chip with a powerful integrated GPU and extra memory bandwidth, AMD is augmenting the ranks of the Ryzen AI Max family with two additional SKUs. These are the Ryzen AI Max+ 392, and the Ryzen AI Max+ 388.

Joining the three existing Max chips, these new SKUs are interesting configurations that forgo the balance of the initial chips in favor of a more lopsided, GPU-heavy configuration – and hence the Max+ branding. Both new chips have a fully-enabled iGPU, with 40 CUs and a peak GPU performance of 60 TFLOPS, just like the flagship Max+ 395. What they give up instead is CPU performance: the 392 is 12 Zen 5 CPU cores, and the 388 is just 8 CPU cores. Put another way, these are versions of the existing Max 385 and 390 chips, but with the full GPU.

Otherwise, the new Ryzen AI Max+ chips are cut from the same Strix Halo silicon as the rest of the family, complete with support for up to 128GB of LPDDR5X memory. And for AMD, that is quite okay, as the Strix Halo chips have been a hot item among AI developers who need a local GPU with access to hundreds of gigabytes of memory. Officially, gaming is still on the marketing menu as well, though signs so far have pointed to a tepid response for Ryzen AI Max gaming PCs.

## Gamers Get Ryzen 7 9850X3D – A Higher Clocked 9800X3D

But for gamers who still need a fast CPU – just without a powerful iGPU bolted to it – AMD’s final consumer hardware announcement of CES 2026 is the Ryzen 7 9850X3D. This, in a nutshell, is a higher clocked version of AMD’s existing Ryzen 7 9800X3D, AMD’s incredibly popular V-cache/X3D-equipped chip for the gaming market.

Joining the existing 9800X3D, this newer SKU fixes about the only complaint one can make about the 9800X3D: Its peak clockspeed. At launch, AMD reserved its best silicon and highest clockspeeds for its 9000X3D chips for the flagship 16 core 9950X3D, which topped out at 5.7GHz versus a more restrained 5.2GHz for the 9800X3D. But, looking to mint an even faster chip for gaming – and one that should unambiguously signal which X3D chip is specifically for gamers – AMD is cranking up the clockspeed on this 8 core part, with it now topping out at 5.6GHz.

The 8% boost in peak clockspeeds is not a massive change overall – it is technically not even the highest clockspeed for an AMD Zen 5 desktop chip. But it means that gamers who want peak clockspeeds no longer need to buy a 12 or 16 core Ryzen 9000X3D chip – all of which means they can forego the additional core scheduling and core parking complexity that multi-CCD Ryzen chips bring. It’s something of a niche market, but as the complexities for multi-CCD chips have never completely gone away, this finally gives AMD a part for customers who just want one really fast Zen 5 CCD with a lot of L3 cache.

AMD has not revealed any pricing information for the Ryzen 7 9850X3D. But the chip is expected to become available later this quarter.

## ROCm 7.2 Launches with Ryzen AI 400 Support

Finally, while not strictly a hardware announcement, one of AMD’s software-related announcements is worth taking note of. The company will be releasing the next point update for their ROCm software stack later this month. ROCm 7.2 will bring official support for Ryzen AI 400 to the software stack for both Windows and Linux, marking a very quick turnaround time for adding a new processor lineup to AMD’s GPU and AI compute software stack.

Admittedly, AMD is really just adding support for additional SKUs of existing silicon to the ROCm suite. But it is none the less a small milestone for the ROCm development group, as their turnaround time has not always been this good.
