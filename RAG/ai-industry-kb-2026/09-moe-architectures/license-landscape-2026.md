---
id: ai-industry-kb-2026/09-moe-architectures/license-landscape-2026
title: "License landscape (2026)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "Apple", "CISA", "DeepSeek", "Huawei", "Meta", "MiniMax", "Moonshot", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-08-05", "2025-12-16", "2026-04", "2026-04-16", "2026-04-24", "2026-06-01", "2026-07", "2026-07-22", "2026-07-27", "2026-08-19", "2026-08-22", "2026-08-25", "2026-09", "2026-09-02", "2026-09-09", "2026-09-15", "2026-09-22"]
keywords: ["license", "apache", "ascend", "attention", "awq", "benchmark", "compute", "consumer", "cost", "decode", "deepseek", "dpo"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5103, 5172]
section: "9. MoE Architectures"
sha256: 639d3efdef00140ee470ad120cf9e1c645db72b23084a48551aa9798935afe30
---

# License landscape (2026)

### License landscape (2026)

- **Apache 2.0:** Qwen3.6-35B-A3B, Gemma 4 family (notable shift from earlier Gemma-specific terms), gpt-oss-120b/20b.
- **MIT:** Kimi K3, GLM-5.3-Flash, MiMo-V2-Flash, MiMo-V2.6-Flash/Pro, DeepSeek-V4 (per later sources; early April 2026 coverage cited Apache 2.0 — UNRESOLVED, do not cite without checking).
- **Custom / community:** MiniMax M3 (MiniMax Community License), GLM-5.3 flagship (glm-5.3 license), Qwen3.8-Max (non-Apache open license).
- **Trend:** the most permissive standard licenses (Apache 2.0, MIT) now cover the majority of open-weight MoE releases; custom licenses cluster at the largest/capability-leading models (M3, GLM-5.3, Qwen3.8-Max).

### Quantization formats used across the 2026 MoE fleet

- **FP8 (e4m3):** default serving format for 300B+ models — GLM-5.3-Flash ships FP8 by default (~306–331 GiB); MiMo-V2.6-Flash weights 172.9 GB in FP8; MiniMax M3 MXFP8 variant ~440 GB; Qwen3.8-Max's FP8 repo is the most-downloaded variant (21,400 downloads vs 17,386 for bf16 as of 2026-08-22 — FP8 leading, indicating serving-motivated fetchers).
- **NVFP4 (group-16):** MiniMax M3 (~245 GB safetensors → ~61 GB/GPU at TP4); vLLM v0.15.0 added NVFP4-CUTLASS kernels; Unsloth Studio exports to NVFP4 since July 2026 (v0.1.481-beta); Axolotl added NVFP4 (4-bit) MoE LoRA in July 2026 via ScatterMoE (W4A16) and SonicMoE (W4A4), including lossless adapter merge back into a plain NVFP4 checkpoint. Third-party nuance (unsloth-cli docs, Sept 2026): NVFP4 exports were live-tested for *serving* via vLLM on Thor (JetPack R38.2.2), not for training — Unsloth/NVFP4 claims to date are export-and-serve claims, not NVFP4-training claims.
- **MXFP4/MXFP8:** native in Kimi K3 (MXFP4 weights / MXFP8 activations; per-expert MXFP4 with E8M0 scale per 32 weights, 33.0M params per expert, 2.72T routed params total) and gpt-oss (MXFP4 native, 60.8 GiB / 12.8 GiB checkpoints).
- **AWQ / Marlin 4-bit:** local-serving standard; documented Mixtral 8x7B vLLM case: 90 GB → 26 GB VRAM, 2.1× throughput, −75% TTFT.
- **GGUF (Q4_K_M and below):** CPU/offload tier — Qwen3.6 20.9 GB Q4 on MacBook Pro; Gemma 4 26B-A4B ~11 GB routed experts on NVMe with ~24 MB/forward hot-loading. Unsloth's Dynamic v3.0 remains the current UD quant generation (launched 2026-08-19); GLM-5.2's Dynamic GGUF was presented by a secondary source as 84% size reduction (~239 GB at 2-bit, ~217 GB at 1-bit) with ~82%/~76% accuracy retention and ~98% at 4-bit, runnable on 256 GB unified-memory machines — [VENDOR-ORIGINATED, UNVERIFIED].
- **Invariance principle (verified):** quantization scales total and active params equally — the total/active ratio (e.g., 5.40× at fp32 and int4 in the OLMoE profile) does not change.

### Attention innovations glossary (what makes 1M context affordable)

