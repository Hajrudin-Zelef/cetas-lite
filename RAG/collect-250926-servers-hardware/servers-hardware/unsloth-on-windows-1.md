---
id: collect-250926-servers-hardware/servers-hardware/unsloth-on-windows-1
title: "Install Unsloth on Windows"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Nvidia", "Unsloth"]
dates: []
keywords: ["claude", "diffusion", "embedding", "fine-tuning", "gpu", "gpus", "license", "lora", "mcp", "memory", "nvidia", "quantization"]
source: docs/RAG/clean4/Unsloth on Windows.md
source_anchor: ""
source_lines: [1, 182]
sha256: 3b9ff6f1b6061e0f6e1f9309c8697d236e96c606b8a7e97dfab45d4bd6135de9
---

# Install Unsloth on Windows

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/install/windows-installation.md).
# Install Unsloth on Windows
Install Unsloth Desktop, Unsloth Studio, or Unsloth Core on Windows.
To install Unsloth Desktop on Windows, follow the steps below.
{% stepper %}
{% step %}
### Download
Download Unsloth for Windows
{% endstep %}
{% step %}
### Install
1. Open the Unsloth installer (`.exe` file)
2. Follow the setup prompts
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
### Install Unsloth manually
The sections below cover the web-based **Unsloth Studio** and the original code-based **Unsloth Core**.
For **Unsloth Core** on Windows, choose Conda, Docker, or WSL below. If PyTorch is already installed, `pip install unsloth` may be enough.
Conda TutorialDocker TutorialWSL Tutorial
### Unsloth Studio
We have launched a new web UI called [Unsloth Studio](/docs/new/studio/install.md) that works on Windows out of the box:
```bash
irm https://unsloth.ai/install.ps1 | iex
```
Use the same command to update.
Then to launch every time:
```bash
unsloth studio -H 0.0.0.0 -p 8888
```
For detailed Unsloth Studio install instructions and requirements, [view our guide](/docs/new/studio/install.md).
Below are installation instructions for the original **Unsloth Core**:
### Method #1 - Windows via Conda:
{% stepper %}
{% step %}
**Install Miniconda (or Anaconda)**
Download Anaconda [here](https://www.anaconda.com/download). Our suggestion is to use [Miniconda](https://www.anaconda.com/docs/getting-started/miniconda/install#quickstart-install-instructions). To use it, first enter Powershell - search "Windows Powershell" in Start:
Then it'll open up Powershell:
Then copy paste the below: CTRL+C, and paste it in Powershell CTRL+V:
{% code overflow="wrap" %}
```ps
Invoke-WebRequest -Uri "https://repo.anaconda.com/miniconda/Miniconda3-latest-Windows-x86_64.exe" -OutFile ".\miniconda.exe"
Start-Process -FilePath ".\miniconda.exe" -ArgumentList "/S" -Wait
del .\miniconda.exe
```
{% endcode %}
Accept the warning and press "Paste anyway" and wait.
It's downloading the installer like below:
After installing, open up open **Anaconda Powershell Prompt** to use Miniconda via Start -> Search for it:
Then you'll see:
{% endstep %}
{% step %}
**Make conda environment**
```bash
conda create --name unsloth_env python==3.12 -y
conda activate unsloth_env
```
**You will see:**
{% endstep %}
{% step %}
**Check `nvidia-smi` to confirm you have a GPU, and look for the CUDA version**
After typing `nvidia-smi` in Powershell, you should see something like below. If you don't have `nvidia-smi` or the below fails to pop up, you need to reinstall [NVIDIA drivers](https://www.nvidia.com/en-us/drivers/).
{% endstep %}
{% step %}
**Install PyTorch**
When running `nvidia-smi` you will see at the top right corner: "CUDA Version: 13.0". Install PyTorch in PowerShell via. Change `130` to your CUDA version - ensure the [version exists](https://pytorch.org/) and matches your CUDA driver version.
{% code overflow="wrap" %}
```bash
pip3 install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu130
```
{% endcode %}
You will see:
Try running this in Python via `python` after PyTorch is installed:
{% code overflow="wrap" %}
```python
import torch
print(torch.cuda.is_available())
A = torch.ones((10, 10), device = "cuda")
B = torch.ones((10, 10), device = "cuda")
A @ B
```
{% endcode %}
You should see a matrix of 10s. Also verify True for the first.
{% endstep %}
{% step %}
**Install Unsloth (only if PyTorch works!)**
{% hint style="danger" %}
**Confirm PyTorch works fine and runs - if not PyTorch is broken and that means your Windows machine might need a re installation of CUDA drivers unfortunately.**
{% endhint %}
In Powershell (after exiting Python via `exit()` , do and wait:
```bash
pip install unsloth
```
{% endstep %}
{% step %}
**Verify Unsloth works**
Now use any script in [Unsloth Notebooks](/docs/get-started/unsloth-notebooks.md) (save to .py file), or use the below basic script:
{% code expandable="true" %}
```python
from unsloth import FastLanguageModel, FastModel
import torch
from trl import SFTTrainer, SFTConfig
from datasets import load_dataset
max_seq_length = 512
url = "https://huggingface.co/datasets/laion/OIG/resolve/main/unified_chip2.jsonl"
dataset = load_dataset("json", data_files = {"train" : url}, split = "train")
model, tokenizer = FastLanguageModel.from_pretrained(
model_name = "unsloth/gemma-3-270m-it",
max_seq_length = max_seq_length, # Choose any for long context!
load_in_4bit = True, # 4-bit quantization. False = 16-bit LoRA.
load_in_8bit = False, # 8-bit quantization
load_in_16bit = False, # 16-bit LoRA
full_finetuning = False, # Use for full fine-tuning.
trust_remote_code = False, # Enable to support new models
# token = "hf_...", # use one if using gated models
)
# Do model patching and add fast LoRA weights
model = FastLanguageModel.get_peft_model(
model,
r = 16,
target_modules = ["q_proj", "k_proj", "v_proj", "o_proj",
"gate_proj", "up_proj", "down_proj",],
lora_alpha = 16,
lora_dropout = 0, # Supports any, but = 0 is optimized
bias = "none", # Supports any, but = "none" is optimized
# [NEW] "unsloth" uses 30% less VRAM, fits 2x larger batch sizes!
use_gradient_checkpointing = "unsloth", # True or "unsloth" for very long context
random_state = 3407,
max_seq_length = max_seq_length,
use_rslora = False, # We support rank stabilized LoRA
loftq_config = None, # And LoftQ
)
trainer = SFTTrainer(
model = model,
train_dataset = dataset,
tokenizer = tokenizer,
args = SFTConfig(
max_seq_length = max_seq_length,
per_device_train_batch_size = 2,
gradient_accumulation_steps = 4,
warmup_steps = 10,
max_steps = 60,
logging_steps = 1,
output_dir = "outputs",
optim = "adamw_8bit",
seed = 3407,
dataset_num_proc = 1,
),
)
trainer.train()
```
{% endcode %}
You should see:
```bash
🦥 Unsloth: Will patch your computer to enable 2x faster free finetuning.
🦥 Unsloth Zoo will now patch everything to make training faster!
==((====))== Unsloth 2026.1.4: Fast Gemma3 patching. Transformers: 4.57.6.
\\ /| NVIDIA GeForce RTX 3060. Num GPUs = 1. Max memory: 12.0 GB. Platform: Windows.
O^O/ \_/ \ Torch: 2.10.0+cu130. CUDA: 8.6. CUDA Toolkit: 13.0. Triton: 3.6.0
\ / Bfloat16 = TRUE. FA [Xformers = 0.0.34. FA2 = False]
"-____-" Free license: http://github.com/unslothai/unsloth
Unsloth: Fast downloading is enabled - ignore downloading bars which are red colored!
Unsloth: Gemma3 does not support SDPA - switching to fast eager.
Unsloth: Making `model.base_model.model.model` require gradients
Unsloth: Tokenizing ["text"] (num_proc=1): 0%| | 0/210289 [00:00, ? examples/s]� Unsloth: Will patch your computer to enable 2x faster free finetuning.
🦥 Unsloth: Will patch your computer to enable 2x faster free finetuning.
```
And training:
{% endstep %}
{% endstepper %}
### Method #2 - Docker:
