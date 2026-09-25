---
id: collect-250926-servers-hardware/servers-hardware/nvida-introduces-rtx-spark-an-arm-soc-for-windows-pcs
title: "nvida-introduces-rtx-spark-an-arm-soc-for-windows-pcs"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Apple", "Intel", "Microsoft", "Nvidia", "Qualcomm", "TSMC"]
dates: []
keywords: ["3nm", "accelerator", "amd", "benchmarks", "blackwell", "consumer", "fp4", "gpu", "gpus", "intel", "lpddr5x", "memory"]
source: docs/RAG/clean4/nvida-introduces-rtx-spark-an-arm-soc-for-windows-pcs.md
source_anchor: ""
source_lines: [1, 69]
sha256: f7e929a73cf1318903ab2f5028620ea690c822c315f0c36e524754e14022c999
---

# nvida-introduces-rtx-spark-an-arm-soc-for-windows-pcs

Headlining NVIDIA’s hardware announcements for this year’s Computex trade show in Taiwan, the company finally revealed its long-awaited (and heavily leaked) Arm SoC for Windows PCs, the RTX Spark. Previously going under the codename N1X, the RTX Spark will mark NVIDIA’s second effort to break into the Windows market, providing a powerful chip based on their Grace Blackwell technologies with up to 20 CPU cores, a sizable GPU, and able to address up to 128GB of RAM.

The RTX Spark will be going into both high-end laptops and small form factor desktops, with NVIDIA briefly showing off devices in both form factors. Systems based on the chip are still a few months away from launch, however, with the first devices based on the RTX Spark not slated to arrive until the fall of this year.


### RTX Spark: Grace Blackwell for Consumer Systems

With Computex traditionally serving as the PC industry’s biggest trade show of the year, it has been a frequent venue for CPU announcements and teases. Working in conjunction with Windows publisher Microsoft, the two companies touted that the chip would usher in a new era of PC.

Yet even with the pomp behind of NVIDIA’s announcement, the company’s announcement only provided limited technical detail on the chip itself – and did not provide any benchmarks or other measured performance figures. In practice, it would be fair to say that this week’s announcement is closer to a teaser than a technical briefing, with NVIDIA forgoing any real discussion of how the RTX Spark compares to rival chips from Intel, AMD, or fellow Windows-on-Arm provider Qualcomm.

From a high level, NVIDIA is touting a full RTX Spark configuration as offering 20 CPU cores, which are custom-built by MediaTek. Meanwhile the GPU side of the chip is a Blackwell-based accelerator with 6144 CUDA cores and around 1PFLOP of FP4 performance, which is similar to the specifications for NVIDIA’s GeForce RTX 5070 video card. The complete chip in turn will be fed by a single LPDDR5X memory pool, with up to 128GB of RAM available.

These high-level chip specifications are also a dead ringer for NVIDIA’s existing GB10 SoC, which has been shipping in the DGX Spark (and identical OEM systems) for roughly the past year.

| **NVIDIA Spark Chips Key Specs** |  |  | 
|  | **N1X (RTX Spark)** | **GB10 (DGX Spark)** | 
| **CPU** | 20 Core Grace CPU | 10x Arm Cortex-X925 10x Arm Cortex-A725 | 
| **GPU** | Blackwell, 48 SMs | Blackwell GB20B GPU – 48 SMs | 
| **Memory** | 128GB LPDDRX5 | 128GB LPDDR5X-8533, Soldered | 
| **Manufacturing Process** | TSMC 3nm | TSMC 3nm | 

The two are so similar that at this point it is not even clear if N1X is a distinct chip (or set of chiplets), or if it is just GB10 rebranded for the consumer market. In the case of the latter, while GB10 systems cannot run Windows, the hardware is otherwise a known quantity when it comes to both CPU and GPU performance.

With that said, even if N1X is a rehash of NVIDIA’s GB10 chip it is not the only RTX Spark chip that NVIDIA will be releasing. The company’s roadmap also calls for a lower-performing version of the chip to be released, offering 400TFLOPS of FP4 GPU performance (~40% of N1X) and an undisclosed CPU configuration. This looks to be a lower-end (or at least, mid-tier) version of NVIDIA’s design, which would presumably go into smaller, cheaper, and lower-powered laptops.

Ultimately with NVIDIA being so light on the technical details for RTX Spark at Computex, we expect that the chip will get a more formal introduction later this year when it is closer to launching.

### SFF Desktops and 14 to 16-inch Laptops, Coming in The Fall

Hardware specifications aside, NVIDIA has already lined up a notable list of partners for the RTX Spark. Acer, ASUS, Dell, Gigabyte, HP, Lenovo, Microsoft, and MSI are all slated to release systems based on the chip this year. And Microsoft in particular, who has been working closely with NVIDIA to bring another Arm CPU vendor into the fold, will be releasing the Surface Laptop Ultra as their new flagship Surface laptop.

None of the vendors listed have published detailed specifications on their laptops, but NVIDIA’s press release noted that RTX Spark laptops will range in size from 14-inches to 16-inches. Furthermore, all of RTX Spark laptops will feature an aluminum chassis and a tandem OLED display.

