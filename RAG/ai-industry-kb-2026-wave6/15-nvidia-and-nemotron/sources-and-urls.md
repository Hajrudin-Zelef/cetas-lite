---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/sources-and-urls
title: "Sources and URLs"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Groq", "Nvidia"]
dates: ["2026-07-27", "2026-08-11"]
keywords: ["accelerator", "agent", "agentic", "agents", "attribution", "benchmark", "benchmarks", "cost", "datacenter", "decode", "diffusion", "distillation"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7703, 7739]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 9824b0db4ea6ff32a30a1f2cf3c68b67889e4a7fcaab8bfc4c2c34f53a7972d4
---

# Sources and URLs

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

## Sources and URLs
- https://www.globenewswire.com/news-release/2026/03/16/3256731/0/en/NVIDIA-Expands-Open-Model-Families-to-Power-the-Next-Wave-of-Agentic-Physical-and-Healthcare-AI.html
- https://www.reuters.com/business/nvidia-is-developing-nemotron-4-open-source-models-information-reports-2026-08-11/
- https://www.reuters.com/business/nvidia-forms-industry-alliance-open-ai-security-after-hugging-face-hack-2026-07-27/
- https://www.globenewswire.com/news-release/2026/08/24/3349940/0/en/nvidia-groq-3-lpx-now-in-full-production-with-world-class-speed-for-agentic-ai.html
- https://fourweekmba.com/ai-nvidia-meta-open-weights-coalition-distillation-policy/
- https://www.unite.ai/nvidia-and-microsoft-back-open-weight-ai-in-joint-letter/
- https://blog.corenexis.com/open-weights-american-ai-leadership
- https://pjfp.com/jensen-huang-x-open-weights-letter/
- https://techxplore.com/news/2026-07-tech-giants-source-ai-alliance.pdf
- https://insideai.news/news/ai-safety/nvidia-launches-open-secure-ai-alliance-after-openai-agent-hack-test/6509/
- https://github.com/vincentzli/clawnews/blob/HEAD/src/content/news/open-weights-fight-ai-cost-floor.mdx


### New sources — expansion

