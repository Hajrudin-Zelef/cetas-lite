---
id: collect-250926-servers-hardware/servers-hardware/thinking-ollama
title: "thinking-ollama"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["reasoning"]
source: docs/RAG/clean4/thinking-ollama.md
source_anchor: ""
source_lines: [1, 31]
sha256: 1601913a492d2ad5a0f15eea0cbc833f4a67c4bb682fe4e098d1ac3d907b126a
---

# thinking-ollama

`thinking` field that separates their reasoning trace from the final answer.
Use this capability to audit model steps, animate the model *thinking*in a UI, or hide the trace entirely when you only need the final response. See the full list of thinking models.

## Discover a model’s thinking controls

Thinking controls vary by model. Use`/api/show` to discover the values a model supports and the value Ollama uses by default:
`thinking` object:
- `values` can contain booleans (`true` or`false` ) for on/off controls. It can also contain model-defined strings for named levels.
- `default` is used when`think` is not set.
- `values: [false]` means the model does not support thinking.
- If `thinking` is omitted, the model has no thinking metadata. The model might conduct thinking based on its existing behavior.

## Enable thinking in API calls

Set the`think` field on a chat or generate request:
- `true` : request thinking output.
- `false` : request no thinking output, if the model permits it.
- `null` : use the model default.
- A string: select a supported level from `thinking.values` . Use the exact value from`/api/show` . Numbers are not supported.

`/api/show` metadata, Ollama applies supported names exactly. Unsupported names use the model default.
The reasoning output and answer use separate fields. Chat returns `message.thinking` and `message.content`. Generate returns `thinking` and `response`.
- cURL
- Python
- JavaScript

## Stream the reasoning trace

Thinking streams interleave reasoning tokens before answer tokens. Detect the first`thinking` chunk to render a “thinking” section, then switch to the final reply once `message.content` arrives.
- Python
- JavaScript
