---
id: etape6-phasec-optics-cabling/00-front-matter/3-7-standards
title: "3.7 Standards"
domain: front-matter
role: reference
task: reference
actors: ["AWS", "Alibaba", "Broadcom", "China", "Cohere", "Meta", "Microsoft", "Nvidia", "Oracle", "United States"]
dates: ["2021-06", "2023-04-05", "2025-03", "2025-03-25", "2026-03-13", "2026-05", "2026-06", "2026-07-14"]
keywords: ["capex", "cost", "cpo", "datacenter", "dsp", "ethernet", "gpu", "hyperscaler", "latency", "lpo", "npo", "nvidia"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [461, 489]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 900d1376fbd51dabc84abc5997d7a823f46a63eecb9367e1fc5c970fe6225ab6
---

# 3.7 Standards

- **LightCounting "Optics for AI" (Jan 2026)**: Ethernet optical transceiver + CPO market for AI clusters: $16.5B in 2025 → **$26B in 2026** (~60% growth both years); Meta and Oracle doubling 2026 capex; growth moderation expected 2027–2031 [independent: https://www.lightcounting.com/newsletter/en/january-2026-optics-for-ai-clusters-366]. Same $16.5B→$26B figures echoed by TrendForce (57.6% YoY) [secondary: https://compoundsemiconductor.net/article/124045/AI_optical_transceiver_market_to_reach_26b_in_2026].
- **LightCounting "Silicon Photonics, LPO/LRO and NPO/CPO" (May 2026)**: 2026 is the first year SiPh-modulator transceivers exceed 50% of the total **$40B** market; NPO/CPO deployments accelerating SiPh adoption; optical chips for transceivers/AOCs/LPO/LRO/NPO/CPO: $4B (2025) → ~$15B (2031, ~4×); total transceiver/AOC/NPO/CPO sales ~$80B by 2031 [independent: https://www.lightcounting.com/newsletter/en/may-2026-silicon-photonics-lpolro-and-npocpo-377].
- **AI cluster optics (scale-up + scale-out)**: from $5B (2024) to >$10B (2026); LPO and CPO adoption in scale-up networks accelerating 2026–2027, high volumes by 2028 [independent: LightCounting via EETimes https://www.eetimes.com/ai-demand-reshapes-optical-connectivity-and-photonics-roadmaps/].
- **LPO/CPO port share**: LPO/CPO ports to account for **>30% of total 800G and 1.6T ports deployed in 2026–2028** (LightCounting, Silicon Photonics report 9th ed.) [independent: https://www.lightwaveonline.com/home/article/55141192/lpo-msa-achieves-multi-vendor-interoperability]. **Conventional retimed pluggables "will continue to dominate the market for the next five years, and probably longer"** [independent, same].
- **LPO market size**: QYResearch forecast — $21.2M (2023) → $2B (2029), 113% CAGR; US+Canada $17.2M → >$1.1B by 2029 [secondary: https://www.lightwaveonline.com/home/article/14310705/linear-pluggable-optics-consortium-to-define-linear-pluggable-optics-lpo-specifications]. Key LPO manufacturers per QYResearch: Eoptolink, Coherent, CIG Tech, Zhongji Innolight, Hisense Broadband, HG Tech; top-5 revenue share ~39.8% in 2023 [secondary, same].
- **800G shipment volumes**: one analyst note cites LightCounting TAM — total transceivers $23.8B (2025); datacenter modules ~$22.8B (2026), of which 800G+1.6T ~$14.6B; 800G+ units ~24M (2025) → ~63M (2026). **Scope disagreement flagged**: MarketsandMarkets pegs DC transceivers at only $9.2B (2025) [secondary: https://github.com/roachx92/equity-watch/blob/HEAD/tickers/AAOI/reports/2026-07-14.md]. No clean DSP-vs-LPO unit split found — **gap flagged**.
- **800G LPO module market report (LP Information)**: covers Coherent, Cisco, Adtran, Lumentum, Innolight as key 800G LPO players; paywalled, no figures extracted [secondary: https://pdf.marketpublishers.com/lpinfo/global-800g-linear-pluggable-optics-module-market-lp.pdf].

### 3.7 Standards

- **LPO MSA**: 100G/lane spec released March 25, 2025 (covers 100/200/400/800G parallel SM links) [official: lpomsa.org press PDF]. 200G/lane spec in development with OIF/IEEE [official]. (See §3.4 for membership.)
- **OIF CEI-112G-Linear / CEI-224G-Linear**: OIF launched the CEI-112G-Linear project (June 2021) — "a linear chip-to-optical-engine interface needed to enable low-power, low-cost, small-form-factor 112G serial optical modules in CPO, NPO and server/GPU applications" [official: https://semiconductor-today.com/news_items/2021/jun/oif-010621.shtml]. CEI-224G Linear demonstrated at ECOC 2024 with AOI [official: AOI press release]. AOI's 1.6T OSFP DR8 LPO designed for OIF CEI-224G Linear [official].
- **OIF 3.2T Co-Packaged Module IA (OIF-Co-Packaging-3.2T-Module-01.0, April 5, 2023)**: industry's first CPO standard; 8×400G optical (FR4/DR4), 32×CEI-112G-XSR host interface (or 32×CEI-56G-XSR backward-compatible), ~140G/mm bandwidth edge density, enables 51.2T aggregate switch; CMIS-based management [official: https://www.oiforum.com/wp-content/uploads/OIF-Co-Packaging-3.2T-Module-01.0.pdf; https://www.businesswire.com/news/home/20230405005071/en/]. Companion: CPO Framework IA and ELSFP (external laser) project [official].
- **IEEE P802.3dj**: defines 200G/400G/800G/1.6T Ethernet at **200 Gb/s per lane** (200G PAM4 signaling for chip-to-chip, chip-to-module, backplane, copper cable, SMF); in Working Group Ballot; **on track for completion in late 2026** (Ethernet Alliance chair Peter Jones) [independent: https://www.networkworld.com/article/4113364/ethernet-groups-keep-2026-focus-on-higher-bandwidth-ai-demands.html; https://convergedigest.com/ieee-p802-3dj-moves-forward-with-200g-to-1-6t-ethernet-working-group-work/]. Early 200G/lane products expected during 2026. **400G/lane**: IEEE 802.3 chartered a 400 Gb/s-per-lane signaling Study Group on 2026-03-13 (up to 500 m SMF) — the path to 3.2T [secondary: https://github.com/jwt625/playground/blob/HEAD/20260521_ieee802.3dj/IEEE-802p3-ai-relevant-subpages.md]. IEEE also began an "802.3 Ethernet Interconnect for AI" assessment (Jan 2026) [secondary: networkworld].
- **Form factor / management**: OSFP MSA Rev 5.1 defines OSFP1600 (8×200G-class lanes); module management via OIF CMIS (current published CMIS 5.3) [secondary: https://roboticsandautomationnews.com/2026/09/22/osfp-modules-the-complete-guide-to-400g-800g-and-1-6t-optical-transceivers-for-ai-and-hyperscale-data-centers/104982/].

### 3.8 Caveats / gaps / conflicts (Wave 3)
1. **No independent lab-measured LPO latency/power figures found** — all numbers are vendor-reported or marketing; DSP ~100 ns vs LPO <1–10 ns spread is vendor claims.
2. **800G DSP module power**: 14–18 W (ascentoptics) vs 8–12 W (edgeoptic, weaker) vs Broadcom official sub-11 W with BCM85812. Report as a range with source tags, not a single figure.
3. **LPO reach**: official Eoptolink spec 500 m–2 km vs blog claim of ~50 m limit — flagged unreliable.
4. **CPO volume timing**: vendor claims (Micas "volume production" 2025; NVIDIA "ramping H2 2026") vs independent analysts (volumes 2028–2030) — unresolved.
5. **MaxLinear / Inphi, Hisense, Source Photonics, Sumitomo, Accelink LPO products**: no 2026 product names sourced — gaps.
6. **No DSP-vs-LPO 800G unit shipment split** found from LightCounting/Cignal AI/Dell'Oro/Yole in accessible sources.
7. **Hyperscaler NPO strategy claims** (Alibaba/Tencent/Meta/Microsoft/Amazon) rest on a single weak Chinese-market source — [unverified].
8. Semtech's "23–25 W retimed" figure is ambiguous (likely 1.6T-class) — flagged.
9. LPO MSA spec announcements (March 2025 press release vs June 2026 Converge Digest piece) may describe the same milestone — flagged as potential duplicate reporting.

---

