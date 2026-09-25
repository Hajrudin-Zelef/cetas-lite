---
id: collect-250926-servers-hardware/servers-hardware/gigabyte-ai-top-atom-nvidia-gb10-review-a-neat-little-box
title: "gigabyte-ai-top-atom-nvidia-gb10-review-a-neat-little-box"
domain: servers-hardware
role: reference
task: reference
actors: ["AWS", "Nvidia", "United States"]
dates: ["2025-11"]
keywords: ["nvidia", "blackwell", "dram", "dram price", "ethernet", "fp4", "gpu", "lpddr5x", "memory", "pricing"]
source: docs/RAG/clean4/gigabyte-ai-top-atom-nvidia-gb10-review-a-neat-little-box.md
source_anchor: ""
source_lines: [1, 67]
sha256: 4d1a9474f346469e174b780b7c4e390bcdf5309fa8d53d2780b279bf619245fd
---

# gigabyte-ai-top-atom-nvidia-gb10-review-a-neat-little-box

One of the more remarkable aspects of NVIDIA’s GB10 SFF workstation program, besides the hardware itself, of course, has been just how many of its traditional partners NVIDIA has lined up to develop and sell semi-custom machines based on NVIDIA’s original design. As a first-generation product with no true parallel within the NVIDIA ecosystem, getting partners to invest in the production of an unproven product category is not an easy feat. But invested they have, and as a result, NVIDIA’s entrance into the miniature workstation/dev box market has landed with a lot more force than it otherwise would have, with the buy-in from its partners sending a strong signal to potential customers that the hardware and the product line are not going to be a flash-in-the-pan type of product.

That brings us to the subject of today’s review. We have looked at a few different GB10 boxes over the past several months, and today we are taking a look at a neat little box that is Gigabyte’s entry into the space: the Gigabyte AI TOP ATOM.

| **GIGABYTE AI TOP ATOM Key Specs (As Reviewed)** |  | 
| **Processors** | NVIDIA GB10 Superchip 10x Arm Cortex-X925 10x Arm Cortex-A725 Blackwell GB20B GPU – 48 SMs | 
| **Operating System** | NVIDIA DGX OS | 
| **Memory** | 128GB LPDDR5X-8533, Soldered | 
| **Storage** | 1TB SSD (PCIe Gen4 x4 M.2-2242, TLC) | 
| **GPU** | NVIDIA Blackwell, 1 PFLOP FP4 AI | 
| **PSU** | 240W External Adapter, USB-C | 
| **Form Factor** | SFF PC | 
| **Dimensions** | 150mm x 150mm x 51mm (5.91 x 5.91 x 2.01 in) | 
| **Weight** | 1.2kg (2.64 lbs) | 
| **Networking** | ConnectX-7 200Gbps Ethernet 10Gb Ethernet | 
| **Wireless** | Wi-Fi 7 (2×2) + Bluetooth 5.4 | 
| **Color** | Black | 
| **Ports** | **Rear:**  4x USB-C 20Gbps (Gen 2×2) w/DP Alt mode, 1x HDMI 2.1a, 1x 10GbE (RJ45), 2x 200GbE (OSFP112) | 

As with the other GB10 systems we have reviewed thus far, this is a small form factor system aimed at the developer market. And for that development market, the theme coming from NVIDIA is standardization: with NVIDIA exercising tight control over all of these DGX Spark-alikes, the core hardware in each and every one thus far has been the same. That means a GB10 SoC paired with 128GB of LPDDR5X memory, an M.2 SSD for storage, a ConnectX-7 NIC for high-bandwidth networking, and all placed in a chassis just a bit over 1 liter in volume.

This makes all of the Spark-alike systems remarkably consistent from one vendor to another, though the trade-off is that it leaves limited room for hardware innovation from the partners outside of cooling and selecting which SSD to install. Instead, innovation has come from the other areas, such as the support and services side of the equation. In Gigabyte’s case, this means partnering with AVADirect in the US to offer next-day on-site repair services.

If you wanted to find the Gigabyte AI TOP ATOM online, here are Amazon Affiliate links to the 4TB model and the 1TB model. Pricing is really challenging right now, as GB10 pricing is being adjusted to reflect DRAM price increases. Case in point: in the week before this article went live, the NVIDIA DGX Spark increased in price from $3999 to $4699. Some of the partner GB10s are now a much better deal with the gap of $700 or more for the 4TB models, at least for now. Since DRAM pricing increases will eventually hit everyone, we are likely in a period where there is a larger gap now than we expect in a month or two.


## Gigabyte AI TOP ATOM External Hardware Overview

Unlike the rather flashy golden DGX Spark, or even ASUS’s own silvered Spark-alike, the AI TOP ATOM starts things off with a far more muted presentation, perhaps more similar to the Dell Pro Max with GB10. Black goes well with anything, and it does not hurt that it helps to hide the significant ventilation these compact GB10 systems require.

