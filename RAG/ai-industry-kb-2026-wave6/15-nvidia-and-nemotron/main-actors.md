---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/main-actors
title: "Main actors"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "CISA", "China", "Groq", "Hugging Face", "Nvidia", "OpenAI", "Poolside", "United States", "Z.ai"]
dates: ["2025-08-25", "2025-12-15", "2025-12-19", "2026-01-05", "2026-03-11", "2026-03-16", "2026-06-01", "2026-06-04", "2026-07-15", "2026-07-21", "2026-07-24", "2026-07-27", "2026-08-03", "2026-08-11", "2026-08-13", "2026-08-24", "2026-09-11"]
keywords: ["accelerator", "acquisition", "agent", "agentic", "agents", "attribution", "benchmark", "benchmarks", "cost", "datacenter", "decode", "diffusion"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7651, 7723]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 0b9b64e8b3e8c771c6d9c1d2bf5d080575a8ee1176b4ce19c357cc87a589d365
---

# Main actors

## Main actors
- **NVIDIA** — author of both July coalition moves (letter co-host, alliance founder); the Nemotron open-model line; Nemotron Coalition; Object-Oriented Agent; the $6B Poolside license [SECONDARY/VENDOR].
- **Jensen Huang** — first-ever X post amplified the July 24 letter [SECONDARY].
- **Nemotron 3 Ultra** — the 550B flagship checkpoint: 90% sparsity, Mamba-2 + LatentMoE + NVFP4 [SECONDARY].
- **Groq** — Groq 3 LPX full production (Aug 24); DOJ "shadow acquisition" probe (Sept 11) [VENDOR/SECONDARY].
- **The Open Secure AI Alliance** — 30+ companies; open-source security tooling for autonomous AI [SECONDARY].
- **GLM 5.2 (Z.ai)** — the Chinese open-weight model that did HF's forensic work in the breach [SECONDARY].

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

- The Lightning sizing dispute (30B/3B vs 31.6B/~3.6B) is a naming-convention artifact worth standardizing: secondary sources round to the "-30B-A3B" checkpoint name while vendor material cites exact counts. Analysts should cite the checkpoint name and note both figures rather than picking one silently. [DIRECTIONAL] [S1][S6]
- Switchyard's router taxonomy (classifier with session affinity, stage router reading tool activity, escalation, residual-stream prefill router) shows NVIDIA productizing the "cheap-first, escalate on difficulty" pattern that 2026 agent stacks converged on — routing is becoming a benchmarked discipline with published cost/accuracy Pareto curves, not folklore. [DIRECTIONAL] [S1][S3]
- The LangChain (−74% cost, 7% frontier calls, ~6 pt tradeoff) and Cognition (within 2.8 pts of Opus 5 at ~28% lower cost) results give the first public, third-party-measured calibration of what intelligent routing is worth on agentic coding benchmarks — useful as a baseline for any 2026 router evaluation. [SECONDARY] [S1]
- NVFP4 + QAD (not PTQ) as the reference 4-bit flow raises the bar for community quantization: matching vendor NVFP4 quality now requires distillation-aware pipelines, and E2M1's precision-exception path is a format-level detail quants must handle. [DIRECTIONAL] [S9][S1]
- Cosmos 3's reasoner/generator split with independent fine-tuning is a deliberate play for robotics researchers: adapt the diffusion generator to your hardware without retraining the VLM brain — a modularity story no single-tower world model offers. [DIRECTIONAL] [S11][S13]
- The Cosmos 3 three-tier ladder (Edge 4B on-device → Nano 16B workstation → Super 64B datacenter) mirrors the LLM industry's size ladder, but for world models — expect 2026 robotics teams to standardize on Nano for development and distill/edge-ify for deployment. [DIRECTIONAL] [S11][S14][S15]
- GR00T N2's extended preview (announced at GTC, availability "by end of year") plus N1.7's early-access-with-commercial-license pattern shows NVIDIA gating humanoid foundation models more tightly than its LLM line — robotics releases move slower and with less public technical detail. [DIRECTIONAL] [S17]
- Alpamayo's OpenMDW-1.1 licensing across the family (including retroactive commercial coverage for earlier R&D releases) is NVIDIA's answer to the "open weights but not really deployable" criticism — and a template other vendors may copy for physical-AI models. [DIRECTIONAL] [S21][S24]
- VoiceChat 11B's frame-locked full-duplex design (12.5 Hz, no external VAD, tool calling mid-conversation) is the first open model to unify the voice-agent stack in one forward pass; the 129 GB H200 footprint shows it is still datacenter-class, so edge voice agents remain cascade-based for now. [DIRECTIONAL] [S25][S27]
- The Nemotron catalog's modality separation (ASR, voicechat, OCR, parse, embed/rerank-VL, safety as distinct models) confirms the "pipeline of models" reality: quoting a single Nemotron number for a multimodal system understates the true latency/cost/failure surface. [SECONDARY] [S8]
- The 5M-vs-50M correction stands: Nemotron's verified HF traction is strong but an order of magnitude below the inflated figure — training samples and downloads must never be merged in adoption narratives. [SECONDARY] [S30]

- Nemotron 3 Super's 449 tok/s at 91.75% RULER@1M makes it the open-model long-context workhorse of 2026: teams that need million-token agentic runs now price against $0.10/$0.50 DeepInfra rates, not frontier-model pricing. [SECONDARY] [S31][S32]
- The 256K-default-config nuance on Super's 1M context is a practical trap for evaluators: published RULER@1M numbers require explicit `VLLM_ALLOW_LONG_MAX_MODEL_LEN=1` and extra VRAM beyond the 8x H100 minimum — benchmark claims should state the config. [SECONDARY] [S32][S36]
- Super's mixed competitive picture (beats GPT-OSS-120B, trades wins with Qwen3.5-122B) confirms the 2026 open-model market is tiered by workload, not by a single leaderboard — long-context retrieval and math favor Super; general knowledge favors Qwen. [SECONDARY] [S32]
- The license-attribution split (Super under "Nemotron Open Model License", Lightning/Cosmos/Alpamayo under OpenMDW-1.1) needs watching: if NVIDIA is mid-rename, downstream compliance docs should cite the checkpoint's own card, not the family name. [DIRECTIONAL] [S32][S1]
- Dynamo's prefill-on-Rubin / decode-on-Groq-LPU split is the first concrete productization of heterogeneous inference fabrics — the "one accelerator for everything" era is ending, and serving stacks must now schedule across ISAs. [DIRECTIONAL] [S40][S41]
- Jetson Thor's four-tier ladder (400–2,070 FP4 TFLOPS, unified JetPack) does for physical AI what the Cosmos size ladder does for world models: a portable scale-down path from datacenter to robot. [DIRECTIONAL] [S39]
- The LG deal (MoU → Q1 2027 unveil → 80MW AI factory 2028) shows physical-AI partnerships are now full-stack: robot, data platform, AI factory and vehicle platform in one agreement. [SECONDARY] [S19]
- Cosmos-Predict2.5/Transfer2.5's open licensing plus RL-refined synthetic data loops close the sim-to-real gap as a software problem — Expect 2026 robotics teams to budget synthetic-data generation alongside real-world collection. [COMMUNITY] [S45]

