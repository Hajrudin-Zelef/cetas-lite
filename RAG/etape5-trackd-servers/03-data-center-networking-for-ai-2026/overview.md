---
id: etape5-trackd-servers/03-data-center-networking-for-ai-2026/overview
title: "Data-Center Networking for AI (2026)"
domain: data-center-networking-for-ai-2026
role: deep-dive
task: reference
actors: ["Broadcom", "Meta", "Nvidia", "TSMC"]
dates: ["2025-06-11", "2025-09", "2025-10-14", "2026-03-12", "2026-06-02", "2026-07", "2026-07-16", "2026-09", "2026-09-22"]
keywords: ["3nm", "asic", "blackwell", "cpo", "dci", "ethernet", "gpu", "hyperscaler", "inference", "latency", "llama", "lpo"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [677, 744]
section: "Data-Center Networking for AI (2026)"
sha256: 63288f9d206e79a6df79357d60ed9065c19b5b31fe460773047732c64634c3fb
---

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

- **Etherlink AI portfolio:** Arista's AI-optimized Ethernet portfolio; customer base expanded to **over 100 customers** (Q2 2026), up from just 4–5 in 2024. [vendor-reported] (cryptobriefing summarizing Arista Q2 2026 earnings)
- During Q2 2026 Arista introduced its **7506c/7 Series AI fabric portfolio** and **1.6 Tbps platforms with liquid-cooled options** (also reported as the **7060XE7 Series 1.6 Tbps platforms**) for the largest AI training/inference clusters. [vendor-reported] [secondary]
- Arista's official Q2 2026 release notes: "Introduced 1.6 Tbps AI fabric platforms, including liquid-cooled options optimized for scale-up, scale-out, and scale-across networks." [official]
- The 7060X6 (Tomahawk 5-based 800G) and 7800R4/7700R4 (modular 800G systems) lines are Arista's shipping 800G portfolio for AI backends; Meta's RoCE-over-Ethernet 24,576-GPU H100 cluster for Llama 3 used Arista 7800 systems side-by-side with a Quantum-2 InfiniBand cluster, both reaching 90%+ network utilization. [secondary]
- **Uncertainty flag:** I did not independently confirm per-SKU shipment volumes for the 7060X6/7800R4/7700R4; the task's "shipments" granularity is not disclosed at SKU level in the sources found. Do not fabricate SKU shipment numbers.

### 1.4 Cisco Nexus 9000 800G / Silicon One

- Cisco introduced **Silicon One G300** — its own 102.4 Tbps networking silicon — debuting in **liquid-cooled N9000 and 8000 series switches**, aimed at competing with Broadcom and NVIDIA switching chips. [secondary] (ainvest)
- Simultaneously, Cisco ships **N9100 series switches explicitly powered by NVIDIA Spectrum-X Ethernet silicon**, running Cisco NX-OS on NVIDIA hardware. [secondary] (ainvest)
- Cisco FY2026 hyperscaler AI design wins: three new wins in Q4 alone — one **Silicon One P200 scale-across** deployment, one **G200 scale-out** project, and one optical line-system deployment; management flagged line-of-sight to more wins over the next six months across G300, G200, P200 and A100 platforms. [vendor-reported] (infotechlead)
- TrendForce platform table (2025 vintage, cited 2026): Cisco Silicon One G200 series at 51.2 Tbps (2023) alongside Marvell Teralynx 10; CPO prototypes. [secondary]

### 1.5 NVIDIA Spectrum-X switches (SN5000-class / SN5600 / Spectrum-X800)

- **Spectrum-X800 Ethernet switch** (SN5600-class): 800 Gb/s per port, 51.2 Tb/s switching capacity (per temperature2.com / TrendForce table). Based on the **Spectrum-4 switch ASIC**; SN5600 = 64-port 800GbE. [secondary]
- **Spectrum-6** (announced July 2026): 102.4 Tbps Ethernet switch system, 2× previous-generation capacity, integrated into the Rubin-architecture stack (Vera CPU, Rubin GPU, NVLink 6, ConnectX-9 SuperNICs, BlueField-4 DPUs, Spectrum-6). [vendor-reported] (cxotoday citing NVIDIA)
- TrendForce roadmap table: Spectrum-X800 (800G/port, 51.2 Tbps) → Spectrum-X1600 (1.6T/port, 102.4 Tbps, paired with Rubin) → Spectrum-X3200 (3.2T/port, 204.8 Tbps). ConnectX-8 (800G, PCIe 6.0) → ConnectX-9 (1.6T, PCIe 7.0) → ConnectX-10 (3.2T, PCIe 8.0). [secondary]

### 1.6 Celestica's position

- Celestica regained #1 in Ethernet AI backend networks in Q1 2026 per Dell'Oro (above). Celestica is NVIDIA's primary Ethernet-switch manufacturing partner for Spectrum-X systems; its ranking reflects white-box/OEM volume into AI fabrics. [independent] [secondary]

---

