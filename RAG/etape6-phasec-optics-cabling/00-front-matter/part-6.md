---
id: etape6-phasec-optics-cabling/00-front-matter/part-6
title: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 6)"
domain: front-matter
role: reference
task: reference
actors: ["Nvidia"]
dates: ["2019-07", "2021-05", "2024-06"]
keywords: ["asic", "datacenter", "ethernet", "gpu", "liquid cooling", "nvidia"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [225, 237]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: fa58f0fbc47707ff77bdac84a1f67c245bbbdbc3c0a30cf6a4f48681fc68a749
---

# Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 6)

- QSFP-DD backward compatible with QSFP+, QSFP28, QSFP56, QSFP112 (DD cage accepts 4-lane modules) [vendor-reported]. Note: inserting QSFP112 into QSFP-DD cage requires host ASIC support for 4×100G PAM4 [vendor-reported].
- **QSFP-DD MSA hardware spec revisions** [secondary — DKN Consulting PDF + fibermall]: Rev 5.0 July 2019; Rev 5.1 Aug 2020; **QSFP-DD800 (8×100G) released 24 May 2021** (qsfptek [secondary]); **Rev 7.0 announced at ECOC 2023** defines QSFP-DD1600 (8×200G PAM4, 1.6T, backward compat with DD800/QSFP56/QSFP28, CMIS 5.x, typical power 20–28 W, riding heatsink); **Rev 7.1 released June 2024** (refines module/connector definitions) [secondary — fibermall.com/blog/1-6t-qsfp-roadmap-qsfp-dd1600-guide.htm].
- Power notes (see also §2.6): early QSFP-DD thermal class 7–12 W [secondary]; QSFP-DD800 modules rated up to 18 W (vendor-reported FS 800G guide); QSFP-DD1600 typical 20–28 W [secondary].

#### OSFP family
- **OSFP** — initially designed 2016 for 8 × 50G = 400G [secondary — qsfptek]; dimensions 22.58 × 107.8 × 13.0 mm vs QSFP-DD 18.35 × 89.4 × 8.5 mm [vendor-reported]; 36 ports per 1U = 14.4 Tbps at 400G [vendor-reported]. Not backward compatible with QSFP+/QSFP28 (larger cage) [vendor-reported].
- **OSFP 4.0** released 2021 — 800G via 8 × 100G [secondary — qsfptek].
- **OSFP MSA Revision 5.1** defines implementations for 50G-, 100G-, and 200G-class signaling per electrical lane → supports 1.6T-class modules (8 × 200G) [secondary — roboticsandautomationnews.com, 22 Sept 2026]. Source: https://roboticsandautomationnews.com/2026/09/22/osfp-modules-the-complete-guide-to-400g-800g-and-1-6t-optical-transceivers-for-ai-and-hyperscale-data-centers/104982/
- Power class: early OSFP 12–15 W [secondary]; 800G OSFP modules rated up to 18 W in vendor tables (see §2.6) [vendor-reported]. 36-port 1U panel delivers 28.8 Tbps at 800G/port [secondary].
- Three thermal tops: finned-top (open top, air-cooled), closed finned-top (air + liquid), flat top (cold-plate liquid cooling) [vendor-reported — FS]. Flat-top OSFP used with NVIDIA DGX H100 Cedar7 GPU links [vendor-reported].
- **OSFP-XD** (eXtra Dense) — MSA-defined 1.6T variant: **16 electrical lanes** (16×100G = 1.6T; future path 16×200G = 3.2T), same cage/panel cutout/thermal as OSFP, up to ~40W design target, up to 32 ports per 1RU / 64 per 2U, supports 100GBASE-CR1 DAC [secondary — Medium/aicplight, 16 Sept 2026; Cosolvic blog]. Sources: https://medium.com/@aicplight888/1-6t-optical-transceiver-form-factor-comparison-osfp1600-vs-osfp-xd-7bc0dcc84670 ; https://cosolvic.com/blog/800g-data-center-connectors-osfp-vs-qsfp-dd-comparison/
- Use-case pattern: OSFP strong in AI/HPC and InfiniBand/NDR (e.g., 800G OSFP DR8 for IB NDR fabrics) in addition to Ethernet; QSFP-DD dominant in Ethernet datacenter [vendor-reported].

