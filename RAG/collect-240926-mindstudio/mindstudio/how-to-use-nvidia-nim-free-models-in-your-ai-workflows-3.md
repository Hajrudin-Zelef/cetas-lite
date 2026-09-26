---
id: collect-240926-mindstudio/mindstudio/how-to-use-nvidia-nim-free-models-in-your-ai-workflows-3
title: "Use exactly as you would any other LangChain LLM"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Nvidia", "OpenAI", "Z.ai"]
dates: []
keywords: ["claude", "context window", "glm", "nvidia", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-to-use-nvidia-nim-free-models-in-your-ai-workflows.md
source_anchor: ""
source_lines: [277, 281]
sha256: acc816f31e82a0370f82f6d1695e024ffd7280955eb901d3da96eb07afc5a603
---

# Use exactly as you would any other LangChain LLM

- NVIDIA NIM's API catalog at build.nvidia.com gives free API access to capable models like GLM-4, with an OpenAI-compatible endpoint that drops into any existing stack.
- Connecting to Claude Code, LangChain, CrewAI, or AutoGen requires changing only two things: the base URL (`https://integrate.api.nvidia.com/v1` ) and the model ID.
- Tiering your model usage — premium models for high-stakes reasoning, free NIM models for repetitive tasks — is one of the most effective ways to reduce AI infrastructure costs.
- Watch for common failure points: wrong model IDs, rate limits on batch jobs, context window mismatches, and system prompts that need tuning for a different model's behavior.
- If you want to use multiple models in workflows without managing API keys and integrations yourself, MindStudio lets you build multi-model workflows visually with 200+ models available out of the box.
