---
id: etape6-phasec-optics-cabling/00-front-matter/2-2-400g-module-types
title: "2.2 400G module types"
domain: front-matter
role: reference
task: reference
actors: ["Lambda"]
dates: ["2024-03", "2026-03", "2026-10"]
keywords: ["cost", "dsp", "ethernet", "lpo", "optics", "parameters", "training", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [238, 285]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 7551dfef0b0d45549ad0e8ad2bccfe546974509bcb1a7f43b99505723fd3bf50
---

# 2.2 400G module types

### 2.2 400G module types

All facts [vendor-reported] from FS (fs.com) and Ascent Optics (ascentoptics.com) guides unless noted; standard refs cross-checked with IEEE sources where marked.

| Type | Reach | Fiber | Connector | Wavelength/optics | Electrical | Typical power | Standard |
|---|---|---|---|---|---|---|---|
| 400G-SR8 | 70 m OM3 / 100 m OM4 | MMF | MPO-16 | 850 nm VCSEL, 8×50G PAM4 | 8×50G PAM4 (400GAUI-8) | ≤10 W | IEEE 802.3cd |
| 400G-SR4 | 70 m OM3 / 100 m OM4 | MMF | MPO-12 | 850 nm VCSEL, 4×100G PAM4 | 8×50G PAM4 | ≤8–8.5 W | IEEE 802.3cd (InfiniBand NDR variant) |
| 400G-SR4.2 | 150 m (OM5; 70 m OM3) | MMF | MPO-12 | 850 + 910 nm (BiDi-ish dual-λ, 4 lanes × 2λ), 8×50G PAM4 | 8×50G PAM4 | ≤12 W | IEEE 802.3cm-class |
| 400G-DR4 | 500 m | SMF (parallel) | MPO-12/APC | 1310 nm, 4×100G PAM4 (parallel fibers, uncooled DFB) | 8×50G PAM4 | ≤9–10 W (FS); 8–10 W (Ascent) | IEEE 802.3bs; 100G Lambda MSA |
| 400G-XDR4 / DR4+ | 2 km | SMF (parallel) | MPO-12/APC | 1310 nm, 4×100G PAM4 | 8×50G PAM4 | ≤9–10 W | MSA/industry spec (not IEEE) |
| 400G-FR4 | 2 km | SMF (duplex) | Duplex LC | **CWDM4: 1271/1291/1311/1331 nm**, 4×100G PAM4 (EML or uncooled DML) | 8×50G PAM4 | ≤9–12 W | 100G Lambda MSA |
| 400G-LR4 | 10 km | SMF (duplex) | Duplex LC | CWDM4 (1271/1291/1311/1331 nm), 4×100G PAM4 EML | 8×50G PAM4 | 12–15 W (Ascent); ≤9–10 W (FS) — **[flagged conflict: vendor tables differ]** | IEEE 802.3bs / 100G Lambda MSA |
| 400G-FR8 | 2 km | SMF | CS | CWDM4, 8×50G PAM4 | 8×50G PAM4 | ~10 W | IEEE 802.3bs |
| 400G-LR8 | 10 km | SMF | Duplex LC | LAN-WDM (~1295/1300/1304/1309 nm), 8×50G PAM4 | 8×50G PAM4 | ~10 W | IEEE 802.3bs |
| 400G-PLR4 | 10 km | SMF (parallel) | MPO-12 | 1310 nm, 4×100G PAM4 | 8×50G PAM4 | ≤9 W | industry |
| 400G-ER8 | 40 km | SMF | Duplex LC | 1310 nm / LAN-WDM, 8×50G PAM4 | 8×50G PAM4 | ≤10 W | industry |

- **DR4 vs FR4 distinction**: DR4 = 4 parallel single-λ 1310 nm fibers per direction (no MUX/DEMUX, lowest cost, easy 4×100G-DR breakout) [vendor-reported]; FR4 = 4 CWDM4 wavelengths MUXed onto a single fiber pair (2 fibers total) [vendor-reported].
- **FR4 uses uncooled DFB/DML + host FEC**; LR4 uses higher-power EML + enhanced DSP — LR4 modules cost ~30–50% more than FR4 [vendor-reported — Ascent]. 100G CWDM4 optics typically cost 30–40% less than 100G LR4 due to uncooled DFB vs EML [vendor-reported — Ascent via Medium].
- **Key MSA/standard references**: IEEE 802.3bs-2017 (200G/400G Ethernet) [independent]; 100G Lambda MSA (100G-per-wavelength DR/FR/LR; "100G Single Lambda MSA") [vendor-reported]; QSFP-DD MSA; OSFP MSA; CMIS Rev 4.x management [vendor-reported].
- Sources: https://ascentoptics.com/blog/400g-qsfp-dd-sr8-dr4-fr4-lr4-transceiver-guide/ ; https://www.fs.com/blog/fs-400g-qsfp-dd-complete-guide-and-solutions-8484.html ; https://ascentoptics.com/blog/400g-osfp-transceiver-types/ ; http://c-light.com/products/details/400G_QSFP112.html (QSFP112 variants: SR4/DR4/FR4/LR4 400G on 4×100G electrical)

### 2.3 800G module types

All electrical: **8 × 100G PAM4 (800GAUI-8, 106.25 Gb/s per lane)** [vendor-reported — FS, expertini]. Facts [vendor-reported] unless noted.

| Type | Reach | Fiber | Connector | Wavelength/optics | Typical max power | Standard |
|---|---|---|---|---|---|---|
| 800G-SR8 | 30 m OM3 / 50 m OM4 (FS QDD-SR8-800G); some sources cite up to 100 m MMF — **[flagged conflict: 50 m vs 100 m claims between vendor tables]** | MMF | MPO-16 APC (QDD); dual MPO-12 APC (OSFP) | 850 nm VCSEL, 8×100G PAM4 | ≤14–16 W | IEEE 802.3df-2024 |
| 800G-2VSR4 | 50 m OM4 | MMF | Dual MTP/MPO-12 APC | 850 nm VCSEL | ≤16 W | industry |
| 800G-DR8 | 500 m | SMF (parallel) | MPO-16 APC (QDD); dual MPO-12 APC (OSFP) | 1310 nm, 8×100G PAM4 (QDD single group; OSFP "dual 4×100G" internally) | ≤16.5 W (DSP); ≤15 W cited elsewhere | IEEE 802.3df-2024 |
| 800G-DR8L | 100 m | SMF | Dual MPO-12 APC | 1310 nm | ≤16.5 W | IEEE 802.3df-2024 |
| 800G-DR8-LPO | 500 m | SMF | Dual MPO-12 APC | 1310 nm, linear (no DSP) | **≤8.5 W** | industry (OIF 112G/224G-Linear lineage) |
| 800G-PSM8 | 100 m | SMF (8-fiber) | MPO-based | CWDM 8λ × 100G | ~14 W class | MSA/industry |
| 800G-2DR4 (=2×DR4) | 500 m | SMF | 2× MPO-12 | 2 groups of 4×100G @1310 nm | ~16–18 W | industry (breakout to 2× 400G-DR4) |
| 800G-XDR8 | 2 km | SMF (parallel) | Dual MPO-12 APC / MPO-16 | 1310 nm, 8×100G PAM4 | ≤18 W | industry |
| 800G-2FR4 (=2×FR4) | 2 km | SMF (duplex) | Dual duplex LC (dual CS variants) | CWDM4 1271/1291/1311/1331 nm, 2× 4×100G | ≤16.5–17 W | IEEE 802.3df-2024 |
| 800G-2LR4 (=2×LR4) | 10 km | SMF (duplex) | Dual duplex LC | CWDM4, 2× 4×100G | ≤18 W | IEEE 802.3df-2024 |
| 800G-PLR8 | 10 km | SMF (parallel) | Dual MPO-12 | 1310 nm, 8×100G PAM4 | ≤18 W | industry |
| 800G-FR8 / LR8 | 2 km / 10 km | SMF duplex | Duplex LC | 4×200G PAM4 (single-λ 200G optics) | ~16–18 W class | IEEE 802.3df-2024 PMDs |

- **IEEE 802.3df-2024** [official — IEEE SA]: "Amendment 9: Media Access Control Parameters for 800 Gb/s and Physical Layers and Management Parameters for 400 Gb/s and 800 Gb/s Operation"; adds Clauses 169–173, Annexes 172A/173A to IEEE Std 802.3-2022; **published 15 March 2024**, 278 pages, ISBN 979-8-8557-0541-6; free via IEEE GET program. Sources: https://standards.ieee.org/ieee/802.3.1/11244/ ; https://www.en-standard.eu/ieee-802.3df-2024-ieee-standard-for-ethernet-amendment-9-media-access-control-parameters-for-800-gb-s-and-physical-layers-and-management-parameters-for-400-gb-s-and-800-gb-s-operation/
- Pre-standard interim: **Ethernet Technology Consortium (ETC) 800G Specification r1.1 (Oct 2021)** served as industry template; 802.3df is "very similar, with relatively minor modifications" — differences include PCS-to-PMA lane-mapping restrictions and link-training lane independence to reduce crosstalk [independent — Spirent/EEWorld, Steve Rumsby]. Sources: https://www.eeworldonline.com/how-ieee-802-3df-brings-800g-ethernet-to-life/ ; https://ethernettechnologyconsortium.org/wp-content/uploads/2021/10/Ethernet-Technology-Consortium_800G-Specification_r1.1.pdf (URL given in source text)
- **IEEE P802.3dj** (200/400/800G + 1.6T on 200G/lane): project title "Amendment: MAC Parameters for 1.6 Tb/s and Physical Layers and Management Parameters for 200/400/800 Gb/s and 1.6 Tb/s Operation" [official — IEEE PAR]. PAR originally projected RevCom submission **March 2026** [official]; Nokia/Nordunet deck (2025) stated "standard to be ratified in October 2026" [vendor-reported — https://nordu.net/ncw2025/static/presentations/ncw2025/11/lightning-talks-2/16-terabit-ethernet-paving-the-road-for-next-gen-ai-datacenters.pdf]; as of Sept 2026, Ethernet Alliance is running 1.6T multi-vendor interoperability showcases at ECOC 2026 (Sept 20–24, Málaga) "alongside IEEE P802.3dj, which is developing specifications extending Ethernet through 1.6 Tb/s using 200 Gb/s-per-lane signaling" — **i.e., not yet ratified as of 22 Sept 2026**; ECOC Sept 23 2026 Market Focus panel to examine "networking beyond 1.6T and industry work toward 400 Gb/s electrical and optical signaling" [secondary — Converge Digest]. Source: https://convergedigest.com/ethernet-alliance-1-6t-ethernet-ecoc-2026/

### 2.4 1.6T early status (as of 22 Sept 2026)

