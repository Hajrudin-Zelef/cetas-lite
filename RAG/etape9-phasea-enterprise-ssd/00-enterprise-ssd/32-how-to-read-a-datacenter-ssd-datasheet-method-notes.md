---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/32-how-to-read-a-datacenter-ssd-datasheet-method-notes
title: "32. How to read a datacenter SSD datasheet (method notes)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: ["2026-09-22"]
keywords: ["datacenter", "latency", "nand", "research", "throughput"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [649, 708]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: f3d284ea679248992ca4a1b5733dfc0f4098409cedce611d0bde84bc7cb30945
---

# 32. How to read a datacenter SSD datasheet (method notes)

## 32. How to read a datacenter SSD datasheet (method notes)

- Headline IOPS/GB/s are best-capacity, high-queue-depth, steady-state numbers — your capacity and QD will differ; check the footnotes for test conditions (FIO job files, preconditioning per SNIA PTS) [independent guidance].
- Sequential TBW >> random TBW (Micron 9550: 201,200 vs 56,064 TBW on 30.72 TB PRO) — match the TBW figure to your workload's randomness [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Power figures: distinguish average RMS (sustained) from max (burst) — thermal design uses max; PSU/cooling budgets use average [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Latency: QD1 4 KB figures (Micron 9550: 60/15 µs) describe interactive workloads; high-QD average latency describes throughput workloads — they are different numbers [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- DWPD without a stated warranty period is meaningless — always pair DWPD with years (5 years is the enterprise norm here) [independent guidance].
- "Up to" capacities: the top SKU may be a different NAND configuration (more dies, higher OP) — per-capacity tables, not headlines, decide purchases [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146).
- SED vs non-SED SKUs have different part numbers and prices — confirm the exact SKU (e.g. Kioxia KCD8DPUG30T7 SED vs base) before comparing prices [secondary](https://corgitech.us/products/lenovo-thinksystem-read-intensive-30-72tb-2-5-pcie-5-0-sed-kioxia-cd8p-r-ssd).
- Reseller specs can lag or lead vendor datasheets (firmware revisions change performance) — the vendor datasheet is authoritative; the reseller listing is the purchasable configuration [independent guidance].

---

*End of Step 9 Phase A — Enterprise SSD hardware. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*

---

## 33. Firmware update and fleet management

- Update paths: vendor tools (Samsung, Micron Storage Executive, Kioxia, WD Dashboard enterprise, Solidigm tools) for plain-channel drives; OEM ISOs (Dell, HPE, Lenovo) for OEM-branded SKUs — mixing channels can brick warranty coverage [independent guidance].
- Signed firmware + secure boot: Micron 9550 (asymmetric roots of trust), Marvell SC5 (FIPS-compliant RoT) — firmware images are authenticated before execution, blocking malicious downgrade [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- NVMe-MI 1.2c out-of-band management enables firmware push and health polling without host OS cooperation — required for lights-out datacenters [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- SPDM 1.2 device authentication (Micron 9550, SK hynix PEB110): the BMC cryptographically verifies each SSD's identity before trusting it — supply-chain attack mitigation [vendor-reported](https://www.storagenewsletter.com/2024/10/09/sk-hynix-unveils-peb110-e1-s-gen5-pcie-up-to-8tb-ssd-for-data-centers/).
- Telemetry: OCP 2.5 standardized log pages (Micron 9550, Samsung PM9D3a) feed fleet analytics — track `percentage_used` and temperature histograms centrally [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Namespace management: 128–512 namespaces per drive (Phison X200: 128; Micron 7450: 132; Micron 9550: 512) — over-provision at the namespace level to tune WAF vs exposed capacity per tenant [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- Sanitize: crypto-erase (SIE/SED SKUs) vs block-erase — SED crypto-erase is near-instant; verify with `nvme sanitize-log` [independent guidance].

---

## 34. Security deep dive

- TCG Opal 2.0/2.02: self-encrypting drive standard on Samsung PM9A3/PM1743/PM9D3a, Micron 9550, Solidigm PS1010, WD SN861, Phison X200 — AES-XTS 256-bit media encryption [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- TCG Enterprise SSC: Samsung PM9A3's enterprise-focused security spec — finer-grained access control than Opal [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- TCG Ruby SSC: Kioxia CD8P option — newer datacenter-oriented spec, simpler than Opal [secondary](https://www.crn.in/news/kioxia-launches-new-pcie-5-0-ssds-for-enterprise-and-data-center-infrastructures/).
- SIE (Sanitize Instant Erase) vs SED: SIE offers crypto-erase without full Opal key management — Kioxia CM7/CD8P offer non-SED, SIE, and SED tiers [secondary](https://www.crn.in/news/kioxia-launches-new-pcie-5-0-ssds-for-enterprise-and-data-center-infrastructures/).
- FIPS 140-3: Kioxia CM7 SED designed for FIPS 140-3 compliance; Micron 9550 L2 certifiable; Solidigm PS1010 FIPS 140-3 compliant; Marvell SC5 FIPS-compliant RoT [secondary](https://www.tweaktown.com/news/87587/kioxia-cm7-series-enterprise-pcie-5-0-ssds-up-to-14gb-sec-30tb/index.html) [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- ISE (Instant Secure Erase): WD SN861 security SKU tier alongside TCG Opal 2.01 [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca232psp9x1-ssd-data-centre-3-2-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041686/).
- Supply chain: SPDM + signed firmware + secure boot form the 2026 baseline against interdiction and counterfeit drives — verify PSID labels on used purchases [independent guidance].

---

## 35. Thermal and power design numbers

| Drive | Idle | Avg/active | Max | Notes |
|---|---|---|---|---|
| Samsung PM9A3 (U.2/E1.S) | — | 11 W read / 13.5 W write | — | M.2: ≤8 W [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf) |
| Samsung PM1743 | — | 19.9 W (1.92 TB) / 25 W (15.36 TB) typical | — | dual-port [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD) |
| Kioxia CM7-V 6.4 TB | — | 25 W active | — | [secondary](https://www.Scan.co.uk/products/kioxia-64tb-cm7-v-u3-sie-pcie-gen5-1x4-2x2-u3-15mm-mix-use-3dwpd-enterprise-ssd) |
| Micron 9550 | — | 18 W seq read / 16 W seq write (avg RMS) | — | [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf) |
| Solidigm D7-PS1010 7.68 TB | 5 W | — | ~22–29 W | [secondary](https://techatlantix.com/blog/post/solidigm-d7-ps1010-review) |
| WD SN861 E1.S | <5 W | 12 W | 21 W | needs 1.5 m/s airflow @ 30 °C for max perf [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf) |
| Phison X200Z | — | <20 W active | — | pSLC [secondary](https://www.thessdreview.com/our-reviews/enterprise/phison-pascari-x200z-gen5-800gb-1-6tb-enterprise-ssd-review-slc-gold-commands-a-lightning-fast-60-dwpd-data-center-ssd/) |
| Marvell SC5 MV-SS1331 | — | — | <8.7 W | controller only [secondary](https://www.blocksandfiles.com/flash/2021/06/01/marvell-launches-first-pcie-5-ssd-controller/1616031) |
| Marvell SC5 MV-SS1333 | — | — | <9.8 W | controller only, 16ch [secondary](https://www.blocksandfiles.com/flash/2021/06/01/marvell-launches-first-pcie-5-ssd-controller/1616031) |

- Design rule: budget cooling for max, PSU for average + margin; Gen5 U.2 drives are 2–3x the power of Gen4 — chassis airflow designed for Gen4 fleets may be inadequate [independent guidance].
- Throttling: enterprise firmware throttles on NAND/controller temperature (~70–85 °C trip points vary) — sustained-throttle performance, not burst, is the honest number [independent guidance].

---

