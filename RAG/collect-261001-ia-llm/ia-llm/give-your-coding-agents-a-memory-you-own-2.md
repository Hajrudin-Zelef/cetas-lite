---
id: collect-261001-ia-llm/ia-llm/give-your-coding-agents-a-memory-you-own-2
title: "give-your-coding-agents-a-memory-you-own"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "memory", "embedding", "open source"]
source: docs/RAG/collect-261001-ia-llm/give-your-coding-agents-a-memory-you-own.md
source_anchor: ""
source_lines: [125, 136]
sha256: bfecd7a2c59ea1b85d734a75f818755a951df8038747d01ce8a4c0c37172ef8b
---

# give-your-coding-agents-a-memory-you-own

*“To think is to forget differences, generalize, make abstractions.”*
— Jorge Luis Borges, *Funes the Memorious*


Your agents already wrote the record. funes lives at
`github.com/huggingface/funes`, one command away
from turning that record into a memory the next agent can read, on whichever machine you
happen to be on.

funes invents little of this. It leans on open-source embedding models good enough to run locally, on Lance's append-only datasets with cheap incremental writes, and on the Hub's caching and content-dedup for datasets. The work is in fitting them into a memory an agent can actually use.

funes is open source too. Open an issue for anything from an install snag to a recall that missed, or an agent you'd like supported.
