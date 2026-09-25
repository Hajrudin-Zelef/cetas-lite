---
id: collect-250926-servers-hardware/servers-hardware/connect-llama-cpp-to-unsloth-run-ggufs-with-llama-server-unsloth-documentation
title: "connect-llama-cpp-to-unsloth-run-ggufs-with-llama-server-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Unsloth"]
dates: []
keywords: ["gguf", "llama", "cost", "inference", "inference engine", "latency", "llama.cpp"]
source: docs/RAG/clean4/connect-llama-cpp-to-unsloth-run-ggufs-with-llama-server-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 39]
sha256: 3b31180b1d0ee6cd3c240f16a42826668d744f977e915d439a772573b70b30f4
---

# connect-llama-cpp-to-unsloth-run-ggufs-with-llama-server-unsloth-documentation

Llama.cpp is an open-source inference engine for running GGUF models efficiently on local hardware, and Unsloth makes it easy to run those models directly into a open-source UI chat interface. By starting a local `llama-server`, you can serve a GGUF model from your machine or Hugging Face, connect it to Unsloth, and use it like any other external chat model.

This guide walks through installing llama.cpp, launching `llama-server`, connecting it to Unsloth, enabling your model, and configuring prompt caching, context length, API keys, FA, and chat templates.

llama-server can load a local .gguf file or download a GGUF model from Hugging Face.

To serve a Hugging Face GGUF repo directly, use the repo and quant name:

`llama-server -hf unsloth/Qwen3.6-27B-GGUF:UD-Q4_K_XL`

If you wish to load a local model, you can also follow the steps below.

Start `llama-server` with the model you want to serve:

This exposes an API endpoint at: `http://localhost:8080/v1`

To require an API key, add:

Open **Settings → Connections**, then click **Add Connection**. Select **llama.cpp**, then enter your server details:

if you did not start llama-server with `--api-key`, leave the API key field empty. Enter the base URL of your server, e.g. `http://localhost:8080/v1`
Click **Load Models** to fetch available model IDs, or enter model IDs manually if your server does not expose `/models`.

Then, after you click **Add Connection,** The models you enabled will now appear under **Connected** in the **Select Model** dropdown.

Prompt caching reduces latency and cost when requests reuse the same long prefix. Use the **Prompt caching** setting in the Unsloth side panel to control caching behaviour for supported connections.

With llama.cpp, prompt caching is enabled by default and can be disabled when starting
`llama-server` with:

The example above only uses the required connection settings. You can add more llama-server arguments depending on your model and hardware.

Common options include:

For the full list of server arguments, see the official llama.cpp server README.

Last updated

Was this helpful?
