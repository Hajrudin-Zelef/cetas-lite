---
id: open-local-models-2026/20-cost-comparison-self-hosting-vs-api/overview
title: "5. COST COMPARISON: SELF-HOSTING vs API"
domain: cost-comparison-self-hosting-vs-api
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Groq", "MiniMax", "Moonshot", "OpenAI", "Z.ai", "vLLM"]
dates: []
keywords: ["cost", "astra", "claude", "deepseek", "gemini", "glm", "gpt-5.6", "gpt-6", "gpu", "gpus", "kimi", "latency"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [975, 1020]
section: "5. COST COMPARISON: SELF-HOSTING vs API"
sha256: ec5014277e519cfab2c7a936df2102ec518506af0d736cdaaa1fd2a8c58dcf5a
---

# 5. COST COMPARISON: SELF-HOSTING vs API

## 5.1 The three-way break-even (2026)

Self-hosting beats a **frontier API** (GPT-5-class, Claude Sonnet 5/Opus, Gemini Pro) at roughly **5–8M tokens/day (~160–256M/month)** at 60–70% GPU utilization. Against a **budget open-weight API** ($0.14–$0.50/M tokens), it **rarely wins on cost at all**. "Self-hosting beats an expensive model long before it beats a cheap one."

| What you'd rent | $/month (Runpod Community Cloud, 720h) | vs API | Break-even volume |
|---|---|---|---|
| RTX 4090 24 GB | $244.80 | Together Llama 3 8B Lite ($0.14/M) | **1.75B tokens/month** (675 tok/s sustained) |
| RTX 4090 24 GB | $244.80 | gpt-4o-mini ($0.375/M) | 653M tokens/month |
| A100 80 GB PCIe | $856.80 | gpt-5.6-terra ($7.00/M blended) | 122M tokens/month |
| H100 80 GB PCIe | $1,432.80 | gpt-6-astra ($30/M blended) | **48M tokens/month** |

*(Rates read Sept 4, 2026; blended = even input/output split.)*

**DeepSeek V4 worked example:** 8× H100 node ≈ **$18,000/month** rented. vs V4-Flash API ($0.28/M output): break-even at **~65B output tokens/month** — a card kept near full load continuously. At 50M tokens/month the API costs ~$14 vs $18,000 self-hosted. "Stay on the API until you are spending well above $15,000–$20,000/month, then evaluate self-hosting."

**European bare-metal (Sept 2026):** from ~€234/mo (20 GB Hetzner GPU server, ≤14B models) → €1,199/mo (96 GB GEX131, 27–70B class) → $2,100–3,100/mo (rented H100). 27–70B self-host breaks even vs Claude Sonnet 5 ($3/$15) at ~10–15M tokens/day; vs Haiku 4.5 ($1/$5) at ~25–45M/day. Add 10–20% of a senior engineer for ops.

**Real deployment case study:** €38,500 hardware + €1,030/mo ops vs €10,900/mo equivalent frontier API → **break-even 3.9 months**, 81% cheaper over 3 years. Below ~€2,000/mo API spend, self-hosting rarely pays on cost alone.

## 5.2 Self-hosted per-token costs (mid-2026)

| Option | Input $/M | Output $/M | Notes |
|---|---|---|---|
| Open ~70B via Groq/Together (managed) | ~$0.60 | ~$0.80 | managed, fast |
| Open ~70B self-hosted (A100) | ~$0.20–0.40 | ~$0.60–1.20 | needs DevOps |
| Open ~8B self-hosted (L4) | ~$0.03–0.05 | ~$0.10–0.20 | high-volume simple tasks |
| vLLM self-hosted (general) | $0.50–1.00 blended | | per privocto |

**API reference (open models, Sept 2026):** DeepSeek V4.1 Flash $0.15/$0.60 off-peak ($0.30/$1.20 peak, $0.003 cache hits); GLM-5.3-Flash $0.15/$0.50; MiMo-V2.6-Pro $0.435/$0.87; MiMo-V2.6-Flash $0.14/$0.28; GLM-5.3 $4.40/M output; Kimi K3 $15/M output; Qwen3.8 Max $6/M output; MiniMax M3 $1.20/M output.

**Flat-rate disruptor — CheapestInference:** Core pool from **$15.29–17.99/mo** unlimited (DeepSeek V4.1 Flash + MiMo v2.5); Frontier pool **$60.35/mo** (GLM-5.3 + MiniMax M3); Flagship pool **$199/mo** per 8h daily block (Kimi K3 + Qwen3.8 Max). "No token caps during your reserved hours."

## 5.3 When local makes sense (decision framework)

**Self-host when:** sustained volume > break-even for your tier; GPU utilization >50–60%; data residency/compliance requires it; you need latency control (local TTFT 20–80ms vs 200–500ms cloud); you need any fine-tune/quantization/engine control; flat predictable cost beats per-token.

**Stay on API when:** volume < ~€2,000/mo spend equivalent; spiky/bursty traffic (can't scale to zero locally); no ops capacity (the real cost is the serving stack + on-call, not the GPUs); you need frontier quality (expect ~1pp quality give-up vs frontier on open models).

**Hidden costs of self-hosting:** idle GPU cost, model updates/drift, monitoring, on-call engineering, networking/storage, power (~$3.50/day for a 400W card), hardware amortization (~$20.55/day for a $25K A100 server over 3 years — "the anchor most 'local is cheaper' calculations ignore").

Sources: http://cloudzy.com/blog/self-hosting-open-weight-llm-gpu-vps-cost/ · https://aireviewrating.com/articles/self-hosting-llm-break-even · https://renezander.com/guides/self-hosted-llm-vs-api/ · https://ecorpit.com/deepseek-v4-self-hosted-vs-api-gpu-cost-break-even-2026/ · https://vallettasoftware.com/blog/post/self-hosted-llm-cost · https://promptcost.org/en/blog/local-llms-total-cost-ownership-2026/ · https://cheapestinference.com/ · https://cheapestinference.com/pools/core/ · https://cheapestinference.com/faq/ · https://cheapestinference.com/reports/state-of-open-weights/

---

