---
id: open-local-models-2026/12-part-10-cross-cutting-observations-for-the-rag/overview
title: "PART 10 — CROSS-CUTTING OBSERVATIONS (for the RAG)"
domain: part-10-cross-cutting-observations-for-the-rag
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Google", "Meta", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Poolside", "United States", "Xiaomi", "Z.ai"]
dates: []
keywords: ["agent", "agentic", "apache", "benchmarks", "blackwell", "claude", "compute", "deepseek", "gemini", "glm", "gpu", "inference"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [704, 740]
section: "PART 10 — CROSS-CUTTING OBSERVATIONS (for the RAG)"
sha256: 49cb3f35645131ea720fad6535b699727189df049c5830f6589ffb45f882861e
---

# PART 10 — CROSS-CUTTING OBSERVATIONS (for the RAG)

## 10.1 Meta vs Chinese open-weights (2026)

The open-weight frontier in 2026 is **Chinese-led**, and Meta is no longer in the flagship race:

- **Top open-weights on AA Index:** MiMo-V2.6-Pro (Xiaomi, **46**), GLM-5.3 (Z.ai, **45–60** depending on index version), Qwen3.8 Max (Alibaba, **45**), Kimi K3 (Moonshot, **44–60**) — all shipping **MIT-licensed** (fully permissive) weights. MiMo-V2.6-Pro: 1.02T MoE / 42B active, 1M context, $0.43/$0.87. DeepSeek V4.1 Flash: 552B MoE, 1M context, MIT, **$0.15/$0.60 off-peak**.
- **Chinese open flagships match closed models:** Kimi K3 scores 60 on the AA Index — 3 points off the top closed model, "the narrowest open-vs-closed gap since February."
- **Meta's position:** Llama 4 (Apr 2025) is its last large open release — a generation behind. Meta's 2026 open contribution is **Glimmer (30B local-agent tier)**, not a flagship; its frontier effort is the **closed Muse Spark line**, which competes credibly on health/visual reasoning and price ($1.25/$4.25 undercuts Opus/GPT tiers) but trails Claude Opus 5, GPT-5.x, and Gemini 3.x on coding and raw intelligence.
- **Licensing contrast:** Chinese labs now ship MIT (Xiaomi, DeepSeek, Z.ai) while Meta's historical position was the restrictive Community License; Glimmer's Apache 2.0 is Meta catching up to the permissive norm.

## 10.2 License liberalization is the 2026 Western story

Gemma moved Terms-of-Use → Apache 2.0 (Apr 2026); NVIDIA's Nemotron 3 ships weights + data + recipes; Poolside's Laguna S 2.1 introduced OpenMDW-1.1 (Linux Foundation-backed); Meta's Glimmer went Apache 2.0. Counter-notes: Mistral introduced a *Modified MIT* for Medium 3.5 (watch this); Mistral's Voxtral TTS is CC BY-NC 4.0 (non-commercial).

## 10.3 Architecture convergence

Hybrid Mamba-Transformer MoE (Nemotron 3, Granite 4.0 H), latent/sparse MoE with ~10% active params (Mistral Large 3, Nemotron 3, Gemma 4 26B-A4B, Poolside Laguna S 2.1), MTP/speculative-decoding heads (Gemma 4, Nemotron 3), 1M-token context as the agentic standard (Nemotron 3, Muse Spark, Laguna S 2.1), 256K+ elsewhere.

## 10.4 US open-weight leadership (AA Intelligence Index, June–Sept 2026)

Nemotron 3 Ultra 48 > Gemma 4 31B 39 > Nemotron 3 Super 36 > gpt-oss-120b 33 — still trailing China's Kimi K2.6 (54) and DeepSeek V4 Pro (~56).

## 10.5 Agentic benchmarks replaced chat benchmarks

SWE-Bench Verified, Terminal-Bench 2.x, τ-bench V3, PinchBench, τ²/τ³-bench, RULER, IFBench dominate vendor reporting in 2026.

## 10.6 Business-model patterns

- **Open weights + paid platform** (Mistral: La Plateforme/Forge/Compute; Poolside pre-NVIDIA: self-hosted enterprise).
- **Proprietary API pivot** (Meta Muse Spark: first paid Meta Model API, Contributor tier trading price for training rights).
- **GPU-vendor open stack** (NVIDIA Nemotron: weights+data+recipes as an on-ramp to NIM/Blackwell inference).
- **Execuhire / licensing deals** as the 2026 exit path (NVIDIA × Poolside: $6B Model Factory license + $1B investment, staff hired).
- **Sovereignty framing** (Mistral's European positioning; Zuckerberg's open-weights-as-US-competitiveness essay; Apertus EU-AI-Act compliance).

---

