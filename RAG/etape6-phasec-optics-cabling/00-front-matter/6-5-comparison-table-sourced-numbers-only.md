---
id: etape6-phasec-optics-cabling/00-front-matter/6-5-comparison-table-sourced-numbers-only
title: "6.5 Comparison table (sourced numbers only)"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "AWS", "Broadcom", "EU", "Meta", "Microsoft", "Nvidia", "OpenAI", "Oracle", "United States", "xAI"]
dates: ["2025-04", "2025-12", "2026-01", "2026-02", "2026-03", "2026-04", "2026-06", "2026-07", "2026-08", "2026-09"]
keywords: ["accelerator", "amd", "aws", "blackwell", "compute", "cost", "dsp", "ethernet", "gpu", "hyperscaler", "latency", "lpo"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1099, 1179]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 3af133d1cc5fae95fd16f1145ea58626afd766842a6e7d1ccf18bf77fa47dd46
---

# 6.5 Comparison table (sourced numbers only)

### 6.5 Comparison table (sourced numbers only)

| Metric | 100G (QSFP28, 4×25G NRZ) | 400G (QSFP-DD/OSFP, 8×50G PAM4) | 800G (OSFP/QSFP-DD, 8×100G PAM4) |
|---|---|---|---|
| Passive DAC max length | 3–5 m (typ. 3 m; 5 m vendor-stretch) | 3 m (QSFP-DD); 2 m (OSFP flat-top); 1.5 m (QSFP112) — FS.com | 2–3 m (breakout); ≤1 m straight OSFP per FS.com listings |
| Passive DAC power | ≤0.5 W (FS.com 100G listing; typ. ~0.1 W) | ≤0.1 W | ≤0.1 W (FS.com; ≤0.005 W on 0.5 m QSFP-DD listing) |
| Passive DAC price (single-unit web list) | US$39–40 (3 m/1 m QSFPTEK/FS.com/ROBOfiber) | US$209 (3 m, Optcore Cisco-compat) | US$161–199 (0.5–1 m, FS.com); US$299 (2 m breakout, FS.com) |
| AOC length range | 1–100 m (OM3/OM4) | 1–100 m (OM4; 70 m OM3 per 200G SR4 data) | sold 1 m+ (FS.com MX listing); same 100 m-class MMF capability [secondary] |
| AOC power | <2.2–2.5 W/end | <8 W/end (straight); <3.5 W/100G-end, 12 W/400G-end (breakout) | ≤14 W (FS.com MX 800G OSFP AOC) |
| AOC price (single-unit web list) | US$204 (3 m, FS.com) | US$705 (3 m, ROBOfiber) – US$1,011 (3 m, FS.com Cisco-compat) | MXN$48,135 (1 m, FS.com MX) |
| AEC length range | n/a (AEC category starts ~200G) | 3/5/7 m (Credo LP SPAN Gen2) | up to 7 m (Credo ZeroFlap; Amphenol 3/5/7 m) |
| AEC power | n/a | 4.5 W/end (Credo LP SPAN Gen2) | <9 W (QSFPTEK) – 12 W/10 W (FS.com EU 800G/400G ends) |
| AEC price (single-unit web list) | n/a | SGD 683.43 (FS.com SG 400G *active copper* 3 m — ACC-class, not full AEC; flag) | US$1,664 (1 m QSFPTEK); €1,568 (4 m FS.com EU); US$1,870–2,504 (3/5/7 m Amphenol) |

**Explicit non-comparability flags:** (a) 400G passive vs active prices mix vendors and currencies (SGD vs USD); (b) FS.com US vs EU vs MX vs SG storefront prices differ by region/tax; (c) AOC prices vary with OEM-coding (Cisco-coded SKUs carry premiums); (d) single-unit web prices ≠ hyperscaler volume pricing — no volume pricing was sourced [unverified-resolution noted].

### 6.6 AI cluster relevance

**Scale-up vs scale-out.** Frontier AI clusters split networking into (1) *scale-up* — GPU-to-GPU inside a node/rack (NVIDIA NVLink: 1.8 TB/s bidirectional per GPU on Blackwell, ns latency; closed/proprietary), and (2) *scale-out* — node-to-node across thousands of racks via RDMA: InfiniBand (NVIDIA Quantum) or RoCEv2 Ethernet (Spectrum-X), at 400G/800G per port [secondary] (hczhu/learning-notes GPU connectivity, June 2026; sungeuns/foundation-model-engineering, September 2026).

