---
id: ai-industry-kb-2026/09-moe-architectures/attention-side-kv-economics-per-model-dated
title: "Attention-side KV economics per model (dated)"
domain: moe-architectures
role: deep-dive
task: finance
actors: ["Alibaba", "Apple", "CISA", "DeepSeek", "Huawei", "Meta", "MiniMax", "Moonshot", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-12-16", "2026-04-16", "2026-04-24", "2026-06-01", "2026-07-27", "2026-08-25", "2026-09", "2026-09-02", "2026-09-09"]
keywords: ["attention", "ascend", "benchmark", "compute", "consumer", "decode", "deepseek", "dpo", "embeddings", "fine-tuning", "fp8", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5145, 5174]
section: "9. MoE Architectures"
sha256: f4eec1377b7430e8072453e53756ecc7783518f429608c5e3b2b719be318792b
---

# Attention-side KV economics per model (dated)

- **Expert offloading (Gemma 4 26B-A4B case study, 2026-04):** only ~4.2–4.8 GB must stay hot in VRAM (embeddings ~1.3 GB, per-layer attention+router ~2.4 GB, shared expert ~50 MB); the 128 routed experts (~11 GB in Q4_K_M) sit on NVMe with ~24 MB loaded per forward pass (8 experts × ~3 MB) — servable on an 8 GB RTX 3070 with flash offloading. Note the asymmetry: MiMo-V2.6-Flash (309B/15B) has NO shared experts — the outlier; offloading economics differ when every expert is routed.
- **Unsloth Studio: MoE expert layers → system memory (v0.1.501-beta, Sept 2026):** automatic GPU placement moves MoE expert layers into system RAM so larger models fit; also added: split models across GPUs, tensor parallelism, per-model-per-quant hardware profiles, improved GGUF memory estimates. This is the documented "modest-hardware MoE" *deployment/fitting* mechanism — it is a placement/inference feature, NOT a training benchmark, and must not be cited as evidence for the 70B-on-24GB training claim.
- **What remains [UNVERIFIED]:** fine-tuning a 70B-class model on a single 24 GB RTX 4090/5090 — no benchmark or report found; the verified 70B config is one 80 GB GPU (Llama 3.3 70B QLoRA). "RTX 4090/5090 clusters" as a training claim is also unverified (consumer-card clusters are PCIe-bandwidth-bound for FSDP2/EP). Adjacent 2026 datapoints: Llama 3.1 8B QLoRA reportedly < 2 h on RTX 4090; gpt-oss-20b reportedly fine-tunable within 12.8 GB [PARTIALLY VERIFIED, vendor/secondary]; Axolotl `quantize_moe_experts` cut GLM-4.7-Flash reserved memory from ~127 GiB to ~23 GiB (secondary-reported); a "Soup"-style layer-streaming approach (Aug 2026) fine-tunes 8B models on a 4 GB laptop GPU; Unsloth's own Qwen3.8 guide states QLoRA works at 24 GB but 27B-class LoRA needs >36 GB, full FT uses 4× more VRAM.
- **Consumer inference datums:** Qwen3.6-35B-A3B Unsloth UD-Q4_K_XL GGUF ~21 GB on a single 24 GB RTX 4090; 20.9 GB Q4 quant on MacBook Pro; Apple Silicon MLX: fine-tune large MoE models (text/images) on Mac, up to 30x faster follow-up turns in long Qwen chats [VENDOR-REPORTED], quantized MLX KV caches use up to 74% less prompt memory [VENDOR-REPORTED], DoRA fine-tuning and more DPO loss types on Apple Silicon (Studio v0.1.807/808-beta, early Sept 2026).

### Attention-side KV economics per model (dated)

- **DeepSeek-V2 (2024):** MLA cut KV cache to ~5% of LLaMA-3-70B's — the reference KV-compression datum all 2026 work is measured against.
- **MiMo-V2-Flash (2025-12-16):** 5:1 SWA/GA with 128-token window + learnable attention-sink bias → ~6× KV reduction; 150 tok/s via MTP self-speculative decoding.
- **Gemma 4 (~2026-04):** 1024-token-window SWA alternating with global attention; shared KV cache; pure-attention (no SSM), so prefill scales to full GPU TFLOPS with no sequential bottleneck.
- **Qwen3.6-35B-A3B (2026-04-16):** 30 Gated DeltaNet linear-attention layers (O(N) in sequence length) + 10 gated attention layers (16Q/2KV, head dim 256); 262K native context, ~1M via YaRN RoPE scaling.
- **MiniMax M3 (2026-06-01):** MSA block-sparse "lightning indexer" (top-k 128-token blocks); ~9× prefill / ~15× decode speedup over M2 at 1M context; ~1/20 per-token compute of the previous generation; 64Q/4KV GQA.
- **Kimi K3 (2026-07-27):** KDA up to 6.3× faster decoding; FlashKDA 1.72–2.22× prefill speedup on H20; 69 KDA + 24 Gated MLA layers in the 93-layer stack.
- **GLM-5.3-Flash (2026-08-25/26):** hybrid sparse+linear attention + mHC; claimed 4.4× smaller KV cache [VENDOR]; 288 routed experts, 8 active per token.
- **DeepSeek-V4-Flash (2026-04-24):** NSA + DSA + token-wise compression [PARTIALLY VERIFIED reconstruction]; up to 384K output tokens; dual thinking/non-thinking modes.
- **MiMo-V2.6-Flash (~2026-09):** 39 sliding-window + 9 global attention layers, 128-token window; five-layer MTP drafter (DFlash-style, 7 tokens ahead per forward pass).
- **Post-hoc KV compression (orthogonal to model choice):** HybridKV up to 7.9×; TurboQuant 3-bit KV ÷6 memory; quantized MLX KV caches "up to 74% less prompt memory" [VENDOR, Unsloth Studio v0.1.807/808-beta, Sept 2026].

### Stability and routing failure modes (dated)

- **2024 lineage → 2026 standard:** DeepSeek-V3 aux-loss-free dynamic bias; expert-choice routing (balanced by construction); Kimi K3 Quantile Balancing — each expert's bias set to the router-score quantile matching its target load, established by a single global histogram/all-reduce. The DeepSeek-V4 reconstruction adds a small sequence-wise balance loss, sqrt(softplus) routing affinity (replacing sigmoid), and deterministic token-ID hash routing on the first three blocks instead of learned routing [PARTIALLY VERIFIED].
- **Failure taxonomy (community tutorial consensus, 2026):** routing collapse, expert imbalance, EP communication instability — MoE training stability rated "medium" vs dense "high".
- **Serving-side imbalance protection:** EPLB (vLLM-Ascend) records expert-traffic maps and rebalances without stop-the-world; dynamic mode avoids TTFT/TPOT spikes; verified on DeepSeek-V3.1/R1; W8A8/W4A8/MXFP4/MXFP8 quant types on Ascend 950.
- **Capacity factor:** C = ⌈α·T·k/N⌉ bounds tokens per expert per step; overflow tokens may be dropped. Token-choice top-k with training-time noise for exploration; orthogonality + variance regularization on routing scores in fine-grained designs.
- **Precision-stability coupling:** Kimi K3 bounds activations (SiTU, ‖f(x)‖∞ ≤ 100) + RMSNorm before up-projection to stabilize MXFP4 training; DeepSeek-V3 pioneered FP8 multi-plane training (forward and backward); Unsloth's 2026-09-09 grouped-GEMM fix corrected silently-wrong gradients in the MoE LoRA path — a live correctness (not performance) incident in the fastest-moving MoE training path of September 2026.
- **Communication and optimizer hazards:** DeepSpeed MoE all-reduce penalties of 3–12 s in problematic large-scale layouts (fixed via multi-rank bucketing + rank placement); per-head Muon optimization (Kimi K3); framework reproducibility hazard — PyTorch 2.14 (2026-09-02) silently changed clamp/min/max boundary subgradients (1→0).
- **Small-scale caveat:** at ~150M active params (OliverSundaram/MoE-Study, 2026), dense beat MoE on almost every benchmark and ran 3× faster at inference — routing overhead dominates at small scale. MoE "earns its complexity" only as sparse capacity compounds (pre-registered 2026-06 ablation: MoE perplexity 5.72 vs dense 22.66 at matched compute, gap widening with training).

### MoE+SSM hybrids and alternative sparsity (dated)

