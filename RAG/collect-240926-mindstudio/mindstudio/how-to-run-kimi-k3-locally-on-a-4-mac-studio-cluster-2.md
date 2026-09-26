---
id: collect-240926-mindstudio/mindstudio/how-to-run-kimi-k3-locally-on-a-4-mac-studio-cluster-2
title: "how-to-run-kimi-k3-locally-on-a-4-mac-studio-cluster"
domain: mindstudio
role: reference
task: reference
actors: ["Moonshot"]
dates: []
keywords: ["kimi", "cost", "gpu", "inference", "tokens per second"]
source: docs/RAG/clean_en/mindstudio/how-to-run-kimi-k3-locally-on-a-4-mac-studio-cluster.md
source_anchor: ""
source_lines: [61, 69]
sha256: 248f5765eef8bcf8039dff4cc4fef12dd589ead15f52398297abd204eeb57223
---

# how-to-run-kimi-k3-locally-on-a-4-mac-studio-cluster

Tensor parallelism, the method used to split the model across machines, requires the model’s dimensions to divide evenly across nodes. That only works cleanly with one, two, four, or eight machines, not three.

### How fast can a local Mac Studio cluster generate text with Kimi K3?

In testing, the cluster generated new tokens at about 14.7 tokens per second after processing prompts at around 238 tokens per second, which is considerably slower than cloud-based inference on GPU-backed infrastructure.

### Is a local Mac Studio cluster cheaper than using a cloud AI service?

Not in the short term. Each Mac Studio used in this setup cost around $16,000, making the full cluster a large one-time investment, while cloud services offering access to Kimi K3 and similar models can start at just a few dollars a month.
