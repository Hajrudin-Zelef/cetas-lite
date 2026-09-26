---
id: collect-240926-mindstudio/mindstudio/why-prompt-rules-can-t-stop-your-ai-agent-from-going-rogue-2
title: "why-prompt-rules-can-t-stop-your-ai-agent-from-going-rogue"
domain: mindstudio
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/clean_en/mindstudio/why-prompt-rules-can-t-stop-your-ai-agent-from-going-rogue.md
source_anchor: ""
source_lines: [57, 79]
sha256: f5e6669f49982445f34633f59193795dcc4ac5834b9a00ddd1cd3bca0d3766fe
---

# why-prompt-rules-can-t-stop-your-ai-agent-from-going-rogue

### What’s the difference between prompt permissioning and tool permissioning?

Prompt permissioning is an instruction telling the agent what it should or shouldn’t do, which the model may or may not follow depending on context. Tool permissioning restricts what the agent is technically capable of doing, usually by limiting the access of the underlying API key, tool, or credential, regardless of what the prompt says.

### What is a scoped API key?

A scoped API key is a credential limited to specific actions or data, rather than full access to a service. For example, a key might allow an agent to draft messages but not send them, or read database records but not modify them.

## One coffee. One working app.

You bring the idea. Remy manages the project.

### Why do AI agents behave inconsistently even with the same instructions?

Language models are non-deterministic, meaning the same prompt and setup can produce different outputs across different runs. Small differences in context, phrasing, or the underlying model version can change how instructions are weighed against available actions.

### How do I know if my agent’s access is too broad?

List every tool, database, file, and API key the agent can use, and ask what the worst-case outcome looks like if it used each one in the most aggressive way possible. If any answer is concerning, narrow the actual permissions rather than adding another instruction.

### Can good testing replace tool-level restrictions?

No. Testing and evaluation tell you how reliably an agent performs within its current capabilities, but they don’t remove the underlying capability. Tool-level restrictions prevent the worst-case action from being possible at all, which testing alone cannot guarantee.
