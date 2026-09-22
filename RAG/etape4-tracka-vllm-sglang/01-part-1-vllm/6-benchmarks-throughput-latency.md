---
id: etape4-tracka-vllm-sglang/01-part-1-vllm/6-benchmarks-throughput-latency
title: "6. Benchmarks (throughput/latency)"
domain: part-1-vllm
role: deep-dive
task: benchmark
actors: ["AMD", "Alibaba", "DeepSeek", "Google", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-03-06", "2026-05", "2026-07-27", "2026-08-21", "2026-08-26"]
keywords: ["benchmark", "benchmarks", "latency", "throughput", "agentic", "amd", "awq", "blackwell", "cost", "decode", "deepseek", "embedding"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [365, 425]
section: "PART 1 — vLLM"
sha256: 365d17d32930301ca9b3923c1c7b87e95625eb58e9a28fd91a3be4a4203b7633
---

# 6. Benchmarks (throughput/latency)

## 6. Benchmarks (throughput/latency)

### 6.1 Official / vendor-reported numbers
- **vLLM blog, 2026-07-27 (Kimi K3):** 118 tok/s without spec decode → **370 tok/s (3.14×) with DSpark**,
  16× NVIDIA GB300 NVL72. [official]
- **vLLM release notes (PR-reported, not independently verified):** Kimi K3 ~5% E2E latency (fused MXFP4
  top-k), 6.6–7.6× kernel speedup (Mamba metadata, one Triton launch), ~17 GiB/GPU saved (shared-expert
  sharding), ~60% better DSpark TTFT (adaptive speculative budget); DeepSeek-V4 routing kernel −2.94% E2E
  TPOT; ROCm W4A4 preshuffled asm GEMM +15% (Llama-3.3-70B MXFP4); B200 FP8 MoE tuning; Hopper low-latency
  GEMM also dispatched on SM100. [official — treat as vendor-reported]
- **NVIDIA GB300 / DeepSeek NVFP4** (third-party collected, unattributed vendor origin): DeepSeek-V3.2
  prefill-only 7,360 tok/GPU/s; DeepSeek-R1 (2× GPU) prefill 22,476 tok/GPU/s; "8–20× over Hopper with
  Blackwell + NVFP4". [secondary — vendor-reported, methodology not shown]
- **NVIDIA Blackwell RTX PRO 6000** (third-party collected): vLLM (NVFP4) 8,033 tok/s, TTFT 10.7 ms vs
  SGLang (GPTQ-INT4) 6,395 tok/s, 100% success at concurrency 128. [secondary]
- **Google Cloud, 2026-08-26 (vLLM TPU embedding):** Qwen3-Embedding-8B on TPU Ironwood (bf16, 16K+ seq,
  TP=4): 83,996 total tok/s, 5.13 req/s; cosine-similarity parity ≥0.999 text / ≥0.995 multimodal.
  [vendor-reported][secondary]

### 6.2 Independent measurements
- **SemiAnalysis InferenceX (Aug 2026)** [independent]:
  - DeepSeek V4 Pro 0813 (1.6T params / 49B active), agentic-coding replay (p50 88k in / p90 272k in tokens):
    MI355X **SGLang** matched B200 **vLLM** on perf/$ **until 2026-08-21**, when vLLM optimizations from
    Inferact and NVIDIA pushed B200 ahead — "a close race rather than a settled one". AMD published its own
    DeepSeek V4 vLLM optimization list (vllm-project/vllm#52911).
  - AMD MI355X Kimi-K2.5 MXFP4: one vLLM PR (#35850, merged 2026-03-06, shipped v0.18) moved 6.6 → **78.9
    tok/s/user** (8k/1k workload), 12.0× interactivity at low batch, 7.7× peak throughput, 15× at
    iso-throughput; peak **2,687 tok/s/GPU** (25 days from baseline to full effect).
  - Kimi K3: between 40–60 s E2E latency, MI355X **ATOM** (AMD vendor engine) beats even GB300 NVL72 vLLM on
    perf/$ — caveat: ATOM is a vendor engine, vLLM is the open comparison.
- **dstack DeepSeek-R1 benchmark (2026)** [independent]:
  - H200: TensorRT-LLM highest online throughput (4,176 tok/s); **vLLM led at concurrencies <128** in online
    throughput + E2E latency; offline: SGLang 6,311 tok/s.
  - MI300X: **vLLM outperformed SGLang in both online and offline throughput and E2E latency**; best online
    4,574 tok/s (vLLM); SGLang better only at concurrencies <32.
- **L40S head-to-head (Medium, Aug 2026)** [independent]: long-context workload — TensorRT-LLM ~2× throughput
  (75.9 vs ~40 tok/s) and 7.5× faster avg TTFT (8.6 s vs vLLM 63.8 s / SGLang 68.6 s); vLLM slight edge over
  SGLang in throughput (41.0 vs 38.2 tok/s) and TTFT; vLLM ITL 68.0 ms vs SGLang 72.6 ms. Quantization sweep
  on vLLM: FP8 vs AWQ vs GPTQ (figures only, no extracted numbers).
- **Single-L4 reproducibility study (inference-bench, GitHub)** [independent]:
  - Short regime (c=64, Qwen2.5-7B): vLLM AWQ **976 tok/s** > SGLang FP16 914 > vLLM FP16 831 > SGLang AWQ 506.
  - A100 sweep (vLLM v0.20.1 vs SGLang, May 2026, Qwen2.5-7B): vLLM FP16 3,102 tok/s @ c=64 vs SGLang 2,141;
    Marlin kernels +70% in v0.20.1 (177 vs 104 tok/s @ c=1 vs v0.8.5); SGLang Marlin collapsed at c=64
    (2,231 vs vLLM 4,762) — "opposite of the L4 pattern where SGLang wins at c=64".
- **ai-lab-benchmarks (Aug 2026, production model Qwen3.6-35B-A3B)** [independent]: NVFP4/vLLM vs GGUF/llama.cpp —
  4.7× throughput (47.1 vs 10.0 items/s), 3.4× faster prompt reading (13,916 vs 4,068 tok/s at 8k), generation
  201.5 vs 167 tok/s; translation quality identical (chrF++ 69.33 vs 69.34). Author noted SGLang setup friction
  (mem-fraction 0.80 vs vLLM 0.90; NVFP4 MoE runner backend flag) — "real setup cost".
- **Architecture-aware AMD study (arXiv 2603.10031, Feb 2026, MI325X ×8, vLLM v0.14.1)** [independent]:
  Llama-3.1-405B 15,944 tok/s vs DeepSeek V3.2 15,343 tok/s peak (text-only); Qwen3-VL-235B 47,873 tok/s
  (vision, incl. image tokens); Kimi-K2.5 7,327 tok/s; AITER required for competitive MLA (+3–5% at high
  concurrency on Llama-405B, MoE/MLA speedups larger).
- ⚠️ Legacy A100 numbers still circulating (17.12 req/s OPT-13B vs HF TGI 0.71 req/s) are from the
  original vLLM paper era — illustrative of PagedAttention's original gains, not 2026 hardware. [secondary]
- ⚠️ Viral "29% gap" claims (SGLang ~16,200 vs vLLM ~12,500 tok/s on H100 Llama-3.1-8B): the primary source
  (RunPod) attributes the gap almost entirely to **prefix-heavy** traffic (RadixAttention prefix reuse);
  on unique prompts the engines are "within a few percent". Do not quote the 29% without the workload
  qualifier. [secondary]

---

