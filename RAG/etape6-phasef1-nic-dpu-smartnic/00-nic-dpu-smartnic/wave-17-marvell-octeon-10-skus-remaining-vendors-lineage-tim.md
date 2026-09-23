---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-17-marvell-octeon-10-skus-remaining-vendors-lineage-tim
title: "Wave 17 — Marvell OCTEON 10 SKUs, remaining vendors, lineage timelines"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "Broadcom", "Google", "Intel", "Microsoft", "Nvidia", "Qualcomm", "TSMC"]
dates: ["2022-05", "2026-09-22"]
keywords: ["amd", "asic", "ethernet", "helios", "inference", "intel", "ipo", "nvidia", "pricing", "research", "serdes"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [604, 653]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: 3b56274f3c03d8e4494bf8fd2cd4c1e78253d818ac5d8ba401100b21a7fae6d8
---

# Wave 17 — Marvell OCTEON 10 SKUs, remaining vendors, lineage timelines

## Wave 17 — Marvell OCTEON 10 SKUs, remaining vendors, lineage timelines

### 17.1 OCTEON 10 SKU detail (from launch coverage)

Four initial configurations, 8–36 Neoverse N2 cores, TSMC 5 nm [secondary] [source: https://fuse.wikichip.org/news/5732/marvell-launches-5nm-octeon-10-dpus-with-neoverse-n2-cores-ai-acceleration/ and https://www.semiaccurate.com/2021/06/28/marvell-announces-their-5nm-octeon-10-dpu/ ]:

| SKU class | Cores | Target | Power | Key I/O |
|---|---|---|---|---|
| CN103 (103) | 8× N2 | Edge/fanless sub-25 W | <25 W | PCIe Gen5, 56G SerDes |
| CN106 / 106S | 24× N2 (+switch) | Mainstream telco/enterprise | mid | 6 DDR5 controllers |
| DPU400 | 36× N2 | Data center | ~60 W | 12 DDR5 controllers @ 5200 MT/s, 8× PCIe Gen5 |
| (4th config) | 24× N2 | Telco CN line | mid | — |

- **CN102/CN103** (later launch): up to 8 Arm N2 cores; SPECint 36.5; **~1.5 SPECint/W**; 25 W (50% lower than prior-gen DPU); 20 Ethernet MACs with MACsec; inline IPsec engine = **50 Gbps IPsec on 50% of one N2 core**; DDR5 up to 5600 MT/s; PCIe 5.0 and 56G SerDes on CN103xx only; offload-DPU or primary-processor roles; fanless designs [official/secondary] [source: https://convergedigest.com/marvell-intros-two-new-octeon-10-dpus/ ].
- "Several of the world's largest networking equipment manufacturers have already incorporated the OCTEON 10 CN102 into product designs" [vendor-reported].
- Integrated **1 Terabit switch** = 16× 50GbE configurable ports (1G–100GE via SerDes combine); 256-bit MACsec on switch; time-sensitive networking (TSN) [secondary] [source: https://fuse.wikichip.org/news/5732/marvell-launches-5nm-octeon-10-dpus-with-neoverse-n2-cores-ai-acceleration/ ].
- Inline ML/AI engine (claimed 100× software-inference boost); VPP hardware accelerators (claimed >5× packet-processing speed); most of IPsec in hardware; vector packet processing over 400G datapath [vendor-reported] [source: https://www.techPowerUp.com/283905/marvell-extends-octeon-leadership-with-industrys-first-5nm-dpus ].
- Marvell claim: **3× performance, 50% lower power vs previous OCTEON** generation; first 5 nm DPU with Neoverse N2, first inline AI/ML acceleration, first integrated 1T switch, first VPP hardware acceleration [vendor-reported].
- PCIe-card development platform: 24 N2 cores, 2×100GbE form factor for enterprise/cloud app development [secondary].
- Primary market per SemiAccurate: 5G RAN (telco), with data-center/cloud secondary; unified software stack from core to edge [secondary] [source: https://www.semiaccurate.com/2021/06/28/marvell-announces-their-5nm-octeon-10-dpu/ ].

### 17.2 Remaining vendors (summary records)

- **Xsight Labs**: Israeli startup; announced "X2" 400G DPU (Arm-based, ~2022); positioned as Intel E810-class alternative. Current 2026 shipping/pricing status **not confirmed** [gap].
- **Kalray**: K200/K200-LP MPPA DPU processors; K200-LP targets service providers building next-gen storage devices (standalone, replaces host processors in some use cases) [secondary] [source: https://medium.com/@Asterfusion/the-most-comprehensive-smartnic-dpu-vendors-with-its-product-line-summary-c8cf9a4aad57 ]. 2026 product status not re-verified [gap].
- **Chelsio**: T7 (400G) / T6 (100G) Terminator-series adapters — TCP/iWARP offload heritage; 2026 line status not re-verified [gap].
- **Ethernity Networks**: FPGA/eASIC-based SmartNIC/5G UPF offload; 2026 status not re-verified [gap].
- **Silicom**: FPGA NICs (N5010/N5110A) for capture/appliances; merchant [secondary].
- **Asterfusion**: Helium SmartNIC (Arm SoC + ASIC + accelerators), SONiC/DASH-oriented [secondary].
- **Napatech** (NT200A02 FPGA SmartNIC): post-2025 ownership status **unconfirmed** — earlier working note linking it to Alphawave was erroneous (Qualcomm bought Alphawave). Flag for recheck [gap].

### 17.3 Lineage timelines

**NVIDIA BlueField lineage:**
Mellanox BlueField-1 (2018, 2×A72) → BlueField-2 (2020, 8×A72, 200G) → BlueField-3 (2022 launch, 16×A78, 400G, PCIe Gen5) → BlueField-4 (2026, 64× Neoverse V2, 800G, PCIe Gen6, ConnectX-9) → BlueField-5 (2028, Feynman) [official/secondary].

**AMD Pensando lineage:**
Capri (DSC, 25/100G) → Elba (DSC-2, 200G) → Salina 400 (dual 400G, 232 P4 engines) → Pollara 400 (UEC NIC, 2026) → Vulcano 800 (800G/Gen6, Helios) [secondary].

**Intel IPU lineage:**
Oak Springs Canyon (Xeon D + Agilex, 2021) → Mount Evans / E2000 (ASIC + N1, Google C3, 2022) → Mount Morgan / Hot Springs Canyon (400G, planned 2023-24 — shipment unconfirmed) → F2000X-PL (unverified) → 800G IPU (planned 2025-26 — unconfirmed) [secondary/gap].

**Broadcom NIC lineage:**
NetXtreme E-Series (BCM57508 Thor 200G) → Thor 2 400G (2024) → Thor Ultra 800G UEC 1.0 (~2025) [secondary].

**Consolidation timeline:** Fungible → Microsoft (Jan 2023, ~$190M); Pensando → AMD (May 2022, $1.9B); Qualcomm → Alphawave Semi ($2.4B, closed ~Q4 2025); Intel/Altera semi-detached for IPO [secondary].

---
*End of Phase F1. 17 waves, append-only. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*

