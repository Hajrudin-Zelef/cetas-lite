---
id: collect-250926-servers-hardware/servers-hardware/nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026-1
title: "nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia", "TSMC"]
dates: []
keywords: ["compute", "nvidia", "rubin", "3nm", "attention", "blackwell", "consumer", "ethernet", "gpu", "gpus", "hbm4", "inference"]
source: docs/RAG/clean4/nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026.md
source_anchor: ""
source_lines: [1, 49]
sha256: b3cb6d832763315a97040a43adc87200579efd81ababa0cbba4ed76233703105
---

# nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026

CES may still informally be the *Consumer* Electronics Show. But that does not mean everyone got the memo – or at least, cares to pay attention to it. Case in point is NVIDIA, who in the first major chipmaker press conference of the day opted for nothing less than to announce the launch of Rubin, the company’s next-generation AI platform.

Revealing that all of the necessary chips are back from the fab and that systems are being brought up in NVIDIA’s labs, NVIDIA is wasting no time in bringing the hardware to market – and even less time in announcing this. The next stop on NVIDIA’s meticulous hardware roadmap, the Rubin GPU and its associated chips are designed to take NVIDIA’s performance and efficiency to the next level, improving per-GPU AI inference performance by 5x, AI training performance by 3.5x, and then feeding those new GPUs with more compute, memory, and networking resources than ever before. In short, Rubin is NVIDIA’s effort to outdo themselves – and to prove that they can do it while keeping their lead over the competition.

Over the last several years NVIDIA has outlined an extensive roadmap for the further development and evolution of their core CPU, GPU, and networking architectures. It is a roadmap that leaves little room for surprises, but at NVIDIA’s scale the company can no longer afford to surprise its customers. So instead, NVIDIA has pivoted to being clear on what to expect, when to expect it, and then ensuring the company hits those deadlines.

Rubin’s launch at CES, in that respect, is an “everything remains on schedule” announcement from NVIDIA. To be sure, the hardware is not shipping yet, and the ramp-up for production wo not even start until the second half of the year. But with working chips in hand, ecosystem partners already putting in their efforts, and more than a bit of bravado, NVIDIA has decided to kick off their 2026 by launching their next-generation platform.

And while NVIDIA is launching Rubin today, this was not a technical presentation from NVIDIA – though a separate technical blog is being published today. Instead, today’s press conference was a high-level look at what NVIDIA has been up to, outlining the chips behind the Rubin platform and the devices they will be going in, all while disclosing few key specifications for their next generation of hardware.


## Rubin Platform: CPU + GPU + DPU + NIC + NVLink + Ethernet Switching

As NVIDIA has been outlining for the past year or so, Rubin is both a GPU architecture and a larger platform. And though NVIDIA remains first and foremost a GPU company, with the Blackwell platform and even more with the Rubin platform, they would really like to sell you the whole system – and indeed a whole SuperPOD, if they can. As a result, the significance for NVIDIA on the development and launch of the Rubin platform is not that they developed one high-end chip – it is that they developed six.

### Rubin GPU

The star of the show is, of course, the Rubin GPU. NVIDIA’s next-generation GPU architecture and the lead implementation of it, the Rubin GPU (part number to be confirmed) is intended to leapfrog NVIDIA’s Blackwell and Blackwell Ultra (GB200) GPUs. NVIDIA is still holding the full details of how they will accomplish this close to their chest, but the big driver for the AI market is an updated transformer engine with support for compression – what NVIDIA refers to as one of their six “technology miracles.”

All told, the big Rubin GPU has been touted as offering 50 PFLOPS of compute for inference using NVIDIA’s NVFP4 format, which is five-times the inference performance of the Blackwell GPU. And while its training gains are not quite as immense, NVIDIA is now touting 35 PFLOPS of compute for training using the same NVFP4 format, which would put it three-and-a-half times ahead of Blackwell.

