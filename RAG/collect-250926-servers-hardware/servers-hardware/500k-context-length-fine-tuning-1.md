---
id: collect-250926-servers-hardware/servers-hardware/500k-context-length-fine-tuning-1
title: "500K Context Length Fine-tuning"
domain: servers-hardware
role: reference
task: reference
actors: ["Unsloth"]
dates: ["2024-04"]
keywords: ["fine-tuning", "context window", "cost", "gpu", "gpus", "memory", "qlora", "research", "throughput", "training"]
source: docs/RAG/clean4/500K Context Length Fine-tuning.md
source_anchor: ""
source_lines: [1, 70]
sha256: 0696587345ed26a42317f36d3f6c996b434c4a4afa89eac86937577608da8af9
---

# 500K Context Length Fine-tuning

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/blog/500k-context-length-fine-tuning.md).
# 500K Context Length Fine-tuning
Learn how to enable >500K token context window fine-tuning with Unsloth.
We’re introducing new algorithms in Unsloth that push the limits of long-context training for **any LLM and VLM**. Training LLMs like gpt-oss-20b can now reach **500K+ context lengths** on a single 80GB H100 GPU, compared to 80K previously with no accuracy degradation.
You can reach >**750K context windows** on a B200 192GB GPU.
> **Try 500K-context gpt-oss-20b fine-tuning on our** [**80GB A100 Colab notebook**](https://colab.research.google.com/github/unslothai/notebooks/blob/main/nb/gpt_oss_$20B$_500K_Context_Fine_tuning.ipynb)**.**
We’ve significantly improved how Unsloth handles memory usage patterns, speed, and context lengths:
* **60% lower VRAM use** with **3.2x longer context** via Unsloth’s new [fused and chunked cross-entropy](#unsloth-loss-refactoring-chunk-and-fuse) loss, with no degradation in speed or accuracy
* Enhanced activation offloading in Unsloth’s [**Gradient Checkpointing**](#unsloth-gradient-checkpointing-enhanced)
* Collabing with Stas Bekman from Snowflake on [Tiled MLP](#tiled-mlp-unlocking-500k), enabling 2× more contexts
Unsloth’s algorithms allows gpt-oss-20b QLoRA (4bit) with 290K context possible on a H100 with no accuracy loss, and 500K+ with Tiled MLP enabled, altogether delivering >**6.4x longer context lengths.**
### 📐 Unsloth Loss Refactoring: Chunk & Fuse
Our new fused loss implementation adds **dynamic sequence chunking**: instead of computing language model head logits and cross-entropies over the entire sequence at once, we process manageable slices along the flattened sequence dimension. This cuts peak memory from GBs to a smaller chunk sizes. Each chunk still runs a fully fused forward + backward pass via `torch.func.grad_and_value` , and retains mixed precision accuracy by upcasting to float32 if necessary. **These changes do not degrade training speed or accuracy.**
The key innovation is that the **chunk size is chosen automatically at runtime** based on available VRAM.
* If you have more free VRAM, larger chunks are used for faster runs
* If you have less VRAM, it increases the number of chunks to avoid memory blowouts.
This **removes manual tuning** and keeps our algorithm robust across old and new GPUs, workloads and different sequence lengths.
{% hint style="success" %}
Due to automatic tuning, **smaller contexts will use more VRAM** (fewer chunks) to **avoid unnecessary overhead**. For the plots above, we adjust the number of loss chunks to reflect realistic VRAM tiers. With 80GB VRAM, this yields >3.2× longer contexts.
{% endhint %}
### 🏁 Unsloth Gradient Checkpointing Enhancements
Our [Unsloth Gradient Checkpointing](https://unsloth.ai/blog/long-context) algorithm, **introduced in April 2024**, quickly became popular and the standard across the industry, having been integrated into most training packages nowadays. It offloads activations to CPU RAM which allowed 10x longer context lengths. Our new enhancements uses CUDA Streams and other tricks to add at most **0.1%** training overhead with no impact on accuracy. Previously it added 1 to 3% training overhead.
{% code expandable="true" %}
```python
# Original Unsloth version released April 2024 - LGPLv3 Licensed
class Unsloth_Offloaded_Gradient_Checkpointer(torch.autograd.Function):
@staticmethod
@torch_amp_custom_fwd
def forward(ctx, forward_function, hidden_states, *args):
ctx.device = hidden_states.device
saved_hidden_states = hidden_states.to("cpu", non_blocking = True)
with torch.no_grad():
output = forward_function(hidden_states, *args)
ctx.save_for_backward(saved_hidden_states)
ctx.forward_function, ctx.args = forward_function, args
return output
@staticmethod
@torch_amp_custom_bwd
def backward(ctx, dY):
(hidden_states,) = ctx.saved_tensors
hidden_states = hidden_states.to(ctx.device, non_blocking = True).detach()
hidden_states.requires_grad_(True)
with torch.enable_grad():
(output,) = ctx.forward_function(hidden_states, *ctx.args)
torch.autograd.backward(output, dY)
return (None, hidden_states.grad,) + (None,)*len(ctx.args)
```
{% endcode %}
By offloading activations as soon as they are produced, we minimize peak activation footprint and free GPU memory exactly when it’s needed. This sharply reduces memory pressure in long-context or large-batch training, where a single decoder layer’s activations can exceed 2 GB.
> **Thus, Unsloth’s new algorithms & Gradient Checkpointing contributes to most improvements (3.2x), enabling 290k-context-length QLoRA GPT-OSS fine-tuning on a single H100.**
### 🔓 Tiled MLP: Unlocking 500K+
With help from [Stas Bekman](https://x.com/StasBekman) (Snowflake), we integrated Tiled MLP from Snowflake’s Arctic Long Sequence Training [paper](https://arxiv.org/abs/2506.13996) and blog post. TiledMLP reduces activation memory and enables much longer sequence lengths by tiling hidden states along the sequence dimension before heavy MLP projections.
**We also introduce a few quality-of-life improvements:**
We preserve RNG state across tiled forward recomputations so dropout and other stochastic ops are consistent between forward and backward replays. This keeps nested checkpointed computations stable and numerically identical.
{% hint style="success" %}
Our implementation auto patches any module named or typed as `mlp`, so **nearly all models with MLP modules are supported out of the box for Tiled MLP.**
{% endhint %}
**Tradeoffs to keep in mind**
TiledMLP saves VRAM at the cost of extra forward passes. Because it lives inside a checkpointed transformer block and is itself written in a checkpoint style, it effectively becomes a nested checkpoint: one **MLP now performs \~3 forward passes and 1 backward pass per step**. In return, we can drop almost all intermediate MLP activations from VRAM while still supporting extremely long sequences.
The plots compare active memory timelines for a single decoder layer’s forward and backward during a long-context training step, without Tiled MLP (left) and with it (right). Without Tiled MLP, peak VRAM occurs during the MLP backward; with Tiled MLP, it shifts to the fused loss calculation. We see \~40% lower VRAM usage, and because the fused loss auto chunks dynamically based on available VRAM, the peak with Tiled MLP would be even smaller on smaller GPUs.
To show cross-entropy loss is not the new bottleneck, we fix its chunk size instead of choosing it dynamically and then double the number of chunks. This significantly reduces the loss-related memory spikes. The max memory now occurs during backward in both cases, and overall timing is similar, though Tiled MLP adds a small overhead: one large GEMM becomes many sequential matmuls, plus the extra forward pass mentioned above.
Overall, the trade-off is worth it: without Tiled MLP, long-context training can require roughly 2× the memory usage, while with **Tiled MLP a single GPU pays only about a 1.3× increase in step time for the same context length.**
**Enabling Tiled MLP in Unsloth:**
```py
model, tokenizer = FastLanguageModel.from_pretrained(
...,
unsloth_tiled_mlp = True,
)
```
Just set `unsloth_tiled_mlp = True` in `from_pretrained` and Tiled MLP is enabled. We follow the same logic as the Arctic paper and choose `num_shards = ceil(seq_len/hidden_size)`. Each tile will operate on sequence lengths which are the same size of the hidden dimension of the model to balance throughput and memory savings.
