---
id: ai-industry-kb-2026-wave6/04-glm-and-z-ai/post-ipo-valuation-arc
title: "Post-IPO valuation arc"
domain: glm-and-z-ai
role: deep-dive
task: finance
actors: ["Anthropic", "China", "Google", "Huawei", "Hugging Face", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "United States", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-01", "2026-04", "2026-06", "2026-06-13", "2026-06-16", "2026-07-13"]
keywords: ["ipo", "valuation", "agent", "agentic", "agents", "ascend", "attention", "benchmarks", "claude", "compute", "context window", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1685, 1754]
section: "§4. GLM and Z.ai"
sha256: 1af19247d665f41cf3cc1c0786ac458f89352f2b590df46051633eabbef5b2ca
---

# Post-IPO valuation arc

### Post-IPO valuation arc
- Stock surged **more than 2,000% since its January listing** (Seoul Economic Daily, ~June 2026); offering price HK$116.2 → stock as high as **HK$2,980**; market cap pushed **above HK$1 trillion (~US$128 billion / ~197 trillion won)** [SECONDARY] (en.sedaily.com).
- Bloomberg (via sedaily): Z.ai weighing a **multibillion-dollar share placement** ahead of the **six-month lock-up expiry (July 8)**; amount could "significantly exceed" the $558M IPO [SECONDARY] (en.sedaily.com).
- TechInAsia: Zhipu seeking **~US$4 billion via 19.8 million shares** at **HK$1,588–1,698** each (up to 13% discount to close); stock **+1,500% since listing**; sole coordinator **China International Capital Corp**; proceeds for **R&D, business expansion, investments, M&A, capital structure optimization, working capital** [SECONDARY] (techinasia.com).
- The GLM-5.2 rally: stock up **~50% in five days** amid praise for GLM-5.2 — AA Intelligence Index **51 points, 4th** after Fable 5, Opus 4.8, GPT-5.5; **first model from China — and among all open-source models — to surpass the 50-point mark** [SECONDARY] (en.sedaily.com).
- Cross-read: the HK$1T cap was reached **within a week of GLM-5.2's release** (per Tom's Hardware, cited by aqalion) — model performance and equity value are directly coupled in this market [SECONDARY] (aqalion.com).


### New verified facts — expansion (continued — GLM-5V-Turbo (native multimodal coding model))

### Release and positioning
- **Released April 1–2, 2026** (coverage dates; Z.ai announcement April 1 via marktechpost; launch write-ups April 2) [SECONDARY] (marktechpost.com; the-decoder.com; ubos.tech).
- **Z.AI's first multimodal coding foundation model**, purpose-built for vision-based coding: design-to-code, GUI agents, native vision for images/videos/designs [SECONDARY] (agentnativedev.medium.com; ubos.tech).
- **Closed-source, API-only** model (not part of the MIT open-weights line) [SECONDARY] (aqalion.com).
- Deep **OpenClaw and Claude Code integration** is a headline feature — the "perceive → plan → execute" loop for autonomous environment interaction [SECONDARY] (marktechpost.com; cometapi.com; ubos.tech).

### Architecture and spec
- **744B-A40B MoE backbone**: ~744B total parameters, ~40B active per token [SECONDARY] (aqalion.com; marktechpost.com).
- **Parameter-count note**: the aqalion review explicitly acknowledges the same 744B-vs-753B tension recorded in §4's contradictions — "Hugging Face model cards for the family list the confirmed total at 753B; treat 744B as Z.ai's official designation" [SECONDARY] (aqalion.com).
- **CogViT** vision encoder — a new encoder built by Z.ai for this model; images and text are processed together **from the start of training**, not as a bolted-on module [VENDOR via secondary] (the-decoder.com; marktechpost.com).
- **MTP**: marktechpost defines it as **Multi-Token Prediction** (multiple tokens predicted per inference step, speeding up output); ubos.tech defines it as **Multimodal-Task-Prompt** — record the expansion conflict [SECONDARY contradiction].
- **200K context window** (202,752 on OpenRouter); **128K max output** [SECONDARY] (cometapi.com; the-decoder.com; aqalion.com).
- **30+ Task Joint Reinforcement Learning** across vision-language-coding tasks (image captioning, VQA, code synthesis, STEM, grounding, video, GUI agents, coding agents) — maintains programming logic while scaling visual perception [VENDOR via secondary] (marktechpost.com; the-decoder.com; ubos.tech).
- Features: **thinking mode, streaming output, function calling, context caching** [SECONDARY] (the-decoder.com; aqalion.com).
- **Training hardware**: 100,000 Huawei Ascend 910B (MindSpore) — the same domestic-chip training claim as GLM-5.2 [SECONDARY] (aqalion.com).
- **Multimodal toolchain** (box drawing, screenshots, website reading incl. image understanding) completes the perception–planning–execution loop; **agentic meta-skills baked into pre-training** to strengthen action prediction early; a **multi-level, controllable, verifiable data system** was built to address the shortage of agent training data [VENDOR via secondary] (the-decoder.com).

### Benchmarks
- **Design2Code: 94.8** (headline, native design-to-code; cf. Opus 4.6 77.3 in the cometapi comparison table) [VENDOR via secondary] (cometapi.com; aqalion.com).
- SOTA on specialized agentic leaderboards: **CC-Bench-V2** (coding/repo exploration), **ZClawBench** (GUI agent interaction); #1 **WebVoyager / AndroidWorld** GUI-agent performance per the cometapi comparison [SECONDARY] (marktechpost.com; cometapi.com; ubos.tech).
- **#5 BridgeBench SpeedBench at 221.2 tokens/sec** — faster than Gemini 3.1 Pro, Claude Sonnet/Opus, and GPT 5.4 [SECONDARY] (agentnativedev.medium.com).
- **"Beats Opus 4.6 on multimodal benchmarks"** is a community/media headline, not an independently verified claim; one FAQ review explicitly cautions the Design2Code numbers are self-reported and unverified [SECONDARY/UNVERIFIED] (agentnativedev.medium.com; wavespeedai medium).

### Pricing
- **$1.20 input / $4.00 output per 1M tokens** [SECONDARY] (aqalion.com; cometapi.com; wavespeedai medium).
- Cost comparison (April 2026): **GPT-4o $2.50/$10.00** — ~2× cheaper input, 2.5× cheaper output; estimated **~$0.004 vs ~$0.027 per typical UI-coding task** [SECONDARY] (wavespeedai medium).
- Native **video input** (short clips alongside images/text in one context) — a stated differentiator over GPT-4o's frame-by-frame workaround [SECONDARY] (wavespeedai medium).
- Access: Z.ai native API (OpenAI-compatible) or **OpenRouter** [SECONDARY] (wavespeedai medium).

### Use cases cited
- Rapid prototyping (Figma → code), legacy UI migration (screenshots → modern React/Vue), **frontend regression checks** (layout misalignment, component overlap, color mismatches), CI pipelines feeding failing screenshots, autonomous web scrapers/form fillers/dashboard builders; early adopters report **70–90% time savings** on frontend tasks [SECONDARY] (cometapi.com).
- The training-data strategy (verifiable multi-level data system + pre-training meta-skills) suggests Z.ai treats agent-data scarcity as the binding constraint, not model scale [DIRECTIONAL].


### New verified facts — expansion (continued — GLM-5.2 deep spec (released June 2026))

### Release mechanics and training hardware
- Initial rollout to **paying subscribers June 13, 2026**; public weights release on **Hugging Face and ModelScope June 16, 2026**, under **MIT** [SECONDARY] (techtimes.com, 2026-07-13).
- Trained **entirely on roughly 100,000 Huawei Ascend 910B chips using the MindSpore framework — no NVIDIA hardware at any stage** [SECONDARY] (techtimes.com, 2026-07-13). This is a stronger, better-sourced Ascend claim than the ones already in §4 — record it as the primary Ascend-training citation.
- Z.ai (then Zhipu AI) has been on the **US Department of Commerce Entity List since January 2025** for its role in advancing Chinese military modernization through AI [SECONDARY] (techtimes.com, 2026-07-13).
- Parameter accounting: **744B total / ~40B active** per inference (techtimes, cryptobriefing); **753B** tensor-class figure (indianexpress, toknow.ai) — same checkpoint-accounting variance already noted in §4 [SECONDARY].

### Architecture details
- **IndexShare**: reuses the identical indexer across every **four sparse attention layers**; reduces per-token compute FLOPs by **2.9×** at long context [SECONDARY] (indianexpress.com citing the GLM-5.2 technical paper).
- **Upgraded Multi-Token Prediction (MTP) layer for speculative decoding**: boosts accepted token length by **up to 20%** during inference [SECONDARY] (indianexpress.com).
- **1M-token context — 5× GLM-5.1's 200K** — with a **131,072-token output limit**, among the largest practically usable open context windows [SECONDARY] (toknow.ai; medium/@cartseoservice).
- **Thinking modes: two tiers — High and Max** — selectable to balance speed vs depth [SECONDARY] (medium/@wenmingtech).
- Predecessor reference: **GLM-5.1 reached 77.8% on SWE-bench Verified** [SECONDARY] (medium/@cartseoservice).
- Positioned as **agent-first rather than chat-first**: function calling, tool use, browser automation, and multi-step API orchestration are the primary design target; conversational chat is secondary [SECONDARY] (medium/@cartseoservice).
- Language coverage: Python, JavaScript, TypeScript, C++, Java, Go, Rust, SQL, plus multilingual natural-language support [SECONDARY] (medium/@cartseoservice).

### Distribution and self-hosting
- Access: Zhipu's **GLM Coding Plan**, standalone API, plus third-party hosts **OpenRouter and Hugging Face Inference Providers**; OpenRouter pricing ran **~$1.40/M input / $4.40/M output** [SECONDARY] (medium/@cartseoservice).
- Repositories: **`zai-org/GLM-5.2`** plus an **FP8 build** on Hugging Face [COMMUNITY] (maxritter/pilot-shell blog).
- Self-hosting reality: **BF16 weights 1.51 TB**; FP8 build needs roughly **744–890 GB VRAM**; community **dynamic-1-bit quants land ~176–180 GB**; **Ascend NPU paths** exist alongside NVIDIA; quantized **GGUF via llama.cpp, Ollama, and LM Studio** [COMMUNITY/SECONDARY] (pilot-shell blog).
- Runnable through **SGLang, vLLM, Transformers, KTransformers, and Unsloth** [COMMUNITY] (pilot-shell blog).
- Z.ai marketed it as **"Pure Open"** — and unlike a hosted API, MIT weights cannot be switched off or geofenced [SECONDARY] (pilot-shell blog analysis).
- Adoption signal: headline claim of **40% of developer tokens** on the measured platform; platform user base **47% American**; **four of the five most-used models were Chinese** [SECONDARY] (techtimes.com, 2026-07-13 — platform-specific measurement, not global share).


