---
id: collect-250926-servers-hardware/servers-hardware/minisforum-ms-02-ultra-review-the-new-home-lab-king-1
title: "minisforum-ms-02-ultra-review-the-new-home-lab-king"
domain: servers-hardware
role: reference
task: reference
actors: ["AWS", "Intel"]
dates: []
keywords: ["cost", "ethernet", "gpu", "intel", "memory", "power delivery"]
source: docs/RAG/clean4/minisforum-ms-02-ultra-review-the-new-home-lab-king.md
source_anchor: ""
source_lines: [1, 53]
sha256: bcd1e996576e095b08ecf61e3bd48094d8e9236e5cafdf22bc3975e54df74b65
---

# minisforum-ms-02-ultra-review-the-new-home-lab-king

Over the past 18 months or so, we have seen several mini-PC vendors put together small form factor systems designed for (or at least marketed to) the AI development crowd. With this being a quickly growing and fairly lucrative market, there continues to be a bit of a race to be in the business of selling the pickaxes in this modern gold rush.

And for specialist builders like Minisforum in particular, we are now seeing them offer multiple takes on the same concept as they tune their offerings to address different sub-sets of the AI market. The company already offers a highly integrated system with the Ryzen AI Max-based Minisforum MS-S1 Max. Now for their latest system, they are going the other direction, offering a rather modular/discrete Intel-based SFF system with their latest design, the MS-02 Ultra.

| **Minisforum MS-02 Ultra Key Specs (As Configured)** |  | 
| **Processors** | Intel Cora Ultra 9 285HX (5.5GHz) | 
| **Operating System** | Windows 11 Pro | 
| **Memory** | 64GB SO-DIMM DDR5-4800 (2x32GB) | 
| **Storage** | 1TB SSD (1TB PCIe 4.0 M.2 2280) | 
| **GPU** | Intel Graphics (Xe-LPG, 4 cores) | 
| **PSU** | 350W Internal PSU | 
| **Form Factor** | Mini-PC (4.8L) | 
| **Dimensions** | 222 x 225 x 97 mm (8.74 x 8.86 x 3.82 in) | 
| **Weight** | 3.45kg (7.6 lbs) | 
| **Wireless** | Wi-Fi 7 (2×2) + Bluetooth 5.4 | 
| **Color** | Black | 
| **Ports** | **Front:**  1x USB-A 10Gbps, 2x USB-C 80Gbps, 1x Combo Audio**Rear:**  1x HDMI 2.1 FRL, 3x USB-A 10Gbps, 1x USB-C 40Gbps, 1x 10GbE LAN (RJ45, RTL8127), 1x 2.5GBE LAN (RJ45, I226-LM), 2x 25 GbE LAND (SFP28, E810) | 

Officially, Minisforum bills this as the successor to the popular MS-01, a 2024-era system based around Intel’s Core 13<sup>th</sup> Gen (Raptor Lake) platform. And while that heritage is easy to see – especially with the use of an Intel platform – the MS-02 Ultra is not just an MS-01 with upgraded guts. It is a whole new design that, at nearly three times the volume of the MS-01, changes virtually everything about the design to allow for more expandability and performance. In short, it is a design that earns its “Ultra” moniker.

So what all are we looking at? In short, the MS-02 Ultra is a sizable mini-PC based on Intel’s Core Ultra 2 series (Arrow Lake-HX) platform. Minisforum is counting on several design aspects to set it apart from the crowd, including room for a dual-slot half-height video card, 80Gbps USB support, and in the high-end 285HX model, ECC memory support and a very rare set of dual 25GbE SFP28 ports. As a result, the MS-02 Ultra leans more towards a workstation in terms of features and component selection, with a copious amount of networking bandwidth to integrate it into larger networks.

If you wanted to find the Minisforum MS-02 Ultra online, here is an Amazon Affiliate link.


## Minisforum MS-02 Ultra External Hardware Overview

Minisforum’s MS-02 Ultra employs a new chassis design from the company that, by and large, is dictated by the system’s support for half-height PCIe cards. Whereas the MS-01 could fit a single card in a horizontal orientation, the MS-02 Ultra can fit 3 cards vertically. And that means the small(ish) form factor PC needs enough vertical clearance to accommodate those cards as well as its Intel desktop-class hardware. By SFF PC standards, this is a pretty sizable machine.

