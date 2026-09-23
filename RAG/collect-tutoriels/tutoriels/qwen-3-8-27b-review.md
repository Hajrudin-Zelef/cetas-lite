---
id: collect-tutoriels/tutoriels/qwen-3-8-27b-review
title: "Qwen3.8-27B Review: The Open-Weight 27B Worth Running"
domain: tutoriels
role: reference
task: review
actors: ["AMD", "Alibaba", "Anthropic", "Hugging Face", "OpenAI"]
dates: ["2026-08", "2026-09-23"]
keywords: ["open-weight", "qwen", "agentic", "amd", "apache", "attention", "benchmark", "benchmarks", "claude", "cost", "fp8", "gguf"]
source: docs/RAG/Collect RAG/07_tutoriels/qwen-3-8-27b-review.md
source_anchor: ""
source_lines: [1, 69]
sha256: 9ba6efde453624d95749a1fa82e7fb939da424f3542e662b35a9dcb3b8c7fee1
---

# Qwen3.8-27B Review: The Open-Weight 27B Worth Running

## Metadata

- **Source** : https://www.orcarouter.ai/blog/qwen-3-8-27b-review
- **Site** : OrcaRouter
- **Type** : Review
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This OrcaRouter review by Alistair Wren (published 15 August 2026) argues that Qwen3.8 27B is the best open-weights 27B model available for self-hosting, and "not particularly close." Weights dropped 14 August 2026 under Apache 2.0: a dense 27B model (28B counting the vision encoder) with 262K native context, native image and video input, and vendor-reported agentic-coding scores at or above Claude Opus 4.6 Max — SWE-bench Pro 61.7, DeepSWE 1.1 42.2, LiveCodeBench v6 90.3. The headline price is zero: download 55.6 GB and run it. However, the caveats are real: independent testing finds it roughly three times slower and more token-hungry than its predecessor, every benchmark on the card is Alibaba-reported, and the 1M context is a hosted-only feature.

The fast facts (checked 15 August 2026): released 14 August 2026 at Qwen/Qwen3.8-27B on Hugging Face, mirrored on ModelScope (some coverage cites August 13, about a week after the Qwen3.8 generation was announced). License is Apache 2.0 with explicit patent grant. Parameters: 27B dense (28B with vision encoder), 64 layers, hidden size 5,120, vocabulary 248,320. Architecture is hybrid attention: 48 Gated DeltaNet linear-attention layers against 16 full Gated Attention layers (a 3:1 split), which is why a dense 27B can carry 262K native context. Context is 262,144 tokens natively, extendable to 1,000,000 via YaRN on the hosted version. Native image and video input, text output. Thinking mode is on by default and disableable, with `reasoning_effort` (low/medium/high) and `preserve_thinking`. Weights are 55.6 GB of BF16 safetensors in 18 shards, with FP8 and community GGUFs.

The model's strongest case is agentic software engineering: Alibaba reports a 3x jump on DeepSWE 1.1 over the previous 27B (42.2 vs 13.3) and a score above Claude Opus 4.6 Max on SWE-bench Pro (61.7 vs 53.4), earning the "Opus at home" nickname. Three further strengths: native multimodal input, 262K native context with lower KV cost via hybrid linear attention, and free permanent weights (runs at 4-bit on a 24 GB card, RTX 3090/4090-class; AMD shipped Day-0 support up to 24.5 tokens/s on Ryzen AI Max+ 395 and 51.8 tokens/s on Radeon AI PRO R9700).

Where it gets ugly: independent testing (one tester's harness, judged by GPT-5.5 via llmcompare) saw 3.8 win 9 of 10 real-world tasks, averaging 8.838 vs 6.862, but using almost three times the tokens and being considerably slower — one task took roughly 600% longer. Also: no third-party benchmarks yet; BF16 needs an 80 GB-class GPU; quantized builds reportedly "lose focus after long context"; 1M context is hosted-only; and "beats Opus" needs qualifiers.

The review places the model in the Qwen 3.8 family (vs Qwen3.6-27B, Qwen3-Coder-30B-A3B, and the 2.4T Qwen3.8-Max), settles the local-vs-API cost question (marginal token cost is zero on your hardware; OrcaRouter lists a free rate-limited tier plus a paid tier at $0.33/M input and $2.40/M output, checked August 15), gives how-to-try steps (free tier, community GGUF Q4_K_M ~17 GB, full precision via `huggingface-cli`), lists who should skip it, and gives the bottom line. The OrcaRouter model page, read that day, showed a 33.3% error rate over the trailing seven days and p50 first-token latency of 225 ms — a brand-new model under early load, not a production contract.

## Key points

- Qwen3.8 27B: dense 27B (28B with vision encoder), Apache 2.0, weights released 14 August 2026.
- 262K native context (1M via YaRN, hosted-only), native image and video input.
- Hybrid attention: 48 Gated DeltaNet linear layers + 16 full Gated Attention layers (3:1).
- Vendor-reported: SWE-bench Pro 61.7, DeepSWE 1.1 42.2, LiveCodeBench v6 90.3 — all Alibaba-reported, not independently reproduced.
- Independent testing: wins 9/10 tasks but ~3x tokens and much slower (one task ~600% longer).
- 55.6 GB BF16 in 18 shards; 4-bit fits a 24 GB card; Q4_K_M GGUF ~17 GB.
- Hosted on OrcaRouter: free rate-limited tier (HTTP 429 past cap), paid $0.33/M input, $2.40/M output.
- Skip if: no GPU, need an SLA, need verified benchmarks, need >262K open-weights context, or throughput-bound.

## Technical data / figures

| Spec | Value |
| --- | --- |
| Model | Qwen3.8-27B |
| Parameters | 27B dense (28B with vision encoder) |
| Layers | 64 |
| Hidden size | 5,120 |
| Vocabulary | 248,320 |
| Attention | 48 Gated DeltaNet linear + 16 full Gated Attention (3:1) |
| Native context | 262,144 tokens |
| Extended context | 1,000,000 (YaRN, hosted only) |
| License | Apache 2.0 (with patent grant) |
| Weights | 55.6 GB BF16 safetensors, 18 shards |
| Release | 14 August 2026 |
| VRAM floor | 24 GB (4-bit) / 80 GB-class (BF16) |
| Q4_K_M GGUF | ~17 GB |

Benchmarks (vendor-reported unless noted):

| Benchmark | Qwen3.8 27B | Comparison |
| --- | --- | --- |
| SWE-bench Pro | 61.7 | Claude Opus 4.6 Max 53.4 |
| DeepSWE 1.1 | 42.2 | Previous 27B 13.3 |
| LiveCodeBench v6 | 90.3 | — |
| Visual reasoning | 85.6 (CoT enabled) | vendor-reported |
| Visual math | 94.6 | vendor-reported |

Hosted pricing: OrcaRouter paid tier $0.33/M input, $2.40/M output; free rate-limited tier. Model page showed 33.3% error rate (7-day) and p50 first-token latency 225 ms.

AMD Day-0 throughput: 24.5 tokens/s (Ryzen AI Max+ 395), 51.8 tokens/s (Radeon AI PRO R9700).

## Why this source matters for the RAG

It is a detailed, critical review of a major open-weight model, pairing vendor claims with independent testing caveats and concrete pricing, hardware, and deployment guidance. It is valuable for RAG questions about open-weight model selection, local inference requirements, and the gap between benchmark claims and real-world performance.
