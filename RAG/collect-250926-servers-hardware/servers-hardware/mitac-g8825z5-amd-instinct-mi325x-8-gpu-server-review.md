---
id: collect-250926-servers-hardware/servers-hardware/mitac-g8825z5-amd-instinct-mi325x-8-gpu-server-review
title: "mitac-g8825z5-amd-instinct-mi325x-8-gpu-server-review"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: ["2025-09-01"]
keywords: ["amd", "gpu", "cost", "gpus", "memory"]
source: docs/RAG/clean4/mitac-g8825z5-amd-instinct-mi325x-8-gpu-server-review.md
source_anchor: ""
source_lines: [1, 77]
sha256: 7c7d678bde19e95b50c22975c6c1926942a9da325e58cdac1d27dfc6bcab4eb1
---

# mitac-g8825z5-amd-instinct-mi325x-8-gpu-server-review

Today, we are taking a look at the MiTAC G8825Z5. This is the company’s AMD Instinct MI325X 8-GPU platform with 2TB of HBM3E memory onboard. What is more, it uses two high-speed AMD EPYC Turin CPUs and has up to twelve PCIe slots for a lot of expansion capability. MiTAC has a unique way of making these servers and we are going to see how they do it.

We are going to take a look at this model, the MiTAC G8825Z5U2BC-325X-575, that we are just going to call the MiTAC G8825Z5 for short. We will also examine this in comparison to other systems. First, we will start with the chassis. Then we are going to get into the GPU comptue tray and the CPU tray.

## MiTAC G8825Z5 Chassis Overview

The system itself if an 8U chassis that has the GPUs on top and everything else on the bottom.

Something nice about the server is that it is easy to service with the major sets of components on trays.

We will get to the trays in subsequent sections, but here is what the chassis looks like from the front with the trays removed.

The top mostly has room for airflow, but then there are also the high-density power and data connectors for the 8-GPU UBB. Something that many do not know, is that the chassis side of the UBB connectors need to be built correctly because the UBB connectors themselves are fairly fragile.

Here is a shot withotu me in it.

On the bottom, here is the motherboard tray area. You can also see the internal rails that the two trays slide on.

On the rear, we get an array of fan modules and then the power supplies.

There are a total of 15 fan modules.

These are hot swappable from the rear.

Here is what the fan modules look like from the other side.

There are also six 3.3kW 80Plus Titanium power supplies for a 4+2 redundant configuration. The rear of the system has no I/O which many customers really like.

Here we can see that installed cables are coming out of the front, even for the management NICs.

Here is the rear where we only need power supplies hooked up.

That is because of the way that is built, and so next, we are going to get inside the motherboard and PCIe tray.

Great deep dive into the MiTAC G8825Z5. The tray-based serviceability and redundant PSU setup are impressive, especially for dense GPU workloads. It is fascinating how designs like this balance airflow and expansion. Thanks for the detailed breakdown.

Roughly, how much would this beast cost?

(I know, I know, if you have to ask…)

@TurboFEM: I’ve been debating about buying a house in La Jolla, CA, or just keeping my apartment in Akron, OH . . . in the end, it’s really not a difficult choice, as I don’t need to move ;)

All well and good, but will it run Crysis?

That last photo is great, a Patrack.

Also great is when the manufacturer has their webpages up for this and it’s brother.

They also have a DOCs page (with PDFs) and a cool YouTube video of an exploded view fly-through:

https://www.mitaccomputing.com/en/products/8U_Rackmount_Server

https://www.mitaccomputing.com/en/documents/Barebones/G8825Z5_G8825Z5U2BC-325X-575#dc

https://www.youtube.com/watch?v=TQBd3mHhtZo

Thanks for bringing this to us, even if we can’t afford it. @TurboFEM 8*30K + Memory+ 2 processors, SSDs and network cards. So, north of 1/2M.

“Laurence ‘GreenReaper’ Parry September 1, 2025 At 10:05 pm

All well and good, but will it run Crysis?”

*How many 4K instances of Crysis can this run?

Nice write-up Patrick!

BTW, what tool do you use to display the PCIe topology like that?

Reach out to ussales@mitaccomputing.com for more information for G8825Z5

@Ash

That is the lstopo command

default output is graphical, but there is also an ASCII output.

Comments are closed.
