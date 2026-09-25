---
id: collect-250926-servers-hardware/servers-hardware/unsloth-via-docker
title: "Install Unsloth via Docker"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Hugging Face", "Nvidia", "Unsloth"]
dates: []
keywords: ["agent", "agents", "amd", "datacenter", "diffusion", "embedding", "embeddings", "fine-tuning", "gguf", "gpu", "gpus", "memory"]
source: docs/RAG/clean4/Unsloth via Docker.md
source_anchor: ""
source_lines: [1, 306]
sha256: 913f14bc8c7f70c1cf2e64c30c8719d4b1bb319c58a3d77eca21fcc178f99c1f
---

# Install Unsloth via Docker

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/install/docker.md).
# Install Unsloth via Docker
Install Unsloth using our official Docker container
Learn how to run and train models with the Unsloth Docker container. No setup required - all dependencies are pre-installed. Just pull the image and start running and training models on your local NVIDIA or AMD GPUs.
| **NVIDIA** Docker: [**`unsloth/unsloth`**](https://hub.docker.com/r/unsloth/unsloth) | **AMD** Docker: [**`unsloth/unsloth-rocm`**](https://hub.docker.com/r/unsloth/unsloth-rocm) |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------- |
{% hint style="success" %}
**NEW Sep 2026: Unsloth Docker container is now updated with Unsloth Studio and AMD.**
Unsloth shares the same cache as notebooks and scripts to avoid unnecessary re-downloads.
{% endhint %}
### Quickstart
This guide section applies to systems with NVIDIA GPUs. If you have [AMD see here](#amd-quickstart).
{% tabs %}
{% tab title="Linux / WSL" %}
To get Unsloth's Docker on NVIDIA GPUs, run the following in terminal:
```bash
docker run -d --name unsloth --gpus all --ipc=host \
--ulimit memlock=-1 --ulimit stack=67108864 \
-p 8000:8000 -p 8888:8888 \
-e JUPYTER_PASSWORD="mypassword" \
-v "$PWD":/workspace/host \
-v "$HOME/.cache/huggingface":/workspace/.cache/huggingface \
-v unsloth-studio:/opt/unsloth-studio \
unsloth/unsloth
```
* Run it from your project folder, because that folder (`$PWD`) becomes `/workspace/host`.
* What the flags do: exposing ports 8000 (Unsloth Studio) and 8888 (JupyterLab); the three `-v` flags keep your files, models and Unsloth Studio data.
* Requirement: NVIDIA driver 570.26 or newer.
{% endtab %}
{% tab title="Windows (PowerShell)" %}
Run the following in Powershell:
```powershell
docker run -d --name unsloth --gpus all --ipc=host `
--ulimit memlock=-1 --ulimit stack=67108864 `
-p 8000:8000 -p 8888:8888 `
-e JUPYTER_PASSWORD="mypassword" `
-v "${PWD}:/workspace/host" `
-v "${HOME}/.cache/huggingface:/workspace/.cache/huggingface" `
-v "unsloth-studio:/opt/unsloth-studio" `
unsloth/unsloth
```
* Run it from your project folder, because that folder (`${PWD}`) becomes `/workspace/host`.
* What the flags do: exposing ports 8000 (Unsloth Studio) and 8888 (JupyterLab); the three `-v` flags keep your files, models and Unsloth Studio data.
* Requirement: NVIDIA driver 570.26 or newer.
{% endtab %}
{% endtabs %}
### Prerequisites
If you do not have Docker installed for NVIDIA Container Toolkit, follow the below:
{% tabs %}
{% tab title="Linux / WSL" %}
Install [Docker Engine](https://docs.docker.com/engine/install/) or [Docker Desktop](https://docs.docker.com/desktop/), if you haven't already:
```bash
curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh
```
If you have an NVIDIA GPU, install [NVIDIA Container Toolkit](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html#installation) if you haven't already:
```bash
curl -fsSL https://raw.githubusercontent.com/unslothai/unsloth/main/docker/install_nvidia_toolkit.sh -o install_nvidia_toolkit.sh && sudo -E bash install_nvidia_toolkit.sh
```
{% endtab %}
{% tab title="Windows" %}
**Install Docker Desktop and the NVIDIA driver.**
Install [Docker Desktop for Windows](https://docs.docker.com/desktop/setup/install/windows-install/) and the latest [NVIDIA Windows driver](https://www.nvidia.com/Download/index.aspx). \
Then update WSL in PowerShell:
```powershell
wsl --update
```
In Docker Desktop settings, select **General** and check that **Use the WSL 2 based engine** is enabled. Wait for the engine to start before continuing.
No separate NVIDIA Container Toolkit installation is needed. [Windows GPU troubleshooting](https://docs.docker.com/desktop/features/gpu/).
{% endtab %}
{% endtabs %}
### AMD Quickstart
This guide section applies to systems with AMD GPUs. If you have [NVIDIA see here](#quickstart).
{% tabs %}
{% tab title="Linux" %}
To get Unsloth's Docker on AMD GPUs, run the following in terminal:
```bash
GPU_FLAGS="--device /dev/kfd"
[ -e /dev/dri ] && GPU_FLAGS="$GPU_FLAGS --device /dev/dri"
for g in video render; do
gid=$(getent group "$g" | cut -d: -f3)
[ -n "$gid" ] && GPU_FLAGS="$GPU_FLAGS --group-add $gid"
done
docker run -d --name unsloth \
$GPU_FLAGS \
--ipc=host \
--ulimit memlock=-1 --ulimit stack=67108864 \
-p 8000:8000 -p 8888:8888 \
-e JUPYTER_PASSWORD="mypassword" \
-v "$PWD":/workspace/host \
-v "$HOME/.cache/huggingface":/workspace/.cache/huggingface \
-v unsloth-studio:/opt/unsloth-studio \
unsloth/unsloth-rocm
```
* Run it from your project folder, because that folder (`$PWD`) becomes `/workspace/host`.
* What the flags do: `$GPU_FLAGS` passes the GPU device nodes and group ids; ports 8000 (Unsloth Studio) and 8888 (JupyterLab); the three `-v` flags keep your files, models and Unsloth Studio data.
* Requirement: a working `amdgpu` driver. Built against ROCm 7.2, covers RDNA1 and newer, plus CDNA.
{% endtab %}
{% tab title="WSL" %}
In WSL, AMD GPUs are reached through WSL2's DXG bridge. Run this inside your WSL2 distro, not PowerShell, WSL2 has no `/dev/kfd` so the GPU is passed as `/dev/dxg` instead:
```bash
docker run -d --name unsloth --device /dev/dxg --ipc=host \
--ulimit memlock=-1 --ulimit stack=67108864 \
-e HSA_ENABLE_DXG_DETECTION=1 \
-e LD_LIBRARY_PATH=/usr/lib/wsl/lib:/opt/rocm/lib \
-p 8000:8000 -p 8888:8888 \
-e JUPYTER_PASSWORD="mypassword" \
-v /usr/lib/wsl/lib:/usr/lib/wsl/lib:ro \
-v "$PWD":/workspace/host \
-v "$HOME/.cache/huggingface":/workspace/.cache/huggingface \
-v unsloth-studio:/opt/unsloth-studio \
unsloth/unsloth-rocm
```
* Run it from inside WSL2 (`wsl -d Ubuntu`), from your project folder, because that folder (`$PWD`) becomes `/workspace/host`.
* What the flags do: `--device /dev/dxg` and the `/usr/lib/wsl/lib` mount pass the GPU through WSL2; ports 8000 (Unsloth Studio) and 8888 (JupyterLab); the three `-v` flags keep your files, models and Unsloth Studio data.
* Requirement: the AMD Windows driver with WSL2 support on the host
{% endtab %}
{% tab title="Windows (PowerShell)" %}
Run the following in PowerShell:
```powershell
docker run -d --name unsloth --device /dev/dxg --ipc=host `
--ulimit memlock=-1 --ulimit stack=67108864 `
-e HSA_ENABLE_DXG_DETECTION=1 `
-e LD_LIBRARY_PATH=/usr/lib/wsl/lib:/opt/rocm/lib `
-p 8000:8000 -p 8888:8888 `
-e JUPYTER_PASSWORD="mypassword" `
-v "/usr/lib/wsl/lib:/usr/lib/wsl/lib:ro" `
-v "${PWD}:/workspace/host" `
-v "${HOME}/.cache/huggingface:/workspace/.cache/huggingface" `
-v "unsloth-studio:/opt/unsloth-studio" `
unsloth/unsloth-rocm
```
* Run it from your project folder, because that folder (`${PWD}`) becomes `/workspace/host`.
* Requirement: an AMD Windows driver with WSL2 support, and Docker Desktop using the WSL 2 engine.
{% endtab %}
{% endtabs %}
If you do not have Docker installed, run the following command:
```bash
curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh
```
AMD needs no container toolkit, only the `amdgpu` driver on the host. On Windows, install Docker Desktop and the AMD driver with WSL2 support.
### 📖 Usage Guides
#### Unsloth Studio Setup
Unsloth Studio comes pre-installed, so you can chat with models, fine-tune them and generate images from the same Docker container.
{% stepper %}
{% step %}
**Start the Unsloth container.** \
Run the full `docker run` command for your OS in **Quickstart** and allow about a minute for startup. Skip this if your container is already running.
```bash
docker start unsloth
```
{% endstep %}
{% step %}
**Find your passwords.** \
If you don’t already know your password, check the container logs:
```bash
docker logs -f unsloth
```
Look for **Unsloth container ready** for links and password details. Press **Ctrl+C** to stop following the logs; the container keeps running.
{% endstep %}
{% step %}
**Sign in**
Open [`http://localhost:8000`](http://localhost:8000) and sign in as `unsloth`. On first setup, use the generated password and choose a new one when prompted. Otherwise, use your existing password.
{% endstep %}
{% step %}
**Forgot your password?**
```bash
docker exec unsloth unsloth studio reset-password --username unsloth
```
This generates a new password, signs out existing sessions and revokes API keys. No restart is needed. These commands assume your container is named `unsloth`. If needed, run `docker ps` to find its name or ID.
{% hint style="info" %}
If Docker runs on another machine (like a cloud GPU or a server), `localhost` means your own computer. You must forward the ports before you open the links. Run `ssh -L 8000:localhost:8000 -L 8888:localhost:8888 user@your-server`. See the [Security Notes](#security-notes) section for details.
{% endhint %}
{% endstep %}
{% endstepper %}
#### **Chat with a model**
{% hint style="info" %}
Unsloth Studio and JupyterLab share the same GPU. A model loaded in Unsloth keeps its GPU memory until you unload it. If you run out of memory, unload the model in Unsloth or stop the notebook kernel.
{% endhint %}
{% stepper %}
{% step %}
**Choose and download a model**
Open **Select model** at the top of the page. Browse **Recommended**, or enter a model name and click **Search Hub** to search Hugging Face. You can also browse models in **Model hub**.
For GGUF models, choose a quantization that fits your available RAM and VRAM. Avoid variants marked **OOM**.
Wait for the model to finish downloading and loading before sending your first message. You can select downloaded models again from **On Device**.
{% hint style="info" %}
Downloaded models are stored in `/workspace/.cache/huggingface`. With the Quickstart volume mount, this cache is shared with notebooks and stays on your computer even if you remove the container.
{% endhint %}
{% endstep %}
{% step %}
**Send a message**
Once the model is ready, type a message and press **Enter**. Try:
> Explain how Docker containers differ from virtual machines.
See the [Unsloth Studio Chat](/docs/new/studio/chat.md) guide for web search and more chat settings.
{% endstep %}
{% endstepper %}
#### **Train a model**
You can train models across text, [vision](https://unsloth.ai/docs/basics/vision-fine-tuning) and [audio](https://unsloth.ai/docs/basics/text-to-speech-tts-fine-tuning), with support for [embeddings](https://unsloth.ai/docs/basics/embedding-finetuning) and image diffusion too. All from the same Docker image.
**Choose a model and dataset**
Select **Train** and open **Configure**. Choose a model and training method, then select a Hugging Face dataset or upload your own.
**Start training**
Review **Parameters** in **Simple** or **Advanced** mode, then click **Start Training**. Track progress, loss and GPU usage in **Current Run**.
**Try your trained model**
When training finishes, click **Compare in Chat**. Test the original and fine-tuned models with the same prompts to see how their responses differ.
See the [Studio training guide](https://unsloth.ai/docs/new/studio/start) for more detail, or use [Data Recipes](https://unsloth.ai/docs/new/studio/data-recipe) to prepare your dataset.
#### **Generate images**
{% stepper %}
{% step %}
**Open Images and choose a model**
Select **Images** in the sidebar. The **Create** workflow opens by default.
Open **Select image model** and choose from **Recommended**, such as Z-Image-Turbo. For GGUF models, pick a quantization that fits your device. **TIGHT** may run slowly; **OOM** is unlikely to fit.
{% endstep %}
{% step %}
**Enter a prompt and generate**
Describe the image you want to create. For example:
> Top-down shot of a koi pond in a Japanese garden, dozens of orange, white and black koi swimming in tight formation, water clear enough to see the stone bottom, red maple leaves floating on the surface. Bright midday light, saturated colour, crisp reflections. Realistic photo, 50mm.
Leave the image settings at their defaults and click **Generate**.
Open your image in the gallery and click **Download** to save it, or **Recipe** to reuse its prompt and settings.
**Transform**, **Inpaint** and **Edit** depend on the loaded model. See the [image diffusion guide](https://unsloth.ai/docs/basics/diffusion-image) for supported workflows and settings.
{% endstep %}
{% endstepper %}
#### **Access JupyterLab**
{% columns %}
{% column %}
The Unsloth Docker image includes JupyterLab and ready-to-use notebooks.
Open [`http://localhost:8888`](http://localhost:8888) and sign in. Use the `JUPYTER_PASSWORD` you chose when starting the container. If you did not set one, use the generated password printed in the container log.
{% hint style="info" %}
JupyterLab passwords are set when the container is created. To change it, remove and recreate the container with a new password.
{% endhint %}
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
After signing in, you’ll see **Unsloth Notebooks**, with folders grouped by model and task.
These folders contain shortcuts to the notebook files in `/workspace/unsloth-notebooks`. Each time the container starts, the notebooks are refreshed from GitHub without overwriting
\
your edits.
Double-click **01 Main Notebooks** to browse examples, or choose a category such as vision, speech or reinforcement learning.

:`.
| Container path | Description | Suggested mount |
| ------------------------------- | ---------------------------------------------------------------- | --------------------------------------------------------------- |
| `/workspace/host` | Your project files | Your project folder, e.g. `"$PWD"` |
| `/workspace/.cache/huggingface` | Downloaded models and datasets | Your host Hugging Face cache, e.g. `"$HOME/.cache/huggingface"` |
| `/opt/unsloth-studio` | Unsloth Studio accounts, chats, trained models, exports and runs | Named volume `unsloth-studio` |
| `/workspace/unsloth-notebooks` | Example notebooks, including your edits | Optional |
| `/workspace/.cache/triton` | Compiled GPU kernels | Optional, speeds up restarts |
* `/workspace` is the working directory and the folder JupyterLab opens in.
* Use a named volume for `/opt/unsloth-studio`, not a host folder. Unsloth Studio needs symlinks there, and a Windows or macOS host folder may not allow them, which stops the container at start.
* To give Unsloth Studio or JupyterLab more of your files, mount more folders under `/workspace`, e.g. `-v /data/datasets:/workspace/datasets`, and refer to them by that container path.
#### Ports
| Port | Service |
| ------ | -------------------------------- |
| `8000` | Unsloth Studio |
| `8888` | JupyterLab |
| `22` | SSH (only when `SSH_KEY` is set) |
Map a container port to any free host port with `-p :`. For example, to enable SSH on host port `2222`:
{% tabs %}
{% tab title="NVIDIA" %}
```bash
docker run -d --name unsloth --gpus all \
-p 8000:8000 -p 8888:8888 -p 2222:22 \
-e SSH_KEY="$(cat ~/.ssh/id_ed25519.pub)" \
-v "$PWD":/workspace/host \
-v "$HOME/.cache/huggingface":/workspace/.cache/huggingface \
-v unsloth-studio:/opt/unsloth-studio \
unsloth/unsloth
```
{% endtab %}
{% tab title="AMD" %}
```bash
GPU_FLAGS="--device /dev/kfd"
[ -e /dev/dri ] && GPU_FLAGS="$GPU_FLAGS --device /dev/dri"
for g in video render; do
gid=$(getent group "$g" | cut -d: -f3)
[ -n "$gid" ] && GPU_FLAGS="$GPU_FLAGS --group-add $gid"
done
docker run -d --name unsloth $GPU_FLAGS \
-p 8000:8000 -p 8888:8888 -p 2222:22 \
-e SSH_KEY="$(cat ~/.ssh/id_ed25519.pub)" \
-v "$PWD":/workspace/host \
-v "$HOME/.cache/huggingface":/workspace/.cache/huggingface \
-v unsloth-studio:/opt/unsloth-studio \
unsloth/unsloth-rocm
```
{% endtab %}
{% endtabs %}
### **🔒 Security Notes**
* The container runs as root. Use `unsloth/unsloth:core` with `--user :` if you need mounted files owned by your host user.
* SSH is off unless `SSH_KEY` or `PUBLIC_KEY` is set. It is key-only, as root, on port 22.
* Published ports listen on every interface. On a cloud host, bind to `127.0.0.1` or use `-e UNSLOTH_STUDIO_SECURE=1`.
* JupyterLab is a full shell. Anyone who can log in to JupyterLab can run any command in the container. Unsloth Studio and JupyterLab serve plain HTTP, so do not expose them to the internet directly.
* Unsloth Studio tools can run code. Studio's server-side tools are on by default and can run commands inside the container. Only mount host folders you are comfortable giving the container access to.
* Environment variables are visible to Docker users. Values passed with `-e` or `--env-file`, including tokens and passwords, show up in `docker inspect`. Anyone with access to the Docker daemon can read them.
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/install/docker.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
