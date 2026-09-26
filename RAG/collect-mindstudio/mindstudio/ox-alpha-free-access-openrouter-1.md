---
id: collect-mindstudio/mindstudio/ox-alpha-free-access-openrouter-1
title: "OX Alpha on OpenRouter: Free Access, Limits, and How Long It Lasts"
domain: mindstudio
role: reference
task: article
actors: ["China", "Nvidia", "OpenRouter", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["agentic", "attention", "benchmark", "benchmarks", "compute", "context window", "cost", "glm", "gpu", "inference", "multimodal", "nvidia"]
source: docs/RAG/Collect RAG/02_mindstudio/ox-alpha-free-access-openrouter.md
source_anchor: ""
source_lines: [1, 70]
sha256: fb0c9ef05163364c1ce03af8495428ef4dca0641facf8ba4ebbfdf61faa97e1f
---

# OX Alpha on OpenRouter: Free Access, Limits, and How Long It Lasts

## Metadata

- **Source** : https://www.mindstudio.ai/blog/ox-alpha-free-access-openrouter
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article examines OX Alpha, a stealth large language model that appeared on OpenRouter offering a 1 million token context window, multimodal input (text, image, and video), zero data retention, and a claimed daily capacity in the trillions of tokens — all at no cost. No company has officially attached its name to the model; OpenRouter's listing describes it only as a frontier model built for efficient coding, sustained agentic work, and real-world production use. That vagueness plus unusually generous usage limits generated significant attention among developers seeking a free way to run high-context, agentic coding workloads.

Free usage limits: the headline number circulating is a daily capacity in the range of 100 trillion tokens, making individual rate limits on OpenRouter extremely generous — described by early testers as "near unlimited." But there are still constraints: OpenRouter typically imposes per-user throttling on free/stealth models to prevent abuse, and free listings are usually temporary by design. Historically, labs testing models anonymously on OpenRouter treat the free window as a limited-time public benchmark, not a permanent offering. Anyone building against OX Alpha should treat access as provisional and avoid depending on it for production.

Why would a lab give away this much compute? Running a 1M-context multimodal model at near-unlimited scale for free requires enormous inference compute, and most AI labs are GPU-constrained, not flush with spare capacity. Theories: (1) Blind evaluation strategy — direct precedent: GLM's maker publicly described a prior anonymous release ("Pony Alpha") as a deliberate move to strip away brand bias and let raw output quality drive feedback; it worked — the stealth model built a reputation purely on performance before being revealed as GLM-5, helping validate that Chinese LLMs could compete at the frontier. If OX Alpha follows this pattern, free unbranded access is a way to generate organic, unbiased benchmarking buzz before a formal launch. (2) New compute backing — speculation about an undisclosed investment or new hardware partner enabling free inference at this scale. (3) Continual learning speculation — unverified theory that OX Alpha might keep updating from live usage rather than staying static after training (Safe Superintelligence, Ilya Sutskever's lab, has been reported working on continual learning); some users claimed signs of mid-deployment updates. Continual learning in deployed LLMs hasn't been demonstrated at this scale by any lab publicly.

Is it as good as GPT 5.5/5.6? The benchmark picture is mixed and inconsistent — itself a useful data point. One early test run on the DeepSWE coding benchmark reportedly scored ~80%, well ahead of comparison models. But other testers' private, contamination-resistant benchmarks showed significant underperformance — landing closer to parity with GPT 5.5-5.6 class models rather than beating them. Separate independent benchmarks placed it just below GPT 5.6 Sol-mid and just above GPT 5.5. Takeaway: OX Alpha appears to be a strong, competitive model in the same tier as current frontier systems, not something that clearly leapfrogs them; dramatic outperformance claims haven't held up consistently across retests; numbers should be treated as provisional since it hasn't been widely tested through a stable API by independent researchers.

What model is it built on? The most consistent theory connects OX Alpha to the GLM family, specifically "GLM 5.3 flash" per some leaks. Two evidence pieces: (1) Pliny the Liberator (known for jailbreaking LLMs to extract system prompts/internal configuration) reportedly found tokenizer evidence linking OX Alpha to the GLM family; (2) a prediction market reportedly assigned ~90% probability that OX Alpha resolves to Zed.ai (the company behind GLM). Prediction markets aggregate informed bets — a meaningful signal, though not proof. Other theories (link to Safe Superintelligence, or an Nvidia collaboration) lack comparable supporting evidence. If the GLM connection holds, there's a practical implication for local inference: a model performing near GPT 5.6 Sol-mid that could eventually run locally on capable hardware (electricity as the main ongoing cost) would be a meaningful development for the open-weight ecosystem.

Practical advice: testing OX Alpha while it's live on OpenRouter is low-risk (free). Treat it as a temporary window — precedent suggests stealth models get unmasked and repositioned commercially once the testing period ends. Build any real workflow on a named, stable model; use OX Alpha for exploratory testing or benchmarking prompts against a frontier-tier system while access lasts.

## Key points

- OX Alpha: stealth model on OpenRouter, no confirmed publisher — 1M token context, multimodal (text/image/video), zero data retention, free with near-unlimited rate limits (claimed ~100 trillion tokens/day).
- Most signals point to a GLM-family connection ("GLM 5.3 flash" per leaks): Pliny the Liberator tokenizer analysis + ~90% prediction-market odds on Zed.ai.
- Benchmarks are inconsistent: ~80% on DeepSWE in one early test; parity/lower on private contamination-resistant retests; roughly between GPT 5.5 and GPT 5.6 Sol-mid elsewhere.
- Leading "why free" theory: a blind evaluation strategy (precedent: "Pony Alpha" → revealed as GLM-5); new compute partner and continual-learning theories unverified.
- Free stealth listings on OpenRouter are typically temporary; treat access as provisional, don't build production on it.
- If it's GLM-family, a near-frontier open-weight model runnable locally would be significant for the open-weight ecosystem.

## Technical data / figures

| Feature | Value |
|---|---|
| Context window | 1 million tokens |
| Modalities | Text, image, video (multimodal input) |
| Data retention | Zero (claimed) |
| Daily capacity claim | ~100 trillion tokens |
| Rate limits | Described as "near unlimited" (per-user throttling still possible) |
| Cost | Free (as of listing) |

| Benchmark signal | Result |
|---|---|
| DeepSWE coding benchmark (one early run) | ~80% (ahead of comparison models) |
| Private contamination-resistant retests | Underperformance; parity with GPT 5.5-5.6 class |
| Independent benchmarks | Just below GPT 5.6 Sol-mid, just above GPT 5.5 |

| Evidence for GLM link | Detail |
|---|---|
| Tokenizer analysis | Pliny the Liberator (jailbreaking researcher) found GLM-family tokenizer evidence |
| Prediction market | ~90% probability OX Alpha resolves to Zed.ai (GLM's company) |
| Leaks | References to a "GLM 5.3 flash" checkpoint |
| Precedent | Prior anonymous "Pony Alpha" release revealed as GLM-5 |

| Theories for free access | Status |
|---|---|
| Blind evaluation strategy | Leading theory (precedent-based) |
| New compute partnership | Speculation, unverified |
| Continual learning | Speculation, unverified (SSI reported working on it) |

## Why this source matters for the RAG

Provides up-to-date facts about a free frontier-tier model available via OpenRouter (context window, modalities, rate limits, benchmark performance, likely origin) — relevant for model selection and cost reduction in agentic/RAG pipelines. Includes cautionary guidance on free stealth-model reliability and benchmark-score provisionality, useful when choosing retrieval/generation models.

## Related context from the article

