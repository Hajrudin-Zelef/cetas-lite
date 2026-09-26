---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/14-storage-cabling-and-connectors
title: "14. Storage cabling and connectors"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["AMD", "Broadcom", "Intel"]
dates: ["2026-04-01"]
keywords: ["amd", "compute", "gpu", "intel", "memory", "optics"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [327, 372]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 1c7fdf80e8b38d8f2fee6a3cbff6d38bc20f5576b044a92dc030179e06fb9624
---

# 14. Storage cabling and connectors

## 14. Storage cabling and connectors

### 14.1 Connector inventory

| Connector | Spec | Use | Notes |
|---|---|---|---|
| SFF-8639 | U.2 | 2.5" NVMe/SAS/SATA bay | 68-pin; U.3 reuses it with SFF-TA-1001 wiring |
| SFF-8643 | Mini-SAS HD | Internal x4 | Previous-gen internal; breakout cables to U.2 still sold |
| SFF-8654 | SlimSAS | Internal x4/x8 | 24G tri-mode standard on Broadcom 95xx/96xx; 8i/4i variants |
| SFF-8674 | — | External | 24G external on 9600-series |
| SFF-8087 | Mini-SAS | Internal x4 | Legacy 6G-era internal |
| MCIO 8i | SFF-TA-1016 | Internal x8 | 74-pin, Gen5/Gen6, "Mini Cool Edge IO" |
| OCuLink | SFF-8611/8612 | Internal/external x4 | Legacy PCIe cabling, largely superseded by MCIO |

- Broadcom 9600-series: "compatible with SlimSAS 24G SFF-8654 internal or SFF-8674 external connectors" [vendor-reported](https://5.imimg.com/data5/SELLER/Doc/2025/11/559109987/XE/CD/PL/228806307/broadcom-24gb-sas-nvme-adapter.pdf).
- Breakout reality: 1 m Broadcom cable "SlimSAS x8 SFF-8654 to two x4 SFF-8643 Mini-SAS HD" (05-60002-00) connects 95xx tri-mode adapters to SAS/SATA/PCIe(NVMe) devices — the cable that lets one adapter generation serve old backplanes [secondary](https://www.scan.co.uk/products/1m-broadcom-05-60002-00-slimsas-x8-sff-8654-to-two-x4-sff-8643-mini-sas-hd-nvme-connection).

### 14.2 MCIO (SFF-TA-1016) — the Gen5/Gen6 cable

- MCIO ("Mini Cool Edge IO"), SFF-TA-1016: 74-pin, x8, 0.60 mm pitch, vertical/right-angle; rated "up to 64 Gb/s PAM4, PCIe Gen 4, PCIe Gen 5, PCIe Gen 6, over 1.0 meter" [vendor-reported](https://www.newegg.ca/p/238-00D8-00027).
- One MCIO 8i carries a full PCIe Gen5 x8 (32 GT/s, "up to 256 Gbps total bandwidth"); supports bifurcation to dual x4 NVMe [vendor-reported](https://datasheet.itscope.com/2.1/t/Hc4xNMicnBB3UhUhPI7m5PtNPzD0l5IEpv3OCf9-3DLS6QFmBJ1X4enr1g6nayNaxVVAqsib773xL6nayNaxVVAqsib773xL6rAddIgK0ysFp4G5BB-pHMRZdELYWY76YkdU7i2THKwwJYzCHH9wzfwyQ9eeDukqo7AP5WTJh8fSFx7-2rfftuzR-I0fP8Q_HPlvGNMEFdCTcgRsloU).
- Typical 2026 use: motherboard/MCIO → 2× U.2 (SFF-8639) breakout with 15-pin SATA power for the drives — the retail cable exists precisely because servers mix MCIO hosts with U.2 bays [secondary](https://www.newegg.com//p/0ZF-0180-00161?Item=9SIB4D1KKB3186&).
- HighPoint's ecosystem: "certified cabling options for direct NVMe and backplane connectivity, including SFF-8639, SFF-TA-1002, SFF-8654, SFF-8643, and MCIO" — one vendor spanning four connector generations [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw).

### 14.3 Signal integrity at Gen5/Gen6

- At 32 GT/s (Gen5) and 64 GT/s (Gen6), trace/cable reach collapses without help: redrivers (analog boost) and retimers (full clock-data recovery) are mandatory on longer paths [independent].
- The PCIe 7.0 Optical Aware Retimer ECN standardizes retimer-based operation "over optical fiber" — copper reach limits are pushing even short-reach storage links toward optics in AI-scale chassis [secondary](https://convergedigest.com/pci-sig-finalizes-pcie-7-0-specification-at-128-0-gt-s/).
- Practical consequence for storage architects: Gen5 NVMe backplanes are co-designed with the motherboard (trace budgets, connector placement); aftermarket "just add cables" expansion tops out at Gen4 speeds unless the cable is explicitly Gen5-rated [independent].

## 15. Server platform integration (2026)

### 15.1 Dell PowerEdge 17G

- 17G is the Intel Xeon 6 / AMD EPYC 9005 generation with PCIe Gen5 throughout and EDSFF E3 as the density play: "The Dell PowerEdge 16G line enabled larger local storage capacity by incorporating the EDSFF-E3 disk drive form factor… EDSFF-E3 support increased local capacity by 60 percent in 1U servers and 33 percent in 2U servers" [secondary](https://www.wwt.com/article/unlocking-the-future-of-the-data-center-the-evolution-of-dell-poweredge-servers).
- Per-model E3.S NVMe Gen5 capacity (Dell/WWT): R470 up to 16 E3.S; R570 up to 32 E3.S; R670 up to 16 E3.S; R770 up to 44 E3.S — "including E3.S NVMe, 2.5" SAS/SATA/NVMe and 3.5" SAS/SATA drives" as versatile configurations [secondary](https://www.wwt.com/article/unlocking-the-future-of-the-data-center-the-evolution-of-dell-poweredge-servers).
- R470/R570 optimized single-socket: "32 x E3.S NVMe Gen5" (1U/2U), "16 x 2.5" SAS4/SATA RAID - Smart Flow in 2U"; R4715/R5715 AMD: "8 x U.2 NVMe Gen4 in 1U" [secondary](https://infohub.delltechnologies.com/static/media/client/7phukh/DAM_e96e704a-a911-4f0a-bec2-8ab602d0c778.pdf).
- R570 configurator chassis menu (dell.com, live 2026): 2.5" NVMe Gen4 direct / PERC 12, 2.5" SAS4/SATA with PERC 12 (H365i/H965i), 2.5" universal drives, 3.5" 12× HDD + 4× E3.S rear direct Gen5 x4, and EDSFF E3.S chassis in 8/16/32-drive variants — Gen5 x4 NVMe direct, Gen5 x2 for the 32-drive variant (lane-count tradeoff), cold-aisle vs hot-aisle serviceability options [secondary](https://www.dell.com/en-au/shop/servers-storage-and-networking/poweredge-r570-rack-server/spd/poweredge-r570/promo_r570_1?view=configurations).
- PERC 12 = H365i / H965i controllers; S160 software RAID also offered [secondary](https://www.dell.com/en-au/shop/servers-storage-and-networking/poweredge-r570-rack-server/spd/poweredge-r570/promo_r570_1?view=configurations).
- CXL 2.0 memory expansion on 17G: R670, R770, R7725, R6725 (16G had CXL 1.1) [secondary](https://www.wwt.com/article/unlocking-the-future-of-the-data-center-the-evolution-of-dell-poweredge-servers).
- PowerEdge XE (AI): E1.S bays (up to 8× 61.44 TB), E3.S, M.2 boot / BOSS-N1, direct NVMe via CX8 mezzanine ("No HBA or PERC, direct NVMe only, OS RAID"), BlueField-3 SuperNIC [secondary](https://secureserve.co.th/uploads/product-datasheet/2026-04-01/69cc9d92b815f.pdf).
- XE8640 (4U GPU): 8× E3.S NVMe = 122.88 TB max, H965i, iDRAC9/10 era management [secondary](https://expresscomputersystems.com/collections/all-computer-components/products/dell-poweredge-xe8640-8-e3-s-edsff-rack-server-chassis).
- Third-party 17G summaries list R670/R6715/R6725 at "up to 20x EDSFF E3.S" and R770 2U at 12× 3.5"/24× 2.5" — treat reseller model-matrix numbers as [secondary] and prefer Dell's configurator/WWT brief where they differ (see §17 conflicts).

### 15.2 HPE ProLiant Compute Gen12

