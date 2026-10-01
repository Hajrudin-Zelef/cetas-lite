---
id: collect-261001-cisco/cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-5
title: "1. Activate your Python virtual environment"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2026-03-26"]
keywords: []
source: docs/RAG/collect-261001-cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-at-662fce.md
source_anchor: ""
source_lines: [661, 816]
sha256: abada304f6f205da0db08f456f65545bb1e816e5f9926048537831d57fed3085
---

# 1. Activate your Python virtual environment

```
# Standard run — reads inventory.yml automatically via ansible.cfg
ansible-playbook ansible-git-catc.yml --vault-password-file .vault_pass
# Interactive vault prompt (if .vault_pass not set up)
ansible-playbook ansible-git-catc.yml --ask-vault-pass
# Override specific inventory variables at runtime
ansible-playbook ansible-git-catc.yml \
  --vault-password-file .vault_pass \
  -e "projectName=MyTestProject" \
  -e "git_branch=feature-xyz"
# Debug mode — prints enriched template list, workflow configs, and CatC sync results
DEBUG=true ansible-playbook ansible-git-catc.yml --vault-password-file .vault_pass
# Syntax check (no connection required)
ansible-playbook ansible-git-catc.yml --syntax-check --vault-password-file .vault_pass
# Dry run — shows what would change without actually connecting to CatC
ansible-playbook ansible-git-catc.yml --check --vault-password-file .vault_pass
```
Set `DEBUG=true` as an environment variable before the playbook command to enable verbose output.

`DEBUG=true ansible-playbook ansible-git-catc.yml --vault-password-file .vault_pass`
| Debug task | Variable / content shown | 
|---|---|
| Execution timestamp | `start_timestamp.date + time` | 
| GitHub repo slug | `git_repo_slug` — confirms URL parsing | 
| Project name | `projectName` — confirms derived or overridden value | 
| Templates in composites | `composite_referenced_templates[]` — names extracted from`.yml` definitions | 
| Sorted template list | `sorted_template_files[].name` — final processing order | 
| Template workflow configs | `template_workflow_configs[]` — full list of config objects before sync | 
| Composite workflow configs | `composite_workflow_configs[]` — full composite config objects | 
| Template sync result | `template_sync_result` — full response from first`template_workflow_manager` call | 
| Composite sync result | `composite_sync_result` — full response from second`template_workflow_manager` call | 

A successful run with 20 regular templates and 1 composite (all already up to date):

```
PLAY [Template Synchronization from Git using Template Workflow Manager] ******
TASK [Set execution timestamp] ************************************************
ok: [catalyst_center]
TASK [Parse GitHub repository slug from git_repo URL] ************************
ok: [catalyst_center]
TASK [Verify GitHub repository is accessible] ********************************
ok: [catalyst_center]
TASK [Verify Git branch exists in repository] ********************************
ok: [catalyst_center]
TASK [Fetch repository file tree from GitHub API] ****************************
ok: [catalyst_center]
TASK [Build template and composite file lists from repository tree] **********
ok: [catalyst_center]
TASK [Fetch template file contents from GitHub] ******************************
ok: [catalyst_center] => (item=DEFN-CLIENT-PORTS.j2)
ok: [catalyst_center] => (item=DEFN-VRF.j2)
...
ok: [catalyst_center] => (item=FABRIC-VRF.j2)
TASK [Fetch last commit info for each template file] *************************
ok: [catalyst_center] => (item=DEFN-CLIENT-PORTS.j2)
...
TASK [Build enriched template file objects] **********************************
ok: [catalyst_center] => (item=DEFN-CLIENT-PORTS.j2)
...
TASK [Synchronize all templates using template_workflow_manager] *************
ok: [catalyst_center]    ← "ok" means templates already exist with matching content
TASK [Synchronize composite templates using template_workflow_manager] ********
ok: [catalyst_center]
TASK [Display synchronization summary] ***************************************
ok: [catalyst_center] =>
  msg:
    - "Template synchronization completed successfully"
    - "Project: Building P0"
    - "Regular templates synced: 20"
    - "Composite templates synced: 1"
    - "Timestamp: 2026-03-26 13:21:01"
PLAY RECAP ****************************
catalyst_center : ok=108  changed=0  unreachable=0  failed=0  skipped=5  rescued=0  ignored=0
```
`ok` vs `changed`:


