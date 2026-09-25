---
id: collect-250926-servers-hardware/servers-hardware/nvidia-vera-rubin-nvl72-rack-and-more-in-the-pegatron-booth-at-gtc-2026
title: "nvidia-vera-rubin-nvl72-rack-and-more-in-the-pegatron-booth-at-gtc-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "rubin", "compute", "gpu", "gpus", "liquid cooling", "nvlink", "rack-scale", "vera rubin"]
source: docs/RAG/clean4/nvidia-vera-rubin-nvl72-rack-and-more-in-the-pegatron-booth-at-gtc-2026.md
source_anchor: ""
source_lines: [1, 21]
sha256: 41ae91a820f70d787fc76c92cd7f84ba30f59abe1c831a0529cb1ff80dc62fc5
---

# nvidia-vera-rubin-nvl72-rack-and-more-in-the-pegatron-booth-at-gtc-2026

At NVIDIA GTC 2026, we also stopped by the Pegatron booth, and it was packed with servers. That included everything from the NVIDIA Vera Rubin NVL72 rack to 2U Rubin NVL8 solutions to high-density PCIe GPU options.

For this, we have a 2-minute short that you can check out above. Let us get to the servers.

## Pegatron Vera Rubin NVL72 Rack and RA4803-72N3 Compute Node

This Pegatron SVR rack solution is built around the NVIDIA Vera Rubin NVL72 architecture. This is a liquid-cooled rack-scale system that unifies 72 Rubin GPUs with 36 Vera CPUs into a single platform designed for AI factory workloads. It also includes the high-speed NVLink switches in the center and through the rear spine, power shelves at the top and bottom, and some ancillary management networking.

Usually, we get to see these installed as full racks, but there was a bit more than that in the booth.

Here we see the RA4803-72N3 compute node installed in the rack. This is Pegatron’s implementation of the NVIDIA Rubin NVL72 compute tray, designed with optimized liquid cooling paths for the high-density GPU configuration.

The front faceplate of the RA4803-72N3 shows a new optimized layout for liquid cooling and serviceability. The design prioritizes accessibility while maintaining the dense compute configuration that makes NVL72 so powerful, including the 800Gbps networking and E1.S SSDs.

If that front faceplate looks familiar, that is because we have seen the I/O layout before. This is the NVIDIA ConnectX-9 SuperNIC that provides the networking backbone for Rubin systems. ConnectX-9 delivers up to 800Gb/s of bandwidth per port, which is critical for keeping all 72 GPUs in an NVL72 rack fed with data. It also has an E1.S SSD. There is a special connector at the rear so that it can be easily assembled and serviced. Also, the entire design is optimized for liquid-cooling.

In the center of the node, we see additional networking cages.

This layout is for the NVIDIA BlueField-4 DPU, shown here, which handles infrastructure processing and offloads networking, storage, and security workloads from the main CPUs. This is becoming a standard component in AI infrastructure to maximize compute efficiency and NVIDIA is leaning on the BueField-4 DPU to accelerate KV caches in the future, which is why the CPU in this generation is much bigger.

Next, let us take a look at some of the PCIe GPU servers on display.
