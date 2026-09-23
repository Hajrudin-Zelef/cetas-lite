---
id: etape6-phasea-vendors-dc/00-front-matter/nn-meraki-ms390-street-pricing-new-complements-round-3-aa
title: "NN. Meraki MS390 street pricing (new — complements round-3 §AA)"
domain: front-matter
role: reference
task: pricing
actors: ["Nvidia", "United States"]
dates: ["2026-05", "2026-07", "2026-08", "2026-08-26", "2026-09", "2026-09-22", "2026-23-10"]
keywords: ["pricing", "compute", "cyber", "distribution", "ethernet", "inference", "license", "neocloud", "nvidia", "research", "revenue", "rubin"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1233, 1311]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: b5d29d8ca09a4fb096c881c2ed5e2158398f105a1027ec33ff5c9164829ef400
---

# NN. Meraki MS390 street pricing (new — complements round-3 §AA)

### NN. Meraki MS390 street pricing (new — complements round-3 §AA)

Earlier passes priced MS150 (§R), MS355/MS450 (§AA); the flagship access-tier **MS390** was missing:

- **MS390-48-HW**: $10,065.48 list / **$7,549.11** street (hummingbirdnetworks, 2026) [secondary]; MS390-48U-HW listed at $10,195.99 (a-power.com) [secondary].
- **MS390-24UX-HW**: $12,716.99 (a-power.com, 2026) [secondary].
- UK: **MS390-24P-HW £3,679.99** (pricerunner.com aggregator, 2026) [secondary].
- India: MS390-48 gray-market listing ₹45,000–₹1,25,000 (indiamart, 2026) — channel artifact, not comparable [secondary].
- License anchors:
  - **LIC-MS390-24E-7Y** (Enterprise 7-yr): MSRP **$3,645.72**, street **$2,311.00**; instant rebate $584 expiring 10/23/2026 (Tech-America) [secondary].
  - **LIC-MS390-48A-1Y** (Advanced 1-yr): MSRP **$2,923.28**, street **$1,853.00** (Tech-America) [secondary].
  - **LIC-MS390-48A-3Y** (Advanced 3-yr): list **$6,576.94** (unisolinternational; request-quote) [secondary].
  - **LIC-MS390-24A-3Y**: **$1,751.90** (hssl.us) [secondary].
- Pattern: ~37–45% discount-to-list at US resellers; consistent with the MS355/MS450 discount pattern in round-3 §AA [secondary].

### OO. NVIDIA networking revenue — Q2 FY2027 reconciliation (corrects round-2 §O)

Round-2 §O cited "networking sales at $11B, +263% YoY" from a July 2026 Spectrum-6 press item. That figure is **stale** — it matches the **Q4 FY2026** print, not the quarter current in July–September 2026 [flag]:

- **Q1 FY2027** (reported ~May 2026): networking revenue **~$14.8B, +199% YoY** (base §2.1) [vendor-reported].
- **Q2 FY2027** (reported **August 26, 2026**, SEC 8-K): total revenue **$96.221B (+106% YoY)**; Data Center **$89.0B (+117% YoY)** (Hyperscale $48.710B, AI Clouds/Industrial/Enterprise $40.313B); Compute & Networking reportable segment **$88.299B (+114% YoY)**; Q3 FY2027 outlook **$108.0B ±2%** [official — SEC filing 0001045810-26-000073; tickergrove.com].
- Earnings-call highlights (Aug 26, 2026): "**networking revenue reached another record and grew 18% sequentially**; **Spectrum-X Ethernet revenue grew 2.6× YoY**" [secondary — convergedigest.com; alphastreet transcript].
- Implied math: $14.8B × 1.18 ≈ **~$17.4B** networking revenue in Q2 FY2027 — an analyst inference, **not a disclosed figure**; NVIDIA does not separately disclose networking revenue as a line item [gap]. Do not quote ~$17.4B as fact.
- Other Q2 FY2027 call context: top-5 hyperscalers expected to spend ~$800B in 2026 / $1.3T in 2027; neocloud partners ~8 GW installed by end-2026 (vs ~3 GW end-2025); Vera Rubin production shipments began August 2026, expected to be NVIDIA's fastest ramp [secondary].
- Consequence for earlier passes: replace the "$11B networking quarter" reading with the sequential-growth framing above; round-3 §HH item 6 noted the same caution.

### PP. Dell campus N-series — 2026 status (no refresh located)

