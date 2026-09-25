---
id: collect-250926-servers-hardware/servers-hardware/budgets-api
title: "Budgets API"
domain: servers-hardware
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["inference"]
source: docs/RAG/clean4/budgets-api.md
source_anchor: ""
source_lines: [1, 80]
sha256: 379b246c1891a6a5f8dad1818d671b1ce79baf739ca9013e9d22ac967fd4ac96
---

# Budgets API

Read effective monthly budgets for every workspace member and manage each memberâs custom budget from an integration.

## Quick start

Create a service account API key with **All** permissions in the Console.
Inference-only keys cannot access budget data or change limits.

```
export CONSOLE_URL="https://opencode.ai/console"
export SERVICE_API_KEY="oc_sk_..."
```
List member budgets:

```
curl --fail-with-body \
  "${CONSOLE_URL}/api/v1/budgets/members" \
  --header "Authorization: Bearer ${SERVICE_API_KEY}" \
  --header "Accept: application/json"
```
## List member budgets

`GET /api/v1/budgets/members`
The response includes every current member, ordered by email. A limit with source `custom` is a member override;
source `default` is inherited from the workspace default. A null limit and source mean the member is unlimited.

```
[
  {
    "user_id": "user_...",
    "email": "alice@example.com",
    "limit_micro_cents": "7525000000",
    "spent_micro_cents": "1250000000",
    "exceeded": false,
    "resets_at": "2026-10-01T00:00:00.000Z",
    "source": "custom",
    "updated_at": "2026-09-03T12:00:00.000Z"
  }
]
```
Monetary response fields are decimal strings in microcents so integrations do not lose precision. There are
100,000,000 microcents per US dollar. Spend and reset timestamps cover the current UTC calendar month. `updated_at` is
null when there is no custom member override.

## Set a custom budget

`PUT /api/v1/budgets/members/:user_id`
Set a $75.25 monthly budget for one member:

```
curl --fail-with-body --request PUT \
  "${CONSOLE_URL}/api/v1/budgets/members/user_..." \
  --header "Authorization: Bearer ${SERVICE_API_KEY}" \
  --header "Content-Type: application/json" \
  --data '{"budget_dollars":75.25}'
```
`budget_dollars` accepts a non-negative US-dollar amount with at most two decimal places. The PUT is idempotent and
returns HTTP 204. Use the member ID returned by the list operation.

## Return to the workspace default

`DELETE /api/v1/budgets/members/:user_id`
Remove the custom override:

```
curl --fail-with-body --request DELETE \
  "${CONSOLE_URL}/api/v1/budgets/members/user_..." \
  --header "Authorization: Bearer ${SERVICE_API_KEY}"
```
Deleting an override does not delete the member. The member inherits the workspace default budget afterward, or becomes unlimited when the workspace has no default. Repeating the request is safe and returns HTTP 204.

## Errors

| Status | Meaning | 
|---|---|
| `400` | Invalid member ID or budget amount. | 
| `401` | Missing, invalid, expired, or revoked service API key. | 
| `403` | The API key does not have permission to manage budgets. | 
| `404` | The member is not in the workspace bound to the API key. |
