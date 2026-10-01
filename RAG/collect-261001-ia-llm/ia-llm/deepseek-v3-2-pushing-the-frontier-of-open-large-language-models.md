---
id: collect-261001-ia-llm/ia-llm/deepseek-v3-2-pushing-the-frontier-of-open-large-language-models
title: "deepseek-v3-2-pushing-the-frontier-of-open-large-language-models"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek", "Google", "Hugging Face", "OpenAI"]
dates: []
keywords: ["deepseek", "agent", "agents", "gemini", "inference", "pricing", "reasoning", "research", "training"]
source: docs/RAG/collect-261001-ia-llm/deepseek-v3-2-pushing-the-frontier-of-open-large-language-models.md
source_anchor: ""
source_lines: [1, 25]
sha256: 8f88a1795be6000b59a4a19de7896efb9acea137cdf61869eba8ec567eb8bb8c
---

# deepseek-v3-2-pushing-the-frontier-of-open-large-language-models

Launching DeepSeek-V3.2 & DeepSeek-V3.2-Speciale â Reasoning-first models built for agents!

- DeepSeek-V3.2: Official successor to V3.2-Exp. Now live on App, Web & API.
- DeepSeek-V3.2-Speciale: Pushing the boundaries of reasoning capabilities. API-only for now.

Tech report: DeepSeek-V3.2 Tech Report - Hugging Face

## World-Leading Reasoning

- V3.2: Balanced inference vs. length. Your daily driver at GPT-5 level performance.
- V3.2-Speciale: Maxed-out reasoning capabilities. Rivals Gemini-3.0-Pro.
- Gold-Medal Performance: V3.2-Speciale attains gold-level results in IMO, CMO, ICPC World Finals & IOI 2025.

Note: V3.2-Speciale dominates complex tasks but requires higher token usage. Currently API-only (no tool-use) to support community evaluation & research.

## Thinking in Tool-Use

- Introduces a new massive agent training data synthesis method covering 1,800+ environments & 85k+ complex instructions.
- DeepSeek-V3.2 is our first model to integrate thinking directly into tool-use, and also supports tool-use in both thinking and non-thinking modes.

## API Update

- V3.2: Same usage pattern as V3.2-Exp.
- V3.2-Speciale: Served via a temporary endpoint: `base_url="https://api.deepseek.com/v3.2_speciale_expires_on_20251215"` . Same pricing as V3.2, no tool calls, available until Dec 15th, 2025, 15:59 (UTC Time).
- V3.2 now supports Thinking in Tool-Use â details: Thinking Mode Guide - API Docs
