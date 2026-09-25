---
id: collect-250926-servers-hardware/servers-hardware/windows-ollama
title: "windows-ollama"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "nvidia"]
source: docs/RAG/clean4/windows-ollama.md
source_anchor: ""
source_lines: [1, 66]
sha256: 7872b05b6fa702e9e3c7059416d4d0779f9ff17718f980ce86ce03dabf3dcfc5
---

# windows-ollama

`ollama` command line is available in `cmd`, `powershell` or your favorite
terminal application. As usual the Ollama API will be served on
`http://localhost:11434`.
## System Requirements

- Windows 10 22H2 or newer, Home or Pro
- NVIDIA 551.61 or newer Drivers if you have an NVIDIA card
- AMD ROCm v7 / HIP7-capable driver stack for ROCm acceleration, or a Vulkan-capable AMD Radeon driver for Vulkan acceleration

Some RDNA2 / Radeon RX 6000 systems, including RX 6800-class cards, may not
expose ROCm v7 on current Windows AMD drivers. Vulkan is enabled by default
and is the recommended fallback for those systems. If a mixed iGPU/dGPU
system selects an unstable Vulkan iGPU, set 

`GGML_VK_VISIBLE_DEVICES` to the
discrete GPU index.
## Filesystem Requirements

The Ollama install does not require Administrator, and installs in your home directory by default. You’ll need at least 4GB of space for the binary install. Once you’ve installed Ollama, you’ll need additional space for storing the Large Language models, which can be tens to hundreds of GB in size. If your home directory doesn’t have enough space, you can change where the binaries are installed, and where the models are stored.
### Changing Install Location

To install the Ollama application in a location different than your home directory, start the installer with the following flag
### Changing Model Location

To change where Ollama stores the downloaded models instead of using your home directory, set the environment variable`OLLAMA_MODELS` in your user account.
1. 
Start the Settings (Windows 11) or Control Panel (Windows 10) application and search for *environment variables* .
2. 
Click on *Edit environment variables for your account* .
3. 
Edit or create a new variable for your user account for `OLLAMA_MODELS` where you want the models stored
4. Click OK/Apply to save.

## API Access

Here’s a quick example showing API access from`powershell`
## Troubleshooting

Ollama on Windows stores files in a few different locations. You can view them in the explorer window by hitting`<Ctrl>+R` and type in:
- `explorer %LOCALAPPDATA%\Ollama` contains logs, and downloaded updates
  - *app.log* contains most resent logs from the GUI application
  - *server.log* contains the most recent server logs
  - *upgrade.log* contains log output for upgrades
- `explorer %LOCALAPPDATA%\Programs\Ollama` contains the binaries (The installer adds this to your user PATH)
- `explorer %HOMEPATH%\.ollama` contains models and configuration
- `explorer %TEMP%` contains temporary executable files in one or more`ollama*` directories

## Uninstall

The Ollama Windows installer registers an Uninstaller application. Under`Add or remove programs` in Windows Settings, you can uninstall Ollama.
## Standalone CLI

The easiest way to install Ollama on Windows is to use the`OllamaSetup.exe`
installer. It installs in your account without requiring Administrator rights.
We update Ollama regularly to support the latest models, and this installer will
help you keep up to date.
If you’d like to install or integrate Ollama as a service, a standalone
`ollama-windows-amd64.zip` zip file is available containing only the Ollama CLI
and GPU library dependencies for Nvidia. Depending on your hardware, you may also
need to download and extract additional packages into the same directory:
- **AMD GPU** :`ollama-windows-amd64-rocm.zip`
- **MLX (CUDA)** :`ollama-windows-amd64-mlx.zip`

`ollama serve` with tools such as
NSSM.
If you are upgrading from a prior version, you should remove the old directories first.
