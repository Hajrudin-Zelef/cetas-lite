---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-round-4-2026-09-22-4
title: "Supplementary / Complementary Research Pass — round 4 (2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "China", "Huawei", "Meta", "Microsoft", "Nvidia", "xAI"]
dates: ["2025-03", "2025-12", "2026-03-05", "2026-06", "2026-06-09", "2026-07-02", "2026-08-04", "2026-09", "2026-09-15", "2026-09-22"]
keywords: ["research", "asic", "attribution", "compute", "cpo", "distribution", "ethernet", "gpu", "liquid cooling", "lpo", "market cap", "memory"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1484, 1533]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 09fb908408d7c56dccb36f35e0ba83182c2abd6d78bbf69792e3efa217d75f23
---

# Supplementary / Complementary Research Pass — round 4 (2026-09-22)

## Supplementary / Complementary Research Pass — round 4 (2026-09-22)

**Scope note:** A base report (§§1–8) and five supplementary passes already exist (§§A–M, §§N–X, round-2 §§N–S, round-3 §§R3-A–R3, round-3 §§AA–II). This round-4 pass adds ONLY material not present in any of them: Arista Networks 2026 financials (competitive context, never covered), Arista Tomahawk-6 1.6T hardware, Huawei Xinghe AI Fabric 2.0 + CloudEngine XH9230 (MWC Barcelona 2026), 2026 CPO ship-status controversy (SemiAnalysis vs NVIDIA CTO), xAI–Dell $5B deal 2026 status finding, Meraki 2026 hardware status-quo finding. Nothing above was modified. Provenance legend from the base applies. All facts as of 2026-09-22.

### R4-A. Arista Networks Q2 2026 results (new — competitive context, never covered)

- **Q2 2026 (reported Aug 4–5, 2026): revenue $3.036B, +37.7% YoY** — Arista's first-ever $3B revenue quarter; non-GAAP EPS **$1.02** (consensus ~$0.88–0.89); GAAP net income $1.213B; adjusted operating margin **49.9%** (up from 47.8% in Q1 2026 and 48.8% a year earlier) [secondary — tickeron, gate.com, insidermonkey, barrons, 2026-08-04/05].
- **Q3 2026 guidance: ~$3.3B revenue** (consensus ~$2.95B), non-GAAP EPS $1.06–$1.08 (consensus ~$0.92), non-GAAP operating margin 48–49% [secondary — tickeron, dailypolitical].
- **Full-year 2026 revenue guidance raised for the third time: now $12.6B** (from $11.5B); implies ~40% YoY growth; management raised as supply availability improved, enabling more shipments [secondary — tickeron, gate.com].
- CEO Jayshree Ullal: "Customers see networking as the central nervous system for infrastructure from the client to campus to data and AI centers" [secondary — barrons]. COO Todd Nightingale: capacity increased in manufacturing and distribution, better component delivery terms negotiated over the last six months (response to Q1 2026 memory/silicon/optics shortages) [secondary — barrons].
- Market reaction: shares +5.2% to $200.37 in early trading; ANET market cap above $250B; Rosenblatt raised price target $210→$280 (Buy) [secondary — barrons, dailypolitical].
- Relevance to this report's vendors: Arista Q2 2026's $3.036B quarter compares to NVIDIA's Q2 DC Ethernet #1 spot ($2.5B per IDC, cited in pass #1 §G) and Cisco's networking quarter ($9.8B product+services, R3 §R3-A) — in branded DC Ethernet switching revenue, the 2026 order is NVIDIA > Arista ≈ Cisco, with Arista the fastest-growing [independent assessment; figures mixed-source, not apples-to-apples]. Source: https://tickeron.com/blogs/arista-networks-anet-delivers-first-3-billion-quarter-with-37-7-revenue-growth-15344/ ; https://www.barrons.com/articles/arista-networks-earnings-stock-price-e421b3c3

### R4-B. Arista 7060XE7 — Tomahawk-6 1.6T AI platform (new competitive hardware angle)

