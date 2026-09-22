---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/xai
title: "xAI"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["Anthropic", "Cohere", "Falcon", "Glasswing", "Google", "Nvidia", "OpenAI", "TII", "United States", "xAI"]
dates: ["2025-08-05", "2025-08-26", "2025-09-30", "2025-12-11", "2025-12-15", "2026-01-05", "2026-01-15", "2026-01-29", "2026-02-12", "2026-02-19", "2026-03-20", "2026-03-31", "2026-04-23", "2026-04-26", "2026-05-20", "2026-06-01", "2026-06-03", "2026-06-10", "2026-07-09", "2026-08-11", "2026-08-12", "2026-08-13", "2026-09-01", "2026-09-02", "2026-09-03", "2026-09-21", "2026-09-22", "2026-09-24", "2026-11-21", "2027-01-01"]
keywords: ["agent", "apache", "astra", "claude", "cohere", "consumer", "diffusion", "fable 5", "gemini", "gpt-5.6", "gpt-6", "grok"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11240, 11316]
section: "§23. Master Model Inventory"
sha256: 54873a29093acd683d7a1b1c6f22405c439c6ecf7c246a341e0af1e3413cddc3
---

# xAI

### xAI

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Grok Code Fast 1 | 2025-08-26/28 | 314B MoE [VENDOR] | 256K | Closed API | Active; no 2026 successor found |
| Grok 4.6 | 2026-08-12 | — | 200K long-ctx tier | Closed API | Active; $2/$6 short ctx [VENDOR] |
| Grok 4.7 | 2026-09-21 | — | 200K tier cliff | Closed API | Active; $2/$6 ≤200K, $4/$12 above [VENDOR] |
| Grok 5 | — | — | — | — | Not shipped as of 2026-09-22 |

### OpenAI

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| gpt-oss | 2025-08 | 120B [SECONDARY] | — | Apache 2.0 · open | Active; 4.3M+ HF downloads [COMMUNITY] |
| GPT-5.2 (Instant/Thinking/Pro) | 2025-12-11 | — | — | Closed API | Active |
| GPT-5.3-Codex-Spark | 2026-02-12 | — | Text-only input | Closed API | Research preview |
| GPT-5.5 / 5.5 Pro | 2026-04-23 | — | — | Closed API | Active; $5/$30, $30/$180 [VENDOR] |
| GPT-5.6 Luna | 2026-07-09 (GA) | — | — | Closed API | Active; $0.20/$1.20 (was $1/$6) [VENDOR] |
| GPT-5.6 Sol | 2026 | — | — | Closed API | Active; $4/$20 promo (std $5/$30) thru 2026-11-21 [VENDOR] |
| GPT-5.6 Terra | 2026-09-02 | — | — | Closed API | Active; $2/$12 [VENDOR] |
| GPT-6 Astra | 2026-09-03 | — | 272K tier cliff | Closed API | Active; $10/$50; AA v4.3 53 (tie) [VENDOR/SECONDARY] |
| GPT-4o | — | — | — | Closed API | Legacy; $2.50/$10 [VENDOR] |
| Sora 2 | 2025-09-30 | — | — | Closed | Consumer wound down 2026-04-26; API end 2026-09-24 [SECONDARY] |

### Anthropic

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Claude Opus 5 | 2026-07 | — | — | Closed API | Active; $5/$25; SWE-bench Verified 97.0% [VENDOR/SECONDARY] |
| Claude Sonnet 5 | 2026 | — | — | Closed API | Active; $2/$10 (rise cancelled 2026-08-11) [SECONDARY] |
| Claude Fable 5.1 | 2026-09 | — | — | Closed API | Active; $10/$50; AA v4.3 53 (tie); TB 4.0 57.9% official #1 [SECONDARY] |
| Claude Mythos 5.1 | 2026-09-01 | — | — | Restricted (Glasswing) | $10/$50; TB 4.0 60.9% [VENDOR] |
| Claude Haiku 4.5 | 2026-09 | — | — | Closed API | Active; $1/$5 [VENDOR] |

