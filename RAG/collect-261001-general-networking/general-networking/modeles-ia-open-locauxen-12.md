---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-12
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Google", "Meta", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Poolside", "United States", "Xiaomi", "Z.ai"]
dates: ["2025-09-30", "2026-09"]
keywords: ["agent", "agentic", "apache", "benchmark", "benchmarks", "blackwell", "capex", "claude", "compute", "deepseek", "gemini", "glm"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [710, 755]
sha256: 94be963bd8eac98b3047041bbb1c185ce8308d8849861d24d7a42f6835be3fe4
---

# ÉTAPE 1 — Open / Local AI Models (EN)

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

# PART 11 — CAVEATS & METHODOLOGY

- Benchmark figures are vendor-reported unless marked as independently measured (Artificial Analysis Intelligence Index, vals.ai, LMArena, third-party tables).
- Harness variance is large: Nemotron 3 Ultra SWE-Bench spans 65–70.4% across five harnesses; Devstral 2 figures conflict across sources (46.8% vs ~72%); Muse Spark 1.0's AA Index is cited as 52 / 43 / 31 across outlets — record methodology version.
- API prices move frequently; all pricing figures are June–September 2026 snapshots from the cited aggregators (llm-stats.com, Artificial Analysis, vendor docs), not live quotes.
- Muse Spark parameter counts and architecture internals are undisclosed by Meta — every figure online is inference or error. "Open weights" claims for the Spark line (Memeburn on 1.2; a French PDF on 1.1) are uncorroborated — treat as unverified.
- The Llama topic attracts fabricated "2026" aggregator content (e.g., a nonexistent "Llama 5" model page). **No Llama 5 exists as of Sept 2026.**
- Poolside's post-NVIDIA pivot direction is undisclosed as of Sept 2026; no August/September 2026 model releases found — Laguna S 2.1 (July 21) remains the current flagship.
- Meta's $115–135B AI capex figure rests on a single weak source; treat as indicative.
- Nemotron 3 Ultra's documented knowledge cutoff is 2025-09-30; exact Mistral Large 3 knowledge cutoff is undisclosed.

---

# PART 12 — KEY SOURCES