- The **Arista 7060XE7 series** (announced 2026, reported via Q2 earnings coverage) is built on Broadcom's **Tomahawk 6** (102.4T) silicon: **100 Tbps system bandwidth, 1.6T single-port speed**, air/liquid/hybrid cooling options; **LPO (linear pluggable optics) support** cutting interconnect power ~60% [secondary — gate.com, 2026-08]. **Meta and Microsoft have endorsed the platform** [secondary — gate.com]. First devices **expected to ship Q4 2026** [secondary — gate.com].
- Competitive read-through: Arista's Tomahawk-6 line competes directly with Dell's Z9964F-ON (same 102.4T class, base §1.1) and NVIDIA Spectrum-X (proprietary), and pressures the Tomahawk-6 white-box platforms added in Dell Enterprise SONiC 4.6 (pass #1 §E) — all three vendors converge on the same silicon generation in H2 2026 [independent assessment].
- Dell'Oro context (reported in coverage): in Q3 2025 **Ethernet accounted for more than two-thirds of AI back-end network switch sales**, up from less than half a year earlier; 800G now dominates shipments/revenue and **1.6T is emerging** as the next standard for AI clusters; Goldman Sachs projects the **AI Ethernet switch market reaching ~$24B by 2029** [secondary via gate.com — Dell'Oro/Goldman Sachs attribution].
- Source: https://www.gate.com/blog/arista-networks-anet-revenue-surpasses-3-billion-stock-pullback-ai-network-growth-sustainability-analysis

### R4-C. Huawei Xinghe AI Fabric 2.0 + CloudEngine XH9230-128DQ-LC (MWC Barcelona 2026; new competitive angle)

- At **MWC Barcelona 2026 (March 5, 2026)**, Huawei's AI DC Innovation Forum unveiled the upgraded **Xinghe AI Fabric 2.0** solution (upgrade from the AI Fabric line introduced in 2018) and the industry's first commercial **51.2T (128 × 400GE) liquid-cooled fixed switch — CloudEngine XH9230-128DQ-LC** [vendor-reported — prnewswire, thefastmode, 2026-03].
- Fabric 2.0 three-layer architecture: Intelligent Brain / Intelligent Connectivity / Intelligent Network Elements; four capabilities: Rock-Solid Architecture 2.0 with AI Eagle-Eye Engine (monitors 200,000 service flows in real time; fault detection in seconds, recovery in minutes); StarryWing Digital Map 2.0 with NetMaster (automation); Xinghuan AI Turbo 2.0 (NPLB network packet load balancing + NSLB network stream load balancing); iFlashboot 2.0 (ultra-fast reboot within 5 seconds) [vendor-reported].
- XH9230-128DQ-LC: **100% liquid cooling for optical modules, claimed 2× industry-average heat-dissipation efficiency**; supports 8 switches per cabinet (claimed 2× cabinet utilization efficiency); paired with full series of **800GE/400GE StarryLink optical modules** (claimed 2× industry-average reliability, ultra-long transmission) [vendor-reported].
- European rollout: Huawei launched Xinghe AI Fabric 2.0 in Europe later in 2026 (customer refs: Migros Turkey, Cineca Italy; Sapienza University of Rome best-practices input); star products: XH9230 128×400GE liquid-cooled switch + StarryLink modules [vendor-reported — prnewswire 302571333].
- Competitive read-through: Huawei answers NVIDIA Spectrum-X and Arista's 400/800G fabric with a liquid-cooled 51.2T fixed platform and explicit load-balancing (NPLB/NSLB) aimed at AI collective traffic — relevant to China-market AI fabric RFPs (per project scope: Chinese models/vendors tracked separately) [independent assessment].
- Sources: https://www.prnewswire.co.uk/news-releases/huawei-unveils-the-upgraded-xinghe-ai-fabric-2-0-solution-for-the-ai-era-302705193.html ; https://www.thefastmode.com/technology-solutions/47518-huawei-introduces-xinghe-ai-fabric-2-0-to-accelerate-ai-data-center-networks ; https://www.prnewswire.co.uk/news-releases/huawei-launches-xinghe-ai-fabric-2-0-in-europe-accelerating-enterprise-digital-transformation-302571333.html

### R4-D. Quantum-X Photonics 2026 ship status + the June 2026 CPO controversy (new; partially closes §7 item 5)

