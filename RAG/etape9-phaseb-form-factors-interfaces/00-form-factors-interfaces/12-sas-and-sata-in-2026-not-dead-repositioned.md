---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/12-sas-and-sata-in-2026-not-dead-repositioned
title: "12. SAS and SATA in 2026 — not dead, repositioned"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Broadcom", "Intel", "Samsung"]
dates: ["2024-03"]
keywords: ["compute", "consumer", "intel", "nand", "serdes"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [277, 326]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 7a52a9901b311f95fb942143354062d75dbc38a5ca251d043a483ef7dcc99cef
---

# 12. SAS and SATA in 2026 — not dead, repositioned

## 12. SAS and SATA in 2026 — not dead, repositioned

### 12.1 SAS-4 (24G) is the current SAS generation

- SAS roadmap: "evolving from 3 Gb/s to 24G with 48G upcoming" [secondary](https://www.eetimes.com/sas-customers-expect-innovation-backwards-compatibility/).
- Tri-mode SerDes (Broadcom 9600 series): 22.5/12/6 Gb/s SAS + 6 Gb/s SATA + PCIe up to 16 GT/s per lane on one PHY; adapters auto-negotiate speed and protocol [vendor-reported](https://WWW.scan.co.uk/products/16-port-broadcom-megaraid-9660-16i-raid-controller-4gb-cache-2x-sas-sff-8654-sata-iii-nvme-raid-0-1).
- 9600-series performance envelope: 6M–6.4M IOPS (4K RR JBOD), 4.5M–6.4M RAID 0/1, 900K–1.1M RAID 5 (4K RW); RAID 0/1/5/6/10/50/60 + JBOD; hardware secure boot + SPDM attestation; UBM-ready (SFF-TA-1005) [vendor-reported](https://5.imimg.com/data5/SELLER/Doc/2025/11/559109987/XE/CD/PL/228806307/broadcom-24gb-sas-nvme-adapter.pdf).
- Scale per controller: "up to 240 SAS/SATA devices or 32 NVMe devices per controller" (9660-16i); HBAs reach 1,024 SAS/SATA or 32 NVMe [vendor-reported](https://www.thetedstore.com/products/ffr-9b16-125-028)[vendor-reported](https://www.scan.co.uk/products/8-port-broadcom-9500-8i-tri-mode-storage-adapter-pcie-gen-40-1-x8-sff-8654-sas-sata-nvme-12-gb-s).
- KIOXIA PM6 24G SAS SSDs on HPE ProLiant deliver "significantly improved performance over 12Gb/s SAS"; HPE+KIOXIA delivered "industry's first 24G SAS SSD for HPE ProLiant servers" [vendor-reported](https://www.businesswire.com/news/home/20210623005830/en/5000252/KIOXIA-America-Showcases-PCIe-4.0-24G-SAS-SSDs-at-HPE-Discover).
- SAS SSD market 2025–2026: competition "centered on SAS-4 (24G SAS) readiness, with Samsung, Seagate, and Western Digital each announcing SAS-4-compliant product roadmaps"; the ">30 TB in standard 2.5-inch" milestone is the 2027–2028 battleground for shelf-replacement TCO [secondary](https://dataintelo.com/report/serial-attached-storage-sas-solid-state-drive-ssd-market).
- Why SAS persists per KIOXIA: "flexibility and manageability… a stable interface, and 24G provides that doubling of bandwidth"; OEMs keep SAS for hybrid SSD+HDD systems — "that ability to stuff SSDs and HDDs into the same system" [secondary](https://www.eetimes.com/sas-customers-expect-innovation-backwards-compatibility/).

### 12.2 SATA: consumer exit, enterprise cold tier

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

- SGPIO (SFF-8485): serial GPIO bit-banging for drive LEDs/activity — the pre-UBM standard, still selectable on UBM chips via pin strapping ("seamless switching between SGPIO and UBM protocols") [vendor-reported](https://www.electronicsforu.com/electronics-projects/reference-design-for-universal-storage-backplane-management).
- SES (SCSI Enclosure Services, 2-wire): the SAS-expander enclosure protocol for slot/LED/thermal management [independent].
- Intel VPP (Virtual Pin Port): per-lane LED/control virtualization for NVMe direct-attach [independent].
- SES/LED/thermal remain the BMC-visible surface: UBM devices expose "I2C communication to Baseboard Management Controller" with up to 8 I2C ports, plus an optional sideband I2C for BMC direct access [vendor-reported](https://www.microsemi.com/product-directory/upcoming-technology/5558-universal-backplane-management-ubm).

### 13.4 Hot-swap, power-disable, surprise removal

- Enterprise bays require: per-slot power control, PCIe reset + reference clock management (UBM's job), and OS/driver surprise-removal handling [secondary](https://www.storagenewsletter.com/2018/06/18/broadcom-co-sponsor-for-sff-ta-1005-universal-backplane-management-spec/).
- NVMe power-disable is a UBM-managed feature on modern backplanes (Microchip EEC1005-UB2 lists it explicitly) [vendor-reported](https://www.microsemi.com/product-directory/upcoming-technology/5558-universal-backplane-management-ubm).
- HighPoint's Gen5 switch adapters advertise "hot-plug and hot-swap NVMe support allows drive replacement while the system remains online" with self-diagnostic LEDs and FRU support [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw).
- Dell 17G E3.S chassis options distinguish "Cold Aisle Supported" vs "Hot Aisle" serviceability — front-serviceable NVMe is a chassis engineering feature, not just a connector [secondary](https://www.dell.com/en-au/shop/servers-storage-and-networking/poweredge-r570-rack-server/spd/poweredge-r570/promo_r570_1?view=configurations).

