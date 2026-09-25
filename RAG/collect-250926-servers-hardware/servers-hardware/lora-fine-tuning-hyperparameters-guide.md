---
id: collect-250926-servers-hardware/servers-hardware/lora-fine-tuning-hyperparameters-guide
title: "LoRA fine-tuning Hyperparameters Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["Meta", "Unsloth"]
dates: []
keywords: ["fine-tuning", "lora", "parameters", "agent", "agents", "attention", "cost", "dpo", "gpu", "gpus", "inference", "llama"]
source: docs/RAG/clean4/LoRA fine-tuning Hyperparameters Guide.md
source_anchor: ""
source_lines: [1, 215]
sha256: 028304802fb6722931ff53151eaffd3517e2446b634eb8d214b39a06c9a98d9a
---

# LoRA fine-tuning Hyperparameters Guide

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/fine-tuning-llms-guide/lora-hyperparameters-guide.md).
# LoRA fine-tuning Hyperparameters Guide
Learn step-by-step the best LLM fine-tuning settings - LoRA rank & alpha, epochs, batch size + gradient accumulation, QLoRA vs. LoRA, target modules, and more.
LoRA hyperparameters are tunable settings that govern how Low-Rank Adaptation [fine-tunes](/docs/get-started/fine-tuning-llms-guide.md) LLMs. With many choices (e.g., learning rate and epochs) and countless combinations, picking the right values is key to accuracy, stability, quality, and fewer hallucinations. Done well, **LoRA can match full fine-tuning performance** while using 4× less VRAM.
You'll learn the best practices for these parameters, based on insights from hundreds of research papers and experiments, and see how they impact the model. **While we recommend using Unsloth's defaults**, understanding these concepts will give you full control.\
\
The goal is to change hyperparameter numbers to increase accuracy while counteracting [**overfitting or underfitting**](#overfitting-poor-generalization-too-specialized). Overfitting occurs when the model memorizes the training data, harming its ability to generalize to new, unseen inputs. The objective is a model that generalizes well, not one that simply memorizes.
{% columns %}
{% column %}
#### :question:But what is LoRA?
In LLMs, we have model weights. Llama 70B has 70 billion numbers. Instead of changing all 70b numbers, we instead add thin matrices A and B to each weight, and optimize those. This means we only optimize 1% of weights.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
## :1234: Key Fine-tuning Hyperparameters
### **Learning Rate**
Defines how much the model’s weights are adjusted during each training step.
* **Higher Learning Rates**: Lead to faster initial convergence but can cause training to become unstable or fail to find an optimal minimum if set too high.
* **Lower Learning Rates**: Result in more stable and precise training but may require more epochs to converge, increasing overall training time. While low learning rates are often thought to cause underfitting, they actually can lead to **overfitting** or even prevent the model from learning.
* **Typical Range**: `2e-4` (0.0002) to `5e-6` (0.000005).\
:green\_square: ***For normal LoRA/QLoRA Fine-tuning***, *we recommend* **`2e-4`** *as a starting point.*\
:blue\_square: ***For Reinforcement Learning** (DPO, GRPO etc.), we recommend* **`5e-6` .**\
:white\_large\_square: ***For Full Fine-tuning,** lower learning rates are generally more appropriate.*
### **Epochs**
The number of times the model sees the full training dataset.
* **More Epochs:** Can help the model learn better, but a high number can cause it to **memorize the training data**, hurting its performance on new tasks.
* **Fewer Epochs:** Reduces training time and can prevent overfitting, but may result in an undertrained model if the number is insufficient for the model to learn the dataset's underlying patterns.
* **Recommended:** 1-3 epochs. For most instruction-based datasets, training for more than 3 epochs offers diminishing returns and increases the risk of overfitting.
### **LoRA or QLoRA**
LoRA uses 16-bit precision, while QLoRA is a 4-bit fine-tuning method.
* **LoRA:** 16-bit fine-tuning. It's slightly faster and slightly more accurate, but consumes significantly more VRAM (4× more than QLoRA). Recommended for 16-bit environments and scenarios where maximum accuracy is required.
* **QLoRA:** 4-bit fine-tuning. Slightly slower and marginally less accurate, but uses much less VRAM (4× less).\
:sloth: *70B LLaMA fits in <48GB VRAM with QLoRA in Unsloth -* [*more details here*](https://unsloth.ai/blog/llama3-3)*.*
### Hyperparameters & Recommendations:

Hyperparameter

Function

Recommended Settings

LoRA Rank (r)

Controls the number of trainable parameters in the LoRA adapter matrices. A higher rank increases model capacity but also memory usage.

8, 16, 32, 64, 128

Choose 16 or 32

LoRA Alpha (lora_alpha)

Scales the strength of the fine-tuned adjustments in relation to the rank (r).

A regularization technique that randomly sets a fraction of LoRA activations to zero during training to prevent overfitting. Not that useful, so we default set it to 0.

0 (default) to 0.1

Weight Decay

A regularization term that penalizes large weights to prevent overfitting and improve generalization. Don't use too large numbers!

0.01 (recommended) - 0.1

Warmup Steps

Gradually increases the learning rate at the start of training.

5-10% of total steps

Scheduler Type

Adjusts the learning rate dynamically during training.

linear or cosine

Seed (random_state)

A fixed number to ensure reproducibility of results.

Any integer (e.g., 42, 3407)

Target Modules

Specify which parts of the model you want to apply LoRA adapters to — either the attention, the MLP, or both.

Attention: q_proj, k_proj, v_proj, o_proj

MLP: gate_proj, up_proj, down_proj

Recommended to target all major linear layers: q_proj, k_proj, v_proj, o_proj, gate_proj, up_proj, down_proj.

## :deciduous\_tree: Gradient Accumulation and Batch Size equivalency
### Effective Batch Size
Correctly configuring your batch size is critical for balancing training stability with your GPU's VRAM limitations. This is managed by two parameters whose product is the **Effective Batch Size**.\
\
**Effective Batch Size** = `batch_size * gradient_accumulation_steps`
* A **larger Effective Batch Size** generally leads to smoother, more stable training.
* A **smaller Effective Batch Size** may introduce more variance.
While every task is different, the following configuration provides a great starting point for achieving a stable **Effective Batch Size** of 16, which works well for most fine-tuning tasks on modern GPUs.
| Parameter | Description | Recommended Setting |
| --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------- |
| **Batch Size** (`batch_size`) |

The number of samples processed in a single forward/backward pass on one GPU.

Primary Driver of VRAM Usage. Higher values can improve hardware utilization and speed up training, but only if they fit in memory.

The number of micro-batches to process before performing a single model weight update.

Primary Driver of Training Time. Allows simulation of a larger batch\_size to conserve VRAM. Higher values increase training time per epoch.

| 8 |
| **Effective Batch Size** (Calculated) | The true batch size used for each gradient update. It directly influences training stability, quality, and final model performance. |

4 to 16 Recommended: 16 (from 2 \* 8)

|
### The VRAM & Performance Trade-off
Assume you want 32 samples of data per training step. Then you can use any of the following configurations:
* `batch_size = 32, gradient_accumulation_steps = 1`
* `batch_size = 16, gradient_accumulation_steps = 2`
* `batch_size = 8, gradient_accumulation_steps = 4`
* `batch_size = 4, gradient_accumulation_steps = 8`
* `batch_size = 2, gradient_accumulation_steps = 16`
* `batch_size = 1, gradient_accumulation_steps = 32`
While all of these are equivalent for the model's weight updates, they have vastly different hardware requirements.
The first configuration (`batch_size = 32`) uses the **most VRAM** and will likely fail on most GPUs. The last configuration (`batch_size = 1`) uses the **least VRAM,** but at the cost of slightly slower training**.** To avoid OOM (out of memory) errors, always prefer to set a smaller `batch_size` and increase `gradient_accumulation_steps` to reach your target **Effective Batch Size**.
### :sloth: Unsloth Gradient Accumulation Fix
Gradient accumulation and batch sizes **are now fully equivalent in Unsloth** due to our bug fixes for gradient accumulation. We have implemented specific bug fixes for gradient accumulation that resolve a common issue where the two methods did not produce the same results. This was a known challenge in the wider community, but for Unsloth users, the two methods are now interchangeable.
[Read our blog post](https://unsloth.ai/blog/gradient) for more details.
Prior to our fixes, combinations of `batch_size` and `gradient_accumulation_steps` that yielded the same **Effective Batch Size** (i.e., `batch_size × gradient_accumulation_steps = 16`) did not result in equivalent training behavior. For example, configurations like `b1/g16`, `b2/g8`, `b4/g4`, `b8/g2`, and `b16/g1` all have an **Effective Batch Size** of 16, but as shown in the graph, the loss curves did not align when using standard gradient accumulation:
After applying our fixes, the loss curves now align correctly, regardless of how the **Effective Batch Size** of 16 is achieved:
## 🦥 **LoRA Hyperparameters in Unsloth**
The following demonstrates a standard configuration. **While Unsloth provides optimized defaults**, understanding these parameters is key to manual tuning.
user\n",
response_part = "model\n",
)
```
## :mag\_right:Training on assistant responses only for vision models, VLMs
For language models, we can use `from unsloth.chat_templates import train_on_responses_only` as described previously. For vision models, use the extra arguments as part of `UnslothVisionDataCollator` just like before!
{% code overflow="wrap" %}
```python
class UnslothVisionDataCollator:
def __init__(
self,
...
# from unsloth.chat_templates import train_on_responses_only
# trainer = train_on_responses_only(
# trainer,
# instruction_part = "<|start_header_id|>user<|end_header_id|>\n\n",
# response_part = "<|start_header_id|>assistant<|end_header_id|>\n\n",
# )
train_on_responses_only = False, # EQUIVALENT to train_on_responses_only for LLMs
instruction_part = None, # EQUIVALENT to train_on_responses_only(instruction_part = ...)
response_part = None, # EQUIVALENT to train_on_responses_only(response_part = ...)
force_match = True, # Match newlines as well!
)
```
{% endcode %}
For example for Llama 3.2 Vision:
```python
UnslothVisionDataCollator(
model, tokenizer,
...
train_on_responses_only = True,
instruction_part = "<|start_header_id|>user<|end_header_id|>\n\n",
response_part = "<|start_header_id|>assistant<|end_header_id|>\n\n",
...
)
```
## :key: **Avoiding Overfitting & Underfitting**
### **Overfitting** (Poor Generalization/Too Specialized)
The model memorizes the training data, including its statistical noise, and consequently fails to generalize to unseen data.
{% hint style="success" %}
If your training loss drops below 0.2, your model is likely **overfitting** — meaning it may perform poorly on unseen tasks.
One simple trick is LoRA alpha scaling — just multiply the alpha value of each LoRA matrix by 0.5. This effectively scales down the impact of fine-tuning.
**This is closely related to merging / averaging weights.**\
You can take the original base (or instruct) model, add the LoRA weights, then divide the result by 2. This gives you an averaged model — which is functionally equivalent to reducing the `alpha` by half.
{% endhint %}
**Solution:**
* **Adjust the learning rate:** A high learning rate often leads to overfitting, especially during short training runs. For longer training, a higher learning rate may work better. It’s best to experiment with both to see which performs best.
* **Reduce the number of training epochs**. Stop training after 1, 2, or 3 epochs.
* **Increase** `weight_decay`. A value of `0.01` or `0.1` is a good starting point.
* **Increase** `lora_dropout`. Use a value like `0.1` to add regularization.
* **Increase batch size or gradient accumulation steps**.
* **Dataset expansion** - make your dataset larger by combining or concatenating open source datasets with your dataset. Choose higher quality ones.
* **Evaluation early stopping** - enable evaluation and stop when the evaluation loss increases for a few steps.
* **LoRA Alpha Scaling** - scale the alpha down after training and during inference - this will make the finetune less pronounced.
* **Weight averaging** - literally add the original instruct model and the finetune and divide the weights by 2.
### **Underfitting** (Too Generic)
The model fails to capture the underlying patterns in the training data, often due to insufficient complexity or training duration.
**Solution:**
* **Adjust the Learning Rate:** If the current rate is too low, increasing it may speed up convergence, especially for short training runs. For longer runs, try lowering the learning rate instead. Test both approaches to see which works best.
* **Increase Training Epochs:** Train for more epochs, but monitor validation loss to avoid overfitting.
* **Increase LoRA Rank** (`r`) and alpha: Rank should at least equal to the alpha number, and rank should be bigger for smaller models/more complex datasets; it usually is between 4 and 64.
* **Use a More Domain-Relevant Dataset**: Ensure the training data is high-quality and directly relevant to the target task.
* **Decrease batch size to 1**. This will cause the model to update more vigorously.
{% hint style="success" %}
Fine-tuning has no single "best" approach, only best practices. Experimentation is key to finding what works for your specific needs. Our notebooks automatically set optimal parameters based on many papers research and our experiments, giving you a great starting point. Happy fine-tuning!
{% endhint %}
***Acknowledgements:** A huge thank you to* [*Eyera*](https://huggingface.co/Orenguteng) *for contributing to this guide!*
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/fine-tuning-llms-guide/lora-hyperparameters-guide.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
