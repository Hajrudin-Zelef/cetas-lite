---
id: collect-250926-servers-hardware/servers-hardware/amds-epyc-venice-instinct-mi455x-helios-hardware-on-display-for-first-time-at-ces-2026-1
title: "amds-epyc-venice-instinct-mi455x-helios-hardware-on-display-for-first-time-at-ces-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia", "TSMC"]
dates: []
keywords: ["amd", "helios", "mi455x", "2nm", "3nm", "attention", "compute", "fp4", "gpu", "gpus", "hbm", "hbm4"]
source: docs/RAG/clean4/amds-epyc-venice-instinct-mi455x-helios-hardware-on-display-for-first-time-at-ces-2026.md
source_anchor: ""
source_lines: [1, 44]
sha256: 8a1c93bb04a9f5794c95d91ba493fcbb3e74bb15395635b78bbefed7fc63e55a
---

# amds-epyc-venice-instinct-mi455x-helios-hardware-on-display-for-first-time-at-ces-2026

Alongside AMD’s numerous client-focused hardware announcements during their CES 2026 keynote, the company also devoted a bit of attention to their data center/server products. Though the company is essentially mid-cycle on its major data center products right now – and thus is not launching anything in the immediate timeframe – AMD still opted to use the attention garnered by hosting the official opening keynote for CES to outline the status of its 2026 server products.

### EPYC “Venice”

Notably, both AMD’s next-generation server CPU and next-generation server GPU were on display in some fashion. First off is Venice, the codename for AMD’s next-generation EPYC processor based on the Zen 6 CPU architecture. While AMD Venice has been a staple of AMD’s data center CPU roadmap for some time now, and the company has previously gone into high-level details about the chip, CES 2026 marked the first occasion where they’ve shown off the actual silicon.

This week’s show and tell did not come with any new information on the specifications for the chip itself. But for the first time we can see how the chip is laid out, something that garnered increased interest after AMD revealed last fall that Venice would rely on 2.5D packaging.

When it launches later this year, Venice will bring several firsts for AMD’s EPYC lineup. Enabled by the aforementioned 2.5D packaging, the heart of the chip is now comprised of what appears to be two I/O Dies (IODs) at the center of the chip. Then directly flanking those IODs are 8 CCDs holding the Zen 6 CPU cores. This is fewer CCDs than found on the current Zen 5 Turin processor (which houses 16 of the things), but it is more CPU cores altogether, as AMD has previously indicated that Venice will have 256 CPU cores. AMD has previously disclosed that the Zen 6 CCDs used in Venice will be built on TSMC’s forthcoming 2nm process.

Otherwise, the biggest question mark with regards to physical hardware is the 2.5D packaging used – if AMD is throwing everything on a large silicon interposer here, or using a finer fan-out solution (ala Navi 31). All of which we expect will be disclosed a bit later in the year as AMD gets closer to shipping the silicon.

### Instinct MI455X (& MI500, too)

Alongside our first look at Venice, AMD CEO Dr. Lisa Su also had its GPU counterpart, MI455X, on-hand and in-hand.

As with Venice, this marks the first time that AMD has shown off the next-generation server GPU. And, fittingly, it is an even bigger deal for the company, both in terms of its sheer size – between its chiplets it is comprised of 320 billion transistors – and in terms of AMD’s plans for expanding their share of the server market.

Unlike Venice, AMD has not sent along any full chip renders. But the chonky chip is big enough to easily see even from the broadcast cameras focused on stage. Multiple chiplets are surrounded by 12 36GB HBM4 memory stacks. AMD tells us that there are 12 chiplets overall – with some covered by the 3D die stacking – which are being fabbed on a mix of TSMC 2nm and 3nm process nodes.

Previously, AMD had disclosed two MI400 series SKUs: the flagship MI455X for AI, as well as the more nebulous MI430X, a chip with faster FP64 for HPC use cases. These two chips, as it turns out, will be joined by the MI440X, which appears to be a scaled-back version of the MI455X. Specifically, AMD is pitching it at 8-way GPU servers, which is in contrast to MI455X’s rack-scale ambitions. There are no further technical details on the chip, but presumably AMD is pruning its power consumption – and possible some I/O – in order to get something that works better in a traditional server node.

Su also used the presentation to reiterate that AMD’s current public roadmap goes out to the MI400 series’ successor, the MI500 series. The company has previously disclosed that the MI500 series will be released in 2027 – in keeping with their annual data center GPU release cadence – and for CES they have delivered some high-level details for the first time. As expected, the chip will be based on a newer CNDA architecture, CDNA 6, and will be fabbed on an “advanced 2nm” process, and will be paired with HBM4e memory. As with the launch of the MI400 series, expect an on-time launch for the MI500 series to be in the latter part of 2027.

### Helios, AMD’s First Rack-Scale System

Ultimately, both Venice and MI455X are meant to be a bigger part of a whole – that whole being AMD’s Helios rack, which was the final data center reveal of the night. When it launches later this year, Helios will be AMD’s first scale-up rack scale system for data center AI processing. As a result, it is a very big deal for AMD, as it is both a vehicle to sell lots of AMD chips stuffed throughout the rack, but also to give them an in-house product to rival NVIDIA’s NVL72 racks.

AMD disclosed Helios back at their 2025 summer data center event, but like the rest of the data center hardware discussed during CES, the show marked the first time it was revealed to the public.

Helios is, in short, a beast. AMD has opted for a double-wide rack for Helios, making it twice the size of a typical server rack. And at nearly 7000 pounds it is quite heavy as well.

As AMD was showing off a full Helios rack, they also had a single compute tray propped up to show off what it is made of.

Helios will include 18 compute trays, each housing 4 MI455X accelerators and a single Venice CPU, giving the complete system a total of 72 GPUs and 18 CPUs – or breaking that down into core counts, about 18,000 compute units and 4,608 CPU cores. All told, the monster rack is slated to offer 2.9 EFLOPS of FP4 compute performance for AI workloads. As well, all of the racks are equipped with networking gear from AMD’s Pensando division, providing the critical scale-out capabilities that make Helios such an important product for AMD.

As with the rest of AMD’s data center product stack, Helios is slated to be released later this year.

AMD’s packaging for Venice is a huge leap forward. Removing the SerDes links between chiplets will increase bandwidth and reduce latencies which should improve multithreaded performance across the chip greatly. I do wonder if the two IO dies used are identical and just rotated 180 degrees with respect to each other. That would keep the design simple and promote economies of scale. What is interesting is that the 256 core figure is likely with the Zen 6c cores for density and not the higher clocked, V-cache capable units which should be 12 core chiplets. That puts an effective limit at 96 cores for those gunning for higher clocks and/or V-cache.

Personally I am surprised at how square the Venice packaging is. I would have thought AMD would leave some room to grow to additional chiplets for things like GPU compute, FPGA, or on-package networking without the need to remove CPU chiplets. I see this as being particularly noteworthy as networking is quickly moving to copackaged optics on the switch side and it is only a matter of time before server chips like this include 800 Gbit/1.6 Tbit networking inside of the socket. The big benefit here is that PCIe lanes off of the IO die will not need to be leveraged for network interfaces while simultaneously providing lower network latency.

Also worth calling out the 16 DIMM slots for Venice in the compute trays. As widely rumored before, it appears that Venice will support 16 channel memory. Without some major changes to memory (see CAMM), this is the effective limit for how many memory channels can hang directly off of a socket. This highlights what DIMMs are for nowadays: pure memory capacity. HBM has taken over for latency, bandwidth and density.

