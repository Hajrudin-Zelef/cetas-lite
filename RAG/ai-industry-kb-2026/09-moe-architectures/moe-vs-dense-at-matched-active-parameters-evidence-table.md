---
id: ai-industry-kb-2026/09-moe-architectures/moe-vs-dense-at-matched-active-parameters-evidence-table
title: "MoE vs dense at matched active parameters: evidence table"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "China", "DeepSeek", "Huawei", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-04-24", "2026-05", "2026-06", "2026-07-27"]
keywords: ["moe", "parameters", "amd", "ascend", "attention", "awq", "benchmark", "benchmarks", "compute", "consumer", "cost", "decode"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5052, 5079]
section: "9. MoE Architectures"
sha256: 3852d2a3bfb91c18ef9e2f381c32eaa131c405c788ff47c8c34ac00ad3b764a4
---

# MoE vs dense at matched active parameters: evidence table

- **Expert parallelism (EP):** distributes experts across GPUs; tokens routed via all-to-all dispatch/combine to the GPU holding their experts. vLLM: `--enable-expert-parallel`, often combined `--tensor-parallel-size 8 --enable-expert-parallel` (DeepSeek V3.2 on 8× H200); DP+EP at high concurrency. SGLang: EP with DeepEP integration; day-0 model-support culture (MiMo-V2-Flash inference code contributed upstream). NVIDIA Dynamo: ships vLLM/SGLang/TensorRT-LLM runtimes with MiniMax-M3 support, KV-aware routing, disaggregated prefill/decode for 1M-context MoE. **Moonshot open-sourced MoonEP (2026-07)** specifically to attack cross-node EP transfer bottlenecks.
- **When EP helps vs hurts — AMD ROCm guidance (verified):** ≤128 concurrent requests → TP=8 gives 40–86% higher throughput; ≥512 → DP=8+EP gives 16–47% higher (7,114 TPS for DeepSeek-R1 at 1024 concurrency); crossover ~256–512. MLA/MQA models (DeepSeek-V2/V3/R1, Kimi-K2.5) → DP+EP mandatory to avoid KV duplication across TP ranks. Ultra-sparse MoE (<1% activation density, e.g., Llama-4-Maverick-class) → EP is 7–12% SLOWER (all-to-all overhead exceeds sparse-compute savings); standard MoE (≥3% density) → EP helps.
- **EPLB — Expert Parallelism Load Balancer:** vLLM-Ascend ships dynamic/static EPLB (record expert-traffic maps, rebalance without stop-the-world) protecting TTFT/TPOT; verified on DeepSeek-V3.1/R1; W8A8, W4A8, MXFP4/MXFP8 quant types supported on Ascend 950.
- **Communication-scale math (reference):** training-side EP dispatch fundamentally requires two any-to-any phases per MoE layer (dispatch token activations to expert ranks, combine outputs back — arXiv 2503.04398); more active experts per token (k) and wider representations (D) scale traffic linearly: per-layer all-to-all ≈ 2 × tokens × k × hidden-dim × (G−1)/G, TB-scale per layer at 64-way EP for V3-class models — the dominant training cost alongside optimizer memory. Inference dispatch per token ≈ k × hidden-dim × 2 (dispatch + combine), amortized across batch. DeepSpeed's communication notes identify MoE all-reduce penalties of **3–12 seconds in problematic large-scale layouts**, fixed via multi-rank bucketing and improved rank placement; the llm-d project documents that sparse dispatch/combine can run over InfiniBand/RoCE while avoiding KV replication — a property of the communication backend (DeepEP over NVSHMEM with GPU-initiated RDMA), not of fine granularity itself.
- **2026 training-side EP maturity:** Axolotl (June 2026) added EP for distributed MoE training via DeepEP, moving DeepEP from the pretraining world into fine-tuning YAML configs; torchtune (May 2026 paper) ships a custom expert-parallel plan on PyTorch DTensor alongside loss parallel (sharding output features over vocab so full logits never materialize) and Ring-Attention context parallel; DeepSpeed's MoE tutorial refresh (~Sept 2026) supports EP combined with data parallelism, ZeRO, model parallelism, and ZeRO-Offload, with layer-specific expert counts and expert-parallel degree.
- **Quantization recipes by tier (2026 practice):** 300B+ flagship serving → FP8 (GLM-5.3-Flash ships FP8 by default, ~306–331 GiB; MiMo-V2.6-Flash weights 172.9 GB FP8 across 65 shards) or NVFP4 (MiniMax M3, ~245 GB → ~61 GB/GPU at TP=4; vLLM v0.15.0 added Marlin/NVFP4-CUTLASS/FP8/INT8 kernels for non-gated MoE). Local/workstation → AWQ/Marlin 4-bit (documented Mixtral 8x7B vLLM case: 90 GB → 26 GB VRAM, 2.1× throughput, −75% TTFT) or MXFP4 (gpt-oss native; Kimi K3 MXFP4 weights with E8M0 scale per 32 weights / MXFP8 activations). Consumer GPU → expert offloading (Gemma 4 26B-A4B hot set ~4.5 GB, 128 routed experts ~11 GB in Q4_K_M on NVMe, ~24 MB loaded per forward pass, servable on 8 GB RTX 3070). CPU/offload → GGUF (Qwen3.6 20.9 GB Q4 on MacBook Pro; Unsloth UD-Q4_K_XL GGUF ~21 GB on a single 24 GB RTX 4090; `save_pretrained_gguf` one-call API with 23-entry quant list and save-time OOM guard `maximum_memory_usage`, default 0.75 — wave2.1 delta, Sept 2026).
- **Speculative decoding via MTP:** training with multi-token-prediction heads gives free inference speedup — MiMo-V2-Flash 2.0–2.6× via self-speculative decoding (accepted length 2.8–3.6 tokens/forward); MiMo-V2.6-Flash 5-layer DFlash-style drafter (7 tokens ahead); DeepSeek-V3 lineage and Qwen3.6 also train MTP heads; Unsloth Studio v0.1.805/806-beta (early Sept 2026) enabled MTP-by-default for Qwen3.8-Flash and GLM-5.3-Flash, claiming up to 2x faster generation [VENDOR-REPORTED]. Asymmetry noted: community ROCm measurements found MTP near-useless on MoE at some configs while helping dense 27B from 45→62 t/s — MoE's sparse compute profile saturates differently.

