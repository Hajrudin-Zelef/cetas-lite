---
id: etape6-phasec-optics-cabling/00-front-matter/10-11-source-list-verbatim-urls
title: "10.11 Source list (verbatim URLs)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "CoreWeave", "Google", "Huawei", "Intel", "Lambda", "Meta", "Microsoft", "Nvidia", "Oracle", "xAI"]
dates: ["2023-01-09", "2023-04-27", "2024-10-16", "2024-10-28", "2025-06-04", "2025-10-08", "2025-10-14", "2025-10-29", "2026-02-20", "2026-09", "2026-09-16", "2026-09-22"]
keywords: ["acquisition", "asic", "compute", "cpo", "ethernet", "full-duplex", "intel", "latency", "lpo", "nvidia", "nvlink", "optics"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1605, 1645]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 5c63a3e9fb5fec89ef348f76f02803deba1ecd7dcd9d56f8ebaa44c38305754f
---

# 10.11 Source list (verbatim URLs)

- **SN5800**: no evidence of a real announced NVIDIA system by 2026-09-22 [unverified].
- **Juniper Express5 / Trio 800G data-center ASIC**: no fresh primary source captured; Juniper's documented 800G systems use merchant silicon (Tomahawk 5/6, Trident 5) [unverified].
- **Arista "7388X5"** as an 800G fixed system: not found [unverified].
- **Celestica 800G model** (e.g., DS-series 800G SKU with ASIC mapping): not found [unverified].
- **ZTE native 800G data-center switch**: not found [unverified].
- **Ruijie RG-S6990-64OC2XS**: reseller-only, no official datasheet; claimed 102.4T figure is full-duplex-implied and non-comparable [unverified].
- **Huawei 16800-X 288 × 800GE**: secondary claim only; official datasheet dual capacity figures (179/387 etc.) need accounting clarification; exact 800G line card unverified [unverified].
- **Cisco N9364F-SG3** details (cages, buffer, power, form factor, ship date): announced Feb 2026, awaiting official datasheet [unverified].
- **Edgecore per-cage optic budget**: 24 W vs 30 W across sources — unresolved [unverified].
- **Cisco N9364E-SG2 max power**: 2,270 W vs 2,779 W across official pages — unresolved [unverified].
- **Q3200-RA aggregate bandwidth**: 57.6 Tbps vs 2 × 28.8 Tbps across mirrored datasheets — inconsistent aggregation [unverified].
- **Tomahawk 6 production ramp scale** through 2026 (volumes, lead customers beyond announced OEM systems): not quantified in captured sources [unverified].
- **Arista 7800R4 line-card ASIC** (Jericho3-AI vs merchant deep-buffer silicon): datasheet excerpt lists SKUs but not the ASIC — do not assert a mapping [unverified].
- **Dell SN6800-LD ASIC identity**: sheet states 102.4 Tb/s modules without naming silicon — do not label "Spectrum-X successor" as fact [unverified].
- **Adtran LiteWave800 0.8W** and **Micas 30% CPO power saving**: vendor claims, not independently measured [unverified].
- **Meta/Microsoft/Google/Lambda/xAI (beyond Colossus) named 800G deployments** with models and quantities: not verified [unverified].

### 10.11 Source list (verbatim URLs)

