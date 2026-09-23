---
id: etape6-tracka-cisco-juniper/01-part-1-cisco/4-cisco-ai-pods-2026-updates
title: "4. Cisco AI PODs — 2026 updates"
domain: part-1-cisco
role: deep-dive
task: reference
actors: ["AMD", "Google", "Intel", "Nvidia", "OpenAI"]
dates: ["2025-03", "2026-05", "2026-05-13", "2026-07-25"]
keywords: ["agent", "amd", "fine-tuning", "gpu", "gpus", "hyperscaler", "intel", "latency", "memory", "neocloud", "nvidia", "nvlink"]
source: docs/RAG/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [72, 124]
section: "PART 1 — CISCO"
sha256: d31e54a6fef56d0488e1dcc13ecc2864c0ed8129122a69f842306a0955a58300
---

# 4. Cisco AI PODs — 2026 updates

## 4. Cisco AI PODs — 2026 updates

### Program overview (2026 data sheet)
Cisco AI PODs = pre-validated, modular, full-stack AI infrastructure under **Cisco Secure AI Factory with NVIDIA** [official]. 2026 data sheet specs:
- GPU density: 32, 64, 128+ scale units; CPUs: dual Intel Xeon Scalable or AMD EPYC; up to 4 TB DDR5/node; NVMe up to 30.7 TB/node [official].
- **Networking: 400G/800G Nexus 9000 series, RoCEv2, lossless, sub-ms latency** [official].
- GPU options: NVIDIA **H100, H200, RTX Pro 6000, L40S, A100**, plus AMD **MI300X** [official].
- Storage: VAST Data, NetApp AFF, Pure Storage FlashArray, Nutanix [official].
- Management: Cisco Intersight + Nexus Dashboard; OS: RHEL, Ubuntu, Rancher, Red Hat OpenShift; security: Hypershield, AI Defense, Isovalent, SSO/AAA [official].
- Ordering SKUs: **AIPOD-POD1** (inferencing/edge), **AIPOD-POD2** (training/fine-tuning) [official].

### 2026 new/updated POD items
1. **Cisco AI POD for Splunk** (Sept 15, 2026, announced at Splunk.conf, Denver) — newest POD configuration within Cisco Secure AI Factory with NVIDIA. Pre-validated stack: Cisco infrastructure + NVIDIA accelerated computing + new AI runtime + Kubernetes, optimized for Splunk AI workloads; targets self-managed on-prem, private cloud, and air-gapped environments. Splunk AI Assistant available now; Agent Launchpad later in 2026; self-hosted models: Cisco Deep Time Series Model, Google Gemma 4, OpenAI GPT-OSS 20B; NVIDIA Nemotron planned. Day-one delivery partners: **Accenture, bitsIO, Wipro, WWT** [secondary via AI Market Watch and eeNews Europe; flagged secondary because Cisco PR not directly found in this pass].
2. **Cisco AI POD Infrastructure for NVIDIA 2-8-9-400 Enterprise Reference Architecture** (2026 CVD) — UCS C885A M8 servers with 8× NVIDIA H200 SXM GPUs, BlueField-3 SuperNICs, Nexus N9364E-SG2-Q backend + N9K-C9364D-GX2A frontend fabrics [official via Cisco design guide].
3. **Internet2 / National Research Platform deployment** (announced ~Sept 21, 2026 via PR Newswire): two AI PODs at two campuses (University of Maryland? — campuses named in full PR; snippet names them via Internet2; specific campus names not captured — flag as gap) adding GPU capacity to the NRP's Nautilus Kubernetes fabric. Combined: **8× NVIDIA H200 GPUs (1.1 TB combined GPU memory), 128 CPU cores, 2 TB system memory, 90 TB NVMe**; each site one UCS C845A M8 chassis with 4× H200 on NVLink [official via PR Newswire release].
4. **Secure AI Factory with NVIDIA** context (announced March 2025, i.e., pre-2026 but foundational): hyperscale-class AI cluster architecture from Cisco + NVIDIA; Nexus One management plane spanning NX-OS and SONiC; Nexus Dashboard AI job monitoring [independent/secondary].

