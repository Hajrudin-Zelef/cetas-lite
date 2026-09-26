---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/qwen-image-3-0-3-0-pro-2026-07-21-no-benchmarks-no-weights-s
title: "Qwen Image 3.0 / 3.0 Pro — 2026-07-21, no benchmarks, no weights [SECONDARY]"
domain: qwen-and-alibaba
role: deep-dive
task: benchmark
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "Google", "Hugging Face", "Mistral", "Nvidia", "OpenAI", "Together AI", "United States"]
dates: ["2026-03-30", "2026-05", "2026-07", "2026-07-15", "2026-07-21", "2026-08-05", "2026-09-22"]
keywords: ["benchmark", "benchmarks", "qwen", "agent", "agentic", "alignment", "apache", "attention", "claude", "consumer", "cost", "gemini"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1202, 1236]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: f64ae99f2ed6a957822dbba7db7beccbc21ebef00ba88e8acce2997ad02acfe5
---

# Qwen Image 3.0 / 3.0 Pro — 2026-07-21, no benchmarks, no weights [SECONDARY]

### Qwen Image 3.0 / 3.0 Pro — 2026-07-21, no benchmarks, no weights [SECONDARY]
- **CONTRADICTION**: release date reported as **2026-07-21** (majority of launch-week coverage) vs **2026-08-05** (siray.ai). The July 21 date is better attested. [SECONDARY]
- Headline spec: **4,500-token prompt input** (up from ~1,000 on Qwen-Image-2.0), vendor-demonstrated **text legibility down to 10 pixels**, **native support for 12 languages**, LaTeX math notation, single-pass dense layouts (nine-panel infographic grids, nested UI mockups, full academic-paper pages, simulated newspaper front pages), and live-data capability (e.g., a weather graphic for a specific city and date). [SECONDARY]
- **Transparency reversal**: no benchmark scores, no model card, no technical report, no parameter count, **no downloadable weights** — a sharp break from Qwen-Image 1.0 (Aug 2025, Apache-2.0 open weights + same-day technical report) and Qwen-Image 2.0 (May 2026, report + Alibaba's own eval leaderboard). Access via Qwen Chat and Alibaba Cloud APIs; reported invite-only at launch with no public API initially. [SECONDARY]
- Lineage context (predates 3.0, do not attribute to it): third-party trackers reported the Qwen-Image line at **88.32 GenEval, 83.84 GSO image-editing, 58.30 LongText-Bench English** (Jan 2026); Alibaba's own **Qwen-Image-Bench placed Qwen Image 2.0 Pro 5th overall**, behind OpenAI and Google models. [SECONDARY]
- Independent hands-on tests (secondary): strengths in dense text rendering and reference-based editing; gaps in landmark and historical-figure accuracy. [SECONDARY]
- **Positioning keyword**: 1.0 = "Precision"; 2.0 = "Precision, Variety, Completeness, Beauty, Authenticity"; 3.0 = **"Real"** (实) — organized into three pillars: **Rich Content, Authentic Details, Deep Knowledge**. Goal: "useful" rather than "good-looking" — deployable productivity tool for documents, interfaces, knowledge-dense visuals. [SECONDARY]
- **Tiers**: `qwen-image-3.0` (standard), `qwen-image-3.0-pro` (flagship). Resolution **512×512 to 2048×2048** (2K PNG). Reference images (editing): **1–3 per request**. Pricing: **from $0.03/image** (high-res); input $0.003/image; Pro i2i **$0.075/generation** (ModelsLab). Availability: Model Studio, Qwen Cloud, kie.ai, OpenArt, ModelsLab. [SECONDARY]
- **Launch sequence**: shipped **2026-07-21** behind invite-only wall; **opened to all Qwen AI platform users 2026-08-05** (OrcaRouter GA tracking). The 07-21/08-05 "contradiction" is **invite-only vs GA**, not a date conflict. [SECONDARY]
- **2.0 efficiency context**: Qwen-Image-2.0 (May 2026) shipped a 4-step variant (down from 40 steps); on Alibaba's arena, 2.0 landed just behind **GPT-Image-2** and **Nano Banana Pro**. [SECONDARY]

### Qwen3.5-Omni — 2026-03-30 [SECONDARY]
- Released **2026-03-30**: multimodal model accepting **text + image + audio + video**, **113 languages**, with "Vibe Coding" capability. [SECONDARY]
- **Architecture**: **Thinker-Talker + Hybrid-Attention MoE** (both Thinker and Talker MoE); **256K context**; **>10 hours continuous audio**; **400+ seconds 720p video** (1 FPS); trained on **>100M hours audio/video data**. [SECONDARY]
- **Variants**: **Plus** (30B-A3B MoE), **Flash** (lightweight MoE), **Light** (dense, open weights). **Qwen3.5-Omni-Plus-Realtime** for streaming. [SECONDARY]
- **Speech**: **113 languages/dialects recognition** (up from 19), **36 languages generation** (up from 10); multi-codebook codec for single-frame synthesis; **ARIA** (dynamic text-speech alignment in streaming decoding). [SECONDARY]
- **Benchmarks**: **215 SOTA** across audio/video understanding (vendor claim — treat with skepticism; self-selected benchmarks favor the releasing lab); Plus **surpasses Gemini 3.1 Pro on audio understanding**, matches on audio-visual. [SECONDARY]
- **Audio-Visual Vibe Coding**: camera + speech/gesture → functional website/game — "building through natural language and visual descriptions." [SECONDARY]
- **Availability**: Qwen Chat, **open weights on Hugging Face/ModelScope**, Alibaba Cloud API (Offline + Realtime). Technical report: arXiv:2604.15804. [SECONDARY]
- **Release cadence note**: Feb (Qwen3.5 397B, Qwen3-Coder-480B, Qwen3-Coder-Next) → Mar 30 (Qwen3.5-Omni) → Mar 31 (Qwen3.6-Plus Preview) — **five major releases in two months**; the Mar 30–31 back-to-back is "wartime shipping." [SECONDARY]

### Regulatory + pricing events — July 2026 [SECONDARY]
- **2026-07-15**: Alibaba's **Qwen app and ByteDance's Doubao simultaneously shut down user-created/custom AI agent features** as China's **'humanlike AI interaction services'** (anthropomorphic-AI/companion) rules took effect. **Qwen offered no migration path** — agent configurations and conversation histories **permanently deleted** (Doubao offered read-only access until Oct 15). Consumer-app change; models/API unaffected. [SECONDARY]
- **Early July 2026** (SCMP): Alibaba cut **Qwen3.7-Max by ~80%** and **Qwen3.7-Plus by ~60%** for international users on **Qoder** during off-peak hours (10pm–8am Beijing = US working hours) — an explicit play for US developer demand amid the Chinese token price war. [SECONDARY]

### Qwen3-Coder line — 480B/30B, SWE-bench Verified 69.6% [SECONDARY]
- **Qwen3-Coder-480B-A35B-Instruct**: 480B MoE / **35B active**, **256K context natively (1M with extrapolation)**, **Apache 2.0**, vendor claim of open-model SOTA on Agentic Coding, Agentic Browser-Use, Agentic Tool-Use, "comparable to Claude Sonnet 4." **SWE-bench Verified: 69.6% (348/500)** via OpenHands (Aug 2025 submission) — per-repo breakdown: django 165/231 (71.4%), sympy 53/75 (70.7%), sphinx 30/44 (68.2%), scikit-learn 27/32 (84.4%), pytest 15/19 (79.0%), xarray 17/22 (77.3%), requests 6/8 (75.0%), matplotlib 20/34 (58.8%), astropy 10/22 (45.5%), pylint 3/10 (30.0%). [SECONDARY]
- Lowest provider pricing (llm-stats.com, 2026-09-22): **$0.30 input / $1.00 output** (Deepinfra). [SECONDARY]
- **Qwen3-Coder-30B-A3B-Instruct**: SWE-bench Multilingual **~33.3–33.8%** (OpenHands/SWE-agent, NVIDIA NeMo skills eval). [SECONDARY]
- **Qwen3-Coder-Next**: smaller/faster hybrid architecture, reported **SWE-bench >70%** at much lower inference cost. [SECONDARY]
- **Training**: **7.5T tokens, 70% code, 358 programming languages**; post-training: SFT + RL (Code RL: execution-driven; Long-Horizon Agent RL: multi-turn planning/tool use); **Constitutional AI** for safety. Qwen built a **20,000-environment system on Alibaba Cloud** for agentic evaluation/training. [SECONDARY]
- **Qwen Code**: open-sourced CLI (forked from Gemini Code), Node.js/npm; integrates with **Claude Code** (DashScope proxy), **Cline**, Ollama, LMStudio, MLX-LM, llama.cpp, KTransformers. [SECONDARY]
- **Together AI pricing**: **$2.00/$2.00** per 1M input/output (480B). [SECONDARY]
- **SWE-bench harness note**: the 69.6% (348/500) is **OpenHands, 500-turn**; standard (non-500-turn) reported as **67.0%** — the harness/turn budget matters, never mix. [SECONDARY]

