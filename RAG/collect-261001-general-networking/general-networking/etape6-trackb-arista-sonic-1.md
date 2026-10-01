---
id: collect-261001-general-networking/general-networking/etape6-trackb-arista-sonic-1
title: "Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: ["2026-01-01", "2026-05", "2026-06", "2026-06-04", "2026-06-09", "2026-08-04", "2026-09-22"]
keywords: ["benchmarks", "capex", "ethernet", "nvidia", "research", "revenue"]
source: docs/RAG/collect-261001-general-networking/etape6_trackB_arista_sonic.md
source_anchor: ""
source_lines: [1, 57]
sha256: 91f17cc98fccab8c0f3b896064ad1063b26d16789047a944ad3385f1389e4f0e
---

# Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)

## Research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 6 (Network & Security), Track B
**Coverage window:** January 1, 2026 → September 22, 2026 (some 2025 context where explicitly marked)
**Scope:** Arista Networks, SONiC (community + commercial distributions), NVIDIA Cumulus Linux, open-networking landscape (UEC/Super Ethernet, Tomahawk 6, white-box), Arista vs NVIDIA Spectrum competitive positioning
**Status:** Research snapshot. Figures are dated snapshots as of September 22, 2026; re-verify against primary sources before use.

### Provenance legend
- **[official]** — vendor's own press release, docs, earnings statement, or regulatory filing.
- **[vendor-reported]** — figure claimed by the vendor (benchmarks, customer counts, adoption) without independent audit.
- **[independent]** — reputable third-party press (Bloomberg, Reuters, CNBC, The Register, Network World, The Next Platform, SiliconANGLE) or independent measurement.
- **[secondary]** — lower-tier press, blogs, analyst writeups, investment sites, community wikis; useful but treat carefully.
- **[unverified]** — single-source, conflicting, or thinly sourced claims; do not treat as fact.

---

## 1. ARISTA NETWORKS

### 1.1 Financial results 2026

**Q2 2026 (reported August 4, 2026) — first $3B quarter**
- Revenue **$3.036 billion**, +12.1% QoQ, **+37.7% YoY** vs Q2 2025 **[official — Arista Q2 2026 earnings release; confirmed by The Next Platform, Sep 20, 2026]**.
- GAAP diluted EPS **$0.95**, non-GAAP diluted EPS **$1.02** (+39.7% YoY from $0.73) **[official]**.
- GAAP operating margin 45.4%, non-GAAP operating margin **49.9%** (vs 44.7%/48.8% in Q2 2025) **[official]**.
- CEO Jayshree Ullal: "As we deliver our first $3 billion quarter in Q2 2026, it is clear that our Arista 2.0 platform strategy is compelling. Customers see networking as the central nervous system for infrastructure from the client to campus to data and AI centers." **[official]**.
- CFO Chantelle Breithaupt (appointed CFO 2025) cited "strong, broad-based growth" **[official]**.
- Full-year 2026 revenue guidance raised to approximately **$12.6 billion**, implying roughly 40% growth **[secondary — Tickeron earnings recap, ~Sep 14, 2026]**.
- Beat Wall Street consensus (revenue consensus ~$2.83B, EPS consensus ~$0.88) **[secondary — Tickeron]**.
- Non-GAAP gross margin slipped to 63.4% on customer mix and higher component costs **[secondary — Tickeron]**.

**Q1 2026 (reported May 2026)**
- Revenue **$2.71 billion**, +35.1% YoY (vs $2.62B expected); non-GAAP EPS **$0.87** vs $0.81 forecast **[secondary — CoinCentral, ~June 9, 2026; ainvest]**. GAAP operating margin 42.7%, non-GAAP 47.8% **[secondary]**.
- Q2 guidance at the time: ~$2.8B **[secondary]**.
- Balance sheet (Mar 31, 2026): $2.79B cash + $9.56B marketable securities; $1.69B operating cash flow in Q1; total liabilities $8.17B, equity $13.49B; low-debt profile **[secondary]**.

**Context (secondary investment-press framing, treat figures as approximate)**
- Zacks "Bull of the Day" (Sep 16, 2026): revenue more than doubled 2022→2025 ($4.38B → $9.00B); GAAP EPS $1.07 → $2.75 over the same period; $13.3B cash and equivalents, $23.7B total assets vs near-zero debt and $8.9B liabilities; five consecutive years of EPS beats; Zacks Rank #1 (Strong Buy) **[secondary — Zacks]**.
- Bank of America raised price target to **$200** (Buy) after the 1.6T launch; average analyst target ~$186.47 **[secondary — CoinCentral]**.
- Major shareholder Andreas Bechtolsheim (co-founder) sold 240,000 shares on June 4, 2026 under a pre-arranged 10b5-1 plan **[secondary — CoinCentral]**.
- AI hyperscalers projected to spend ~$800B in AI-related capex in 2026 **[secondary — Zacks]**.

### 1.2 AI networking revenue and Etherlink portfolio

**AI revenue trajectory (vendor-reported/management statements — handle with care)**
- Management expects AI revenue to reach **at least $3.6B in 2026**, supported by scale-up, scale-out and scale-across deployments **[secondary — Zacks, Sep 16, 2026]**.
- Earlier framing (late 2025 / Feb 2026 reports): Arista **doubled its AI networking revenue goal to $3.25B for FY2026**, up from ~$1.5B in FY2025 **[secondary — Financial Freedom is a Journey, Feb 2026]**.
- ⚠️ The $3.6B figure supersedes the $3.25B figure; use the most recent (Sep 2026) as the current target. Both are management guidance, not independently audited.
- Etherlink AI fabric portfolio: **more than 100 cumulative customers** as of Q2 2026, up from "only a handful of early adopters in 2024" **[secondary — Tickeron Q2 2026 recap; Zacks]**.

**Etherlink AI — platform positioning**
- Arista's AI Ethernet fabric portfolio spans **400G, 800G and (since June 2026) 1.6T** switching **[secondary — The CODEW, Aug 2026]**.
- Technical foundation: RoCEv2 + PFC + DCQCN congestion control, Arista dynamic load balancing (DLB), plus per-packet spraying, congestion signaling, and trimming features aligned with the Ultra Ethernet transport trajectory — described as a "bridge today (RoCEv2 + PFC + DLB) with UEC plumbing being added" **[secondary — LLM-Systems-Wiki Arista/Etherlink page, updated Aug 2026; vendor-reported]**.
- **7700R4 Distributed Etherlink Switch**: single-hop distributed architecture designed to scale to **more than 30,000 connected accelerators** while preserving lossless, deterministic forwarding — an alternative to conventional multi-tier spine-leaf for very large clusters **[secondary — The CODEW]**.
- EOS implements RoCEv2, DCQCN, and priority flow control alongside Arista's own load-balancing for AI workloads **[secondary — The CODEW]**.