**Where copper (DAC/AEC) sits.** In-rack and rack-to-ToR links (<3–5 m) are predominantly copper — passive DAC where reach allows, AEC where 112G-lane signal integrity, thinness, or 5–7 m reach demands it. The SONiC community's Ethernet scale-up architecture doc (updated September 2026) states: "Intra-rack links (<3m) primarily use copper; optical is reserved for high-density or future 1.6T+ deployments" and notes copper DAC/AEC reach "<5m at 400G" [secondary] (sonic-net/sonic Ethernet Scale-Up AI Cluster Architecture doc).

**NVIDIA DGX/HGX.** DGX SuperPOD reference architectures (H100/H200) use NVIDIA Quantum QM9700 NDR 400G InfiniBand for compute and storage fabrics (8× NDR400 per system), with DAC/AOC as the physical media options on those IB links [official] (docs.nvidia.com DGX H100/H200 SuperPOD reference architectures). NVIDIA's own 800G OSFP NDR DAC SKUs (MCP7Y00-Nxxx, 1–3 m, 26–30AWG) are built for Quantum-2 switch ↔ ConnectX-7 adapter links [vendor-reported].

**Why AEC matters for AI backends:** (1) 800G passive DACs are thick/rigid (26AWG) and short — airflow and routing pain in dense GPU racks; AECs use 32–34AWG ("route like Cat6") [vendor-reported]; (2) lossless RDMA backends are sensitive to "soft link flaps" — Credo's ZeroFlap AECs target zero-flap operation, endorsed by xAI for 100k+ GPU builds [official]; (3) cost/power: AEC ≈ half the power of AOC-class optics and up to $1,000/GPU claimed saving [vendor-reported]; (4) rack-scale designs (e.g., GB200 NVL72 with NVLink scale-up domains) push scale-out NIC/switch links to 800G, where AEC is the copper vehicle [secondary].

**Meta/OCP.** Meta's OCP contributions focus on optics for AI clusters (2×400G FR4 LITE for ≤500 m, 400G DR4 OSFP for NIC-to-switch, OCP Summit 2025) and co-founded ESUN (Ethernet for Scale-Up Networking, OCP Summit 2025, with AMD/Arista/Broadcom/Cisco/HPE/Marvell/Microsoft/NVIDIA/OpenAI/Oracle) to standardize Ethernet for scale-up accelerator interconnects [independent] (convergedigest.com, Meta OCP Summit 2025 coverage). Meta also backs the Optical Scale-Up (OCI) MSA for future optical scale-up domains beyond copper backplanes [independent] (lightwaveonline.com, OFC 2026 OCI consortium coverage, March 2026). No source in this research documents Meta deploying AECs specifically — [unverified].

**Market momentum.** Astera Labs' Taurus Ethernet smart-cable-module revenue grew >4× in 2025 as datacenters moved 400G→800G, with a multi-year Amazon/AWS agreement (up to $6.5B purchases over 7 years, warrants for 3.3M shares) tied to Trainium/Inferentia connectivity [secondary/financial press, February 2026]; 650 Group forecasts AECs "will quickly replace direct attached copper" at hyperscalers [independent-analyst-via-vendor-PR].

### 6.7 Vendor pricing (representative SKUs — verbatim, with dates/URLs)

