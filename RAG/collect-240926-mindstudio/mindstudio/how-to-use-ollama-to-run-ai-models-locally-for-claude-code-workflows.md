---
id: collect-240926-mindstudio/mindstudio/how-to-use-ollama-to-run-ai-models-locally-for-claude-code-workflows
title: "Terminal 1"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "DeepSeek", "Google", "Microsoft", "Mistral", "Nvidia", "OpenAI"]
dates: []
keywords: ["agentic", "agents", "claude", "context window", "cost", "deepseek", "gemini", "gpu", "gpus", "inference", "llama", "llama.cpp"]
source: docs/RAG/clean_en/mindstudio/how-to-use-ollama-to-run-ai-models-locally-for-claude-code-workflows.md
source_anchor: ""
source_lines: [1, 309]
sha256: 682f5bb3e4bab8f26ad2fe4dc1ff3c693ce149e7e7f96c4d57e0357dae0e6c51
---

# Terminal 1

<!-- source: https://www.mindstudio.ai/blog/ollama-run-ai-models-locally-claude-code-workflows -->

## Why Running AI Models Locally Is Worth Your Time

API costs add up fast. If you’re using Claude, GPT-4, or Gemini for coding workflows, it’s easy to burn through hundreds of dollars a month — especially when running agentic tasks that make dozens of calls per session. That’s where Ollama comes in.

Ollama lets you run powerful open-source AI models locally on your own hardware, with no API costs, no data leaving your machine, and no rate limits. When paired with Claude Code — Anthropic’s agentic coding tool — you get a capable local AI coding workflow that costs nothing to run after setup.

This guide covers exactly how to set up Ollama, connect it to Claude Code, and get a working local AI backend running on your machine. We’ll also cover which models perform best for coding tasks, common issues you’ll run into, and where tools like MindStudio fit if you want to extend your workflows beyond the terminal.

## What Ollama Actually Is

Ollama is an open-source runtime for running large language models on your local machine. It handles all the complexity of model management — downloading weights, quantization, GPU/CPU allocation, and serving — behind a simple CLI interface.

You pull a model with one command. You run it with another. That’s most of the interface.

Under the hood, Ollama:

- Manages model files in a local library (similar to Docker images)
- Uses `llama.cpp` for efficient inference across CPUs, Apple Silicon, and NVIDIA GPUs
- Exposes a local REST API at `http://localhost:11434`
- Provides an OpenAI-compatible endpoint at `http://localhost:11434/v1`

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

That last point is critical. The OpenAI-compatible endpoint is what makes it possible to plug Ollama into tools originally designed for commercial APIs.

## Installing Ollama and Pulling Models

### Install Ollama

Installation is straightforward on all major platforms.

**macOS:**

`brew install ollama`
Or download the installer from the Ollama website.

**Linux:**

`curl -fsSL https://ollama.com/install.sh | sh`
**Windows:**
Download the installer from the Ollama website. Native Windows support has improved significantly in recent releases.

After installation, start the Ollama service:

`ollama serve`
This starts the local server on port 11434. Keep this running in the background while you work.

### Pull a Model

`ollama pull gemma3:12b`
Ollama’s model library includes dozens of models. For coding workflows, here are the practical choices:

| Model | Size | Best For | 
|---|---|---|
| `gemma3:12b` | ~8GB | Strong general coding, good tool use | 
| `llama3.1:8b` | ~5GB | Fast responses, decent code quality | 
| `llama3.1:70b` | ~40GB | Near-frontier quality (needs 48GB+ RAM) | 
| `qwen2.5-coder:14b` | ~9GB | Purpose-built for coding tasks | 
| `deepseek-coder-v2:16b` | ~10GB | Strong at code generation and debugging | 
| `phi4:14b` | ~9GB | Microsoft’s model, surprisingly capable | 
| `mistral:7b` | ~4GB | Lightweight, fast iteration | 

For most developers running on a MacBook Pro M-series or a mid-range GPU setup, `qwen2.5-coder:14b` or `gemma3:12b` offer the best balance of quality and speed.

### Verify the API is Running

`curl http://localhost:11434/api/tags`
You should see a JSON list of your installed models. If you get a connection error, run `ollama serve` first.

## How Claude Code Connects to a Local Backend

Claude Code (the `claude` CLI from Anthropic) is built to call Anthropic’s API. By default, it sends requests to `https://api.anthropic.com`. But it respects the `ANTHROPIC_BASE_URL` environment variable — which lets you redirect those requests to a different endpoint.

The catch: Ollama uses OpenAI’s API format, not Anthropic’s. Claude Code sends requests in Anthropic’s message format. So you can’t point Claude Code directly at Ollama without a translation layer.

