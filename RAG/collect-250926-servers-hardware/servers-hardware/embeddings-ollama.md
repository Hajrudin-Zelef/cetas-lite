---
id: collect-250926-servers-hardware/servers-hardware/embeddings-ollama
title: "embeddings-ollama"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["embedding", "embeddings"]
source: docs/RAG/clean4/embeddings-ollama.md
source_anchor: ""
source_lines: [1, 25]
sha256: 78dc8a5ac4d1af05454e86ea3aa95c9c2eff02d6bb93144ee1123816deba0e82
---

# embeddings-ollama

## Recommended models

## Generate embeddings

- CLI
- cURL
- Python
- JavaScript

Generate embeddings directly from the command line:You can also pipe text to generate embeddings:Output is a JSON array.

The 

`/api/embed` endpoint returns L2‑normalized (unit‑length) vectors.
## Generate a batch of embeddings

Pass an array of strings to`input`.
- cURL
- Python
- JavaScript

## Tips

- Use cosine similarity for most semantic search use cases.
- Use the same embedding model for both indexing and querying.
