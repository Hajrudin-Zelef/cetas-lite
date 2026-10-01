---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-1
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Intel", "Moonshot", "Nvidia", "SGLang", "vLLM"]
dates: ["2023-02-09", "2026-01-20", "2026-01-24", "2026-01-29", "2026-02-01", "2026-02-04", "2026-02-25", "2026-03-07", "2026-03-11", "2026-03-20", "2026-03-31", "2026-04-03", "2026-04-18", "2026-04-27", "2026-05-04", "2026-05-10", "2026-05-15", "2026-05-29", "2026-06-05", "2026-06-15", "2026-06-29", "2026-07-11", "2026-07-14", "2026-07-27", "2026-08-10", "2026-08-11", "2026-08-26", "2026-09-09", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "amd", "apache", "benchmarks", "deepseek", "intel", "kimi", "kv cache", "license", "memory"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [1, 99]
sha256: 7df3937f4255f2cbfd4ff1e424657ace49452e13e3994d0b5c4c01bc511c36ff
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)
## Research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 4 (Infra inference / training), Track A
**Coverage window:** February 1, 2026 → September 22, 2026
**Engines:** vLLM (vllm-project) and SGLang (sgl-project)

This file combines two full research reports produced as independent passes:

- **Part 1 — vLLM** (`_draft_etape4_vllm.md`): 609 lines. Research date 2026-09-22.
- **Part 2 — SGLang** (`_draft_etape4_sglang.md`): 394 lines. Research date 2026-09-22.
- **Part 3 — Synthesis** (head-to-head comparison, 2026 combined timeline, shared ecosystem,
  unified uncertainty log) — added below, after Part 2.

### Provenance legend (used throughout both parts)
- **[official]** — the project itself: GitHub repo, GitHub releases API, official docs, project blog.
- **[vendor-reported]** — hardware/software vendor claims (NVIDIA, AMD, Intel, Red Hat, cloud providers).
- **[independent]** — reputable third-party measurements (SemiAnalysis InferenceX, dstack, community benchmarks).
- **[secondary]** — press, blogs, third-party summaries, aggregator repos.
- **[unverified]** — could not be confirmed from an authoritative source; treat as uncertain.

Prices, versions, and star counts are dated snapshots of 2026-09-22; re-verify before operational use.

---

# PART 1 — vLLM

# vLLM Inference-Serving Framework — Research Report
**Research date:** 2026-09-22
**Project:** RAG data collection, Step 4 (AI infra: inference/training)
**Language:** English

> Provenance tags used throughout: `[official]` (vLLM project itself: GitHub, docs, blog),
> `[vendor-reported]` (NVIDIA/AMD/Intel/Red Hat/vendor sources),
> `[independent]` (independent benchmarks/studies),
> `[secondary]` (press, blogs, third-party summaries),
> `[unverified]` (could not be confirmed from an authoritative source).
> Uncertainties are marked explicitly with ⚠️.

---

## 1. Identity & origins

- **What:** vLLM is an open-source, high-throughput LLM inference and serving engine built around
  **PagedAttention** (KV cache organized into fixed-size physical blocks addressed via per-sequence
  block tables, OS-virtual-memory-style) and **continuous (iteration-level) batching**. [official][secondary]
- **Origin:** Created at UC Berkeley's **Sky Computing Lab** (SOSP 2023 paper: "Efficient Memory Management
  for Large Language Model Serving with PagedAttention"); associated with the **LMSYS** team (Chatbot Arena).
  [official][secondary]
- **License:** Apache-2.0, no CLA requirement, no relicense history to date. [secondary]
- **GitHub:** https://github.com/vllm-project/vllm [official]
  - **Stars: 92,444 — Forks: 22,525 — Open issues: 8,335** (GitHub API, read 2026-09-22 19:4x UTC;
    pushed_at 2026-09-22T19:39:14Z; repo created 2023-02-09). [official]
- ⚠️ Ecosystem-scale claim: listed as the backend for many inference-as-a-service platforms and the de-facto
  open-source standard for LLM serving (widely repeated; directionally consistent with star/fork data, but a
  "de-facto standard" is a characterization, not a measurable fact). [secondary]

---

## 2. Release history 2026 (all dates from GitHub Releases API, 2026-09-22)

Release cadence in 2026 is roughly bi-weekly majors plus occasional patches. Full official list
(version — published date) [official]:

| Version | Published | Notes |
|---|---|---|
| **v0.30.0** | 2026-09-22 | Latest as of research date |
| v0.29.0 | 2026-09-09 | MRV2 default for all models |
| v0.28.0 | 2026-08-26 | Kimi-K3 perf push; DeepSeek V4 sparse MLA |
| v0.27.1 | 2026-08-11 | patch |
| v0.27.0 | 2026-08-10 | |
| v0.26.0 | 2026-07-27 | Inkling family; DeepSeek-V4 perf push |
| v0.25.1 | 2026-07-14 | patch |
| v0.25.0 | 2026-07-11 | |
| v0.24.0 | 2026-06-29 | |
| v0.23.0 | 2026-06-15 | |
| v0.22.1 | 2026-06-05 | |
| v0.22.0 | 2026-05-29 | |
| v0.21.0 | 2026-05-15 | |
| v0.20.2 | 2026-05-10 | patch |
| v0.20.1 | 2026-05-04 | patch |
| v0.20.0 | 2026-04-27 | |
| v0.19.1 | 2026-04-18 | patch |
| v0.19.0 | 2026-04-03 | |
| v0.18.1 | 2026-03-31 | patch |
| v0.18.0 | 2026-03-20 | |
| v0.17.1 | 2026-03-11 | patch |
| v0.17.0 | 2026-03-07 | |
| v0.16.0 | 2026-02-25 | |
| v0.15.1 | 2026-02-04 | patch |
| v0.15.0 | 2026-01-29 | |
| v0.14.1 | 2026-01-24 | patch |
| v0.14.0 | 2026-01-20 | First 2026 release |

Source: GitHub Releases API `https://api.github.com/repos/vllm-project/vllm/releases` (read 2026-09-22).
Release-notes pages: https://github.com/vllm-project/vllm/releases/tag/v0.30.0 etc. [official]

### 2.1 v0.30.0 — 2026-09-22 (762 commits, 315 contributors, 104 new) [official]

