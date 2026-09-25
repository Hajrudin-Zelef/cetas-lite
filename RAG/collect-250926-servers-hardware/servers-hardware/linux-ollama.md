---
id: collect-250926-servers-hardware/servers-hardware/linux-ollama
title: "linux-ollama"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "gpu"]
source: docs/RAG/clean4/linux-ollama.md
source_anchor: ""
source_lines: [1, 51]
sha256: 6fe64205545f3969b559079a5ad0aad8a992816cdc29ea380f2e21c1fc09d75e
---

# linux-ollama

## Install

To install Ollama, run the following command:
## Manual install

If you are upgrading from a prior version, you should remove the old libraries
with 

`sudo rm -rf /usr/lib/ollama` first.
### AMD GPU install

If you have an AMD GPU, also download and extract the additional ROCm package:
### ARM64 install

Download and extract the ARM64-specific package:
### Adding Ollama as a startup service (recommended)

Create a user and group for Ollama:`/etc/systemd/system/ollama.service`:
### Install CUDA drivers (optional)

Download and install CUDA. Verify that the drivers are installed by running the following command, which should print details about your GPU:
### Install AMD ROCm drivers (optional)

Download and Install ROCm v7.
### Start Ollama

Start Ollama and verify it is running:
While AMD has contributed the 

`amdgpu` driver upstream to the official linux
kernel source, the version is older and may not support all ROCm features. We
recommend you install the latest driver from
https://www.amd.com/en/support/download/linux-drivers.html for best support of your Radeon
GPU.
## Customizing

To customize the installation of Ollama, you can edit the systemd service file or the environment variables by running:`/etc/systemd/system/ollama.service.d/override.conf`:
## Updating

Update Ollama by running the install script again:
## Installing specific versions

Use`OLLAMA_VERSION` environment variable with the install script to install a specific version of Ollama, including pre-releases. You can find the version numbers in the releases page.
For example:
## Viewing logs

To view logs of Ollama running as a startup service, run:
## Uninstall

Remove the ollama service:`/usr/local/lib`, `/usr/lib`, or `/lib`):
`/usr/local/bin`, `/usr/bin`, or `/bin`):
