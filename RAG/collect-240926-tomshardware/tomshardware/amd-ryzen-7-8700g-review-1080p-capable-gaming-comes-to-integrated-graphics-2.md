---
id: collect-240926-tomshardware/tomshardware/amd-ryzen-7-8700g-review-1080p-capable-gaming-comes-to-integrated-graphics-2
title: "amd-ryzen-7-8700g-review-1080p-capable-gaming-comes-to-integrated-graphics"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "United States"]
dates: []
keywords: ["amd", "accelerator", "benchmarks", "compute", "cost", "gpu", "gpus", "inference", "intel", "memory"]
source: docs/RAG/clean_en/tomshardware/amd-ryzen-7-8700g-review-1080p-capable-gaming-comes-to-integrated-graphics.md
source_anchor: ""
source_lines: [64, 124]
sha256: d634135a76e35e458a66b4e25887f0c3bb9c1e697b43e5edcf5d2c753d30d715
---

# amd-ryzen-7-8700g-review-1080p-capable-gaming-comes-to-integrated-graphics

The previous-gen 5000G series chips supported DDR4-3200 and snapped into the AM4 platform, forging a true budget system. The value proposition isn't quite as clear with Ryzen 8000G: These chips drop into the AM5 platform with DDR5-5200 memory (and plenty of overclocking headroom). The AM5 motherboard ecosystem remains pricey, and DDR5 memory is still more expensive than DDR4. The Ryzen 8000G chips support 600-series chipsets, and this class of chip is best paired with the B650 chipset, though the A620 is also attractive.

The Ryzen 7 8700G and Ryzen 5 8600G’s XDNA AI engine runs at 1.6 GHz, a 60% improvement over the inaugural Ryzen 7040 mobile series (those processors never came to the desktop PC). The XDNA engine delivers roughly the same 16 TOPS of INT8 (only) performance as the mobile variants, though it is possible some workloads could generate more performance as a result of the PC’s higher power thresholds. AMD says the chips' CPU, GPU, and XDNA engines combine to deliver up to 39 TOPS of overall AI inference performance.

The standard Ryzen 7000 desktop processors expose 24 usable PCIe 5.0 lanes to the user, but the Phoenix products only expose 16 usable PCIe 4.0 lanes, which is a big step back on available bandwidth due to both fewer lanes and a reduction in PCIe interface speed. However, the x8 PCIe 4.0 connection to the CPU won't be a constraint with current GPUs — though make no mistake, these chips aren't really meant to be used with a discrete GPU — and the system has two x4 NVMe SSD connections available. This should be enough connectivity for a lower-end platform.

The previous-gen Ryzen 7 5700G and Ryzen 5 5600G come with less powerful Vega graphics with either 7 or 8 CUs. However, only one compute unit (CU) and 100 MHz separated the graphics engines on the prior-gen models. In contrast, the new Ryzen chips have a much larger gap: The 8700G's Radeon 780M iGPU has 12 CU compared to the 8600G's Radeon 760M with eight CU. This resulted in a larger performance gap between the two models than we saw with the prior gen. However, since memory bandwidth is the primary constraint for the iGPUs, overclocking the lower-end model could help level the playing field. 

## Phoenix 2

AMD uses the Phoenix 2 die with the Ryzen Ryzen 5 8500G and Ryzen 3 8300G, meaning these two lower chips have a mix of both standard Zen 4 cores and slower density-optimized Zen 4c cores. They also don't come with the AI accelerator, so they won’t have the Ryzen AI badge on their product box. We aren't reviewing these chips yet, but they'll be under the microscope soon. Given the massive difference in architecture, it is wise to understand the differences if you're considering these as lower-cost options.

AMD’s Ryzen 5 8500G has two Zen 4 cores paired with four density-optimized Zen 4c cores, while the Ryzen 3 8300G has one Zen 4 core with three Zen 4c cores. As with Intel's E-cores, AMD's Zen 4c cores are designed to take up less space on a processor die than the regular Zen 4 cores while providing enough horsepower for less demanding tasks. This saves power and delivers more computing performance per square millimeter than was previously possible (deep dive here). But the similarities end there.

