---
id: ai-industry-kb-2026/22-robotics/robot-foundation-models-the-25b-brain-bet
title: "Robot foundation models: the $25B+ \"brain\" bet"
domain: robotics
role: deep-dive
task: robotics
actors: ["AWS", "Anthropic", "China", "Google", "Hugging Face", "Nvidia", "OpenAI", "Samsung"]
dates: ["2025-09", "2025-09-25", "2026-01-14", "2026-03-27", "2026-06", "2026-06-01", "2026-06-24", "2026-07", "2026-07-16", "2026-08-19", "2026-09-01", "2026-09-18"]
keywords: ["acquisition", "blackwell", "chatgpt", "fine-tuning", "gemini", "humanoid", "ipo", "lawsuit", "liability", "nvidia", "omni", "pricing"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [10727, 10745]
section: "22. Robotics"
sha256: 868a2f22a3e1cff2fe9f2d011f912ea24bc3234af4d0b2147d175bcd4b8f8bf8
---

# Robot foundation models: the $25B+ "brain" bet

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

