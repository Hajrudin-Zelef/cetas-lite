---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/3-kioxia-enterprise-ssds
title: "3. Kioxia enterprise SSDs"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["datacenter", "dram", "latency", "nand", "pricing", "research"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [63, 122]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 61fb9fb0e20a1c4ba41181cbaa67dae5df31389b98ba996d74ef23402458ca77
---

# 3. Kioxia enterprise SSDs

## 3. Kioxia enterprise SSDs

### 3.1 CM7 — dual-port Gen5 performance flagship

- PCIe 5.0 + NVMe 2.0 compliant (PCI-SIG workshops); listed on PCI-SIG Integrator's List; up to 14,000 MB/s read on x4 [secondary](https://www.tweaktown.com/news/95016/kioxia-ssds-achieve-pcie-5-0-compliance-up-to-30tb-capacity-and-14-000-mb/index.html).
- NAND: Kioxia 112-layer BiCS FLASH (BiCS5) 3D TLC; proprietary controller + firmware [secondary](https://www.Scan.co.uk/products/kioxia-64tb-cm7-v-u3-sie-pcie-gen5-1x4-2x2-u3-15mm-mix-use-3dwpd-enterprise-ssd).
- Form factors: 2.5" 15 mm (U.3, SFF-TA-1001) and EDSFF E3.S 7.5 mm [secondary](https://www.shidirect.com/Product/46112649/KIOXIA-CM7-R-Series-SSD).
- Two classes: CM7-R read-intensive, 1 DWPD, up to 30.72 TB; CM7-V mixed-use, 3 DWPD, up to 12.80 TB; dual-port design for HA [official-derived](https://americas.kioxia.com/content/dam/kioxia/shared/business/ssd/enterprise-ssd/asset/KIOXIA_CM7_Namespace1_Performance_Brief.pdf).
- Performance: CM7-R up to 2,700K random read / 310K random write IOPS; CM7-V 6.4 TB example: 14,000/6,750 MB/s, 2,450K/550K IOPS, 25 W active, MTBF 2.5M h [secondary](https://www.Scan.co.uk/products/kioxia-64tb-cm7-v-u3-sie-pcie-gen5-1x4-2x2-u3-15mm-mix-use-3dwpd-enterprise-ssd) [secondary](https://www.directdial.com/us/item/kioxia-7-68tb-cm7-r-series-enterprise-2-5-nvme-ssd-solid-state-drive/kcmyxrug7t68).
- Kioxia lab numbers (performance brief, Oct 2023): FIO 100% seq read 14.5 GB/s, seq write ~7.1 GB/s, random read ~2,814K IOPS across namespace sizes — smaller namespaces can raise random-write IOPS [vendor-reported](https://americas.kioxia.com/content/dam/kioxia/shared/business/ssd/enterprise-ssd/asset/KIOXIA_CM7_Namespace1_Performance_Brief.pdf).
- Features: flash die failure protection, PLP, end-to-end data protection, SR-IOV, CMB, multistream writes, SGL, TCG Opal SED designed for FIPS 140-3 compliance, OCP Datacenter NVMe SSD v2.0 support (not all requirements) [secondary](https://www.tweaktown.com/news/87587/kioxia-cm7-series-enterprise-pcie-5-0-ssds-up-to-14gb-sec-30tb/index.html) [secondary](https://www.Scan.co.uk/products/kioxia-64tb-cm7-v-u3-sie-pcie-gen5-1x4-2x2-u3-15mm-mix-use-3dwpd-enterprise-ssd).
- 5-year limited warranty typical on reseller listings [secondary](https://www.shidirect.com/Product/46112649/KIOXIA-CM7-R-Series-SSD).

### 3.2 CM9 — teased 2026 successor (BiCS8 CBA)

- There is no CM8: Kioxia skipped it deliberately, tying SSD nomenclature to a skipped BiCS generation; CM9 jumps to BiCS8 218-layer with CBA (CMOS directly Bonded to Array) [secondary](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).
- Teased gains vs CM7: ~+65% random write, +55% random read, +95% sequential write; +55% seq-read and +75% seq-write perf-per-watt; headline 3.4M random read IOPS — ahead of Micron 9550 Pro (3.3M), FADU Echo (3.2M), SK hynix PS1010/Solidigm D7-PS1010/PS1030 (3.1M), Phison Pascari (3.0M) [vendor-reported](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).
- Possibility raised of a QLC BiCS8 U.2 variant reaching ~124 TB using 2 Tbit QLC chips (already shipped to Pure Storage for DirectFlash Modules) — roadmap speculation, not a product [secondary](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).

### 3.3 CD8P — single-port Gen5 datacenter line

- PCIe 5.0 (32 GT/s x4), NVMe 2.0, NVMe-MI 1.1d, OCP Datacenter NVMe SSD spec (not all requirements); single-port design; in-house controller; BiCS5 112L TLC [secondary](https://www.crn.in/news/kioxia-launches-new-pcie-5-0-ssds-for-enterprise-and-data-center-infrastructures/).
- CD8P-R (read intensive): up to 30.72 TB, 1 DWPD, up to 2,000K/200K IOPS, 12,000/5,500 MB/s; 60–80% higher seq read than Gen4 CD8-R [secondary](https://www.digitec.ch/en/s1/product/kioxia-cd8p-r-series-kcd8xpug7t68-ssd-7680-gb-ssd-53141468).
- CD8P-V (mixed use): up to 12.8 TB, 3 DWPD, up to 2,000K/400K IOPS [secondary](https://www.galaxus.ch/en/s1/product/kioxia-x121-cd8p-v-dssd-u2-pcie-sie-3200-gb-25-ssd-49654389).
- Latency: 99.999th percentile <250 µs standard random read; <1.8 ms OLTP-style mixed workloads [vendor-reported](https://www.crn.in/news/kioxia-launches-new-pcie-5-0-ssds-for-enterprise-and-data-center-infrastructures/).
- Security options: non-SED, SIE, SED (TCG Opal and Ruby SSCs); end-to-end data protection, PLP, flash die failure recovery [secondary](https://www.crn.in/news/kioxia-launches-new-pcie-5-0-ssds-for-enterprise-and-data-center-infrastructures/).
- 2026 pricing example (Lenovo-branded SED): CD8P-R 30.72 TB at $21,685.99 ≈ $706/TB — channel pricing [secondary](https://corgitech.us/products/lenovo-thinksystem-read-intensive-30-72tb-2-5-pcie-5-0-sed-kioxia-cd8p-r-ssd).
- CD8 (Gen4, prior gen) is the slower single-port datacenter predecessor [secondary](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).

---

## 4. Micron enterprise SSDs

### 4.1 Micron 9550 — Gen5 flagship (in-house stack)

- PCIe Gen5 x4, NVMe 2.0b; Micron's own controller, NAND, DRAM and firmware [vendor-reported](https://www.edn.com/data-center-ssds-achieve-blistering-speeds/).
- NAND: Micron 232-layer 3D TLC (G8) [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Two classes: 9550 PRO read-intensive 1 DWPD (3.84/7.68/15.36/30.72 TB); 9550 MAX mixed-use 3 DWPD (3.2/6.4/12.8/25.6 TB) [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Form factors: U.2 15 mm, E3.S 7.5 mm, E1.S 15 mm [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Performance: 14,000 MB/s seq read all SKUs; seq write 6,000–10,000 MB/s by capacity/class; random read up to 3,300 KIOPS; random write up to 900 KIOPS (MAX); 70/30 mixed 500–1,400 KIOPS [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Latency (QD1 4 KB): 60 µs read / 15 µs write typical [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Endurance TBW (random): PRO 7,008–56,064 TBW; MAX 17,520–140,160 TBW; (sequential): PRO 29,400–201,200 TBW; MAX 37,200–282,600 TBW — theoretical maxima, workload-dependent [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Reliability: MTTF 2.0M h @ 0–55 °C and 2.5M h @ 0–50 °C; UBER <1 sector per 10^17 bits read; 5-year warranty [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Security/data: SED SKUs, TCG Opal Rev 2.02, FIPS 140-3 L2 certifiable, digitally signed firmware, SPDM 1.2/SHA-512/RSA compliance options, TAA options; OCP 2.0 (r21) compliant, OCP 2.5 telemetry; NVMe-MI 1.2c; 512 namespaces; enterprise PLP; 512/4096-byte sectors [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- Power: seq read avg RMS up to 18 W; seq write up to 16 W [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Vendor positioning: "industry's fastest" datacenter SSD claim at launch (2024) — vendor-reported, see Section 14 [vendor-reported](https://www.edn.com/data-center-ssds-achieve-blistering-speeds/).
- 2026 street price example: 9550 PRO 7.68 TB U.2 at €1,597 excl. VAT ≈ €208/TB (computeruniverse) — one data point, not a market average [secondary](https://www.computeruniverse.net/en/p/3314-02V).

### 4.2 Micron 7450 — Gen4 mainstream

- PCIe Gen4 x4, NVMe 1.4b; Micron 176-layer 3D TLC NAND [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- PRO (1 DWPD, up to 15.36 TB) / MAX (3 DWPD, up to 12.8 TB); form factors U.3, E1.S, M.2 2280/22110 [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Performance: 6,800/5,600 MB/s seq; up to 1,000,000 random read IOPS / 400,000 write; 80 µs read / 15 µs write typical latency [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Endurance: PRO up to 28,000 TBW @ 1 DWPD; MAX up to 70,000 TBW @ 3 DWPD; MTTF 2M device hours; UBER <1/10^17; full PLP; 132 namespaces [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Security: TCG Opal Rev 2.01, digitally signed firmware; asymmetric roots of trust / secure boot / secure execution environment on MAX listings [secondary](https://eu.shi.com/Product/45125244/Micron-7450-MAX-SSD).
- Reseller example: 7450 Pro 7.68 TB U.3, 6,800/5,600 MB/s, 1,000K/180K IOPS, 1 DWPD, 2M h MTBF [secondary](https://www.bigw.com.au/product/micron-7450-pro-7-68tb-gen4-nvme-enterprise-ssd-u-3-6800-5600-mb-s-r-w-1000k-180k-iops-25700tbw-1dwpd-2m-hrs-mtbf-server-data-centre-5yrs/p/9902768775).
- Micron 7500 (Gen4 mainstream, 2024 launch): no specs captured in this research pass — gap, Section 14 [unverified].

---

