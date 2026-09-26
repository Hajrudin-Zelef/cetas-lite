---
id: collect-240926-mindstudio/mindstudio/how-to-self-host-an-open-weight-ai-stack-for-enterprise-in-under-a-day-deepseek-1
title: "how-to-self-host-an-open-weight-ai-stack-for-enterprise-in-under-a-day-deepseek-"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "Microsoft", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["deepseek", "open-weight", "agent", "agents", "attribution", "benchmark", "benchmarks", "claude", "compute", "consumer", "context window", "cost"]
source: docs/RAG/clean_en/mindstudio/how-to-self-host-an-open-weight-ai-stack-for-enterprise-in-under-a-day-deepseek-.md
source_anchor: ""
source_lines: [1, 117]
sha256: 76275f486fafe4ee228159043fd73655a78c11d9568137b81926bcbf21bfde66
---

# how-to-self-host-an-open-weight-ai-stack-for-enterprise-in-under-a-day-deepseek-

<!-- source: https://www.mindstudio.ai/blog/self-host-open-weight-ai-stack-enterprise-deepseek-v4 -->

## You Can Cut Your AI Inference Bill by 3x Before Lunch

Your team is spending real money on API tokens. Not “noticeable” money — money that shows up in budget reviews and makes finance ask uncomfortable questions. If you’re running document summarization, customer support agents, or internal RAG pipelines against GPT-5.5 at $5/M input tokens and $30/M output tokens, the math gets ugly fast at any meaningful scale.

Here’s what changes in under a day: you self-host DeepSeek V4 with Qwen embedding models and Llama 4 Scout or Maverick handling your local RAG and agent workloads. The result is near-equivalent benchmark performance for most enterprise use cases at a fraction of the cost — DeepSeek V4 runs at $1.74/M input tokens and $3.48/M output tokens when served from your own infrastructure, and the Qwen embedding models that power your retrieval layer cost essentially nothing beyond electricity once they’re running locally.

This isn’t a theoretical exercise. The models exist, the tooling is mature enough to deploy in a working day, and the business case is straightforward for any team spending more than a few thousand dollars a month on inference.

## Why Your Token Bill Is Probably 3x Higher Than It Needs to Be

### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

DeepSeek V4 is open-weight, carries a 1 million token context window, and benchmarks at near-parity with GPT-5.4 on math and Q&A tasks. Compare that to Claude Opus 4.7 at $5/M input and $25/M output, or Gemini 3.1 at $2/M input and $12/M output. For the overwhelming majority of enterprise workloads — document processing, structured data extraction, internal Q&A, support ticket routing — the frontier models are overkill.

The irony is that China’s GPU export restrictions accidentally made DeepSeek cheaper to serve than any American model. Because DeepSeek couldn’t access the most powerful Nvidia chips during training, they were forced to find compute-efficient training methods. Those efficiency gains translate directly into lower serving costs. The constraint became a structural advantage.

For enterprise teams, this creates a specific opportunity: you can run DeepSeek V4 on your own servers, keep all data on-premises, and pay only for the hardware and electricity. The security and compliance arguments alone justify the setup time for many regulated industries. The cost savings are a bonus.

The stack we’re building here pairs DeepSeek V4 as the primary reasoning model with Qwen embedding models for retrieval and Llama 4 Scout (or Maverick, if you need the heavier mixture-of-experts model) for agent tasks that benefit from a more efficient local model. If you want a deeper comparison of how Qwen models stack up against other open-weight options, the Gemma 4 vs Qwen 3.5 open-weight comparison covers the tradeoffs in detail.

## What You Need Before You Start

**Hardware.** DeepSeek V4 is too large for consumer GPUs — you need either a cloud VM with serious VRAM (an A100 or H100 instance), an on-premises server with multiple high-memory GPUs, or an Nvidia DGX Spark if you want an appliance-style local deployment. For the embedding layer, Qwen embedding models are small enough to run on almost anything with a GPU, including a Mac Studio with 128GB unified memory.

