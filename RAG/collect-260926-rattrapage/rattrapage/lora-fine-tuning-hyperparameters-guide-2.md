---
id: collect-260926-rattrapage/rattrapage/lora-fine-tuning-hyperparameters-guide-2
title: "LoRA fine-tuning Hyperparameters Guide"
domain: rattrapage
role: reference
task: reference
actors: ["Unsloth"]
dates: []
keywords: ["fine-tuning", "lora", "parameters", "agent", "agents", "cost", "gpus", "inference", "llama", "memory", "open source", "research"]
source: docs/RAG/lot-rattrapage/fine-tuning/LoRA fine-tuning Hyperparameters Guide.md
source_anchor: ""
source_lines: [118, 214]
sha256: dbd54e956736e3439f9ed1649974ae744901129c75463837a70e93bd0805f33d
---

# LoRA fine-tuning Hyperparameters Guide

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
