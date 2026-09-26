---
id: collect-250926-servers-hardware/servers-hardware/amd-intros-instinct-mi350p-accelerator-cdna-4-comes-to-pcie-cards-1
title: "amd-intros-instinct-mi350p-accelerator-cdna-4-comes-to-pcie-cards"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["accelerator", "amd", "compute", "gpu", "gpus", "hbm", "inference", "memory", "nvidia"]
source: docs/RAG/clean4/amd-intros-instinct-mi350p-accelerator-cdna-4-comes-to-pcie-cards.md
source_anchor: ""
source_lines: [1, 46]
sha256: 59c6dff5dc02fde00fb3f3cf8eea3c0e3611913cac003e2d9b322bfa4ecfd187
---

# amd-intros-instinct-mi350p-accelerator-cdna-4-comes-to-pcie-cards

AMD this morning is launching a new member of their Instinct MI350 series of AI accelerators with a particularly interesting product: a PCIe card. Marking AMD’s first Instinct PCIe card product in nearly half a decade, the new MI350P sees AMD bring its current generation accelerator architecture to conventional PCIe cards. With it, the company is targeting customers who want to do on-premises AI inference, but either cannot support the high thermal and power density of current server nodes, or want to integrate the accelerators with existing hardware. In short, customers who are not buying AI hardware by the rack.

To accomplish this, AMD is essentially taking one of its MI350X accelerators and cutting it in half, resulting in a card with half as many compute resources, half as much memory, and perhaps most importantly, a bit over half of the power consumption. The end result is a scaled-down card that offers all of the AI functionality of the CDNA 4 architecture that underpins the rest of the series, but in a chip that is small enough and light enough in power needs to fit on an air-cooled PCIe card.

| **AMD Instinct MI350 Series Key Specs** |  |  | 
| **GPU** | **MI350P** | **MI350X** | 
| **Compute Units** | 128 | 256 | 
| **Matrix Cores** | 512 | 1024 | 
| **Peak Engine Clock** | 2200MHz | 2200MHz | 
| **Memory** | 144GB HBM3E | 288GB HBM3E | 
| **Memory Bandwidth** | 4TB/sec (8Gbps x 4096-bits) | 8TB/sec (8Gbps x 8192-bits) | 
| **Matrix Perf (MXFP8)** | 2.3 PFLOPS | 4.6 PFLOPS | 
| **I/O** | PCIe Gen5 x16 | PCIe Gen5 x16 7x Infinity Fabric (x16) | 
| **TBP** | 600W (Optional: 450W) | 1000W | 
| **Form Factor** | PCIe CEM, 10.5-inch FHFL DS | OAM | 
| **Architecture** | CDNA 4 | CDNA 4 | 


## Addressing the PCIe Market Hole

With the high demand for high performance server AI accelerators for what has now been the past several years, both market leader NVIDIA and long-time rival AMD have been primarily focusing on delivering server hardware by the node – or more recently, the whole rack. With both companies selling modular accelerators destined for compute trays and racks as fast as they can make them, they have had little need to focus on much else.

In practice, most of the demand for AI accelerators is for modular (OAM/SXM) accelerators that can go into modern compute nodes. These systems offer the highest hardware densities as well as the best scale-up and scale-out functionality thanks to the heavy use of proprietary GPU interconnects. But this leaves customers who need less hardware or are operating data centers that cannot accommodate an 11kW compute node of GPUs in a bind; it has left a hole in the overall market for AI accelerators.

For the past couple of years, both companies have been trying to plug this hole with products based on their respective workstation graphics GPUs, in AMD’s case, the Radeon AI Pro series. These graphics-based products offer a lot of the same AI functionality, but not all of it. In particular, they do not offer the higher performance of a flagship server GPU, nor the critical memory capacity and bandwidth afforded by its HBM. Consequently, while these graphics-based products can fill some needs for PCIe-based accelerators, they are not a perfect replacement by any means.

So for the first time since 2022’s Instinct MI210, AMD is releasing an Instinct PCIe card based on its current-generation architecture in order to fill this hole. The fundamental idea with the product is to give traditional server and data center customers access to the same caliber of AI hardware as their high-end modular parts, but in a readily replaceable and upgradable form factor. It is a bit more of a niche market these days, but with compute rack density and subsequent power and cooling needs presenting an issue for older data centers, it remains a very relevant one. And as an added bonus for AMD, it is a niche that rival NVIDIA is not currently addressing (nor has indicated they will be addressing), giving the company access to a market in which they will immediately be the front-runner.

## Instinct MI350P: Half a MI350X At Half the Power

Diving into the hardware itself, as alluded to earlier, the MI350P is essentially half of one of AMD’s flagship MI350X accelerators. When AMD first informed us about the product, I had assumed the PCIe card was being built with salvaged chips that did not meet the specifications for use in an MI350X accelerator. But once AMD sent over some details about the hardware, the reality became much more interesting.

In short, AMD is not using salvaged MI350X chips for this product. Instead, they are building a smaller chip especially for use on the MI350P by leveraging the original’s use of chiplets to make a smaller chip out of the same silicon. Whereas the MI350X was built from two I/O dies (IODs), each with four accelerator complex dies (XCDs) stacked on top (for a total of 8 XCDs), the MI350P’s chip is half of that. It is a single IOD with four XCDs, which is clocked identically to the MI350X and, at peak performance figures, offers half of the performance of AMD’s modular accelerator.

With the reduction in IODs also comes a reduction in memory capacity and bandwidth. 8 HBM3E stacks has become 4, resulting in a card with 144GB of HBM3E memory and a total memory capacity of 4TB/second – once again half as much as the MI350X. As a result, the MI350P is an almost perfect scale-down of the Mi350X, offering around half of the performance and half of the memory capacity.

On paper, the only aspect of the hardware that has not been halved is power consumption. Whereas the MI350X has a typical board power (TBP) rating of 1000W (and 1400W for the MI355X), the MI350P is rated for 600W. As a more standardized form factor, 600W is a magic number for PCIe cards as it is the limit defined by the PCIe CEM spec itself, with AMD opting to run the card as hot and fast as the spec allows. Even then, as not all servers can handle 600W PCIe cards, AMD is also offering a 450W TBP mode, which shaves off some performance to further bring down power consumption.

Physically, the card is a very standard and by-the-books full height full length (FHFL) dual slot card, which in both TBP configurations is designed to be air cooled. Typical for server cards, this is an entirely passive cooler design, with one large heatsink running the length of the card that is designed to be cooled by airflow coming from the server chassis itself. This configuration means that up to 8 cards can be installed in a single server tray.

Notably, however, AMD is not exposing its GPU-to-GPU Infinity Fabric links in any way on the MI350P. As a result, multi-card setups are limited to just using the PCIe bus (PCIe Gen5 x16) to communicate between the two of them. This is arguably the single biggest tradeoff of the MI350P versus the MI350X, as it limits the use of AI models larger than a single card’s memory pool. The end result is that an 8-card setup is a better fit for running 8 models than it is for one large model spread over 8 GPUs.

And speaking of multiple models, it is worth noting that the PCIe member of the Instinct MI350 family is also retaining the series’ partitioning support. With half as many XCDs, that limit is reduced accordingly to just 4 partitions, but otherwise a CPX configuration is the same 1 XCD with 36GB of memory setup as it is on the MI350X.

Finally, as far as the overall performance of the card is concerned, AMD has taken an interesting (and welcome) step for the MI350P by publishing both peak and typical/delivered performance figures for the card, something they did not do for the MI350X launch.

