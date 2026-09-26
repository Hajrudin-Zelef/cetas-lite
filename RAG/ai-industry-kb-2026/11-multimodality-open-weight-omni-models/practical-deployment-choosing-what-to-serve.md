---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/practical-deployment-choosing-what-to-serve
title: "Practical deployment — choosing what to serve"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["AMD", "AWS", "Alibaba", "Apple", "China", "DeepSeek", "Google", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Unsloth", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-01"]
keywords: ["agent", "agentic", "agents", "apache", "aws", "benchmarks", "compute", "cost", "deepseek", "diffusion", "embeddings", "fine-tuning"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6055, 6081]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: edede67f4d2533948cb860fa30031e55808bca2b1e54aa849049bf22d0ebb76b
---

# Practical deployment — choosing what to serve

- **vLLM-Omni**: repo `vllm-project/vllm-omni`; community release Nov 2025; **v0.14.0 (Jan 2026)** first stable; **v0.16.0 (Feb 2026)** rebased on upstream vLLM v0.16.0; paper arXiv 2602.02204. Serves text, image, video, audio **and diffusion models** (DiT image/video) from one framework, one API, one deployment; OpenAI-compatible routes (`/v1/chat/completions`, `/v1/audio/speech`, `/v1/audio/generate`, `/v1/images/generations`, `/v1/videos`). Model coverage: Qwen3-Omni / Qwen3-TTS, Bagel, MiMo-Audio, GLM-Image, diffusion stack; platforms CUDA/ROCm/NPU/XPU; pipelined AR + diffusion execution. Multimodal flags: `--mm-encoder-tp-mode data` (vision tower data-parallel, avoids all-reduce), `--enable-vit-cuda-graph`. AWS Deep Learning Containers ship `vllm:omni-cuda` images. (Cross-ref §6 — one line: vLLM-Omni is the dedicated omni serving framework; production-grade as of January 2026.)
- **SGLang — broadest day-0 multimodal coverage**: `--enable-multimodal`; `--mm-process-config` for per-modality input limits; `--keep-mm-feature-on-device` (latency vs VRAM tradeoff; default moves feature tensors to CPU). **sglang-jax** has a standalone multimodal subsystem: text-to-image/video (Wan), vision-language (Qwen2.5-VL), **omni (Qwen3-Omni)**, audio (MiMo Audio); in-model path for regular AR multimodal models. Day-0 wins: Qwen3.6, Kimi K2.6, GLM-5.1, MiniMax M2.5/M2.6; native Kimi K2.5 (MoonViT + PatchMerger); Inkling-Small with MTP speculative decoding. Caveat: new hybrid checkpoints (e.g. DeepSeek-V4-Flash + MoonViT combos) sometimes need external model packages + narrow patches. **RadixAttention note**: multimodal mode auto-disables the radix cache — visual token sequences almost never share prefixes. Multi-node omni: Qwen3.8-Max (2.4T/95B, Aug 2026) runs on SGLang across 8×A800 with DeepEP MoE expert parallelism; MiMo-V2.6-Flash on SGLang TP16/DP2.
- **llama.cpp / GGUF — mature vision, experimental audio, nascent video**: `tools/mtmd` subsystem; vision encoder runs as a separate model, embeddings injected into the LLM token stream; **mmproj sidecars**; supported: LLaVA, MiniCPM-V, Pixtral, **Gemma 3/4**, Qwen2.5-VL, dots3-note, Qwen3.8-Flash-Next (qwen4exp) image handling. Audio: Whisper / MERaLiON-2 encoders; Gemma 4 E2B's GGUF projector reports **vision+audio** via mtmd; OpenAI-style `input_audio` message parts for Ultravox/Qwen2.5-Omni-class models. Video: `--video-*` CLI flags added (#24318); end-to-end video **not qualified** — requires `MTMD_VIDEO` builds + FFmpeg; workaround: extract frames, send as images. Quantization: MoE experts quantize like all weights; **no expert-parallelism fabric** (wide-EP is the llm-d/NVIDIA Dynamo layer); `vllm-gguf-plugin` brings experimental GGUF to vLLM. Unsloth Dynamic 2.0 GGUFs (Qwen3.x, GLM-5.3, Kimi, DeepSeek-V4, Gemma 4); UD-IQ2_XXS squeezes models like Qwen3.8-27B into 8 GB.
- **MLX / Apple Silicon and local tooling**: mlx-vlm (Blaizzy) vision fine-tuning + inference on M-series (Gemma 4 Vision, Qwen2.5-VL); mlx-community quants (Qwen3-VL-4B 3-bit ~3 GB, 8B 4-bit ~6 GB, 30B-A3B 6-bit ~20 GB); Whisper/Parakeet STT and Kokoro/Chatterbox TTS via mlx-audio; H3 MLX 8-bit port within days of weights. Unsloth claims image/video diffusion and multimodal training generally; Apple-Silicon video-understanding support specifically is community-led (mlx-vlm), not a confirmed Unsloth product capability. [UNVERIFIED — as narrowed in source.]
- **Ollama**: `gemma4:12b-it-qat` (7.2 GB), `gemma4:e2b-it-qat` (4.3 GB), `qwen3.5:9b` (6.6 GB, native vision) — one-command local multimodal.