### Customer deployments found
- **Luma AI** (generative video startup): anchor customer for the AMD/Cisco/HUMAIN JV's first 100 MW Saudi deployment [independent via Reuters-cited reports — see §8].
- **Internet2/NRP**: two university campuses (above).
- Hyperscalers: named hyperscaler wins are **not publicly disclosed** by Cisco (orders reported in aggregate). Flag as a gap: no named hyperscaler Cisco win found in sources; the "large customer wins" story is mostly aggregate-order figures, not logos.

---

## 5. Financials — AI infrastructure orders/revenue (earnings calls)

Cisco's fiscal year ends late July (FY2026 ended July 25, 2026). All figures below from earnings coverage [independent, multiple outlets consistent]; exact Cisco IR press-release text not directly verified in this pass — flag as independently-reported-but-earnings-sourced.

### Q3 FY2026 (reported May 13, 2026)
- Revenue: **$15.84B** (+12% YoY, record), adjusted EPS $1.06 [independent via The Weighted Average].
- **$1.9B in hyperscaler AI infrastructure orders in Q3** (triple YoY); **$5.3B booked YTD**; raised FY2026 AI-orders target to **~$9B** (4.5× FY2025), from prior $5B [independent via Barron's/biggo].
- **Raised FY2026 AI infrastructure revenue guidance to $4B** (from $3B) [independent].
- Segment color: total networking product orders +50%+ YoY; data-center switching orders +40%+; campus networking +25%+ [independent via The Weighted Average].
- Same release announced **~4,000 job cuts** over the following three quarters (restructuring alongside AI growth) [independent].
- Q3 also included: Acacia record quarter — **750,000+ 400G and 40,000 800G coherent pluggable optics shipped**; five new hyperscaler design wins; Acacia projected to **grow 200% in FY2026** per CEO [independent via Lightwave].
- Zacks analyst note (May 2026): Cisco expected **>$3B AI infra revenue from hyperscalers in FY2026**; pipeline of **>$2.5B orders** from sovereign/neocloud/enterprise; Q2 FY2026 took **$350M** from neocloud/sovereign/enterprise (separate from $2.1B hyperscaler orders in that quarter) [secondary; investor-relations deck corroborates the $350M Q2 non-hyperscaler AI orders figure — official via Cisco IR deck].

### Q4 FY2026 (reported Aug 12, 2026)
- Revenue: **$17.3B** (+18% YoY, record); non-GAAP EPS **$1.22** (+23%); beat consensus ($16.8B / $1.17) [independent via Barron's].
- **AI infrastructure orders from hyperscalers: $4B in Q4**; **$9.3B for FY2026** (4.5× FY2025; above the $9B target) [independent, consistently reported].
- **AI infrastructure revenue (hyperscalers): ~$4B in FY2026** [independent].
- Networking revenue: **$9.79B in Q4** (+28% YoY) [independent]. Product orders +35% YoY; hyperscaler orders triple digits; networking product orders +40% (8th consecutive quarter of double-digit growth); data-center networking orders +~25% YoY for FY2026 [independent].
- **FY2027 guidance**: total revenue **$72.2–73.4B** (midpoint ~15% growth off $63.3B FY2026 revenue); non-GAAP EPS **$5.05–5.11**; non-GAAP operating margin ~35%; **AI infrastructure revenue from hyperscalers guided to $7.5B** [independent]. 4–5 percentage points of FY2027 revenue growth expected from price increases; memory costs pressuring gross margin [independent via Trefis].
- Margin/headwind caveats: gross margin pressured by higher hardware mix and memory costs; ~60% of AI orders are Silicon One systems, 40% optics [secondary via Trefis]; CEO argues the orders are still highly profitable (low incremental sales expense) [secondary].
- **Nexus switch orders tagged for AI deployments: up >85% sequentially** (CEO Chuck Robbins, Q4 call) [secondary via Benzinga/cloudfront mirror — treat as secondary since not from a transcript).
- Market reaction: stock fell after-hours despite beat [independent via Barron's/Blockonomi].

### Trailing history (context)
- Q1 FY2026: **$1.3B** hyperscaler AI infra orders + **>$200M** from neocloud/sovereign/enterprise [secondary via Motley Fool].
- "Cisco raised AI infrastructure revenue guidance in FY2026 from $2B to $3B" (earlier in year per Network World) — superseded by $4B actual [independent].

---

