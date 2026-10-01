---
id: collect-261001-ia-llm/ia-llm/build-low-latency-multilingual-voice-agents-open-weights-full-deployment-control-with-nvid-2
title: "Build Low-Latency Multilingual Voice Agents: Open Weights & Full Deployment Control with NVIDIA Magpie TTS"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "Nvidia"]
dates: []
keywords: ["agent", "agents", "latency", "nvidia", "open weights", "voice", "attention", "fine-tuning", "gpu", "inference", "license", "multimodal"]
source: docs/RAG/collect-261001-ia-llm/build-low-latency-multilingual-voice-agents-open-weights-full-deployment-control-with-nvidia-magpie-.md
source_anchor: ""
source_lines: [92, 138]
sha256: be1728abe3f070c40a7bdadd0539fbee4fb78050946b0e50be3476fc82031a2a
---

# Build Low-Latency Multilingual Voice Agents: Open Weights & Full Deployment Control with NVIDIA Magpie TTS

For enterprises building production voice AI, this control over deployment, performance, and customization is often what matters most.

Voice AI in production is a system of models, not a single one. Magpie TTS is part of the NVIDIA Nemotron Voice Agent Developer Example, a reference implementation showing how purpose-built speech, language, and reasoning models work together as a coordinated system — so you can build always-on voice agents, not just better-sounding speech.

Developers can combine:

- Nemotron Speech for streaming speech recognition
- Magpie TTS for natural multilingual speech synthesis
- Nemotron language and multimodal models for reasoning, tool calling, and multimodal understanding
- NVIDIA NIM for GPU-optimized, production-ready inference microservices
- NeMo for customization and fine-tuning

The Nemotron Voice Agent developer example provides an end-to-end reference implementation that developers can clone, customize, and deploy in hours. It includes production patterns for:

- Real-time interruptible (barge-in) conversations
- Multimodal voice agents with vision understanding
- Multi-agent orchestration and tool calling
- Multilingual voice interactions
- Sub-second end-to-end latency using NVIDIA NIM

Rather than assembling individual components from scratch, developers can start from a complete reference architecture and adapt it to their own applications.

**Try the model**

**Deploy to production**

- NVIDIA Magpie Multilingual TTS NIM — optimized inference containers

**Customize for your domain**

- NVIDIA NeMo Speech — fine-tuning and training

**Build complete voice agents**

**Open weights and license**

- Model card on Hugging Face — open weights under the NVIDIA Open Model License.

Recommended inference configuration:

```
cfg_scale = 2.5          # classifier-free guidance — raise for tighter text adherence
temperature = 0.6
top_k = 80
apply_attention_prior = True
prior_epsilon = 0.1
```