- **No 2026 hardware refresh of the Dell N-series was located** in 2026 sources; the current lines remain **N2200-ON** (access, up to 48× 1/2.5G RJ45, OS6) and **N3200-ON** (multigigabit 1/2.5/5/10G, OS6 + OS10 options), with **N4000** (aggregation/core) still referenced in reseller/educational materials [secondary — etb-tech.com, Sep 2026; official Dell support pages].
- Dell's official positioning: N3200-ON = "power-efficient and resilient 1GbE and 1/2.5/5/10GbE multigigabit switching solution for advanced Layer 3 distribution for offices and campus networks" [official — dell.com support].
- External PSU options (MPS-1S/MPS-3S) documented for N2200-ON/N3200-ON/E3200-ON [official — Dell installation PDFs].
- Treat 2026 continuation as assumed [unverified]; no 2026 campus-switch announcement from Dell was located (contrast with HPE's NRF 2026 CX 6000 additions in the base timeline).

### QQ. Round-4 verification log (new open items)

1. UFM Enterprise / UFM Cyber-AI / UFM Telemetry exact 2026 release versions and licensing model — the base cites 6.26.1; this pass covers only the security bulletin [gap].
2. NVIDIA networking revenue as a disclosed line item — not disclosed; the ~$17.4B Q2 FY2027 figure is an inference from "+18% sequential" and must not be quoted as fact [flag].
3. DOCA SDK 3.x vs DPF v26.x version correlation — not documented; do not conflate [gap].
4. AOS-CX 10.18.0001/10.18.1002 branch maturity (10.18 is the newest branch in the Sep 2026 bulletin) — feature content not reviewed in this pass [gap].
5. MS390 UK reseller prices vs US list — channel-dependent, not directly comparable [secondary].
6. Dell N-series 2026 refresh — none located; continuation assumed [unverified].
7. Round-2 §O "$11B networking quarter" — superseded by this pass (§OO); flagged as stale press reuse of the Q4 FY26 figure.

### RR. Round-4 sources (verbatim URLs)

- https://securityonline.info/nvidia-security-updates-ufm-dgx/
- https://www.techPowerUp.com/268829/nvidia-unveils-ai-platform-to-minimize-downtime-in-supercomputing-data-centers
- https://docs.nvidia.com/doca/sdk/general-support/index.html
- https://docs.nvidia.com/doca/archive/3-2-0/doca-gpunetio/index.html
- https://catalog.ngc.nvidia.com/orgs/nvidia/doca/resources/doca_hbn/3.4.0/version-history
- https://docs.nvidia.com/doca/sdk/snap-virtio-fs-service-release-notes/index.html
- https://network.nvidia.com/pdf/eol/LCR-000453.pdf
- https://network.nvidia.com/pdf/eol/LCR-000694.pdf
- https://network.nvidia.com/pdf/eol/LCR-000916.pdf
- https://github.com/soliddowant/infra-mk3/blob/HEAD/docs/setup-sx6036.md
- https://www.techtimes.com/articles/326586/20260904/hpe-patches-arubaos-cx-two-independent-no-credentials-rce-paths-found-same-bulletin.htm
- https://socradar.io/blog/cve-2026-73749-hpe-arubaos-cx-rce/
- https://cvetodo.com/news/hpe-patches-critical-98-score-rce-flaw-cve-2026-73749-in-aos-cx-switch-platform
- https://www.redlegg.com/blog/security-bulletin-hpe-aruba-networking-aos-cx
- https://www.securityweek.com/critical-hpe-aos-cx-vulnerability-allows-admin-password-resets/
- https://www.hummingbirdnetworks.com/cisco-meraki-48-port-l3-gigabit-cloud-based-switch-ms390-48-hw
- https://a-power.com/product/cisco-meraki-ms390-48-managed-l3-gigabit-ethernet-white-1u-poe-ms390-48u-hw/
- https://a-power.com/product/cisco-meraki-ms390-24-managed-l3-gigabit-ethernet-1u-poe-ms390-24ux-hw/
- https://www.tech-america.com/item/meraki-enterprise-7-years-enterprise-support/lic-ms390-24e-7y
- https://www.tech-america.com/item/meraki-advanced-license-and-support/lic-ms390-48a-1y
- https://www.unisolinternational.com/product/meraki-ms390-48-port-advanced-license-and-s/
- https://hssl.us/meraki-ms390-24port-advanced-lics-sup-3yr-cisco-meraki-usa-meraki-ms-series-advanced-3yr-cisco-meraki-usa-meraki-ms-series-advanced-3-year-lic-ms390-24a-3yr/
- https://tickergrove.com/stories/nvidia-q2-fy2027-results
- https://www.sec.gov/Archives/edgar/data/1045810/000104581026000073/q2fy27pr.htm
- https://convergedigest.com/nvidia-q2-fy2027-data-center-vera-rubin-ai-infrastructure/
- https://news.alphastreet.com/nvidia-corporation-nvda-q2-2027-earnings-call-transcript/
- https://www.dell.com/support/product-details/en-hk/product/networking-n3200-series
- https://www.etb-tech.com/networking/dell-switches/dell-n-series
- https://www.delltechnologies.com/asset/de-ch/products/networking/technical-support/dell-networking-n3200-powerswitch-specsheet.pdf

**Round-4 collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in §QQ above. Existing sections §§1–8, §§A–M, §§N–X, "round 2" §§N–S, and §§AA–II were not modified.

---

