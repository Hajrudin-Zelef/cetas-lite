---
id: collect-240926-mindstudio/mindstudio/how-to-use-ollama-to-run-ai-models-locally-for-claude-code-workflows-2
title: "Terminal 1"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Mistral", "OpenAI"]
dates: []
keywords: ["agents", "claude", "context window", "cost", "deepseek", "gpu", "inference", "llama", "mistral", "quantization", "qwen", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-to-use-ollama-to-run-ai-models-locally-for-claude-code-workflows.md
source_anchor: ""
source_lines: [187, 309]
sha256: 70a3e3b9c7b62051c31c3992d2452e23a9e85ffa451c2819539680937ea4b139
---

# Terminal 1

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
