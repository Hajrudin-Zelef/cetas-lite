---
id: collect-260926-rattrapage/rattrapage/nvidia-nemotron-3-ultra-how-to-run-locally-3
title: "NVIDIA Nemotron 3 Ultra - How To Run Locally"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/lot-rattrapage/servers-reviews/NVIDIA Nemotron 3 Ultra - How To Run Locally.md
source_anchor: ""
source_lines: [219, 239]
sha256: aea72d815c98e5c8e7350b320ff7cd40dc179a8eba32c10157aa8b5efde0ef4b
---

# NVIDIA Nemotron 3 Ultra - How To Run Locally

| Multi-Challenge                               |         63.8        |          42.5         |        63.0       |        63.1       |  63.9 |  64.1 | 63.5 |
| **Long Context**                              |                     |                       |                   |                   |       |       |      |
| AA-LCR                                        |         65.4        |          69.8         |        66.9       |        70.2       |  68.3 |  67.3 | 62.7 |
| RULER (1M)                                    |         94.7        |           --          |         --        |         --        |  90.1 |  94.2 | 87.7 |
| Longbench v2 (≤ 1M)                           |         61.9        |           --          |         --        |         --        |  68.9 |  62.1 | 57.0 |
| **Multilingual**                              |                     |                       |                   |                   |       |       |      |
| MMLU-ProX (avg en/de/fr/es/it/ja/zh/hi/pt/ko) |         83.0        |          78.4         |        85.8       |        85.0       |  86.4 |  85.6 | 84.3 |
| WMT24++ (en→xx)                               |         83.7        |          82.8         |        84.4       |        84.5       |  86.8 |  85.9 | 85.9 |
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/models/nemotron-3-ultra.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
