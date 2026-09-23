---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/7-western-digital-ultrastar-dc-sn861
title: "7. Western Digital Ultrastar DC SN861"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["AMD", "Intel", "Meta", "Microsoft", "Samsung", "TSMC"]
dates: []
keywords: ["amd", "asic", "compute", "consumer", "datacenter", "dram", "inference", "intel", "latency", "nand", "training"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [173, 227]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 3b05eac824dc3e01d083e8a2ae5a0bc0998b555109f3739c8076ce76da625d31
---

# 7. Western Digital Ultrastar DC SN861

## 7. Western Digital Ultrastar DC SN861

- PCIe Gen5 x4; NVMe 2.0/2.02 (U.2 datasheets) / NVMe 1.4b (E1.S datasheet) — spec-sheet revision variance, Section 14; NVMe-MI 1.2c; Flexible Data Placement (FDP); OCP 2.0 supportive [secondary](https://data.kosatec.de) [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf).
- U.2 15 mm: 1.92/3.84/7.68 TB (1 DWPD class), 3.2/6.4/12.8 TB (3 DWPD class); read up to 13,700 MB/s; write up to 7,500 MB/s (7.68 TB) / 3,600 (1.92 TB); random read up to 3,300K IOPS; random write 165K–800K by capacity/class [secondary](https://www.singular.com.cy/wd-ultrastar-dc-sn861-wus6ba119psp9x3-ssd-data-centre-192-tb-internal-25-u2-pcie-50-x4-nvme.html) [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca232psp9x1-ssd-data-centre-3-2-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041686/).
- E1.S 15 mm: 1.92/3.84/7.68 TB; read 12,100–13,700 MB/s; write 3,400–7,000 MB/s; random read 1,550K–2,850K; random write 140K–235K; 1 DWPD; 12 W avg / 21 W max operating, <5 W idle [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf).
- Latency: 65 µs read / 8 µs write (U.2); 70/10 µs (E1.S) — different SKUs/methods [secondary](https://www.singular.com.cy/wd-ultrastar-dc-sn861-wus6ba138psp9x3-ssd-data-centre-384-tb-internal-25-u2-pcie-50-x4-nvme.html?sl=el).
- Reliability: MTTF 2.5M h (projected); UBER 1 per 10^17; AFR 0.35% projected; 5-year warranty [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf).
- Features: PLP, end-to-end data path protection, TCG Opal 2.01 / ISE / SE security SKUs, high QoS [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca264psp9x1-ssd-data-centre-6-4-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041687/).
- Target workloads: AI training/inference, ML/deep learning, hyperscale cloud, HPC, big data [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf).
- 2026 price example: 3.2 TB U.2 at £572.81 excl. VAT ≈ £179/TB (Converge UK) — one data point [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca232psp9x1-ssd-data-centre-3-2-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041686/).

---

## 8. Phison Pascari — the controller vendor's own SSD brand

- Pascari is Phison's first-party enterprise SSD brand (launched 2024), distinct from its controller-ASIC business; IMAGIN+ is the companion custom-design service for hyperscalers [secondary](https://www.tweaktown.com/news/98293/pascari-is-new-enterprise-brand-from-phison-the-biggest-behind-scenes-name-in-storage/index.html).
- Lineup: X-Series (performance U.2/E3.S), D-Series (datacenter M.2/E1.S), S-Series (SATA 2.5"), B-Series (boot U.2/E1.S) [secondary](https://www.tweaktown.com/news/98293/pascari-is-new-enterprise-brand-from-phison-the-biggest-behind-scenes-name-in-storage/index.html).
- X200E (Gen5): PS5302-X2 16-channel controller + SK hynix 176L 4D eTLC + 8 GB DDR4-3200 DRAM; 1.6/3.2/6.4/12.8/25.6 TB (U.2), up to 30.72 TB family; up to 14,800/8,700 MB/s seq; up to 3,200K/930K random IOPS; 1 or 3 DWPD; 28% overprovisioning (6.4 TB: 5,961.63 GB usable); PLP, ISE/TCG Opal 2.0, AES-XTS 256, end-to-end protection, SECDED, sanitize, NVMe-MI/SMBus; 5-year warranty [vendor-reported](https://www.phisonenterprise.com/wp-content/uploads/2025/01/PascariProductBrochure_X200_010925.pdf) [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html).
- Independent review (TechPowerUp): 6.4 TB X200E measured 3,200K random read / 880K random write IOPS, 35,040 TBW @ 3 DWPD, MTBF 2.5M h; price at review $1,309.99 ≈ $205/TB — older price, not 2026 [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html).
- X200Z (pSLC "SLC gold"): TLC dies run in pseudo-SLC mode (~1/3 capacity); 800 GB/1.6/3.2 TB U.2 dual-port 2x2; up to 14.8/9.5 GB/s; 3.1M/950K IOPS; 60 DWPD — cache-tier drive; <20 W active [secondary](https://www.thessdreview.com/our-reviews/enterprise/phison-pascari-x200z-gen5-800gb-1-6tb-enterprise-ssd-review-slc-gold-commands-a-lightning-fast-60-dwpd-data-center-ssd/).
- X200 dual-port option: PCIe 5.0 1x4 single-port or 2x2 dual-port [vendor-reported](https://www.phisonenterprise.com/wp-content/uploads/2025/01/PascariProductBrochure_X200_010925.pdf).

---

## 9. SSD controllers (enterprise)

### 9.1 Marvell Bravera SC5

- Industry's first PCIe 5.0 SSD controller (announced 2021); two SKUs: MV-SS1331 (8 NAND channels, <8.7 W) and MV-SS1333 (16 channels, <9.8 W), both 20x20 mm [secondary](https://www.blocksandfiles.com/flash/2021/06/01/marvell-launches-first-pcie-5-ssd-controller/1616031).
- Targets: 14 GB/s seq read, 9 GB/s seq write, 2M/1M random IOPS; <6 µs latency with granular arbitration; NAND up to 1,600 MT/s; supports SLC/MLC/TLC/QLC from all vendors [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- Compute: 10 Arm cores (Cortex-R8/R7/M-class mix per press; Marvell brief shows Cortex-R8 firmware cores + Cortex-M7 security engine), hardware elastic SLA enforcer, 5th-gen NANDEdge LDPC ECC, hardware RAID engine, DMA controllers [secondary](https://www.tomshardware.com/news/marvell-announced-pcie-gen5-ssd-controllers).
- Virtualization: 16 physical functions / 32 virtual functions (SR-IOV); DRAM DDR4-3200 or LPDDR4X-4266 (72-bit with ECC) [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- Security: FIPS-compliant root of trust, AES-256, multi-key revocation, end-to-end data path protection [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- Flexibility: first flash controllers to enable SEF, ZNS, Open Channel, multiple usage models without hardware changes; NVMe 1.4b; dual-port 2x2 option; launch partners incl. Kioxia, SK hynix, AMD, Intel, Microsoft, Meta [secondary](https://www.guru3d.com/story/marvell-announces-bravera-sc5-ssd-controller-family-supports-pcie-gen-5-(17gbs)/?print).
- 16-channel in 20x20 mm enabled the industry's first 16CH E1.S ("ruler") SSDs [secondary](https://www.guru3d.com/story/marvell-announces-bravera-sc5-ssd-controller-family-supports-pcie-gen-5-(17gbs)/?print).

### 9.2 Phison enterprise controllers

- Phison PS5302-X2 ("X2"): 16-channel Gen5 controller behind Pascari X200 — 14,800 MB/s read demonstrated in first-party drive [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html).
- Phison E26 (consumer/client Gen5, PS5026-E26) is the client counterpart; enterprise uses the X2 line — do not conflate [secondary](https://www.tweaktown.com/news/98293/pascari-is-new-enterprise-brand-from-phison-the-biggest-behind-scenes-name-in-storage/index.html).
- Phison supplies controllers to Kioxia, Micron, WD and Samsung client lines historically — the "biggest behind-the-scenes name in storage" [secondary](https://www.tweaktown.com/news/98293/pascari-is-new-enterprise-brand-from-phison-the-biggest-behind-scenes-name-in-storage/index.html).

### 9.3 In-house controllers

- Samsung: proprietary controllers on PM1743/PM9A3/PM9D3a (Samsung is vertically integrated: NAND + controller + DRAM + firmware) [vendor-reported](https://www.businesswire.com/news/home/20211222005474/en/5120105/Samsung-Develops-High-Performance-PCIe-5.0-SSD-for-Enterprise-Servers).
- Kioxia: in-house developed controller on CD8P/CM7 [secondary](https://www.crn.in/news/kioxia-launches-new-pcie-5-0-ssds-for-enterprise-and-data-center-infrastructures/).
- Micron: in-house controller/NAND/DRAM/firmware on 9550 [vendor-reported](https://www.edn.com/data-center-ssds-achieve-blistering-speeds/).
- SK hynix: Atomos Prime ACNT08M on Solidigm D7-PS1010 (7 nm TSMC, ARM Cortex-R8, 16ch @ 1,600 MT/s) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146).

### 9.4 InnoGrit — gap

- InnoGrit (Tacoma IG5669 Gen4, etc.) is a relevant enterprise controller vendor but no 2026-corroborated data was retrieved in this pass — Section 14 [unverified].

---

