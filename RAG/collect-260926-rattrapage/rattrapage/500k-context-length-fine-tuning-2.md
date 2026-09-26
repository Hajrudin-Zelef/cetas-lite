---
id: collect-260926-rattrapage/rattrapage/500k-context-length-fine-tuning-2
title: "500K Context Length Fine-tuning"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["fine-tuning", "agent", "agents", "gpu", "memory", "research", "training"]
source: docs/RAG/lot-rattrapage/fine-tuning/500K Context Length Fine-tuning.md
source_anchor: ""
source_lines: [71, 91]
sha256: 15eed1b01a1fa1a26af08674bc77e084d09093ef42aca73288a0328063568ad4
---

# 500K Context Length Fine-tuning

We also discussed how Tiled MLP effectively does 3 forward passes and 1 backward, compared to normal gradient checkpointing which does 2 forward passes and 1 backward with Stas Bekman and [DeepSpeed](https://github.com/deepspeedai/DeepSpeed/pull/7664) provided a doc update for Tiled MLP within DeepSpeed.
{% hint style="success" %}
Next time fine-tuning runs out of memory, try turning on `unsloth_tiled_mlp = True`. This should save some VRAM as long as the context length is longer than the LLM's hidden dimension.
{% endhint %}
***
**With our latest update, it is possible to now reach 1M context length with a smaller model on a single GPU!**
**Try 500K-context gpt-oss-20b fine-tuning on our** [**80GB A100 Colab notebook**](https://colab.research.google.com/github/unslothai/notebooks/blob/main/nb/gpt_oss_$20B$_500K_Context_Fine_tuning.ipynb)**.**
If you've made it this far, we're releasing a new blog on our latest improvements in training speed this week so stay tuned by joining our [Reddit r/unsloth](https://www.reddit.com/r/unsloth/) or our Docs.
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/blog/500k-context-length-fine-tuning.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
