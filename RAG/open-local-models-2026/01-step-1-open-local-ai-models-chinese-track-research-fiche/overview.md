---
id: open-local-models-2026/01-step-1-open-local-ai-models-chinese-track-research-fiche/overview
title: "STEP 1 — Open / Local AI Models (Chinese track) — Research Fiche"
domain: step-1-open-local-ai-models-chinese-track-research-fiche
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Hugging Face", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "StepFun", "United States", "Unsloth", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-02-16", "2026-04", "2026-05", "2026-05-26", "2026-08-03", "2026-08-12", "2026-09", "2026-09-22"]
keywords: ["research", "agentic", "apache", "astra", "attention", "attribution", "awq", "benchmark", "benchmarks", "claude", "cost", "datacenter"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [7, 63]
section: "STEP 1 — Open / Local AI Models (Chinese track) — Research Fiche"
sha256: 400ea7d8cf54cb7e69e488056c6531bf1a4a5eaea05881ac2909d920974d8c0b
---

# STEP 1 — Open / Local AI Models (Chinese track) — Research Fiche
**Compiled: 2026-09-22 · Output: English · For RAG ingestion**
*All benchmark figures are vendor-reported unless marked as independently measured (Artificial Analysis, BenchLM, independent harnesses). Always store the provenance tag with the number.*

---

## 0. LANDSCAPE OVERVIEW — the open-weight frontier (September 2026)

- The open-weight frontier is **owned by Chinese labs**: DeepSeek, Moonshot AI (Kimi), Z.ai (GLM), Alibaba (Qwen), Xiaomi (MiMo), MiniMax, Tencent (Hy3), StepFun (Step 5 Preview). **Meta has left the open frontier** — Llama 5 has not shipped (forecast 2027); Meta pivoted to its first closed frontier model, Muse Spark (April 2026).
- **Artificial Analysis Intelligence Index v4.3 (Sept 2026) snapshot** — closed leaders: Fable 5.1 = GPT-6 Astra = 53, Opus 5 = 51, Muse Spark 1.3 = 48, GPT-5.6 Sol = 47. Open-weight leaders: **MiMo-V2.6-Pro = 46** (tied with closed Grok 4.7), **Qwen3.8 Max 0902 = 45, GLM-5.3 = 45, Kimi K3 = 44** (level with Grok 4.6), **DeepSeek V4.1 Flash ≈ 39–40, DeepSeek V4.1 Pro = 36**.
  - *Methodology warning: AA Index v4.1/v4.1.1 numbers (e.g., Kimi K3 57.1, GLM-5.3-Flash 57, Qwen3.6-Max-Preview 52, Qwen3.7-Max 56.6) are NOT directly comparable to v4.3 numbers (K3 44). Record the index version with every figure.*
- **OpenRouter**: Chinese open models went from ~2% of usage (2025) to >50% (mid-2026); US frontier models fell from 70% to 30%. Crossover week: Feb 9–15, 2026.
- **The "open" bifurcation**: leading open models are trillion-parameter datacenter-scale MoEs (GLM-5.3 753B, Kimi K3 2.8T, Qwen3.8 Max 2.4T, MiMo-V2.6-Pro 1.02T) needing 8–64+ accelerators; genuinely small models (Qwen3.6-27B, Qwen3.8-27B) run on one GPU, with little in between.
- **License tightening arc**: every Chinese family moved from MIT/Apache permissive terms toward custom licenses with revenue-gated MaaS clauses (Kimi K3: $20M; MiniMax M3: $20M; GLM-5.3: $10B security review; MiniMax H3: territorial restriction). GLM-5.3-Flash (MIT) is the notable counter-example.
- **SWE-bench milestones**: DeepSeek V4 Pro 80.6% Verified; MiniMax M3 80.5% Verified; GLM-5.2 62.1% Pro (vs GPT-5.5 58.6%); MiMo-V2.6-Pro DeepSWE v1.1 71.9.

---

## PART A — QWEN (Alibaba / Tongyi Qianwen)

**Naming convention (2026):** Qwen3.5 / 3.6 / 3.7 / 3.8 are successive flagship generations, not point releases. "Max" = top tier of each generation (proprietary until 3.8).

### A1. Qwen3.5-397B-A17B — released February 16, 2026
- **Architecture:** Sparse MoE — **397B total, 17B active/token**; 512 routed experts + 1 shared, 10 routed + 1 shared activated. **Hybrid attention:** 60 layers as 15 repeating groups of [3× Gated DeltaNet → MoE, 1× Gated Attention → MoE]. Multi-Token Prediction (MTP), MRoPE, vocab 248,320.
- **Context:** 262,144 native, extensible to 1M+ via YaRN. Extremely small KV cache (~3.9 GB at 128K ctx).
- **Multimodal:** Native vision-language via early-fusion training (text + image + video in); 201 languages/dialects.
- **Reasoning:** Thinking mode on by default, toggleable via API.
- **License:** **Apache 2.0**.
- **Hugging Face:** `Qwen/Qwen3.5-397B-A17B` (BF16), `Qwen/Qwen3.5-397B-A17B-FP8`; community quants (Unsloth, QuantTrio AWQ). Smaller sibling: Qwen3.5-27B.
- **Benchmarks (vendor):** SWE-bench Verified 76.2 · Pro 50.9 · Multilingual 69.3 · Terminal-Bench 2.0 52.5 · SkillsBench 30.0 · MMLU-Pro 87.8 · SuperGPQA 70.4 · GPQA Diamond 88.4 · HLE 28.7 · LiveCodeBench v6 83.6 · AIME26 93.3 · IFBench 76.5 (best in class) · BrowseComp 78.6 (best in class).
- **Distinguishing:** First open-weight Max-scale generation; Alibaba claimed 8.6–19× decoding throughput vs Qwen3-Max.

### A2. Qwen3.6 family (April 2026)
- **Qwen3.6-Plus** (Apr 2, 2026): 1M context, always-on reasoning, agentic coding + visual coding; integrates OpenClaw, Claude Code, Cline. Bailian pricing ~¥2/M in, ¥12/M out.
- **Qwen3.6-Max-Preview** (Apr 20, 2026): **closed weights**; 35B total / 3B active MoE; 256K context; #1 at launch on six benchmarks (SWE-bench Pro, Terminal-Bench 2.0, SkillsBench, SciCode, QwenClawBench, QwenWebBench). AA Index v4.1.1: **52**.
- **Qwen3.6-27B** (Apr 22, 2026): first **dense** open-weight Qwen3.6; **Apache 2.0**; HF `Qwen/Qwen3.6-27B`; 256K context; multimodal; runs on 1× H100/H200. SWE-bench Verified **77.2** (beats Qwen3.5-397B), Pro 53.5, Terminal-Bench 2.0 59.3.
- **Qwen3.6-35B-A3B:** MoE 35B/3B active; SWE-bench Verified 73.4.

### A3. Qwen3.7-Max — unveiled May 26, 2026 (API GA May 21–23)
- **Architecture:** Sparse MoE, **>1T parameters**, 1M context, 65,536 max output. Tailored for agentic workflows (35-hour autonomous runs claimed).
- **License:** **Proprietary / closed weights** (API only).
- **Benchmarks:** AA Index v4.1.1 **56.6** — global #5, #1 Chinese model (May 2026); IFBench 79.1 (record).
- **Pricing:** Alibaba Cloud ¥12/M in / ¥36/M out; OpenRouter **$2.50/$7.50**. Supports OpenAI- and Anthropic-Messages API formats.

### A4. Qwen3.8-Max — Preview July 19–20, 2026; GA August 3, 2026
- **Architecture:** Sparse MoE + hybrid attention — **2.4T total, ~95B active/token** (~4% activation); 92 layers, hidden 8192, **512 experts (10 routed + 1 shared)**; vocab 248,320.
- **Context:** 1M tokens (~991K max input; 131K max output; 262K reasoning budget). Multimodal in (text/image/video), text out — API only.
- **Pricing (API):** **$2/M input, $6/M output**. Updated checkpoint **Qwen3.8-Max-0902** (Sept 2): #1 Code Arena.
- **Benchmarks (vendor):** Terminal-Bench 2.1 86.6 · PaperBench 93.0 · SWE-bench Pro 67.7 · FrontierSWE 73.5 · DeepSWE 1.1 56.6 · GPQA Diamond 92.6. AA Index v4.3: **45** (AA-Briefcase Elo 1622, GDPval 1689, cost/task **$5.41** — the expensive end of the tier).
- **OPEN WEIGHTS — VERIFIED SHIPPED:** Published **August 12, 2026** on HF as **`Qwen/Qwen3.8-2.4T-A95B`** and **`-FP8`** (also ModelScope; Unsloth GGUF quants). First-ever open weights for a Max-tier Qwen flagship.
  - **License catch:** NOT Apache 2.0 — custom **"qwen3.8-max" license**: attribution above 100M MAU or $20M monthly revenue; MaaS/AI-assistant businesses above $50M trailing revenue need a separate commercial license.
  - **Capability gap vs API:** open checkpoint is **text-only, thinking-mode only** — no vision, no native 1M context (262,144 native, extensible to ~1,010,000).
  - **Hardware:** BF16 ≈ 4.89 TB, FP8 ≈ half; cheapest usable quant ≈ 450 GB combined RAM+VRAM; NVIDIA reference = GB300 NVL72 rack.
- **Companion:** **Qwen3.8-27B** — dense 27.8B, **Apache 2.0**, Terminal-Bench 73.0, ~14–16 GB VRAM at 4-bit (RTX 4090 friendly); HF `Qwen/Qwen3.8-27B`.

---

