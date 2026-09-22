---
id: etape4-tracka-vllm-sglang/03-part-3-synthesis-vllm-vs-sglang-2026-state/overview
title: "PART 3 — SYNTHESIS: vLLM vs SGLang (2026 state)"
domain: part-3-synthesis-vllm-vs-sglang-2026-state
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Alibaba", "Cohere", "Crusoe", "DeepSeek", "Google", "Hugging Face", "Intel", "Lambda", "Meta", "Mistral", "Moonshot", "Nebius", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM", "xAI"]
dates: ["2023-02-09", "2024-01-08", "2026-01-01", "2026-01-09", "2026-01-16", "2026-01-20", "2026-01-21", "2026-01-23", "2026-02-19", "2026-02-24", "2026-02-25", "2026-04-06", "2026-04-27", "2026-05", "2026-05-05", "2026-05-16", "2026-06-13", "2026-07-10", "2026-07-25", "2026-07-27", "2026-08-08", "2026-08-21", "2026-08-22", "2026-08-26", "2026-09-01", "2026-09-05", "2026-09-09", "2026-09-18", "2026-09-22"]
keywords: ["sglang", "vllm", "amd", "apache", "attention", "attribution", "aws", "cohere", "compute", "decode", "deepseek", "diffusion"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [1039, 1113]
section: "PART 3 — SYNTHESIS: vLLM vs SGLang (2026 state)"
sha256: ae0ee1aff23577064376e3a7e0202e9ca840dd1493948fc12ca6bd53cc49c9ef
---

# PART 3 — SYNTHESIS: vLLM vs SGLang (2026 state)

**Synthesized 2026-09-22 from the two research passes above.**

## 3.1 Snapshot comparison (2026-09-22)

| Dimension | vLLM | SGLang |
|---|---|---|
| Repo | `vllm-project/vllm` (created 2023-02-09) | `sgl-project/sglang` (created 2024-01-08) |
| License | Apache-2.0, no CLA | Apache-2.0 |
| GitHub stars / forks | **92,444 / 22,525** [official] | **36,323 / 9,061** [official] |
| Latest 2026 release | **v0.30.0** (2026-09-22), bi-weekly cadence | **v0.5.20** (2026-09-18), ~2–4 week cadence |
| Core innovation | PagedAttention (block-paged KV) | RadixAttention (radix-tree prefix KV) |
| Engine architecture | V1 engine (V0 removed since v0.11); **Model Runner V2 default since v0.29.0** (MRV1 deprecated, removal target v0.32) [official] | Zero-overhead batch scheduler (v0.4); piecewise/breakable CUDA graphs default (v0.5.10); **unified radix tree default since v0.5.19** [official] |
| Origin | UC Berkeley Sky Computing Lab + LMSYS (SOSP 2023) | UC Berkeley (Ion Stoica's lab) 2023, hosted by LMSYS non-profit |
| Commercial backers | Red Hat AI Inference Server (vLLM + Neural Magic); nm-vllm with SLAs; VMware AI Factory runtime [vendor-reported] | **RadixArk** spin-out (Aug 2025; TechCrunch Jan 2026: ~$400M valuation, Accel-led; $100M seed May 2026 single-sourced [unverified]); paid hosting [secondary] |
| Sponsors | a16z, Dropbox, Sequoia, Skywork AI, ZhenFund (cash); compute from Alibaba Cloud, AMD, Anyscale, AWS, Crusoe, Databricks, DeepInfra, Google Cloud, Intel, Lambda, Nebius, NVIDIA, Replicate, Roblox, RunPod (README list — verify canonical) [official] | a16z Open Source AI Grant (2025), PyTorch Ecosystem member (Mar 2025) [official] |
| Governance | Community-led, no foundation umbrella (third-party assessment Jul 2026) [secondary]; vLLM Production Stack Helm/operator; llm-d (Red Hat/Google/IBM/NVIDIA) builds on vLLM [secondary] | LMSYS non-profit hosts the project; formal security process weaker — Aug-2026 CERT/CC disclosure of 6 vulns; no SECURITY.md at that time [secondary][unverified fix status] |
| Production users (project- or third-party-stated) | Meta, LinkedIn, Mistral, Hugging Face, Uber (third-party claims, unconfirmed) [unverified]; Tesla, Cohere via llm-d [secondary] | xAI (community deck), Cursor, LinkedIn, clouds/GPU providers (README list — production depth unverified) [unverified] |
| Scale claims (self-reported) | "De-facto open standard" characterization [secondary] | "Trillions of tokens/day", "400,000+ GPUs worldwide", "de facto industry standard" [official][unverified] |

## 3.2 Performance picture (independent + vendor, qualified)

- **Prefix-heavy high-concurrency (H100 80GB, Llama 3.1 8B, PremAI 2026):** SGLang ~16,200 tok/s vs
  vLLM ~12,500 tok/s — SGLang +29%, but RunPod's primary attribution is **prefix-heavy traffic only**;
  on unique prompts the engines are "within a few percent" [secondary].
- **Spheron (Llama 3.3 70B FP8, H100):** SGLang +29% total throughput, +117% output-token throughput,
  TTFT −23%, ITL −15%; gap shrinks to +2–5% with unique prompts [secondary].
- **Single unique prompt (community, ~2025-06):** vLLM 60.0 vs SGLang 52.7 tok/s (vLLM ~1.1×) [secondary, dated].
- **DeepSeek family:** SGLang claims 3.1× on DeepSeek-V3 via optimized MLA backends [secondary]; both
  engines had **day-0 support for DeepSeek-V4** (vLLM v0.28-ish stack; SGLang v0.5.12, 2026-05-16) [official].
- **SemiAnalysis InferenceX (Aug 2026) — DeepSeek V4 Pro 0813:** MI355X **SGLang** matched B200 **vLLM**
  on perf/$ until 2026-08-21, when vLLM optimizations (Inferact, NVIDIA) pushed B200 ahead — "a close
  race, not a settled one" [independent].
- **dstack DeepSeek-R1 (H200):** vLLM led at concurrency <128 (online throughput + E2E latency); SGLang
  6,311 tok/s offline. **MI300X:** vLLM beat SGLang in online and offline throughput and E2E latency [independent].
- **Long-context L40S:** TensorRT-LLM ~2× throughput and 7.5× faster TTFT than both; vLLM slight edge
  over SGLang (41.0 vs 38.2 tok/s) [independent].
- **Official GB300 claims:** vLLM blog — Kimi K3 118 → **370 tok/s (3.14×)** with DSpark on 16× GB300 NVL72
  [official]. SGLang lmsys blog (2026-02-19) — "Unlocking 25x Inference Performance on GB300 NVL72" [official].
- ⚠️ "29% gap" numbers must always be quoted **with the prefix-heavy workload qualifier**.

## 3.3 2026 combined timeline (both engines)

| Date | vLLM | SGLang |
|---|---|---|
| 2026-01-01 | — | v0.5.7 (day-0 Mimo-V2-Flash) [official] |
| 2026-01-09 | — | SGLang Model Gateway v0.3.1: 10–12× faster cache-aware routing [official; date secondary] |
| 2026-01-16 | — | lmsys blog: SGLang-Diffusion 2.5× faster than Nov-2025 release [official] |
| 2026-01-20 | **v0.14.0** — first 2026 release [official] | — |
| 2026-01-21 | — | TechCrunch: SGLang spins out as **RadixArk** (~$400M valuation, Accel) [secondary] |
| 2026-01-23 | — | v0.5.8 (diffusion speedups up to 1.5×) [official] |
| 2026-02-19 | — | lmsys blog: 25× inference perf on GB300 NVL72 [official] |
| 2026-02-24 | — | v0.5.9 (LoRA load/compute overlap, −78% TTFT) [official] |
| 2026-02-25 | v0.16.0 — V1 confirmed as only engine [secondary] | — |
| 2026-04-06 | — | v0.5.10 (piecewise CUDA graphs default, Elastic NIXL-EP, HiSparse) [official] |
| 2026-04-27 | v0.20.0 [official] | — |
| 2026-05-05 | — | v0.5.11 (CUDA 13 + Torch 2.11 default) [official] |
| 2026-05-16 | — | **v0.5.12 — DeepSeek-V4 day-0** [official] |
| 2026-06-13 | — | v0.5.13 (Nemotron 3 Ultra day-0) [official] |
| 2026-07-10 | v0.25.0 [official] | v0.5.15 (GLM-5.2 NVFP4 production tuning) [official] |
| 2026-07-25 | — | v0.5.16 (DSpark confidence-driven spec decode) [official] |
| 2026-07-27 | **v0.26.0** — Inkling family; DeepSeek-V4 perf push; vLLM Kimi-K3 blog: 370 tok/s w/ DSpark (3.14×), 16× GB300 NVL72 [official] | lmsys blog: RadixArk + Google bring full SGLang to TPUs [official] |
| 2026-08-08 | — | **v0.5.17 — Kimi K3 day-0** [official] |
| 2026-08-21 | vLLM optimizations push B200 ahead of MI355X-SGLang on DeepSeek V4 perf/$ (SemiAnalysis) [independent] | — |
| 2026-08-22 | — | v0.5.18 (710 PRs) [official; highlights not pulled] |
| 2026-08-26 | **v0.28.0** — Kimi-K3 perf push, DeepSeek V4 sparse MLA, KV disk offload, bitsandbytes→out-of-tree [breaking] [official] | — |
| 2026-08-26 | Google Cloud: vLLM TPU embedding pipelines GA on Ironwood [vendor-reported] | — |
| 2026-08 | — | CERT/CC-coordinated disclosure of 6 SGLang vulnerabilities [secondary] |
| 2026-09-01 | VMware AI Factory ships vLLM-based runtime [secondary] | — |
| 2026-09-05 | — | **v0.5.19** — beam search, DeepEP v2, unified radix tree default, AMD Lean attention [official] |
| 2026-09-09 | **v0.29.0** — Model Runner V2 default, 10 archs removed, RL weight sync [breaking] [official] | — |
| 2026-09-18 | — | **v0.5.20** — RL sampling masks, SGLang Simulator, CUDA 12 retired, XPU/MUSA images [official] |
| 2026-09-22 | **v0.30.0** — Fast Start IPC weight cache, Gumbel-max watermarking, HiSparse, DeepSeek-V4.1-Flash [breaking] [official] | — |

