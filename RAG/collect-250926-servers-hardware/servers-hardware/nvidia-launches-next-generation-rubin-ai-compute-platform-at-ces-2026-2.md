---
id: collect-250926-servers-hardware/servers-hardware/nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026-2
title: "nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["compute", "nvidia", "rubin", "blackwell", "cost", "ethernet", "gpu", "gpus", "hbm4", "inference", "kv cache", "lpddr5x"]
source: docs/RAG/clean4/nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026.md
source_anchor: ""
source_lines: [50, 95]
sha256: 766d728f28e785c36ebe7bbf1460a75446fbbafafc5239274bc7db9523f9d261
---

# nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026

### Spectrum-6 Ethernet Switching & Co-packaged Optics

Finally, tying all the rest of the Rubin platform hardware together will be NVIDIA’s latest Spectrum Ethernet switch, the Spectrum-6 (chip #6). This will be the heart of NVIDIA’s Spectrum-X switches, which for the first time for the company will implement co-packaged optics to rein in on power consumption.

Separate from NVIDIA’s various GPU boxes, the two key products for Ethernet switching will be SN6800 and SN6810 switches. A beast in any sense of the word, the SN6800 will offer 512 ports of 800G Ethernet or 2048 ports at 200G, for an aggregate bandwidth of 409.6 Tb/second. Meanwhile the smaller SN6810 will feature 128 ports of 800G or 512 ports of 200G – one quarter that of the SN6800, and at 102.4 Tb/sec aggregate bandwidth, one quarter of the throughput as well.

With co-packaged optics as the significant technical innovation here, NVIDIA is hoping to crack the formula for both reducing power consumption and improving the reliability of these large, high-performance switches. In short, by using a shared laser source and silicon photonics to modulate those lasers, the Spectrum-X switches will be able to provide high-speed optical networking with fewer fragile parts and without the high power cost of traditional optical networking.

NVIDIA certainly has high hopes for the switches, and is touting some very impressive numbers to go with them. According to the company, Spectrum-X switches with co-packaged optics will be able to offer 5x the power efficiency and 10x the reliability of an equivalent traditional networking switch.

## The Hardware: Vera Rubin NVL72 and HGX Rubin NVL8

All of these new chips, in turn, will be going into various NVIDIA systems. Ahead of Rubin’s full commercial availability, NVIDIA is re-affirming that they will be offering two types of Rubin systems. For those going all-in on the NVIDIA ecosystem, there is an updated version of NVIDIA’s rack-scale NVL72 system. Meanwhile for those who need to keep a toe in the world of x86, NVIDIA will be offering a next-generation 8-way HGX carrier design, the HGX Rubin NVL8.

Like its predecessor, Grace Blackwell NVL72, a single Vera Rubin NVL72 rack (previously known as NVL144) will be comprised of 72 Vera Rubin GPUs (144 GPU dies) as well as 36 Vera CPUs. As this is a purely scale-up solution, all of the GPUs are connected to each other via NVLInk switches. The sum-total of all of this hardware is a whopping 220 trillion transistors, and undoubtedly an obscene amount of power, as well.

With regards to performance figures, virtually all the high-level numbers show the same level of improvement as a single Rubin GPU. So a Rubin NVL72 rack offers 5x the inference performance, 3.5x the training performance, etc. Or to put that in numbers, 3.6 EFLOPS of compute for inference and 2.5 EFLOPS of compute for training. Backing that performance will be a sum total of 54TB of LPDDR5X (2.5x that of GB200 NVL72), 20.7TB of HBM4 memory (1.5x that of GB200 NVL72), and a cumulative 1.6PB/second of HBM4 bandwidth (2.8x that of GB200 NVL72).

Outside of the intricate chips that comprise an NVL72 rack, NVIDIA has also apparently gone back to the drawing board to simplify the design of the rack itself – with the Blackwell generation having generated plenty of murmurs of issues. The big shift here is that NVIDIA has moved to a fully cable free modular tray design, which massively cuts down on the amount of time needed to stand up an NVL72 rack. According to NVIDIA, the assembly of a single rack has gone from 100 minutes on Blackwell to just 6 minutes on Rubin. And by removing so many cables altogether, NVIDIA is expecting this to improve reliability as well, as the cables themselves were potential failure points.

Coupled with some further improvements to NVIDIA’s NVLink technology and their second-generation RAS engine, NVIDIA is promising that Vera Rubin NVL72 can achieve zero downtime for health checks and network maintenance.

### One More Tool: NVIDIA Inference Context Memory Storage Platform

Going beyond a single NVL72 rack means moving to scale-out computing, and with that NVIDIA increasingly relies on its networking technologies to help get the job done. Besides the aforementioned ConnectX-9 NICs, BlueField DPUs, and Spectrum-X switches, NVIDIA is also adding one more tool to their toolbox: a key value (KV) caching system called the NVIDIA Inference Context Memory Storage Platform.

A further use case for NVIDIA’s networking hardware, NVIDIA’s context memory storage platform is intended to serve as a KV cache for inferencing. As with other KV cache ideas (e.g. Enfabrica), the basic idea is to insert another tier of storage specifically to store key value pairs used in inference. Because the amount of context data generated by modern models is so significant – especially multistep models – there is too much context data to be stored at the node level. That leaves system operators with the option to either toss the data and recompute it later, or store it in some fashion. A KV cache does just this, allowing for context data to be saved and (relatively) quickly retrieved for further use rather than recomputed.

Essentially, NVIDIA is pitching this as POD-scale optimization technology to work around current bottlenecks in inference performance. At a high level, they are promising as much as a 5x improvement in inference performance and 5x better power efficiency – a significant improvement indeed.

As for the hardware used, this would leverage NVIDIA’s networking products, but it would not be an NVIDIA box ala the DGX boxes or a Spectrum-X switch. Instead, NVIDIA would provide the parts and let their partners pick up the rest, assembling context memory storage nodes out of SSDs for storage and BlueField/ConnectX hardware to link back to the respective AI nodes.

Even though it is not an NVIDIA product per-se, NVIDIA is going all-in with their context memory storage technology. Not only is NVIDIA developing the necessary network hardware, but they have added the necessary software functionality to the CUDA stack as well, exposing it in Dynamo, DOCA, and other frameworks. It is only fitting then that they will be using it themselves for their biggest tier of systems, the SuperPOD.

### DGX SuperPODs: Now for Rubin

With a new GPU architecture and new NVL72 racks, NVIDIA’s hardware stack is completed by the grand amalgamation of all of that hardware, the DGX SuperPOD. As with the Blackwell generation, the SuperPOD is intended to both serve as a blueprint/proof-of-concept for scale-out GPU systems, but also a commercial product for NVIDIA to sell to customers.

NVIDIA will be offering two SuperPOD configurations to potential customers. The first, based around their Vera Rubin NVL72 racks, incorporates 8 of those racks, NVIDIA’s Specturm-X Ethernet and Quantum-X800 InfiniBand networking switches, BlueFlield 4 DPUs, and storage nodes for context memory storage. This is intended to be Vera Rubin at its highest level.

By the numbers, a single DGX SuperPOD with DGX Vera Rubin NVL72 will incorporate 8 NVL72 racks, for a total of 576 GPUs, 288 CPUs, and some 600TB of memory. In terms of total throughput, NVIDIA is touting an enormous 28.8 EFLOPS of compute performance at NVFP4 precision, and with enough NVLInk bandwidth that model partitioning within a rack will no longer be necessary.

Alternatively, for x86 users NVIDIA will also be offering SuperPODs built around DGX Rubin NVL8 nodes. This is a less dense option overall, with 64 NVL8 nodes being combined into a single SuperPOD to offer 512 GPUs. Otherwise this is a similar scale-out design, heavily leveraging NVIDIA’s hardware to scale DGX Rubin NVL8 beyond a single rack of nodes.

## Rubin Platform: Available in Second Half Of 2026

