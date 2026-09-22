---
id: etape4-trackd-unsloth-training/00-unsloth-training/overview
title: "Step 4 — Track D: Unsloth + Training / Fine-Tuning Tooling (2026)"
domain: step-4-track-d-unsloth-training-fine-tuning-tooling-2026
role: deep-dive
task: training
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "China", "DeepSeek", "Google", "Hugging Face", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "Samsung", "Unsloth", "Z.ai"]
dates: ["2026-03", "2026-03-23", "2026-07-20", "2026-09", "2026-09-22"]
keywords: ["fine-tuning", "training", "agent", "agents", "amd", "apache", "attention", "benchmark", "benchmarks", "claude", "consumer", "deepseek"]
source: docs/RAG/etape4_trackD_unsloth_training.md
source_anchor: ""
source_lines: [1, 65]
section: "Step 4 — Track D: Unsloth + Training / Fine-Tuning Tooling (2026)"
sha256: 37e0b7f96d9ce497197a192b0dd5b8fd6742d490cb6312deb189f6ee6402e80b
---

# Step 4 — Track D: Unsloth + Training / Fine-Tuning Tooling (2026)

**Coverage window:** January 1 → September 22, 2026
**Report date:** September 22, 2026
**Status:** research snapshot; prices/versions are September-2026 snapshots. Re-verify against official repos before use.

## Provenance legend
- `[official]` — project's own repo, docs, changelog, or release notes.
- `[vendor-reported]` — claim made by the vendor/team without independent audit.
- `[independent]` — third-party measurement, reputable press, or peer analysis.
- `[secondary]` — blogs, aggregators, community docs (GitHub knowledge bases, AI China News); useful but unverified.
- `[unverified]` — single-source or conflicting; treat as uncertain.

---

## 1. UNSLOTH AI

### 1.1 Company and funding (as of September 2026)

