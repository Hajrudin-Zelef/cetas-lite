---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/13-comparison-matrix-gen5-flagships-top-skus
title: "13. Comparison matrix (Gen5 flagships, top SKUs)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: []
keywords: ["benchmark", "compute", "datacenter", "memory", "nand"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [290, 355]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: f87096e8a9d175306e832df59b93ceee3a783905bdaa5618dc5ce4e275d5d477
---

# 13. Comparison matrix (Gen5 flagships, top SKUs)

## 13. Comparison matrix (Gen5 flagships, top SKUs)

| Drive | Interface | NAND | Top capacity | Seq R/W (MB/s) | Rand R/W (KIOPS) | DWPD | MTBF | Notes |
|---|---|---|---|---|---|---|---|---|
| Samsung PM1743 | Gen5 x4 | TLC | 15.36 TB | 14,000 / 7,100 | 2,500 / 360 | 1 | 2.5M h | dual-port, E3.S+U.2 [secondary](https://www.shidirect.com/product/45832454/THINKSYSTEM-2.5IN-U.3-PM1743-15.36TB-READ-INTENSIVE-NVME-PCIE-5.0) |
| Samsung PM9D3a | Gen5 x4 | TLC | 30.72 TB | 12,000 / 6,800 | 2,000 / 400 | 1 | [unverified] | OCP 2.5, U.2/E1.S/E3.S [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/) |
| Kioxia CM7-R | Gen5 x4 | 112L TLC | 30.72 TB | 14,000 / 6,750 | 2,700 / 310 | 1 | 2.5M h | dual-port, SR-IOV [secondary](https://www.directdial.com/us/item/kioxia-7-68tb-cm7-r-series-enterprise-2-5-nvme-ssd-solid-state-drive/kcmyxrug7t68) |
| Kioxia CD8P-R | Gen5 x4 | 112L TLC | 30.72 TB | 12,000 / 5,500 | 2,000 / 200 | 1 | 2.5M h | single-port [secondary](https://www.digitec.ch/en/s1/product/kioxia-cd8p-r-series-kcd8xpug7t68-ssd-7680-gb-ssd-53141468) |
| Micron 9550 PRO | Gen5 x4 | 232L TLC | 30.72 TB | 14,000 / 7,600 | 2,800 / 380 | 1 | 2.0–2.5M h | in-house stack, OCP 2.0/2.5 [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf) |
| Micron 9550 MAX | Gen5 x4 | 232L TLC | 25.6 TB | 14,000 / 7,600 | 2,800 / 750 | 3 | 2.0–2.5M h | mixed-use [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf) |
| Solidigm D7-PS1010 | Gen5 x4 | 176L TLC | 15.36 TB | 14,500 / 9,300 | 3,100 / 400 | 1 | 2.5M h | UBER 1E-18 [secondary](https://techatlantix.com/blog/post/solidigm-d7-ps1010-review) |
| SK hynix PS1010 | Gen5 x4 | [unverified] | [unverified] | [unverified] | 3,100 / [unverified] | [unverified] | [unverified] | prior gen [secondary](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142) |
| WD SN861 | Gen5 x4 | [unverified] | 12.8 TB (3DWPD) | 13,700 / 7,500 | 3,300 / 800 | 1–3 | 2.5M h | FDP, OCP 2.0 [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca264psp9x1-ssd-data-centre-6-4-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041687/) |
| Phison X200E | Gen5 x4 | 176L eTLC | 30.72 TB fam. | 14,800 / 8,700 | 3,200 / 930 | 1–3 | 2.5M h | dual-port option [vendor-reported](https://www.phisonenterprise.com/wp-content/uploads/2025/01/PascariProductBrochure_X200_010925.pdf) |
| Kioxia CM9 | Gen5 x4 | BiCS8 CBA | [unverified] | [unverified] | 3,400 / [unverified] | [unverified] | [unverified] | teased only [vendor-reported](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142) |

- QLC capacity leaders: Solidigm D5-P5336 122.88 TB (Gen4); Sandisk UltraQLC SN670 256 TB and Kioxia LC9 245.8 TB announced (2 Tb BiCS8 QLC); Solidigm 245 TB QLC roadmap end-2026 [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- Benchmark non-comparability: random-IOPS figures use different queue depths, block sizes, and steady-state definitions (SNIA Enterprise PTS) — compare only same-methodology numbers [independent guidance].

---

## 14. Gaps, conflicts, and unverified claims register

1. [conflict] Samsung PM9A3 peak specs: product brief 6,800/4,000 MB/s + 1,000K/180K IOPS vs samsung.com page 6,900/4,100 + 1,100K/200K — likely best-capacity-case variance; capacity list also differs (brief: 960 GB–15.36 TB; page: 480 GB–15.36 TB U.2, 960 GB–7.68 TB U.2 in another listing). Treated as variance, flagged [official].
2. [conflict] WD SN861 NVMe revision: U.2 datasheets say NVMe 2.0/2.02; E1.S datasheet says NVMe 1.4b — spec-sheet revision variance, flagged [secondary].
3. [conflict] Micron 9550 PRO random-write IOPS: product brief table lists 280–400 KIOPS by capacity while the technical product spec text cites "up to 720 KIOPS" — different measurement definitions; flagged [official].
4. [conflict] Solidigm D7-PS1010 random read: 2,800K (TechPowerUp 7.68 TB) vs 3,100K (TechAtlantix) vs 3,200K (Tom's Hardware E1 variant) — capacity/form-factor variance; flagged [secondary].
5. [conflict] Samsung PM1743 read speed: 13,000 MB/s (2021 announcement) vs 14,000 MB/s (2026 reseller listings) — announcement target vs shipping product; flagged [vendor-reported].
6. [gap] SK hynix PE8111: no retrieved source confirms specs — not found [unverified].
7. [gap] Micron 7500: no specs captured in this pass [unverified].
8. [gap] InnoGrit enterprise controllers: no 2026-corroborated data retrieved [unverified].
9. [gap] Solidigm D7-P5810 (SLC, write-intensive): mentioned as sibling line, specs not captured [unverified].
10. [gap] SK hynix PEB110 mass-production status: announced for 2Q25 pending qualification; no 2026 confirmation retrieved [vendor-reported].
11. [gap] Solidigm 245 TB QLC SSD: roadmap "by end of 2026", not a shipped product [secondary].
12. [gap] Kioxia CM9: teased performance figures only; no datasheet, capacity, or ship date [vendor-reported].
13. [gap] Sandisk UltraQLC SN670 256 TB / Kioxia LC9 245.8 TB: announced, specs not captured [secondary].
14. [unverified] Reseller $/TB figures (SHI, Converge, computeruniverse, CorgiTech) are single-channel snapshots; not market averages [secondary].
15. [unverified] Lenovo-branded PM1743 1.92 TB at $8,968 looks like extreme channel markup — treat as channel artifact, not NAND-market signal [secondary].
16. [vendor-claim] "World's fastest datacenter SSD" (Micron 9550, 2024 launch) vs Kioxia CM9 teased 3.4M IOPS vs Phison X200 14,800 MB/s — marketing claims from different dates, not comparable [vendor-reported].
17. [unverified] Samsung V10 400+ layer TLC production date — demoed ISSCC 2025, no production announcement [secondary].
18. [unverified] JEDEC NL-SSD standard implied by SK hynix AIN-D — JEDEC has said nothing publicly [secondary].

---

## 15. Glossary

- **DWPD**: Drive Writes Per Day — full-capacity overwrites per day over warranty; 1 = read-intensive, 3 = mixed-use, 25–60 = pSLC cache tier.
- **TBW/PBW**: total terabytes/petabytes writable over warranty; TBW ≈ capacity × DWPD × 365 × years.
- **UBER**: Uncorrectable Bit Error Rate — enterprise standard 1 sector per 10^17 bits read (Solidigm PS1010: 1E-18).
- **MTBF/MTTF**: mean time between/to failures — population statistic, not a per-drive lifespan promise.
- **WAF**: write amplification factor — NAND writes ÷ host writes; OP, TRIM, ZNS/FDP lower it.
- **Overprovisioning (OP)**: spare NAND beyond advertised capacity (7–28% typical enterprise).
- **PLP**: power-loss protection — capacitor-backed flush of in-flight writes.
- **P/E cycles**: program/erase cycles per cell — TLC ~3,000, QLC ~1,000, PLC <500 (indicative).
- **CBA**: CMOS directly Bonded to Array — Kioxia/SanDisk hybrid-bonding architecture.
- **CoP/PUC**: Cell-on-Periphery / Peripheral Under Cell — logic under the array (Samsung, SK hynix).
- **pSLC**: pseudo-SLC — TLC/QLC dies run in SLC mode for endurance/speed at ~1/3 capacity.
- **ZNS / FDP**: Zoned Namespaces / Flexible Data Placement — host-managed write placement reducing WAF.
- **SR-IOV / CMB**: single-root I/O virtualization / controller memory buffer — NVMe virtualization features.
- **SPDM**: Security Protocol and Data Model — device authentication standard.
- **TCG Opal / Ruby / Enterprise**: self-encrypting-drive security specifications.
- **OCP 2.0/2.5**: Open Compute Project datacenter NVMe SSD specification + telemetry.
- **EDSFF**: Enterprise and Datacenter SSD Form Factor (E1.S/E3.S) — replacing 2.5" U.2 in new servers.

---

