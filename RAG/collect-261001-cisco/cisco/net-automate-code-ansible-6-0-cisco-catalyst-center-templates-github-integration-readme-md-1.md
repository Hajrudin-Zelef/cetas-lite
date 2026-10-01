---
id: collect-261001-cisco/cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-1
title: "1. Activate your Python virtual environment"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-at-662fce.md
source_anchor: ""
source_lines: [1, 119]
sha256: 4cdff76e17f06869a0c2210f1ebb441f70d3d1c8bad67e2470f002454883b53a
---

# 1. Activate your Python virtual environment

**Playbook:** `ansible-git-catc.yml`

**Task files:** `process-template.yml`, `process-composite.yml`

**Module (CatC):** `cisco.dnac.template_workflow_manager`

**Module (GitHub):** `ansible.builtin.uri`

**Minimum Catalyst Center version:** 2.3.7.6

**Minimum Ansible version:** 2.15

**Authors:** Igor Manassypov — Systems Engineer (imanassy@cisco.com)

**Copyright © 2024–2026 Cisco Systems, Inc. All rights reserved.**


This playbook implements a **general-purpose GitOps workflow** for Cisco Catalyst Center template management. It can synchronize **any collection of Jinja2 templates** stored in a GitHub repository — simple flat collections, modular libraries, or fully nested composite templates — directly to a Catalyst Center Template Project without requiring a local clone.

Point the playbook at any GitHub repository subfolder containing `.j2` files and it will fetch the content, enrich each template with Git commit metadata, determine the correct processing order, and sync everything to CatC. No code changes are needed to switch between different template collections.

**About the BGP EVPN example used throughout this document:**

The working example in this playbook is the CatalystCenter-BGP-EVPN-VXLAN template collection. It is used here because it exercises the full range of playbook capabilities — particularly **composite (nested) templates** where one top-level template wraps an ordered set of member templates, and **cross-template Jinja2 includes** where configuration templates pull in shared data-definition and macro-library files at render time. A simpler flat template collection (just `.j2` files, no `.yml` composite definitions) works identically with zero additional configuration.


| Capability | Description | 
|---|---|
| **No local clone** | All content fetched at runtime via the GitHub REST API and raw content URLs | 
| **Idempotent sync** | `state: merged` — creates templates that don't exist yet; updates those that have changed | 
| **Automatic ordering** | Template processing order derived from composite definitions — no hard-coded lists to maintain | 
| **Composite templates** | Builds ordered multi-template composites with full `containingTemplates` resolution | 
| **Git metadata in CatC** | Commit timestamp, message, and author written to the template description field in CatC | 
| **Optional diff headers** | Git patch embedded as Jinja2 comments ( `{## ... ##}` ) at the top of each template | 
| **Private repo support** | Optional `git_token` for authenticated API calls (also raises rate limits from 60 to 5,000/hr) | 

| Platform | Module | Purpose in this playbook | Module Docs | 
|---|---|---|---|
| Cisco Catalyst Center | cisco.dnac.template_workflow_manager | Create or update template and composite template objects in CatC | cisco.dnac 6.46.0: template_workflow_manager | 
| GitHub API | ansible.builtin.uri | Repository verification, branch checks, tree listing, file/commit/diff fetch | ansible-core: uri | 

| Phase | HTTP | Endpoint | Why it is used | API Docs | 
|---|---|---|---|---|
| Repository access check | GET | https://api.github.com/repos/{owner}/{repo} | Validate repository exists and token access is valid | GitHub REST: Repositories API | 
| Branch validation | GET | https://api.github.com/repos/{owner}/{repo}/branches/{branch} | Confirm requested branch is available | GitHub REST: Branches API | 
| Tree discovery | GET | https://api.github.com/repos/{owner}/{repo}/git/trees/{branch}?recursive=1 | Enumerate candidate template and composite files | GitHub REST: Git trees API | 
| Raw template/composite content | GET | https://raw.githubusercontent.com/{owner}/{repo}/{branch}/{path} | Pull source template text used for CatC sync | GitHub docs: Raw file URLs | 
| Commit metadata and diff | GET | https://api.github.com/repos/{owner}/{repo}/commits?... | Build template summary metadata and optional diff header | GitHub REST: Commits API | 
| CatC template sync | module-managed | Template-programmer endpoints used by template_workflow_manager | Apply merged template/composite state in target CatC project | CatC 2.3.7.9: API Reference | 

