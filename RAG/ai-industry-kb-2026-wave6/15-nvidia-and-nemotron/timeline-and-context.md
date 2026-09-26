---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/timeline-and-context
title: "Timeline and context"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Anthropic", "CISA", "China", "Groq", "Hugging Face", "Nvidia", "OpenAI", "Poolside", "United States"]
dates: ["2025-08-25", "2025-12-15", "2025-12-19", "2026-01-05", "2026-03-11", "2026-03-16", "2026-06-01", "2026-06-04", "2026-07-15", "2026-07-21", "2026-07-24", "2026-07-27", "2026-08-03", "2026-08-11", "2026-08-13", "2026-08-24", "2026-09-11"]
keywords: ["acquisition", "agent", "agentic", "benchmark", "cost", "decode", "diffusion", "disaggregated", "distillation", "fp4", "full-duplex", "humanoid"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7659, 7702]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: de4fa9307f063e108b95d89937ffe8f23b7bcbb9230561fc6d18f0d804beae70
---

# Timeline and context

## Timeline and context
- **2025-12-15** — Nemotron 3 Nano [VENDOR].
- **2026-03** — Nemotron-Cascade 2 (IMO 2025 gold) [SECONDARY].
- **2026-03-16** — GTC 2026 open-model-family announcements [VENDOR].
- **2026-06-01** — Nemotron 3 Ultra launch (Computex) [SECONDARY].
- **2026-06-04** — Nemotron 3 Ultra weights [SECONDARY].
- **2026-07-21** — OpenAI discloses the autonomous-agent HF breach; FBI alerted [SECONDARY].
- **2026-07-24** — "Open Weights and American AI Leadership" letter (25 orgs; Huang's first X post) [SECONDARY].
- **2026-07-27** — Open Secure AI Alliance launches (NVIDIA + 30+ companies) [SECONDARY].
- **2026-08-11** — Nemotron 3.5 Lightning + NeMo Switchyard; Reuters covers Nemotron 4-in-development report [SECONDARY].
- **2026-08-24** — Groq 3 LPX full production [VENDOR].
- **2026-09-11** — DOJ "shadow acquisition" investigation reported [SECONDARY].


### New verified timeline entries — expansion

- 2025-12-19 — NVIDIA developer forum records 5M Hugging Face downloads for the Nemotron family; the verified figure, not 50M+. [SECONDARY] [S30]
- 2026-01-05/06 — NVIDIA launches the Alpamayo family at CES 2026: Alpamayo 1 (10B reasoning VLA), AlpaSim simulator, 1,700+ hours of driving datasets; Lucid, JLR and Uber named as adopters. [SECONDARY] [S20][S22]
- 2026-06-01 — NVIDIA announces Cosmos 3 at GTC Taipei (one source says ~May 31): two-tower MoT omnimodal world models; Nano 16B and Super 64B open weights under OpenMDW-1.1. [SECONDARY] [S11][S12][S13]
- 2026-06-01 — NVIDIA announces Alpamayo 2 Super (34B VLA: 32B Cosmos-3-Super-Reasoner-based VLM + 2.3B diffusion action decoder), OpenMDW-1.1, availability expected summer 2026; also the Isaac GR00T Reference Humanoid Robot. [SECONDARY] [S21][S23][S18]
- 2026-07-15/20 — Cosmos 3 Edge (4B on-device world model) launches; weights land on Hugging Face July 20 (`nvidia/Cosmos3-Edge`), targeting Jetson Thor 15 Hz control. [SECONDARY] [S14][S15]
- 2026-08-03 — NVIDIA publishes NemotronLabs VoiceChat 11B on Hugging Face under OpenMDW 1.1: end-to-end full-duplex speech-to-speech at 12.5 Hz frame lock. [SECONDARY] [S25]
- 2026-08-11/12 — NVIDIA releases Nemotron 3.5 Lightning (30B/3B per secondary; 31.6B/~3.6B per vendor-adjacent) and the open-source NeMo Switchyard router; LangChain's 145-task benchmark shows −74% cost vs Opus 4.8 at ~6 pts accuracy tradeoff. [SECONDARY] [S1][S2]
- 2026 (GTC keynote window) — GR00T N1.7 enters early access with commercial licensing; GR00T N2 previewed (DreamZero world-action model, >2x success claim, No.1 MolmoSpaces/RoboArena), slated for availability by end of 2026. [SECONDARY] [S17]
- NVFP4 (E2M1, block-16, dual-scale) and the Nemotron-3-Ultra-NVFP4 QAD checkpoint workflow are documented in NVIDIA's Model Optimizer blog — the production 4-bit path for Nemotron 3-class models. [VENDOR] [S9]

- 2025-08-25 — Jetson AGX Thor (T5000) reaches general availability: 2,070 FP4 TFLOPS, 128GB, 130W. [SECONDARY] [S37]
- 2026-01 — First Cosmos generation context: Huang's CES unveiling of world foundation models trained on 20M hours of video (Nano/Super/Ultra). [SECONDARY] [S46]
- 2026-03-11 — Nemotron 3 Super released at GTC (secondary catalogs; GTC announcements dated 2026-03-16 in base corpus): 120B/~12B, 1M context, 449 tok/s, SWE-Bench Verified 60.47%. [SECONDARY] [S31][S32]
- 2026-03 — GTC 2026: NVIDIA introduces disaggregated inference via Dynamo (prefill on Vera Rubin, decode on Groq LPU racks); TrendForce coverage dates the announcement to GTC week. [SECONDARY] [S40]
- 2026-07-15 — Jetson Thor family expands to four tiers (T3000/T2000 announced, shipping Q1 2027; T5000/T4000 already shipping). [COMMUNITY] [S39]
- 2026-08-13 — LG Group–NVIDIA MoU signed at Santa Clara: joint bipedal humanoid for Q1 2027 unveiling; Vera Rubin AI factory pilot H1 2027; 80MW Cheonan AI factory H1 2028. [SECONDARY] [S19]

## Implications
1. **The 90%-sparsity hybrid (Mamba-2 + LatentMoE + NVFP4) is NVIDIA's open-model technical bet** — efficiency as the moat for agentic inference, paired with Groq 3 LPX on the hardware side.
2. **Coalition-building is the product** — NVIDIA authored both July moves; the letter (distillation fight) and the alliance (security tooling) are policy infrastructure, not model releases.
3. **The breach inverted the safety argument** — the July 21 incident is now the canonical exhibit for "defenders need inspectable models," and a Chinese open model did the work US closed models refused.
4. **License, don't acquire, when the target is a model factory** — the $6B Poolside license (not acquisition) is the template for capability transfer without ownership.
5. **"In development" is not "announced"** — Nemotron 4's ≥1T report is a training-status leak, not a launch; treat all derived claims accordingly.
6. **Member-list variance is normal in fast coalitions** — three outlets, three partial lists for the alliance; cite "NVIDIA + 30+ companies," never a closed list.


### New verified implications — expansion

