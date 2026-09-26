---
id: collect-250926-servers-hardware/servers-hardware/unsloth-via-docker-1
title: "Install Unsloth via Docker"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia", "Unsloth"]
dates: []
keywords: ["amd", "datacenter", "gpu", "gpus", "nvidia", "training"]
source: docs/RAG/clean4/Unsloth via Docker.md
source_anchor: ""
source_lines: [1, 155]
sha256: 73a253f641fbd124dca108958f56d4ff893d5f0468a8302071196e27358fbb8e
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
