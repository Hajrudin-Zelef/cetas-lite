---
id: collect-250926-servers-hardware/servers-hardware/faq-ollama-2
title: "faq-ollama"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "flash attention", "gpu", "gpus", "inference", "memory", "quantization"]
source: docs/RAG/clean4/faq-ollama.md
source_anchor: ""
source_lines: [124, 193]
sha256: 271af46ff394d757f67d3c5fc3c3f49f79bff06fc33ccb93792960089a003e43
---

# faq-ollama

By default models are kept in memory for 5 minutes before being unloaded. This allows for quicker response times if you’re making numerous requests to the LLM. If you want to immediately unload a model from memory, use the`ollama stop` command:
`keep_alive` parameter with the `/api/generate` and `/api/chat` endpoints to set the amount of time that a model stays in memory. The `keep_alive` parameter can be set to:
- a duration string (such as “10m” or “24h”)
- a number in seconds (such as 3600)
- any negative number which will keep the model loaded in memory (e.g. -1 or “-1m”)
- ‘0’ which will unload the model immediately after generating a response

`OLLAMA_KEEP_ALIVE` environment variable when starting the Ollama server. The `OLLAMA_KEEP_ALIVE` variable uses the same parameter types as the `keep_alive` parameter types mentioned above. Refer to the section explaining how to configure the Ollama server to correctly set the environment variable.
The `keep_alive` API parameter with the `/api/generate` and `/api/chat` API endpoints will override the `OLLAMA_KEEP_ALIVE` setting.
## How do I manage the maximum number of requests the Ollama server can queue?

If too many requests are sent to the server, it will respond with a 503 error indicating the server is overloaded. You can adjust how many requests may be queued by setting`OLLAMA_MAX_QUEUE`.
## How does Ollama handle concurrent requests?

Ollama supports two levels of concurrent processing. If your system has sufficient available memory (system memory when using CPU inference, or VRAM for GPU inference) then multiple models can be loaded at the same time. For a given model, if there is sufficient available memory when the model is loaded, it is configured to allow parallel request processing. If there is insufficient available memory to load a new model request while one or more models are already loaded, all new requests will be queued until the new model can be loaded. As prior models become idle, one or more will be unloaded to make room for the new model. Queued requests will be processed in order. When using GPU inference new models must be able to completely fit in VRAM to allow concurrent model loads. Parallel request processing for a given model results in increasing the context size by the number of parallel requests. For example, a 2K context with 4 parallel requests will result in an 8K context and additional memory allocation. The following server settings may be used to adjust how Ollama handles concurrent requests on most platforms:
- `OLLAMA_MAX_LOADED_MODELS` - The maximum number of models that can be loaded concurrently provided they fit in available memory. The default is 3 * the number of GPUs or 3 for CPU inference.
- `OLLAMA_NUM_PARALLEL` - The maximum number of parallel requests each model will process at the same time, default 1.  Required RAM will scale by`OLLAMA_NUM_PARALLEL` *`OLLAMA_CONTEXT_LENGTH` .
- `OLLAMA_MAX_QUEUE` - The maximum number of requests Ollama will queue when busy before rejecting additional requests. The default is 512

## How does Ollama load models on multiple GPUs?

When loading a new model, Ollama evaluates the required VRAM for the model against what is currently available. If the model will entirely fit on any single GPU, Ollama will load the model on that GPU. This typically provides the best performance as it reduces the amount of data transferring across the PCI bus during inference. If the model does not fit entirely on one GPU, then it will be spread across all the available GPUs.
## How can I enable Flash Attention?

Flash Attention is a feature of most modern models that can significantly reduce memory usage as the context size grows. Ollama uses Flash Attention automatically when the selected backend and devices support it. To force Flash Attention on, set`OLLAMA_FLASH_ATTENTION=1` when starting the Ollama server. To disable it, set `OLLAMA_FLASH_ATTENTION=0`.
## How can I set the quantization type for the K/V cache?

The K/V context cache can be quantized to significantly reduce memory usage when Flash Attention is enabled. To use quantized K/V cache with Ollama you can set the following environment variable:
- `OLLAMA_KV_CACHE_TYPE` - The quantization type for the K/V cache. Default is`f16` .

Currently this is a global option - meaning all models will run with the
specified quantization type.

- `f16` - high precision and memory usage (default).
- `q8_0` - 8-bit quantization, uses approximately 1/2 the memory of`f16` with a very small loss in precision, this usually has no noticeable impact on the model’s quality (recommended if not using f16).
- `q4_0` - 4-bit quantization, uses approximately 1/4 the memory of`f16` with a small-medium loss in precision that may be more noticeable at higher context sizes.

## Where can I find my Ollama Public Key?

Your
**Ollama Public Key**is the public part of the key pair that lets your local Ollama instance talk to ollama.com. You’ll need it to:

- Push models to Ollama
- Pull private models from Ollama to your machine
- Run models hosted in Ollama Cloud

### How to Add the Key

- 
**Sign-in via the Settings page** in the**Mac** and**Windows App**
- 
**Sign‑in via CLI**

- **Manually copy & paste** the key on the**Ollama Keys** page:
https://ollama.com/settings/keys

### Where the Ollama Public Key lives

Replace <username> with your actual Windows user name.

## How can I stop Ollama from starting when I login to my computer?

Ollama for Windows and macOS register as a login item during installation. You can disable this if you prefer not to have Ollama automatically start. Ollama will respect this setting across upgrades, unless you uninstall the application.
**Windows**

- In `Task Manager` go to the`Startup apps` tab, search for`ollama` then click`Disable`

**MacOS**

- Open `Settings` and search for “Login Items”, find the`Ollama` entry under`Allow in the Background` , then click the slider to disable.