`ok` — CatC already has the template with matching content; no changes made.
`changed` — Template was created or updated.- The high
`ok` count (~108) is normal — each template goes through ~5 Ansible tasks internally inside `template_workflow_manager`.

This playbook is **Step 6** in the full lab automation chain. Templates must exist in CatC before a network profile (Step 7) can bind them to a site, and network profiles must exist before devices can be provisioned (Step 8).

```
1.0 Site Hierarchy
2.0 Settings
3.0 Credentials
4.0 Device Discovery
5.0 Assign to Site
6.0 Templates (this playbook) ←── syncs Jinja2 templates from GitHub to CatC
7.0 Network Profile            ←── binds templates to site hierarchy
8.0 Provision Composite        ←── deploys composite template to managed devices
```
Running playbooks out of order will result in errors. For example, running 8.0 before 6.0 will fail because the composite `BGP-EVPN-BUILD.j2` does not yet exist in CatC.


```
git_repo (GitHub URL)
    │
    ▼ Stage 1 — GET /repos/{slug}/git/trees/{branch}?recursive=1
raw tree[] (all files in repo)
    filter: path starts with git_repo_subfolder/
    ├─ filter: ends with .j2  → api_template_files[]
    └─ filter: ends with .yml → api_composite_files[]
    │
    ▼ Stage 2 — per file: GET raw content + GET last commit metadata
enriched_template_files[]  = [{ name, path, content, commit_message, diff_content }]
enriched_composite_files[] = [{ name, path, content (YAML parsed) }]
    │
    ▼ Stage 3 — composite .yml parsed → composite_referenced_templates[]
    template NOT in composite_referenced → regular_template_list[]   ← DEFN-*, FUNC-*
    template     in composite_referenced → priority_template_list[]  ← FABRIC-*
sorted_template_files = regular_template_list + priority_template_list
    │                ← un-referenced templates first, composite members last
    ▼ Stage 4 — include_tasks: process-template.yml / process-composite.yml
template_workflow_configs[]  ← one entry per .j2 file
composite_workflow_configs[] ← one entry per .yml composite file
    │
    ▼ Stage 5 — cisco.dnac.template_workflow_manager (state: merged)
    call 1: templates  → POST /dna/intent/api/v1/template-programmer/project/{id}/template
    call 2: composites → POST /dna/intent/api/v1/template-programmer/project/{id}/template
                         (composite: true, containing_templates: [...])
```
**Before — GitHub tree API response (truncated):**

```
{
  "tree": [
    { "path": "BGP_EVPN/DayNTemplates/DEFN-LOOPBACKS.j2",  "type": "blob" },
    { "path": "BGP_EVPN/DayNTemplates/FABRIC-NVE.j2",      "type": "blob" },
    { "path": "BGP_EVPN/DayNTemplates/BGP-EVPN-BUILD.yml",  "type": "blob" },
    { "path": "BGP_EVPN/DayNTemplates/README.md",           "type": "blob" }
  ]
}
```
Only entries under `git_repo_subfolder/` matching `.j2` or `.yml` are kept. All other file types (`.md`, `.png`, `.json`, etc.) are silently ignored.


**After — filtered file lists:**

```
{
  "api_template_files":  ["BGP_EVPN/DayNTemplates/DEFN-LOOPBACKS.j2", "BGP_EVPN/DayNTemplates/FABRIC-NVE.j2"],
  "api_composite_files": ["BGP_EVPN/DayNTemplates/BGP-EVPN-BUILD.yml"]
}
```
**After — template ordering decision (Stage 3):**

```
BGP-EVPN-BUILD.yml defines containing_templates: [FABRIC-NVE.j2, ...]
regular_template_list  → [DEFN-LOOPBACKS.j2]   ← not referenced by any composite
priority_template_list → [FABRIC-NVE.j2]        ← referenced in BGP-EVPN-BUILD.yml
sorted_template_files  = [DEFN-LOOPBACKS.j2, FABRIC-NVE.j2]
```
**After — `template_workflow_configs[0]`** (submitted in Stage 5 call 1):

