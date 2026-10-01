---
id: collect-261001-cisco/cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-3
title: "1. Activate your Python virtual environment"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2022-11-28"]
keywords: []
source: docs/RAG/collect-261001-cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-at-662fce.md
source_anchor: ""
source_lines: [258, 438]
sha256: 3aa31683213bac03192fd928cb514195aba9b4127bf0b0de1ec3bf9bcc55d3fa
---

# 1. Activate your Python virtual environment

```
{# At the top of FABRIC-VRF.j2 #}
{% include "{{ TEMPLATE_PROJECT_NAME }}/DEFN-VRF.j2" %}
{% include "{{ TEMPLATE_PROJECT_NAME }}/FUNC-VRF-LOOKUP.j2" %}
```
`{{ TEMPLATE_PROJECT_NAME }}` is replaced by the value of `projectName` from `inventory.yml` (e.g., `Building P0`) before the template is uploaded to CatC. This is why DEFN-* and FUNC-* templates must exist in CatC *before* the FABRIC-* templates that include them can be rendered correctly.

The playbook runs in five sequential stages. All five happen within a single Ansible play targeting the `catalyst_center` inventory host.

Source: `DIAGRAMS/logical-flow.mmd`

Re-render: `mmdc -i DIAGRAMS/logical-flow.mmd -o DIAGRAMS/logical-flow.png --scale 3 && /usr/bin/sips -Z 4000 DIAGRAMS/logical-flow.png`


Before fetching any templates, the playbook validates that the repository and branch are accessible. This avoids wasting time on authentication failures or typos in `inventory.yml`.

**URL parsing:**

```
git_repo: "https://github.com/imanassypov/CatalystCenter-BGP-EVPN-VXLAN.git"
                │
                ▼  regex_replace (strip https://github.com/ and .git)
git_repo_slug = "imanassypov/CatalystCenter-BGP-EVPN-VXLAN"
```
**Pre-flight API calls:**

```
GET https://api.github.com/repos/{git_repo_slug}
    ← 200 OK: repo accessible
    ← 404: fail — "Repository not found: verify git_repo in inventory"
GET https://api.github.com/repos/{git_repo_slug}/branches/{git_branch}
    ← 200 OK: branch exists
    ← 404: fail — "Branch not found: verify git_branch in inventory"
```
**Fetch recursive file tree:**

```
GET https://api.github.com/repos/{git_repo_slug}/git/trees/{git_branch}?recursive=1
    Headers:
      Accept: application/vnd.github+json
      X-GitHub-Api-Version: 2022-11-28
      Authorization: Bearer {git_token}   ← only if git_token defined
    ← {
        "tree": [
          {"type": "blob", "path": "BGP EVPN/DEFN-VRF.j2",          "sha": "..."},
          {"type": "blob", "path": "BGP EVPN/FABRIC-VRF.j2",         "sha": "..."},
          {"type": "blob", "path": "BGP EVPN/BGP-EVPN-BUILD.yml",    "sha": "..."},
          ...
        ]
      }
```
**File filtering (applied to `tree[]`):**

```
git_repo_subfolder = "BGP EVPN"
git_tree_prefix    = "BGP EVPN/"
All tree entries where:
  type == "blob"
  AND path starts with "BGP EVPN/"
  AND path ends with ".j2"   → api_template_files[]   (20 files)
All tree entries where:
  type == "blob"
  AND path starts with "BGP EVPN/"
  AND path ends with ".yml"  → api_composite_files[]  (1 file)
```
For each file identified in Stage 1, the playbook fetches the content and — for templates — the latest commit metadata.

**For each `.j2` template:**

```
# 1. Fetch raw template content
GET https://raw.githubusercontent.com/{slug}/{branch}/{path}
    ← raw Jinja2 text (e.g., the full FABRIC-VRF.j2 configuration template)
# 2. Fetch last commit metadata
GET https://api.github.com/repos/{slug}/commits
    ?path={path}&per_page=1&sha={branch}
    ← [
        {
          "sha": "abc123...",
          "commit": {
            "author": {
              "name": "Igor Manassypov",
              "date": "2026-03-15T14:22:10Z"
            },
            "message": "Add FABRIC-CLIENT-PORTS template"
          }
        }
      ]
# 3. (Optional — when include_diff_header: true)
GET https://api.github.com/repos/{slug}/commits/{sha}
    ← {
        "files": [
          {
            "filename": "BGP EVPN/FABRIC-CLIENT-PORTS.j2",
            "patch": "@@ -0,0 +1,42 @@\n+{# Client port configuration ... #}"
          }
        ]
      }
```
**For each `.yml` composite definition:**

