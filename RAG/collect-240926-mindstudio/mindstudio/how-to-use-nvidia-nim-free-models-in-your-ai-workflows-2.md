---
id: collect-240926-mindstudio/mindstudio/how-to-use-nvidia-nim-free-models-in-your-ai-workflows-2
title: "Use exactly as you would any other LangChain LLM"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "China", "Nvidia", "OpenAI", "Z.ai"]
dates: []
keywords: ["agent", "agentic", "agents", "benchmarks", "claude", "context window", "cost", "embedding", "embeddings", "glm", "gpu", "gpus"]
source: docs/RAG/clean_en/mindstudio/how-to-use-nvidia-nim-free-models-in-your-ai-workflows.md
source_anchor: ""
source_lines: [141, 276]
sha256: 6a6da72943484940ad714ef1cfd17f5ad28863fc63d147fb93e204663cb0159e
---

# Use exactly as you would any other LangChain LLM

```
from langgraph.graph import StateGraph
from langchain_openai import ChatOpenAI
# Expensive model for planning
planner = ChatOpenAI(model="gpt-4o", api_key="your-openai-key")
# Free NIM model for data extraction
extractor = ChatOpenAI(
    model="zhipuai/glm-4",
    api_key="nvapi-xxxxxxxxxxxxxxxx",
    base_url="https://integrate.api.nvidia.com/v1"
)
def planning_node(state):
    return {"plan": planner.invoke(state["task"])}
def extraction_node(state):
    return {"extracted": extractor.invoke(state["document"])}
```
This pattern lets you build a cost-aware agent graph where high-value reasoning goes to premium models and repetitive extraction or formatting goes to free NIM models.

## Integrating NIM with Other Agentic Frameworks

The OpenAI-compatible API means NVIDIA NIM works with virtually every popular AI framework.

### CrewAI

```
from crewai import LLM
nim_model = LLM(
    model="openai/zhipuai/glm-4",  # CrewAI prefixes with "openai/"
    api_key="nvapi-xxxxxxxxxxxxxxxx",
    base_url="https://integrate.api.nvidia.com/v1"
)
```
Assign this LLM to any crew member that handles lower-complexity tasks like data normalization, formatting outputs, or generating boilerplate content.

### AutoGen

```
config_list = [
    {
        "model": "zhipuai/glm-4",
        "api_key": "nvapi-xxxxxxxxxxxxxxxx",
        "base_url": "https://integrate.api.nvidia.com/v1",
        "api_type": "openai"
    }
]
```
Pass this config to any AutoGen agent that doesn’t require the most capable model in your setup.

### Direct HTTP Calls

For any tool that doesn’t have built-in SDK support, you can make raw HTTP requests:

```
curl https://integrate.api.nvidia.com/v1/chat/completions \
  -H "Authorization: Bearer nvapi-xxxxxxxxxxxxxxxx" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "zhipuai/glm-4",
    "messages": [{"role": "user", "content": "Hello, what can you do?"}],
    "max_tokens": 256
  }'
```
## Common Mistakes and How to Avoid Them

### Using the Wrong Model ID

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

The most common source of errors when connecting NVIDIA NIM is passing an incorrect model identifier. Each model has a specific ID in the format `organization/model-name`. Check the catalog page for the exact string — don’t guess.

If you get a `model not found` error, that’s almost always the issue.

### Not Checking Rate Limits on Free Models

Free models on NVIDIA NIM have rate limits. For most development use cases these are generous, but if you’re running batch jobs, you can hit them. Check the specific model’s limits in the catalog and add appropriate retry logic with exponential backoff.

### Expecting Identical Output Quality Across Models

GLM-4 and similar free models are capable, but they’re not identical to GPT-4o or Claude Sonnet. Test your specific use cases before committing to routing a task to a free model. Some tasks — structured JSON extraction, code generation for specific frameworks, nuanced reasoning — may produce noticeably different results.

A practical approach: run a few dozen examples through both a premium model and your chosen NIM model, and compare outputs. That tells you whether the quality tradeoff is acceptable for your use case.

### Forgetting Context Window Limits

Different models have different context window sizes. If you’re routing tasks that involve long documents or conversation history to a NIM model, verify that its context window can handle your inputs. Sending a 50,000-token input to a model with a 4,096-token limit will fail.

### Passing System Prompts That Assume a Specific Model’s Behavior

If you’ve tuned system prompts for Claude or GPT-4o, they may not translate directly to a different model. Test your prompts with the NIM model you’re using and adjust where needed.

## Practical Use Cases for NVIDIA NIM Free Models

Here are the task categories where free NIM models consistently deliver value:

**Text classification and routing** — Categorizing support tickets, tagging content, routing messages to the right queue. These tasks don’t require the best model.

**Summarization** — Condensing long documents, meeting transcripts, or articles. Free models handle this well in most cases.

**Data extraction and transformation** — Pulling structured fields from unstructured text, formatting data for downstream systems.

**Translation and multilingual processing** — GLM-4 is particularly strong on Chinese/English tasks.

**First-pass content drafts** — Generating first drafts that a more capable model (or human) refines.

**Embedding generation** — If you're building a RAG pipeline, NVIDIA NIM also offers embedding models. Running embeddings on free or low-cost endpoints can significantly reduce vector database build costs.

## Frequently Asked Questions

### What models are free on NVIDIA NIM?

NVIDIA's API catalog offers a mix of free models and credit-based models. Free models don't consume your credit balance and include options from partners like Zhipu AI (GLM-4 family). The exact list of free models changes as NVIDIA adds and updates the catalog, so check build.nvidia.com for the current state. Each model page shows its pricing clearly.

### Is NVIDIA NIM the same as running a model locally?

No. NVIDIA NIM via the API catalog is a hosted inference service — you're calling NVIDIA's servers, not running anything on your own hardware. NVIDIA also offers NIM as a self-hosted deployment option (packaged as Docker containers you run on your own NVIDIA GPU infrastructure), but that's a separate product from the free API access described in this article.

### Can I use NVIDIA NIM free models in production?

For light production workloads, yes — but with caveats. Free tier access typically comes with rate limits and is subject to NVIDIA's terms of service. For high-volume production use, you'd likely need to move to a paid tier or consider self-hosting NIM on your own infrastructure. Check NVIDIA's current terms and rate limits before relying on free tier access for mission-critical workflows.

### How does GLM-4 compare to GPT-4o or Claude?

GLM-4 is a strong model for its size, particularly for instruction-following and bilingual Chinese/English tasks. For straightforward tasks — summarization, classification, extraction, translation — it performs comparably to much larger models. For complex multi-step reasoning, nuanced writing, or code generation for specialized domains, premium models like GPT-4o and Claude Sonnet generally outperform it. The right approach is to test on your specific tasks rather than rely on general benchmarks.

### Do I need an NVIDIA GPU to use NVIDIA NIM free models?

No. The free API access is fully hosted by NVIDIA. You're making HTTP requests to their endpoint — your own hardware doesn't matter. NVIDIA GPUs are only required if you're self-hosting NIM containers in your own infrastructure.

### What's the rate limit on NVIDIA NIM free models?

Rate limits vary by model and account tier. Free accounts typically get a set number of requests per minute and a daily request cap. Check the specific model's page in the catalog for current limits. For most development and testing purposes, the free limits are sufficient. If you're running batch jobs or high-frequency inference, you may need to implement queuing and retry logic.

## Key Takeaways

