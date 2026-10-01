---
id: collect-261001-general-networking/general-networking/etape6-trackb-arista-sonic-4
title: "Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Broadcom", "Huawei", "Meta", "Microsoft", "Nvidia", "OpenAI", "TSMC"]
dates: ["2026-03-12"]
keywords: ["3nm", "asic", "cpo", "ethernet", "gpu", "inference", "latency", "lpo", "nvidia", "optics", "rack-scale", "revenue"]
source: docs/RAG/collect-261001-general-networking/etape6_trackB_arista_sonic.md
source_anchor: ""
source_lines: [188, 245]
sha256: a18590ef4ced55bf64b5224e0871f4622ca63c8be0e480457b131f4a8d65e50b
---

# Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)

- **Broadcom began volume shipments** of the Tomahawk 6 switch series on **March 12, 2026** — first 102.4 Tbps single-chip Ethernet switch; **TSMC 3nm**; 512×200G SerDes channels; ~250ns latency (UE-compatible variant); Cognitive Routing 2.0, Global Load Balancing to attack incast congestion **[secondary — financialcontent/marketminute, Mar 26, 2026; TechPowerUp]**.
- Positioning: unified scale-up and scale-out AI networks; support for 100G/200G SerDes and co-packaged optics (CPO); designed for clusters of **more than one million XPUs**; open specifications for easy integration **[secondary — TechPowerUp]**.
- **Arista and Cisco integrated TH6** into their latest chassis — reporting "flatter" topologies: large AI clusters built with only **two tiers** instead of three or four, cutting optical transceiver count and GPU-to-GPU latency **[secondary — financialcontent]**.
- **Initial deployments scheduled for clusters of 100,000+ XPUs** **[secondary — StorageReview]**.
- **DriveNets** built AI Fabric switches around the Tomahawk 6 ASIC, targeting AI infrastructures with hundreds of thousands of XPUs; systems expected to ship **Q3 2026** — a visible system-level demand signal for TH6 **[secondary — ainvest, Jul 2026]**.
- Broadcom Q1 FY2026 (Mar 2026): total revenue $19.31B (+29.5%); AI infrastructure revenue **$8.4B (+106%)**; networking segment revenue +60% YoY, with Tomahawk 6 "effectively secured Broadcom's technological lead in data center fabrics"; confirmed **OpenAI as sixth custom-silicon (XPU) customer**; Anthropic $21B "Ironwood Racks" order **[secondary — financialcontent, Mar 6, 2026]**.
- ⚠️ "Beat competitors to the 100T era by at least two full quarters" is a secondary outlet's characterization, not independently verified.

### 4.3 White-box / ODM switching

- IDC no longer breaks out ODM Ethernet revenue publicly; **The Next Platform estimates ODMs at $3.56B in Q2 2026, up 2.43× YoY** — its own estimate, **[unverified as to exact figure; directionally consistent with AI buildout]** **[secondary — The Next Platform, Sep 20, 2026]**.
- Open-hardware ecosystem for SONiC deployments: Accton/Edgecore, Celestica, Wistron, Cisco, NVIDIA (Aviz-certified images) **[secondary — Aviz]**.
- SONiC enables NOS choice on white-box fleets; Dell has promoted this model for enterprise **[independent — The Register]**.

---

## 5. ARISTA vs NVIDIA SPECTRUM — COMPETITIVE POSITIONING

### 5.1 Market-share data (IDC, via press)

**Q1 2026 — NVIDIA takes the data-center Ethernet lead**
- **NVIDIA became the #1 data-center Ethernet switch vendor**: ~$2.1B revenue, **+193% YoY**, **21.5% share** — ahead of Arista (20.7%) and Cisco (17.8%) **[secondary — ainvest, citing IDC; remio.ai]**.
- NVIDIA's share climbed from <4% two years earlier to 21.5% in one quarter; the vehicle is **Spectrum-X** **[secondary]**.
- Overall Ethernet switching market: **$15.4B in Q1 2026, +39.8% YoY** **[secondary — remio.ai]**.
- NVIDIA's data-center Ethernet share was ~11% in late 2025 / 11.6% in Q3 2025 **[secondary]**.

