---
id: collect-250926-servers-hardware/servers-hardware/3x-faster-llm-training-with-unsloth-kernels-packing-2
title: "3x Faster LLM Training with Unsloth Kernels + Packing"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Unsloth"]
dates: []
keywords: ["training", "agent", "agents", "attention", "fine-tuning", "flash attention", "gpus", "llama", "memory"]
source: docs/RAG/clean4/3x Faster LLM Training with Unsloth Kernels + Packing.md
source_anchor: ""
source_lines: [66, 125]
sha256: d6d4b49910a3f5c18c7162e48078d46e06fdf3122e0eb96b5c239e5ea5e33d36
---

# 3x Faster LLM Training with Unsloth Kernels + Packing

The first graph (above) plots progress on `yahma/alpaca-cleaned` with `max_length = 2048`, Unsloth new with packing + kernels (maroon) vs. Unsloth old (gray). Both are trained with `max_steps = 500`, but we plot the x-axis in wall-clock time. Notice that we train on nearly 40% of an epoch in the packed case in the same amount of steps (and only a bit more wall-clock time) that it takes to train less than 5% of an epoch in the unpacked case.
Similarly, the 2nd graph (above) plots loss from the same runs, this time plotted with training steps on the x-axis. Notice that the losses match in scale and trend, but the loss in the packing case is less variable since the model is seeing more tokens per training step.
### :sparkles:How to enable packing?
**Update Unsloth first and padding free is done by default**! So all training is immediately 1.1 to 2x faster with 30% less memory usage at least and 0 change in loss curve metric!
{% code overflow="wrap" %}
```bash
pip install --upgrade --force-reinstall --no-cache-dir --no-deps unsloth
pip install --upgrade --force-reinstall --no-cache-dir --no-deps unsloth_zoo
```
{% endcode %}
We also support Flash Attention 3 via Xformers, SDPA support, Flash Attention 2, and this works on old GPUs (Tesla T4, RTX 2080) and new GPUs like H100s, B200s etc! Sample packing works *regardless of choice of attention backend or model family*, so enjoy the same speedups previously had with these fast attention implementations!
If you want to enable explicit packing, then add `packing = True` to enable up to 5x faster training!
{% hint style="warning" %}
Note `packing=True` will change the training loss and will make the dataset number of rows truncated, since multiple short sequences are packed into 1 sequence. You might see the number of examples in the dataset shrink.
To not get different training loss numbers, simply set `packing=False` and we will enable auto padding-free, which already makes training faster!
{% endhint %}
```python
from unsloth import FastLanguageModel
from trl import SFTTrainer, SFTConfig
model, tokenizer = FastLanguageModel.from_pretrained(
"unsloth/Qwen3-14B",
)
trainer = SFTTrainer(
model = model,
processing_class = tokenizer,
train_dataset = dataset,
args = SFTConfig(
per_device_train_batch_size = 1,
max_length = 4096,
…,
packing = True, # required to enable sample packing!
),
)
trainer.train()
```
All our notebooks are automatically faster (no need to do anything). See [Unsloth Notebooks](/docs/get-started/unsloth-notebooks.md)
{% columns %}
{% column %}
Qwen3 14B faster:
{% embed url="" %}
{% endcolumn %}
{% column %}
Llama 3.1 Conversational faster:
{% embed url="" %}
{% endcolumn %}
{% endcolumns %}
Thank you! If you're interested, see our [500K Context Training](/docs/blog/500k-context-length-fine-tuning.md) blog, [Memory Efficient RL](/docs/get-started/reinforcement-learning-rl-guide/memory-efficient-rl.md) blog and [Long Context gpt-oss](/docs/models/gpt-oss-how-to-run-and-fine-tune/long-context-gpt-oss-training.md) blog for more topics on kernels and performance gains!
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/blog/3x-faster-training-packing.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
