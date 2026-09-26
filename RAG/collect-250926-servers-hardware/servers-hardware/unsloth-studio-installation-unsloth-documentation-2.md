---
id: collect-250926-servers-hardware/servers-hardware/unsloth-studio-installation-unsloth-documentation-2
title: "unsloth-studio-installation-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Google", "Hugging Face", "Nvidia", "Unsloth"]
dates: []
keywords: ["fine-tuning", "gguf", "gpu", "gpus", "inference", "llama", "llama.cpp", "nvidia", "parameters"]
source: docs/RAG/clean4/unsloth-studio-installation-unsloth-documentation.md
source_anchor: ""
source_lines: [144, 222]
sha256: 8c7b0502a0d81564502267a966e1857c4db3668093d06d6ed3606c6fa0faf677
---

# unsloth-studio-installation-unsloth-documentation

- **macOS, WSL, Linux:**`rm -rf ~/.unsloth/studio`
- **Windows (PowerShell):**`Remove-Item -Recurse -Force "$HOME\.unsloth\studio"`

**3. Remove everything Unsloth-related:**

- **macOS, WSL, Linux:**`rm -rf ~/.unsloth`
- **Windows (PowerShell):**`Remove-Item -Recurse -Force "$HOME\.unsloth"`

Note: Step 3 deletes everything in history, chats, model checkpoints, and exports. This cannot be undone.

**4. Remove shortcuts and symlinks:**

**macOS:**

**Linux:** 

**WSL / Windows (PowerShell):**

**5. Remove the CLI command:**

- **macOS, Linux, WSL:**`rm -f ~/.local/bin/unsloth`
- **Windows (PowerShell):** The installer added the venv's`Scripts` directory to your User PATH. To remove it, open Settings → System → About → Advanced system settings → Environment Variables, find`Path` under User variables, and remove the entry pointing to`.unsloth\studio\...\Scripts` .

You can delete old model files either from the bin icon in model search or by removing the relevant cached model folder from the default Hugging Face cache directory. By default, Hugging Face uses `~/.cache/huggingface/hub/` on macOS/Linux/WSL and `C:\Users\<username>\.cache\huggingface\hub\` on Windows.

- **MacOS, Linux, WSL:**`~/.cache/huggingface/hub/`
- **Windows:**`%USERPROFILE%\.cache\huggingface\hub\`

If `HF_HUB_CACHE` or `HF_HOME` is set, use that location instead. On Linux and WSL, `XDG_CACHE_HOME` can also change the default cache root.

**Apr 1 update:** You can now select an existing folder for Unsloth to detect from.

**Mar 27 update:** Unsloth Studio now **automatically detects older / pre-existing models** downloaded from Hugging Face, LM Studio etc.

**Manual instructions:** Unsloth Studio detects models downloaded to your Hugging Face Hub cache `(C:\Users{your_username}.cache\huggingface\hub)`. If you have GGUF models downloaded through LM Studio, note that these are stored in `C:\Users{your_username}.cache\lm-studio\models` **OR**`C:\Users{your_username}\lm-studio\models` . Sometimes when they are not visible, you will need to move or copy those .gguf files into your Hugging Face Hub cache directory (or another path accessible to llama.cpp) for Unsloth Studio to load them.

After fine-tuning a model or adapter in Unsloth, you can export it to GGUF and run local inference with **llama.cpp** directly in Unsloth Chat. Unsloth Studio is powered by llama.cpp and Hugging Face.

We’ve created a free Google Colab notebook so you can explore all of Unsloth’s features on Colab’s T4 GPUs. You can train and run most models up to 22B parameters, and switch to a larger GPU for bigger models. Just Click 'Run all' and the UI should pop up after installation.

Once installation is complete, scroll to **Start Unsloth Studio** and click **Open Unsloth Studio** in the white box shown on the left:

**Scroll further down, to see the actual UI.**

Sometimes the Unsloth link may return an error. This happens because you might have disabled cookies or you're using an adblocker or Mozilla. You can still access the UI by scrolling below the button.

Google Colab also expects you to stay on the Colab page; if it detects inactivity, it may shut down the GPU session.

Python version error

`sudo apt install python3.12 python3.12-venv` version 3.11 up to, but not including, 3.14

`nvidia-smi not found`

Install NVIDIA drivers from https://www.nvidia.com/Download/index.aspx

`nvcc not found` (CUDA)

`sudo apt install nvidia-cuda-toolkit` or add `/usr/local/cuda/bin` to PATH

llama-server build failed

Non-fatal, Unsloth still works, GGUF inference won't be available. Install `cmake` and re-run setup to fix.

`cmake not found`

`sudo apt install cmake`

`git not found`

`sudo apt install git`

Build failed

Delete `~/.unsloth/llama.cpp` and re-run setup

Last updated

Was this helpful?
