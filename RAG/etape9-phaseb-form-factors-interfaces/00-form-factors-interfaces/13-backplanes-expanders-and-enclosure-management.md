---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/13-backplanes-expanders-and-enclosure-management
title: "13. Backplanes, expanders, and enclosure management"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Broadcom", "Intel"]
dates: ["2024-03"]
keywords: ["compute", "consumer", "intel", "nand"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [291, 314]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 71c81ba378eb092932c28b20ec46b7b1ce9cb5bd2356a18772d2dc9351465dff
---

# 13. Backplanes, expanders, and enclosure management

- Consumer: no SATA IV ever; SATA SSDs losing price/performance vs NVMe and external USB SSDs; motherboard SATA ports shrinking, "especially in smaller form factor systems and newer platforms" [secondary](https://www.networkworld.com/article/4121569/reports-of-satas-demise-are-overblown-but-the-technology-is-aging-fast.html)[secondary](https://www.techradar.com/pro/large-external-ssds-are-now-cheaper-than-internal-ones-as-4tb-sata-ssd-face-extinction-due-to-negligible-price-difference).
- Enterprise: 20–30 TB SATA HDDs still ship in volume for cold storage at ~$25/TB flat across 12/16/20 TB [secondary](https://www.networkworld.com/article/4121569/reports-of-satas-demise-are-overblown-but-the-technology-is-aging-fast.html)[secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- Enterprise SATA SSDs are being replaced by value-SAS (KIOXIA RM5 "Welcome to Life after SATA") because "SATA's 6Gb/s performance ceiling" bottlenecks flash [vendor-reported](https://www.businesswire.com/news/home/20210623005830/en/5000252/KIOXIA-America-Showcases-PCIe-4.0-24G-SAS-SSDs-at-HPE-Discover)[secondary](https://www.eetimes.com/sas-customers-expect-innovation-backwards-compatibility/).
- 2026 price reality (Aug 2026): new enterprise NVMe U.2 at $300–$1,172/TB vs nearline SATA $25/TB vs used 3.84 TB enterprise SATA/SAS SSDs $78–$127/TB — "the used channel is priced off a different clock than the NAND market"; no NAND relief expected "before the second half of 2027" per TrendForce [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).

## 13. Backplanes, expanders, and enclosure management

### 13.1 Tri-mode controllers and universal backplanes

- A universal (U.3) backplane needs three things: tri-mode controller, SFF-8639 connectors with SFF-TA-1001 wiring, and the Universal Backplane Management framework [secondary](https://www.storagereview.com/news/evolving-storage-with-sff-ta-1001-u-3-universal-drive-bays?amp).
- Broadcom's tri-mode stack pairs the controller with "Universal Backplane Management (UBM) ready (SFF-TA-1005)" backplanes; 9600-series works on both SlimSAS 24G SFF-8654 and existing SFF-8639 (U.2) backplanes [vendor-reported](https://5.imimg.com/data5/SELLER/Doc/2025/11/559109987/XE/CD/PL/228806307/broadcom-24gb-sas-nvme-adapter.pdf).
- HPE Gen12's tri-mode backplane "supports SATA, SAS4 plus NVMe SSD devices" and can direct-attach NVMe to motherboard M-XIO connectors without a RAID card [secondary](https://www.itpro.com/infrastructure/servers-and-storage/hpe-proliant-compute-dl325-gen12-review-a-deceptively-small-and-powerful-1p-rack-server-with-a-huge-core-count).
- SAS expanders still matter for HDD-heavy chassis (single controller fanning to 240+ devices); NVMe has no expander equivalent — scaling is via PCIe switches or more CPU lanes [independent].

### 13.2 UBM (SFF-TA-1005) — one management language

- UBM unifies "SGPIO (SFF-8485), 2-Wire SES and 2-Wire PCA9555 usages" plus Intel VPP into a single backplane management spec driven by the SNIA SFF TA TWG [secondary](https://www.storagenewsletter.com/2018/06/18/broadcom-co-sponsor-for-sff-ta-1005-universal-backplane-management-spec/).
- UBM benefits: "holistic support of SAS, SATA and PCIe (NVMe) backplanes; supports storage topologies such as controller-, expander- and switch-attach; high-speed lane mapping descriptions; device slot and device presence information; management of LED states and slot power control; management of PCIe Reset and Reference Clock mechanisms" [secondary](https://www.storagenewsletter.com/2018/06/18/broadcom-co-sponsor-for-sff-ta-1005-universal-backplane-management-spec/).
- Reference silicon: Microchip EEC1005-UB2 — Cortex-M4 UBM controller, SFF-TA-1005 v1.4 compliant, tri-mode NVMe/SAS/SATA on one device, up to 16 drives per chip, multiple backplanes per chassis, FRU-programmable configs, "customizable LED patterns for diagnostics", secure boot + firmware update via AES-256/ECDSA-P-256/SHA-256, NVMe hot-plug and power-disable support; priced at $7.33 in 10k quantities (March 2024) [vendor-reported](https://www.storagenewsletter.com/2024/03/14/microchip-eec1005-ub2-universal-backplane-management-controller/)[vendor-reported](https://www.eeworldonline.com/ubm-controllers-targeted-at-data-canter-and-storage-management-applications/).
- LED behavior follows SFF-8489 IBPI (International Blinking Pattern Interpretation) by default, with custom blink patterns via FRU [vendor-reported](https://www.microsemi.com/product-directory/upcoming-technology/5558-universal-backplane-management-ubm).
- UBM also covers the EDSFF drive types (U.2, U.3 and EDSFF) — one controller family across bay types [vendor-reported](https://www.storagenewsletter.com/2024/03/14/microchip-eec1005-ub2-universal-backplane-management-controller/).

### 13.3 The legacy alphabet: SGPIO, SES, VPP

