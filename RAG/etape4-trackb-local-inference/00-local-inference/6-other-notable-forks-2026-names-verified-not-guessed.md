---
id: etape4-trackb-local-inference/00-local-inference/6-other-notable-forks-2026-names-verified-not-guessed
title: "6. Other notable forks (2026) — names verified, not guessed"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "DeepSeek", "Hugging Face", "Moonshot", "Nvidia", "OpenAI"]
dates: ["2025-06", "2026-01-27", "2026-05", "2026-06-04", "2026-08-01", "2026-08-31", "2026-09", "2026-09-07", "2026-09-14", "2026-09-15"]
keywords: ["agentic", "amd", "benchmark", "claude", "compute", "deepseek", "embeddings", "gguf", "gpus", "inference", "kimi", "lean"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [181, 228]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 41453bd4a6299a9a248e29694eb74eb766a249e99d3fc56e582ced916a13e44e
---

# 6. Other notable forks (2026) — names verified, not guessed

## 6. Other notable forks (2026) — names verified, not guessed

| Fork | Focus | Verification |
|---|---|---|
| **artomyuan/llama.cpp-rocm** | ROCm-first fork: ROCmFPX quants (GGML types 100–107), TurboQuant, mtmd-grounders multimodal vision (LocateAnything + DeepSeek4V dual projector); synced to upstream 97e4ca735 on 2026-09-14; claims pp512 +105% / pp2048 +109% (measured 2026-09-15) | [secondary] https://github.com/artomyuan/llama.cpp-rocm/blob/HEAD/CHANGELOG.en.md |
| **wszhoho/llama-cpp-turboquant-dflash2** | TurboQuant (TURBO2_0/3_0/4_0 = enum 43/44/45) + DFlash/DFlash2/DSpark drafts + DeepSeek-V4 ecosystem; tracks upstream closely (Qwen3.5, Gemma4, Kimi-K3 support) | [secondary] https://github.com/wszhoho/llama-cpp-turboquant-dflash2 |
| **prismml-eng/llama.cpp** | OpenVINO-oriented fork (docs/backend/OPENVINO.md with Ubuntu install scripts for OpenVINO 2026.3) | [secondary] https://github.com/prismml-eng/llama.cpp/blob/HEAD/docs/backend/OPENVINO.md |
| **kyuz0/amd-strix-halo-toolboxes** | ROCm toolboxes + llama.cpp builds for Strix Halo (gfx1151), incl. **EngramHalo** variant for Qwen3.8-Flash-Next's 51B engram table; rocm-10.0 images (Aug 2026) | [secondary] https://github.com/sypherin/strix-halo-setup/blob/HEAD/docs/qwen3.8-flash-next-engramhalo-stability.md; https://github.com/ashebanow/nix-config/commit/7f93e6e59ffc6fb2bd7d8eab53e5960d7670bd0e |
| **nathanw1014/strix-halo-llamacpp** | Strix Halo branch: DFlash2 validation on gfx1151/Vulkan (Aug 2026) | [secondary] https://github.com/nathanw1014/strix-halo-llamacpp/blob/HEAD/docs/dflash2-strix.md |
| **pugant/strix-halo-llamacpp-lab**, **pugant/strix-nebulosa** | ROCmFP4-STRIX_LEAN preset origin; ships fork source tree `rocmfpx/` | [secondary] https://github.com/pugant/strix-halo-llamacpp-lab |
| **mrlordcat/llama.cpp-rdna-lab** | RDNA/Strix research docs (spec-decoding status, DFlash via `--spec-draft-device ROCm1`) | [secondary] https://github.com/mrlordcat/llama.cpp-rdna-lab/blob/HEAD/docs/research/SPEC_DECODING_STATUS.md |
| **Anbeeld/beellama.cpp** | Source of the DFlash port into llama.cpp (`c6dfa39e3`) | [secondary] (referenced by mrlordcat docs) |
| **poisonxa16/pxq_llama.cpp** | Documents itself as a fork tracking ik_llama.cpp (DELTA-SINCE-IK.md); mirrors ik README | [secondary] https://github.com/poisonxa16/pxq_llama.cpp/blob/HEAD/docs/README-upstream-ik_llama.md |
| **black6spdz/ik_llama.cpp** | Stale mirror of ik (last updated ~212 days ago) | [secondary] https://github.com/black6spdz/ik_llama.cpp |
| **tao71-ai/llama-cpp-python-jamepeng** | llama-cpp-python fork, syncs upstream frequently (2026-08-31: load_mode + Qwen3.8-Flash-Next + DFlash2 NVFP4 patch) | [secondary] https://github.com/tao71-ai/llama-cpp-python-jamepeng/blob/HEAD/CHANGELOG.md |
| **dushyant30suthar/opencode-llama.cpp** | Experiment notes repo (DFlash vs MTP measurements, exllamav3 cross-checks) | [secondary] https://github.com/dushyant30suthar/opencode-llama.cpp/blob/HEAD/docs/HANDOVER-2026-08-01.md |

Note: Tom's/ServeTheHome-level press does not cover these forks individually; the list above is verified by repo existence + 2026-dated content. Anything else (e.g., hypothetical "llama.cpp-openvino") was NOT verified and is excluded per the no-guessing rule.

---
## 7. Server mode (llama-server) — features as of September 2026

Current-master README (tools/server/README.md) lists [official] (https://github.com/ggml-org/llama.cpp/blob/master/tools/server/README.md):

- **OpenAI-compatible API**: chat completions, **responses** (`/v1/responses`), and embeddings routes; multimodal (vision/audio/video) via the OpenAI-compatible API; accepts `data:` URLs for media input (added in v0.4.0 window, #27735).
- **Anthropic Messages API-compatible** chat completions; **reranking endpoint** (#9510); assistant-message prefilling à la Claude API; `preserve_reasoning` **on by default** (v0.4.0, #28174).
- **Function calling / tool use for ~any model**; prefilled assistant tool calls now rejected (v0.4.0, #27626); **MCP**: web UI replaced "MCP overrides" with a **tool policy** (v0.4.0, #27745); server README documents **MCP stdio servers** and CORS defaults (#26847, v0.1.2 window).
- **Parallel decoding with multi-user support, continuous batching, speculative decoding** (incl. DFlash2/synthetic spec-acceptance options).
- Serving ops: per-slot context limits (v0.4.0, #24124); `/metrics` accessible even during sleep (#27376); model endpoints can be made private when auth is enabled (#26347); `dedup-cache-models` preset (#27346); router mode lazy-loads `startup_models` (#27424); docker-repo (`--docker-repo`, e.g. `gemma3`) and **Hugging Face shorthand** `--hf-repo <user>/<model>[:quant]` (default Q4_K_M; mmproj auto-download) — `llama serve -hf ggml-org/Qwen3.5-0.8B-GGUF`.
- CLI-serving flags of note: `--lazy-mode/-lzm` (on-demand tensor read), `--load-mode` (auto/mmap/mlock/mmap+mlock/dio), `--fit/--fit-target/--fit-ctx` (auto-fit to device memory), `-ncmoe`/`-ncffn` (keep first N MoE/FFN layers on CPU), `-sm {none,layer,row,tensor}`, `--spec-draft-type-k/v` KV types, `--video-*` params.
- Web UI: redesigned (June 2025, legacy available via `--path ./examples/server/public_legacy`); tabbed chat navigation (v0.3.0); chat form actions UX, agentic-response grouping/copy (v0.4.0); pytest-xdist server tests (v0.4.0).

## 8. Community & press coverage (2026)

### Tom's Hardware
- **"Nvidia DGX Spark review: the GB10 Superchip … beats out AMD's Ryzen AI Max+ 395"** (~2026-01-27, "238 days ago" at crawl): uses **llama.cpp as the standard LLM inference benchmark platform** ("broadly compatible … well-documented and tunable benchmarking interface") vs Corsair AI Workstation 300 (Ryzen AI Max+ 395). Tested prompt-processing/prefill and token generation on llama-3.1-8B, Gemma 3 12B/27B, plus a token-generation **power-efficiency test: 4096 tokens from GPT-OSS 120b measuring joules/token**, with GPT-OSS 20B on both platforms for comparison [independent] (https://www.tomshardware.com/pc-components/gpus/nvidia-dgx-spark-review/3).
- Also: OpenClaw on Beelink SER10 MAX (2026-08) used llama.cpp at **10.64 tok/s** for a general-knowledge query [independent] (https://www.tomshardware.com/tech-industry/artificial-intelligence/setting-up-openclaw-isnt-as-straightforward-as-the-internet-wants-you-to-think-running-local-ai-on-humble-hardware).

### ServeTheHome
- **No dedicated llama.cpp benchmark article was found on servethehome.com in 2026.** Community coverage lives in the **STH Forums**: ES Xeon Discussion thread (May 2026 posts — dual-Xeon + ik_llama.cpp DeepSeek-R1 numbers cited above), ASUS Pro WS W790E-SAGE / Xeon SPR-SP thread (DeepSeek-R1 on W790 + RTX 4090 via llama.cpp). STH main site adjacent: MLPerf Inference v5.0 results (June 4, 2026) and a Lenovo ThinkPad X1 Carbon Gen 14 review using **MLPerf Client 1.6.1** (OpenVINO path, not llama.cpp) [independent/secondary] (https://forums.servethehome.com/index.php?threads/es-xeon-discussion.5031/page-205; https://www.servethehome.com/lenovo-thinkpad-x1-carbon-gen-14-review/).

### Other press/independent numbers
- **TechRadar** (gpt-oss launch, Aug 2025 window): NVIDIA RTX AI PCs + llama.cpp — RTX 5090 **282 tok/s** on gpt-oss-20b vs M3 Ultra 116 vs 7900 XTX 102 [independent/vendor-reported] (https://www.techradar.com/ai-platforms-assistants/gpt-oss-20b-performance-faster-pc-rtx-nvidia).
- **RunAIHome** (2026-09-07): DFlash2 guide with honest 1.8–2.7× guidance [secondary].
- Community hardware guides (vucense, Medium data-science-collective, mayhemcode) routinely use llama.cpp numbers as the reference engine; key framing: VRAM capacity is the first constraint; Strix Halo is a capacity play (128 GB unified, ~256 GB/s → ~6 t/s ceiling on 70B Q4_K_M) [secondary] (https://vucense.com/tech-reviews/compute-chips/local-llm-hardware-2026-strix-halo-m5-ultra-rtx-5090-70b-models/).
- arXiv 2601.14277 (Jan 2026): unified GGUF quantization evaluation on Llama-3.1-8B-Instruct (quality table above) [independent] (https://arxiv.org/pdf/2601.14277v1).
- arXiv 2511.05502 (Nov 2025): Apple Silicon engine comparison incl. llama.cpp ≈150 t/s on M2 Ultra [independent] (https://arxiv.org/pdf/2511.05502v1.pdf).

