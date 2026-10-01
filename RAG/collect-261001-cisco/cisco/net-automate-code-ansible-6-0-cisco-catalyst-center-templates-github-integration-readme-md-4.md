---
id: collect-261001-cisco/cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-4
title: "1. Activate your Python virtual environment"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2022-11-28"]
keywords: []
source: docs/RAG/collect-261001-cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-at-662fce.md
source_anchor: ""
source_lines: [439, 660]
sha256: 6b8dd9d0559c1caa01081783e7643940b1dfa8045648a8c62b6be29401d2dfb0
---

# 1. Activate your Python virtual environment

```
sorted_template_files[] (20 entries)
  └── include_tasks: process-template.yml  (runs 20 times)
        └── appends one entry to template_workflow_configs[]
enriched_composite_files[] (1 entry)
  └── include_tasks: process-composite.yml (runs 1 time)
        └── appends one entry to composite_workflow_configs[]
```
See the Task File Reference section for exact payload structures.

Two sequential calls to `cisco.dnac.template_workflow_manager` with `state: merged`:

```
Call 1 — Regular templates (all 20)
─────────────────────────────────────
cisco.dnac.template_workflow_manager:
  state: merged
  config: "{{ template_workflow_configs }}"   ← list of 20 entries
  For each template in the list:
    Does it exist in CatC project?
    ├── No  → CREATE template (POST to /template-programmer/project/{id}/template)
    └── Yes → UPDATE template (PUT with new content + new version)
Call 2 — Composite template (1 entry, after Call 1 completes)
──────────────────────────────────────────────────────────────
cisco.dnac.template_workflow_manager:
  state: merged
  config: "{{ composite_workflow_configs }}"  ← list of 1 entry
  Does it exist in CatC project?
  ├── No  → CREATE composite (all member templates from Call 1 now guaranteed to exist)
  └── Yes → UPDATE composite (member list re-synced)
```
Both calls are **idempotent**: if a template already exists in CatC with identical content, the module performs no action on it (Ansible shows `ok` instead of `changed`).

Called once per template in the `sorted_template_files` loop. Receives the `template_file` loop variable.

**Input — `template_file` object:**

```
name: "FABRIC-VRF.j2"
path: "BGP EVPN/FABRIC-VRF.j2"
content: "{% include \"{{ TEMPLATE_PROJECT_NAME }}/DEFN-VRF.j2\" %}\n..."
commit_message: "2026-03-15T14:22:10Z | Add FABRIC-CLIENT-PORTS template [Igor Manassypov]"
diff_content: ""   # empty when include_diff_header: false
```
**Processing steps:**

```
Step 1 — Resolve project name placeholder
  Replace every occurrence of {{ TEMPLATE_PROJECT_NAME }} in template content
  with the value of projectName from inventory (e.g., "Building P0")
  Before: {% include "{{ TEMPLATE_PROJECT_NAME }}/DEFN-VRF.j2" %}
  After:  {% include "Building P0/DEFN-VRF.j2" %}
Step 2 — Wrap diff (if include_diff_header: true)
  Each line of diff_content is wrapped in a Jinja2 comment:
  "@@ -10,5 +10,6 @@" → "{## @@ -10,5 +10,6 @@ ##}"
  The full wrapped diff becomes diff_content_raw.
Step 3 — Assemble final template content
  template_content = diff_content_raw + template_content_raw
  (diff header at top, Jinja2 configuration content below)
Step 4 — Build the config entry
  Append to template_workflow_configs[]:
```
**Output — one entry appended to `template_workflow_configs[]`:**

```
configuration_templates:
  template_name: "FABRIC-VRF.j2"
  project_name: "Building P0"
  language: "JINJA"
  template_content: |
    {% include "Building P0/DEFN-VRF.j2" %}
    {% include "Building P0/FUNC-VRF-LOOKUP.j2" %}
    ...
template_description: "2026-03-15T14:22:10Z | Add FABRIC-CLIENT-PORTS template [Igor Manassypov]"
  device_types:
    - product_family: "Switches and Hubs"
      product_series: "Cisco Catalyst 9500 Series Switches"
    - product_family: "Switches and Hubs"
      product_series: "Cisco Catalyst 9300 Series Switches"
  software_type: "IOS"
  software_variant: "XE"
  software_version: null
  template_params: []
  failure_policy: "ABORT_TARGET_ON_ERROR"
  version: "1.0"
  tags: []
```
Called once per composite definition file. Receives the `composite_file` loop variable.

