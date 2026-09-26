---
id: collect-250926-servers-hardware/servers-hardware/asus-pro-ws-w890e-sage-se-motherboard-review-1
title: "asus-pro-ws-w890e-sage-se-motherboard-review"
domain: servers-hardware
role: reference
task: reference
actors: ["AWS", "Intel"]
dates: []
keywords: ["ethernet", "intel", "memory", "power delivery"]
source: docs/RAG/clean4/asus-pro-ws-w890e-sage-se-motherboard-review.md
source_anchor: ""
source_lines: [1, 61]
sha256: 585db626bd045506eb0f161bed3dd6b5de5aba6accea3106821c679f7bf3704e
---

# asus-pro-ws-w890e-sage-se-motherboard-review

Earlier this year, Intel launched its Xeon 600 workstation processors, marking the transition of its workstation platform to the newest Granite Rapids architecture. The release of the new CPUs was also accompanied by an update to the rest of Intel’s supporting hardware, including a new socket (LGA 4710) and the new W890 chipset, which in turn brought a new generation of workstation motherboards to support the new platforms.

Today we are looking at one of the new high-end motherboards for the new Xeon 600 platform, ASUS’s Pro WS W890E SAGE SE. One of two motherboards ASUS launched using the new socket, the SAGE SE is their premium offering, offering additional features and a flat-out larger motherboard for a premium price tag. The Xeon 600-generation motherboard is the direct successor to ASUS’s previous-generation Pro WS W790E-SAGE SE, which, for its time, was one of the better-received Sapphire Rapids workstation motherboards.

| **ASUS Pro WS W890E-SAGE SE Key Specs** |  | 
| **Processor Family** | Xeon 600 (Granite Rapids) | 
| **Socket** | LGA4710-2 | 
| **Chipset** | W890 | 
| **Memory** | 8x DDR5 RDIMM Slots (1DPC) | 
| **PCIe Slots** | 6x PCIe x16 (Gen5 x16 Electrical) 1x PCIe x16 (Gen5 x8 Electrical) | 
| **Storage** | 2x M.2 22110 (Gen5 x4) 2x M.2 2280 (Gen5 x4) 1x MCIO x8 (Gen5 x4/x4) 2x SlimSAS (Gen4 x4) 4x SATA 6Gbps | 
| **Ethernet** | 2x 10GbE (Intel E610-XAT2) | 
| **Rear USB** | 2x 40Gbps USB-C 6x 10Gbps USB-A | 
| **Front USB** | 1x 20Gbps USB-C 2x 5Gbps USB-A 2x 480Mbps USB-A | 
| **Audio** | Realtek ALC1220P | 
| **BMC** | ASPEED AST2600 | 
| **Form Factor** | EEB (E-ATX) | 

Meanwhile, for DIY workstation builders in particular, the W890E-SAGE lineup is effectively the default option for equipping a new system. ASUS’s boards are the only Xeon 600 motherboards readily available in the North American market, and, as the boards the company sampled for the launch of the Xeon 600 processors, they are the de facto retail motherboard for Intel’s new workstation platform. This puts ASUS in the enviable position of controlling that market segment, but it also means the overall success of the Xeon 600 in the DIY market rests heavily on ASUS’s shoulders.

If you wanted to find the ASUS Pro WS W890E-SAGE SE online, here is an Amazon Affiliate link.


## ASUS Pro WS W890E-SAGE SE Hardware Overview

As ASUS’s premium W890 motherboard, the Pro WS is looking to scratch every itch of the retail/DIY workstation market. On the one hand, it is a high-end motherboard that exposes the full functionality of the Xeon 600 chips and the W890 chipset, offering support for a full 8-channel memory configuration, copious I/O bandwidth, and even a BMC for business use. On the other hand, ASUS is leaning into the overclockable nature of the X-series SKUs by adding a fairly robust set of features specifically for the overclocking market. The end result is a board that offers something for almost everyone and, in the process, earns its status as a high-end workstation motherboard.

The significant feature set means the W890E-SAGE SE is a sizable motherboard. ASUS has gone with the full 12-inch by 13-inch EEB (E-ATX) form factor, taking up almost every last square inch of space on the board to lay down a slot or socket of some sort.