Unlike Intel, AMD employs the same microarchitecture and supports the same features with its smaller cores as it does with the larger cores. However, the Zen 4c cores do run at lower clock speeds and deliver less peak performance than standard cores. Notably, the maximum boost frequency of the Zen 4c cores is actually lower than the base frequency of the standard Zen 4 cores, which stands in contrast to Intel's approach. We'll put this arrangement to the test soon.

Both of these chips come with a Radeon 740M iGPU with a mere four CUs. The Phoenix 2 chips come with 10 usable PCIe 4.0 lanes, with four dedicated to graphics and the other six split across M.2, USB, and WiFi. AMD says the M.2 SSD should have a x4 connection with the rest of the I/O split across two lanes, but the end configuration is up to the ODM.

For now, let's take a look at how the Ryzen 7 8700G fares in our gaming benchmarks. 

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Current page: The Return of the APU

Next Page AMD Ryzen 7 5700G Hyper-RX, Power Consumption, Overclocking, Test Setup
- 
AMD Ryzen 7 5700G Hyper-RX, Power Consumption, Overclocking, Test SetupReply
 Hyper-RX -> HYPR-RX
 "The Ryzen 7 8700G utterly destroys the previous-gen Ryzen 7 8700G"
 
 This review appears to be much more favorable to the 8700G. AnandTech did all testing at DDR5-5200.
 
 Gamers Nexus found an issue with STAPM being enabled and degrading performance. That's Skin Temperature Aware Power Management, which obviously is not relevant to desktop APUs in a desktop computer.
 
 There have been good bundles in the US with 7600X/7800X3D + motherboard + DDR5. Micro Center obviously but also Newegg.
 
 https://slickdeals.net/f/17261224
 If you do live near a Micro Center, take note of the stupidly low open box prices on DDR5-6000 kits. These returns aren't happening because the memory is bad. Maybe it's because they were included in so many bundles:
https://www.microcenter.com/category/4294966965/desktop-memory-ram?storeid=181
- 
The included AI accelerator needs to be exposed in Gaming. It Will become a break through if the Ryzen 8700G GPU+NPU can be combined with discrete Graphics like Radeon RX 7600 to accelerate ray tracing effects in Gaming in the same FPS like Ray tracing effects disabled.Reply
- 
For the price of an 8700g, 32gb ddr5 ram, and motherboard, you could build something like a 5700x, 32gb ddr4, b550 board, and an RX 6600, that would destroy the IGP in the 8700g. The price is simply too high.Reply
- 
Replyjxdking said:8700G is weird.
 Only with half the L3 cache of 7700x, it doesn't perform well in game even with a dgpu.
Some of those issues may be related to what GN discovered, as mentioned earlier.
- 
Reply
 The kids in their 'VANS T-Shirts' at GameStop are of a different opinion regarding all of the hoopla being now offered on the various tech channel reviews. Their argument is that the 8700G will not in real life have the same juice or capability of a (2019) GTX 1650 mobile and which so far allowed them playing quite satisfactory 80%+ of their mostly outdated and now starkly reduced on-sale games. And no matter what AMD is now promising the new 'Phoenix' chips are being capable off!suryasans said:I will wait another year to buy Ryzen 8000G series as these APUs are still not a good value in terms of price/performance.
 
 For me these new Phoenix editions are essentially pointless as well. The three-star rating here in way telling the story! The Phoenix line I also think represents a niche product and the niche here is even smaller than with the once mighty Threadripper. Making me wonder why AMD would even bring this kind of new product to the market! Talking about niches:
 Office PC: Too much GPU performance
 Gaming PC: Much too little GPU performance
 Parents PC: Too much GPU performance
 Multimedia PC: Marginal at best
Children's PC (simple games): PossibleAccording to current rumors, Zen5 x3D will not come onto the market until 2025 and which would be my first consideration all things considered. Finally it will be curious to see in see how AMD sells or will market these new APUs through their strategic partners.
- 
Reply
 16 TOPS is weak. It has an efficiency advantage in laptops when it can be used. I don't think there's any chance that it can make upscaling or raytracing better.suryasans said:The included AI accelerator needs to be exposed in Gaming. It Will become a break through if the Ryzen 8700G GPU+NPU can be combined with discrete Graphics like Radeon RX 7600 to accelerate ray tracing effects in Gaming in the same FPS like Ray tracing effects disabled.
 
