---
id: collect-261001-ia-llm/ia-llm/etape4-trackb-local-inference-6
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "DeepSeek", "Meta", "Nvidia", "OpenAI", "vLLM"]
dates: ["2026-07", "2026-09", "2026-09-07"]
keywords: ["inference", "llama", "llama.cpp", "agent", "attention", "benchmark", "chatgpt", "claude", "compute", "deepseek", "embedding", "funding"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [222, 289]
sha256: 2b5636a57f576df7351932698ba02763902c3b8e43be1dad1fc5beab275d97ef
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)

### Other press/independent numbers
- **TechRadar** (gpt-oss launch, Aug 2025 window): NVIDIA RTX AI PCs + llama.cpp — RTX 5090 **282 tok/s** on gpt-oss-20b vs M3 Ultra 116 vs 7900 XTX 102 [independent/vendor-reported] (https://www.techradar.com/ai-platforms-assistants/gpt-oss-20b-performance-faster-pc-rtx-nvidia).
- **RunAIHome** (2026-09-07): DFlash2 guide with honest 1.8–2.7× guidance [secondary].
- Community hardware guides (vucense, Medium data-science-collective, mayhemcode) routinely use llama.cpp numbers as the reference engine; key framing: VRAM capacity is the first constraint; Strix Halo is a capacity play (128 GB unified, ~256 GB/s → ~6 t/s ceiling on 70B Q4_K_M) [secondary] (https://vucense.com/tech-reviews/compute-chips/local-llm-hardware-2026-strix-halo-m5-ultra-rtx-5090-70b-models/).
- arXiv 2601.14277 (Jan 2026): unified GGUF quantization evaluation on Llama-3.1-8B-Instruct (quality table above) [independent] (https://arxiv.org/pdf/2601.14277v1).
- arXiv 2511.05502 (Nov 2025): Apple Silicon engine comparison incl. llama.cpp ≈150 t/s on M2 Ultra [independent] (https://arxiv.org/pdf/2511.05502v1.pdf).

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

## 1. Executive summary

- Ollama remains the dominant local-LLM runner in 2026: #1 most-starred conversational-AI open-source project (178,535 stars, Aug 2026 [secondary]), ~8.9M monthly developers, 85% of Fortune 500 [vendor-reported via TechCrunch, 9 Jul 2026], ~1M installs/week [vendor-reported].
- Funding: **$65M Series B led by Theory Ventures announced 9 July 2026**; $88M raised total (prior $15M Series A, Benchmark) [independent].
- 2026 release cadence: v0.21 (Apr) → v0.30 (Jun, engine rebuild on llama.cpp + MLX) → v0.32 (Jul–Aug, Muse Glimmer, coding-agent "launch" integrations) → v0.33 (late Aug) → **v0.34.2 (15 Sep 2026, latest at cut-off)** [official/secondary].
- Strategic pivot of 2026: from pure local runner to **local + cloud hybrid** — Ollama Cloud (`:cloud` models, per-token pricing, off-peak discounts), **coding-agent integrations** (`ollama launch claude/codex/pi/openclaw/hermes/dsh/muse`), and first-party **desktop apps** wired into ChatGPT Desktop and Claude Desktop.
- Pricing shifted during 2026 from GPU-time billing (still reported Jul 2026) to **transparent per-token pricing with usage credits**: Free $0 / Pro $20/mo ($60 usage) / Max $100/mo ($300 usage) / Team $500/mo ($1,000 shared) [secondary, Sep 2026]; Team pricing data conflicts with a seat-based figure (see §7, flagged).
- Competition: LM Studio (GUI/desktop lane, proprietary, paid enterprise) and vLLM (server-throughput lane, up to 2.6x throughput claims) are the two named challengers; 2026 coverage consistently ranks Ollama #1 for developer API use but notes weak concurrency and no built-in auth.

---

## 2. Company, team and funding