- **Controversy (June 9, 2026):** SemiAnalysis published a report flagging **CPO packaging-yield risk** (scenario: 0.95^32 ≈ 19.4% system yield for a 32-engine switch ASIC), triggering a photonics selloff (AAOI −17%, COHR −11%, LITE −8% that day) [secondary — github life-os research memo summarizing 2026-06-09/2026-07-02 events].
- **NVIDIA response:** CTO Gilad Shainer publicly disputed the yield thesis on June 9, stating **CPO is already shipping and ramping in H2 2026** [secondary — same memo; vendor-reported via CTO statement].
- **Independent reality check:** Yole Group forecasts **large-scale CPO deployment realistically in 2028–2030, not 2026–2027**; the sector assessment treats NVIDIA's H2 2026 ramp claim as disputed/unverified at volume [secondary — same memo].
- Product layer detail consistent with prior passes: Quantum-X Photonics = InfiniBand XDR-class CPO (144×800G, 115T + SHARP v4 14.4 TFLOPS in-network compute, liquid-cooled; announced "early 2026"); Spectrum-X Photonics = SN6810 (128×800G, 102.4T) + SN6800 (512×800G, 409.6T), both H2 2026 per original GTC 2025 announcement [secondary — dataconomy, techspot, biggo summaries of NVIDIA announcement].
- **Status conclusion:** shipping-vs-ramp dispute unresolved; no independent 2026 deployment/customer evidence for Quantum-X Photonics located [unverified — §7 item 5 remains open, now better characterized].
- Sources: https://github.com/47tzp4ydc9-cmyk/life-os/blob/HEAD/investment-os/narrative/research/silicon-photonics.md ; https://www.techpowerup.com/334337/nvidia-commercializes-silicon-photonics-with-infiniband-and-ethernet-switches ; https://www.techspot.com/news/109190-nvidia-turns-silicon-photonics-supercharge-next-gen-ai.html

### R4-E. Spectrum-X Photonics ship status — 2026 finding (partially addresses §7 item 4)

- The only 2026 ship-window evidence remains NVIDIA's original announcement: **SN6810/SN6800 shipping H2 2026** (GTC 2025; reiterated in press coverage through 2026). No H2 2026 shipping confirmation, named customer, or GA press release located as of September 22, 2026 [unverified — §7 item 4 remains open].
- Technical anchors restated for completeness (not new): SN6810 = single Spectrum-X CPO device, 128×800G, 102.4T; SN6800 = 512×800G, 409.6T; multi-chip package (central packet-processing engine + SerDes chiplets), 224G/lane, liquid-cooled; 3.5× power efficiency / 9W vs 30W per port vendor-claimed [secondary — techpowerup, nextplatform].

### R4-F. xAI–Dell $5B deal — 2026 status finding (§7 item 3 remains open)

- Searched September 2026 news: **no 2026 confirmation that the xAI–Dell $5B+ GB200 server deal (Bloomberg, Feb 2025) was finalized** [unverified — finding].
- Latest references: a December 2025 article still cites "Dell ... benefiting from a $5 billion hardware commitment and ongoing maintenance agreements" for Colossus expansion, and Morgan Stanley analyst Lewis Cheng upgraded NVIDIA/Dell on xAI procurement cadence [secondary — applyingai.com, 2025-12]. A March 2025 article said xAI "has secured a $5 billion deal with Dell" [secondary — gearmusk.com]. None is a 2026 deal-closing announcement [unverified].
- Context 2026: xAI expanded Colossus (Memphis million-GPU ambitions); Dell raised $5B in senior notes September 15, 2026 (4 tranches, 2029–2037) — financing activity, not deal evidence [secondary — minichart.com.sg].
- Sources: https://www.sdxcentral.com/news/musks-xai-considering-second-data-center-5bn-dell-chip-deal/ ; https://applyingai.com/2025/12/elon-musks-xai-expands-colossus-ai-data-center-near-memphis-3-stocks-set-to-benefit/ ; https://www.minichart.com.sg/2026/09/15/dell-raises-us5-billion-in-senior-notes-offering/

