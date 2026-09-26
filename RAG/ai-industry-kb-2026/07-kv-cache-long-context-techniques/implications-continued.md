---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/implications-continued
title: "Implications (continued)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Groq", "Intel", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-01-22", "2026-02-19", "2026-03-16", "2026-03-21", "2026-03-25", "2026-04-15", "2026-04-24", "2026-05-05", "2026-05-28", "2026-06-01", "2026-06-12", "2026-06-13", "2026-08-13", "2026-08-26", "2026-08-29", "2026-09-17", "2026-09-22", "2026-10-20"]
keywords: ["attention", "benchmark", "benchmarks", "claude", "decode", "deepseek", "foundry", "fp8", "gemini", "glm", "gpt-5.6", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3658, 3692]
section: "7. KV Cache & Long-Context Techniques"
sha256: 94ecdd1281574f754b457f8fc22fa7d5ca6ab1e764b40e235fdb78e1dce12c11
---

# Implications (continued)

| Date | Event |
|---|---|
| 2025-12 | DeepSeek-V3.2 introduces DeepSeek Sparse Attention (learned lightning indexer, top-K 2048) |
| 2025-12 | NVIDIA acquires Groq technology/team (~$20B) — decode hardware enters NVIDIA's roadmap |
| 2026-01-22 | Inferact launches: $150M seed at $800M valuation to commercialize vLLM (a16z + Lightspeed) |
| 2026-03-21 | HF TGI archived read-only (maintenance mode) — MLA existed only on the Intel-Gaudi branch, never mainline CUDA |
| 2026 (v0.25) | PagedAttention removed from vLLM's new internal engine — legacy path only from here on |
| 2026-02 | SGLang blog: "Unlocking 25× inference performance on GB300 NVL72" [VENDOR headline; baseline unaudited] |
| 2026-02 | Qwen3.5-MoE ships: Gated-DeltaNet/Mamba + 256-expert MoE hybrid |
| 2026-02-19 | Gemini 3.1 Pro preview: 1M input / 64–66K output — 1M enters flagship-closed tier |
| 2026-03 | LMCache v0.3.15: engine-independent KV caching goes hardware-portable (H100/H200/B200/MI300X) |
| 2026-03-25 | Google Research blog announces TurboQuant (ICLR 2026): 3-bit KV, 6× memory, 8× attention speedup |
| 2026-03-16–17 | GTC 2026: Groq 3 LPU unveiled (150 TB/s SRAM decode; LPX rack 35×/MW claim [VENDOR]); Rubin CPX with GDDR7 for prefill |
| 2026-04 | vLLM FP8-KV benchmarks: 54% ITL slope vs BF16, break-even ~7K tokens, sub-0.3% accuracy |
| 2026-04 | FoveatedKV independent benchmark (M3 Max, Apr 4): importance-adaptive fp16/fp8/INT4 tiering |
| 2026-04-15 | Qwen3.6-35B-A3B (~): Gated-DeltaNet + Gated-Attention + MoE at **256K** — the 1M-tier exception |
| 2026-04-24 | DeepSeek V4 (Pro/Flash): 1M context, MIT license, CSA+HCA; SGLang day-zero support |
| 2026-05-05 | RadixArk formal launch: $100M seed (Accel-led, $400M post-money) to commercialize SGLang |
| 2026-05-28 | Claude Opus 4.8: 1M / 128K (secondary-sourced; Foundry caps at 200K on some tiers) |
| 2026-06-01 | MiniMax M3: 1M context open-weight launch (512K guaranteed floor) |
| 2026-06-12 | NVIDIA Nemotron-3-Ultra-550B-A55B: Mamba-2 hybrid MoE, MTP, 1M, NVFP4, OpenMDW-1.1 |
| 2026-06-13/17 | GLM-5.2 API release / open weights: IndexShare (2.9× FLOP cut at 1M), 1M/65K |
| 2026-07 | RadixArk expands Google TPU partnership (SGLang-JAX); SK Hynix $28B IPO filing reported |
| 2026-07 | SGLang day-zero support for Kimi K3 |
| 2026-06 | Zamba2-VL: Mamba2–Transformer hybrid VLM, ~10× TTFT cut (Jun 2026, MarkTechPost) — hybrid-side datapoint |
| 2026-08 | GPT-5.6 family: 1.05M context with pricing cut |
| 2026-08-13 | DeepSeek V4-Pro GA checkpoint (V4-Pro-0813) — secondary-sourced |
| 2026-08-26 | vLLM v0.28.0 tagged (see §6 for engine detail: sparse MLA, tiered KV offload to disk) |
| 2026-08-29 | Futurum Group: vLLM the "de facto open-source LLM inference engine" |
| 2026-09-17 | Lightbits RDMA-paged KV post: 8K→38.8 ms TTFT, 10.5M-token restore in seconds |
| 2026-09-22 | LMCache docs: 7.43× second-run CPU offload; llm-d p2p guide (−55.8%→−88.2%); this wave's FP8-KV default synthesis |
| 2026-10-20–21 | PyTorch Conference NA, San Jose (post-cutoff): vLLM across program tracks — program positioning only as of 2026-09-22 |

## Implications (continued)

