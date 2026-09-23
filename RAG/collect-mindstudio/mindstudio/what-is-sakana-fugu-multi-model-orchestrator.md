---
id: collect-mindstudio/mindstudio/what-is-sakana-fugu-multi-model-orchestrator
title: "What Is Sakana Fugu? The Multi-Model Orchestrator Explained"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Google", "OpenAI", "Sakana"]
dates: ["2026-06", "2026-09-23"]
keywords: ["fugu", "sakana", "attention", "benchmark", "benchmarks", "claude", "compute", "consumer", "cost", "distribution", "gemini", "inference"]
source: docs/RAG/Collect RAG/02_mindstudio/what-is-sakana-fugu-multi-model-orchestrator.md
source_anchor: ""
source_lines: [1, 56]
sha256: ac3cbb47234ce42d27d3b46a7d0eeffb714a613008a15db1816e3c18c00abc4a
---

# What Is Sakana Fugu? The Multi-Model Orchestrator Explained

## Metadata

- **Source**: https://www.mindstudio.ai/blog/what-is-sakana-fugu-multi-model-orchestrator
- **Site**: MindStudio
- **Type**: Article
- **Language**: en
- **Verification status**: ✅ reachable
- **Collection date**: 2026-09-23

## Full summary

This article explains **Sakana Fugu**, a multi-model LLM orchestration system from **Sakana AI** that automatically routes each incoming prompt to the most appropriate language model. The core problem it solves: no single model wins at everything — GPT-4o crushes complex reasoning but burns tokens fast, while a smaller model handles simple classification cheaply but fumbles nuanced tasks. Rather than picking one model or building fragile custom routing logic, Fugu uses a learned router.

**Sakana AI background.** Tokyo-based research company founded in 2023 by Llion Jones (one of the original co-authors of the "Attention Is All You Need" Transformer paper) and David Ha (former Head of Research at Google Brain). "Sakana" means "fish" in Japanese; model names follow aquatic themes (Fugu = puffer fish). Rather than competing on raw parameter count, Sakana AI pursues evolutionary model merging and adaptive routing — extracting more from existing compute rather than scaling up.

**How Fugu works.** Fugu is a learned router sitting in front of a pool of language models. The key insight: most queries don't need the most powerful model. A simple summarization, factual lookup, or basic classification can be handled by a smaller, cheaper model; expensive models are reserved for genuinely hard problems. The system uses a **two-tier architecture**: Tier 1 (fast lane) — a lightweight, low-cost model for queries within its capability, optimized for speed and cost; Tier 2 (capable lane) — a more powerful model for complex reasoning, long-context understanding, and nuanced generation. The router is a trained classifier taking the prompt as input and outputting a routing decision.

**How the router learns.** It isn't a keyword filter or prompt-length heuristic — it's trained on actual model outputs: running a large set of queries through both tiers, labeling which the Tier 1 model handled adequately versus fell short on, and training a lightweight classifier to predict those outcome labels. The classifier picks up signals including task-type signals (multi-step reasoning vs single-step retrieval), linguistic complexity (sentence structure, domain vocabulary, ambiguity), output-requirement signals (structured output, code, long-form → higher tier), and context length. Routing adds minimal latency (typically milliseconds) and no extra inference cost.

**Benchmark results.** Fugu approaches the performance of exclusively using the top-tier model while running a substantial share of queries (often the majority) on the cheaper Tier 1 model. The cost-performance tradeoff curve: at the high-accuracy end, Fugu matches/approaches full Tier 2 performance while still routing 20–40% of queries to Tier 1; at the cost-optimized end, it routes 70–80% of queries to Tier 1 with acceptable accuracy loss; the sweet spot is the middle. Compared to naive baselines, Fugu: matches or approaches "always Tier 2" at lower average cost; significantly outperforms "always Tier 1" on complex query subsets; and consistently outperforms random routing at every cost target.

**Why multi-model orchestration matters in production.** Token economics at scale (a 40% average cost-per-query reduction can mean tens of thousands of dollars monthly for large deployments); latency variability (Tier 1 models are faster, smoothing response-time distribution); model specialization (different models best for different tasks, hidden behind the routing layer); and future-proofing (swap in new models without rewriting application logic).

**Limitations.** Router accuracy isn't perfect (individual query routing is fallible); latency overhead (small but nonzero); domain dependency (a router trained on general benchmarks may not generalize to niche domains like legal or medical); maintenance burden (calibration may drift as models update, requiring retraining with fresh data).

**Alternatives.** **RouteLLM** (LMSys): open-source routing framework using human preference data to train classifiers. **Martian** and **Unify**: commercial routing services across providers. **Manual step-level routing**: workflow builders like MindStudio where different steps use different models based on task type — no classifier, no automation, more transparency.

**When to use Fugu-style orchestration.** High-volume inference workloads (tens of thousands of queries/day); mixed-complexity query distributions; cost-sensitive products (consumer products, per-query API pricing); acceptable accuracy tradeoffs; stable query distributions.

## Key points

- Sakana Fugu is a multi-model LLM orchestrator routing each prompt to the best model tier using a learned classifier.
- Two-tier design: Tier 1 (fast, cheap) and Tier 2 (capable, expensive), with automatic escalation based on predicted difficulty.
- The router is trained on empirical model outputs, not proxy heuristics — it learns prompt signals that predict when the cheap model suffices.
- Benchmark results: approaches top-tier accuracy while routing 20–80% of queries to Tier 1 depending on the threshold.
- Outperforms naive baselines (always Tier 2, always Tier 1, random routing) at every cost target.
- Benefits: cost reduction at scale, lower latency, model specialization, future-proofing.
- Limitations: imperfect routing accuracy, latency overhead, domain dependency, maintenance/calibration drift.
- Alternatives: RouteLLM (open source), Martian, Unify (commercial), and manual step-level model assignment.

## Technical data / figures

- Founded: 2023, Tokyo; founders Llion Jones ("Attention Is All You Need" co-author) and David Ha (ex-Google Brain).
- Architecture: two-tier — Tier 1 (lightweight/low-cost, speed-optimized) + Tier 2 (powerful, complex reasoning).
- Router: trained classifier; latency typically milliseconds; no extra inference cost.
- Routing signals: task type, linguistic complexity, output requirements, context length.
- Benchmark tradeoff curve: high-accuracy end routes 20–40% to Tier 1 (≈full Tier 2 accuracy); cost-optimized end routes 70–80% to Tier 1 (acceptable accuracy loss).
- Baselines outperformed: always Tier 2, always Tier 1, random routing.
- Model examples referenced: GPT-4o, Claude, Gemini.
- Alternatives: RouteLLM (LMSys, open source), Martian, Unify (commercial).

## Why this source matters for the RAG

Provides detailed, current (June 2026) technical coverage of Sakana Fugu — a specific multi-model routing system with named architecture, two-tier design, training method, benchmark figures, and alternatives. This is precise, up-to-date knowledge that lets the RAG answer accurately about model routing, orchestration, and cost-performance tradeoffs without hallucinating on this recent system.
