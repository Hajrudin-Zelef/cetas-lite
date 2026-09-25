---
id: collect-250926-servers-hardware/servers-hardware/d-matrix-joins-the-nvidia-nvlink-fusion-platform
title: "d-matrix-joins-the-nvidia-nvlink-fusion-platform"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Meta", "Nvidia", "Qualcomm"]
dates: []
keywords: ["nvidia", "nvlink", "accelerator", "agi", "amd", "dram", "ethernet", "inference", "intel", "memory", "rack-scale"]
source: docs/RAG/clean4/d-matrix-joins-the-nvidia-nvlink-fusion-platform.md
source_anchor: ""
source_lines: [1, 17]
sha256: d146b2c72710c4d2946e2cac53277ab452dce4a3919d47702c237c5d2e162d4a
---

# d-matrix-joins-the-nvidia-nvlink-fusion-platform

This week, d-Matrix and NVIDIA had a big announcement. The companies announced that d-Matrix would bring its next-generation XPUs to the NVLink Fusion platform. A closer look at the announcement revealed a few interesting details.

## d-Matrix Joins the NVIDIA NVLink Fusion Platform

We covered the d-Matrix Raptor 3D-DRAM Accelerator for Generative Inference at Hot Chips 2026. The company has an interesting memory technology that was a very popular piece at the conference, easily cracking our top 5 pieces from the show. Part of the plan, of course, is scaling beyond one Raptor XPU to multiple XPUs and then to larger clusters.

This week, the company and NVIDIA announced that it will use NVLink Fusion technology, which allows the company to scale up in the MGX racks.

About a year ago, we covered the d-Matrix JetStream 400G Ethernet Card for Data Center Scale AI Inference. That was an interesting announcement, since d-Matrix was trying to show it could do networking along with the new AI accelerator.

With the new announcement, it looks like d-Matrix is using not just the NVLink Fusion chiplets, but also in the announcement was the other networking technologies. d-Matrix also seems to be using NVIDIA Vera CPUs, ConnectX-9, and BlueField-4 DPUs, NVLink switch trays, and Spectrum-X switching.

## Final Words

d-Matrix has a lot of momentum right now. NVLink Fusion gives the company a relatively standardized way to scale up and out without reinventing networking on FPGAs. It also gives the company a rack-scale platform that hyper-scalers and neoclouds understand how to deploy and maintain. For NVIDIA, even if d-Matrix gets customer wins for its accelerators, NVIDIA gets to provide the interconnect and CPUs found within that deployment. Notably, if d-Matrix wins a deployment, it is a CPU socket that AMD, Intel, and Qualcomm do not win, nor does the Arm AGI CPU.

Interesting development for AI infrastructure. The move toward tighter accelerator and interconnect integration shows how important high-speed data movement has become for scaling AI workloads.
