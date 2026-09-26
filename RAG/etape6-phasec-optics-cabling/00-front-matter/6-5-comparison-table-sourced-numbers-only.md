---
id: etape6-phasec-optics-cabling/00-front-matter/6-5-comparison-table-sourced-numbers-only
title: "6.5 Comparison table (sourced numbers only)"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "AWS", "Broadcom", "EU", "Meta", "Microsoft", "Nvidia", "OpenAI", "Oracle", "United States", "xAI"]
dates: ["2026-01", "2026-02", "2026-03", "2026-06", "2026-07", "2026-08", "2026-09"]
keywords: ["accelerator", "amd", "aws", "blackwell", "compute", "cost", "ethernet", "gpu", "hyperscaler", "latency", "nvidia", "nvlink"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1099, 1153]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 2bf559a6d32da152e62080f4377008ab96611709192a699e94e703e28013f497
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

