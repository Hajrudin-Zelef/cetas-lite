---
id: collect-250926-servers-hardware/servers-hardware/lora-fine-tuning-hyperparameters-guide-1
title: "LoRA fine-tuning Hyperparameters Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["Meta", "Unsloth"]
dates: []
keywords: ["fine-tuning", "lora", "parameters", "attention", "dpo", "gpu", "gpus", "llama", "memory", "qlora", "research", "training"]
source: docs/RAG/clean4/LoRA fine-tuning Hyperparameters Guide.md
source_anchor: ""
source_lines: [1, 117]
sha256: e09424f0c8242b337b872768b644af69eeca0dd67c6e5d6e36cc2173f7932ced
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

