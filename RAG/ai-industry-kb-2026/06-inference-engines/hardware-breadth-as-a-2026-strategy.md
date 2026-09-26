---
id: ai-industry-kb-2026/06-inference-engines/hardware-breadth-as-a-2026-strategy
title: "Hardware breadth as a 2026 strategy"
domain: inference-engines
role: deep-dive
task: hardware
actors: ["AMD", "Alibaba", "Apple", "Broadcom", "DeepSeek", "Google", "Huawei", "Hugging Face", "Intel", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM", "xAI"]
dates: ["2025-08", "2025-10-29", "2026-01", "2026-01-22", "2026-02", "2026-04", "2026-05", "2026-05-05", "2026-06", "2026-07", "2026-09", "2026-09-20"]
keywords: ["amd", "apache", "ascend", "attention", "attribution", "benchmark", "consumer", "datacenter", "deepseek", "diffusion", "fine-tuning", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2707, 2751]
section: "6. Inference Engines"
sha256: d8ac32049a642a7beb656b9e9d74c9a33511f6ba35e7080a76be74d58880c6bb
---

# Hardware breadth as a 2026 strategy

### Hardware breadth as a 2026 strategy

- The v0.5.19 cookbook texture is the strategy made visible: Kimi-K3 on Ascend A3, Kimi-K2.7-Code-MXFP4 and Qwen3.5 MXFP4 on MI355X, MiniMax-H3 and Ling-3.0-flash on DGX Spark with consumer-GPU tuning guides, GLM-5.3 and PaddleOCR-VL deployment guides.
- HiCache host-memory tier appears in the Qwen3.5 MXFP4-on-MI355X recipe — the host-memory KV tier is reaching AMD consumer-adjacent hardware, not just datacenter NVIDIA.
- Diffusion hardware (NVIDIA, AMD, Intel XPU, Ascend NPU, Apple Silicon via MPS, Moore Threads MTT S5000) is the broadest hardware list of any 2026 diffusion-serving project, inherited from SGLang's multi-backend work rather than built per-modality.
- The July 2026 RadixArk–Google TPU partnership expansion (SGLang-JAX as the commercial vehicle) plus NVentures + AMD + MediaTek in the seed round read as deliberate hardware neutrality — the counterpart to vLLM's hardware breadth documented in 06a.

### SGLang-JAX — native TPU engine

- SGLang-JAX is a separate JAX/XLA inference engine for native Google TPU inference: continuous batching, prefix caching, tensor/expert parallelism, speculative decoding, and optimized TPU kernels.
- Date correction: announced 2025-10-29 — 2026 activity (RadixArk + Google partnership expansion, covered by TechCrunch in July 2026) is the expansion, not the birth.
- In 2026 SGLang's own documentation identifies SGLang-JAX as its TPU backend, and RadixArk's July 2026 Google TPU partnership expansion confirms it as the commercial vehicle.
- Dedup: TPU hardware specifics belong to §15; this part records only the engine-backend fact.

### SGLang day-zero model support

- Verified from the SGLang news log: day-zero support for DeepSeek-V4 (April 2026), Nemotron and Higgs Audio additions (June 2026), and day-zero support for Kimi K3 (July 2026).
- DeepSeek-V4 lineage note: SGLang shipped an implementation of DeepSeek Sparse Attention (V3.2, Dec 2025 — learned "lightning indexer" selecting top-K, commonly 2,048 tokens) including FP8 Q8KV8 sparse-MLA prefill (PR #30514).
- DSA nuance: sparse attention reduces attention computation; it does not intrinsically shrink the KV cache, because any historical token may be selected by the indexer later — the cache must retain full-fidelity KVs (DeepSeek V4 handles this separately via hybrid CSA+HCA: V4-Pro uses 27% of V3.2 inference FLOPs and 10% of KV-cache size at 1M tokens; V4-Flash pushes to 10% FLOPs / 7% KV [VENDOR figures via felloai]).
- See also the cross-engine race note under v0.5.19: time-to-serve is the 2026 differentiation metric.

### Claim audits: "2.5x cache hit rate", "400,000+ GPUs", "25x GB300"

- "RadixAttention cache hit rate up to 2.5x higher than the competition": [UNVERIFIED] as stated — no source in this research expresses RadixAttention's cache hit rate as a "2.5x vs competition" ratio.
- Likely origin of the confusion: a 2026-09-20 architecture review reports "up to 2.5x throughput gain under strict JSON constraints" for jump-forward decoding (bypassing autoregressive forward passes on grammar-constrained spans) — a throughput figure, not a cache-hit-rate figure.
- Nearest verifiable figures to cite instead: sgl-router 3.8x cache hit rate vs round-robin load balancing (baseline is round-robin, not "the competition's prefix caching"); cache-aware load balancer 1.9x throughput / 3.8x hit rate (v0.4 blog [VENDOR]); the SGLang paper's cache-aware longest-shared-prefix-first scheduler provably reaching the optimal cache hit rate and in practice hitting ~96% of it; Chatbot Arena deployment measured 52.4% (LLaVA-Next-34B) and 74.1% (Vicuna-33B).
- "Deployed on 400,000+ GPUs worldwide": [UNVERIFIED] — repeated verbatim across 2026 secondary roundups and community skill docs, likely originating from SGLang ecosystem marketing; no primary release-note citation found in either research wave.
- The figure was repeated again in RadixArk launch coverage ("hundreds of thousands of GPUs"); RadixArk's "trillions of tokens daily" is [ATTRIBUTION] to company launch materials (Business Wire), not independently audited — keep the same caveat.
- "Unlocking 25x Inference Performance with SGLang on NVIDIA GB300 NVL72": [VENDOR] — the SGLang project news log carries a real, dateable February 2026 article with that title; however, what "25×" means (throughput, per-GPU performance, or another metric) and against which baseline could not be audited from the headline — always quote with "per SGLang's own benchmark" attribution and a baseline caveat.

### RadixArk — SGLang's commercial spinoff

- Date correction: the $400M valuation was REPORTED in January 2026 (TechCrunch's 2026-01-22 Inferact piece: SGLang commercialization talks were "seeking a $400M valuation"); the FORMAL LAUNCH was 2026-05-05 — do not present January as the launch.
- Launch terms (2026-05-05, Business Wire): $100 million seed round at $400 million post-money, led by Accel, co-led by Spark Capital; participants NVentures (NVIDIA's venture arm), Salience Capital, A&E Investments, HOF Capital, Walden Catalyst Ventures, AMD, LDV Partners, WTT Investment, MediaTek.
- Angels: Igor Babuschkin (xAI co-founder), Lip-Bu Tan (Intel CEO), Hock Tan (Broadcom CEO), John Schulman (OpenAI co-founder, Thinking Machines Lab), Soumith Chintala (PyTorch creator), Olivier Pomel (Datadog co-founder), Thomas Wolf (Hugging Face co-founder), William Fedus (Periodic Labs), Robert Nishihara (Anyscale co-founder), Eric Zelikman (humans&), Logan Kilpatrick (Gemini product lead).
- Founders: Ying Sheng and Banghua Zhu — AI infrastructure veterans from xAI and NVIDIA; SGLang was created in 2023 by Sheng and collaborators.
- The company will "steward SGLang" — SGLang remains open source under Apache 2.0 — while building a commercial end-to-end platform covering training proprietary models, fine-tuning open models, reinforcement learning, and inference at scale.
- It also introduced Miles, its own large-scale reinforcement-learning framework (named in the 2026 RL stack alongside verl, AReaL, slime, and Tunix).
- Pre-launch precursor: RadixArk was already announced as a startup around August 2025 with an Accel-led round at ~$400M valuation reported — the May 2026 event is the formal, term-confirmed launch.
- Consolidation consequence: RadixArk is SGLang's commercial counterpart to vLLM's Red Hat/IBM commercial layer (vLLM's own January 2026 commercialization vehicle, Inferact — see 06a) — it narrows the "enterprise packaging gap" and reframes it as a two-commercial-stack comparison; as of September 2026 it is still early relative to Red Hat's shipped enterprise server, with paid-hosting tiers planned.
- The investor list (NVentures + AMD + MediaTek in one round) supports the "hardware-neutral" reading of SGLang's trajectory.
- January 2026 commercialization wave: on the same day (2026-01-22) the press carried both SGLang commercialization talks "seeking a $400M valuation" and Inferact's $150M-at-$800M launch for vLLM (a16z + Lightspeed, Simon Mo CEO) — the inference-duopoly pricing event; Inferact detail lives in 06a.

### MLA native support and the "4x batch size per GPU" question

