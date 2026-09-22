---
id: open-local-models-2026/08-part-6-microsoft-phi/overview
title: "PART 6 — MICROSOFT PHI"
domain: part-6-microsoft-phi
role: deep-dive
task: reference
actors: ["Alibaba", "Apple", "DeepSeek", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["agent", "benchmarks", "cost", "deepseek", "distribution", "foundry", "license", "llama", "llama.cpp", "mistral", "mit license", "multimodal"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [608, 618]
section: "PART 6 — MICROSOFT PHI"
sha256: 8b5bd30648c6e355fe5ce36c53cf71fcff64f7147ccb5568e8266f2f2d63c3be
---

# PART 6 — MICROSOFT PHI

**Status: incremental in 2026 — no Phi-5.** Two independent 2026 surveys confirm "Phi-5 unreleased" as of mid-2026.

- **Phi-4-reasoning-vision-15B** (Mar 2026) — current latest. Selective reasoning (model decides when to think), ~200B multimodal tokens; fits 12GB VRAM. MIT license.
- **Phi-4** (Dec 12, 2024, 14B dense, MIT); **Phi-4-mini** (3.8B); **Phi-4-multimodal** (5.6B, unified text+speech+vision). All MIT-licensed, HF + Ollama (`ollama pull phi-4`) + LM Studio.
- **Positioning as small local models:** Microsoft's SLM line is the reference "runs anywhere" family. **Foundry Local** reached GA (announced at Build 2026, milestone Apr 9): a ~20MB embeddable runtime (Windows/macOS Apple Silicon/Linux x64, no cloud, no per-token cost, OpenAI-compatible API, ONNX Runtime backend claiming 3.9× throughput over llama.cpp). Its curated GA catalog ships **Phi**, Qwen, DeepSeek, Mistral, Whisper — Phi is the first-party small-model default for on-device Windows/agent scenarios.
- **Note:** Phi-4-class models are now outclassed by newer small open models (e.g., Qwen3.5-9B, Gemma 4 E4B) on benchmarks; their value is MIT licensing + Microsoft distribution, not leadership.

---

