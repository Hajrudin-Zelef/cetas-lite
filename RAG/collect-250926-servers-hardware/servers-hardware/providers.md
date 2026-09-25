---
id: collect-250926-servers-hardware/servers-hardware/providers
title: "Providers"
domain: servers-hardware
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI", "OpenRouter", "vLLM", "xAI"]
dates: []
keywords: ["aws", "bedrock", "copilot", "cost", "embedding", "foundry", "gemini", "pricing", "reasoning", "vllm"]
source: docs/RAG/clean4/providers.md
source_anchor: ""
source_lines: [1, 268]
sha256: 0c3051888ccaa718924ad0c5227261cd9964b2f274e587eb4c25cd2188b401ed
---

# Providers

OpenCode includes a provider and model catalog from models.dev. For most providers, connect an account first, then select a model.

## Setup

Run `/connect`, choose a provider, and enter its credentials. Then run `/models` to select one of its models.

```
/connect
/models
```
## Go

OpenCode Go is an optional subscription for coding models tested by the OpenCode team. Subscribe in the
console, copy your API key, then connect it as **OpenCode Go**.

```
/connect
# Select OpenCode Go, then paste your API key.
/models
```
See the Go guide for usage limits, endpoints, and privacy details.

## Custom

Add a provider when its API is not already in the catalog. This OpenAI-compatible example defines the credential, runtime package, endpoint, and first model together.

The `providers` object is keyed by the provider ID used in model references, such as `acme/qwen3-coder`.

| Field | Purpose | 
|---|---|
| `name` | Display name. | 
| `env` | Ordered environment variable names that can provide the credential. | 
| `package` | Runtime provider package. | 
| `canonical` | Built-in provider ID whose catalog defaults this provider inherits. | 
| `settings` | Typed OpenCode controls and JSON options passed to the runtime package. | 
| `headers` | String-valued HTTP headers added to requests. | 
| `body` | JSON fields merged into request bodies. | 
| `models` | Models to add or override, keyed by the OpenCode model ID. | 

## Endpoint

Override `settings.baseURL` to send a catalog provider through a proxy or compatible endpoint. Its existing package,
models, and connection still apply.

`settings` is package-specific. A field only has an effect when the selected package supports it.

## Requests

Use `headers` for additional HTTP headers and `body` for JSON fields that should be merged into every request body.

Both fields can also be set on a model or variant when only part of a providerâs traffic needs the override.

## Packages

The `package` field selects the runtime that communicates with a provider. Use the compatible runtime for APIs that
implement the OpenAI request format.

Native package options include:

- `@opencode/ai/providers/openai`
- `@opencode/ai/providers/openai/chat`
- `@opencode/ai/providers/openai/responses`
- `@opencode/ai/providers/openai-compatible`
- `@opencode/ai/providers/openai-compatible/responses`
- `@opencode/ai/providers/anthropic`
- `@opencode/ai/providers/anthropic-compatible`
- `@opencode/ai/providers/google`
- `@opencode/ai/providers/google-vertex`
- `@opencode/ai/providers/google-vertex/gemini`
- `@opencode/ai/providers/google-vertex/chat`
- `@opencode/ai/providers/google-vertex/responses`
- `@opencode/ai/providers/google-vertex/messages`
- `@opencode/ai/providers/azure`
- `@opencode/ai/providers/azure/chat`
- `@opencode/ai/providers/azure/responses`
- `@opencode/ai/providers/amazon-bedrock`
- `@opencode/ai/providers/amazon-bedrock/mantle`
- `@opencode/ai/providers/amazon-bedrock/mantle/chat`
- `@opencode/ai/providers/amazon-bedrock/mantle/responses`
- `@opencode/ai/providers/openrouter`
- `@opencode/ai/providers/xai`

You can also use an npm package such as `@acme/opencode-provider` or an absolute `file://` URL for a local package.

## Models

Add or override models in a providerâs `models` map. The map key is the model ID used by OpenCode; `modelID` changes the
ID sent to the provider.

