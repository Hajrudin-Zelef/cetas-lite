---
id: collect-250926-servers-hardware/servers-hardware/unsloth-via-docker-3
title: "Install Unsloth via Docker"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Hugging Face", "Nvidia", "Unsloth"]
dates: []
keywords: ["agent", "agents", "amd", "gpu", "gpus", "nvidia"]
source: docs/RAG/clean4/Unsloth via Docker.md
source_anchor: ""
source_lines: [239, 306]
sha256: b613298df89048da9432438ec31a16187313f172c54cadf0d211e0647c17607b
---

# Install Unsloth via Docker

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
