---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-10
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["Broadcom", "China", "Nvidia", "TSMC"]
dates: ["2025-06-11", "2025-09", "2025-10-14", "2026-03-12", "2026-06-02", "2026-07", "2026-07-16", "2026-09", "2026-09-22"]
keywords: ["3nm", "blackwell", "cpo", "dci", "ethernet", "gpu", "latency", "lpo", "nvidia", "optics", "pricing", "rack-scale"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [663, 719]
sha256: d529f3cc602de2e07ca74e3cea5cb99aaac3b9f491d49e6269d6a5ed1b35664e
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

1. **HPE vs branded-OEM share shift**: IDC Q2 2026 — branded OEMs are capturing a growing share of AI infrastructure (ODM Direct share 60.6% → 53.9% YoY); HPE +46.0% YoY to $5.87B [independent — IDC]. This validates the OEM AI-server wave these vendors ride.
2. **GB300/Vera Rubin timing**: Wccftech (Dec 2025-era report) noted GB300 mass-production concerns and CSP preference for mature HGX 8-GPU systems; GB300 shipments projected +129% YoY in 2026 [secondary — https://wccftech.com/nvidia-blackwell-ultra-ai-servers-to-lead-the-ai-infrastructure-race-moving-into-2026/]. Rubin rack-scale platforms (Giga, ASUS) are staged for **H2 2026** [secondary].
3. **Prices**: no public list pricing was found for any AI server in this wave (all enterprise quote-based) — systematic gap.
4. **Model-number uncertainty**: "ESC N8A-E12" (ASUS) and "Cray XD675" (HPE) appear in the task brief but were not corroborated in 2026 sources; do not use them as confirmed current SKUs.
5. **xFusion product detail**: the weakest area; needs Chinese-language source work (company site, WeChat, CSRC filings) — recommended follow-up by a follow-on research task.
6. **MSI customer/financial detail**: also a gap — MSI's AI-server business scale and 2026 design wins were not found publicly.

**Report written to**: `~/workspace/rag_collect/_draft_etape5_hpe_others.md` (this file)


---

## §3 Data-center networking

# Data-Center Networking for AI (2026)

**Status:** as of 22 September 2026. Report compiled 2026-09-22 (UTC).
**Scope:** switch silicon and systems for AI fabrics (200/400/800 GbE), NVIDIA Spectrum-X / Spectrum-XGS platform, InfiniBand vs Ethernet, UEC / Super Ethernet, optical interconnect (CPO, LPO, 800G/1.6T transceivers, NVIDIA silicon photonics), Arista and Cisco 2026 AI networking business, DCI / 800G adoption figures and 1.6T timeline.
**Provenance key:** every bullet carries a tag — `[official]` = vendor press release / investor-relations page; `[vendor-reported]` = company claim via earnings call or company blog (not independently audited); `[independent]` = independent press (ServeTheHome, Data Center Dynamics, SemiEngineering, EE Times, etc.) or analyst firm data reproduced faithfully; `[secondary]` = trade/financial press or analyst summaries; `[unverified]` = plausible but not independently confirmed by a named source. **Conflicts and gaps are flagged explicitly; identifiers are only quoted when found in a source.**

---

## 0. Executive summary

- 800G is the mainstream AI-fabric generation in 2026; 1.6T is shipping in early volume and is expected to ramp in H2 2026. [independent] [secondary]
- Dell'Oro Group (Q1 2026 quarterly snapshot, published June 2, 2026): Ethernet switch sales in AI backend networks "more than doubled" YoY and were "about two-thirds of data center switch sales in AI clusters"; InfiniBand sales "more than tripled" in the quarter, supported by the 800 Gbps switch ramp with NVIDIA's Blackwell Ultra platform; Dell'Oro cautioned the IB rebound may partly be brownfield upgrades, not greenfield expansion. Same report: Celestica #1 in Ethernet AI backend networks, followed closely by NVIDIA, Arista third, Cisco fourth (largest share gain in the quarter). [independent] (analyst data cited via gpusmith summary of delloro.com)
- NVIDIA is now the largest data center Ethernet switching vendor per Dell'Oro figures cited by secondary press (~$2.1B quarterly switch revenue, 21.5% share in 2026, up from <4% two years ago). [secondary]
- Broadcom's Tomahawk 6 (102.4 Tbps, 512×200G SerDes, 3nm) reached production volume March 12, 2026. [official]
- NVIDIA announced Spectrum-6 (102.4 Tbps Ethernet switch, part of the Rubin stack) in July 2026 and Spectrum-XGS Ethernet (scale-across, software/firmware upgrade, ~2× NCCL throughput claim) — both [official].
- UEC Specification 1.0 released June 11, 2025; 1.0.1 in September 2025; current 1.0.3 as of July 16, 2026 (per community wiki; confirm against ultraethernet.org). Broadcom Thor Ultra NIC (Oct 2025) was the first NIC built to the spec; Keysight–Broadcom demoed UEC LLR/CBFC at 800GE at OFC 2026. [independent] [unverified]
- CPO entered the market in 2026 with early products from NVIDIA and Broadcom, but large-scale high-volume CPO is widely expected only ~2028–2030; LPO is shipping in volume in 2026 (millions of units expected). [independent]
- TrendForce: 800G+ optical transceiver shipments ~24M units (2025) → ~63M units (2026), 2.6×; 1.6T shipments ~2.5M (2025) → 20M+ by end-2026. EML laser capacity is a bottleneck; NVIDIA has pre-allocated large EML capacity. [secondary]
- Arista Q2 2026: record $3.036B revenue (+37.7% YoY), raised FY2026 AI-fabrics revenue target to at least $3.5B (at least $3.6B per Zacks), overall FY2026 target ~$11.5B (Aug 2026); one secondary outlet (remio.ai, Sept 19, 2026) claims a raise to ~$12.6B — **conflicting, unverified**. Etherlink customer base >100. [official] [secondary]
- Cisco FY2026: $9.3B AI infrastructure orders from hyperscalers (~4.5× FY2025), ~$4B recognized revenue, FY2027 AI revenue guide $7.5B; Silicon One G300 (102.4 Tbps) introduced; Cisco also sells NVIDIA Spectrum-X-powered N9100 switches. [vendor-reported] [secondary]

---

## 1. 200/400/800 GbE switch adoption in AI clusters

### 1.1 Adoption picture (Dell'Oro Q1 2026)

- Dell'Oro Group's most recent quarterly snapshot (published June 2, 2026) on AI backend networks: "Ethernet switch sales in AI backend networks more than doubled and accounted for about two-thirds of data center switch sales in AI clusters during the first quarter 2026," while "InfiniBand sales more than tripled during the quarter, supported by the ramp of 800 Gbps switches shipping with NVIDIA's Blackwell Ultra platform." [independent] (cited via gpusmith summary of delloro.com)
- Dell'Oro's analysts cautioned the InfiniBand rebound "may partly reflect upgrades of the installed base in brownfield deployments rather than greenfield expansion." [independent]
- Vendor share in Ethernet AI backend networks (Q1 2026): "Celestica regained the leading position, followed closely by NVIDIA," with "Arista ranked third" and "Cisco recorded the largest share gain during the quarter and ranked fourth." [independent]
- Same source: "800 Gbps switches accounted for the vast majority of the Ethernet switch shipments and revenues in AI backend networks during the quarter," with 1600 Gbps switches "only beginning to sample" and expected to "ramp in the second half of 2026." [independent]
- Current-generation switch silicon across both IB (Quantum-X800) and Ethernet (Spectrum-X, Tomahawk 5/6, Silicon One G200) sits at roughly the same 800 Gb/s per-port tier. [secondary] (temperature2.com)

### 1.2 Broadcom Tomahawk 5 / Tomahawk 6

- **Tomahawk 6 family (TH6):** production volume shipments announced **March 12, 2026** [official] (GlobeNewswire via PR). Key specs: 102.4 Tbps single-chip switching capacity (2× Tomahawk 5); 512 × 200G SerDes (or 1024 × 100G); 3nm (TSMC, per secondary reporting); support for 100G and 200G SerDes; CPO (co-packaged optics) option; AI routing features incl. Cognitive Routing 2.0 and Global Load Balancing; targets >1M-XPU clusters; 128K-XPU fabrics in two tiers; single-chip 512-XPU connectivity. [official]
- Sameh Boujelbene (VP, Dell'Oro Group), quoted in Broadcom's release: TH6 addresses "large-scale, low-latency fabrics with fewer switch tiers and reduced optical complexity." [vendor-reported]
- Arista and Cisco have integrated TH6 into latest chassis per secondary coverage (stocktitan/financialcontent). [secondary]
- **Tomahawk 5 (TH5):** 51.2 Tbps, 100G/200G SerDes, the prior generation; base of Arista's 7060X6-class and other 800G platforms. [secondary]
- **Thor Ultra NIC:** Broadcom announced October 14, 2025 as the first NIC built to the UEC 1.0 specification. [secondary] (temperature2.com citing Broadcom)

### 1.3 Arista (7060X6, 7800R4, 7700R4, 7060XE7, Etherlink)