A good chunk of that space is dominated by the CPU socket and RDIMM slots, and for obvious reasons. The LGA4710-2 socket is massive on its own thanks to the large number of I/O and power pins it contains, and then all of those pins have to run somewhere. ASUS’s high-end W890 board goes for the full 8-channel memory treatment, surrounding the CPU socket with 4 RDIMM slots above and below it to route what is effectively a 512-bit memory data bus back to the CPU.

Flanking the socket and memory in turn is a rather significant VRM/MOSFET arrangement, with ASUS running 16 phases for the CPU alone. While the highest Xeon 600 SKU is rated for a toasty maximum turbo power of 420 Watts, the overclocker-friendly design of the board means that ASUS has gone above what a stock chip would require for a power delivery system. Which means the SAGE SE can deliver plenty of power, but it occupies a decent amount of board real estate to do so.

Between those VRMs and the RDIMM slots, ASUS is packing a lot of high-power, high-heat components very close together. As a result, in addition to placing a rather sizable heatsink on top of the VRMs, ASUS also includes multiple fans to help keep the VRMs and RDIMM slots cool. Baked into the left side of the board (nearest the rear I/O ports) is a fan to keep the bulk of the system’s VRMs cool.

And on the other side of the CPU socket, ASUS offers an optional cooling fan assembly, which the company calls their AngleBoost. Sitting atop the heatsink that covers the rest of the VRMs, the AngleBoost fan kit is a pair of small fans designed to be mounted at an angle, thus living up to its name. This provides extra airflow for the heatsink itself and directs airflow to all 8 RDIMM slots.

The other big I/O pin expenditure from a board design perspective is the SAGE SE’s 7 PCIe x16 slots.

When paired with a high-end Xeon 600 processor (more than 16 CPU cores), all of these slots are available, and all but one are full PCIe Gen5 x16 slots, both physically and electrically. The one outlier is slot #6 (second from left), which is an x16 slot that electrically runs at Gen5 x8 speeds. All of these are armored slots for additional rigidity, or as ASUS refers to them, Safeslots.

The board also offers limited CXL 2.0 support. Specifically, PCIe slots 1, 3, 5, and 7 all support the expansion standard, reflecting the I/O quirks of the Granite Rapids CPU itself.

From this view, we can also catch a glimpse of a black heatsink that is the board’s network controller.

Meanwhile, the SAGE SE also offers an extensive collection of storage options. Not counting the PCIe slots or any kind of external storage, there are four different types of storage interfaces on ASUS’s motherboard, three of which are PCIe-based.

A good deal of the board’s remaining floor space is allocated to M.2 SSD sockets. ASUS has equipped the SAGE SE with four such sockets, each offering PCIe Gen5 x4 host connections. Two of the sockets support 2242/2260/2280-size SSDs, while the other two sockets go one size larger and support extra-long 22110 M.2 SSDs.

ASUS provides heatsinks for all of these M.2 sockets. A single L-shaped bracket covers three of the sockets on the lower portion of the board, and there are thermal pads both on the motherboard and on the heatsink to ensure that both single- and double-sided SSDs can be cooled.

The final M.2 slot near the top of its board features its own heatsink and mounting mechanism, which ASUS calls its Q-Release Slim heatsink. This mechanism is entirely toolless, allowing it to be removed from the board entirely to more easily install an M.2 SSD inside it.

Meanwhile, for larger storage devices such as U.2 or E3-form-factor SSDs, ASUS has included a single MCIO connector. This connector is attached directly to the CPU and provides a pair of PCIe Gen5 x4 connections, allowing two additional SSDs to be installed in this fashion.

For slightly older SSDs, there are also a pair of SlimSAS ports. These hang off of the W890 chipset, and each provides a PCIe Gen4 x4 connection.

Finally, a quartet of SATA 6Gbps ports is also found on the SAGE SE, allowing for SATA SSDs, HDDs, and optical drives to be used with the motherboard as well. Altogether, this means that the SAGE SE can connect to 12 internal storage devices before relying on using multiplexors or the PCIe slots.

The power requirements of all of this hardware, in turn, are quite significant. And to that end, there are several power sockets of note on the SAGE SE.