For Llama 4 Scout and Maverick, the mixture-of-experts architecture means the active parameter count per token is lower than the total model size suggests. Scout is the more practical choice for most teams — it runs on hardware that’s actually purchasable without a procurement process.

**Software prerequisites.** You need:

- Docker (for containerized serving)
- Ollama or vLLM (your inference runtime — Ollama for getting started quickly, vLLM when you’re serving a team)
- Python 3.10+ for the RAG pipeline
- A vector database: Postgres with pgvector is the production default; SQLite with sqlite-vec works for smaller deployments

**Accounts and access.** Pull the DeepSeek V4 weights from Hugging Face. You’ll need a Hugging Face account and enough disk space — plan for several hundred GB depending on quantization level. Qwen embedding models are available from the same source and are substantially smaller.

**Knowledge baseline.** You should be comfortable with the command line, understand what an embedding is at a conceptual level, and have some familiarity with REST APIs. You don’t need to understand the model architecture.

## The Setup, Step by Step

**Step 1: Stand up your inference runtime**

Install Ollama on your target machine. For Linux:

`curl -fsSL https://ollama.ai/install.sh | sh`
Start the Ollama service and verify it’s running on port 11434. Now you have a local OpenAI-compatible API endpoint that other tools can talk to.

If you’re deploying for a team rather than a single workstation, use vLLM instead — it handles batching and concurrent requests properly. The Ollama path is faster to get running; vLLM is the right answer when serving becomes infrastructure.

**Step 2: Pull DeepSeek V4 and your embedding model**

```
ollama pull deepseek-v4
ollama pull qwen:embedding
```
The DeepSeek V4 pull will take a while depending on your connection. While it’s downloading, pull Llama 4 Scout:

`ollama pull llama4:scout`
Once complete, run a quick sanity check:

`ollama run deepseek-v4 "Summarize the key risks in a standard enterprise software contract in three sentences."`
If you get a coherent response, the model is serving correctly. Now you have a working local inference layer.

**Step 3: Set up your vector store**

For a production RAG setup, Postgres with pgvector is the right choice. Install pgvector and create your embeddings table:

```
CREATE EXTENSION IF NOT EXISTS vector;
CREATE TABLE document_chunks (
  id SERIAL PRIMARY KEY,
  content TEXT,
  embedding vector(1536),
  source TEXT,
  created_at TIMESTAMP DEFAULT NOW()
);
CREATE INDEX ON document_chunks USING ivfflat (embedding vector_cosine_ops);
```
The Qwen embedding models output 1536-dimensional vectors by default — match your vector column dimension to whatever the model actually produces, which you can verify by running a test embedding and checking the output shape.

For smaller deployments or local experimentation, SQLite with sqlite-vec works and requires no separate database server. The how to run local AI models with Claude Code to cut costs post covers the SQLite path in more detail if that’s your starting point.

Now you have a vector store ready to receive embeddings.

**Step 4: Build the ingestion pipeline**

Your ingestion pipeline needs to: load documents, chunk them sensibly, generate embeddings via your local Qwen model, and store them in your vector database.

```
import requests
import psycopg2
import json
OLLAMA_URL = "http://localhost:11434/api/embeddings"
EMBEDDING_MODEL = "qwen:embedding"
def embed_text(text: str) -> list[float]:
    response = requests.post(OLLAMA_URL, json={
        "model": EMBEDDING_MODEL,
        "prompt": text
    })
    return response.json()["embedding"]
def chunk_document(text: str, chunk_size: int = 512, overlap: int = 64) -> list[str]:
    words = text.split()
    chunks = []
    for i in range(0, len(words), chunk_size - overlap):
        chunk = " ".join(words[i:i + chunk_size])
        chunks.append(chunk)
    return chunks
```
The chunking strategy matters more than most people expect. PDFs need different handling than plain text. Meeting transcripts need speaker attribution preserved. Code needs symbol-aware chunking. A naive word-count splitter will work well enough to validate the pipeline, but plan to revisit this before you call the system production-ready.

