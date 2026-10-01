---
id: collect-261001-ia-llm/ia-llm/eduhsouza-pixel-opnsense-agent-skill-blob-head-skills-opnsense-references-api-co-18525c02
title: "eduhsouza-pixel-opnsense-agent-skill-blob-head-skills-opnsense-references-api-co-18525c02"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/eduhsouza-pixel-opnsense-agent-skill-blob-head-skills-opnsense-references-api-co-18525c02.md
source_anchor: ""
source_lines: [1, 86]
sha256: aafb1c9a2e058b94b9ffdec45c5de9697bf65dbfc43bc488fec8650ba415bff5
---

# eduhsouza-pixel-opnsense-agent-skill-blob-head-skills-opnsense-references-api-co-18525c02

```
https://<host>/api/<module>/<controller>/<command>[/<param>[/<param>...]]
```
Auth is HTTP Basic with the API **key as username** and **secret as password**.
Only `GET` and `POST` exist. Bodies and responses are `application/json`.

**The controller segment is literal.** It is the PHP class name converted to
snake_case, and there is no fallback:

| PHP class | URL segment | 
|---|---|
| `DNatController` | `d_nat` (not`dnat` , not`nat` ) | 
| `SourceNatController` | `source_nat` | 
| `OneToOneController` | `one_to_one` | 
| `AliasUtilController` | `alias_util` | 
| `VlanSettingsController` | `vlan_settings` (not`vlan` ) | 

**The command segment accepts both forms.** The published reference lists
`add_rule`; the GUI calls `addRule`. Phalcon camelizes the action name, so both
reach `addRuleAction`. Acronyms round-trip oddly in the docs —
`get_alias_u_u_i_d` is `getAliasUUID`.

Two exceptions where the docs' own form is the only one: `routes/routes` uses
`addroute`, `delroute`, `getroute`, `searchroute`, `setroute`, `toggleroute` —
all lowercase, no separator.

The reference generator states plainly that it collects endpoints and their
**"most likely call method"** — it is inferred from the code, not declared. So
the method column is wrong in places: `wireguard/client/add_client` is published
as `GET` when adding a peer is obviously a `POST`.

Follow the semantics, not the table: reads are `GET`, and anything that creates,
changes, deletes or triggers is `POST`. The client prints a note on a mismatch
and sends the request regardless.

`docs.opnsense.org` generates its endpoint list per PHP class, so **actions
inherited from a base class are missing**. The clearest case:
`FilterBaseController` provides `apply` to `filter`, `d_nat`, `source_nat`,
`one_to_one` and `npt`, but the reference only shows it under `filter_base`.
`POST /api/firewall/filter/apply` works and is the correct call.

The bundled index patches these back in and marks them `[inherited, undocumented]`. When you need something the index rejects but you believe
exists, `--force` sends it anyway.

1. **Browser devtools.** Open the GUI page, filter network requests by`/api/` ,
and read the real request payload. Nearly every endpoint is used by the GUI,
so this reveals both the route and the exact body shape.
2. **On the firewall itself:** Remember to check the parent class too — that is wherefind /usr/local/opnsense/mvc/app/controllers -path "*/Api/*Controller.php" grep -n "public function .*Action" <path-to-controller>.php `apply` lives.

Controllers built on `ApiMutableModelControllerBase` share one contract.

**Search** — `POST`, returns a paged grid:

`{"current": 1, "rowCount": 20, "searchPhrase": "", "sort": {}}``{"total": 42, "rowCount": 20, "current": 1, "rows": [{"uuid": "...", "...": "..."}]}`
**Get one** — `GET .../get_item/<uuid>`. With no UUID it returns a blank model
with defaults, which is the reliable way to learn every field a type accepts
before writing one.

**Add / set** — `POST`, wrapped in a node named after the model:

`{"rule": {"interface": "lan", "action": "pass", "enabled": "1"}}`
The wrapper name varies (`rule`, `alias`, `host`, `pipe`, `server`). Get the
blank model first and mirror its top-level key.

**Selection fields are dictionaries, not strings.** A `GET` returns
`{"lan": {"value": "LAN", "selected": 1}, "wan": {...}}`; you write back only
the selected key, comma-separated for multi-select. Round-tripping a `get`
straight into a `set` without flattening these is the most common cause of a
silent field reset.

| Response | Meaning | 
|---|---|
| `{"result": "saved", "uuid": "..."}` | Written to `config.xml` . Not live yet. | 
| `{"result": "failed", "validations": {...}}` | Nothing written. Keys name the bad fields. | 
| `{"result": "deleted"}` | Removed from `config.xml` . Not live yet. | 
| `{"status": "ok"}` | An apply/reconfigure succeeded. | 
| HTTP 401 | Wrong credential type, or the key's user lacks the privilege. | 
| HTTP 404 | Wrong controller segment far more often than a missing feature. | 

A 200 with `result: failed` is still a failure — check the body, not just the
status code.

API keys inherit the privileges of the user that owns them. A key for a user
without `page-firewall-rules` gets 401 on filter endpoints even though the key
is valid. For automation, create a dedicated user with only the pages it needs
rather than reusing root.
