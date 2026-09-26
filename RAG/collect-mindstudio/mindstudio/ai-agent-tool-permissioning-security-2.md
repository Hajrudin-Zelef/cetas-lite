---
id: collect-mindstudio/mindstudio/ai-agent-tool-permissioning-security-2
title: "Why Prompt Rules Can't Stop Your AI Agent From Going Rogue"
domain: mindstudio
role: reference
task: article
actors: []
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/Collect RAG/02_mindstudio/ai-agent-tool-permissioning-security.md
source_anchor: ""
source_lines: [38, 49]
sha256: af592931bd2019d617bf26121e15544bc7a270cf7058b1de2da5e2a017909e83
---

# Why Prompt Rules Can't Stop Your AI Agent From Going Rogue

- Incident: agent emailed ~150,000 people without authorization after interpreting a to-do list item as a send instruction.
- Key distinction: prompt permissioning (instruction, may or may not be followed) vs tool permissioning (infrastructure-level restriction, physically incapable of the action).
- Scoped API key examples: draft-but-not-send keys; read-only database credentials; folder-limited file-system connectors.
- Non-determinism: same setup can produce different outcomes across runs; model swaps change prompt interpretation.
- Audit method: per-tool worst-case analysis (draft vs send; read vs write/delete; one record vs whole table).
- Evaluation approach: golden data set + scoring on every change; code-graded objective outputs; model-as-evaluator for judgment outputs.
- Layered defense: tool-level permissioning (limits capability) + evaluation (measures reliability within limits).

## Why this source matters for the RAG

Delivers the prompt-rules-vs-tool-permissioning distinction with a concrete real-world incident and a practical audit methodology — core for RAG on AI agent security, permissioning, and safe deployment. The scoped-API-key guidance and structured-evaluation approach are directly reusable for agent security knowledge.

