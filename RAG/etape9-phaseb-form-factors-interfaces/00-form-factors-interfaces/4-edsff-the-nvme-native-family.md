---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/4-edsff-the-nvme-native-family
title: "4. EDSFF — the NVMe-native family"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: ["2026-04-01"]
keywords: ["alignment", "compute", "cost", "datacenter", "gpu", "gpus", "nand"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [46, 95]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 69ef90010bb6f23c1174ec4a88b70adf77c1dd2051188fb24196d9f38e5f4169
---

# 4. EDSFF — the NVMe-native family

## 4. EDSFF — the NVMe-native family

### 4.1 Family map (SNIA)

- EDSFF = Enterprise and Datacenter SSD Form Factor, developed in the SNIA SFF TA work groups; E1 for datacenter, E3 for enterprise storage [official via SNIA members doc](https://members.snia.org/document/dl/53687).
- SNIA's own usage split: **E3 = enterprise storage; E1.S = datacenter compute; E1.L = datacenter storage** [secondary](https://dev-snia-org.pantheonsite.io/sites/default/files/2025-10/SNIA-SDC25-Constantine-Why-Another-EDSFF-SSD.pdf).
- Design goals shared across the family: "Same pinout/functions = Same connector; Same connector = Same signal integrity; Same signal integrity = Same PCIe speeds" — one connector ecosystem instead of per-form-factor revalidation [secondary](https://dev-snia-org.pantheonsite.io/sites/default/files/2025-10/SNIA-SDC25-Constantine-Why-Another-EDSFF-SSD.pdf).
- Mechanical specs: E1.S = SFF-TA-1006, E1.L = SFF-TA-1007, E3 (.S/.L) = SFF-TA-1008, pin/signal = SFF-TA-1009 [official](https://members.snia.org/document/dl/53687).
- The EDSFF connector can supply up to 6.6 A continuously, i.e. 79.2 W at nominal 12 V — the ceiling above which host power limits, airflow and device thermals further constrain [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).
- E3.x devices support x4, x8 and even x16 PCIe link widths ("The current drive uses eight lanes of PCI Express Gen4"; PCIe Gen5 drives use higher power), breaking the x4-only assumption of U.2 [secondary](https://www.networkworld.com/article/968945/kioxia-demonstrates-new-high-capacity-ssd-form-factor.html).

### 4.2 E1.S — the 1U compute drive

- E1.S is "optimized for the NVMe drive design for use across all data center and edge systems to scale as mainstream storage" [secondary](https://eetimes.com/nvme-drives-ready-to-embrace-own-form-factors/).
- Thickness variants: 5.9 / 8.01 / 9.5 / 15 / 25 mm. Initial slot power limit: **12 W** for the 5.9/8.01 mm thin variants, **25 W** for 9.5/15/25 mm [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).
- For E1.S/E1.L "SNIA does not specify a separate maximum sustained power below the connector limit" — sustained power is a platform/airflow negotiation, not a fixed spec number [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).
- Samsung PM9A3 (PCIe Gen4 x4, E1.S, 960 GB–7.68 TB, ~3,200 MB/s seq read, 180K random-write IOPS class) was positioned as "the most sought-after storage solution on the market for tier one and tier two cloud datacenter servers", with a collaborative Inspur server reference design [secondary](https://www.electronicsweekly.com/news/business/ssd-snia-based-e1-s-form-factor-2020-05/).
- For datacenters on U.2, "the E1.S form factor will allow hardware engineers to add more SSDs per server, freeing up additional 1U space and lowering the total cost of ownership" [secondary](https://www.electronicsweekly.com/news/business/ssd-snia-based-e1-s-form-factor-2020-05/).
- Dell PowerEdge XE AI servers list "Up to 8 x EDSFF E1.S (61.44 TB)" — showing E1.S carrying the highest-capacity drives in GPU nodes [secondary](https://secureserve.co.th/uploads/product-datasheet/2026-04-01/69cc9d92b815f.pdf).

### 4.3 E1.L — the ruler for capacity

- E1.L ("ruler") thickness variants 9.5 / 18 mm; initial slot power limit **25 W** [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).
- E1.L is the long-ruler format for maximum NAND packages per drive — the capacity-optimized datacenter-storage member of the family [secondary](https://dev-snia-org.pantheonsite.io/sites/default/files/2025-10/SNIA-SDC25-Constantine-Why-Another-EDSFF-SSD.pdf).
- 2026 sighting: DapuStor Roealsen6 R6060 **E1.L 245.76 TB** SSD reviewed by TweakTown — the quarter-petabyte ruler drive is real product, not concept [secondary](https://www.tweaktown.com/reviews/11605/kioxia-cm9-r-15-36tb-e3-s-enterprise-ssd-performance-king-at-1-dwpd/index.html).
- Kioxia LC9 245 TB-class SSDs are being partnered into Dell PowerEdge servers (per TweakTown's 2026 coverage links), indicating hyperscale/OEM pull for 200 TB+ rulers [secondary](https://www.tweaktown.com/news/105256/kioxia-cm9-series-announced-first-pcie-5-0-nvme-ssds-with-8th-gen-bics-flash-technology/index.html).

### 4.4 E3.S / E3.L — the enterprise workhorse

- E3.S 1T: recommended air-cooled sustained maximum **25 W**; E3.S 2T: **40 W**; E3.L 1T: **40 W**; E3.L 2T: **70 W** — these are "cooling-design recommendations, not the normal consumption of every E3 SSD" [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).
- E3 breaks "free from the design limitations of the 2.5" form factor by supporting higher power budgets (up to 40 watts) and better signal integrity to deliver the performance promised by PCIe Gen 5.0 and beyond" [vendor-reported](https://www.businesswire.com/news/home/20200630005242/en/4781834/KIOXIA-Demonstrates-New-EDSFF-SSD-Form-Factor-Purpose-Built-for-Servers-and-Storage).
- E3.x is a **multi-device** standard: "one common connector, this innovative form factor standard for PCIe technology-based devices, such as NVMe SSDs, graphics processing units (GPUs), and network interface cards (NICs)" — E3.S slots can host non-storage PCIe devices [vendor-reported](https://www.businesswire.com/news/home/20200630005242/en/4781834/KIOXIA-Demonstrates-New-EDSFF-SSD-Form-Factor-Purpose-Built-for-Servers-and-Storage).
- TweakTown on E3.S in 2026: "The E3 form factor is becoming increasingly popular because it is more efficient than conventional U.2 in terms of footprint/density and has a more modern interface designed specifically for NVMe SSDs. However, cooling can be a bit more challenging at a slim 7.5mm thickness" — their 15.36 TB Kioxia CM9-R ran at 56 °C under sustained write with conventional air cooling [independent](https://www.tweaktown.com/reviews/11605/kioxia-cm9-r-15-36tb-e3-s-enterprise-ssd-performance-king-at-1-dwpd/index.html).
- SFF-TA-1034 compatibility wrinkle: an E3 device with a 4C+ connector "will be treated like an OCP NIC 3.0 device in a PMM slot"; E1.S/E1.L mechanicals are "not considered compatible" with the SFF-TA-1034 host slot (thickness/length keep-in violations), and E3 in a 1034 slot needs a host-provided carrier for alignment/retention/extraction [official](https://members.snia.org/document/dl/53687).
- Pin configuration for E3-in-1034 can use the GND/EDSFF_DETECT pin, configured "prior to or simultaneous with 12 V power being applied" [official](https://members.snia.org/document/dl/53687).

### 4.5 EDSFF power/density quick-reference

| Form factor | Thickness options | Initial slot power | Air-cooled sustained max | Connector ceiling |
|---|---|---|---|---|
| E1.S thin | 5.9 / 8.01 mm | 12 W | (not separately specified) | 79.2 W @ 12 V |
| E1.S | 9.5 / 15 / 25 mm | 25 W | (not separately specified) | 79.2 W @ 12 V |
| E1.L | 9.5 / 18 mm | 25 W | (not separately specified) | 79.2 W @ 12 V |
| E3.S 1T | 7.5 mm class | 25 W | 25 W | platform-limited |
| E3.S 2T | 16.8 mm class | 40 W | 40 W | platform-limited |
| E3.L 1T | — | 40 W | 40 W | platform-limited |
| E3.L 2T | — | 70 W | 70 W | platform-limited |

*Source: SNIA EDSFF power guidance via [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html). Interpret carefully: "The initial slot power limit, recommended sustained maximum, and connector capability are different values and should not be treated as the SSD's normal power consumption."*