### MoE vs dense at matched active parameters: evidence table

| Experiment | Scale | Result |
|---|---|---|
| Controlled from-scratch (OliverSundaram/MoE-Study, 2026) | ~150M active, matched data/compute/hardware | Dense beat MoE on almost every benchmark, 3× faster inference — routing overhead dominates at small scale |
| Pre-registered ablation (2026-06) | matched compute | MoE perplexity 5.72 vs dense 22.66 (4.0×), gap widening with training — sparse capacity compounds |
| OLMoE-1B-7B vs OLMo-1B (same family) | ~1B active | Active params agree within 8.9%; resident memory differs 5.9× — active ≠ total |
| Nemotron 3.5 Lightning (30B/3B MoE+Mamba-2) vs Gemma 4 31B dense (Artificial Analysis, 2026-08) | ~30B total | 235–494 vs 37–222 t/s; $0.22/M vs $0.40/M output; capability 24 vs 30 — speed/$ win, capability cost |
| Qwen3.6-35B-A3B (3B active) vs 27B dense, same RTX 5070 (practitioner) | workstation | 61 vs 7 t/s — the active-compute illustration |
| Qwen3.6 (35B/3B) vs Gemma 4 (25.2B/3.8B) (community benchmarks) | similar active compute | 73.4% vs 17.4% SWE-bench Verified — post-training decides, architecture enables |

### Non-NVIDIA serving hardware (dated)

- **2026-04-24:** DeepSeek-V4 launches on Huawei Ascend chips first (no NVIDIA required) — the first major open MoE debuting on non-NVIDIA hardware; vLLM-Ascend ships dynamic/static EPLB verified on DeepSeek-V3.1/R1, with W8A8/W4A8/MXFP4/MXFP8 quant types supported on Ascend 950.
- **2026-07-27:** Kimi K3 runs on AMD GPUs — MI355X serving benchmarks exist; Moonshot open-sources MoonEP for cross-node EP transfers.
- **2026-08:** GLM-5.3-Flash's stealth week ("Ox Alpha") served entirely on domestically produced Chinese AI chips while topping OpenRouter coding charts.
- **2026-09:** Unsloth Studio betas make AMD Vulkan-by-default (+20% prefill, +23% prompt processing, +8% generation on Strix Halo [VENDOR]) and add RDNA1/2 support; Strix iGPU BIOS guidance claims 3x faster inference with more iGPU VRAM [VENDOR]; third-party unsloth-cli docs report NVFP4/AWQ exports live-tested for *serving* (not training) via vLLM on Thor (JetPack R38.2.2).

### Benchmark scoreboard (selected, with dates and caveats)