In fact, the AI TOP ATOM pretty much goes all-in on the concept. You will not find a single port or button on the front of the system, and even the power button is in the rear. So other than Gigabyte’s logo, it is all business in the front.

Otherwise, with dimensions of 150 x 150 x 51mm (5.91 x 5.91 x 2.01 inch), Gigabyte’s take on the Spark is every bit as tiny as NVIDIA’s original design, and uses a similar front-to-back cooling design.

With nothing of substance on the sides, we will rotate around to the back of the AI TOP ATOM to take a look at what it offers.

If you have seen any of our other GB10 reviews thus far, then you know what you will find here. Starting from the left is the system’s power button, followed by four USB Type-C ports. All of these are 20Gbps ports, running in USB 3.2 Gen 2×2 under the hood. The left-most port is also the power-input port for the system, taking full advantage of the 240 Watt EPR limit for USB-C. The remaining three USB ports are then for any peripherals that need to be hooked up to the system, including external displays, thanks to support for DisplayPort 2.1 via DP alt-mode for USB-C.

Compared to some of the other GB10 system designs, Gigabyte’s labeling here is pretty barebones. The USB-C ports are labeled as such, but do not indicate their speed or use any official USB iconography.

Next to the USB-C ports is the system’s sole dedicated display output: an HDMI 2.1 port. NVIDIA has done quite a bit since the initial introduction a few months ago to support more resolutions, which makes this much more of a useful desktop PC than it was in November 2025.

Moving to the right again, we reach the part of the GB10 design that is arguably more interesting, with its trio of network ports.

First off is a Realtek RTL8127 10GbE port for local networking. In a typical mini-PC, this would be the fastest port you would find on a system, but of course, the Spark aims much higher.

Which is why it is rounded out by a pair of 200Gbps QSFP112 ports. Connected in turn to the mini-PC’s integrated NVIDIA ConnectX-7 NIC, the high-performance NIC is one of the marquee features of the GB10 boxes, with our own Patrick Kennedy often contextualizing GB10 boxes as a ConnectX-7 NIC with a really good SoC and RAM bolted on.

The internal architecture of the Spark design does make getting the most out of these QSFP ports a bit tricky, however. As we have outlined in previous articles, 200Gbps of networking needs a PCIe 5.0 x8 connection to be fully fed, but GB10 can only provide PCIe 5.0 root ports with at most a x4 connection. As a result, the NIC is connected to a pair of x4 ports. This provides the same aggregate bandwidth as an x8 connection, but it imposes network topology restrictions due to the NIC appearing to the rest of the system as four network interfaces. In many ways, this is similar to what we saw with our NVIDIA ConnectX-8 C8240 800G Dual 400G NIC Review, having two PCIe Gen5 x16 connections and two 400Gbps ports. The GB10 is like a small version of that.

Ultimately, the purpose of including a high-end ConnectX-7 NIC is to allow NVIDIA GB10 boxes to scale out, similar to their big iron brethren. 200Gbps of networking bandwidth is not nearly as much as a full-fledged GB200/GB300-based server, but it gives developers access to more processing power and a way to see how their software and models will perform on a scale-out setup. The most common setup we expect to see is two-way systems, using a single cable offering 200Gbps of bandwidth between GB10 systems. At the same time, with a network switch, it is possible to scale the whole cluster out to several machines. We already reviewed the MikroTik CRS812 DDQ, which is great for 2-4 NVIDIA GB10 machines as well as storage. We also have two MikroTik CRS804 DDQ switches that we are using with this machine in an 8x GB10 cluster. We put a note in the forums on that one.

Not pictured here, the AI TOP ATOM also offers one final networking option with an integrated Wi-Fi 7 (2×2) + Bluetooth 5.4 radio.

The rest of the back of the AI TOP ATOM, in turn, is dominated by the exhaust vents towards the top of the chassis.

Flipping the ATOM on its back, we also get a quick look at the bottom of the machine. This is for the most part a sealed machine from the bottom; Gigabyte has placed a single intake vent near the front of the system. The bottom itself serves as a stand to keep the rest of the system elevated, keeping said vent free from obstruction.

Finally, here is a quick look at the external power supply included with ATOM. This is a powerful 240 Watt USB-C adapter from Delta – and a common fixture among Spark-alike systems. According to NVIDIA documentation, the GB10 chip itself has a TDP of 140 Watts,

Now let us get inside the system, or at least, as much as we can.

I would like to see a comparision of 19″ rack-mount kits for these and other GB10 systems.

Get the MyElectronics rack for the GB10s. It is *fantastic*. Nothing else even comes close.

What’s the point of reviewing all of these Nvidia boxes? Are they not all the same board from Nvidia?
