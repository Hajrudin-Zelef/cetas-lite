---
id: collect-250926-servers-hardware/servers-hardware/sglang-deployment-inference-guide-unsloth-documentation
title: "OPTIONAL use a virtual environment"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Mistral", "Nvidia", "OpenAI", "SGLang", "Unsloth"]
dates: []
keywords: ["benchmarks", "fine-tuning", "fp8", "gguf", "gpu", "gpus", "inference", "latency", "llama", "lora", "memory", "mistral"]
source: docs/RAG/clean4/sglang-deployment-inference-guide-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 97]
sha256: 74aa14c551fd419aac5aaebb8da374ef15dbf6e8fc49e4065463e28f80b31624
---

# OPTIONAL use a virtual environment

You can serve any LLM or fine-tuned model via SGLang for low-latency, high-throughput inference. SGLang supports text, image/video model inference on any GPU setup, with support for some GGUFs.

To install SGLang and Unsloth on NVIDIA GPUs, you can use the below in a virtual environment (which won't break your other Python libraries)

```
# OPTIONAL use a virtual environment
python -m venv unsloth_env
source unsloth_env/bin/activate
# Install Rust, outlines-core then SGLang
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source $HOME/.cargo/env && sudo apt-get install -y pkg-config libssl-dev
pip install --upgrade pip && pip install uv
uv pip install "sglang" && uv pip install unsloth
```
For **Docker** setups run:

```
docker run --gpus all \
    --shm-size 32g \
    -p 30000:30000 \
    -v ~/.cache/huggingface:/root/.cache/huggingface \
    --env "HF_TOKEN=<secret>" \
    --ipc=host \
    lmsysorg/sglang:latest \
    python3 -m sglang.launch_server --model-path unsloth/Llama-3.1-8B-Instruct --host 0.0.0.0 --port 30000
```
Note if you see the below, update Rust and outlines-core as specified in SGLang

If you see a Flashinfer issue like below:

Remove the flashinfer cache via `rm -rf .cache/flashinfer` and also the directory listed in the error message ie `rm -rf ~/.cache/flashinfer`

To deploy any model like for example unsloth/Llama-3.2-1B-Instruct, do the below in a separate terminal (otherwise it'll block your current terminal - you can also use tmux):

You can then use the OpenAI Chat completions library to call the model (in another terminal or using tmux):

And you will get `2 + 2 = 4.`

After fine-tuning Fine-tuning Guide or using our notebooks at Unsloth Notebooks, you can save or deploy your models directly through SGLang within a single workflow. An example Unsloth finetuning script for eg:

**To save to 16-bit for SGLang, use:**

**To save just the LoRA adapters**, either use:

Or just use our builtin function to do that:

Below is a step-by-step tutorial with instructions for training the gpt-oss-20b using Unsloth and deploying it with SGLang. It includes performance benchmarks across multiple quantization formats.

To deploy models with FP8 online quantization which allows 30 to 50% more throughput and 50% less memory usage with 2x longer context length supports with SGLang, you can do the below:

You can also use `--kv-cache-dtype fp8_e5m2` which has a larger dynamic range which might solve FP8 inference issues if you see them. Or use our pre-quantized float8 quants listed in https://huggingface.co/unsloth/models?search=-fp8 or some are listed below:

Below is some code you can run to test the performance speed of your finetuned model:

Then in another terminal or via tmux:

You will see the benchmarking run like below:

We used a B200x1 GPU with gpt-oss-20b and got the below results (~2,500 tokens throughput)

8/1024/1024

0.40

3.59

20,718.95

2,562.87

8/8192/1024

0.42

3.74

154,459.01

2,473.84

See https://docs.sglang.ai/advanced_features/server_arguments.html for server arguments for SGLang.

You can also use SGLang in offline mode (ie not a server) inside a Python interactive environment.

SGLang also interestingly supports GGUFs! **Qwen3 MoE is still under construction, but most dense models (Llama 3, Qwen 3, Mistral etc) are supported.**

First install the latest gguf python package via:

Then for example in offline mode SGLang, you can do:

First download the specific GGUF file like below:

Then serve the specific file `Qwen3-32B-UD-Q4_K_XL.gguf` and use `--served-model-name unsloth/Qwen3-32B` and also we need the HuggingFace compatible tokenizer via `--tokenizer-path`

Last updated

Was this helpful?
