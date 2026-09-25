---
id: collect-250926-servers-hardware/servers-hardware/wild-density-supermicro-2ou-8x-nvidia-b300-server-at-ocp-2025
title: "wild-density-supermicro-2ou-8x-nvidia-b300-server-at-ocp-2025"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["nvidia", "amd", "compute", "gpu", "gpus", "intel", "liquid cooling"]
source: docs/RAG/clean4/wild-density-supermicro-2ou-8x-nvidia-b300-server-at-ocp-2025.md
source_anchor: ""
source_lines: [1, 33]
sha256: f56eb14785b3932079c2d105cc2166f004670173a659f4c475060cb1dd5ba82c
---

# wild-density-supermicro-2ou-8x-nvidia-b300-server-at-ocp-2025

At OCP 2025, we saw many AI servers. This one was somewhat wild, though. Supermicro had a 2OU system with eight NVIDIA B300 GPUs. If you can power this, then it is awesome density, potentially fitting twice as many B300 GPUs in a rack versus the NVL72 racks.

## Wild Density Supermicro 2OU 8x NVIDIA B300 Server at OCP 2025

Here is the system that is notable since it is only 2OU in height. With the NVIDIA B300 generation, NVIDIA’s HGX B300 platform swaps PCIe switches and retimers for NVIDIA ConnectX-8 networking. As a result, we have eight 800Gbps ports on the front. Supermicro is taking advantage of this since it can have all front I/O and in the ORv3 rack it gets power and liquid cooling blind mating at the rear.

For storage, there are two M.2 boot drives and eight E1.S SSDs, one for each GPU.

We also get slots for NVIDIA BlueField-3 DPUs and the management I/O and another AIOM/ OCP NIC 3.0 slot on the right. Powering this are two Intel Xeon 6 CPUs (Xeon 6700/ Xeon 6500 series) and Supermicro is doing a ton of liquid cooling internally to make this all work.

Just for some sense of why this matters, here is the 8U air-cooled version where we get very similar front I/O, just on a much larger 8U 19″ rack faceplate.

We asked about how much power the liquid cooled version saves, and it is something like 10%. Another key though is the density. With 8x NVIDIA B300 GPUs in 2OU, this is a system for folks that want to put well over 100 GPUs in a rack.

## Final Words

This is just neat to see. It also shows some of the advantage of the ConnectX-8 integration on the NVIDIA HGX B300 8-GPU baseboard over the previous HGX generations. We did not get to show inside the systems because this was a unique take on the B300 generation. Inside, the other components are liquid-cooled to ensure that the system can fit in such a tight footprint.

## Bonus: Liquid-Cooled AMD Instinct MI355X

As a quick bonus, Supermicro also had its popular 4U liquid-cooled GPU platform with a twist. This version has eight AMD Instinct MI355X GPUs in the system.

20U, I think someone has made a typo

No – this is OU and not 0U. This is Open Compute rack that is 21″ and not 19″

Perhaps this should read “20 OU (Open Rack Units) aka 96mm high”

Perhaps this should read “2 OU (Open Rack Units) aka 96mm high”

No, the industry standard term is OU.

Should 19″ servers start using the term “2 U (EIA 19″ rack units) aka 87mm high” as well?
