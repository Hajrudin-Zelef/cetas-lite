---
id: etape6-phasec-optics-cabling/00-front-matter/10-2-fixed-51-2t-systems-64-800g-class
title: "10.2 Fixed 51.2T systems (64 × 800G class)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "Nvidia"]
dates: ["2023-10", "2024-10-16", "2025-10", "2026-02"]
keywords: ["asic", "cpo", "dci", "ethernet", "full-duplex", "hbm", "inference", "latency", "nvidia", "rack-scale", "serdes"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1510, 1532]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: d779d8ac7ce4a5c56e6d766b8c242e466d8c8379ee6ac81fe2bddd8736ecff35
---

# 10.2 Fixed 51.2T systems (64 × 800G class)

### 10.2 Fixed 51.2T systems (64 × 800G class)

**Arista 7060X6 (Tomahawk 5).** DCS-7060X6-64PE: 64 × 800G OSFP800, 2RU, 51.2 Tbps, up to 128 × 400G via breakout, 165 MB buffer, latency from 700 ns; Arista positions the family for AI backend and data-center applications [official] (arista.com 7060X6 series). DCS-7060X6-64PE-B: same 64 × 800G / 51.2T class with breakout distinctions in maximum 100G density vs 64PE [official] (quick-look). DCS-7060X6-32PE: 32 × 800G OSFP800, 1RU, 25.6 Tbps, up to 64 × 400G, 84 MB buffer, latency from 700 ns [official]. Arista's 1.6T-rack-scale **7060XE7** family (Tomahawk 6, 102.4T), four variants: 7060XE7-64PS (64 × 1.6T, air-cooled, 4RU, planned Q4 2026), 7060XE7-64PRS (64 × 1.6T, air/liquid options, early 2027 indicated for some variants), 7060XE7-64PRS-RV3-L (specialized 2OU liquid-cooled ORv3, fanless, 224G SerDes, planned Q1 2027), 7060XE7-128PE (128 × 800G, air-cooled 4RU, 100G SerDes, backward compatibility, early 2027 reporting) [independent] (networkworld.com; corroborated by sdxcentral.com).

