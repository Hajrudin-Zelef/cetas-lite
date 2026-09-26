---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-round-3a-cisco-fy2
title: "Supplementary / Complementary Research Pass — round 3a: Cisco FY2026 results, NVLink-6 scale-up, HPE Q3 FY2026, Meraki SKUs/API, Dell N3200 OS matrix (2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "AWS", "Nvidia", "Oracle", "United States"]
dates: ["2026-03-22", "2026-07-13", "2026-07-22", "2026-08", "2026-08-12", "2026-09-02", "2026-09-22"]
keywords: ["nvlink", "research", "acquisition", "alignment", "asic", "aws", "backlog", "chiplet", "compute", "custom silicon", "ethernet", "fp8"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [691, 736]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: d869157ed105abeb79b1a5b7605bea3261c87c2273b039fc62f1b47b476f4725
---

# Supplementary / Complementary Research Pass — round 3a: Cisco FY2026 results, NVLink-6 scale-up, HPE Q3 FY2026, Meraki SKUs/API, Dell N3200 OS matrix (2026-09-22)

## Supplementary / Complementary Research Pass — round 3a: Cisco FY2026 results, NVLink-6 scale-up, HPE Q3 FY2026, Meraki SKUs/API, Dell N3200 OS matrix (2026-09-22)

**Scope note:** A base report (§§1–8), a first supplementary pass (§§A–M), a second supplementary pass (§§N–X), and a round-2 pass (§§N–S) already exist. This round-3a pass adds only material not present in any of them, and is distinct from the other two "round 3" sections in this file (written by parallel research passes the same day: §§AA–II and §§Y–AG). All earlier sections were left untouched. Provenance legend from the base applies. All facts as of 2026-09-22.

### R3-A. Cisco FY2026 results (Aug 12, 2026) — networking supercycle numbers

None of the earlier sections contain Cisco's FY2026 results; this fills that gap [secondary — multiple trade summaries of the official release, 2026-08-12/13]:

- **Q4 FY2026**: total revenue **$17.3B (+18% YoY)**; non-GAAP EPS $1.22 (+23%); **networking revenue $9.8B (+28% YoY)**; networking product orders **+40%**, eighth consecutive quarter of double-digit growth; total product orders +35% (+25% ex-hyperscalers); security +14%; collaboration +12%; observability +6%; non-GAAP gross margin 66.3% (hardware-mix-driven decline) [secondary — Cisco release via tickeron, fierce-network, financial-news].
- **AI infrastructure**: **$4B in AI infrastructure orders in Q4 FY26** from hyperscale cloud customers; **$9.3B for the full fiscal year** (4.5× FY2025); ~$4B of that converted to FY26 revenue (orders-to-revenue lag, large remaining backlog); FY2027 AI infrastructure revenue estimated at **$7.5B** [secondary — mirrorreview, fierce-network].
- **Full FY2026**: revenue **$63.3B (+12%)**; GAAP net income $13.3B (+30%); non-GAAP EPS $4.33 (+14%); "highest revenue, operating margin, and earnings per employee in 30 years" per Cisco [secondary — cxfoundation].
- **FY2027 guidance**: total revenue **$72.2–73.4B** (~15% growth at midpoint); non-GAAP EPS $5.05–$5.11; Q1 FY27 revenue guided $18.0–18.2B, EPS $1.32–$1.34; 4–5 points of FY27 growth expected from price increases [secondary — tickeron, business-news-today].
- Analyst read-throughs: Morningstar — Cisco is "converting AI demand into revenue faster than anticipated"; Dell'Oro (Jimmy Yu) — no significant lead-time increases for Cisco despite tight supply, "a benefit of higher vertical integration"; 650 Group (Alan Chris DePuy) — Cisco leadership indicated it **plans to move away from merchant silicon completely by FY2029**, leaning into Silicon One; triple-digit YoY orders from 4 separate hyperscalers; ~$1B F4Q26 Acacia optics order [secondary — fierce-network].

### R3-B. NVIDIA NVLink Switch — 6th generation / scale-up fabric (new angle)

The report covers scale-out (Spectrum-X) and InfiniBand scale-out, but NVLink **scale-up** networking was only a table row in §5. New detail, 2026 [official — NVIDIA technical blog, nvidia.com NVLink page; secondary — ServeTheHome, gpusmith]:

- **NVLink 6** (Rubin generation, unveiled CES 2026, full production mid-2026): **3.6 TB/s bidirectional per GPU** (2× NVLink 5's 1.8 TB/s), over **36 links**; **400G custom SerDes**; **>14× PCIe Gen6 bandwidth** [official — NVIDIA technical blog, 2026; secondary — gpusmith, 2026-07-22].
- **Vera Rubin NVL72**: 72 GPUs in a single all-to-all NVLink domain = **260 TB/s aggregate scale-up bandwidth** (vs 130 TB/s for GB200/GB300 NVL72 on NVLink 5); **up to 9 NVLink 6 switches per NVL72 rack**; NVLink 6 switch trays (4 switch chips per tray, 28.8 TB/s per tray) connected via ~5,000-cable NVLink spine [official — NVIDIA technical blog; secondary — naddod.medium, 2026].
- **In-network compute on scale-up**: 14.4 TFLOPS FP8 per switch tray; 130 TFLOPS aggregate in the rack; accelerates all-reduce/reduce/broadcast collectives; 3× lower GPU-to-GPU latency and 10× higher packet rate vs off-the-shelf Ethernet alternatives (vendor-claimed) [official — NVIDIA technical blog].
- **Resiliency features new in gen 6**: hot-swappable switch trays, control-plane resilience, support for partially populated racks, zero-downtime maintenance / rack-level RAS [official — nvidia.com; secondary — servethehome CES 2026].
- **Roadmap**: NVLink scale-up domains to **1,152 GPUs**; optical scale-up alongside copper so NVLink domains can expand beyond a single rack [official — NVIDIA technical blog; secondary — github stock-research memo, 2026-07-13].
- **NVLink Fusion** (2025–2026): NVLink Fusion chiplet lets hyperscalers/custom-ASIC designers drop NVLink connectivity into custom silicon — Vera Rubin NVLink Switch tray + 400G custom SerDes; up to **72 custom ASICs all-to-all at 3.6 TB/s each**; AWS integrating NVLink Fusion for Trainium4 deployment (announced ~2025-11, continuing into 2026) [official — NVIDIA technical blog, "AWS Integrates AI Infrastructure with NVIDIA NVLink Fusion"].
- Note: the NVLink 6 Switch chip is **liquid-cooled** (28.8 TB/s in a tight tray footprint) [secondary — ServeTheHome, CES 2026].

### R3-C. HPE Networking Q3 FY2026 — new 2026 detail beyond base §3.6

Base §3.6 had the headline numbers. New detail from the Sep 2, 2026 earnings call [secondary — ainvest/Zacks/themarketsdaily earnings-call coverage, 2026-09-02/03]:

- Q3 FY26: networking revenue **$2.89B** (+74.9% reported, +10% normalized); operating margin **22%**; normalized orders **+36%** (~3.5× faster than revenue — supply constraints limited conversion); campus & branch revenue **$1.44B** (+8% normalized); routing **$788M** (+23% — "on- and off-ramp AI network infrastructure"); security +12%; **data-center networking revenue declined 6%** on supply-constrained shipment timing [vendor-reported via earnings call].
- **Oracle deal**: HPE announced an **expanded multiyear, multi-gigawatt collaboration with Oracle** — routers, switches (QFX switching, PTX routing), software and AI operations for Oracle's AI cloud infrastructure buildout [vendor-reported via earnings call].
- To meet order growth, HPE **more than doubled networking purchase commitments quarter-over-quarter**; management expects supply alignment to improve in Q4 FY26 but remain a constraint into FY27 [vendor-reported via earnings call].
- Portfolio pipeline: recently launched **direct liquid-cooled Tomahawk-6-based switch**; **Helios platform** "expected to bring further acceleration" (announced as upcoming) [vendor-reported via earnings call].
- DOJ settlement: U.S. federal court approved HPE's settlement with the Department of Justice related to the Juniper acquisition in August 2026; integration ahead of schedule; **$600M annual run-rate of Juniper-related savings targeted by end of FY2028** [secondary — themarketsdaily].

### R3-D. Meraki MS355 — full spec/pricing detail (new)

Base §4.1 had the family-level table. Customs product-information sheets (2026-03-22) give per-SKU specs [official — customs.ciscomeraki.systems; secondary — reseller specs]:

| SKU | Access ports | Uplinks | Switching capacity | Notes |
|---|---|---|---|---|
| MS355-24X-HW | 16× 1G + 8× mGbE (up to 10G) | 4× 10G SFP+ + 2× 40G QSFP+ | 352 Gbps | UPoE 740W; 2× dedicated 100G QSFP28 stack ports (400 Gbps stacking) |
| MS355-24X2-HW | 24× mGbE (up to 10G) | 4× 10G SFP+ + 2× 40G QSFP+ | 640 Gbps | UPoE 740W |
| MS355-48X2-HW | 24× 1G + 24× mGbE (up to 10G) | 4× 10G SFP+ + 2× 40G QSFP+ | 688 Gbps | UPoE 740W (60W/port); 512 Mpps; 32K MAC; dual hot-swap PSU; 3× hot-swap fans; 8-unit physical stacking |

