---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-round-4-2026-09-22
title: "Supplementary / Complementary Research Pass — round 4 (2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: ["Cohere", "Intel", "Nvidia", "TSMC"]
dates: ["2026-02", "2026-05", "2026-08-26", "2026-09-22"]
keywords: ["research", "agentic", "alignment", "attribution", "cpo", "distribution", "dsp", "ethernet", "intel", "latency", "nvidia", "optics"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1098, 1151]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 186511a608ff4a3360317a3340c6ef7da6efeb29622ee9502ee340326671f0c8
---

# Supplementary / Complementary Research Pass — round 4 (2026-09-22)

## Supplementary / Complementary Research Pass — round 4 (2026-09-22)

**Scope note:** A fourth additive pass, independently researched the same day. It contains only material not present in the base §§1–8, Supplementary Pass #1 (§§A–M), Pass #2 (§§N–X), the round-2 pass (second §§N–S), or the three round-3 passes (§§R3-A–R3-H, §§AA–II, §§Y–AG). Earlier sections were left untouched. Read the base report and prior passes first; the provenance legend from the base applies here.

### R4-A. NVIDIA Q2 FY2027 earnings (Aug 26, 2026) — networking callouts and quarterly sequence correction

Round-2 §O cited a "$11B networking quarter (+263% YoY)" as coming from the Q2 FY2027 earnings call. The call highlights published by Converge Digest paint a different, newer picture, so the quarterly sequence is corrected here [vendor-reported — earnings call; secondary — Converge Digest]:

