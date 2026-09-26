---
id: collect-260926-rattrapage/rattrapage/fine-tuning-llms-on-nvidia-dgx-station-with-unsloth
title: "Fine-Tuning LLMs on NVIDIA DGX Station with Unsloth"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia", "Unsloth"]
dates: []
keywords: ["fine-tuning", "nvidia", "agent", "agents", "attention", "blackwell", "compute", "gguf", "gpu", "memory", "moe", "nvlink"]
source: docs/RAG/lot-rattrapage/fine-tuning/Fine-Tuning LLMs on NVIDIA DGX Station with Unsloth.md
source_anchor: ""
source_lines: [1, 136]
sha256: 22dd2d7d55d1c4cf773c03300e2bd7f7f1120ceb250552d27be9a09ab32e996a
---

# Fine-Tuning LLMs on NVIDIA DGX Station with Unsloth

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/blog/dgx-station.md).
# Fine-Tuning LLMs on NVIDIA DGX Station with Unsloth
NVIDIA DGX Station tutorial on how to fine-tune with notebooks from Unsloth.
You can now train LLMs locally on your NVIDIA DGX Station with [Unsloth](https://github.com/unslothai/unsloth). DGX Station has more than **\~200GB VRAM** and over **700GB of unified GPU / CPU memory** and combines a Grace CPU and a Blackwell GPU in a tightly connected system designed for large-scale AI workloads. Linked by NVLink-C2C, the CPU and GPU remain distinct but work together far more efficiently than in a traditional CPU-GPU setup.
In this guide, we’ll use Unsloth notebooks train [Qwen3.5](#qwen3.5-35b-a3b-fine-tuning) and [gpt-oss-120b](#gpt-oss-120b-fine-tuning) on DGX Station. Thank you to NVIDIA for providing some early access DGX Station hardware to test Unsloth on!
### Quickstart
You will need `python3` installed and in particular the dev headers are needed. On our system we have `python 3.12` so we will install the 3.12 dev headers.
```bash
sudo apt update
sudo apt install python3.12-dev
```
Then create a fresh virtual environment to install [Unsloth](https://github.com/unslothai/unsloth). This way we minimize dependency conflicts and preserve the state of the current working environment.
{% code overflow="wrap" %}
```bash
python3 -m venv .unsloth
source .unsloth/bin/activate
pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu130
```
{% endcode %}
{% hint style="warning" %}
First install `torch` from the `cuda 13` index otherwise we could get the CPU version or a mismatch in architecture and capabilities!
{% endhint %}

Now we can install Unsloth:
```bash
pip install unsloth
```

Now lets install `xformers` and optionally build `flash-attention` from source. Both packages take time so please be patient while they build.
{% code overflow="wrap" expandable="true" %}
```bash
pip install --no-deps --no-build-isolation xformers==0.0.33.post1
# Optionally flash-attn
# Clone and build (targets sm_100 for B300)
git clone https://github.com/Dao-AILab/flash-attention
cd flash-attention
# B300 = sm_100, set arch explicitly
TORCH_CUDA_ARCH_LIST="10.0" MAX_JOBS=8 pip install . --no-build-isolation
cd ..
```
{% endcode %}

{% columns %}
{% column %}
For Qwen 3.5 MoE we’ll want to download two kernel packages `flash-linear-attention` and `causal-conv1d` to make it fast.
{% code overflow="wrap" expandable="true" %}
```bash
pip install --no-build-isolation flash-linear-attention causal_conv1d==1.6.0
```
{% endcode %}
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
If you don’t already have a notebook client, install one. For this guide we will use Jupyter Notebook:
{% code overflow="wrap" expandable="true" %}
```bash
cd ..
pip install notebook
pip install ipywidgets
```
{% endcode %}
Finally we download the actual Unsloth notebooks to run. There are 250+ notebooks for LLM Training as well as Python scripts.
{% code overflow="wrap" expandable="true" %}
```bash
git clone https://github.com/unslothai/notebooks.git
cd notebooks
```
{% endcode %}
### Training Tutorials
{% columns %}
{% column %}
Now we can launch Jupyter Notebook and navigate to the UI on a browser.
{% code overflow="wrap" expandable="true" %}
```bash
jupyter notebook
```
{% endcode %}
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
{% columns %}
{% column %}
Copy and paste the `localhost` site with token parameter and paste into your browser. You should see something like:
The `nb` folder has all the notebooks to run.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
#### Qwen3.5-35B-A3B Training
{% columns %}
{% column %}
Open the file `nb/Qwen3_5_MoE.ipynb`. Skip past the installation section since we already installed everything we need beforehand. Navigate to the Unsloth section and start executing cells from there.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
{% columns %}
{% column %}
The notebook covers model setup, dataset preparation, and trainer configuration. Each step can take some time as we are downloading a very large model, initializing billions of weights, and further optimizing to make it run fast.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
Training is very fast with the default setting. On the DGX Station there is plenty of memory so you can play with the default training hyper parameters to really push the memory and compute. Once done training you can save the model for later, push the model to Hugging Face Hub to share with others, or export to a quantized format.
#### gpt-oss-120b Training
{% columns %}
{% column %}
Open the file `nb/gpt-oss-(120B)_A100-Fine-tuning.ipynb`. Skip past the installation section since we already installed the prerequisites and navigate to the Unsloth section. We can start running the notebook from there. The notebook will use around 72 GB of GPU memory and take about 10 minutes.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
{% columns %}
{% column %}
Each cell can take some time to run as we need to download the model, initialize the weights, and further optimize for a fast experience. The notebook goes through dataset preprocessing and trainer setup. Once we get to the `trainer.train()` cell and execute training begins.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
Now that it’s complete we can save the model for later use, push to Hugging Face Hub to share with the world, or export it to GGUF format.
Read more about NVIDIA's DGX Station at
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/blog/dgx-station.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
