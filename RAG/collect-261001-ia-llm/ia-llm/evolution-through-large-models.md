---
id: collect-261001-ia-llm/ia-llm/evolution-through-large-models
title: "evolution-through-large-models"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["research", "training"]
source: docs/RAG/collect-261001-ia-llm/evolution-through-large-models.md
source_anchor: ""
source_lines: [1, 1]
sha256: d9714a7fdfe1d25e2432c249df6bad9c67866baad2530da76c06dea0edd010de
---

# evolution-through-large-models

This paper pursues the insight that large language models (LLMs) trained to generate code can vastly improve the effectiveness of mutation operators applied to programs in genetic programming (GP). Because such LLMs benefit from training data that includes sequential changes and modifications, they can approximate likely changes that humans would make. To highlight the breadth of implications of such evolution through large models (ELM), in the main experiment ELM combined with MAP-Elites generates hundreds of thousands of functional examples of Python programs that output working ambulating robots in the Sodarace domain, which the original LLM had never seen in pre-training. These examples then help to bootstrap training a new conditional language model that can output the right walker for a particular terrain. The ability to bootstrap new models that can output appropriate artifacts for a given context in a domain where zero training data was previously available carries implications for open-endedness, deep learning, and reinforcement learning. These implications are explored here in depth in the hope of inspiring new directions of research now opened up by ELM.
