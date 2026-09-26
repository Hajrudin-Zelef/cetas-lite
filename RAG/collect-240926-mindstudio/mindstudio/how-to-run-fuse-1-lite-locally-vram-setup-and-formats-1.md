---
id: collect-240926-mindstudio/mindstudio/how-to-run-fuse-1-lite-locally-vram-setup-and-formats-1
title: "how-to-run-fuse-1-lite-locally-vram-setup-and-formats"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Apple", "vLLM"]
dates: []
keywords: ["agents", "attention", "compute", "consumer", "distillation", "fine-tuning", "gguf", "gpu", "inference", "inference engine", "llama", "llama.cpp"]
source: docs/RAG/clean_en/mindstudio/how-to-run-fuse-1-lite-locally-vram-setup-and-formats.md
source_anchor: ""
source_lines: [1, 100]
sha256: 4c81d2e8e00823c6670078d78f22c7d092a69340aeeaf08e9845473ff9b648db
---

# how-to-run-fuse-1-lite-locally-vram-setup-and-formats

<!-- source: https://www.mindstudio.ai/blog/run-fuse-1-lite-locally -->

## What is fuse-1 Lite and what does it take to run it?

fuse-1 Lite is a 5.72B parameter coding model that runs comfortably on consumer hardware, with VRAM needs ranging from about 3.36 GB in 4-bit quantized form up to roughly 12 GB in full bfloat16 precision. It combines a small 2.6B host model (LiquidAI’s LFM2.5-2.6B) with 960 coding-specialized experts pulled from a much larger Qwen model. That design keeps the footprint small while aiming to preserve coding-specific capability, and it’s available across five deployment formats: native transformers, bitsandbytes quantization, GGUF, MLX, and vLLM.

## TL;DR

- **fuse-1 Lite** is a 5.72B parameter mixture-of-experts model built by transplanting 960 coding experts from Qwen3.6-35B-A3B onto a frozen LFM2.5-2.6B host, with only a lightweight router trained on top.
- **VRAM requirements scale from 3.36 GB to about 12 GB** depending on precision: 4-bit NF4 needs the least, 8-bit sits in the middle, and bfloat16 or full-precision formats need the most.
- **Five deployment paths exist** : transformers with bitsandbytes, GGUF for llama.cpp, MLX for Apple Silicon, and a vLLM plugin for high-throughput serving, each with different setup complexity.
- **GGUF and MLX require custom code** , not stock builds. The GGUF format uses a custom`fuse3` architecture that needs a patched llama.cpp fork, and MLX loading needs`trust_remote_code=True` with a custom model file.
- **The cheapest way to try it is 4-bit bitsandbytes quantization** , which fits in about 3.36 GB and runs on hardware as modest as a T4, RTX 3060, or M2 Pro.
- **A built-in toggle lets you disable the coding experts** at runtime, switching the model back to plain LFM2 behavior for general text tasks without reloading weights.
- **The model card lists training costs around $3 total** , using 300 steps of router training on 55 examples, which is unusually cheap for a model of this class.

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

## How is fuse-1 Lite built differently from a typical fine-tune?

Most small coding models are produced by distillation (training a small model to mimic a larger teacher) or by fine-tuning an existing base on code data. fuse-1 Lite does neither. It takes actual expert weights from Qwen3.6-35B-A3B, a mixture-of-experts model, and grafts 960 of them (32 per layer across 30 layers) directly onto LFM2.5-2.6B’s decoder layers as frozen residual additions.

On top of that, a small router (about 2 million trainable parameters total, including a per-layer scale factor) is trained to decide which of the transplanted experts fire for a given token. The host model and the transplanted experts stay frozen throughout. Only the router and scale values get updated, which is why training took just 300 steps and roughly 8 minutes on an L4 GPU according to the model’s documentation.

The practical result: the model behaves like LFM2.5 most of the time, and only “activates” its coding experts when it detects code-relevant tokens. The model card shows this pattern explicitly, with experts effectively disabled (scale near zero) in 19 of 30 layers, and concentrated activity in layers associated with algorithmic reasoning, logic flow, and code synthesis, peaking at layer 19.

## What VRAM do you actually need for each format?

Here’s the breakdown from the model’s published specs:

