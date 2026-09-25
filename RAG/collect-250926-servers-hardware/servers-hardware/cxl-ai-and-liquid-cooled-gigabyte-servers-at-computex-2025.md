---
id: collect-250926-servers-hardware/servers-hardware/cxl-ai-and-liquid-cooled-gigabyte-servers-at-computex-2025
title: "cxl-ai-and-liquid-cooled-gigabyte-servers-at-computex-2025"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["compute", "accelerator", "ethernet", "gpu", "gpus", "intel", "liquid cooling", "memory", "nvidia"]
source: docs/RAG/clean4/cxl-ai-and-liquid-cooled-gigabyte-servers-at-computex-2025.md
source_anchor: ""
source_lines: [1, 54]
sha256: 998bc36f4275e1350941af44dee0df4184347823e8896e7eb907db6f4f50159a
---

# cxl-ai-and-liquid-cooled-gigabyte-servers-at-computex-2025

At Computex 2025, we saw a lot of servers. Gigabyte’s booth had a good mix of them, including a few unique designs. We saw a neat CXL and NVMe server, along with liquid-cooled and air-cooled AI servers. There were even some neat multi-node servers, like a HPC-focused liquid-cooled 1U 2-node server and a 10-node 3U server.

Here is a quick 1-minute short on a few of the highlights:


## Gigabyte R284-A91-AAL3 CXL and NVMe Server

Gigabyte showed off the R284-A91-AAL3 with two Intel Xeon 6900 series processors with a twist in this 2U E3.S platform.

Using E3.S form factors, you can have both NVMe SSDs as well as up to sixteen Micron CXL memory expansion modules.

In our photos, we saw the neat little CXL tray design. This is what the Micron CZ122 128GB CXL module goes into.

Each Intel Xeon 6900 series CPU gets 12 memory channels plus eight additional CXL memory modules worth of capacity and bandwidth.

There is more storage including dual M.2 slots on the motherboard of this 2U server.

At the rear we get power supplies, OCP NIC slots, and riser slots, but there are two more 2.5″ bays at the rear for even more storage.

There is certainly a lot going on in this server.

## Gigabyte 4U AI Liquid-Cooled Servers

Gigabyte showed off a number of AI servers. Those included 4U 8x GPU compute servers that are liquid-cooled.

Gigabyte has a number of AI accelerator options but by using liquid cooling, these can only take 2U of rack space instead of 6-8U. As a result, the 4U servers can be stacked into dense AI racks. It may seem like a small feature, but the shorter profile of an AI accelerator coldplate versus a heatsink dramatically reduces the height of each system, allowing for copper DACs to be used instead of optical connections across more nodes. Those copper DACs decrease costs while increasing reliability.

This appears to be an Intel Xeon-based AI server with CoolIT liquid cooling that gives high-density compute, memory, and networking capabilities.

This is not a completely liquid cooled design. Instead, the main heat sources are liquid cooled. The DIMMs and SSDs for example, are still air cooled, but fewer fans can be used.

This is a design we have seen GIGABYTE use with many AI servers recently, but the a difference here is that foam blocks are inserted instead of some of the fans because cooling requirements are so much lower.

4U is a big deal since companies can put 8-10 in a modern AI rack and maximize density with 80 GPUs in a rack. It was neat to see the GIGABYTE GIGAPOD nodes like this at Computex.

## Gigabyte G893-SG1-AAX1 Intel Gaudi 3 Server

In addition to the liquid cooled AI systems, we saw an air cooled Gigabyte Intel Gaudi 3 AI server.

A huge feature in this server is the eight OSFP ports. The Intel Gaudi 3 utilizes Ethernet for its scale-out. As a result, instead of NICs, the Ethernet comes from the OAM accelerator. We showed how NVIDIA is starting to migrate in the direction of Intel Gaudi.

This server has eight Intel Gaudi 3 AI accelerators in an OAM form factor that connect to those OSFP ports

There was the CPU host compute sled here as well with its 10GbE, management interfaces, and eight 2.5″ NVMe bays.

Here is a quick look at GIGABYTE’s AI accelerator host CPU platform.

You can see that, like many other acceleration platforms, this is using the 8-channel memory Xeons instead of the 12-channel because it allows for 16 DIMMs per CPU or 32 DIMMs total.

Apropos of the B343-X40-AAS1: how are blade servers doing these days? In principle it seems like they could be a lot more capable than they used to be(or a lot less dependent on proprietary black-box glue logic to be as capable; you can’t necessarily just plug a multi-host NIC or one of the more complicated PCIe topologies into a random motherboard and expect it to work; but those are now standards that you can make work); but I can’t remember the last time I heard one mentioned as an object of any interest.

Are they still out there, plugging away reliably in some niche; or did the economics of being wholly at the vendor’s mercy for network modules and theoretically more elegant chassis-level management just not really survive the squeeze between ‘just use big VM host’ if you want a bunch of hosts that look like they have fast networking and flexible resource allocation between them and ‘commodity 2U4n or other high density’ that is slightly less elegant but cheaper and faster moving if you just want a bunch of nodes?

Comments are closed.
