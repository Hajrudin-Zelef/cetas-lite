---
id: collect-240926-huggingface/huggingface/unsloth-gpt-oss-20b-gguf-hugging-face
title: "Read our How to Run gpt-oss Guide here!"
domain: huggingface
role: reference
task: reference
actors: ["Google", "Hugging Face", "OpenAI", "Unsloth", "vLLM"]
dates: []
keywords: ["agentic", "apache", "benchmarks", "consumer", "fine-tuning", "gguf", "gpu", "inference", "latency", "license", "llama", "llama.cpp"]
source: docs/RAG/clean_en/huggingface/unsloth-gpt-oss-20b-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 135]
sha256: d760601ceadaabd217feb727a0a9198c3a5e12e48fd509f1995e1fef48b6e48e
---

# Read our How to Run gpt-oss Guide here!

<!-- source: https://huggingface.co/unsloth/gpt-oss-20b-GGUF -->

# Read our How to Run gpt-oss Guide here!

    **See our collection for all versions of gpt-oss including GGUF, 4-bit & 16-bit formats.**
  

    *Learn to run gpt-oss correctly - Read our Guide.*
  

   *See Unsloth Dynamic 2.0 GGUFs for our quantization benchmarks.*
  

# ✨ Read our gpt-oss Guide here!

- Fine-tune gpt-oss-20b for free using our Google Colab notebook
- Read our Blog about gpt-oss support: unsloth.ai/blog/gpt-oss
- View the rest of our notebooks in our docs here.
- Thank you to the llama.cpp team for their work on supporting this model. We wouldn't be able to release quants without them!

The F32 quant is MXFP4 upcasted to BF16 for every single layer and is unquantized.


  **Try gpt-oss** ·
  **Guides** ·
  **System card** ·
  **OpenAI blog**

Welcome to the gpt-oss series, OpenAI’s open-weight models designed for powerful reasoning, agentic tasks, and versatile developer use cases.

We’re releasing two flavors of the open models:

- `gpt-oss-120b` — for production, general purpose, high reasoning use cases that fits into a single H100 GPU (117B parameters with 5.1B active parameters)
- `gpt-oss-20b` — for lower latency, and local or specialized use cases (21B parameters with 3.6B active parameters)

Both models were trained on our harmony response format and should only be used with the harmony format as it will not work correctly otherwise.

This model card is dedicated to the smaller `gpt-oss-20b` model. Check out `gpt-oss-120b` for the larger model.


- **Permissive Apache 2.0 license:** Build freely without copyleft restrictions or patent risk—ideal for experimentation, customization, and commercial deployment.
- **Configurable reasoning effort:** Easily adjust the reasoning effort (low, medium, high) based on your specific use case and latency needs.
- **Full chain-of-thought:** Gain complete access to the model’s reasoning process, facilitating easier debugging and increased trust in outputs. It’s not intended to be shown to end users.
- **Fine-tunable:** Fully customize models to your specific use case through parameter fine-tuning.
- **Agentic capabilities:** Use the models’ native capabilities for function calling, web browsing, Python code execution, and Structured Outputs.
- **Native MXFP4 quantization:** The models are trained with native MXFP4 precision for the MoE layer, making`gpt-oss-120b` run on a single H100 GPU and the`gpt-oss-20b` model run within 16GB of memory.

You can use `gpt-oss-120b` and `gpt-oss-20b` with Transformers. If you use the Transformers chat template, it will automatically apply the harmony response format. If you use `model.generate` directly, you need to apply the harmony format manually using the chat template or use our openai-harmony package.

To get started, install the necessary dependencies to setup your environment:

```
pip install -U transformers kernels torch 
```
Once, setup you can proceed to run the model by running the snippet below:

```
from transformers import pipeline
import torch
model_id = "openai/gpt-oss-20b"
pipe = pipeline(
    "text-generation",
    model=model_id,
    torch_dtype="auto",
    device_map="auto",
)
messages = [
    {"role": "user", "content": "Explain quantum mechanics clearly and concisely."},
]
outputs = pipe(
    messages,
    max_new_tokens=256,
)
print(outputs[0]["generated_text"][-1])
```
Alternatively, you can run the model via `Transformers Serve` to spin up a OpenAI-compatible webserver:

```
transformers serve
transformers chat localhost:8000 --model-name-or-path openai/gpt-oss-20b
```
vLLM recommends using uv for Python dependency management. You can use vLLM to spin up an OpenAI-compatible webserver. The following command will automatically download the model and start the server.

```
uv pip install --pre vllm==0.10.1+gptoss \
    --extra-index-url https://wheels.vllm.ai/gpt-oss/ \
    --extra-index-url https://download.pytorch.org/whl/nightly/cu128 \
    --index-strategy unsafe-best-match
vllm serve openai/gpt-oss-20b
```
To learn about how to use this model with PyTorch and Triton, check out our reference implementations in the gpt-oss repository.

If you are trying to run gpt-oss on consumer hardware, you can use Ollama by running the following commands after installing Ollama.

```
# gpt-oss-20b
ollama pull gpt-oss:20b
ollama run gpt-oss:20b
```
If you are using LM Studio you can use the following commands to download.

```
# gpt-oss-20b
lms get openai/gpt-oss-20b
```
Check out our awesome list for a broader collection of gpt-oss resources and inference partners.

You can download the model weights from the Hugging Face Hub directly from Hugging Face CLI:

```
# gpt-oss-20b
huggingface-cli download openai/gpt-oss-20b --include "original/*" --local-dir gpt-oss-20b/
pip install gpt-oss
python -m gpt_oss.chat model/
```
You can adjust the reasoning level that suits your task across three levels:

- **Low:** Fast responses for general dialogue.
- **Medium:** Balanced speed and detail.
- **High:** Deep and detailed analysis.

The reasoning level can be set in the system prompts, e.g., "Reasoning: high".

The gpt-oss models are excellent for:

- Web browsing (using built-in browsing tools)
- Function calling with defined schemas
- Agentic operations like browser tasks

Both gpt-oss models can be fine-tuned for a variety of specialized use cases.

This smaller model `gpt-oss-20b` can be fine-tuned on consumer hardware, whereas the larger `gpt-oss-120b` can be fine-tuned on a single H100 node.

- Downloads last month
- 527,966