**Q2 2026 — the gap widens (IDC via The Next Platform, Sep 20, 2026)**
- Total Ethernet switch revenue grew strongly; **NVIDIA's Ethernet revenue hit $3.86B (2.8×)**, driven by the GenAI systems business and **Spectrum-X attach rates** on DGX nodes and NVL72 rack-scale systems **[secondary — The Next Platform]**.
- Arista: **$2.53B, +37.6%**; Cisco: $5.43B (+36.2% overall, data-center +43% in Q1); HPE (incl. Juniper): $1.15B (+6.1%); Huawei: $1.55B (+28.6%); ODMs est. $3.56B (2.43×) **[secondary — The Next Platform; Q1 HPE/Huawei figures from IDC via ainvest]**.
- Data-center Ethernet in Q2: **$12.3B, +64.5% YoY**, +23% sequentially; ~129.5M ports sold into the data center (analyst estimate); speeds above 200G are overwhelmingly data-center/AI **[secondary — The Next Platform]**.

### 5.2 Architecture comparison: Etherlink vs Spectrum-X

| Dimension | Arista Etherlink | NVIDIA Spectrum-X |
|---|---|---|
| Switch ASIC | Merchant (Tomahawk 5 / Tomahawk 6 / Jericho 3-AI) | In-house Spectrum-4 (+ Spectrum-6 on roadmap) |
| Endpoint / NIC | Third-party NICs (no own NIC) | BlueField-3 SuperNIC / ConnectX; DPUs bundled |
| Fabric stack | RoCEv2 + PFC + DCQCN today; UEC/NSCC/RCCC roadmap | Integrated platform: Spectrum switches + BlueField DPUs + LinkX cables + CUDA stack |
| Congestion control | DCQCN now; UEC NSCC/RCCC later | Closed-loop TCC via SuperNIC |
| Multipath | DLB, packet spraying, MRC/CSIG (EOS) | Per-packet spraying + TCC |
| UEC standing | **Founding/steering member** | Member (~2024), pushes Spectrum-X |
| Optics | LPO/XPO, multi-vendor | Integrated; Spectrum Photonic (TSMC COUPE silicon photonics) roadmap |
| Best-fit | Open multi-vendor AI fabrics; merchant-silicon + UEC trajectory; no NVIDIA networking lock-in | NVIDIA-GPU AI factories; pre-optimized full stack |

**Comparison sources:** LLM-Systems-Wiki Arista/Etherlink page (secondary, vendor-adjacent technical detail) and StorageReview TH6/Spectrum-Photonic comparison (secondary). Treat table cells as secondary-sourced where marked.

### 5.3 Analyst framing of the competitive dynamic

- **The bullish-on-Ethernet case**: by early 2026, Ethernet had "largely caught up to NVIDIA's proprietary InfiniBand in terms of latency and congestion management," capturing **>65% of new AI back-end deployments**; Meta, Microsoft, Amazon, Alphabet validated Ethernet (RoCE) as equivalent performance at lower TCO vs InfiniBand **[secondary — Tamar Securities, Jan 2026]**.
- **The bearish-on-Arista-multiple case**: Arista's stock is priced on the smallest, most contested slice of its revenue (the AI back end), while NVIDIA — which Arista "spent two decades building its identity on" — just became #1 in data-center Ethernet. "Hyperscalers are no longer buying networking as an isolated silo... It is not as good as it looks when the customer stops buying switches and starts buying GPU clusters that happen to include switches." **[secondary — ainvest, Aug 2026]**.
- **Arista's response**: openness and operational consistency — one EOS binary across DC/AI/routing/campus; mix-and-match accelerators and NICs from multiple suppliers; founding UEC member betting Ethernet becomes the default for AI networking, not just scale-out **[secondary — remio.ai; The CODEW]**.
- **NVIDIA's advantage**: vertical integration — GPU + switch chip + switch + SuperNIC + cables + software as one pre-optimized, single-support-number AI factory; IDC's read: integrated GPU-plus-networking co-design is winning AI deals **[secondary — ainvest; The Next Platform]**.
- "AI back-end" market expected to surpass **$15B annually by end of 2026**; trend shifting from inference to large-scale training **[secondary — Tamar Securities]**.

---

## 6. OPEN VERIFICATION ITEMS