**Input — `composite_file` object:**

```
name: "BGP-EVPN-BUILD.yml"
path: "BGP EVPN/BGP-EVPN-BUILD.yml"
content: |
  templates:
    - name: "FABRIC-VRF.j2"
    - name: "FABRIC-LOOPBACKS.j2"
    ...
```
**Processing steps:**

```
Step 1 — Parse YAML content
  composite_def = content | from_yaml
  → {
      templates: [
        {name: "FABRIC-VRF.j2"},
        {name: "FABRIC-LOOPBACKS.j2"},
        ...
      ]
    }
Step 2 — Resolve composite name
  composite_name = composite_def.composite_name
                ?? filename with .yml replaced by .j2
                = "BGP-EVPN-BUILD.j2"
Step 3 — Build containing_templates_list[]
  For each template name in composite_def.templates, create an entry:
  {
    name:             "FABRIC-VRF.j2"
    composite:        false
    project_name:     "Building P0"
    language:         "JINJA"
    description:      "description"
    device_types:     [...]
    software_type:    "IOS"
    software_variant: "XE"
    templateParams:   []
    tags:             []
  }
Step 4 — Build the composite config entry
  Append to composite_workflow_configs[]
Step 5 — Reset containing_templates_list = []
  (required — prevents member list from accumulating across composites)
```
**Output — one entry appended to `composite_workflow_configs[]`:**

```
configuration_templates:
  template_name: "BGP-EVPN-BUILD.j2"
  project_name: "Building P0"
  composite: true
  language: "JINJA"
  template_content: ""           # always empty — rendered content comes from members
  template_description: "Composite template synced from Git repository"
  device_types:
    - product_family: "Switches and Hubs"
      product_series: "Cisco Catalyst 9500 Series Switches"
    - product_family: "Switches and Hubs"
      product_series: "Cisco Catalyst 9300 Series Switches"
  software_type: "IOS"
  software_variant: "XE"
  software_version: null
  template_params: []
  failure_policy: "ABORT_TARGET_ON_ERROR"
  version: "1.0"
  containing_templates:
    - name: "FABRIC-VRF.j2"
      composite: false
      project_name: "Building P0"
      language: "JINJA"
      description: "description"
      device_types: [...]
      software_type: "IOS"
      software_variant: "XE"
      templateParams: []
      tags: []
    - name: "FABRIC-LOOPBACKS.j2"
      # ... same structure ...
    # ... 7 more FABRIC-* entries ...
  tags: []
```
```
GET https://api.github.com/repos/{owner}/{repo}/git/trees/{branch}?recursive=1
Accept: application/vnd.github+json
X-GitHub-Api-Version: 2022-11-28
Authorization: Bearer ghp_...   (if git_token defined)
```
Response (abbreviated):

```
{
  "sha": "abc123...",
  "tree": [
    {"path": "BGP EVPN/DEFN-VRF.j2",         "type": "blob", "sha": "..."},
    {"path": "BGP EVPN/FABRIC-VRF.j2",        "type": "blob", "sha": "..."},
    {"path": "BGP EVPN/BGP-EVPN-BUILD.yml",   "type": "blob", "sha": "..."}
  ],
  "truncated": false
}
```
If `truncated: true`, the repository has more than 100,000 files. This playbook does not handle recursive tree pagination — use `git_repo_subfolder` to scope the scan to a subdirectory.


```
GET https://api.github.com/repos/{owner}/{repo}/commits
    ?path=BGP%20EVPN/FABRIC-VRF.j2&per_page=1&sha=main
```
Response (abbreviated):

```
[
  {
    "sha": "def456...",
    "commit": {
      "author": {
        "name": "Igor Manassypov",
        "date": "2026-03-15T14:22:10Z"
      },
      "message": "feat: add client-ports template"
    }
  }
]
```
The `cisco.dnac.template_workflow_manager` module translates each `configuration_templates` entry into the appropriate CatC Template Programmer API calls. The `state: merged` behaviour:

- Template does **not** exist in project →`POST /dna/intent/api/v1/template-programmer/project/{projectId}/template`
- Template **exists** and content has changed →`PUT /dna/intent/api/v1/template-programmer/template/{templateId}` +`POST .../template/version` (commit)
- Template exists and content is identical → no-op (`ok` in Ansible output)

All commands assume you are in the playbook directory with the virtual environment activated.

