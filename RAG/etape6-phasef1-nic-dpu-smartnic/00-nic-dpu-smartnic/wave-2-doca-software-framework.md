---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-2-doca-software-framework
title: "Wave 2 — DOCA software framework"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "Broadcom", "Google", "Intel", "Nvidia", "Oracle"]
dates: ["2022-05", "2026-09-22"]
keywords: ["accelerator", "agent", "amd", "asic", "benchmark", "cybersecurity", "ethernet", "gpu", "gpus", "helios", "hyperscaler", "inference"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [56, 113]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: dbfefd6b42fe18e5e247bfc627dacd1572c088aaf42a79e865d2811dc602623a
---

# Wave 2 — DOCA software framework

## Wave 2 — DOCA software framework

- **DOCA** is NVIDIA's data-center-on-a-chip architecture / SDK: runtime environment, orchestration tools for provisioning/updating/monitoring thousands of DPUs, libraries, APIs and applications (DPI, load balancing) [official].
- DOCA SDK 1.0 released alongside BlueField-3 (Nov 2021); documented versions now extend to DOCA 2.9.2 (docs.nvidia.com archive) [official].
- The **doca-platform** GitHub repo (crawled 2026-09-22) lists tested BlueField-3 DPU models for DPF (DOCA Platform Framework): B3240, B3220, B3210; container ecosystem versions pinned (OpenBao v2.6.1, OTel collector 0.157.0, Prometheus v3.13.2, prometheus-operator v0.93.0) [official] [source: https://github.com/nvidia/doca-platform/blob/HEAD/docs/public/platform-support.md ].
- DOCA underpins NVIDIA's "Inference Context Memory Storage Platform" (KV cache), tightly integrated with the **NIXL** library and **Dynamo** software [vendor-reported].
- DOCA backward compatibility / application portability across BlueField generations is a stated NVIDIA goal [vendor-reported].
- Morpheus AI cybersecurity framework can run on BlueField-3 as a telemetry agent (announced at BF-3 launch) [vendor-reported].
- No 2026 DOCA major-version number located in this research wave; the docs archive at 2.9.2 suggests 2.x line is current [gap].

---
*End of Waves 1–2. Continuing with AMD Pensando, Intel IPU, Marvell/Broadcom/Fungible, offloads deep-dive and pricing matrix.*

## Wave 3 — AMD Pensando portfolio

### 3.1 Origins and architecture

- Pensando Systems was founded in 2017 and acquired by AMD in **May 2022** for ~$1.9 billion [secondary].
- Pensando's core is a **P4-programmable** packet-processing architecture: the Distributed Services Card (DSC) line integrates P4 MPU engines with Arm cores for network, security and storage offload [vendor-reported].
- Salina 400 DPU: announced at AMD Advancing AI (2024); **400G networking with dual 400GE PCIe Gen 5**, **232 P4 MPU engines**, up to **128 GB DDR5**, **16× Arm Neoverse-N1 CPU cores** [vendor-reported] [source: https://techly.relevanpress.com/amd-announces-pensando-salina-400-dpu-_-pollara-40--129573/ ].

### 3.2 Pollara 400 AI NIC (shipping 2025–2026)

- **AMD Pensando Pollara 400** — industry's first NIC compliant with the **Ultra Ethernet Consortium (UEC) 1.0 specification**; announced 2024, began shipping 2025; Oracle is the first CSP to deploy [vendor-reported/secondary] [source: https://www.networkworld.com/article/4011095/amd-rolls-out-first-ultra-ethernet-compliant-nic.html ].
- Designed for PCIe Gen5 systems; supports **Ultra Ethernet, RDMA, and RCCL** (AMD's alternative to NCCL) for scale-out collective communication [vendor-reported].
- Fully programmable RDMA transport with hardware-based congestion control; intelligent multipathing and path-aware congestion control to avoid congested paths; fast failover [vendor-reported].
- RoCEv2 compatibility and interoperability with other NICs; GPU-to-GPU intelligent routing reduces latency [vendor-reported].
- AMD roadmap (2023–2027, per BigGo Finance reporting of AMD whitepaper, Sept 2026): **Pollara 400G NIC remains the per-GPU NIC through 2026**, one per GPU; optical modules stay at 400G DR4/DR8 through 2026 [secondary] [source: https://finance.biggo.com/news/1bad314f-c6c1-4284-b172-100bc3930dbf ].

### 3.3 Vulcano 800 AI NIC (2026)

- **AMD Pensando Vulcano 800** — 800 Gbps Ethernet NIC announced 2025/2026, UEC-compliant, PCIe Gen6, targeting rack-scale GPU clusters; directly challenges NVIDIA ConnectX-8 SuperNIC [secondary] [source: https://infotechlead.com/networking/amd-pensando-vulcano-800-challenges-nvidia-connectx-8-as-ai-networking-race-heats-up-broadcom-arista-and-marvell-intensify-competition-97488 ].
- Improves fault tolerance and accelerates collective GPU communication by reducing training bottlenecks [vendor-reported].
- AMD roadmap: **upgrade to "Pensando 800G" in 2027**, doubling bandwidth vs Pollara; 800G DR4 optics in 2027 [secondary].
- AMD's claim of 33% lower switching costs with Vulcano vs a competing 800G NIC on Broadcom Tomahawk 6 is based on AMD's internal comparison, not an independent benchmark [secondary].

### 3.4 Helios / MI400 systems (2026)

- CES 2026: AMD Helios AI rack systems use **Pensando Pollara 400G and Vulcano 800G NICs** for scale-out Ultra Ethernet networking alongside MI400-series GPUs (MI430X, MI440X, MI455X), EPYC Venice CPUs, Infinity Fabric and UALink for scale-up [vendor-reported/secondary] [source: https://www.techtimes.com/articles/313781/20260106/ces-2026-amd-details-helios-ai-rack-next-gen-instinct-mi400-gpus.htm ].
- AMD's roadmap targets 256 GPUs per 3-rack Pod by the MI500X era (2027) with co-packaged copper/optics; Pollara-class NICs remain the scale-out vehicle through 2026 [secondary].

## Wave 4 — Intel IPU line

### 4.1 Portfolio and roadmap

- Intel's IPU (Infrastructure Processing Unit) portfolio circa 2022–2023: three technologies — **Big Springs Canyon** (Xeon-D SoC with Ethernet), **Oak Springs Canyon** (FPGA-based gen 2, Xeon D + Intel Agilex FPGA, 2× 100GbE, PCIe Gen4), and **Mount Evans** (ASIC-based gen 2, 200G class) [secondary] [source: https://www.blocksandfiles.com/ai-ml/2023/10/04/intel-semi-detaching-altera-business-for-future-ipo/1590209 ].
- **Mount Evans / Intel IPU E2000**: Intel's first ASIC IPU, co-designed with Google Cloud; **200 Gbps** class; 16× Arm Neoverse N1 cores; P4 packet-processing logic; Intel QuickAssist-based accelerators; hardware-accelerated NVM storage interface (derived from Optane) to emulate NVMe devices; supports vSwitch offload, firewalls, virtual routing; shipping to Google and other service providers from 2022 [vendor-reported/secondary] [source: https://www.datacenterdynamics.com/en/news/intel-and-google-cloud-jointly-launch-data-center-accelerator-chip/ ].
- Google Cloud deploys E2000 in **C3 VMs** (4th Gen Xeon Sapphire Rapids): 200 Gbps low-latency networking encrypted at line rate with the open-source **PSP** protocol; C3 claimed 20% faster than C2 [vendor-reported].
- **Oak Springs Canyon** (officially C6000X, built by Silicom): Xeon D + Agilex FPGA; 2× 100GbE; PCIe Gen4 x16; 16 GB DDR4; Open vSwitch, NVMe-oF, RoCEv2 offloads; network virtualization function offload; shipping to Google and service providers [secondary].
- Roadmap (Intel Vision 2022, through 2026): gen-3 **400G IPUs — Mount Morgan (ASIC) and Hot Springs Canyon (FPGA)** slated 2023–2024; **800G IPUs in 2025–2026** [vendor-reported] [source: https://www.hpcwire.com/2022/05/10/intel-extends-ipu-roadmap-through-2026/ ].
- Software: **IPDK** (Infrastructure Programmer Development Kit) — common programming model for Intel's ASIC and FPGA IPUs using P4, DPDK, SPDK; Intel's stated goal was to make IPDK vendor-neutral for other DPUs too [vendor-reported] [source: https://www.servethehome.com/intel-ipu-plans-revealed-for-800gbps-ipus-in-2025-dpu/ ].

### 4.2 2026 status and strategic context [gap/unverified]

- No 2026 confirmation found in this research wave that Mount Morgan / Hot Springs Canyon shipped on the original 2023–2024 schedule, nor that an 800G IPU launched in 2025–2026 [gap]. Intel's IPU messaging has been quiet; the NEX business unit leadership changed (Sachin Katti appointed SVP/GM NEX, Feb 2023) and Intel semi-detached the **Altera** (FPGA) business toward a future IPO [secondary].
- The task scope mentioned **Intel F2000X-PL** — no facts on this part number were located in this wave; treat as [unverified] pending confirmation.
- Broader context: UEC 1.0 spec (mid-2025) and hyperscaler preference for open Ethernet (AMD Pollara, NVIDIA Spectrum-X) has shifted the DPU/IPU competitive frame toward AI networking, where Intel currently has no announced shipping 400G/800G IPU [analysis].

