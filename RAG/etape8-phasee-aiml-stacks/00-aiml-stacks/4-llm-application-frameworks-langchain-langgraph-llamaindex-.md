---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/4-llm-application-frameworks-langchain-langgraph-llamaindex-
title: "4. LLM application frameworks: LangChain, LangGraph, LlamaIndex, Haystack, DSPy"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: reference
actors: ["OpenAI"]
dates: ["2025-10", "2026-02", "2026-05-14", "2026-05-22", "2026-06", "2026-06-22", "2026-08-24"]
keywords: ["llama", "agent", "agentic", "agents", "distribution", "license", "mit license", "packaging", "research"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [267, 355]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: 020f2413a1b41c73417f600f23d07f39c6fc78dda7c5ef769de2568d2feabc37
---

# 4. LLM application frameworks: LangChain, LangGraph, LlamaIndex, Haystack, DSPy

## 4. LLM application frameworks: LangChain, LangGraph, LlamaIndex, Haystack, DSPy

### 4.1 LangChain and LangGraph — 1.0 and the agent platform

- Multiple secondary sources report **LangChain and LangGraph reached 1.0
  in October 2025**, marking API stability after the 0.x era [secondary].
- The documented positioning: LangGraph is the stateful, durable execution
  layer (graphs, checkpoints, human-in-the-loop, long-running agents);
  LangChain is the higher-level component/agent layer on top [secondary].
- A secondary June-2026 source claimed LangChain 1.3.8 and
  `langchain-core` 1.4.6 — not verified against PyPI; treat as [unverified].
- **langchain-community was archived on May 22, 2026** (v0.4.2 final),
  ending the community-integration package as a separate distribution
  [secondary].
- CVE-2026-4539 affected LangChain (per a June 2026 newsletter roundup);
  details were not captured — gap G-04 [secondary].
- LangSmith is LangChain Inc.'s commercial observability/evaluation
  platform for agents; the LangGraph Platform hosts deployed graphs
  [secondary].

### 4.2 LlamaIndex — the data/RAG framework

- LlamaIndex is positioned around ingestion, document processing, indexing,
  hybrid retrieval, and RAG pipelines, complementing LangGraph's durable
  agent-workflow focus [secondary].
- Secondary release reports: 0.14.22 (May 14, 2026), 0.14.23 (June 24,
  2026), and 0.14.24 (mentioned August 24, 2026) — not verified against
  PyPI; treat the series as [secondary].
- By mid-2026 LlamaIndex had repositioned as an "agentic document and OCR
  platform," reflecting the shift from pure RAG indexing to document
  agents [secondary].
- **LlamaIndex Workflows 1.0** shipped June 22, 2026, adding the
  event-driven orchestration layer for multi-step agent pipelines
  [secondary].
- `llama-index-llms-google-genai` and similar provider packs show the
  per-provider packaging model (core + provider integrations) [secondary].

### 4.3 Haystack

- Haystack (by deepset) is the open-source, pipeline-oriented LLM framework
  with a strong production/NLP heritage; 2.x is the current architecture
  generation [secondary].
- No exact 2026 version was captured in this pass — gap G-05 [unverified].
- Its niche vs. LangChain/LlamaIndex: explicit pipeline DAGs, strong
  document-store abstractions, and deepset's managed platform [secondary].

### 4.4 DSPy — "programming, not prompting"

- DSPy (Stanford NLP, MIT license) treats LLM interactions as optimizable
  programs: developers declare typed `Signature`s and compose `Module`s;
  optimizers (teleprompters) compile them into prompts and few-shot
  examples [secondary].
- Reported versions: **3.1.3 (February 2026)** and **3.2.0 (April 21,
  2026)**; 3.2.0 absorbed GEPA and RLM-style optimizers, added more
  flexible `BetterTogether` chaining, and shipped production plumbing
  (caching, tracing, checkpointed optimizer state, MLflow hooks)
  [secondary].
- Three-layer architecture per a February-2026 research note: Module layer
  (`dspy.Module`, `Signature`, `Predict`), Adapter layer (`ChatAdapter`,
  `JSONAdapter`, `XMLAdapter`), Client layer (`dspy.LM`, LiteLLM backend,
  40+ providers) [secondary].
- ~22K–33.7K GitHub stars reported across sources (range conflicts) —
  recorded as conflict C-03; treat popularity figures as [secondary].
- DSPy 3.x requires Python 3.10+ (3.9 dropped in 3.0) [secondary].

### 4.5 RAG framework decision matrix

| Framework | Primary strength | Execution model | Best fit |
|---|---|---|---|
| LangGraph 1.x | Durable stateful agents | Checkpointed graphs, HITL | Long-running, human-in-loop agents [secondary] |
| LangChain 1.x | Components, integrations | Chains/agents over LangGraph | Fast agent prototyping [secondary] |
| LlamaIndex 0.14.x | Ingestion, indexing, retrieval | Pipelines + Workflows 1.0 | Document-heavy RAG [secondary] |
| Haystack 2.x | Production NLP pipelines | Pipeline DAGs | Enterprise search/RAG [secondary] |
| DSPy 3.x | Prompt/program optimization | Compiled signatures | Quality-critical LLM programs [secondary] |

### 4.6 Complementary pieces

- **LiteLLM** is the de-facto provider-unification proxy/SDK (OpenAI-style
  API across 100+ providers); DSPy uses it as its client backend
  [secondary].
- **Instructor / Outlines / Guidance** cover structured output and
  constrained generation; they compose with the frameworks above rather
  than replacing them [secondary].
- **LangSmith, Langfuse, Arize Phoenix, Braintrust** are the
  observability/eval layer most of these frameworks integrate with
  [secondary].

---

