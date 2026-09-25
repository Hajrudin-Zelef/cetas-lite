---
id: collect-250926-servers-hardware/servers-hardware/streaming-ollama
title: "streaming-ollama"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["reasoning", "tool calling"]
source: docs/RAG/clean4/streaming-ollama.md
source_anchor: ""
source_lines: [1, 13]
sha256: 9424d6f2b7ca8d4a8bfa7b5bfe9f2aca8d51cede32dd6e381305e80b542c428d
---

# streaming-ollama

`stream` parameter to `True`.
## Key streaming concepts

1. Chatting: Stream partial assistant messages. Each chunk includes the `content` so you can render messages as they arrive.
2. Thinking: Thinking-capable models emit a `thinking` field alongside regular content in each chunk. Detect this field in streaming chunks to show or hide reasoning traces before the final answer arrives.
3. Tool calling: Watch for streamed `tool_calls` in each chunk, execute the requested tool, and append tool outputs back into the conversation.

## Handling streamed chunks

 It is necessary to accumulate the partial fields in order to maintain the history of the conversation. This is particularly important for tool calling where the thinking, tool call from the model, and the executed tool result must be passed back to the model in the next request. 

- Python
- JavaScript
