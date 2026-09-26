---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/23-ai-workloads-and-qlc-economics
title: "23. AI workloads and QLC economics"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: finance
actors: ["Samsung"]
dates: []
keywords: ["cost", "datacenter", "inference", "ipo", "liquid cooling", "memory", "nand", "training"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [498, 525]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 000158d603947d420c04f2e9e37dfc4a6549d5612d33547e7f2389750939aaa8
---

# 23. AI workloads and QLC economics

- Endurance remaining: `nvme smart-log` → `percentage_used` (100 = worn out); compare `data_units_written` × 512,000 bytes to rated TBW; <70% used is the usual comfort zone for TLC [independent guidance].
- Error counters: any non-zero `media_errors` or `num_err_log_entries` growth is a reject; `available_spare` must be well above `available_spare_threshold` [independent guidance].
- Power-on hours vs wear: high POH with low `data_units_written` = lightly used (good); low POH with high writes = hammered (bad) — never use POH alone [independent guidance].
- Critical warnings byte: bit 0 (available spare), bit 1 (temperature), bit 2 (device reliability), bit 3 (read-only), bit 4 (volatile memory backup failed = PLP capacitor issue) — reject on bits 3/4 [independent guidance].
- Unsafe shutdowns: high counts suggest missing/failed PLP protection in prior deployment or dirty power — acceptable if SMART is otherwise clean [independent guidance].
- Firmware: check vendor firmware revision against current; Samsung/Kioxia/Micron OEM-branded drives may need OEM update ISOs (Dell/HPE/Lenovo) — plain-channel drives update via vendor tools [independent guidance].
- Sector format: enterprise drives may ship 520/528-byte or 4096-byte sectors — `nvme format --lbaf` to 512/4096 as needed (destroys data) [independent guidance].
- Namespace/SED state: `sedutil-cli` or `nvme` security commands to verify no TCG lock remains; locked drives are bricks without the PSID (printed on the label — photograph it) [independent guidance].
- Physical: U.2 drives need SFF-8643/SFF-8654 (SlimSAS) cables or U.2 backplanes; M.2→U.2 adapters exist but verify PCIe bifurcation support on the motherboard; 7 mm vs 15 mm z-height must match the bay/caddy [independent guidance].
- Thermal: Gen4 U.2 (8–14 W) runs on passive chassis airflow; Gen5 U.2 (19–29 W) wants direct airflow or heatsink kits — throttling starts ~70–75 °C NAND temp [independent guidance].
- Warranty/returns: used enterprise drives are usually sold as-is; factor one spare drive per RAID set into the budget [independent guidance].
- Price sanity: compare $/TB against the 2026 used band ($78–$127/TB for 3.84 TB SATA/SAS enterprise) and new NVMe band ($300–$1,172/TB) — anything far outside needs an explanation (OEM markup, scarcity, or a scam) [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).

---

## 23. AI workloads and QLC economics

- AI storage tiers show read-to-write ratios up to 20:1+, which matches QLC's asymmetric profile (fast sequential reads, weak random writes) — the core reason hyperscalers absorb QLC despite ~1,000 P/E cycles [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Solidigm D5-P5336 (122.88 TB): up to 9:1 rack-space reduction and ~90% lower storage power vs hybrid HDD/TLC NAS configs; ~25 W peak per drive; optimized for mid-sized I/O patterns in object storage and AI data pipelines [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Kioxia AI angle: CM9 teased explicitly "aimed at AI workloads" with up to 65% random-write and 95% sequential-write gains over CM7 [vendor-reported](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).
- SK hynix AIN-D family: high-capacity, low-cost, low-power eSSD intended to replace HDDs in AI datacenters; capacity undisclosed; JEDEC NL-SSD standard implied [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- Samsung PM9D3a positioning: "built for AI, hyperscale, and cloud" with OCP 2.5 compliance — hyperscale qualification (OCP) is the gating factor for AI-cluster SSD selection [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Liquid cooling arrives in storage: Solidigm's liquid-cooled D7-PS1010 E1 variant with wrap-around cold plate targets AI datacenters where air-cooling dense Gen5 E1.S is marginal [secondary](https://www.tomshardware.com/pc-components/ssds/solidigm-touts-industrys-first-liquid-cooled-enterprise-ssd-d7-ps1010-is-an-e-1-pcie-5-0-drive-with-a-wrap-around-cold-plate).
- Checkpointing workloads (LLM training): bursty sequential writes favor high sustained-write TLC (Micron 9550 MAX 10 GB/s-class, Phison X200) over QLC [independent guidance].
- Inference/serving (model weights, KV caches): read-heavy sequential/random-read — QLC sweet spot; 256 TB-class QLC drives (Sandisk SN670, Kioxia LC9 announced) target exactly this [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).

---

