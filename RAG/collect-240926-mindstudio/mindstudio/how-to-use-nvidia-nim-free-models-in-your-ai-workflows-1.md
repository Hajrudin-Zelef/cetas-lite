---
id: collect-240926-mindstudio/mindstudio/how-to-use-nvidia-nim-free-models-in-your-ai-workflows-1
title: "Use exactly as you would any other LangChain LLM"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "Mistral", "Nvidia", "OpenAI", "TensorRT-LLM", "Z.ai"]
dates: []
keywords: ["agent", "agentic", "agents", "claude", "cost", "embeddings", "glm", "gpus", "inference", "latency", "llama", "mistral"]
source: docs/RAG/clean_en/mindstudio/how-to-use-nvidia-nim-free-models-in-your-ai-workflows.md
source_anchor: ""
source_lines: [1, 140]
sha256: 4d1a1cba67c07bd7eeca3ecbb9cfbf9686f421bf422f5cc946f4ddee16a9871e
---

# Use exactly as you would any other LangChain LLM

<!-- source: https://www.mindstudio.ai/blog/nvidia-nim-free-models-ai-workflows -->

## What NVIDIA NIM Actually Offers (And Why It Matters for Cost)

Running AI workloads at scale gets expensive fast. Whether you’re building autonomous agents, running batch processing pipelines, or just experimenting with new models, API costs add up quickly. NVIDIA NIM free models offer a way to access capable, production-grade inference without burning through your budget — and connecting them to tools like Claude Code or any agentic framework is more straightforward than most people realize.

NVIDIA NIM (short for NVIDIA Inference Microservices) is an optimized inference platform that makes large language models and other AI models accessible via API. What makes it interesting for developers is that NVIDIA’s API catalog at build.nvidia.com includes a selection of models available for free — including models like GLM-4 that punch well above their weight for many common tasks.

This guide covers what the free tier looks like, how to get your API key, and how to wire these models into Claude Code, LangChain, or any other agentic tool you’re already using.

## Understanding NVIDIA NIM and Its Free Model Catalog

NVIDIA NIM packages optimized AI model inference into API-accessible endpoints. The underlying runtime is tuned specifically for NVIDIA GPUs using TensorRT-LLM, which means you get faster, lower-latency responses compared to generic inference setups.

The model catalog at build.nvidia.com hosts hundreds of models spanning:

- **Large language models** — Llama 3, Mistral, Mixtral, Phi-3, Qwen, and more
- **Multimodal models** — models that handle images alongside text
- **Code-focused models** — models optimized for code generation and analysis
- **Specialized models** — embeddings, reranking, and domain-specific variants

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

### What “Free” Actually Means on NVIDIA NIM

When you create a free account on NVIDIA’s API catalog, you get access to a credits pool that lets you try any model in the catalog. Some models are designated as free — meaning they don’t consume credits — while others give you a generous initial credit balance to test with.

GLM-4, developed by Zhipu AI and available on the platform, is one of the notable free-tier options. GLM-4 is a strong bilingual (Chinese/English) model that handles reasoning, instruction-following, and code tasks well. For teams doing multilingual work or looking for a capable alternative to pricier closed models, it’s worth serious consideration.

### The OpenAI-Compatible API Advantage

Here’s what makes NVIDIA NIM practical to integrate: its API is designed to be OpenAI-compatible. That means any tool, framework, or application that supports a custom base URL can point to NVIDIA NIM instead of OpenAI — with no other code changes required.

You change two things: the base URL and the model name. Everything else stays the same.

## Setting Up Your NVIDIA NIM API Access

Before you can use any NIM model, you need credentials. The setup takes about five minutes.

### Step 1: Create an Account

Go to build.nvidia.com and sign up for a free account. You’ll need to verify your email. Once you’re in, you’ll land on the API catalog.

### Step 2: Generate an API Key

Navigate to your account settings and find the API Key section. Generate a new key and save it somewhere secure — you’ll only see it once. It’ll look like `nvapi-xxxxxxxxxxxxxxxx`.

### Step 3: Explore the Model Catalog

Before writing any code, it’s worth browsing the catalog to understand what’s available. Each model page shows:

- The model identifier (what you’ll pass as the `model` parameter)
- Supported endpoints (chat completions, embeddings, etc.)
- Whether it’s free or costs credits
- A built-in playground to test it immediately

For GLM-4, the model ID on the NIM platform follows the format `zhipuai/glm-4` or a versioned variant — check the specific page for the exact string, as this matters when you make API calls.

### Step 4: Check the Endpoint URL

NVIDIA NIM’s base URL for OpenAI-compatible API calls is:

`https://integrate.api.nvidia.com/v1`
That’s the URL you’ll use as your custom base URL across all the tools and integrations covered below.

## Connecting NVIDIA NIM to Claude Code

Claude Code is Anthropic’s agentic coding assistant that runs in your terminal. It supports custom model configurations, which makes it possible to route specific tasks to a NVIDIA NIM model instead of — or alongside — Claude.

### How Claude Code Handles Custom Models

Claude Code uses a configuration file and environment variables to manage model routing. For connecting external OpenAI-compatible providers, you set environment variables that tell Claude Code where to send requests and what credentials to use.

### Setting Up the Connection

**Option 1: Environment variables**

Set these in your terminal session or `.bashrc`/`.zshrc` file:

```
export NVIDIA_NIM_API_KEY="nvapi-xxxxxxxxxxxxxxxx"
export NVIDIA_NIM_BASE_URL="https://integrate.api.nvidia.com/v1"
```
When using Claude Code with a workflow that calls an LLM backend, you can configure the underlying LLM client to use these values.

**Option 2: Direct SDK integration**

If you’re building custom tools or scripts that Claude Code calls as part of an agentic loop, you can use the OpenAI Python SDK pointed at NVIDIA NIM:

```
from openai import OpenAI
client = OpenAI(
    api_key="nvapi-xxxxxxxxxxxxxxxx",
    base_url="https://integrate.api.nvidia.com/v1"
)
response = client.chat.completions.create(
    model="zhipuai/glm-4",  # use the exact model ID from the catalog
    messages=[
        {"role": "user", "content": "Explain this function in one paragraph."}
    ],
    temperature=0.3,
    max_tokens=512
)
print(response.choices[0].message.content)
```
Because it’s OpenAI-compatible, the response structure is identical to what you’d get from OpenAI. Your existing parsing logic doesn’t need to change.

### When to Route Tasks to NIM Free Models

Not every task in an agentic workflow needs your most expensive model. A practical pattern is to tier your model usage:

- **High-stakes reasoning, planning, or synthesis** : Use a premium model (Claude, GPT-4o)
- **Repetitive classification, summarization, or formatting tasks** : Use a free NIM model
- **Multilingual processing** : GLM-4 is particularly strong here

This tiering approach can reduce your API spend by 40–70% on typical agentic workflows without meaningful quality degradation on the tasks routed to free models.

## Using NVIDIA NIM Free Models in LangChain and LangGraph

LangChain has native support for any OpenAI-compatible endpoint. Wiring up NVIDIA NIM takes a single configuration change.

### LangChain Setup

```
from langchain_openai import ChatOpenAI
nim_llm = ChatOpenAI(
    model="zhipuai/glm-4",
    api_key="nvapi-xxxxxxxxxxxxxxxx",
    base_url="https://integrate.api.nvidia.com/v1",
    temperature=0.2
)
# Use exactly as you would any other LangChain LLM
response = nim_llm.invoke("Summarize this document in three bullet points.")
```
### Using It in a LangGraph Agent

If you’re building multi-agent systems with LangGraph, you can assign different nodes to different models:

