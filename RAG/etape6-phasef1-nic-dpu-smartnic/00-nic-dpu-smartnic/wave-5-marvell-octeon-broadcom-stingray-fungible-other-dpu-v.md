---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-5-marvell-octeon-broadcom-stingray-fungible-other-dpu-v
title: "Wave 5 — Marvell OCTEON, Broadcom Stingray, Fungible, other DPU vendors"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "AWS", "Broadcom", "Google", "Intel", "Meta", "Microsoft", "Nvidia"]
dates: ["2021-06", "2023-01-09"]
keywords: ["accelerator", "acquisition", "amd", "asic", "aws", "compute", "datacenter", "disaggregated", "ethernet", "hyperscaler", "intel", "latency"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [114, 168]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: aa100f65cc3f855a8b704285f9ff1df78d0f353a7e0e6780be40b663f44c588b
---

# Wave 5 — Marvell OCTEON, Broadcom Stingray, Fungible, other DPU vendors

## Wave 5 — Marvell OCTEON, Broadcom Stingray, Fungible, other DPU vendors

### 5.1 Marvell OCTEON 10 DPU family

- OCTEON 10: industry's first 5nm DPU family, first with **Arm Neoverse N2** cores and **Armv9 + SVE2** (2×128b vector units); announced June 2021 by Marvell [official/secondary] [source: https://www.techPowerUp.com/283905/marvell-extends-octeon-leadership-with-industrys-first-5nm-dpus ].
- 3× compute performance and 50% lower power vs previous OCTEON generation (vendor claim) [vendor-reported].
- Key hardware blocks: up to **1 terabit integrated switch**, true **inline crypto** (most of IPsec in hardware), **VPP (vector packet processing) hardware accelerators**, ML/AI inline accelerator (100× vs software per Marvell), eHSM/TrustZone secure boot, DDR5, **PCIe 5.0**, 56G SerDes, 20 Ethernet MACs with MACsec support [official] [source: http://www.marvell.com/content/dam/marvell/en/public-collateral/embedded-processors/marvell-octeon-10-dpu-platform-product-brief.pdf ].
- CN102/CN103 parts: 8–24 cores on telco CN line; **DPU400** datacenter part with 36 cores; up to 16× 50GbE ports (2× 400GbE or redundant 400GbE pairs); ~60W for the 36-core part per SemiAccurate analysis [secondary/independent].
- Example board: **ET3608-2P2S** (cloudswit.ch) with OCTEON 10 CN103XX — 8× Neoverse N2 @ 2.7 GHz, 8 MB L2 + 16 MB L3, DDR5 5600, 100 Gbps bidirectional; platform runs SONiC+VPP [secondary] [source: https://cloudswit.ch/blogs/marvell-octeon-10-cn103xx-open-router-gateway/ ].
- Positioning: telco/5G RAN primary market plus cloud/enterprise/carrier datacenters; 100/200 Gbps intelligent processing (routing), 80/160 Gbps firewall/IPSec/SSL-TLS [vendor-reported].
- Marvell's 2026 AI connectivity push: 1.6T optical DSPs (Ara platform, Ara T), but DPU-side OCTEON roadmap beyond OCTEON 10 was not detailed in this wave [gap]. Note: **OCTEON 12** mentioned in the task scope was not confirmed in search results — treat as [unverified].

### 5.2 Broadcom Stingray

- Broadcom Stingray SmartNIC/DPU family (Stingray PS225/PS1100, based on NetXtreme + Arm cores) was Broadcom's DPU answer; public roadmap updates since 2023 have been minimal [gap].
- As of 2026, Broadcom's networking focus visible in press is on **Jericho/Tomahawk/Ramona switch silicon** (Tomahawk 6: 102.4 Tbps, 128× 800GbE, production 2026) and custom AI ASICs (Google TPU partnership); Stingray DPU positioning vs BlueField/Pensando is unclear in current materials [analysis].
- No 2026 Stingray product announcements located [gap].

### 5.3 Fungible → Microsoft (acquired Jan 2023)

- Microsoft announced acquisition of Fungible on **2023-01-09**; terms undisclosed; reported ~**$190 million** [secondary] [source: https://www.datacenterknowledge.com/deals/microsoft-extends-data-center-ip-with-acquisition-of-dpu-startup-fungible ].
- Fungible (founded 2016 by Pradeep Sindhu and Bertrand Serlet; CEO Eric Hayes from 2021) had raised **$300M+** (SoftBank Vision Fund, Battery, Mayfield); ~250 employees [secondary].
- Assets: DPU microprocessor, **FunOS** software, **Fungible Storage Cluster** (all-flash NVMe-over-TCP disaggregated storage); FunPX partner program (2022) [secondary].
- Microsoft's stated intent: team joins Azure datacenter infrastructure engineering; focus on multiple DPU solutions, network innovation, hardware systems; reflects "long-term differentiated investments in datacenter infrastructure" [official].
- Post-acquisition product status (2026): no public standalone Fungible DPU product line; IP absorbed into Microsoft's internal DPU efforts (e.g., Azure Maia-adjacent infrastructure); direct hyperscaler-DPU competition noted at the time vs NVIDIA/AMD [analysis].

### 5.4 Other DPU/SmartNIC vendors (2026 landscape)

- **Napatech** (Copenhagen): FPGA-based SmartNICs for telco/cloud; acquired by **Alphawave Semi** (announced 2024–2025, ~$250M); NT200/NT400 series [secondary] — acquisition detail [unverified in this wave].
- **Xsight Labs** (Israel): E1 DPU ("X1") for cloud/hyperscale; Hot Chips 2026 comparisons referenced E1 vs BlueField-4 (V2 vs N2 cores) [independent] [source: https://www.servethehome.com/nvidia-bluefield-4-processor-at-hot-chips-2026/ ].
- **Ethernity Networks**: FPGA/ASIC 5G/edge UPF offload; small-cap vendor [secondary].
- **Kalray** (France): MPPA DPU for storage/compute offload [secondary].
- **AWS Nitro**: custom in-house DPU lineage (Nitro v5/v6) — not sold standalone; underpins EC2 virtualization, EBS, ENA networking [secondary].
- **Google IPU** (with Intel E2000 as above) and **Meta MTIA-adjacent** network offload are in-house efforts, not merchant products [analysis].
- **Chelsio**: T6/T7 adapters with iWARP RDMA, crypto offload; niche [secondary].
- DPU market sizing: Allied Market Research via DataCenterKnowledge — **$554M in 2021, projected $5.5B by 2031** at 27% CAGR [secondary]; 2026 actuals not located [gap].

---
*End of Waves 3–5. Next: offloads deep-dive, pricing matrix, competitive landscape, open items.*

## Wave 6 — NVIDIA ConnectX NIC line (DPU-adjacent, 2026)

### 6.1 ConnectX generations and positioning

- NVIDIA's ConnectX SmartNIC family is the ASIC foundation inside BlueField DPUs (BlueField-2 integrated ConnectX-6 Dx; BlueField-3 integrates ConnectX-7-class networking) [official/secondary].
- **ConnectX-8 SuperNIC**: up to **800 Gb/s** total bandwidth; InfiniBand 800/400/200/100 Gb/s, Ethernet 400/200/100/50/25 Gb/s (single-port 800G OSFP, dual-port 400G QSFP112); **PCIe Gen6 x16** (up to 48 lanes on datasheet); **NVIDIA Multi-Host** (up to 4 hosts); 16 million I/O channels; PCIe switch DPC; 256–4096B MTU [vendor-reported via datasheet copy] [source: https://cdn.pathfactory.com/assets/10412/contents/1195564/153be327-b398-454e-b6af-6a0aed28d960.pdf ].
- ConnectX-8 offloads: stateless TCP offload (IP/TCP/UDP checksum), LSO/LRO/GRO/TSS/RSS, **SR-IOV**, **ASAP²** (Accelerated Switching and Packet Processing) for SDN/VNF — OVS acceleration, overlay (VXLAN/GENEVE/NVGRE) acceleration, connection tracking (L4 firewall) + NAT, hierarchical QoS, header rewrite, flow mirroring, flow-based statistics/aging; RDMA + RoCEv2 acceleration; advanced programmable congestion control; **GPUDirect RDMA / GPUDirect Storage**; network compute: packet reordering, MPI acceleration, burst buffer offload, collectives, enhanced atomics, rendezvous protocol offload [vendor-reported].
- Portfolio: PCIe HHHL 1-port OSFP, 2-port QSFP112, dual mezzanine; form factors HHHL / Mezzanine / OCP 3.0; models **C8180** (single-port OSFP, XDR 800G; part 900-9X81E-00EX-DT0) and **C8240** (dual-port QSFP112, NDR, SocketDirect; part 900-9X81Q-00CN-ST0) [vendor-reported via reseller data].
- ConnectX-8 vs ConnectX-6 (vendor/reseller claims): 2× bandwidth, 50% lower latency for AI workloads, better power efficiency [vendor-reported].
- AMD's Vulcano 800 is positioned directly against ConnectX-8 SuperNIC (identical 800G headline bandwidth; differentiation on software programmability, congestion management, TCO) [secondary].

### 6.2 Street pricing (2026 snapshots)

All prices below are third-party retail/wholesale snapshots, not NVIDIA list prices; dates are crawl dates of the listings (Sept 2026 unless noted).

