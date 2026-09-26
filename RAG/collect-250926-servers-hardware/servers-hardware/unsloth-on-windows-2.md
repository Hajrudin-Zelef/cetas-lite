---
id: collect-250926-servers-hardware/servers-hardware/unsloth-on-windows-2
title: "Install Unsloth on Windows"
domain: servers-hardware
role: reference
task: reference
actors: ["Hugging Face", "Intel", "Microsoft", "Nvidia", "Unsloth", "vLLM"]
dates: []
keywords: ["blackwell", "datacenter", "fine-tuning", "gpu", "gpus", "intel", "nvidia", "training", "vllm"]
source: docs/RAG/clean4/Unsloth on Windows.md
source_anchor: ""
source_lines: [183, 322]
sha256: 56d9b9273e9d8d63c62c08402f5c70d36f99dab0518d345ec3463a4e329ccadd
---

# Install Unsloth on Windows

Docker might be the easiest way for Windows users to get started with Unsloth as there is no setup needed or dependency issues. [**`unsloth/unsloth`**](https://hub.docker.com/r/unsloth/unsloth) is Unsloth's only Docker image. For [Blackwell](/docs/blog/fine-tuning-llms-with-blackwell-rtx-50-series-and-unsloth.md) and 50-series GPUs, use this same image - no separate image needed.
For installation instructions, please follow our [Docker guide](/docs/blog/how-to-fine-tune-llms-with-unsloth-and-docker.md), otherwise here is a quickstart guide:
{% stepper %}
{% step %}
**Install Docker and NVIDIA Container Toolkit.**
Install Docker via [Linux](https://docs.docker.com/engine/install/) or [Desktop](https://docs.docker.com/desktop/) (other). Then install [NVIDIA Container Toolkit](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html#installation):
```bash
export NVIDIA_CONTAINER_TOOLKIT_VERSION=1.17.8-1
sudo apt-get update && sudo apt-get install -y \
nvidia-container-toolkit=${NVIDIA_CONTAINER_TOOLKIT_VERSION} \
nvidia-container-toolkit-base=${NVIDIA_CONTAINER_TOOLKIT_VERSION} \
libnvidia-container-tools=${NVIDIA_CONTAINER_TOOLKIT_VERSION} \
libnvidia-container1=${NVIDIA_CONTAINER_TOOLKIT_VERSION}
```
{% endstep %}
{% step %}
**Run the container.**
[**`unsloth/unsloth`**](https://hub.docker.com/r/unsloth/unsloth) is Unsloth's only Docker image.
```bash
docker run -d -e JUPYTER_PASSWORD="mypassword" \
-p 8888:8888 -p 2222:22 \
-v $(pwd)/work:/workspace/work \
--gpus all \
unsloth/unsloth
```
{% endstep %}
{% step %}
**Access Jupyter Lab**
Go to [http://localhost:8888](http://localhost:8888/) and open Unsloth. Access the `unsloth-notebooks` tabs to see Unsloth notebooks.
{% endstep %}
{% step %}
**Start training with Unsloth**
If you're new, follow our step-by-step [Fine-tuning Guide](/docs/get-started/fine-tuning-llms-guide.md), [RL Guide](/docs/get-started/reinforcement-learning-rl-guide.md) or just save/copy any of our premade [notebooks](/docs/get-started/unsloth-notebooks.md).
{% endstep %}
{% step %}
**Docker issues - GPU not discovered?**
Try doing WSL via [#method-2-wsl](#method-2-wsl "mention")
{% endstep %}
{% endstepper %}
### Method #3 - WSL:
{% stepper %}
{% step %}
**Install WSL**
Open up Command Prompt, the Terminal, and install Ubuntu. Set the password if asked.
```bash
wsl.exe --install Ubuntu-24.04
wsl.exe -d Ubuntu-24.04
```
{% endstep %}
{% step %} **If you did NOT do (1), so you already installed WSL****, enter WSL by typing `wsl` and ENTER in the command prompt**
```bash
wsl
```
{% endstep %}
{% step %}
**Install Python**
{% code overflow="wrap" %}
```bash
sudo apt update
sudo apt install python3 python3-full python3-pip python3-venv -y
```
{% endcode %}
{% endstep %}
{% step %}
**Install PyTorch**
{% code overflow="wrap" %}
```bash
pip install torch torchvision --force-reinstall --index-url https://download.pytorch.org/whl/cu130
```
{% endcode %}
If you encounter permission issues, use `–break-system-packages` so `pip install torch torchvision --force-reinstall --index-url https://download.pytorch.org/whl/cu130 –break-system-packages`
{% endstep %}
{% step %}
**Install Unsloth and Jupyter Notebook**
{% code overflow="wrap" %}
```bash
pip install unsloth jupyter
```
{% endcode %}
If you encounter permission issues, use `–-break-system-packages` so `pip install unsloth jupyter –-break-system-packages`
{% endstep %}
{% step %}
**Launch Unsloth via Jupyter Notebook**
{% code overflow="wrap" %}
```bash
jupyter notebook
```
{% endcode %}
Then open up our notebooks within [Unsloth Notebooks](/docs/get-started/unsloth-notebooks.md)and load them up! You can also go to Colab notebooks and download > download .ipynb and load them.
![](https://3215535692-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FxhOjnexMCB3dmuQFQ2Zq%2Fuploads%2FVbqNWsG2CCHKJJjrnU4s%2Funknown.png?alt=media\&token=854a6d0e-fc84-4e44-bf8e-4bf254801692)
{% endstep %}
{% endstepper %}
{% hint style="warning" %}
If you're using GRPO or plan to use vLLM, currently vLLM does not support Windows directly but only via WSL or Linux.
{% endhint %}
### **Troubleshooting /** Advanced
For **advanced installation instructions** or if you see weird errors during installations:
1. Install `torch` and `triton`. Go to to install it. For example `pip install torch torchvision torchaudio triton`
2. Confirm if CUDA is installated correctly. Try `nvcc`. If that fails, you need to install `cudatoolkit` or CUDA drivers.
3. If using an Intel GPU, you will need to follow our [Intel Windows guide](/docs/get-started/install/intel.md#windows-only-runtime-configurations)
4. Install `xformers` manually. You can try installing `vllm` and seeing if `vllm` succeeds. Check if `xformers` succeeded with `python -m xformers.info` Go to . Another option is to install `flash-attn` for Ampere GPUs.
5. Double check that your versions of Python, CUDA, CUDNN, `torch`, `triton`, and `xformers` are compatible with one another. The [PyTorch Compatibility Matrix](https://github.com/pytorch/pytorch/blob/main/RELEASE.md#release-compatibility-matrix) may be useful.
6. Finally, install `bitsandbytes` and check it with `python -m bitsandbytes`
7. If Unsloth is not detecting or using your GPU and if you are using our Docker container on Windows, your CUDA toolkit version `nvcc --version` should match the version of CUDA shown by nvidia-smi on the host GPU support for Docker containers on Windows is not automatic. [You need to follow Docker's guide](https://docs.docker.com/desktop/features/gpu/).
## Uninstall
#### Unsloth Desktop
1. Close Unsloth if it is running.
2. Open **Settings** in Windows.
3. Go to **Apps** → **Installed apps**.
4. Find **Unsloth**, select the three-dot menu, then choose **Uninstall**.
5. Follow the prompts to remove the app.
Uninstalling Unsloth Desktop does not remove downloaded models or your files.
#### Unsloth Studio (manual installation)
The recommended way to fully remove Unsloth Studio is the uninstall script for your OS. It stops any running servers, removes the app, CLI command, launcher data, shortcuts, and platform-specific entries (macOS `.app` bundle + Launch Services; Windows Start Menu + registry + PATH):
```ps1
irm https://raw.githubusercontent.com/unslothai/unsloth/main/scripts/uninstall.ps1 | iex
```
#### Manual uninstall
If you prefer to remove only specific parts:
**1. Remove app only** (keeps history, chats, checkpoints, and exports intact):
* `Remove-Item -Recurse -Force "$HOME\.unsloth\studio\unsloth_studio"`
**2. Remove Unsloth entirely** (keeps other Unsloth tools intact):
* `Remove-Item -Recurse -Force "$HOME\.unsloth\studio"`
**3. Remove everything Unsloth-related:**
* `Remove-Item -Recurse -Force "$HOME\.unsloth"`
{% hint style="warning" %}
Note: Step 3 deletes everything in history, chats, model checkpoints, and exports. This cannot be undone.
{% endhint %}
**4. Remove shortcuts and symlinks:**
```shellscript
Remove-Item -Force "$HOME\Desktop\Unsloth Studio.lnk"
Remove-Item -Force "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Unsloth Studio.lnk"
```
**5. Remove the CLI command:**
* **Windows (PowerShell):** The installer added the venv's `Scripts` directory to your User PATH. To remove it, open Settings → System → About → Advanced system settings → Environment Variables, find `Path` under User variables, and remove the entry pointing to `.unsloth\studio\...\Scripts`.
{% hint style="info" %}
Note: Steps 1-5 dont touch your downloaded HF model files. See Deleting cached HF model files below if you want to reclaim that space.
{% endhint %}
### **Deleting cached HF model files**
You can delete old model files either from the bin icon in model search or by removing the relevant cached model folder from the default Hugging Face cache directory. By default, Hugging Face uses `~/.cache/huggingface/hub/` on macOS/Linux/WSL and `C:\Users\\.cache\huggingface\hub\` on Windows.
