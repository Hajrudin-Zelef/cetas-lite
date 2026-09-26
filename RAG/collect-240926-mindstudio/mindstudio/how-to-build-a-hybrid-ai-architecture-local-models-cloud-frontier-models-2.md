---
id: collect-240926-mindstudio/mindstudio/how-to-build-a-hybrid-ai-architecture-local-models-cloud-frontier-models-2
title: "Install Ollama"
domain: mindstudio
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "Google", "Lambda", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "aws", "benchmark", "benchmarks", "claude", "context window", "cost", "embedding", "embeddings", "gemini"]
source: docs/RAG/clean_en/mindstudio/how-to-build-a-hybrid-ai-architecture-local-models-cloud-frontier-models.md
source_anchor: ""
source_lines: [106, 228]
sha256: 8d7046666b0189901af89f16a2237f79a634326fbb7e255b0c5442542bbf5279
---

# Install Ollama

- Ambiguity level
- Required reasoning depth
- Context length
- Output sensitivity (does a mistake here matter a lot?)

Based on the score, the request routes to the appropriate tier. This is more adaptive but adds a small amount of latency and complexity.

### 3. Two-stage routing with advisor pattern

This is one of the most effective patterns for quality-sensitive workflows. A cheaper or local model generates a first-pass response. A frontier model then reviews, critiques, or validates that response — without having to generate from scratch.

The Anthropic Advisor Strategy using Opus with Haiku or Sonnet is a concrete implementation of this idea. The expensive model’s job is to review, not originate — which dramatically cuts token usage while preserving output quality.

For teams building on top of multi-model routing infrastructure, this pattern integrates cleanly.

## Infrastructure: Running Local Models

You have two main options for running local models: on your own hardware, or on a private cloud VM.

### On your own hardware (Ollama)

## One coffee. One working app.

You bring the idea. Remy manages the project.

Ollama is the simplest way to run local models. It handles model management, VRAM optimization, and provides an OpenAI-compatible API endpoint. A machine with a modern GPU (or even Apple Silicon) can run 7B–14B parameter models comfortably.

Setup is minimal:

```
# Install Ollama
curl -fsSL https://ollama.com/install.sh | sh
# Pull a model
ollama pull qwen3:7b
# Start serving
ollama serve
```
Your local endpoint is then available at `http://localhost:11434/v1`, which is OpenAI-API compatible. Any code that calls OpenAI’s API can point to this endpoint instead with a one-line change.

### On a cloud VM

If you don’t have the right hardware locally, a GPU cloud VM (AWS g4dn, GCP A2, or a provider like RunPod or Lambda Labs) gives you the same flexibility. Spin up when needed, run models, shut down. At $0.50–$2/hour for a capable GPU instance, this is often far cheaper than frontier API calls at scale.

### Connecting local models to your agent stack

If you’re using MindStudio as your agent platform, connecting local LLMs to your AI agents is straightforward via the local model tunnel. The same applies to local image models if your workflow involves image analysis or generation.

## Choosing Your Frontier Models

For the cloud tier, you’re choosing between Claude, GPT-5, and Gemini. Each has different strengths. The best AI models for agentic workflows in 2026 is a useful reference if you’re evaluating options in depth.

In a hybrid architecture, your frontier model selection matters primarily for the tasks that actually reach it. A few things to prioritize:

- **Reasoning quality on your specific task type** — benchmark on representative examples, not general benchmarks.
- **Context window** — for synthesis and analysis tasks, you want 100K+ context.
- **Output format reliability** — for structured output tasks, some models are more consistent at following JSON schemas.
- **Latency** — for user-facing workflows, even a small latency difference at the frontier tier matters if users are waiting on it.

Cost matters too, but in a hybrid setup, your frontier model should be handling a minority of total requests. The ROI on quality at that tier is usually worth it.

## A Practical Architecture Example

Here’s how a document processing pipeline might be structured:

**Input: user uploads a PDF document**

1. **Local model (Whisper or document parser)** — Extract text from the PDF. No need for a frontier API call.
2. **Local embedding model** — Generate embeddings and store in vector database for retrieval.
3. **Local model (7B classifier)** — Classify the document type (contract, invoice, report, etc.) and extract structured fields.
4. **Conditional routing:**  - If document type is “routine invoice” → local model extracts line items, validates totals, outputs JSON. Done.
  - If document type is “complex contract” → frontier model (Claude Opus) analyzes the full document, identifies unusual clauses, and writes a plain-language summary.
5. **Local model** — Format the frontier model’s output for downstream systems (CRM, email, Slack notification).

In this pipeline, the frontier API is called only for the complex contract analysis step. Everything else runs locally. For a team processing 1,000 documents per day, the cost difference between “frontier model for everything” and this hybrid approach can be 5–10x.

## Managing Costs and Token Budgets

Even in a hybrid setup, frontier model costs can creep up. A few practices that help:

### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

**Set explicit token budgets per workflow.** Don’t let frontier model calls run unbounded. AI agent token budget management is worth reading if you’re building systems where agents can make multiple sequential calls.

**Compress context before frontier model calls.** Use a local model to summarize or filter content before it goes to your frontier API. Sending 50K tokens when 5K would do is expensive and slow.

**Cache frontier model responses for repeated inputs.** Semantic caching — where you cache based on embedding similarity rather than exact string match — can eliminate a significant portion of redundant frontier model calls.

**Monitor per-task costs, not just aggregate spend.** You need to know which specific workflow steps are generating the most frontier API spend. Without that visibility, optimization is guesswork.

The sub-agent era is pushing AI labs to release smaller, faster models specifically designed for the high-volume, low-complexity layers of agent pipelines. That trend plays directly into hybrid architecture design — as capable sub-agent models improve, the boundary of what belongs on the local tier keeps expanding.

## Where Remy Fits

Remy is relevant here for teams building AI-powered applications — particularly where you want the underlying infrastructure to support a hybrid model stack without having to wire it up yourself.

Remy applications are full-stack: real backend, typed SQL database, auth, deployment. If you’re building an application that incorporates AI workflows — document processing, intelligent routing, multi-step agents — you can describe that in the spec and let Remy compile the backend and frontend together.

The underlying infrastructure runs on MindStudio, which supports over 200 AI models from frontier providers and open-source options. That means you’re not locked to a single provider at any tier. You can describe a workflow in your spec that uses a local embedding model for retrieval and Claude Opus for synthesis, and the infrastructure handles the routing.

If you’re evaluating this kind of setup, try Remy at goremy.ai.

## Common Mistakes

### Using frontier models for high-volume, narrow tasks

The most expensive mistake. If you’re calling Claude Opus to classify support ticket severity at 10,000 tickets per day, you’re spending orders of magnitude more than necessary. A fine-tuned 7B model will match or exceed frontier performance on a specific classification task.

### Treating all local models as interchangeable

They’re not. A model good at text classification may be poor at structured JSON extraction. A model good at summarization may hallucinate on entity extraction. Benchmark each model on each task type you’re using it for.

### No fallback when local models fail

Local models can produce unexpected outputs or fail under certain input conditions. Always have a fallback path — either retry logic, a different local model, or escalation to a frontier model for validation.

### Ignoring latency requirements

