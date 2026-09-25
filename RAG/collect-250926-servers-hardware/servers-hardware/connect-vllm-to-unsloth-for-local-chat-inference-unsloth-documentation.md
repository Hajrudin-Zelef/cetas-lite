---
id: collect-250926-servers-hardware/servers-hardware/connect-vllm-to-unsloth-for-local-chat-inference-unsloth-documentation
title: "connect-vllm-to-unsloth-for-local-chat-inference-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["OpenAI", "Unsloth", "vLLM"]
dates: []
keywords: ["inference", "vllm", "gpu", "memory", "reasoning"]
source: docs/RAG/clean4/connect-vllm-to-unsloth-for-local-chat-inference-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 39]
sha256: b4223e04d4a6ea61eaf3dabb708598e57dfe1b0873dd04c4fdaeba5d2c61d6ee
---

# connect-vllm-to-unsloth-for-local-chat-inference-unsloth-documentation

Learn how to connect **vLLM to** **Unsloth** using vLLM’s **OpenAI-compatible API** so you can serve models and chat with them locally inside a open-source UI chat interface. This guide walks through installing vLLM, launching a local vLLM server, configuring the API base URL, loading available model IDs, and selecting your hosted vLLM model.

By the end, your vLLM-served models will appear alongside local models, giving you a fast and flexible way to run external LLM inference from a UI chat interface.

Install vLLM first so you can run the `vllm serve` command. Follow the official vLLM install guide for your platform and hardware.

After installing, check that vLLM works in your terminal: `vllm --help`

Open **Settings → Connections**, then click **Add Connection**.

Select **vLLM**, then enter your server details.

Enter your vLLM server details:

- **API key:** leave empty unless you started vLLM with --api-key
- **Base URL:** for example, http://localhost:8000/v1
- **Reasoning model:** enable this if the served model supports thinking
- **Model IDs:** click**Load Models** , or enter custom IDs manually

After you click **Add Connection**, the models you enabled will appear under **Connection** in the model dropdown.

The example above uses the core serving settings. You can add more vllm serve arguments depending on your model and hardware.

Common options include:

```
vllm serve unsloth/gemma-4-26B-A4B-it \
  --dtype auto \
  --host 0.0.0.0 \
  --port 8000 \
  --api-key token-abc123 \
  --max-model-len 8192 \
  --gpu-memory-utilization 0.9
```
For the full list of vLLM server arguments, see the official vLLM OpenAI-compatible server docs.

Last updated

Was this helpful?
