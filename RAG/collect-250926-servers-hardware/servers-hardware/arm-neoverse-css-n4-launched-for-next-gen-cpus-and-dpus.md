---
id: collect-250926-servers-hardware/servers-hardware/arm-neoverse-css-n4-launched-for-next-gen-cpus-and-dpus
title: "arm-neoverse-css-n4-launched-for-next-gen-cpus-and-dpus"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["3nm", "agi", "chiplet", "compute", "fp8", "memory"]
source: docs/RAG/clean4/arm-neoverse-css-n4-launched-for-next-gen-cpus-and-dpus.md
source_anchor: ""
source_lines: [1, 15]
sha256: a8a66e3347fa641034bf82b75fc63b54552d010d5be1e7baf4665edb5803d433
---

# arm-neoverse-css-n4-launched-for-next-gen-cpus-and-dpus

Arm announced a new compute subsystem. The Arm Neoverse CSS N4 brings up to 128 cores, PCIe Gen7, and more to future servers. We have seen Neoverse CSS in more products over the past few years, so this is an important launch for the industry.

## Arm Neoverse CSS N4 Launched for Next-Gen CPUs and DPUs

The new Arm Neoverse N4 CSS is built for 3nm process and supports 8 to 128 cores per die. With Neoverse N4 cores, Armv9.3 architecture becomes the baseline. These cores are designed to operate at up to 3.8GHz given their focus on power-efficient compute, but also bring support for FP8 and MMLA. For the core interconnect, there is an Arm Neoverse CMN S4 Interconnect as the fabric for the CSS N4 subsystem. CMN S4 supports up to a 16×16 mesh interconnect and a fully coherent multi-chiplet connectivity using CHI C2C. This fabric is what ties the CPU cores, memory, I/O, die-to-die interfaces, and so forth together.

On the memory side, this new generation also brings support for both LPDDR6 memory as well as DDR5 and MRDIMMs at 8000-12000MT/s speeds. For I/O, there are up to 128 lanes of PCIe Gen7 which also tells us that this is designed for CPUs coming out in at least one generation from now.

One of the neatest N-series CSS implementations we have seen recently is the XSight Labs E1 64-Core Arm 800G DPU. This is an Arm Neoverse CSS N2 design with 64 Arm Neoverse N2 cores.

For a startup like XSight Labs, using a Neoverse CSS design makes a lot of sense, since Arm does much of the base design work, helping a smaller team get to market faster.

## Final Words

Hopefully we get more details on the Neoverse CSS N4 as we get closer to seeing products with the new IP. While it is often the Neoverse V-series line making headlines, the N-series of cores has become very popular in product lines like DPUs and network appliances. Arm is also making its own CPUs now, with the Arm AGI CPU that is using Neoverse V3 cores. We are going to have more on that soon on STH.
