---
id: etape6-phasef2-fpga/00-fpga/3-amd-alveo-accelerator-smartnic-cards
title: "3. AMD Alveo accelerator/SmartNIC cards"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AMD", "TSMC", "United States"]
dates: ["2026-01-14", "2026-05-19", "2026-09", "2026-09-22"]
keywords: ["accelerator", "amd", "acquisition", "asic", "datacenter", "dsp", "inference", "latency", "memory", "pricing"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [79, 127]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 5d76022a45d167336c57c746bde9b30d08ab434e905e30a5de96a8c54ad3fe5b
---

# 3. AMD Alveo accelerator/SmartNIC cards

## 3. AMD Alveo accelerator/SmartNIC cards

### 3.1 Alveo U45 (current network-accelerator recommendation)

- AMD's official SN1000 product page states: *"While the Alveo SN1022 adaptive SmartNIC is actively being supported for current users, AMD recommends that all new designs target the **Alveo U45** network accelerator"* — i.e., SN1000 family is in support/maintenance mode; U45 is the forward path `[official]` — source: https://www.amd.com/en/products/accelerators/alveo/sn1000/a-sn1022-p4.html `[official]`

### 3.2 Alveo SN1000 (SN1022) SmartNIC

- 100G FPGA-based composable SmartNIC on **Xilinx XCU26** (16 nm UltraScale+), **8 or 16 Arm A72 cores**, dual-QSFP **10/25/100Gb/s**, PCIe Gen4, FHHL, **75 W** `[vendor-reported]` — source: https://www.hpcwire.com/2021/02/24/xilinx-launches-alveo-sn1000-smartnic/ `[secondary]`
- Part number **A-SN1022-P4** `[official]`
- Software: composable data plane, virtio/ef100 with vDPA, DPDK, SPDK, OpenOnload, Linux XDP, embedded OVS offload, accelerated virtio-blk to Ceph/TCP `[vendor-reported]` — source: https://www.xilinx.com/publications/technology-briefs/xilinx-alveo-sn1000-technical-brief.pdf `[official]`
- Community HLS plugins (AES-128-CTR, VXLAN parser) built with Vitis HLS 2020.2 target the SN1000 plugin region `[secondary]` — source: https://github.com/noreaaa/smartnic_plugins `[secondary]`

### 3.3 Alveo U25

- 25G SmartNIC-class Alveo card; community Q&A (AMD Adaptive Support, crawled ~18 days before 2026-09-22) compares U25 vs SN1000 services (MAE, Vitis P4) and SoC Arm accessibility constraints — indicates both lines still have active users asking migration questions `[secondary]` — source: https://adaptivesupport.amd.com/s/question/0D52E00006ihQPhSAM/differences-between-the-alveo-u25-and-sn1000?language=en_US `[secondary]`

### 3.4 Alveo U280

- Launch pricing (2018): standard **US$5,995**, premium **US$9,995**; 8 GB HBM2 on premium variant `[secondary]` — source: https://www.anandtech.com/show/13399/xilinx-announces-alveo-u280-datacenter-accelerator-8gb-of-hbm2 `[secondary]`
- Secondary-market snapshot (eBay, September 2026 listings): used U280 cards around **US$999–1,450** depending on seller/condition — resale listings, not OEM pricing `[secondary]`

### 3.5 Alveo U250 / U200 / U50

- U250 launch price (Oct 2018): from **US$8,995** `[secondary]` — source: https://www.anandtech.com/show/13468/xilinx-announces-alveo-u200-u250-datacenter-accelerators `[secondary]`
- Secondary-market snapshots (eBay, September 2026): U250 listings roughly **US$2,250–5,479** (seller/condition dependent); U200 around **US$1,350**; these are resale listings, not OEM MSRP `[secondary]` — sources: https://www.ebay.com/itm/326896993846, https://www.ebay.com/p/4052111222, https://www.ebay.com/itm/227221217698 `[secondary]`
- Alinx-branded AXU-series carrier/eval boards (AXU2CGB etc.) listed ~**US$900–1,050** on eBay September 2026 — third-party dev boards, not AMD list price `[secondary]`

### 3.6 Alveo UL3524 / UL3422 (HFT)

- UL3524: purpose-built FPGA accelerator for ultra-low-latency electronic trading — **64 low-latency transceivers, 780K LUTs, 1,680 DSP slices**, Vivado flow, FINN integration for ML inference `[vendor-reported]` — source: https://www.globenewswire.com/fr/news-release/2023/09/27/2750440/0/en/AMD-Unveils-Purpose-Built-FPGA-Based-Accelerator-for-Ultra-Low-Latency-Electronic-Trading.html `[vendor-reported]`
- Exegy completed acquisition of **NovaSparks** on **2026-01-14**, consolidating FPGA-enabled ultra-low-latency market-data/normalization products; Exegy had previously acquired Vela Trading Systems and Enyx `[secondary]` — source: https://mondovisione.com/media-and-resources/news/exegy-acquires-novasparks-inc-extending-its-leadership-in-ultra-low-latency-fi-2026114/ `[secondary]`

### 3.7 Alveo V80 (see §2.4)

- MSRP **US$9,495**; 32 GB HBM2e, 820 GB/s, 4× QSFP56, 190 W `[vendor-reported]`

### 3.8 Alveo U30 (media)

- Secondary-market snapshot (eBay, September 2026): U30 media-accelerator cards listed ~**US$99–280** — resale listings only, no current OEM price found `[secondary]`

## 4. Achronix Speedster7t

- **AC7t800**: TSMC 7 nm, **711K logic elements**, **864 machine-learning processors (MLPs)**, **12 Tb/s 2D network-on-chip**, **6× GDDR6** subsystems (1.5 Tb/s external-memory bandwidth), **2× 400GbE** channels, **PCIe Gen5 x16** `[vendor-reported]` — sources: https://www.edn.com/fpga-is-optimized-for-high-bandwidth-workloads/ and https://www.achronix.com/documentation/speedster7t-fpga-datasheet-ds015 `[vendor-reported]`
- Speedster7t datasheet **version 2.3** listed with release date **2026-05-19** — current documentation as of 2026-09-22 `[official]`
- VectorPath S7t-VG6 accelerator card (BittWare): Speedster7t device, GDDR6, **1× QSFP-DD** for 400G or multiple 100G links `[secondary]` — source: https://www.zerif.co.uk/fpga-main-board-accelerators/achronix/bittware-s7t-vg6-achronix-speedster7t `[secondary]`
- Achronix positions Speedster7t for high-bandwidth workloads: AI/ML inference, networking/security offload, 5G/ORAN, computational storage `[vendor-reported]`
- Achronix Speedcore embedded-FPGA IP also offered for ASIC integration (separate product line from standalone Speedster7t) `[official]`