### Google

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Gemini 3 Pro | 2025–2026 | — | — | Closed API | Active; FACTS Suite 68.8 (Dec 2025 run) [VENDOR/SECONDARY] |
| Gemini 3.1 Pro | 2026-02-19 (pricing) | — | 200K tier | Closed API | Active; $2/$12 ≤200K, $4/$18 above [VENDOR] |
| Gemini 3.5 Flash-Lite | 2026-09 | — | — | Closed API | Active; $0.30/$2.50 [SECONDARY] |
| Gemini 3.7/3.8 Flash | 2026 (3.8 intro 2026-08-13) | — | — | Closed API | Active; $0.75/$3.75 intro (doubles 2027-01-01) [VENDOR] |
| Gemma 4 (26B-A4B, 31B) | 2026-03-31 (weights) | 26B/3.8–4B; 32.7B dense [SECONDARY] | 256K | Apache 2.0 · open | Active; first Gemma off Gemma Terms |
| Gemma 4 (E2B, E4B) | 2026-03-31 | 5.1B/2.3B; 8.0B/4.5B [SECONDARY] | 128K | Apache 2.0 · open | Active |
| Gemma 4 12B "Unified" | 2026-06-03 | 11.95B dense [SECONDARY] | 256K | Apache 2.0 · open | Active; encoder-free multimodal |
| DiffusionGemma 26B-A4B | 2026-06-10 | 26B/~3.8B active (diffusion) [SECONDARY] | 256-token canvas | Apache 2.0 · open | Experimental; ~1,000 tok/s on H100 [VENDOR] |
| TranslateGemma (4B/12B/27B) | 2026-01-15 | 4B/12B/27B [SECONDARY] | — | Gemma Terms of Use | Active; not Apache 2.0 |
| Genie 3 | announced 2025-08-05 | — | — | Proprietary | Research preview; Project Genie 2026-01-29 |

### NVIDIA

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Nemotron 3 Nano | 2025-12-15 | 31.6B/3.2B (128 experts) [SECONDARY] | 1M | Open weight (NVIDIA) | Active; first MoE in the line |
| Nemotron 3 Super | 2026-03 | ~120B/12.7B [SECONDARY] | — | Open weight (NVIDIA) | Active |
| Nemotron 3 Ultra | 2026-06-01 | 550B/55B (Mamba-2 + LatentMoE + NVFP4) [VENDOR] | — | Open weight (NVIDIA) | Active; AA 48 at launch ("highest US-origin open model") [VENDOR] |
| Nemotron 3.5 Lightning | 2026-08-11 | 30B/3B [SECONDARY] | — | Open weight (NVIDIA) | Active; agent execution-layer model |
| Nemotron-Cascade 2 (30B-A3B) | ~2026-03-20/31 | 30B/3B [SECONDARY] | — | NVIDIA Open Model License | Active; IMO 2025 gold at 3B active [SECONDARY] |
| Nemotron 4 (1T) | — | ≥1T (reported) | — | — | In development only; not announced/shipped [SECONDARY] |

### Cohere

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Command A+ | 2026-05-20 | 218B/25B MoE [SECONDARY] | — | Apache 2.0 · open | Active; Cohere's first Apache-2.0 model |
| Command A / R / R+ | — | — | — | CC-BY-NC (research-only) | Legacy open posture |
| Command R+ | 2026-09 | — | — | Closed API | Active; $3/$15 [SECONDARY] |
| Cohere Rerank 4 | 2026 | — | 32K [SECONDARY] vs 4K [COMMUNITY] | — | Active; context figure unresolved |

### TII / Falcon

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Falcon-H1R / Falcon-H1 Arabic | 2026-01-05 | H1R 7B [SECONDARY] | — | — | Active; contradicts "no 2026 Falcon release" |
| Falcon Perception / Falcon OCR | ~2026-04 | — | — | — | Active (vision models) |
| Falcon-Perception-300M | 2026-07 | 300M | — | — | Active |