- Unsloth AI (unsloth.ai), founded by brothers **Daniel Han** and **Michael Han** (Daniel — ex-NVIDIA background; LinkedIn `danielhanchen`) — Australia-based founders, corporate office **720 Clipper Street, San Francisco, CA** (per PitchBook; office is real-estate HQ/listing, team largely remote) [secondary — PitchBook company profile, Crunchbase].
- **Confirmed funding: ~$500K pre-seed** (Y Combinator pre-seed; Redpoint Ventures scout; Samsung NEXT; Pioneer Fund). An investor scorecard (Zoo Capital EIR, March 23, 2026) lists "~$500K seed (Redpoint Ventures scout, Samsung NEXT)" with 57K stars and 150M+ downloads at that date [secondary — https://github.com/lucy-cxy/agentvc-index/blob/main/cases/2026-03-23_unsloth.md]. Crunchbase lists 8 investors incl. **Pioneer Fund and Samsung NEXT**, headcount 11–50 [secondary]. PitchBook lists 12 investors: **Lightspeed Venture Partners, Transpose Platform Management, Y Combinator, Brightwing Capital** and others [secondary].
- **No confirmed large seed/Series A announced** as of Sep 22, 2026. Some older dev.to commentary claimed "no seed round yet" (that piece is from ~Feb 2025; stale). Treat any "raised $X" figure beyond the ~$500K pre-seed as **[unverified]** — Crunchbase/PitchBook figures are paywall-obfuscated in public view.
- Notable: an unusually small raise relative to adoption (~500K→investor ratio is widely commented on). ⚠️ "Official fine-tuning partner" claims (Google/OpenAI/Meta/NVIDIA) appear in the investor scorecard but lack official partnership press releases — **[unverified]**.
- Community adoption snapshot: the skill-card survey (Sep 17, 2026) puts Unsloth at **34k+ GitHub stars** (repo `unslothai/unsloth`); March 2026 scorecard: 57K stars / 150M+ downloads — ⚠️ inconsistent counts across sources; star count is a moving target [secondary].

### 1.2 Product line (2026)

Unsloth now ships three surfaces [official — unsloth.ai docs + staging READMEs]:
- **Unsloth Core** — the Python library (`pip install unsloth`), code-first fine-tuning.
- **Unsloth Studio** — web UI: dataset building ("Data Recipes" from PDFs/CSVs/DOCX), export to GGUF/NVFP4/FP8, RL configs [official].
- **Unsloth Desktop** — local desktop app (Mac/Windows/Linux), 100% offline-capable, GGUF/Safetensors; one-line installer `curl -fsSL https://unsloth.ai/install.sh | sh` [official/vendor-reported].
- **Unsloth Start** (2026): connects coding agents to local models with one command: `unsloth start claude --model unsloth/Qwen3.8-27B-GGUF:UD-Q4_K_XL`; supports Claude Code, OpenAI Codex, DeepSeek Harness (`dsh`), Hermes Agent, OpenCode, OpenClaw [official — README, Sep 2026].
- Docker image: `hub.docker.com/r/unsloth/unsloth` [official].

### 1.3 Speed / VRAM claims (vendor vs independent)

Vendor claims (all `[vendor-reported]` unless noted):
- "2–5x faster training with 50–80% less VRAM, zero accuracy loss" via custom fused Triton kernels + manual backprop derivation [vendor-reported].
- **MoE LLMs: "12x faster with 35% less VRAM"** — DeepSeek, GLM, Qwen, gpt-oss [vendor-reported — https://unsloth.ai/docs/new/faster-moe].
- **New RoPE & MLP Triton kernels + padding-free packing: "3x faster training, 30% less VRAM"** [vendor-reported — https://unsloth.ai/docs/new/3x-faster-training-packing].
- **7x longer-context RL** vs other setups via new batching algorithms [vendor-reported — https://unsloth.ai/docs/new/grpo-long-context].
- **500K context**: fine-tuning a 20B model with >500K context on an 80GB GPU [vendor-reported — https://unsloth.ai/docs/blog/500k-context-length-fine-tuning].
- **Embedding models: ~1.8–3.3x faster embedding fine-tuning** [vendor-reported].
- Qwen3.6: MTP for 1.4–2.2x faster inference; NVFP4 quants for supported GPUs [vendor-reported].
- Dynamic 2.0 4-bit quantization: claims record 5-shot MMLU & Aider Polyglot scores vs other quants [vendor-reported, single-vendor benchmarks].
- Independent verification: **not found** in this research window — no published third-party benchmark reproduced Unsloth's 2–5x/70% figures against transformers+FA2 baselines. A community voice-fine-tuning study (Ryan Baumann, Aug 17, 2026) is building such a comparison (Gemma 4 26B-A4B on Unsloth vs `mlx_lm.lora` on Apple Silicon) but had not published final numbers at research time [independent-in-progress]. → **flag: speed/VRAM claims remain vendor-only; cite with caution in RAG.**

### 1.4 2026 feature releases (chronological, from official docs/news)

- **GRPO / RL stack**: rule-based GRPO (no separate critic), **Dr. GRPO, DAPO, BNPO, GSPO** variants; **FP8 reinforcement learning** on consumer GPUs; **Vision RL (VLM GRPO)**; long-context RL (7x claim) [official docs].
- **TTS**: text-to-speech fine-tuning incl. `sesame/csm-1b`, STT `openai/whisper-large-v3`; 2026 model support extended to MiniMax-Music3 / Higgs / MOSS audio models [official]. ⚠️ Per-community docs, Whisper support is for inference/adaptation recipes, not full TTS training — exact coverage not independently verified.
- **Vision fine-tuning**: Qwen3-VL, Gemma 3, Llama 3.2 Vision recipes; VLM GRPO [official].
- **500K-context blog** (2026): 20B model at >500K ctx on 80GB [official].
- **AMD ROCm**: formal AMD GPU support added **2026-07-20** [secondary — community SKILL.md; official docs page not directly fetched — verify].
- **DoRA** lightweight fine-tuning method support [secondary].
- **GGUF Dynamic v3.0** quants [secondary].
- **Auto compaction** of long chats in Desktop; remote/LAN access preview with QR; MCP support; Projects; Deep Research Mode; Parallel Chat [secondary — community skill cards; not all independently confirmed].
- **Model support added in 2026** (from official README news): Qwen3.5 (0.8B–112B-A10B), Qwen3.6 (+MTP, NVFP4), Qwen3.8-27B, Qwen3.8-Flash-Next (125B), DiffusionGemma, GLM-5.2, GLM-5.3-Flash, DeepSeek-V4-0731, Kimi K3 (2.8T total / 104B active MoE, 1M ctx), Gemma 4 12B (+QAT), Muse Glimmer (Meta), MiniMax audio models, gpt-oss RL + flex-attention long-context [official].
- **Apple Silicon / MLX**: one community card says "MLX (Apple Silicon) — not yet supported" while another (later, Chinese) lists "MLX optimized support" — **conflicting; [unverified]** whether native MLX fine-tuning shipped in 2026.

### 1.5 Licensing
- Core library: **Apache 2.0** (per investor scorecard) [secondary]. Unsloth's pre-quantized model uploads on Hugging Face carry per-model licenses; commercial products: **Unsloth Pro** (multi-GPU, enhanced speed) and **Unsloth Studio/Desktop** (desktop app) — pricing not published in fetched sources [secondary/unverified]. ⚠️ Confirm current license text in repo LICENSE before legal use.

---

