---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-round-4-2026-09-22-3
title: "Supplementary / Complementary Research Pass — round 4 (2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-05", "2026-07-25", "2026-08-12", "2026-09-22"]
keywords: ["research", "cost", "hyperscaler", "license", "pricing", "revenue"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1312, 1359]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: de4808eb206cd13da1da4c3c66fff8d6dd4ee308aed4d4ed20a1fe5720abc43b
---

# Supplementary / Complementary Research Pass — round 4 (2026-09-22)

## Supplementary / Complementary Research Pass — round 4 (2026-09-22)

**Scope note:** Base §§1–8 plus earlier supplementary passes already exist, including parallel "round 4" sibling passes that were appended concurrently by other research processes (covering CVE-2026-23813/MS390/C9550 among other topics — I checked for overlap and kept only items unique to this pass: §§JJ and NN). Nothing above this section was modified. The provenance legend from the base applies.

### JJ. Cisco Q4 FY2026 results — August 12, 2026 (new networking context for the Meraki parent)

Cisco's Q4 FY2026 (quarter ended July 25, 2026; reported 2026-08-12) was a beat-and-raise print driven by AI infrastructure demand — relevant context for Meraki's enterprise strategy (§4) [independent — Morningstar/Dow Jones Newswires; WSJ]:

- Revenue **$17.25B (+17.6% YoY)**, vs $16.84B consensus; GAAP profit **$3.86B ($0.97/share)**, vs $2.55B ($0.64) a year earlier; adjusted EPS **$1.22** vs $1.17 expected [independent].
- Product revenue **+24% YoY**; services flat [independent].
- **AI infrastructure orders from hyperscalers: ~$4B in Q4 alone**, bringing the **FY2026 total to $9.3B** — nearly 5× the FY2025 total of $2B (which itself had doubled the original $1B target) [vendor-reported — earnings call via Morningstar].
- FY2027 guidance: revenue **$72.2–73.4B**, adjusted EPS **$5.05–5.11** (both above consensus $69.12B / $4.83); Q1 FY27 guided revenue $18.0–18.2B [vendor-reported].
- Workforce: Cisco said in May 2026 it would cut **<5% of headcount** (thousands of roles) to reallocate resources toward AI; restructuring cost up to ~$1B [independent — WSJ].
- Context for §4: Cisco's networking business is growing on hyperscaler/AI-fabric orders even as the Meraki line (enterprise access/aggregation, covered §§4.1–4.4) stayed hardware-static through 2026 — the two faces of Cisco's networking story [independent assessment].

### NN. AOS-CX Feature Pack subscription licensing (new detail)

A subscription model detail missing from base §3 — HPE's CX feature-pack structure as documented in 2026 [official — HPE Aruba feature-pack deployment guide]:

- Advanced AOS-CX features are subscription-gated via **per-platform advanced feature packs** (1/3/5/7/10-year terms plus 90-day eval), covering: application-based policies; reflexive policies (port-access/GBP clients); MACsec extensions for the WAN (8360, 9300S only); queue statistics monitoring (8325/8325H/9xxx); inband flow analyzer and flow telemetry (9xxx); queue congestion monitoring (8325/8325H/8325P); hosting of HPE-certified applications (10040/9xxx) [official].
- Relevance: Aruba's "all-inclusive licensing" claim (base §5.4) applies to base AOS-CX; advanced telemetry/MACsec/flow features are subscription add-ons — nuance for the CX pricing picture [independent assessment].

### OO. Round-4 verification log (new open items)

1. Cisco Q4 FY2026 networking-segment revenue split (networking vs security vs collaboration) — product revenue +24% confirmed; segment breakout not extracted in this pass [gap].
2. AOS-CX 10.17 feature delta vs 10.16 (what's new beyond security fixes) — release notes not parsed [gap].
3. Meraki MS210/MS225 street-pricing anchors — not yet located in any pass [gap].
4. C9550 street pricing and license tiers (Essentials/Advantage) — not located [gap].
5. Enterprise SONiC H2 2026 release (4.6.1/4.7.0) — status not checked in this pass [gap].
6. Aruba CX 6300/6400 access-tier street pricing — not systematically located [gap].
7. Overlap audit across this file's concurrent supplementary passes (round 3b / parallel round-4 passes / Pass #3) — not done; some topics (e.g. MS390 pricing, C9550, AOS-CX CVEs) are covered in more than one section [known gap — flagged for the curator].

### PP. Round-4 sources (verbatim URLs)

- https://www.morningstar.com/news/dow-jones/2026081211847/cisco-reports-higher-4q-profit-as-ai-orders-roll-in
- https://www.wsj.com/business/earnings/cisco-reports-higher-fourth-quarter-profit-as-ai-orders-roll-in-e82db3b4
- https://WWW.ZACKS.COM/stock/news/2712980/ciscos-q4-earnings-beat-estimates-revenues-rise-yy-shares-down
- https://arubanetworking.hpe.com/techdocs/AOS-CX/10.16/PDF/feature-pack.pdf

**Round-4 collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in §OO above. Existing sections §§1–8, §§A–M, §§N–X, "round 2" §§N–S, and "round 3" §§AA–II were not modified.


---

## Supplementary / Complementary Research Pass #3 — 2026-09-22

**Scope note:** Two additive passes already exist above (Supplementary Pass §§A–M and Supplementary Pass #2 §§N–X, plus a "round 2" §§N–S). This third pass adds only material not present in §§1–8 or any earlier pass. Earlier sections were left untouched. Read the base report and earlier passes first; the provenance legend from the base applies here.