- **Q2 FY2027** (ended Jul 26, 2026; reported 2026-08-26): total revenue **$96.2B (+106% YoY, +18% sequential)**; Data Center **$89.0B (+117% YoY)** — ~92% of company revenue; GAAP gross margin 75.0% [vendor-reported — NVIDIA 8-K/exhibit 99.1].
- **Networking revenue "reached another record and grew 18% sequentially"; Spectrum-X Ethernet revenue grew 2.6× YoY** [vendor-reported — earnings call via Converge Digest, 2026-08-26]. No absolute Q2 FY2027 networking dollar figure was located in the call highlights — do not infer one from the Q1 FY2027 $14.8B figure in base §2.1 [gap].
- Forward context from the same call: customer forecasts indicate demand could support ~100% FY2028 revenue growth, but NVIDIA expects ~70% because of supply constraints; supply is expected to remain a bottleneck **at least through the end of FY2028**; Vera Rubin expected to account for ~20% of Data Center revenue in Q3 FY2027 [vendor-reported — earnings call].
- **Correction flag:** the "$11B / +263%" figure cited in §O belongs to the **Q4 FY2026** print (base §2.1's quarterly sequence: $4.9B Q2 FY26 → $11B Q4 FY26 → $14.8B Q1 FY27 → Q2 FY27 record +18% sequential). The §O attribution to the Q2 FY2027 call is therefore superseded [secondary].

### R4-B. Cisco Cloud Control — Cisco Live 2026 unified management plane (new angle)

Not covered in any prior pass; this is the 2026 convergence story for Cisco's management surface, adjacent to Meraki cloud management (base §4):

- At Cisco Live 2026, Jeetu Patel (President & CPO) stated: **"Every Cisco product you know will be managed from Cisco Cloud Control"** — Catalyst, Meraki, Nexus, ACI, security, collaboration, Splunk under one management plane, with agentic AI runbooks that execute fixes, not just recommend them [secondary — YouTube/Tech Updates summary of Cisco Live 2026].
- **Cisco IQ** (AI support layer): 2,036 customers onboarded vs 800 expected; 88% first-try support routing; included in support contracts customers already pay for [secondary].
- Caveats: pricing and migration timelines for Cloud Control were **not announced** at the event — same dashboard-promise caution noted for DNA Center/Intersight/Meraki in earlier eras [secondary]. This is a strategy announcement, not a shipping milestone [unverified].

### R4-C. Cisco C9550 Fixed Core Smart Switches — orderable May 2026 (new hardware)

- **Cisco C9550 Series Fixed Core Smart Switches — orderable since May 2026**, powered by **Cisco Silicon One**: up to **four 400G uplinks, 6.4 Tbps switching capacity, 3.9 Bpps**; positioned for campus cores that are becoming "machine communication fabrics" for AI traffic; **quantum-safe**; managed through Cisco Cloud Control [secondary — LinkedIn/Cisco Live 2026 recap]. No Cisco list price located [gap].
- Note: distinct from the C9350 access smart switches (§AA) and the C9610 modular chassis (§AA/base §4) — the C9550 is the fixed-form **core** tier [secondary].

### R4-D. Cisco Live 2026 campus switching updates (new detail)

- **Catalyst 9300X and 9500X platform updates**: new line cards and software updates, **improved MACsec encryption performance**, enhanced **SRv6** (Segment Routing over IPv6) support in campus environments, **new 100G uplink options** for high-density distribution layers; IOS-XE updates with AI-driven assurance engine improvements and telemetry streaming for third-party monitoring integration [secondary — thenetworkdna.com, Cisco Live 2026 recap].
- **Meraki-adjacent announcements** at the same event: new **MX security-appliance models** with improved throughput and built-in SD-WAN; **Meraki Go** enhancements for very-small-business; deeper Meraki↔Cisco Security Cloud policy integration across Catalyst and Meraki-managed environments [secondary].
- Wi-Fi 7 APs manageable through **both Catalyst Center and Meraki Dashboard** — another data point on the enterprise/SMB management convergence [secondary].

### R4-E. Catalyst 9300-M cloud-managed switches — 15 new models (Feb 2026, Cisco Live Amsterdam)

- Cisco's February 2026 announcement: **15 new Catalyst 9300-M models**, combining Catalyst 9300 hardware with Meraki cloud management; **90W Cisco UPOE+** options, high-density multigigabit, fiber options, and **up to 1TB physical stacking**; L7 application visibility, Adaptive Policy, and Meraki-dashboard troubleshooting built in [official — Cisco Blogs, "Delivering Simplified Operations: The Latest from Cisco Live Amsterdam"].
- Positioning: 9300-M is the access-tier counterpart to the C9350 smart-switch line — cloud-native Catalyst with full Catalyst hardware capability [official].

### R4-F. Spectrum-X Photonics ecosystem detail (GTC 2025 announcement — partner scope)

Extends base §2.3 and Pass #3 §EE (CPO technical specs) with the vendor-ecosystem scope, which prior passes did not enumerate:

- NVIDIA's photonics ecosystem partners: **TSMC, Browave, Coherent, Corning Incorporated, Fabrinet, Foxconn, Lumentum, SENKO, SPIL, Sumitomo Electric Industries, TFC Communication** [official — NVIDIA press release, investor.nvidia.com].
- TSMC COUPE (Compact Universal Photonic Engine) alignment: first generation stacks a 65nm EIC with a photonic IC via TSMC SoIC-X packaging [official].
- Claimed economics: **9 W/port with CPO vs ~30 W/port with pluggable DSP optics**; resiliency **10×**, signal integrity **64×**, deployment **~30% faster** due to fewer components [vendor-reported — NVIDIA via tech press].
- Note: pluggable transceiver partners (Coherent, Eoptolink, Fabrinet, Innolight) continue to be named alongside photonics — NVIDIA is running CPO and pluggable in parallel, not as a replacement play [official].

### R4-G. Dell Z9432F-ON spec depth + distribution evidence (new)

- Full spec-sheet detail for **Z9432F-ON** (1RU, Trident4-X11): **64→32× 400G QSFP56-DD** (2× 10G SFP+), 12.8 Tbps, 5.2 Bpps, 132 MB packet buffer, **<850 ns latency**, programmable NPL pipeline; Intel **Denverton C3758 8-core @ 2.2 GHz**; **32 GB DDR4 ECC**; 156K MAC; **400K IPv4 ALPM routes / 300K IPv6**; max power **1,404 W** (typical 900 W); airflow options PSU-to-IO and IO-to-PSU [official — Dell spec sheet, delltechnologies.com].
- Distribution: third-party resellers (e.g., hssl.us) carry the Z9432F-ON as **quote-only** with 2–5+ week ETA — consistent with the base finding that Dell publishes no list prices for Z-series switches [secondary].
- Current 2026 line-up confirmation: Dell's data-center-switch store page (dell.com) lists Z9864F-ON, Z9664F-ON, Z9432F-ON, and the S5448F/S5232F/S5296F/S5248F/S5224F/S5212F families as the current PowerSwitch range [official — Dell store].

### R4-H. Round-4 verification log (new open items)

