---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/longcat-2-0-preview-2026-04-20
title: "LongCat-2.0-Preview (2026-04-20)"
domain: longcat-and-meituan
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "China", "DeepSeek", "Google", "Hugging Face", "LongCat", "Moonshot", "OpenAI", "Z.ai"]
dates: ["2026-01", "2026-04-20", "2026-06-30", "2026-09-02", "2026-09-10"]
keywords: ["accelerator", "agent", "agentic", "attention", "benchmark", "claude", "deepseek", "distribution", "gemini", "glm", "gpu", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2733, 2786]
section: "§6. LongCat and Meituan"
delta_of: ai-industry-kb-2026
sha256: b7a1c6880c959b9b5c199bec7859c11457029b1be9d286cb7128cd1d1e837cff
---

# LongCat-2.0-Preview (2026-04-20)

### LongCat-2.0-Preview (2026-04-20)
- Preview launch: agent development with native tool calling, multi-step reasoning, long-context tasks; strong in code generation, automation workflows, complex instruction execution; deeply integrated with **Claude Code, OpenClaw, OpenCode, Kilo Code** [VENDOR] (ChangeLog).
- Quota: **5,000,000 tokens/day** initial; submitting valid model feedback can earn quota refreshes up to **120,000,000 tokens/day** [VENDOR].
- Integration list here (Claude Code/OpenClaw/OpenCode/Kilo Code) vs the 2.0 release entry which adds **Hermes** — record both [VENDOR].

### LongCat-2.0 release and billing (2026-06-30)
- LongCat-2.0 release; billing went live with two options: **Token Pack** (fixed quota upfront, valid **30 calendar days** from purchase — short-term high-volume) and **API Pay-As-You-Go** (top-up balance, charged on actual consumption — variable workloads) [VENDOR] (ChangeLog).
- 2.0 positioned as: **trillion parameters, 1M long context**; native tool calling and multi-step reasoning; superior coding; deeply compatible with **Claude Code, Hermes, OpenClaw, OpenCode, Kilo Code** [VENDOR].

### Enterprise services (2026-09-02 / 2026-09-10)
- **Enterprise Verification** (legal-representative face verification or corporate bank transfer); verified enterprises get **self-service invoicing** (VAT special invoices) and **tiered discounts on Pay-As-You-Go** plus dedicated services [VENDOR] (ChangeLog).
- Two changelog entries (Sept 2 and Sept 10) both titled "Introducing Enterprise Services" — the Sept 10 entry adds that prior benefits remain unchanged and points to FAQ docs; treat as rollout refinement, not a duplicate launch [VENDOR].


### New verified facts — expansion (continued — LongCat-Flash-Thinking-2601 deep spec)

### Model identity and training recipe
- **LongCat-Flash-Thinking-2601**: 560B total / **27B activated on average per token**; described as a powerful, efficient MoE reasoning model with strong agentic reasoning capability [VENDOR] (huggingface.co/meituan-longcat/LongCat-Flash-Thinking-2601; arXiv 2601.16725).
- Technical report: **arXiv 2601.16725** (January 2026) [SECONDARY] (arxiv.org/pdf/2601.16725v1.pdf).
- Pretraining **largely follows the LongCat-Flash-Chat recipe, retaining the original data distribution** to preserve general reasoning, then extends toward large-scale agentic reasoning through a **mid-training stage** [VENDOR] (arXiv 2601.16725 abstract).
- Mid-training rationale: agentic behaviors involve long-horizon trajectories with proactive tool invocations, but such interaction patterns are **extremely scarce in real-world corpora** (mostly natural language) — so the model is exposed to **moderate-scale synthesized structured agentic trajectories** during mid-training as initialization for RL [VENDOR] (arXiv 2601.16725).
- Developed by the **DORA system**: an efficient **distributed RL framework supporting asynchronous training and flexible accelerator usage** for stability and efficiency [VENDOR] (github.com/meituan-longcat/LongCat-Flash-Thinking).
- Two-phase pipeline: **(1) Long CoT Cold-Start Training** — curriculum learning during mid-training to bolster intrinsic capabilities, then SFT on reasoning-intensive and agentic data; **(2) RL stage** [VENDOR] (LongCat-Flash-Thinking GitHub).
- Equipped for **formal reasoning and agentic reasoning**: mathematics, logic, programming, automatic theorem proving, tool use [VENDOR] (LongCat-Flash-Thinking GitHub).
- Training strategy per Chinese press: **"environment expansion + multi-environment reinforcement learning"** with diversified high-intensity environments; **noise injected into training data** to harden robustness against API call failures and missing data [SECONDARY] (aibase.com/news/24677).
- New evaluation method proposed by the team: **automated task synthesis** — randomly generating complex tasks from keywords and evaluating the model on them; 2601 maintained leading performance across randomly generated tasks [SECONDARY] (aibase.com/news/24677).
- Opened: **model weights, inference code, and online experience**; distributed via **GitHub, Hugging Face, and ModelScope**; try at **https://longcat.ai** [VENDOR/SECONDARY] (aibase.com; HF model card).

