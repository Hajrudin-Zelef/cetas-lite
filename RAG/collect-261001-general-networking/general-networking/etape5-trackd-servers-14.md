---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-14
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Meta", "Microsoft", "Nvidia", "Oracle", "xAI"]
dates: ["2026-06", "2026-07", "2026-07-16"]
keywords: ["cpo", "dci", "ethernet", "gpu", "gpus", "hyperscaler", "latency", "memory", "neocloud", "nvidia", "optics", "revenue"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [888, 959]
sha256: 73388204d4b90ad04262282584df20fd92fc12278c81d4c2131525209f5951dd
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

- At Q2 2026: management raised the **AI-fabrics revenue target to at least $3.5B for FY2026** (more than double the prior year); Zacks (Sept 16, 2026) quotes "at least $3.6B" supported by scale-up, scale-out and scale-across deployments. [vendor-reported] [secondary]
- Overall FY2026 revenue outlook raised to **approximately $11.5B** (≈27.7% growth) per August-2026 coverage. [secondary] (cryptobriefing, fxempire)
- **Conflict flag:** remio.ai (published ~Sept 19, 2026) states Arista "raised its 2026 revenue target to $12.6 billion" after Q2, with Q3 guidance ~$3.3B — this conflicts with the $11.5B figure from August coverage. No corroborating source found. Possible misread (e.g., annualizing Q3 guide) or a later raise not yet covered elsewhere. **Treat $12.6B as [unverified]/conflicting.**
- Etherlink AI customer base: **over 100 customers** (vs 4–5 in 2024); Microsoft and Meta identified as key customers. [vendor-reported] [secondary]
- Consensus estimates (TIKR, Aug 2026): $12.67B (2026), $16.16B (2027), ~$26.54B by 2030; 2021–2025 revenue grew $2.95B → $9.0B. [secondary]
- Zacks (Sept 16, 2026): stock +47% YTD 2026; 52-week range $114.52–$214.89; AI revenue at least $3.6B in 2026; projected FY26 revenue growth ~39%. [secondary]

### 6.3 Product launches (2026)

- 1.6 Tbps AI fabric platforms (7060XE7 series), incl. liquid-cooled options; 7506c/7 Series AI fabric portfolio; 7th-gen 7280R series for AI workloads. [vendor-reported] [secondary]
- Arista positioned across scale-up, scale-out, and scale-across networks. [official]

### 6.4 xAI/Colossus relationship

- xAI Colossus chose **NVIDIA Spectrum-X Ethernet**, not Arista — a design win for NVIDIA, not Arista. Arista's hyperscaler AI exposure is via Microsoft, Meta, and other cloud builders (RoCE/Ethernet fabrics, e.g., Meta's Arista 7800-based 24,576-GPU cluster). [official] [secondary]
- Competitive pressure: NVIDIA's integrated stack (GPUs + Spectrum-X + SuperNICs + optics + software) is the structural challenge to Arista's open-Ethernet position, per secondary analysis. [secondary]

---

## 7. Cisco: AI networking (Silicon One, Nexus), 2026 wins and revenue

### 7.1 FY2026 results (year ended July 2026; reported ~Aug 2026)

- Q4 FY2026 revenue **$17.25B** (+18% YoY, record); full-year revenue **$63.33B** (+12% vs $56.65B FY2025); product revenue $48.30B, services ~$15.03B. [secondary] (infotechlead)
- **AI infrastructure orders from hyperscalers: $9.3B in FY2026** (~4.5× the ~$2B of FY2025), incl. **$4B in Q4 alone** (record quarter for AI orders). [vendor-reported]
- Order mix: ~60% Silicon One-based systems, ~40% optics. [vendor-reported]
- **Recognized AI infrastructure revenue from hyperscalers: ~$4B in FY2026** (vs $1B FY2025); AI infra was ~6% of total revenue (vs <2% in FY2025). [vendor-reported]
- **FY2027 guide: total revenue $72.2B–$73.4B (~15% growth at midpoint); AI infrastructure revenue from hyperscalers $7.5B** (CFO said the rest of the business would still grow ~10%); Cisco retired the order-booking target framework in favor of a formal revenue target. [vendor-reported] (trefis, ainvest)
- Price increases added ~5 points to FY2026 Q4 revenue growth (CFO); 4–5 points planned for FY2027 (front-half weighted); memory costs and hardware mix pressuring gross margin. [vendor-reported] (trefis)
- Non-hyperscaler AI: >$400M of AI infra orders from neocloud, sovereign and enterprise customers in Q4; full-year >$1B from those customers. [vendor-reported]

### 7.2 Design wins (Q4 FY2026)

