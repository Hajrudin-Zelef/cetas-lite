---
id: collect-250926-servers-hardware/servers-hardware/fine-tuning-llms-on-amd-gpus-with-unsloth
title: "Fine-tuning LLMs on AMD GPUs with Unsloth Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Alibaba", "Mistral", "Nvidia", "Unsloth"]
dates: []
keywords: ["amd", "fine-tuning", "gpus", "agent", "agents", "inference", "llama", "memory", "mistral", "mxfp4", "nvidia", "reasoning"]
source: docs/RAG/clean4/Fine-tuning LLMs on AMD GPUs with Unsloth.md
source_anchor: ""
source_lines: [1, 118]
sha256: fe18891d34214fffffcd6a008da5ae4c50190beed2a2a7c0627e0cf0700a6d3b
---

# Fine-tuning LLMs on AMD GPUs with Unsloth Guide

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/install/amd.md).
# Fine-tuning LLMs on AMD GPUs with Unsloth Guide
Learn how to fine-tune large language models (LLMs) on AMD GPUs with Unsloth.
Fine-tune LLMs up to 2x faster with \~70% less memory on AMD hardware, no NVIDIA required. Unsloth supports AMD Radeon RDNA 3/3.5/4 (RX 6000–9000 series) on both Windows and Linux as well as data center GPUs including the MI300X (192GB).
{% stepper %}
{% step %}
#### **One-line installer**
To install Unsloth on AMD, the easiest way is to download our [Desktop app](/docs/desktop.md).
Download for WindowsDownload for Linux\
\
For manual installation:
**Linux :**
```bash
curl -fsSL https://unsloth.ai/install.sh | sh
```
**Windows (PowerShell):**
```powershell
irm https://unsloth.ai/install.ps1 | iex
```
The manual steps below are for users who want to install the Unsloth Python library on AMD with the required dependencies.
{% endstep %}
{% step %}
#### **Make a new isolated environment (Optional)**
To not break any system packages, you can make an isolated pip environment. Reminder to check what Python version you have! It might be `pip3`, `pip3.13`, `python3`, `python.3.13` etc.
**Linux:** 3.13 shown; any 3.11-3.13 works everywhere (3.10 works for manual installs
{% code overflow="wrap" %}
```bash
# Linux — swap 3.13 for whichever 3.10-3.13 you have
apt update && apt install python3.13-venv -y
python3.13 -m venv unsloth_env
source unsloth_env/bin/activate
pip install uv
```
{% endcode %}
**Windows (PowerShell):**
3.13 shown; use 3.12 if installing the unsloth\[rocm72-torch291] extra
```shellscript
py -3.13 -m venv unsloth_env
unsloth_env\Scripts\Activate.ps1
pip install uv
```
{% endstep %}
{% step %}
#### **Install PyTorch**
Planning to install Unsloth with an AMD extra (`unsloth[rocm72-torch291]` etc., see the next section)? Those bundle a matching PyTorch, so you can **skip this section**. Install PyTorch here only if you want to manage it yourself or your ROCm version isn't covered by an extra.
**Linux:** \
Install PyTorch with ROCm support from the PyTorch index. Check your ROCm version via `amd-smi version` (look for the `ROCm version:` line), then change `https://download.pytorch.org/whl/rocm7.1` to match it. ROCm 6.0 or newer is required.
{% code overflow="wrap" %}
```bash
uv pip install "torch>=2.4,<2.11.0" "torchvision<0.26.0" "torchaudio<2.11.0" \
--index-url https://download.pytorch.org/whl/rocm7.1 --upgrade --force-reinstall
```
{% endcode %}
ROCm 7.2 ships newer wheels (torch 2.11), so on ROCm 7.2 use this instead:
```bash
uv pip install "torch>=2.11.0,<2.12.0" torchvision torchaudio \
--index-url https://download.pytorch.org/whl/rocm7.2 --upgrade --force-reinstall
```
Available index tags are `rocm6.0`, `rocm6.1`, `rocm6.2`, `rocm6.3`, `rocm6.4`, `rocm7.0`, `rocm7.1`, and `rocm7.2`. ROCm 6.5-6.9 has no dedicated wheels (use `rocm6.4`), and ROCm 7.3+ uses `rocm7.2`.
*The version caps prevent accidentally pulling torch 2.11+ which only has ROCm 7.2 wheels and will break things. Update `rocm7.1` to match your detected version as before.*
We also wrote a single terminal command to extract the correct ROCM version if it helps.
```bash
ROCM_TAG="$({ command -v amd-smi >/dev/null 2>&1 && amd-smi version 2>/dev/null | awk -F'ROCm version: ' 'NF>1{split($2,a,"."); print "rocm"a[1]"."a[2]; ok=1; exit} END{exit !ok}'; } || { [ -r /opt/rocm/.info/version ] && awk -F. '{print "rocm"$1"."$2; exit}' /opt/rocm/.info/version; } || { command -v hipconfig >/dev/null 2>&1 && hipconfig --version 2>/dev/null | awk -F': *' '/HIP version/{split($2,a,"."); print "rocm"a[1]"."a[2]; ok=1; exit} END{exit !ok}'; } || { command -v dpkg-query >/dev/null 2>&1 && ver="$(dpkg-query -W -f="${Version}\n" rocm-core 2>/dev/null)" && [ -n "$ver" ] && awk -F'[.-]' '{print "rocm"$1"."$2; exit}' <<<"$ver"; } || { command -v rpm >/dev/null 2>&1 && ver="$(rpm -q --qf '%{VERSION}\n' rocm-core 2>/dev/null)" && [ -n "$ver" ] && awk -F'[.-]' '{print "rocm"$1"."$2; exit}' <<<"$ver"; })"; [ -n "$ROCM_TAG" ] && uv pip install "torch>=2.4,<2.11.0" "torchvision<0.26.0" "torchaudio<2.11.0" --index-url "https://download.pytorch.org/whl/$ROCM_TAG" --upgrade --force-reinstall
```
*Note: If your ROCm version is 7.2 or higher, replace `$ROCM_TAG` in the command above with `rocm7.1,` no PyTorch wheels exist yet for 7.2+.*

You can also use our :ledger:[automatic kernel gen RL notebook](https://github.com/unslothai/notebooks/blob/main/nb/AMD-gpt_oss_$20B$_GRPO_BF16.ipynb) also with gpt-oss to auto create matrix multiplication kernels in Python. The notebook also devices multiple methods to counteract reward hacking.
{% columns %}
{% column width="50%" %}
The prompt we used to auto create these kernels was:
{% code overflow="wrap" %}
````
Create a new fast matrix multiplication function using only native Python code.
You are given a list of list of numbers.
Output your new function in backticks using the format below:
```
python
def matmul(A, B):
return ...
```
````
{% endcode %}
{% endcolumn %}
{% column width="50%" %}
The RL process learns for example how to apply the Strassen algorithm for faster matrix multiplication inside of Python.
{% endcolumn %}
{% endcolumns %}
### :books:AMD Free One-click notebooks
AMD provides one-click notebooks equipped with **free 192GB VRAM MI300X GPUs** through their Dev Cloud. Train large models completely for free (no signup or credit card required):
* [Qwen3 (32B)](https://amd-ai-academy.com/github/unslothai/notebooks/blob/main/nb/Qwen3_$32B$_A100-Reasoning-Conversational.ipynb)
* [Llama 3.3 (70B)](https://amd-ai-academy.com/github/unslothai/notebooks/blob/main/nb/AMD-Llama3.3_$70B$_A100-Conversational.ipynb)
* [Qwen3 (14B)](https://amd-ai-academy.com/github/unslothai/notebooks/blob/main/nb/AMD-Qwen3_$14B$-Reasoning-Conversational.ipynb)
* [Mistral v0.3 (7B)](https://amd-ai-academy.com/github/unslothai/notebooks/blob/main/nb/AMD-Mistral_v0.3_$7B$-Alpaca.ipynb)
* [GPT OSS MXFP4 (20B)](https://amd-ai-academy.com/github/unslothai/notebooks/blob/main/nb/AMD-GPT_OSS_MXFP4_$20B$-Inference.ipynb) - Inference
* [Gemma4 (E2B)](https://amd-ai-academy.com/github/unslothai/notebooks/blob/main/nb/Gemma4_$E2B$_Reinforcement_Learning_Sudoku_Game.ipynb) - RL Sudoku
* Unsloth Studio
{% embed url="" %}
You can use any Unsloth notebook by prepending in [Unsloth Notebooks](/docs/get-started/unsloth-notebooks.md) by changing the link from \
to
{% columns %}
{% column width="33.33333333333333%" %}
{% endcolumn %}
{% column width="66.66666666666667%" %}
{% endcolumn %}
{% endcolumns %}
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/install/amd.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
