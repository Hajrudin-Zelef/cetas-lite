---
id: collect-250926-servers-hardware/servers-hardware/faq-ollama-1
title: "faq-ollama"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["context window", "gpu", "memory", "mistral", "nvidia"]
source: docs/RAG/clean4/faq-ollama.md
source_anchor: ""
source_lines: [1, 123]
sha256: dd2af67d1fac6b14af784c13f3f38767a6fd85739bcf07ec5706fdda1fb0aba9
---

# faq-ollama

## How can I upgrade Ollama?

Ollama on macOS and Windows will automatically download updates. Click on the taskbar or menubar item and then click “Restart to update” to apply the update. Updates can also be installed by downloading the latest version manually. On Linux, re-run the install script:
## How can I view the logs?

Review the Troubleshooting docs for more about using logs.
## Is my GPU compatible with Ollama?

Please refer to the GPU docs.
## How can I specify the context window size?

By default, Ollama uses a context window size of 4096 tokens. This can be overridden with the`OLLAMA_CONTEXT_LENGTH` environment variable. For example, to set the default context window to 8K, use:
`ollama run`, use `/set parameter`:
`num_ctx` parameter:
## How can I tell if my model was loaded onto the GPU?

Use the`ollama ps` command to see what models are currently loaded into memory.
**Output**:

`Processor` column will show which memory the model was loaded into:
- `100% GPU` means the model was loaded entirely into the GPU
- `100% CPU` means the model was loaded entirely in system memory
- `48%/52% CPU/GPU` means the model was loaded partially onto both the GPU and into system memory

## How do I configure Ollama server?

Ollama server can be configured with environment variables.
### Setting environment variables on Mac

If Ollama is run as a macOS application, environment variables should be set using`launchctl`:
1. 
For each environment variable, call `launchctl setenv` .
2. Restart Ollama application.

### Setting environment variables on Linux

If Ollama is run as a systemd service, environment variables should be set using`systemctl`:
1. 
Edit the systemd service by calling `systemctl edit ollama.service` . This will open an editor.
2. 
For each environment variable, add a line `Environment` under section`[Service]` :
3. Save and exit.
4. 
Reload `systemd` and restart Ollama:

### Setting environment variables on Windows

On Windows, Ollama inherits your user and system environment variables.
1. First Quit Ollama by clicking on it in the task bar.
2. 
Start the Settings (Windows 11) or Control Panel (Windows 10) application and search for *environment variables* .
3. 
Click on *Edit environment variables for your account* .
4. 
Edit or create a new variable for your user account for `OLLAMA_HOST` ,`OLLAMA_MODELS` , etc.
5. Click OK/Apply to save.
6. Start the Ollama application from the Windows Start menu.

## How do I use Ollama behind a proxy?

Ollama pulls models from the Internet and may require a proxy server to access the models. Use`HTTPS_PROXY` to redirect outbound requests through the proxy. Ensure the proxy certificate is installed as a system certificate. Refer to the section above for how to use environment variables on your platform.
Avoid setting 

`HTTP_PROXY`. Ollama does not use HTTP for model pulls, only
HTTPS. Setting `HTTP_PROXY` may interrupt client connections to the server.
### How do I use Ollama behind a proxy in Docker?

The Ollama Docker container image can be configured to use a proxy by passing`-e HTTPS_PROXY=https://proxy.example.com` when starting the container.
Alternatively, the Docker daemon can be configured to use a proxy. Instructions are available for Docker Desktop on macOS, Windows, and Linux, and Docker daemon with systemd.
Ensure the certificate is installed as a system certificate when using HTTPS. This may require a new Docker image when using a self-signed certificate.
## Does Ollama send my prompts and answers back to ollama.com?

Ollama runs locally. We don’t see your prompts or data when you run locally. When using cloud-hosted models, we process your prompts and responses to provide the service but do not store or log that content and never train on it. We collect basic account info and limited usage metadata to provide the service that does not include prompt or response content. We don’t sell your data. You can delete your account anytime.
## How do I disable Ollama Cloud features?

Ollama can run in local only mode by disabling Ollama’s cloud features. By turning off Ollama’s cloud features, you will lose the ability to use Ollama’s cloud models and web search. Set`disable_ollama_cloud` in `~/.ollama/server.json`:
`Ollama cloud disabled: true`.
## How can I expose Ollama on my network?

Ollama binds 127.0.0.1 port 11434 by default. Change the bind address with the`OLLAMA_HOST` environment variable.
Refer to the section above for how to set environment variables on your platform.
## How can I use Ollama with a proxy server?

Ollama runs an HTTP server and can be exposed using a proxy server such as Nginx. To do so, configure the proxy to forward requests and optionally set required headers (if not exposing Ollama on the network). For example, with Nginx:
## How can I use Ollama with ngrok?

Ollama can be accessed using a range of tunneling apps. For example with Ngrok:
## How can I use Ollama with Cloudflare Tunnel?

To use Ollama with Cloudflare Tunnel, use the`--url` and `--http-host-header` flags:
## How can I allow additional web origins to access Ollama?

Ollama allows cross-origin requests from`127.0.0.1` and `0.0.0.0` by default. Additional origins can be configured with `OLLAMA_ORIGINS`.
For browser extensions, you’ll need to explicitly allow the extension’s origin pattern. Set `OLLAMA_ORIGINS` to include `chrome-extension://*`, `moz-extension://*`, and `safari-web-extension://*` if you wish to allow all browser extensions access, or specific extensions as needed:
## Where are models stored?

- macOS: `~/.ollama/models`
- Linux: `/usr/share/ollama/.ollama/models`
- Windows: `C:\Users\%username%\.ollama\models`

### How do I set them to a different location?

If a different directory needs to be used, set the environment variable`OLLAMA_MODELS` to the chosen directory.
On Linux using the standard installer, the 

`ollama` user needs read and write access to the specified directory. To assign the directory to the `ollama` user run `sudo chown -R ollama:ollama <directory>`.
## How can I use Ollama in Visual Studio Code?

Install the Ollama extension to use Ollama models in VS Code Chat. See the VS Code integration guide for setup and troubleshooting.
## How do I use Ollama with GPU acceleration in Docker?

The Ollama Docker container can be configured with GPU acceleration in Linux or Windows (with WSL2). This requires the nvidia-container-toolkit. See ollama/ollama for more details. GPU acceleration is not available for Docker Desktop in macOS due to the lack of GPU passthrough and emulation.
## Why is networking slow in WSL2 on Windows 10?

This can impact both installing Ollama, as well as downloading models. Open`Control Panel > Networking and Internet > View network status and tasks` and click on `Change adapter settings` on the left panel. Find the `vEthernet (WSL)` adapter, right click and select `Properties`.
Click on `Configure` and open the `Advanced` tab. Search through each of the properties until you find `Large Send Offload Version 2 (IPv4)` and `Large Send Offload Version 2 (IPv6)`. *Disable*both of these properties.

## How can I preload a model into Ollama to get faster response times?

If you are using the API you can preload a model by sending the Ollama server an empty request. This works with both the`/api/generate` and `/api/chat` API endpoints.
To preload the mistral model using the generate endpoint, use:
## How do I keep a model loaded in memory or make it unload immediately?