**Arista 7800R4 modular + 7280R4 fixed.** 7800R4 modular family (4/8/12/16-slot chassis) reaches up to 576 ports of 800GbE in a single system, using 36-port 800GbE OSFP line cards (DCS-7800R4-36PE-LC / -LC# spare/in-chassis variants, plus C- and K-series SKUs); one chassis footnote marks a variant "available in mid-2025" [official] (arista.com 7800R4 AI-Spine datasheet). **HyperPort**: a single 3.2 Tbps high-speed interface for scale-across inter-DC interconnection, claimed to cut AI bandwidth-flow job completion time by 44% versus traditional load balancing over four separate 800G ports [official] (Arista investor PR, 2025). **7280R4** fixed: 32-port 800GbE system for AI/DC spine or backbone routing, and a 64 × 100GbE with 10 × 800GbE system for AI/DC leaf [official]. Availability per October 2025 announcement: 7800R4 modular systems shipping now, two new 7800R4 line cards shipping now, two new 7280R4 platforms shipping now, new 7020R4 chassis platforms Q1 2026, 7800R4 with HyperPort Q1 2026 [official]. No verified model named "7388X5" as an 800G fixed system was found — do not use without a source [unverified].

**Cisco Nexus 9364E family (Silicon One).**
- N9364E-SG2-Q: 64 × 800G QSFP-DD, Silicon One G200, 2RU, 51.2 Tbps one-way ASIC capacity [official].
- N9364E-SG2-O: 64 × 800G OSFP with 256 MB shared on-die buffer [official] (cisco.com N9364E-SG2 datasheet).
- N9364E-SG2X-Q / -O: 64 × 800G QSFP-DD or OSFP, 256 MB shared buffer, typical power ~972 W, maximum 2,779 W [official] (static-cisco.com SG2X datasheet).
- N9364E-SP2R-Q / -O: 64 × 800G (QSFP-DD or OSFP), Silicon One P200, 3RU, 144 MB shared on-die buffer + 16 GB HBM, typical power 2,500 W, maximum 5,100 W, reported 102.4 Tbps bidirectional (= 51.2T one-way — do not compare with unidirectional ASIC figures) [official] (cisco.com SP2R datasheet).
- N9364F-SG3 (Silicon One G300): reported 64 ports supporting 800G/1.6T at 102.4 Tbps, announced February 2026; exact cage type, buffer, power, form factor, commercial ship date unverified against an official datasheet [vendor-reported].

**Dell PowerSwitch.** **PowerSwitch Z9864F-ON**: 64 × 800GbE OSFP112, 2RU, 51.2 Tbps half-duplex/one-way fabric capacity, 20.3 Bpps, sub-700 ns latency, 165.2 MB packet buffer, Dell Enterprise SONiC, targeting AI/ML, cloud, Web 2.0, general data-center; breakout up to 320 logical ports per Dell/reseller literature [official] (delltechnologies.com Z9864F-ON spec sheet; corroborated by dell.com UK Z-series page).

**NVIDIA SN5600 family (Spectrum-4).** SN5600: 64 × 800GbE OSFP, 51.2 Tbps, 33.3 Bpps, 2RU, Spectrum-4, 160 MB buffer, 8 × 100G PAM4 lanes per native 800G port, typical switch power ~940 W with passive copper [vendor-reported] (SN5600 datasheet mirror). Mirrored SN5000-series datasheet lists SN5600 ordering SKUs 920-9N42F-00RI-7N0 (ONIE + NOS authentication), 920-9N42F-00RI-5N0 (ONIE), 920-9N42F-00RI-7C0 (Cumulus Linux authentication), each 64 × OSFP 800GbE with 1 × SFP28 management port, 2 × AC PSUs, x86 CPU [vendor-reported]. SN5600D and SN5610 appear in NVIDIA platform literature as the same 64 × 800G / 51.2T class; exact management/CPU/power distinctions were not fully captured [secondary]. No verified SN5800/SN6800 (air-cooled 51.2T) specifications captured; the only "SN6800" found is Dell's liquid-cooled CPO SN6800-LD — do not infer a conventional SN6800 [unverified].

**Juniper / HPE Juniper QFX.** QFX5241-64OD: 64 × 800G OSFP, 51.2 Tbps, 2U, shallow buffer, leaf/spine/end-of-row; QFX5241-64QD is the QSFP-DD twin at the same class [official] (juniper.net Junos release notes 25.2X100-D10). QFX5240-line breakout: 128 × 400G, 256 × 200G, 320 × 100G; 64-port versions use Tomahawk 5 with 165 MB buffer [official]. **QFX5140**: Broadcom Trident 5, 16 Tbps, 24 × 400G QSFP112 plus 8 × 800G OSFP800, 1RU — inference/storage and general leaf/spine rather than maximum-radix backend [official] (HPE QFX5140 datasheet). **QFX5250** (Tomahawk 6): 102.4 Tbps unidirectional / 204.8 Tbps bidirectional, 64 × 1.6T OSFP-RHS with breakout to 128 × 800G / 256 × 400G / 512 × 200/100G, 2OU, 200 MB unified buffer, 100%-liquid-cooled fanless variant, 3,375 W typical / 3,435 W maximum, 10 lpm coolant flow; air-cooled version also mentioned [official] (HPE QFX5250 document). **HPE Aruba gap**: no credible native Aruba CX-branded 800G switch found; contemporary material describes Aruba CX portfolio as spanning through 400GbE, and CX 10040 is not evidence of an 800G switch — it is a DPU/services-oriented platform focused on 100G server connectivity; do not invent "CX 10040 800G" [independent] (hpe.com Aruba CX page). HPE's 800G/1.6T data-center portfolio is now represented by HPE Juniper Networking QFX systems following HPE's Juniper integration — report the Aruba-CX-native gap honestly [independent].

**Edgecore (whitebox).** DCS560 / AIS800-64O: 64 × 800G OSFP800, Tomahawk 5, 51.2 Tbps, 2RU, SONiC/ONIE, up to 128 × 400G, targeting AI/HPC/RoCE and DCI; OSFP (AIS800-64O) and QSFP-DD800 (AIS800-64D) variants announced for OCP/Fyuz demos in October 2023 [independent] (thefastmode.com). Reported max system power for AIS800-64O is 2,775 W; reported per-cage optic budget varies by source between 24 W and 30 W — preserve as source-specific and flag the discrepancy rather than resolving it [secondary] (stordis.com product page). Edgecore's 2024 catalog corroborates the AIS800 positioning in its data-center lineup [vendor-reported].

**Micas (whitebox).** M2-W6940-64OC: 64 × 800GbE OSFP, Tomahawk 5, 2RU, 51.2 Tbps one-way / 102.4 Tbps full-duplex, targeting RoCEv2 and lossless Ethernet; announced globally shipping 2024-10-16 [official] (globenewswire.com). Micas also demonstrated a Tomahawk 5 Bailly CPO whitebox with 64 × 800G optical ports (2×FR4) and a claimed 30% lower power — vendor-reported until independently measured [official].

