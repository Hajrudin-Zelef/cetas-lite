---
id: collect-mindstudio/mindstudio/ox-alpha-free-access-opencode
title: "Ox Alpha Free on OpenCode: Pricing, Limits, and How Long It Lasts"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Moonshot", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-09-23"]
keywords: ["pricing", "agent", "agentic", "benchmark", "benchmarks", "context window", "deepseek", "fine-tuning", "glm", "grok", "grok 4", "kimi"]
source: docs/RAG/Collect RAG/02_mindstudio/ox-alpha-free-access-opencode.md
source_anchor: ""
source_lines: [1, 47]
sha256: 9cf6c327cc7f4336e4b2b13f1fd886e49f3f58b2066ab8f16b5531bda4b267a1
---

# Ox Alpha Free on OpenCode: Pricing, Limits, and How Long It Lasts

---

# Ox Alpha Free on OpenCode: Pricing, Limits, and How Long It Lasts

## Metadata

- **Source** : https://www.mindstudio.ai/blog/ox-alpha-free-access-opencode
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article covers Ox Alpha, an anonymous "stealth" model that appeared on OpenCode with no confirmed provider attached. OpenCode lists it as free with a 1 million token context window, multimodal support, zero data retention, and a stated daily capacity of 100 trillion tokens. No lab has officially claimed it, but early testing suggests frontier-level performance, and community investigation points toward Zhipu AI's GLM model family.

The free access window runs about a week from launch, with the cutoff around August 27th. After that, the model is expected to either be formally revealed and priced or disappear if it was a testing exercise. This is a common pattern for stealth model drops: labs use free windows to gather real-world usage data and community feedback before committing to launch. The 100 trillion token daily capacity figure is less a hard cap and more a signal that the free tier is not artificially throttled; combined with the 1M context window, users can run large codebases and long agent sessions without aggressive rate limiting. Zero data retention is part of the offering, important for teams testing on proprietary code.

Benchmark results: On Kingbench (coding, math, generation tasks), Ox Alpha scored 70/80 (87.5%), second on the leaderboard behind GLM 5.3 (91.25%), ahead of Kimi K2 variants, Qwen 3.8 Max (81.25%), and Opus 4.8 (80%). It scored perfectly on a Three.js contact lens rendering case, an SVG generation task, a hard math permutation problem, and a Gemma fine-tuning task. Separately, developer Ben Davis ran Ox Alpha through a 10-task subset of Deep SWE (software engineering): Ox Alpha scored 80%, vs GLM 5.3 at 62%, Grok 4.6 at 62%, and GPT 5.6 Sol at 52%. Ox Alpha also solved a "Marriott task" in a single attempt where GLM 5.3, GPT 5.6 Sol, and Grok 4.6 all failed on four attempts. Caveat: small 10-task sample.

On origin: no lab has claimed the model, but Ben Davis's investigation (≈90% confidence) rests on technical fingerprints: video encoding token counts matching GLM 5V Turbo exactly (frame sampling, duration scaling ~147 tokens/sec, per-frame resolution scaling); tokenizer counts matching GLM 5.3 exactly across 25 prompts; emoji-heavy writing style consistent with GLM/Qwen families; rejection of audio input matching GLM 5V behavior (ruling out MiMo). This process of elimination excluded DeepSeek, Qwen, Xiaomi, and Western labs. No official confirmation exists.

Worth trying during the free window because there's no downside: 1M context, strong benchmarks across one-shot generation and agentic coding, no data retention concerns. Interesting wrinkle: Ox Alpha scored slightly below GLM 5.3 on general benchmarks but notably above it on agentic coding — if it's a next-gen GLM checkpoint, this suggests a model tuned more toward agentic/coding workflows than general-purpose generation.

## Key points

- Ox Alpha is a free, anonymous stealth model on OpenCode with 1M token context, multimodal support, zero data retention, and ~100 trillion tokens/day capacity.
- Free access period runs ~1 week, ending around August 27th; then reveal/pricing or disappearance.
- Kingbench score 87.5% (70/80) — 2nd overall behind GLM 5.3, ahead of Qwen 3.8 Max (81.25%) and Opus 4.8 (80%).
- Deep SWE 10-task subset: Ox Alpha 80% vs GLM 5.3 62%, Grok 4.6 62%, GPT 5.6 Sol 52%; solved Marriott task in one attempt where others failed four times.
- Fingerprinting (video tokens ~147 tokens/sec, tokenizer, style, audio rejection) points to a next-gen GLM from Zhipu AI at ~90% confidence — unconfirmed.
- Pattern: this is the 5th stealth release in a series; prior ones were all eventually claimed by Chinese AI labs.
- Likely tuned for agentic/coding workflows rather than general one-shot generation.

## Technical data / figures

- Context window: 1,000,000 tokens; multimodal; zero data retention; 100 trillion tokens/day stated capacity.
- Kingbench: Ox Alpha 87.5% (70/80); GLM 5.3 91.25%; Qwen 3.8 Max 81.25%; Opus 4.8 80%.
- Perfect Kingbench tasks: Three.js contact lens rendering, SVG generation, hard permutation math, Gemma fine-tuning.
- Deep SWE 10-task subset: Ox Alpha 80%, GLM 5.3 62%, Grok 4.6 62%, GPT 5.6 Sol 52%.
- Fingerprinting signals: ~147 tokens/sec video duration scaling matching GLM 5V Turbo; tokenizer matches GLM 5.3 over 25 prompts; audio input rejected.

## Why this source matters for the RAG

Documents a stealth-model release pattern and fingerprinting methodology used to identify anonymous models — directly useful for RAG entries on model provenance, benchmarking, and the GLM model family. It also provides concrete cross-model benchmark comparisons (Kingbench, Deep SWE) and free-access economics.

