---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/overview
title: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "Baidu", "Broadcom", "California", "Intel", "Microsoft", "Nvidia"]
dates: ["2025-10", "2026-09-22"]
keywords: ["amd", "intel", "nvidia", "pricing", "accelerator", "blackwell", "compute", "cybersecurity", "ethernet", "gpu", "gpus", "inference"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [1, 55]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: 6761f01781c9a1093b1600f0edb28b558b008d4aa6b752893ce2d2d194f52baa
---

# Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)

- **Scope:** Network interface cards (NIC), Data Processing Units (DPU), Infrastructure Processing Units (IPU) and SmartNICs: NVIDIA BlueField line, AMD Pensando portfolio, Intel IPU roadmap, Marvell OCTEON, Broadcom Stingray, Fungible/Microsoft, other DPU vendors; offload deep-dive (RDMA/RoCE, NVMe-oF, TLS/IPsec crypto, vSwitch/OVS, VirtIO); street pricing and competitive landscape 2026.
- **Date / cutoff:** 2026-09-22.
- **Method:** read-only web research (browser_search, browser_open); no live-browser visits, no purchases, nothing sent externally. No identifiers guessed; all SKUs, URLs, versions and prices are verbatim from tool results.
- **Provenance legend:** `[official]` vendor documentation/press; `[vendor-reported]` vendor claims without independent verification; `[independent]` third-party testing or analysis; `[secondary]` press/reseller reporting; `[unverified]` single-source or unconfirmed claims.
- **Line target:** ≥ 750 lines. Append-only: each wave adds sections; earlier waves are not modified.

---

## Wave 1 — NVIDIA BlueField family: BF-2, BF-3, BF-4

### 1.1 BlueField-2 (previous generation, still widely available)

- BlueField-2 DPU cards remain in retail channels in 2026, e.g. NVIDIA BlueField-2 DPU E-Series SmartNIC **MBF2M516A-CECOT** — Crypto enabled with Secure Boot, PCIe 4.0 x16, 2× 100GbE QSFP56, 16 GB DDR — listed at **$2,299.00** (MSRP $3,143.00) on SHI [secondary] [source: https://www.shi.com/product/44101934/NVIDIA-BlueField-2-DPU-E-Series-SmartNIC-MBF2M516A-CECOT ].
- NVIDIA BlueField-2 Ethernet DPU **MBF2M516A-CENOT** — Crypto enabled, PCIe 4.0 x8, 2× 25GbE SFP56, 8 GB RAM — **$2,043.00** on SHI [secondary] [source: https://www.shi.com/product/43607275/NVIDIA-BlueField-2-Ethernet-DPU-MBF2M516A-CENOT ].
- P-Series variant **MBF2H332A-AECOT** — 2× 25GbE SFP56, PCIe 4.0 x8, crypto enabled with secure boot, 1-year warranty — listed on SHI Germany (de.shi.com) [secondary].
- E-Series dual-port 100GbE variant **MBF2M516A-EECOT** at **$2,299.00** on SHI [secondary].
- Secondary market: used BlueField-2 200GbE (**BF2M515A**) at **$379.99** on eBay (239 sold, seller Pomona California) [secondary] [source: https://www.ebay.com/itm/125970047180 ] — shows steep depreciation of the previous generation.
- UK reseller Comms Express lists E-Series **MBF2M516A-CEEOT** / **MBF2M516C-CESOT** (100GbE dual-port QSFP56, PCIe Gen4 x16, 16 GB DDR, 1GbE OOB management) as price-on-application, 2–3 week delivery [secondary].

### 1.2 BlueField-3 (current shipping generation)

- BlueField-3 is NVIDIA's third-generation DPU ("data-center infrastructure-on-a-chip"), delivering up to **400 Gb/s Ethernet or NDR InfiniBand** [official] [source: https://nvidianews.nvidia.com/news/nvidia-extends-data-center-infrastructure-processing-roadmap-with-bluefield-3 ].
- CPU complex: up to **16× Armv8.2+ A78 "Hercules" cores**; 10× the accelerated compute power of the previous generation and 4× the cryptography acceleration [official].
- First DPU to support **PCIe Gen5** and time-synchronized data-center acceleration [official].
- Offloads: NVMe-oF / NVMe/TCP, decompression, erasure coding, SNAP storage virtualization; networking: RoCE, Zero-Touch RoCE, ASAP² SDN acceleration, VXLAN/GENEVE/NVGRE; virtualization: SR-IOV, VirtIO acceleration, overlay offloads, connection tracking; HPC/AI: GPUDirect, GPUDirect Storage, MPI tag matching [vendor-reported].
- Security engine: secure boot, PKA root-of-trust, MACsec/IPsec/TLS, flash encryption [vendor-reported].
- Programmable datapath: 16-core / 256-thread accelerator for SDN & NFV [vendor-reported].
- Official part numbers (from NVIDIA DOCA installation guide, supported/tested platforms) [official] [source: https://docs.nvidia.com/doca/archive/2-9-2/doca+installation+guide+for+linux/index.html ]:
  - **900-9D3B6-00CN-AB0** — BlueField-3 B3240 P-Series dual-slot FHHL; 400GbE/NDR IB default; dual QSFP112; PCIe Gen5 x16 (x16 extension option); 16 Arm cores; 32 GB DDR; BMC; crypto enabled.
  - **900-9D3B6-00CV-AAH** — BlueField-3 B3220 P-Series FHHL; 200GbE default / NDR200 IB; dual QSFP112; 16 Arm cores; 32 GB DDR; crypto enabled.
  - **900-9D3B6-00SN-AB0** / **900-9D3B6-00SV-AA0** — same B3240/B3220 with crypto disabled.
  - **900-9D3B4-00CC-EA0** / **900-9D3B4-00SC-EA0** — BlueField-3 B3210L E-series FHHL **SuperNIC**; 100GbE default / HDR100 IB; dual QSFP112; PCIe Gen5 x16; 8 Arm cores; 16 GB DDR.
  - **900-9D3B4-00EN-EA0** / **900-9D3B4-00PN-EA0** — BlueField-3 B3140L E-series FHHL SuperNIC; 400GbE / NDR IB; single QSFP112; 8 Arm cores; 16 GB DDR.
- Retail pricing (India reseller Grabnpay, crawled 2026-09-22): BlueField-3 DPU 16-core Arm, 32 GB DDR5, dual QSFP112, PCIe Gen5 — **Rs. 205,500.00** (was Rs. 245,980.00), SKU 900-9D3B6-00SC-AA, 20 in stock [secondary] [source: https://www.grabnpay.in/products/nvidia-bluefield-3-dpu-16-core-arm-32gb-ddr5-dual-qsfp112-pcie-gen5-crypto-options-enterprise-networking-infiniband ].
  - Note: Grabnpay description says "Dual DDR5 controllers + 16GB onboard DDR5" while NVIDIA docs list 32 GB on the P-series — reseller copy may refer to a different sub-model; treat with caution [unverified].
- Ecosystem adoption (from NVIDIA's BF-3 launch, 2022): server vendors Dell Technologies, Inspur, Lenovo, Supermicro integrating BlueField DPUs; CSPs Baidu, JD.com, UCloud; hybrid cloud partners Canonical, Red Hat, VMware; cybersecurity Fortinet, Guardicore; storage DDN, NetApp, WekaIO; edge Cloudflare, F5, Juniper Networks [official].

### 1.3 BlueField-4 (announced Oct 2025; 2026 generation)

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

