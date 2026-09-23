---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/10-experiment-tracking-and-evaluation-w-b-mlflow-lm-eval-ins
title: "10. Experiment tracking and evaluation: W&B, MLflow, lm-eval, Inspect"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: reference
actors: ["CoreWeave", "OpenAI", "vLLM"]
dates: ["2025-06", "2025-11", "2026-04", "2026-06-29"]
keywords: ["agent", "agentic", "agents", "benchmarks", "leaderboard", "pricing", "tool use", "vllm"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [662, 734]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: ea01f05087d3af6505570a76c8590f53a0591dbe9c56a66f5180ae9205ce52a1
---

# 10. Experiment tracking and evaluation: W&B, MLflow, lm-eval, Inspect

## 10. Experiment tracking and evaluation: W&B, MLflow, lm-eval, Inspect

### 10.1 Weights & Biases (W&B)

- Secondary sources agree **W&B became a CoreWeave company in 2025**
  [secondary].
- Reported transaction values conflict (~$1.4B vs ~$1.7B) — recorded as
  conflict C-04; the price is [secondary]/reported, not definitive.
- A 2026 interview reports **ARIA launched June 29, 2026 in public
  preview** (an AI-assistant/product surface inside W&B) [secondary].
- No verified 2026 platform version or pricing-tier changes captured —
  gap G-17 [unverified].
- W&B's role in the 2026 stack: experiment tracking, sweeps, model
  registry, and the Weave tracing layer for LLM/agent observability
  [secondary].

### 10.2 MLflow

- **MLflow 3.0 (June 2025)** redesigned the platform around GenAI and
  agent workflows; releases ran through **3.11 (April 2026)** deepening
  agent support across authoring, serving, tracing, evaluation, and
  gateway routing [secondary].
- `ResponsesAgent` is the recommended agent interface since 3.0
  (OpenAI Responses API schema), superseding `ChatModel`/`ChatAgent`;
  agents are logged, versioned, and served as `RegisteredModel`/`pyfunc`
  artifacts — there is no separate agent entity type [secondary].
- The experimental **Agent Server (3.6.0+)** is a FastAPI host for
  `ResponsesAgent` agents with decorator registration (`@invoke`,
  `@stream`), validation, tracing, and async streaming [secondary].
- **LoggedModel** and the Prompt Registry give GenAI apps auditable
  code+prompt+config+trace snapshots; tracing integrates with
  OpenTelemetry [secondary].
- MLflow 3.x is the open-source counterweight to W&B for teams that want
  self-hosted tracking with first-class LLM/agent support [secondary].

### 10.3 lm-eval (EleutherAI lm-evaluation-harness)

- The harness is the de-facto standard for academic LLM benchmarking
  (MMLU, GSM8K, HumanEval, HellaSwag, TruthfulQA, ARC, and 60+ more
  tasks), distributed as the `lm_eval` PyPI package [secondary].
- Release evidence: **v0.4.9.2** (~November 2025, Python 3.10 minimum,
  AIME/MATH500/LongBench-v2 wave of tasks) and **v0.4.11** (~February
  2026, Windows ML backend, BEAR probe) [secondary].
- The README announces the v0.4.0 line with config-based task creation,
  Jinja2 prompt design, vLLM and MPS backend support, and Open LLM
  Leaderboard task groups [secondary].
- The base install no longer bundles transformers/torch; extras like
  `lm_eval[hf]` / `lm_eval[vllm]` add backends [secondary].
- Task versioning matters: task config versions change across releases
  (e.g., `mgsm_direct` 3.0→4.0), so results must cite the harness
  version and task versions to be comparable [secondary].

### 10.4 Inspect (UK AI Safety Institute)

- Inspect is UK AISI's open-source framework for agentic/LLM evaluations,
  built around evals-as-code with solvers, scorers, and sandboxed tool
  use [secondary].
- No verified 2026 version captured in this pass — gap G-18 [unverified].
- It complements lm-eval: lm-eval covers academic benchmarks, Inspect
  covers agentic/safety evaluations with tool use [secondary].

### 10.5 Tracking/eval decision matrix

| Tool | Primary job | Hosting | Best fit |
|---|---|---|---|
| W&B (+Weave) | Experiment tracking, sweeps, LLM tracing | SaaS (CoreWeave) | Teams wanting managed tracking [secondary] |
| MLflow 3.x | Tracking, registry, agent serving/tracing | Self-hosted / Databricks | OSS-first, agent-centric LLMOps [secondary] |
| lm-eval 0.4.x | Academic benchmarks | Local/CI | Standardized model comparisons [secondary] |
| Inspect | Agentic/safety evals | Local | Evals with tool use, safety [secondary] |
| LangSmith/Langfuse | Agent observability | SaaS/self-hosted | Framework-adjacent tracing [secondary] |

---

