---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-llms-with-openai-codex-2
title: "How to Run Local LLMs with OpenAI Codex"
domain: servers-hardware
role: reference
task: reference
actors: ["OpenAI", "Unsloth"]
dates: []
keywords: ["agents", "chatgpt", "gpu", "inference", "llama", "llama.cpp", "sandbox"]
source: docs/RAG/clean4/How to Run Local LLMs with OpenAI Codex.md
source_anchor: ""
source_lines: [116, 207]
sha256: b4d8a6803c9ee5e1d3158d869673e3a4595fa15aec3cc06139362a101f1280c1
---

# How to Run Local LLMs with OpenAI Codex

|
| `requires_openai_auth` | `false` makes Codex skip the "Sign in with ChatGPT" screen for this provider. |
{% hint style="warning" %}
OpenAI removed `wire_api = "chat"` support. Always use `wire_api = "responses"`. If you set `wire_api = "chat"`, Codex refuses to start with `` `wire_api = "chat"` is no longer supported. How to fix: set `wire_api = "responses"` in your provider config. ``
{% endhint %}
{% hint style="info" %}
You can create multiple profile files, one for each Unsloth model you swap between. Launch the one you want with `codex --profile `.
{% endhint %}
{% endstep %}
{% step %}
#### Set the API key env var
Use the same env var name you wrote in `env_key`. In the Unsloth Studio example above, `env_key = "UNSLOTH_STUDIO_AUTH_TOKEN"`, so set `UNSLOTH_STUDIO_AUTH_TOKEN` in the same terminal you will run Codex from:
{% code title="MacOS / Linux / WSL" %}
```bash
export UNSLOTH_STUDIO_AUTH_TOKEN=YOUR_TOKEN
```
{% endcode %}
{% code title="Windows PowerShell" %}
```powershell
$env:UNSLOTH_STUDIO_AUTH_TOKEN = "YOUR_TOKEN"
```
{% endcode %}
If you renamed `env_key`, rename the variable in the commands too. For example, a llama.cpp profile that uses `env_key = "LLAMA_CPP_API_KEY"` needs `LLAMA_CPP_API_KEY`, not `UNSLOTH_STUDIO_AUTH_TOKEN`.
**Session vs Persistent:** the commands above apply to the current terminal only. To persist:
* **MacOS / Linux / WSL:** add the `export` line to `~/.bashrc` (bash) or `~/.zshrc` (zsh).
* **Windows:** run `setx UNSLOTH_STUDIO_AUTH_TOKEN "YOUR_TOKEN"` once, or add the `$env:` line to your PowerShell `$PROFILE`.
{% hint style="warning" %}
**Running Codex inside WSL with Unsloth on Windows?** WSL is a separate network namespace, so `localhost` from inside WSL doesn't reach Unsloth. Edit your `config.toml` to use the Windows host IP instead:
```bash
# Get the Windows host IP from inside WSL
ip route | grep default | awk '{print $3}'
```
Then set `base_url = "http://:8888/v1"`. If you have WSL2 mirrored networking enabled (`.wslconfig` → `networkingMode=mirrored`), `localhost` works as on native Windows.
{% endhint %}
{% endstep %}
{% step %}
#### **Launch Codex**
```bash
mkdir my-project && cd my-project
codex --oss --profile unsloth_api
```
{% hint style="info" %}
**First launch in a new directory** Codex asks *"Do you trust the contents of this directory?"* - pick *Yes, continue.* This is the per-cwd trust prompt, not the ChatGPT login (that one is skipped because of \`requires\_openai\_auth = false\`). Subsequent launches in the same directory skip this prompt.
{% endhint %}
{% hint style="info" %}
**Seeing `Model metadata for` unsloth/gemma-4-26B-A4B `not found. Defaulting to fallback metadata`?** Codex ships with a built-in table of context windows, tool support, and input modalities for OpenAI's own models. For anything else - it falls back to safe defaults. The warning fires once per session for every non-OpenAI slug. Everything still works, you can ignore it.
**To fix it:** add `model_context_window = 131072` to the top of `~/.codex/config.toml` so Codex uses Gemma 4's real 128K context instead of its fallback guess. For full control over tool support and input modalities too, point `model_catalog_json` inside `[profiles.unsloth_api]` at a JSON file containing a custom `ModelInfo` entry for your slug.
{% endhint %}
The `--profile unsloth_api` flag tells Codex to load `~/.codex/unsloth_api.config.toml`, which selects the Unsloth Studio provider and model. Add `--oss` to run through Codex's local OSS provider flow. The model name appears in Codex's status bar.
Add `--search` to enable web search:
```bash
codex --oss --profile unsloth_api --search
```
To bypass all approval prompts **(BEWARE this will make Codex do and execute code however it likes without any approvals!)**:
{% code overflow="wrap" %}
```bash
codex --oss --profile unsloth_api --search --dangerously-bypass-approvals-and-sandbox
```
{% endcode %}
{% endstep %}
{% endstepper %}
### Try a real task
Try this prompt to install and run a simple Unsloth finetune:
{% code overflow="wrap" %}
```
You can only work in the cwd project/. Do not search for AGENTS.md - this is it.
Install Unsloth via a virtual environment via uv. See
https://unsloth.ai/docs/get-started/install/pip-install on how (get it and read).
Then do a simple Unsloth finetuning run described in
https://github.com/unslothai/unsloth. You have access to 1 GPU.
```
{% endcode %}
and if we wait a little longer, you will see a successfully fine-tuned model with Unsloth!
### Disconnect or revert
Launch Codex without `-p unsloth_api` and it'll use its default provider. Or delete the `[profiles.unsloth_api]` and `[model_providers.unsloth_api]` blocks from `~/.codex/config.toml`.
```bash
unset UNSLOTH_STUDIO_AUTH_TOKEN
```
You can leave Unsloth Studio running or shut it down. It doesn't intercept anything when stopped.
### Troubleshooting
| Symptom | Likely cause | Fix |
| ------------------------------------------ | ------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `Model metadata for ... not found` | Non-OpenAI slug, no built-in metadata | Harmless warning. To silence the side-effects, set `model_context_window = 131072` in `~/.codex/config.toml`, or point |
| Codex says it's GPT | Codex injects an OpenAI-referencing system prompt; local models mirror it | Not a routing bug. Verify via Unsloth's activity panel. Override the system prompt to change self-report. |
| `Connection refused` | Unsloth isn't running or wrong port | Confirm Unsloth is up at `http://localhost:8888`; check `base_url` in `config.toml` |
| `wire_api = "chat" is no longer supported` | Legacy `wire_api = "chat"` in config | Switch to `wire_api = "responses"` |
| `model not found` | Model ID typo | `GET http://localhost:8888/v1/models` and copy the exact ID |
| OOM mid-generation | Context too large for VRAM | Reduce context in Unsloth **Settings → Inference**, or use a smaller quant |
| Codex shows "Sign in with ChatGPT" picker |

Launched bare codex (no --oss)

