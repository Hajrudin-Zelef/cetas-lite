---
id: etape6-phasec-optics-cabling/00-front-matter/part-13
title: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 13)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "Cohere"]
dates: ["2022-03"]
keywords: ["optics", "2nm", "cpo", "dsp", "ethernet", "latency", "lpo"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [414, 419]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 172f2a8910d11b7d7fd18fd49f15445c9d2d73da39be03f7e8fc57afce2b0bae
---

# Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 13)

#### 3.3.1 DSP makers
- **Broadcom**: optical PHY portfolio spans 16 nm → 7 nm → 5 nm nodes; five flagship DSP platforms: **Centenario, Jesko, Portofino, Gemera, Cygnus**, covering 100G/400G/800G/1.6T PMDs [secondary: https://jrsekwele.co.za/kiloptics/files/1-6t-optical-module-yield---JR-Sekwele-Optical-Networks-&-Photonic-Group_Sat-21-Jun-2025-4058.pdf]. **BCM85812** (OFC 2023): industry's first 5 nm 100G/lane optical PAM-4 DSP PHY with integrated TIA + laser driver; 8×106G PAM4 (53 Gbaud), for 800G DR8 / 2×400G FR4 / 800G AOC; compliant with OIF 3.2T co-packaged specs [official: https://www.broadcom.com/products/ethernet-connectivity/phy-and-poe/optical/bcm85812; https://docs.broadcom.com/doc/85812-PB]. **BCM85822**: 5 nm 800G (8:4) PHY, 8×106G client → 4×226G PAM4 line (113 Gbaud), supports 400G DR1.2 and 1.6T DR4.2, OSFP-XD form factor [official: https://www.broadcom.com/products/ethernet-connectivity/phy-and-poe/optical/bcm85822].
- **Marvell**: PAM4 optical DSP line spanning 800G **Spica** (first 800G PAM4 DSP to ship in volume, March 2022 production release; InnoLight shipped 800G optics on Spica to early adopters) [official via stockwatch reprint: http://www.stockwatch.com/news/item/u-prsf80426-u!mrvl-20220303/u/mrvl]. 1.6T generation: **Nova** (first 5 nm 200G/lane 1.6T DSP, 2023), **Ara** (3 nm 1.6T platform, 2024), and **Libra** (2 nm 800G DSP, announced earlier in 2026) [official via Marvell ECOC 2026 announcement: https://www.financialcontent.com/article/bizwire-2026-9-21-marvell-to-showcase-industry-first-2nm-optical-technology-demos-for-ai-data-center-infrastructure-at-ecoc-2026]. Nova/Ara support Ethernet + InfiniBand [secondary: https://www.ainvest.com/news/marvell-owns-dsp-moves-ai-networks-copper-light-packaged-fork-decides-worth-2609/]. ECOC 2026 demos (Sept 20–24, Málaga): industry-first 2 nm optical interconnects — 400G/lane optical PAM4 for 3.2T, 800G ZR/ZR+ pluggable with MACsec, 1.6T ZR, coherent-lite O-band, and 102.4T CPO platform [official]. **Note: ECOC items are demonstrations; Marvell gave no availability dates** [secondary: https://aistockwire.com/blog/marvell-mrvl-2nm-optical-demos-ecoc-2026; https://biztechweekly.com/at-ecoc-2026-marvells-2nm-optical-demos-point-toward-3-2t-ai-networks-not-yet-deployments/].
- **Semtech**: **FiberEdge®** linear driver/TIA portfolio; demonstrated 200G PAM4 Mach–Zehnder modulated DFB laser driven by Semtech differential driver for 1.6T (ECOC 2023, with Coherent) [official via Coherent press release: https://www.globenewswire.com/en/news-release/2023/10/02/2752425/11543/en/Coherent-to-Demonstrate-Next-Generation-Transceiver-and-Semiconductor-Laser-Technology-for-800G-and-1-6T-Transmission-in-AI-Networks-at-ECOC-2023.html]. Semtech also co-chairs the LPO MSA and ACC-MSA [secondary: https://www.lightwaveonline.com/home/article/55141192/lpo-msa-achieves-multi-vendor-interoperability; https://convergedigest.com/new-industry-msa-targets-low-power-copper-interconnect-for-800g-and-1-6t/]. **MaxLinear and Inphi (now Marvell)**: named in task but no 2026 product specifics surfaced in searches — **no sourced product names found; flagged as gap**.
- Alphawave Semi: PAM4 DSP PHY IP used with InnoLight LPO OSFP in OFC 2024 demo (PCIe 6.0 over optics) [official press: https://www.silicon.co.uk/press-release/alphawave-semi-and-innolight-collaborate-to-demonstrate-low-latency-linear-pluggable-optics-with-pcie-6-0-subsystem-solution-for-high-performance-ai-infrastructure-at-ofc-2024].