The solution is **LiteLLM**, a proxy that accepts Anthropic-format requests and translates them for any downstream provider — including Ollama.

The flow looks like this:

`Claude Code → LiteLLM Proxy (localhost:8000) → Ollama (localhost:11434) → Local Model`
This setup works reliably and takes about 10 minutes to configure.

## Setting Up the LiteLLM Proxy

### Install LiteLLM

LiteLLM is a Python package. Install it with pip:

`pip install litellm[proxy]`
Or using pipx for isolated installation:

`pipx install litellm[proxy]`
### Create a Config File

Create a file called `litellm_config.yaml`:

```
model_list:
  - model_name: claude-3-5-sonnet-20241022
    litellm_params:
      model: ollama/qwen2.5-coder:14b
      api_base: http://localhost:11434
  - model_name: claude-3-haiku-20240307
    litellm_params:
      model: ollama/gemma3:12b
      api_base: http://localhost:11434
general_settings:
  master_key: sk-local-dev-key
```
The `model_name` values here match what Claude Code will send in its API requests. By naming them after real Anthropic models, Claude Code doesn’t need any modification — it thinks it’s talking to Anthropic, but LiteLLM routes the requests to your local Ollama instance.

### Start the Proxy

`litellm --config litellm_config.yaml --port 8000`
You’ll see output confirming the proxy is running on `localhost:8000`.

## Configuring Claude Code to Use Ollama

With Ollama and LiteLLM both running, configure Claude Code to use your local proxy.

### Set Environment Variables

```
export ANTHROPIC_BASE_URL=http://localhost:8000
export ANTHROPIC_API_KEY=sk-local-dev-key
```
### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

The API key value doesn’t matter for local use — you just need to provide something so Claude Code doesn’t error on authentication. Use whatever string you set as `master_key` in your LiteLLM config.

### Add to Your Shell Profile

To make this permanent, add both lines to your `.zshrc`, `.bashrc`, or equivalent:

```
echo 'export ANTHROPIC_BASE_URL=http://localhost:8000' >> ~/.zshrc
echo 'export ANTHROPIC_API_KEY=sk-local-dev-key' >> ~/.zshrc
source ~/.zshrc
```
### Test the Connection

Run Claude Code in a project directory:

`claude`
Try a simple prompt like "explain what this function does" or "add error handling to this file." If you see a response, the connection is working. Response times will be slower than the Anthropic API — expect 5–30 seconds depending on your hardware and model size.

## Choosing the Right Model for Coding Workflows

Not all open-source models handle coding tasks equally. Here's what to know before picking one.

### Tool Use Support Is Critical

Claude Code relies heavily on tool use (function calling) to interact with your filesystem, run commands, and make targeted edits. Many smaller models don't support tool use well, or at all.

Models with strong tool use support for local coding workflows:

- **Qwen 2.5 Coder** (any size) — Built for code, solid tool use
- **Llama 3.1 8B/70B** — Function calling is well-implemented
- **Gemma 3 12B** — Good instruction following, improving tool use
- **Mistral 7B/Nemo** — Decent tool use, fast

Avoid models that weren't instruction-tuned or that lack explicit function-calling support if you're doing anything beyond simple chat.

### Context Window Matters

Claude Code often sends large context windows — full file contents, project trees, conversation history. Check your model's context length:

- Qwen 2.5 Coder: 128K tokens
- Llama 3.1: 128K tokens
- Gemma 3: 128K tokens
- Mistral 7B: 32K tokens

For typical coding sessions, 32K is usually sufficient. But if you're working in large codebases, prioritize models with longer context windows.

### Hardware Requirements

A rough guide for running these models comfortably:

- **8GB VRAM / 16GB unified RAM (M2 MacBook):** Stick to 7B–8B models, or 4-bit quantized versions of 13B models
- **16GB VRAM / 32GB unified RAM (M3 Pro/Max):** 12B–14B models run well
- **24GB+ VRAM:** 30B+ models become viable
- **48GB+ RAM:** 70B models in 4-bit quantization

If performance feels slow, try the `q4_K_M` quantized version of your model (e.g., `ollama pull qwen2.5-coder:14b-q4_K_M`). Quantized models trade a small amount of quality for significant speed gains.

## Practical Workflow Tips

### Run Both Servers in the Background

Set up Ollama and LiteLLM to start automatically, or use a process manager like `tmux` or `screen` to keep them running:

```
# Terminal 1
ollama serve
# Terminal 2
litellm --config ~/litellm_config.yaml --port 8000
```
### Switch Between Local and Cloud Models

You can toggle between your local Ollama backend and the real Anthropic API by swapping environment variables:

