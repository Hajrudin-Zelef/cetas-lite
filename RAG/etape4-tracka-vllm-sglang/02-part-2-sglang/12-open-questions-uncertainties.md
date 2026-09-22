---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/12-open-questions-uncertainties
title: "12. Open questions / uncertainties"
domain: part-2-sglang
role: deep-dive
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-01-01", "2026-01-09", "2026-01-16", "2026-01-21", "2026-01-23", "2026-02-19", "2026-02-24", "2026-03-28", "2026-04-06", "2026-05-05", "2026-05-16", "2026-06-13", "2026-06-26", "2026-07-10", "2026-07-25", "2026-08-08", "2026-08-22", "2026-09-05", "2026-09-18", "2026-09-22"]
keywords: ["agentic", "attention", "awq", "benchmarks", "compute", "consumer", "decode", "deepseek", "diffusion", "disclosure", "fp8", "glm"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [976, 1038]
section: "PART 2 — SGLang"
sha256: d028e9970bb404fd59c694308852489fdb4bfe38cff06fbb26ebbf8027015523
---

# 12. Open questions / uncertainties

## 12. Open questions / uncertainties

1. **v0.5.18 highlights** — release-notes body not pulled; only size (710 PRs / 212 contributors) confirmed [official].
2. **v0.5.17 Kimi-K3 details** — release body truncated in API fetch; "2.8T-parameter multimodal LatentMoE (896 experts, top-16…)" is a partial quote.
3. **Gateway-v0.3.1 date** (2026-01-09) is [secondary]-sourced only.
4. **RadixArk $100M seed figure** — single-sourced summary; only the ~$400M valuation is TechCrunch-sourced.
5. **Security fixes** — six Aug-2026 CERT/CC-coordinated issues; fix/landed-version status as of 2026-09-22 not verified.
6. **No Artificial Analysis, ServeTheHome, or SiliconANGLE** SGLang benchmarks/articles were found in this pass.
7. **Adoption/GPU/token counts** ("400,000 GPUs", "trillions of tokens/day") are project self-reports — no independent audit found.
8. **SM120 INT4 correctness issue (#21132)** — community report predates v0.5.16 NVFP4 fixes; current status not verified.
9. **vLLM star/fork counts** and **SGLang enterprise pricing/SLAs** were not pulled (out of scope for this pass).

---

## 13. 2026 news & blog timeline (official unless noted)

- **2026-01-01** — v0.5.7 released (day-0 Mimo-V2-Flash) [official].
- **2026-01-09** — SGLang Model Gateway v0.3.1: 10–12× faster cache-aware routing, 99% memory reduction [official; date [secondary]].
- **2026-01-16** — lmsys.org blog: "SGLang-Diffusion: Two Months In" — up to 2.5× faster than its Nov-2025 initial release [official].
- **2026-01-21** — TechCrunch: "Project SGLang spins out as RadixArk with $400M valuation" (round led by Accel; announced Aug 2025; founders Ying Sheng + Banghua Zhu) [secondary].
- **2026-01-23** — v0.5.8: diffusion speedups up to 1.5× [official].
- **2026-02-19** — lmsys.org blog: "Unlocking 25x Inference Performance with SGLang on NVIDIA GB300 NVL72" [official].
- **2026-02-24** — v0.5.9: LoRA weight-load/compute overlap (−78% TTFT) [official].
- **2026-03-28** — v0.5.10rc0 prerelease [official].
- **2026-04** — lmsys.org blog: "DeepSeek-V4 on Day 0: From Fast Inference to Verified RL with SGLang and Miles" [official].
- **2026-04-06** — v0.5.10: piecewise CUDA graphs default, Elastic NIXL-EP, HiSparse, FlashInfer MXFP8, transformers 5.3.0 [official].
- **2026-05-05** — v0.5.11: CUDA 13 + Torch 2.11 default [official].
- **2026-05-16** — v0.5.12: DeepSeek-V4 day-0 support [official].
- **2026-06** — lmsys.org blog: "The next generation of speculative decoding: DFlash and Spec V2" [official].
- **2026-06** — lmsys.org blog: day-0 support for Nemotron 3 Ultra / Nemotron 3 Super / Higgs Audio v3 TTS [official].
- **2026-06-13** — v0.5.13: Nemotron 3 Ultra day-0 [official].
- **2026-06-26** — v0.5.14: GLM-5.2, LFM2.5, Kimi-K2.7-Code [official].
- **2026-07** — lmsys.org blog: "Serving GLM5.2 NVFP4 agentic workloads with SGLang: Reaching 500 TPS in two weeks" [official].
- **2026-07** — lmsys.org blog: "SGLang and Miles add day-0 support for Kimi K3" [official].
- **2026-07** — lmsys.org blog: "RadixArk and Google bring full SGLang features to TPUs" [official].
- **2026-07-10** — v0.5.15: GLM-5.2 NVFP4 production tuning [official].
- **2026-07-25** — v0.5.16: DSpark speculative decoding [official].
- **2026-08** — SGLang provides day-0 support for OpenAI gpt-oss (noted 2025-08 in README news; gpt-oss PD-decode prefix reuse for SWA hybrids landed v0.5.19) [official].
- **2026-08** — CERT/CC-coordinated disclosure of six SGLang vulnerabilities; mitigations published by community (pickle IPC, dumper port, API key, network isolation) [secondary].
- **2026-08-08** — v0.5.17: Kimi K3 day-0 [official].
- **2026-08-22** — v0.5.18 [official].
- **2026-09-05** — v0.5.19: beam search, DeepEP v2, unified radix tree default [official].
- **2026-09-18** — v0.5.20: RL sampling masks, SGLang Simulator, CUDA 12 retired [official].
- **2026-09** — SGLang-Omni v0.1.6 on PyPI; day-0 AuK/AuK-Flash audio models [official].

## 14. Practical deployment notes (from docs & community, 2026)

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