- Three new hyperscaler AI design wins: one **Silicon One P200 scale-across** deployment, one **G200 scale-out** project, one **optical line-system** deployment; four leading hyperscale customers; line of sight to more wins over six months across G300, G200, P200, A100 and optics. [vendor-reported]
- Acacia optics business: >$1B of orders in a single quarter (Q4). [vendor-reported] (ainvest)
- Products: Silicon One **G300** (102.4 Tbps) in liquid-cooled N9000 and 8000 series; **N9100** series powered by NVIDIA Spectrum-X silicon with NX-OS; Series 8000 routers; Acacia coherent pluggables. [secondary] [vendor-reported]

### 7.3 Strategic position / caveats

- Cisco lost the DC switching crown to Arista; NVIDIA then overtook both in DC Ethernet switching share (Dell'Oro figures cited by secondary press). [secondary]
- Commentary (ainvest, editorial): Cisco's strategy mixes own-silicon (G300) with reselling NVIDIA's Spectrum-X platform — described as defensive. [secondary, opinion]
- Margin caution: the AI orders are in a competitive, buyer-leverage-heavy segment; Cisco's operating margin (TTM 25.4%) vs Dell 9.6% noted as the offset. [secondary] (trefis)

---

## 8. Data center interconnect: 800G adoption, 1.6T timeline

- **800G is the mainstream AI-fabric generation:** "the vast majority of Ethernet switch shipments and revenues in AI backend networks" in Q1 2026 were 800G (Dell'Oro). [independent]
- **Optical volume:** 800G+ shipments 24M units (2025) → ~63M (2026) (TrendForce). [secondary]
- **1.6T timeline:** 1600G switches "only beginning to sample," expected to "ramp in the second half of 2026" (Dell'Oro, June 2026). 1.6T optical shipments 2.5M (2025) → 20M+ by end-2026 (OFC-2026 analyst confirmation); "1.6T is shipping now" with NVIDIA and Google integrating, Meta and Oracle slated next (eetimes, Sept 2026). [independent] [secondary]
- IEEE 802.3dj (200G/lane, covering 200G/400G/800G/1.6T) on track for completion late 2026; early 200G/lane products expected during 2026; 400G/lane project next. [independent] (networkworld)
- **DCI / scale-across:** emerging as a distinct revenue layer — NVIDIA Spectrum-XGS (available now, SW/FW upgrade), Cisco P200 silicon for scale-across, Arista 1.6T scale-across platforms. [official] [vendor-reported]
- **3.2T:** development underway around 400G-per-lane designs (Kozlov, eetimes); Spectrum-X3200/Quantum-X3200 on NVIDIA's roadmap (TrendForce). [secondary]

---

## 9. Key uncertainties and explicit flags

1. **Arista FY2026 guide conflict:** $11.5B (Aug 2026 coverage) vs $12.6B (remio.ai, Sept 19, 2026). Unresolved. [unverified]
2. **UEC 1.0.3 currency:** the "1.0.3 (July 16, 2026) current" claim comes from a community wiki, not ultraethernet.org. Verify before citing. [unverified]
3. **UEC production adoption:** no confirmed hyperscaler production deployment of UEC/UET found; 2026 evidence is interop demos, NICs, switches, test gear. [unverified]
4. **NVIDIA silicon-photonics/CPO shipping:** roadmap language (Spectrum-X Ethernet Photonics H2 2026; Quantum-X early 2026) vs SemiAnalysis yield concerns; Shainer's "shipping and ramping H2 2026" is a vendor claim. Large-scale CPO volume realistically 2028–2030 per Yole. [unverified]
5. **Quantum-3:** no evidence located; roadmap shows Quantum-2 → Quantum-X800 → Quantum-X1600. Do not assert a Quantum-3 product. [unverified]
6. **Spectrum-XGS announcement date:** "available now" release; Hot Chips mention suggests Sept 2025 but date not verified from NVIDIA's page. [unverified]
7. **Market-size figures conflict:** transceiver TAM estimates differ widely by firm (LightCounting ~$23.8B 2025 vs MarketsandMarkets ~$9.2B DC 2025). Treat as directional. [secondary]
8. **DriveNets claims** (6% better than IB, 15% better than Spectrum-X on NCCL): single vendor blog using SemiAnalysis test data, unaudited. [vendor-reported]
9. **Quantum-X800 latency/bandwidth table figures** (ascentoptics: <100 ns, 115.2 Tb/s vs 51.2 Tb/s): secondary product-page comparison, internally inconsistent with other latency quotes; quote with caution. [secondary]

---

## 10. Sources (URLs as found)

