---
id: collect-260926-rattrapage/rattrapage/unsloth-on-linux
title: "Install Unsloth on Linux"
domain: rattrapage
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "Unsloth", "vLLM"]
dates: []
keywords: ["agent", "agents", "claude", "diffusion", "embedding", "mcp", "quantization", "vllm"]
source: docs/RAG/lot-rattrapage/fine-tuning/Unsloth on Linux.md
source_anchor: ""
source_lines: [1, 109]
sha256: d0057a44593c9ea3b3b18b89065c72c0f55c127cdc8d741b907ead1d271057fd
---

# Install Unsloth on Linux

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/install/linux.md).
# Install Unsloth on Linux
Install Unsloth Desktop, Unsloth Studio, or Unsloth Core on Linux.
To install Unsloth Desktop on Linux, follow the steps below.
{% stepper %}
{% step %}
### Download
Download Unsloth for Linux
{% endstep %}
{% step %}
### Install
1. Open the Unsloth installer (`.deb` file)
2. Select Install
3. Launch the app and wait for installation to complete
{% endstep %}
{% step %}
### Choose a model
Open Select model or Model hub, choose a model and quantization, then download it. Once it finishes, start chatting - no setup required.
{% endstep %}
{% step %}
### Unsloth is now ready
To start chatting, type a message and press Enter.
* **Connect tools:** [Claude Code](/docs/basics/claude-code.md), [Codex](/docs/basics/codex.md), web search, [MCP](/docs/basics/mcp.md) and more
* **Train models:** Fine-tune text, diffusion, embedding, and more
* **Generate media:** Create and train images, video, TTS locally
{% endstep %}
{% endstepper %}
### Advanced installation
Use these options for developer builds, nightly Unsloth Studio, or code-based Unsloth Core.
#### Install Unsloth Studio from the main repo
```bash
git clone https://github.com/unslothai/unsloth
cd unsloth
./install.sh --local
unsloth studio -H 0.0.0.0 -p 8888
```
#### Install nightly Unsloth Studio
```bash
git clone https://github.com/unslothai/unsloth
cd unsloth
git checkout nightly
./install.sh --local
unsloth studio -H 0.0.0.0 -p 8888
```
#### Install Unsloth Core with uv
```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
uv venv unsloth_env --python 3.13
source unsloth_env/bin/activate
uv pip install unsloth --torch-backend=auto
```
Or install the latest pip release directly:
```bash
pip install unsloth
```
To install vLLM with Unsloth:
```bash
uv pip install unsloth vllm --torch-backend=auto
```
For virtual environments, CUDA and Torch-specific commands, and advanced pip options:
[**View the advanced pip installation guide**](https://unsloth.ai/docs/get-started/install/pip-install)
### Install Unsloth Studio manually
The command below installs **Unsloth Studio**, the browser-based web UI. It does not install the native Desktop app.
```bash
curl -fsSL https://unsloth.ai/install.sh | sh
```
Use the same command to update Unsloth Studio.
#### Launch Unsloth Studio
```bash
unsloth studio -H 0.0.0.0 -p 8888
```
For detailed Unsloth Studio install instructions and requirements, [view the Unsloth Studio guide](https://unsloth.ai/docs/new/studio/install).
### Uninstall
#### Unsloth Desktop
**Ubuntu or Debian (`.deb`)**
1. Close Unsloth if it is running.
2. Remove the package:
```bash
sudo apt remove unsloth
```
**AppImage**
1. Close Unsloth if it is running.
2. Delete the downloaded `.AppImage` file.
Uninstalling Unsloth Desktop does not remove downloaded models or your files.
#### Unsloth Studio (manual installation)
Run the Linux uninstall script:
```bash
curl -fsSL https://raw.githubusercontent.com/unslothai/unsloth/main/scripts/uninstall.sh | sh
```
This removes Unsloth Studio, its CLI command, launcher data, and Linux desktop shortcuts. It does not remove downloaded Hugging Face models.
#### Delete downloaded model files
The default Hugging Face cache is:
```bash
~/.cache/huggingface/hub/
```
If `HF_HUB_CACHE`, `HF_HOME`, or `XDG_CACHE_HOME` is set, use that location instead.
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/install/linux.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
