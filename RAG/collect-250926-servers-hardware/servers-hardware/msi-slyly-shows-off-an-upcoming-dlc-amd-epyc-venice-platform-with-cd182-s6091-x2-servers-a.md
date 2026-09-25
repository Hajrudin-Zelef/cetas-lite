---
id: collect-250926-servers-hardware/servers-hardware/msi-slyly-shows-off-an-upcoming-dlc-amd-epyc-venice-platform-with-cd182-s6091-x2-servers-a
title: "msi-slyly-shows-off-an-upcoming-dlc-amd-epyc-venice-platform-with-cd182-s6091-x2-servers-and-racks"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "compute", "intel", "liquid cooling", "memory", "nvidia"]
source: docs/RAG/clean4/msi-slyly-shows-off-an-upcoming-dlc-amd-epyc-venice-platform-with-cd182-s6091-x2-servers-and-racks.md
source_anchor: ""
source_lines: [1, 57]
sha256: 673f2f1f1f28a054db1b5990bf03ad7a18343e118c48458161c2889e9f538245
---

# msi-slyly-shows-off-an-upcoming-dlc-amd-epyc-venice-platform-with-cd182-s6091-x2-servers-and-racks

One of the all-time classic activities at Computex is playing “spot the unreleased hardware.” While processor vendors such as Intel, AMD, and NVIDIA largely try to keep everything under wraps until the formal launch of their new hardware, the significant hardware ecosystem required to bootstrap the launch of these new platforms means that their OEM partners need to get an early start on board and system development. Which, in turn, means that as the launch of a new platform approaches, they get increasingly eager to show off their upcoming hardware, or at least as much of it as they can get away with.

This year’s Computex was no exception. With AMD set to launch their 6<sup>th</sup> generation “Venice” EPYC platform later this year, some of AMD’s partners have already completed their product designs ahead of that launch. Placed out in public with a wink and a nod, one such system we found this year was at MSI’s booth, where the company had an upcoming 1OU2N form factor server node on display: the CD182-S6091-X2 (DLC).

## MSI CD182-S6091-X2 (DLC)

Designed as a dual- socket system for an undisclosed AMD processor, the CD182-S6091-X2 (DLC) is an ambitious next-generation server node design from MSI that is intended for high-density installations. With two nodes in each 1OU chassis, the CD182-S6091-X2 (DLC) is designed to allow four AMD processors and their complete memory complement to fit inside a single unit of OCP’s v3 rack standard.

Key to enabling this level of density (and driving the “DLC” in the product name) is the use of direct liquid cooling. AMD has not disclosed the TDPs of Venice processors, but with the 5<sup>th</sup> generation chips already reaching as high as 500 Watts each, and more DIMM channels in Venice (16 channels shown in this system), it is easy to imagine how important liquid-cooling will become.

In fact, liquid cooling has been a common theme throughout the server booths at this year’s show, as system TDPs continue to creep up and server vendors are looking to push the envelope on both chip density and cooling performance. With air reaching its limits, liquid cooling is becoming all but necessary to achieve those higher performance densities.

For MSI, liquid cooling also makes it convenient for trade shows by hiding the CPUs in use beneath the system’s massive cold plates. Each dual-socket node features the cold plates in a serial setup, with the CPUs on their own loop. Meanwhile, a second loop is run for the similarly liquid-cooled memory, which, at 32 RDIMMs per node, generates a not-insignificant amount of heat on its own. There also appears to be some leak-detection circuitry for the CPU cold plates, though MSI has not disclosed the full extent of its capabilities.

All of this liquid cooling, in turn, is to enable the high density of the node. Despite only being half the volume of a single OU rack unit, MSI has been able to not only place two high-end EPYC processors in the server nodes, but they have also been able to fully populate all of the DIMM channels as well. Which means 16 channels of DDR5 RAM per processor in a 1DPC configuration, with half going on each side of the CPU.

The end result is that the CD182-S6091-X2 comes close to being all CPU and compute, as those two items alone take up most of the board space even on the expanded OCPv3 design.

Still, this design leaves enough room for a few expansion cards. On the networking front, each node contains an OCP 3.0 slot for installing the customer’s NIC of choice, which in turn is wired up via the next-generation PCIe Gen6 lanes that Venice will be introducing. On top of this, each node sports both a FHHL and a HHHL PCIe slot, with both of those running at PCIe Gen6 x16 thanks to the large number of PCIe lanes a dual socket setup brings with it.

Also towards the front of the system, MSI has installed a quartet of E1.S drive slots, allowing the already-dense server to include a respectable amount of local storage. Like the PCIe slots, this is PCIe Gen6-based, so the system is ready for next-generation SSDs such as Micron’s 9650 SSDs.

Rounding out the server is the usual collection of server management hardware. Each CD182-S6091-X2 node includes an ASPEED AST2700 BMC and a pair of 1GbE ports. That is neat since in the next generation, we are getting a newer BMC generation.

The larger chassis holding the individual nodes is a rather plain affair, and intentionally so.

Besides providing plumbing for the liquid-cooling loops and forced-air cooling for the other components, there is not much else there. In fact, you will not even find a power supply; the CD182-S6091-X2 (DLC) is powered via a 48VDC busbar, allowing MSI to save precious space by avoiding a power supply and further simplifying the process of installing (and removing) not only the individual nodes from the chassis, but the entire chassis itself.

## MSI’s ORv3 Liquid Cooled Rack

So where are these CD182-S6091-X2 (DLC) servers meant to be installed, anyhow? For Computex 2026, MSI had that covered as well with their ORv3 Liquid Cooled Rack.

Capable of housing 28 of the high-density servers, MSI’s ORv3 liquid-cooled rack is a 44OU rack that is designed to accommodate up to 100kW worth of high-density computing hardware. Which in this case means 112 Venice processors and their 1792 RDIMMs.

Besides the compute nodes, MSI’s demo rack was equipped with a pair of 55kW Chicony power shelves to energize the busbar running through the rack, as well as a 100kW Auras CDU, which is designed to hook into a larger cooling loop for liquid-to-liquid cooling.

MSI also had a basic networking setup for the rack installed, with a single 32-port 100GbE switch for the nodes themselves, as well as a pair of 1GbE 48-port switches for the out-of-band management network.

This is a big rack.

### Final Words

While AMD’s 6th-generation EPYC “Venice” processors are still a few months off, the excitement for them at Computex was palpable. And so was OEMs’ eagerness to show off the hardware they have developed to house them.

As evidenced by MSI’s server lineup, the OEM has significant ambitions for AMD’s upcoming hardware. With a combination of a new CPU architecture, PCIe Gen6 connectivity, and an even wider 16-channel (1024-bit memory bus, Venice is slated to bring a lot to the table in terms of performance and capacity. And MSI, in turn, is gearing up to pack Venice into even larger hardware setups thanks to its liquid-cooled 1OU2N CD182-S6091-X2 (DLC) servers and accompanying ORv3 racks. Suffice it to say, Venice brings a lot to look forward to.

“Which means 16 channels of DDR5 RAM per server in a 1DPC configuration, with half going on each side of the CPU.”

I guess this is technically correct? But that sentence feels wrong somehow. Knowing it’s 16-channels per CPU, but saying 16 channels per server. Not sure I really have a good point here. Just that the sentence somehow feels “off”.

@ James.

It’s actually 2 servers (nodes) in a 1OU chassis, so the sentence is correct. See the first image “MSI CD182 S6091 X2 (DLC) Hero” at the top of the article which shows two two nodes.

@James

That’s fair criticism. I’ve cleaned up the wording to be clearer.
