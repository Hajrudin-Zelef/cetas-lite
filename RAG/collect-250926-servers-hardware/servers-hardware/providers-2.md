---
id: collect-250926-servers-hardware/servers-hardware/providers-2
title: "Providers"
domain: servers-hardware
role: reference
task: reference
actors: ["AWS", "Microsoft", "OpenAI", "OpenRouter", "vLLM", "xAI"]
dates: []
keywords: ["aws", "bedrock", "copilot", "embedding", "reasoning", "vllm"]
source: docs/RAG/clean4/providers.md
source_anchor: ""
source_lines: [192, 268]
sha256: 864057b68d693ec8095a6166747d55077f28b39251d43c242e13a88d6c84070c
---

# Providers

```
/connect
# Select GitHub Copilot, then Login with GitHub Copilot.
/models
```
For GitHub Enterprise, enter the deployment URL or domain when prompted. OpenCode uses a Copilot API endpoint returned by GitHub when available; otherwise it derives the endpoint from the enterprise domain.

`company.ghe.com`
The connected account needs Copilot Chat access. If GitHub reports no entitlement, sign up for Copilot Free or ask the organization to assign a Copilot seat, then connect again. OpenCode fetches the accountâs current model list after login and whenever the active Copilot account changes.

## Ollama

Start Ollama and pull a model. OpenCode probes `http://127.0.0.1:11434`, discovers completion models, and adds them to
`/models` without an account connection.

```
ollama serve
ollama pull qwen3:8b
```
Point OpenCode at a remote or proxied Ollama server with `settings.baseURL`. Include `/v1`; OpenCode derives the native
`/api/tags` and `/api/show` discovery paths from it. `apiKey` is optional and is sent as a bearer token to discovery and
model requests.

Discovery refreshes periodically and keeps the last successful inventory during a temporary outage. Embedding-only
models do not appear because OpenCode only adds models whose Ollama metadata includes the `completion` capability.

## Runtimes

LM Studio and vLLM also have built-in local discovery. Their default endpoints are
`http://127.0.0.1:1234/v1` and `http://127.0.0.1:8000/v1`.

LM Studio discovers language models from `/api/v1/models`. vLLM checks `/health` before reading `/v1/models`; its
discovery cannot infer tool support, so discovered vLLM models start with tools disabled. For another OpenAI-compatible
runtime, use the custom provider recipe and list its models explicitly.

## Gateways

OpenRouter uses its native runtime and catalog. Provide `OPENROUTER_API_KEY` to the OpenCode server or connect an
OpenRouter account, then refer to models with the full OpenRouter model ID after the provider prefix.

Keep the native OpenRouter package when overriding its endpoint or routing through an OpenRouter-compatible gateway. This preserves OpenRouter request and reasoning behavior.

For a gateway that exposes an OpenAI-compatible API but has its own model inventory, define a custom provider instead.
The configuration key becomes the provider prefix and each `models` key becomes a selectable model ID.

## Errors

Provider and model errors usually identify the failed stage. Check the server process environment and the exact model ID before changing packages.

| Error or symptom | Check | 
|---|---|
| `No model is available for session ...` | No enabled model is currently selectable. Finish the provider-specific setup, then choose a model in `/models` . | 
| `Model unavailable: provider/model` | The provider is inactive, the model ID is absent or disabled, or dynamic discovery no longer returns it. For custom aliases, check the `models` map key rather than`modelID` . | 
| `Cannot initialize provider/model: NAME is required to resolve the provider endpoint` | A `${NAME}` placeholder remains in`baseURL` . Set that variable on the OpenCode server or replace the template with a complete endpoint. | 
| `Azure resource name is missing` | Set `providers.azure.settings.resourceName` ,`AZURE_RESOURCE_NAME` , or a complete`settings.baseURL` . | 
| Vertex does not appear | Set a resolvable project in provider settings or a supported project variable. ADC alone does not activate the provider. | 
| Bedrock does not appear | Provide a profile or another supported AWS credential-chain input. `AWS_REGION` by itself only selects a region. | 
| Ollama has no models | Confirm the server exposes `/api/tags` and`/api/show` at the path derived from`baseURL` , and that`/api/show` reports the`completion` capability. | 
| Copilot has no models | Confirm the active OAuth account has Copilot Chat access. A failed model sync is logged as `failed to sync GitHub Copilot models` . | 

V2 rejects the retired provider IDs with a direct replacement. Use `azure/<model>` instead of
`azure-cognitive-services/<model>`, and `google-vertex/<model>` instead of `google-vertex-anthropic/<model>`.

## WebSockets

OpenAI, xAI, and supported Azure Responses models can keep one WebSocket connection open per session. Consecutive steps reuse the unchanged request prefix and only send content added since the previous response, reducing uploads in long sessions.

WebSocket behavior follows these rules:

- Built-in providers opt supported routes in according to their own policy.
- Provider `settings.transport: "websocket"` enables it;`"http"` disables it.
- `"websocket"` on a route without a WebSocket channel logs a warning and falls back to HTTP.
- OpenAI provider compaction uses the same connection.
- xAI continues from stored responses only. With its default `store: false` , each step is sent in full over the reused
connection.
- A closed socket reconnects on the next step. If the connection cannot open, the session continues over HTTP.
- Provider plugins with `http.request` or`http.response` hooks stay on HTTP so each request remains observable.
