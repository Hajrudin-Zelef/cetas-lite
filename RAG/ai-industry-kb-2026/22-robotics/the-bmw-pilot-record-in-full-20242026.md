---
id: ai-industry-kb-2026/22-robotics/the-bmw-pilot-record-in-full-20242026
title: "The BMW pilot record in full (2024–2026)"
domain: robotics
role: deep-dive
task: robotics
actors: ["China", "Intel", "Microsoft", "Nvidia", "OpenAI", "Qualcomm", "Samsung", "xAI"]
dates: ["2025-03", "2025-09", "2026-01", "2026-05-26", "2026-06", "2026-06-23", "2026-07-26", "2026-08", "2026-09", "2026-09-16"]
keywords: ["benchmark", "consumer", "distribution", "gpus", "humanoid", "intel", "lawsuit", "liability", "nvidia", "robotics", "series c", "throughput"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [10619, 10655]
section: "22. Robotics"
sha256: f74e16169b7b70697732f49b0ce31e26ee3234fdf26453c1b000879b2e04c2fa
---

# The BMW pilot record in full (2024–2026)

- **2024-01 — Figure 02 deployment begins at BMW Spartanburg [VERIFIED].** (Pre-period baseline.)
- **2025-02 — OpenAI partnership ended; in-house Helix VLA announced [VERIFIED].** (Pre-period baseline.)
- **2025-03 — BotQ manufacturing facility unveiled: 12K units/yr line capacity [VERIFIED, VENDOR].** (Pre-period baseline.)
- **2025-09 — Series C >$1B at $39B**, led by Parkway Venture Capital + Brookfield; investors incl. NVIDIA, Microsoft, Bezos, Intel, Samsung, Qualcomm, Salesforce [VERIFIED]. (Pre-period; included because the flat secondary is a 2026 market signal.)
- **2025-10 — Figure 03 unveiled**: 35 DOF, tactile-sensor hands, palm cameras, $20K target price, TIME Best Invention 2025 [VERIFIED]. (Pre-period baseline.)
- **2025-11 — BMW Spartanburg Figure 02 pilot completes: 11 months, 30,000+ vehicles, 90,000+ parts; federal whistleblower safety lawsuit filed [VERIFIED].** The lawsuit is a live watch item; no outcome verified in this pass.
- **2026-01 — Helix 02 released [VENDOR]:** billed as "full-body autonomy." Architecture: System 2 = 7B VLM running 7–9 Hz; System 1 = 80M-parameter policy at 200 Hz; fully onboard on dual GPUs.
- **2026-04 — BotQ at ~1 robot/90 min (~240/month); reported as a 24× ramp in 120 days to ~1 robot/hour [VENDOR].** CEO-announced snapshot; not audited throughput. [SHIPMENT-class claim at best.]
- **2026-05-26 — Commercial deal with Catalyst Brands**: large-scale deployment at a Reno retail distribution center (Brookfield portfolio link) [SECONDARY]. [PAID DEPLOYMENT per reporting; scale unverified.]
- **2026-06-23 — ~740 robots deployed vs. 660 employees** ("more robots than staff" milestone) [VENDOR, CEO-announced].
- **2026-06 — 40 Figure 03 units deployed at BMW Spartanburg** in logistics sequencing (sorting unsorted components into sequencing trolleys, handed to tugger trains/STRs); Leipzig pilot announced for summer 2026 — billed as the first Physical AI deployment in European automotive production [SECONDARY via carscoops, quoting BMW VP Ulrich Wieland]. [PILOT-to-PAID-DEPLOYMENT transition class; do not present as unsupervised.]
- **2026-07-26 — 1,000th Figure 03 built at BotQ [VENDOR].** [SHIPMENT milestone.]
- **2026-08 — Figure 03 autonomous ladder-climbing demo**; CEO Adcock declares wheeled robots an "utter dead end" [DEMO, VENDOR].
- **Corrections:** The Figure 02→03 transition at BMW was a **use-case change** (body-shop sheet-metal insertion → logistics sequencing), not just a hardware swap. Setter's $39B secondary valuation (Q2 2026) equals the September 2025 primary — **flat, not up**.
- **Figure 03 spec snapshot [VENDOR]:** 5'6", 60 kg, 35 DOF, 100 N·m joint torque, 4.9 kg hands, 2.3 kWh battery, $20K target price.

### The BMW pilot record in full (2024–2026)

- **Jan 2024 — Figure 02 deployment begins at BMW Spartanburg [VERIFIED]** (pre-period baseline): body-shop sheet-metal insertion use case.
- **Nov 2025 — Figure 02 pilot completes [VERIFIED]:** 11 months, 30,000+ vehicles, 90,000+ parts handled. The only completed Western automotive humanoid pilot with public numbers — the benchmark against which all 2026 "factory deployment" claims must be read.
- **Nov 2025 — federal whistleblower safety lawsuit filed [VERIFIED filing, outcome unknown]:** live watch item. Its resolution is the liability precedent for humanoid factory pilots; do not assert an outcome.
- **June 2026 — 40 Figure 03 at BMW Spartanburg [SECONDARY]:** use-case change from sheet-metal insertion to logistics sequencing (sorting unsorted components into sequencing trolleys, handed to tugger trains/STRs). The hardware changed (02→03) but the more important change was the task — logistics sequencing is a more structured, more supervisable task than body-shop insertion.
- **Summer 2026 — Leipzig pilot announced [SECONDARY]:** billed as the first Physical AI deployment in European automotive production (carscoops, quoting BMW VP Ulrich Wieland). Execution unverified in this pass.
- **Reading [DIRECTIONAL]:** the BMW record is the industry's best-documented deployment arc — a completed pilot with numbers, a live safety suit, a second-generation rollout on an easier task, and a European expansion. It is also the record that shows how slow "deployment" really is: 11 months for the first pilot, with the 40-unit follow-on still task-supervised. Any 2026 claim of faster deployment elsewhere should be read against this timeline.

### Tesla Optimus: the 2026 factory year — read against the forecast record

- **The forecast record [VERIFIED via RoboZaps timeline review, citing Electrek/AP/Omdia]:** "~10,000 robots in 2025" (Q4 2024 call) → Omdia via AP: fewer than 500 shipped in 2025; DigiTimes-derived: ~1,000 assembled by mid-2025. "5,000 in 2025, ~50,000 in 2026" (March 2025 all-hands) → 2025 missed several-fold; 2026 guidance withdrawn as "impossible to predict." "Robots doing useful factory work" (2024) → **Musk, January 2026: "not in usage in our factories in a material way."** "Production start late July/August 2026" (Q1 2026 call) → softened on July 22 to "soon"/"later this year."
- **2026-05 — Tesla winds down Model S/X at Fremont** to convert floor space into a dedicated Optimus line; JPMorgan analysts (August tour) confirmed ~4-month conversion, targeting **1M units/yr design capacity** [SECONDARY]. **Design capacity is not production** — never cite "1M Optimus" as production.
- **2026-09-16/17 — Tesla robotics team audits Ningbo suppliers** (Bloomberg via RobotAIGeek, cited by Teslarati): certified three Chinese mass-production partners — Tuopu Group (actuators/chassis), Ningbo Joyson Electronic (sensors), Zhejiang Sanhua Intelligent Controls (thermal). Fresh orders reported as an initial batch of ~5,000 units [SECONDARY]. Sanhua context: Teslarati reported (Oct 2025) a ~$685M Sanhua linear-actuator order, estimated by watchers at ~180,000 robots' worth [SECONDARY].
- **September 2026 supply-chain reports [SECONDARY — nextbigfuture/teslarati; treat as vendor-adjacent rumor until Tesla confirms]:** ~1,000 units/week by late September, ramping to 2,000–2,500/week by end-2026 → 100,000–125,000/yr run rate into 2027.
- **Optimus Gen 3 teased** by Musk ("so many improvements") — capabilities undisclosed [VENDOR].
- **Optimus specs [VENDOR]:** ~5'8", 57 kg, 22 hand DOF, $20,000–$30,000/unit target price; consumer availability "end of 2027" conditional on reliability/safety/functionality. Status: **pre-production** (PILOT at most; no material factory usage admitted).
- **RAG discipline:** the 2025 miss record is the correct Bayesian prior for any Tesla robot date. Discount his dates per his record.

### 1X Neo: the first consumer humanoid deployment

