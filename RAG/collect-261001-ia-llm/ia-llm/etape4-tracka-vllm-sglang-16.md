---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-16
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Cohere", "Crusoe", "DeepSeek", "Google", "Hugging Face", "Intel", "Lambda", "Meta", "MiniMax", "Mistral", "Moonshot", "Nebius", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM", "xAI"]
dates: ["2023-02-09", "2024-01-08", "2026-02-19", "2026-05", "2026-05-16", "2026-08-21", "2026-09-18", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "amd", "apache", "attention", "attribution", "awq", "aws", "cohere", "compute", "consumer"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [1023, 1082]
sha256: 0a234a1439dff20c55835472229643d1004b0c3c87b3cc0504c624cf3064ee32
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

- **Install:** `pip install "sglang[srt]"` (PyPI; `sglang` package; Python ≥ 3.10; CUDA 13.x driver needed for recent releases); Docker `lmsysorg/sglang`; from-source for unreleased features [official/secondary].
- **Launch:** `python -m sglang.launch_server --model <hf-id> --tp 2` or `sglang serve`; default port 30000; OpenAI-compatible endpoints [official/secondary].
- **Key server flags (2026):** `--enable-hierarchical-cache` (HiCache), `--moe-a2a-backend {deepep,deepep_v2,mooncake,mori}`, `--moe-runner-backend {triton,flashinfer_cutlass,flashinfer_trtllm,flashinfer_megamoe,cutlass}`, `--speculative-algorithm {EAGLE,DFLASH,DSPARK,NEXTN,NGRAM,STANDALONE}`, `--quantization {fp8,awq,gptq_marlin,...}`, `--kv-cache-dtype fp8_e4m3`, `--enable-layernorm-sp`, `--enable-lean-attention`, `--sampling-mask-max-tokens`, `--enable-response-store`, `--reasoning-parser`, `--tool-call-parser`, `--enable-metrics` (Prometheus `sglang:*`) [official/secondary].
- **Env knobs seen in the wild:** `SGLANG_ALLOW_OVERWRITE_LONGER_CONTEXT_LEN=1`, `SGLANG_DISABLE_CUDNN_CHECK=1`, `VLLM_WORKER_MULTIPROC_METHOD=spawn`, `SGLANG_USE_PICKLE_IPC=false` (security), `SGLANG_DISABLE_LEAN_ATTENTION=1`, `SGLANG_OPT_KDA_FUSED_ACCEPT_STATE=1` [secondary].
- **Multi-node:** PD disaggregation topologies (prefill TP4 + decode DEP4 cited with ~5× TPS/GPU on Qwen3.5 [official]); NIXL/Mooncake RDMA backends; per-scheduler load sockets feed the Model Gateway/router [official].
- **Observability:** `/server_info`, `/metrics` (Prometheus), per-token weight-version spans in generation metadata (v0.5.19), e2e latency metadata in Rust server [official/secondary].
- **Known operational gotchas (community, 2026):** TP model-loading barrier hardcodes a 480 s timeout (`UNBALANCED_MODEL_LOADING_TIMEOUT_S`) — large quantized checkpoints on slow storage can die mid-boot (community patch proposes `SGLANG_UNBALANCED_MODEL_LOADING_TIMEOUT_S` env override) [secondary](https://github.com/mattbucci/2x-3090-ga102-300-a1-sglang-inference/blob/HEAD/scripts/upstream-pr/049-load-timeout-env-PR.md); `/get_server_info` deprecated in favor of `/server_info` (v0.5.10) [official]; `reasoning_tokens` was zero before v0.5.10 [official].
- **Consumer-GPU recipes (cookbooks):** Qwen3.8-27B on RTX 5090 / RTX PRO 6000 / DGX Spark (re-measured v0.5.19); MiniMax-H3 on 24 GB GPUs; Ling-3.0-flash on DGX Spark; Qwen3.5 MXFP4 on MI355X with FP8 KV cache or HiCache host tier [official].

---

*End of report. Research conducted 2026-09-22. All dated claims carry provenance tags; see §12 for open questions.*


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

