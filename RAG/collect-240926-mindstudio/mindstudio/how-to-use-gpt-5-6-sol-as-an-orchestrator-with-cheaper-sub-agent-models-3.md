---
id: collect-240926-mindstudio/mindstudio/how-to-use-gpt-5-6-sol-as-an-orchestrator-with-cheaper-sub-agent-models-3
title: "how-to-use-gpt-5-6-sol-as-an-orchestrator-with-cheaper-sub-agent-models"
domain: mindstudio
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agent", "sol", "agents", "cost", "gpt-5.6", "luna", "terra"]
source: docs/RAG/clean_en/mindstudio/how-to-use-gpt-5-6-sol-as-an-orchestrator-with-cheaper-sub-agent-models.md
source_anchor: ""
source_lines: [253, 257]
sha256: a3992f2873c3b14c505827df52cadab47a0e1d384d4b9b0928f231664a877798
---

# how-to-use-gpt-5-6-sol-as-an-orchestrator-with-cheaper-sub-agent-models

- GPT-5.6 Sol’s value in a multi-agent system is in orchestration — planning, routing, and synthesis — not in executing every task.
- Luna and Terra handle mid-complexity and simple tasks at significantly lower cost without sacrificing output quality when prompts are well-written.
- The savings compound at scale: 60–80% cost reduction per run is achievable with well-structured routing.
- Sub-agent prompts must be self-contained — don’t rely on sub-agents having context they weren’t given.
- Always include an orchestrator review pass for any workflow where output quality matters.
