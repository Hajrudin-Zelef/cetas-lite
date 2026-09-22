---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/new-verified-facts-expansion-continued-longcat-flash-thinkin
title: "New verified facts — expansion (continued — LongCat-Flash-Thinking-2601 deep spec)"
domain: longcat-and-meituan
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "China", "DeepSeek", "Google", "Huawei", "Hugging Face", "LongCat", "Meituan", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Z.ai"]
dates: ["2025-09", "2026-01", "2026-06-30", "2026-07-01", "2026-07-05"]
keywords: ["accelerator", "agent", "agentic", "amd", "ascend", "asic", "attention", "attribution", "benchmark", "claude", "compute", "context window"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2747, 2823]
section: "§6. LongCat and Meituan"
delta_of: ai-industry-kb-2026
sha256: 5fda600dbc5ecf228d51138f16185193da3a296cf1a6d4248ff875bbac04f5f8
---

# New verified facts — expansion (continued — LongCat-Flash-Thinking-2601 deep spec)

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


### New verified facts — expansion

### LongCat-Flash (original) — technical spec from the technical report
- **560B total parameters**; dynamic active compute **18.6B–31.3B, average ~27B**; **128K context window** [SECONDARY] (huggingface.co/docs/transformers model_doc/longcat_flash; akihikowatanabe/paper_notes #4995).
- Architecture: **shortcut-connected MoE (ScMoE)**; **zero-computation experts** implemented as **identity/skip experts** — experts that can be skipped entirely, so the dynamic active range is load-dependent rather than fixed top-k [SECONDARY] (transformers model doc; paper notes).
- Reported throughput: **>100 tokens/second** [SECONDARY] (transformers model doc).
- Vendor MMLU: **89.71** [VENDOR] (transformers model doc reporting vendor figures).
- Technical report: arXiv **2509.01322** (September 2025) [SECONDARY] (arxiv.org/pdf/2509.01322).
- The report's evaluation harness covers: MMLU, MMLU-Pro, ArenaHard, CEval, CMMLU (general); IFEval, COLLIE, Meeseeks (instruction following); MATH500, AIME24, AIME25, BeyondAIME (math); GPQA-diamond, DROP, ZebraLogic, GraphWalks (reasoning); HumanEval+, MBPP+, LiveCodeBench (2024.08–2025.05), SWE-Bench-Verified, TerminalBench (coding) [SECONDARY] (arXiv report §4.4.1).

### LongCat-Flash-Prover — formal-mathematics variant
- **LongCat-Flash-Prover**: 560B-class **open-weight MoE specialized for Lean4** theorem proving [COMMUNITY] (github.com/brndngln/longcat-flash-prover).
- Training framework: **Hybrid-Experts Iteration Framework** [SECONDARY] (paper notes #4995).
- RL method: **HisPO** (history-aware policy optimization) with **sequence-level and token-level gradient masking** to counter policy staleness and train/inference mismatch [SECONDARY] (paper notes #4995).
- Anti-reward-hacking: **theorem consistency and legality checks** in the reward pipeline [SECONDARY] (paper notes #4995).
- Vendor results: **MiniF2F-Test 97.1% with 72 attempts; ProverBench 70.8%; PutnamBench 41.5% with ≤220 attempts** — all [VENDOR] (paper notes #4995).
- Implication: the Prover variant shows the LongCat program extending beyond agentic coding into formal mathematics — a second pillar alongside the coding flagship [DIRECTIONAL].

### LongCat-2.0 — architecture deep spec
- **Shortcut-connected MoE (ScMoE)**; dynamic active **33B–56B per token, average ~48B** of 1.6T total [SECONDARY] (marktechpost.com, 2026-07-05; cryptobriefing.com).
- **Zero-computation experts** with **PID-controller adjustment** of the dynamic expert budget [SECONDARY] (marktechpost.com, 2026-07-05; venturebeat.com).
- **LongCat Sparse Attention** with three indexing mechanisms: **Streaming-aware, Cross-Layer, and Hierarchical indexing** — the source of the native 1M-token context, cutting long-context cost from quadratic toward linear [SECONDARY] (marktechpost.com, 2026-07-05).
- **135B-parameter N-gram embedding module**: expands the core embedding space roughly **100×**, capturing dense local token relationships and accelerating large-batch inference by reducing memory I/O bottlenecks [SECONDARY] (venturebeat.com).
- Post-training: **MOPD (Multi-Teacher Optimization via Mixture of Specialized Experts)** — fuses three teacher groups instead of blending raw human feedback into one reward function: **Agent Experts** (structural execution, tool invocation, multi-turn API parameter parsing, self-correcting loops), **Reasoning Experts** (multi-hop logic, chain-of-thought engineering, math, STEM), **Interaction Experts** [VENDOR via secondary] (venturebeat.com, 2026-07).
- Systems stack: **six-dimensional parallelism, disaggregated prefill/decode, super-kernels, L2 weight prefetch** [SECONDARY] (marktechpost.com, 2026-07-05; venturebeat.com).

### LongCat-2.0 — training, hardware, and release mechanics
- Training-token report: **35T+ tokens** appears in secondary coverage; the vendor's own materials reviewed did not confirm the figure in the sources checked — record as **[VENDOR, unconfirmed in reviewed sources]** and do not present as established [DIRECTIONAL].
- Hardware: training and inference ran on a **50,000-card domestic AI ASIC cluster with no NVIDIA hardware** (no A100/H100, no AMD MI300X); Meituan claims it is a **first of its scale for domestic Chinese hardware** [VENDOR via secondary] (cryptobriefing.com; marktechpost.com, 2026-07-05).
- Hardware-identity conflict: one French secondary source explicitly names **Huawei Ascend 910B**, but stronger sources say only "domestic chips/ASICs" — **do not promote "Ascend 910B" beyond a single-source [UNVERIFIED] attribution** [DIRECTIONAL] (ayinedjimi-consultants.fr PDF vs cryptobriefing.com/marktechpost.com).
- Unveiled **June 30, 2026** (cryptobriefing "reveals" piece); weights release **July 1, 2026** [SECONDARY] (cryptobriefing.com).
- License: **MIT** [SECONDARY] (marktechpost.com, 2026-07-05).
- Weights on Hugging Face under the **`meituan-longcat`** organization; GitHub repo **`meituan-longcat/LongCat-2.0`** [VENDOR] (github.com/meituan-longcat/LongCat-2.0).
- Chat access: official website **https://longcat.ai/** [VENDOR] (LongCat-2.0 GitHub README).
- Deployment: **GPU via SGLang cookbook; NPU via SGLang-FluentLLM** — the NPU path is a first-class vendor-supported target, consistent with the domestic-chip training story [VENDOR] (LongCat-2.0 GitHub README).
- The model had been **leading OpenRouter** (by usage) ahead of the open release — community traction preceded the weights drop [SECONDARY] (venturebeat.com headline reporting).