**FS.com US (fs.com, category crawls ~July 2026):**
- 100G QSFP28 4×25G NRZ passive DAC 1 m — US$40.00 (21.5K sold) — fs.com 100g-qsfp28-dac category.
- 100G QSFP28 4×25G NRZ passive DAC breakout 3 m (QSFP28→4×SFP28) — US$136.00.
- 100G QSFP28 AOC 3 m — US$204.00 (<2.2 W) — same page.
- 800G QSFP-DD passive DAC 0.5 m — US$199.00; 800G OSFP passive DAC 0.5 m — US$161.00; 800G OSFP passive 1 m — US$186.00 — fs.com 800g-dac-aoc category.
- 800G OSFP→2×400G OSFP passive breakout 2 m — US$299.00; →2×400G QSFP112 2 m — US$299.00; →4×200G OSFP 2 m — US$461.00; →4×200G QSFP112 2 m — US$324.00.
- 800G OSFP→2×400G QSFP112 *active* breakout DAC 3 m — US$1,136.00 (≤1.5W/≤0.6W).
- 800G OSFP→4×200G QSFP112 *active* breakout DAC 5 m — US$1,536.00 (≤1.7W/≤0.3W).
- 400G QSFP-DD AOC 3 m, Cisco-compat (P/N QDD-400G-AO03, SKU 146374, Broadcom chip, <8 W/end) — US$1,011.00 — fs.com product 146374 (crawled ~January 2026).
- 400G QSFP-DD active copper 3 m, Arista-compat (P/N QDD-400G-AC03, SKU 177377, Macom chip, 2.5 W) — SGD 683.43 GST incl. (FS.com Singapore) — fs.com/sg product 177377.
- 400G QSFP-DD→4×100G QSFP56 AOC breakout 3 m (P/N QDD-400G-4QAO03, SKU 150593, Inphi chip) — AUD 2,504.70 GST incl. (FS.com Australia) — fs.com/au product 150593.
- 800G OSFP→2×400G OSFP AEC 4 m (P/N OSFP-800G-2OFLAE04, SKU 312079, retimer, 12 W/10 W) — €1,568.00 VAT excl. (€1,865.92 incl.) (FS.com Europe) — fs.com/eu-en product 312079.
- 800G OSFP AEC 1 m — MXN$21,151 (≤12 W); 800G OSFP AOC 1 m — MXN$48,135 (≤14 W) (FS.com Mexico) — fs.com/mx 800g-1.6t-osfp-qsfp-dd category.

**QSFPTEK (qsfptek.com, September 2026):**
- Cisco QSFP-100G-CU3M-compat 100G QSFP28 passive DAC 3 m (QT-Q28-PC3, 30AWG) — US$39.00, 471 sold, stock dated 12 September 2026 — qsfptek.com product 99242.
- Generic 800G OSFP→2×400G OSFP AEC breakout 1 m (Product No. 103796, <9 W, post-FEC BER <1e-15) — US$1,664.10, stock dated 16 September 2026 — qsfptek.com product 103796.

**Cables on Demand (Amphenol AECs, in stock ~August 2026):**
- Amphenol SF-NND1JA0303-003M: 800G OSFP→2×400G OSFP AEC 3 m, 32AWG — $1,869.93 (1–50 qty) — cablesondemand.com QSFP-DD/OSFP cables page.
- Amphenol SF-NND1JE0305-005M: 5 m, 30AWG — $2,131.02.
- Amphenol SF-NND1JH0317-007M: 7 m, 28AWG — $2,503.51.

**Others:**
- Optcore Cisco QDD-400-CU3M-compat 400G QSFP-DD passive DAC 3 m, 26AWG — US$209.00 excl. VAT — optcore.net QDD-400G-DAC-P3M (page updated ~700 days ago; still live ~April 2026 crawl).
- ROBOfiber/DataInterfaces (US, crawled ~September 2026): 100G QSFP28 passive DAC 1 m $40.00; 100G→4×SFP28 breakout 3 m $99.00; 400G QSFP-DD AOC 3 m $705.00; 400G→4×QSFP56 passive DAC 3 m $330.00; 100G→4×SFP28 AOC breakout 3 m $201.00 — datainterfaces.com.

**Volume-price context note:** 800G DR8 optical transceivers at volume (1000+ units) list $1,000–1,400, LPO $700–900 [secondary] (saastisfy.fr price list, April 2025) — context only; transceivers are not DAC/AOC/AEC and are included solely to frame the "AEC ≈ half of optics" claim.

### 6.8 Reliability / compatibility

**EEPROM coding.** Every DAC/AOC/AEC carries an EEPROM (SFF-8636 for QSFP28-class; CMIS for QSFP-DD/OSFP-class) readable over I2C, programmed with vendor name, part number, length, and compliance codes [vendor-reported] (Dawnray, FS.com). Third-party vendors (FS.com, QSFPTEK, Optcore) operate in-house coding labs that program cables to mimic OEM identities (Cisco, Arista, Juniper, NVIDIA/Mellanox, Dell EMC…); FS.com: "Our in-house coding facility programs all of our parts to standard OEM specs for compatibility" [vendor-reported] (FS datasheet 20240428120743z6yjda.pdf). **Vendor lock-in is real:** "Some switches enforce vendor coding (lock-in)" — always confirm vendor-approved/coded cables for critical deployments [secondary] (network-switch.com). AECs additionally run a microcontroller implementing CMIS (Microchip META-DX2C SDK implements CMIS 5.2; Credo CLOS AEC uses CMIS) — AECs present a much richer management interface (telemetry, diagnostics, firmware) than passive DACs [official] (Microchip; Credo CLOS brief).

