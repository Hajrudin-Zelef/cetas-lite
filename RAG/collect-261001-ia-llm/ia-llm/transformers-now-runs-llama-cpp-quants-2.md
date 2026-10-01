---
id: collect-261001-ia-llm/ia-llm/transformers-now-runs-llama-cpp-quants-2
title: "transformers-now-runs-llama-cpp-quants"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Apple", "Hugging Face"]
dates: []
keywords: ["llama", "attention", "cost", "decode", "flash attention", "gguf", "gpu", "inference", "llama.cpp", "memory", "moe", "multimodal"]
source: docs/RAG/collect-261001-ia-llm/transformers-now-runs-llama-cpp-quants.md
source_anchor: ""
source_lines: [127, 193]
sha256: adb3d5cdbdbcd715e6637c632ea772c308e52342366a391e38cefe51b8c28755
---

# transformers-now-runs-llama-cpp-quants

When GGML and llama.cpp joined Hugging Face, we described their complementary roles: llama.cpp provides a foundation for local inference, while transformers provides a foundation for model definition. GGUF support brings those two closer together.

**llama.cpp remains our recommended engine when your priority is efficient local inference.** Its dedicated runtime, memory management, and broad hardware support are built around that goal. This integration gives developers a convenient way to work with the same GGUF checkpoints inside transformers:

- **Experiment with GGUF in Python and PyTorch.** Inspect intermediate activations with hooks, modify a model's forward pass, or prototype custom layers using familiar PyTorch tools.
- **Evaluate GGUF models.** Use your existing transformers evaluation workflows to measure the quality of quantized checkpoints.
- **Validate GGUF conversions.** For us as developers, loading the original checkpoint and its GGUF conversion in transformers makes it easier to check that the weights were converted correctly, accounting for quantization error.
- **Try new decoding ideas.** Use custom logits processors and stopping criteria with`generate` , or write your own generation loop in Python.
- **Fine-tune from a GGUF checkpoint.** Dequantize the weights and continue with a standard transformers training workflow.

For that last case, use `GgufConfig(dequantize=True)`:

```
import torch
from transformers import AutoModelForCausalLM, GgufConfig
model = AutoModelForCausalLM.from_pretrained(
    "unsloth/Qwen3.5-4B-GGUF",
    gguf_file="Qwen3.5-4B-Q4_K_M.gguf",
    quantization_config=GgufConfig(dequantize=True),
    dtype=torch.bfloat16,
)
```
**The bigger opportunity is bringing ggml's performance to models that llama.cpp does not support.**

transformers already provides the PyTorch implementations of these architectures. With ggml kernels and quantization schemes available in PyTorch, we can work toward accelerating their supported operations without first implementing the entire model in llama.cpp. This is especially useful for new architectures, research models, and custom variants that may never receive a dedicated llama.cpp implementation.

That opportunity extends beyond the GGUF format itself. A kernel operates on tensors; it does not require the whole model to come from a GGUF file. The same building blocks can be integrated into other transformers models and loading workflows. This also opens a path to other modalities: computer vision models, audio models, and multimodal models could reuse compatible attention, normalization, and matrix multiplication kernels without first having a full implementation in llama.cpp. Each architecture still needs integration and validation; the initial GGUF examples here cover text generation.

We also wanted to show how far we can get while keeping the model and generation loop in Python. **With the right kernels and an efficient generation loop, Python and PyTorch can deliver strong local inference performance.** The kernels handle the heavy computation, while the generation loop keeps the GPU busy by avoiding unnecessary synchronization.

Our focus was to make eager execution fast without requiring `torch.compile`. For interactive use, we wanted a quick start and a steady stream of tokens, without compilation pauses or recompilation when input shapes change. The two main pieces of that work are the kernels and `generate` itself.

A kernel is a small program that performs an operation on the GPU. PyTorch supplies general-purpose implementations; a specialized kernel can do less work, combine several operations, or read quantized weights directly in their stored format.

The `kernels` library lets us distribute compatible builds of ggml's Metal kernels on the Hub and call them from transformers. That brings ggml's work into the PyTorch model without replacing the model with a separate inference runtime.

| Kernel | What it does | 
|---|---|
| `ggml-quantization` | Reads packed quantized weights for matrix operations, including the selected experts in an MoE model. It avoids expanding the whole weight matrix before each decode operation. | 
| `ggml-norm` | Fuses normalization operations, including the zero-centered RMSNorm used by Qwen3.5 and Qwen3.8. | 
| `ggml-attn` | Provides ggml's Metal flash attention for prompt processing and token decoding. | 
| `ggml-gated-delta-net` | Accelerates the gated delta network used in the linear-attention layers of the Qwen3.5 and Qwen3.8 hybrid architectures. | 
| `topk` | Selects the experts for each token in an MoE model, combining softmax and top-k routing. This is our own Metal implementation. | 

The first four packages build on ggml's kernels; the top-k kernel addresses a separate bottleneck in MoE routing. Together they reduce the GPU work needed for each generated token.

To show the contribution of the layer kernels, we compare the same packed GGUF checkpoints with and without them. The quantization kernel stays enabled in both configurations: disabling it would also change how weights are represented and would measure a different tradeoff.

Faster kernels only help if the GPU has work to do. During generation, the CPU schedules GPU operations and controls the loop that produces the next token. Reading a result back from the GPU can force the CPU to wait until queued operations finish. Repeating even a small wait for every token can noticeably reduce throughput.

Two changes address this in `generate`, which results in improvements for all transformers models (not just when running GGUF files):

- **Drop an unnecessary attention mask early (#48814).** When a supported decoder-only input has no padding, its all-ones padding mask can be removed at the start of generation. Downstream attention code no longer needs to inspect that mask repeatedly to determine whether it can be skipped. Causal attention is still preserved.
- **Defer the stopping check (#47975).** On supported paths,`generate` copies the stopping decision asynchronously and consumes it on the following step. The CPU can keep scheduling work while the GPU runs. Streaming tokens use the same approach, and any extra step past the stopping condition is removed from the result.

These changes improve the generation loop around the model, so their usefulness extends beyond GGUF. They complement the kernel work: kernels reduce the cost of an operation, while fewer synchronization points let CPU scheduling and GPU execution overlap.

These measurements keep all layer kernels enabled; the bars isolate the changes to the generation loop.

The initial target is a single interactive conversation on Apple Silicon. There are a few boundaries to keep in mind:

- **The packed inference path is MPS-only for now.** GGUF import through dequantization remains a separate option; support for the file format does not imply that packed kernels are available on every device.
- **Padding and batching still need work.** Unpadded inputs benefit from the mask optimization described above. Padded batches cannot take the same shortcut and can have lower performance. We want to extend the work to`generate_batch` on MPS.
- **Architecture coverage is limited.** The packed loader currently covers the Qwen3.5 dense and MoE architectures, including compatible Qwen3.8 checkpoints. Adding support for other architectures is relatively straightforward, and we’ll expand coverage gradually.

If you have a GGUF model you would like to use in transformers, open an issue with the checkpoint and your use case. That will help us prioritize support for the models people are running locally.

