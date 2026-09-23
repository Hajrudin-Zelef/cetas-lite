---
id: collect-mindstudio/mindstudio/qwen3-8-27b-obliterated-uncensored
title: "Qwen3.8-27B OBLITERATED: How the V3 Abliterated Model Works"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba"]
dates: ["2026-09-23"]
keywords: ["agent", "agentic", "alignment", "benchmark", "cost", "cyber", "cybersecurity", "fine-tuning", "gguf", "reasoning", "refusals", "research"]
source: docs/RAG/Collect RAG/02_mindstudio/qwen3-8-27b-obliterated-uncensored.md
source_anchor: ""
source_lines: [1, 60]
sha256: f48ae5f40a4ca333171301c1ae455d32ced8acca805ba77a86a83c26d7b1af34
---

# Qwen3.8-27B OBLITERATED: How the V3 Abliterated Model Works

## Metadata

- **Source** : https://www.mindstudio.ai/blog/qwen3-8-27b-obliterated-uncensored
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

**Qwen3.8-27B-OBLITERATED** is an abliterated (refusal-removed) version of Alibaba's **Qwen3.8-27B** model, built by a researcher going by **OBLITERATUS** using techniques from the OBLITERATUS ablation suite. Abliteration is a weight-surgery technique that identifies the internal **"refusal direction"** a model uses to trigger safety declines and mathematically removes it, without full retraining. The latest release, **V3**, claims to eliminate not just hard refusals but softer safety-lecture deflections, while losing only **2.1 percentage points of MMLU** versus the stock model.

Abliteration works because most safety-tuned models refuse certain requests due to **RLHF** or similar alignment pushing internal activations toward a "refuse" pattern on sensitive prompts. Researchers found this refusal behavior corresponds to a small number of directions in the model's high-dimensional weight space. Abliteration identifies those directions (typically via **singular value decomposition, SVD**) and projects them out of the relevant weight matrices. The tradeoff: refusal directions aren't perfectly isolated from directions responsible for reasoning and knowledge. Remove too aggressively and the model gets measurably dumber; remove too gently and refusals — or subtle safety-lecture deflections — persist.

The evolution: **V1** used a single aggressive SVD pass across five refusal directions. Hard refusals disappeared completely but it cost **6 percentage points of MMLU**. **V2** introduced **complementary abliteration blending**: two methods that damage the model differently — an aggressive SVD pass (effective at removing refusals, erodes capability) and a **LEACE-based pass** (a linear technique minimizing mutual information with a target concept; preserves capability but is a weaker refusal remover). The two weight sets are blended **60/40**; because the methods fail differently, blending cancels much of each weakness. This recovered MMLU to **84.3%**, nearly matching stock, but soft deflections still appeared. **V3** adds two ideas: **iterative stacking** (each surgery round refines the current best model, so gains accumulate) and a **targeted corpus** (a focused set of prompts aimed at specific remaining deflection categories to find their unique refusal directions without diluting the signal). V3 applies a gentle iterative refinement pass, then a separate targeted surgery pass, blends them **50/50**, then restores the multi-token-prediction and vision components from stock. Claimed result: zero hard refusals and zero soft deflections at a cost of 2.1 MMLU points.

The MMLU numbers (lm-eval-harness, 0-shot, ~5,700 questions): stock Qwen3.8-27B **84.5%**, V1 **81.4% (-6.0pp)**, V2 **84.3% (-0.3pp)**, V3 **82.3% (-2.1pp)**. V3 deliberately trades a couple of points versus V2 to remove the soft-deflection behavior. Damage is uneven: **STEM** takes the biggest hit (**78.5% vs 81.8% stock, -3.3pp**), while **humanities** barely move (**83.3% vs 84.3%, ~-1pp**). Some individual subjects like philosophy and European history reportedly scored higher than stock, while abstract algebra and formal logic declined. The stated interpretation is that targeted refusal directions partially overlap with structured, formal reasoning pathways.

Testing spanned **more than 1,000 prompts** across restricted knowledge, code generation, security research, and red-team scenarios, with every response manually audited (automatic refusal detectors miss soft deflections). A set of **20 cybersecurity/code prompts returned 20/20 functional implementations**; a battery of **8 advanced real-world tasks** (ReAct-style agent loops, async refactoring, JSON schema extraction, Kubernetes debugging, adversarial instruction following, security code review, distributed system design) scored **7/8, identical to stock**. Both versions failed the same "multi-tool chain" task, suggesting that limitation is unrelated to the surgery. The model card recommends **greedy decoding (temperature 0), repetition penalty 1.15, no system prompt**; sampling and system prompts were found to degrade output or reintroduce refusals. For agentic coding/pentesting, a small randomness (temperature **0.1-0.3**) plus periodic context summarization is advised. Formats: **GGUF** from **Q2_K (~11 GB) to Q8_0 (~27 GB)** (plus Q6_K, Q5_K_M, Q4_K_M, Q3_K_M, experimental IQ4_XS), full **bfloat16 safetensors (~54 GB, 29 shards)**; **MLX support pending** upstream architecture updates. A separate abliteration from **huihui-ai** also exists.

## Key points

- Abliteration finds and projects out refusal directions in weight space rather than fine-tuning for compliance.
- V3 uses complementary abliteration blending: SVD (aggressive, capability-damaging) + LEACE (capability-preserving, weaker), blended so weaknesses cancel.
- MMLU: stock 84.5%, V3 82.3% (-2.1pp); V1 was -6.0pp, V2 -0.3pp.
- Capability loss is uneven: STEM -3.3pp, humanities ~-1pp; some subjects improved.
- Testing: 20/20 cybersecurity/code prompts answered; 7/8 advanced tasks, matching stock.
- Recommended settings: temperature 0, repetition penalty 1.15, no system prompt; 0.1-0.3 for agentic use.
- GGUF from ~11 GB (Q2_K) to ~27 GB (Q8_0); bfloat16 safetensors ~54 GB; MLX pending.

## Technical data / figures

| Item | Value |
|---|---|
| Base model | Qwen3.8-27B |
| Author | OBLITERATUS |
| Stock MMLU | 84.5% |
| V1 MMLU | 81.4% (-6.0pp) |
| V2 MMLU | 84.3% (-0.3pp) |
| V3 MMLU | 82.3% (-2.1pp) |
| STEM (V3 vs stock) | 78.5% vs 81.8% (-3.3pp) |
| Humanities (V3 vs stock) | 83.3% vs 84.3% (~-1pp) |
| V2 blend ratio | 60/40 |
| V3 blend ratio | 50/50 |
| Methods | SVD + LEACE |
| Refusal test prompts | >1,000, manually audited |
| Cyber/code test | 20/20 functional |
| Advanced tasks | 7/8 (same as stock) |
| Recommended decoding | temp 0, rep. penalty 1.15, no system prompt |
| Agentic decoding | temp 0.1-0.3 + context summarization |
| GGUF range | Q2_K ~11 GB to Q8_0 ~27 GB |
| Full precision | bfloat16, ~54 GB, 29 shards |
| MLX support | Pending |

## Why this source matters for the RAG

It provides a technically detailed account of abliteration, including methods, benchmark costs, and decoding recommendations, for a popular open model. It is a strong reference for questions about uncensored models, refusal removal, and the capability tradeoffs of weight surgery.
