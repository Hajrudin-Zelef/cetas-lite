---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-39
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Baseten", "Cerebras", "CoreWeave", "EU", "Fireworks AI", "Google", "Groq", "Hugging Face", "Meta", "Mistral", "Nebius", "Nvidia", "Together AI", "xAI"]
dates: ["2026-03"]
keywords: ["hyperscaler", "agent", "backlog", "benchmarks", "blackwell", "capex", "cost", "distribution", "fine-tuning", "funding", "funding round", "gpu"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1634, 1719]
sha256: f498bd6220ceb517047bff6949e28aeea0bf1081676e3f85bdeda39b40a995de
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

- **Vera Rubin platform:** NVIDIA's next-gen GPU platform after Blackwell; the March 2026 Thinking Machines deal commits ≥1 GW of Vera Rubin systems with deployment starting 2027 [independent].
- **Cost scale:** industry estimates ~$50B for 1 GW-class AI capacity [secondary] — the capex bar for next-gen training clusters.
- **DGX Spark:** personal AI supercomputer (GB10 Grace Blackwell Superchip desktop) shipping 2026; NIM-compatible local inference target for developers running Nemotron-class models at the desk [secondary].
- **Strategic read:** NVIDIA is simultaneously (a) selling the picks and shovels (Rubin), (b) giving away the models (Nemotron open weights), and (c) buying the distribution (Hugging Face, pending) — a full-stack capture of the open-model inference path. [research finding]

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

## §13 — INFERENCE HYPERSCALERS: CEREBRAS, SAMBANOVA, TOGETHER AI, FIREWORKS AI, NEBIUS, COREWEAVE

> **2026 at a glance — inference hyperscalers:** Cerebras IPO'd (May 13); SambaNova ($1B), Together AI ($800M) and Fireworks ($1.505B) raised in a nine-day July window; Nebius built out with reported Meta orders; CoreWeave levered its backlog into September converts; Together/Fireworks/Baseten/Modal/Databricks all served Inkling at launch.

### 13.1 Cerebras
- **Cerebras IPO** — the wafer-scale AI chip company's **initial public offering** took place during the window (Track D) — landmark liquidity event for the AI-silicon sector. [secondary]
- **WSE (Wafer-Scale Engine)** lineage: CS-3 systems; wafer-scale architecture as the differentiator vs. GPU clusters. [secondary]
- **Cerebras Inference** — cloud inference API on wafer-scale hardware; 2026 expansion. [secondary]
- Post-IPO partnerships and enterprise traction (Track D).

### 13.2 SambaNova
- **Series F** financing round during the window (Track D) — continued capitalization of the reconfigurable-dataflow (RDU) AI platform. [secondary]
- **SambaNova Cloud** inference offerings; sovereign-AI and enterprise positioning. [secondary]

### 13.3 Together AI
- **Series C** financing during the window (Track D). [secondary]
- Inference + fine-tuning API platform; open-model hosting leadership. [secondary]
- **Together AI is a separate company from Cerebras and Groq** — independent inference API provider. [research note]

### 13.4 Fireworks AI
- **Series D** financing during the window (Track D). [secondary]
- Fast inference API for open models; function-calling / agent infrastructure. [secondary]

### 13.5 Nebius
- **Capital stack** activity in 2026 (Track D): continued financing of the Nebius AI-cloud buildout (post-Volozh/Yandex split). [secondary]
- **Meta orders**: Nebius reported **Meta as a customer / order flow** in 2026 (see §4.4). [secondary]
- Positioned as the "neocloud" challenger alongside CoreWeave. [secondary]

