---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-17
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "DeepSeek", "Google", "Groq", "Hugging Face", "Intel", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["agent", "agentic", "agents", "amd", "astra", "aws", "benchmarks", "blackwell", "claude", "consumer", "cost", "deepseek"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [941, 1024]
sha256: 8c04d09e39c42bf5005046b6e2447c9306056252f2d03a98ad63605f1b1675de
---

# ÉTAPE 1 — Open / Local AI Models (EN)

Sources: https://github.com/chsasank/blackwell · https://github.com/0xsero/blackwell-gpu-wiki/blob/HEAD/docs/blackwell/nvfp4-deep-dive.md · https://github.com/kubesimplify/website/blob/HEAD/content/blog/day-4-quantization-demystified-bf16-fp8-nvfp4-mxfp4-int4-gguf-and-why-it-all-matters.md · https://github.com/codehalwell/fable-skills/blob/HEAD/skills/llm-inference-optimization/SKILL.md · https://github.com/notwitcheer/llm-bench-rig/blob/HEAD/reports/fp4-consumer-blackwell.md

---

# 4. vLLM vs SGLang FOR SELF-HOSTING — 2026 RECAP

## 4.1 The two-contender market

- **TGI is out:** Hugging Face's Text Generation Inference accepts bug fixes only since Dec 2025 — no new features. HF Inference Endpoints now default to **vLLM**, with SGLang as the alternative. Two real contenders remain.

| Dimension | vLLM | SGLang |
|---|---|---|
| Origin / core trick | UC Berkeley — **PagedAttention** | LMSYS — **RadixAttention** (prefix tree) |
| Best for | Production multi-tenant serving; high-concurrency API; predictable latency | Structured-output workloads; multi-turn agents; shared-prefix workloads |
| 2026 benchmarks (H100) | Llama 3.3 70B FP8: 120→2,400 tok/s (conc. 1→100); Llama 3.1 8B: ~12,500 tok/s | Llama 3.3 70B FP8: 125→2,460 tok/s; Llama 3.1 8B: **~16,200 tok/s (+29%)** |
| Structured output | Noticeable overhead at high batch | **4.7× multi-turn speedup**, 99.8% JSON validity, overlapped mask gen |
| Memory (agent workloads) | baseline | **~47% less** (40 GB vs 75 GB) |
| Single-user structured | baseline | **6.6× faster** than vLLM |
| Hardware | NVIDIA, AMD, Intel, AWS Trainium, TPU | NVIDIA, AMD (400K+ GPUs deployed) |
| Day-0 model support 2026 | NVIDIA Nemotron 3 Ultra (Jun 2026); GLM-5.2 recipe targets vLLM 0.23.0 (stable) | **DeepSeek-V4 infer+RL (Apr 2026)**; GLM-5.2 (incl. NVFP4 Blackwell ckpt); Kimi-K2.7-Code; MiMo; Nemotron-H; LFM2.5 |
| Maturity | Very high; mature docs, Helm charts | High; Docker-first |

**2026 decision rule (community consensus):** vLLM for high-concurrency API serving where predictable latency matters; SGLang for agent loops, multi-turn conversations, and structured generation where prefix caching pays. "Ollama vs vLLM is easy; vLLM vs SGLang depends on your workload" — vLLM beats Ollama 16–29× in aggregate throughput at concurrency >10 (Ollama times out at ~20 concurrent).

## 4.2 VRAM examples in production serving

- **GLM-5.1 FP8 ≈ 800 GB** → 8× H200/H20 node (vLLM or SGLang, tensor-parallel).
- **DeepSeek V4.1 Flash:** 614 GB floor → 8× H200 (1,128 GB) or GB200 NVL4; vLLM ≥ 0.30 / SGLang day-zero / NVIDIA Dynamo.
- **Llama 3.3 70B FP8 on H100:** single GPU serves 100 concurrent at ~2,400–2,460 tok/s either engine.

Sources: https://techsy.io/en/blog/vllm-vs-sglang · https://privocto.com/blog/vllm-sglang · https://lushbinary.com/blog/glm-5-2-self-hosting-open-weights-vllm-guide/ · https://github.com/chipi/agentic-ai-homelab/blob/HEAD/docs/reading/self-hosting-llms.md · https://github.com/adilshamim8/genai-roadmap-with-notes-and-projects/blob/HEAD/fine-tuning-and-self-hosting/03-vllm-sglang-tgi.md

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

# 6. OPENROUTER DATA — OPEN MODELS DOMINATE 2026

## 6.1 Platform scale

