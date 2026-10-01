---
id: collect-261001-ia-llm/ia-llm/granite-4-2-llms-how-they-re-built-5
title: "Default: previous thinking is stripped to save context"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI", "vLLM"]
dates: ["2026-09-11"]
keywords: ["cost", "reasoning", "vllm"]
source: docs/RAG/collect-261001-ia-llm/granite-4-2-llms-how-they-re-built.md
source_anchor: ""
source_lines: [474, 526]
sha256: 6468febfbc4e0c19816b1c2c71d6bb435dcd16f07750e7a9b5754a609fbcc945
---

# Default: previous thinking is stripped to save context

```
{
  "providers": {
    "vllm": {
      "baseUrl": "http://localhost:8000/v1",
      "api": "openai-completions",
      "apiKey": "EMPTY",
      "compat": {
        "supportsDeveloperRole": false,
        "supportsReasoningEffort": false
      },
      "models": [
        {
          "id": "granite-4.2-30b",
          "name": "Granite 4.2 30B",
          "reasoning": true,
          "input": ["text"],
          "contextWindow": 131072,
          "maxTokens": 8192,
          "samplingParams": {
            "temperature": 1.0,
            "top_p": 0.95
          },
          "cost": { "input": 0, "output": 0, "cacheRead": 0, "cacheWrite": 0 }
        }
      ]
    }
  }
}
```
**Run:**

```
pi
```
Then select the `granite-4.2-30b` model with `/model` or `Ctrl+L` in the interactive session.

For full documentation, see pi.dev/docs.

OpenHands is an AI software engineer that can plan, write code, and execute commands.

1. **Install and launch OpenHands** following the official installation guide.
2. **Configure the LLM** in the OpenHands settings with:
  - **Model:**`granite-4.2-30b`
  - **Base URL:**`http://localhost:8000/v1`
  - **API Key:** your vLLM`--api-key` value

**Note:** The `openai/` prefix is required when connecting to OpenAI-compatible endpoints like vLLM. Refer to the OpenHands local LLM documentation for detailed setup instructions, troubleshooting, and alternative installation methods.


**Resources:**

*Last updated: September 11, 2026 to clarify base models.*
