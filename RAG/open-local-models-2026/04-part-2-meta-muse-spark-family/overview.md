---
id: open-local-models-2026/04-part-2-meta-muse-spark-family/overview
title: "PART 2 — META MUSE SPARK FAMILY"
domain: part-2-meta-muse-spark-family
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Apple", "Google", "Hugging Face", "Meta", "OpenAI", "OpenRouter", "United States", "vLLM", "xAI"]
dates: ["2026-04-08", "2026-07-09", "2026-08-05", "2026-08-10", "2026-08-31", "2026-09-02", "2026-09-04", "2026-09-08"]
keywords: ["muse", "muse spark", "agent", "agentic", "agi", "apache", "benchmark", "benchmarks", "claude", "consumer", "context window", "cost"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [368, 460]
section: "PART 2 — META MUSE SPARK FAMILY"
sha256: b905a1ec1763c63c1a96990330b7ee00603e9614e689f6de8d8c0b16b88b2f0d
---

# PART 2 — META MUSE SPARK FAMILY

*(Merged from two independent research passes; discrepancies across outlets are flagged inline rather than smoothed over.)*

## 2.1 What it is

**Muse Spark** is Meta's proprietary frontier reasoning-model family and the **first product of Meta Superintelligence Labs (MSL)**, the division led by Chief AI Officer **Alexandr Wang** (joined Meta from Scale AI in 2025; Meta invested ~$14–14.3B in Scale AI to bring him over). Debuted **April 8, 2026**. Internally codenamed **"Avocado"**.

It is a **natively multimodal reasoning model**: text, image, video, PDF, audio in; text out. Built around long-horizon agentic work: tool use, computer use, MCP servers, custom skills, multi-agent orchestration (main agent delegating to parallel subagents, or running as a subagent). Meta says it was built after a **9-month rebuild of its AI stack "from scratch"** (new infrastructure, new architecture, new data pipelines).

## 2.2 Release timeline (through Sept 2026)

| Version | Date | Notable |
|---|---|---|
| **Muse Spark 1.0** | 2026-04-08 | First MSL model; Meta AI app/meta.ai free tier; API invite-only |
| **Muse Spark 1.1** | 2026-07-09 | First **paid Meta Model API** (public preview, US developers at launch); 1M context; Thinking mode in Meta AI app |
| **Muse Spark 1.2** | 2026-08-05 | Co-released with **Muse Code** terminal coding agent (co-trained with it); global public-preview API expansion; OpenRouter listing; API out of beta 2026-08-31 (subscription option added) |
| **Muse Glimmer 30B** (sibling, open weights) | ~2026-08-10 | 30B dense local agentic model, Apache 2.0, HF: `meta-models/Muse-Glimmer-30B` |
| **Muse Spark 1.3** | 2026-09-02 | Coding/long-context flagship; two variants — `muse-spark-1.3` (xhigh, public), `muse-spark-1.3-max` (limited partner preview); reasoning_effort up to "max" (live 2026-09-04); paid API opened Sept 3 |
| **"Muse" consumer app** | 2026-09-08 | Personal agent app running Muse Spark 1.3 on per-user Meta-hosted VMs; free / $20/mo / $100/mo tiers |

**Meta Model API** (Meta's first metered inference API) public + paid since **July 9, 2026**.

## 2.3 Specs

- **Parameters:** Not disclosed by Meta for any Muse Spark version (Artificial Analysis FAQ: "Meta has not disclosed the model size or parameter count"). Every figure online is inference or error.
- **Architecture:** Described by third parties as a dense multimodal foundation model / hybrid reasoning transformer.
- **Context window:** **1M tokens** for 1.1/1.2/1.3 (standard figure). *Discrepancies:* Artificial Analysis lists the original Muse Spark at **260K tokens**; one dev.to comparison table lists 1.2 at 128K (131,072) — minority report, likely error.
- **Reasoning control:** `reasoning_effort` parameter from `minimal` to `xhigh` (1.3 adds `max`); reasoning tokens billed as output. Consumer app modes: **Instant, Thinking, Contemplating** (Contemplating = multi-agent parallel reasoning).
- **License:** **Proprietary, closed weights, API-only** — Meta's first paid model API, a sharp break from the open-weight Llama tradition. *Discrepancies:* (a) one French consultant PDF claims 1.1 is open-weights under a "Llama 5 Community License" with HF deployment via vLLM/Ollama — not corroborated by any major outlet; (b) Memeburn reported Meta "will open the weights for Muse Spark 1.2," again not corroborated by Meta's own blog. Treat "open weights" claims for the Spark line as **unverified**.

## 2.4 Availability

- **Consumer surfaces:** Meta AI app and meta.ai (free); rolled out across **WhatsApp, Instagram, Facebook, Messenger**, and Ray-Ban Meta AI glasses.
- **Developer API:** Meta Model API, base URL `api.meta.ai/v1`, supports **OpenAI Chat Completions and Anthropic Messages formats**; model names `muse-spark-1.1`, `muse-spark-1.2`, `muse-spark-1.2-contributor`, `muse-spark-1.3`. Built-in web-search grounding (`{"type": "web_search"}` tool). New accounts get a **$20 one-time credit**. Zero-data-retention arrangements available on request. OpenRouter listing from 1.2. Early partners: Replit, Cline, Box.
- **Muse Code** (terminal coding agent, macOS/Linux beta, launched Aug 5, 2026): parallel subagents in isolated git worktrees; replay-safe local event log; default skills `/plan`, `/grill`, `/goal`. Competes with Claude Code/Codex; Wang frames the strategy as differentiating "by price rather than capabilities".
- **Hugging Face:** Muse Spark weights are **not** on HF (proprietary). Sibling **Muse Glimmer 30B** is open on HF under Apache 2.0 (see §2.7).

## 2.5 API pricing

| Item | Price |
|---|---|
| Input | **$1.25 / 1M tokens** |
| Output | **$4.25 / 1M tokens** (reasoning tokens billed as output) |
| Cached input | **$0.15 / 1M tokens** |
| Blended rate (7:2:1 cache/input/output) | ~$0.78 / 1M tokens (Artificial Analysis) |
| **Contributor tier** (`muse-spark-1.2-contributor` / `muse-spark-1.3-contributor`) | ~**$0.10 in / $0.20 out per 1M** ($0.002 cached input) — in exchange for Meta training rights on prompts/completions; rate-limited on rolling 5-hour token window |

Competitive position: undercuts GPT-5.6 Sol ($5/$30) and Claude Sonnet 5 ($2/$10), slightly above GPT-5.6 Luna ($1/$6) and Claude Haiku 4.5. Output runs ~30%+ faster than Claude Opus 5-class models; aggressive pricing aimed at high-volume agent workloads.

## 2.6 Benchmarks

**Artificial Analysis Intelligence Index (independent):**
- Muse Spark 1.0: **52** (top-5 globally at launch, behind GPT-5.4 57, Gemini 3.1 Pro 57, Claude Opus 4.6 53). *Inconsistencies flagged:* TechTimes later cited 43 for 1.0; AA's own FAQ page estimates 31.
- Muse Spark 1.1: **53**; cost/task ~$0.29.
- Muse Spark 1.2: **57** (v4.1.1).
- Muse Spark 1.3 (xhigh): **61**; 1.3 (max): **62** — behind only Claude Fable 5.1 (max, 66) and Claude Opus 5 (max, 63); tied with GPT-5.6 Sol (max) and Grok 4.6 (high). 1.3 gains come from agentic work + science. 1.3 (max): Tau3-Bench Banking **52%** (vs 47% xhigh), GDPval-AA v2 **1,754 Elo** (vs 1,709). Output speed 232.6 tok/s, TTFT 23.48s.

**1.0 launch-window results:** **HealthBench Hard 42.8%** — best of any frontier model at the time (vs GPT-5.4 40.1, Gemini 3.1 Pro 20.6); CharXiv Reasoning **86.4**; MMMU-Pro **80.5%**; Humanity's Last Exam **58%** (Contemplating mode) / 39.9% (third-party); weak spots: Terminal-Bench 2.0 **59.0** (vs GPT-5.4's 75.1), ARC-AGI-2 **42.5%** (vs Gemini 3.1 Pro 76.5%), GDPval-AA Elo **1427–1444** (vs GPT-5.4 1672). **SWE-bench Pro 55** — tied #1 with GPT-5.4 xHigh on the Scale leaderboard. Token-efficient: 58M output tokens on the AA index vs 157M (Claude Opus 4.6).

**1.2:** Terminal-Bench 2.1 **82.9%** (2nd to Claude Opus 5's 86.7); GDPval-AA v2 **1631 Elo** (#5 of all models tested, ahead of Claude Opus 4.8); SWE-bench Verified ~**54.2%** (third-party table); cost/task ~**$0.40** (among most cost-efficient at its level).

**1.3 (Meta-reported):** DeepSWE v1.1 **75.4%** (113 real-world tasks, 91 repos, 5 languages); Terminal-Bench 2.1 **88.8%** (tied with GPT-5.6 Sol); SWEAtlas CodeBase QnA **59.4%**; MRCR long-context retrieval **98.5%** (256K–512K) / **98.1%** (512K–1M); DeepSearchQA **89.4%** (vs GPT-5.6 93.0%); Agentic IF Index **57.8** (vs GPT-5.6 60.5). Efficiency: ~20% fewer tool calls, ~25% fewer tokens than 1.2. Trails Claude Opus 5 on JobBench, OSWorld 2.0, AutomationBench, GDPval-AA v2 in Meta's own table.

## 2.7 Muse Glimmer 30B — Meta returns to open weights (Apache 2.0)

**Release:** August 10, 2026.

| Attribute | Value |
|---|---|
| Architecture | **Dense 30B** (~29.6B), distilled from Muse Spark 1.2 via logit distillation; dedicated perception encoder; natively multimodal (text+image) |
| Context window | **120K+ tokens** (131K per one source) |
| License | **Apache 2.0** — commercial use, modification, redistribution; no MAU cutoff, no AUP |
| HF weights | `meta-models/Muse-Glimmer-30B` (separate `meta-models` org; +GGUF, -assistant, ExecuTorch-PTE variants) |
| Design goal | Agentic local model: fits **<20 GB** quantized on a single 24GB consumer GPU or Apple Silicon Mac; runs offline; **no first-party Meta API** |

**Benchmarks (Meta vendor-reported):** SWE-Bench Pro **51.2%**, AIME **94.7** (2026), GPQA **83.5**, SWE-Bench Verified **76%**, WildCodeBench **47.6%**; MCP Atlas **75.5–75.8** (sources vary slightly) vs Gemma4-31B 54.2 vs Qwen3.6-27B 62.5; OSWorld-Verified **65.9** vs Qwen3.6-27B 75.6. Bundled **DFlash** speculative drafter: RTX 5090 **74.9 → 233.4 tok/s** (3.1×); M5 Max 26.6 → 50.2 tok/s. vs Qwen3.6-27B: Qwen is the stronger pure coder (77.2% SWE-bench Verified); Glimmer wins on native multimodality and agent tooling.

## 2.8 Relationship to Llama

- Muse Spark is **not** an open-weights release and has **no migration path from Llama** ("fundamentally different deployment models"). Meta's stated rationale: Llama 4 (Apr 2025) failed commercially and its launch was marred by the LMArena benchmark scandal (acknowledged by Yann LeCun after his departure, Jan 2026).
- Alexandr Wang has promised to **open-source future (bigger) versions**; existing Llama models remain available but are in maintenance mode while frontier investment flows to Muse.
- One French source claims Muse Spark = instruction-tuned/RLHF version of a "Llama 5" base under a Llama 5 Community License — treat as unverified/single-source.

## 2.9 What distinguishes Muse Spark

- First paid Meta model API — strategic pivot from open weights to monetization under Wang/MSL; fast release cadence (1.0→1.3 in 5 months).
- Agentic-first design (context compaction over 1M tokens, multi-agent orchestration, computer use) with aggressive pricing aimed at high-volume agent workloads.
- Co-training of model + agent harness (Muse Spark 1.2 ↔ Muse Code).
- Strongest on: health/vision (1.0), long-horizon coding + long-context retrieval (1.3). Weakest on: abstract reasoning (ARC-AGI-2), open-ended agentic search vs OpenAI/Anthropic flagships.

---

