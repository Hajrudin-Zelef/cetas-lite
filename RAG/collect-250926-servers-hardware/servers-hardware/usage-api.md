---
id: collect-250926-servers-hardware/servers-hardware/usage-api
title: "Usage API"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["claude", "cost", "deepseek", "inference", "parameters", "reasoning"]
source: docs/RAG/clean4/usage-api.md
source_anchor: ""
source_lines: [1, 124]
sha256: c28cbfef5b3372b2f79ab99cdd4803f008e94ac6f8b0ffe106726927002f91a4
---

# Usage API

Export usage records as CSV for a workspace, member, service account, or model. Use the versioned endpoint in scripts, reporting jobs, and billing integrations.

## Quick start

Create a service account API key in the Console. This endpoint accepts service account keys only; user session tokens are rejected.

```
export CONSOLE_URL="https://opencode.ai/console"
export SERVICE_API_KEY="oc_sk_..."
```
Export the last seven days of workspace usage:

```
curl --fail-with-body --get \
  "${CONSOLE_URL}/api/v1/usage/export" \
  --header "Authorization: Bearer ${SERVICE_API_KEY}" \
  --header "Accept: text/csv" \
  --data-urlencode "scope=organization" \
  --data-urlencode "range=7d" \
  --output "usage-organization-7d.csv"
```
A successful request writes `usage-organization-7d.csv`. Remove `--output` to print the CSV in your terminal instead.

## Endpoint and scopes

`GET /api/v1/usage/export`
Every request requires `scope` and `range`. Supported ranges are `24h`, `7d`, and `30d`. They start at midnight UTC
rather than being rolling windows. The full range is returned as a single streamed CSV file, regardless of size.

| Scope | Additional parameters | Result | 
|---|---|---|
| `organization` | None | All usage in the workspace. | 
| `member` | `user_email` | Usage attributed to one member. | 
| `service_account` | `service_account_id` | Usage attributed to one service account. | 
| `model` | `provider` and`model` | Usage for one provider and model pair. | 

## Export examples

One member:

```
curl --fail-with-body --get \
  "${CONSOLE_URL}/api/v1/usage/export" \
  --header "Authorization: Bearer ${SERVICE_API_KEY}" \
  --header "Accept: text/csv" \
  --data-urlencode "scope=member" \
  --data-urlencode "range=24h" \
  --data-urlencode "user_email=alice@example.com" \
  --output "usage-member-24h.csv"
```
One service account:

```
curl --fail-with-body --get \
  "${CONSOLE_URL}/api/v1/usage/export" \
  --header "Authorization: Bearer ${SERVICE_API_KEY}" \
  --header "Accept: text/csv" \
  --data-urlencode "scope=service_account" \
  --data-urlencode "range=30d" \
  --data-urlencode "service_account_id=svcacct_..." \
  --output "usage-service-account-30d.csv"
```
One provider and model:

```
curl --fail-with-body --get \
  "${CONSOLE_URL}/api/v1/usage/export" \
  --header "Authorization: Bearer ${SERVICE_API_KEY}" \
  --header "Accept: text/csv" \
  --data-urlencode "scope=model" \
  --data-urlencode "range=7d" \
  --data-urlencode "provider=anthropic" \
  --data-urlencode "model=claude-sonnet-4-5" \
  --output "usage-model-7d.csv"
```
Provider is the provider catalog key, such as `anthropic`, `deepseek`, or `opencode`. A valid filter with no matching
records still returns HTTP 200 with a header-only CSV file.

## CSV response

Results are newest first. Empty optional values are blank cells. `cost_micro_cents` is the amount charged by Console
in microcents, where 100,000,000 equals one dollar. Managed-inference usage carries its charge; free, BYOK, and
unclassified legacy usage have a zero charge.

Organization and member exports include Websearch charges. Search rows identify `web-search` in
the `service` column and leave provider, model, and token fields blank. Model and provider filters select inference
records only.

| Field | Description | 
|---|---|
| `id` | Unique record ID. New service records use a `service:` prefix; inference and historical IDs are numeric. | 
| `user_email` | Current email of the member that generated the usage, when applicable. | 
| `service_account_name` | Current name of the service account that generated the usage, when applicable. | 
| `app` | Client app title, or its referrer when no title was reported. | 
| `provider` | Provider catalog key, such as `openai` ,`anthropic` ,`deepseek` , or`opencode` . | 
| `model` | Model identifier reported for the request. | 
| `input_tokens` | Number of input tokens consumed. | 
| `output_tokens` | Number of output tokens generated. | 
| `reasoning_tokens` | Number of reasoning tokens reported by the provider. | 
| `cache_read_tokens` | Number of input tokens read from the provider cache. | 
| `cache_write_5m_tokens` | Number of tokens written to a five-minute cache. | 
| `cache_write_1h_tokens` | Number of tokens written to a one-hour cache. | 
| `reasoning_mode` | Reasoning mode used, such as `effort` ,`adaptive` ,`budget` , or`disabled` . | 
| `reasoning_effort` | Provider-specific reasoning effort value, when supplied. | 
| `reasoning_budget_tokens` | Requested reasoning-token budget, when supplied. | 
| `reasoning_source` | Source from which the reasoning configuration was derived. | 
| `billing_source` | How the request was funded, such as `managed-inference` ,`credit` ,`byok` , or`free` . | 
| `cost_micro_cents` | Amount charged by Console in microcents. Divide by 100,000,000 to convert to USD. | 
| `created_at` | When the usage was recorded, as an ISO 8601 UTC timestamp. | 
| `service` | Service identifier, currently `web-search` . Blank for model requests. | 
| `quantity` | Number of billable operations represented by the record; normally one. | 

## Errors

| Status | Meaning | 
|---|---|
| `400` | Missing parameters, invalid values, or filters that do not match the selected scope. | 
| `401` | Missing, invalid, expired, or revoked service API key. | 
| `403` | The authenticated service account is not allowed to read usage. | 

Curl uses `--fail-with-body` in these examples so authentication and validation failures produce a non-zero exit code
while preserving any error response.