| Backend | Precision | VRAM/Memory | Hardware tier | 
|---|---|---|---|
| Transformers | bfloat16 | ~12 GB | L4, A10G, RTX 4090 | 
| Transformers | 8-bit | 6.00 GB | T4, L4, RTX 3060 | 
| Transformers | 4-bit (NF4) | 3.36 GB | T4, RTX 3060, M2 Pro | 
| vLLM | bfloat16 | ~12 GB | A10G, A100, H100 | 
| MLX | float16 | ~12 GB | M1 Pro+, M2, M3, M4 | 
| llama.cpp | F16 GGUF | ~11.4 GB | Any CPU/GPU | 
| llama.cpp | Q4_K_M GGUF | ~4 GB | Any CPU/GPU | 

The pattern is straightforward: full precision costs roughly 12 GB no matter which serving stack you pick, while quantization (either bitsandbytes NF4 on the transformers path, or Q4_K_M on the GGUF path) cuts that down to somewhere between 3 and 4 GB. If you’re VRAM-constrained, quantized GGUF or bitsandbytes 4-bit are the practical entry points. If you have a 12 GB+ card and want the best fidelity to the original weights, bfloat16 through transformers, vLLM, or MLX is the way to go.

## Is bitsandbytes quantization the easiest way to get started?

For most people with a single GPU and a standard Python environment, yes. It requires no custom builds or forks, just standard `transformers`, `torch`, and `bitsandbytes` packages.

4-bit NF4 quantization (double-quantized, bfloat16 compute dtype) brings VRAM down to 3.36 GB, which fits on entry-level cards like a T4 or RTX 3060, or even an Apple M2 Pro under unified memory. 8-bit quantization sits at 6.00 GB, offering a middle ground between quality and footprint.

Setup looks like this:

```
from transformers import AutoModelForCausalLM, AutoTokenizer, BitsAndBytesConfig
import torch
bnb_config = BitsAndBytesConfig(
    load_in_4bit=True,
    bnb_4bit_quant_type="nf4",
    bnb_4bit_compute_dtype=torch.bfloat16,
    bnb_4bit_use_double_quant=True,
)
model = AutoModelForCausalLM.from_pretrained(
    "Akahsizrr/fuse-1-Lite",
    quantization_config=bnb_config,
    device_map="auto",
    trust_remote_code=True,
)
tokenizer = AutoTokenizer.from_pretrained("Akahsizrr/fuse-1-Lite")
```
Note that `trust_remote_code=True` is required regardless of quantization level, because fuse-1 Lite ships a custom `Fuse3ForCausalLM` architecture class rather than a stock transformers model.

## One coffee. One working app.

You bring the idea. Remy manages the project.

## What’s involved in running fuse-1 Lite via GGUF or MLX?

This is where things get more involved than a typical GGUF model. fuse-1 Lite's GGUF release uses a custom `fuse3` architecture that stock llama.cpp cannot load. To run it, you need a llama.cpp fork patched with Fuse3 support, which the GGUF repository provides alongside a C++ graph builder (`src/models/fuse3.cpp`) and a Python converter script. The patched implementation reuses LFM2's existing attention and short-conv code, then bolts on the expert MoE block (router, top-k selection, SwiGLU experts, scaling, residual addition) after each augmented layer's dense feed-forward network.

Once built, inference looks like standard llama.cpp usage:

```
./llama-cli -m fuse-1-Lite-f16.gguf \
  -p "Write a Python function to check if a number is prime." \
  -n 512 --temp 0.1
```
MLX, for Apple Silicon users, follows a similar pattern. The MLX repo includes a custom model file (`fuse3_mlx.py`) that extends MLX's native LFM2 implementation with the same expert augmentation logic. Loading it requires `trust_remote_code=True`:

```
from mlx_lm import load, generate
model, tokenizer = load("Akahsizrr/fuse-1-Lite-MLX", trust_remote_code=True)
```
Both formats need roughly 11 to 12 GB at full precision (F16), or you can drop to a quantized GGUF variant like Q4_K_M for about 4 GB.

## Is vLLM serving supported, and how does the plugin work?

Yes, via a dedicated plugin rather than native support. The model card notes that vLLM's optimized inference engine doesn't natively handle the custom MoE augmentation, so a separate plugin package registers a `Fuse3ForCausalLM` model class through vLLM's `ModelRegistry` and `general_plugins` entry point. It reuses vLLM's existing LFM2 attention and short-conv layers, adding the expert MoE block after each layer's FFN, similar to the llama.cpp approach.

Setup involves installing the plugin, then serving as usual:

