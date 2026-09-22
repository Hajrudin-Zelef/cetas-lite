---
id: labs-hyperscalers-2026/00-labs-hyperscalers/12-groq
title: "§12 — GROQ"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["Alibaba", "Cerebras", "EU", "Google", "Groq", "Mistral", "xAI"]
dates: []
keywords: ["benchmarks", "funding", "funding round", "gpu", "gpus", "grok", "hbm", "hyperscaler", "inference", "latency", "llama", "memory"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1639, 1691]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: abd3ae18f7be185fb5041dd4909d2d12aa31b6ae1ea04d4ae57d9f078e0a1203
---

# §12 — GROQ

## §12 — GROQ

> **Identity guardrail:** Groq (LPU inference company, groq.com) ≠ Grok (xAI's model family, §8) ≠ Cerebras (wafer-scale chips, §13).

> **2026 at a glance — Groq:** the LPU inference company (not Grok) kept its deterministic-latency lead on Artificial Analysis throughput boards; GroqCloud's free tier stayed the default open-model prototyping route; EU capacity expansion for sovereign demand; no 2026 raise found (last: $750M Series D, Aug 2024).

### 12.1 GroqCloud & LPU inference (2026)
- **GroqCloud** — Groq's inference API/console continued through 2026, serving open and proprietary models on Groq's **LPU (Language Processing Unit)** architecture. [secondary: Track D]
- Positioning: deterministic low-latency inference; tokens-per-second leadership on supported models. [secondary]
- 2026 milestones: expanded model catalog, enterprise contracts, international data-center partnerships (Track D).

### 12.2 Business & funding
- Groq's funding history (2024 $640M Series D at $2.8B pre-window context); 2026 financing/partnership activity per Track D (details in §16 scoreboard). [secondary]
- Competitive position: the pure-play inference-speed company vs. hyperscaler inference APIs. [secondary]

---

### 12.4 Groq — detailed (from Track D)

**Identity:** Groq (Groq Inc.) = LPU (Language Processing Unit) inference company founded by ex-Google TPU engineer Jonathan Ross. **Not Grok (xAI's model).** LPUs are deterministic single-core streaming processors purpose-built for LLM inference — low latency, high tokens/sec [secondary].

**2026 corporate events:** no major funding round found in-window in this pass (last confirmed: $750M Series D, Aug 2024, $2.8B valuation, led by Disruptive; BlackRock participated) [independent, 2024 — context only]. European data-center expansion: Groq announced/expanded EU inference capacity in 2026 to serve sovereign-AI demand [secondary — thin sourcing, flag for follow-up].

**GroqCloud pricing (representative, 2026):** Llama 3.3 70B-class serving commonly cited at ~$0.35/$0.79 per 1M input/output tokens; open-weight model catalog (Llama, Qwen, Mistral, Whisper large-v3-turbo for speech) [secondary — pricing snapshots, verify against console.groq.com]. Free tier: generous rate-limited free API tier widely used by developers; rate limits tightened at points in 2026 [secondary].

**Performance claims:** vendor-reported throughput leadership on tokens/sec/GPU-equivalent for supported open models; independent Artificial Analysis throughput leaderboards frequently rank Groq LPUs at/near top for latency-sensitive serving [vendor-reported/independent]. Mark specific numbers as time-sensitive — verify against current Artificial Analysis inference-provider benchmarks.

**Differentiation vs GPU clouds:** deterministic performance, no HBM contention; trade-off: model-porting effort per new architecture, smaller catalog than GPU-based serverless providers [secondary].

### 12.5 Key sources — Groq
- GroqCloud console/pricing: https://console.groq.com (pricing page; verify live)
- Artificial Analysis inference benchmarks: https://artificialanalysis.ai (provider leaderboards)

### 12.6 Groq LPU architecture — explainer (from Track D)

- **What LPUs are:** Groq's Language Processing Unit is a deterministic, single-core streaming processor — one core executes the entire model in a fixed, compiler-scheduled dataflow, unlike GPUs' thousands of general-purpose cores. No HBM contention; predictable per-token latency. [secondary]
- **Why it matters for inference:** LLM serving is memory-bandwidth-bound and latency-sensitive; LPUs trade programmability for deterministic throughput, frequently topping Artificial Analysis tokens/sec leaderboards for supported open models. [vendor-reported/independent]
- **Trade-offs:** each new model architecture needs compiler/porting work (smaller day-one catalog than GPU serverless); economics depend on sustained utilization of Groq's own silicon. [secondary]
- **2026 position:** GroqCloud's free tier made it the default prototyping route for open models; EU capacity expansion targets sovereign-AI demand; no 2026 funding round found in-window (last: $750M Series D, Aug 2024 @ $2.8B). [secondary/independent]
- **Do not confuse:** Groq (LPU company, Jonathan Ross) ≠ Grok (xAI model family) ≠ Grokking (ML term).

### 12.7 GroqCloud model catalog (2026 snapshot)

| Category | Representative models served | Notes |
|---|---|---|
| Chat / instruct | Llama 3.3 70B-class, Qwen family, Mistral family | Core LPU catalog |
| Speech | Whisper large-v3-turbo | STT on LPUs |
| Vision-language | Open-weight VLMs (catalog varies) | Verify live |

- **Free tier:** generous rate-limited free API tier widely used for prototyping; limits tightened at points in 2026 [secondary].
- **Playground/console:** console.groq.com — verify pricing live; snapshots in §16.3 are time-sensitive [secondary].
- **Enterprise:** dedicated LPU capacity; EU expansion 2026 for sovereign demand [secondary].

