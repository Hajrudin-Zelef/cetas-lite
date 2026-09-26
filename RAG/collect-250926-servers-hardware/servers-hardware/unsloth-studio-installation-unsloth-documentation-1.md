---
id: collect-250926-servers-hardware/servers-hardware/unsloth-studio-installation-unsloth-documentation-1
title: "unsloth-studio-installation-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Apple", "Intel", "Nvidia", "Unsloth"]
dates: []
keywords: ["amd", "blackwell", "gguf", "gpu", "gpus", "intel", "llama", "llama.cpp", "nvidia", "training"]
source: docs/RAG/clean4/unsloth-studio-installation-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 143]
sha256: ff3c27a3bc6cf08462c14639e142b0e1e48f184e906bbded38012b7e44241e52
---

# unsloth-studio-installation-unsloth-documentation

Unsloth Studio works on Windows, Linux, WSL and MacOS. You should use the same installation process on every device, although the system requirements may differ by device.

The easiest way to install Unsloth Studio is with the native Desktop app. Download it for your operating system:

The rest of this page installs **Unsloth Studio** manually.

- **CPU: Unsloth still works without a GPU** , but for Chat + Data Recipes.
- **Training:** Works on**NVIDIA** ,**Intel** ,**AMD** GPUs and Mac devices

Launch the `terminal` from Mac, then install Unsloth by entering the command below.

`curl -fsSL https://unsloth.ai/install.sh | sh`
Unsloth will start setting up the environment and installing the required packages as shown below. Type **Y** and Press `Enter` when asked if you want to allow Unsloth to start now. This will start Unsloth on your local **8888** port.

Open your browser of choice and go to the `http://127.0.0.1:8888` URL. If this is your first time installing Unsloth, you will be prompted to create a new password. After, the Unsloth app should now open on the Chat Page as shown below.

**Launch Unsloth securely with HTTPS and Cloudflare.** Unsloth now provides a secure way to launch Unsloth over HTTPS through a free Cloudflare tunnel. Use the below (works in Windows, Mac & Linux):

You can start training and running models immediately. You can view our more detailed step-by-step guide to get started below:

Get Started
Open the Start Menu, search for `PowerShell`, and launch it. Copy & paste the install command below:

it will begin installing automatically. After installation finishes, PowerShell will ask if you want to start Unsloth Studio**.**

You can also launch it with the following command:

Open `http://127.0.0.1:8888` in your browser. On first launch, create a new password to continue to the Chat page. **Unsloth Studio** is now installed and ready to use.

The Unsloth app should now open on the Chat Page as shown below.

**Launch Unsloth securely with HTTPS and Cloudflare.** Unsloth now provides a secure way to launch Unsloth over HTTPS through a free Cloudflare tunnel. Use the below (works in Windows, Mac & Linux):

You can start training and running models immediately. You can view our more detailed step-by-step guide to get started below:

Get Started
Open your terminal application. You can launch it by pressing `Ctrl + Alt + T`, or by searching for `Terminal` in your system's application menu.

Click the Windows Start Menu, type the name of your installed distro (e.g. `Ubuntu`), then open it.

On **WSL**, make sure your **NVIDIA drivers** are installed on **Windows** (not inside WSL) and that the **CUDA toolkit** is installed inside your WSL distro. See the System Requirements below for details.

To install, copy and run the install command:

Then:

1. Click inside the terminal window
2. Paste the command with `Ctrl + Shift + V`
3. Press `Enter`

Unsloth will start setting up the environment and installing the required packages as shown below. Type **Y** and Press `Enter` when asked if you want to allow Unsloth to start now. This will start Unsloth on your local **8888** port.

Open your browser of choice and type `http://127.0.0.1:8888` in the URL box. If this is your first time installing Unsloth, you will be forwarded to `http://127.0.0.1:8888/change-password` page as shown below:

**Launch Unsloth securely with HTTPS and Cloudflare.** Unsloth now provides a secure way to launch Unsloth over HTTPS through a free Cloudflare tunnel. Use the below (works in Windows, Mac & Linux):

You can start training and running models immediately. You can view our more detailed step-by-step guide to get started below:

Get Started
To update Unsloth Studio use the same commands as install:

**macOS, WSL, Linux:**

