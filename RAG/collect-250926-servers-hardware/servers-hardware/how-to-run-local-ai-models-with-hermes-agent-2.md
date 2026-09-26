---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-ai-models-with-hermes-agent-2
title: "How to Run Local AI Models with Hermes Agent"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Unsloth"]
dates: []
keywords: ["agent", "agents", "gguf", "reasoning"]
source: docs/RAG/clean4/How to Run Local AI Models with Hermes Agent.md
source_anchor: ""
source_lines: [154, 204]
sha256: a66031f332a0682c59a0ad681de86fe73a02c1033d013cdfd85bd02704654a70
---

# How to Run Local AI Models with Hermes Agent

| **API base URL** | `http://localhost:8888/v1` *(your Unsloth port + `/v1`)* |
| **API key** | Your `sk-unsloth-…` key |
| **Detected model: … Use this model?** | `Y` *(Hermes auto-detects the model via `GET /v1/models`)* |
| **Context length in tokens** | *(leave blank for auto-detect)* |
| **Display name** | Anything you like, e.g. `unsloth-api` |
Hermes verifies the endpoint against `/v1/models` and confirms the detected model before continuing.
**4. Accept defaults for the remaining prompts** (TTS, tools, messaging gateway, agent settings) you can reconfigure any of them later. Hermes writes everything to `~/.hermes/config.yaml` and `~/.hermes/.env`.
**5. Launch Hermes:**
```bash
hermes
```
The startup banner shows your Unsloth model name in the status bar (e.g. `unsloth/Qwen3.6-27B-GGUF`), and the prompt is ready for input.
{% hint style="info" %}
To reconfigure just the model later, run `hermes setup model`. To edit the config file directly, `hermes config edit` opens `~/.hermes/config.yaml` in your `$EDITOR`.
{% endhint %}
### Optional: tune the Unsloth server
`unsloth run` starts the local API server and loads a model for your app to connect to. You can also customize how the server behaves when starting it.
```bash
# Serve Hermes (--disable-tools passes the agent's own tools through)
unsloth run \
--model unsloth/gemma-4-26B-A4B-it-GGUF \
--disable-tools \
--reasoning off \
-p 8888
```
{% hint style="warning" %}
Use `--disable-tools` when driving Hermes (or any external agent with its own tools). By default Unsloth Studio runs its own server-side tools, which swallows the agent's tool calls, so Hermes answers but never runs its tools. `--disable-tools` switches to passthrough, so Hermes's own tools are used.
{% endhint %}
Use `--reasoning off` to turn thinking off, or `--reasoning on` to turn it on for models that support reasoning.
```bash
# Expose the API on your local network
unsloth run \
--model unsloth/gemma-4-26B-A4B-it-GGUF \
-H 0.0.0.0 \
-p 8888
```
This starts the server on `0.0.0.0:8888`, allowing other devices on your local network to connect. `-p` changes which port the server runs on. If you want phones, laptops, or other devices on your network to connect to the API server, start it with `-H 0.0.0.0`.
Some apps may still override generation settings for individual requests. For more advanced runtime configuration, see the main [API tuning](https://unsloth.ai/docs/basics/api#unsloth-run-command) section.
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/integrations/hermes-agent.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