nasdaq.com Broadcom Tomahawk 5 shipment; stocktitan.com AVGO Tomahawk 5; telecomlead.com Broadcom AI fabric; storagenewsletter.com 2025-06-04 Tomahawk 6; globenewswire.com 2025-10-08 Tomahawk 6 Davisson CPO; networkworld.com Arista 1.6T rack-scale; networkworld.com Broadcom Jericho3-AI 2023; nextplatform.com Jericho3-AI 2023-04-27; telecomtv.com Broadcom AI fabric; cisco.com N9364E-SG2-Q install guide; cisco.com N9364E-SP2R datasheet; storagenewsletter.com 2026-02-20 Cisco G300; aapnews.aap.com.au G300 syndication; cisco.com N9364E-SG2 datasheet; static-cisco.com N9364E-SG2X datasheet; marvell.com Teralynx 10 brief; convergedigest.com Marvell 51.2T shipments; servethehome.com Teralynx 10 preview 2023-03; networkworld.com networking roundup; convergedigest.com NVIDIA Spectrum-4 51.2T; NVIDIA Spectrum-X800 solution brief (rackcdn.com); NVIDIA Spectrum Ethernet datasheet (website-files.com); naddod.medium.com N9570-128QC Feb 2026; cdn.blueally.com Dell SN6800-LD spec sheet; hardwarenation.com SN5000 series datasheet mirror; cdn.dellshop.ru SN5600 datasheet mirror; nasdaq.com Intel exits network switching 2023-01; intel.com Tofino 6.4T retired; arista.com 7060X6 series; arista.com 7060X6 quick-look; sdxcentral.com Arista 1.6T race; arista.com 7800R4 AI-Spine datasheet; investors.arista.com 2025 800G announcement; delltechnologies.com Z9864F-ON spec sheet; dell.com UK Z-series; juniper.net Junos 25.2X100-D10 QFX5241; cdnhpe.getuniqcli.com QFX5240 brief; hpe.com QFX5140 datasheet; hpe.com QFX5250 document; hpe.com Aruba CX page; thefastmode.com Edgecore AIS800 Oct 2023; ftp.romsat.ua Edgecore 2024 catalog; stordis.com AIS800-64O page; globenewswire.com Micas M2-W6940-64OC 2024-10-16; globenewswire.com Micas Bailly CPO OCP 2024; ufispace.com AI solutions; eo-sgp-cos.ruijie.com RG-S6980-64QC datasheet 2026-06; network-switch.com Ruijie RG-S6990-64OC2XS; e.huawei.com CloudEngine 16800-X datasheet; huaweicentral.com 800GE core switch; FS.COM eu-en product 250955 N9600-64OD; fs.com uk product 250955; fs.com blog N-series bare-metal switches; resource.fs.com N9550-64D datasheet 2025-09; gpusmith.com Quantum-X800 datasheet mirror; highendcomputing.co.uk Q3401-RD datasheet; koicomputers.com Quantum-X800 PDF; FS.com blog 800G InfiniBand future; medium.com aicplight888 800G→2×400G cabling; globenewswire.com NVIDIA xAI Colossus 2024-10-28; investor.nvidia.com Colossus press release; sdxcentral.com xAI Colossus expansion; techradar.com Colossus 100k Hopper; github.com sungeuns foundation-model-engineering ch.6; oracle.com AI World OCI roundup; convergedigest.com Oracle Acceleron RoCE; stocktitan.net ORCL Zettascale10; nationalcioreview.com Oracle Q1 FY2027; hpcwire.com CoreWeave Vera Rubin NVL72 2026-09-16; investor.nvidia.com Spectrum-XGS 2025; ascentoptics.com 800G power consumption; thefastmode.com Adtran LiteWave800 Mar 2026; thefastmode.com FS 800G LPO launch; aithority.com FS 800G LPO; charratcommunication.fr Fiji LPO pluggable PDF (Jan 2026); gpusmith.com SN5600 datasheet mirror; datacentrenews.in xAI Colossus; financialcontent.com bizwire Arista 2025-10-29 announcement.

---

## Wave 11 — 800G NICs/DPUs and AI-Cluster Deployments

*Research date: 2026-09-22. Single-writer wave. All factual lines carry one evidence tag: [official] = primary vendor/standards source, [vendor-reported] = vendor claim via press release/partner, [independent] = independent press/expert analysis, [secondary] = press/reseller/aggregator/resale data, [unverified] = single-source or uncorroborated claim.*

### 11.0 Method notes and headline corrections

- This wave covers two tracks: Part A (800G NIC/DPU hardware) and Part B (AI-cluster deployments).
- Correction: Broadcom "Thor 2" is NOT an 800G NIC — Thor 2 (BCM57608) is a 400G PCIe Gen5 adapter; Broadcom's 800G AI NIC is **Thor Ultra**, launched October 14, 2025 and sampling to select customers as of September 2026 [independent].
- Correction: Broadcom "Stingray PS1100R" is NOT a relevant 800G DPU — PS1100R is a 100GbE storage adapter (2020 vintage), not a 400G or 800G DPU [official].
- Correction: Intel E830 is NOT 800G — the Intel E830 family ceiling found is 200GbE [official].
- Correction: the "9,216 MPO cables" figure for a GB200 NVL72 scalable unit is misstated — the cited vendor source says 9,216 active fiber strands (4,608 compute-fabric strands), not 9,216 MPO cables; internal NVLink fabric is passive copper with zero optical fiber [vendor-reported].
- Fungible is NOT an independent 800G NIC/DPU supplier: Microsoft announced the acquisition on January 9, 2023 [vendor-reported].
- Marvell: no verified public merchant 800G NIC/DPU product was found; Marvell's public AI products in this space are optical modules, DSPs, and electro-optics, not host NICs [independent].
- Pricing figures from FS.com, NADDOD, and resale forums are channel/retail list prices, not NVIDIA list prices, and typically exclude optics, cables, support, tax, and volume discounts; they are not comparable across vendors without BOM normalization [secondary].
- Vendor performance claims (throughput, efficiency, latency) are presented as vendor claims, not independent measurements, unless otherwise noted [independent].

### 11.1 NVIDIA ConnectX-8 SuperNIC (800G, shipping)