**Failure modes (by type).** Passive DAC: mechanical — conductor fatigue from tight bends (respect min. bend radius: 33–72 mm depending on AWG/length), connector mating-cycle wear, crosstalk/insertion-loss degradation. Thick 26AWG bundles also create airflow/thermal issues in dense racks (a deployment-level "failure" via overheating neighbors) [secondary]. AOC: VCSEL/laser wear-out and fiber-connector contamination/damage; no re-termination — one bad end scraps the whole assembly [secondary]. AEC: active-silicon failure modes (retimer chip, PMIC, thermal) — but far fewer components than an optical link. Credo claims up to 100M hours MTBF and "100 times better reliability" than optical solutions for its 800G AECs [vendor-reported] — **flag: vendor marketing claim, not independently verified.** Link-flap sensitivity in AI backends: Credo's ZeroFlap family specifically targets "zero soft link flaps" for lossless RDMA AI fabrics [official].

**BER considerations at 100G PAM4 lanes.** IEEE 802.3 requires pre-FEC BER ≤ 2.4×10⁻⁴ for 100G-PAM4-class clauses (e.g., 100GBASE-CR2/KP4), with KP4 RS(544,514) Reed-Solomon FEC mandatory on 400G and most 100G-PAM4 electrical/optical interfaces, delivering post-FEC BER <10⁻¹²–10⁻¹⁵ [official-via-secondary] (Tektronix PAM4 primer; EDN 400G FEC coverage). Raw (pre-FEC) BER on PAM4 links "can easily reach 10⁻⁶ to 10⁻⁴" vs the traditional 10⁻¹² NRZ target; RS-FEC provides ~7–8 dB coding gain [secondary] (fibermall.com FEC explainer, December 2025). Cable-vendor BER claims: 400G passive DAC "BER better than 1E-15" (Vchung) [vendor-reported]; 800G AEC "BER (Post-FEC) <10⁻¹⁵" (QSFPTEK) [vendor-reported]; Credo retimer/DSP cables are specified against KP4 FEC thresholds (optional inner Hamming (128,120) FEC in some DSPs) [official] (Credo Bluebird brief). Practical implication: at 100G/lane PAM4, DAC/AEC link budgets assume FEC is ON; a passive DAC that meets pre-FEC 2.4e-4 at its rated length/loss will deliver effectively error-free post-FEC operation. Switch ASICs (Broadcom Tomahawk/Jericho, NVIDIA Spectrum, Cisco Silicon One) terminate FEC in hardware; disabling FEC breaks auto-negotiation on standards-compliant ports [secondary]. AEC retimers regenerate the signal and can terminate/monitor FEC, which is why AECs hold BER margin at 5–7 m where passive copper's pre-FEC BER would exceed the KP4 correction threshold (mechanism per vendor DSP descriptions; specific threshold margins per cable are vendor-proprietary — [unverified] at per-SKU level).

### 6.9 Flags & conflicts summary (Wave 6)
1. **Credo "Dove" chip** — could not verify; do not use. **Seagull** verified only as optical DSP, not AEC retimer.
2. **100G QSFP28 passive max length** — 3 m standard vs 5 m vendor-stretch (Optcore); unresolved.
3. **ACC chip placement** — Rx-end only (Fibermall) vs both ends (VEEX); unresolved, likely vendor-dependent.
4. **800G "DAC hits the wall" (Credo 2021)** vs 2026 reality of widely sold 800G passive DACs (FS.com, QSFPTEK) — vendor positioning vs market fact.
5. **AEC retimer latency** — no cable-level ns figure sourced; ~100 ns/hop is industry lore [unverified]; only verified figure is Credo Bluebird *optical* DSP <40 ns/direction.
6. **All prices** are single-unit web list prices (crawled Jul–Sept 2026), region/currency/SKU-dependent, and not comparable to hyperscaler volume pricing; several comparisons above are explicitly directional.
7. **Credo reliability/power/cost claims** (100M-hr MTBF, 100× reliability, half-power vs optics, $1,000/GPU saving) are vendor-reported marketing figures without independent verification.
8. **10G/25G passive DAC lengths and TE/Molex AEC SKUs** were not directly sourced — omitted or flagged rather than invented.

---