```
# Use local Ollama backend
export ANTHROPIC_BASE_URL=http://localhost:8000
export ANTHROPIC_API_KEY=sk-local-dev-key
# Use real Anthropic API
unset ANTHROPIC_BASE_URL
export ANTHROPIC_API_KEY=your-real-anthropic-key
```
This is useful when you need Claude's full capabilities for a complex task but want to use local models for routine work.

### Use `.env` Files Per Project

Rather than global shell variables, create a `.env` file in each project directory:

```
ANTHROPIC_BASE_URL=http://localhost:8000
ANTHROPIC_API_KEY=sk-local-dev-key
```
Then load it before running Claude Code:

`set -a; source .env; set +a; claude`
## Troubleshooting Common Issues

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

### Claude Code Returns Connection Errors

Check that both Ollama and LiteLLM are running. Test each independently:

```
# Test Ollama
curl http://localhost:11434/api/tags
# Test LiteLLM proxy
curl http://localhost:8000/health
```
### Responses Are Truncated or Stop Mid-Output

This usually means the model hit its context limit or the LiteLLM proxy timed out. Try:

1. Switching to a model with a larger context window
2. Increasing LiteLLM's timeout settings in the config
3. Working in smaller files or splitting your request

### Tool Use Fails or Model Ignores Tool Calls

Not all models implement function calling the same way. If your model keeps ignoring tool use prompts:

1. Switch to Qwen 2.5 Coder or Llama 3.1 — these have the most consistent tool use implementations
2. Check if the model variant you downloaded is the instruct version (e.g., `llama3.1:8b` is instruct by default in Ollama)

### Model Pulls Fail or Are Slow

Large models (14B+) can take 20–40 minutes to download on slower connections. Ollama resumes interrupted downloads, so if it fails, just run the pull command again.

## Frequently Asked Questions

### Can you use Ollama with Claude Code directly without LiteLLM?

Not cleanly. Claude Code sends requests in Anthropic's API format, and Ollama's API uses a different format (OpenAI-compatible). Without a translation proxy like LiteLLM, the requests will fail or produce malformed responses. LiteLLM is the standard solution here, and setup takes about 10 minutes.

### What models does Ollama support in 2025?

Ollama supports hundreds of models through its model library, including Llama 3.1/3.2, Gemma 3, Qwen 2.5/2.5-Coder, Mistral, Phi-4, DeepSeek Coder V2, WizardCoder, and many others. You can browse the full list at ollama.com/library. New models are typically added within days of their public release.

### Is running AI locally actually faster than using the API?

It depends on your hardware. On a high-end MacBook Pro M3 Max or a machine with a dedicated GPU (24GB+ VRAM), local inference on a 13B model can be faster than waiting for remote API responses during peak hours. On older hardware or when running CPU-only inference, it will be noticeably slower. Smaller quantized models (7B at q4) often feel snappy even on mid-range hardware.

### Does Ollama support multi-modal models?

Yes. Ollama supports models with vision capabilities, including `llava`, `llama3.2-vision`, and `moondream`. These let you process images alongside text. Not all multi-modal features work perfectly through proxy setups — test your specific use case.

### How do you keep Ollama models up to date?

Ollama doesn’t auto-update models. To get the latest version of a model, pull it again:

`ollama pull qwen2.5-coder:14b`
Ollama will download only the changed layers if the model has been updated.

### Is this setup suitable for production or team use?

For individual development use, yes. For team or production use, you’d need to set up a shared Ollama server (rather than running it per-machine), manage access controls, and account for the fact that local models are generally less capable than frontier models like Claude 3.5 Sonnet for complex reasoning tasks. This setup works best as a cost-free option for routine coding assistance and experimentation.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

## Key Takeaways

- **Ollama is a local model runtime** that handles downloading, quantization, and serving open-source models through a simple CLI and REST API.
- **Claude Code supports custom API endpoints** via the`ANTHROPIC_BASE_URL` environment variable — but you need LiteLLM as a translation proxy between Claude Code’s Anthropic format and Ollama’s OpenAI format.
- **The setup takes about 15–20 minutes** : install Ollama, pull a model, install LiteLLM, create a config, start both servers, set two environment variables.
- **Model choice matters** : For coding workflows, prioritize models with strong tool use support — Qwen 2.5 Coder and Llama 3.1 are the most reliable options.
- **This is a zero-cost local alternative** for routine coding assistance, with the tradeoff of slower responses and lower capability compared to frontier models.

If you want to take local AI workflows further — connecting Ollama-generated outputs to external services, building team-facing tools, or automating multi-step processes without writing glue code — MindStudio supports Ollama natively and handles the integration layer for you.
