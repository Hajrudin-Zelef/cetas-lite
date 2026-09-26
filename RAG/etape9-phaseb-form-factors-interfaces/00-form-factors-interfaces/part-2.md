---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/part-2
title: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle) (part 2)"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: []
keywords: []
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [37, 45]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 0aeee12c2e0f71a7947103ce00c406806280c9b0f95fbd0b4c21ba6bbb6b6d57
---

# Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle) (part 2)

- U.3 is defined by SFF-TA-1001, titled "Specification for Universal x4 Link Definition for SFF-8639", co-authored by Broadcom and HPE and driven through the SNIA SFF Technology Affiliate [secondary](https://www.storagenewsletter.com/2017/11/17/broadcom-server-storage-hba-and-megaraid-solutions-with-sff-ta-1001-u-3-reference-platform/).
- One bay, three protocols: "U.3 allows SAS and SATA HDDs and SAS, SATA, and NVMe SSDs to operate in a single bay without the complexity of wiring for multiple protocols. It enables single, dual, and wide-port SAS, SATA, and x1, x2 or x4 NVMe devices to all work on the same shared signals and connectors" [secondary](https://www.globenewswire.com/news-release/2017/11/07/1176767/0/en/Broadcom-Extends-Server-Storage-Leadership-with-Comprehensive-SFF-TA-1001-U-3-Reference-Platform.html).
- Mechanically backward compatible with U.2/SFF-8639; the difference is pin usage, slot detection, and host/backplane wiring per SFF-TA-1001 [secondary](https://www.storagereview.com/news/evolving-storage-with-sff-ta-1001-u-3-universal-drive-bays?amp).
- The U.3 tri-mode platform has three components: (1) tri-mode controller, (2) SFF-8639 connector on drive and backplane, (3) Universal Backplane Management framework (SFF-TA-1005/UBM) [secondary](https://www.storagereview.com/news/evolving-storage-with-sff-ta-1001-u-3-universal-drive-bays?amp).
- The tri-mode controller's "auto-sense capability" determines which of the three interface protocols is being serviced, "eliminating the need for OEMs to use one controller dedicated to SAS and SATA and a different controller for NVMe" [secondary](https://www.storagereview.com/news/evolving-storage-with-sff-ta-1001-u-3-universal-drive-bays?amp).
- KIOXIA CM6/CD6 were "SFF-TA-1001 conformant (U.3) SSDs to support universal drive bay systems that can accept NVMe, SAS and SATA SSDs" — KIOXIA positioned this as "a replacement SATA technology" path for HPE ProLiant [vendor-reported](https://www.businesswire.com/news/home/20210623005830/en/5000252/KIOXIA-America-Showcases-PCIe-4.0-24G-SAS-SSDs-at-HPE-Discover)[vendor-reported](https://blog-us.kioxia.com/post/2020/03/05/the-fastest-available-u-3-pcie-4-0-nvme-ssds-are-now-shipping).
- Broadcom's 9600-series tri-mode adapters "are compatible with existing PCI Express SFF-8639 Module (U.2) backplanes", protecting U.2 investments while moving to 24G SAS [vendor-reported](https://WWW.scan.co.uk/products/16-port-broadcom-megaraid-9660-16i-raid-controller-4gb-cache-2x-sas-sff-8654-sata-iii-nvme-raid-0-1).
- U.3 vs EDSFF in 2026: U.3 wins on fleet compatibility and mixed media (HDD+SSD in one chassis); EDSFF wins on per-rack NVMe density and forward PCIe headroom — Dell/HPE sell both side by side [independent].

