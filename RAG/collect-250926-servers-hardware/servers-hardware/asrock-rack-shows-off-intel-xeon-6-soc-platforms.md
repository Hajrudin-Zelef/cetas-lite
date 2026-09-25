---
id: collect-250926-servers-hardware/servers-hardware/asrock-rack-shows-off-intel-xeon-6-soc-platforms
title: "asrock-rack-shows-off-intel-xeon-6-soc-platforms"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "cost", "memory", "pricing"]
source: docs/RAG/clean4/asrock-rack-shows-off-intel-xeon-6-soc-platforms.md
source_anchor: ""
source_lines: [1, 39]
sha256: f6b990e417e2fff99872caf855856b95efebb63000328f23e37c596282740233
---

# asrock-rack-shows-off-intel-xeon-6-soc-platforms

The Intel Xeon 6 SoC is the next-generation Xeon D, but it is one that we have not seen overly often at this point. ASRock Rack showcased the Xeon 6 SoC in a 2U platform, along with the motherboard, at Computex 2025.

## ASRock Rack Shows off Intel Xeon 6 SoC Platforms

The motherboard shown was the ASRock Rack GNRDD8HM-2O. This is a custom half-width form factor meant for higher-density applications.

Here you can see the two QSFP28 ports for 100GbE, the ASPEED BMC, and a riser slot.

In the middle, you can see two E1.S slot headers and an internal M.2 slot.

The heatsink is a passively cooled version with four DDR5 DIMM slots placed and four are not populated.

This all makes more sense when you see the ASRock Rack 2U4N-EGB which was also on display. As the name suggests, this is a 2U 4-node chassis, but there is something else that is neat in the system.

That riser slot allows for a 2U height node that adds extra PCIe slots to the system. There are also the two E1.S slots that have the drive trays here.

On the other side we have the double stacked nodes. It seems like there are two options.

On the right side, there are power supplies.

Here is a shot of the bare Xeon 6 SoC package in the system.

It seems like in late Q3 or in Q4 2025 when the higher-end Xeon 6 SoC parts become available that this also has the option to add those.

## Final Words

This looks like a project that was done for a customer or two. It also looks pretty neat since it is a unique design. Hopefully, it gets greenlighted. One challenge is that the Intel Xeon 6 SoC is a significantly higher cost option than many of the non-SoC versions. For those who are using the lower power, integrated NIC, and offloads, it is a valuable platform, but Intel has moved the platform out of the original range of the Xeon D-1500 pricing ranges. Still, it is a neat design we thought folks might like seeing.

How is this an SoC? Doesn’t an SoC have to include graphics capabilities embedded on package at least enough to drive a display?

@Mark why? The BMC already provides that.

Then that means this is inherently NOT an SoC as one of the ‘system’ components (graphics) is NOT ‘on the chip’.

Mark, this is definitely a “system” on a chip, but maybe not a desktop or laptop system. It has the required I/O and processing for network systems. SOC does not necessarily require or imply graphics as part of the system. I’d be more likely to argue it isn’t a SOC because the memory is off chip.

All CPUs have I/O in them nowadays. Every single one. Are all CPUs SoCs?

Comments are closed.
