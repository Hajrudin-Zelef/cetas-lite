---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/5-other-physical-formats-in-the-2026-datacenter
title: "5. Other physical formats in the 2026 datacenter"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["AMD"]
dates: []
keywords: ["datacenter", "amd", "compute", "consumer", "packaging"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [96, 140]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 4aa9a7b07412647fc4cb46bcae78fa1f3fc0d55ad782b3e7bb0dba50f2f45753
---

# 5. Other physical formats in the 2026 datacenter

## 5. Other physical formats in the 2026 datacenter

### 5.1 2.5" SATA/SAS — the legacy workhorse

- 2.5" 7 mm (client) and 15 mm (enterprise) z-heights; the enterprise 15 mm bay is the shared mechanical home of U.2/U.3 NVMe, SAS and SATA [independent].
- SATA III tops out at 6 Gb/s; "SATA III in 2009, but there never was a SATA IV. There was just nibbling around the edges with incremental updates as momentum and emphasis shifted to PCI Express and NVMe" [secondary](https://www.networkworld.com/article/4121569/reports-of-satas-demise-are-overblown-but-the-technology-is-aging-fast.html).
- Consumer SATA SSDs are being squeezed out: a 4 TB SATA SSD at ~$300 vs a 4 TB PCIe Gen4 NVMe at ~$341 and 4 TB external USB at ~$280, with motherboard SATA ports shrinking — "4TB SATA SSDs are approaching the end of their practical relevance" [secondary](https://www.techradar.com/pro/large-external-ssds-are-now-cheaper-than-internal-ones-as-4tb-sata-ssd-face-extinction-due-to-negligible-price-difference).
- Enterprise SATA survives on $/TB for cold data: "Nearline SATA is flat at $25/TB across 12 TB, 16 TB and 20 TB drives" and 20–30 TB SATA HDDs from Seagate/WD "are apparently still in wide use in cloud data centers for things like cold storage" [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/)[secondary](https://www.networkworld.com/article/4121569/reports-of-satas-demise-are-overblown-but-the-technology-is-aging-fast.html).
- KIOXIA's "value SAS" SSD line is explicitly positioned as "a replacement SATA technology" for servers where SATA SSDs bottleneck the CPU [vendor-reported](https://www.businesswire.com/news/home/20210623005830/en/5000252/KIOXIA-America-Showcases-PCIe-4.0-24G-SAS-SSDs-at-HPE-Discover).

### 5.2 3.5" LFF — bulk HDD territory

- 3.5" large-form-factor bays remain the home of high-capacity HDDs (12–24+ bays in 2U, e.g. Dell R570/R770 3.5" chassis options, HPE DL380 Gen12 12× LFF configs) [secondary](https://www.dell.com/en-au/shop/servers-storage-and-networking/poweredge-r570-rack-server/spd/poweredge-r570/promo_r570_1?view=configurations)[secondary](https://farnamco.net/wp-content/uploads/2026/01/HPE-ProLiant-Compute-Family-Guide.pdf).
- Hybrid 2026 pattern: 3.5" HDD bulk tier + rear E3.S NVMe cache tier (HPE DL340e Gen12: "up to 24 LFF 3.5-inch drives up front … and up to 10 EDSFF NVMe drives in the rear for a hybrid capacity-plus-cache config") [secondary](https://www.youtube.com/shorts/ACX3Z1LOihU).

### 5.3 AIC / HHHL add-in cards

- The PCIe add-in card (half-height half-length, CEM connector) predates U.2 for NVMe; EDSFF's rationale explicitly notes AIC "requires PCIe AIC slots for other devices and has limited hot-plug capabilities while also taking up a great deal of space" [secondary](https://eetimes.com/nvme-drives-ready-to-embrace-own-form-factors/).
- 2026 AIC role: PCIe switch adapters (e.g. HighPoint Rocket 1624A: Gen5 x16 host, dual MCIO, up to 16 NVMe SSDs), retimer/redriver cards, and specialty accelerators — not primary SSD packaging for new servers [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw).

### 5.4 OCP slots and DC-MHS — the modular chassis

- OCP NIC 3.0 / DC-MHS PMM slots share the EDSFF connector lineage: an E3 device with 4C+ connector in an SFF-TA-1034 slot "will be treated like an OCP NIC 3.0 device in a PMM slot" [official](https://members.snia.org/document/dl/53687).
- DC-MHS (Datacenter Modular Hardware System) splits the server into Host Processor Module (HPM, compute) and Datacenter Secure Control Module (DC-SCM, management/security), so CPU and BMC/security lifecycles evolve independently [secondary](https://medium.com/codex/from-traditional-servers-to-modular-servers-understanding-ocp-dc-mhs-hpm-and-dc-scm-2a55471033b8).
- MSI's DC-MHS AMD platforms (OCP APAC 2025) ship 12× E3.S PCIe 5.0 NVMe bays per node with DC-SCM2 management modules, for 21" ORv3 and 19" EIA racks [vendor-reported](https://www.msi.com/news/detail/MSI-Showcases-DC-MHS-and-MGX-Server-Platforms--for-Cloud-Scale-and-AI-Infrastructure-at-OCP-APAC-2025-146777).
- HPE Gen12 hybrid front cages accept "two front OCP NICs as optional" in the same cage structure as SFF/E3.S drives — storage and network sharing front-bay real estate [secondary](https://buy.hpe.com/my/en/compute/hpe-proliant-compute-gen12-servers/c/c001030).

## 6. Form-factor comparison matrix (2026 enterprise)

| Form factor | Typical interface | Lanes | Power envelope | Hot-swap | Density note | 2026 sweet spot |
|---|---|---|---|---|---|---|
| M.2 2280/22110 | NVMe | x4 | ~8–10 W | No (board/carrier) | Low | Boot/OS (BOSS, NS204i-u) |
| 2.5" SATA | SATA 6G | — | ~5–9 W | Yes | Medium | Cold tier, legacy fleets |
| 2.5" SAS (24G) | SAS-4 | x1/x2/wide | ~9–15 W | Yes | Medium | Mixed HDD/SSD, dual-port HA |
| U.2 (SFF-8639) | NVMe | x4 | 25 W | Yes | Medium | Installed NVMe fleets |
| U.3 (SFF-TA-1001) | NVMe/SAS/SATA | x1/x2/x4 | 25 W | Yes | Medium | Universal/tri-mode bays |
| E1.S | NVMe | x4 | 12–25 W | Yes | High (1U) | 1U compute nodes |
| E1.L | NVMe | x4 | 25 W | Yes | Very high (capacity) | 200 TB+ ruler drives |
| E3.S | NVMe | x4/x8 | 25–40 W | Yes | Very high | Density NVMe, AI storage |
| E3.L | NVMe | x4/x8/x16 | 40–70 W | Yes | High (capacity) | All-flash arrays |
| 3.5" LFF | SAS/SATA | — | ~10–15 W | Yes | Low (per TB high) | Bulk HDD cold storage |
| AIC HHHL | PCIe | x8/x16 | 75–300 W | Limited | Low | Switch/retimer cards |

*Power figures: EDSFF from SNIA guidance [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html); others are typical-class envelopes [independent]. Always validate against the specific drive's datasheet — "should not be treated as the SSD's normal power consumption."*

