---
id: ai-industry-kb-2026/09-moe-architectures/quantization-and-serving-format-timeline-dated
title: "Quantization and serving-format timeline (dated)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "CISA", "DeepSeek", "Meta", "MiniMax", "Moonshot", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-08-05", "2025-12-16", "2026-04-24", "2026-06-01", "2026-07-22", "2026-08-19", "2026-08-22", "2026-09-15", "2026-09-22"]
keywords: ["quantization", "attention", "awq", "compute", "cost", "decode", "deepseek", "fp4", "fp8", "gguf", "glm", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5121, 5144]
section: "9. MoE Architectures"
sha256: 80186f72ecfa47b0b71833b92b2f174fe414ad027992c74e82aba78aa4c909ab
---

# Quantization and serving-format timeline (dated)

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