- GitHub calls are direct uri tasks; Catalyst Center writes are module-managed.
- The section above lists the primary endpoint families used by the workflow stages.

| Requirement | Version | Where to Get It | 
|---|---|---|
| Python | ≥ 3.9 | python.org | 
| Ansible | ≥ 2.15 | `pip install ansible` | 
| `cisco.dnac` Ansible collection | 6.46.0 (tested) | `ansible-galaxy collection install -r requirements.yml` | 
| `community.general` collection | ≥ 8.0 | Same `requirements.yml` | 
| `dnacentersdk` Python SDK | ≥ 2.8.6 | `pip install -r requirements.txt` | 
| Cisco Catalyst Center | ≥ 2.3.7.6 | — | 
| GitHub repository | public or private | — | 

```
6.0-Cisco-Catalyst-Center-Templates-Github-integration/
├── ansible.cfg              # Sets inventory = inventory.yml (no -i flag needed)
├── ansible-git-catc.yml     # Main playbook — 5-stage GitOps sync
├── process-template.yml     # Included task: builds one regular template config
├── process-composite.yml    # Included task: builds one composite template config
├── inventory.yml            # CatC connection + Git repo parameters
├── vault.yml                # Encrypted credentials (gitignored, never commit plain)
├── vault.yml.example        # Vault template — safe to commit
├── requirements.txt         # Python package dependencies
├── requirements.yml         # Ansible Galaxy collection dependencies
└── DIAGRAMS/
    ├── logical-flow.mmd     # Mermaid flowchart source
    └── logical-flow.png     # Rendered diagram (embedded below)
```
```
# 1. Activate your Python virtual environment
cd 6.0-Cisco-Catalyst-Center-Templates-Github-integration/
source ../.venv/bin/activate      # adjust path as needed
# 2. Install Python dependencies
pip install -r requirements.txt
# 3. Install Ansible Galaxy collections
ansible-galaxy collection install -r requirements.yml
# 4. Fix directory permissions (one-time — Ansible ignores ansible.cfg in world-writable dirs)
chmod o-w .
# 5. Set up vault credentials (see Configuration section below)
cp vault.yml.example vault.yml
# Edit vault.yml, then encrypt it:
ansible-vault encrypt vault.yml
echo 'your_vault_password_here' > .vault_pass
chmod 600 .vault_pass
```
All parameters live in `inventory.yml`. Nothing is hard-coded in the playbook.

These variables authenticate to the CatC REST API via the `cisco.dnac` collection:

| Variable | Example | Description | 
|---|---|---|
| `dnac_host` | `198.18.129.100` | CatC management IP or hostname | 
| `dnac_port` | `443` | HTTPS port (default: 443) | 
| `dnac_version` | `2.3.7.9` | CatC version — controls which SDK API paths are used. Set to the highest version that the SDK knows, even if your appliance is newer. | 
| `dnac_verify` | `false` | Set `true` for production (requires a trusted cert).`false` disables TLS verification for lab. | 
| `dnac_debug` | `true` | Enables SDK debug output to `dnac.log` | 
| `dnac_log_level` | `INFO` | Log verbosity: `DEBUG` ,`INFO` ,`WARNING` ,`ERROR` | 
| `dnac_log` | `true` | Writes SDK log to `dnac.log` in the playbook directory | 

| Variable | Example | Description | 
|---|---|---|
| `git_repo` | `https://github.com/org/repo.git` | Full GitHub repository URL | 
| `git_branch` | `main` | Branch to read templates from | 
| `git_repo_subfolder` | `BGP EVPN` | Subfolder within the repo to scan for `.j2` and`.yml` files. Leave blank to scan the entire repository root. | 
| `git_api_base_url` | `https://api.github.com` | Override for GitHub Enterprise instances | 

**Note on `git_token`:** This variable is intentionally absent from `inventory.yml` and defined only in `vault.yml`. See the Vault section below.


