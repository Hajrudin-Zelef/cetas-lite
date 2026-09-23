---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/20-interface-timeline-1990-2026
title: "20. Interface timeline (1990 → 2026)"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: []
keywords: []
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [567, 611]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: b9a496eeb6572f17fff3504a8ca2f21aac9c06c5e423114e82a128e8cf4a2d26
---

# 20. Interface timeline (1990 → 2026)

## 20. Interface timeline (1990 → 2026)

| Year | Milestone | Notes |
|---|---|---|
| 1990s | Parallel ATA / SCSI | Legacy parallel buses |
| 2003 | SATA 1.0 (1.5 Gb/s) | Serial ATA debut |
| 2003 | SAS 1.0 (3 Gb/s) | Serial Attached SCSI |
| 2008–09 | SATA III (6 Gb/s) | Last SATA generation — "SATA III in 2009, but there never was a SATA IV" [secondary](https://www.networkworld.com/article/4121569/reports-of-satas-demise-are-overblown-but-the-technology-is-aging-fast.html) |
| 2010 | PCIe 3.0 | 8 GT/s NRZ |
| 2011 | NVMe 1.0 work begins | NVM Express Work Group formed |
| 2013 | NVMe 1.0 released | First NVMe spec over PCIe |
| 2013–14 | 12G SAS (SAS-3) | Still in wide use in 2026 |
| 2014 | NVMe 1.1, NVMe-oF work starts | Fabric concept begins |
| 2015 | M.2 NVMe client SSDs ship | Client NVMe takeoff |
| 2016 | NVMe-oF 1.0 (RDMA/FC) | First fabric spec |
| 2017 | PCIe 4.0 spec | 16 GT/s |
| 2018 | NVMe/TCP work; U.3 (SFF-TA-1001) | TCP transport + universal bay |
| 2019 | PCIe 5.0 spec; NVMe 1.4 | 32 GT/s |
| 2020 | EDSFF demos (E1.S/E1.L/E3) | SNIA form factors debut [secondary](https://www.businesswire.com/news/home/20200630005242/en/4781834/KIOXIA-Demonstrates-New-EDSFF-SSD-Form-Factor-Purpose-Built-for-Servers-and-Storage) |
| 2021 | NVMe 2.0 family; 24G SAS (SAS-4) | Spec refactor; SAS-4 silicon |
| 2022 | PCIe 6.0 spec; Samsung/WD ZNS MOU | PAM4 era; D2PF alliance [secondary](https://www.eetimes.com/samsung-western-digital-unite-around-zoned-storage) |
| 2023 | Samsung FDP whitepaper (TP4146) | Host-directed placement defined [vendor-reported](https://download.semiconductor.samsung.com/resources/white-paper/FDP_Whitepaper_102423_Final.pdf) |
| 2024 | NVMe 2.1 (Aug) | FDP, live migration, computational [secondary](https://www.techpowerup.com/325310/nvm-express-releases-nvme-2-1-specifications) |
| 2025 | NVMe 2.3 (Aug); PCIe 7.0 (Jun) | Rapid Path Failure Recovery; 128 GT/s [official](https://pcisig.com/faq?field_category_value%5B%5D=pci_express_7.0&keys=) |
| 2026 | NVMe 2.4 (Aug); Gen5 servers mainstream | PQC, voltage monitoring; Dell 17G, HPE Gen12 [official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage) |
| 2028 (planned) | PCIe 8.0; PCIe 7.0 products | 256 GT/s target [secondary](https://convergedigest.com/pci-sig-targets-2028-release-of-pcie-8-0-at-256-gt-s/) |

*Early entries (1990–2016) are well-established industry history [independent]; dated spec releases above are cited per row.*

## 21. Drive-type ↔ bay-type interoperability matrix

| Drive \ Bay | SATA bay | SAS bay | U.2 NVMe bay | U.3 universal bay | E1.S / E3.S slot |
|---|---|---|---|---|---|
| SATA SSD/HDD | ✅ | ✅ (SAS controllers accept SATA) | ❌ | ✅ | ❌ |
| SAS SSD/HDD | ❌ | ✅ | ❌ | ✅ | ❌ |
| U.2 NVMe | ❌ | ❌ | ✅ | ✅ (x4 NVMe) | ❌ |
| U.3 NVMe | ❌ | ❌ | ✅ (x4, if wired x4) | ✅ (x1/x2/x4) | ❌ |
| E1.S / E3.S | ❌ | ❌ | ❌ | ❌ | ✅ (same family) |
| M.2 | ❌ | ❌ | ❌ (needs carrier/adapter) | ❌ | ❌ |

- U.3 bay accepts all three protocols on the same SFF-8639 connector thanks to SFF-TA-1001 wiring + tri-mode controller + UBM [secondary](https://www.storagereview.com/news/evolving-storage-with-sff-ta-1001-u-3-universal-drive-bays?amp).
- U.2 drive in U.3 bay: works at x4 NVMe; U.2 lacks the x1/x2 negotiation modes — "the difference between U.2 and U.3 is that U.3 uses a new controller specification to support x1, x2, and x4 links, while U.2 only supports x4" [secondary](https://www.storagereview.com/news/evolving-storage-with-sff-ta-1001-u-3-universal-drive-bays?amp).
- EDSFF and U.2/U.3 are mechanically and electrically incompatible — different connectors (4C+ vs SFF-8639); no passive adapter changes this at enterprise power levels [independent].
- M.2 → U.2 works via powered adapters/carriers (common in whitebox), but M.2 is not hot-swap by itself [independent].