| Field | Purpose | 
|---|---|
| `modelID` | Model or deployment ID sent to the provider. | 
| `name` | Display name. | 
| `family` | Model family used for grouping related models. | 
| `package` | Runtime override for this model. | 
| `settings` | Model-level OpenCode controls and package-specific JSON settings. | 
| `headers` | Additional string-valued request headers. | 
| `body` | Additional JSON request body fields. | 
| `capabilities` | Tool support plus accepted input and output media types. | 
| `compatibility` | Request and response compatibility overrides. | 
| `variants` | Named variants with their own `settings` ,`headers` , and`body` . | 
| `cost` | Input, output, and optional cache pricing per million tokens. | 
| `limit` | Context, input, and output token limits. | 
| `disabled` | Removes the model from selection when `true` . | 

See Models for selection, defaults, capabilities, limits, costs, and variants.

## Azure

Azureâs standard catalog endpoint needs a resource name in addition to its credential. Set `settings.resourceName`
once, then connect an API key or use the Microsoft Entra ID session from the Azure CLI as described in
Provider accounts.

Find the **Resource name** in the Azure portal or
Microsoft Foundry. It is also the first part of endpoints such as
`https://my-models.openai.azure.com/` and `https://my-models.services.ai.azure.com/`.

```
az cognitiveservices account list \
  --query "[].{name:name,resourceGroup:resourceGroup}" \
  --output table
```
To use Entra ID, install the Azure CLI, sign in, then
choose **Microsoft Entra ID (Azure CLI)** when connecting Azure.

`az login`
For a resource in another tenant or subscription, select both explicitly.

```
az login --tenant TENANT_ID
az account set --subscription NAME_OR_ID
```
Instead of configuration, `AZURE_RESOURCE_NAME` supplies the resource name to the OpenCode server. The legacy
`AZURE_COGNITIVE_SERVICES_RESOURCE_NAME` variable also works.

OpenCode does not query Azure management APIs or discover deployments. If a deployment does not match its catalog
model name, map an OpenCode model ID to the deployment with `modelID`.

Your identity needs one of these roles:

- **Cognitive Services OpenAI User** for Azure OpenAI models.
- **Cognitive Services User** for other Foundry models.

If a request uses a token from the wrong tenant, sign in again with the required tenant.

`az login --tenant TENANT_ID`
## Bedrock

Amazon Bedrock uses the AWS default credential chain. A named AWS profile is the simplest durable setup; OpenCode also recognizes access-key environments, web identity, and container credentials.

```
aws configure sso --profile work
aws sso login --profile work
```
Select the profile and region in provider settings. A configured `profile`, `AWS_PROFILE`, `AWS_ACCESS_KEY_ID`, web
identity token file, or container credential URI activates the provider; a region alone does not.

Without an explicit region, OpenCode uses `AWS_REGION`, then `AWS_DEFAULT_REGION`, then `us-east-1`. For a private or
VPC endpoint, set `baseURL` while keeping the same profile and region.

Bedrock API keys use `AWS_BEARER_TOKEN_BEDROCK`. Other AWS variables feed SigV4 and are not stored as OpenCode API-key
accounts.

## Vertex

Google Vertex uses Application Default Credentials (ADC) and needs a project before its models become available. Create local ADC, then set the project and location in provider settings.

`gcloud auth application-default login`
OpenCode also resolves the project from `GOOGLE_VERTEX_PROJECT`, `GOOGLE_CLOUD_PROJECT`, `GCP_PROJECT`, or
`GCLOUD_PROJECT`, in that order. It resolves the location from `GOOGLE_VERTEX_LOCATION`, `GOOGLE_CLOUD_LOCATION`, or
`VERTEX_LOCATION`, and defaults to `us-central1`.

`GOOGLE_CLOUD_PROJECT=my-project GOOGLE_VERTEX_LOCATION=europe-west4 opencode --standalone`
Service accounts work through the same ADC path. Point `GOOGLE_APPLICATION_CREDENTIALS` at the service-account JSON
file and still provide a project through settings or one of the project variables above.

```
GOOGLE_APPLICATION_CREDENTIALS=/secure/vertex.json \
GOOGLE_CLOUD_PROJECT=my-project \
opencode --standalone
```
Use the single `google-vertex` provider ID for Gemini, Anthropic, and OpenAI-compatible Vertex catalog models. The old
`google-vertex-anthropic` provider ID is unavailable in V2.

## Copilot

GitHub Copilot supports device OAuth rather than manual API-key entry. Connect **GitHub Copilot**, choose GitHub.com or
GitHub Enterprise, finish the device flow, then open `/models`.

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
