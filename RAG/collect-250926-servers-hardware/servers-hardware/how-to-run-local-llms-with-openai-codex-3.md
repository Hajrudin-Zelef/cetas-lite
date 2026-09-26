---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-llms-with-openai-codex-3
title: "How to Run Local LLMs with OpenAI Codex"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Apple", "Google", "OpenAI", "Unsloth"]
dates: []
keywords: ["agent", "agentic", "agents", "gguf", "gpu", "inference", "kv cache", "llama", "llama.cpp", "parameters", "quantization", "tool calling"]
source: docs/RAG/clean4/How to Run Local LLMs with OpenAI Codex.md
source_anchor: ""
source_lines: [208, 314]
sha256: a1bd95d98ab500c3ea5d850a13ddc383abb9057c5520a8301568f6a577aefee8
---

# How to Run Local LLMs with OpenAI Codex

| Quit (Ctrl+C), then re-launch with `codex --oss --profile unsloth_api`. Custom providers skip that |
| Tool calling unreliable | Need self-healing fallback | Unsloth's [self-healing tool calls](file:///1382377/new/studio/#execute-code--heal-tool-calling) are on by default |
| WSL: `Connection refused` to `localhost` | WSL network namespace | Use the Windows host IP in `base_url`, or enable WSL2 mirrored networking |
## 🦙 Llama.cpp Tutorial
We can also use `llama.cpp` directly. We need to deploy `llama-server` which is an open-source framework for running and serving LLMs efficiently on Mac, Linux and Windows devices. The model will be served on **port 8001** with all agent tool calls routed through that single OpenAI-compatible endpoint.
{% hint style="info" %}
The llama.cpp endpoint will be on **port 8001** instead of `8888` (Unsloth Studio's default). Adjust your Codex `base_url` accordingly in `~/.codex/config.toml`.
{% endhint %}
{% stepper %}
{% step %}
#### **Install llama.cpp**
We need to install `llama.cpp` to deploy/serve local LLMs to use in Codex. We follow the official build instructions for correct GPU bindings and maximum performance. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.
```bash
apt-get update
apt-get install pciutils build-essential cmake curl libcurl4-openssl-dev git-all -y
git clone https://github.com/ggml-org/llama.cpp
cmake llama.cpp -B llama.cpp/build \
-DBUILD_SHARED_LIBS=OFF -DGGML_CUDA=ON
cmake --build llama.cpp/build --config Release -j --clean-first \
--target llama-cli llama-mtmd-cli llama-server llama-gguf-split
cp llama.cpp/build/bin/llama-* llama.cpp
```
{% endstep %}
{% step %}
#### **Download and use models locally**
Download the model via the `hf` CLI (`pip install huggingface_hub hf_transfer`). We use the **UD-Q4\_K\_XL** quant for the best size/accuracy balance. You can find all Unsloth GGUF uploads in our [Collection here](file:///1382377/get-started/unsloth-model-catalog.md). If downloads get stuck, see [https://hugging-face-hub-xet-debugging.md](https://hugging-face-hub-xet-debugging.md "mention").
```bash
hf download unsloth/gemma-4-26B-A4B-it-GGUF \
--local-dir unsloth/gemma-4-26B-A4B-it-GGUF \
--include "*UD-Q4_K_XL*"
```
{% hint style="info" %}
**Want vision support?** Add `--include "*mmproj-BF16*"` to also pull the vision projector, then pass `--mmproj unsloth/gemma-4-26B-A4B-it-GGUF/mmproj-BF16.gguf` to `llama-server`. Codex itself is text-only, so this is optional.
{% endhint %}
{% hint style="success" %}
We used `unsloth/gemma-4-26B-A4B-it-GGUF`, but you can use anything like `unsloth/Qwen3.6-35B-A3B-GGUF` - see [Qwen3.6-35B-A3B](/docs/models/qwen3.6.md).
{% endhint %}
{% endstep %}
{% step %}
#### **Start the Llama-server**
To deploy Gemma-4-26B-A4B for agentic workloads, we use `llama-server`. We apply Google's recommended sampling parameters (`temp 1.0`, `top_p 0.95`, `top_k 64`) and enable `--jinja` for proper tool calling support.
Run this command in a new terminal (use `tmux` or open a new terminal). The below should **fit comfortably in a 24GB GPU (RTX 4090)** at \~18GB. `--fit on` will also auto offload, but if you see bad performance, reduce `--ctx-size`.
```bash
./llama.cpp/llama-server \
--model unsloth/gemma-4-26B-A4B-it-GGUF/gemma-4-26B-A4B-it-UD-Q4_K_XL.gguf \
--alias "unsloth/gemma-4-26B-A4B" \
--temp 1.0 \
--top-p 0.95 \
--top-k 64 \
--port 8001 \
--kv-unified \
--cache-type-k q8_0 --cache-type-v q8_0 \
--batch-size 4096 --ubatch-size 1024
```
{% hint style="info" %}
We used `--cache-type-k q8_0 --cache-type-v q8_0` for KV cache quantization to reduce VRAM use. If you see reduced quality, use `bf16` instead (`--cache-type-k bf16 --cache-type-v bf16`), but VRAM doubles.
{% endhint %}
{% hint style="success" %}
**Disabling thinking** can improve performance for agentic coding tasks. Gemma 4 enables thinking by default via the chat template - to disable it, add the following flag to the llama-server command:
**MacOS / Linux / WSL:**
`--chat-template-kwargs '{"enable_thinking":false}'`
**Windows PowerShell:**
`--chat-template-kwargs "{\"enable_thinking\":false}"`
{% endhint %}
{% endstep %}
{% step %}
#### **Point Codex at port 8001**
Edit your `~/.codex/config.toml` to use the llama-server port:
{% code title="\~/.codex/config.toml" %}
```toml
[model_providers.llama_cpp]
name = "llama.cpp"
base_url = "http://localhost:8001/v1"
env_key = "LLAMA_CPP_API_KEY"
wire_api = "responses"
```
{% endcode %}
Then launch with the new profile:
```bash
codex --oss llama_cpp
```
Since llama-server doesn't require a real key, you can set the auth token to anything:
{% code title="MacOS / Linux / WSL" %}
```bash
export LLAMA_CPP_API_KEY=sk-no-key-required
```
{% endcode %}
{% code title="Windows PowerShell" %}
```powershell
$env:LLAMA_CPP_API_KEY = "sk-no-key-required"
```
{% endcode %}
{% endstep %}
{% endstepper %}
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/basics/codex.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
