---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/part-2
title: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing) (part 2)"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["Nvidia"]
dates: ["2025-10", "2026-09-22"]
keywords: ["nvidia", "pricing", "blackwell", "compute", "ethernet", "gpu", "gpus", "inference", "kv cache", "memory", "rack-scale", "rubin"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [42, 55]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: 8153c10cc5f9ef7ed8da412ef2e7696de5e9f9230ba7096521fa6ce110f23c51
---

# Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing) (part 2)

- Announced at **NVIDIA GTC Washington D.C., October 2025** (GTC DC) as the next-generation DPU, succeeding BlueField-3; doubles throughput from 400 Gb/s to **800 Gb/s** [secondary] [source: https://www.servethehome.com/nvidia-bluefield-4-with-64-arm-cores-and-800g-networking-announced-for-2026/ ].
- **NVIDIA stated: "expected to launch in early availability as part of NVIDIA Vera Rubin platforms in 2026"** [vendor-reported].
- Key silicon: combines an NVIDIA **Grace CPU with 64 Arm Neoverse V2 cores at 1.7 GHz** and **ConnectX-9 networking**; 126 billion transistors [vendor-reported].
- Memory: **128 GB LPDDR5** (275 GB/s bandwidth per Hot Chips 2026 slides); host link **PCIe Gen6 x16** [vendor-reported/independent] [source: https://www.servethehome.com/nvidia-bluefield-4-processor-at-hot-chips-2026/ ].
- Networking: **800G Ethernet across 200G PAM4 SerDes**, inline encryption; also supports InfiniBand [vendor-reported].
- Hot Chips 2026 (Aug 2026): NVIDIA positioned BF-4 as the "AI-native scale-in infrastructure engine" — 7 Tb/s aggregate bandwidth on the "AI DPU" vs 200 Gb/s-class cloud DPUs; Vera Rubin needs 7 Tb/s per compute tray (four 1.6 Tb/s GPU scale-out links + 800 Gb/s scale-in to the DPU) [independent/vendor-reported].
- Positioning: essential building block for "AI factories" — HPC clusters for AI combining CPUs, Blackwell GPUs and BlueField-4 DPUs; Spectrum-X Ethernet co-design [vendor-reported].
- Jensen Huang on stage at GTC DC 2025: a driver for the new DPU is accelerating **KV cache functions** in NVIDIA systems (inference context for LLMs) [secondary].
- CES 2026: NVIDIA positioned BF-4 as the backbone of the **Inference Context Memory Storage Platform** — AI-native storage for KV cache sharing across rack-scale clusters, up to 5× power efficiency vs traditional storage (vendor claim), integrated with DOCA, NIXL library and Dynamo software, Spectrum-X for RDMA-based KV cache access [vendor-reported] [source: https://nvidianews.nvidia.com/news/nvidia-bluefield-4-powers-new-class-of-ai-native-storage-infrastructure-for-the-next-frontier-of-ai ].
- Storage partners building BF-4 platforms (announced H1 2026): AIC, Cloudian, DDN, Dell Technologies, HPE, Hitachi Vantara, IBM, Nutanix, Pure Storage, Supermicro, VAST Data, WEKA; availability **H2 2026** [vendor-reported].
- Features: multi-tenant networking, AI runtime security, cloud elasticity, native DOCA microservices (containerized services), multiservice architecture with service-function chaining [vendor-reported].
- GTC Washington reveal: "6× the compute power" and support for AI factories up to 4× larger than with BlueField-3 [vendor-reported].
- No public retail pricing found for BF-4 as of 2026-09-22 [gap].