Consequently, the bulk of the front of the MS-02 Ultra chassis is reserved for airflow – and we will see why in better detail once we crack the system open. That leaves a relatively small portion along the front edge for expansion ports and buttons. Here we will find a single USB-A 10Gbps port, and a pair of deceptively powerful USB-C ports.

I say “deceptively” because, atypically for an SFF PC, these are 80Gbps USB-C (USB4 v2) ports being driven by Intel’s JHL9580 (Barlow Ridge) Thunderbolt 5 controller. As a result, they are not only far faster than USB-C ports in many other systems (which frequently top out at just 10Gbps), but they also support PCIe tunneling, 15W power delivery, and Thunderbolt networking. In practice, they are Thunderbolt 5 ports in all but name: Minisforum does not advertise the system as being Thunderbolt 5 capable, implying that they have not received official certification for the standard. Still, all of Thunderbolt 5’s major features are supported thanks to the Intel controller.

USB ports aside, rounding out the front panel is the obligatory 3.5mm combo jack and a power button.

Zooming back out and taking a look at the rest of the system, the MS-02 Ultra is designed to work in both horizontal and vertical orientations.

Along with rubber feet on the bottom for the horizontal layout, the system also has a set of feet on its right side, allowing it to safely and securely sit vertically.

Flipping over to the back, we get to see the other piece of the picture with regards to the rest of the various I/O ports and expansion brackets for the system, along with all of the exhaust vents that are helping to keep the guts of the MS-02 cool.

For all of the MS-02 models, Minisforum is including a trio of 10Gbps USB-A ports, as well as a 40Gbps USB-C port that is being driven by the SoC’s integrated controller. This USB-C port pulls double-duty as the system’s rear DisplayPort output as well, thanks to DP alt mode. Otherwise, if you need a dedicated and traditional display output, the system offers a single HDMI 2.1 port that is driven by the SoC’s integrated graphics.

All MS-02 systems also come with a pair of RJ45 Ethernet ports. On the left is a 2.5GbE port driven by Intel’s i226-LM controller. And to the right of that is an even faster 10GbE port that is driven by Realtek’s RTL8127 controller, the first real budget 10GbE controller on the market that we are quickly seeing peripheral and system manufacturers adopt.

If 10Gb Ethernet is still too slow for you, then you are in luck: on the high-end 285HX model like we are reviewing today, Minisforum also offers a pair of SFP28 25GbE ports that are situated on a PCIe expansion card. These connectors are driven by an Intel E810 controller, and suffice it to say, they give the MS-02 Ultra an ample amount of networking bandwidth, as well as access to more enterprise-focused features like RDMA support. Which, in turn, is a big part of the reason that Minisforum is pitching the MS-02 Ultra at the AI developer market (though admittedly, almost everyone is pitching products at the AI developer market these days).

The inclusion of SFP+ ports on the original MS-01 was one of that system’s more novel features. Here, the added SFP28 ports add a lot of cost and considerable power consumption, which is why it is only standard on the higher-end model.

I/O ports aside, the final connector you will find on the back of the MS-02 Ultra is a standard C14 power connector. Minisforum has integrated a 350 Watt power supply into the system – another major upgrade over the MS-01 – so gone is the need for an external power adapter and DC barrel connector. Though this does come at the cost of some space.

Meanwhile, our look at the MS-02’s chassis gives us a hint at the PCIe expandability offered by the system. Along with the PCIe slot that Minisforum’s 25GbE networking card sits in, there are two more unoccupied PCIe slots sitting behind blank PCIe brackets. This is another major upgrade for the MS-02 Ultra as compared to the MS-01, which only had a single expansion slot. The dual slots – or rather the double-width bay – affords the room to install a modern video card, as development of single-slot half-height cards has seemingly come to a halt. As a result, in order to be able to accommodate newer dual-slot video cards, Minisforum has expanded its PCIe bay size accordingly. The cost of that is significantly upscaling the size of the system to match.

