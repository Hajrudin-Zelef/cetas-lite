---
id: collect-250926-servers-hardware/servers-hardware/how-to-connect-ollama-to-unsloth-unsloth-documentation
title: "how-to-connect-ollama-to-unsloth-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "OpenAI", "Unsloth"]
dates: []
keywords: ["gguf", "quantization", "qwen"]
source: docs/RAG/clean4/how-to-connect-ollama-to-unsloth-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 80]
sha256: 604356bf5058110b1ddb5b5b525a1d77f57d9825db8816e7ad5e0fca4eea7745
---

# how-to-connect-ollama-to-unsloth-unsloth-documentation

Ollama lets you run local LLMs on your own hardware, and Unsloth makes it easy to connect and run those models directly into a open-source UI chat interface. In this guide, you’ll learn how to install Ollama, run native Ollama models or GGUF models from Hugging Face, connect Ollama to Unsloth, and start chatting with local AI models.

Whether you want to use models like Qwen, import a GGUF file, or expose your local Ollama server through an OpenAI-compatible endpoint, this walkthrough covers the full setup from installation to first chat.

You can choose a model in two common ways:

- Use a GGUF model from Hugging Face, then copy the Ollama command from **Use this model** .

For an Ollama model, pull and run it:

```
ollama pull qwen3.6:35b-a3b
ollama run qwen3.6:35b-a3b
```
If the Ollama app or service is not already running, start it first:

`ollama serve`
If you are using a GGUF model from Hugging Face, the easiest way to get the command is from the model page.

Open the model you want to use, click **Use this model**, then choose **Ollama** from the local apps list. Pick the quantization you want from the dropdown, then copy the generated command.

For example, with Ollama:

`ollama run hf.co/unsloth/Qwen3.6-35B-A3B-GGUF:UD-Q4_K_XL`
This helps avoid mistakes with the repo name or quantization tag.

Open **Settings → Connections**, then click **Add Connection**.

Select **Ollama**, then enter your connection details:

Use the Ollama URL shown in the Unsloth form. In most local setups, this is:

`http://localhost:11434`
If Unsloth asks for an OpenAI-compatible base URL, use:

`http://localhost:11434/v1`
Ollama normally does not need an API key. Leave the API key field empty unless you are using a proxy that requires one.

Click **Load Models** to fetch the models running in Ollama, or enter the **model ID** yourself, for example `qwen3.6`.

Use these while setting up the model you want to expose to Unsloth:

`ollama run qwen3.6:35b-a3b`

Run a model and open an interactive chat

`ollama pull qwen3.6:35b-a3b`

Download a model without starting chat

`ollama ls`

List downloaded models

`ollama ps`

List models currently running

`ollama stop qwen3.6:35b-a3b`

Stop a running model

`ollama rm qwen3.6:35b-a3b`

Remove a downloaded model

`ollama serve`

Start the Ollama server

If you are importing a local GGUF into Ollama, create a `Modelfile`, then run:

`ollama create -f Modelfile`
If Ollama is not detected, make sure the Ollama app or service is running. Then click **Load Models** again in Unsloth.

For the full command list, see the Ollama CLI reference.

Last updated

Was this helpful?
