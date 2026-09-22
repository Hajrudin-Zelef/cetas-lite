---
id: ai-industry-kb-2026/22-robotics/key-dated-facts
title: "Key dated facts"
domain: robotics
role: deep-dive
task: robotics
actors: ["AWS", "Anthropic", "Apple", "California", "China", "Google", "Hugging Face", "Intel", "Microsoft", "Nvidia", "OpenAI", "Qualcomm", "Samsung", "United States", "xAI"]
dates: ["2024-06", "2025-03", "2025-09", "2025-09-25", "2025-10-28", "2026-01", "2026-01-14", "2026-03-27", "2026-04", "2026-05-26", "2026-06", "2026-06-01", "2026-06-23", "2026-06-24", "2026-07", "2026-07-16", "2026-07-26", "2026-08", "2026-08-19", "2026-09", "2026-09-01", "2026-09-09", "2026-09-16", "2026-09-18", "2026-12"]
keywords: ["acquisition", "benchmark", "blackwell", "chatgpt", "consumer", "disclosure", "distribution", "energy", "fine-tuning", "gemini", "gpus", "humanoid"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [10584, 10745]
section: "22. Robotics"
sha256: 7db7913c2f23a34c203613d81b56be5c94733ee30879f014461d88924e5fef3c
---

# Key dated facts

## Key dated facts

### The 2026 taxonomy rule (applies to every claim in this chapter)

Robotics reporting in 2026 systematically blurs six distinct things. This chapter enforces these labels:

- **DEMO** — scripted or teleoperated showing; no customer revenue. Examples: Figure 03 ladder climbing (Aug 2026, [VENDOR]); the T800 cage fight (2026-09-18, remote-piloted per the organizers' own disclosures); IFA 2026 punch feints.
- **PILOT** — customer-site evaluation, free or paid-for-data. Examples: BMW Figure 02 (2024–2025, completed); Airbus purchasing Walker S2 units for concept testing (Jan 2026); Amazon's Digit activity (R&D only).
- **PAID DEPLOYMENT** — paying customer on commercial terms. Examples: GXO Spanx Digit RaaS (running since June 2024, 100,000+ totes moved); Toyota Motor Manufacturing Canada's 7-Digit commercial agreement (Feb 2026); Catalyst Brands/Reno.
- **SHIPMENT** — units that left the factory. Examples: Unitree 5,511 shipped (2025); the 1,000th Figure 03 at BotQ (2026-07-26); ~19,100 global H1 2026 (Smart Analytics Global).
- **SALE** — revenue recognized. Example: Unitree 5,215 sold (2025) — 296 fewer than shipped.
- **PRODUCTIVE DEPLOYMENT** — doing real, useful, unsupervised work at scale. Examples: SemiAnalysis: ≤250 Unitree units (2025); BMW Spartanburg: 40 Figure 03 in logistics sequencing (June 2026).

Rule of thumb: classify **one level down from wherever the headline puts it** — toward the stricter category — until a filing or the customer confirms the higher level. Headlines that say "deployed in factories" almost always mean shipped-to-a-pilot or teleoperated.

### The SPAC as an instrument (Agility, 2026-06-24)

- **Mechanics [VERIFIED]:** merger with Churchill Capital Corp XI at ~$2.5B valuation; expected gross proceeds ~$600–620M (~$420M trust + $200M PIPE led by Foxconn); ticker AGLT; expected close later 2026. The first pure-play humanoid robotics listing on a US exchange.
- **The $300M contracted Digit v5 orders [VENDOR]:** multi-year RaaS agreements with warrants vesting on deployment — the SPAC's revenue story is forward-contracted, not recognized. GeekWire's filing review (revenue "a small fraction" of the implied valuation) is the counterweight.
- **Amazon's posture inside the SPAC record:** $150M invested in the 2022 Series B; spokesperson confirmed deployments held for v5 (cooperative safety for high-traffic warehouses). The largest potential customer is explicitly waiting — a disclosed risk factor for the RaaS story.
- **The Foxconn $200M PIPE [VERIFIED]:** the contract manufacturer that builds the world's iPhones is buying into humanoid RaaS — and separately appears as a UBTech Walker S2 customer [SECONDARY]. Foxconn is long humanoids on both sides of the Pacific.
- **Reading [DIRECTIONAL]:** the SPAC prices paid hours (65,000+, 100,000+ totes) as a growth asset. The close — expected later 2026 — is the first public-market referendum on whether RaaS hours convert to RaaS revenue at venture valuations. Score it against the December Digit v5 deliveries, not the announcement-day headlines.

### Unitree: IPO and the filing that reframes the industry

- **2026-08-19 — Unitree IPO, STAR Market [VERIFIED]:** Listed ~40.45 million shares at RMB 150.80 (10% float) for an issue valuation of ~RMB 61B (~$9B). First day: ~600% intraday spike, close ~+460% → **~$51B market cap**. The raise was ¥6.1B (~$904M); this is not the market cap. [DEMO-vs-deployment classification: the pop is a capital-markets event, not a capability signal.]
- **Filing (2025 figures, anchor evidence):** 5,716 produced / 5,511 shipped / 5,215 sold (pure humanoids, wheeled dual-arm excluded); RMB 867.8M humanoid robot revenue; average recognized price RMB 166,400/unit; Jan–Sep 2025 revenue mix **73.6% research/education, 17.39% commercial/consumer, 9.01% industry**; enterprise reception/guidance was 50–70% of the industry-application slice; only RMB 15.702M (29.29% of the industry slice) from clearly operational scenarios (manufacturing, inspection, logistics). SemiAnalysis (June 2026): ≤250 units in productive industrial pilots/deployments in 2025.
- **2026-06-01, Computex — NVIDIA + Unitree "Isaac Root" research platform [VERIFIED/VENDOR]:** Unitree H2 as flagship hardware for NVIDIA's Isaac Root initiative, with Jetson Thor (Blackwell) and Isaac GR00T models; 31 DOF (25 in Sharpa Robotics dexterous hands); targeting Ai2, ETH Zurich, Stanford, UCSD. Notably no Chinese institutions were listed.
- **2026 lineup & US prices (Aug 2026 listings [SECONDARY]):** R1 AIR $4,900 (cheapest walking humanoid), R1 $5,900, G1 $13,500, H2 $29,900, H2 Plus $100,000; Go2 quadruped from $1,600.
- **H2 specs [VENDOR/second-hand]:** 6 ft, ~150 lb, Intel Core i5 base, optional Jetson AGX Thor (2,070 TOPS) on the Edu variant, 360 N·m leg torque, 3-hour quick-swap battery.
- **H1 2026 shipments:** ~5,900 (SAG) / >7,000 (Counterpoint); R1 shipments began January 2026.
- [PILOT vs SHIPMENT classification: "thousands sold" is filing-verified, but the revenue-mix data reclassifies most of them as research/education PILOT-class usage, not industrial deployment.]

### Figure AI: BotQ ramp, BMW, Helix 02

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

- **2025-10-28 — Neo pre-orders opened** (1X Technologies, Norway/California); WSJ's Joanna Stern hands-on: "I tried the first humanoid home robot" [VERIFIED]. (Pre-period baseline.)
- **2026 — Deliveries began** (company: US priority, Q3 2026 wave per trade press) [VERIFIED as started; wave completion unverified].
- **Specs:** 165 cm, 30 kg, tendon-driven, soft knit exterior, 22 DOF per hand, ~4h battery, NVIDIA chips, Redwood AI VLA.
- **Pricing [VERIFIED]:** **$20,000 purchase or $499/month subscription** ($200 refundable deposit).
- **Expert Mode (disclosed teleoperation):** when Neo cannot do a task, a 1X human expert teleoperates it to teach it — the privacy tradeoff Stern flagged (a company rep may watch through Neo's cameras). Disclosed by the company, not hidden — but it means early "autonomy" demos are teleop-assisted by design.
- **Safety pitch:** compliant tendons, HIC <250 (head-impact criterion), quieter than a fridge, IP44 body / IP68 hands (trade-press spec sheet, unverified primary).
- **The consumer-humanoid evidence bar [DIRECTIONAL]:** Neo clears the bar no industrial humanoid has: a disclosed price, an open order book, and deliveries to paying non-corporate customers. What it does not clear — and no 2026 consumer review claims it does — is unsupervised task completion at human reliability. The privacy tradeoff (Expert Mode) is the honest version of the teleoperation that industrial vendors leave undisclosed. On the taxonomy, Neo is DEPLOYMENT (consumer, revenue-recognized) with teleop-assisted autonomy — the most honestly labeled product of 2026.
- **Competitive context:** no other 2026 consumer home humanoid has verified deliveries. UBTech's U1 consumer line was announced but is not evidenced here as delivered; the Optimus consumer date is "end of 2027" [VENDOR, conditional]. Neo stands alone in its category — a fact that says as much about the category's immaturity as about 1X.
- **Distinguish:** Neo is home-first (tidying, laundry, dishwashing — no cooking, not waterproof body per launch coverage), unlike the industrial humanoids. It is the **only 2026 consumer home humanoid with verified deliveries**. [DEPLOYMENT class — consumer, not industrial.]
- Safety incident footnote (pattern context, 2025): a robot slapping a child mid-dance-demo (viral clips) — included only as context for the teleop-safety discussion.

### Agility Robotics: SPAC, Digit v5, paid hours

- **2026-06-24 — SPAC merger announced [VERIFIED]:** Churchill Capital Corp XI at **~$2.5B valuation**; expected gross proceeds **~$600–620M** (~$420M trust + **$200M PIPE led by Foxconn**); ticker **AGLT**; expected close later 2026. First pure-play humanoid robotics listing on a US exchange. GeekWire flagged that filings show revenue is still a small fraction of the implied valuation.
- **Digit (logistics specialist):** 175 cm, 65 kg, 28 DOF, 16 kg payload, ~8h battery, NVIDIA Jetson AGX Thor. Paid RaaS at **GXO's Spanx facility (Georgia) since June 2024 — 100,000+ totes moved** [PAID DEPLOYMENT]. **Toyota Motor Manufacturing Canada: Feb 2026 commercial agreement for 7 Digits** (Woodstock, Ontario; RAV4 plant) [PAID DEPLOYMENT]. Schaeffler (Cheraw, SC — 25-lb bearing baskets, 8h/day cycle, purchase agreement); Mercado Libre (San Antonio); **Amazon: pilot/R&D only** — spokesperson confirmed Amazon deployments held for v5 (cooperative safety for high-traffic warehouses); Amazon invested $150M in the 2022 Series B.
- **Scale:** **65,000+ operational hours across 9 customer facilities** [VENDOR]; outside observers estimate dozens of units — exact unit counts undisclosed.
- **Digit v5 (H2 2026) [VENDOR]:** 23 kg payload, 22-hour battery, 7.2 ft reach, **NVIDIA Halos safety architecture** (first humanoid with Halos; announced June 2026 at Automate) — marketed as the first "cooperatively safe" humanoid (no safety cage); early customers from **December 2026**. **$300M in contracted v5 orders** (multi-year RaaS incl. warrants vesting on deployment) [VENDOR].
- **RoboFab (Salem, Oregon):** 70,000 sq ft, 10,000 units/yr peak capacity [VENDOR, SHIPMENT capacity not output]. **Fremont, CA 60,000 sq ft AI training hub opened July 2026** (near Tesla's Optimus factory — CEO Peggy Johnson: "having Tesla in the same area is a good thing").
- **Pricing discipline:** SPAC deck models **~$8,500/month per Digit** RaaS (illustrative); the "$30/hr" figure is the human-labor comparator, not Digit's price; $250K was the 2020 purchase price; $45K listings are wrong (RoboZaps).

### UBTech Walker S2: industrial-deployment leader by audited units

- **FY2025 filing [VERIFIED]:** **1,079 full-size humanoids delivered** (>80% industrial); humanoid segment revenue **RMB 820.6M** (+2,204% YoY, now the largest segment); total revenue RMB 2.0B (+53%); net loss narrowed to RMB 790M. Listed HKEX 9880.
- **Walker S2** (mass-production model, deliveries from Nov 2025): autonomous battery swap, 4th-gen hands, reported 52 DOF [SECONDARY]; customers: **BYD, Geely, FAW-Volkswagen, Audi FAW, BAIC New Energy, Dongfeng Liuzhou, Foxconn, SF Express, Texas Instruments**; **Airbus purchased units Jan 2026** for early concept testing [PILOT]; **¥264M (~$37M) contract for customs/patrol at China–Vietnam border crossings** (deliveries from Dec 2025, work underway mid-2026) [SECONDARY].
- **Capacity guidance:** 5,000 (2026) → 10,000 (2027); Dongfeng Liuzhou factory hitting a **10-minute production pace** (Sept 2026) [SECONDARY]; 1,000th Walker S2 off the line targeted late Dec 2025.
- **No Walker S3/S2 Pro** announced as of July 2026 — investor-blog names only. Newest launches are Walker S2 (industrial) and the U1 consumer line.
- **Specs [SECONDARY]:** 176 cm, 76 kg, 15 kg payload, hot-swap battery, enterprise pricing.
- **Framing preserved:** RoboZaps' verdict — *"By commercial evidence the S2 leads the industry; by independently verified capability, no humanoid yet leads anyone."* [DIRECTIONAL assessment.]

### AgiBot (Zhiyuan): the shipment leader with thin disclosure

- **H1 2026 shipments: ~8,400 (SAG) / ~9,700 (Counterpoint)** — #1 globally by both trackers [VERIFIED as market leader by trackers; company details secondary].
- **Backed by CATL** (battery giant); Shanghai-based. With Unitree holds **>70% of the global market** by tracker counts.
- **Disclosure gap:** deeper company-level disclosures (audited unit/revenue splits) were not found in this pass — unlike Unitree and UBTech, AgiBot's figures reach us only through market trackers. Flag for follow-up. [UNVERIFIED at filing level.]

### EngineAI: combat as marketing, teleoperation as fine print

- **T800 specs [VENDOR]:** 173 cm, 75–85 kg, 29 joints, peak joint torque up to 450 N·m.
- **2026-07-16, Shenzhen — URKL (Ultimate Robot Knock-out League)** [VERIFIED as staged event]: billed as the world's first humanoid robot fighting tournament; 32 teams (China, US, Poland, Singapore); T800 as standard platform; **10 kg pure-gold belt (~¥10M / ~$1.4M)** for the champions; finals Dec 2026; opener saw a T800 ("Matador") decapitated yet keep fighting on its torso system. Elon Musk shared the clip: "Robot fights are fun." Donnie Yen attended. [DEMO — entertainment spectacle.]
- **2026-09-09, Riyadh — Hero Esports' CyberHero** kickboxing event: 10 T800s, Middle East's first humanoid kickboxing show [DEMO].
- **2026-09-18, San Francisco — REK (Robot Entertainment Kombat)** [VERIFIED]: first billed human-vs-full-size-humanoid fight — creator **Frankie LaPenna vs EngineAI T800**; the robot won; LaPenna left with a **reported hand/wrist injury**. **Critical: the T800 was remote-piloted** (TMZ; REK's own site sells "real pilots and real punches"). A person beat a person using a machine body. **This is the one verified 2026 human injury from robot combat.** [DEMO — teleoperated entertainment; do not file under autonomy.]
- **2026 World Robot Conference (Beijing):** two T800s boxed ("Doujiang" vs "Luzhu"); one recovered from a knockdown (Channel 4 News) [DEMO].
- **IFA 2026 (Berlin):** untethered T800 doing punch feints and heel kicks near attendees; PrimeBOT Q1 marching into a kneecap; webpronews alarm piece on unrestrained fighting robots on a trade-show floor [DEMO — safety-relevant].
- **Pattern context (2025, pre-period, included as context):** Russia Volga Electronics store — a customer shoved a humanoid, it "adopted a kung-fu stance" and kicked/punched until restrained (NY Post, Sept 2025); Unitree handler groin-kick during testing.
- **Framing rule:** combat is a durability/marketing spectacle; treat all "robot beats human" headlines as **teleoperated entertainment** unless autonomy is evidenced.

### Galbot, Leju, and the invisible middle

- **The leader-board fact:** both H1 2026 trackers rank Galbot third and Leju fifth (behind AgiBot, Unitree; ahead of/behind UBTech), in the same order. This is the only verifiable statement about either company in this pass.
- **What is not known:** no 2026 company-level disclosures, no filing-level unit/revenue splits, no named customer deployments, and no pricing were found for Galbot or Leju in this pass. [UNVERIFIED at company level; tracker-level rank only.]
- **Why it matters [DIRECTIONAL]:** two companies in the global top five are nearly invisible in Western press — the coverage gap is itself a finding about where the industry's demand base sits. The China-domestic market (hospitals, malls, reception, logistics) absorbs volumes that never appear in English-language reporting. Any "global humanoid market" narrative built from Western sources alone is structurally incomplete.
- **Follow-up:** dedicated passes on Galbot and Leju company disclosures are needed before the next consolidation; their tracker rank makes them too large to leave as footnotes.

### Manufacturing and supply-chain ramp (2026)

- **Figure BotQ:** unveiled March 2025 at 12K units/yr line capacity [VENDOR]; ~1 robot/90 min in April 2026, reported as a **24× ramp in 120 days to ~1 robot/hour** [VENDOR]; **1,000th Figure 03 on 2026-07-26** [VENDOR]; ~740 robots deployed vs. 660 employees (June 2026) [VENDOR]. Throughput claims are CEO-announced snapshots, not audited throughput.
- **UBTech–Dongfeng Liuzhou:** 1,000th Walker S2 targeted late Dec 2025; capacity guidance **5,000 (2026) → 10,000 (2027)**; **10-minute production pace** reported September 2026 [SECONDARY]. The China–Vietnam border contract (¥264M, ~$37M [SECONDARY], deliveries from Dec 2025, work underway mid-2026) is UBTech's non-factory revenue proof point.
- **Agility RoboFab (Salem, Oregon):** 70,000 sq ft, **10,000 units/yr peak capacity** [VENDOR]; **Fremont, CA 60,000 sq ft AI training hub opened July 2026** — deliberately sited near Tesla's Optimus factory (CEO Peggy Johnson: "having Tesla in the same area is a good thing").
- **Tesla–Ningbo chain (Sept 16–17, 2026) [SECONDARY]:** certified Tuopu Group (actuators/chassis), Ningbo Joyson Electronic (sensors), Zhejiang Sanhua Intelligent Controls (thermal); initial batch of ~5,000 units reported. Sanhua's ~$685M linear-actuator order (Oct 2025) is the upstream bellwether — estimated at ~180,000 robots' worth by watchers.
- **Scale asymmetry [DIRECTIONAL]:** the combined named 2026 line capacities (BotQ 12K, RoboFab 10K peak, UBTech guidance 5K, Optimus 1M design) describe installed tooling, not output. Actual 2026 humanoid output evidence tops out at the low tens of thousands globally (trackers), and productive deployments in the low hundreds. Capacity headlines are the single most misread number in 2026 robotics reporting.

### Helix 02 and the autonomy-architecture debate

- **Helix 02 (Jan 2026) [VENDOR]:** Figure's System 2 / System 1 architecture — a 7B VLM running at 7–9 Hz for reasoning and a 80M-parameter policy at 200 Hz for control, fully onboard on dual GPUs. The claim of "full-body autonomy" is vendor-announced; no independent benchmark of Helix 02's unsupervised task success was found in this pass.
- **The VLA split:** Figure (Helix, in-house after ending the OpenAI partnership Feb 2025), 1X (Redwood AI), Physical Intelligence (pi0.5/pi0.7, cross-embodiment claims), Skild (Skild Brain, "omni-bodied," trained on human video + Isaac Lab/Cosmos simulation), Google DeepMind (Gemini Robotics 1.5 + ER 1.5, 2025 baseline). 2026 did not settle which approach generalizes; it only priced them — Skild $14B+, PI >$11B — on vendor-reported claims.
- **Adcock's "utter dead end" (Aug 2026) [VENDOR]:** Figure's CEO declared wheeled robots an "utter dead end" at the ladder-climbing demo. The industry note: wheels remain dominant in warehouse automation (AMRs), and the claim is a positioning statement, not an empirical result. Read as competitive rhetoric, not taxonomy.
- **Training-data economics [DIRECTIONAL]:** PI cites 10K+ robot-hours across 7 platforms; Skild cites human video + physics simulation (Isaac Lab, Cosmos WFM). Teleoperation (1X Expert Mode) doubles as a data-collection pipeline. Whoever accumulates the largest labeled manipulation dataset — not who ships the most bodies — may own the margin layer. This is the strongest form of the "brain" thesis.
- **Open models as counterweight:** PI's openpi is open-source; NVIDIA's GR00T N1 is an open humanoid foundation model (2025); Hugging Face's LeRobot stack (see §16) underpins the community tier. The open tier sets the floor; the closed tier prices the ceiling. [COMMUNITY]

### Safety, teleoperation, and liability — the 2026 record

- **2026-09-18, San Francisco — REK fight [VERIFIED, DEMO]:** the first billed human-vs-full-size-humanoid fight; Frankie LaPenna vs. EngineAI T800; robot won; **reported hand/wrist injury to LaPenna**. The T800 was remote-piloted (TMZ; REK's own site sells "real pilots and real punches"). Framing rule: a person beat a person using a machine body. This is the **only verified 2026 human injury** from robot combat.
- **IFA 2026 (Berlin) [VERIFIED as event, DEMO]:** untethered T800 doing punch feints and heel kicks near attendees; PrimeBOT Q1 marching into a kneecap; webpronews alarm piece on unrestrained fighting robots on a trade-show floor. Safety-relevant: no barriers between high-torque humanoids and the public.
- **2026 World Robot Conference (Beijing) [DEMO]:** two T800s boxed ("Doujiang" vs "Luzhu"); one recovered from a knockdown (Channel 4 News).
- **URKL Shenzhen opener (2026-07-16) [DEMO]:** a T800 ("Matador") decapitated yet continuing to fight on its torso system — marketed as durability proof; no safety information attached.
- **The teleoperation pattern [DIRECTIONAL, disclosed cases]:** 1X's Expert Mode (disclosed); REK's remote pilots (disclosed); SemiAnalysis's "heavily teleoperated" industrial units (June 2026); Figure's logistics deployments described as structured tasks with human oversight (carscoops). The industry's "autonomy" in 2026 is substantially supervised autonomy — investors and reporters who do not ask "who is holding the joystick" will misread the evidence.
- **2025-11 — BMW whistleblower suit [VERIFIED filing, outcome unknown]:** federal safety whistleblower lawsuit filed at the close of the Figure 02 pilot. Live watch item for 2027; do not assert its outcome. Its significance: the first federal safety-litigation instrument attached to a humanoid factory pilot.
- **CVE-2026-25874 (LeRobot PolicyServer pickle RCE, CVSS 9.8, unpatched as of Sept 2026) — cross-ref §16 only:** the software-supply-chain layer of robot safety. Not duplicated here per the §16 dedup rule.
- **2025 pattern incidents (pre-period context):** Unitree handler groin-kick during testing; a robot slapping a child mid-dance-demo (viral clips, 2025); Russia Volga Electronics store (Sept 2025) — a shoved humanoid "adopting a kung-fu stance" and striking until restrained (NY Post). Included as pattern context for why trade-press alarm pieces (IFA 2026) gained traction.
- **NVIDIA Halos (June 2026, Digit v5) [VENDOR]:** the year's substantive safety-engineering answer — cooperative-safety architecture removing the safety cage for high-traffic warehouses; Amazon explicitly holds its Digit deployment for v5. Safety is thus both a liability story (fights, suits) and a commercial unlock (Halos).

### Robot foundation models: the $25B+ "brain" bet

- **2026-01-14 — Skild AI Series C [VERIFIED]:** **~$1.4B at $14B+** (Business Wire): led by **SoftBank**; NVentures (NVIDIA), Macquarie Capital, Bezos Expeditions, Disruptive, 1789 Capital; returning Lightspeed, Felicis, Coatue, Sequoia; strategic Samsung, LG, Schneider Electric, CommonSpirit, Salesforce Ventures. Total raised $2.2B+. **Skild Brain**: "omni-bodied" foundation model (any robot, any task) trained on human video + physics simulation (Isaac Lab, Cosmos WFM); **~$30M revenue within months in 2025** [VENDOR]. Founders: Deepak Pathak, Abhinav Gupta (CMU Robotics Institute). [DIRECTIONAL: venture is pricing foundation models as the value layer.]
- **2026-03-27 — Physical Intelligence [VERIFIED reports]:** in talks to raise ~**$1B at >$11B valuation** (Bloomberg/TechCrunch) — nearly doubling the $5.6B Series B (Nov 2025) in four months. Participants per Bloomberg: Founders Fund, Lightspeed, returning Thrive/Lux. Early-stage, details changeable. Founders: Karol Hausman (ex-DeepMind), Sergey Levine (UC Berkeley), Chelsea Finn, Lachy Groom; ~80 employees; "ChatGPT for robots" (Levine). **pi0.5 (Apr 2026)**: claimed first cross-embodiment generalization without per-robot fine-tuning; pi0.7: zero-shot task generalization [VENDOR]. 10K+ hours robot data across 7 platforms; openpi open-source. **Anthropic–PI acquisition talks** (spring 2026, The Information, July 2026) — **denied by CEO Hausman**. OpenAI is a PI investor.
- **Gemini Robotics 1.5 — BASELINE CORRECTION [VERIFIED date]:** announced **2025-09-25 — NOT a 2026 release.** Multiple outlets date it September 2025. Included only as baseline: VLA (Gemini Robotics 1.5) + embodied-reasoning planner (Gemini Robotics-ER 1.5, tool-calling incl. web search); cross-embodiment transfer (Aloha 2 → Franka → **Apptronik Apollo**); trusted testers + AI Studio rollout. **No 2026 DeepMind robotics model release found in this pass** — flag as a coverage gap, not as absence.
- **NVIDIA 2026 robotics [VERIFIED]:** Isaac Root research platform (Computex 2026-06-01, Unitree H2 flagship hardware, Jetson Thor Blackwell + Isaac GR00T models); **Halos safety architecture** (first humanoid integration: Agility Digit v5, announced June 2026 at Automate); GR00T N1 (open humanoid foundation model) and Jetson Thor general availability date to **2025** (GTC/CES) — 2026 is the deployment/integration year, not the announcement year.
- **Cross-ref (not duplicated):** Hugging Face robotics — Pollen Robotics acquisition (Apr 2025) → Reachy Mini launch (Jul 2025; $299 Lite / $449 Full) → 3,000 units + Huang CES keynote (Jan 2026) → **HopeJR (66 actuated DOF, ~$3,000) + Reachy Mini refresh ($250–$300) unveiled 2026-09-01** → CVE-2026-25874 (LeRobot async PolicyServer pickle RCE, CVSS 9.8, unpatched as of Sept 2026) — covered in **§16**; no new HF-robotics facts in this pass.
- **2026 capital events in robotics (one-line cross-refs to §20):** Skild AI Series C $1.4B at $14B+ (2026-01-14); Physical Intelligence ~$1B at >$11B (2026-03-27, in talks); Agility SPAC $2.5B (2026-06-24); Unitree IPO ¥6.1B raised / ~$51B debut cap (2026-08-19); SoftBank→ABB Robotics $5.4B (announced **Oct 2025**, close expected mid-to-late 2026); Figure Series C >$1B at $39B (2025-09; flat $39B Setter secondary Q2 2026).

