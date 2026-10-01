---
id: collect-261001-cisco/cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-6
title: "1. Activate your Python virtual environment"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-at-662fce.md
source_anchor: ""
source_lines: [817, 864]
sha256: 07d96cfc0a31cfd8be5953d65d362a44c30ee08a6435c6f7d3a908876af6ce1e
---

# 1. Activate your Python virtual environment

```
{
  "configuration_templates": {
    "template_name":        "DEFN-LOOPBACKS.j2",
    "project_name":         "Building P0",
    "language":             "JINJA",
    "template_content":     "...",
    "template_description": "Template synced from Git Building P0 | add loopback definitions",
    "device_types":         [{ "product_family": "Switches and Hubs", "product_series": "Cisco Catalyst 9000 Series" }],
    "software_type":        "IOS",
    "software_variant":     "XE",
    "composite":            false,
    "failure_policy":       "ABORT_TARGET_ON_ERROR",
    "version":              "1.0"
  }
}
```
**After — `composite_workflow_configs[0]`** (submitted in Stage 5 call 2):

```
{
  "configuration_templates": {
    "template_name":        "BGP-EVPN-BUILD.j2",
    "project_name":         "Building P0",
    "language":             "JINJA",
    "composite":            true,
    "template_content":     "",
    "containing_templates": [
      { "name": "DEFN-LOOPBACKS.j2", "composite": false, "project_name": "Building P0" },
      { "name": "FABRIC-NVE.j2",     "composite": false, "project_name": "Building P0" }
    ]
  }
}
```
Templates are synced in two separate `template_workflow_manager` calls: all individual templates first, then all composite templates. This guarantees every member template exists in Catalyst Center before the composite that references it is created or updated.

| Symptom | Likely Cause | Resolution | 
|---|---|---|
| `No inventory was parsed` | `ansible.cfg` is being ignored | Run `chmod o-w .` — Ansible ignores`ansible.cfg` in world-writable directories | 
| `Repository not found` (404 from GitHub) | `git_repo` URL incorrect, or repo is private without a token | Verify the URL in `inventory.yml` ; set`git_token` in`vault.yml` for private repos | 
| `Branch not found` (404 from GitHub) | `git_branch` does not exist in the repository | Verify the branch name in `inventory.yml` | 
| `HTTP Error 403: rate limit exceeded` | No `git_token` set; GitHub limits unauthenticated calls to 60/hr per IP | Set `git_token` in`vault.yml` — authenticated calls allow 5,000/hr | 
| `NCTP10073: syntax error in template` | CatC Jinja2 parser limitation — some Python Jinja2 syntax is not supported (e.g., `not X in Y` vs`X not in Y` ) | Rewrite the affected line in the source template; `not in` may need to be written as`not X in Y` | 
| Composite created but device config is wrong | Member templates applied in wrong order | Check the `templates:` list order in your`.yml` composite definition — the order maps directly to the device config rendering sequence | 
| Template update not reflected in CatC | Template content appears identical to CatC | `template_workflow_manager` compares content — ensure the file actually changed in Git before running the sync | 
| `VAULT_CiPHER_SUITE unavailable` | Python `cryptography` package not installed | Run `pip install -r requirements.txt` | 
| Missing templates after partial run | Playbook was interrupted mid-run (process killed, connection lost) | Re-run the playbook — `state: merged` is idempotent; completed templates are skipped, missing ones are created | 
| Wrong templates synced (wrong project or subfolder) | `projectName` or`git_repo_subfolder` misconfigured | Verify both in `inventory.yml` ; use`DEBUG=true` to check`projectName` and the sorted template list |
