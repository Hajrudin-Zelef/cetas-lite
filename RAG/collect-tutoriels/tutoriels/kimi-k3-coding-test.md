---
id: collect-tutoriels/tutoriels/kimi-k3-coding-test
title: "Kimi K3 Coding Test: A Reproducible Repository and Agent Evaluation"
domain: tutoriels
role: reference
task: tutorial
actors: ["MiniMax", "Moonshot", "OpenAI"]
dates: ["2026-07", "2026-09-23"]
keywords: ["agent", "kimi", "agents", "benchmark", "cost", "gpt-5.6", "pricing", "reasoning", "research", "sol", "tool use"]
source: docs/RAG/Collect RAG/07_tutoriels/kimi-k3-coding-test.md
source_anchor: ""
source_lines: [1, 75]
sha256: d244e352abb44ec6be3324c705720363f715eeca2c7cdc94d927211413a71755
---

# Kimi K3 Coding Test: A Reproducible Repository and Agent Evaluation

## Metadata

- **Source** : https://poyo.ai/hub/kimi-k3-coding-test
- **Site** : Poyo
- **Type** : Tutorial
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Poyo.ai team article (published 21 July 2026, ~10 min read) provides a reproducible test plan for evaluating Kimi K3 as a coding agent, rather than pretending that a single unpublished run is a universal verdict. Kimi K3 is marketed for long-horizon coding, terminal orchestration, visual software development, and large repositories, so the authors argue a one-prompt code demo cannot test those claims. A useful evaluation must give the model real state, allow mistakes, require verification, and measure whether it finishes within scope. The plan is designed so K3 can be compared with another model (specifically GPT-5.6 Sol) under the same harness, tools, repository commit, time budget, and scoring rules. No fabricated scores are presented.

The article first defines the dimensions a production coding agent must be measured on: correctness, repository understanding, persistence, tool use, scope control, verification, visual reasoning, safety, and efficiency. It then lists the environment fields that must be published before scores: model ID and provider, evaluation date, reasoning effort, agent harness and version, OS and hardware, repository URL and exact commit, available tools, network permissions, time/turn limits, context-compaction policy, maximum completion setting, number of runs per task, human interventions, and token/cost accounting method. A key K3-specific note is that it requires complete assistant state in later turns; discarding required history means the harness, not the model, is being tested.

A scoring rubric is defined with six weighted categories: functional correctness 35%, verification quality 20%, scope control 15%, tool use and recovery 15%, code quality 10%, and efficiency 5%. Severe failures (deleting unrelated data, exposing credentials, falsely claiming tests passed, changing security boundaries without authorization) are defined separately and should not be averaged away by code style.

The plan then specifies eight tests: (1) large-repository navigation (read-only trace of cross-directory behavior), (2) multi-file bug fix with hidden traps, (3) terminal-agent recovery from controlled failures, (4) screenshot-to-frontend iteration, (5) small playable game, (6) research-to-code workflow, (7) tool-failure recovery, and (8) long-session state preservation. Each test includes the task, score criteria, and hidden traps. Test 8 is explicitly required by K3's documentation: run a correct client returning complete assistant messages versus an intentionally incomplete client retaining only final content, to quantify the integration failure Moonshot describes.

Finally, the article explains measuring cost per verified success (total API spend across all attempts divided by verified successes), fair comparison rules (same task text, repo state, tool permissions, time limit, verification; at least three runs; report median and worst-case), a results table template, ten common evaluation mistakes, what would count as a strong K3 result, and an FAQ. A strong result is finishing correct tasks with bounded tools, preserving user changes, reporting failures honestly, and using long context without unnecessary cost.

## Key points

- The article is a reproducible test methodology, not a benchmark verdict; it publishes no fabricated scores.
- Nine evaluation dimensions defined, from correctness to efficiency.
- Scoring rubric: correctness 35%, verification 20%, scope 15%, tools 15%, quality 10%, efficiency 5%.
- Eight concrete tests, including long-session state preservation, required by K3's docs.
- K3 requires complete assistant messages in multi-turn/tool workflows; dropping history destabilizes performance.
- Cost is measured as total API spend / verified successes, not per-token price alone.
- Three runs minimum; report median and worst case; severe failures tracked separately.
- Comparing cache-hit cost with cache-miss cost is listed as a common mistake.

## Technical data / figures

Scoring rubric:

| Category | Weight |
| --- | --- |
| Functional correctness | 35% |
| Verification quality | 20% |
| Scope control | 15% |
| Tool use and recovery | 15% |
| Code quality | 10% |
| Efficiency | 5% |

The eight tests:

1. Large-repository navigation
2. Multi-file bug fix
3. Terminal-agent recovery
4. Screenshot-to-frontend iteration
5. Small playable game
6. Research-to-code workflow
7. Tool-failure recovery
8. Long-session state preservation

Cost formula:
```
verified-success cost =
  total API spend across all attempts / verified successes
```

Reference pricing shown on the page:

| Model | Provider | Price |
| --- | --- | --- |
| Kimi K3 | Moonshot AI | $2.28 / 1M input tokens |
| GPT-5.6 | OpenAI | $0.056 / 1M input tokens |
| MiniMax H3 / Hailuo 03 | MiniMax | $0.105 / second |

## Why this source matters for the RAG

It supplies a rigorous, reusable evaluation framework and rubric for coding agents, with K3-specific integration warnings (complete assistant state, long context) that are essential for anyone benchmarking or deploying Kimi K3. It is valuable for RAG queries about model evaluation methodology, scoring, and fair comparisons.