**Windows (PowerShell):**

Unsloth Studio works directly on Windows without WSL. To train models, make sure your system satisfies these requirements:

**Requirements**

- Windows 10 or Windows 11 (64-bit)
- NVIDIA GPU with drivers installed
- **Git** :`winget install --id Git.Git -e --source winget`
- **Python** : version 3.11 up to, but not including, 3.14
- Work inside a Python environment such as **uv** ,**venv** , or**conda/mamba**

Unsloth Studio works on Mac devices for Chat and training and all features. You can run MLX models with Unsloth.

- macOS 12 Monterey or newer (Intel or Apple Silicon)
- Work inside a Python environment such as **uv** ,**venv** , or**conda/mamba**

- Ubuntu 20.04+ or similar distro (64-bit)
- NVIDIA GPU with drivers installed
- CUDA toolkit (12.4+ recommended, 12.8+ for blackwell)
- Git: `sudo apt install git`
- Python: version 3.11 up to, but not including, 3.14
- Work inside a Python environment such as **uv** ,**venv** , or**conda/mamba**

Our Docker image now works for Unsloth! We're working on Mac compatibility.

- Pull our latest Unsloth container image: `docker pull unsloth/unsloth`
- Run the container via:

For more information, see here.

- Access your studio instance at `http://localhost:8000` or external ip address`http://external_ip_address:8000/`

Unsloth Studio supports CPU devices for Chat for GGUF models and Data Recipes (Export coming very soon)

- Same as the ones mentioned above for Linux (except for NVIDIA GPU drivers) and MacOS.

To install into an isolated location (its own virtual env, `auth/`, `studio.db`, cache and llama.cpp build), set `UNSLOTH_STUDIO_HOME` and pass it again at launch:

Then to update :

To install into an isolated location (its own virtual env, `auth/`, `studio.db`, cache and llama.cpp build), set `UNSLOTH_STUDIO_HOME` and pass it again at launch:

Then to update :

By default `unsloth studio` binds to `127.0.0.1` (this machine only). To reach it from another device, pick one of:

- `--secure` (recommended): serve**only** through a free Cloudflare HTTPS link. Unsloth stays bound to localhost and the tunnel provides the public URL; it fails closed (does not start) if the tunnel can't come up, so the raw port is never exposed.

- `-H 0.0.0.0` : bind the raw port on all network interfaces, reachable from anywhere on the network. Only use this on a trusted network.

Server-side tools (web search, Python and terminal code execution) run as your user and are on by default. Anyone who can reach the server with the API key can run code on this machine, so keep your API key private and pass `--disable-tools` when exposing Unsloth.

Installer options can be passed as environment variables. On macOS, Linux and WSL place the variable after the pipe so the shell passes it to `sh`; on Windows set it with `$env:` before piping to `iex`.

Skip PyTorch (GGUF-only mode):

Pin the Python version:

Install to a custom location with `UNSLOTH_STUDIO_HOME`:

Cap Unsloth's native CPU thread pools on high-core hosts: `UNSLOTH_CPU_THREADS=8 unsloth studio -p 8888`.

The recommended way to fully remove Unsloth Studio is the matching uninstall script for your OS. It stops any running servers, removes the install dir, the launcher data dir, the desktop shortcut, and any platform-specific entries (macOS `.app` bundle + Launch Services on Mac; Start Menu, `HKCU\Software\Unsloth` registry key and user `PATH` entries on Windows):

-  **MacOS, WSL, Linux:**

-  **Windows (PowerShell):**

If you only want to drop the install dir and keep the launcher/shortcut for a later reinstall, you can instead run `rm -rf ~/.unsloth/studio` (Mac/Linux/WSL) or `Remove-Item -Recurse -Force "$HOME\.unsloth\studio"` (Windows). The model cache at `~/.cache/huggingface` is not touched by any of these.

If you prefer to remove only specific parts:

**1. Remove app only** (keeps history, chats, checkpoints, and exports intact):

- **macOS, WSL, Linux:**`rm -rf ~/.unsloth/studio/unsloth_studio`
- **Windows (PowerShell):**`Remove-Item -Recurse -Force "$HOME\.unsloth\studio\unsloth_studio"`

**2. Remove Unsloth entirely** (keeps other Unsloth tools intact):

