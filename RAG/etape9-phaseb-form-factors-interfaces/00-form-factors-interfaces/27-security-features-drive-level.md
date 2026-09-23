---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/27-security-features-drive-level
title: "27. Security features (drive-level)"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Broadcom", "Samsung"]
dates: ["2026-08", "2026-09-22"]
keywords: ["compute", "datacenter", "ethernet", "latency", "memory", "nand", "research", "serdes"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [669, 751]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 9194f499508832537492d072abaa02e283fdd09a4753e97df7d5ca4178daaf48
---

# 27. Security features (drive-level)

## 27. Security features (drive-level)

- **SED (Self-Encrypting Drive)**: TCG Opal/Enterprise hardware encryption; enterprise SSDs typically ship SED-capable SKUs with instant secure erase [independent].
- **NVMe 2.4 Post-Quantum Cryptography**: new in the August 2026 spec family — "foundational enablers to remain secure in the post-quantum era" [official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage).
- **Secure boot & attestation**: Broadcom 9600-series ships hardware secure boot + SPDM attestation; Microchip UBM controller has AES-256/ECDSA-P-256/SHA-256 secure boot and firmware update [vendor-reported](https://5.imimg.com/data5/SELLER/Doc/2025/11/559109987/XE/CD/PL/228806307/broadcom-24gb-sas-nvme-adapter.pdf)[vendor-reported](https://www.storagenewsletter.com/2024/03/14/microchip-eec1005-ub2-universal-backplane-management-controller/).
- **Sanitize**: NVMe Format/Sanitize commands for crypto-erase and block-erase at decommission [independent].
- [unverified] Per-SKU SED/FIPS availability for specific 2026 EDSFF drives — confirm against the drive vendor's datasheet before procurement.

## 28. Firmware and lifecycle management

- **NVMe-MI firmware update**: the standard in-band/out-of-band path for drive firmware; drives like the KIOXIA CM9 expose NVMe-MI 1.2c [vendor-reported](http://www.techpowerup.com/336800/kioxia-announces-first-enterprise-nvme-ssd-with-8th-gen-bics-flash-technology).
- **Dual-image firmware**: enterprise drives commonly carry two firmware slots (active + staged) so updates don't brick the device [independent].
- **OEM firmware gates**: Dell/HPE/Supermicro qualify specific drive firmware revisions per platform — mixing retail firmware into OEM backplanes can break UBM/LED/telemetry behavior; always check the vendor's support matrix [independent].
- **BMC integration**: UBM gives the BMC I2C visibility into slot presence, power, PCIe reset and LED state — the hardware foundation of iDRAC/iLO drive health reporting [vendor-reported](https://www.microsemi.com/product-directory/upcoming-technology/5558-universal-backplane-management-ubm).

## 29. Thermal and airflow notes

- EDSFF power classes directly size the thermal design: E1.S thin 12 W vs thick 25 W, E3.S 25–40 W, E3.L 40–70 W sustained recommendations [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).
- Ruler drives (E1.L) expose large surface area for airflow in 1U; the SNIA guidance notes airflow directionality in the mechanical spec [independent].
- Dell 17G chassis options distinguish cold-aisle vs hot-aisle serviceability for E3.S — front-bay airflow is engineered per chassis variant [secondary](https://www.dell.com/en-au/shop/servers-storage-and-networking/poweredge-r570-rack-server/spd/poweredge-r570/promo_r570_1?view=configurations).
- NVMe 2.3 Power Limit Config lets hosts cap device power — "particularly important for older systems with limited power capabilities" — a thermal tool as well as an electrical one [secondary](https://markets.financialcontent.com/sweetwaterreporter/article/bizwire-2025-8-5-nvm-express-publishes-set-of-nvme-specifications-enabling-new-capabilities-for-ai-cloud-enterprise-and-client-storage).
- Gen5 SSD controllers run hot under sustained 14 GB/s sequential loads — datasheet operating-temperature ranges assume the server's airflow; homelab/whitebox users adding Gen5 U.2 via adapters should verify cooling [independent].

## 30. One-page cheat sheet (2026 enterprise storage hardware)

| Question | Answer |
|---|---|
| Default server SSD interface | NVMe over PCIe Gen5 x4 |
| Default server SSD form factor | E3.S for density; U.2/U.3 for installed fleets |
| Boot drives | 2× M.2 NVMe (Dell BOSS-N1, HPE NS204i-u V2), not data bays |
| Universal bay | U.3 (SFF-TA-1001): SATA + SAS + NVMe x1/x2/x4 |
| 1U dense compute | E1.S (12–25 W) |
| Capacity ruler | E1.L (245 TB class, 25 W) |
| All-flash array | E3.L (40–70 W) |
| Tri-mode controller | Broadcom 95xx/96xx (SAS 24G + SATA + NVMe) |
| Backplane management | UBM (SFF-TA-1005); legacy SGPIO/SES/VPP |
| Internal Gen5 cable | MCIO 8i (SFF-TA-1016), ~1 m copper |
| Legacy internal cable | SFF-8643 Mini-SAS HD; SlimSAS SFF-8654 for 24G |
| Shared flash over Ethernet | NVMe/TCP (no special fabric); RDMA/RoCE for lowest latency |
| Legacy SAN path | FC-NVMe on Gen6 FC HBAs |
| Sequential-write optimization | ZNS (zones) or FDP (TP4146, host hints) |
| HA without SAS | Dual-port NVMe + ANA multipath |
| Spec to design against | NVMe 2.4 family (Aug 2026); PCIe 7.0 spec ratified, products ~2028 |
| Cold data $/TB | 3.5" SATA HDD ~$25/TB |
| 2026 NAND market | Tight; no relief expected before H2 2027 per TrendForce [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/) |
| Modular chassis | OCP DC-MHS (HPM + DC-SCM); E3.S 4C+ slots via SFF-TA-1034 |
| CXL on servers | CXL 2.0 memory expansion (Dell 17G); shares PCIe 5.0 SerDes with NVMe |

*End of file content. Research cutoff 2026-09-22. Single writer on this path; no other workspace files modified.*

## 31. SKUs referenced in this file (quick table)

### 31.1 Flagship NVMe SSDs (2026)

| SSD | Interface | Form factors | Key specs (as announced) |
|---|---|---|---|
| KIOXIA CM9-R | PCIe 5.0, NVMe 2.0, NVMe-MI 1.2c | 2.5" / E3.S | BiCS FLASH Gen 8 TLC; up to 14.8 GB/s read, 11 GB/s write (preliminary); dual-port; 15.36 TB class; OCP Datacenter NVMe SSD 2.5 compliant [vendor-reported](http://www.techpowerup.com/336800/kioxia-announces-first-enterprise-nvme-ssd-with-8th-gen-bics-flash-technology) |
| Samsung PM1731a | NVMe (ZNS) | 2.5" | ZNS SSD; WAF close to 1; up to 4× endurance vs conventional NVMe [vendor-reported](https://www.enterprisestorageforum.com/news/samsung-zns-ssd/) |
| KIOXIA PM6 | 24G SAS | 2.5" | SAS-4 SSD; industry's first 24G SAS SSD for HPE ProLiant [vendor-reported](https://www.businesswire.com/news/home/20210623005830/en/5000252/KIOXIA-America-Showcases-PCIe-4.0-24G-SAS-SSDs-at-HPE-Discover) |
| KIOXIA RM5 (value SAS) | 12G SAS | 2.5" | Positioned as SATA replacement for servers [vendor-reported](https://www.businesswire.com/news/home/20210623005830/en/5000252/KIOXIA-America-Showcases-PCIe-4.0-24G-SAS-SSDs-at-HPE-Discover) |

### 31.2 Broadcom tri-mode controllers referenced

| Controller | Host | Ports | Perf (4K) |
|---|---|---|---|
| 9500-8i | PCIe Gen4 x8 | 8 internal (SFF-8654) | 12G SAS class [vendor-reported](https://www.scan.co.uk/products/8-port-broadcom-9500-8i-tri-mode-storage-adapter-pcie-gen-40-1-x8-sff-8654-sas-sata-nvme-12-gb-s) |
| 9500-16i / 9500-8e | PCIe Gen4 x16 | 16 internal / 16 external | Tri-mode 12G [secondary](https://www.thetedstore.com/products/ffr-9b16-125-028) |
| 9600-16i / 24i / 32i | PCIe Gen4 | 16/24/32 (SFF-8654) | 24G tri-mode, up to 240 SAS/SATA or 32 NVMe per controller [vendor-reported](https://www.thetedstore.com/products/ffr-9b16-125-028) |
| MegaRAID 9660-16i | PCIe Gen4 | 16 (2× SFF-8654) | 6M IOPS JBOD; 4.5M RAID 0/1; 900K RAID 5 [vendor-reported](https://5.imimg.com/data5/SELLER/Doc/2025/11/559109987/XE/CD/PL/228806307/broadcom-24gb-sas-nvme-adapter.pdf) |
| MegaRAID 9670W-16i | PCIe Gen4 | 16 | 6.4M IOPS JBOD; 6.4M RAID 0/1; 1.1M RAID 5 [vendor-reported](https://5.imimg.com/data5/SELLER/Doc/2025/11/559109987/XE/CD/PL/228806307/broadcom-24gb-sas-nvme-adapter.pdf) |

- Cable example: Broadcom 05-60002-00 — 1 m SlimSAS x8 SFF-8654 to 2× x4 SFF-8643 Mini-SAS HD (bridges 95xx adapters to older backplanes) [secondary](https://www.scan.co.uk/products/1m-broadcom-05-60002-00-slimsas-x8-sff-8654-to-two-x4-sff-8643-mini-sas-hd-nvme-connection).
- PCIe switch example: HighPoint Rocket 1624A — Broadcom PEX89048, Gen5 x16 host, 2× MCIO 8i, up to 16 NVMe SSDs [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw).
- UBM controller example: Microchip EEC1005-UB2 — SFF-TA-1005 v1.4, Cortex-M4, up to 16 drives, $7.33 @10k (Mar 2024) [vendor-reported](https://www.storagenewsletter.com/2024/03/14/microchip-eec1005-ub2-universal-backplane-management-controller/).

## 32. Document revision note

- **v1.0 — 2026-09-22**: Initial single-writer completion. 31 sections, all research current to 2026-09-22.
- Waves: header + EDSFF (§1–4) → M.2/U.2/U.3 (§5 header region) → other form factors, comparison matrix, NVMe spec family (§5–7) → command sets, NVMe-oF, advanced features (§8–10) → PCIe generations, SAS/SATA (§11–12) → backplanes, cabling (§13–14) → server integration, decision guide, gaps, glossary, sources (§15–19) → timeline, interop, lane budget, retimers, CXL, software (§20–25) → capacity, security, firmware, thermal (§26–29) → cheat sheet (§30) → SKU tables (§31).
- No other workspace files were modified during this task.
- Suggested maintenance: re-verify NVMe 2.4 drive availability, PCIe 6.0 SSD launches, and SAS-48G ratification when revisiting after Q1 2027.

- RAG ingestion: chunk on `##` section boundaries; each section is self-contained with its own citations.
