---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/5-solidigm-enterprise-ssds
title: "5. Solidigm enterprise SSDs"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["TSMC"]
dates: ["2024-08-06"]
keywords: ["cost", "datacenter", "dram", "ipo", "memory", "nand", "research", "throughput"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [123, 172]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 2e2b42bff6911d92cbbc7df659bcf24caa11041c3edb574b71b6082ce3b143bc
---

# 5. Solidigm enterprise SSDs

## 5. Solidigm enterprise SSDs

### 5.1 D7-PS1010 / D7-PS1030 — Gen5 TLC performance line

- PCIe 5.0 x4, NVMe 2.0; SK hynix 176-layer 3D TLC charge-trap NAND (V7); Solidigm Atomos Prime controller (ACNT08M), 7 nm TSMC, 16 channels @ 1,600 MT/s, ARM Cortex-R8, DDR4 DRAM cache, pseudo-SLC cache [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-1-9-tb.d2148).
- Model-number convention: third digit = DWPD — PS1010 = 1 DWPD, PS1030 = 3 DWPD [secondary](https://www.6donline.com/solidigm-d7-ps1010-and-d7-ps1030-pcie-5-0-and-176l-tlc-datacenter-ssd-performance-play/).
- Capacities: 1.92/3.84/7.68/15.36 TB (PS1010); 1.6/3.2/6.4/12.8 TB (PS1030); U.2 15 mm and E3.S 7.5 mm [secondary](https://www.6donline.com/solidigm-d7-ps1010-and-d7-ps1030-pcie-5-0-and-176l-tlc-datacenter-ssd-performance-play/).
- Performance: 14,500 MB/s seq read; seq write 4,100 MB/s (1.6/1.92 TB) → 8,200 (3.2/3.84 TB) → 9,300 MB/s (≥6.4 TB); random read up to 3,100K IOPS (7.68 TB: 2,800K/400K per TechPowerUp; Tom's Hardware cites 3.2M/315K for the E1 variant); 5-year warranty [secondary](https://www.6donline.com/solidigm-d7-ps1010-and-d7-ps1030-pcie-5-0-and-176l-tlc-datacenter-ssd-performance-play/) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146) [secondary](https://www.tomshardware.com/pc-components/ssds/solidigm-touts-industrys-first-liquid-cooled-enterprise-ssd-d7-ps1010-is-an-e-1-pcie-5-0-drive-with-a-wrap-around-cold-plate).
- Reliability: MTBF 2.5M h (+25% vs D7-P5x20); UBER 1E-18 (10x better than prior gen); 14,016 TBW on 7.68 TB (1 DWPD) [secondary](https://www.6donline.com/solidigm-d7-ps1010-and-d7-ps1030-pcie-5-0-and-176l-tlc-datacenter-ssd-performance-play/) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146).
- Overprovisioning: ~14.5% (e.g. 1,039.4 GB on 7.68 TB SKU) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146).
- Security: AES-256, TCG Opal 2.02, FIPS 140-3 compliant; end-to-end data protection + PLP; power 5 W idle / ~22–29 W max (7.68 TB) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146) [secondary](https://techatlantix.com/blog/post/solidigm-d7-ps1010-review).
- Launch: August 6, 2024; still in production at cutoff [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146).
- Industry's first liquid-cooled enterprise SSD variant: E1.S/E.1 D7-PS1010 with wrap-around cold plate for AI datacenters (Tom's Hardware, ~Sept 2025) [secondary](https://www.tomshardware.com/pc-components/ssds/solidigm-touts-industrys-first-liquid-cooled-enterprise-ssd-d7-ps1010-is-an-e-1-pcie-5-0-drive-with-a-wrap-around-cold-plate).
- Predecessors: D7-P5520/P5620 (Gen4, 144L floating-gate TLC) — PS10x0 improves high-QD random IOPS 1.8–2.8x and seq throughput 2.0–2.2x [secondary](https://www.6donline.com/solidigm-d7-ps1010-and-d7-ps1030-pcie-5-0-and-176l-tlc-datacenter-ssd-performance-play/).

### 5.2 D5-P5336 — QLC high-capacity line

- QLC line for read-heavy workloads; 61.44 TB and 122.88 TB in standard U.2 form factor (PCIe Gen4) on 192-layer QLC NAND [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Positioning: up to 9:1 rack-space reduction and ~90% lower storage power vs hybrid HDD/TLC NAS configurations; ~25 W peak on 122 TB SKU; independent review (StorageReview 2025): "optimized for mid-sized I/O patterns commonly found in object storage and AI data pipelines" [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Roadmap: >200-layer QLC successor and a 245 TB QLC SSD anticipated by end of 2026 — vendor roadmap, not shipped product [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- QLC vs TLC endurance: ~1,000 P/E cycles vs ~3,000 — acceptable for AI read-heavy tiers, disqualifying for write-hot workloads [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Corporate: Solidigm ETF filed on Wall Street ahead of a September 4 IPO deadline as SK hynix seeks to consolidate its AI-memory position [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Sibling: D7-P5810, SLC-based, for extremely write-intensive workloads — specs not captured, gap Section 14 [unverified].

---

## 6. SK hynix enterprise SSDs

### 6.1 PS1010 (prior-gen Gen5)

- Mass-produced Gen5 datacenter SSD; rated 3.1M random read IOPS in competitive comparisons [secondary](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).
- SK hynix presents PEB110 as complementing the PS1010 portfolio for AI-era NAND demand [vendor-reported](https://www.thepickool.com/sk-hynix-develops-high-performance-ssd-for-data-centers/).

### 6.2 PEB110 E1.S — Gen5 E1.S (announced 2024-09)

- PCIe Gen5 (up to 32 GT/s); ~2x performance and >30% better power efficiency than its predecessor; built on 238-layer 4D NAND [vendor-reported](https://en.prnasia.com/releases/global/sk-hynix-develops-peb110-e1-s-for-data-centers-460322.shtml).
- Capacities 2/4/8 TB; OCP 2.5 support; first SK hynix datacenter SSD with SPDM (Security Protocol and Data Model) for secure server authentication/monitoring [vendor-reported](https://www.storagenewsletter.com/2024/10/09/sk-hynix-unveils-peb110-e1-s-gen5-pcie-up-to-8tb-ssd-for-data-centers/).
- Mass production planned 2Q25 pending customer qualification (announced Sept 2024; qualification-status at 2026 cutoff not confirmed — gap) [vendor-reported](https://en.prnasia.com/releases/global/sk-hynix-develops-peb110-e1-s-for-data-centers-460322.shtml).
- PE8111 (task-scope SKU): not found in any retrieved source — gap, Section 14 [unverified].

### 6.3 NAND technology: 321L, 321L QLC, split-cell PLC research

- SK hynix 321-layer Gen9 (4D PUC): mass-produced; 1 Tb TLC die (~20 Gb/mm²) [secondary](https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding).
- World's first 321-layer QLC 2 Tb die: mass production started, product shipping 1H 2026 [vendor-reported](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- AIN-D (AI NAND) product family teased: high-capacity low-cost low-power eSSD intended to replace HDDs; JEDEC NL-SSD standard implied but not confirmed; capacity undisclosed; PLC not mentioned as its technology [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- Split-cell 5-bit "Multi-Site Cell" (MSC) NAND demoed at IEDM Dec 2025: cell split into two half-sites of 6 voltage states each combining to 5 bits; claimed 20x faster reads than conventional PLC, better endurance via lower voltage stress; working wafers shown; manufacturability under evaluation — research demo, not a product [secondary](https://www.trendforce.com/news/2026/01/16/news-sk-hynix-unveils-5-bit-nand-that-splits-cells-delivers-20x-faster-reads/) [secondary](https://www.techradar.com/pro/samsungs-biggest-ssd-rival-has-a-clever-plan-to-make-ssd-bigger-but-not-cheaper-split-cell-5-bit-flash-sounds-like-raid-0-on-hard-drives-but-for-ssd).
- Conventional PLC (5 bits/cell, 32 voltage states): not commercial — read reliability too low, endurance too short; would add ~25% die capacity vs QLC [secondary](https://www.trendforce.com/news/2026/01/16/news-sk-hynix-unveils-5-bit-nand-that-splits-cells-delivers-20x-faster-reads/).

---

