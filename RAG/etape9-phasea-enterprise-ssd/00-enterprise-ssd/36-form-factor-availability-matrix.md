---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/36-form-factor-availability-matrix
title: "36. Form-factor availability matrix"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung", "United States"]
dates: ["2026-09-22"]
keywords: ["benchmark", "datacenter", "ipo", "memory", "nand", "optics", "research"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [709, 759]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: d6d11cfe9b1d7fc2f4e4211a216a752794f76650c8a6d6becc4ca17c616c6c0e
---

# 36. Form-factor availability matrix

## 36. Form-factor availability matrix

| Drive | U.2 15 mm | U.3 | E1.S | E3.S | M.2 |
|---|---|---|---|---|---|
| Samsung PM9A3 | yes | — | yes | — | yes (2280/22110) [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf) |
| Samsung PM1743 | yes | — | — | yes | — [vendor-reported](https://www.businesswire.com/news/home/20211222005474/en/5120105/Samsung-Develops-High-Performance-PCIe-5.0-SSD-for-Enterprise-Servers) |
| Samsung PM9D3a | yes | — | yes (9.5 mm) | yes | — [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/) |
| Kioxia CM7 | yes | yes (SFF-TA-1001) | — | yes (7.5 mm) | — [secondary](https://www.shidirect.com/Product/46112649/KIOXIA-CM7-R-Series-SSD) |
| Kioxia CD8P | yes (U.3) | — | — | — | — [secondary](https://www.crn.in/news/kioxia-launches-new-pcie-5-0-ssds-for-enterprise-and-data-center-infrastructures/) |
| Micron 9550 | yes | — | yes (15 mm) | yes (7.5 mm) | — [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf) |
| Micron 7450 | yes | yes | yes | — | yes (2280/22110) [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf) |
| Solidigm D7-PS1010 | yes | — | (E.1 liquid-cooled variant) | yes | — [secondary](https://www.tomshardware.com/pc-components/ssds/solidigm-touts-industrys-first-liquid-cooled-enterprise-ssd-d7-ps1010-is-an-e-1-pcie-5-0-drive-with-a-wrap-around-cold-plate) |
| Solidigm D5-P5336 | yes | — | — | — | — [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm) |
| SK hynix PEB110 | — | — | yes | — | — [vendor-reported](https://en.prnasia.com/releases/global/sk-hynix-develops-peb110-e1-s-for-data-centers-460322.shtml) |
| WD SN861 | yes | — | yes (15 mm) | — | — [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf) |
| Phison X200E | yes | — | — | yes | — [vendor-reported](https://www.phisonenterprise.com/wp-content/uploads/2025/01/PascariProductBrochure_X200_010925.pdf) |

---

## 37. Related RAG cross-references

- Step 7 Phase E (Ceph, Rook, ZFS, backup): pair drive selection here with the software stack there — BlueStore DB/WAL placement, ZFS special vdevs/SLOG/L2ARC guidance in Section 17 builds on it.
- Step 6 networking (NVMe-oF): Gen5 drives saturating 14 GB/s need 200 GbE+ fabrics per target to avoid network bottlenecks — see Phase C/D optics and architecture files.
- Step 5 servers (HPE, Dell, Supermicro): backplane form-factor support (U.2/U.3/E3.S tri-mode) and Gen5 retimer requirements constrain which drives fit which chassis.
- Step 8 (planned): power/cooling design numbers in Section 35 feed datacenter power budgeting.

---

## 15b. Glossary additions

- **4D NAND**: SK hynix marketing term for PUC (Periphery Under Cell) — logic beneath the array [secondary](https://www.storagenewsletter.com/2024/10/09/sk-hynix-unveils-peb110-e1-s-gen5-pcie-up-to-8tb-ssd-for-data-centers/).
- **BiCS**: Kioxia/SanDisk 3D NAND brand (BiCS5 = 112L, BiCS8 = 218L CBA, BiCS10 = 332L CBA) [secondary](https://www.techtimes.com/articles/323160/20260805/kioxia-squeezes-38-more-data-per-chip-samsung-68-fewer-layers.htm).
- **V-NAND**: Samsung 3D NAND brand (V6 = 128L, V9 = ~290L, V10 = 400+L) [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- **Toggle DDR / Toggle DDR6.0**: NAND interface standard; BiCS10 reaches 4.8 GB/s [secondary](https://www.techtimes.com/articles/323160/20260805/kioxia-squeezes-38-more-data-per-chip-samsung-68-fewer-layers.htm).
- **ONFI**: Open NAND Flash Interface — alternative NAND interface standard [independent].
- **LDPC**: Low-Density Parity-Check — modern NAND ECC (Marvell NANDEdge 5th gen) [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- **FTL**: Flash Translation Layer — firmware mapping logical to physical NAND pages [independent].
- **GC**: garbage collection — background reclaim of invalid pages; source of WAF [independent].
- **TRIM / Deallocate**: host hints that pages are invalid — lowers WAF [independent].
- **QD**: queue depth — outstanding I/Os; high QD needed to hit headline IOPS [independent].
- **Steady state**: preconditioned performance per SNIA PTS — the only comparable benchmark state [independent].
- **PLP capacitor aging**: electrolytic/supercap PLP elements degrade — `volatile memory backup failed` warning (SMART bit 4) flags it [independent].
- **PSID**: physical security ID printed on SED labels — required for PSID revert of locked drives [independent].
- **TAA**: Trade Agreements Act — US federal procurement compliance option on Micron 9550 SKUs [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- **AFR**: annualized failure rate — population failure metric (WD SN861: 0.35% projected) [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf).
- **EDSFF**: Enterprise and Datacenter SSD Form Factor family (E1.S/E3.S) [secondary](https://www.tweaktown.com/news/95016/kioxia-ssds-achieve-pcie-5-0-compliance-up-to-30tb-capacity-and-14-000-mb/index.html).
- **NVMe-oF**: NVMe over Fabrics — remote NVMe via RDMA/TCP; Gen5 drives need 200 GbE+ per target [independent].

---

*End of Step 9 Phase A — Enterprise SSD hardware. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*
