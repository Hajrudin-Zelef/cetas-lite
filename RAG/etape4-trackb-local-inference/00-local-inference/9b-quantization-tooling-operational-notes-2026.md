---
id: etape4-trackb-local-inference/00-local-inference/9b-quantization-tooling-operational-notes-2026
title: "9b. Quantization tooling & operational notes (2026)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: quantization
actors: ["Alibaba", "DeepSeek"]
dates: ["2026-09"]
keywords: ["quantization", "attention", "benchmark", "deepseek", "embedding", "gguf", "gpus", "inference", "kv cache", "llama", "llama.cpp", "memory"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [229, 276]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 9ea57795a6ea313996b580126446ef6a5cff87eb5c12d0f455e9ebd9744d7117
---

# 9b. Quantization tooling & operational notes (2026)

## 9b. Quantization tooling & operational notes (2026)
- **llama-quantize** (2026): RAM cap via new `max_buf_size` quantize param (#27795); row-slab streaming to avoid thread starvation (#27830); memory optimization by evicting weights after each layer (#22877, v0.2.0 window); combined with `--lazy-mode`/`-lzm` and the "prevent RAM peaking at load" fix (#27483) — 2026 is the year llama.cpp fixed its RAM spikes on both load and quantize paths [official] (v0.4.0, v0.2.0 changelogs).
- **KV-cache quantization**: allowed types unchanged (f32/f16/bf16/q8_0/q4_0/q4_1/iq4_nl/q5_0/q5_1); new per-model draft KV-cache variants in code (`llama-kv-cache-dsv4.cpp` draft ring, `llama-kv-cache-iswa.cpp`, `-dsa.cpp`, `-msa.cpp`) track new attention types (DSpark, DSA-ISWA, MSA) rather than new quant types [official/secondary] (v0.3.0 dots3-note; wszhoho fork README).
- **Multimodal (mtmd/libmtmd)**: vision/audio in llama-server; video params `--video-*`; `mtmd_tokenize_from_parts()`; DeepSeek-V4-Flash-Vision-Exp; dots3-note vision+audio; WebP via ffmpeg; Pillow-accurate resize; `--mmproj-device`; SAM deepseek-ocr conv2d with F32 im2col [official] (v0.3.0/v0.4.0 changelogs).
- **Finetune tool**: fixed "no KV cache" bug (#27199, v0.4.0 window) [official].
- **CI/release infra**: signed release artifacts + attestations (v0.2.0); release.sh + nightly-tag.txt; BoringSSL 0.20260903.0 vendored (v0.4.0); cpp-httplib 0.54.0 [official].

## 9. Uncertainties & gaps (flagged explicitly)
1. **Feb–Jul 2026 b-tag milestones**: only sampled (b7964/b7973/b8018/b8119/b8299/b8461/b10485). A complete commit-level timeline was out of scope; numbers are from changelogs/build records, not direct tag inspection.
2. **KleidiAI macOS build** is listed DISABLED in current release assets — reason not established [unverified].
3. **GGUF spec version**: no GGUF *container* version bump was found in 2026 changelogs; tokenizer/vocab metadata additions (integer scores, DFlash2 keys) documented above are the confirmed changes. Absence of a version bump is an inference from changelog silence — treat as [unverified].
4. **Tom's Hardware DGX Spark benchmark numbers** (charts) could not be extracted (page renders charts as images); the methodology/model list is confirmed, numeric pp/tg values are not.
5. **APEX I-Compact** quant format: seen only in one community benchmark table; no format spec verified.
6. **Qwen3.8-Flash-Next "engram"** (51B per-layer embedding table, SSD lazy-read) is described by a community fork (EngramHalo), not by upstream docs — details are [secondary].
7. Fork performance claims (e.g., artomyuan pp512 +105%) are fork-author-reported; not independently reproduced here.

## 10. Key source URLs
- https://github.com/ggml-org/llama.cpp/releases/tag/v0.4.0
- https://github.com/ggml-org/llama.cpp/releases/tag/v0.3.0
- https://github.com/ggml-org/llama.cpp/releases/tag/v0.2.0
- https://github.com/ggml-org/llama.cpp/releases/tag/b11001
- https://github.com/ggml-org/llama.cpp/blob/master/tools/server/README.md
- https://github.com/ggml-org/ggml/discussions/1579 (release versioning policy)
- https://github.com/ikawrakow/ik_llama.cpp
- https://github.com/artomyuan/llama.cpp-rocm/blob/HEAD/CHANGELOG.en.md
- https://github.com/wszhoho/llama-cpp-turboquant-dflash2
- https://www.tomshardware.com/pc-components/gpus/nvidia-dgx-spark-review/3
- https://www.techradar.com/ai-platforms-assistants/gpt-oss-20b-performance-faster-pc-rtx-nvidia
- https://runaihome.com/blog/qwen38-27b-dflash2-speculative-decoding-guide-2026/
- https://github.com/nathanw1014/strix-halo-llamacpp/blob/HEAD/docs/dflash2-strix.md
- https://arxiv.org/pdf/2601.14277v1 (quant evaluation)
- https://llama.app

---
*End of report. File: ~/workspace/rag_collect/etape4_draft_llamacpp.md*


---

## Part B — Ollama


> RAG data-collection project, Step 4, track B, part 2. Research cut-off: 22 September 2026.
> Provenance tags: `[official]` = Ollama/official docs & GitHub releases; `[vendor-reported]` = Ollama statements via press; `[independent]` = reputable press; `[secondary]` = third-party summaries/docs/repos quoting official material; `[unverified]` = single-source, low-confidence.
> **Rule followed: no version numbers or model names were guessed; everything below is sourced from tool output with URLs.**

---

