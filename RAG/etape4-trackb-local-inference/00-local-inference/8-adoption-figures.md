---
id: etape4-trackb-local-inference/00-local-inference/8-adoption-figures
title: "8. Adoption figures"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["AMD", "Anthropic", "Apple", "Hugging Face", "Intel", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["accelerator", "agents", "amd", "benchmarks", "chatgpt", "claude", "gguf", "gpu", "intel", "llama", "llama.cpp", "memory"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [547, 602]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 2b2b6115dee94af79dc7f42deacc3309e67a164d689a2b674ec594d992bf3094
---

# 8. Adoption figures

## 8. Adoption figures

| Metric | Value | Date | Provenance |
|---|---|---|---|
| GitHub stars | **178,535** (#1 in conversational-AI tracked set) | Aug 2026 | [secondary: dappros/state-of-conversational-ai 2026-Q3] |
| GitHub stars | ~177,000 | 28 Jul 2026 | [secondary: agents-radar AI OSS trends] |
| GitHub stars / forks | **176,000 / ~17,000** | 9 Jul 2026 | [independent: TechCrunch] |
| GitHub stars | 169,752 | 23 Apr 2026 | [secondary: AppSelfHost] |
| GitHub stars | ~164,300 | 7 Mar 2026 | [secondary: github-topstar daily] |
| Monthly developers | **~8.9M** | Jul 2026 | [vendor-reported via TechCrunch] |
| Fortune 500 penetration | **85%** | Jul 2026 | [vendor-reported via TechCrunch] |
| Installs | **~1M/week** | Jul 2026 | [vendor-reported via ollama-herd research citing TechCrunch/TFN] |
| Monthly downloads | 100K (Q1 2023) → **52M (Q1 2026)**, 520x | Mar 2026 | [unverified: blackroados research notes — single source, methodology unclear] |
| Rank vs peers (stars, Aug 2026) | ollama 178,535 > langflow 153,239 > dify 152,436 > open-webui 148,767 > langchain 144,246 > llama.cpp 123,903 > vLLM 89,044 | Aug 2026 | [secondary: dappros report] |
| GGUF models on Hugging Face | 135,000 | Mar 2026 | [unverified: blackroados notes] |

- Sources: https://techcrunch.com/2026/07/09/popular-open-source-ai-developer-tool-ollama-raises-65m-grows-to-nearly-9m-users/ ; https://github.com/dappros/state-of-conversational-ai/blob/HEAD/2026-q3/report/report.md ; https://github.com/geeks-accelerator/ollama-herd/blob/HEAD/docs/research/mlx-vs-ollama-adoption-2026.md

---

## 9. Desktop app, engine, quantization, hardware backends

### Desktop app
- Native apps for **macOS and Windows** (download at ollama.com/download); an **Apps page** connects third-party apps — sign-in handled by Ollama [official: docs.ollama.com/cloud].
- **New desktop onboarding flow** on first launch (v0.32.15, Aug 2026) [secondary].
- **First-run setup** in CLI (`ollama`) offering sign-in or local-only; state shared with desktop app (v0.34.2, Sep 2026) [secondary].
- `ollama://apps` deep links open the desktop Apps page (v0.34.2) [secondary].
- **ChatGPT Desktop integration** (v0.34.0, Sep 2026): use Ollama models inside ChatGPT Desktop, configured from the Ollama macOS app [official].
- **Claude Desktop** configurable with Ollama as third-party gateway provider (v0.33.0, Aug 2026) [secondary].
- Cached resolved model metadata → TTFT ~995 ms → ~524 ms (v0.32.15 benchmarks) [secondary].

### Engines
- **llama.cpp**: direct integration since v0.30.0 (Jun 2026) — no longer via GGML layer; GGUF-native [secondary].
- **MLX runner** (Apple Silicon): DFlash speculative decoding, NVFP4 support, MLX safetensors model creation GA in v0.34.1 [official/secondary].
- **MLX-C** referenced in v0.33.3 engine updates [official].

### Hardware backends & GPU support [official: docs/windows.mdx unless noted]
- **Windows** (10 22H2+, native app, no admin required): NVIDIA CUDA (551.61+ drivers), AMD **ROCm v7 / HIP7** driver stack, or **Vulkan** for AMD Radeon; Vulkan is enabled by default and the recommended fallback where ROCm v7 isn't exposed (e.g. RDNA2/RX 6000) [official: https://github.com/ollama/ollama/blob/HEAD/docs/windows.mdx].
- **macOS Apple Silicon**: Metal (native); **macOS Intel**: CPU only [secondary: whisperjav research].
- **Linux x64**: CUDA, ROCm, Vulkan (experimental); **Docker**: CUDA, ROCm [secondary].
- **NVIDIA Jetson**: CC 87 / CUDA v13 support (v0.30.11) [official].
- **Windows on Arm**: Ollama ships Windows-on-Arm builds (noted working on NVIDIA RTX Spark / N1X at launch, Jun 2026) [secondary: awesomeagents.ai].
- GPU discovery: two-phase bootstrap, multi-backend detection (CUDA/ROCm/Metal/Vulkan) with dedup, auto VRAM-based layer offloading, CPU fallback [secondary].
- Mixed iGPU/dGPU control: `GGML_VK_VISIBLE_DEVICES` (Vulkan indices); fixed inverted iGPU/dGPU Vulkan classification on Windows hybrid graphics (v0.30.11) [official].
- Known 2026 pain points (community): Vulkan backend numerically wrong outputs on RDNA 4 (gfx1201, e.g. RX 9070 XT) — community workarounds use standalone llama.cpp ROCm 7 builds [secondary: tokenpal doc, Apr 2026 debug session].

### Quantization & model creation
- Primary format **GGUF**; Q4_K_M the community sweet spot [secondary].
- **`ollama create` from MLX safetensors no longer experimental** (v0.34.1); **GGUF creation now requires llama.cpp tooling** for conversion/quantization [official].
- **NVFP4** models supported on MLX (Metal quality bug fixed v0.32.5) [secondary].
- **ModelOpt** checkpoints with float32 global scale: fused kernel → +7.9% prefill on qwen3.6:27b (M5 Max) in v0.32.10-rc0 [secondary].
- Speculative decoding: unified/tuned in mlxrunner (v0.30.11); draft-model output heads quantized at requested type (v0.32.4); memory-growth fix for MLX speculative decoding (v0.34.2) [official/secondary].
- Hugging Face direct: `ollama run hf.co/user/repo:Q4_K_M` [secondary].

---

