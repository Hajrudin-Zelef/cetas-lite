---
id: collect-260926-rattrapage/rattrapage/how-to-run-local-llms-with-claude-code-4
title: "How to Run Local LLMs with Claude Code"
domain: rattrapage
role: reference
task: reference
actors: ["Unsloth"]
dates: []
keywords: ["claude", "agent", "agents", "gpu", "inference", "kv cache"]
source: docs/RAG/lot-rattrapage/ai-llm/How to Run Local LLMs with Claude Code.md
source_anchor: ""
source_lines: [362, 387]
sha256: 54786778f7efdb07da1a719b21a4d991d06eacf858fe003ea6ca5c2718a3b0e9
---

# How to Run Local LLMs with Claude Code

Try this prompt to install and run a simple Unsloth finetune:
{% code overflow="wrap" %}
```
You can only work in the cwd project/. Do not search for CLAUDE.md - this is it. Install Unsloth via a virtual environment via uv. Use `python -m venv unsloth_env` then `source unsloth_env/bin/activate` if possible. See https://unsloth.ai/docs/get-started/install/pip-install on how (get it and read). Then do a simple Unsloth finetuning run described in https://github.com/unslothai/unsloth. You have access to 1 GPU.
```
{% endcode %}
After waiting a bit, Unsloth will be installed in a venv via uv, and loaded up:
and finally you will see a successfully finetuned model with Unsloth!
{% hint style="warning" %}
If you see `Unable to connect to API (ConnectionRefused)` , remember to unset `ANTHROPIC_BASE_URL` via `unset ANTHROPIC_BASE_URL`
If you find open models to be 90% slower, [see here first](#fixing-90-slower-inference-in-claude-code) to fix KV cache being invalidated.
{% endhint %}
[^1]: Must use this!
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/basics/claude-code.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
