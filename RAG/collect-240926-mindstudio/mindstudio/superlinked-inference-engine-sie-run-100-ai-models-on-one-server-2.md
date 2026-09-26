---
id: collect-240926-mindstudio/mindstudio/superlinked-inference-engine-sie-run-100-ai-models-on-one-server-2
title: "superlinked-inference-engine-sie-run-100-ai-models-on-one-server"
domain: mindstudio
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "embedding", "embeddings", "gpus", "license", "open source"]
source: docs/RAG/clean_en/mindstudio/superlinked-inference-engine-sie-run-100-ai-models-on-one-server.md
source_anchor: ""
source_lines: [81, 85]
sha256: 9273e556cbeed016aca2471f634a66b09c51cd6401e407e93bb0a2dd10811b03
---

# superlinked-inference-engine-sie-run-100-ai-models-on-one-server

The core engine is open source under the Apache 2 license and free to self-host on your own hardware. Superlinked also sells a managed cloud version of the same engine for teams that prefer not to manage their own GPUs.

### Does text generation use the same server as embeddings and reranking?

Not exactly. In the demonstrated setup, text generation runs through a different backend deployed via Docker, separate from the embedding, reranking, and extraction server, but it uses the same model catalog and the same client SDK.
