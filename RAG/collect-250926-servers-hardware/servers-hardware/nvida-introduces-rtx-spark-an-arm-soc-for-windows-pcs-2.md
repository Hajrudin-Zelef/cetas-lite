---
id: collect-250926-servers-hardware/servers-hardware/nvida-introduces-rtx-spark-an-arm-soc-for-windows-pcs-2
title: "nvida-introduces-rtx-spark-an-arm-soc-for-windows-pcs"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Apple", "Nvidia"]
dates: []
keywords: ["amd", "consumer", "gpus", "lpddr5x", "memory", "nvidia", "nvlink", "rubin", "vera rubin"]
source: docs/RAG/clean4/nvida-introduces-rtx-spark-an-arm-soc-for-windows-pcs.md
source_anchor: ""
source_lines: [51, 69]
sha256: 961b297f49a195b0b573f5dd42438c1f2c72221ffd3d958ae1c72586694f0929
---

# nvida-introduces-rtx-spark-an-arm-soc-for-windows-pcs

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