```
GET https://raw.githubusercontent.com/{slug}/{branch}/{path}
    ← raw YAML text content of BGP-EVPN-BUILD.yml
```
**Result — enriched objects assembled:**

```
# enriched_template_files[] — one entry per .j2 file
- name: "FABRIC-VRF.j2"
  path: "BGP EVPN/FABRIC-VRF.j2"
  content: "{% include \"{{ TEMPLATE_PROJECT_NAME }}/DEFN-VRF.j2\" %}\n..."
  commit_message: "2026-03-15T14:22:10Z | Add FABRIC-CLIENT-PORTS template [Igor Manassypov]"
  diff_content: "@@ -10,5 +10,6 @@..."   # empty string if include_diff_header: false
# enriched_composite_files[] — one entry per .yml file
- name: "BGP-EVPN-BUILD.yml"
  path: "BGP EVPN/BGP-EVPN-BUILD.yml"
  content: "templates:\n  - name: \"FABRIC-VRF.j2\"\n  ..."
```
**`catc_template_summary_maxchar`** — the `commit_message` field is truncated to this length (default 1024 characters) before being written to the template description in CatC. The format is `<date> | <message> [<author>]`.


This is the key intelligence of the playbook. Rather than maintaining a static list of template names in a specific order, the playbook **derives the processing order automatically** from the composite definition files.

**Why order matters:**

| Order | Why required | 
|---|---|
| DEFN-* and FUNC-* processed first | FABRIC-* templates use `{% include "Project/DEFN-VRF.j2" %}` — the referenced template must exist in CatC before CatC can validate the FABRIC template | 
| FABRIC-* processed second | The composite template's `containingTemplates` references FABRIC-* templates by name — they must exist in CatC before the composite can be created | 
| Composite processed last | Has a hard dependency on all its member FABRIC-* templates already existing in CatC | 

**How the ordering is determined:**

```
Step 1: Parse each enriched_composite_files[].content as YAML
        → extract the list of template names under "templates:"
        composite_referenced_templates = [
          "FABRIC-VRF.j2",
          "FABRIC-LOOPBACKS.j2",
          "FABRIC-L3OUT.j2",
          "FABRIC-NVE.j2",
          "FABRIC-MCAST.j2",
          "FABRIC-EVPN.j2",
          "FABRIC-OVERLAY.j2",
          "FABRIC-CLIENT-PORTS.j2",
          "FABRIC-NAC.j2"
        ]
Step 2: For each template in enriched_template_files[]:
  template.name IN composite_referenced_templates?
  ├── YES → priority_template_list[]   (FABRIC-* templates)
  └── NO  → regular_template_list[]   (DEFN-* and FUNC-* templates)
Step 3: sorted_template_files = regular_template_list + priority_template_list
        Processing order:
          1. DEFN-CLIENT-PORTS.j2       ← regular (not in composite)
          2. DEFN-L3OUT.j2              ← regular
          3. DEFN-LOOPBACKS.j2          ← regular
          4. DEFN-MCAST.j2              ← regular
          5. DEFN-NAC.j2                ← regular
          6. DEFN-OVERLAY.j2            ← regular
          7. DEFN-ROLES.j2              ← regular
          8. DEFN-VNIOFFSETS.j2         ← regular
          9. DEFN-VRF.j2                ← regular
         10. FUNC-CLIENT-PORTS.j2       ← regular
         11. FUNC-VRF-LOOKUP.j2         ← regular
         12. FABRIC-CLIENT-PORTS.j2     ← priority (in composite)
         13. FABRIC-EVPN.j2             ← priority
         14. FABRIC-L3OUT.j2            ← priority
         15. FABRIC-LOOPBACKS.j2        ← priority
         16. FABRIC-MCAST.j2            ← priority
         17. FABRIC-NAC.j2              ← priority
         18. FABRIC-NVE.j2              ← priority
         19. FABRIC-OVERLAY.j2          ← priority
         20. FABRIC-VRF.j2              ← priority
Step 4: Composites are processed separately (after all 20 regular templates)
          21. BGP-EVPN-BUILD.j2         ← composite
```
Adding or removing templates from the repository automatically adjusts the ordering — no playbook changes required.

The playbook loops over the sorted template list and composite list, calling an included task file for each item. Each task file appends one configuration entry to a running list:

