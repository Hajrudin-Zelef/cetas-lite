---
id: ai-industry-kb-2026-wave6/07-minimax/minimax-m2-7-self-evolving-agent-model
title: "MiniMax M2.7 — self-evolving agent model"
domain: minimax
role: deep-dive
task: agents
actors: ["AMD", "Alibaba", "Anthropic", "China", "Google", "Hugging Face", "MiniMax", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2025-12-30", "2026-01-01", "2026-01-09", "2026-04", "2026-04-12", "2026-06-03", "2026-06-18"]
keywords: ["agent", "amd", "attention", "benchmark", "benchmarks", "blackwell", "claude", "context window", "gemini", "glm", "gpu", "gqa"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3225, 3295]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: 8a25a9a8ff22c6379e765e73c98f9fd7ef8f73feaefff3ccb9439aa77085c031
---

# MiniMax M2.7 — self-evolving agent model

### MiniMax M2.7 — self-evolving agent model
- **Released early April 2026**; **officially open-sourced April 12, 2026** with weights on Hugging Face [SECONDARY] (marktechpost.com, 2026-04-12; particula.tech).
- **Self-evolving training**: post-trained with **100+ rounds of autonomous scaffold optimization on the open-source OpenClaw harness**, yielding a **30% performance improvement** without increasing base model size; MiniMax calls it the first model to actively participate in its own development [VENDOR via secondary] (marktechpost.com; particula.tech; venturebeat.com).
- **Activates only 10B parameters per token**; roughly **3× the throughput of Claude Opus 4.6** at similar SWE-Bench-family quality [SECONDARY] (particula.tech).
- Vendor benchmarks (all [VENDOR]): **SWE-Pro 56.22%** (matching GPT-5.3-Codex); **Terminal Bench 2 57.0%**; **SWE-Bench Verified 78%**; **SWE Multilingual 76.5%** [SECONDARY reporting vendor] (marktechpost.com; particula.tech; bofai docs).
- **GDPval-AA: Elo 1495** — claimed highest among open-source-accessible models [VENDOR] (news.saerio.com).
- **AA-Omniscience Index: +1**, described as a massive leap from **M2.5's −40** [VENDOR] (news.saerio.com).
- **Hallucination rate 34%**, vs **46% for Claude Sonnet 4.6** and **50% for Gemini 3.1 Pro Preview** [VENDOR] (news.saerio.com).
- **MM Claw: 97% skill adherence** across 40 complex skills (each >2,000 tokens), a substantial improvement over M2.5 [VENDOR] (news.saerio.com; marktechpost.com).
- Reasoning considered **equivalent to GLM-5 while using 20% fewer output tokens** [VENDOR] (news.saerio.com).
- **Medal rate 66.6%** in MiniMax's trials — tying Google's Gemini 3.1, approaching Claude Opus 4.6 SOTA [VENDOR] (venturebeat.com; news.saerio.com).
- Internal deployment: handles **30–50% of MiniMax's own RL research-team workflows autonomously**; company goal is **full autonomy in model training and inference architecture without human involvement** [VENDOR] (venturebeat.com; news.saerio.com).
- Finance use case: autonomously reads annual reports and earnings-call transcripts, cross-references research reports, designs assumptions, builds revenue-forecast models, and produces **templated PPT and Word research reports** like a junior analyst [SECONDARY] (marktechpost.com).
- Spec: **204.8K context window / 131.1K max output**; **dynamic tool search, multi-agent handoffs, dependency tracking** across parallel workstreams; SWE Multilingual 76.5% [COMMUNITY] (bofai/docs model page).
- **Text-only input — no multimodal image/video support** [COMMUNITY] (bofai/docs).
- Independent caveat: **BridgeBench (vibe coding) shows regression from M2.5** [COMMUNITY] (bofai/docs).
- **License conflict**: MarkTechPost reports M2.7 **"officially open source"** with weights on Hugging Face [SECONDARY], while the bofai/docs mirror states **"open weights are released under a non-commercial license, and commercial use requires a separate agreement"** [COMMUNITY] — record both; the checkpoint license file decides.
- Third-party reference pricing (B.AI credits/token): **input 0.30 / cache-write 0.375 / cache-read 0.06 / output 1.20** — reference prices, actual billing may be lower via bonuses [COMMUNITY] (bofai/docs).

### Platform mechanics (all models)
- **Dual protocol compatibility**: MiniMax serves both the **OpenAI (`/v1/chat/completions`)** and **Anthropic (`/anthropic`)** protocols simultaneously — migrate existing code with a one-line change [COMMUNITY] (vibe-investing guide).
- **Coding Plan subscription**: developer-only usage-based plan, **10–20× cheaper than OpenAI/Anthropic** [COMMUNITY] (vibe-investing guide).
- Tooling ecosystem setup: **VS Code (Cline / Claude Code / Continue / Kilo Code), JetBrains, OpenClaw, Cursor, Zed** [COMMUNITY] (vibe-investing guide).


### New verified facts — expansion

### MiniMax-M3 — full config-sheet architecture (from the released checkpoint config)
- **60 layers total; first 3 layers dense, layers 4–60 MoE; sparse attention disabled in the first 3 layers and enabled in layers 4–60** [SECONDARY] (semianalysisai inferencex-app model page reading the official config.json; morphllm.com).
- **128 routed experts, top-4 per token, one shared expert** [SECONDARY] (inferencex-app; morphllm.com).
- MSA (MiniMax Sparse Attention) index config: **Top-16 blocks, block size 128, 4 index heads, index dimension 128**; **GQA-group-specific block selection** [SECONDARY] (inferencex-app).
- **7 MTP modules** (num_mtp_modules: 7) plus one next-token-prediction layer [SECONDARY] (inferencex-app).
- Activation/norm: custom **SwiGLU variant** (`hidden_act: swigluoai`, swiglu_alpha 1.702, swiglu_limit 7.0) and **RMSNorm/Gemma-style norm** [SECONDARY] (inferencex-app).
- **RoPE applied to half of each 128-dim head** [SECONDARY] (inferencex-app).
- MSA technical report: **arXiv:2606.13392** [SECONDARY] (morphllm.com).
- Open-source MSA kernel: **github.com/MiniMax-AI/MSA** [SECONDARY] (morphllm.com).

### MiniMax-M3 — inference, checkpoints, and deployment
- Two weight builds: standard full build and **MiniMax-M3-MXFP8** at **~440 GB**; **MXFP8 targets NVIDIA Blackwell and AMD Instinct**, while the **BF16 full build targets NVIDIA Hopper (H200)** — check the serving recipe against the GPU generation before deploying [SECONDARY] (bytedance-iaas sglang cookbook; aitoolgrade.com review).
- Deployment frameworks confirmed in community/vendor docs: **SGLang, vLLM, Transformers, KTransformers, unsloth, ATOM** [SECONDARY] (sglang cookbook; startupfortune-style coverage; huggingface.co/MiniMaxAI/MiniMax-M3 card).
- Model card: **~428B total / ~23B active**, **1M-token context**, with a **guaranteed usable floor of 512K tokens** noted in secondary coverage [SECONDARY] (huggingface.co/MiniMaxAI/MiniMax-M3; medium/gptproto via morphllm context — the 512K floor is single-source, mark [UNVERIFIED]).
- Three thinking modes supported via the `thinking` parameter: **enabled / adaptive / disabled**; reasoning traces wrapped in **`<mm:think>...</mm:think>`** tags [VENDOR via secondary] (huggingface.co/MiniMaxAI/MiniMax-M3 card; sglang cookbook).
- **Native tool calling with a custom namespace-token XML format** and **`reasoning_content` parsing** confirmed in the SGLang integration — agent scaffolding must use this format, not generic OpenAI-style tool calls [SECONDARY] (bytedance-iaas sglang cookbook).

### MiniMax-M3 — production training-token contradiction (recorded, unresolved)
- One line of coverage quotes MiniMax (via MLQ.ai, June 3, 2026) as training on **roughly 100 trillion tokens using interleaved text and image sequences from the beginning** [SECONDARY, vendor-quoted].
- A separate June 18, 2026 report states the production M3's training-token budget **"had not been publicly disclosed"** as of that date, and clarifies that the widely repeated **109B-parameter / 3T-token** figures describe the **MSA research testbed model, not production M3** — the confusion traces to a June 17 MarkTechPost piece that mixed the two [SECONDARY] (techtimes.com, 2026-06-18).
- Record both; do not merge. The 109B/3T numbers belong to the research model in every citation [DIRECTIONAL].

### MiniMax-M3 — MSA speedup disambiguation
- Three speedup figure sets are in circulation and refer to **different experiment setups** — never quote one without its setup [DIRECTIONAL]:
  - **9.7× / 15.6×**: production-diagram figures [SECONDARY] (inferencex-app).
  - **9× / 15×**: rounded figures from the launch model card [SECONDARY] (morphllm.com).
  - **14.2× / 7.6×**: measured on the **109B research model** (the 3T-token testbed), not production M3 [SECONDARY] (techtimes.com, 2026-06-18).
- Headline vendor framing: MSA is positioned as the sparse-attention mechanism that makes the 428B/23B operating point economical [VENDOR via secondary] (morphllm.com).

### MiniMax-M3 — vendor benchmark set (all [VENDOR] unless independently rerun)
- **SWE-bench Verified 80.5; SWE-bench Pro 59.0; Terminal-Bench 2.1 66.0; MCP Atlas 74.2; OSWorld-Verified 70.06; BrowseComp 83.5; KernelBench Hard 28.8** [SECONDARY reporting vendor] (inferencex-app; codingfleet.com; pasqualepillitteri.it).
- Keep SWE-bench **Verified** (80.5) and SWE-bench **Pro** (59.0) strictly separate — different suites [DIRECTIONAL].
- Benchmark-hygiene example worth preserving: one comparison source **incorrectly juxtaposes M3's Terminal-Bench 2.1 score with GPT-5.5's Terminal-Bench 2.0 score and explicitly warns they are not comparable** — cite this as the canonical example of why version labels are mandatory [SECONDARY] (codingfleet.com).

### MiniMax-M3 — API pricing
- Standard API: **$0.60/M input / $2.40/M output**, cached input **$0.15/M** [SECONDARY] (aitoolgrade.com review; codingfleet.com).
- Launch promotion: **$0.30/M input / $1.20/M output**; promo cached-input figure **$0.06/M** appears in secondary breakdowns [SECONDARY] (aitoolgrade.com; YouTube breakdown via search coverage).
- Available through **MiniMax Code** (first-party coding agent), the **Token Plan**, and the **API** [SECONDARY] (aitoolgrade.com; secondary coverage quoting MiniMax).

### MiniMax — Hong Kong IPO facts
- Target: **up to HK$4.19B** from **25.4M shares at HK$151–165**; debut planned **January 9, 2026**; implied valuation **~$6.5B** [SECONDARY] (reuters.com, 2025-12-30).
- Cornerstone investors reportedly included **Alibaba and ADIA** (Abu Dhabi Investment Authority) [SECONDARY] (reuters.com, 2025-12-30).
- Part of a **year-end 2025 Hong Kong IPO rush** of Chinese AI firms [SECONDARY] (bworldonline.com, 2026-01-01).

