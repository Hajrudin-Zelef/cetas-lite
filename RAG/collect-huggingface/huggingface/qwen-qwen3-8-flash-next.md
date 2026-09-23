---
id: collect-huggingface/huggingface/qwen-qwen3-8-flash-next
title: "Qwen3.8-Flash-Next - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["qwen", "agentic", "agents", "attention", "benchmarks", "claude", "compute", "deepseek", "embedding", "embeddings", "latency", "license"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3.8-Flash-Next.md
source_anchor: ""
source_lines: [1, 53]
sha256: 34313ba81fe0adfe746ca4f61fc069732ddf0e8580113419ebad382a294f2270
---

# Qwen3.8-Flash-Next - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3.8-Flash-Next
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3.8-Flash-Next is an experimental preview from Alibaba Qwen of the architecture that will underpin Qwen4, rethinking how core LLM components interact at scale. It is a causal language model with a vision encoder (Image-Text-to-Text): 125B parameters with 6B activated, plus 51B n-gram embedding and 4B MTP; hidden dimension 2560; 48 layers; layout 12 x (3 x (Gated DeltaNet -> MoE) -> 1 x (Qwen Sparse Attention -> MoE)). Key innovations: (1) Hybrid Attention with QSA — Gated DeltaNet + Qwen Sparse Attention operating at the micro-block level (budget 512 blocks or 2048 tokens), cutting long-context latency; (2) Gated Residual — element-wise read gate and per-branch scalar write gate on widened residual streams (4 branches, bottleneck rank 320); (3) N-gram Embedding — 20,000,000 bigram/trigram entries indexed at layer 2 for parameter scaling without MoE compute; (4) Tailored training recipe — Muon and AdamW per weight category, refitted scaling laws, no batch-size warmup. MoE: 512 experts, 10 routed + 1 shared active, expert intermediate 640. Context: 262,144 native, extensible to 1M (YaRN static scaling with factor 4.0, mrope_section 11/11/10, rope_theta 1e7, partial_rotary_factor 0.25).

Benchmarks (best in bold, mostly beating Qwen3.8-27B, Qwen3.7-Plus, DeepSeek-V4-Flash-0731, Claude-Opus-4.6): DeepSWE 1.1 58.7, SWE-bench Pro 62.5, SWE-bench Multilingual 81.0, CoWorkBench 73.9, JobBench 55.7, Agents' Last Exam 51.2, Toolathlon Verified 73.5, IFBench 81.3, GPQA Diamond 91.7, HLE 35.9, LiveCodeBench v6 91.9. Vision-language: ClawEval-MM pass@3 64.4, RecreationBench 49.9, AndroidWorld 84.5, OSWorld 2.0 binary 19.4/partial 52.3, Vision2Web 64.0, ERQA 72.3, LVBench 76.6, RealWorldQA 88.5, MathVision (w/o CI) 90.6, CharXiv RQ 84.6.

Usage: thinking mode by default (enable_thinking/preserve_thinking/reasoning_effort with xhigh/medium/low). Thinking sampling temp 1.0, top_p 0.95, top_k 20; non-thinking temp 0.7, top_p 0.8, presence_penalty 1.5. License: qwen-community-1.0. Hub: 180B params BF16/I64; 807,550 monthly downloads; 269 quantizations. Deploy via SGLang, vLLM, TokenSpeed, Transformers; official Qwen3.8-Flash (1M default) on Qwen Cloud.

## Key points

- Experimental preview of the Qwen4 architecture, first open-weight release.
- 125B total / 6B activated + 51B n-gram embedding + 4B MTP.
- Hybrid attention: Gated DeltaNet + Qwen Sparse Attention (micro-block level).
- Gated Residual streams and N-gram Embedding for efficient parameter scaling.
- Muon/AdamW tailored training recipe with refitted scaling laws.
- 262K native context, extensible to 1M via YaRN.
- Beats frontier baselines on most language and vision-language benchmarks.
- qwen-community-1.0 license; thinking mode default.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | Qwen |
| Model name | Qwen3.8-Flash-Next |
| Architecture | Hybrid (Gated DeltaNet + QSA) causal LM + vision encoder |
| Total params | 125B (6B activated) + 51B n-gram emb + 4B MTP |
| Hidden size | 2560 |
| Layers | 48 |
| Experts | 512 (10 routed + 1 shared) |
| N-gram embedding | 20,000,000 (bigrams/trigrams, layer 2) |
| Context length | 262,144 native; up to 1M (YaRN) |
| License | qwen-community-1.0 |
| Weight formats | BF16, I64 |
| Benchmarks | SWE-bench Pro 62.5, GPQA-D 91.7, LiveCodeBench 91.9, AndroidWorld 84.5, LVBench 76.6 |
| Sampling | Thinking: temp 1.0 / top_p 0.95; non-thinking: temp 0.7 / top_p 0.8 |
| Deployment | SGLang, vLLM, TokenSpeed, Transformers |
| Downloads/month | 807,550 |

## Why this source matters for the RAG

This card documents a radically new hybrid architecture (sparse attention + gated delta net + n-gram embeddings) that delivers frontier agentic performance at just 6B active parameters, making it central reference material for retrieval on architectural innovation, long-context efficiency, and efficient frontier models.
