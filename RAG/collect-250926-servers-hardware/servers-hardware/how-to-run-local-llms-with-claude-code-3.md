---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-llms-with-claude-code-3
title: "How to Run Local LLMs with Claude Code"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Hugging Face", "OpenAI", "Unsloth", "Z.ai"]
dates: []
keywords: ["claude", "agent", "agentic", "gguf", "glm", "gpu", "inference", "kv cache", "llama", "llama.cpp", "mcp", "parameters"]
source: docs/RAG/clean4/How to Run Local LLMs with Claude Code.md
source_anchor: ""
source_lines: [239, 361]
sha256: f9e059d3bc2eb5556882bbd56acb88b7d6ad3acbbc292223495a526bce28e9d7
---

# How to Run Local LLMs with Claude Code

After you submit the prompt, the agent will search the web, evaluate findings, and write the final report. This may take a few minutes.
Some workflows may require you to approve actions or answer follow up prompts.
{% hint style="info" %}
Some workflows may require you to approve actions or answer follow-up prompts.
{% endhint %}
Once complete, the generated `sft_report.md` will look similar to this.
{% hint style="warning" %}
If you see `Unable to connect to API (ConnectionRefused)` , remember to unset `ANTHROPIC_BASE_URL` via `unset ANTHROPIC_BASE_URL`
If you find open models to be 90% slower, [see here first](#fixing-90-slower-inference-in-claude-code) to fix KV cache being invalidated.
{% endhint %}
### Optional: shrink the system prompt
Claude Code was built for Anthropic's hosted models, so its default system prompt is large. On local models you can trim it for faster responses and better KV-cache reuse by adding two flags when you launch:
{% code overflow="wrap" %}
```shellscript
claude --model unsloth/gemma-4-26B-A4B-it-GGUF --bare --exclude-dynamic-system-prompt-sections
```
{% endcode %}
{% hint style="info" %}
`--bare` skips auto-discovery of hooks, skills, plugins, MCP servers and CLAUDE.md (Claude keeps Bash and file read/edit), and `--exclude-dynamic-system-prompt-sections` moves per-machine sections out of the prompt prefix. Both shrink the prompt and improve KV-cache reuse, which makes local models noticeably faster. They are optional and do not change the connection setup above.
{% endhint %}
### Optional: tune the Unsloth server
Claude Code uses the model running in Unsloth. You can customize how the server behaves when starting it.
```bash
# Serve for a coding agent: --disable-tools passes the agent's own tools through
unsloth run \
  --model unsloth/gemma-4-26B-A4B-it-GGUF:UD-Q4_K_XL \
  --disable-tools \
  --reasoning off \
  -p 8888
```
{% hint style="warning" %}
Use `--disable-tools` when driving Claude Code (or any external coding agent). By default Unsloth Studio runs its own server-side tools, which swallows the agent's tool calls, so Claude Code answers but never edits files. `--disable-tools` switches to passthrough, so Claude Code's own Write/Edit/Bash tools are used.
{% endhint %}
Use `--reasoning off` to turn thinking off, or `--reasoning on` to turn it on for models that support reasoning.
```bash
# Expose the API on your local network
unsloth run \
  --model unsloth/gemma-4-26B-A4B-it-GGUF:UD-Q4_K_XL \
  -H 0.0.0.0 \
  -p 8888
```
This starts the server on `0.0.0.0:8888`, allowing other devices on your local network to connect.
Use `-p` to change which port the server runs on. Use `-H 0.0.0.0` if you want phones, laptops, or other devices on your network to connect.
For more advanced runtime configuration, see the main [API tuning](https://unsloth.ai/docs/basics/api#unsloth-run-command) section.
## 🦙 Llama.cpp Tutorial
Before we begin, we firstly need to complete setup for the specific model you're going to use. We use `llama.cpp` which is an open-source framework for running LLMs on your Mac, Linux, Windows etc. devices. Llama.cpp contains `llama-server` which allows you to serve and deploy LLMs efficiently. The model will be served on port 8001, with all agent tools routed through a single OpenAI-compatible endpoint.
#### Qwen3.5 Tutorial
We'll be using [Qwen3.5](/docs/models/qwen3.5.md)-35B-A3B and specific settings for fast accurate coding tasks. If you don't have enough VRAM and want a **smarter** model, **Qwen3.5-27B** is a great choice, but it will be \~2x slower, or you can use other Qwen3.5 variants like 9B, 4B or 2B.
{% hint style="info" %}
Use Qwen3.5-27B if you want a **smarter** model or if you don't have enough VRAM. It will be \~2x slower than 35B-A3B however. Or you can use [**Qwen3-Coder-Next**](/docs/models/qwen3-coder-next.md) which is fantastic if you have enough VRAM.
{% endhint %}
{% stepper %}
{% step %}
#### Install llama.cpp
We need to install `llama.cpp` to deploy/serve local LLMs to use in Claude Code etc. We follow the official build instructions for correct GPU bindings and maximum performance. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.
```bash
apt-get update
apt-get install pciutils build-essential cmake curl libcurl4-openssl-dev git-all -y
git clone https://github.com/ggml-org/llama.cpp
cmake llama.cpp -B llama.cpp/build \
    -DBUILD_SHARED_LIBS=OFF -DGGML_CUDA=ON
cmake --build llama.cpp/build --config Release -j --clean-first --target llama-cli llama-mtmd-cli llama-server llama-gguf-split
cp llama.cpp/build/bin/llama-* llama.cpp
```
{% endstep %}
{% step %}
#### Download and use models locally
Download the model via `huggingface_hub` in Python (after installing via `pip install huggingface_hub hf_transfer`). We use the **UD-Q4\_K\_XL** quant for the best size/accuracy balance. You can find all Unsloth GGUF uploads in our [Collection here](/docs/get-started/unsloth-model-catalog.md). If downloads get stuck, see [Hugging Face Hub, XET debugging](/docs/basics/troubleshooting-and-faqs/hugging-face-hub-xet-debugging.md)
```bash
hf download unsloth/Qwen3.5-35B-A3B-GGUF \
    --local-dir unsloth/Qwen3.5-35B-A3B-GGUF \
    --include "*UD-Q4_K_XL*" # Use "*UD-Q2_K_XL*" for Dynamic 2bit
```
{% hint style="success" %}
We used `unsloth/Qwen3.5-35B-A3B-GGUF` , but you can use another variant like 27B or any other model like `unsloth/`[`Qwen3-Coder-Next`](/docs/models/qwen3-coder-next.md)`-GGUF`.
{% endhint %}
{% endstep %}
{% step %}
#### Start the Llama-server
To deploy Qwen3.5 for agentic workloads, we use `llama-server`. We apply [Qwen's recommended sampling parameters](/docs/models/qwen3.5.md#recommended-settings) for thinking mode: `temp 0.6`, `top_p 0.95` , `top-k 20`. Keep in mind these numbers change if you use non-thinking mode or other tasks.
Run this command in a new terminal (use `tmux` or open a new terminal). The below should **fit perfectly in a 24GB GPU (RTX 4090) (uses 23GB)** `--fit on` will also auto offload, but if you see bad performance, reduce `--ctx-size` .
{% hint style="info" %}
We used `--cache-type-k q8_0 --cache-type-v q8_0` for KV cache quantization for less VRAM usage. For full precision, use `--cache-type-k bf16 --cache-type-v bf16` .Note bf16 KV Cache might be slightly slower on some machines.
{% endhint %}
```bash
./llama.cpp/llama-server \
    --model unsloth/Qwen3.5-35B-A3B-GGUF/Qwen3.5-35B-A3B-UD-Q4_K_XL.gguf \
    --alias "unsloth/Qwen3.5-35B-A3B" \
    --temp 0.6 \
    --top-p 0.95 \
    --top-k 20 \
    --min-p 0.00 \
    --port 8001 \
    --kv-unified \
    --cache-type-k q8_0 --cache-type-v q8_0
```
{% hint style="success" %}
You can also disable thinking for Qwen3.5 which can improve performance for agentic coding stuff. To disable thinking with llama.cpp add this to the llama-server command:
`--chat-template-kwargs "{\"enable_thinking\": false}"`
{% endhint %}
{% endstep %}
{% endstepper %}
### Start Claude Code with llama-server
{% hint style="success" %}
We used `unsloth/GLM-4.7-Flash-GGUF` , but you can use anything like `unsloth/Qwen3.6-27B-GGUF`.
{% endhint %}
{% hint style="warning" %}
See [#fixing-90-slower-inference-in-claude-code](#fixing-90-slower-inference-in-claude-code "mention") first to fix open models being 90% slower due to KV Cache invalidation.
{% endhint %}
Navigate to your project folder (`mkdir project ; cd project`) and run:
```bash
claude --model unsloth/GLM-4.7-Flash
```
To use Qwen3.6-35B-A3B, simply change it to:
```bash
claude --model unsloth/Qwen3.6-35B-A3B
```
To set Claude Code to execute commands without any approvals do **(BEWARE this will make Claude Code do and execute code however it likes without any approvals!)**
{% code overflow="wrap" %}
```bash
claude --model unsloth/GLM-4.7-Flash --dangerously-skip-permissions
```
{% endcode %}
