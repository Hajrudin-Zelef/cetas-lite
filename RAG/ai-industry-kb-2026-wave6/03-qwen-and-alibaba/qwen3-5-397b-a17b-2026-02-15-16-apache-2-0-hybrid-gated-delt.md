---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/qwen3-5-397b-a17b-2026-02-15-16-apache-2-0-hybrid-gated-delt
title: "Qwen3.5-397B-A17B — 2026-02-15/16, Apache 2.0, hybrid Gated DeltaNet MoE [SECONDARY]"
domain: qwen-and-alibaba
role: deep-dive
task: architecture
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "OpenRouter", "Together AI"]
dates: ["2026-02-15", "2026-02-16", "2026-03-31", "2026-04", "2026-04-20"]
keywords: ["apache", "moe", "qwen", "agent", "agentic", "agents", "attention", "awq", "benchmarks", "claude", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1140, 1158]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: 2ffc6d102433c51078dfb3f673486d191ba6e7d502f3b8b44dd06fb81400e3ed
---

# Qwen3.5-397B-A17B — 2026-02-15/16, Apache 2.0, hybrid Gated DeltaNet MoE [SECONDARY]

### Qwen3.5-397B-A17B — 2026-02-15/16, Apache 2.0, hybrid Gated DeltaNet MoE [SECONDARY]
- **CONTRADICTION**: release date reported as **2026-02-15** (Puter developer docs) vs **2026-02-16** (Medium/dig.watch launch coverage). The model card is the tiebreaker to fetch; the corpus should use 2026-02-15/16 as a range until then. [SECONDARY]
- Architecture (per a Hugging Face community card transcribing the spec): **397B total / 17B active**, **512 routed experts with 10 routed + 1 shared active per token**, 60 layers in the pattern `15 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE))`, Gated DeltaNet with 64 V/16 QK linear-attention heads (dim 128), Gated Attention with 32 Q / 2 KV heads (dim 256, RoPE-64), expert intermediate dim 1024, embedding size 248320 (padded — consistent with the reported ~250k vocabulary). **Multi-step MTP** trained. [SECONDARY]
- Context: **262,144 natively, extensible to 1,010,000 tokens** — the 1M-scale extensibility is native to the architecture, not a serving hack. [SECONDARY]
- Positioning: **native vision-language foundation model via early-fusion training** on multimodal tokens; claimed **cross-generational parity with Qwen3** while beating Qwen3-VL on reasoning, coding, agents, and visual understanding; **8.6–19× faster decoding than Qwen3-Max** via the Gated DeltaNet + sparse MoE stack; **201 languages and dialects** (up from Qwen3's 119); **Apache 2.0** license with weights on Hugging Face and ModelScope. [SECONDARY]
- Hosted tier: **Qwen3.5-Plus** on Alibaba Cloud Model Studio — 1M context, built-in adaptive tool use. [SECONDARY]
- Vendor-claimed benchmarks (from the HF card, label as vendor claims): MMLU-Pro **87.8**, MMLU-Redux **94.9**, SuperGPQA **70.4**, C-Eval **93.0**, IFEval **92.6**, IFBench **76.5**, MultiChallenge **67.6**, GPQA **88.4**, HLE **28.7**, HLE-Verified **37.6**, LiveCodeBench v6 **83.6**, HMMT Feb 25 **94.8**, HMMT Nov 25 **92.7**, AIME26 **91.3**, BFCL-V4 **72.9**, TAU2-Bench **86.7**, MCP-Mark **46.1**, BrowseComp **69.0/78.6**, BrowseComp-zh **70.3**, WideSearch **74.0**, Seal-0 **46.9**, MMMLU **88.5**. The card tabulates Qwen3.5-397B-A17B against GPT-5.2, Claude 4.5 Opus, Gemini 3 Pro, Qwen3-Max-Thinking, and K2.5-1T-A32B. [SECONDARY]
- Community serving: on **4× RTX PRO 6000 (96GB)**, a local-inference recipe reports **259 tok/s single-user** (AWQ + b12x decode + EAGLE MTP), 180 tok/s NVFP4 + MTP, and 1550 tok/s at 64 concurrent users — evidence the 397B/17B footprint is servable on prosumer 4-GPU hardware. [COMMUNITY]
- **Release mechanics**: announced **2026-02-16** via the Qwen blog ("Qwen3.5: Towards Native Multimodal Agents"); GitHub QwenLM/Qwen3.5 news entry dated **2026-02-16** confirms "The first release includes a 397B-A17B MoE model." Third-party trackers: OpenRouter lists **February 16**, Artificial Analysis lists **February 17** — the 02-15/02-16 contradiction in the corpus narrows to a **02-16/02-17** range with the official blog as tiebreaker (02-16). [SECONDARY]
- **Pricing** (third-party): **$0.60 input / $3.60 output** per 1M (Together AI, Novita, OpenRouter); Puter docs list **Feb 15, 2026** release date (the 02-15 source). Max output **64–66K tokens**. Input: text/image/video; output: text. Function calling + structured output supported. [SECONDARY]
- **Competitive frame** (sci-tech-today): vs **Kimi K2.5** (1T/32B, $0.60/$3.00, Modified MIT, AA 46) vs **DeepSeek V3.2** (685B/37B, $0.28/$0.42, text-only) — Qwen3.5 at **AA 45**, the only one with **native vision-language + 201 languages + Apache 2.0**. [SECONDARY]
- **Architecture detail** (per SemiAnalysis inferencex-app model card): **3:1 Gated DeltaNet to Gated Attention hybrid stack**, **262K native context**, **Apache 2.0 weights**. [SECONDARY]
- **Qwen3-VL-235B-A22B** (companion vision model): **8.6× faster** than Qwen3-Max at 32K tokens, **19× faster** at 256K tokens — the efficiency story behind the native vision-language claim. [SECONDARY]

### Qwen3.6 — first closed-weights flagship, April 2026 [SECONDARY]
- **2026-04-20** — **Qwen3.6-Max-Preview** released: the first Qwen flagship to ship **closed-weights only**, breaking the Apache-2.0 pattern. Vendor-claimed top ranks on **six coding/agent benchmarks**: SWE-Bench Pro, Terminal-Bench 2.0, SkillsBench, QwenClawBench, QwenWebBench, SciCode. Artificial Analysis Intelligence Index: **52**. Specs: **260K context**, OpenAI + Anthropic API compatible, **`preserve_thinking`** for multi-turn agents. [SECONDARY]
- **Qwen3.6-35B-A3B FP8** — the **first open-weight release in the 3.6 series**: 35B total / **3B active**, Gated DeltaNet + gated-attention hybrid, vision input, **262K native context extensible to 1M**, `preserve_thinking` option, served in FP8 on Together AI. Benchmarks (Together AI model card): **SWE-bench Verified 73.4%**, LiveCodeBench v6 80.4%, Terminal-Bench 2.0 51.5%, MMMU 81.7%, MathVista (mini) 86.4%, GPQA Diamond 84.1%, HLE 20%, SciCode 36%, Terminal-Bench 2.1 45%. [SECONDARY]
- **Qwen3.6-Plus Preview** (2026-03-31): agentic-coding + reasoning flagship with 1M context — the March 30/31 double drop with Qwen3.5-Omni. [SECONDARY]

