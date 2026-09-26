---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/11-2-what-gen5-means-for-storage-in-2026
title: "11.2 What Gen5 means for storage in 2026"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Broadcom", "Samsung"]
dates: ["2025-06-11"]
keywords: ["accelerator", "compute", "consumer", "gpus", "memory", "research", "serdes"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [257, 290]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 9a9db3ab872684b7add25e09c8124ea208c769a34d826c6926a9346063a579df
---

# 11.2 What Gen5 means for storage in 2026

- PCIe 6.0 was "introduced in 2022 but has not yet seen widespread implementation in consumer products"; 7.0 products are "unlikely… before 2028" — spec ratification runs ~3 years ahead of volume [secondary](https://themunicheye.com/pci-express-7-0-finalized-data-transfer-speeds-23001).
- PCIe 7.0 official facts: 128.0 GT/s raw, "up to 512.0 GB/s bi-directional bandwidth via a sixteen lane (x16) configuration", PAM4 signaling, Flit-based encoding, improved power efficiency, full backward compatibility [official](https://pcisig.com/faq?field_category_value%5B%5D=pci_express_7.0&keys=).
- PAM4 arrived with 6.0 ("allows each transfer to convey two bits of data using four voltage levels, necessitating more sophisticated controller designs"); 7.0 keeps PAM4 + Flit [secondary](https://themunicheye.com/pci-express-7-0-finalized-data-transfer-speeds-23001)[secondary](https://www.eetimes.com/pcie-7-0-keeps-pace-with-ai-demands/).
- PCIe 7.0 was released to members on June 11, 2025; PCI-SIG president Al Yanes: "It takes three years to develop a specification. It takes three years for innovation. It takes three years for feedback on the previous technology" [secondary](https://convergedigest.com/pci-sig-finalizes-pcie-7-0-specification-at-128-0-gt-s/)[secondary](https://www.eetimes.com/pcie-7-0-keeps-pace-with-ai-demands/).
- **Optical Aware Retimer ECN**: alongside 7.0, PCI-SIG updated PCIe 6.4 and 7.0 "to enable standardized PCIe operation over optical fiber using retimer-based solutions" — "a pivotal step toward mainstream adoption of optical links in data center and AI accelerator topologies" [secondary](https://convergedigest.com/pci-sig-finalizes-pcie-7-0-specification-at-128-0-gt-s/).
- PCIe underpins CXL ("PCIe serves as the foundation for Compute Express Link, enabling the connection of additional memory modules") — storage and memory expansion ride the same SerDes [secondary](https://themunicheye.com/pci-express-7-0-finalized-data-transfer-speeds-23001).

### 11.2 What Gen5 means for storage in 2026

- PCIe 5.0 x4 ≈ 16 GB/s raw per drive link — the reason flagship Gen5 SSDs (KIOXIA CM9: 14.8 GB/s seq read) finally saturate what Gen4 x4 (8 GB/s) could not [vendor-reported](http://www.techpowerup.com/336800/kioxia-announces-first-enterprise-nvme-ssd-with-8th-gen-bics-flash-technology)[independent].
- All 2026 server platforms in this research are PCIe Gen5 native: Dell 17G, HPE Gen12, Supermicro Petascale, MSI DC-MHS [secondary].
- Gen5 signal integrity forces retimers/redrivers on longer traces and cables: Icy Dock's PCIe 5.0 x8 MCIO adapter integrates a "PCIe Redriver" plus 100 MHz clock buffer "for stable enterprise timing" [vendor-reported](https://datasheet.itscope.com/2.1/t/Hc4xNMicnBB3UhUhPI7m5PtNPzD0l5IEpv3OCf9-3DLS6QFmBJ1X4enr1g6nayNaxVVAqsib773xL6nayNaxVVAqsib773xL6rAddIgK0ysFp4G5BB-pHMRZdELYWY76YkdU7i2THKwwJYzCHH9wzfwyQ9eeDukqo7AP5WTJh8fSFx7-2rfftuzR-I0fP8Q_HPlvGNMEFdCTcgRsloU).
- Note: the datasheet URL above is a long vendor CDN link — cited verbatim as returned; treat link rot risk as [unverified] for long-term retrieval.

### 11.3 Bifurcation, switches, retimers

- PCIe bifurcation (splitting x16 into x4/x4/x4/x4 etc.) is how motherboards feed multiple M.2/U.2 NVMe drives from one slot; bifurcation is a **motherboard** capability — adapter cards like the Icy Dock MB409A5 "itself does not perform bifurcation" [vendor-reported](https://datasheet.itscope.com/2.1/t/Hc4xNMicnBB3UhUhPI7m5PtNPzD0l5IEpv3OCf9-3DLS6QFmBJ1X4enr1g6nayNaxVVAqsib773xL6nayNaxVVAqsib773xL6rAddIgK0ysFp4G5BB-pHMRZdELYWY76YkdU7i2THKwwJYzCHH9wzfwyQ9eeDukqo7AP5WTJh8fSFx7-2rfftuzR-I0fP8Q_HPlvGNMEFdCTcgRsloU).
- PCIe switch adapters multiply NVMe density: HighPoint Rocket 1624A — PCIe Gen5 x16 host, dual MCIO 8i ports, "supports up to 16 NVMe SSDs" per adapter, configurable downstream 1x16 down to 16x1, hot-plug/hot-swap, direct P2P between GPUs/NICs/SSDs bypassing the CPU [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw)[vendor-reported](https://electronicsbuzz.in/highpoint-unveils-comprehensive-pcie-gen-5-mcio-expansion-ecosystem/).
- Native OS support (Linux/Windows, x86 and ARM) with no proprietary driver overhead is the selling point of switch-based NVMe expansion [vendor-reported](https://electronicsbuzz.in/highpoint-unveils-comprehensive-pcie-gen-5-mcio-expansion-ecosystem/).

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

