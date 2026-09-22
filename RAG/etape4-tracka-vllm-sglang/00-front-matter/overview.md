---
id: etape4-tracka-vllm-sglang/00-front-matter/overview
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-02-01", "2026-09-22"]
keywords: ["inference", "sglang", "vllm", "amd", "benchmarks", "intel", "nvidia", "research", "training"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [1, 25]
section: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
sha256: 0af4fe353327b877ad5101335b6ac83581fb753f7c0e65b4b94dc3d63d91bcdb
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