### LongCat-Flash-Thinking-2601 — vendor benchmark table (HF model card; [VENDOR])
Comparison set in the card: DeepSeek-V3.2-Thinking (671B/37B), Kimi-K2-Thinking (1T/32B), Qwen3-235B-A22B-Thinking-2507 (235B/22B), GLM-4.7-Thinking (355B/32B), Claude-Opus-4.5-Thinking, Gemini-3-Pro, GPT-5.2-Thinking-xhigh. Footnotes: `*` cited from official report, `†` reproduced results, `‡` multi-value (attempts variants).
- **Mathematical Reasoning w/ Tools**: AIME-25 (Avg@16) **99.6 / 100.0‡**; HMMT-25 (Avg@16) 93.4 / 97.5‡; IMO-AnswerBench (Avg@4) 78.6 / **86.8‡**; AMO-Bench EN (Avg@16) 61.6 / 66.0‡; AMO-Bench CH (Avg@16) 56.8 / 67.5‡.
- **Agentic Search**: BrowseComp (Pass@1) 56.6 / **73.1**; BrowseComp-zh (Pass@1) **69.0 / 77.7**; RW Search (Pass@1) 79.5.
- **Agentic Tool Using**: τ²-Retail (Avg@4) 88.6; τ²-Airline (Avg@4) **76.5**; τ²-Telecom (Avg@4) **99.3**; τ²-Avg (Avg@4) 88.2; τ²-Noise (Avg@4) **67.1**; VitaBench (Avg@4) 29.3; VitaBench-Noise (Avg@4) 20.5; Random Complex Tasks (Avg@4) **35.8**.
- **General QA**: HLE text-only (w/o tools) 25.2; GPQA-Diamond (Avg@16) 80.5 / 85.2‡.
- Programming: **LCB (LiveCodeBench) evaluation 82.8**, ranking among the top models in its category [SECONDARY] (aibase.com/news/24677).
- AIME-25 **perfect 100** highlighted in launch coverage as consolidating its leading position in mathematical reasoning [SECONDARY] (aibase.com).

### LongCat-Flash-Thinking-ZigZag — sparse-attention variant
- **~50% of full-attention layers replaced with SSA (sparse) layers**; remaining layers retain **MLA-based full attention**; layer-level (not head-level) sparsity avoids computational imbalance and GPU thread divergence [SECONDARY] (arXiv 2601.16725 via proxy text).
- Sparsification procedure: a calibrated dataset estimates relative attention-layer importance; the **lowest-importance subset is replaced with SSA layers**; then continued long-context mid-training [SECONDARY] (arXiv 2601.16725).
- Attention params: **block size 128, 1 sink block + 7 local blocks = 1,024-token effective span per layer** [SECONDARY] (arXiv 2601.16725).
- **YaRN-based positional encoding extension** enables extrapolation to **1M-token context** [SECONDARY] (arXiv 2601.16725).
- Yields about a **1.5× end-to-end inference speedup** while preserving reasoning and agentic benchmark performance [SECONDARY] (arXiv 2601.16725).
- Alternating sparse/full layers create a **zigzag-shaped connectivity path** along the sequence — global information preserved through cross-layer composition despite per-layer sparsity [SECONDARY] (arXiv 2601.16725).
- In the vendor comparison table, **ZigZag has sparse attention ✅ while 2601 has ❌** — ZigZag is the sparse-attention counterpart of the same base [VENDOR] (HF ZigZag model card).
- ZigZag card figures (selected): AIME-25 99.2; HMMT-25 93.5; AMO-Bench EN 60.4; CH 58.3; BrowseComp 55.2; BrowseComp-zh 71.9; τ²-Retail 86.8; Airline 76.5; Telecom 97.4; Avg 86.9; HLE 25.8; GPQA-Diamond 80.6 [VENDOR] (huggingface.co/meituan-longcat/LongCat-Flash-Thinking-ZigZag).
- Hugging Face: **meituan-longcat/LongCat-Flash-Thinking-ZigZag** [VENDOR].

### Community Apple-silicon port
- **inferencerlabs/LongCat-Flash-Thinking-2601-MLX-5.5bit**: community MLX quant with measured perplexities — q8.5: **1.128**, q6.5: **1.128**, q5.5: **1.141**, q4.5: **1.168**, q3.5: **1.900**, q2.5: **41.293** (collapse below q3.5) [COMMUNITY] (Hugging Face).
- Tested on **M3 Ultra 512GB RAM**: single inference **~23 tok/s** @ 1000 tokens; batched **~30 tok/s** across two inferences; **~362 GB memory** [COMMUNITY].
- The quant was **archived/removed from HF due to storage restrictions** — availability is not guaranteed [COMMUNITY].


