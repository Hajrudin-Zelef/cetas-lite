---
id: collect-250926-servers-hardware/servers-hardware/nvidia-announces-dsx-ready-qualification-program-for-data-center-power-and-cooling-hardwar
title: "nvidia-announces-dsx-ready-qualification-program-for-data-center-power-and-cooling-hardware"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "distribution", "energy", "nvlink"]
source: docs/RAG/clean4/nvidia-announces-dsx-ready-qualification-program-for-data-center-power-and-cooling-hardware.md
source_anchor: ""
source_lines: [1, 15]
sha256: d53300d2bf86802111d0f013a091d3949e05669524a35f098a6c1d9e45dc7a98
---

# nvidia-announces-dsx-ready-qualification-program-for-data-center-power-and-cooling-hardware

NVIDIA this week has announced that they are launching a new qualification program for DSX AI factory power and cooling parts. The DSX Ready program will see NVIDIA offer qualification guidelines for both battery energy storage systems (BESS) and cooling distribution units (CDUs) – both classes of data center parts that NVIDIA itself does not provide – with the goal of giving customers clear guidance on what parts are confirmed to work with NVIDIA’s DSX racks in the process of bringing up their own systems. NVIDIA is initially starting with batteries and cooling, but the company has indicated that they expect to expand the DSX Ready program in the future.

With the DSX Ready program, NVIDIA is looking to flesh out their current server hardware certification program with a wide array of components and a stronger degree of branding behind it. The company already offers the “NVIDIA product qualified” label for CDUs, and this expands that in terms of requirements and the types of components offered. The ultimate goal being to move the installation of DGX server racks into an even more turn-key solution than it already is today, by not only working with their OEM partners to provide racks of the critical hardware as part of their DGX AI factory blueprint, but also giving them clear guidance on whose ancillary parts have been confirmed to work with those rack scale systems.

The DSX Ready program marks the latest in a long list of ecosystem plays from NVIDIA in and around the data center business. At this juncture the company has several major ecosystems across software and hardware, including of course the CUDA software stacks, but also their MGX platform for modular server designs, their NVLink ecosystem for cache coherent network fabrics, and the broader DSX AI ecosystem for their rack scale systems.

Notably, the DSX Ready qualification process varies depending on the type of gear – so it is not a one-size-fits-all process. CDU vendors are allowed to self-qualify based on meeting NVIDIA’s function and performance specifications (though they still need NVIDIA’s final approval). BESS vendors, on the other hand, face a more involved process that requires running qualification tests and then submitting data to NVIDIA for them to review and sign off on – presumably owing to the more complex nature of electrical work and the tighter requirements thereof. Meanwhile, the company is not disclosing what (if anything) it is charging for participating in the DSX Ready ecosystem and using the branding.

For the launch of the program, LG Electronics, LiquidStack and Vertiv are the first qualified CDU vendors. Meanwhile Hitachi Energy, LG Energy Solution and Tesla are the first qualified BESS vendors.

## Final Words

The DSX Ready program is live as of today. NVIDIA has also strongly hinted that CDUs and BESSes will not be the only data center components eventually covered by the program, taking care to note that this week’s announcement was for the “initial” categories of parts. So we should expect to see further categories of data center parts added to the DSX Ready program in due time.

I’m thinking that some of these qualifications need to be tied to particular generations of hardware due to the rapid increase of individual rack power consumption. Obviously the closer the equipment is to the individual rack the more it makes sense to add this context. Items like PDUs and CDUs for a row of racks make sense for generational tagging.