### Practical deployment — choosing what to serve

- **Rule 1**: check the engine coverage matrix before choosing a model, not after — vLLM-Omni for production omni + diffusion, SGLang for fastest day-0 support, llama.cpp/mtmd for local/edge (video still the gap).
- **Rule 2**: "open weights" ≠ "runs locally" — MiMo-Pro (MIT) and MiniMax M3 are data-center class; Gemma 4 E2B/E4B are phone-class. Qwen3.8-Omni-Flash has no weights at all.
- **Rule 3**: budget the modality encoders as a separate line item (ViT/audio towers often kept at higher precision; SGLang offloads their features to CPU by default).
- **Rule 4**: KV cache at 1M context is the binding constraint — prefer GQA/MLA/shared-KV/HybridKV/TurboQuant architectures for long omni prompts; DeepSeek V4.1-Flash's 890 bytes/token MXFP4 KV is the current state of the art for cheap context carrying.
- **Rule 5**: for token pricing, audio ~6.25 tok/sec and video fps×resolution dominate long-context omni costs; pair `$ per token` with `tokens per task` (reasoning-agent fan-out can multiply the billed surface even as unit prices fall).
- **Rule 6**: verify the license file in the HF repo before commercial use — M3's license is disputed, H3's excludes major territories, K3 has a MaaS resale gate, and H3 forbids training other models on its outputs.

## Implications

1. **The "omni gap" inverted on audio.** Open weights now match or beat closed models on audio understanding and short-clip multimodal tasks; the remaining moat is long-horizon visual agents (10–17× on VTB) and agentic video reasoning. For RAG builders: native video/audio input + 1M context collapses the old "transcribe then chunk" pipeline — long meeting recordings and video go directly into the model; binding constraints become token pricing (audio ~6.25 tok/sec) and perception quality on long temporal reasoning.
2. **Licensing improved dramatically in 2026 — and got complicated.** Apache 2.0 (Qwen3-Omni, Gemma 4) and MIT (MiMo-V2.6, GLM-5.3-Flash, DeepSeek V4.1-Flash) now cover frontier-class omni models; exceptions: Qwen3.8-Omni-Flash (API-only), MiniMax M3 (disputed), MiniMax H3 (territorially restricted), Kimi K3 (MaaS resale gate). The RAG's license taxonomy needs a *territorially-restricted community license* category — "open weights" alone is no longer a sufficient label.
3. **Training transparency as a differentiator.** Xiaomi's livestreamed $3M+ RL run (60K sandboxes, public dashboard) sets a new bar for open-weight credibility — but vendor benchmarks remain unreproduced; treat decimals as decoration, provenance-tag every score.
4. **Serving converged on three stacks.** vLLM-Omni (production omni + diffusion), SGLang (fastest day-0 coverage), llama.cpp/mtmd (local/edge; video still the gap). The practical rule: check the engine coverage matrix before choosing a model, not after.
5. **China leads open omni.** Alibaba, MiniMax, Xiaomi, Zhipu, Moonshot released the flagship 2026 omni models; Google's Gemma 4 is the Western counterweight. Several were trained/served entirely on domestic Chinese accelerators (GLM-5.3-Flash: ~100,000 domestic chips).
6. **The any-to-any board is no longer closed-only.** MiniMax H3 is the first open model to top an AI video ranking — but its license bars most Western labs from local deployment, so the closed lead is now a *productization-and-licensing* lead, not a capability lead. Vendor compression claims need denominator discipline (the DeepSeek "437×" episode: record the denominator alongside the headline, prefer the replacement-model ~4× comparison).
7. **Test-time compute is now system design.** DeepSeek V4.1-Flash's KV engineering (CSA2 + MXFP4 + bounded replay) exists to make 1M-token agentic sessions affordable, and its price card (50× cache-hit/miss gap) monetizes exactly that — but third-party analysis (yage.ai) shows sub-agent fan-out can surge *total* token consumption even as per-token cost falls: always pair `$ per token` with `tokens per task`.

## Sources and URLs

