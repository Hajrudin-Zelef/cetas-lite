---
id: collect-240926-huggingface/huggingface/openai-gpt-oss-20b-hugging-face
title: "gpt-oss-20b"
domain: huggingface
role: reference
task: reference
actors: ["AMD", "Hugging Face", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["agentic", "amd", "apache", "consumer", "fine-tuning", "gpu", "inference", "latency", "leaderboard", "license", "memory", "moe"]
source: docs/RAG/clean_en/huggingface/openai-gpt-oss-20b-hugging-face.md
source_anchor: ""
source_lines: [1, 138]
sha256: ad65b2ee6977bad82ad00f9d8ae7a5095c33cf4da60fa704e5c6538f6515a77c
---

# gpt-oss-20b

<!-- source: https://huggingface.co/openai/gpt-oss-20b -->

**Try gpt-oss** ·
  **Guides** ·
  **Model card** ·
  **OpenAI blog**

Welcome to the gpt-oss series, OpenAI’s open-weight models designed for powerful reasoning, agentic tasks, and versatile developer use cases.

We’re releasing two flavors of these open models:

- `gpt-oss-120b` — for production, general purpose, high reasoning use cases that fit into a single 80GB GPU (like NVIDIA H100 or AMD MI300X) (117B parameters with 5.1B active parameters)
- `gpt-oss-20b` — for lower latency, and local or specialized use cases (21B parameters with 3.6B active parameters)

Both models were trained on our harmony response format and should only be used with the harmony format as it will not work correctly otherwise.

This model card is dedicated to the smaller `gpt-oss-20b` model. Check out `gpt-oss-120b` for the larger model.


- **Permissive Apache 2.0 license:** Build freely without copyleft restrictions or patent risk—ideal for experimentation, customization, and commercial deployment.
- **Configurable reasoning effort:** Easily adjust the reasoning effort (low, medium, high) based on your specific use case and latency needs.
- **Full chain-of-thought:** Gain complete access to the model’s reasoning process, facilitating easier debugging and increased trust in outputs. It’s not intended to be shown to end users.
- **Fine-tunable:** Fully customize models to your specific use case through parameter fine-tuning.
- **Agentic capabilities:** Use the models’ native capabilities for function calling, web browsing, Python code execution, and Structured Outputs.
- **MXFP4 quantization:** The models were post-trained with MXFP4 quantization of the MoE weights, making`gpt-oss-120b` run on a single 80GB GPU (like NVIDIA H100 or AMD MI300X) and the`gpt-oss-20b` model run within 16GB of memory. All evals were performed with the same MXFP4 quantization.

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

```
@misc{openai2025gptoss120bgptoss20bmodel,
      title={gpt-oss-120b & gpt-oss-20b Model Card}, 
      author={OpenAI},
      year={2025},
      eprint={2508.10925},
      archivePrefix={arXiv},
      primaryClass={cs.CL},
      url={https://arxiv.org/abs/2508.10925}, 
}
```
- Downloads last month
- 6,692,738

## Spaces using openai/gpt-oss-20b 100

## Collection including openai/gpt-oss-20b

## Paper for openai/gpt-oss-20b

- Idavidrein/gpqa · Diamond leaderboard
- GPQA Diamond View evaluation results source58.59<sup>*</sup>
- Reasoning: medium, With tools View evaluation resultssource67.1<sup>*</sup>
- Reasoning: medium View evaluation resultssource66<sup>*</sup>
- SWE-bench/SWE-bench_Verified · Swe Bench Resolved leaderboard
- Reasoning: medium View evaluation resultssource53.2<sup>*</sup>
