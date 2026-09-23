---
id: etape6-tracka-cisco-juniper/01-part-1-cisco/overview
title: "PART 1 — CISCO"
domain: part-1-cisco
role: deep-dive
task: reference
actors: ["Broadcom", "Nvidia"]
dates: ["2026-02-10", "2026-09-15", "2026-09-22"]
keywords: ["cpo", "energy", "ethernet", "gpu", "hbm", "hyperscaler", "latency", "memory", "nvidia", "optics", "pricing", "research"]
source: docs/RAG/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [7, 71]
section: "PART 1 — CISCO"
sha256: 8ace20bc6e45b12a49e2abaa6df2ce1904df34e7a3ebb4afcbdaf4806dc71a91
---

# PART 1 — CISCO

# Cisco Enterprise/Data-Center Networking 2026 — Research Report

Coverage as of: September 22, 2026.
Prepared by subagent research, read-only web research. No live browser navigation performed.

**Legend:** [official] = Cisco press release, blog, data sheet, earnings release, investor deck | [vendor-reported] = partner (DDN, NetApp, WWT, CDW, Computacenter) quotes in Cisco PR | [independent] = trade press / analyst outlets | [secondary] = aggregators, blogs, opinion pieces, investor blogs | [unverified] = single-source or low-reliability claims; treat with caution.

---

## 1. Executive Summary

2026 was Cisco's breakout AI-networking year, centered on the February 10, 2026 launch of the **Silicon One G300 (102.4 Tbps)** chip at Cisco Live EMEA in Amsterdam, followed by AgenticOps software (Cisco Cloud Control) at Cisco Live Las Vegas (June 2–4, 2026) and a September 15, 2026 Splunk AI POD announcement at Splunk.conf in Denver. Financially, Cisco reported **$9.3B in FY2026 AI infrastructure orders from hyperscalers (4.5× FY2025)** and **~$4B in related revenue** [official via multiple trade press; vendor-leaning but earnings-sourced], guiding to **$7.5B AI infra revenue in FY2027** [independent via Barron's/Trefis coverage of earnings call]. Roughly 60% of AI orders were Silicon One systems, 40% optics [secondary via Trefis]. M&A in 2026 skewed toward AI security/observability (Galileo, Astrix Security, Teleport partnership, Astrix ~$400M [unverified]) rather than data-center switching silicon. No 204.8T chip announcement from Cisco was found in 2026; the top announced density remains 102.4T (G300).

---

## 2. Silicon One — G300 and roadmap

### Silicon One G300 (flagship 2026 announcement)
- **Announced February 10, 2026** at Cisco Live EMEA, Amsterdam [official] (press release via stockhouse/PR Newswire mirrors; original at Cisco newsroom).
- **102.4 Tbps switching capacity**, single-chip; designed for massive "scale-out" AI clusters [official].
- Powers new **Nexus 9000 and Cisco 8000** fixed and modular Ethernet systems, in air- and **100% liquid-cooled** configurations [official].
- Specs per Cisco PR coverage: deterministic, low-latency, power-efficient; supports **64× 1.6T (1600 GbE)** port configurations [independent via eWeek].
- **"Industry's largest on-chip buffer"** claim [vendor-reported] — per quotes in the Cisco PR from WWT ("102.4 terabits of capacity with the industry's largest on-chip buffer"). Exact buffer size in MB/GB was **not disclosed** in the materials found — flagged as a gap.
- G300 marketing framed around "Intelligent Collective Networking": shared packet buffers, proactive telemetry, path-based load balancing to boost GPU utilization and job completion times [independent via SDxCentral/HostNoc summarizing Cisco PR]. These are **marketing terms**; exact mechanisms (beyond Intelligent Packet Flow, ECN, PFC, DLB — see §6) were not technically specified in found sources — flag as uncertainty.
- Partners validating the stack in the PR: **DDN (Sven Oehme, CTO), NetApp (Syam Nair, CPO), WWT (Neil Anderson), CDW (Brian Campbell), Computacenter (Thomas Berger)** [vendor-reported].
- New systems announced alongside: **N9364F-SG3** (G300-powered Nexus 9000, 64 ports 800G/1.6T, 102.4T total) and **N9364E-SP2R-X** (P200-powered, 64 ports 800G/1.6T, 51.2T total) [vendor-reported via CDW quote in PR; treat SKU list as vendor-reported].
- Liquid-cooled G300 systems claim "nearly 70% energy efficiency improvement," delivering in a single system the bandwidth that previously required 6 prior-generation systems [official]. Baseline of comparison not specified — flagged as an apples-to-oranges caveat.

### P200 / G200 (prior generation, context)
- Silicon One **P200** (51.2T) powered 51.2T Nexus 9000 systems announced for hyperscale deployments in 2025; at Cisco Live EMEA 2026 Cisco announced an **expanded portfolio of P200-powered systems** for spine, core, and interconnect roles (neoclouds, SPs, enterprises), plus expanded OS support on 8223 systems [independent via CRN/SDxCentral].
- Five new hyperscaler design wins reported in Q3 FY2026 (per Lightwave) included "systems powered by Silicon One P200 and G200 chips" [secondary]. Note: G200 details not covered in this research pass; **no G500 announcement found as of Sept 2026** — the task's "G300/G500 (or newer)" premise: only G300 is confirmed [independent].

### Port-density announcements (2026)
- 51.2T, 102.4T: confirmed (P200, G300) [official/independent].
- **204.8T: no Cisco announcement found in 2026 research** [independent]. Flagged: if 204.8T-class news exists, it comes from competitors (e.g., Broadcom Tomahawk-class roadmaps), not Cisco's public announcements in the covered sources.

---

## 3. Nexus 9000 new SKUs (2026)

### Nexus 9364F-SG3 (G300)
- G300-powered fixed system; 64 ports of 800G/1.6T; 102.4T total throughput [vendor-reported]. Positioned for hyperscale/enterprise AI scale-out [official/independent].

### Nexus 9364E-SP2R-X (P200)
- P200-powered; 64 ports 800G/1.6T; 51.2T total throughput [vendor-reported]. Announced at Cisco Live EMEA 2026 [independent via CRN].

### Nexus 9332D-H2R (deep-buffer 400G)
- **32-port 400G QSFP-DD**, 1RU, 25.6 Tbps of bandwidth [official via Cisco data sheet].
- **First switch in the Nexus 9000 fixed-switch portfolio with deep-buffer capability**: 80 MB on-die packet buffer + **8 GB high-bandwidth memory (HBM)** [official].
- MACsec on all 32 ports; PTP Class C timing accuracy [official].
- Power: 2000W AC/DC/HV PSUs; typical 1015 W, maximum 1900 W; operating temp 0–40°C [official].
- Cisco positioning: for AI/ML fabrics, with "the hardware and software capabilities to provide the right latency, congestion-management mechanisms, and telemetry to meet the requirements" of AI workloads [official via product page text].
- **Price (third-party listing)**: Zones listed N9K-C9332D-H2R at **$128,731.99** [secondary]. List price from Cisco not found — treat as retailer pricing, may vary.
- NOTE: The data sheet found is ©2026; the product was introduced earlier (the SKU predates 2026 per some © dates) — flag that it is not strictly a "2026 launch" but is Cisco's current deep-buffer 400G fixed option and featured in 2026 AI-cluster white papers.

### Other SKUs referenced in 2026 materials
- **N9364E-SG2-Q** (Nexus 9300, 64× 800G QSFP-DD spine-leaf) — appears in the 2026 Cisco AI POD design for NVIDIA 2-8-9-400 Enterprise Reference Architecture (6 units in backend fabric) [official via Cisco CVD PDF].
- **N9K-C9364D-GX2A** (64-port 400GE) — used as frontend/spine-leaf in AI POD BOM and in Cisco's own 256-GPU (H100) and 512-GPU (H200) reference cluster designs [official via Cisco white paper/CVD].
- **Nexus 9364E-SG2** (high-density 800G aggregation, 400/200/100G breakout support) [independent via Network World, referencing Cisco blog].
- Cisco used a **rail-optimized, non-blocking backend network** for its 256× NVIDIA H100 GPU cluster (24× Nexus 9364D-GX2A per 512× H200 design) [official via Cisco AI-clusters white paper].

---

