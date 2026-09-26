---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/part-4
title: "Step 9 — Enterprise SSD Hardware (Phase A) (part 4)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["latency", "nand", "research"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [113, 122]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 940571b1232d7e8d8154a0cf1192d7800da4d9224455d51d1a38472bc002ec8c
---

# Step 9 — Enterprise SSD Hardware (Phase A) (part 4)

- PCIe Gen4 x4, NVMe 1.4b; Micron 176-layer 3D TLC NAND [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- PRO (1 DWPD, up to 15.36 TB) / MAX (3 DWPD, up to 12.8 TB); form factors U.3, E1.S, M.2 2280/22110 [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Performance: 6,800/5,600 MB/s seq; up to 1,000,000 random read IOPS / 400,000 write; 80 µs read / 15 µs write typical latency [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Endurance: PRO up to 28,000 TBW @ 1 DWPD; MAX up to 70,000 TBW @ 3 DWPD; MTTF 2M device hours; UBER <1/10^17; full PLP; 132 namespaces [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Security: TCG Opal Rev 2.01, digitally signed firmware; asymmetric roots of trust / secure boot / secure execution environment on MAX listings [secondary](https://eu.shi.com/Product/45125244/Micron-7450-MAX-SSD).
- Reseller example: 7450 Pro 7.68 TB U.3, 6,800/5,600 MB/s, 1,000K/180K IOPS, 1 DWPD, 2M h MTBF [secondary](https://www.bigw.com.au/product/micron-7450-pro-7-68tb-gen4-nvme-enterprise-ssd-u-3-6800-5600-mb-s-r-w-1000k-180k-iops-25700tbw-1dwpd-2m-hrs-mtbf-server-data-centre-5yrs/p/9902768775).
- Micron 7500 (Gen4 mainstream, 2024 launch): no specs captured in this research pass — gap, Section 14 [unverified].

---