- **MSA — MiniMax Sparse Attention (MiniMax M3, 2026-06):** block-sparse "lightning indexer" attention selecting top-k 128-token blocks; decode cost roughly flat in context length; ~9× prefill / ~15× decode speedup over M2 at 1M context; ~1/20 per-token compute of the previous generation.
- **KDA — Kimi Delta Attention (Kimi K3, 2026-07):** up to 6.3× faster decoding; paired with Attention Residuals (AttnRes, ~25% higher training efficiency); optimized kernel FlashKDA (1.72–2.22× prefill on H20). 69 KDA + 24 Gated MLA layers in the 93-layer stack.
- **Gated DeltaNet (Qwen3.6, 2026-04; Qwen3.8-Max, 2026-08):** linear attention with delta-rule memory update and gating; O(N) in sequence length; interleaved 3:1 with gated standard attention in Qwen3.6 (30 DeltaNet + 10 attention layers); hybrid stack in Qwen3.8-Max (128 linear heads for V / 16 for QK). Unsloth's Sept 2026 Qwen3.8 guide adds "Flash Linear Attention kernels" for Gated DeltaNet training.
- **Gated Attention (Qwen3.6/3.8):** standard softmax attention with learned gating; 16Q/2KV heads (Qwen3.6), 64Q/4KV heads (Qwen3.8-Max).
- **MLA — Multi-head Latent Attention (DeepSeek lineage):** compresses KV into a latent vector; DeepSeek-V2's MLA cut KV cache to ~5% of LLaMA-3-70B's. Kimi K3 uses Gated MLA in 24 of 93 layers. V4-era evolution: CSA/HCA hybrid attention replacing the earlier MLA design [PARTIALLY VERIFIED, secondary reconstructions].
- **NSA — Native Sparse Attention (DeepSeek-V4, 2026-04):** sparse attention native to training (not a post-hoc approximation); Flash variant adds DSA + token-wise compression.
- **SWA hybrids (MiMo, Gemma 4, GLM-5.3-Flash):** sliding-window local attention interleaved with global layers; MiMo's 5:1 SWA/GA with 128-token window + learnable attention-sink bias → ~6× KV reduction; Gemma 4 alternates 1024-window SWA with global attention; GLM-5.3-Flash claims 4.4× smaller KV cache via its hybrid sparse+linear stack.
- **mHC — Manifold-Constrained Hyper-Connections (GLM-5.3-Flash, 2026-08):** connection-scaling technique claimed to improve training efficiency at 320B scale.
- **SiTU — Sigmoid Tanh Unit (Kimi K3, 2026-07):** replaces SiLU/GELU in the 2.8T model, bounded activations for low-precision stability.

### Quantization and serving-format timeline (dated)

- **2025-08-05:** gpt-oss ships MXFP4 native (60.8 GiB 120b / 12.8 GiB 20b) — sets the native sub-8-bit template every 2026 release follows.
- **2025-12-16:** MiMo-V2-Flash trains in FP8 mixed precision and ships day-0 SGLang support — FP8 moves from experiment to production format.
- **2026-04-24:** DeepSeek-V4 routed expert weights reportedly deployed in FP4 [PARTIALLY VERIFIED, secondary reconstructions].
- **2026-06-01:** MiniMax M3 publishes NVFP4 (~245 GB safetensors → ~61 GB/GPU at TP4) and MXFP8 (~440 GB) checkpoints; FP8/MXFP8 and NVFP4 variants all exist.
- **2026-07:** Axolotl adds NVFP4 (4-bit) MoE LoRA via ScatterMoE (W4A16) and SonicMoE (W4A4), including lossless adapter merge back into a plain NVFP4 checkpoint; Unsloth Studio v0.1.481-beta adds NVFP4/FP8/imatrix-GGUF export with DeepSeek-V4 support; Kimi K3 trains natively in MXFP4 weights / MXFP8 activations (E8M0 scale per 32 weights).
- **2026-07-22:** MarkTechPost 4-framework comparison — B200 gpt-oss 8K context: Unsloth 47.43 GB vs Transformers v5 73.80 GB; 16K: Transformers v5 OOMs vs Unsloth 55.13 GB.
- **2026 (vLLM v0.15.0):** Marlin/NVFP4-CUTLASS/FP8/INT8 kernels for non-gated MoE land in vLLM.
- **2026-08:** GLM-5.3-Flash ships FP8 by default (~306–331 GiB); Unsloth Dynamic v3.0 launches (2026-08-19), remains the current UD quant generation as of 2026-09-22; Qwen3.8-Max's FP8 repo out-downloads bf16 21,400 vs 17,386 (2026-08-22) — serving-motivated fetchers vote with downloads.
- **2026-09:** Unsloth's GGUF one-call export API documented (`save_pretrained_gguf`, 23-entry quant list, multi-quant in one call, `push_to_hub_gguf`, save-time OOM guard `maximum_memory_usage` default 0.75); third-party unsloth-cli 0.7.0/0.7.1 (2026-09-15) operationalizes container-backed merged-16bit/4bit/GGUF/AWQ/NVFP4 export with quantization-loss eval (`sloth eval --model DIR`); GLM-5.2 Dynamic GGUF secondary figures: ~239 GB at 2-bit / ~82% accuracy retention, ~217 GB at 1-bit / ~76%, ~98% at 4-bit, runnable on 256 GB unified-memory machines [VENDOR-ORIGINATED, UNVERIFIED].

### MoE on modest hardware: documented mechanisms

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

