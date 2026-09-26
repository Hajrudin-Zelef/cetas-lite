---
id: ai-industry-kb-2026/22-robotics/galbot-leju-and-the-invisible-middle
title: "Galbot, Leju, and the invisible middle"
domain: robotics
role: deep-dive
task: robotics
actors: ["China", "Google", "Hugging Face", "Nvidia", "OpenAI", "United States", "xAI"]
dates: ["2025-03", "2026-04", "2026-06", "2026-07", "2026-07-16", "2026-07-26", "2026-09", "2026-09-09", "2026-09-18"]
keywords: ["benchmark", "gemini", "gpus", "humanoid", "liability", "nvidia", "omni", "pricing", "reasoning", "revenue", "robotics", "throughput"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [10693, 10726]
section: "22. Robotics"
sha256: 7fcd35b6fa4ef23affe56d0fd159d50bec527f6c17e47cb2b9c34ecd05d8d23e
---

# Galbot, Leju, and the invisible middle

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