Like its predecessor, Rubin is a dual die chip, with two Rubin GPU dies on a single package. Fabbed on TSMC’s 3nm process, the two dies are “reticle sized”. Those dies, in turn, are paired with up to 288GB of HBM4 memory, the same capacity available today with Blackwell Ultra. But compared to the HBM3e memory used on Blackwell, Vera’s HBM4 will afford a total of 22 TB/second of memory bandwidth, a surprising 2.8x more than what Blackwell can offer.

In regards to transistor counts, the Rubin GPU measures in at 336 billion transistors, 1.6x that of Blackwell. Just how much power the chip will consume remains to be seen, however, as NVIDIA has yet to disclose the power consumption for the GPU. But they are claiming that it will offer eight-times the performance-per-watt of Blackwell when it comes to inference, which can be juxtaposed to the claimed 5x increase in performance.

### NVLink Switch

The Rubin GPU will in turn once again rely on NVLInk to reach out to its neighbors to form larger scale-up clusters of GPUs. NVLink 6 doubles the available bandwidth from NVLink 5, with each GPU now offering 3.6 TB/second of NVLink bandwidth. And with it, NVIDIA has developed a new NVLink Switch chip (chip #2).

As you would expect for a doubled transfer rate, NVIDIA has needed to move to 400Gbps SerDes for NVLink 6. And with the NVLink 6 Switch chip offering full all-to-all bandwidth to each GPU hooked up to it, that gives the switch chip a total of 28.8 GB/second of bandwidth. That much networking traffic in a tight space also generates quite a lot of heat, and as a result the NVLink 6 Switch chip must be liquid cooled.

### Vera CPU

Orchestrating all of this will be Vera, NVIDIA’s new high-end ARM-based CPU (chip #3), and the first part of the titular Vera Rubin duo. Each Vera CPU has been confirmed to be comprised of 88 CPU cores, codenamed Olympus, which is an NVIDIA-custom design that implements the Arm v9.2-A architecture. Internal architectural details are limited at the moment, but NVIDIA has confirmed that it is an SMT-capable design – allowing for 176 threads – via NVIDIA’s spatial multi-threading technology. At a high level, NVIDIA is touting that Vera will offer twice the data processing performance and twice the compression performance of Grace.

Vera will be attached to up to 1.5TB of LPDDR5X memory (3x that of Grace) using SOCAMM modules, which NVIDIA developed in conjunction with Micron. The move to modular memory is designed to resolve one of the few drawbacks of Grace-based GB200, namely its use of unchangeable soldered-down memory. All told, a Vera CPU will have access to 1.2 TB/second of memory bandwidth, a bit over 2x the bandwidth of Grace.

A typical DGX node will feature one Vera CPU that is paired with two Rubin GPUs, and connected to those GPUs via NVLink-C2C, the latest generation slated to offer 1.8 TB/second of bandwidth.

Vera will also be the final piece of the puzzle for NVIDIA to offer rack-scale confidential computing. While Blackwell could already handle encrypted workloads, Grace could not, limiting the size of the confidential domain to just the GPU. But with Vera now fully in sync with NVIDIA’s confidential computing technologies on Rubin, it will be possible to have the entire rack encrypted.

### ConnectX-9 NIC & BlueField 4 DPU

Providing more traditional Ethernet connectivity to the Rubin platform will fall on the latest generation of NVIDIA’s ConnectX NIC and associated BlueField DPU technology.

The ConnectX-9 NIC (chip #4) will offer 1.6 Tb/second of networking bandwidth by utilizing 200G PAM4 SerDes. And in larger, multi-rack configurations, it will provide the basis for scale-out networking for larger clusters.

Its partner in crime (prevention) is the BlueField 4 DPU (chip #5). The latest BlueField is a mix of old and new, combining a 64 core Grace CPU with its own ConnectX-9 NIC. NVIDIA is touting BlueField 4 as offering twice the bandwidth, three-times the memory bandwidth, and six-times the compute performance as BlueField 3.

