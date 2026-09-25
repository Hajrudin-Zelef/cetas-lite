---
id: collect-250926-servers-hardware/servers-hardware/supermicro-nvidia-gb200-nvl72-system-at-computex-2024
title: "supermicro-nvidia-gb200-nvl72-system-at-computex-2024"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia", "United States"]
dates: []
keywords: ["nvidia", "blackwell", "distribution", "gpu", "gpus", "liquid cooling", "lpddr", "memory", "nvlink"]
source: docs/RAG/clean4/supermicro-nvidia-gb200-nvl72-system-at-computex-2024.md
source_anchor: ""
source_lines: [1, 35]
sha256: 87a4e59eb58b5facd9fc02b0e409b16d8a3a392a38ad62cd3729a46aa0997195
---

# supermicro-nvidia-gb200-nvl72-system-at-computex-2024

Recently, the team in Tapei had the opportunity to see Supermicro’s version of the NVIDIA GB200 NVL72. You may have seen this system in a recent video when we discussed Why Servers Are Using So Much Power. We thought since we showed it in our power video as an example, we should also show it in a main site feature.

## Supermicro NVIDIA GB200 NVL72 at Computex 2024

The Supermicro NVIDIA GB200 NVL72 is Supermicro’s flavor of the NVIDIA Grace Blackwell 200 NVLink 72 GPU rack.

On top, we get a Supermicro logo, networking, and power.

We then get ten dual-node GB200 1U chassis.

The Grace Blackwell GB200 node has connections in the back for the NVLink backplane. These are half-width nodes so that two can fit side-by-side in a 1U rackmount chassis.

There are two Blackwell GPUs.

We then get the NVIDIA Grace CPU with LPDDR memory alongside it.

And the I/O connectivity at the bottom of the node.

In the middle, we get the NVLink switches that are used to connect the systems. At the bottom, we get eight more dual GB200 nodes. Eighteen 1U chassis, each with two GB200 assemblies, each of wich has two Blackwell GPUs gives us a total of 72 GPUs.

Below that we get power supplies, and a coolant distribution unit built by Supermicro.

This is a similar CDU to the design we saw in our Supermicro Custom Liquid Cooling Rack A Look at the Cooling Distribution piece.

## Final Words

This rack is a cluster that uses ~120kW, or about as much power each hour, that is stored in the 123kWh Tesla Cybertruck battery. In the video, we discussed how the GB200 NVL72 rack, in these terms, is like taking a ~7000lb truck over 300 miles. We expect higher-power data centers to utilize these GB200 NVL72 solutions in the coming months. That is one of the reasons we wanted to focus on the power side of the equation, given how much we have done on liquid cooling. This rack alone uses more power than 80 of our original STH hosting racks could deliver.

This piece also kicks-off a “big” week at STH. Stay tuned for more.

Put another way….

“This rack is a cluster that uses ~120kW, or about as much power as 100 US households.”

Comments are closed.