NVIDIA had several of the closed laptops on stage for their Computex keynote, and with the aim for RTX Spark to be a premium laptop offering, all of them look quite similar to each other. Though not identical in design, all of the laptops are similar in coloring, logo placement, and more.

Meanwhile the RTX Spark will also be going into small form factor desktops. NVIDIA had a few of these on display as well, with Jensen Huang briefly holding up an MSI model. On the whole it looks nearly identical to MSI’s existing GB10 system, which given the silicon similarities between N1X and GB10 is not too surprising. Notably, these was nothing bigger than a SFF PC demonstrated, so it would seem all the systems that RTX Spark is going into will be highly integrated systems.

Overall, this does leave me with the distinct impression that NVIDIA is holding the reins very tight for the RTX Spark, and that they may be dictating designs to the OEMs rather than the other way around. If so, this would make the RTX Spark partner program a great deal like NVIDIA’s GB10 partner program, which has seen OEMs all release clones of the DGX Spark under their own branding. This has ensured a very consistent experience among all of the GB10 systems, but it also has meant that OEMs are left with little room for innovation or differentiation.

In any case, RTX Spark laptops and desktops are slated to become available this fall, with more details to come.

### A Second Shot at Windows-on-Arm

With the announcement of the RTX Spark, this marks NVIDIA’s second major effort at breaking into the Windows-on-Arm market. The company previously had a go at this over a decade ago, when Microsoft first started their Windows-on-Arm efforts alongside Windows 8 and the Tegra 3-powered Surface RT tablet. This effort was ultimately a flop, and while NVIDIA dropped out of participating in the Windows-on-Arm ecosystem for quite some time, Microsoft has continued to develop their technologies alongside current partner Qualcomm.

As a result, while NVIDIA’s second effort to secure a beachhead is far from guaranteed, they’re starting in a much better position this time around. On the software side of matters a significant number of major Windows applications have been ported to/compiled for Arm, with everything from Chrome and Firefox to Photoshop and VLC. And for those applications that are still only available as x86 binaries, Microsoft’s Prism x86 emulator recently added AVX and AVX2 support, which further improves the compatibility of the x86-on-Arm emulator.

More interesting, perhaps, will be the state of gaming on RTX Spark systems, which NVIDIA is clearly making a focus from the start. Games as a whole remain one of the largest hold-outs for releasing Arm binaries, though that is slowly improving as additional developers have jumped on board, and critical anti-cheat solutions are getting native Arm support as well. Still, games have the potential to be a rough spot for emulation since they are so performance sensitive, so it will be interesting to see what the real-world performance implications are once RTX Spark devices ship, and how NVIDIA is going to be throwing around its massive weight in the gaming space to influence more developers to release Arm-native games.

Otherwise, while NVIDIA is holding off on performance comparisons and competitive claims at the moment, the specifications for the RTX Spark chips put them in the same ballpark as AMD’s own big-chip SoC, the Ryzen AI Max series (Strix Halo) as well as Apple’s larger M5 chip configurations. Ironically, AMD’s chips have been almost entirely coopted for building DGX Spark-like SFF workstations, so with a robust laptop lineup NVIDIA is not immediately facing any direct competitors in the Windows space. All of which leaves Apple as the player to beat as far as single-chip solutions are concerned.

### A Roadmap for the Future: More Sparks to Fly

Rounding out their RTX Spark announcement, NVIDIA also made sure to underscore that while the RTX Spark was their first consumer Windows-on-Arm chip in this era, it will not be the last. The company’s chip roadmap now includes Spark chips alongside its other CPUs, GPUs, and networking chips.

The high-level roadmap is light on details, but it outlines NVIDIA developing further Spark chips based on its Vera Rubin architecture as well as the Rosa Feynman architecture that will succeed it. The roadmap also reveals that the Vera Rubin Sparks will be paired with LPDDR6 memory, while the memory tech used for the Rosa Feynman chips is unspecified.

Notably, NVIDIA’s roadmap also has both a big and small Spark chips slated for each generation – so the two RTX Spark configurations for this generation are not a one-off occurrence. They will however come in the second-half of each generation; NVIDIA’s roadmap is clear to illustrate that the Vera Rubin and Rosa Feynman Spark chips will not land until 2028 and 2030 respectively, maintaining a two-year cadence and launching roughly a year after NVIDIA’s big silicon GPUs and CPUs are released.

256 bit memory for N1X/GB10.

I’m confused, the DGX Spark has a 256 bit memory interface and LPDDR5X and a 273 GB/s memory bandwidth. Is that 600 GB/s NVLink C2C trying to say that it will have 600 GB/s of memory bandwidth, or is it going to sill have that 273 GB/s limitation?

@JL

600GB/sec is solely die-to-die bandwidth. It has nothing to do with the LPDDR5X memory bandwidth.

At this point NVIDIA has not confirmed how much memory bandwidth RTX Spark chips will have.
