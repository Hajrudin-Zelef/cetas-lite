---
id: collect-250926-servers-hardware/servers-hardware/inference-and-deployment
title: "Inference & Deployment"
domain: servers-hardware
role: reference
task: reference
actors: ["Unsloth"]
dates: []
keywords: ["inference", "agent", "agents", "inference engine"]
source: docs/RAG/clean4/inference-and-deployment.md
source_anchor: ""
source_lines: [1, 18]
sha256: 8eb83d7473eb3763c13c8522e486645dfeea19000a6729bbfc17fa247d06d499
---

# Inference & Deployment

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/basics/inference-and-deployment.md).
# Inference & Deployment
Learn how to save your finetuned model so you can run it in your favorite inference engine.
You can also run your fine-tuned models by using [Unsloth's 2x faster inference](/docs/basics/inference-and-deployment/unsloth-inference.md).

---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/basics/inference-and-deployment.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
